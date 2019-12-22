package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chadgrant/service-status/api"
	"github.com/chadgrant/service-status/api/repository"
	"github.com/gorilla/mux"
)

type ServiceHandler struct {
	repo repository.ServiceRepository
}

func NewServiceHandler(repo repository.ServiceRepository) *ServiceHandler {
	return &ServiceHandler{
		repo: repo,
	}
}

func (h *ServiceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	svcs, err := h.repo.GetAll(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnJson(w, r, svcs)
}

func (h *ServiceHandler) GetForEnvironment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	env := vars["environment"]

	svcs, err := h.repo.GetForEnvironment(r.Context(), env)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnJson(w, r, svcs)
}

func (h *ServiceHandler) Add(w http.ResponseWriter, r *http.Request) {
	var s api.Service
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t := time.Now().UTC()
	s.Created = &t

	if err := h.repo.Add(r.Context(), &s); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/service/%s", s.Friendly))
	w.WriteHeader(http.StatusCreated)
}

func (h *ServiceHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	friendly := vars["friendly"]

	var s api.Service
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t := time.Now().UTC()
	s.Friendly = friendly
	s.Updated = &t

	if err := h.repo.Update(r.Context(), friendly, &s); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
