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
)

// NewGetSystemGcScheduleParams creates a new GetSystemGcScheduleParams object
// with the default values initialized.
func NewGetSystemGcScheduleParams() *GetSystemGcScheduleParams {

	return &GetSystemGcScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemGcScheduleParamsWithTimeout creates a new GetSystemGcScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSystemGcScheduleParamsWithTimeout(timeout time.Duration) *GetSystemGcScheduleParams {

	return &GetSystemGcScheduleParams{

		timeout: timeout,
	}
}

// NewGetSystemGcScheduleParamsWithContext creates a new GetSystemGcScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSystemGcScheduleParamsWithContext(ctx context.Context) *GetSystemGcScheduleParams {

	return &GetSystemGcScheduleParams{

		Context: ctx,
	}
}

// NewGetSystemGcScheduleParamsWithHTTPClient creates a new GetSystemGcScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSystemGcScheduleParamsWithHTTPClient(client *http.Client) *GetSystemGcScheduleParams {

	return &GetSystemGcScheduleParams{
		HTTPClient: client,
	}
}

/*GetSystemGcScheduleParams contains all the parameters to send to the API endpoint
for the get system gc schedule operation typically these are written to a http.Request
*/
type GetSystemGcScheduleParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get system gc schedule params
func (o *GetSystemGcScheduleParams) WithTimeout(timeout time.Duration) *GetSystemGcScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system gc schedule params
func (o *GetSystemGcScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system gc schedule params
func (o *GetSystemGcScheduleParams) WithContext(ctx context.Context) *GetSystemGcScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system gc schedule params
func (o *GetSystemGcScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system gc schedule params
func (o *GetSystemGcScheduleParams) WithHTTPClient(client *http.Client) *GetSystemGcScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system gc schedule params
func (o *GetSystemGcScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemGcScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
