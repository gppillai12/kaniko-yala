// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
)

// GetProjectsProjectIDScannerReader is a Reader for the GetProjectsProjectIDScanner structure.
type GetProjectsProjectIDScannerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDScannerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDScannerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsProjectIDScannerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectsProjectIDScannerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectsProjectIDScannerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectsProjectIDScannerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDScannerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectsProjectIDScannerOK creates a GetProjectsProjectIDScannerOK with default headers values
func NewGetProjectsProjectIDScannerOK() *GetProjectsProjectIDScannerOK {
	return &GetProjectsProjectIDScannerOK{}
}

/*GetProjectsProjectIDScannerOK handles this case with default header values.

The details of the scanner registration.
*/
type GetProjectsProjectIDScannerOK struct {
	Payload *legacy.ScannerRegistration
}

func (o *GetProjectsProjectIDScannerOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/scanner][%d] getProjectsProjectIdScannerOK  %+v", 200, o.Payload)
}

func (o *GetProjectsProjectIDScannerOK) GetPayload() *legacy.ScannerRegistration {
	return o.Payload
}

func (o *GetProjectsProjectIDScannerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legacy.ScannerRegistration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDScannerBadRequest creates a GetProjectsProjectIDScannerBadRequest with default headers values
func NewGetProjectsProjectIDScannerBadRequest() *GetProjectsProjectIDScannerBadRequest {
	return &GetProjectsProjectIDScannerBadRequest{}
}

/*GetProjectsProjectIDScannerBadRequest handles this case with default header values.

Bad project ID
*/
type GetProjectsProjectIDScannerBadRequest struct {
}

func (o *GetProjectsProjectIDScannerBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/scanner][%d] getProjectsProjectIdScannerBadRequest ", 400)
}

func (o *GetProjectsProjectIDScannerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDScannerUnauthorized creates a GetProjectsProjectIDScannerUnauthorized with default headers values
func NewGetProjectsProjectIDScannerUnauthorized() *GetProjectsProjectIDScannerUnauthorized {
	return &GetProjectsProjectIDScannerUnauthorized{}
}

/*GetProjectsProjectIDScannerUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetProjectsProjectIDScannerUnauthorized struct {
}

func (o *GetProjectsProjectIDScannerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/scanner][%d] getProjectsProjectIdScannerUnauthorized ", 401)
}

func (o *GetProjectsProjectIDScannerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDScannerForbidden creates a GetProjectsProjectIDScannerForbidden with default headers values
func NewGetProjectsProjectIDScannerForbidden() *GetProjectsProjectIDScannerForbidden {
	return &GetProjectsProjectIDScannerForbidden{}
}

/*GetProjectsProjectIDScannerForbidden handles this case with default header values.

Request is not allowed
*/
type GetProjectsProjectIDScannerForbidden struct {
}

func (o *GetProjectsProjectIDScannerForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/scanner][%d] getProjectsProjectIdScannerForbidden ", 403)
}

func (o *GetProjectsProjectIDScannerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDScannerNotFound creates a GetProjectsProjectIDScannerNotFound with default headers values
func NewGetProjectsProjectIDScannerNotFound() *GetProjectsProjectIDScannerNotFound {
	return &GetProjectsProjectIDScannerNotFound{}
}

/*GetProjectsProjectIDScannerNotFound handles this case with default header values.

The requested object is not found
*/
type GetProjectsProjectIDScannerNotFound struct {
}

func (o *GetProjectsProjectIDScannerNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/scanner][%d] getProjectsProjectIdScannerNotFound ", 404)
}

func (o *GetProjectsProjectIDScannerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDScannerInternalServerError creates a GetProjectsProjectIDScannerInternalServerError with default headers values
func NewGetProjectsProjectIDScannerInternalServerError() *GetProjectsProjectIDScannerInternalServerError {
	return &GetProjectsProjectIDScannerInternalServerError{}
}

/*GetProjectsProjectIDScannerInternalServerError handles this case with default header values.

Internal server error happened
*/
type GetProjectsProjectIDScannerInternalServerError struct {
}

func (o *GetProjectsProjectIDScannerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/scanner][%d] getProjectsProjectIdScannerInternalServerError ", 500)
}

func (o *GetProjectsProjectIDScannerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
