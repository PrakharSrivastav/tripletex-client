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

	models "github.com/PrakharSrivastav/Petstore/models"
)

// NewPutBankReconciliationIDAdjustmentParams creates a new PutBankReconciliationIDAdjustmentParams object
// with the default values initialized.
func NewPutBankReconciliationIDAdjustmentParams() *PutBankReconciliationIDAdjustmentParams {
	var ()
	return &PutBankReconciliationIDAdjustmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutBankReconciliationIDAdjustmentParamsWithTimeout creates a new PutBankReconciliationIDAdjustmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutBankReconciliationIDAdjustmentParamsWithTimeout(timeout time.Duration) *PutBankReconciliationIDAdjustmentParams {
	var ()
	return &PutBankReconciliationIDAdjustmentParams{

		timeout: timeout,
	}
}

// NewPutBankReconciliationIDAdjustmentParamsWithContext creates a new PutBankReconciliationIDAdjustmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutBankReconciliationIDAdjustmentParamsWithContext(ctx context.Context) *PutBankReconciliationIDAdjustmentParams {
	var ()
	return &PutBankReconciliationIDAdjustmentParams{

		Context: ctx,
	}
}

// NewPutBankReconciliationIDAdjustmentParamsWithHTTPClient creates a new PutBankReconciliationIDAdjustmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutBankReconciliationIDAdjustmentParamsWithHTTPClient(client *http.Client) *PutBankReconciliationIDAdjustmentParams {
	var ()
	return &PutBankReconciliationIDAdjustmentParams{
		HTTPClient: client,
	}
}

/*PutBankReconciliationIDAdjustmentParams contains all the parameters to send to the API endpoint
for the put bank reconciliation ID adjustment operation typically these are written to a http.Request
*/
type PutBankReconciliationIDAdjustmentParams struct {

	/*Body
	  Adjustments

	*/
	Body []*models.BankReconciliationAdjustment
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put bank reconciliation ID adjustment params
func (o *PutBankReconciliationIDAdjustmentParams) WithTimeout(timeout time.Duration) *PutBankReconciliationIDAdjustmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put bank reconciliation ID adjustment params
func (o *PutBankReconciliationIDAdjustmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put bank reconciliation ID adjustment params
func (o *PutBankReconciliationIDAdjustmentParams) WithContext(ctx context.Context) *PutBankReconciliationIDAdjustmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put bank reconciliation ID adjustment params
func (o *PutBankReconciliationIDAdjustmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put bank reconciliation ID adjustment params
func (o *PutBankReconciliationIDAdjustmentParams) WithHTTPClient(client *http.Client) *PutBankReconciliationIDAdjustmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put bank reconciliation ID adjustment params
func (o *PutBankReconciliationIDAdjustmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put bank reconciliation ID adjustment params
func (o *PutBankReconciliationIDAdjustmentParams) WithBody(body []*models.BankReconciliationAdjustment) *PutBankReconciliationIDAdjustmentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put bank reconciliation ID adjustment params
func (o *PutBankReconciliationIDAdjustmentParams) SetBody(body []*models.BankReconciliationAdjustment) {
	o.Body = body
}

// WithID adds the id to the put bank reconciliation ID adjustment params
func (o *PutBankReconciliationIDAdjustmentParams) WithID(id int32) *PutBankReconciliationIDAdjustmentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put bank reconciliation ID adjustment params
func (o *PutBankReconciliationIDAdjustmentParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutBankReconciliationIDAdjustmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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