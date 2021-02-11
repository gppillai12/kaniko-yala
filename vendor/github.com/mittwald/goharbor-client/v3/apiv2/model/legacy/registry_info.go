// Code generated by go-swagger; DO NOT EDIT.

package legacy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RegistryInfo The registry info contains the base info and capability declarations of the registry
//
// swagger:model RegistryInfo
type RegistryInfo struct {

	// The description
	Description string `json:"description,omitempty"`

	// The filters that the registry supports
	SupportedResourceFilters []*FilterStyle `json:"supported_resource_filters"`

	// The triggers that the registry supports
	SupportedTriggers []string `json:"supported_triggers"`

	// The registry type
	Type string `json:"type,omitempty"`
}

// Validate validates this registry info
func (m *RegistryInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSupportedResourceFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistryInfo) validateSupportedResourceFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportedResourceFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.SupportedResourceFilters); i++ {
		if swag.IsZero(m.SupportedResourceFilters[i]) { // not required
			continue
		}

		if m.SupportedResourceFilters[i] != nil {
			if err := m.SupportedResourceFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supported_resource_filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistryInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistryInfo) UnmarshalBinary(b []byte) error {
	var res RegistryInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
