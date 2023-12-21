package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func parseFrom(w http.ResponseWriter, r *http.Request, o interface{}) error {

	body := http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(body)
	err := dec.Decode(&o)

	return err
}

func numberParam(r *http.Request, key string) int {
	value := chi.URLParam(r, key)
	num, _ := strconv.Atoi(value)

	return num
}
