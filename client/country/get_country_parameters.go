// Code generated by go-swagger; DO NOT EDIT.

package country

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

// NewGetCountryParams creates a new GetCountryParams object
// with the default values initialized.
func NewGetCountryParams() *GetCountryParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCountryParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCountryParamsWithTimeout creates a new GetCountryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCountryParamsWithTimeout(timeout time.Duration) *GetCountryParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCountryParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetCountryParamsWithContext creates a new GetCountryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCountryParamsWithContext(ctx context.Context) *GetCountryParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCountryParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetCountryParamsWithHTTPClient creates a new GetCountryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCountryParamsWithHTTPClient(client *http.Client) *GetCountryParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCountryParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetCountryParams contains all the parameters to send to the API endpoint
for the get country operation typically these are written to a http.Request
*/
type GetCountryParams struct {

	/*Code
	  List of IDs

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

// WithTimeout adds the timeout to the get country params
func (o *GetCountryParams) WithTimeout(timeout time.Duration) *GetCountryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get country params
func (o *GetCountryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get country params
func (o *GetCountryParams) WithContext(ctx context.Context) *GetCountryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get country params
func (o *GetCountryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get country params
func (o *GetCountryParams) WithHTTPClient(client *http.Client) *GetCountryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get country params
func (o *GetCountryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCode adds the code to the get country params
func (o *GetCountryParams) WithCode(code *string) *GetCountryParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the get country params
func (o *GetCountryParams) SetCode(code *string) {
	o.Code = code
}

// WithCount adds the count to the get country params
func (o *GetCountryParams) WithCount(count *int64) *GetCountryParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get country params
func (o *GetCountryParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the get country params
func (o *GetCountryParams) WithFields(fields *string) *GetCountryParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get country params
func (o *GetCountryParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get country params
func (o *GetCountryParams) WithFrom(from *int64) *GetCountryParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get country params
func (o *GetCountryParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the get country params
func (o *GetCountryParams) WithID(id *string) *GetCountryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get country params
func (o *GetCountryParams) SetID(id *string) {
	o.ID = id
}

// WithSorting adds the sorting to the get country params
func (o *GetCountryParams) WithSorting(sorting *string) *GetCountryParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get country params
func (o *GetCountryParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetCountryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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