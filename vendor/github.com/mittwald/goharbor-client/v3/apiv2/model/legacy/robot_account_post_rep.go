// Code generated by go-swagger; DO NOT EDIT.

package legacy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RobotAccountPostRep robot account post rep
//
// swagger:model RobotAccountPostRep
type RobotAccountPostRep struct {

	// the name of robot account
	Name string `json:"name,omitempty"`

	// the token of robot account
	Token string `json:"token,omitempty"`
}

// Validate validates this robot account post rep
func (m *RobotAccountPostRep) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RobotAccountPostRep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RobotAccountPostRep) UnmarshalBinary(b []byte) error {
	var res RobotAccountPostRep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
