/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// VideoV1RecordingSettings struct for VideoV1RecordingSettings
type VideoV1RecordingSettings struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the stored Credential resource
	AwsCredentialsSid *string `json:"aws_credentials_sid,omitempty"`
	// The URL of the AWS S3 bucket where the recordings are stored
	AwsS3Url *string `json:"aws_s3_url,omitempty"`
	// Whether all recordings are written to the aws_s3_url
	AwsStorageEnabled *bool `json:"aws_storage_enabled,omitempty"`
	// Whether all recordings are stored in an encrypted form
	EncryptionEnabled *bool `json:"encryption_enabled,omitempty"`
	// The SID of the Public Key resource used for encryption
	EncryptionKeySid *string `json:"encryption_key_sid,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
