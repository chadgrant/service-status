package server

import(
	"context"
	"time"

	"github.com/chadgrant/service-status/api/repository"
	"github.com/chadgrant/service-status/api/generated"
	"github.com/chadgrant/service-status/api"
)

type EnvironmentServer struct {
	repo repository.EnvironmentRepository
}

func NewEnvironmentServer(repo repository.EnvironmentRepository) *EnvironmentServer {
	return &EnvironmentServer{repo:repo}
}

func (s *EnvironmentServer) Add(ctx context.Context, env *generated.Environment) (*generated.Environment, error) {
	t := time.Now().UTC()
	e := copyGeneratedEnvironment(env)
	e.Created = &t
	if err:= s.repo.Add(ctx, e);err != nil {
		return nil,err
	}
	return copyEnvironment(e),nil
}

func (s *EnvironmentServer) Get(ctx context.Context, req *generated.GetEnvironmentsRequest) (*generated.EnvironmentsPaged, error) {
	envs,err := s.repo.GetAll(ctx)
	if err != nil {
		return nil,err
	}
	return &generated.EnvironmentsPaged{ Results: copyEnvironments(envs) }, nil
}

func (s *EnvironmentServer) Update(ctx context.Context, env *generated.Environment) (*generated.Environment, error) {
	t := time.Now().UTC()
	e := copyGeneratedEnvironment(env)
	e.Updated = &t
	if err := s.repo.Update(ctx, env.Key, e);err != nil {
		return nil,err
	}
	return copyEnvironment(e),nil
}

func copyGeneratedEnvironment(g *generated.Environment) *api.Environment {
	return &api.Environment {
		Name : g.Name,
		Friendly: g.Friendly,
		Active: g.Active,
		Sort: int(g.Sort),
		Created: g.Created,
		Updated: g.Updated,
	}
}

func copyEnvironments(e []*api.Environment) []*generated.Environment {
	res := make([]*generated.Environment,len(e))
	for i,g := range e {
		res[i] = copyEnvironment(g)
	}
	return res
}

func copyEnvironment(e *api.Environment) *generated.Environment {
	return &generated.Environment {
		Name : e.Name,
		Friendly: e.Friendly,
		Active: e.Active,
		Sort: int32(e.Sort),
		Created: e.Created,
		Updated: e.Updated,
	}
}