/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// VerifyV2AccessToken struct for VerifyV2AccessToken
type VerifyV2AccessToken struct {
	// Account Sid.
	AccountSid *string `json:"account_sid,omitempty"`
	// The date this access token was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// Unique external identifier of the Entity
	EntityIdentity *string `json:"entity_identity,omitempty"`
	// A human readable description of this factor.
	FactorFriendlyName *string `json:"factor_friendly_name,omitempty"`
	// The Type of the Factor
	FactorType *string `json:"factor_type,omitempty"`
	// Verify Service Sid.
	ServiceSid *string `json:"service_sid,omitempty"`
	// A string that uniquely identifies this Access Token.
	Sid *string `json:"sid,omitempty"`
	// Generated access token.
	Token *string `json:"token,omitempty"`
	// How long, in seconds, the access token is valid.
	Ttl *int `json:"ttl,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
