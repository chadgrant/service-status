package handlers

import (
	"fmt"
	"strconv"
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

func getVarOrDefault(vars map[string]string, name string, defaultVal string) string {
	if v, ok := vars[name]; ok && len(v) > 0 {
		return v
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
