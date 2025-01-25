package response

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func Success(w http.ResponseWriter, message any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	responseMap := map[string]any{"success": true, "data": message}

	write(responseMap, w)
}

func Error(w http.ResponseWriter, status int, message any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	responseMap := map[string]any{"success": false, "error": message}

	write(responseMap, w)
}

func write(responseMap map[string]any, w http.ResponseWriter) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(responseMap); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	trimmedResponse := strings.TrimSpace(buf.String())
	if _, err := w.Write([]byte(trimmedResponse)); err != nil {
		log.Printf("Failed to write response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
