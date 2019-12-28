package server

import(
	"context"
	"time"

	"github.com/chadgrant/service-status/api"
	"github.com/chadgrant/service-status/api/repository"
	"github.com/chadgrant/service-status/api/generated"
)

type ServiceServer struct {
	repo repository.ServiceRepository
}

func NewServiceServer(repo repository.ServiceRepository) *ServiceServer {
	return &ServiceServer{repo:repo}
}

func (s *ServiceServer) Add(ctx context.Context, svc *generated.Service) (*generated.Service, error) {
	t := time.Now().UTC()
	a := copyGeneratedService(svc)
	a.Created = &t
	if err:= s.repo.Add(ctx, a);err != nil {
		return nil,err
	}
	return copyService(a),nil
}

func (s *ServiceServer) Get(ctx context.Context, req *generated.GetServicesRequest) (*generated.ServicesPaged, error) {
	svcs,err := s.repo.GetAll(ctx)
	if err != nil {
		return nil,err
	}
	resp:= &generated.ServicesPaged{ Results: copyServices(svcs) }
	return resp, nil
}

func (s *ServiceServer) GetForEnvironment(ctx context.Context, req *generated.GetServicesRequest) (*generated.ServicesPaged, error) {
	svcs,err := s.repo.GetForEnvironment(ctx, req.Environment)
	if err != nil {
		return nil,err
	}
	resp:= &generated.ServicesPaged{ Results: copyServices(svcs) }
	return resp, nil
}

func (s *ServiceServer) Update(ctx context.Context, svc *generated.Service) (*generated.Service, error) {
	t := time.Now().UTC()
	u := copyGeneratedService(svc)
	u.Updated = &t
	if err := s.repo.Update(ctx, svc.Key, u);err != nil {
		return nil,err
	}
	return copyService(u),nil
}

func copyGeneratedService(g *generated.Service) *api.Service {
	return &api.Service{
		Name : g.Name,
		Friendly: g.Friendly,
		Active: g.Active,
		Status: g.Status,
		Sort: int(g.Sort),
		Created: g.Created,
		Updated: g.Updated,
	}
}

func copyServices(e []*api.Service) []*generated.Service {
	res := make([]*generated.Service,len(e))
	for i,g := range e {
		res[i] = copyService(g)
	}
	return res
}

func copyService(s *api.Service) *generated.Service {
	return &generated.Service {
		Name : s.Name,
		Friendly: s.Friendly,
		Active: s.Active,
		Status: s.Status,
		Sort: int32(s.Sort),
		Created: s.Created,
		Updated: s.Updated,
	}
}
