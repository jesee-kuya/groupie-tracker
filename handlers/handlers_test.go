package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		wantStatus int
		wantBody   string
	}{
		{"Home GET", "GET", "/", http.StatusOK, ""},
		{"Home POST", "POST", "/", http.StatusMethodNotAllowed, ""},
		{"Artist GET", "GET", "http://www.example.com/artist?id=2", http.StatusOK, "SOJA"},
		{"Location GET", "GET", "http://www.example.com/location?id=1", http.StatusOK, "Queen"},
		{"Dates GET", "GET", "http://www.example.com/dates?id=3", http.StatusOK, "Pink Floyd"},
		{"Relations GET", "GET", "http://www.example.com/relations?id=9", http.StatusOK, "ACDC"},
		{"Search GET", "GET", "/search", http.StatusOK, ""},
		{"Unknown Path", "GET", "/unknown", http.StatusNotFound, ""},
	}

	err := os.Chdir("../")
	if err != nil {
		t.Error("Failed to cd")
		t.FailNow()
		return
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rr := httptest.NewRecorder()

			Handler(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.wantStatus)
			}

			want := fmt.Sprintf("<h1>%s</h1>", tt.wantBody)
			got := ""
			if tt.wantBody != "" && !strings.Contains(rr.Body.String(), want) {
				t.Errorf("Handler returned unexpected header: got %v want %v", got, want)
			}
		})
	}
}
