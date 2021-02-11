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

// NewGetReplicationExecutionsIDTasksTaskIDLogParams creates a new GetReplicationExecutionsIDTasksTaskIDLogParams object
// with the default values initialized.
func NewGetReplicationExecutionsIDTasksTaskIDLogParams() *GetReplicationExecutionsIDTasksTaskIDLogParams {
	var ()
	return &GetReplicationExecutionsIDTasksTaskIDLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReplicationExecutionsIDTasksTaskIDLogParamsWithTimeout creates a new GetReplicationExecutionsIDTasksTaskIDLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReplicationExecutionsIDTasksTaskIDLogParamsWithTimeout(timeout time.Duration) *GetReplicationExecutionsIDTasksTaskIDLogParams {
	var ()
	return &GetReplicationExecutionsIDTasksTaskIDLogParams{

		timeout: timeout,
	}
}

// NewGetReplicationExecutionsIDTasksTaskIDLogParamsWithContext creates a new GetReplicationExecutionsIDTasksTaskIDLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReplicationExecutionsIDTasksTaskIDLogParamsWithContext(ctx context.Context) *GetReplicationExecutionsIDTasksTaskIDLogParams {
	var ()
	return &GetReplicationExecutionsIDTasksTaskIDLogParams{

		Context: ctx,
	}
}

// NewGetReplicationExecutionsIDTasksTaskIDLogParamsWithHTTPClient creates a new GetReplicationExecutionsIDTasksTaskIDLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReplicationExecutionsIDTasksTaskIDLogParamsWithHTTPClient(client *http.Client) *GetReplicationExecutionsIDTasksTaskIDLogParams {
	var ()
	return &GetReplicationExecutionsIDTasksTaskIDLogParams{
		HTTPClient: client,
	}
}

/*GetReplicationExecutionsIDTasksTaskIDLogParams contains all the parameters to send to the API endpoint
for the get replication executions ID tasks task ID log operation typically these are written to a http.Request
*/
type GetReplicationExecutionsIDTasksTaskIDLogParams struct {

	/*ID
	  The execution ID.

	*/
	ID int64
	/*TaskID
	  The task ID.

	*/
	TaskID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get replication executions ID tasks task ID log params
func (o *GetReplicationExecutionsIDTasksTaskIDLogParams) WithTimeout(timeout time.Duration) *GetReplicationExecutionsIDTasksTaskIDLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get replication executions ID tasks task ID log params
func (o *GetReplicationExecutionsIDTasksTaskIDLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get replication executions ID tasks task ID log params
func (o *GetReplicationExecutionsIDTasksTaskIDLogParams) WithContext(ctx context.Context) *GetReplicationExecutionsIDTasksTaskIDLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get replication executions ID tasks task ID log params
func (o *GetReplicationExecutionsIDTasksTaskIDLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get replication executions ID tasks task ID log params
func (o *GetReplicationExecutionsIDTasksTaskIDLogParams) WithHTTPClient(client *http.Client) *GetReplicationExecutionsIDTasksTaskIDLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get replication executions ID tasks task ID log params
func (o *GetReplicationExecutionsIDTasksTaskIDLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get replication executions ID tasks task ID log params
func (o *GetReplicationExecutionsIDTasksTaskIDLogParams) WithID(id int64) *GetReplicationExecutionsIDTasksTaskIDLogParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get replication executions ID tasks task ID log params
func (o *GetReplicationExecutionsIDTasksTaskIDLogParams) SetID(id int64) {
	o.ID = id
}

// WithTaskID adds the taskID to the get replication executions ID tasks task ID log params
func (o *GetReplicationExecutionsIDTasksTaskIDLogParams) WithTaskID(taskID int64) *GetReplicationExecutionsIDTasksTaskIDLogParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the get replication executions ID tasks task ID log params
func (o *GetReplicationExecutionsIDTasksTaskIDLogParams) SetTaskID(taskID int64) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetReplicationExecutionsIDTasksTaskIDLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param task_id
	if err := r.SetPathParam("task_id", swag.FormatInt64(o.TaskID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
