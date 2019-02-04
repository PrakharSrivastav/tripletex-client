// Code generated by go-swagger; DO NOT EDIT.

package customer

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

// NewPostCustomerParams creates a new PostCustomerParams object
// with the default values initialized.
func NewPostCustomerParams() *PostCustomerParams {
	var ()
	return &PostCustomerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomerParamsWithTimeout creates a new PostCustomerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomerParamsWithTimeout(timeout time.Duration) *PostCustomerParams {
	var ()
	return &PostCustomerParams{

		timeout: timeout,
	}
}

// NewPostCustomerParamsWithContext creates a new PostCustomerParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomerParamsWithContext(ctx context.Context) *PostCustomerParams {
	var ()
	return &PostCustomerParams{

		Context: ctx,
	}
}

// NewPostCustomerParamsWithHTTPClient creates a new PostCustomerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomerParamsWithHTTPClient(client *http.Client) *PostCustomerParams {
	var ()
	return &PostCustomerParams{
		HTTPClient: client,
	}
}

/*PostCustomerParams contains all the parameters to send to the API endpoint
for the post customer operation typically these are written to a http.Request
*/
type PostCustomerParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.Customer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post customer params
func (o *PostCustomerParams) WithTimeout(timeout time.Duration) *PostCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customer params
func (o *PostCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customer params
func (o *PostCustomerParams) WithContext(ctx context.Context) *PostCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customer params
func (o *PostCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customer params
func (o *PostCustomerParams) WithHTTPClient(client *http.Client) *PostCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customer params
func (o *PostCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post customer params
func (o *PostCustomerParams) WithBody(body *models.Customer) *PostCustomerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post customer params
func (o *PostCustomerParams) SetBody(body *models.Customer) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
