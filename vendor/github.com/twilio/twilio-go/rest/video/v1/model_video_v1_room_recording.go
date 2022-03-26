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

import (
	"time"
)

// VideoV1RoomRecording struct for VideoV1RoomRecording
type VideoV1RoomRecording struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The codec used for the recording
	Codec *string `json:"codec,omitempty"`
	// The file format for the recording
	ContainerFormat *string `json:"container_format,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The duration of the recording in seconds
	Duration *int `json:"duration,omitempty"`
	// A list of SIDs related to the Recording
	GroupingSids *map[string]interface{} `json:"grouping_sids,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The URL of the media file associated with the recording when stored externally
	MediaExternalLocation *string `json:"media_external_location,omitempty"`
	// The number of milliseconds between a point in time that is common to all rooms in a group and when the source room of the recording started
	Offset *int `json:"offset,omitempty"`
	// The SID of the Room resource the recording is associated with
	RoomSid *string `json:"room_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The size of the recorded track in bytes
	Size *int `json:"size,omitempty"`
	// The SID of the recording source
	SourceSid *string `json:"source_sid,omitempty"`
	// The status of the recording
	Status *string `json:"status,omitempty"`
	// The name that was given to the source track of the recording
	TrackName *string `json:"track_name,omitempty"`
	// The recording's media type
	Type *string `json:"type,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
