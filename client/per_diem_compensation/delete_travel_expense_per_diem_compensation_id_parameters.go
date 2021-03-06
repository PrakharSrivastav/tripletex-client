// Code generated by go-swagger; DO NOT EDIT.

package per_diem_compensation

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

// NewDeleteTravelExpensePerDiemCompensationIDParams creates a new DeleteTravelExpensePerDiemCompensationIDParams object
// with the default values initialized.
func NewDeleteTravelExpensePerDiemCompensationIDParams() *DeleteTravelExpensePerDiemCompensationIDParams {
	var ()
	return &DeleteTravelExpensePerDiemCompensationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTravelExpensePerDiemCompensationIDParamsWithTimeout creates a new DeleteTravelExpensePerDiemCompensationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTravelExpensePerDiemCompensationIDParamsWithTimeout(timeout time.Duration) *DeleteTravelExpensePerDiemCompensationIDParams {
	var ()
	return &DeleteTravelExpensePerDiemCompensationIDParams{

		timeout: timeout,
	}
}

// NewDeleteTravelExpensePerDiemCompensationIDParamsWithContext creates a new DeleteTravelExpensePerDiemCompensationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTravelExpensePerDiemCompensationIDParamsWithContext(ctx context.Context) *DeleteTravelExpensePerDiemCompensationIDParams {
	var ()
	return &DeleteTravelExpensePerDiemCompensationIDParams{

		Context: ctx,
	}
}

// NewDeleteTravelExpensePerDiemCompensationIDParamsWithHTTPClient creates a new DeleteTravelExpensePerDiemCompensationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTravelExpensePerDiemCompensationIDParamsWithHTTPClient(client *http.Client) *DeleteTravelExpensePerDiemCompensationIDParams {
	var ()
	return &DeleteTravelExpensePerDiemCompensationIDParams{
		HTTPClient: client,
	}
}

/*DeleteTravelExpensePerDiemCompensationIDParams contains all the parameters to send to the API endpoint
for the delete travel expense per diem compensation ID operation typically these are written to a http.Request
*/
type DeleteTravelExpensePerDiemCompensationIDParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete travel expense per diem compensation ID params
func (o *DeleteTravelExpensePerDiemCompensationIDParams) WithTimeout(timeout time.Duration) *DeleteTravelExpensePerDiemCompensationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete travel expense per diem compensation ID params
func (o *DeleteTravelExpensePerDiemCompensationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete travel expense per diem compensation ID params
func (o *DeleteTravelExpensePerDiemCompensationIDParams) WithContext(ctx context.Context) *DeleteTravelExpensePerDiemCompensationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete travel expense per diem compensation ID params
func (o *DeleteTravelExpensePerDiemCompensationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete travel expense per diem compensation ID params
func (o *DeleteTravelExpensePerDiemCompensationIDParams) WithHTTPClient(client *http.Client) *DeleteTravelExpensePerDiemCompensationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete travel expense per diem compensation ID params
func (o *DeleteTravelExpensePerDiemCompensationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete travel expense per diem compensation ID params
func (o *DeleteTravelExpensePerDiemCompensationIDParams) WithID(id int32) *DeleteTravelExpensePerDiemCompensationIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete travel expense per diem compensation ID params
func (o *DeleteTravelExpensePerDiemCompensationIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTravelExpensePerDiemCompensationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
