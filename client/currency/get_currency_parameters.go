// Code generated by go-swagger; DO NOT EDIT.

package currency

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

// NewGetCurrencyParams creates a new GetCurrencyParams object
// with the default values initialized.
func NewGetCurrencyParams() *GetCurrencyParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCurrencyParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCurrencyParamsWithTimeout creates a new GetCurrencyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCurrencyParamsWithTimeout(timeout time.Duration) *GetCurrencyParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCurrencyParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetCurrencyParamsWithContext creates a new GetCurrencyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCurrencyParamsWithContext(ctx context.Context) *GetCurrencyParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCurrencyParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetCurrencyParamsWithHTTPClient creates a new GetCurrencyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCurrencyParamsWithHTTPClient(client *http.Client) *GetCurrencyParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCurrencyParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetCurrencyParams contains all the parameters to send to the API endpoint
for the get currency operation typically these are written to a http.Request
*/
type GetCurrencyParams struct {

	/*Code
	  Currency codes

	*/
	Code *string
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
	/*ID
	  List of IDs

	*/
	ID *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get currency params
func (o *GetCurrencyParams) WithTimeout(timeout time.Duration) *GetCurrencyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get currency params
func (o *GetCurrencyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get currency params
func (o *GetCurrencyParams) WithContext(ctx context.Context) *GetCurrencyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get currency params
func (o *GetCurrencyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get currency params
func (o *GetCurrencyParams) WithHTTPClient(client *http.Client) *GetCurrencyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get currency params
func (o *GetCurrencyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCode adds the code to the get currency params
func (o *GetCurrencyParams) WithCode(code *string) *GetCurrencyParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the get currency params
func (o *GetCurrencyParams) SetCode(code *string) {
	o.Code = code
}

// WithCount adds the count to the get currency params
func (o *GetCurrencyParams) WithCount(count *int64) *GetCurrencyParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get currency params
func (o *GetCurrencyParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the get currency params
func (o *GetCurrencyParams) WithFields(fields *string) *GetCurrencyParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get currency params
func (o *GetCurrencyParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get currency params
func (o *GetCurrencyParams) WithFrom(from *int64) *GetCurrencyParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get currency params
func (o *GetCurrencyParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the get currency params
func (o *GetCurrencyParams) WithID(id *string) *GetCurrencyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get currency params
func (o *GetCurrencyParams) SetID(id *string) {
	o.ID = id
}

// WithSorting adds the sorting to the get currency params
func (o *GetCurrencyParams) WithSorting(sorting *string) *GetCurrencyParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get currency params
func (o *GetCurrencyParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetCurrencyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Code != nil {

		// query param code
		var qrCode string
		if o.Code != nil {
			qrCode = *o.Code
		}
		qCode := qrCode
		if qCode != "" {
			if err := r.SetQueryParam("code", qCode); err != nil {
				return err
			}
		}

	}

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

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
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