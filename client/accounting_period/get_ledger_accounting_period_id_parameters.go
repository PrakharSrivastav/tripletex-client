// Code generated by go-swagger; DO NOT EDIT.

package accounting_period

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

// NewGetLedgerAccountingPeriodIDParams creates a new GetLedgerAccountingPeriodIDParams object
// with the default values initialized.
func NewGetLedgerAccountingPeriodIDParams() *GetLedgerAccountingPeriodIDParams {
	var ()
	return &GetLedgerAccountingPeriodIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLedgerAccountingPeriodIDParamsWithTimeout creates a new GetLedgerAccountingPeriodIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLedgerAccountingPeriodIDParamsWithTimeout(timeout time.Duration) *GetLedgerAccountingPeriodIDParams {
	var ()
	return &GetLedgerAccountingPeriodIDParams{

		timeout: timeout,
	}
}

// NewGetLedgerAccountingPeriodIDParamsWithContext creates a new GetLedgerAccountingPeriodIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLedgerAccountingPeriodIDParamsWithContext(ctx context.Context) *GetLedgerAccountingPeriodIDParams {
	var ()
	return &GetLedgerAccountingPeriodIDParams{

		Context: ctx,
	}
}

// NewGetLedgerAccountingPeriodIDParamsWithHTTPClient creates a new GetLedgerAccountingPeriodIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLedgerAccountingPeriodIDParamsWithHTTPClient(client *http.Client) *GetLedgerAccountingPeriodIDParams {
	var ()
	return &GetLedgerAccountingPeriodIDParams{
		HTTPClient: client,
	}
}

/*GetLedgerAccountingPeriodIDParams contains all the parameters to send to the API endpoint
for the get ledger accounting period ID operation typically these are written to a http.Request
*/
type GetLedgerAccountingPeriodIDParams struct {

	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ledger accounting period ID params
func (o *GetLedgerAccountingPeriodIDParams) WithTimeout(timeout time.Duration) *GetLedgerAccountingPeriodIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ledger accounting period ID params
func (o *GetLedgerAccountingPeriodIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ledger accounting period ID params
func (o *GetLedgerAccountingPeriodIDParams) WithContext(ctx context.Context) *GetLedgerAccountingPeriodIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ledger accounting period ID params
func (o *GetLedgerAccountingPeriodIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ledger accounting period ID params
func (o *GetLedgerAccountingPeriodIDParams) WithHTTPClient(client *http.Client) *GetLedgerAccountingPeriodIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ledger accounting period ID params
func (o *GetLedgerAccountingPeriodIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get ledger accounting period ID params
func (o *GetLedgerAccountingPeriodIDParams) WithFields(fields *string) *GetLedgerAccountingPeriodIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get ledger accounting period ID params
func (o *GetLedgerAccountingPeriodIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get ledger accounting period ID params
func (o *GetLedgerAccountingPeriodIDParams) WithID(id int32) *GetLedgerAccountingPeriodIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get ledger accounting period ID params
func (o *GetLedgerAccountingPeriodIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLedgerAccountingPeriodIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
