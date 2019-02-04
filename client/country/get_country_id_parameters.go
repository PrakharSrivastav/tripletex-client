// Code generated by go-swagger; DO NOT EDIT.

package country

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

// NewGetCountryIDParams creates a new GetCountryIDParams object
// with the default values initialized.
func NewGetCountryIDParams() *GetCountryIDParams {
	var ()
	return &GetCountryIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCountryIDParamsWithTimeout creates a new GetCountryIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCountryIDParamsWithTimeout(timeout time.Duration) *GetCountryIDParams {
	var ()
	return &GetCountryIDParams{

		timeout: timeout,
	}
}

// NewGetCountryIDParamsWithContext creates a new GetCountryIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCountryIDParamsWithContext(ctx context.Context) *GetCountryIDParams {
	var ()
	return &GetCountryIDParams{

		Context: ctx,
	}
}

// NewGetCountryIDParamsWithHTTPClient creates a new GetCountryIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCountryIDParamsWithHTTPClient(client *http.Client) *GetCountryIDParams {
	var ()
	return &GetCountryIDParams{
		HTTPClient: client,
	}
}

/*GetCountryIDParams contains all the parameters to send to the API endpoint
for the get country ID operation typically these are written to a http.Request
*/
type GetCountryIDParams struct {

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

// WithTimeout adds the timeout to the get country ID params
func (o *GetCountryIDParams) WithTimeout(timeout time.Duration) *GetCountryIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get country ID params
func (o *GetCountryIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get country ID params
func (o *GetCountryIDParams) WithContext(ctx context.Context) *GetCountryIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get country ID params
func (o *GetCountryIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get country ID params
func (o *GetCountryIDParams) WithHTTPClient(client *http.Client) *GetCountryIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get country ID params
func (o *GetCountryIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get country ID params
func (o *GetCountryIDParams) WithFields(fields *string) *GetCountryIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get country ID params
func (o *GetCountryIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get country ID params
func (o *GetCountryIDParams) WithID(id int32) *GetCountryIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get country ID params
func (o *GetCountryIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCountryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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