package server

import (
	"context"
	"time"

	"github.com/chadgrant/service-status/api"
	"github.com/chadgrant/service-status/api/generated"
	"github.com/chadgrant/service-status/api/repository"
)

type ServiceServer struct {
	serviceMapper
	repo repository.ServiceRepository
}

type serviceMapper struct{}

func (*serviceMapper) Generated(r *api.Service) *generated.Service {
	return &generated.Service{
		Name:     r.Name,
		Friendly: r.Friendly,
		Active:   r.Active,
		Status:   r.Status,
		Sort:     int32(r.Sort),
		Created:  r.Created,
		Updated:  r.Updated,
	}
}

func (m *serviceMapper) GeneratedMany(src []*api.Service) []*generated.Service {
	dest := make([]*generated.Service, len(src))
	for i, v := range src {
		dest[i] = m.Generated(v)
	}
	return dest
}

func (*serviceMapper) Real(g *generated.Service) *api.Service {
	return &api.Service{
		Name:     g.Name,
		Friendly: g.Friendly,
		Active:   g.Active,
		Status:   g.Status,
		Sort:     int(g.Sort),
		Created:  g.Created,
		Updated:  g.Updated,
	}
}

func NewServiceServer(repo repository.ServiceRepository) *ServiceServer {
	return &ServiceServer{
		repo:          repo,
		serviceMapper: serviceMapper{},
	}
}

func (s *ServiceServer) Add(ctx context.Context, req *generated.Service) (*generated.Service, error) {
	t := time.Now().UTC()
	svc := s.Real(req)
	svc.Created = &t
	if err := s.repo.Add(ctx, svc); err != nil {
		return nil, err
	}
	return s.Generated(svc), nil
}

func (s *ServiceServer) Get(ctx context.Context, req *generated.GetServicesRequest) (*generated.ServicesPaged, error) {
	svcs, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	resp := &generated.ServicesPaged{Results: s.GeneratedMany(svcs)}
	return resp, nil
}

func (s *ServiceServer) GetForEnvironment(ctx context.Context, req *generated.GetServicesRequest) (*generated.ServicesPaged, error) {
	svcs, err := s.repo.GetForEnvironment(ctx, req.Environment)
	if err != nil {
		return nil, err
	}
	resp := &generated.ServicesPaged{Results: s.GeneratedMany(svcs)}
	return resp, nil
}

func (s *ServiceServer) Update(ctx context.Context, req *generated.Service) (*generated.Service, error) {
	t := time.Now().UTC()
	svc := s.Real(req)
	svc.Updated = &t
	if err := s.repo.Update(ctx, req.Key, svc); err != nil {
		return nil, err
	}
	return s.Generated(svc), nil
}
