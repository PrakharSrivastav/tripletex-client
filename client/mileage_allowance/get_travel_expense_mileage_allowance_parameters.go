// Code generated by go-swagger; DO NOT EDIT.

package mileage_allowance

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

// NewGetTravelExpenseMileageAllowanceParams creates a new GetTravelExpenseMileageAllowanceParams object
// with the default values initialized.
func NewGetTravelExpenseMileageAllowanceParams() *GetTravelExpenseMileageAllowanceParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTravelExpenseMileageAllowanceParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTravelExpenseMileageAllowanceParamsWithTimeout creates a new GetTravelExpenseMileageAllowanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTravelExpenseMileageAllowanceParamsWithTimeout(timeout time.Duration) *GetTravelExpenseMileageAllowanceParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTravelExpenseMileageAllowanceParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetTravelExpenseMileageAllowanceParamsWithContext creates a new GetTravelExpenseMileageAllowanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTravelExpenseMileageAllowanceParamsWithContext(ctx context.Context) *GetTravelExpenseMileageAllowanceParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTravelExpenseMileageAllowanceParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetTravelExpenseMileageAllowanceParamsWithHTTPClient creates a new GetTravelExpenseMileageAllowanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTravelExpenseMileageAllowanceParamsWithHTTPClient(client *http.Client) *GetTravelExpenseMileageAllowanceParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTravelExpenseMileageAllowanceParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetTravelExpenseMileageAllowanceParams contains all the parameters to send to the API endpoint
for the get travel expense mileage allowance operation typically these are written to a http.Request
*/
type GetTravelExpenseMileageAllowanceParams struct {

	/*AmountFrom
	  From and including

	*/
	AmountFrom *float64
	/*AmountTo
	  To and excluding

	*/
	AmountTo *float64
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
	/*DepartureLocation
	  Containing

	*/
	DepartureLocation *string
	/*Destination
	  Containing

	*/
	Destination *string
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*IsCompanyCar
	  Equals

	*/
	IsCompanyCar *bool
	/*KmFrom
	  From and including

	*/
	KmFrom *float64
	/*KmTo
	  To and excluding

	*/
	KmTo *float64
	/*RateCategoryID
	  Equals

	*/
	RateCategoryID *string
	/*RateFrom
	  From and including

	*/
	RateFrom *float64
	/*RateTo
	  To and excluding

	*/
	RateTo *float64
	/*RateTypeID
	  Equals

	*/
	RateTypeID *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string
	/*TravelExpenseID
	  Equals

	*/
	TravelExpenseID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithTimeout(timeout time.Duration) *GetTravelExpenseMileageAllowanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithContext(ctx context.Context) *GetTravelExpenseMileageAllowanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithHTTPClient(client *http.Client) *GetTravelExpenseMileageAllowanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmountFrom adds the amountFrom to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithAmountFrom(amountFrom *float64) *GetTravelExpenseMileageAllowanceParams {
	o.SetAmountFrom(amountFrom)
	return o
}

// SetAmountFrom adds the amountFrom to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetAmountFrom(amountFrom *float64) {
	o.AmountFrom = amountFrom
}

// WithAmountTo adds the amountTo to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithAmountTo(amountTo *float64) *GetTravelExpenseMileageAllowanceParams {
	o.SetAmountTo(amountTo)
	return o
}

// SetAmountTo adds the amountTo to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetAmountTo(amountTo *float64) {
	o.AmountTo = amountTo
}

// WithCount adds the count to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithCount(count *int64) *GetTravelExpenseMileageAllowanceParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetCount(count *int64) {
	o.Count = count
}

// WithDateFrom adds the dateFrom to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithDateFrom(dateFrom *string) *GetTravelExpenseMileageAllowanceParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetDateFrom(dateFrom *string) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithDateTo(dateTo *string) *GetTravelExpenseMileageAllowanceParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetDateTo(dateTo *string) {
	o.DateTo = dateTo
}

// WithDepartureLocation adds the departureLocation to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithDepartureLocation(departureLocation *string) *GetTravelExpenseMileageAllowanceParams {
	o.SetDepartureLocation(departureLocation)
	return o
}

// SetDepartureLocation adds the departureLocation to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetDepartureLocation(departureLocation *string) {
	o.DepartureLocation = departureLocation
}

// WithDestination adds the destination to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithDestination(destination *string) *GetTravelExpenseMileageAllowanceParams {
	o.SetDestination(destination)
	return o
}

// SetDestination adds the destination to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetDestination(destination *string) {
	o.Destination = destination
}

// WithFields adds the fields to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithFields(fields *string) *GetTravelExpenseMileageAllowanceParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithFrom(from *int64) *GetTravelExpenseMileageAllowanceParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetFrom(from *int64) {
	o.From = from
}

// WithIsCompanyCar adds the isCompanyCar to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithIsCompanyCar(isCompanyCar *bool) *GetTravelExpenseMileageAllowanceParams {
	o.SetIsCompanyCar(isCompanyCar)
	return o
}

// SetIsCompanyCar adds the isCompanyCar to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetIsCompanyCar(isCompanyCar *bool) {
	o.IsCompanyCar = isCompanyCar
}

// WithKmFrom adds the kmFrom to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithKmFrom(kmFrom *float64) *GetTravelExpenseMileageAllowanceParams {
	o.SetKmFrom(kmFrom)
	return o
}

// SetKmFrom adds the kmFrom to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetKmFrom(kmFrom *float64) {
	o.KmFrom = kmFrom
}

// WithKmTo adds the kmTo to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithKmTo(kmTo *float64) *GetTravelExpenseMileageAllowanceParams {
	o.SetKmTo(kmTo)
	return o
}

// SetKmTo adds the kmTo to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetKmTo(kmTo *float64) {
	o.KmTo = kmTo
}

// WithRateCategoryID adds the rateCategoryID to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithRateCategoryID(rateCategoryID *string) *GetTravelExpenseMileageAllowanceParams {
	o.SetRateCategoryID(rateCategoryID)
	return o
}

// SetRateCategoryID adds the rateCategoryId to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetRateCategoryID(rateCategoryID *string) {
	o.RateCategoryID = rateCategoryID
}

// WithRateFrom adds the rateFrom to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithRateFrom(rateFrom *float64) *GetTravelExpenseMileageAllowanceParams {
	o.SetRateFrom(rateFrom)
	return o
}

// SetRateFrom adds the rateFrom to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetRateFrom(rateFrom *float64) {
	o.RateFrom = rateFrom
}

// WithRateTo adds the rateTo to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithRateTo(rateTo *float64) *GetTravelExpenseMileageAllowanceParams {
	o.SetRateTo(rateTo)
	return o
}

// SetRateTo adds the rateTo to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetRateTo(rateTo *float64) {
	o.RateTo = rateTo
}

// WithRateTypeID adds the rateTypeID to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithRateTypeID(rateTypeID *string) *GetTravelExpenseMileageAllowanceParams {
	o.SetRateTypeID(rateTypeID)
	return o
}

// SetRateTypeID adds the rateTypeId to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetRateTypeID(rateTypeID *string) {
	o.RateTypeID = rateTypeID
}

// WithSorting adds the sorting to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithSorting(sorting *string) *GetTravelExpenseMileageAllowanceParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithTravelExpenseID adds the travelExpenseID to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) WithTravelExpenseID(travelExpenseID *string) *GetTravelExpenseMileageAllowanceParams {
	o.SetTravelExpenseID(travelExpenseID)
	return o
}

// SetTravelExpenseID adds the travelExpenseId to the get travel expense mileage allowance params
func (o *GetTravelExpenseMileageAllowanceParams) SetTravelExpenseID(travelExpenseID *string) {
	o.TravelExpenseID = travelExpenseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTravelExpenseMileageAllowanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AmountFrom != nil {

		// query param amountFrom
		var qrAmountFrom float64
		if o.AmountFrom != nil {
			qrAmountFrom = *o.AmountFrom
		}
		qAmountFrom := swag.FormatFloat64(qrAmountFrom)
		if qAmountFrom != "" {
			if err := r.SetQueryParam("amountFrom", qAmountFrom); err != nil {
				return err
			}
		}

	}

	if o.AmountTo != nil {

		// query param amountTo
		var qrAmountTo float64
		if o.AmountTo != nil {
			qrAmountTo = *o.AmountTo
		}
		qAmountTo := swag.FormatFloat64(qrAmountTo)
		if qAmountTo != "" {
			if err := r.SetQueryParam("amountTo", qAmountTo); err != nil {
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

	if o.DepartureLocation != nil {

		// query param departureLocation
		var qrDepartureLocation string
		if o.DepartureLocation != nil {
			qrDepartureLocation = *o.DepartureLocation
		}
		qDepartureLocation := qrDepartureLocation
		if qDepartureLocation != "" {
			if err := r.SetQueryParam("departureLocation", qDepartureLocation); err != nil {
				return err
			}
		}

	}

	if o.Destination != nil {

		// query param destination
		var qrDestination string
		if o.Destination != nil {
			qrDestination = *o.Destination
		}
		qDestination := qrDestination
		if qDestination != "" {
			if err := r.SetQueryParam("destination", qDestination); err != nil {
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

	if o.IsCompanyCar != nil {

		// query param isCompanyCar
		var qrIsCompanyCar bool
		if o.IsCompanyCar != nil {
			qrIsCompanyCar = *o.IsCompanyCar
		}
		qIsCompanyCar := swag.FormatBool(qrIsCompanyCar)
		if qIsCompanyCar != "" {
			if err := r.SetQueryParam("isCompanyCar", qIsCompanyCar); err != nil {
				return err
			}
		}

	}

	if o.KmFrom != nil {

		// query param kmFrom
		var qrKmFrom float64
		if o.KmFrom != nil {
			qrKmFrom = *o.KmFrom
		}
		qKmFrom := swag.FormatFloat64(qrKmFrom)
		if qKmFrom != "" {
			if err := r.SetQueryParam("kmFrom", qKmFrom); err != nil {
				return err
			}
		}

	}

	if o.KmTo != nil {

		// query param kmTo
		var qrKmTo float64
		if o.KmTo != nil {
			qrKmTo = *o.KmTo
		}
		qKmTo := swag.FormatFloat64(qrKmTo)
		if qKmTo != "" {
			if err := r.SetQueryParam("kmTo", qKmTo); err != nil {
				return err
			}
		}

	}

	if o.RateCategoryID != nil {

		// query param rateCategoryId
		var qrRateCategoryID string
		if o.RateCategoryID != nil {
			qrRateCategoryID = *o.RateCategoryID
		}
		qRateCategoryID := qrRateCategoryID
		if qRateCategoryID != "" {
			if err := r.SetQueryParam("rateCategoryId", qRateCategoryID); err != nil {
				return err
			}
		}

	}

	if o.RateFrom != nil {

		// query param rateFrom
		var qrRateFrom float64
		if o.RateFrom != nil {
			qrRateFrom = *o.RateFrom
		}
		qRateFrom := swag.FormatFloat64(qrRateFrom)
		if qRateFrom != "" {
			if err := r.SetQueryParam("rateFrom", qRateFrom); err != nil {
				return err
			}
		}

	}

	if o.RateTo != nil {

		// query param rateTo
		var qrRateTo float64
		if o.RateTo != nil {
			qrRateTo = *o.RateTo
		}
		qRateTo := swag.FormatFloat64(qrRateTo)
		if qRateTo != "" {
			if err := r.SetQueryParam("rateTo", qRateTo); err != nil {
				return err
			}
		}

	}

	if o.RateTypeID != nil {

		// query param rateTypeId
		var qrRateTypeID string
		if o.RateTypeID != nil {
			qrRateTypeID = *o.RateTypeID
		}
		qRateTypeID := qrRateTypeID
		if qRateTypeID != "" {
			if err := r.SetQueryParam("rateTypeId", qRateTypeID); err != nil {
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

	if o.TravelExpenseID != nil {

		// query param travelExpenseId
		var qrTravelExpenseID string
		if o.TravelExpenseID != nil {
			qrTravelExpenseID = *o.TravelExpenseID
		}
		qTravelExpenseID := qrTravelExpenseID
		if qTravelExpenseID != "" {
			if err := r.SetQueryParam("travelExpenseId", qTravelExpenseID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
