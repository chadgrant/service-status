// Code generated by goa v3.0.9, DO NOT EDIT.
//
// environment gRPC client CLI support package
//
// Command:
// $ goa gen servicestatus/design

package client

import (
	"encoding/json"
	"fmt"
	environment "servicestatus/gen/environment"
	environmentpb "servicestatus/gen/grpc/environment/pb"
)

// BuildAddPayload builds the payload for the environment add endpoint from CLI
// flags.
func BuildAddPayload(environmentAddMessage string) (*environment.EnvironmentBase, error) {
	var err error
	var message environmentpb.AddRequest
	{
		if environmentAddMessage != "" {
			err = json.Unmarshal([]byte(environmentAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"active\": true,\n      \"friendly\": \"development\",\n      \"key\": \"Pariatur ut.\",\n      \"name\": \"Production\",\n      \"sort\": 0\n   }'")
			}
		}
	}
	v := &environment.EnvironmentBase{
		Name:     message.Name,
		Friendly: message.Friendly,
		Active:   message.Active,
		Sort:     message.Sort,
	}
	if message.Key != "" {
		v.Key = &message.Key
	}
	if message.Sort == 0 {
		v.Sort = 0
	}
	return v, nil
}

// BuildUpdatePayload builds the payload for the environment update endpoint
// from CLI flags.
func BuildUpdatePayload(environmentUpdateMessage string) (*environment.EnvironmentBase, error) {
	var err error
	var message environmentpb.UpdateRequest
	{
		if environmentUpdateMessage != "" {
			err = json.Unmarshal([]byte(environmentUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"active\": true,\n      \"friendly\": \"development\",\n      \"key\": \"In inventore ut.\",\n      \"name\": \"Production\",\n      \"sort\": 0\n   }'")
			}
		}
	}
	v := &environment.EnvironmentBase{
		Name:     message.Name,
		Friendly: message.Friendly,
		Active:   message.Active,
		Sort:     message.Sort,
	}
	if message.Key != "" {
		v.Key = &message.Key
	}
	if message.Sort == 0 {
		v.Sort = 0
	}
	return v, nil
}