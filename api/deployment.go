package api

import "time"

type Deployment struct {
	ID          string     `json:"id" db:"id"`
	Service     string     `json:"service,omitempty" db:"service"`
	Environment string     `json:"environment,omitempty" db:"environment"`
	BuildNumber string     `json:"build_number" db:"build_number"`
	Created     *time.Time `json:"created" db:"created"`
	Updated     *time.Time `json:"updated,omitempty" db:"updated"`
}
