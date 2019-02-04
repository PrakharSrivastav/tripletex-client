// Code generated by go-swagger; DO NOT EDIT.

package department

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

// NewPutDepartmentIDParams creates a new PutDepartmentIDParams object
// with the default values initialized.
func NewPutDepartmentIDParams() *PutDepartmentIDParams {
	var ()
	return &PutDepartmentIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDepartmentIDParamsWithTimeout creates a new PutDepartmentIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDepartmentIDParamsWithTimeout(timeout time.Duration) *PutDepartmentIDParams {
	var ()
	return &PutDepartmentIDParams{

		timeout: timeout,
	}
}

// NewPutDepartmentIDParamsWithContext creates a new PutDepartmentIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDepartmentIDParamsWithContext(ctx context.Context) *PutDepartmentIDParams {
	var ()
	return &PutDepartmentIDParams{

		Context: ctx,
	}
}

// NewPutDepartmentIDParamsWithHTTPClient creates a new PutDepartmentIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDepartmentIDParamsWithHTTPClient(client *http.Client) *PutDepartmentIDParams {
	var ()
	return &PutDepartmentIDParams{
		HTTPClient: client,
	}
}

/*PutDepartmentIDParams contains all the parameters to send to the API endpoint
for the put department ID operation typically these are written to a http.Request
*/
type PutDepartmentIDParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.Department
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put department ID params
func (o *PutDepartmentIDParams) WithTimeout(timeout time.Duration) *PutDepartmentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put department ID params
func (o *PutDepartmentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put department ID params
func (o *PutDepartmentIDParams) WithContext(ctx context.Context) *PutDepartmentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put department ID params
func (o *PutDepartmentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put department ID params
func (o *PutDepartmentIDParams) WithHTTPClient(client *http.Client) *PutDepartmentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put department ID params
func (o *PutDepartmentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put department ID params
func (o *PutDepartmentIDParams) WithBody(body *models.Department) *PutDepartmentIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put department ID params
func (o *PutDepartmentIDParams) SetBody(body *models.Department) {
	o.Body = body
}

// WithID adds the id to the put department ID params
func (o *PutDepartmentIDParams) WithID(id int32) *PutDepartmentIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put department ID params
func (o *PutDepartmentIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDepartmentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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