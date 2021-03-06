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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// NewPutCustomerCategoryIDParams creates a new PutCustomerCategoryIDParams object
// with the default values initialized.
func NewPutCustomerCategoryIDParams() *PutCustomerCategoryIDParams {
	var ()
	return &PutCustomerCategoryIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomerCategoryIDParamsWithTimeout creates a new PutCustomerCategoryIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomerCategoryIDParamsWithTimeout(timeout time.Duration) *PutCustomerCategoryIDParams {
	var ()
	return &PutCustomerCategoryIDParams{

		timeout: timeout,
	}
}

// NewPutCustomerCategoryIDParamsWithContext creates a new PutCustomerCategoryIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomerCategoryIDParamsWithContext(ctx context.Context) *PutCustomerCategoryIDParams {
	var ()
	return &PutCustomerCategoryIDParams{

		Context: ctx,
	}
}

// NewPutCustomerCategoryIDParamsWithHTTPClient creates a new PutCustomerCategoryIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomerCategoryIDParamsWithHTTPClient(client *http.Client) *PutCustomerCategoryIDParams {
	var ()
	return &PutCustomerCategoryIDParams{
		HTTPClient: client,
	}
}

/*PutCustomerCategoryIDParams contains all the parameters to send to the API endpoint
for the put customer category ID operation typically these are written to a http.Request
*/
type PutCustomerCategoryIDParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.CustomerCategory
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customer category ID params
func (o *PutCustomerCategoryIDParams) WithTimeout(timeout time.Duration) *PutCustomerCategoryIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customer category ID params
func (o *PutCustomerCategoryIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customer category ID params
func (o *PutCustomerCategoryIDParams) WithContext(ctx context.Context) *PutCustomerCategoryIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customer category ID params
func (o *PutCustomerCategoryIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customer category ID params
func (o *PutCustomerCategoryIDParams) WithHTTPClient(client *http.Client) *PutCustomerCategoryIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customer category ID params
func (o *PutCustomerCategoryIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put customer category ID params
func (o *PutCustomerCategoryIDParams) WithBody(body *models.CustomerCategory) *PutCustomerCategoryIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put customer category ID params
func (o *PutCustomerCategoryIDParams) SetBody(body *models.CustomerCategory) {
	o.Body = body
}

// WithID adds the id to the put customer category ID params
func (o *PutCustomerCategoryIDParams) WithID(id int32) *PutCustomerCategoryIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put customer category ID params
func (o *PutCustomerCategoryIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomerCategoryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
