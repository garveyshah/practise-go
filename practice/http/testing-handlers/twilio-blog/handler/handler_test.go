package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProcessDataHandler(t *testing.T) {
	validData := []byte(`{"name": "John Doe", "email": "john.doe@example.com"}`)

	req := httptest.NewRequest(http.MethodPost, "/process-data", bytes.NewBuffer(validData))
	req.Header.Set("Content-Type", "application/json")
	fmt.Println()

	fmt.Println("req -> ", req)
	fmt.Println()



	w := httptest.NewRecorder()
	fmt.Println()

	fmt.Println("w ->", w)
	fmt.Println()

	ProcessDataHandler(w, req)
	fmt.Println()

	fmt.Println("w ->", w)
	fmt.Println()


	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, w.Code)
	}

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	if response["status"] != "success" || response["message"] != "Data processed successfully" {
		t.Errorf("Unexpected response body: %v", response)
	}

	req = httptest.NewRequest(http.MethodGet, "/process-data", nil)
	w = httptest.NewRecorder()
	ProcessDataHandler(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %v, got %v", http.StatusMethodNotAllowed, w.Code)
	}

	invalidJSON := []byte(`{"name" : "John Doe", "email.doe@example.com"`) // Misssing closing brace
	req = httptest.NewRequest(http.MethodPost, "/process-data", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	ProcessDataHandler(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, w.Code)
	}
	missingData := []byte(`{"name": "John Doe}`)

	req = httptest.NewRequest(http.MethodPost, "/process-data", bytes.NewBuffer(missingData))
	req.Header.Set("Content-TYpe", "application/json")
	w = httptest.NewRecorder()
	ProcessDataHandler(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, w.Code)
	}
}
