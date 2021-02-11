// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Repository repository
//
// swagger:model Repository
type Repository struct {

	// The count of the artifacts inside the repository
	ArtifactCount int64 `json:"artifact_count,omitempty"`

	// The creation time of the repository
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creation_time,omitempty"`

	// The description of the repository
	Description string `json:"description,omitempty"`

	// The ID of the repository
	ID int64 `json:"id,omitempty"`

	// The name of the repository
	Name string `json:"name,omitempty"`

	// The ID of the project that the repository belongs to
	ProjectID int64 `json:"project_id,omitempty"`

	// The count that the artifact inside the repository pulled
	PullCount int64 `json:"pull_count,omitempty"`

	// The update time of the repository
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this repository
func (m *Repository) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Repository) validateCreationTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_time", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Repository) validateUpdateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Repository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Repository) UnmarshalBinary(b []byte) error {
	var res Repository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
