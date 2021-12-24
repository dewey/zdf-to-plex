/*
 * Twilio - Api
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

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateApplication'
type CreateApplicationParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`. The default value is the account's default API version.
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// A descriptive string that you create to describe the new application. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The URL we should call using a POST method to send message status information to your application.
	MessageStatusCallback *string `json:"MessageStatusCallback,omitempty"`
	// The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`.
	SmsFallbackMethod *string `json:"SmsFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs while retrieving or executing the TwiML from `sms_url`.
	SmsFallbackUrl *string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`.
	SmsMethod *string `json:"SmsMethod,omitempty"`
	// The URL we should call using a POST method to send status information about SMS messages sent by the application.
	SmsStatusCallback *string `json:"SmsStatusCallback,omitempty"`
	// The URL we should call when the phone number receives an incoming SMS message.
	SmsUrl *string `json:"SmsUrl,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// Whether we should look up the caller's caller-ID name from the CNAM database (additional charges apply). Can be: `true` or `false`.
	VoiceCallerIdLookup *bool `json:"VoiceCallerIdLookup,omitempty"`
	// The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	VoiceFallbackMethod *string `json:"VoiceFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`.
	VoiceFallbackUrl *string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`.
	VoiceMethod *string `json:"VoiceMethod,omitempty"`
	// The URL we should call when the phone number assigned to this application receives a call.
	VoiceUrl *string `json:"VoiceUrl,omitempty"`
}

func (params *CreateApplicationParams) SetPathAccountSid(PathAccountSid string) *CreateApplicationParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateApplicationParams) SetApiVersion(ApiVersion string) *CreateApplicationParams {
	params.ApiVersion = &ApiVersion
	return params
}
func (params *CreateApplicationParams) SetFriendlyName(FriendlyName string) *CreateApplicationParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateApplicationParams) SetMessageStatusCallback(MessageStatusCallback string) *CreateApplicationParams {
	params.MessageStatusCallback = &MessageStatusCallback
	return params
}
func (params *CreateApplicationParams) SetSmsFallbackMethod(SmsFallbackMethod string) *CreateApplicationParams {
	params.SmsFallbackMethod = &SmsFallbackMethod
	return params
}
func (params *CreateApplicationParams) SetSmsFallbackUrl(SmsFallbackUrl string) *CreateApplicationParams {
	params.SmsFallbackUrl = &SmsFallbackUrl
	return params
}
func (params *CreateApplicationParams) SetSmsMethod(SmsMethod string) *CreateApplicationParams {
	params.SmsMethod = &SmsMethod
	return params
}
func (params *CreateApplicationParams) SetSmsStatusCallback(SmsStatusCallback string) *CreateApplicationParams {
	params.SmsStatusCallback = &SmsStatusCallback
	return params
}
func (params *CreateApplicationParams) SetSmsUrl(SmsUrl string) *CreateApplicationParams {
	params.SmsUrl = &SmsUrl
	return params
}
func (params *CreateApplicationParams) SetStatusCallback(StatusCallback string) *CreateApplicationParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateApplicationParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateApplicationParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *CreateApplicationParams) SetVoiceCallerIdLookup(VoiceCallerIdLookup bool) *CreateApplicationParams {
	params.VoiceCallerIdLookup = &VoiceCallerIdLookup
	return params
}
func (params *CreateApplicationParams) SetVoiceFallbackMethod(VoiceFallbackMethod string) *CreateApplicationParams {
	params.VoiceFallbackMethod = &VoiceFallbackMethod
	return params
}
func (params *CreateApplicationParams) SetVoiceFallbackUrl(VoiceFallbackUrl string) *CreateApplicationParams {
	params.VoiceFallbackUrl = &VoiceFallbackUrl
	return params
}
func (params *CreateApplicationParams) SetVoiceMethod(VoiceMethod string) *CreateApplicationParams {
	params.VoiceMethod = &VoiceMethod
	return params
}
func (params *CreateApplicationParams) SetVoiceUrl(VoiceUrl string) *CreateApplicationParams {
	params.VoiceUrl = &VoiceUrl
	return params
}

// Create a new application within your account
func (c *ApiService) CreateApplication(params *CreateApplicationParams) (*ApiV2010Application, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Applications.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ApiVersion != nil {
		data.Set("ApiVersion", *params.ApiVersion)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.MessageStatusCallback != nil {
		data.Set("MessageStatusCallback", *params.MessageStatusCallback)
	}
	if params != nil && params.SmsFallbackMethod != nil {
		data.Set("SmsFallbackMethod", *params.SmsFallbackMethod)
	}
	if params != nil && params.SmsFallbackUrl != nil {
		data.Set("SmsFallbackUrl", *params.SmsFallbackUrl)
	}
	if params != nil && params.SmsMethod != nil {
		data.Set("SmsMethod", *params.SmsMethod)
	}
	if params != nil && params.SmsStatusCallback != nil {
		data.Set("SmsStatusCallback", *params.SmsStatusCallback)
	}
	if params != nil && params.SmsUrl != nil {
		data.Set("SmsUrl", *params.SmsUrl)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.VoiceCallerIdLookup != nil {
		data.Set("VoiceCallerIdLookup", fmt.Sprint(*params.VoiceCallerIdLookup))
	}
	if params != nil && params.VoiceFallbackMethod != nil {
		data.Set("VoiceFallbackMethod", *params.VoiceFallbackMethod)
	}
	if params != nil && params.VoiceFallbackUrl != nil {
		data.Set("VoiceFallbackUrl", *params.VoiceFallbackUrl)
	}
	if params != nil && params.VoiceMethod != nil {
		data.Set("VoiceMethod", *params.VoiceMethod)
	}
	if params != nil && params.VoiceUrl != nil {
		data.Set("VoiceUrl", *params.VoiceUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Application{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteApplication'
type DeleteApplicationParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteApplicationParams) SetPathAccountSid(PathAccountSid string) *DeleteApplicationParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete the application by the specified application sid
func (c *ApiService) DeleteApplication(Sid string, params *DeleteApplicationParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Applications/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
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

// Optional parameters for the method 'FetchApplication'
type FetchApplicationParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchApplicationParams) SetPathAccountSid(PathAccountSid string) *FetchApplicationParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch the application specified by the provided sid
func (c *ApiService) FetchApplication(Sid string, params *FetchApplicationParams) (*ApiV2010Application, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Applications/{Sid}.json"
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

	ps := &ApiV2010Application{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListApplication'
type ListApplicationParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The string that identifies the Application resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListApplicationParams) SetPathAccountSid(PathAccountSid string) *ListApplicationParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListApplicationParams) SetFriendlyName(FriendlyName string) *ListApplicationParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListApplicationParams) SetPageSize(PageSize int) *ListApplicationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListApplicationParams) SetLimit(Limit int) *ListApplicationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Application records from the API. Request is executed immediately.
func (c *ApiService) PageApplication(params *ListApplicationParams, pageToken, pageNumber string) (*ListApplicationResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Applications.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
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

	ps := &ListApplicationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Application records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListApplication(params *ListApplicationParams) ([]ApiV2010Application, error) {
	if params == nil {
		params = &ListApplicationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageApplication(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010Application

	for response != nil {
		records = append(records, response.Applications...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListApplicationResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListApplicationResponse)
	}

	return records, err
}

// Streams Application records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamApplication(params *ListApplicationParams) (chan ApiV2010Application, error) {
	if params == nil {
		params = &ListApplicationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageApplication(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010Application, 1)

	go func() {
		for response != nil {
			for item := range response.Applications {
				channel <- response.Applications[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListApplicationResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListApplicationResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListApplicationResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListApplicationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateApplication'
type UpdateApplicationParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`. The default value is your account's default API version.
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The URL we should call using a POST method to send message status information to your application.
	MessageStatusCallback *string `json:"MessageStatusCallback,omitempty"`
	// The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`.
	SmsFallbackMethod *string `json:"SmsFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs while retrieving or executing the TwiML from `sms_url`.
	SmsFallbackUrl *string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`.
	SmsMethod *string `json:"SmsMethod,omitempty"`
	// Same as message_status_callback: The URL we should call using a POST method to send status information about SMS messages sent by the application. Deprecated, included for backwards compatibility.
	SmsStatusCallback *string `json:"SmsStatusCallback,omitempty"`
	// The URL we should call when the phone number receives an incoming SMS message.
	SmsUrl *string `json:"SmsUrl,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// Whether we should look up the caller's caller-ID name from the CNAM database (additional charges apply). Can be: `true` or `false`.
	VoiceCallerIdLookup *bool `json:"VoiceCallerIdLookup,omitempty"`
	// The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	VoiceFallbackMethod *string `json:"VoiceFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`.
	VoiceFallbackUrl *string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`.
	VoiceMethod *string `json:"VoiceMethod,omitempty"`
	// The URL we should call when the phone number assigned to this application receives a call.
	VoiceUrl *string `json:"VoiceUrl,omitempty"`
}

func (params *UpdateApplicationParams) SetPathAccountSid(PathAccountSid string) *UpdateApplicationParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateApplicationParams) SetApiVersion(ApiVersion string) *UpdateApplicationParams {
	params.ApiVersion = &ApiVersion
	return params
}
func (params *UpdateApplicationParams) SetFriendlyName(FriendlyName string) *UpdateApplicationParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateApplicationParams) SetMessageStatusCallback(MessageStatusCallback string) *UpdateApplicationParams {
	params.MessageStatusCallback = &MessageStatusCallback
	return params
}
func (params *UpdateApplicationParams) SetSmsFallbackMethod(SmsFallbackMethod string) *UpdateApplicationParams {
	params.SmsFallbackMethod = &SmsFallbackMethod
	return params
}
func (params *UpdateApplicationParams) SetSmsFallbackUrl(SmsFallbackUrl string) *UpdateApplicationParams {
	params.SmsFallbackUrl = &SmsFallbackUrl
	return params
}
func (params *UpdateApplicationParams) SetSmsMethod(SmsMethod string) *UpdateApplicationParams {
	params.SmsMethod = &SmsMethod
	return params
}
func (params *UpdateApplicationParams) SetSmsStatusCallback(SmsStatusCallback string) *UpdateApplicationParams {
	params.SmsStatusCallback = &SmsStatusCallback
	return params
}
func (params *UpdateApplicationParams) SetSmsUrl(SmsUrl string) *UpdateApplicationParams {
	params.SmsUrl = &SmsUrl
	return params
}
func (params *UpdateApplicationParams) SetStatusCallback(StatusCallback string) *UpdateApplicationParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *UpdateApplicationParams) SetStatusCallbackMethod(StatusCallbackMethod string) *UpdateApplicationParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *UpdateApplicationParams) SetVoiceCallerIdLookup(VoiceCallerIdLookup bool) *UpdateApplicationParams {
	params.VoiceCallerIdLookup = &VoiceCallerIdLookup
	return params
}
func (params *UpdateApplicationParams) SetVoiceFallbackMethod(VoiceFallbackMethod string) *UpdateApplicationParams {
	params.VoiceFallbackMethod = &VoiceFallbackMethod
	return params
}
func (params *UpdateApplicationParams) SetVoiceFallbackUrl(VoiceFallbackUrl string) *UpdateApplicationParams {
	params.VoiceFallbackUrl = &VoiceFallbackUrl
	return params
}
func (params *UpdateApplicationParams) SetVoiceMethod(VoiceMethod string) *UpdateApplicationParams {
	params.VoiceMethod = &VoiceMethod
	return params
}
func (params *UpdateApplicationParams) SetVoiceUrl(VoiceUrl string) *UpdateApplicationParams {
	params.VoiceUrl = &VoiceUrl
	return params
}

// Updates the application&#39;s properties
func (c *ApiService) UpdateApplication(Sid string, params *UpdateApplicationParams) (*ApiV2010Application, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Applications/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ApiVersion != nil {
		data.Set("ApiVersion", *params.ApiVersion)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.MessageStatusCallback != nil {
		data.Set("MessageStatusCallback", *params.MessageStatusCallback)
	}
	if params != nil && params.SmsFallbackMethod != nil {
		data.Set("SmsFallbackMethod", *params.SmsFallbackMethod)
	}
	if params != nil && params.SmsFallbackUrl != nil {
		data.Set("SmsFallbackUrl", *params.SmsFallbackUrl)
	}
	if params != nil && params.SmsMethod != nil {
		data.Set("SmsMethod", *params.SmsMethod)
	}
	if params != nil && params.SmsStatusCallback != nil {
		data.Set("SmsStatusCallback", *params.SmsStatusCallback)
	}
	if params != nil && params.SmsUrl != nil {
		data.Set("SmsUrl", *params.SmsUrl)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.VoiceCallerIdLookup != nil {
		data.Set("VoiceCallerIdLookup", fmt.Sprint(*params.VoiceCallerIdLookup))
	}
	if params != nil && params.VoiceFallbackMethod != nil {
		data.Set("VoiceFallbackMethod", *params.VoiceFallbackMethod)
	}
	if params != nil && params.VoiceFallbackUrl != nil {
		data.Set("VoiceFallbackUrl", *params.VoiceFallbackUrl)
	}
	if params != nil && params.VoiceMethod != nil {
		data.Set("VoiceMethod", *params.VoiceMethod)
	}
	if params != nil && params.VoiceUrl != nil {
		data.Set("VoiceUrl", *params.VoiceUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Application{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}