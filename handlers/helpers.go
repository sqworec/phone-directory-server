package handlers

import (
	"encoding/json"
	"net/http"
)

func parseFrom(w http.ResponseWriter, r *http.Request, o interface{}) error {
	body := http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(body)
	err := dec.Decode(&o)

	return err
}
