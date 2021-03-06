// Code generated by go-swagger; DO NOT EDIT.

package posting

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

// NewGetLedgerPostingOpenPostParams creates a new GetLedgerPostingOpenPostParams object
// with the default values initialized.
func NewGetLedgerPostingOpenPostParams() *GetLedgerPostingOpenPostParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetLedgerPostingOpenPostParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLedgerPostingOpenPostParamsWithTimeout creates a new GetLedgerPostingOpenPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLedgerPostingOpenPostParamsWithTimeout(timeout time.Duration) *GetLedgerPostingOpenPostParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetLedgerPostingOpenPostParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetLedgerPostingOpenPostParamsWithContext creates a new GetLedgerPostingOpenPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLedgerPostingOpenPostParamsWithContext(ctx context.Context) *GetLedgerPostingOpenPostParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetLedgerPostingOpenPostParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetLedgerPostingOpenPostParamsWithHTTPClient creates a new GetLedgerPostingOpenPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLedgerPostingOpenPostParamsWithHTTPClient(client *http.Client) *GetLedgerPostingOpenPostParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetLedgerPostingOpenPostParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetLedgerPostingOpenPostParams contains all the parameters to send to the API endpoint
for the get ledger posting open post operation typically these are written to a http.Request
*/
type GetLedgerPostingOpenPostParams struct {

	/*AccountID
	  Element ID

	*/
	AccountID *int32
	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*CustomerID
	  Element ID

	*/
	CustomerID *int32
	/*Date
	  Invoice date. Format is yyyy-MM-dd (to and excl.).

	*/
	Date string
	/*DepartmentID
	  Element ID

	*/
	DepartmentID *int32
	/*EmployeeID
	  Element ID

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
	/*ProductID
	  Element ID

	*/
	ProductID *int32
	/*ProjectID
	  Element ID

	*/
	ProjectID *int32
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string
	/*SupplierID
	  Element ID

	*/
	SupplierID *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithTimeout(timeout time.Duration) *GetLedgerPostingOpenPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithContext(ctx context.Context) *GetLedgerPostingOpenPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithHTTPClient(client *http.Client) *GetLedgerPostingOpenPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithAccountID(accountID *int32) *GetLedgerPostingOpenPostParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetAccountID(accountID *int32) {
	o.AccountID = accountID
}

// WithCount adds the count to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithCount(count *int64) *GetLedgerPostingOpenPostParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetCount(count *int64) {
	o.Count = count
}

// WithCustomerID adds the customerID to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithCustomerID(customerID *int32) *GetLedgerPostingOpenPostParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetCustomerID(customerID *int32) {
	o.CustomerID = customerID
}

// WithDate adds the date to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithDate(date string) *GetLedgerPostingOpenPostParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetDate(date string) {
	o.Date = date
}

// WithDepartmentID adds the departmentID to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithDepartmentID(departmentID *int32) *GetLedgerPostingOpenPostParams {
	o.SetDepartmentID(departmentID)
	return o
}

// SetDepartmentID adds the departmentId to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetDepartmentID(departmentID *int32) {
	o.DepartmentID = departmentID
}

// WithEmployeeID adds the employeeID to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithEmployeeID(employeeID *int32) *GetLedgerPostingOpenPostParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetEmployeeID(employeeID *int32) {
	o.EmployeeID = employeeID
}

// WithFields adds the fields to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithFields(fields *string) *GetLedgerPostingOpenPostParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithFrom(from *int64) *GetLedgerPostingOpenPostParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetFrom(from *int64) {
	o.From = from
}

// WithProductID adds the productID to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithProductID(productID *int32) *GetLedgerPostingOpenPostParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetProductID(productID *int32) {
	o.ProductID = productID
}

// WithProjectID adds the projectID to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithProjectID(projectID *int32) *GetLedgerPostingOpenPostParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetProjectID(projectID *int32) {
	o.ProjectID = projectID
}

// WithSorting adds the sorting to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithSorting(sorting *string) *GetLedgerPostingOpenPostParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithSupplierID adds the supplierID to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) WithSupplierID(supplierID *int32) *GetLedgerPostingOpenPostParams {
	o.SetSupplierID(supplierID)
	return o
}

// SetSupplierID adds the supplierId to the get ledger posting open post params
func (o *GetLedgerPostingOpenPostParams) SetSupplierID(supplierID *int32) {
	o.SupplierID = supplierID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLedgerPostingOpenPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID int32
		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := swag.FormatInt32(qrAccountID)
		if qAccountID != "" {
			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
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

	if o.CustomerID != nil {

		// query param customerId
		var qrCustomerID int32
		if o.CustomerID != nil {
			qrCustomerID = *o.CustomerID
		}
		qCustomerID := swag.FormatInt32(qrCustomerID)
		if qCustomerID != "" {
			if err := r.SetQueryParam("customerId", qCustomerID); err != nil {
				return err
			}
		}

	}

	// query param date
	qrDate := o.Date
	qDate := qrDate
	if qDate != "" {
		if err := r.SetQueryParam("date", qDate); err != nil {
			return err
		}
	}

	if o.DepartmentID != nil {

		// query param departmentId
		var qrDepartmentID int32
		if o.DepartmentID != nil {
			qrDepartmentID = *o.DepartmentID
		}
		qDepartmentID := swag.FormatInt32(qrDepartmentID)
		if qDepartmentID != "" {
			if err := r.SetQueryParam("departmentId", qDepartmentID); err != nil {
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

	if o.ProductID != nil {

		// query param productId
		var qrProductID int32
		if o.ProductID != nil {
			qrProductID = *o.ProductID
		}
		qProductID := swag.FormatInt32(qrProductID)
		if qProductID != "" {
			if err := r.SetQueryParam("productId", qProductID); err != nil {
				return err
			}
		}

	}

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID int32
		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := swag.FormatInt32(qrProjectID)
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

	if o.SupplierID != nil {

		// query param supplierId
		var qrSupplierID int32
		if o.SupplierID != nil {
			qrSupplierID = *o.SupplierID
		}
		qSupplierID := swag.FormatInt32(qrSupplierID)
		if qSupplierID != "" {
			if err := r.SetQueryParam("supplierId", qSupplierID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
