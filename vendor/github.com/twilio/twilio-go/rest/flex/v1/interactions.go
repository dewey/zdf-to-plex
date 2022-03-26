/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Optional parameters for the method 'CreateInteraction'
type CreateInteractionParams struct {
	// The Interaction's channel.
	Channel *map[string]interface{} `json:"Channel,omitempty"`
	// The Interaction's routing logic.
	Routing *map[string]interface{} `json:"Routing,omitempty"`
}

func (params *CreateInteractionParams) SetChannel(Channel map[string]interface{}) *CreateInteractionParams {
	params.Channel = &Channel
	return params
}
func (params *CreateInteractionParams) SetRouting(Routing map[string]interface{}) *CreateInteractionParams {
	params.Routing = &Routing
	return params
}

// Create a new Interaction.
func (c *ApiService) CreateInteraction(params *CreateInteractionParams) (*FlexV1Interaction, error) {
	path := "/v1/Interactions"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Channel != nil {
		v, err := json.Marshal(params.Channel)

		if err != nil {
			return nil, err
		}

		data.Set("Channel", string(v))
	}
	if params != nil && params.Routing != nil {
		v, err := json.Marshal(params.Routing)

		if err != nil {
			return nil, err
		}

		data.Set("Routing", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Interaction{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) FetchInteraction(Sid string) (*FlexV1Interaction, error) {
	path := "/v1/Interactions/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Interaction{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
