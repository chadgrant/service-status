package api

import "time"

// Environment entity
// swagger:parameters addenvironment
// swagger:parameters putenvironment
type Environment struct {
	Name     string     `json:"name" db:"name"`
	Friendly string     `json:"friendly" db:"friendly_name"`
	Active   bool       `json:"active" db:"active"`
	Sort     int        `json:"sort" db:"sort"`
	Created  *time.Time `json:"created" db:"created"`
	Updated  *time.Time `json:"updated,omitempty" db:"updated"`
}
