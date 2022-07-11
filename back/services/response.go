package services

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

func UnauthorizedErrResponse(error string, w http.ResponseWriter) {
	type servererrdata struct {
		Statuscode int    `json:"status"`
		Message    string `json:"msg"`
	}
	response := &servererrdata{Statuscode: 401, Message: error}

	//Send header, status code and output to writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(response)
}

func ServerResponse(msg string, w http.ResponseWriter) {
	type serverdata struct {
		Statuscode int    `json:"status"`
		Message    string `json:"msg"`
	}
	response := &serverdata{Statuscode: 201, Message: msg}

	//Send header, status code and output to writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(response)
}
