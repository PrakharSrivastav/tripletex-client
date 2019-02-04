// Code generated by go-swagger; DO NOT EDIT.

package prospect

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

// NewGetCrmProspectParams creates a new GetCrmProspectParams object
// with the default values initialized.
func NewGetCrmProspectParams() *GetCrmProspectParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCrmProspectParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCrmProspectParamsWithTimeout creates a new GetCrmProspectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCrmProspectParamsWithTimeout(timeout time.Duration) *GetCrmProspectParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCrmProspectParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetCrmProspectParamsWithContext creates a new GetCrmProspectParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCrmProspectParamsWithContext(ctx context.Context) *GetCrmProspectParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCrmProspectParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetCrmProspectParamsWithHTTPClient creates a new GetCrmProspectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCrmProspectParamsWithHTTPClient(client *http.Client) *GetCrmProspectParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetCrmProspectParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetCrmProspectParams contains all the parameters to send to the API endpoint
for the get crm prospect operation typically these are written to a http.Request
*/
type GetCrmProspectParams struct {

	/*ClosedDateFrom
	  From and including

	*/
	ClosedDateFrom *string
	/*ClosedDateTo
	  To and excluding

	*/
	ClosedDateTo *string
	/*ClosedReason
	  Equals

	*/
	ClosedReason *string
	/*Competitor
	  Containing

	*/
	Competitor *string
	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*CreatedDateFrom
	  From and including

	*/
	CreatedDateFrom *string
	/*CreatedDateTo
	  To and excluding

	*/
	CreatedDateTo *string
	/*CustomerID
	  Equals

	*/
	CustomerID *string
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
	/*IsClosed
	  Equals

	*/
	IsClosed *bool
	/*Name
	  Containing

	*/
	Name *string
	/*ProjectID
	  Equals

	*/
	ProjectID *string
	/*ProjectOfferID
	  Equals

	*/
	ProjectOfferID *string
	/*ProspectType
	  Equals

	*/
	ProspectType *string
	/*SalesEmployeeID
	  Equals

	*/
	SalesEmployeeID *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get crm prospect params
func (o *GetCrmProspectParams) WithTimeout(timeout time.Duration) *GetCrmProspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get crm prospect params
func (o *GetCrmProspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get crm prospect params
func (o *GetCrmProspectParams) WithContext(ctx context.Context) *GetCrmProspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get crm prospect params
func (o *GetCrmProspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get crm prospect params
func (o *GetCrmProspectParams) WithHTTPClient(client *http.Client) *GetCrmProspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get crm prospect params
func (o *GetCrmProspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClosedDateFrom adds the closedDateFrom to the get crm prospect params
func (o *GetCrmProspectParams) WithClosedDateFrom(closedDateFrom *string) *GetCrmProspectParams {
	o.SetClosedDateFrom(closedDateFrom)
	return o
}

// SetClosedDateFrom adds the closedDateFrom to the get crm prospect params
func (o *GetCrmProspectParams) SetClosedDateFrom(closedDateFrom *string) {
	o.ClosedDateFrom = closedDateFrom
}

// WithClosedDateTo adds the closedDateTo to the get crm prospect params
func (o *GetCrmProspectParams) WithClosedDateTo(closedDateTo *string) *GetCrmProspectParams {
	o.SetClosedDateTo(closedDateTo)
	return o
}

// SetClosedDateTo adds the closedDateTo to the get crm prospect params
func (o *GetCrmProspectParams) SetClosedDateTo(closedDateTo *string) {
	o.ClosedDateTo = closedDateTo
}

// WithClosedReason adds the closedReason to the get crm prospect params
func (o *GetCrmProspectParams) WithClosedReason(closedReason *string) *GetCrmProspectParams {
	o.SetClosedReason(closedReason)
	return o
}

// SetClosedReason adds the closedReason to the get crm prospect params
func (o *GetCrmProspectParams) SetClosedReason(closedReason *string) {
	o.ClosedReason = closedReason
}

// WithCompetitor adds the competitor to the get crm prospect params
func (o *GetCrmProspectParams) WithCompetitor(competitor *string) *GetCrmProspectParams {
	o.SetCompetitor(competitor)
	return o
}

// SetCompetitor adds the competitor to the get crm prospect params
func (o *GetCrmProspectParams) SetCompetitor(competitor *string) {
	o.Competitor = competitor
}

// WithCount adds the count to the get crm prospect params
func (o *GetCrmProspectParams) WithCount(count *int64) *GetCrmProspectParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get crm prospect params
func (o *GetCrmProspectParams) SetCount(count *int64) {
	o.Count = count
}

// WithCreatedDateFrom adds the createdDateFrom to the get crm prospect params
func (o *GetCrmProspectParams) WithCreatedDateFrom(createdDateFrom *string) *GetCrmProspectParams {
	o.SetCreatedDateFrom(createdDateFrom)
	return o
}

// SetCreatedDateFrom adds the createdDateFrom to the get crm prospect params
func (o *GetCrmProspectParams) SetCreatedDateFrom(createdDateFrom *string) {
	o.CreatedDateFrom = createdDateFrom
}

// WithCreatedDateTo adds the createdDateTo to the get crm prospect params
func (o *GetCrmProspectParams) WithCreatedDateTo(createdDateTo *string) *GetCrmProspectParams {
	o.SetCreatedDateTo(createdDateTo)
	return o
}

// SetCreatedDateTo adds the createdDateTo to the get crm prospect params
func (o *GetCrmProspectParams) SetCreatedDateTo(createdDateTo *string) {
	o.CreatedDateTo = createdDateTo
}

// WithCustomerID adds the customerID to the get crm prospect params
func (o *GetCrmProspectParams) WithCustomerID(customerID *string) *GetCrmProspectParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get crm prospect params
func (o *GetCrmProspectParams) SetCustomerID(customerID *string) {
	o.CustomerID = customerID
}

// WithDescription adds the description to the get crm prospect params
func (o *GetCrmProspectParams) WithDescription(description *string) *GetCrmProspectParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the get crm prospect params
func (o *GetCrmProspectParams) SetDescription(description *string) {
	o.Description = description
}

// WithFields adds the fields to the get crm prospect params
func (o *GetCrmProspectParams) WithFields(fields *string) *GetCrmProspectParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get crm prospect params
func (o *GetCrmProspectParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get crm prospect params
func (o *GetCrmProspectParams) WithFrom(from *int64) *GetCrmProspectParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get crm prospect params
func (o *GetCrmProspectParams) SetFrom(from *int64) {
	o.From = from
}

// WithIsClosed adds the isClosed to the get crm prospect params
func (o *GetCrmProspectParams) WithIsClosed(isClosed *bool) *GetCrmProspectParams {
	o.SetIsClosed(isClosed)
	return o
}

// SetIsClosed adds the isClosed to the get crm prospect params
func (o *GetCrmProspectParams) SetIsClosed(isClosed *bool) {
	o.IsClosed = isClosed
}

// WithName adds the name to the get crm prospect params
func (o *GetCrmProspectParams) WithName(name *string) *GetCrmProspectParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get crm prospect params
func (o *GetCrmProspectParams) SetName(name *string) {
	o.Name = name
}

// WithProjectID adds the projectID to the get crm prospect params
func (o *GetCrmProspectParams) WithProjectID(projectID *string) *GetCrmProspectParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get crm prospect params
func (o *GetCrmProspectParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithProjectOfferID adds the projectOfferID to the get crm prospect params
func (o *GetCrmProspectParams) WithProjectOfferID(projectOfferID *string) *GetCrmProspectParams {
	o.SetProjectOfferID(projectOfferID)
	return o
}

// SetProjectOfferID adds the projectOfferId to the get crm prospect params
func (o *GetCrmProspectParams) SetProjectOfferID(projectOfferID *string) {
	o.ProjectOfferID = projectOfferID
}

// WithProspectType adds the prospectType to the get crm prospect params
func (o *GetCrmProspectParams) WithProspectType(prospectType *string) *GetCrmProspectParams {
	o.SetProspectType(prospectType)
	return o
}

// SetProspectType adds the prospectType to the get crm prospect params
func (o *GetCrmProspectParams) SetProspectType(prospectType *string) {
	o.ProspectType = prospectType
}

// WithSalesEmployeeID adds the salesEmployeeID to the get crm prospect params
func (o *GetCrmProspectParams) WithSalesEmployeeID(salesEmployeeID *string) *GetCrmProspectParams {
	o.SetSalesEmployeeID(salesEmployeeID)
	return o
}

// SetSalesEmployeeID adds the salesEmployeeId to the get crm prospect params
func (o *GetCrmProspectParams) SetSalesEmployeeID(salesEmployeeID *string) {
	o.SalesEmployeeID = salesEmployeeID
}

// WithSorting adds the sorting to the get crm prospect params
func (o *GetCrmProspectParams) WithSorting(sorting *string) *GetCrmProspectParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get crm prospect params
func (o *GetCrmProspectParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *GetCrmProspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClosedDateFrom != nil {

		// query param closedDateFrom
		var qrClosedDateFrom string
		if o.ClosedDateFrom != nil {
			qrClosedDateFrom = *o.ClosedDateFrom
		}
		qClosedDateFrom := qrClosedDateFrom
		if qClosedDateFrom != "" {
			if err := r.SetQueryParam("closedDateFrom", qClosedDateFrom); err != nil {
				return err
			}
		}

	}

	if o.ClosedDateTo != nil {

		// query param closedDateTo
		var qrClosedDateTo string
		if o.ClosedDateTo != nil {
			qrClosedDateTo = *o.ClosedDateTo
		}
		qClosedDateTo := qrClosedDateTo
		if qClosedDateTo != "" {
			if err := r.SetQueryParam("closedDateTo", qClosedDateTo); err != nil {
				return err
			}
		}

	}

	if o.ClosedReason != nil {

		// query param closedReason
		var qrClosedReason string
		if o.ClosedReason != nil {
			qrClosedReason = *o.ClosedReason
		}
		qClosedReason := qrClosedReason
		if qClosedReason != "" {
			if err := r.SetQueryParam("closedReason", qClosedReason); err != nil {
				return err
			}
		}

	}

	if o.Competitor != nil {

		// query param competitor
		var qrCompetitor string
		if o.Competitor != nil {
			qrCompetitor = *o.Competitor
		}
		qCompetitor := qrCompetitor
		if qCompetitor != "" {
			if err := r.SetQueryParam("competitor", qCompetitor); err != nil {
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

	if o.CreatedDateFrom != nil {

		// query param createdDateFrom
		var qrCreatedDateFrom string
		if o.CreatedDateFrom != nil {
			qrCreatedDateFrom = *o.CreatedDateFrom
		}
		qCreatedDateFrom := qrCreatedDateFrom
		if qCreatedDateFrom != "" {
			if err := r.SetQueryParam("createdDateFrom", qCreatedDateFrom); err != nil {
				return err
			}
		}

	}

	if o.CreatedDateTo != nil {

		// query param createdDateTo
		var qrCreatedDateTo string
		if o.CreatedDateTo != nil {
			qrCreatedDateTo = *o.CreatedDateTo
		}
		qCreatedDateTo := qrCreatedDateTo
		if qCreatedDateTo != "" {
			if err := r.SetQueryParam("createdDateTo", qCreatedDateTo); err != nil {
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

	if o.ProjectOfferID != nil {

		// query param projectOfferId
		var qrProjectOfferID string
		if o.ProjectOfferID != nil {
			qrProjectOfferID = *o.ProjectOfferID
		}
		qProjectOfferID := qrProjectOfferID
		if qProjectOfferID != "" {
			if err := r.SetQueryParam("projectOfferId", qProjectOfferID); err != nil {
				return err
			}
		}

	}

	if o.ProspectType != nil {

		// query param prospectType
		var qrProspectType string
		if o.ProspectType != nil {
			qrProspectType = *o.ProspectType
		}
		qProspectType := qrProspectType
		if qProspectType != "" {
			if err := r.SetQueryParam("prospectType", qProspectType); err != nil {
				return err
			}
		}

	}

	if o.SalesEmployeeID != nil {

		// query param salesEmployeeId
		var qrSalesEmployeeID string
		if o.SalesEmployeeID != nil {
			qrSalesEmployeeID = *o.SalesEmployeeID
		}
		qSalesEmployeeID := qrSalesEmployeeID
		if qSalesEmployeeID != "" {
			if err := r.SetQueryParam("salesEmployeeId", qSalesEmployeeID); err != nil {
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