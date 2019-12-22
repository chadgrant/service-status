package repository

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/chadgrant/service-status/api"
	"github.com/chadgrant/service-status/api/repository/mysql"
	"github.com/google/uuid"
)

func TestEnvironmentRepo(t *testing.T) {

	if len(os.Getenv("TEST_INTEGRATION")) == 0 {
		t.Log("Skipping integration tests, TEST_INTEGRATION environment variable not set")
		return
	}

	repo := mysql.NewEnvironmentRepository("localhost", 3306, "docker", "password", "service_status")

	t.Run("Add", func(t *testing.T) {
		testAddEnvironment(repo, t)
	})

	t.Run("Update", func(t *testing.T) {
		testUpdateEnvironment(repo, t)
	})

	t.Run("GetAll", func(t *testing.T) {
		testGetAllEnvironments(repo, t)
	})
}

func testAddEnvironment(repo EnvironmentRepository, t *testing.T) {
	n := uuid.New()
	t2 := time.Now().UTC()
	e := &api.Environment{Name: "Test " + n.String(), Friendly: "test-" + n.String(), Created: &t2, Active: false}
	if err := repo.Add(context.Background(), e); err != nil {
		t.Error(err)
	}
}

func testUpdateEnvironment(repo EnvironmentRepository, t *testing.T) {
	ctx := context.Background()
	n := uuid.New()
	t2 := time.Now().UTC()
	env := &api.Environment{Name: "Test " + n.String(), Friendly: "test-" + n.String(), Created: &t2, Active: false}
	if err := repo.Add(ctx, env); err != nil {
		t.Fatal(err)
	}

	all, err := repo.GetAll(ctx)
	if err != nil {
		t.Fatal(err)
	}

	new, err := find(all, func(e *api.Environment) bool { return e.Friendly == env.Friendly })
	if err != nil {
		t.Fatal(err)
	}

	t3 := time.Now().UTC()
	new.Name = "Updated " + new.Name
	env.Active = false
	new.Updated = &t3
	if err := repo.Update(ctx, new.Friendly, new); err != nil {
		t.Error(err)
	}

	all, err = repo.GetAll(ctx)
	if err != nil {
		t.Error(err)
	}

	u, err := find(all, func(e *api.Environment) bool { return e.Friendly == env.Friendly })
	if err != nil {
		t.Error(err)
	}

	if u.Name != new.Name || u.Active {
		t.Errorf("env not updated")
	}
}

func testGetAllEnvironments(repo EnvironmentRepository, t *testing.T) {
	e, err := repo.GetAll(context.Background())
	if err != nil {
		t.Error(err)
	}

	if len(e) == 0 {
		t.Errorf("no environments returned")
	}
}

func find(arr []*api.Environment, test func(*api.Environment) bool) (*api.Environment, error) {
	for _, v := range arr {
		if test(v) {
			return v, nil
		}
	}
	return nil, fmt.Errorf("did not find environment")
}
