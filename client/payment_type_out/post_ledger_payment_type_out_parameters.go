// Code generated by go-swagger; DO NOT EDIT.

package payment_type_out

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

// NewPostLedgerPaymentTypeOutParams creates a new PostLedgerPaymentTypeOutParams object
// with the default values initialized.
func NewPostLedgerPaymentTypeOutParams() *PostLedgerPaymentTypeOutParams {
	var ()
	return &PostLedgerPaymentTypeOutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLedgerPaymentTypeOutParamsWithTimeout creates a new PostLedgerPaymentTypeOutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLedgerPaymentTypeOutParamsWithTimeout(timeout time.Duration) *PostLedgerPaymentTypeOutParams {
	var ()
	return &PostLedgerPaymentTypeOutParams{

		timeout: timeout,
	}
}

// NewPostLedgerPaymentTypeOutParamsWithContext creates a new PostLedgerPaymentTypeOutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLedgerPaymentTypeOutParamsWithContext(ctx context.Context) *PostLedgerPaymentTypeOutParams {
	var ()
	return &PostLedgerPaymentTypeOutParams{

		Context: ctx,
	}
}

// NewPostLedgerPaymentTypeOutParamsWithHTTPClient creates a new PostLedgerPaymentTypeOutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLedgerPaymentTypeOutParamsWithHTTPClient(client *http.Client) *PostLedgerPaymentTypeOutParams {
	var ()
	return &PostLedgerPaymentTypeOutParams{
		HTTPClient: client,
	}
}

/*PostLedgerPaymentTypeOutParams contains all the parameters to send to the API endpoint
for the post ledger payment type out operation typically these are written to a http.Request
*/
type PostLedgerPaymentTypeOutParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.PaymentTypeOut

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ledger payment type out params
func (o *PostLedgerPaymentTypeOutParams) WithTimeout(timeout time.Duration) *PostLedgerPaymentTypeOutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ledger payment type out params
func (o *PostLedgerPaymentTypeOutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ledger payment type out params
func (o *PostLedgerPaymentTypeOutParams) WithContext(ctx context.Context) *PostLedgerPaymentTypeOutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ledger payment type out params
func (o *PostLedgerPaymentTypeOutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ledger payment type out params
func (o *PostLedgerPaymentTypeOutParams) WithHTTPClient(client *http.Client) *PostLedgerPaymentTypeOutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ledger payment type out params
func (o *PostLedgerPaymentTypeOutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post ledger payment type out params
func (o *PostLedgerPaymentTypeOutParams) WithBody(body *models.PaymentTypeOut) *PostLedgerPaymentTypeOutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post ledger payment type out params
func (o *PostLedgerPaymentTypeOutParams) SetBody(body *models.PaymentTypeOut) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostLedgerPaymentTypeOutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
