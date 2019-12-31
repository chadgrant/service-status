package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/chadgrant/servicestatus/api"
	"github.com/chadgrant/servicestatus/api/repository"
	"github.com/gorilla/mux"
)

type DeploymentHandler struct {
	repo repository.DeploymentRepository
}

type PagedDeploys struct {
	Page    int               `json:"page"`
	Size    int               `json:"size"`
	Total   int               `json:"total"`
	Next    string            `json:"next,omitempty"`
	Prev    string            `json:"prev,omitempty"`
	Results []*api.Deployment `json:"results"`
}

func NewDeploymentHandler(repo repository.DeploymentRepository) *DeploymentHandler {
	return &DeploymentHandler{
		repo: repo,
	}
}

func (h *DeploymentHandler) GetPaged(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := getIntVarOrDefault(vars, "page", 1)
	size := getIntVarOrDefault(vars, "size", 25)

	t, deploys, err := h.repo.GetPaged(r.Context(), page, size)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnJSON(w, r, &PagedDeploys{Page: page, Size: size, Total: t, Results: deploys,
		Next: nextLink("/deployments", page, size, t),
		Prev: prevLink("/deployments", page, size)})
}

func (h *DeploymentHandler) GetForEnvironmentPaged(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	env := vars["environment"]
	page := getIntVarOrDefault(vars, "page", 1)
	size := getIntVarOrDefault(vars, "size", 25)

	t, deploys, err := h.repo.GetForEnvironmentPaged(r.Context(), env, page, size)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnJSON(w, r, &PagedDeploys{Page: page, Size: size, Total: t, Results: deploys,
		Next: nextLink(fmt.Sprintf("/environment/%s/deployments", env), page, size, t),
		Prev: prevLink(fmt.Sprintf("/environment/%s/deployments", env), page, size)})
}

func (h *DeploymentHandler) GetForServicePaged(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	env := vars["environment"]
	svc := vars["service"]
	page := getIntVarOrDefault(vars, "page", 1)
	size := getIntVarOrDefault(vars, "size", 25)

	t, deploys, err := h.repo.GetForServicePaged(r.Context(), env, svc, page, size)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnJSON(w, r, &PagedDeploys{Page: page, Size: size, Total: t, Results: deploys,
		Next: nextLink(fmt.Sprintf("/environment/%s/service/%s/deployments", env, svc), page, size, t),
		Prev: prevLink(fmt.Sprintf("/environment/%s/service/%s/deployments", env, svc), page, size)})
}

func (h *DeploymentHandler) Add(w http.ResponseWriter, r *http.Request) {
	var d api.Deployment
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t := time.Now().UTC()
	d.Created = &t

	if err := h.repo.Add(r.Context(), &d); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/service/%s/deployment/%s", d.Service, d.ID))
	w.WriteHeader(http.StatusCreated)
}
