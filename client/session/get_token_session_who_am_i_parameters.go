// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTokenSessionWhoAmIParams creates a new GetTokenSessionWhoAmIParams object
// with the default values initialized.
func NewGetTokenSessionWhoAmIParams() *GetTokenSessionWhoAmIParams {
	var ()
	return &GetTokenSessionWhoAmIParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTokenSessionWhoAmIParamsWithTimeout creates a new GetTokenSessionWhoAmIParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTokenSessionWhoAmIParamsWithTimeout(timeout time.Duration) *GetTokenSessionWhoAmIParams {
	var ()
	return &GetTokenSessionWhoAmIParams{

		timeout: timeout,
	}
}

// NewGetTokenSessionWhoAmIParamsWithContext creates a new GetTokenSessionWhoAmIParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTokenSessionWhoAmIParamsWithContext(ctx context.Context) *GetTokenSessionWhoAmIParams {
	var ()
	return &GetTokenSessionWhoAmIParams{

		Context: ctx,
	}
}

// NewGetTokenSessionWhoAmIParamsWithHTTPClient creates a new GetTokenSessionWhoAmIParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTokenSessionWhoAmIParamsWithHTTPClient(client *http.Client) *GetTokenSessionWhoAmIParams {
	var ()
	return &GetTokenSessionWhoAmIParams{
		HTTPClient: client,
	}
}

/*GetTokenSessionWhoAmIParams contains all the parameters to send to the API endpoint
for the get token session who am i operation typically these are written to a http.Request
*/
type GetTokenSessionWhoAmIParams struct {

	/*Fields
	  Fields filter pattern

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get token session who am i params
func (o *GetTokenSessionWhoAmIParams) WithTimeout(timeout time.Duration) *GetTokenSessionWhoAmIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get token session who am i params
func (o *GetTokenSessionWhoAmIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get token session who am i params
func (o *GetTokenSessionWhoAmIParams) WithContext(ctx context.Context) *GetTokenSessionWhoAmIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get token session who am i params
func (o *GetTokenSessionWhoAmIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get token session who am i params
func (o *GetTokenSessionWhoAmIParams) WithHTTPClient(client *http.Client) *GetTokenSessionWhoAmIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get token session who am i params
func (o *GetTokenSessionWhoAmIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get token session who am i params
func (o *GetTokenSessionWhoAmIParams) WithFields(fields *string) *GetTokenSessionWhoAmIParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get token session who am i params
func (o *GetTokenSessionWhoAmIParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetTokenSessionWhoAmIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
