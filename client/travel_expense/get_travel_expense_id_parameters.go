// Code generated by go-swagger; DO NOT EDIT.

package travel_expense

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

// NewGetTravelExpenseIDParams creates a new GetTravelExpenseIDParams object
// with the default values initialized.
func NewGetTravelExpenseIDParams() *GetTravelExpenseIDParams {
	var ()
	return &GetTravelExpenseIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTravelExpenseIDParamsWithTimeout creates a new GetTravelExpenseIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTravelExpenseIDParamsWithTimeout(timeout time.Duration) *GetTravelExpenseIDParams {
	var ()
	return &GetTravelExpenseIDParams{

		timeout: timeout,
	}
}

// NewGetTravelExpenseIDParamsWithContext creates a new GetTravelExpenseIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTravelExpenseIDParamsWithContext(ctx context.Context) *GetTravelExpenseIDParams {
	var ()
	return &GetTravelExpenseIDParams{

		Context: ctx,
	}
}

// NewGetTravelExpenseIDParamsWithHTTPClient creates a new GetTravelExpenseIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTravelExpenseIDParamsWithHTTPClient(client *http.Client) *GetTravelExpenseIDParams {
	var ()
	return &GetTravelExpenseIDParams{
		HTTPClient: client,
	}
}

/*GetTravelExpenseIDParams contains all the parameters to send to the API endpoint
for the get travel expense ID operation typically these are written to a http.Request
*/
type GetTravelExpenseIDParams struct {

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

// WithTimeout adds the timeout to the get travel expense ID params
func (o *GetTravelExpenseIDParams) WithTimeout(timeout time.Duration) *GetTravelExpenseIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get travel expense ID params
func (o *GetTravelExpenseIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get travel expense ID params
func (o *GetTravelExpenseIDParams) WithContext(ctx context.Context) *GetTravelExpenseIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get travel expense ID params
func (o *GetTravelExpenseIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get travel expense ID params
func (o *GetTravelExpenseIDParams) WithHTTPClient(client *http.Client) *GetTravelExpenseIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get travel expense ID params
func (o *GetTravelExpenseIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get travel expense ID params
func (o *GetTravelExpenseIDParams) WithFields(fields *string) *GetTravelExpenseIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get travel expense ID params
func (o *GetTravelExpenseIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get travel expense ID params
func (o *GetTravelExpenseIDParams) WithID(id int32) *GetTravelExpenseIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get travel expense ID params
func (o *GetTravelExpenseIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTravelExpenseIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
