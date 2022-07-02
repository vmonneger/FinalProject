package middlewares

import (
	"encoding/json"
	"net/http"
)

func ServerErrResponse(error string, w http.ResponseWriter) {
	type servererrdata struct {
		Statuscode int    `json:"status"`
		Message    string `json:"msg"`
	}
	response := &servererrdata{Statuscode: 500, Message: error}

	//Send header, status code and output to writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(response)
}
