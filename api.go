package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
	store         Storage
}

func NewAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (server *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(server.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(server.handleGetAccount))

	log.Println("JSON API Server running on port:", server.listenAddress)
	http.ListenAndServe(server.listenAddress, router)
}

func (server *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return server.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return server.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return server.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}
func (server *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	// account := NewAccount("Maccbrainy", "Michael")
	id := mux.Vars(r)["id"]
	fmt.Printf(id)
	return WriteJSON(w, http.StatusOK, &Account{})
}
func (server *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (server *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (server *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error { return nil }

type ApiError struct {
	Error string
}
type apiFunc func(http.ResponseWriter, *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//Handle the error
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})

		}
	}
}
