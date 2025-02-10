package functional_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"rupamic-arch/common/middlewares"
	"testing"
)

func HttpHandlerTester(w http.ResponseWriter, r *http.Request) {
	log.Println("HttpHandlerTester Called.")
	userId, ok := r.Context().Value(middlewares.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}
	roles, ok := r.Context().Value(middlewares.RolesKey).([]string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}
	if len(userId) < 1 {
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}
	if len(roles) < 1 {
		http.Error(w, "Roles not found", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "HttpHandlerTester called, UserID %s", userId)
}
func TestAuthAndLogMiddlewares(t *testing.T) {
	userId := "12345678901234"
	roles := []string{"Admin", "SuperDamin"}
	token, err := middlewares.CreateToken(userId, roles)
	if err != nil {
		t.Errorf("Test failed: got %v, want %v", err, nil)
	}
	r := httptest.NewRequest("GET", "/test/path", nil)
	if err != nil {
		t.Errorf("Test failed: got %v, want %v", err, nil)
	}
	r.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	h := middlewares.AuthMiddleware(middlewares.LogMiddleware(http.HandlerFunc(HttpHandlerTester)))
	h.ServeHTTP(rr, r)

	if rr.Code != http.StatusOK {
		t.Errorf("Test failed: got %v, want %v", rr.Code, http.StatusOK)
	}
	expectedResponse := fmt.Sprint("HttpHandlerTester called, UserID ", userId)

	if rr.Body.String() != expectedResponse {
		t.Errorf("Test failed: got %v, want %v", rr.Body.String(), expectedResponse)
	}
}
