// Code generated by go-swagger; DO NOT EDIT.

package details

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

// NewPutEmployeeEmploymentDetailsIDParams creates a new PutEmployeeEmploymentDetailsIDParams object
// with the default values initialized.
func NewPutEmployeeEmploymentDetailsIDParams() *PutEmployeeEmploymentDetailsIDParams {
	var ()
	return &PutEmployeeEmploymentDetailsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutEmployeeEmploymentDetailsIDParamsWithTimeout creates a new PutEmployeeEmploymentDetailsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutEmployeeEmploymentDetailsIDParamsWithTimeout(timeout time.Duration) *PutEmployeeEmploymentDetailsIDParams {
	var ()
	return &PutEmployeeEmploymentDetailsIDParams{

		timeout: timeout,
	}
}

// NewPutEmployeeEmploymentDetailsIDParamsWithContext creates a new PutEmployeeEmploymentDetailsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutEmployeeEmploymentDetailsIDParamsWithContext(ctx context.Context) *PutEmployeeEmploymentDetailsIDParams {
	var ()
	return &PutEmployeeEmploymentDetailsIDParams{

		Context: ctx,
	}
}

// NewPutEmployeeEmploymentDetailsIDParamsWithHTTPClient creates a new PutEmployeeEmploymentDetailsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutEmployeeEmploymentDetailsIDParamsWithHTTPClient(client *http.Client) *PutEmployeeEmploymentDetailsIDParams {
	var ()
	return &PutEmployeeEmploymentDetailsIDParams{
		HTTPClient: client,
	}
}

/*PutEmployeeEmploymentDetailsIDParams contains all the parameters to send to the API endpoint
for the put employee employment details ID operation typically these are written to a http.Request
*/
type PutEmployeeEmploymentDetailsIDParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.EmploymentDetails
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put employee employment details ID params
func (o *PutEmployeeEmploymentDetailsIDParams) WithTimeout(timeout time.Duration) *PutEmployeeEmploymentDetailsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put employee employment details ID params
func (o *PutEmployeeEmploymentDetailsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put employee employment details ID params
func (o *PutEmployeeEmploymentDetailsIDParams) WithContext(ctx context.Context) *PutEmployeeEmploymentDetailsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put employee employment details ID params
func (o *PutEmployeeEmploymentDetailsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put employee employment details ID params
func (o *PutEmployeeEmploymentDetailsIDParams) WithHTTPClient(client *http.Client) *PutEmployeeEmploymentDetailsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put employee employment details ID params
func (o *PutEmployeeEmploymentDetailsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put employee employment details ID params
func (o *PutEmployeeEmploymentDetailsIDParams) WithBody(body *models.EmploymentDetails) *PutEmployeeEmploymentDetailsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put employee employment details ID params
func (o *PutEmployeeEmploymentDetailsIDParams) SetBody(body *models.EmploymentDetails) {
	o.Body = body
}

// WithID adds the id to the put employee employment details ID params
func (o *PutEmployeeEmploymentDetailsIDParams) WithID(id int32) *PutEmployeeEmploymentDetailsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put employee employment details ID params
func (o *PutEmployeeEmploymentDetailsIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutEmployeeEmploymentDetailsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
