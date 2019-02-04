// Code generated by go-swagger; DO NOT EDIT.

package municipality

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

// NewGetMunicipalityParams creates a new GetMunicipalityParams object
// with the default values initialized.
func NewGetMunicipalityParams() *GetMunicipalityParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetMunicipalityParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMunicipalityParamsWithTimeout creates a new GetMunicipalityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMunicipalityParamsWithTimeout(timeout time.Duration) *GetMunicipalityParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetMunicipalityParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetMunicipalityParamsWithContext creates a new GetMunicipalityParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMunicipalityParamsWithContext(ctx context.Context) *GetMunicipalityParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetMunicipalityParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetMunicipalityParamsWithHTTPClient creates a new GetMunicipalityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMunicipalityParamsWithHTTPClient(client *http.Client) *GetMunicipalityParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetMunicipalityParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetMunicipalityParams contains all the parameters to send to the API endpoint
for the get municipality operation typically these are written to a http.Request
*/
type GetMunicipalityParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get municipality params
func (o *GetMunicipalityParams) WithTimeout(timeout time.Duration) *GetMunicipalityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get municipality params
func (o *GetMunicipalityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get municipality params
func (o *GetMunicipalityParams) WithContext(ctx context.Context) *GetMunicipalityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get municipality params
func (o *GetMunicipalityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get municipality params
func (o *GetMunicipalityParams) WithHTTPClient(client *http.Client) *GetMunicipalityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get municipality params
func (o *GetMunicipalityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get municipality params
func (o *GetMunicipalityParams) WithCount(count *int64) *GetMunicipalityParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get municipality params
func (o *GetMunicipalityParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the get municipality params
func (o *GetMunicipalityParams) WithFields(fields *string) *GetMunicipalityParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get municipality params
func (o *GetMunicipalityParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get municipality params
func (o *GetMunicipalityParams) WithFrom(from *int64) *GetMunicipalityParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get municipality params
func (o *GetMunicipalityParams) SetFrom(from *int64) {
	o.From = from
}

// WithSorting adds the sorting to the get municipality params
func (o *GetMunicipalityParams) WithSorting(sorting *string) *GetMunicipalityParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get municipality params
func (o *GetMunicipalityParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetMunicipalityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

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

	if o.From != nil {

		// query param from
		var qrFrom int64
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := swag.FormatInt64(qrFrom)
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.Sorting != nil {

		// query param sorting
		var qrSorting string
		if o.Sorting != nil {
			qrSorting = *o.Sorting
		}
		qSorting := qrSorting
		if qSorting != "" {
			if err := r.SetQueryParam("sorting", qSorting); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
