package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type APIServer struct {
	listenAddr string
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//handle error here
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

// NewAPIServer creates a new APIServer
func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{listenAddr: listenAddr}
}

// Start starts the APIServer
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

	log.Println("JSON API server running on port ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	} else if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	} else if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed")
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {

	id := mux.Vars(r)["id"]

	fmt.Println("ID: ", id)

	//db.get(id)...

	//account := NewAccount("John", "Doe")
	return WriteJSON(w, http.StatusOK, &Account{})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (s *APIServer) handleTransferAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}
