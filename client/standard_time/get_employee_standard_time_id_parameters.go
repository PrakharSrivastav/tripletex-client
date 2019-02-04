// Code generated by go-swagger; DO NOT EDIT.

package standard_time

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

// NewGetEmployeeStandardTimeIDParams creates a new GetEmployeeStandardTimeIDParams object
// with the default values initialized.
func NewGetEmployeeStandardTimeIDParams() *GetEmployeeStandardTimeIDParams {
	var ()
	return &GetEmployeeStandardTimeIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmployeeStandardTimeIDParamsWithTimeout creates a new GetEmployeeStandardTimeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmployeeStandardTimeIDParamsWithTimeout(timeout time.Duration) *GetEmployeeStandardTimeIDParams {
	var ()
	return &GetEmployeeStandardTimeIDParams{

		timeout: timeout,
	}
}

// NewGetEmployeeStandardTimeIDParamsWithContext creates a new GetEmployeeStandardTimeIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmployeeStandardTimeIDParamsWithContext(ctx context.Context) *GetEmployeeStandardTimeIDParams {
	var ()
	return &GetEmployeeStandardTimeIDParams{

		Context: ctx,
	}
}

// NewGetEmployeeStandardTimeIDParamsWithHTTPClient creates a new GetEmployeeStandardTimeIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmployeeStandardTimeIDParamsWithHTTPClient(client *http.Client) *GetEmployeeStandardTimeIDParams {
	var ()
	return &GetEmployeeStandardTimeIDParams{
		HTTPClient: client,
	}
}

/*GetEmployeeStandardTimeIDParams contains all the parameters to send to the API endpoint
for the get employee standard time ID operation typically these are written to a http.Request
*/
type GetEmployeeStandardTimeIDParams struct {

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

// WithTimeout adds the timeout to the get employee standard time ID params
func (o *GetEmployeeStandardTimeIDParams) WithTimeout(timeout time.Duration) *GetEmployeeStandardTimeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get employee standard time ID params
func (o *GetEmployeeStandardTimeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get employee standard time ID params
func (o *GetEmployeeStandardTimeIDParams) WithContext(ctx context.Context) *GetEmployeeStandardTimeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get employee standard time ID params
func (o *GetEmployeeStandardTimeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get employee standard time ID params
func (o *GetEmployeeStandardTimeIDParams) WithHTTPClient(client *http.Client) *GetEmployeeStandardTimeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get employee standard time ID params
func (o *GetEmployeeStandardTimeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get employee standard time ID params
func (o *GetEmployeeStandardTimeIDParams) WithFields(fields *string) *GetEmployeeStandardTimeIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get employee standard time ID params
func (o *GetEmployeeStandardTimeIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get employee standard time ID params
func (o *GetEmployeeStandardTimeIDParams) WithID(id int32) *GetEmployeeStandardTimeIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get employee standard time ID params
func (o *GetEmployeeStandardTimeIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmployeeStandardTimeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
