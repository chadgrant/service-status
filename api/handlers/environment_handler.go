package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/chadgrant/go/http/infra"
	"github.com/chadgrant/service-status/api"
	"github.com/chadgrant/service-status/api/repository"
	"github.com/gorilla/mux"
)

type EnvironmentHandler struct {
	repo repository.EnvironmentRepository
}

func NewEnvironmentHandler(repo repository.EnvironmentRepository) *EnvironmentHandler {
	return &EnvironmentHandler{
		repo: repo,
	}
}

func (h *EnvironmentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	envs, err := h.repo.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnJson(w, r, envs)
}

func (h *EnvironmentHandler) Add(w http.ResponseWriter, r *http.Request) {
	var e api.Environment
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t := time.Now().UTC()
	e.Created = &t

	if err := h.repo.Add(r.Context(), &e); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/environment/%s", e.Friendly))
	w.WriteHeader(http.StatusCreated)
}

func (h *EnvironmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	friendly := vars["friendly"]

	var e api.Environment
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t := time.Now().UTC()
	e.Friendly = friendly
	e.Updated = &t

	if err := h.repo.Update(r.Context(), friendly, &e); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func returnJson(w http.ResponseWriter, r *http.Request, o interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(o); err != nil {
		infra.Error(w, r, http.StatusInternalServerError, err)
	}
}
