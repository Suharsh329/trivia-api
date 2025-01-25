package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	message := map[string]any{"data": "test"}

	Success(w, message)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var responseMap map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&responseMap); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	success, ok := responseMap["success"].(bool)
	if !ok || !success {
		t.Errorf("Expected success to be true, got %v", responseMap["success"])
	}

	data, ok := responseMap["data"].(map[string]any)
	if !ok || data["data"] != "test" {
		t.Errorf("Expected data to be 'test', got %v", responseMap["data"])
	}
}

func TestError(t *testing.T) {
	w := httptest.NewRecorder()
	message := "test error"

	Error(w, http.StatusBadRequest, message)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	var responseMap map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&responseMap); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if responseMap["success"].(bool) {
		t.Errorf("Expected success to be false, got %v", responseMap["success"])
	}

	if responseMap["error"] != message {
		t.Errorf("Expected error to be '%s', got %v", message, responseMap["error"])
	}
}
