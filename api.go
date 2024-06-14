package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
}

type ApiError struct {
	Error string
}
type apiFunc func(http.ResponseWriter, *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
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

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

func (server *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(server.handleAccount))
}

func (server *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error    { return nil }
func (server *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error { return nil }
func (server *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (server *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (server *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error { return nil }
