/*
 * Twilio - Accounts
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListCredentialPublicKeyResponse struct for ListCredentialPublicKeyResponse
type ListCredentialPublicKeyResponse struct {
	Credentials []AccountsV1CredentialPublicKey `json:"credentials,omitempty"`
	Meta        ListCredentialAwsResponseMeta   `json:"meta,omitempty"`
}
