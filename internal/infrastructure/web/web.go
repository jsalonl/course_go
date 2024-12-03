package web

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, httpStatus int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(data)
}
