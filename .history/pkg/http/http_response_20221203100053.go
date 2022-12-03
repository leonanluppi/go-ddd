package response

import (
	"encoding/json"
	"net/http"
)

func response(res http.ResponseWriter, status int, v any) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(v)
}

func Created(rw http.ResponseWriter, v any) {
	response(rw, http.StatusCreated, v)
}

func Deleted(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNoContent)
}

func OK(rw http.ResponseWriter, v any) {
	response(rw, http.StatusOK, v)
}

type ResponseError struct {
	Message []string `json:"message"`
}

func Error(res http.ResponseWriter, status int, errs ...error) {
	value := ResponseError{
		Message: make([]string, len(errs)),
	}
	for k, err := range errs {
		value.Message[k] = err.Error()
	}
	response(res, status, value)
}
