// Code generated by go-swagger; DO NOT EDIT.

package type_operations

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

// NewGetSalaryTypeParams creates a new GetSalaryTypeParams object
// with the default values initialized.
func NewGetSalaryTypeParams() *GetSalaryTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetSalaryTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSalaryTypeParamsWithTimeout creates a new GetSalaryTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSalaryTypeParamsWithTimeout(timeout time.Duration) *GetSalaryTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetSalaryTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetSalaryTypeParamsWithContext creates a new GetSalaryTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSalaryTypeParamsWithContext(ctx context.Context) *GetSalaryTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetSalaryTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetSalaryTypeParamsWithHTTPClient creates a new GetSalaryTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSalaryTypeParamsWithHTTPClient(client *http.Client) *GetSalaryTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetSalaryTypeParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetSalaryTypeParams contains all the parameters to send to the API endpoint
for the get salary type operation typically these are written to a http.Request
*/
type GetSalaryTypeParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*Description
	  Containing

	*/
	Description *string
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
	/*Name
	  Containing

	*/
	Name *string
	/*Number
	  Containing

	*/
	Number *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get salary type params
func (o *GetSalaryTypeParams) WithTimeout(timeout time.Duration) *GetSalaryTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get salary type params
func (o *GetSalaryTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get salary type params
func (o *GetSalaryTypeParams) WithContext(ctx context.Context) *GetSalaryTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get salary type params
func (o *GetSalaryTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get salary type params
func (o *GetSalaryTypeParams) WithHTTPClient(client *http.Client) *GetSalaryTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get salary type params
func (o *GetSalaryTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get salary type params
func (o *GetSalaryTypeParams) WithCount(count *int64) *GetSalaryTypeParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get salary type params
func (o *GetSalaryTypeParams) SetCount(count *int64) {
	o.Count = count
}

// WithDescription adds the description to the get salary type params
func (o *GetSalaryTypeParams) WithDescription(description *string) *GetSalaryTypeParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the get salary type params
func (o *GetSalaryTypeParams) SetDescription(description *string) {
	o.Description = description
}

// WithFields adds the fields to the get salary type params
func (o *GetSalaryTypeParams) WithFields(fields *string) *GetSalaryTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get salary type params
func (o *GetSalaryTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get salary type params
func (o *GetSalaryTypeParams) WithFrom(from *int64) *GetSalaryTypeParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get salary type params
func (o *GetSalaryTypeParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the get salary type params
func (o *GetSalaryTypeParams) WithID(id *string) *GetSalaryTypeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get salary type params
func (o *GetSalaryTypeParams) SetID(id *string) {
	o.ID = id
}

// WithName adds the name to the get salary type params
func (o *GetSalaryTypeParams) WithName(name *string) *GetSalaryTypeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get salary type params
func (o *GetSalaryTypeParams) SetName(name *string) {
	o.Name = name
}

// WithNumber adds the number to the get salary type params
func (o *GetSalaryTypeParams) WithNumber(number *string) *GetSalaryTypeParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get salary type params
func (o *GetSalaryTypeParams) SetNumber(number *string) {
	o.Number = number
}

// WithSorting adds the sorting to the get salary type params
func (o *GetSalaryTypeParams) WithSorting(sorting *string) *GetSalaryTypeParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get salary type params
func (o *GetSalaryTypeParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetSalaryTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Description != nil {

		// query param description
		var qrDescription string
		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {
			if err := r.SetQueryParam("description", qDescription); err != nil {
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

	if o.Number != nil {

		// query param number
		var qrNumber string
		if o.Number != nil {
			qrNumber = *o.Number
		}
		qNumber := qrNumber
		if qNumber != "" {
			if err := r.SetQueryParam("number", qNumber); err != nil {
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
