// Code generated by go-swagger; DO NOT EDIT.

package details

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

// NewGetEmployeeEmploymentDetailsIDParams creates a new GetEmployeeEmploymentDetailsIDParams object
// with the default values initialized.
func NewGetEmployeeEmploymentDetailsIDParams() *GetEmployeeEmploymentDetailsIDParams {
	var ()
	return &GetEmployeeEmploymentDetailsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmployeeEmploymentDetailsIDParamsWithTimeout creates a new GetEmployeeEmploymentDetailsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmployeeEmploymentDetailsIDParamsWithTimeout(timeout time.Duration) *GetEmployeeEmploymentDetailsIDParams {
	var ()
	return &GetEmployeeEmploymentDetailsIDParams{

		timeout: timeout,
	}
}

// NewGetEmployeeEmploymentDetailsIDParamsWithContext creates a new GetEmployeeEmploymentDetailsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmployeeEmploymentDetailsIDParamsWithContext(ctx context.Context) *GetEmployeeEmploymentDetailsIDParams {
	var ()
	return &GetEmployeeEmploymentDetailsIDParams{

		Context: ctx,
	}
}

// NewGetEmployeeEmploymentDetailsIDParamsWithHTTPClient creates a new GetEmployeeEmploymentDetailsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmployeeEmploymentDetailsIDParamsWithHTTPClient(client *http.Client) *GetEmployeeEmploymentDetailsIDParams {
	var ()
	return &GetEmployeeEmploymentDetailsIDParams{
		HTTPClient: client,
	}
}

/*GetEmployeeEmploymentDetailsIDParams contains all the parameters to send to the API endpoint
for the get employee employment details ID operation typically these are written to a http.Request
*/
type GetEmployeeEmploymentDetailsIDParams struct {

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

// WithTimeout adds the timeout to the get employee employment details ID params
func (o *GetEmployeeEmploymentDetailsIDParams) WithTimeout(timeout time.Duration) *GetEmployeeEmploymentDetailsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get employee employment details ID params
func (o *GetEmployeeEmploymentDetailsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get employee employment details ID params
func (o *GetEmployeeEmploymentDetailsIDParams) WithContext(ctx context.Context) *GetEmployeeEmploymentDetailsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get employee employment details ID params
func (o *GetEmployeeEmploymentDetailsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get employee employment details ID params
func (o *GetEmployeeEmploymentDetailsIDParams) WithHTTPClient(client *http.Client) *GetEmployeeEmploymentDetailsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get employee employment details ID params
func (o *GetEmployeeEmploymentDetailsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get employee employment details ID params
func (o *GetEmployeeEmploymentDetailsIDParams) WithFields(fields *string) *GetEmployeeEmploymentDetailsIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get employee employment details ID params
func (o *GetEmployeeEmploymentDetailsIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get employee employment details ID params
func (o *GetEmployeeEmploymentDetailsIDParams) WithID(id int32) *GetEmployeeEmploymentDetailsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get employee employment details ID params
func (o *GetEmployeeEmploymentDetailsIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmployeeEmploymentDetailsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
