// Code generated by go-swagger; DO NOT EDIT.

package external

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

// NewGetProductExternalParams creates a new GetProductExternalParams object
// with the default values initialized.
func NewGetProductExternalParams() *GetProductExternalParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetProductExternalParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductExternalParamsWithTimeout creates a new GetProductExternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductExternalParamsWithTimeout(timeout time.Duration) *GetProductExternalParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetProductExternalParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetProductExternalParamsWithContext creates a new GetProductExternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProductExternalParamsWithContext(ctx context.Context) *GetProductExternalParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetProductExternalParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetProductExternalParamsWithHTTPClient creates a new GetProductExternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProductExternalParamsWithHTTPClient(client *http.Client) *GetProductExternalParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetProductExternalParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetProductExternalParams contains all the parameters to send to the API endpoint
for the get product external operation typically these are written to a http.Request
*/
type GetProductExternalParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*ElNumber
	  List of valid el numbers

	*/
	ElNumber *string
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*IsInactive
	  Equals

	*/
	IsInactive *bool
	/*Name
	  Containing

	*/
	Name *string
	/*NrfNumber
	  List of valid nrf numbers

	*/
	NrfNumber *string
	/*OrganizationNumber
	  Wholesaler organization number. Mandatory if Wholesaler is not selected. If Wholesaler is selected, this field is ignored.

	*/
	OrganizationNumber *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string
	/*Wholesaler
	  Wholesaler

	*/
	Wholesaler *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get product external params
func (o *GetProductExternalParams) WithTimeout(timeout time.Duration) *GetProductExternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get product external params
func (o *GetProductExternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get product external params
func (o *GetProductExternalParams) WithContext(ctx context.Context) *GetProductExternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get product external params
func (o *GetProductExternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get product external params
func (o *GetProductExternalParams) WithHTTPClient(client *http.Client) *GetProductExternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get product external params
func (o *GetProductExternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get product external params
func (o *GetProductExternalParams) WithCount(count *int64) *GetProductExternalParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get product external params
func (o *GetProductExternalParams) SetCount(count *int64) {
	o.Count = count
}

// WithElNumber adds the elNumber to the get product external params
func (o *GetProductExternalParams) WithElNumber(elNumber *string) *GetProductExternalParams {
	o.SetElNumber(elNumber)
	return o
}

// SetElNumber adds the elNumber to the get product external params
func (o *GetProductExternalParams) SetElNumber(elNumber *string) {
	o.ElNumber = elNumber
}

// WithFields adds the fields to the get product external params
func (o *GetProductExternalParams) WithFields(fields *string) *GetProductExternalParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get product external params
func (o *GetProductExternalParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get product external params
func (o *GetProductExternalParams) WithFrom(from *int64) *GetProductExternalParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get product external params
func (o *GetProductExternalParams) SetFrom(from *int64) {
	o.From = from
}

// WithIsInactive adds the isInactive to the get product external params
func (o *GetProductExternalParams) WithIsInactive(isInactive *bool) *GetProductExternalParams {
	o.SetIsInactive(isInactive)
	return o
}

// SetIsInactive adds the isInactive to the get product external params
func (o *GetProductExternalParams) SetIsInactive(isInactive *bool) {
	o.IsInactive = isInactive
}

// WithName adds the name to the get product external params
func (o *GetProductExternalParams) WithName(name *string) *GetProductExternalParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get product external params
func (o *GetProductExternalParams) SetName(name *string) {
	o.Name = name
}

// WithNrfNumber adds the nrfNumber to the get product external params
func (o *GetProductExternalParams) WithNrfNumber(nrfNumber *string) *GetProductExternalParams {
	o.SetNrfNumber(nrfNumber)
	return o
}

// SetNrfNumber adds the nrfNumber to the get product external params
func (o *GetProductExternalParams) SetNrfNumber(nrfNumber *string) {
	o.NrfNumber = nrfNumber
}

// WithOrganizationNumber adds the organizationNumber to the get product external params
func (o *GetProductExternalParams) WithOrganizationNumber(organizationNumber *string) *GetProductExternalParams {
	o.SetOrganizationNumber(organizationNumber)
	return o
}

// SetOrganizationNumber adds the organizationNumber to the get product external params
func (o *GetProductExternalParams) SetOrganizationNumber(organizationNumber *string) {
	o.OrganizationNumber = organizationNumber
}

// WithSorting adds the sorting to the get product external params
func (o *GetProductExternalParams) WithSorting(sorting *string) *GetProductExternalParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get product external params
func (o *GetProductExternalParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithWholesaler adds the wholesaler to the get product external params
func (o *GetProductExternalParams) WithWholesaler(wholesaler *string) *GetProductExternalParams {
	o.SetWholesaler(wholesaler)
	return o
}

// SetWholesaler adds the wholesaler to the get product external params
func (o *GetProductExternalParams) SetWholesaler(wholesaler *string) {
	o.Wholesaler = wholesaler
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductExternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ElNumber != nil {

		// query param elNumber
		var qrElNumber string
		if o.ElNumber != nil {
			qrElNumber = *o.ElNumber
		}
		qElNumber := qrElNumber
		if qElNumber != "" {
			if err := r.SetQueryParam("elNumber", qElNumber); err != nil {
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

	if o.NrfNumber != nil {

		// query param nrfNumber
		var qrNrfNumber string
		if o.NrfNumber != nil {
			qrNrfNumber = *o.NrfNumber
		}
		qNrfNumber := qrNrfNumber
		if qNrfNumber != "" {
			if err := r.SetQueryParam("nrfNumber", qNrfNumber); err != nil {
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

	if o.Wholesaler != nil {

		// query param wholesaler
		var qrWholesaler string
		if o.Wholesaler != nil {
			qrWholesaler = *o.Wholesaler
		}
		qWholesaler := qrWholesaler
		if qWholesaler != "" {
			if err := r.SetQueryParam("wholesaler", qWholesaler); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
