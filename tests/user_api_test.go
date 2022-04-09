package tests

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Users struct {
	App http.Handler
}

func (u *Users) Create(t *testing.T) {
	var created map[string]interface{}
	jsonBody := `
        {
            "username": "peterpan",
            "email": "peterpan@gmail.com", 
            "password": "1234", 
            "re_password": "1234", 
            "is_active": true
        }
    `
	err, httpCode := callApi(u.App, http.MethodPost, "/users", jsonBody, &created)
	if err != nil {
		t.Fatal("posting: error call api")
	}
	if http.StatusCreated != httpCode {
		t.Fatalf("posting: expected status code %v, got %v", http.StatusCreated, httpCode)
	}

	c := created["data"].(map[string]interface{})

	if c["id"] == "" || c["id"] == nil {
		t.Fatal("expected non-empty product id")
	}

	want := map[string]interface{}{
		"status_code":    "CDC-200",
		"status_message": "OK",
		"data": map[string]interface{}{
			"id":        c["id"],
			"email":     "peterpan@gmail.com",
			"is_active": false,
			"username":  "peterpan",
		},
	}

	if diff := cmp.Diff(want, created); diff != "" {
		t.Fatalf("Response did not match expected. Diff:\n%s", diff)
	}
}
