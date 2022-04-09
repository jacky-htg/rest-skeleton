package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"rest/router"
	"rest/schema"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	db, teardown := NewUnit(t)
	defer teardown()

	if err := schema.Seed(db); err != nil {
		t.Fatal(err)
	}

	log := log.New(os.Stderr, "TEST : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// api test for users
	{
		users := Users{App: router.API(db, log)}
		t.Run("APiUsersCreate", users.Create)
	}
}

func callApi(h http.Handler, method string, path string, jsonBody string, data *map[string]interface{}) (error, int) {
	body := strings.NewReader(jsonBody)
	req := httptest.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	h.ServeHTTP(resp, req)

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err, 0
	}

	return nil, resp.Code
}
