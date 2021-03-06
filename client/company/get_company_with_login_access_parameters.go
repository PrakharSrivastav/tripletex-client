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

// NewGetCompanyWithLoginAccessParams creates a new GetCompanyWithLoginAccessParams object
// with the default values initialized.
func NewGetCompanyWithLoginAccessParams() *GetCompanyWithLoginAccessParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCompanyWithLoginAccessParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCompanyWithLoginAccessParamsWithTimeout creates a new GetCompanyWithLoginAccessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCompanyWithLoginAccessParamsWithTimeout(timeout time.Duration) *GetCompanyWithLoginAccessParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCompanyWithLoginAccessParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetCompanyWithLoginAccessParamsWithContext creates a new GetCompanyWithLoginAccessParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCompanyWithLoginAccessParamsWithContext(ctx context.Context) *GetCompanyWithLoginAccessParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCompanyWithLoginAccessParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetCompanyWithLoginAccessParamsWithHTTPClient creates a new GetCompanyWithLoginAccessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCompanyWithLoginAccessParamsWithHTTPClient(client *http.Client) *GetCompanyWithLoginAccessParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCompanyWithLoginAccessParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetCompanyWithLoginAccessParams contains all the parameters to send to the API endpoint
for the get company with login access operation typically these are written to a http.Request
*/
type GetCompanyWithLoginAccessParams struct {

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

// WithTimeout adds the timeout to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) WithTimeout(timeout time.Duration) *GetCompanyWithLoginAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) WithContext(ctx context.Context) *GetCompanyWithLoginAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) WithHTTPClient(client *http.Client) *GetCompanyWithLoginAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) WithCount(count *int64) *GetCompanyWithLoginAccessParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) WithFields(fields *string) *GetCompanyWithLoginAccessParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) WithFrom(from *int64) *GetCompanyWithLoginAccessParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) SetFrom(from *int64) {
	o.From = from
}

// WithSorting adds the sorting to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) WithSorting(sorting *string) *GetCompanyWithLoginAccessParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get company with login access params
func (o *GetCompanyWithLoginAccessParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetCompanyWithLoginAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
