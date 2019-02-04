// Code generated by go-swagger; DO NOT EDIT.

package voucher

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

// NewGetLedgerVoucherNonPostedParams creates a new GetLedgerVoucherNonPostedParams object
// with the default values initialized.
func NewGetLedgerVoucherNonPostedParams() *GetLedgerVoucherNonPostedParams {
	var (
		countDefault              = int64(1000)
		fromDefault               = int64(0)
		includeNonApprovedDefault = bool(false)
	)
	return &GetLedgerVoucherNonPostedParams{
		Count:              &countDefault,
		From:               &fromDefault,
		IncludeNonApproved: includeNonApprovedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLedgerVoucherNonPostedParamsWithTimeout creates a new GetLedgerVoucherNonPostedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLedgerVoucherNonPostedParamsWithTimeout(timeout time.Duration) *GetLedgerVoucherNonPostedParams {
	var (
		countDefault              = int64(1000)
		fromDefault               = int64(0)
		includeNonApprovedDefault = bool(false)
	)
	return &GetLedgerVoucherNonPostedParams{
		Count:              &countDefault,
		From:               &fromDefault,
		IncludeNonApproved: includeNonApprovedDefault,

		timeout: timeout,
	}
}

// NewGetLedgerVoucherNonPostedParamsWithContext creates a new GetLedgerVoucherNonPostedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLedgerVoucherNonPostedParamsWithContext(ctx context.Context) *GetLedgerVoucherNonPostedParams {
	var (
		countDefault              = int64(1000)
		fromDefault               = int64(0)
		includeNonApprovedDefault = bool(false)
	)
	return &GetLedgerVoucherNonPostedParams{
		Count:              &countDefault,
		From:               &fromDefault,
		IncludeNonApproved: includeNonApprovedDefault,

		Context: ctx,
	}
}

// NewGetLedgerVoucherNonPostedParamsWithHTTPClient creates a new GetLedgerVoucherNonPostedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLedgerVoucherNonPostedParamsWithHTTPClient(client *http.Client) *GetLedgerVoucherNonPostedParams {
	var (
		countDefault              = int64(1000)
		fromDefault               = int64(0)
		includeNonApprovedDefault = bool(false)
	)
	return &GetLedgerVoucherNonPostedParams{
		Count:              &countDefault,
		From:               &fromDefault,
		IncludeNonApproved: includeNonApprovedDefault,
		HTTPClient:         client,
	}
}

/*GetLedgerVoucherNonPostedParams contains all the parameters to send to the API endpoint
for the get ledger voucher non posted operation typically these are written to a http.Request
*/
type GetLedgerVoucherNonPostedParams struct {

	/*ChangedSince
	  Only return elements that have changed since this date and time

	*/
	ChangedSince *string
	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*DateFrom
	  From and including

	*/
	DateFrom *string
	/*DateTo
	  To and excluding

	*/
	DateTo *string
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*IncludeNonApproved
	  Include non-approved vouchers in the result.

	*/
	IncludeNonApproved bool
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) WithTimeout(timeout time.Duration) *GetLedgerVoucherNonPostedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) WithContext(ctx context.Context) *GetLedgerVoucherNonPostedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) WithHTTPClient(client *http.Client) *GetLedgerVoucherNonPostedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangedSince adds the changedSince to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) WithChangedSince(changedSince *string) *GetLedgerVoucherNonPostedParams {
	o.SetChangedSince(changedSince)
	return o
}

// SetChangedSince adds the changedSince to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) SetChangedSince(changedSince *string) {
	o.ChangedSince = changedSince
}

// WithCount adds the count to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) WithCount(count *int64) *GetLedgerVoucherNonPostedParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) SetCount(count *int64) {
	o.Count = count
}

// WithDateFrom adds the dateFrom to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) WithDateFrom(dateFrom *string) *GetLedgerVoucherNonPostedParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) SetDateFrom(dateFrom *string) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) WithDateTo(dateTo *string) *GetLedgerVoucherNonPostedParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) SetDateTo(dateTo *string) {
	o.DateTo = dateTo
}

// WithFields adds the fields to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) WithFields(fields *string) *GetLedgerVoucherNonPostedParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) WithFrom(from *int64) *GetLedgerVoucherNonPostedParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) SetFrom(from *int64) {
	o.From = from
}

// WithIncludeNonApproved adds the includeNonApproved to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) WithIncludeNonApproved(includeNonApproved bool) *GetLedgerVoucherNonPostedParams {
	o.SetIncludeNonApproved(includeNonApproved)
	return o
}

// SetIncludeNonApproved adds the includeNonApproved to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) SetIncludeNonApproved(includeNonApproved bool) {
	o.IncludeNonApproved = includeNonApproved
}

// WithSorting adds the sorting to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) WithSorting(sorting *string) *GetLedgerVoucherNonPostedParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get ledger voucher non posted params
func (o *GetLedgerVoucherNonPostedParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetLedgerVoucherNonPostedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChangedSince != nil {

		// query param changedSince
		var qrChangedSince string
		if o.ChangedSince != nil {
			qrChangedSince = *o.ChangedSince
		}
		qChangedSince := qrChangedSince
		if qChangedSince != "" {
			if err := r.SetQueryParam("changedSince", qChangedSince); err != nil {
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

	if o.DateFrom != nil {

		// query param dateFrom
		var qrDateFrom string
		if o.DateFrom != nil {
			qrDateFrom = *o.DateFrom
		}
		qDateFrom := qrDateFrom
		if qDateFrom != "" {
			if err := r.SetQueryParam("dateFrom", qDateFrom); err != nil {
				return err
			}
		}

	}

	if o.DateTo != nil {

		// query param dateTo
		var qrDateTo string
		if o.DateTo != nil {
			qrDateTo = *o.DateTo
		}
		qDateTo := qrDateTo
		if qDateTo != "" {
			if err := r.SetQueryParam("dateTo", qDateTo); err != nil {
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

	// query param includeNonApproved
	qrIncludeNonApproved := o.IncludeNonApproved
	qIncludeNonApproved := swag.FormatBool(qrIncludeNonApproved)
	if qIncludeNonApproved != "" {
		if err := r.SetQueryParam("includeNonApproved", qIncludeNonApproved); err != nil {
			return err
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