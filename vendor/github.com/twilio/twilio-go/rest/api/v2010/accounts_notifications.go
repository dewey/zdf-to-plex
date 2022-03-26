/*
 * Twilio - Api
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

// Optional parameters for the method 'FetchNotification'
type FetchNotificationParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Notification resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchNotificationParams) SetPathAccountSid(PathAccountSid string) *FetchNotificationParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a notification belonging to the account used to make the request
func (c *ApiService) FetchNotification(Sid string, params *FetchNotificationParams) (*ApiV2010NotificationInstance, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Notifications/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010NotificationInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListNotification'
type ListNotificationParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Notification resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Only read notifications of the specified log level. Can be:  `0` to read only ERROR notifications or `1` to read only WARNING notifications. By default, all notifications are read.
	Log *int `json:"Log,omitempty"`
	// Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date.
	MessageDate *string `json:"MessageDate,omitempty"`
	// Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date.
	MessageDateBefore *string `json:"MessageDate&lt;,omitempty"`
	// Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date.
	MessageDateAfter *string `json:"MessageDate&gt;,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListNotificationParams) SetPathAccountSid(PathAccountSid string) *ListNotificationParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListNotificationParams) SetLog(Log int) *ListNotificationParams {
	params.Log = &Log
	return params
}
func (params *ListNotificationParams) SetMessageDate(MessageDate string) *ListNotificationParams {
	params.MessageDate = &MessageDate
	return params
}
func (params *ListNotificationParams) SetMessageDateBefore(MessageDateBefore string) *ListNotificationParams {
	params.MessageDateBefore = &MessageDateBefore
	return params
}
func (params *ListNotificationParams) SetMessageDateAfter(MessageDateAfter string) *ListNotificationParams {
	params.MessageDateAfter = &MessageDateAfter
	return params
}
func (params *ListNotificationParams) SetPageSize(PageSize int) *ListNotificationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListNotificationParams) SetLimit(Limit int) *ListNotificationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Notification records from the API. Request is executed immediately.
func (c *ApiService) PageNotification(params *ListNotificationParams, pageToken, pageNumber string) (*ListNotificationResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Notifications.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Log != nil {
		data.Set("Log", fmt.Sprint(*params.Log))
	}
	if params != nil && params.MessageDate != nil {
		data.Set("MessageDate", fmt.Sprint(*params.MessageDate))
	}
	if params != nil && params.MessageDateBefore != nil {
		data.Set("MessageDate<", fmt.Sprint(*params.MessageDateBefore))
	}
	if params != nil && params.MessageDateAfter != nil {
		data.Set("MessageDate>", fmt.Sprint(*params.MessageDateAfter))
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

	ps := &ListNotificationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Notification records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListNotification(params *ListNotificationParams) ([]ApiV2010Notification, error) {
	response, err := c.StreamNotification(params)
	if err != nil {
		return nil, err
	}

	records := make([]ApiV2010Notification, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Notification records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamNotification(params *ListNotificationParams) (chan ApiV2010Notification, error) {
	if params == nil {
		params = &ListNotificationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageNotification(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010Notification, 1)

	go func() {
		for response != nil {
			responseRecords := response.Notifications
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListNotificationResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListNotificationResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListNotificationResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListNotificationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
