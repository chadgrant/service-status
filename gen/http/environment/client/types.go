// Code generated by goa v3.0.9, DO NOT EDIT.
//
// environment HTTP client types
//
// Command:
// $ goa gen servicestatus/design

package client

import (
	environment "servicestatus/gen/environment"
	environmentviews "servicestatus/gen/environment/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "environment" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Name of environment
	Name string `form:"name" json:"name" xml:"name"`
	// url friendly name of environment used for REST based operations
	Friendly string `form:"friendly" json:"friendly" xml:"friendly"`
	// is environment active
	Active bool `form:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// useful for sorting in UI
	Sort uint32 `form:"sort,omitempty" json:"sort,omitempty" xml:"sort,omitempty"`
	// used for update case of changing key, PUT /environment/key
	Key *string `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
}

// UpdateRequestBody is the type of the "environment" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// Name of environment
	Name string `form:"name" json:"name" xml:"name"`
	// url friendly name of environment used for REST based operations
	Friendly string `form:"friendly" json:"friendly" xml:"friendly"`
	// is environment active
	Active bool `form:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// useful for sorting in UI
	Sort uint32 `form:"sort,omitempty" json:"sort,omitempty" xml:"sort,omitempty"`
}

// ListResponseBody is the type of the "environment" service "list" endpoint
// HTTP response body.
type ListResponseBody []*EnvironmentResponse

// UpdateNotFoundResponseBody is the type of the "environment" service "update"
// endpoint HTTP response body for the "not_found" error.
type UpdateNotFoundResponseBody struct {
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// friendly name of environment
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// EnvironmentResponse is used to define fields on response body types.
type EnvironmentResponse struct {
	// Name of environment
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// url friendly name of environment used for REST based operations
	Friendly *string `form:"friendly,omitempty" json:"friendly,omitempty" xml:"friendly,omitempty"`
	// is environment active
	Active *bool `form:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// useful for sorting in UI
	Sort *uint32 `form:"sort,omitempty" json:"sort,omitempty" xml:"sort,omitempty"`
	// timestamp of when environment was created
	Created *string `form:"created,omitempty" json:"created,omitempty" xml:"created,omitempty"`
	// timestamp of when environment was updated
	Updated *string `form:"updated,omitempty" json:"updated,omitempty" xml:"updated,omitempty"`
}

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "environment" service.
func NewAddRequestBody(p *environment.EnvironmentBase) *AddRequestBody {
	body := &AddRequestBody{
		Name:     p.Name,
		Friendly: p.Friendly,
		Active:   p.Active,
		Sort:     p.Sort,
		Key:      p.Key,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "environment" service.
func NewUpdateRequestBody(p *environment.EnvironmentBase) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Name:     p.Name,
		Friendly: p.Friendly,
		Active:   p.Active,
		Sort:     p.Sort,
	}
	return body
}

// NewListEnvironmentCollectionOK builds a "environment" service "list"
// endpoint result from a HTTP "OK" response.
func NewListEnvironmentCollectionOK(body ListResponseBody) environmentviews.EnvironmentCollectionView {
	v := make([]*environmentviews.EnvironmentView, len(body))
	for i, val := range body {
		v[i] = &environmentviews.EnvironmentView{
			Name:     val.Name,
			Friendly: val.Friendly,
			Active:   val.Active,
			Sort:     val.Sort,
			Created:  val.Created,
			Updated:  val.Updated,
		}
		if val.Active == nil {
			var tmp bool = true
			v[i].Active = &tmp
		}
		if val.Sort == nil {
			var tmp uint32 = 0
			v[i].Sort = &tmp
		}
	}
	return v
}

// NewUpdateNotFound builds a environment service update endpoint not_found
// error.
func NewUpdateNotFound(body *UpdateNotFoundResponseBody) *environment.NotFound {
	v := &environment.NotFound{
		Message: *body.Message,
		Name:    *body.Name,
	}
	return v
}

// ValidateUpdateNotFoundResponseBody runs the validations defined on
// update_not_found_response_body
func ValidateUpdateNotFoundResponseBody(body *UpdateNotFoundResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

// ValidateEnvironmentResponse runs the validations defined on
// EnvironmentResponse
func ValidateEnvironmentResponse(body *EnvironmentResponse) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Friendly == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("friendly", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 150 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 150, false))
		}
	}
	if body.Friendly != nil {
		if utf8.RuneCountInString(*body.Friendly) > 150 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.friendly", *body.Friendly, utf8.RuneCountInString(*body.Friendly), 150, false))
		}
	}
	if body.Sort != nil {
		if *body.Sort < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.sort", *body.Sort, 0, true))
		}
	}
	if body.Sort != nil {
		if *body.Sort > 5000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.sort", *body.Sort, 5000, false))
		}
	}
	if body.Created != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.created", *body.Created, goa.FormatDateTime))
	}
	if body.Updated != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.updated", *body.Updated, goa.FormatDateTime))
	}
	return
}
