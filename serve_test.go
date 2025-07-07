package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/al-masood/book_server/server"
)

var testBooks = []map[string]interface{}{
	{
		"uuid":        "1",
		"name":        "a",
		"authorList":  []string{"b"},
		"publishDate": "24-02-2002",
		"isbn":        "1001-1001-1001-1002",
	},
	{
		"uuid":        "2",
		"name":        "c",
		"authorList":  []string{"d", "e"},
		"publishDate": "25-03-2003",
		"isbn":        "1001-1001-1001-1003",
	},
	{
		"uuid":        "2",
		"name":        "e",
		"authorList":  []string{"f"},
		"publishDate": "26-04-2004",
		"isbn":        "1001-1001-1001-1004",
	},
}

func setupTestServer() *server.Server {
	return server.NewServer(false)
}

func createRandomBooks(testServer *server.Server) {
	for i := 0; i <= 1; i++ {
		body, _ := json.Marshal(testBooks[i])
		req := httptest.NewRequest(http.MethodPost, "/api/v1/books", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()

		testServer.Router.ServeHTTP(res, req)
	}
}

func TestCreateBook(t *testing.T) {
	testServer := setupTestServer()

	for i := 0; i <= 1; i++ {
		body, err := json.Marshal(testBooks[i])
		if err != nil {
			t.Fatalf("Failed to marshal book data: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/books", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()

		testServer.Router.ServeHTTP(res, req)

		if res.Code != http.StatusCreated {  // Usually Create returns 201 Created
			t.Errorf("Failed to create book, status code: %d, response body: %s", res.Code, res.Body.String())
		}
	}
}



func TestGetAllBooks(t *testing.T) {
	testServer := setupTestServer()

	createRandomBooks(testServer)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/books", nil)
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	testServer. Router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Error getting all books")
	}
}


func TestGetBooksById(t *testing.T) {
	testServer := setupTestServer()

	createRandomBooks(testServer)

	id := "1"

	req := httptest.NewRequest(http.MethodGet, "/api/v1/books/"+id, nil)
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	testServer. Router.ServeHTTP(res, req)

	if res.Code != http.StatusOK{
		t.Errorf("Error getting by id")
	}

	id = "3"

	req = httptest.NewRequest(http.MethodGet, "/api/v1/books/"+id, nil)
	req.Header.Set("Content-Type", "application/json")

	res = httptest.NewRecorder()

	testServer. Router.ServeHTTP(res, req)

	if res.Code != http.StatusBadRequest{
		t.Errorf("Error getting book with invalid id")
	}
}

func TestUpdateBook(t *testing.T) {
	testServer := setupTestServer()

	createRandomBooks(testServer)

	id := "1"

	req := httptest.NewRequest(http.MethodGet, "/api/v1/books/"+id, nil)
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	testServer. Router.ServeHTTP(res, req)

	if res.Code != http.StatusOK{
		t.Errorf("Error updating book")
	}
}

func TestDeleteBook(t *testing.T) {
	testServer := setupTestServer()

	createRandomBooks(testServer)

	id := "1"

	req := httptest.NewRequest(http.MethodGet, "/api/v1/books/"+id, nil)
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	testServer. Router.ServeHTTP(res, req)

	if res.Code != http.StatusOK{
		t.Errorf("Error deleting book")
	}
}

func TestGetToken(t *testing.T) {
	testServer := setupTestServer()

	createRandomBooks(testServer)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/get-token", nil)
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	testServer. Router.ServeHTTP(res, req)

	if res.Code != http.StatusOK{
		t.Errorf("Error getting token")
	}
}
