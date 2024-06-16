package handler

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"your_project/service"
)

type PersonHandler struct {
	Service *service.PersonService
}

func (h *PersonHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/pessoas", h.CreatePerson).Methods("POST")
	router.HandleFunc("/pessoas/{id}", h.GetPerson).Methods("GET")
	router.HandleFunc("/pessoas", h.SearchPeople).Methods("GET")
	// Add additional routes here
}

func (h *PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	// Implementation goes here
}

func (h *PersonHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	// Implementation goes here
}

func (h *PersonHandler) SearchPeople(w http.ResponseWriter, r *http.Request) {
	// Implementation goes here
}
