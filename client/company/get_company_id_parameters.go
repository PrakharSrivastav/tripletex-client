// Code generated by go-swagger; DO NOT EDIT.

package company

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

// NewGetCompanyIDParams creates a new GetCompanyIDParams object
// with the default values initialized.
func NewGetCompanyIDParams() *GetCompanyIDParams {
	var ()
	return &GetCompanyIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCompanyIDParamsWithTimeout creates a new GetCompanyIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCompanyIDParamsWithTimeout(timeout time.Duration) *GetCompanyIDParams {
	var ()
	return &GetCompanyIDParams{

		timeout: timeout,
	}
}

// NewGetCompanyIDParamsWithContext creates a new GetCompanyIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCompanyIDParamsWithContext(ctx context.Context) *GetCompanyIDParams {
	var ()
	return &GetCompanyIDParams{

		Context: ctx,
	}
}

// NewGetCompanyIDParamsWithHTTPClient creates a new GetCompanyIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCompanyIDParamsWithHTTPClient(client *http.Client) *GetCompanyIDParams {
	var ()
	return &GetCompanyIDParams{
		HTTPClient: client,
	}
}

/*GetCompanyIDParams contains all the parameters to send to the API endpoint
for the get company ID operation typically these are written to a http.Request
*/
type GetCompanyIDParams struct {

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

// WithTimeout adds the timeout to the get company ID params
func (o *GetCompanyIDParams) WithTimeout(timeout time.Duration) *GetCompanyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get company ID params
func (o *GetCompanyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get company ID params
func (o *GetCompanyIDParams) WithContext(ctx context.Context) *GetCompanyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get company ID params
func (o *GetCompanyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get company ID params
func (o *GetCompanyIDParams) WithHTTPClient(client *http.Client) *GetCompanyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get company ID params
func (o *GetCompanyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get company ID params
func (o *GetCompanyIDParams) WithFields(fields *string) *GetCompanyIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get company ID params
func (o *GetCompanyIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get company ID params
func (o *GetCompanyIDParams) WithID(id int32) *GetCompanyIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get company ID params
func (o *GetCompanyIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCompanyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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