// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewGetProjectParams creates a new GetProjectParams object
// with the default values initialized.
func NewGetProjectParams() *GetProjectParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetProjectParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectParamsWithTimeout creates a new GetProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProjectParamsWithTimeout(timeout time.Duration) *GetProjectParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetProjectParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetProjectParamsWithContext creates a new GetProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProjectParamsWithContext(ctx context.Context) *GetProjectParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetProjectParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetProjectParamsWithHTTPClient creates a new GetProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProjectParamsWithHTTPClient(client *http.Client) *GetProjectParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetProjectParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetProjectParams contains all the parameters to send to the API endpoint
for the get project operation typically these are written to a http.Request
*/
type GetProjectParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*CustomerID
	  Equals

	*/
	CustomerID *string
	/*DepartmentID
	  List of IDs

	*/
	DepartmentID *string
	/*EmployeeInProjectID
	  List of IDs

	*/
	EmployeeInProjectID *string
	/*EndDateFrom
	  From and including

	*/
	EndDateFrom *string
	/*EndDateTo
	  To and excluding

	*/
	EndDateTo *string
	/*ExternalAccountsNumber
	  Containing

	*/
	ExternalAccountsNumber *string
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
	/*IsClosed
	  Equals

	*/
	IsClosed *bool
	/*IsOffer
	  Equals

	*/
	IsOffer *bool
	/*Name
	  Containing

	*/
	Name *string
	/*Number
	  Equals

	*/
	Number *string
	/*ProjectManagerID
	  List of IDs

	*/
	ProjectManagerID *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string
	/*StartDateFrom
	  From and including

	*/
	StartDateFrom *string
	/*StartDateTo
	  To and excluding

	*/
	StartDateTo *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get project params
func (o *GetProjectParams) WithTimeout(timeout time.Duration) *GetProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project params
func (o *GetProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project params
func (o *GetProjectParams) WithContext(ctx context.Context) *GetProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project params
func (o *GetProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project params
func (o *GetProjectParams) WithHTTPClient(client *http.Client) *GetProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project params
func (o *GetProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get project params
func (o *GetProjectParams) WithCount(count *int64) *GetProjectParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get project params
func (o *GetProjectParams) SetCount(count *int64) {
	o.Count = count
}

// WithCustomerID adds the customerID to the get project params
func (o *GetProjectParams) WithCustomerID(customerID *string) *GetProjectParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get project params
func (o *GetProjectParams) SetCustomerID(customerID *string) {
	o.CustomerID = customerID
}

// WithDepartmentID adds the departmentID to the get project params
func (o *GetProjectParams) WithDepartmentID(departmentID *string) *GetProjectParams {
	o.SetDepartmentID(departmentID)
	return o
}

// SetDepartmentID adds the departmentId to the get project params
func (o *GetProjectParams) SetDepartmentID(departmentID *string) {
	o.DepartmentID = departmentID
}

// WithEmployeeInProjectID adds the employeeInProjectID to the get project params
func (o *GetProjectParams) WithEmployeeInProjectID(employeeInProjectID *string) *GetProjectParams {
	o.SetEmployeeInProjectID(employeeInProjectID)
	return o
}

// SetEmployeeInProjectID adds the employeeInProjectId to the get project params
func (o *GetProjectParams) SetEmployeeInProjectID(employeeInProjectID *string) {
	o.EmployeeInProjectID = employeeInProjectID
}

// WithEndDateFrom adds the endDateFrom to the get project params
func (o *GetProjectParams) WithEndDateFrom(endDateFrom *string) *GetProjectParams {
	o.SetEndDateFrom(endDateFrom)
	return o
}

// SetEndDateFrom adds the endDateFrom to the get project params
func (o *GetProjectParams) SetEndDateFrom(endDateFrom *string) {
	o.EndDateFrom = endDateFrom
}

// WithEndDateTo adds the endDateTo to the get project params
func (o *GetProjectParams) WithEndDateTo(endDateTo *string) *GetProjectParams {
	o.SetEndDateTo(endDateTo)
	return o
}

// SetEndDateTo adds the endDateTo to the get project params
func (o *GetProjectParams) SetEndDateTo(endDateTo *string) {
	o.EndDateTo = endDateTo
}

// WithExternalAccountsNumber adds the externalAccountsNumber to the get project params
func (o *GetProjectParams) WithExternalAccountsNumber(externalAccountsNumber *string) *GetProjectParams {
	o.SetExternalAccountsNumber(externalAccountsNumber)
	return o
}

// SetExternalAccountsNumber adds the externalAccountsNumber to the get project params
func (o *GetProjectParams) SetExternalAccountsNumber(externalAccountsNumber *string) {
	o.ExternalAccountsNumber = externalAccountsNumber
}

// WithFields adds the fields to the get project params
func (o *GetProjectParams) WithFields(fields *string) *GetProjectParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get project params
func (o *GetProjectParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get project params
func (o *GetProjectParams) WithFrom(from *int64) *GetProjectParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get project params
func (o *GetProjectParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the get project params
func (o *GetProjectParams) WithID(id *string) *GetProjectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get project params
func (o *GetProjectParams) SetID(id *string) {
	o.ID = id
}

// WithIsClosed adds the isClosed to the get project params
func (o *GetProjectParams) WithIsClosed(isClosed *bool) *GetProjectParams {
	o.SetIsClosed(isClosed)
	return o
}

// SetIsClosed adds the isClosed to the get project params
func (o *GetProjectParams) SetIsClosed(isClosed *bool) {
	o.IsClosed = isClosed
}

// WithIsOffer adds the isOffer to the get project params
func (o *GetProjectParams) WithIsOffer(isOffer *bool) *GetProjectParams {
	o.SetIsOffer(isOffer)
	return o
}

// SetIsOffer adds the isOffer to the get project params
func (o *GetProjectParams) SetIsOffer(isOffer *bool) {
	o.IsOffer = isOffer
}

// WithName adds the name to the get project params
func (o *GetProjectParams) WithName(name *string) *GetProjectParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get project params
func (o *GetProjectParams) SetName(name *string) {
	o.Name = name
}

// WithNumber adds the number to the get project params
func (o *GetProjectParams) WithNumber(number *string) *GetProjectParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get project params
func (o *GetProjectParams) SetNumber(number *string) {
	o.Number = number
}

// WithProjectManagerID adds the projectManagerID to the get project params
func (o *GetProjectParams) WithProjectManagerID(projectManagerID *string) *GetProjectParams {
	o.SetProjectManagerID(projectManagerID)
	return o
}

// SetProjectManagerID adds the projectManagerId to the get project params
func (o *GetProjectParams) SetProjectManagerID(projectManagerID *string) {
	o.ProjectManagerID = projectManagerID
}

// WithSorting adds the sorting to the get project params
func (o *GetProjectParams) WithSorting(sorting *string) *GetProjectParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get project params
func (o *GetProjectParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithStartDateFrom adds the startDateFrom to the get project params
func (o *GetProjectParams) WithStartDateFrom(startDateFrom *string) *GetProjectParams {
	o.SetStartDateFrom(startDateFrom)
	return o
}

// SetStartDateFrom adds the startDateFrom to the get project params
func (o *GetProjectParams) SetStartDateFrom(startDateFrom *string) {
	o.StartDateFrom = startDateFrom
}

// WithStartDateTo adds the startDateTo to the get project params
func (o *GetProjectParams) WithStartDateTo(startDateTo *string) *GetProjectParams {
	o.SetStartDateTo(startDateTo)
	return o
}

// SetStartDateTo adds the startDateTo to the get project params
func (o *GetProjectParams) SetStartDateTo(startDateTo *string) {
	o.StartDateTo = startDateTo
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CustomerID != nil {

		// query param customerId
		var qrCustomerID string
		if o.CustomerID != nil {
			qrCustomerID = *o.CustomerID
		}
		qCustomerID := qrCustomerID
		if qCustomerID != "" {
			if err := r.SetQueryParam("customerId", qCustomerID); err != nil {
				return err
			}
		}

	}

	if o.DepartmentID != nil {

		// query param departmentId
		var qrDepartmentID string
		if o.DepartmentID != nil {
			qrDepartmentID = *o.DepartmentID
		}
		qDepartmentID := qrDepartmentID
		if qDepartmentID != "" {
			if err := r.SetQueryParam("departmentId", qDepartmentID); err != nil {
				return err
			}
		}

	}

	if o.EmployeeInProjectID != nil {

		// query param employeeInProjectId
		var qrEmployeeInProjectID string
		if o.EmployeeInProjectID != nil {
			qrEmployeeInProjectID = *o.EmployeeInProjectID
		}
		qEmployeeInProjectID := qrEmployeeInProjectID
		if qEmployeeInProjectID != "" {
			if err := r.SetQueryParam("employeeInProjectId", qEmployeeInProjectID); err != nil {
				return err
			}
		}

	}

	if o.EndDateFrom != nil {

		// query param endDateFrom
		var qrEndDateFrom string
		if o.EndDateFrom != nil {
			qrEndDateFrom = *o.EndDateFrom
		}
		qEndDateFrom := qrEndDateFrom
		if qEndDateFrom != "" {
			if err := r.SetQueryParam("endDateFrom", qEndDateFrom); err != nil {
				return err
			}
		}

	}

	if o.EndDateTo != nil {

		// query param endDateTo
		var qrEndDateTo string
		if o.EndDateTo != nil {
			qrEndDateTo = *o.EndDateTo
		}
		qEndDateTo := qrEndDateTo
		if qEndDateTo != "" {
			if err := r.SetQueryParam("endDateTo", qEndDateTo); err != nil {
				return err
			}
		}

	}

	if o.ExternalAccountsNumber != nil {

		// query param externalAccountsNumber
		var qrExternalAccountsNumber string
		if o.ExternalAccountsNumber != nil {
			qrExternalAccountsNumber = *o.ExternalAccountsNumber
		}
		qExternalAccountsNumber := qrExternalAccountsNumber
		if qExternalAccountsNumber != "" {
			if err := r.SetQueryParam("externalAccountsNumber", qExternalAccountsNumber); err != nil {
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

	if o.IsClosed != nil {

		// query param isClosed
		var qrIsClosed bool
		if o.IsClosed != nil {
			qrIsClosed = *o.IsClosed
		}
		qIsClosed := swag.FormatBool(qrIsClosed)
		if qIsClosed != "" {
			if err := r.SetQueryParam("isClosed", qIsClosed); err != nil {
				return err
			}
		}

	}

	if o.IsOffer != nil {

		// query param isOffer
		var qrIsOffer bool
		if o.IsOffer != nil {
			qrIsOffer = *o.IsOffer
		}
		qIsOffer := swag.FormatBool(qrIsOffer)
		if qIsOffer != "" {
			if err := r.SetQueryParam("isOffer", qIsOffer); err != nil {
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

	if o.ProjectManagerID != nil {

		// query param projectManagerId
		var qrProjectManagerID string
		if o.ProjectManagerID != nil {
			qrProjectManagerID = *o.ProjectManagerID
		}
		qProjectManagerID := qrProjectManagerID
		if qProjectManagerID != "" {
			if err := r.SetQueryParam("projectManagerId", qProjectManagerID); err != nil {
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

	if o.StartDateFrom != nil {

		// query param startDateFrom
		var qrStartDateFrom string
		if o.StartDateFrom != nil {
			qrStartDateFrom = *o.StartDateFrom
		}
		qStartDateFrom := qrStartDateFrom
		if qStartDateFrom != "" {
			if err := r.SetQueryParam("startDateFrom", qStartDateFrom); err != nil {
				return err
			}
		}

	}

	if o.StartDateTo != nil {

		// query param startDateTo
		var qrStartDateTo string
		if o.StartDateTo != nil {
			qrStartDateTo = *o.StartDateTo
		}
		qStartDateTo := qrStartDateTo
		if qStartDateTo != "" {
			if err := r.SetQueryParam("startDateTo", qStartDateTo); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
