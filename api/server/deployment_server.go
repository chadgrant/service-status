package server

import(
	"context"
	"time"

	"github.com/chadgrant/service-status/api"
	"github.com/chadgrant/service-status/api/repository"
	"github.com/chadgrant/service-status/api/generated"
)

type DeploymentServer struct {
	repo repository.DeploymentRepository
}

func NewDeploymentServer(repo repository.DeploymentRepository) *DeploymentServer {
	return &DeploymentServer{repo:repo}
}

func (s *DeploymentServer) Get(ctx context.Context, req *generated.GetDeploymentsRequest) (*generated.DeploymentsPaged, error) {
	t,depls,err := s.repo.GetPaged(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil,err
	}
	return &generated.DeploymentsPaged{ 
		Results: copyDeployments(depls), 
		Total: int32(t), 
		Page: req.Page, 
		PageSize: req.PageSize, 
	}, nil
}

func (s *DeploymentServer) GetForEnvironment(ctx context.Context, req *generated.GetDeploymentsRequest) (*generated.DeploymentsPaged, error) {
	t,depls,err := s.repo.GetForEnvironmentPaged(ctx, req.Environment, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil,err
	}
	return &generated.DeploymentsPaged{ 
		Results: copyDeployments(depls), 
		Total: int32(t), 
		Page: req.Page, 
		PageSize: req.PageSize,
	}, nil
}

func (s *DeploymentServer) GetForService(ctx context.Context, req *generated.GetDeploymentsRequest) (*generated.DeploymentsPaged, error) {
	t,depls,err := s.repo.GetForServicePaged(ctx, req.Environment, req.Service, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil,err
	}
	return &generated.DeploymentsPaged{ 
		Results: copyDeployments(depls), 
		Total: int32(t), 
		Page: req.Page, 
		PageSize: req.PageSize ,
	}, nil
}

func (s *DeploymentServer) Add(ctx context.Context, depl *generated.Deployment) (*generated.Deployment, error) {
	t := time.Now().UTC()
	a := copyGeneratedDeployment(depl)
	a.Created = &t
	if err:= s.repo.Add(ctx, a);err != nil {
		return nil,err
	}
	return copyDeployment(a),nil
}

func copyGeneratedDeployment(g *generated.Deployment) *api.Deployment {
	return &api.Deployment {
		ID : g.ID,
		Service: g.Service,
		Environment: g.Environment,
		BuildNumber: g.BuildNumber,
		Created: g.Created,
		Updated: g.Updated,
	}
}

func copyDeployments(e []*api.Deployment) []*generated.Deployment {
	res := make([]*generated.Deployment,len(e))
	for i,g := range e {
		res[i] = copyDeployment(g)
	}
	return res
}

func copyDeployment(d *api.Deployment) *generated.Deployment {
	return &generated.Deployment {
		ID : d.ID,
		Service: d.Service,
		Environment: d.Environment,
		BuildNumber: d.BuildNumber,
		Created: d.Created,
		Updated: d.Updated,
	}
}
