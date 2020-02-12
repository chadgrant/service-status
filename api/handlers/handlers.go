package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chadgrant/go-http-infra/infra"
)

func getIntVarOrDefault(vars map[string]string, name string, defaultVal int) int {
	if v, ok := vars[name]; ok && len(v) > 0 {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i
		}
	}
	return defaultVal
}

func nextLink(url string, page, size, total int) string {
	if page*size > total {
		return ""
	}
	return fmt.Sprintf("%s?page=%d&size=%d", url, page+1, size)
}

func prevLink(url string, page, size int) string {
	if page <= 1 {
		return ""
	}
	return fmt.Sprintf("%s?page=%d&size=%d", url, page-1, size)
}

func returnJSON(w http.ResponseWriter, r *http.Request, o interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(o); err != nil {
		infra.Error(w, r, http.StatusInternalServerError, err)
	}
}
