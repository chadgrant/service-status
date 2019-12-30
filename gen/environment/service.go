// Code generated by goa v3.0.9, DO NOT EDIT.
//
// environment service
//
// Command:
// $ goa gen servicestatus/design

package environment

import (
	"context"
	environmentviews "servicestatus/gen/environment/views"
)

// The environment service performs operations on environments
type Service interface {
	// List all environments
	List(context.Context) (res EnvironmentCollection, err error)
	// adds a new environment and returns the url of where to retrieve it
	Add(context.Context, *EnvironmentBase) (res string, err error)
	// update an existing environment
	Update(context.Context, *EnvironmentBase) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "environment"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"list", "add", "update"}

// EnvironmentCollection is the result type of the environment service list
// method.
type EnvironmentCollection []*Environment

// EnvironmentBase is the payload type of the environment service add method.
type EnvironmentBase struct {
	// Name of environment
	Name string
	// url friendly name of environment used for REST based operations
	Friendly string
	// is environment active
	Active bool
	// useful for sorting in UI
	Sort uint32
	// used for update case of changing key, PUT /environment/key
	Key *string
}

type Environment struct {
	// Name of environment
	Name string
	// url friendly name of environment used for REST based operations
	Friendly string
	// is environment active
	Active bool
	// useful for sorting in UI
	Sort uint32
	// timestamp of when environment was created
	Created *string
	// timestamp of when environment was updated
	Updated *string
}

// NotFound when an environment does not exist.
type NotFound struct {
	// Message of error
	Message string
	// friendly name of environment
	Name string
}

// Error returns an error description.
func (e *NotFound) Error() string {
	return "NotFound when an environment does not exist."
}

// ErrorName returns "NotFound".
func (e *NotFound) ErrorName() string {
	return e.Message
}

// NewEnvironmentCollection initializes result type EnvironmentCollection from
// viewed result type EnvironmentCollection.
func NewEnvironmentCollection(vres environmentviews.EnvironmentCollection) EnvironmentCollection {
	var res EnvironmentCollection
	switch vres.View {
	case "default", "":
		res = newEnvironmentCollection(vres.Projected)
	}
	return res
}

// NewViewedEnvironmentCollection initializes viewed result type
// EnvironmentCollection from result type EnvironmentCollection using the given
// view.
func NewViewedEnvironmentCollection(res EnvironmentCollection, view string) environmentviews.EnvironmentCollection {
	var vres environmentviews.EnvironmentCollection
	switch view {
	case "default", "":
		p := newEnvironmentCollectionView(res)
		vres = environmentviews.EnvironmentCollection{Projected: p, View: "default"}
	}
	return vres
}

// newEnvironmentCollection converts projected type EnvironmentCollection to
// service type EnvironmentCollection.
func newEnvironmentCollection(vres environmentviews.EnvironmentCollectionView) EnvironmentCollection {
	res := make(EnvironmentCollection, len(vres))
	for i, n := range vres {
		res[i] = newEnvironment(n)
	}
	return res
}

// newEnvironmentCollectionView projects result type EnvironmentCollection to
// projected type EnvironmentCollectionView using the "default" view.
func newEnvironmentCollectionView(res EnvironmentCollection) environmentviews.EnvironmentCollectionView {
	vres := make(environmentviews.EnvironmentCollectionView, len(res))
	for i, n := range res {
		vres[i] = newEnvironmentView(n)
	}
	return vres
}

// newEnvironment converts projected type Environment to service type
// Environment.
func newEnvironment(vres *environmentviews.EnvironmentView) *Environment {
	res := &Environment{
		Created: vres.Created,
		Updated: vres.Updated,
	}
	if vres.Friendly != nil {
		res.Friendly = *vres.Friendly
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Active != nil {
		res.Active = *vres.Active
	}
	if vres.Sort != nil {
		res.Sort = *vres.Sort
	}
	if vres.Active == nil {
		res.Active = true
	}
	if vres.Sort == nil {
		res.Sort = 0
	}
	return res
}

// newEnvironmentView projects result type Environment to projected type
// EnvironmentView using the "default" view.
func newEnvironmentView(res *Environment) *environmentviews.EnvironmentView {
	vres := &environmentviews.EnvironmentView{
		Name:     &res.Name,
		Friendly: &res.Friendly,
		Active:   &res.Active,
		Sort:     &res.Sort,
		Created:  res.Created,
		Updated:  res.Updated,
	}
	return vres
}