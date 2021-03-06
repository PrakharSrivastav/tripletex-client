// Code generated by go-swagger; DO NOT EDIT.

package mileage_allowance

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

// NewPutTravelExpenseMileageAllowanceIDParams creates a new PutTravelExpenseMileageAllowanceIDParams object
// with the default values initialized.
func NewPutTravelExpenseMileageAllowanceIDParams() *PutTravelExpenseMileageAllowanceIDParams {
	var ()
	return &PutTravelExpenseMileageAllowanceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutTravelExpenseMileageAllowanceIDParamsWithTimeout creates a new PutTravelExpenseMileageAllowanceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutTravelExpenseMileageAllowanceIDParamsWithTimeout(timeout time.Duration) *PutTravelExpenseMileageAllowanceIDParams {
	var ()
	return &PutTravelExpenseMileageAllowanceIDParams{

		timeout: timeout,
	}
}

// NewPutTravelExpenseMileageAllowanceIDParamsWithContext creates a new PutTravelExpenseMileageAllowanceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutTravelExpenseMileageAllowanceIDParamsWithContext(ctx context.Context) *PutTravelExpenseMileageAllowanceIDParams {
	var ()
	return &PutTravelExpenseMileageAllowanceIDParams{

		Context: ctx,
	}
}

// NewPutTravelExpenseMileageAllowanceIDParamsWithHTTPClient creates a new PutTravelExpenseMileageAllowanceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutTravelExpenseMileageAllowanceIDParamsWithHTTPClient(client *http.Client) *PutTravelExpenseMileageAllowanceIDParams {
	var ()
	return &PutTravelExpenseMileageAllowanceIDParams{
		HTTPClient: client,
	}
}

/*PutTravelExpenseMileageAllowanceIDParams contains all the parameters to send to the API endpoint
for the put travel expense mileage allowance ID operation typically these are written to a http.Request
*/
type PutTravelExpenseMileageAllowanceIDParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.MileageAllowance
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put travel expense mileage allowance ID params
func (o *PutTravelExpenseMileageAllowanceIDParams) WithTimeout(timeout time.Duration) *PutTravelExpenseMileageAllowanceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put travel expense mileage allowance ID params
func (o *PutTravelExpenseMileageAllowanceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put travel expense mileage allowance ID params
func (o *PutTravelExpenseMileageAllowanceIDParams) WithContext(ctx context.Context) *PutTravelExpenseMileageAllowanceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put travel expense mileage allowance ID params
func (o *PutTravelExpenseMileageAllowanceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put travel expense mileage allowance ID params
func (o *PutTravelExpenseMileageAllowanceIDParams) WithHTTPClient(client *http.Client) *PutTravelExpenseMileageAllowanceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put travel expense mileage allowance ID params
func (o *PutTravelExpenseMileageAllowanceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put travel expense mileage allowance ID params
func (o *PutTravelExpenseMileageAllowanceIDParams) WithBody(body *models.MileageAllowance) *PutTravelExpenseMileageAllowanceIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put travel expense mileage allowance ID params
func (o *PutTravelExpenseMileageAllowanceIDParams) SetBody(body *models.MileageAllowance) {
	o.Body = body
}

// WithID adds the id to the put travel expense mileage allowance ID params
func (o *PutTravelExpenseMileageAllowanceIDParams) WithID(id int32) *PutTravelExpenseMileageAllowanceIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put travel expense mileage allowance ID params
func (o *PutTravelExpenseMileageAllowanceIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutTravelExpenseMileageAllowanceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
