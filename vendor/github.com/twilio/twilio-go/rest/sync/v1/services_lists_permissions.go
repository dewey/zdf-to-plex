/*
 * Twilio - Sync
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
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Delete a specific Sync List Permission.
func (c *ApiService) DeleteSyncListPermission(ServiceSid string, ListSid string, Identity string) error {
	path := "/v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ListSid"+"}", ListSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific Sync List Permission.
func (c *ApiService) FetchSyncListPermission(ServiceSid string, ListSid string, Identity string) (*SyncV1SyncListPermission, error) {
	path := "/v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ListSid"+"}", ListSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncListPermission{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSyncListPermission'
type ListSyncListPermissionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSyncListPermissionParams) SetPageSize(PageSize int) *ListSyncListPermissionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSyncListPermissionParams) SetLimit(Limit int) *ListSyncListPermissionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SyncListPermission records from the API. Request is executed immediately.
func (c *ApiService) PageSyncListPermission(ServiceSid string, ListSid string, params *ListSyncListPermissionParams, pageToken, pageNumber string) (*ListSyncListPermissionResponse, error) {
	path := "/v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ListSid"+"}", ListSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncListPermissionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SyncListPermission records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSyncListPermission(ServiceSid string, ListSid string, params *ListSyncListPermissionParams) ([]SyncV1SyncListPermission, error) {
	response, err := c.StreamSyncListPermission(ServiceSid, ListSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]SyncV1SyncListPermission, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams SyncListPermission records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSyncListPermission(ServiceSid string, ListSid string, params *ListSyncListPermissionParams) (chan SyncV1SyncListPermission, error) {
	if params == nil {
		params = &ListSyncListPermissionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSyncListPermission(ServiceSid, ListSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan SyncV1SyncListPermission, 1)

	go func() {
		for response != nil {
			responseRecords := response.Permissions
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListSyncListPermissionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSyncListPermissionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSyncListPermissionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncListPermissionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSyncListPermission'
type UpdateSyncListPermissionParams struct {
	// Whether the identity can delete the Sync List. Default value is `false`.
	Manage *bool `json:"Manage,omitempty"`
	// Whether the identity can read the Sync List and its Items. Default value is `false`.
	Read *bool `json:"Read,omitempty"`
	// Whether the identity can create, update, and delete Items in the Sync List. Default value is `false`.
	Write *bool `json:"Write,omitempty"`
}

func (params *UpdateSyncListPermissionParams) SetManage(Manage bool) *UpdateSyncListPermissionParams {
	params.Manage = &Manage
	return params
}
func (params *UpdateSyncListPermissionParams) SetRead(Read bool) *UpdateSyncListPermissionParams {
	params.Read = &Read
	return params
}
func (params *UpdateSyncListPermissionParams) SetWrite(Write bool) *UpdateSyncListPermissionParams {
	params.Write = &Write
	return params
}

// Update an identity&#39;s access to a specific Sync List.
func (c *ApiService) UpdateSyncListPermission(ServiceSid string, ListSid string, Identity string, params *UpdateSyncListPermissionParams) (*SyncV1SyncListPermission, error) {
	path := "/v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ListSid"+"}", ListSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Manage != nil {
		data.Set("Manage", fmt.Sprint(*params.Manage))
	}
	if params != nil && params.Read != nil {
		data.Set("Read", fmt.Sprint(*params.Read))
	}
	if params != nil && params.Write != nil {
		data.Set("Write", fmt.Sprint(*params.Write))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncListPermission{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
