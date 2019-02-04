// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteProjectListParams creates a new DeleteProjectListParams object
// with the default values initialized.
func NewDeleteProjectListParams() *DeleteProjectListParams {
	var ()
	return &DeleteProjectListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectListParamsWithTimeout creates a new DeleteProjectListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProjectListParamsWithTimeout(timeout time.Duration) *DeleteProjectListParams {
	var ()
	return &DeleteProjectListParams{

		timeout: timeout,
	}
}

// NewDeleteProjectListParamsWithContext creates a new DeleteProjectListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProjectListParamsWithContext(ctx context.Context) *DeleteProjectListParams {
	var ()
	return &DeleteProjectListParams{

		Context: ctx,
	}
}

// NewDeleteProjectListParamsWithHTTPClient creates a new DeleteProjectListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProjectListParamsWithHTTPClient(client *http.Client) *DeleteProjectListParams {
	var ()
	return &DeleteProjectListParams{
		HTTPClient: client,
	}
}

/*DeleteProjectListParams contains all the parameters to send to the API endpoint
for the delete project list operation typically these are written to a http.Request
*/
type DeleteProjectListParams struct {

	/*Ids
	  ID of the elements

	*/
	Ids string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete project list params
func (o *DeleteProjectListParams) WithTimeout(timeout time.Duration) *DeleteProjectListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete project list params
func (o *DeleteProjectListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete project list params
func (o *DeleteProjectListParams) WithContext(ctx context.Context) *DeleteProjectListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete project list params
func (o *DeleteProjectListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete project list params
func (o *DeleteProjectListParams) WithHTTPClient(client *http.Client) *DeleteProjectListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete project list params
func (o *DeleteProjectListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the delete project list params
func (o *DeleteProjectListParams) WithIds(ids string) *DeleteProjectListParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the delete project list params
func (o *DeleteProjectListParams) SetIds(ids string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param ids
	qrIds := o.Ids
	qIds := qrIds
	if qIds != "" {
		if err := r.SetQueryParam("ids", qIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
