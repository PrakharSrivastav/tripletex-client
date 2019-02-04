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
)

// NewDeleteTravelExpenseMileageAllowanceIDParams creates a new DeleteTravelExpenseMileageAllowanceIDParams object
// with the default values initialized.
func NewDeleteTravelExpenseMileageAllowanceIDParams() *DeleteTravelExpenseMileageAllowanceIDParams {
	var ()
	return &DeleteTravelExpenseMileageAllowanceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTravelExpenseMileageAllowanceIDParamsWithTimeout creates a new DeleteTravelExpenseMileageAllowanceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTravelExpenseMileageAllowanceIDParamsWithTimeout(timeout time.Duration) *DeleteTravelExpenseMileageAllowanceIDParams {
	var ()
	return &DeleteTravelExpenseMileageAllowanceIDParams{

		timeout: timeout,
	}
}

// NewDeleteTravelExpenseMileageAllowanceIDParamsWithContext creates a new DeleteTravelExpenseMileageAllowanceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTravelExpenseMileageAllowanceIDParamsWithContext(ctx context.Context) *DeleteTravelExpenseMileageAllowanceIDParams {
	var ()
	return &DeleteTravelExpenseMileageAllowanceIDParams{

		Context: ctx,
	}
}

// NewDeleteTravelExpenseMileageAllowanceIDParamsWithHTTPClient creates a new DeleteTravelExpenseMileageAllowanceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTravelExpenseMileageAllowanceIDParamsWithHTTPClient(client *http.Client) *DeleteTravelExpenseMileageAllowanceIDParams {
	var ()
	return &DeleteTravelExpenseMileageAllowanceIDParams{
		HTTPClient: client,
	}
}

/*DeleteTravelExpenseMileageAllowanceIDParams contains all the parameters to send to the API endpoint
for the delete travel expense mileage allowance ID operation typically these are written to a http.Request
*/
type DeleteTravelExpenseMileageAllowanceIDParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete travel expense mileage allowance ID params
func (o *DeleteTravelExpenseMileageAllowanceIDParams) WithTimeout(timeout time.Duration) *DeleteTravelExpenseMileageAllowanceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete travel expense mileage allowance ID params
func (o *DeleteTravelExpenseMileageAllowanceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete travel expense mileage allowance ID params
func (o *DeleteTravelExpenseMileageAllowanceIDParams) WithContext(ctx context.Context) *DeleteTravelExpenseMileageAllowanceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete travel expense mileage allowance ID params
func (o *DeleteTravelExpenseMileageAllowanceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete travel expense mileage allowance ID params
func (o *DeleteTravelExpenseMileageAllowanceIDParams) WithHTTPClient(client *http.Client) *DeleteTravelExpenseMileageAllowanceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete travel expense mileage allowance ID params
func (o *DeleteTravelExpenseMileageAllowanceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete travel expense mileage allowance ID params
func (o *DeleteTravelExpenseMileageAllowanceIDParams) WithID(id int32) *DeleteTravelExpenseMileageAllowanceIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete travel expense mileage allowance ID params
func (o *DeleteTravelExpenseMileageAllowanceIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTravelExpenseMileageAllowanceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
