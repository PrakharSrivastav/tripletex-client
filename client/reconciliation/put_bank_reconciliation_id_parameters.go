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

// NewPutBankReconciliationIDParams creates a new PutBankReconciliationIDParams object
// with the default values initialized.
func NewPutBankReconciliationIDParams() *PutBankReconciliationIDParams {
	var ()
	return &PutBankReconciliationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutBankReconciliationIDParamsWithTimeout creates a new PutBankReconciliationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutBankReconciliationIDParamsWithTimeout(timeout time.Duration) *PutBankReconciliationIDParams {
	var ()
	return &PutBankReconciliationIDParams{

		timeout: timeout,
	}
}

// NewPutBankReconciliationIDParamsWithContext creates a new PutBankReconciliationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutBankReconciliationIDParamsWithContext(ctx context.Context) *PutBankReconciliationIDParams {
	var ()
	return &PutBankReconciliationIDParams{

		Context: ctx,
	}
}

// NewPutBankReconciliationIDParamsWithHTTPClient creates a new PutBankReconciliationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutBankReconciliationIDParamsWithHTTPClient(client *http.Client) *PutBankReconciliationIDParams {
	var ()
	return &PutBankReconciliationIDParams{
		HTTPClient: client,
	}
}

/*PutBankReconciliationIDParams contains all the parameters to send to the API endpoint
for the put bank reconciliation ID operation typically these are written to a http.Request
*/
type PutBankReconciliationIDParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.BankReconciliation
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put bank reconciliation ID params
func (o *PutBankReconciliationIDParams) WithTimeout(timeout time.Duration) *PutBankReconciliationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put bank reconciliation ID params
func (o *PutBankReconciliationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put bank reconciliation ID params
func (o *PutBankReconciliationIDParams) WithContext(ctx context.Context) *PutBankReconciliationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put bank reconciliation ID params
func (o *PutBankReconciliationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put bank reconciliation ID params
func (o *PutBankReconciliationIDParams) WithHTTPClient(client *http.Client) *PutBankReconciliationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put bank reconciliation ID params
func (o *PutBankReconciliationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put bank reconciliation ID params
func (o *PutBankReconciliationIDParams) WithBody(body *models.BankReconciliation) *PutBankReconciliationIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put bank reconciliation ID params
func (o *PutBankReconciliationIDParams) SetBody(body *models.BankReconciliation) {
	o.Body = body
}

// WithID adds the id to the put bank reconciliation ID params
func (o *PutBankReconciliationIDParams) WithID(id int32) *PutBankReconciliationIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put bank reconciliation ID params
func (o *PutBankReconciliationIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutBankReconciliationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
