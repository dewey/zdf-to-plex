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

// ApiV2010CallEvent struct for ApiV2010CallEvent
type ApiV2010CallEvent struct {
	// Call Request.
	Request *map[string]interface{} `json:"request,omitempty"`
	// Call Response with Events.
	Response *map[string]interface{} `json:"response,omitempty"`
}
