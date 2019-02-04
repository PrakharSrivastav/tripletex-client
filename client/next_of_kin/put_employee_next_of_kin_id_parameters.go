// Code generated by go-swagger; DO NOT EDIT.

package next_of_kin

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

// NewPutEmployeeNextOfKinIDParams creates a new PutEmployeeNextOfKinIDParams object
// with the default values initialized.
func NewPutEmployeeNextOfKinIDParams() *PutEmployeeNextOfKinIDParams {
	var ()
	return &PutEmployeeNextOfKinIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutEmployeeNextOfKinIDParamsWithTimeout creates a new PutEmployeeNextOfKinIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutEmployeeNextOfKinIDParamsWithTimeout(timeout time.Duration) *PutEmployeeNextOfKinIDParams {
	var ()
	return &PutEmployeeNextOfKinIDParams{

		timeout: timeout,
	}
}

// NewPutEmployeeNextOfKinIDParamsWithContext creates a new PutEmployeeNextOfKinIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutEmployeeNextOfKinIDParamsWithContext(ctx context.Context) *PutEmployeeNextOfKinIDParams {
	var ()
	return &PutEmployeeNextOfKinIDParams{

		Context: ctx,
	}
}

// NewPutEmployeeNextOfKinIDParamsWithHTTPClient creates a new PutEmployeeNextOfKinIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutEmployeeNextOfKinIDParamsWithHTTPClient(client *http.Client) *PutEmployeeNextOfKinIDParams {
	var ()
	return &PutEmployeeNextOfKinIDParams{
		HTTPClient: client,
	}
}

/*PutEmployeeNextOfKinIDParams contains all the parameters to send to the API endpoint
for the put employee next of kin ID operation typically these are written to a http.Request
*/
type PutEmployeeNextOfKinIDParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.NextOfKin
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put employee next of kin ID params
func (o *PutEmployeeNextOfKinIDParams) WithTimeout(timeout time.Duration) *PutEmployeeNextOfKinIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put employee next of kin ID params
func (o *PutEmployeeNextOfKinIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put employee next of kin ID params
func (o *PutEmployeeNextOfKinIDParams) WithContext(ctx context.Context) *PutEmployeeNextOfKinIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put employee next of kin ID params
func (o *PutEmployeeNextOfKinIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put employee next of kin ID params
func (o *PutEmployeeNextOfKinIDParams) WithHTTPClient(client *http.Client) *PutEmployeeNextOfKinIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put employee next of kin ID params
func (o *PutEmployeeNextOfKinIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put employee next of kin ID params
func (o *PutEmployeeNextOfKinIDParams) WithBody(body *models.NextOfKin) *PutEmployeeNextOfKinIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put employee next of kin ID params
func (o *PutEmployeeNextOfKinIDParams) SetBody(body *models.NextOfKin) {
	o.Body = body
}

// WithID adds the id to the put employee next of kin ID params
func (o *PutEmployeeNextOfKinIDParams) WithID(id int32) *PutEmployeeNextOfKinIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put employee next of kin ID params
func (o *PutEmployeeNextOfKinIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutEmployeeNextOfKinIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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