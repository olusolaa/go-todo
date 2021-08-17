package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/olusolaa/go-todo/pkg/models"
	"github.com/olusolaa/go-todo/pkg/util"
	"net/http"
	"net/http/httptest"
	"testing"
)
func CreateTodo(w http.ResponseWriter, r *http.Request)  {
	createTodo := &models.Todo{}
	util.ParseBody(r, createTodo)
	t:= createTodo.CreateTodo()
	res, _ :=json.Marshal(t)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}


func GetTodo(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	NewTodos := models.GetAllTodo(code)
	res, _ := json.Marshal(NewTodos)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func NewServer(port string) *http.Server {
	addr := fmt.Sprintf(":%s", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}



var (
	port = "8080"
)

func TestHandler(t *testing.T) {
	expected := []byte("Hello World")

	req, err := http.NewRequest("GET", buildUrl("/"), nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	handler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Response code was %v; want 200", res.Code)
	}

	if bytes.Compare(expected, res.Body.Bytes()) != 0 {
		t.Errorf("Response body was '%v'; want '%v'", expected, res.Body)
	}
}

func buildUrl(path string) string {
	return urlFor("http", port, path)
}

func urlFor(scheme string, serverPort string, path string) string {
	return scheme + "://localhost:" + serverPort + path
}


func main() {

}
