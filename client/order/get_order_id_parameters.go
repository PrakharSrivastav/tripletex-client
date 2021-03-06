// Code generated by go-swagger; DO NOT EDIT.

package order

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

// NewGetOrderIDParams creates a new GetOrderIDParams object
// with the default values initialized.
func NewGetOrderIDParams() *GetOrderIDParams {
	var ()
	return &GetOrderIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrderIDParamsWithTimeout creates a new GetOrderIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrderIDParamsWithTimeout(timeout time.Duration) *GetOrderIDParams {
	var ()
	return &GetOrderIDParams{

		timeout: timeout,
	}
}

// NewGetOrderIDParamsWithContext creates a new GetOrderIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrderIDParamsWithContext(ctx context.Context) *GetOrderIDParams {
	var ()
	return &GetOrderIDParams{

		Context: ctx,
	}
}

// NewGetOrderIDParamsWithHTTPClient creates a new GetOrderIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrderIDParamsWithHTTPClient(client *http.Client) *GetOrderIDParams {
	var ()
	return &GetOrderIDParams{
		HTTPClient: client,
	}
}

/*GetOrderIDParams contains all the parameters to send to the API endpoint
for the get order ID operation typically these are written to a http.Request
*/
type GetOrderIDParams struct {

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

// WithTimeout adds the timeout to the get order ID params
func (o *GetOrderIDParams) WithTimeout(timeout time.Duration) *GetOrderIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get order ID params
func (o *GetOrderIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get order ID params
func (o *GetOrderIDParams) WithContext(ctx context.Context) *GetOrderIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get order ID params
func (o *GetOrderIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get order ID params
func (o *GetOrderIDParams) WithHTTPClient(client *http.Client) *GetOrderIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get order ID params
func (o *GetOrderIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get order ID params
func (o *GetOrderIDParams) WithFields(fields *string) *GetOrderIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get order ID params
func (o *GetOrderIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get order ID params
func (o *GetOrderIDParams) WithID(id int32) *GetOrderIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get order ID params
func (o *GetOrderIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrderIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
