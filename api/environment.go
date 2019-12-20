package api

import "time"

type Environment struct {
	Name     string     `json:"name"`
	Friendly string     `json:"friendly"`
	Active   bool       `json:"active"`
	Sort     int        `json:"sort"`
	Created  *time.Time `json:"created"`
	Updated  *time.Time `json:"updated,omitempty"`
}
