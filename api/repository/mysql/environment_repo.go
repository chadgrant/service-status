package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/chadgrant/service-status/api"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
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
	envs := make([]*api.Environment, 0)

	db, err := sql.Open(r.driver, r.connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.QueryContext(ctx, "select name,friendly_name,cast(active as unsigned),sort,created,updated from environment order by sort desc, name asc")
	if err != nil {
		return envs, err
	}

	for rows.Next() {
		e := &api.Environment{}
		var nt mysql.NullTime
		var b uint8
		if err := rows.Scan(&e.Name, &e.Friendly, &b, &e.Sort, &e.Created, &nt); err != nil {
			return envs, err
		}

		e.Active = b > 0

		if nt.Valid {
			e.Updated = &nt.Time
		}
		envs = append(envs, e)
	}

	return envs, nil
}

func (r *EnvironmentRepository) Add(ctx context.Context, env *api.Environment) error {
	return r.Update(ctx, env.Friendly, env)
}

func (r *EnvironmentRepository) Update(ctx context.Context, friendly string, env *api.Environment) error {
	db, err := sql.Open(r.driver, r.connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.PrepareContext(ctx, "call environment_upsert(?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	var b uint8
	if env.Active {
		b = 1
	}

	_, err = stmt.ExecContext(ctx, friendly, env.Name, env.Friendly, b, env.Sort)
	if err != nil {
		return err
	}

	return nil
}
