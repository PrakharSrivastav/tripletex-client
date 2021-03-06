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

// NewGetTimesheetEntryRecentActivitiesParams creates a new GetTimesheetEntryRecentActivitiesParams object
// with the default values initialized.
func NewGetTimesheetEntryRecentActivitiesParams() *GetTimesheetEntryRecentActivitiesParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTimesheetEntryRecentActivitiesParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTimesheetEntryRecentActivitiesParamsWithTimeout creates a new GetTimesheetEntryRecentActivitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTimesheetEntryRecentActivitiesParamsWithTimeout(timeout time.Duration) *GetTimesheetEntryRecentActivitiesParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTimesheetEntryRecentActivitiesParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetTimesheetEntryRecentActivitiesParamsWithContext creates a new GetTimesheetEntryRecentActivitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTimesheetEntryRecentActivitiesParamsWithContext(ctx context.Context) *GetTimesheetEntryRecentActivitiesParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTimesheetEntryRecentActivitiesParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetTimesheetEntryRecentActivitiesParamsWithHTTPClient creates a new GetTimesheetEntryRecentActivitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTimesheetEntryRecentActivitiesParamsWithHTTPClient(client *http.Client) *GetTimesheetEntryRecentActivitiesParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTimesheetEntryRecentActivitiesParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetTimesheetEntryRecentActivitiesParams contains all the parameters to send to the API endpoint
for the get timesheet entry recent activities operation typically these are written to a http.Request
*/
type GetTimesheetEntryRecentActivitiesParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*EmployeeID
	  ID of employee to find activities for. Defaults to ID of token owner.

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
	/*ProjectID
	  ID of project to find activities for

	*/
	ProjectID int32
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) WithTimeout(timeout time.Duration) *GetTimesheetEntryRecentActivitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) WithContext(ctx context.Context) *GetTimesheetEntryRecentActivitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) WithHTTPClient(client *http.Client) *GetTimesheetEntryRecentActivitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) WithCount(count *int64) *GetTimesheetEntryRecentActivitiesParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) SetCount(count *int64) {
	o.Count = count
}

// WithEmployeeID adds the employeeID to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) WithEmployeeID(employeeID *int32) *GetTimesheetEntryRecentActivitiesParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) SetEmployeeID(employeeID *int32) {
	o.EmployeeID = employeeID
}

// WithFields adds the fields to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) WithFields(fields *string) *GetTimesheetEntryRecentActivitiesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) WithFrom(from *int64) *GetTimesheetEntryRecentActivitiesParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) SetFrom(from *int64) {
	o.From = from
}

// WithProjectID adds the projectID to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) WithProjectID(projectID int32) *GetTimesheetEntryRecentActivitiesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) SetProjectID(projectID int32) {
	o.ProjectID = projectID
}

// WithSorting adds the sorting to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) WithSorting(sorting *string) *GetTimesheetEntryRecentActivitiesParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get timesheet entry recent activities params
func (o *GetTimesheetEntryRecentActivitiesParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetTimesheetEntryRecentActivitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param projectId
	qrProjectID := o.ProjectID
	qProjectID := swag.FormatInt32(qrProjectID)
	if qProjectID != "" {
		if err := r.SetQueryParam("projectId", qProjectID); err != nil {
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
