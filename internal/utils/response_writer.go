// The response writter is specificially defined for chi and gorilla-mux router for conciseness.
package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponseWriter(w http.ResponseWriter,statusCode int,response interface{}){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&response)
}