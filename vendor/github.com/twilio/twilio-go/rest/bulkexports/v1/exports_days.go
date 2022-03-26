/*
 * Twilio - Bulkexports
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

// Fetch a specific Day.
func (c *ApiService) FetchDay(ResourceType string, Day string) error {
	path := "/v1/Exports/{ResourceType}/Days/{Day}"
	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)
	path = strings.Replace(path, "{"+"Day"+"}", Day, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'ListDay'
type ListDayParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListDayParams) SetPageSize(PageSize int) *ListDayParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListDayParams) SetLimit(Limit int) *ListDayParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Day records from the API. Request is executed immediately.
func (c *ApiService) PageDay(ResourceType string, params *ListDayParams, pageToken, pageNumber string) (*ListDayResponse, error) {
	path := "/v1/Exports/{ResourceType}/Days"

	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

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

	ps := &ListDayResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Day records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDay(ResourceType string, params *ListDayParams) ([]BulkexportsV1Day, error) {
	response, err := c.StreamDay(ResourceType, params)
	if err != nil {
		return nil, err
	}

	records := make([]BulkexportsV1Day, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Day records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDay(ResourceType string, params *ListDayParams) (chan BulkexportsV1Day, error) {
	if params == nil {
		params = &ListDayParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageDay(ResourceType, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan BulkexportsV1Day, 1)

	go func() {
		for response != nil {
			responseRecords := response.Days
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListDayResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListDayResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListDayResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDayResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
