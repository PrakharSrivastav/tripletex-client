// Code generated by go-swagger; DO NOT EDIT.

package payment_type

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

// NewGetTravelExpensePaymentTypeIDParams creates a new GetTravelExpensePaymentTypeIDParams object
// with the default values initialized.
func NewGetTravelExpensePaymentTypeIDParams() *GetTravelExpensePaymentTypeIDParams {
	var ()
	return &GetTravelExpensePaymentTypeIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTravelExpensePaymentTypeIDParamsWithTimeout creates a new GetTravelExpensePaymentTypeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTravelExpensePaymentTypeIDParamsWithTimeout(timeout time.Duration) *GetTravelExpensePaymentTypeIDParams {
	var ()
	return &GetTravelExpensePaymentTypeIDParams{

		timeout: timeout,
	}
}

// NewGetTravelExpensePaymentTypeIDParamsWithContext creates a new GetTravelExpensePaymentTypeIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTravelExpensePaymentTypeIDParamsWithContext(ctx context.Context) *GetTravelExpensePaymentTypeIDParams {
	var ()
	return &GetTravelExpensePaymentTypeIDParams{

		Context: ctx,
	}
}

// NewGetTravelExpensePaymentTypeIDParamsWithHTTPClient creates a new GetTravelExpensePaymentTypeIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTravelExpensePaymentTypeIDParamsWithHTTPClient(client *http.Client) *GetTravelExpensePaymentTypeIDParams {
	var ()
	return &GetTravelExpensePaymentTypeIDParams{
		HTTPClient: client,
	}
}

/*GetTravelExpensePaymentTypeIDParams contains all the parameters to send to the API endpoint
for the get travel expense payment type ID operation typically these are written to a http.Request
*/
type GetTravelExpensePaymentTypeIDParams struct {

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

// WithTimeout adds the timeout to the get travel expense payment type ID params
func (o *GetTravelExpensePaymentTypeIDParams) WithTimeout(timeout time.Duration) *GetTravelExpensePaymentTypeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get travel expense payment type ID params
func (o *GetTravelExpensePaymentTypeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get travel expense payment type ID params
func (o *GetTravelExpensePaymentTypeIDParams) WithContext(ctx context.Context) *GetTravelExpensePaymentTypeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get travel expense payment type ID params
func (o *GetTravelExpensePaymentTypeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get travel expense payment type ID params
func (o *GetTravelExpensePaymentTypeIDParams) WithHTTPClient(client *http.Client) *GetTravelExpensePaymentTypeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get travel expense payment type ID params
func (o *GetTravelExpensePaymentTypeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get travel expense payment type ID params
func (o *GetTravelExpensePaymentTypeIDParams) WithFields(fields *string) *GetTravelExpensePaymentTypeIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get travel expense payment type ID params
func (o *GetTravelExpensePaymentTypeIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get travel expense payment type ID params
func (o *GetTravelExpensePaymentTypeIDParams) WithID(id int32) *GetTravelExpensePaymentTypeIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get travel expense payment type ID params
func (o *GetTravelExpensePaymentTypeIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTravelExpensePaymentTypeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
