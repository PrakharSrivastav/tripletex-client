// Code generated by go-swagger; DO NOT EDIT.

package category

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

// NewPostCustomerCategoryParams creates a new PostCustomerCategoryParams object
// with the default values initialized.
func NewPostCustomerCategoryParams() *PostCustomerCategoryParams {
	var ()
	return &PostCustomerCategoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomerCategoryParamsWithTimeout creates a new PostCustomerCategoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomerCategoryParamsWithTimeout(timeout time.Duration) *PostCustomerCategoryParams {
	var ()
	return &PostCustomerCategoryParams{

		timeout: timeout,
	}
}

// NewPostCustomerCategoryParamsWithContext creates a new PostCustomerCategoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomerCategoryParamsWithContext(ctx context.Context) *PostCustomerCategoryParams {
	var ()
	return &PostCustomerCategoryParams{

		Context: ctx,
	}
}

// NewPostCustomerCategoryParamsWithHTTPClient creates a new PostCustomerCategoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomerCategoryParamsWithHTTPClient(client *http.Client) *PostCustomerCategoryParams {
	var ()
	return &PostCustomerCategoryParams{
		HTTPClient: client,
	}
}

/*PostCustomerCategoryParams contains all the parameters to send to the API endpoint
for the post customer category operation typically these are written to a http.Request
*/
type PostCustomerCategoryParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.CustomerCategory

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post customer category params
func (o *PostCustomerCategoryParams) WithTimeout(timeout time.Duration) *PostCustomerCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customer category params
func (o *PostCustomerCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customer category params
func (o *PostCustomerCategoryParams) WithContext(ctx context.Context) *PostCustomerCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customer category params
func (o *PostCustomerCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customer category params
func (o *PostCustomerCategoryParams) WithHTTPClient(client *http.Client) *PostCustomerCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customer category params
func (o *PostCustomerCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post customer category params
func (o *PostCustomerCategoryParams) WithBody(body *models.CustomerCategory) *PostCustomerCategoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post customer category params
func (o *PostCustomerCategoryParams) SetBody(body *models.CustomerCategory) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomerCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
