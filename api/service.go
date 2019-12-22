package api

import "time"

type Service struct {
	Name     string     `json:"name" db:"name"`
	Friendly string     `json:"friendly" db:"friendly_name"`
	Active   bool       `json:"active" db:"active"`
	Status   string     `json:"status" db:"status"`
	Sort     int        `json:"sort" db:"sort"`
	Created  *time.Time `json:"created" db:"created"`
	Updated  *time.Time `json:"updated,omitempty" db:"updated"`
}

type ServiceEnvironment struct {
	Service
	Deployment *Deployment `json:"deployment"`
}
