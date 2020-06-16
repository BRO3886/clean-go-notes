package utils

import (
	"encoding/json"
	"net/http"
)

//ResponseWrapper response wrapper
func ResponseWrapper(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	// JsonifyHeader(w)
	w.Header().Add("Content-Type", "application/json;  charset=utf-8")
	_ = json.NewEncoder(w).Encode(map[string]string{"message": msg})
}

//JsonifyHeader to set content type as json
func JsonifyHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json;  charset=utf-8")
}

//WrapData to wrap data function
func WrapData(w http.ResponseWriter, v map[string]interface{}) {
	// JsonifyHeader(w)
	w.Header().Add("Content-Type", "application/json;  charset=utf-8")
	_ = json.NewEncoder(w).Encode(v)
}
