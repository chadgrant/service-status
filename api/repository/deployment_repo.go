package repository

import (
	"context"

	"github.com/chadgrant/service-status/api"
)

type DeploymentRepository interface {
	GetPaged(context.Context, int, int) (int, []*api.Deployment, error)
	GetForEnvironmentPaged(context.Context, string, int, int) (int, []*api.Deployment, error)
	GetForServicePaged(context.Context, string, string, int, int) (int, []*api.Deployment, error)
	Add(context.Context, *api.Deployment) error
}
