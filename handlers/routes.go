package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"phone-directory-server/data"
)

type HTTPHandler struct {
	r  *chi.Mux
	db *data.DB
}

type ResponseID struct {
	ID int `json:"id"`
}

func NewHTTPHandler(db *data.DB) *HTTPHandler {

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-type", "X-CSRF-Token", "Remote-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(c.Handler)

	h := HTTPHandler{
		r:  r,
		db: db,
	}

	h.initRoutes()

	return &h
}

func (h *HTTPHandler) initRoutes() {

	//get all employees
	h.r.Get("/employees", func(w http.ResponseWriter, r *http.Request) {

		employees, err := h.db.Employee.Employee()
		if err != nil {
			http.Error(w, "500 Internal server error", http.StatusInternalServerError)
			log.Printf("[ERROR] Employee: %s", err)
			return
		}

		jsonResponse(w, employees)

	})

	//post new employee
	h.r.Post("/employees", func(w http.ResponseWriter, r *http.Request) {
		employee := data.Employee{}
		parseFrom(w, r, &employee)

		id, err := h.db.Employee.AddEmployee(&employee)
		if err != nil {
			http.Error(w, "500 Internal server error", http.StatusInternalServerError)
			log.Printf("[ERROR] AddEmployee: %s", err)
			return
		}

		jsonResponse(w, id)
	})

	h.r.Delete("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := numberParam(r, "id")

		fmt.Println(id)
		err := h.db.Employee.DeleteEmployee(id)

		if err != nil {
			http.Error(w, "500 Internal server error", http.StatusInternalServerError)
			log.Printf("[ERROR] DeleteEmployee: %s", err)
			return
		}

		response := ResponseID{
			ID: id,
		}

		jsonResponse(w, response)
	})

	//h.r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
	//	employee := data.Employee{}
	//	parseFrom(w, r, &employee)
	//
	//	foundEmployees, err := h.db.Employee.SearchEmployee(&employee)
	//	if err != nil {
	//		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
	//		log.Printf("[ERROR] SearchEmployees: %s", err)
	//		return
	//	}
	//
	//	jsonResponse(w, foundEmployees)
	//})
}

func (h *HTTPHandler) ServeHTTP(port int) {
	fmt.Printf("Server is listening on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), h.r)
}

func jsonResponse(w http.ResponseWriter, data any) {
	bytes, err := json.Marshal(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}
