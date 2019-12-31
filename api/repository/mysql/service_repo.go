package mysql

import (
	"context"
	"fmt"

	"github.com/chadgrant/servicestatus/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ServiceRepository struct {
	driver  string
	connStr string
}

func NewServiceRepository(host string, port int, user, password, dbname string) *ServiceRepository {
	return &ServiceRepository{
		driver:  "mysql",
		connStr: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbname),
	}
}

func (r *ServiceRepository) GetAll(ctx context.Context) ([]*api.Service, error) {
	return r.getMany(ctx, "select name,friendly_name,cast(active as unsigned) as active,sort,status,created,updated from service order by sort desc, name asc")
}

func (r *ServiceRepository) GetForEnvironment(ctx context.Context, environment string) ([]*api.Service, error) {
	sql := fmt.Sprintf(
		`select s.name,s.friendly_name,cast(s.active as unsigned) as active,s.sort,s.status,s.created,s.updated
		 from service s
		 inner join service_environment se on s.id = se.service_id
		 inner join environment e on se.environment_id = e.id
		 where e.friendly_name='%s'
		 order by s.sort desc, s.name asc`, environment)
	return r.getMany(ctx, sql)
}

func (r *ServiceRepository) Add(ctx context.Context, svc *api.Service) error {
	return r.Update(ctx, svc.Friendly, svc)
}

func (r *ServiceRepository) Update(ctx context.Context, friendly string, svc *api.Service) error {
	db, err := sqlx.Connect(r.driver, r.connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, "call service_upsert(?, ?, ?, ?, ?, ?);", friendly, svc.Name, svc.Friendly, svc.Status, svc.Active, svc.Sort)
	return err
}

func (r *ServiceRepository) getMany(ctx context.Context, sql string) ([]*api.Service, error) {
	db, err := sqlx.Connect(r.driver, r.connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	svcs := []*api.Service{}
	if err := db.SelectContext(ctx, &svcs, sql); err != nil {
		return nil, err
	}
	return svcs, nil
}
