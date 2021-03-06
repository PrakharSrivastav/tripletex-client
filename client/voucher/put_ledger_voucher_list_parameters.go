// Code generated by go-swagger; DO NOT EDIT.

package voucher

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

	models "github.com/PrakharSrivastav/Petstore/models"
)

// NewPutLedgerVoucherListParams creates a new PutLedgerVoucherListParams object
// with the default values initialized.
func NewPutLedgerVoucherListParams() *PutLedgerVoucherListParams {
	var (
		sendToLedgerDefault = bool(true)
	)
	return &PutLedgerVoucherListParams{
		SendToLedger: &sendToLedgerDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLedgerVoucherListParamsWithTimeout creates a new PutLedgerVoucherListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLedgerVoucherListParamsWithTimeout(timeout time.Duration) *PutLedgerVoucherListParams {
	var (
		sendToLedgerDefault = bool(true)
	)
	return &PutLedgerVoucherListParams{
		SendToLedger: &sendToLedgerDefault,

		timeout: timeout,
	}
}

// NewPutLedgerVoucherListParamsWithContext creates a new PutLedgerVoucherListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLedgerVoucherListParamsWithContext(ctx context.Context) *PutLedgerVoucherListParams {
	var (
		sendToLedgerDefault = bool(true)
	)
	return &PutLedgerVoucherListParams{
		SendToLedger: &sendToLedgerDefault,

		Context: ctx,
	}
}

// NewPutLedgerVoucherListParamsWithHTTPClient creates a new PutLedgerVoucherListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLedgerVoucherListParamsWithHTTPClient(client *http.Client) *PutLedgerVoucherListParams {
	var (
		sendToLedgerDefault = bool(true)
	)
	return &PutLedgerVoucherListParams{
		SendToLedger: &sendToLedgerDefault,
		HTTPClient:   client,
	}
}

/*PutLedgerVoucherListParams contains all the parameters to send to the API endpoint
for the put ledger voucher list operation typically these are written to a http.Request
*/
type PutLedgerVoucherListParams struct {

	/*Body
	  JSON representing updates to object. Should have ID and version set.

	*/
	Body []*models.Voucher
	/*SendToLedger
	  Should the voucher be sent to ledger? Requires the "Advanced Voucher" permission.

	*/
	SendToLedger *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put ledger voucher list params
func (o *PutLedgerVoucherListParams) WithTimeout(timeout time.Duration) *PutLedgerVoucherListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put ledger voucher list params
func (o *PutLedgerVoucherListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put ledger voucher list params
func (o *PutLedgerVoucherListParams) WithContext(ctx context.Context) *PutLedgerVoucherListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put ledger voucher list params
func (o *PutLedgerVoucherListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put ledger voucher list params
func (o *PutLedgerVoucherListParams) WithHTTPClient(client *http.Client) *PutLedgerVoucherListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put ledger voucher list params
func (o *PutLedgerVoucherListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put ledger voucher list params
func (o *PutLedgerVoucherListParams) WithBody(body []*models.Voucher) *PutLedgerVoucherListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put ledger voucher list params
func (o *PutLedgerVoucherListParams) SetBody(body []*models.Voucher) {
	o.Body = body
}

// WithSendToLedger adds the sendToLedger to the put ledger voucher list params
func (o *PutLedgerVoucherListParams) WithSendToLedger(sendToLedger *bool) *PutLedgerVoucherListParams {
	o.SetSendToLedger(sendToLedger)
	return o
}

// SetSendToLedger adds the sendToLedger to the put ledger voucher list params
func (o *PutLedgerVoucherListParams) SetSendToLedger(sendToLedger *bool) {
	o.SendToLedger = sendToLedger
}

// WriteToRequest writes these params to a swagger request
func (o *PutLedgerVoucherListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.SendToLedger != nil {

		// query param sendToLedger
		var qrSendToLedger bool
		if o.SendToLedger != nil {
			qrSendToLedger = *o.SendToLedger
		}
		qSendToLedger := swag.FormatBool(qrSendToLedger)
		if qSendToLedger != "" {
			if err := r.SetQueryParam("sendToLedger", qSendToLedger); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
