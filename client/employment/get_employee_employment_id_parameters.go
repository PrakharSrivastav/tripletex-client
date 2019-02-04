// Code generated by go-swagger; DO NOT EDIT.

package employment

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

// NewGetEmployeeEmploymentIDParams creates a new GetEmployeeEmploymentIDParams object
// with the default values initialized.
func NewGetEmployeeEmploymentIDParams() *GetEmployeeEmploymentIDParams {
	var ()
	return &GetEmployeeEmploymentIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmployeeEmploymentIDParamsWithTimeout creates a new GetEmployeeEmploymentIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmployeeEmploymentIDParamsWithTimeout(timeout time.Duration) *GetEmployeeEmploymentIDParams {
	var ()
	return &GetEmployeeEmploymentIDParams{

		timeout: timeout,
	}
}

// NewGetEmployeeEmploymentIDParamsWithContext creates a new GetEmployeeEmploymentIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmployeeEmploymentIDParamsWithContext(ctx context.Context) *GetEmployeeEmploymentIDParams {
	var ()
	return &GetEmployeeEmploymentIDParams{

		Context: ctx,
	}
}

// NewGetEmployeeEmploymentIDParamsWithHTTPClient creates a new GetEmployeeEmploymentIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmployeeEmploymentIDParamsWithHTTPClient(client *http.Client) *GetEmployeeEmploymentIDParams {
	var ()
	return &GetEmployeeEmploymentIDParams{
		HTTPClient: client,
	}
}

/*GetEmployeeEmploymentIDParams contains all the parameters to send to the API endpoint
for the get employee employment ID operation typically these are written to a http.Request
*/
type GetEmployeeEmploymentIDParams struct {

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

// WithTimeout adds the timeout to the get employee employment ID params
func (o *GetEmployeeEmploymentIDParams) WithTimeout(timeout time.Duration) *GetEmployeeEmploymentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get employee employment ID params
func (o *GetEmployeeEmploymentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get employee employment ID params
func (o *GetEmployeeEmploymentIDParams) WithContext(ctx context.Context) *GetEmployeeEmploymentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get employee employment ID params
func (o *GetEmployeeEmploymentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get employee employment ID params
func (o *GetEmployeeEmploymentIDParams) WithHTTPClient(client *http.Client) *GetEmployeeEmploymentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get employee employment ID params
func (o *GetEmployeeEmploymentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get employee employment ID params
func (o *GetEmployeeEmploymentIDParams) WithFields(fields *string) *GetEmployeeEmploymentIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get employee employment ID params
func (o *GetEmployeeEmploymentIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get employee employment ID params
func (o *GetEmployeeEmploymentIDParams) WithID(id int32) *GetEmployeeEmploymentIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get employee employment ID params
func (o *GetEmployeeEmploymentIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmployeeEmploymentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
