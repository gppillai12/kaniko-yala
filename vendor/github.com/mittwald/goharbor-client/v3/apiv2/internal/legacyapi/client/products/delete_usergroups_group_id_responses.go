// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteUsergroupsGroupIDReader is a Reader for the DeleteUsergroupsGroupID structure.
type DeleteUsergroupsGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsergroupsGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUsergroupsGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUsergroupsGroupIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUsergroupsGroupIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUsergroupsGroupIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUsergroupsGroupIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUsergroupsGroupIDOK creates a DeleteUsergroupsGroupIDOK with default headers values
func NewDeleteUsergroupsGroupIDOK() *DeleteUsergroupsGroupIDOK {
	return &DeleteUsergroupsGroupIDOK{}
}

/*DeleteUsergroupsGroupIDOK handles this case with default header values.

User group deleted successfully.
*/
type DeleteUsergroupsGroupIDOK struct {
}

func (o *DeleteUsergroupsGroupIDOK) Error() string {
	return fmt.Sprintf("[DELETE /usergroups/{group_id}][%d] deleteUsergroupsGroupIdOK ", 200)
}

func (o *DeleteUsergroupsGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsergroupsGroupIDBadRequest creates a DeleteUsergroupsGroupIDBadRequest with default headers values
func NewDeleteUsergroupsGroupIDBadRequest() *DeleteUsergroupsGroupIDBadRequest {
	return &DeleteUsergroupsGroupIDBadRequest{}
}

/*DeleteUsergroupsGroupIDBadRequest handles this case with default header values.

The user group id is invalid.
*/
type DeleteUsergroupsGroupIDBadRequest struct {
}

func (o *DeleteUsergroupsGroupIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /usergroups/{group_id}][%d] deleteUsergroupsGroupIdBadRequest ", 400)
}

func (o *DeleteUsergroupsGroupIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsergroupsGroupIDUnauthorized creates a DeleteUsergroupsGroupIDUnauthorized with default headers values
func NewDeleteUsergroupsGroupIDUnauthorized() *DeleteUsergroupsGroupIDUnauthorized {
	return &DeleteUsergroupsGroupIDUnauthorized{}
}

/*DeleteUsergroupsGroupIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type DeleteUsergroupsGroupIDUnauthorized struct {
}

func (o *DeleteUsergroupsGroupIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /usergroups/{group_id}][%d] deleteUsergroupsGroupIdUnauthorized ", 401)
}

func (o *DeleteUsergroupsGroupIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsergroupsGroupIDForbidden creates a DeleteUsergroupsGroupIDForbidden with default headers values
func NewDeleteUsergroupsGroupIDForbidden() *DeleteUsergroupsGroupIDForbidden {
	return &DeleteUsergroupsGroupIDForbidden{}
}

/*DeleteUsergroupsGroupIDForbidden handles this case with default header values.

Only admin has this authority.
*/
type DeleteUsergroupsGroupIDForbidden struct {
}

func (o *DeleteUsergroupsGroupIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /usergroups/{group_id}][%d] deleteUsergroupsGroupIdForbidden ", 403)
}

func (o *DeleteUsergroupsGroupIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsergroupsGroupIDInternalServerError creates a DeleteUsergroupsGroupIDInternalServerError with default headers values
func NewDeleteUsergroupsGroupIDInternalServerError() *DeleteUsergroupsGroupIDInternalServerError {
	return &DeleteUsergroupsGroupIDInternalServerError{}
}

/*DeleteUsergroupsGroupIDInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type DeleteUsergroupsGroupIDInternalServerError struct {
}

func (o *DeleteUsergroupsGroupIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /usergroups/{group_id}][%d] deleteUsergroupsGroupIdInternalServerError ", 500)
}

func (o *DeleteUsergroupsGroupIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
