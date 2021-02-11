// Code generated by go-swagger; DO NOT EDIT.

package legacy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReplicationTrigger replication trigger
//
// swagger:model ReplicationTrigger
type ReplicationTrigger struct {

	// trigger settings
	TriggerSettings *TriggerSettings `json:"trigger_settings,omitempty"`

	// The replication policy trigger type. The valid values are manual, event_based and scheduled.
	Type string `json:"type,omitempty"`
}

// Validate validates this replication trigger
func (m *ReplicationTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTriggerSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplicationTrigger) validateTriggerSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.TriggerSettings) { // not required
		return nil
	}

	if m.TriggerSettings != nil {
		if err := m.TriggerSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger_settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplicationTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicationTrigger) UnmarshalBinary(b []byte) error {
	var res ReplicationTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
