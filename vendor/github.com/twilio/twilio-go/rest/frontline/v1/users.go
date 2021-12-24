/*
 * Twilio - Frontline
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

// Fetch a frontline user
func (c *ApiService) FetchUser(Sid string) (*FrontlineV1User, error) {
	path := "/v1/Users/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FrontlineV1User{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateUser'
type UpdateUserParams struct {
	// The avatar URL which will be shown in Frontline application.
	Avatar *string `json:"Avatar,omitempty"`
	// The string that you assigned to describe the User.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether the User is available for new conversations. Set to `false` to prevent User from receiving new inbound conversations if you are using [Pool Routing](https://www.twilio.com/docs/frontline/handle-incoming-conversations#3-pool-routing).
	IsAvailable *bool `json:"IsAvailable,omitempty"`
	// Current state of this user. Can be either `active` or `deactivated`.
	State *string `json:"State,omitempty"`
}

func (params *UpdateUserParams) SetAvatar(Avatar string) *UpdateUserParams {
	params.Avatar = &Avatar
	return params
}
func (params *UpdateUserParams) SetFriendlyName(FriendlyName string) *UpdateUserParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateUserParams) SetIsAvailable(IsAvailable bool) *UpdateUserParams {
	params.IsAvailable = &IsAvailable
	return params
}
func (params *UpdateUserParams) SetState(State string) *UpdateUserParams {
	params.State = &State
	return params
}

// Update an existing frontline user
func (c *ApiService) UpdateUser(Sid string, params *UpdateUserParams) (*FrontlineV1User, error) {
	path := "/v1/Users/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Avatar != nil {
		data.Set("Avatar", *params.Avatar)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IsAvailable != nil {
		data.Set("IsAvailable", fmt.Sprint(*params.IsAvailable))
	}
	if params != nil && params.State != nil {
		data.Set("State", *params.State)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FrontlineV1User{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}