/*
 * Twilio - Taskrouter
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

// Optional parameters for the method 'FetchWorkersRealTimeStatistics'
type FetchWorkersRealTimeStatisticsParams struct {
	// Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
	TaskChannel *string `json:"TaskChannel,omitempty"`
}

func (params *FetchWorkersRealTimeStatisticsParams) SetTaskChannel(TaskChannel string) *FetchWorkersRealTimeStatisticsParams {
	params.TaskChannel = &TaskChannel
	return params
}

func (c *ApiService) FetchWorkersRealTimeStatistics(WorkspaceSid string, params *FetchWorkersRealTimeStatisticsParams) (*TaskrouterV1WorkersRealTimeStatistics, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/RealTimeStatistics"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.TaskChannel != nil {
		data.Set("TaskChannel", *params.TaskChannel)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkersRealTimeStatistics{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
