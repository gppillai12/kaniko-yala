// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRegistriesIDReader is a Reader for the PutRegistriesID structure.
type PutRegistriesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRegistriesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutRegistriesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRegistriesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRegistriesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutRegistriesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutRegistriesIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutRegistriesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutRegistriesIDOK creates a PutRegistriesIDOK with default headers values
func NewPutRegistriesIDOK() *PutRegistriesIDOK {
	return &PutRegistriesIDOK{}
}

/*PutRegistriesIDOK handles this case with default header values.

Updated registry successfully.
*/
type PutRegistriesIDOK struct {
}

func (o *PutRegistriesIDOK) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] putRegistriesIdOK ", 200)
}

func (o *PutRegistriesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRegistriesIDBadRequest creates a PutRegistriesIDBadRequest with default headers values
func NewPutRegistriesIDBadRequest() *PutRegistriesIDBadRequest {
	return &PutRegistriesIDBadRequest{}
}

/*PutRegistriesIDBadRequest handles this case with default header values.

The registry is associated with policy which is enabled.
*/
type PutRegistriesIDBadRequest struct {
}

func (o *PutRegistriesIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] putRegistriesIdBadRequest ", 400)
}

func (o *PutRegistriesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRegistriesIDUnauthorized creates a PutRegistriesIDUnauthorized with default headers values
func NewPutRegistriesIDUnauthorized() *PutRegistriesIDUnauthorized {
	return &PutRegistriesIDUnauthorized{}
}

/*PutRegistriesIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type PutRegistriesIDUnauthorized struct {
}

func (o *PutRegistriesIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] putRegistriesIdUnauthorized ", 401)
}

func (o *PutRegistriesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRegistriesIDNotFound creates a PutRegistriesIDNotFound with default headers values
func NewPutRegistriesIDNotFound() *PutRegistriesIDNotFound {
	return &PutRegistriesIDNotFound{}
}

/*PutRegistriesIDNotFound handles this case with default header values.

Registry does not exist.
*/
type PutRegistriesIDNotFound struct {
}

func (o *PutRegistriesIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] putRegistriesIdNotFound ", 404)
}

func (o *PutRegistriesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRegistriesIDConflict creates a PutRegistriesIDConflict with default headers values
func NewPutRegistriesIDConflict() *PutRegistriesIDConflict {
	return &PutRegistriesIDConflict{}
}

/*PutRegistriesIDConflict handles this case with default header values.

Registry name is already used.
*/
type PutRegistriesIDConflict struct {
}

func (o *PutRegistriesIDConflict) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] putRegistriesIdConflict ", 409)
}

func (o *PutRegistriesIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRegistriesIDInternalServerError creates a PutRegistriesIDInternalServerError with default headers values
func NewPutRegistriesIDInternalServerError() *PutRegistriesIDInternalServerError {
	return &PutRegistriesIDInternalServerError{}
}

/*PutRegistriesIDInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PutRegistriesIDInternalServerError struct {
}

func (o *PutRegistriesIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] putRegistriesIdInternalServerError ", 500)
}

func (o *PutRegistriesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
