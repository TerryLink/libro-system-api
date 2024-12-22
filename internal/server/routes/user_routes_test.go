package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllusers(t *testing.T) {

	r := gin.New()
	r.GET("/", getAllusers)
	// Create a test HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Serve the HTTP request
	r.ServeHTTP(rr, req)
	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	// Check the response body
	expected := `{"status":"ok","books":[{ID: "1", Username: "Golang", Email: "Google", Password: "abc"},	{ID: "2", Username: "Java", Email: "Oracle", Password: "abcd"},{ID: "3", Username: "Python", Email: "Python Software Foundation", Password: "abcdefg"}]}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
