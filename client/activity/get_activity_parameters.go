// Code generated by go-swagger; DO NOT EDIT.

package activity

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

// NewGetActivityParams creates a new GetActivityParams object
// with the default values initialized.
func NewGetActivityParams() *GetActivityParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetActivityParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetActivityParamsWithTimeout creates a new GetActivityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetActivityParamsWithTimeout(timeout time.Duration) *GetActivityParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetActivityParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetActivityParamsWithContext creates a new GetActivityParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetActivityParamsWithContext(ctx context.Context) *GetActivityParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetActivityParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetActivityParamsWithHTTPClient creates a new GetActivityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetActivityParamsWithHTTPClient(client *http.Client) *GetActivityParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetActivityParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetActivityParams contains all the parameters to send to the API endpoint
for the get activity operation typically these are written to a http.Request
*/
type GetActivityParams struct {

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
	/*IsChargeable
	  Equals

	*/
	IsChargeable *bool
	/*IsGeneral
	  Equals

	*/
	IsGeneral *bool
	/*IsProjectActivity
	  Equals

	*/
	IsProjectActivity *bool
	/*IsTask
	  Equals

	*/
	IsTask *bool
	/*Name
	  Containing

	*/
	Name *string
	/*Number
	  Equals

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

// WithTimeout adds the timeout to the get activity params
func (o *GetActivityParams) WithTimeout(timeout time.Duration) *GetActivityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get activity params
func (o *GetActivityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get activity params
func (o *GetActivityParams) WithContext(ctx context.Context) *GetActivityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get activity params
func (o *GetActivityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get activity params
func (o *GetActivityParams) WithHTTPClient(client *http.Client) *GetActivityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get activity params
func (o *GetActivityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get activity params
func (o *GetActivityParams) WithCount(count *int64) *GetActivityParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get activity params
func (o *GetActivityParams) SetCount(count *int64) {
	o.Count = count
}

// WithDescription adds the description to the get activity params
func (o *GetActivityParams) WithDescription(description *string) *GetActivityParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the get activity params
func (o *GetActivityParams) SetDescription(description *string) {
	o.Description = description
}

// WithFields adds the fields to the get activity params
func (o *GetActivityParams) WithFields(fields *string) *GetActivityParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get activity params
func (o *GetActivityParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get activity params
func (o *GetActivityParams) WithFrom(from *int64) *GetActivityParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get activity params
func (o *GetActivityParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the get activity params
func (o *GetActivityParams) WithID(id *string) *GetActivityParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get activity params
func (o *GetActivityParams) SetID(id *string) {
	o.ID = id
}

// WithIsChargeable adds the isChargeable to the get activity params
func (o *GetActivityParams) WithIsChargeable(isChargeable *bool) *GetActivityParams {
	o.SetIsChargeable(isChargeable)
	return o
}

// SetIsChargeable adds the isChargeable to the get activity params
func (o *GetActivityParams) SetIsChargeable(isChargeable *bool) {
	o.IsChargeable = isChargeable
}

// WithIsGeneral adds the isGeneral to the get activity params
func (o *GetActivityParams) WithIsGeneral(isGeneral *bool) *GetActivityParams {
	o.SetIsGeneral(isGeneral)
	return o
}

// SetIsGeneral adds the isGeneral to the get activity params
func (o *GetActivityParams) SetIsGeneral(isGeneral *bool) {
	o.IsGeneral = isGeneral
}

// WithIsProjectActivity adds the isProjectActivity to the get activity params
func (o *GetActivityParams) WithIsProjectActivity(isProjectActivity *bool) *GetActivityParams {
	o.SetIsProjectActivity(isProjectActivity)
	return o
}

// SetIsProjectActivity adds the isProjectActivity to the get activity params
func (o *GetActivityParams) SetIsProjectActivity(isProjectActivity *bool) {
	o.IsProjectActivity = isProjectActivity
}

// WithIsTask adds the isTask to the get activity params
func (o *GetActivityParams) WithIsTask(isTask *bool) *GetActivityParams {
	o.SetIsTask(isTask)
	return o
}

// SetIsTask adds the isTask to the get activity params
func (o *GetActivityParams) SetIsTask(isTask *bool) {
	o.IsTask = isTask
}

// WithName adds the name to the get activity params
func (o *GetActivityParams) WithName(name *string) *GetActivityParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get activity params
func (o *GetActivityParams) SetName(name *string) {
	o.Name = name
}

// WithNumber adds the number to the get activity params
func (o *GetActivityParams) WithNumber(number *string) *GetActivityParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get activity params
func (o *GetActivityParams) SetNumber(number *string) {
	o.Number = number
}

// WithSorting adds the sorting to the get activity params
func (o *GetActivityParams) WithSorting(sorting *string) *GetActivityParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get activity params
func (o *GetActivityParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetActivityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IsChargeable != nil {

		// query param isChargeable
		var qrIsChargeable bool
		if o.IsChargeable != nil {
			qrIsChargeable = *o.IsChargeable
		}
		qIsChargeable := swag.FormatBool(qrIsChargeable)
		if qIsChargeable != "" {
			if err := r.SetQueryParam("isChargeable", qIsChargeable); err != nil {
				return err
			}
		}

	}

	if o.IsGeneral != nil {

		// query param isGeneral
		var qrIsGeneral bool
		if o.IsGeneral != nil {
			qrIsGeneral = *o.IsGeneral
		}
		qIsGeneral := swag.FormatBool(qrIsGeneral)
		if qIsGeneral != "" {
			if err := r.SetQueryParam("isGeneral", qIsGeneral); err != nil {
				return err
			}
		}

	}

	if o.IsProjectActivity != nil {

		// query param isProjectActivity
		var qrIsProjectActivity bool
		if o.IsProjectActivity != nil {
			qrIsProjectActivity = *o.IsProjectActivity
		}
		qIsProjectActivity := swag.FormatBool(qrIsProjectActivity)
		if qIsProjectActivity != "" {
			if err := r.SetQueryParam("isProjectActivity", qIsProjectActivity); err != nil {
				return err
			}
		}

	}

	if o.IsTask != nil {

		// query param isTask
		var qrIsTask bool
		if o.IsTask != nil {
			qrIsTask = *o.IsTask
		}
		qIsTask := swag.FormatBool(qrIsTask)
		if qIsTask != "" {
			if err := r.SetQueryParam("isTask", qIsTask); err != nil {
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
