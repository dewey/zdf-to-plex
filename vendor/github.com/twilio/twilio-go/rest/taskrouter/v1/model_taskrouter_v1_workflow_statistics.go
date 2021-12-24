/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskrouterV1WorkflowStatistics struct for TaskrouterV1WorkflowStatistics
type TaskrouterV1WorkflowStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// An object that contains the cumulative statistics for the Workflow
	Cumulative *map[string]interface{} `json:"cumulative,omitempty"`
	// An object that contains the real-time statistics for the Workflow
	Realtime *map[string]interface{} `json:"realtime,omitempty"`
	// The absolute URL of the Workflow statistics resource
	Url *string `json:"url,omitempty"`
	// Returns the list of Tasks that are being controlled by the Workflow with the specified SID value
	WorkflowSid *string `json:"workflow_sid,omitempty"`
	// The SID of the Workspace that contains the Workflow
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
