package mysql

import (
	"context"
	"fmt"

	"github.com/chadgrant/servicestatus/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DeploymentRepository struct {
	driver  string
	connStr string
}

func NewDeploymentRepository(host string, port int, user, password, dbname string) *DeploymentRepository {
	return &DeploymentRepository{
		driver:  "mysql",
		connStr: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&multiStatements=true&interpolateParams=true", user, password, host, port, dbname),
	}
}

func (r *DeploymentRepository) GetPaged(ctx context.Context, page, size int) (int, []*api.Deployment, error) {
	return r.getPaged(ctx, "", "", page, size)
}

func (r *DeploymentRepository) GetForEnvironmentPaged(ctx context.Context, environment string, page, size int) (int, []*api.Deployment, error) {
	return r.getPaged(ctx, environment, "", page, size)
}

func (r *DeploymentRepository) GetForServicePaged(ctx context.Context, environment, service string, page, size int) (int, []*api.Deployment, error) {
	return r.getPaged(ctx, environment, service, page, size)
}

func (r *DeploymentRepository) Add(ctx context.Context, deploy *api.Deployment) error {
	db, err := sqlx.Connect(r.driver, r.connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.QueryxContext(ctx, "call deployment_insert(?, ?, ?, @nid); select @nid;", deploy.Environment, deploy.Service, deploy.BuildNumber)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&deploy.ID)
		if err != nil {
			return err
		}
	}
	return rows.Err()
}

func (r *DeploymentRepository) getPaged(ctx context.Context, environment, service string, page, size int) (int, []*api.Deployment, error) {
	db, err := sqlx.Connect(r.driver, r.connStr)
	if err != nil {
		return 0, nil, err
	}
	defer db.Close()

	rows, err := db.QueryxContext(ctx, `call deployment_paged(?,?,?,?,@total); select @total;`, environment, service, page, size)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	deploys := make([]*api.Deployment, 0)

	for rows.Next() {
		d := &api.Deployment{}
		if err := rows.Scan(&d.ID, &d.Environment, &d.Service, &d.BuildNumber, &d.Created, &d.Updated); err != nil {
			return 0, nil, err
		}
		deploys = append(deploys, d)
	}

	var t int
	if rows.NextResultSet() {
		for rows.Next() {
			if err := rows.Scan(&t); err != nil {
				return 0, nil, err
			}
		}
	}

	return t, deploys, rows.Err()
}
