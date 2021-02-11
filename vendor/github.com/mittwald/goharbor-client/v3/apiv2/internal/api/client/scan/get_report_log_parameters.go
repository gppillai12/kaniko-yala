// Code generated by go-swagger; DO NOT EDIT.

package scan

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

// NewGetReportLogParams creates a new GetReportLogParams object
// with the default values initialized.
func NewGetReportLogParams() *GetReportLogParams {
	var ()
	return &GetReportLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReportLogParamsWithTimeout creates a new GetReportLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReportLogParamsWithTimeout(timeout time.Duration) *GetReportLogParams {
	var ()
	return &GetReportLogParams{

		timeout: timeout,
	}
}

// NewGetReportLogParamsWithContext creates a new GetReportLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReportLogParamsWithContext(ctx context.Context) *GetReportLogParams {
	var ()
	return &GetReportLogParams{

		Context: ctx,
	}
}

// NewGetReportLogParamsWithHTTPClient creates a new GetReportLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReportLogParamsWithHTTPClient(client *http.Client) *GetReportLogParams {
	var ()
	return &GetReportLogParams{
		HTTPClient: client,
	}
}

/*GetReportLogParams contains all the parameters to send to the API endpoint
for the get report log operation typically these are written to a http.Request
*/
type GetReportLogParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*ProjectName
	  The name of the project

	*/
	ProjectName string
	/*Reference
	  The reference of the artifact, can be digest or tag

	*/
	Reference string
	/*ReportID
	  The report id to get the log

	*/
	ReportID string
	/*RepositoryName
	  The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -> a%252Fb

	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get report log params
func (o *GetReportLogParams) WithTimeout(timeout time.Duration) *GetReportLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get report log params
func (o *GetReportLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get report log params
func (o *GetReportLogParams) WithContext(ctx context.Context) *GetReportLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get report log params
func (o *GetReportLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get report log params
func (o *GetReportLogParams) WithHTTPClient(client *http.Client) *GetReportLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get report log params
func (o *GetReportLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get report log params
func (o *GetReportLogParams) WithXRequestID(xRequestID *string) *GetReportLogParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get report log params
func (o *GetReportLogParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectName adds the projectName to the get report log params
func (o *GetReportLogParams) WithProjectName(projectName string) *GetReportLogParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the get report log params
func (o *GetReportLogParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithReference adds the reference to the get report log params
func (o *GetReportLogParams) WithReference(reference string) *GetReportLogParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the get report log params
func (o *GetReportLogParams) SetReference(reference string) {
	o.Reference = reference
}

// WithReportID adds the reportID to the get report log params
func (o *GetReportLogParams) WithReportID(reportID string) *GetReportLogParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the get report log params
func (o *GetReportLogParams) SetReportID(reportID string) {
	o.ReportID = reportID
}

// WithRepositoryName adds the repositoryName to the get report log params
func (o *GetReportLogParams) WithRepositoryName(repositoryName string) *GetReportLogParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get report log params
func (o *GetReportLogParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	// path param reference
	if err := r.SetPathParam("reference", o.Reference); err != nil {
		return err
	}

	// path param report_id
	if err := r.SetPathParam("report_id", o.ReportID); err != nil {
		return err
	}

	// path param repository_name
	if err := r.SetPathParam("repository_name", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
