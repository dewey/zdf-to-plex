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

// VideoV1RoomParticipantPublishedTrack struct for VideoV1RoomParticipantPublishedTrack
type VideoV1RoomParticipantPublishedTrack struct {
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Whether the track is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The track type
	Kind *string `json:"kind,omitempty"`
	// The track name
	Name *string `json:"name,omitempty"`
	// The SID of the Participant resource with the published track
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// The SID of the Room resource where the track is published
	RoomSid *string `json:"room_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
