// Code generated by go-swagger; DO NOT EDIT.

package accommodation_allowance

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

// NewDeleteTravelExpenseAccommodationAllowanceIDParams creates a new DeleteTravelExpenseAccommodationAllowanceIDParams object
// with the default values initialized.
func NewDeleteTravelExpenseAccommodationAllowanceIDParams() *DeleteTravelExpenseAccommodationAllowanceIDParams {
	var ()
	return &DeleteTravelExpenseAccommodationAllowanceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTravelExpenseAccommodationAllowanceIDParamsWithTimeout creates a new DeleteTravelExpenseAccommodationAllowanceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTravelExpenseAccommodationAllowanceIDParamsWithTimeout(timeout time.Duration) *DeleteTravelExpenseAccommodationAllowanceIDParams {
	var ()
	return &DeleteTravelExpenseAccommodationAllowanceIDParams{

		timeout: timeout,
	}
}

// NewDeleteTravelExpenseAccommodationAllowanceIDParamsWithContext creates a new DeleteTravelExpenseAccommodationAllowanceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTravelExpenseAccommodationAllowanceIDParamsWithContext(ctx context.Context) *DeleteTravelExpenseAccommodationAllowanceIDParams {
	var ()
	return &DeleteTravelExpenseAccommodationAllowanceIDParams{

		Context: ctx,
	}
}

// NewDeleteTravelExpenseAccommodationAllowanceIDParamsWithHTTPClient creates a new DeleteTravelExpenseAccommodationAllowanceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTravelExpenseAccommodationAllowanceIDParamsWithHTTPClient(client *http.Client) *DeleteTravelExpenseAccommodationAllowanceIDParams {
	var ()
	return &DeleteTravelExpenseAccommodationAllowanceIDParams{
		HTTPClient: client,
	}
}

/*DeleteTravelExpenseAccommodationAllowanceIDParams contains all the parameters to send to the API endpoint
for the delete travel expense accommodation allowance ID operation typically these are written to a http.Request
*/
type DeleteTravelExpenseAccommodationAllowanceIDParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete travel expense accommodation allowance ID params
func (o *DeleteTravelExpenseAccommodationAllowanceIDParams) WithTimeout(timeout time.Duration) *DeleteTravelExpenseAccommodationAllowanceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete travel expense accommodation allowance ID params
func (o *DeleteTravelExpenseAccommodationAllowanceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete travel expense accommodation allowance ID params
func (o *DeleteTravelExpenseAccommodationAllowanceIDParams) WithContext(ctx context.Context) *DeleteTravelExpenseAccommodationAllowanceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete travel expense accommodation allowance ID params
func (o *DeleteTravelExpenseAccommodationAllowanceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete travel expense accommodation allowance ID params
func (o *DeleteTravelExpenseAccommodationAllowanceIDParams) WithHTTPClient(client *http.Client) *DeleteTravelExpenseAccommodationAllowanceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete travel expense accommodation allowance ID params
func (o *DeleteTravelExpenseAccommodationAllowanceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete travel expense accommodation allowance ID params
func (o *DeleteTravelExpenseAccommodationAllowanceIDParams) WithID(id int32) *DeleteTravelExpenseAccommodationAllowanceIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete travel expense accommodation allowance ID params
func (o *DeleteTravelExpenseAccommodationAllowanceIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTravelExpenseAccommodationAllowanceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
