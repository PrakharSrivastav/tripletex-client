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

	models "github.com/PrakharSrivastav/Petstore/models"
)

// NewPutProjectListParams creates a new PutProjectListParams object
// with the default values initialized.
func NewPutProjectListParams() *PutProjectListParams {
	var ()
	return &PutProjectListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutProjectListParamsWithTimeout creates a new PutProjectListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutProjectListParamsWithTimeout(timeout time.Duration) *PutProjectListParams {
	var ()
	return &PutProjectListParams{

		timeout: timeout,
	}
}

// NewPutProjectListParamsWithContext creates a new PutProjectListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutProjectListParamsWithContext(ctx context.Context) *PutProjectListParams {
	var ()
	return &PutProjectListParams{

		Context: ctx,
	}
}

// NewPutProjectListParamsWithHTTPClient creates a new PutProjectListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutProjectListParamsWithHTTPClient(client *http.Client) *PutProjectListParams {
	var ()
	return &PutProjectListParams{
		HTTPClient: client,
	}
}

/*PutProjectListParams contains all the parameters to send to the API endpoint
for the put project list operation typically these are written to a http.Request
*/
type PutProjectListParams struct {

	/*Body
	  JSON representing updates to object. Should have ID and version set.

	*/
	Body []*models.Project

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put project list params
func (o *PutProjectListParams) WithTimeout(timeout time.Duration) *PutProjectListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put project list params
func (o *PutProjectListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put project list params
func (o *PutProjectListParams) WithContext(ctx context.Context) *PutProjectListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put project list params
func (o *PutProjectListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put project list params
func (o *PutProjectListParams) WithHTTPClient(client *http.Client) *PutProjectListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put project list params
func (o *PutProjectListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put project list params
func (o *PutProjectListParams) WithBody(body []*models.Project) *PutProjectListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put project list params
func (o *PutProjectListParams) SetBody(body []*models.Project) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutProjectListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
