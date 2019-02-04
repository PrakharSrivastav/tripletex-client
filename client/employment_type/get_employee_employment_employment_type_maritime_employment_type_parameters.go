// Code generated by go-swagger; DO NOT EDIT.

package employment_type

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

// NewGetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams creates a new GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams object
// with the default values initialized.
func NewGetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams() *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParamsWithTimeout creates a new GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParamsWithTimeout(timeout time.Duration) *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParamsWithContext creates a new GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParamsWithContext(ctx context.Context) *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParamsWithHTTPClient creates a new GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParamsWithHTTPClient(client *http.Client) *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams contains all the parameters to send to the API endpoint
for the get employee employment employment type maritime employment type operation typically these are written to a http.Request
*/
type GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams struct {

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
	/*Type
	  maritimeEmploymentType

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) WithTimeout(timeout time.Duration) *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) WithContext(ctx context.Context) *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) WithHTTPClient(client *http.Client) *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) WithCount(count *int64) *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) WithFields(fields *string) *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) WithFrom(from *int64) *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) SetFrom(from *int64) {
	o.From = from
}

// WithSorting adds the sorting to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) WithSorting(sorting *string) *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithType adds the typeVar to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) WithType(typeVar *string) *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get employee employment employment type maritime employment type params
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
