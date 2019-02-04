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

// NewGetTravelExpenseCostIDParams creates a new GetTravelExpenseCostIDParams object
// with the default values initialized.
func NewGetTravelExpenseCostIDParams() *GetTravelExpenseCostIDParams {
	var ()
	return &GetTravelExpenseCostIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTravelExpenseCostIDParamsWithTimeout creates a new GetTravelExpenseCostIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTravelExpenseCostIDParamsWithTimeout(timeout time.Duration) *GetTravelExpenseCostIDParams {
	var ()
	return &GetTravelExpenseCostIDParams{

		timeout: timeout,
	}
}

// NewGetTravelExpenseCostIDParamsWithContext creates a new GetTravelExpenseCostIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTravelExpenseCostIDParamsWithContext(ctx context.Context) *GetTravelExpenseCostIDParams {
	var ()
	return &GetTravelExpenseCostIDParams{

		Context: ctx,
	}
}

// NewGetTravelExpenseCostIDParamsWithHTTPClient creates a new GetTravelExpenseCostIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTravelExpenseCostIDParamsWithHTTPClient(client *http.Client) *GetTravelExpenseCostIDParams {
	var ()
	return &GetTravelExpenseCostIDParams{
		HTTPClient: client,
	}
}

/*GetTravelExpenseCostIDParams contains all the parameters to send to the API endpoint
for the get travel expense cost ID operation typically these are written to a http.Request
*/
type GetTravelExpenseCostIDParams struct {

	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get travel expense cost ID params
func (o *GetTravelExpenseCostIDParams) WithTimeout(timeout time.Duration) *GetTravelExpenseCostIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get travel expense cost ID params
func (o *GetTravelExpenseCostIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get travel expense cost ID params
func (o *GetTravelExpenseCostIDParams) WithContext(ctx context.Context) *GetTravelExpenseCostIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get travel expense cost ID params
func (o *GetTravelExpenseCostIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get travel expense cost ID params
func (o *GetTravelExpenseCostIDParams) WithHTTPClient(client *http.Client) *GetTravelExpenseCostIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get travel expense cost ID params
func (o *GetTravelExpenseCostIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get travel expense cost ID params
func (o *GetTravelExpenseCostIDParams) WithFields(fields *string) *GetTravelExpenseCostIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get travel expense cost ID params
func (o *GetTravelExpenseCostIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get travel expense cost ID params
func (o *GetTravelExpenseCostIDParams) WithID(id int32) *GetTravelExpenseCostIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get travel expense cost ID params
func (o *GetTravelExpenseCostIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTravelExpenseCostIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
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