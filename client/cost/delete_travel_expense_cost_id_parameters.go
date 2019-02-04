// Code generated by go-swagger; DO NOT EDIT.

package cost

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

// NewDeleteTravelExpenseCostIDParams creates a new DeleteTravelExpenseCostIDParams object
// with the default values initialized.
func NewDeleteTravelExpenseCostIDParams() *DeleteTravelExpenseCostIDParams {
	var ()
	return &DeleteTravelExpenseCostIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTravelExpenseCostIDParamsWithTimeout creates a new DeleteTravelExpenseCostIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTravelExpenseCostIDParamsWithTimeout(timeout time.Duration) *DeleteTravelExpenseCostIDParams {
	var ()
	return &DeleteTravelExpenseCostIDParams{

		timeout: timeout,
	}
}

// NewDeleteTravelExpenseCostIDParamsWithContext creates a new DeleteTravelExpenseCostIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTravelExpenseCostIDParamsWithContext(ctx context.Context) *DeleteTravelExpenseCostIDParams {
	var ()
	return &DeleteTravelExpenseCostIDParams{

		Context: ctx,
	}
}

// NewDeleteTravelExpenseCostIDParamsWithHTTPClient creates a new DeleteTravelExpenseCostIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTravelExpenseCostIDParamsWithHTTPClient(client *http.Client) *DeleteTravelExpenseCostIDParams {
	var ()
	return &DeleteTravelExpenseCostIDParams{
		HTTPClient: client,
	}
}

/*DeleteTravelExpenseCostIDParams contains all the parameters to send to the API endpoint
for the delete travel expense cost ID operation typically these are written to a http.Request
*/
type DeleteTravelExpenseCostIDParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete travel expense cost ID params
func (o *DeleteTravelExpenseCostIDParams) WithTimeout(timeout time.Duration) *DeleteTravelExpenseCostIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete travel expense cost ID params
func (o *DeleteTravelExpenseCostIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete travel expense cost ID params
func (o *DeleteTravelExpenseCostIDParams) WithContext(ctx context.Context) *DeleteTravelExpenseCostIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete travel expense cost ID params
func (o *DeleteTravelExpenseCostIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete travel expense cost ID params
func (o *DeleteTravelExpenseCostIDParams) WithHTTPClient(client *http.Client) *DeleteTravelExpenseCostIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete travel expense cost ID params
func (o *DeleteTravelExpenseCostIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete travel expense cost ID params
func (o *DeleteTravelExpenseCostIDParams) WithID(id int32) *DeleteTravelExpenseCostIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete travel expense cost ID params
func (o *DeleteTravelExpenseCostIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTravelExpenseCostIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
