// Code generated by goa v3.0.9, DO NOT EDIT.
//
// environment gRPC client types
//
// Command:
// $ goa gen servicestatus/design

package client

import (
	environment "servicestatus/gen/environment"
	environmentviews "servicestatus/gen/environment/views"
	environmentpb "servicestatus/gen/grpc/environment/pb"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// NewListRequest builds the gRPC request type from the payload of the "list"
// endpoint of the "environment" service.
func NewListRequest() *environmentpb.ListRequest {
	message := &environmentpb.ListRequest{}
	return message
}

// NewListResult builds the result type of the "list" endpoint of the
// "environment" service from the gRPC response type.
func NewListResult(message *environmentpb.EnvironmentCollection) environmentviews.EnvironmentCollectionView {
	result := make([]*environmentviews.EnvironmentView, len(message.Field))
	for i, val := range message.Field {
		result[i] = &environmentviews.EnvironmentView{
			Name:     &val.Name,
			Friendly: &val.Friendly,
			Active:   &val.Active,
		}
		if val.Sort != 0 {
			result[i].Sort = &val.Sort
		}
		if val.Created != "" {
			result[i].Created = &val.Created
		}
		if val.Updated != "" {
			result[i].Updated = &val.Updated
		}
		if val.Sort == 0 {
			var tmp uint32 = 0
			result[i].Sort = &tmp
		}
	}
	return result
}

// NewAddRequest builds the gRPC request type from the payload of the "add"
// endpoint of the "environment" service.
func NewAddRequest(payload *environment.EnvironmentBase) *environmentpb.AddRequest {
	message := &environmentpb.AddRequest{
		Name:     payload.Name,
		Friendly: payload.Friendly,
		Active:   payload.Active,
		Sort:     payload.Sort,
	}
	if payload.Key != nil {
		message.Key = *payload.Key
	}
	return message
}

// NewAddResult builds the result type of the "add" endpoint of the
// "environment" service from the gRPC response type.
func NewAddResult(message *environmentpb.AddResponse) string {
	result := message.Field
	return result
}

// NewUpdateRequest builds the gRPC request type from the payload of the
// "update" endpoint of the "environment" service.
func NewUpdateRequest(payload *environment.EnvironmentBase) *environmentpb.UpdateRequest {
	message := &environmentpb.UpdateRequest{
		Name:     payload.Name,
		Friendly: payload.Friendly,
		Active:   payload.Active,
		Sort:     payload.Sort,
	}
	if payload.Key != nil {
		message.Key = *payload.Key
	}
	return message
}

// ValidateEnvironmentCollection runs the validations defined on
// EnvironmentCollection.
func ValidateEnvironmentCollection(message *environmentpb.EnvironmentCollection) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateEnvironment1(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateEnvironment1 runs the validations defined on Environment1.
func ValidateEnvironment1(message *environmentpb.Environment1) (err error) {
	if utf8.RuneCountInString(message.Name) > 150 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 150, false))
	}
	if utf8.RuneCountInString(message.Friendly) > 150 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.friendly", message.Friendly, utf8.RuneCountInString(message.Friendly), 150, false))
	}
	if message.Sort < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.sort", message.Sort, 0, true))
	}
	if message.Sort > 5000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.sort", message.Sort, 5000, false))
	}
	err = goa.MergeErrors(err, goa.ValidateFormat("message.created", message.Created, goa.FormatDateTime))

	err = goa.MergeErrors(err, goa.ValidateFormat("message.updated", message.Updated, goa.FormatDateTime))

	return
}
