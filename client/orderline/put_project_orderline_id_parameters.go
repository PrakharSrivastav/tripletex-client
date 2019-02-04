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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// NewPutProjectOrderlineIDParams creates a new PutProjectOrderlineIDParams object
// with the default values initialized.
func NewPutProjectOrderlineIDParams() *PutProjectOrderlineIDParams {
	var ()
	return &PutProjectOrderlineIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutProjectOrderlineIDParamsWithTimeout creates a new PutProjectOrderlineIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutProjectOrderlineIDParamsWithTimeout(timeout time.Duration) *PutProjectOrderlineIDParams {
	var ()
	return &PutProjectOrderlineIDParams{

		timeout: timeout,
	}
}

// NewPutProjectOrderlineIDParamsWithContext creates a new PutProjectOrderlineIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutProjectOrderlineIDParamsWithContext(ctx context.Context) *PutProjectOrderlineIDParams {
	var ()
	return &PutProjectOrderlineIDParams{

		Context: ctx,
	}
}

// NewPutProjectOrderlineIDParamsWithHTTPClient creates a new PutProjectOrderlineIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutProjectOrderlineIDParamsWithHTTPClient(client *http.Client) *PutProjectOrderlineIDParams {
	var ()
	return &PutProjectOrderlineIDParams{
		HTTPClient: client,
	}
}

/*PutProjectOrderlineIDParams contains all the parameters to send to the API endpoint
for the put project orderline ID operation typically these are written to a http.Request
*/
type PutProjectOrderlineIDParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.ProjectOrderLine
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put project orderline ID params
func (o *PutProjectOrderlineIDParams) WithTimeout(timeout time.Duration) *PutProjectOrderlineIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put project orderline ID params
func (o *PutProjectOrderlineIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put project orderline ID params
func (o *PutProjectOrderlineIDParams) WithContext(ctx context.Context) *PutProjectOrderlineIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put project orderline ID params
func (o *PutProjectOrderlineIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put project orderline ID params
func (o *PutProjectOrderlineIDParams) WithHTTPClient(client *http.Client) *PutProjectOrderlineIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put project orderline ID params
func (o *PutProjectOrderlineIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put project orderline ID params
func (o *PutProjectOrderlineIDParams) WithBody(body *models.ProjectOrderLine) *PutProjectOrderlineIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put project orderline ID params
func (o *PutProjectOrderlineIDParams) SetBody(body *models.ProjectOrderLine) {
	o.Body = body
}

// WithID adds the id to the put project orderline ID params
func (o *PutProjectOrderlineIDParams) WithID(id int32) *PutProjectOrderlineIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put project orderline ID params
func (o *PutProjectOrderlineIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutProjectOrderlineIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
