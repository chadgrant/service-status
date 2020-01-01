package repository

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/chadgrant/servicestatus/api"
	"github.com/chadgrant/servicestatus/api/repository/mysql"
	"github.com/google/uuid"
)

func TestServiceRepo(t *testing.T) {

	if len(os.Getenv("TEST_INTEGRATION")) == 0 {
		t.Log("Skipping integration tests, TEST_INTEGRATION environment variable not set")
		return
	}

	repo := mysql.NewServiceRepository("localhost", 3306, "docker", "password", "service_status")

	t.Run("Add", func(t *testing.T) {
		testAddService(repo, t)
	})

	t.Run("Update", func(t *testing.T) {
		testUpdateService(repo, t)
	})

	t.Run("GetAll", func(t *testing.T) {
		testGetAllServices(repo, t)
	})

	t.Run("GetForEnvironment", func(t *testing.T) {
		testGetServicesForEnvironment(repo, t)
	})
}

func testAddService(repo ServiceRepository, t *testing.T) {
	n := uuid.New()
	t2 := time.Now().UTC()
	s := &api.Service{Name: "Test " + n.String(), Friendly: "test-" + n.String(), Status: "testing", Created: &t2, Active: false}
	if err := repo.Add(context.Background(), s); err != nil {
		t.Error(err)
	}
}

func testUpdateService(repo ServiceRepository, t *testing.T) {
	ctx := context.Background()
	n := uuid.New()
	t2 := time.Now().UTC()
	svc := &api.Service{Name: "Test " + n.String(), Friendly: "test-" + n.String(), Status: "testing", Created: &t2, Active: false}
	if err := repo.Add(ctx, svc); err != nil {
		t.Fatal(err)
	}

	all, err := repo.GetAll(ctx)
	if err != nil {
		t.Fatal(err)
	}

	new, err := findservice(all, func(s *api.Service) bool { return s.Friendly == svc.Friendly })
	if err != nil {
		t.Fatal(err)
	}

	t3 := time.Now().UTC()
	new.Name = "Updated " + new.Name
	svc.Active = false
	new.Updated = &t3
	if err := repo.Update(ctx, new.Friendly, new); err != nil {
		t.Error(err)
	}

	all, err = repo.GetAll(ctx)
	if err != nil {
		t.Error(err)
	}

	u, err := findservice(all, func(s *api.Service) bool { return s.Friendly == svc.Friendly })
	if err != nil {
		t.Error(err)
	}

	if u.Name != new.Name || u.Active {
		t.Errorf("env not updated")
	}
}

func testGetAllServices(repo ServiceRepository, t *testing.T) {
	e, err := repo.GetAll(context.Background())
	if err != nil {
		t.Error(err)
	}

	if len(e) == 0 {
		t.Errorf("no environments returned")
	}
}

func testGetServicesForEnvironment(repo ServiceRepository, t *testing.T) {
	e, err := repo.GetForEnvironment(context.Background(), "development")
	if err != nil {
		t.Error(err)
	}

	if len(e) == 0 {
		t.Errorf("no environments returned")
	}
}

func findservice(arr []*api.Service, test func(*api.Service) bool) (*api.Service, error) {
	for _, v := range arr {
		if test(v) {
			return v, nil
		}
	}
	return nil, fmt.Errorf("did not find service")
}
