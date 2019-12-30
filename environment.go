package servicestatus

import (
	"context"
	"log"
	environment "servicestatus/gen/environment"
	"strings"
	"time"
)

// environment service example implementation.
// The example methods log the requests and return zero values.
type environmentsrvc struct {
	logger       *log.Logger
	environments environment.EnvironmentCollection
}

// NewEnvironment returns the environment service implementation.
func NewEnvironment(logger *log.Logger) environment.Service {
	s := &environmentsrvc{
		logger:       logger,
		environments: make([]*environment.Environment, 0),
	}

	t := time.Now().UTC().Format(time.RFC3339)

	s.environments = append(s.environments, &environment.Environment{
		Name:     "Development",
		Friendly: "dev",
		Active:   true,
		Sort:     1,
		Created:  &t,
	})

	return s
}

// List all environments
func (s *environmentsrvc) List(ctx context.Context) (res environment.EnvironmentCollection, err error) {
	s.logger.Print("environment.list")
	return s.environments, nil
}

// adds a new environment and returns the url of where to retrieve it
func (s *environmentsrvc) Add(ctx context.Context, p *environment.EnvironmentBase) (res string, err error) {
	s.logger.Print("environment.add")

	t := time.Now().UTC().Format(time.RFC3339)

	s.environments = append(s.environments, &environment.Environment{
		Name:     p.Name,
		Friendly: p.Friendly,
		Active:   p.Active,
		Sort:     p.Sort,
		Created:  &t,
	})

	return p.Friendly, nil
}

// updated an existing environment
func (s *environmentsrvc) Update(ctx context.Context, p *environment.EnvironmentBase) (err error) {
	s.logger.Print("environment.update")

	t := time.Now().UTC().Format(time.RFC3339)

	e := s.Find(coalesce(p.Key, p.Friendly))
	if e == nil {
		return &environment.NotFound{
			Message: "environment not found",
			Name:    coalesce(p.Key, p.Friendly),
		}
	}
	e.Friendly = p.Friendly
	e.Name = p.Name
	e.Sort = p.Sort
	e.Active = p.Active
	e.Updated = &t

	return nil
}

func (s *environmentsrvc) Find(friendly string) *environment.Environment {
	for _, v := range s.environments {
		if strings.EqualFold(v.Friendly, friendly) {
			return v
		}
	}
	return nil
}

// func (s *environmentsrvc) Replace(friendly string, e *environment.Environment) bool {
// 	for i,v := range s.environments {
// 		if strings.EqualFold(v.Friendly,friendly) {
// 			s.environments[i] = e
// 			return true
// 		}
// 	}
// 	return false
// }

func coalesce(strs ...interface{}) string {
	for _, s := range strs {
		switch v := s.(type) {
		case string:
			if len(v) > 0 {
				return v
			}
		case *string:
			if v != nil && len(*v) > 0 {
				return *v
			}
		}
	}
	return ""
}
