// Code generated by go-swagger; DO NOT EDIT.

package reconciliation

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

// NewGetBankReconciliationLastClosedParams creates a new GetBankReconciliationLastClosedParams object
// with the default values initialized.
func NewGetBankReconciliationLastClosedParams() *GetBankReconciliationLastClosedParams {
	var ()
	return &GetBankReconciliationLastClosedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBankReconciliationLastClosedParamsWithTimeout creates a new GetBankReconciliationLastClosedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBankReconciliationLastClosedParamsWithTimeout(timeout time.Duration) *GetBankReconciliationLastClosedParams {
	var ()
	return &GetBankReconciliationLastClosedParams{

		timeout: timeout,
	}
}

// NewGetBankReconciliationLastClosedParamsWithContext creates a new GetBankReconciliationLastClosedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBankReconciliationLastClosedParamsWithContext(ctx context.Context) *GetBankReconciliationLastClosedParams {
	var ()
	return &GetBankReconciliationLastClosedParams{

		Context: ctx,
	}
}

// NewGetBankReconciliationLastClosedParamsWithHTTPClient creates a new GetBankReconciliationLastClosedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBankReconciliationLastClosedParamsWithHTTPClient(client *http.Client) *GetBankReconciliationLastClosedParams {
	var ()
	return &GetBankReconciliationLastClosedParams{
		HTTPClient: client,
	}
}

/*GetBankReconciliationLastClosedParams contains all the parameters to send to the API endpoint
for the get bank reconciliation last closed operation typically these are written to a http.Request
*/
type GetBankReconciliationLastClosedParams struct {

	/*AccountID
	  Account ID

	*/
	AccountID int32
	/*After
	  Format is yyyy-MM-dd

	*/
	After *string
	/*Fields
	  Fields filter pattern

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get bank reconciliation last closed params
func (o *GetBankReconciliationLastClosedParams) WithTimeout(timeout time.Duration) *GetBankReconciliationLastClosedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bank reconciliation last closed params
func (o *GetBankReconciliationLastClosedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bank reconciliation last closed params
func (o *GetBankReconciliationLastClosedParams) WithContext(ctx context.Context) *GetBankReconciliationLastClosedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bank reconciliation last closed params
func (o *GetBankReconciliationLastClosedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bank reconciliation last closed params
func (o *GetBankReconciliationLastClosedParams) WithHTTPClient(client *http.Client) *GetBankReconciliationLastClosedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bank reconciliation last closed params
func (o *GetBankReconciliationLastClosedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get bank reconciliation last closed params
func (o *GetBankReconciliationLastClosedParams) WithAccountID(accountID int32) *GetBankReconciliationLastClosedParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get bank reconciliation last closed params
func (o *GetBankReconciliationLastClosedParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithAfter adds the after to the get bank reconciliation last closed params
func (o *GetBankReconciliationLastClosedParams) WithAfter(after *string) *GetBankReconciliationLastClosedParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get bank reconciliation last closed params
func (o *GetBankReconciliationLastClosedParams) SetAfter(after *string) {
	o.After = after
}

// WithFields adds the fields to the get bank reconciliation last closed params
func (o *GetBankReconciliationLastClosedParams) WithFields(fields *string) *GetBankReconciliationLastClosedParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get bank reconciliation last closed params
func (o *GetBankReconciliationLastClosedParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetBankReconciliationLastClosedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param accountId
	qrAccountID := o.AccountID
	qAccountID := swag.FormatInt32(qrAccountID)
	if qAccountID != "" {
		if err := r.SetQueryParam("accountId", qAccountID); err != nil {
			return err
		}
	}

	if o.After != nil {

		// query param after
		var qrAfter string
		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {
			if err := r.SetQueryParam("after", qAfter); err != nil {
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
