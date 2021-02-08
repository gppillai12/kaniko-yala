// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetUsergroupsGroupIDParams creates a new GetUsergroupsGroupIDParams object
// with the default values initialized.
func NewGetUsergroupsGroupIDParams() *GetUsergroupsGroupIDParams {
	var ()
	return &GetUsergroupsGroupIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsergroupsGroupIDParamsWithTimeout creates a new GetUsergroupsGroupIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsergroupsGroupIDParamsWithTimeout(timeout time.Duration) *GetUsergroupsGroupIDParams {
	var ()
	return &GetUsergroupsGroupIDParams{

		timeout: timeout,
	}
}

// NewGetUsergroupsGroupIDParamsWithContext creates a new GetUsergroupsGroupIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsergroupsGroupIDParamsWithContext(ctx context.Context) *GetUsergroupsGroupIDParams {
	var ()
	return &GetUsergroupsGroupIDParams{

		Context: ctx,
	}
}

// NewGetUsergroupsGroupIDParamsWithHTTPClient creates a new GetUsergroupsGroupIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsergroupsGroupIDParamsWithHTTPClient(client *http.Client) *GetUsergroupsGroupIDParams {
	var ()
	return &GetUsergroupsGroupIDParams{
		HTTPClient: client,
	}
}

/*GetUsergroupsGroupIDParams contains all the parameters to send to the API endpoint
for the get usergroups group ID operation typically these are written to a http.Request
*/
type GetUsergroupsGroupIDParams struct {

	/*GroupID
	  Group ID

	*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get usergroups group ID params
func (o *GetUsergroupsGroupIDParams) WithTimeout(timeout time.Duration) *GetUsergroupsGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get usergroups group ID params
func (o *GetUsergroupsGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get usergroups group ID params
func (o *GetUsergroupsGroupIDParams) WithContext(ctx context.Context) *GetUsergroupsGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get usergroups group ID params
func (o *GetUsergroupsGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get usergroups group ID params
func (o *GetUsergroupsGroupIDParams) WithHTTPClient(client *http.Client) *GetUsergroupsGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get usergroups group ID params
func (o *GetUsergroupsGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the get usergroups group ID params
func (o *GetUsergroupsGroupIDParams) WithGroupID(groupID int64) *GetUsergroupsGroupIDParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get usergroups group ID params
func (o *GetUsergroupsGroupIDParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsergroupsGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param group_id
	if err := r.SetPathParam("group_id", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
