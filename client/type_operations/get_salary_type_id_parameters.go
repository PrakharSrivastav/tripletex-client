// Code generated by go-swagger; DO NOT EDIT.

package type_operations

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

// NewGetSalaryTypeIDParams creates a new GetSalaryTypeIDParams object
// with the default values initialized.
func NewGetSalaryTypeIDParams() *GetSalaryTypeIDParams {
	var ()
	return &GetSalaryTypeIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSalaryTypeIDParamsWithTimeout creates a new GetSalaryTypeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSalaryTypeIDParamsWithTimeout(timeout time.Duration) *GetSalaryTypeIDParams {
	var ()
	return &GetSalaryTypeIDParams{

		timeout: timeout,
	}
}

// NewGetSalaryTypeIDParamsWithContext creates a new GetSalaryTypeIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSalaryTypeIDParamsWithContext(ctx context.Context) *GetSalaryTypeIDParams {
	var ()
	return &GetSalaryTypeIDParams{

		Context: ctx,
	}
}

// NewGetSalaryTypeIDParamsWithHTTPClient creates a new GetSalaryTypeIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSalaryTypeIDParamsWithHTTPClient(client *http.Client) *GetSalaryTypeIDParams {
	var ()
	return &GetSalaryTypeIDParams{
		HTTPClient: client,
	}
}

/*GetSalaryTypeIDParams contains all the parameters to send to the API endpoint
for the get salary type ID operation typically these are written to a http.Request
*/
type GetSalaryTypeIDParams struct {

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

// WithTimeout adds the timeout to the get salary type ID params
func (o *GetSalaryTypeIDParams) WithTimeout(timeout time.Duration) *GetSalaryTypeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get salary type ID params
func (o *GetSalaryTypeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get salary type ID params
func (o *GetSalaryTypeIDParams) WithContext(ctx context.Context) *GetSalaryTypeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get salary type ID params
func (o *GetSalaryTypeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get salary type ID params
func (o *GetSalaryTypeIDParams) WithHTTPClient(client *http.Client) *GetSalaryTypeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get salary type ID params
func (o *GetSalaryTypeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get salary type ID params
func (o *GetSalaryTypeIDParams) WithFields(fields *string) *GetSalaryTypeIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get salary type ID params
func (o *GetSalaryTypeIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get salary type ID params
func (o *GetSalaryTypeIDParams) WithID(id int32) *GetSalaryTypeIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get salary type ID params
func (o *GetSalaryTypeIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSalaryTypeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
