// Code generated by go-swagger; DO NOT EDIT.

package time_clock

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

// NewGetTimesheetTimeClockParams creates a new GetTimesheetTimeClockParams object
// with the default values initialized.
func NewGetTimesheetTimeClockParams() *GetTimesheetTimeClockParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTimesheetTimeClockParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTimesheetTimeClockParamsWithTimeout creates a new GetTimesheetTimeClockParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTimesheetTimeClockParamsWithTimeout(timeout time.Duration) *GetTimesheetTimeClockParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTimesheetTimeClockParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetTimesheetTimeClockParamsWithContext creates a new GetTimesheetTimeClockParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTimesheetTimeClockParamsWithContext(ctx context.Context) *GetTimesheetTimeClockParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTimesheetTimeClockParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetTimesheetTimeClockParamsWithHTTPClient creates a new GetTimesheetTimeClockParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTimesheetTimeClockParamsWithHTTPClient(client *http.Client) *GetTimesheetTimeClockParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTimesheetTimeClockParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetTimesheetTimeClockParams contains all the parameters to send to the API endpoint
for the get timesheet time clock operation typically these are written to a http.Request
*/
type GetTimesheetTimeClockParams struct {

	/*ActivityID
	  List of IDs

	*/
	ActivityID *string
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
	/*EmployeeID
	  List of IDs

	*/
	EmployeeID *string
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*HourID
	  List of IDs

	*/
	HourID *string
	/*ID
	  List of IDs

	*/
	ID *string
	/*IsRunning
	  Equals

	*/
	IsRunning *bool
	/*ProjectID
	  List of IDs

	*/
	ProjectID *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithTimeout(timeout time.Duration) *GetTimesheetTimeClockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithContext(ctx context.Context) *GetTimesheetTimeClockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithHTTPClient(client *http.Client) *GetTimesheetTimeClockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActivityID adds the activityID to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithActivityID(activityID *string) *GetTimesheetTimeClockParams {
	o.SetActivityID(activityID)
	return o
}

// SetActivityID adds the activityId to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetActivityID(activityID *string) {
	o.ActivityID = activityID
}

// WithCount adds the count to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithCount(count *int64) *GetTimesheetTimeClockParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetCount(count *int64) {
	o.Count = count
}

// WithDateFrom adds the dateFrom to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithDateFrom(dateFrom *string) *GetTimesheetTimeClockParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetDateFrom(dateFrom *string) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithDateTo(dateTo *string) *GetTimesheetTimeClockParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetDateTo(dateTo *string) {
	o.DateTo = dateTo
}

// WithEmployeeID adds the employeeID to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithEmployeeID(employeeID *string) *GetTimesheetTimeClockParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetEmployeeID(employeeID *string) {
	o.EmployeeID = employeeID
}

// WithFields adds the fields to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithFields(fields *string) *GetTimesheetTimeClockParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithFrom(from *int64) *GetTimesheetTimeClockParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetFrom(from *int64) {
	o.From = from
}

// WithHourID adds the hourID to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithHourID(hourID *string) *GetTimesheetTimeClockParams {
	o.SetHourID(hourID)
	return o
}

// SetHourID adds the hourId to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetHourID(hourID *string) {
	o.HourID = hourID
}

// WithID adds the id to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithID(id *string) *GetTimesheetTimeClockParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetID(id *string) {
	o.ID = id
}

// WithIsRunning adds the isRunning to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithIsRunning(isRunning *bool) *GetTimesheetTimeClockParams {
	o.SetIsRunning(isRunning)
	return o
}

// SetIsRunning adds the isRunning to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetIsRunning(isRunning *bool) {
	o.IsRunning = isRunning
}

// WithProjectID adds the projectID to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithProjectID(projectID *string) *GetTimesheetTimeClockParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithSorting adds the sorting to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) WithSorting(sorting *string) *GetTimesheetTimeClockParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get timesheet time clock params
func (o *GetTimesheetTimeClockParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetTimesheetTimeClockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActivityID != nil {

		// query param activityId
		var qrActivityID string
		if o.ActivityID != nil {
			qrActivityID = *o.ActivityID
		}
		qActivityID := qrActivityID
		if qActivityID != "" {
			if err := r.SetQueryParam("activityId", qActivityID); err != nil {
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

	if o.EmployeeID != nil {

		// query param employeeId
		var qrEmployeeID string
		if o.EmployeeID != nil {
			qrEmployeeID = *o.EmployeeID
		}
		qEmployeeID := qrEmployeeID
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

	if o.HourID != nil {

		// query param hourId
		var qrHourID string
		if o.HourID != nil {
			qrHourID = *o.HourID
		}
		qHourID := qrHourID
		if qHourID != "" {
			if err := r.SetQueryParam("hourId", qHourID); err != nil {
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

	if o.IsRunning != nil {

		// query param isRunning
		var qrIsRunning bool
		if o.IsRunning != nil {
			qrIsRunning = *o.IsRunning
		}
		qIsRunning := swag.FormatBool(qrIsRunning)
		if qIsRunning != "" {
			if err := r.SetQueryParam("isRunning", qIsRunning); err != nil {
				return err
			}
		}

	}

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID string
		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {
			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
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
