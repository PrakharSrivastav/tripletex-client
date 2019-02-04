// Code generated by go-swagger; DO NOT EDIT.

package per_diem_compensation

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

// NewGetTravelExpensePerDiemCompensationParams creates a new GetTravelExpensePerDiemCompensationParams object
// with the default values initialized.
func NewGetTravelExpensePerDiemCompensationParams() *GetTravelExpensePerDiemCompensationParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTravelExpensePerDiemCompensationParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTravelExpensePerDiemCompensationParamsWithTimeout creates a new GetTravelExpensePerDiemCompensationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTravelExpensePerDiemCompensationParamsWithTimeout(timeout time.Duration) *GetTravelExpensePerDiemCompensationParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTravelExpensePerDiemCompensationParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewGetTravelExpensePerDiemCompensationParamsWithContext creates a new GetTravelExpensePerDiemCompensationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTravelExpensePerDiemCompensationParamsWithContext(ctx context.Context) *GetTravelExpensePerDiemCompensationParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTravelExpensePerDiemCompensationParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewGetTravelExpensePerDiemCompensationParamsWithHTTPClient creates a new GetTravelExpensePerDiemCompensationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTravelExpensePerDiemCompensationParamsWithHTTPClient(client *http.Client) *GetTravelExpensePerDiemCompensationParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &GetTravelExpensePerDiemCompensationParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*GetTravelExpensePerDiemCompensationParams contains all the parameters to send to the API endpoint
for the get travel expense per diem compensation operation typically these are written to a http.Request
*/
type GetTravelExpensePerDiemCompensationParams struct {

	/*Address
	  Containing

	*/
	Address *string
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
	/*CountFrom
	  From and including

	*/
	CountFrom *int32
	/*CountTo
	  To and excluding

	*/
	CountTo *int32
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*IsDeductionForBreakfast
	  Equals

	*/
	IsDeductionForBreakfast *bool
	/*IsDinnerDeduction
	  Equals

	*/
	IsDinnerDeduction *bool
	/*IsLunchDeduction
	  Equals

	*/
	IsLunchDeduction *bool
	/*Location
	  Containing

	*/
	Location *string
	/*OvernightAccommodation
	  Equals

	*/
	OvernightAccommodation *string
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

// WithTimeout adds the timeout to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithTimeout(timeout time.Duration) *GetTravelExpensePerDiemCompensationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithContext(ctx context.Context) *GetTravelExpensePerDiemCompensationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithHTTPClient(client *http.Client) *GetTravelExpensePerDiemCompensationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithAddress(address *string) *GetTravelExpensePerDiemCompensationParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetAddress(address *string) {
	o.Address = address
}

// WithAmountFrom adds the amountFrom to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithAmountFrom(amountFrom *float64) *GetTravelExpensePerDiemCompensationParams {
	o.SetAmountFrom(amountFrom)
	return o
}

// SetAmountFrom adds the amountFrom to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetAmountFrom(amountFrom *float64) {
	o.AmountFrom = amountFrom
}

// WithAmountTo adds the amountTo to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithAmountTo(amountTo *float64) *GetTravelExpensePerDiemCompensationParams {
	o.SetAmountTo(amountTo)
	return o
}

// SetAmountTo adds the amountTo to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetAmountTo(amountTo *float64) {
	o.AmountTo = amountTo
}

// WithCount adds the count to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithCount(count *int64) *GetTravelExpensePerDiemCompensationParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetCount(count *int64) {
	o.Count = count
}

// WithCountFrom adds the countFrom to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithCountFrom(countFrom *int32) *GetTravelExpensePerDiemCompensationParams {
	o.SetCountFrom(countFrom)
	return o
}

// SetCountFrom adds the countFrom to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetCountFrom(countFrom *int32) {
	o.CountFrom = countFrom
}

// WithCountTo adds the countTo to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithCountTo(countTo *int32) *GetTravelExpensePerDiemCompensationParams {
	o.SetCountTo(countTo)
	return o
}

// SetCountTo adds the countTo to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetCountTo(countTo *int32) {
	o.CountTo = countTo
}

// WithFields adds the fields to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithFields(fields *string) *GetTravelExpensePerDiemCompensationParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithFrom(from *int64) *GetTravelExpensePerDiemCompensationParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetFrom(from *int64) {
	o.From = from
}

// WithIsDeductionForBreakfast adds the isDeductionForBreakfast to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithIsDeductionForBreakfast(isDeductionForBreakfast *bool) *GetTravelExpensePerDiemCompensationParams {
	o.SetIsDeductionForBreakfast(isDeductionForBreakfast)
	return o
}

// SetIsDeductionForBreakfast adds the isDeductionForBreakfast to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetIsDeductionForBreakfast(isDeductionForBreakfast *bool) {
	o.IsDeductionForBreakfast = isDeductionForBreakfast
}

// WithIsDinnerDeduction adds the isDinnerDeduction to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithIsDinnerDeduction(isDinnerDeduction *bool) *GetTravelExpensePerDiemCompensationParams {
	o.SetIsDinnerDeduction(isDinnerDeduction)
	return o
}

// SetIsDinnerDeduction adds the isDinnerDeduction to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetIsDinnerDeduction(isDinnerDeduction *bool) {
	o.IsDinnerDeduction = isDinnerDeduction
}

// WithIsLunchDeduction adds the isLunchDeduction to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithIsLunchDeduction(isLunchDeduction *bool) *GetTravelExpensePerDiemCompensationParams {
	o.SetIsLunchDeduction(isLunchDeduction)
	return o
}

// SetIsLunchDeduction adds the isLunchDeduction to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetIsLunchDeduction(isLunchDeduction *bool) {
	o.IsLunchDeduction = isLunchDeduction
}

// WithLocation adds the location to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithLocation(location *string) *GetTravelExpensePerDiemCompensationParams {
	o.SetLocation(location)
	return o
}

// SetLocation adds the location to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetLocation(location *string) {
	o.Location = location
}

// WithOvernightAccommodation adds the overnightAccommodation to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithOvernightAccommodation(overnightAccommodation *string) *GetTravelExpensePerDiemCompensationParams {
	o.SetOvernightAccommodation(overnightAccommodation)
	return o
}

// SetOvernightAccommodation adds the overnightAccommodation to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetOvernightAccommodation(overnightAccommodation *string) {
	o.OvernightAccommodation = overnightAccommodation
}

// WithRateCategoryID adds the rateCategoryID to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithRateCategoryID(rateCategoryID *string) *GetTravelExpensePerDiemCompensationParams {
	o.SetRateCategoryID(rateCategoryID)
	return o
}

// SetRateCategoryID adds the rateCategoryId to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetRateCategoryID(rateCategoryID *string) {
	o.RateCategoryID = rateCategoryID
}

// WithRateFrom adds the rateFrom to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithRateFrom(rateFrom *float64) *GetTravelExpensePerDiemCompensationParams {
	o.SetRateFrom(rateFrom)
	return o
}

// SetRateFrom adds the rateFrom to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetRateFrom(rateFrom *float64) {
	o.RateFrom = rateFrom
}

// WithRateTo adds the rateTo to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithRateTo(rateTo *float64) *GetTravelExpensePerDiemCompensationParams {
	o.SetRateTo(rateTo)
	return o
}

// SetRateTo adds the rateTo to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetRateTo(rateTo *float64) {
	o.RateTo = rateTo
}

// WithRateTypeID adds the rateTypeID to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithRateTypeID(rateTypeID *string) *GetTravelExpensePerDiemCompensationParams {
	o.SetRateTypeID(rateTypeID)
	return o
}

// SetRateTypeID adds the rateTypeId to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetRateTypeID(rateTypeID *string) {
	o.RateTypeID = rateTypeID
}

// WithSorting adds the sorting to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithSorting(sorting *string) *GetTravelExpensePerDiemCompensationParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithTravelExpenseID adds the travelExpenseID to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) WithTravelExpenseID(travelExpenseID *string) *GetTravelExpensePerDiemCompensationParams {
	o.SetTravelExpenseID(travelExpenseID)
	return o
}

// SetTravelExpenseID adds the travelExpenseId to the get travel expense per diem compensation params
func (o *GetTravelExpensePerDiemCompensationParams) SetTravelExpenseID(travelExpenseID *string) {
	o.TravelExpenseID = travelExpenseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTravelExpensePerDiemCompensationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Address != nil {

		// query param address
		var qrAddress string
		if o.Address != nil {
			qrAddress = *o.Address
		}
		qAddress := qrAddress
		if qAddress != "" {
			if err := r.SetQueryParam("address", qAddress); err != nil {
				return err
			}
		}

	}

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

	if o.CountFrom != nil {

		// query param countFrom
		var qrCountFrom int32
		if o.CountFrom != nil {
			qrCountFrom = *o.CountFrom
		}
		qCountFrom := swag.FormatInt32(qrCountFrom)
		if qCountFrom != "" {
			if err := r.SetQueryParam("countFrom", qCountFrom); err != nil {
				return err
			}
		}

	}

	if o.CountTo != nil {

		// query param countTo
		var qrCountTo int32
		if o.CountTo != nil {
			qrCountTo = *o.CountTo
		}
		qCountTo := swag.FormatInt32(qrCountTo)
		if qCountTo != "" {
			if err := r.SetQueryParam("countTo", qCountTo); err != nil {
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

	if o.IsDeductionForBreakfast != nil {

		// query param isDeductionForBreakfast
		var qrIsDeductionForBreakfast bool
		if o.IsDeductionForBreakfast != nil {
			qrIsDeductionForBreakfast = *o.IsDeductionForBreakfast
		}
		qIsDeductionForBreakfast := swag.FormatBool(qrIsDeductionForBreakfast)
		if qIsDeductionForBreakfast != "" {
			if err := r.SetQueryParam("isDeductionForBreakfast", qIsDeductionForBreakfast); err != nil {
				return err
			}
		}

	}

	if o.IsDinnerDeduction != nil {

		// query param isDinnerDeduction
		var qrIsDinnerDeduction bool
		if o.IsDinnerDeduction != nil {
			qrIsDinnerDeduction = *o.IsDinnerDeduction
		}
		qIsDinnerDeduction := swag.FormatBool(qrIsDinnerDeduction)
		if qIsDinnerDeduction != "" {
			if err := r.SetQueryParam("isDinnerDeduction", qIsDinnerDeduction); err != nil {
				return err
			}
		}

	}

	if o.IsLunchDeduction != nil {

		// query param isLunchDeduction
		var qrIsLunchDeduction bool
		if o.IsLunchDeduction != nil {
			qrIsLunchDeduction = *o.IsLunchDeduction
		}
		qIsLunchDeduction := swag.FormatBool(qrIsLunchDeduction)
		if qIsLunchDeduction != "" {
			if err := r.SetQueryParam("isLunchDeduction", qIsLunchDeduction); err != nil {
				return err
			}
		}

	}

	if o.Location != nil {

		// query param location
		var qrLocation string
		if o.Location != nil {
			qrLocation = *o.Location
		}
		qLocation := qrLocation
		if qLocation != "" {
			if err := r.SetQueryParam("location", qLocation); err != nil {
				return err
			}
		}

	}

	if o.OvernightAccommodation != nil {

		// query param overnightAccommodation
		var qrOvernightAccommodation string
		if o.OvernightAccommodation != nil {
			qrOvernightAccommodation = *o.OvernightAccommodation
		}
		qOvernightAccommodation := qrOvernightAccommodation
		if qOvernightAccommodation != "" {
			if err := r.SetQueryParam("overnightAccommodation", qOvernightAccommodation); err != nil {
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