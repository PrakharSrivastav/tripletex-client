// Code generated by go-swagger; DO NOT EDIT.

package division

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

// NewPutDivisionListParams creates a new PutDivisionListParams object
// with the default values initialized.
func NewPutDivisionListParams() *PutDivisionListParams {
	var ()
	return &PutDivisionListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDivisionListParamsWithTimeout creates a new PutDivisionListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDivisionListParamsWithTimeout(timeout time.Duration) *PutDivisionListParams {
	var ()
	return &PutDivisionListParams{

		timeout: timeout,
	}
}

// NewPutDivisionListParamsWithContext creates a new PutDivisionListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDivisionListParamsWithContext(ctx context.Context) *PutDivisionListParams {
	var ()
	return &PutDivisionListParams{

		Context: ctx,
	}
}

// NewPutDivisionListParamsWithHTTPClient creates a new PutDivisionListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDivisionListParamsWithHTTPClient(client *http.Client) *PutDivisionListParams {
	var ()
	return &PutDivisionListParams{
		HTTPClient: client,
	}
}

/*PutDivisionListParams contains all the parameters to send to the API endpoint
for the put division list operation typically these are written to a http.Request
*/
type PutDivisionListParams struct {

	/*Body
	  JSON representing updates to object. Should have ID and version set.

	*/
	Body []*models.Division

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put division list params
func (o *PutDivisionListParams) WithTimeout(timeout time.Duration) *PutDivisionListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put division list params
func (o *PutDivisionListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put division list params
func (o *PutDivisionListParams) WithContext(ctx context.Context) *PutDivisionListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put division list params
func (o *PutDivisionListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put division list params
func (o *PutDivisionListParams) WithHTTPClient(client *http.Client) *PutDivisionListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put division list params
func (o *PutDivisionListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put division list params
func (o *PutDivisionListParams) WithBody(body []*models.Division) *PutDivisionListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put division list params
func (o *PutDivisionListParams) SetBody(body []*models.Division) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutDivisionListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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