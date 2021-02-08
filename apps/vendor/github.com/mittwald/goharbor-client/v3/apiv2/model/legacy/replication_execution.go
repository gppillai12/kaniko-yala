// Code generated by go-swagger; DO NOT EDIT.

package legacy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReplicationExecution The replication execution
//
// swagger:model ReplicationExecution
type ReplicationExecution struct {

	// The end time
	EndTime string `json:"end_time,omitempty"`

	// The count of failed tasks
	Failed int64 `json:"failed,omitempty"`

	// The ID
	ID int64 `json:"id,omitempty"`

	// The count of in_progress tasks
	InProgress int64 `json:"in_progress,omitempty"`

	// The policy ID
	PolicyID int64 `json:"policy_id,omitempty"`

	// The start time
	StartTime string `json:"start_time,omitempty"`

	// The status
	Status string `json:"status,omitempty"`

	// The status text
	StatusText string `json:"status_text,omitempty"`

	// The count of stopped tasks
	Stopped int64 `json:"stopped,omitempty"`

	// The count of succeed tasks
	Succeed int64 `json:"succeed,omitempty"`

	// The total count of all tasks
	Total int64 `json:"total,omitempty"`

	// The trigger mode
	Trigger string `json:"trigger,omitempty"`
}

// Validate validates this replication execution
func (m *ReplicationExecution) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReplicationExecution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicationExecution) UnmarshalBinary(b []byte) error {
	var res ReplicationExecution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
