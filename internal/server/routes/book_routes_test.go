package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// test book routes

// test getBooks route
func TestGetBooks(t *testing.T) {
	r := gin.New()
	r.GET("/", getBooks)
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
	// Expected: {"status":"ok","books":[{"ID":"1","Title":"Golang","Author":"Google","Quantity":10},{"ID":"2","Title":"Java","Author":"Oracle","Quantity":20},{"ID":"3","Title":"Python","Author":"Python Software Foundation","Quantity":30}]}
	expected := `{"status":"ok","books":[{"ID":"1","Title":"Golang","Author":"Google","Quantity":10},{"ID":"2","Title":"Java","Author":"Oracle","Quantity":20},{"ID":"3","Title":"Python","Author":"Python Software Foundation","Quantity":30}]}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
