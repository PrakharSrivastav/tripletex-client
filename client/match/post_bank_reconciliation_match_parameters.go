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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// NewPostBankReconciliationMatchParams creates a new PostBankReconciliationMatchParams object
// with the default values initialized.
func NewPostBankReconciliationMatchParams() *PostBankReconciliationMatchParams {
	var ()
	return &PostBankReconciliationMatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostBankReconciliationMatchParamsWithTimeout creates a new PostBankReconciliationMatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostBankReconciliationMatchParamsWithTimeout(timeout time.Duration) *PostBankReconciliationMatchParams {
	var ()
	return &PostBankReconciliationMatchParams{

		timeout: timeout,
	}
}

// NewPostBankReconciliationMatchParamsWithContext creates a new PostBankReconciliationMatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostBankReconciliationMatchParamsWithContext(ctx context.Context) *PostBankReconciliationMatchParams {
	var ()
	return &PostBankReconciliationMatchParams{

		Context: ctx,
	}
}

// NewPostBankReconciliationMatchParamsWithHTTPClient creates a new PostBankReconciliationMatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostBankReconciliationMatchParamsWithHTTPClient(client *http.Client) *PostBankReconciliationMatchParams {
	var ()
	return &PostBankReconciliationMatchParams{
		HTTPClient: client,
	}
}

/*PostBankReconciliationMatchParams contains all the parameters to send to the API endpoint
for the post bank reconciliation match operation typically these are written to a http.Request
*/
type PostBankReconciliationMatchParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.BankReconciliationMatch

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post bank reconciliation match params
func (o *PostBankReconciliationMatchParams) WithTimeout(timeout time.Duration) *PostBankReconciliationMatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post bank reconciliation match params
func (o *PostBankReconciliationMatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post bank reconciliation match params
func (o *PostBankReconciliationMatchParams) WithContext(ctx context.Context) *PostBankReconciliationMatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post bank reconciliation match params
func (o *PostBankReconciliationMatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post bank reconciliation match params
func (o *PostBankReconciliationMatchParams) WithHTTPClient(client *http.Client) *PostBankReconciliationMatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post bank reconciliation match params
func (o *PostBankReconciliationMatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post bank reconciliation match params
func (o *PostBankReconciliationMatchParams) WithBody(body *models.BankReconciliationMatch) *PostBankReconciliationMatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post bank reconciliation match params
func (o *PostBankReconciliationMatchParams) SetBody(body *models.BankReconciliationMatch) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostBankReconciliationMatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
