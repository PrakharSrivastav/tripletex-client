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

// NewDeleteBankReconciliationIDParams creates a new DeleteBankReconciliationIDParams object
// with the default values initialized.
func NewDeleteBankReconciliationIDParams() *DeleteBankReconciliationIDParams {
	var ()
	return &DeleteBankReconciliationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBankReconciliationIDParamsWithTimeout creates a new DeleteBankReconciliationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteBankReconciliationIDParamsWithTimeout(timeout time.Duration) *DeleteBankReconciliationIDParams {
	var ()
	return &DeleteBankReconciliationIDParams{

		timeout: timeout,
	}
}

// NewDeleteBankReconciliationIDParamsWithContext creates a new DeleteBankReconciliationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteBankReconciliationIDParamsWithContext(ctx context.Context) *DeleteBankReconciliationIDParams {
	var ()
	return &DeleteBankReconciliationIDParams{

		Context: ctx,
	}
}

// NewDeleteBankReconciliationIDParamsWithHTTPClient creates a new DeleteBankReconciliationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteBankReconciliationIDParamsWithHTTPClient(client *http.Client) *DeleteBankReconciliationIDParams {
	var ()
	return &DeleteBankReconciliationIDParams{
		HTTPClient: client,
	}
}

/*DeleteBankReconciliationIDParams contains all the parameters to send to the API endpoint
for the delete bank reconciliation ID operation typically these are written to a http.Request
*/
type DeleteBankReconciliationIDParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete bank reconciliation ID params
func (o *DeleteBankReconciliationIDParams) WithTimeout(timeout time.Duration) *DeleteBankReconciliationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete bank reconciliation ID params
func (o *DeleteBankReconciliationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete bank reconciliation ID params
func (o *DeleteBankReconciliationIDParams) WithContext(ctx context.Context) *DeleteBankReconciliationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete bank reconciliation ID params
func (o *DeleteBankReconciliationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete bank reconciliation ID params
func (o *DeleteBankReconciliationIDParams) WithHTTPClient(client *http.Client) *DeleteBankReconciliationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete bank reconciliation ID params
func (o *DeleteBankReconciliationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete bank reconciliation ID params
func (o *DeleteBankReconciliationIDParams) WithID(id int32) *DeleteBankReconciliationIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete bank reconciliation ID params
func (o *DeleteBankReconciliationIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBankReconciliationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
