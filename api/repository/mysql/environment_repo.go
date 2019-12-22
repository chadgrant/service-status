package mysql

import (
	"context"
	"fmt"

	"github.com/chadgrant/service-status/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type EnvironmentRepository struct {
	driver  string
	connStr string
}

func NewEnvironmentRepository(host string, port int, user, password, dbname string) *EnvironmentRepository {
	return &EnvironmentRepository{
		driver:  "mysql",
		connStr: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbname),
	}
}

func (r *EnvironmentRepository) GetAll(ctx context.Context) ([]*api.Environment, error) {
	db, err := sqlx.Connect(r.driver, r.connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	envs := []*api.Environment{}
	if err := db.SelectContext(ctx, &envs, "select name,friendly_name,cast(active as unsigned) as active,sort,created,updated from environment order by sort desc, name asc"); err != nil {
		return nil, err
	}
	return envs, nil
}

func (r *EnvironmentRepository) Add(ctx context.Context, env *api.Environment) error {
	return r.Update(ctx, env.Friendly, env)
}

func (r *EnvironmentRepository) Update(ctx context.Context, friendly string, env *api.Environment) error {
	db, err := sqlx.Connect(r.driver, r.connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, "call environment_upsert(?, ?, ?, ?, ?);", friendly, env.Name, env.Friendly, env.Active, env.Sort)
	return err
}
