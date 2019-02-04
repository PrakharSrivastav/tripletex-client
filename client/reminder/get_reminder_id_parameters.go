// Code generated by go-swagger; DO NOT EDIT.

package reminder

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

// NewGetReminderIDParams creates a new GetReminderIDParams object
// with the default values initialized.
func NewGetReminderIDParams() *GetReminderIDParams {
	var ()
	return &GetReminderIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReminderIDParamsWithTimeout creates a new GetReminderIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReminderIDParamsWithTimeout(timeout time.Duration) *GetReminderIDParams {
	var ()
	return &GetReminderIDParams{

		timeout: timeout,
	}
}

// NewGetReminderIDParamsWithContext creates a new GetReminderIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReminderIDParamsWithContext(ctx context.Context) *GetReminderIDParams {
	var ()
	return &GetReminderIDParams{

		Context: ctx,
	}
}

// NewGetReminderIDParamsWithHTTPClient creates a new GetReminderIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReminderIDParamsWithHTTPClient(client *http.Client) *GetReminderIDParams {
	var ()
	return &GetReminderIDParams{
		HTTPClient: client,
	}
}

/*GetReminderIDParams contains all the parameters to send to the API endpoint
for the get reminder ID operation typically these are written to a http.Request
*/
type GetReminderIDParams struct {

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

// WithTimeout adds the timeout to the get reminder ID params
func (o *GetReminderIDParams) WithTimeout(timeout time.Duration) *GetReminderIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get reminder ID params
func (o *GetReminderIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get reminder ID params
func (o *GetReminderIDParams) WithContext(ctx context.Context) *GetReminderIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get reminder ID params
func (o *GetReminderIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get reminder ID params
func (o *GetReminderIDParams) WithHTTPClient(client *http.Client) *GetReminderIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get reminder ID params
func (o *GetReminderIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get reminder ID params
func (o *GetReminderIDParams) WithFields(fields *string) *GetReminderIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get reminder ID params
func (o *GetReminderIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get reminder ID params
func (o *GetReminderIDParams) WithID(id int32) *GetReminderIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get reminder ID params
func (o *GetReminderIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetReminderIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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