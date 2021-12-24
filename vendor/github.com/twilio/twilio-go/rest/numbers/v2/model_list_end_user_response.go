/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEndUserResponse struct for ListEndUserResponse
type ListEndUserResponse struct {
	Meta    ListBundleResponseMeta `json:"meta,omitempty"`
	Results []NumbersV2EndUser     `json:"results,omitempty"`
}
