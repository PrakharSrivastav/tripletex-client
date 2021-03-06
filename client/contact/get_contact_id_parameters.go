// Code generated by go-swagger; DO NOT EDIT.

package contact

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

// NewGetContactIDParams creates a new GetContactIDParams object
// with the default values initialized.
func NewGetContactIDParams() *GetContactIDParams {
	var ()
	return &GetContactIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContactIDParamsWithTimeout creates a new GetContactIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContactIDParamsWithTimeout(timeout time.Duration) *GetContactIDParams {
	var ()
	return &GetContactIDParams{

		timeout: timeout,
	}
}

// NewGetContactIDParamsWithContext creates a new GetContactIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContactIDParamsWithContext(ctx context.Context) *GetContactIDParams {
	var ()
	return &GetContactIDParams{

		Context: ctx,
	}
}

// NewGetContactIDParamsWithHTTPClient creates a new GetContactIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContactIDParamsWithHTTPClient(client *http.Client) *GetContactIDParams {
	var ()
	return &GetContactIDParams{
		HTTPClient: client,
	}
}

/*GetContactIDParams contains all the parameters to send to the API endpoint
for the get contact ID operation typically these are written to a http.Request
*/
type GetContactIDParams struct {

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

// WithTimeout adds the timeout to the get contact ID params
func (o *GetContactIDParams) WithTimeout(timeout time.Duration) *GetContactIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contact ID params
func (o *GetContactIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contact ID params
func (o *GetContactIDParams) WithContext(ctx context.Context) *GetContactIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contact ID params
func (o *GetContactIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contact ID params
func (o *GetContactIDParams) WithHTTPClient(client *http.Client) *GetContactIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contact ID params
func (o *GetContactIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get contact ID params
func (o *GetContactIDParams) WithFields(fields *string) *GetContactIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get contact ID params
func (o *GetContactIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get contact ID params
func (o *GetContactIDParams) WithID(id int32) *GetContactIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get contact ID params
func (o *GetContactIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetContactIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
