// Code generated by go-swagger; DO NOT EDIT.

package supplier

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

// NewGetSupplierParams creates a new GetSupplierParams object
// with the default values initialized.
func NewGetSupplierParams() *GetSupplierParams {
	var (
		countDefault      = int64(1000)
		fromDefault       = int64(0)
		isInactiveDefault = bool(false)
	)
	return &GetSupplierParams{
		Count:      &countDefault,
		From:       &fromDefault,
		IsInactive: &isInactiveDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSupplierParamsWithTimeout creates a new GetSupplierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSupplierParamsWithTimeout(timeout time.Duration) *GetSupplierParams {
	var (
		countDefault      = int64(1000)
		fromDefault       = int64(0)
		isInactiveDefault = bool(false)
	)
	return &GetSupplierParams{
		Count:      &countDefault,
		From:       &fromDefault,
		IsInactive: &isInactiveDefault,

		timeout: timeout,
	}
}

// NewGetSupplierParamsWithContext creates a new GetSupplierParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSupplierParamsWithContext(ctx context.Context) *GetSupplierParams {
	var (
		countDefault      = int64(1000)
		fromDefault       = int64(0)
		isInactiveDefault = bool(false)
	)
	return &GetSupplierParams{
		Count:      &countDefault,
		From:       &fromDefault,
		IsInactive: &isInactiveDefault,

		Context: ctx,
	}
}

// NewGetSupplierParamsWithHTTPClient creates a new GetSupplierParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSupplierParamsWithHTTPClient(client *http.Client) *GetSupplierParams {
	var (
		countDefault      = int64(1000)
		fromDefault       = int64(0)
		isInactiveDefault = bool(false)
	)
	return &GetSupplierParams{
		Count:      &countDefault,
		From:       &fromDefault,
		IsInactive: &isInactiveDefault,
		HTTPClient: client,
	}
}

/*GetSupplierParams contains all the parameters to send to the API endpoint
for the get supplier operation typically these are written to a http.Request
*/
type GetSupplierParams struct {

	/*AccountManagerID
	  List of IDs

	*/
	AccountManagerID *string
	/*ChangedSince
	  Only return elements that have changed since this date and time

	*/
	ChangedSince *string
	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*Email
	  Equals

	*/
	Email *string
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
	/*InvoiceEmail
	  Equals

	*/
	InvoiceEmail *string
	/*IsInactive
	  Equals

	*/
	IsInactive *bool
	/*IsWholesaler
	  Equals

	*/
	IsWholesaler *bool
	/*OrganizationNumber
	  Equals

	*/
	OrganizationNumber *string
	/*ShowProducts
	  Equals

	*/
	ShowProducts *bool
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string
	/*SupplierNumber
	  List of IDs

	*/
	SupplierNumber *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get supplier params
func (o *GetSupplierParams) WithTimeout(timeout time.Duration) *GetSupplierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get supplier params
func (o *GetSupplierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get supplier params
func (o *GetSupplierParams) WithContext(ctx context.Context) *GetSupplierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get supplier params
func (o *GetSupplierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get supplier params
func (o *GetSupplierParams) WithHTTPClient(client *http.Client) *GetSupplierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get supplier params
func (o *GetSupplierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountManagerID adds the accountManagerID to the get supplier params
func (o *GetSupplierParams) WithAccountManagerID(accountManagerID *string) *GetSupplierParams {
	o.SetAccountManagerID(accountManagerID)
	return o
}

// SetAccountManagerID adds the accountManagerId to the get supplier params
func (o *GetSupplierParams) SetAccountManagerID(accountManagerID *string) {
	o.AccountManagerID = accountManagerID
}

// WithChangedSince adds the changedSince to the get supplier params
func (o *GetSupplierParams) WithChangedSince(changedSince *string) *GetSupplierParams {
	o.SetChangedSince(changedSince)
	return o
}

// SetChangedSince adds the changedSince to the get supplier params
func (o *GetSupplierParams) SetChangedSince(changedSince *string) {
	o.ChangedSince = changedSince
}

// WithCount adds the count to the get supplier params
func (o *GetSupplierParams) WithCount(count *int64) *GetSupplierParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get supplier params
func (o *GetSupplierParams) SetCount(count *int64) {
	o.Count = count
}

// WithEmail adds the email to the get supplier params
func (o *GetSupplierParams) WithEmail(email *string) *GetSupplierParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the get supplier params
func (o *GetSupplierParams) SetEmail(email *string) {
	o.Email = email
}

// WithFields adds the fields to the get supplier params
func (o *GetSupplierParams) WithFields(fields *string) *GetSupplierParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get supplier params
func (o *GetSupplierParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get supplier params
func (o *GetSupplierParams) WithFrom(from *int64) *GetSupplierParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get supplier params
func (o *GetSupplierParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the get supplier params
func (o *GetSupplierParams) WithID(id *string) *GetSupplierParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get supplier params
func (o *GetSupplierParams) SetID(id *string) {
	o.ID = id
}

// WithInvoiceEmail adds the invoiceEmail to the get supplier params
func (o *GetSupplierParams) WithInvoiceEmail(invoiceEmail *string) *GetSupplierParams {
	o.SetInvoiceEmail(invoiceEmail)
	return o
}

// SetInvoiceEmail adds the invoiceEmail to the get supplier params
func (o *GetSupplierParams) SetInvoiceEmail(invoiceEmail *string) {
	o.InvoiceEmail = invoiceEmail
}

// WithIsInactive adds the isInactive to the get supplier params
func (o *GetSupplierParams) WithIsInactive(isInactive *bool) *GetSupplierParams {
	o.SetIsInactive(isInactive)
	return o
}

// SetIsInactive adds the isInactive to the get supplier params
func (o *GetSupplierParams) SetIsInactive(isInactive *bool) {
	o.IsInactive = isInactive
}

// WithIsWholesaler adds the isWholesaler to the get supplier params
func (o *GetSupplierParams) WithIsWholesaler(isWholesaler *bool) *GetSupplierParams {
	o.SetIsWholesaler(isWholesaler)
	return o
}

// SetIsWholesaler adds the isWholesaler to the get supplier params
func (o *GetSupplierParams) SetIsWholesaler(isWholesaler *bool) {
	o.IsWholesaler = isWholesaler
}

// WithOrganizationNumber adds the organizationNumber to the get supplier params
func (o *GetSupplierParams) WithOrganizationNumber(organizationNumber *string) *GetSupplierParams {
	o.SetOrganizationNumber(organizationNumber)
	return o
}

// SetOrganizationNumber adds the organizationNumber to the get supplier params
func (o *GetSupplierParams) SetOrganizationNumber(organizationNumber *string) {
	o.OrganizationNumber = organizationNumber
}

// WithShowProducts adds the showProducts to the get supplier params
func (o *GetSupplierParams) WithShowProducts(showProducts *bool) *GetSupplierParams {
	o.SetShowProducts(showProducts)
	return o
}

// SetShowProducts adds the showProducts to the get supplier params
func (o *GetSupplierParams) SetShowProducts(showProducts *bool) {
	o.ShowProducts = showProducts
}

// WithSorting adds the sorting to the get supplier params
func (o *GetSupplierParams) WithSorting(sorting *string) *GetSupplierParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get supplier params
func (o *GetSupplierParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithSupplierNumber adds the supplierNumber to the get supplier params
func (o *GetSupplierParams) WithSupplierNumber(supplierNumber *string) *GetSupplierParams {
	o.SetSupplierNumber(supplierNumber)
	return o
}

// SetSupplierNumber adds the supplierNumber to the get supplier params
func (o *GetSupplierParams) SetSupplierNumber(supplierNumber *string) {
	o.SupplierNumber = supplierNumber
}

// WriteToRequest writes these params to a swagger request
func (o *GetSupplierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountManagerID != nil {

		// query param accountManagerId
		var qrAccountManagerID string
		if o.AccountManagerID != nil {
			qrAccountManagerID = *o.AccountManagerID
		}
		qAccountManagerID := qrAccountManagerID
		if qAccountManagerID != "" {
			if err := r.SetQueryParam("accountManagerId", qAccountManagerID); err != nil {
				return err
			}
		}

	}

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

	if o.Email != nil {

		// query param email
		var qrEmail string
		if o.Email != nil {
			qrEmail = *o.Email
		}
		qEmail := qrEmail
		if qEmail != "" {
			if err := r.SetQueryParam("email", qEmail); err != nil {
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

	if o.InvoiceEmail != nil {

		// query param invoiceEmail
		var qrInvoiceEmail string
		if o.InvoiceEmail != nil {
			qrInvoiceEmail = *o.InvoiceEmail
		}
		qInvoiceEmail := qrInvoiceEmail
		if qInvoiceEmail != "" {
			if err := r.SetQueryParam("invoiceEmail", qInvoiceEmail); err != nil {
				return err
			}
		}

	}

	if o.IsInactive != nil {

		// query param isInactive
		var qrIsInactive bool
		if o.IsInactive != nil {
			qrIsInactive = *o.IsInactive
		}
		qIsInactive := swag.FormatBool(qrIsInactive)
		if qIsInactive != "" {
			if err := r.SetQueryParam("isInactive", qIsInactive); err != nil {
				return err
			}
		}

	}

	if o.IsWholesaler != nil {

		// query param isWholesaler
		var qrIsWholesaler bool
		if o.IsWholesaler != nil {
			qrIsWholesaler = *o.IsWholesaler
		}
		qIsWholesaler := swag.FormatBool(qrIsWholesaler)
		if qIsWholesaler != "" {
			if err := r.SetQueryParam("isWholesaler", qIsWholesaler); err != nil {
				return err
			}
		}

	}

	if o.OrganizationNumber != nil {

		// query param organizationNumber
		var qrOrganizationNumber string
		if o.OrganizationNumber != nil {
			qrOrganizationNumber = *o.OrganizationNumber
		}
		qOrganizationNumber := qrOrganizationNumber
		if qOrganizationNumber != "" {
			if err := r.SetQueryParam("organizationNumber", qOrganizationNumber); err != nil {
				return err
			}
		}

	}

	if o.ShowProducts != nil {

		// query param showProducts
		var qrShowProducts bool
		if o.ShowProducts != nil {
			qrShowProducts = *o.ShowProducts
		}
		qShowProducts := swag.FormatBool(qrShowProducts)
		if qShowProducts != "" {
			if err := r.SetQueryParam("showProducts", qShowProducts); err != nil {
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

	if o.SupplierNumber != nil {

		// query param supplierNumber
		var qrSupplierNumber string
		if o.SupplierNumber != nil {
			qrSupplierNumber = *o.SupplierNumber
		}
		qSupplierNumber := qrSupplierNumber
		if qSupplierNumber != "" {
			if err := r.SetQueryParam("supplierNumber", qSupplierNumber); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
