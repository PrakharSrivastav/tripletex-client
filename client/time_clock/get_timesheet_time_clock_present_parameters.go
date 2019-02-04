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

// NewGetTimesheetTimeClockPresentParams creates a new GetTimesheetTimeClockPresentParams object
// with the default values initialized.
func NewGetTimesheetTimeClockPresentParams() *GetTimesheetTimeClockPresentParams {
	var ()
	return &GetTimesheetTimeClockPresentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTimesheetTimeClockPresentParamsWithTimeout creates a new GetTimesheetTimeClockPresentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTimesheetTimeClockPresentParamsWithTimeout(timeout time.Duration) *GetTimesheetTimeClockPresentParams {
	var ()
	return &GetTimesheetTimeClockPresentParams{

		timeout: timeout,
	}
}

// NewGetTimesheetTimeClockPresentParamsWithContext creates a new GetTimesheetTimeClockPresentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTimesheetTimeClockPresentParamsWithContext(ctx context.Context) *GetTimesheetTimeClockPresentParams {
	var ()
	return &GetTimesheetTimeClockPresentParams{

		Context: ctx,
	}
}

// NewGetTimesheetTimeClockPresentParamsWithHTTPClient creates a new GetTimesheetTimeClockPresentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTimesheetTimeClockPresentParamsWithHTTPClient(client *http.Client) *GetTimesheetTimeClockPresentParams {
	var ()
	return &GetTimesheetTimeClockPresentParams{
		HTTPClient: client,
	}
}

/*GetTimesheetTimeClockPresentParams contains all the parameters to send to the API endpoint
for the get timesheet time clock present operation typically these are written to a http.Request
*/
type GetTimesheetTimeClockPresentParams struct {

	/*EmployeeID
	  Employee ID. Defaults to ID of token owner.

	*/
	EmployeeID *int32
	/*Fields
	  Fields filter pattern

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get timesheet time clock present params
func (o *GetTimesheetTimeClockPresentParams) WithTimeout(timeout time.Duration) *GetTimesheetTimeClockPresentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get timesheet time clock present params
func (o *GetTimesheetTimeClockPresentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get timesheet time clock present params
func (o *GetTimesheetTimeClockPresentParams) WithContext(ctx context.Context) *GetTimesheetTimeClockPresentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get timesheet time clock present params
func (o *GetTimesheetTimeClockPresentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get timesheet time clock present params
func (o *GetTimesheetTimeClockPresentParams) WithHTTPClient(client *http.Client) *GetTimesheetTimeClockPresentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get timesheet time clock present params
func (o *GetTimesheetTimeClockPresentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmployeeID adds the employeeID to the get timesheet time clock present params
func (o *GetTimesheetTimeClockPresentParams) WithEmployeeID(employeeID *int32) *GetTimesheetTimeClockPresentParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the get timesheet time clock present params
func (o *GetTimesheetTimeClockPresentParams) SetEmployeeID(employeeID *int32) {
	o.EmployeeID = employeeID
}

// WithFields adds the fields to the get timesheet time clock present params
func (o *GetTimesheetTimeClockPresentParams) WithFields(fields *string) *GetTimesheetTimeClockPresentParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get timesheet time clock present params
func (o *GetTimesheetTimeClockPresentParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetTimesheetTimeClockPresentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}