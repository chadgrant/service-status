package repository

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/chadgrant/service-status/api"
	"github.com/chadgrant/service-status/api/repository/mysql"
)

func TestDeploymentRepo(t *testing.T) {

	if len(os.Getenv("TEST_INTEGRATION")) == 0 {
		t.Log("Skipping integration tests, TEST_INTEGRATION environment variable not set")
		return
	}

	repo := mysql.NewDeploymentRepository("localhost", 3306, "docker", "password", "service_status")

	t.Run("Add", func(t *testing.T) {
		testAddDeployment(repo, t)
	})

	t.Run("GetPaged", func(t *testing.T) {
		testGetPagedDeployments(repo, t)
	})

	t.Run("GetPagedEnvironment", func(t *testing.T) {
		testGetPagedDeploymentsForEnvironment(repo, t)
	})

	t.Run("GetPagedService", func(t *testing.T) {
		testGetPagedDeploymentsForService(repo, t)
	})
}

func testAddDeployment(repo DeploymentRepository, t *testing.T) {
	t2 := time.Now().UTC()
	d := &api.Deployment{Environment: "development", Service: "sample_service", BuildNumber: "6.6.666", Created: &t2}
	if err := repo.Add(context.Background(), d); err != nil {
		t.Error(err)
	}
	if len(d.ID) == 0 {
		t.Errorf("id was not set on deploy")
	}
}

func testGetPagedDeployments(repo DeploymentRepository, t *testing.T) {
	tot, deploys, err := repo.GetPaged(context.Background(), 1, 25)
	if err != nil {
		t.Error(err)
	}

	if len(deploys) == 0 {
		t.Errorf("no results returned")
	}

	if tot <= 0 {
		t.Errorf("total is less than or equal to zero")
	}
}

func testGetPagedDeploymentsForEnvironment(repo DeploymentRepository, t *testing.T) {
	tot, deploys, err := repo.GetForEnvironmentPaged(context.Background(), "development", 1, 25)
	if err != nil {
		t.Error(err)
	}

	if len(deploys) == 0 {
		t.Errorf("no results returned")
	}

	if tot <= 0 {
		t.Errorf("total is less than or equal to zero")
	}
}

func testGetPagedDeploymentsForService(repo DeploymentRepository, t *testing.T) {
	tot, deploys, err := repo.GetForServicePaged(context.Background(), "development", "sample_service", 1, 25)
	if err != nil {
		t.Error(err)
	}

	if len(deploys) == 0 {
		t.Errorf("no results returned")
	}

	if tot <= 0 {
		t.Errorf("total is less than or equal to zero")
	}
}
