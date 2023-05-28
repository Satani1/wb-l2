package json

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type mapJSON map[string]interface{}

// JSON sends json response
func JSON(w http.ResponseWriter, r *http.Request, status int, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

// ErrorJSON sends error as json
func ErrorJSON(w http.ResponseWriter, r *http.Request, status int, err error, details string) {
	JSON(w, r, status, mapJSON{"ERROR": err.Error(), "Details": details})
}

// NoContentJSON sends no content response
func NoContentJSON(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}