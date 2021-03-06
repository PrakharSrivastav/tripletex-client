// Code generated by go-swagger; DO NOT EDIT.

package voucher_type

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

// NewGetLedgerVoucherTypeParams creates a new GetLedgerVoucherTypeParams object
// with the default values initialized.
func NewGetLedgerVoucherTypeParams() *GetLedgerVoucherTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetLedgerVoucherTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLedgerVoucherTypeParamsWithTimeout creates a new GetLedgerVoucherTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLedgerVoucherTypeParamsWithTimeout(timeout time.Duration) *GetLedgerVoucherTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetLedgerVoucherTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetLedgerVoucherTypeParamsWithContext creates a new GetLedgerVoucherTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLedgerVoucherTypeParamsWithContext(ctx context.Context) *GetLedgerVoucherTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetLedgerVoucherTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetLedgerVoucherTypeParamsWithHTTPClient creates a new GetLedgerVoucherTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLedgerVoucherTypeParamsWithHTTPClient(client *http.Client) *GetLedgerVoucherTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetLedgerVoucherTypeParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetLedgerVoucherTypeParams contains all the parameters to send to the API endpoint
for the get ledger voucher type operation typically these are written to a http.Request
*/
type GetLedgerVoucherTypeParams struct {

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
	/*Name
	  Containing

	*/
	Name *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) WithTimeout(timeout time.Duration) *GetLedgerVoucherTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) WithContext(ctx context.Context) *GetLedgerVoucherTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) WithHTTPClient(client *http.Client) *GetLedgerVoucherTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) WithCount(count *int64) *GetLedgerVoucherTypeParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) WithFields(fields *string) *GetLedgerVoucherTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) WithFrom(from *int64) *GetLedgerVoucherTypeParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) SetFrom(from *int64) {
	o.From = from
}

// WithName adds the name to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) WithName(name *string) *GetLedgerVoucherTypeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) SetName(name *string) {
	o.Name = name
}

// WithSorting adds the sorting to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) WithSorting(sorting *string) *GetLedgerVoucherTypeParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get ledger voucher type params
func (o *GetLedgerVoucherTypeParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetLedgerVoucherTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
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
