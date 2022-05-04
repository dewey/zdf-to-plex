/*
 * Twilio - Insights
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

// InsightsV1ConferenceParticipant struct for InsightsV1ConferenceParticipant
type InsightsV1ConferenceParticipant struct {
	// Account SID.
	AccountSid *string `json:"account_sid,omitempty"`
	// Call direction of the participant.
	CallDirection *string `json:"call_direction,omitempty"`
	// Unique SID identifier of the call.
	CallSid *string `json:"call_sid,omitempty"`
	// Call status of the call that generated the participant.
	CallStatus *string `json:"call_status,omitempty"`
	// The Call Type of this Conference Participant.
	CallType *string `json:"call_type,omitempty"`
	// Call SIDs coached by this participant.
	CoachedParticipants *[]string `json:"coached_participants,omitempty"`
	// The Conference Region of this Conference Participant.
	ConferenceRegion *string `json:"conference_region,omitempty"`
	// Conference SID.
	ConferenceSid *string `json:"conference_sid,omitempty"`
	// ISO alpha-2 country code of the participant.
	CountryCode *string `json:"country_code,omitempty"`
	// Participant durations in seconds.
	DurationSeconds *int `json:"duration_seconds,omitempty"`
	// Object containing information of actions taken by participants. Nested resource URLs.
	Events *map[string]interface{} `json:"events,omitempty"`
	// Caller ID of the calling party.
	From *string `json:"from,omitempty"`
	// Boolean. Indicated whether participant was a coach.
	IsCoach *bool `json:"is_coach,omitempty"`
	// Boolean. Indicates whether participant had startConferenceOnEnter=true or endConferenceOnExit=true.
	IsModerator *bool `json:"is_moderator,omitempty"`
	// The Jitter Buffer Size of this Conference Participant.
	JitterBufferSize *string `json:"jitter_buffer_size,omitempty"`
	// ISO 8601 timestamp of participant join event.
	JoinTime *time.Time `json:"join_time,omitempty"`
	// The user-specified label of this participant.
	Label *string `json:"label,omitempty"`
	// ISO 8601 timestamp of participant leave event.
	LeaveTime *time.Time `json:"leave_time,omitempty"`
	// Object. Contains participant quality metrics.
	Metrics *map[string]interface{} `json:"metrics,omitempty"`
	// Estimated time in queue at call creation.
	OutboundQueueLength *int `json:"outbound_queue_length,omitempty"`
	// Actual time in queue (seconds).
	OutboundTimeInQueue *int `json:"outbound_time_in_queue,omitempty"`
	// Twilio region where the participant media originates.
	ParticipantRegion *string `json:"participant_region,omitempty"`
	// SID for this participant.
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// Processing state of the Participant Summary.
	ProcessingState *string `json:"processing_state,omitempty"`
	// Participant properties and metadata.
	Properties *map[string]interface{} `json:"properties,omitempty"`
	// Called party.
	To *string `json:"to,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}