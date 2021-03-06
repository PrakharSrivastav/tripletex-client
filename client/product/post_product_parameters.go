// Code generated by go-swagger; DO NOT EDIT.

package product

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

// NewPostProductParams creates a new PostProductParams object
// with the default values initialized.
func NewPostProductParams() *PostProductParams {
	var ()
	return &PostProductParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostProductParamsWithTimeout creates a new PostProductParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostProductParamsWithTimeout(timeout time.Duration) *PostProductParams {
	var ()
	return &PostProductParams{

		timeout: timeout,
	}
}

// NewPostProductParamsWithContext creates a new PostProductParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostProductParamsWithContext(ctx context.Context) *PostProductParams {
	var ()
	return &PostProductParams{

		Context: ctx,
	}
}

// NewPostProductParamsWithHTTPClient creates a new PostProductParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostProductParamsWithHTTPClient(client *http.Client) *PostProductParams {
	var ()
	return &PostProductParams{
		HTTPClient: client,
	}
}

/*PostProductParams contains all the parameters to send to the API endpoint
for the post product operation typically these are written to a http.Request
*/
type PostProductParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.Product

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post product params
func (o *PostProductParams) WithTimeout(timeout time.Duration) *PostProductParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post product params
func (o *PostProductParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post product params
func (o *PostProductParams) WithContext(ctx context.Context) *PostProductParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post product params
func (o *PostProductParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post product params
func (o *PostProductParams) WithHTTPClient(client *http.Client) *PostProductParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post product params
func (o *PostProductParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post product params
func (o *PostProductParams) WithBody(body *models.Product) *PostProductParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post product params
func (o *PostProductParams) SetBody(body *models.Product) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
