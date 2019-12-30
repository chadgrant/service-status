// Code generated by goa v3.0.9, DO NOT EDIT.
//
// HTTP request path constructors for the environment service.
//
// Command:
// $ goa gen servicestatus/design

package client

import (
	"fmt"
)

// ListEnvironmentPath returns the URL path to the environment service list HTTP endpoint.
func ListEnvironmentPath() string {
	return "/v1/environment"
}

// AddEnvironmentPath returns the URL path to the environment service add HTTP endpoint.
func AddEnvironmentPath() string {
	return "/v1/environment"
}

// UpdateEnvironmentPath returns the URL path to the environment service update HTTP endpoint.
func UpdateEnvironmentPath(key string) string {
	return fmt.Sprintf("/v1/environment/%v", key)
}
