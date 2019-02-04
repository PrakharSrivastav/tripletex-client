// Code generated by go-swagger; DO NOT EDIT.

package leave_of_absence_type

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

// NewGetEmployeeEmploymentLeaveOfAbsenceTypeParams creates a new GetEmployeeEmploymentLeaveOfAbsenceTypeParams object
// with the default values initialized.
func NewGetEmployeeEmploymentLeaveOfAbsenceTypeParams() *GetEmployeeEmploymentLeaveOfAbsenceTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetEmployeeEmploymentLeaveOfAbsenceTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmployeeEmploymentLeaveOfAbsenceTypeParamsWithTimeout creates a new GetEmployeeEmploymentLeaveOfAbsenceTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmployeeEmploymentLeaveOfAbsenceTypeParamsWithTimeout(timeout time.Duration) *GetEmployeeEmploymentLeaveOfAbsenceTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetEmployeeEmploymentLeaveOfAbsenceTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetEmployeeEmploymentLeaveOfAbsenceTypeParamsWithContext creates a new GetEmployeeEmploymentLeaveOfAbsenceTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmployeeEmploymentLeaveOfAbsenceTypeParamsWithContext(ctx context.Context) *GetEmployeeEmploymentLeaveOfAbsenceTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetEmployeeEmploymentLeaveOfAbsenceTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetEmployeeEmploymentLeaveOfAbsenceTypeParamsWithHTTPClient creates a new GetEmployeeEmploymentLeaveOfAbsenceTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmployeeEmploymentLeaveOfAbsenceTypeParamsWithHTTPClient(client *http.Client) *GetEmployeeEmploymentLeaveOfAbsenceTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetEmployeeEmploymentLeaveOfAbsenceTypeParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetEmployeeEmploymentLeaveOfAbsenceTypeParams contains all the parameters to send to the API endpoint
for the get employee employment leave of absence type operation typically these are written to a http.Request
*/
type GetEmployeeEmploymentLeaveOfAbsenceTypeParams struct {

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

// WithTimeout adds the timeout to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) WithTimeout(timeout time.Duration) *GetEmployeeEmploymentLeaveOfAbsenceTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) WithContext(ctx context.Context) *GetEmployeeEmploymentLeaveOfAbsenceTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) WithHTTPClient(client *http.Client) *GetEmployeeEmploymentLeaveOfAbsenceTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) WithCount(count *int64) *GetEmployeeEmploymentLeaveOfAbsenceTypeParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) WithFields(fields *string) *GetEmployeeEmploymentLeaveOfAbsenceTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) WithFrom(from *int64) *GetEmployeeEmploymentLeaveOfAbsenceTypeParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) SetFrom(from *int64) {
	o.From = from
}

// WithSorting adds the sorting to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) WithSorting(sorting *string) *GetEmployeeEmploymentLeaveOfAbsenceTypeParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get employee employment leave of absence type params
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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