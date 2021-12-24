/*
 * Twilio - Video
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
	"time"

	"github.com/twilio/twilio-go/client"
)

func (c *ApiService) DeleteRoomRecording(RoomSid string, Sid string) error {
	path := "/v1/Rooms/{RoomSid}/Recordings/{Sid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchRoomRecording(RoomSid string, Sid string) (*VideoV1RoomRecording, error) {
	path := "/v1/Rooms/{RoomSid}/Recordings/{Sid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1RoomRecording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRoomRecording'
type ListRoomRecordingParams struct {
	// Read only the recordings with this status. Can be: `processing`, `completed`, or `deleted`.
	Status *string `json:"Status,omitempty"`
	// Read only the recordings that have this `source_sid`.
	SourceSid *string `json:"SourceSid,omitempty"`
	// Read only recordings that started on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// Read only Recordings that started before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone.
	DateCreatedBefore *time.Time `json:"DateCreatedBefore,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRoomRecordingParams) SetStatus(Status string) *ListRoomRecordingParams {
	params.Status = &Status
	return params
}
func (params *ListRoomRecordingParams) SetSourceSid(SourceSid string) *ListRoomRecordingParams {
	params.SourceSid = &SourceSid
	return params
}
func (params *ListRoomRecordingParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListRoomRecordingParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListRoomRecordingParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListRoomRecordingParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListRoomRecordingParams) SetPageSize(PageSize int) *ListRoomRecordingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRoomRecordingParams) SetLimit(Limit int) *ListRoomRecordingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of RoomRecording records from the API. Request is executed immediately.
func (c *ApiService) PageRoomRecording(RoomSid string, params *ListRoomRecordingParams, pageToken, pageNumber string) (*ListRoomRecordingResponse, error) {
	path := "/v1/Rooms/{RoomSid}/Recordings"

	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.SourceSid != nil {
		data.Set("SourceSid", *params.SourceSid)
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
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

	ps := &ListRoomRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists RoomRecording records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRoomRecording(RoomSid string, params *ListRoomRecordingParams) ([]VideoV1RoomRecording, error) {
	if params == nil {
		params = &ListRoomRecordingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRoomRecording(RoomSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VideoV1RoomRecording

	for response != nil {
		records = append(records, response.Recordings...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRoomRecordingResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListRoomRecordingResponse)
	}

	return records, err
}

// Streams RoomRecording records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRoomRecording(RoomSid string, params *ListRoomRecordingParams) (chan VideoV1RoomRecording, error) {
	if params == nil {
		params = &ListRoomRecordingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRoomRecording(RoomSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VideoV1RoomRecording, 1)

	go func() {
		for response != nil {
			for item := range response.Recordings {
				channel <- response.Recordings[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRoomRecordingResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListRoomRecordingResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListRoomRecordingResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoomRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}