// Code generated by go-swagger; DO NOT EDIT.

package employee

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

// NewPostEmployeeListParams creates a new PostEmployeeListParams object
// with the default values initialized.
func NewPostEmployeeListParams() *PostEmployeeListParams {
	var ()
	return &PostEmployeeListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostEmployeeListParamsWithTimeout creates a new PostEmployeeListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostEmployeeListParamsWithTimeout(timeout time.Duration) *PostEmployeeListParams {
	var ()
	return &PostEmployeeListParams{

		timeout: timeout,
	}
}

// NewPostEmployeeListParamsWithContext creates a new PostEmployeeListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostEmployeeListParamsWithContext(ctx context.Context) *PostEmployeeListParams {
	var ()
	return &PostEmployeeListParams{

		Context: ctx,
	}
}

// NewPostEmployeeListParamsWithHTTPClient creates a new PostEmployeeListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostEmployeeListParamsWithHTTPClient(client *http.Client) *PostEmployeeListParams {
	var ()
	return &PostEmployeeListParams{
		HTTPClient: client,
	}
}

/*PostEmployeeListParams contains all the parameters to send to the API endpoint
for the post employee list operation typically these are written to a http.Request
*/
type PostEmployeeListParams struct {

	/*Body
	  JSON representing a list of new object to be created. Should not have ID and version set.

	*/
	Body []*models.Employee

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post employee list params
func (o *PostEmployeeListParams) WithTimeout(timeout time.Duration) *PostEmployeeListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post employee list params
func (o *PostEmployeeListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post employee list params
func (o *PostEmployeeListParams) WithContext(ctx context.Context) *PostEmployeeListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post employee list params
func (o *PostEmployeeListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post employee list params
func (o *PostEmployeeListParams) WithHTTPClient(client *http.Client) *PostEmployeeListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post employee list params
func (o *PostEmployeeListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post employee list params
func (o *PostEmployeeListParams) WithBody(body []*models.Employee) *PostEmployeeListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post employee list params
func (o *PostEmployeeListParams) SetBody(body []*models.Employee) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostEmployeeListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
