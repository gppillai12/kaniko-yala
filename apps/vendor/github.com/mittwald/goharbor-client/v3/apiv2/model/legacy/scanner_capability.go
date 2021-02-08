// Code generated by go-swagger; DO NOT EDIT.

package legacy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScannerCapability scanner capability
//
// swagger:model ScannerCapability
type ScannerCapability struct {

	// consumes mime types
	ConsumesMimeTypes []string `json:"consumes_mime_types"`

	// produces mime types
	ProducesMimeTypes []string `json:"produces_mime_types"`
}

// Validate validates this scanner capability
func (m *ScannerCapability) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScannerCapability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScannerCapability) UnmarshalBinary(b []byte) error {
	var res ScannerCapability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
