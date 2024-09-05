package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}

	}
}

type APIServer struct {
	listenAddress string
	store         storage
}

func NewAPIServer(listenAddr string, store storage) *APIServer {

	return &APIServer{
		listenAddress: listenAddr,
		store:         store,
	}
}

func (s *APIServer) Run() {

	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleAccount))

	log.Println("JSON API server running on port:", s.listenAddress)

	http.ListenAndServe(s.listenAddress, router)

}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetAccounts(w, r)
	}
	if r.Method == "GET" && mux.Vars(r)["id"] != "" {
		return s.handleGetAccountByID(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("Method not allowed %S", r.Method)
}

func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {

	id := mux.Vars(r)["id"]

	account, err := s.store.GetAccountByID(id)
	if err != nil {

		return err
	}

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {

	createAccReq := new(CreateAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(createAccReq); err != nil {

		return err
	}

	account := NewAccount(createAccReq.FirstName, createAccReq.LastName)

	if err := s.store.CreateAccount(account); err != nil {

		return err
	}

	return WriteJSON(w, http.StatusCreated, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	id := vars["id"]
	err := s.store.DeleteAccount(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusNoContent, nil)
}

func (s *APIServer) handleTransferAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}

//GET ALL ACCOUNTS

func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {

	accounts, err := s.store.GetAccounts()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, accounts)
}
