package controller

import (
	"encoding/json"
	"net/http"
)

type errorDefault struct {
	Message string `json:"message"`
}

func responseJsonErr(w *http.ResponseWriter, err error, status_code int) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(status_code)
	e := errorDefault{err.Error()}
	json.NewEncoder(*w).Encode(e)
}

func responseCreated(w *http.ResponseWriter, body interface{}) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusCreated)
	json.NewEncoder(*w).Encode(body)
}

func responseOk(w *http.ResponseWriter, body interface{}) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusOK)
	json.NewEncoder(*w).Encode(body)
}
