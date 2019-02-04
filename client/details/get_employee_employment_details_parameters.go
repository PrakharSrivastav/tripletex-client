// Code generated by go-swagger; DO NOT EDIT.

package details

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

// NewGetEmployeeEmploymentDetailsParams creates a new GetEmployeeEmploymentDetailsParams object
// with the default values initialized.
func NewGetEmployeeEmploymentDetailsParams() *GetEmployeeEmploymentDetailsParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetEmployeeEmploymentDetailsParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmployeeEmploymentDetailsParamsWithTimeout creates a new GetEmployeeEmploymentDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmployeeEmploymentDetailsParamsWithTimeout(timeout time.Duration) *GetEmployeeEmploymentDetailsParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetEmployeeEmploymentDetailsParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetEmployeeEmploymentDetailsParamsWithContext creates a new GetEmployeeEmploymentDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmployeeEmploymentDetailsParamsWithContext(ctx context.Context) *GetEmployeeEmploymentDetailsParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetEmployeeEmploymentDetailsParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetEmployeeEmploymentDetailsParamsWithHTTPClient creates a new GetEmployeeEmploymentDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmployeeEmploymentDetailsParamsWithHTTPClient(client *http.Client) *GetEmployeeEmploymentDetailsParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetEmployeeEmploymentDetailsParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetEmployeeEmploymentDetailsParams contains all the parameters to send to the API endpoint
for the get employee employment details operation typically these are written to a http.Request
*/
type GetEmployeeEmploymentDetailsParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*EmploymentID
	  Element ID

	*/
	EmploymentID *int32
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

// WithTimeout adds the timeout to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) WithTimeout(timeout time.Duration) *GetEmployeeEmploymentDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) WithContext(ctx context.Context) *GetEmployeeEmploymentDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) WithHTTPClient(client *http.Client) *GetEmployeeEmploymentDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) WithCount(count *int64) *GetEmployeeEmploymentDetailsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) SetCount(count *int64) {
	o.Count = count
}

// WithEmploymentID adds the employmentID to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) WithEmploymentID(employmentID *int32) *GetEmployeeEmploymentDetailsParams {
	o.SetEmploymentID(employmentID)
	return o
}

// SetEmploymentID adds the employmentId to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) SetEmploymentID(employmentID *int32) {
	o.EmploymentID = employmentID
}

// WithFields adds the fields to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) WithFields(fields *string) *GetEmployeeEmploymentDetailsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) WithFrom(from *int64) *GetEmployeeEmploymentDetailsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) SetFrom(from *int64) {
	o.From = from
}

// WithSorting adds the sorting to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) WithSorting(sorting *string) *GetEmployeeEmploymentDetailsParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get employee employment details params
func (o *GetEmployeeEmploymentDetailsParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmployeeEmploymentDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EmploymentID != nil {

		// query param employmentId
		var qrEmploymentID int32
		if o.EmploymentID != nil {
			qrEmploymentID = *o.EmploymentID
		}
		qEmploymentID := swag.FormatInt32(qrEmploymentID)
		if qEmploymentID != "" {
			if err := r.SetQueryParam("employmentId", qEmploymentID); err != nil {
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