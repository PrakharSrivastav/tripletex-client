// Code generated by go-swagger; DO NOT EDIT.

package category

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

// NewGetCustomerCategoryIDParams creates a new GetCustomerCategoryIDParams object
// with the default values initialized.
func NewGetCustomerCategoryIDParams() *GetCustomerCategoryIDParams {
	var ()
	return &GetCustomerCategoryIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomerCategoryIDParamsWithTimeout creates a new GetCustomerCategoryIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomerCategoryIDParamsWithTimeout(timeout time.Duration) *GetCustomerCategoryIDParams {
	var ()
	return &GetCustomerCategoryIDParams{

		timeout: timeout,
	}
}

// NewGetCustomerCategoryIDParamsWithContext creates a new GetCustomerCategoryIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomerCategoryIDParamsWithContext(ctx context.Context) *GetCustomerCategoryIDParams {
	var ()
	return &GetCustomerCategoryIDParams{

		Context: ctx,
	}
}

// NewGetCustomerCategoryIDParamsWithHTTPClient creates a new GetCustomerCategoryIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomerCategoryIDParamsWithHTTPClient(client *http.Client) *GetCustomerCategoryIDParams {
	var ()
	return &GetCustomerCategoryIDParams{
		HTTPClient: client,
	}
}

/*GetCustomerCategoryIDParams contains all the parameters to send to the API endpoint
for the get customer category ID operation typically these are written to a http.Request
*/
type GetCustomerCategoryIDParams struct {

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

// WithTimeout adds the timeout to the get customer category ID params
func (o *GetCustomerCategoryIDParams) WithTimeout(timeout time.Duration) *GetCustomerCategoryIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customer category ID params
func (o *GetCustomerCategoryIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customer category ID params
func (o *GetCustomerCategoryIDParams) WithContext(ctx context.Context) *GetCustomerCategoryIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customer category ID params
func (o *GetCustomerCategoryIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customer category ID params
func (o *GetCustomerCategoryIDParams) WithHTTPClient(client *http.Client) *GetCustomerCategoryIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customer category ID params
func (o *GetCustomerCategoryIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get customer category ID params
func (o *GetCustomerCategoryIDParams) WithFields(fields *string) *GetCustomerCategoryIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get customer category ID params
func (o *GetCustomerCategoryIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get customer category ID params
func (o *GetCustomerCategoryIDParams) WithID(id int32) *GetCustomerCategoryIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get customer category ID params
func (o *GetCustomerCategoryIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomerCategoryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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