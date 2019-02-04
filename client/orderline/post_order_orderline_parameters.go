// Code generated by go-swagger; DO NOT EDIT.

package orderline

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

// NewPostOrderOrderlineParams creates a new PostOrderOrderlineParams object
// with the default values initialized.
func NewPostOrderOrderlineParams() *PostOrderOrderlineParams {
	var ()
	return &PostOrderOrderlineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOrderOrderlineParamsWithTimeout creates a new PostOrderOrderlineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOrderOrderlineParamsWithTimeout(timeout time.Duration) *PostOrderOrderlineParams {
	var ()
	return &PostOrderOrderlineParams{

		timeout: timeout,
	}
}

// NewPostOrderOrderlineParamsWithContext creates a new PostOrderOrderlineParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOrderOrderlineParamsWithContext(ctx context.Context) *PostOrderOrderlineParams {
	var ()
	return &PostOrderOrderlineParams{

		Context: ctx,
	}
}

// NewPostOrderOrderlineParamsWithHTTPClient creates a new PostOrderOrderlineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOrderOrderlineParamsWithHTTPClient(client *http.Client) *PostOrderOrderlineParams {
	var ()
	return &PostOrderOrderlineParams{
		HTTPClient: client,
	}
}

/*PostOrderOrderlineParams contains all the parameters to send to the API endpoint
for the post order orderline operation typically these are written to a http.Request
*/
type PostOrderOrderlineParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.OrderLine

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post order orderline params
func (o *PostOrderOrderlineParams) WithTimeout(timeout time.Duration) *PostOrderOrderlineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post order orderline params
func (o *PostOrderOrderlineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post order orderline params
func (o *PostOrderOrderlineParams) WithContext(ctx context.Context) *PostOrderOrderlineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post order orderline params
func (o *PostOrderOrderlineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post order orderline params
func (o *PostOrderOrderlineParams) WithHTTPClient(client *http.Client) *PostOrderOrderlineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post order orderline params
func (o *PostOrderOrderlineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post order orderline params
func (o *PostOrderOrderlineParams) WithBody(body *models.OrderLine) *PostOrderOrderlineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post order orderline params
func (o *PostOrderOrderlineParams) SetBody(body *models.OrderLine) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostOrderOrderlineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
