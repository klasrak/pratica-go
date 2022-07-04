package httphelpers

import (
	"encoding/json"
	"net/http"
)

func BadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("content-type", "application/json")
	w.Write(ToJson(ApiError{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}))
}

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ToJson(v interface{}) []byte {
	b, err := json.Marshal(v)

	if err != nil {
		panic(err)
	}

	return b
}
