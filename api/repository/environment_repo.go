package repository

import (
	"context"

	"github.com/chadgrant/service-status/api"
)

type EnvironmentRepository interface {
	GetAll(context.Context) ([]*api.Environment, error)
	Add(context.Context, *api.Environment) error
	Update(context.Context, string, *api.Environment) error
}
