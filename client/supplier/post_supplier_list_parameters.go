// Code generated by go-swagger; DO NOT EDIT.

package supplier

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

// NewPostSupplierListParams creates a new PostSupplierListParams object
// with the default values initialized.
func NewPostSupplierListParams() *PostSupplierListParams {
	var ()
	return &PostSupplierListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSupplierListParamsWithTimeout creates a new PostSupplierListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSupplierListParamsWithTimeout(timeout time.Duration) *PostSupplierListParams {
	var ()
	return &PostSupplierListParams{

		timeout: timeout,
	}
}

// NewPostSupplierListParamsWithContext creates a new PostSupplierListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSupplierListParamsWithContext(ctx context.Context) *PostSupplierListParams {
	var ()
	return &PostSupplierListParams{

		Context: ctx,
	}
}

// NewPostSupplierListParamsWithHTTPClient creates a new PostSupplierListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSupplierListParamsWithHTTPClient(client *http.Client) *PostSupplierListParams {
	var ()
	return &PostSupplierListParams{
		HTTPClient: client,
	}
}

/*PostSupplierListParams contains all the parameters to send to the API endpoint
for the post supplier list operation typically these are written to a http.Request
*/
type PostSupplierListParams struct {

	/*Body
	  JSON representing a list of new object to be created. Should not have ID and version set.

	*/
	Body []*models.Supplier

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post supplier list params
func (o *PostSupplierListParams) WithTimeout(timeout time.Duration) *PostSupplierListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post supplier list params
func (o *PostSupplierListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post supplier list params
func (o *PostSupplierListParams) WithContext(ctx context.Context) *PostSupplierListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post supplier list params
func (o *PostSupplierListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post supplier list params
func (o *PostSupplierListParams) WithHTTPClient(client *http.Client) *PostSupplierListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post supplier list params
func (o *PostSupplierListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post supplier list params
func (o *PostSupplierListParams) WithBody(body []*models.Supplier) *PostSupplierListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post supplier list params
func (o *PostSupplierListParams) SetBody(body []*models.Supplier) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSupplierListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
