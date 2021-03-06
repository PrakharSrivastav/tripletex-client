// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// NewPutEventSubscriptionIDParams creates a new PutEventSubscriptionIDParams object
// with the default values initialized.
func NewPutEventSubscriptionIDParams() *PutEventSubscriptionIDParams {
	var ()
	return &PutEventSubscriptionIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutEventSubscriptionIDParamsWithTimeout creates a new PutEventSubscriptionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutEventSubscriptionIDParamsWithTimeout(timeout time.Duration) *PutEventSubscriptionIDParams {
	var ()
	return &PutEventSubscriptionIDParams{

		timeout: timeout,
	}
}

// NewPutEventSubscriptionIDParamsWithContext creates a new PutEventSubscriptionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutEventSubscriptionIDParamsWithContext(ctx context.Context) *PutEventSubscriptionIDParams {
	var ()
	return &PutEventSubscriptionIDParams{

		Context: ctx,
	}
}

// NewPutEventSubscriptionIDParamsWithHTTPClient creates a new PutEventSubscriptionIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutEventSubscriptionIDParamsWithHTTPClient(client *http.Client) *PutEventSubscriptionIDParams {
	var ()
	return &PutEventSubscriptionIDParams{
		HTTPClient: client,
	}
}

/*PutEventSubscriptionIDParams contains all the parameters to send to the API endpoint
for the put event subscription ID operation typically these are written to a http.Request
*/
type PutEventSubscriptionIDParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.Subscription
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put event subscription ID params
func (o *PutEventSubscriptionIDParams) WithTimeout(timeout time.Duration) *PutEventSubscriptionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put event subscription ID params
func (o *PutEventSubscriptionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put event subscription ID params
func (o *PutEventSubscriptionIDParams) WithContext(ctx context.Context) *PutEventSubscriptionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put event subscription ID params
func (o *PutEventSubscriptionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put event subscription ID params
func (o *PutEventSubscriptionIDParams) WithHTTPClient(client *http.Client) *PutEventSubscriptionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put event subscription ID params
func (o *PutEventSubscriptionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put event subscription ID params
func (o *PutEventSubscriptionIDParams) WithBody(body *models.Subscription) *PutEventSubscriptionIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put event subscription ID params
func (o *PutEventSubscriptionIDParams) SetBody(body *models.Subscription) {
	o.Body = body
}

// WithID adds the id to the put event subscription ID params
func (o *PutEventSubscriptionIDParams) WithID(id int32) *PutEventSubscriptionIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put event subscription ID params
func (o *PutEventSubscriptionIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutEventSubscriptionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
