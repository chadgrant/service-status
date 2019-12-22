package repository

import (
	"context"

	"github.com/chadgrant/service-status/api"
)

type ServiceRepository interface {
	GetAll(context.Context) ([]*api.Service, error)
	GetForEnvironment(context.Context, string) ([]*api.Service, error)
	Add(context.Context, *api.Service) error
	Update(context.Context, string, *api.Service) error
}
