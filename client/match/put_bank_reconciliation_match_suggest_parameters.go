// Code generated by go-swagger; DO NOT EDIT.

package match

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

// NewPutBankReconciliationMatchSuggestParams creates a new PutBankReconciliationMatchSuggestParams object
// with the default values initialized.
func NewPutBankReconciliationMatchSuggestParams() *PutBankReconciliationMatchSuggestParams {
	var ()
	return &PutBankReconciliationMatchSuggestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutBankReconciliationMatchSuggestParamsWithTimeout creates a new PutBankReconciliationMatchSuggestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutBankReconciliationMatchSuggestParamsWithTimeout(timeout time.Duration) *PutBankReconciliationMatchSuggestParams {
	var ()
	return &PutBankReconciliationMatchSuggestParams{

		timeout: timeout,
	}
}

// NewPutBankReconciliationMatchSuggestParamsWithContext creates a new PutBankReconciliationMatchSuggestParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutBankReconciliationMatchSuggestParamsWithContext(ctx context.Context) *PutBankReconciliationMatchSuggestParams {
	var ()
	return &PutBankReconciliationMatchSuggestParams{

		Context: ctx,
	}
}

// NewPutBankReconciliationMatchSuggestParamsWithHTTPClient creates a new PutBankReconciliationMatchSuggestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutBankReconciliationMatchSuggestParamsWithHTTPClient(client *http.Client) *PutBankReconciliationMatchSuggestParams {
	var ()
	return &PutBankReconciliationMatchSuggestParams{
		HTTPClient: client,
	}
}

/*PutBankReconciliationMatchSuggestParams contains all the parameters to send to the API endpoint
for the put bank reconciliation match suggest operation typically these are written to a http.Request
*/
type PutBankReconciliationMatchSuggestParams struct {

	/*BankReconciliationID
	  Element ID

	*/
	BankReconciliationID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put bank reconciliation match suggest params
func (o *PutBankReconciliationMatchSuggestParams) WithTimeout(timeout time.Duration) *PutBankReconciliationMatchSuggestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put bank reconciliation match suggest params
func (o *PutBankReconciliationMatchSuggestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put bank reconciliation match suggest params
func (o *PutBankReconciliationMatchSuggestParams) WithContext(ctx context.Context) *PutBankReconciliationMatchSuggestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put bank reconciliation match suggest params
func (o *PutBankReconciliationMatchSuggestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put bank reconciliation match suggest params
func (o *PutBankReconciliationMatchSuggestParams) WithHTTPClient(client *http.Client) *PutBankReconciliationMatchSuggestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put bank reconciliation match suggest params
func (o *PutBankReconciliationMatchSuggestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBankReconciliationID adds the bankReconciliationID to the put bank reconciliation match suggest params
func (o *PutBankReconciliationMatchSuggestParams) WithBankReconciliationID(bankReconciliationID int32) *PutBankReconciliationMatchSuggestParams {
	o.SetBankReconciliationID(bankReconciliationID)
	return o
}

// SetBankReconciliationID adds the bankReconciliationId to the put bank reconciliation match suggest params
func (o *PutBankReconciliationMatchSuggestParams) SetBankReconciliationID(bankReconciliationID int32) {
	o.BankReconciliationID = bankReconciliationID
}

// WriteToRequest writes these params to a swagger request
func (o *PutBankReconciliationMatchSuggestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param bankReconciliationId
	qrBankReconciliationID := o.BankReconciliationID
	qBankReconciliationID := swag.FormatInt32(qrBankReconciliationID)
	if qBankReconciliationID != "" {
		if err := r.SetQueryParam("bankReconciliationId", qBankReconciliationID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}