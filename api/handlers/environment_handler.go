package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/chadgrant/service-status/api"
	"github.com/chadgrant/service-status/api/repository"
	"github.com/gorilla/mux"
)

type (
	EnvironmentHandler struct {
		repo repository.EnvironmentRepository
	}

	// wrapper response to avoid returning an array (security)
	// swagger:response EnvironmentsResponse
	EnvironmentsResponse struct {
		// in:body
		Results []*api.Environment `json:"results"`
	}
)

func NewEnvironmentHandler(repo repository.EnvironmentRepository) *EnvironmentHandler {
	return &EnvironmentHandler{
		repo: repo,
	}
}

// GetAll gets all environments
// swagger:route GET /environments environments getenvironments
// Gets all environments 2
// responses:
//		200: EnvironmentsResponse
func (h *EnvironmentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	envs, err := h.repo.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnJSON(w, r, envs)
}

// Add adds an environment
// swagger:route POST /environments environments addenvironment
// Adds an environment
// responses:
//		201:
//			description: the resource was created, check location header for location to retrieve resource
//		400:
//			description: errors deserializing or validating request
//		500:
//			description: server error
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

// Update updates an environment
// swagger:route PUT /environments environments putenvironment
// Updates an environment
// responses:
//		204:
//			description: the resource was sucessfully updated.
//		400:
//			description: errors deserializing or validating request
//		500:
//			description: server error
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
