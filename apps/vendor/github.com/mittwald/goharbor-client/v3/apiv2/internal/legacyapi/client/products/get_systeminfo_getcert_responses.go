// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetSysteminfoGetcertReader is a Reader for the GetSysteminfoGetcert structure.
type GetSysteminfoGetcertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSysteminfoGetcertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSysteminfoGetcertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSysteminfoGetcertNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSysteminfoGetcertInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSysteminfoGetcertOK creates a GetSysteminfoGetcertOK with default headers values
func NewGetSysteminfoGetcertOK() *GetSysteminfoGetcertOK {
	return &GetSysteminfoGetcertOK{}
}

/*GetSysteminfoGetcertOK handles this case with default header values.

Get default root certificate successfully.
*/
type GetSysteminfoGetcertOK struct {
}

func (o *GetSysteminfoGetcertOK) Error() string {
	return fmt.Sprintf("[GET /systeminfo/getcert][%d] getSysteminfoGetcertOK ", 200)
}

func (o *GetSysteminfoGetcertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSysteminfoGetcertNotFound creates a GetSysteminfoGetcertNotFound with default headers values
func NewGetSysteminfoGetcertNotFound() *GetSysteminfoGetcertNotFound {
	return &GetSysteminfoGetcertNotFound{}
}

/*GetSysteminfoGetcertNotFound handles this case with default header values.

Not found the default root certificate.
*/
type GetSysteminfoGetcertNotFound struct {
}

func (o *GetSysteminfoGetcertNotFound) Error() string {
	return fmt.Sprintf("[GET /systeminfo/getcert][%d] getSysteminfoGetcertNotFound ", 404)
}

func (o *GetSysteminfoGetcertNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSysteminfoGetcertInternalServerError creates a GetSysteminfoGetcertInternalServerError with default headers values
func NewGetSysteminfoGetcertInternalServerError() *GetSysteminfoGetcertInternalServerError {
	return &GetSysteminfoGetcertInternalServerError{}
}

/*GetSysteminfoGetcertInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetSysteminfoGetcertInternalServerError struct {
}

func (o *GetSysteminfoGetcertInternalServerError) Error() string {
	return fmt.Sprintf("[GET /systeminfo/getcert][%d] getSysteminfoGetcertInternalServerError ", 500)
}

func (o *GetSysteminfoGetcertInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
