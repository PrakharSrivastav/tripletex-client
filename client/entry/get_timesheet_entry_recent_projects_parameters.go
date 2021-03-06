// Code generated by go-swagger; DO NOT EDIT.

package entry

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

// NewGetTimesheetEntryRecentProjectsParams creates a new GetTimesheetEntryRecentProjectsParams object
// with the default values initialized.
func NewGetTimesheetEntryRecentProjectsParams() *GetTimesheetEntryRecentProjectsParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTimesheetEntryRecentProjectsParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTimesheetEntryRecentProjectsParamsWithTimeout creates a new GetTimesheetEntryRecentProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTimesheetEntryRecentProjectsParamsWithTimeout(timeout time.Duration) *GetTimesheetEntryRecentProjectsParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTimesheetEntryRecentProjectsParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetTimesheetEntryRecentProjectsParamsWithContext creates a new GetTimesheetEntryRecentProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTimesheetEntryRecentProjectsParamsWithContext(ctx context.Context) *GetTimesheetEntryRecentProjectsParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTimesheetEntryRecentProjectsParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetTimesheetEntryRecentProjectsParamsWithHTTPClient creates a new GetTimesheetEntryRecentProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTimesheetEntryRecentProjectsParamsWithHTTPClient(client *http.Client) *GetTimesheetEntryRecentProjectsParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTimesheetEntryRecentProjectsParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetTimesheetEntryRecentProjectsParams contains all the parameters to send to the API endpoint
for the get timesheet entry recent projects operation typically these are written to a http.Request
*/
type GetTimesheetEntryRecentProjectsParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*EmployeeID
	  ID of employee with recent project hours Defaults to ID of token owner.

	*/
	EmployeeID *int32
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

// WithTimeout adds the timeout to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) WithTimeout(timeout time.Duration) *GetTimesheetEntryRecentProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) WithContext(ctx context.Context) *GetTimesheetEntryRecentProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) WithHTTPClient(client *http.Client) *GetTimesheetEntryRecentProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) WithCount(count *int64) *GetTimesheetEntryRecentProjectsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) SetCount(count *int64) {
	o.Count = count
}

// WithEmployeeID adds the employeeID to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) WithEmployeeID(employeeID *int32) *GetTimesheetEntryRecentProjectsParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) SetEmployeeID(employeeID *int32) {
	o.EmployeeID = employeeID
}

// WithFields adds the fields to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) WithFields(fields *string) *GetTimesheetEntryRecentProjectsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) WithFrom(from *int64) *GetTimesheetEntryRecentProjectsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) SetFrom(from *int64) {
	o.From = from
}

// WithSorting adds the sorting to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) WithSorting(sorting *string) *GetTimesheetEntryRecentProjectsParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get timesheet entry recent projects params
func (o *GetTimesheetEntryRecentProjectsParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetTimesheetEntryRecentProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EmployeeID != nil {

		// query param employeeId
		var qrEmployeeID int32
		if o.EmployeeID != nil {
			qrEmployeeID = *o.EmployeeID
		}
		qEmployeeID := swag.FormatInt32(qrEmployeeID)
		if qEmployeeID != "" {
			if err := r.SetQueryParam("employeeId", qEmployeeID); err != nil {
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
