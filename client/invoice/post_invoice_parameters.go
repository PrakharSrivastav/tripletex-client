// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewPostInvoiceParams creates a new PostInvoiceParams object
// with the default values initialized.
func NewPostInvoiceParams() *PostInvoiceParams {
	var (
		sendToCustomerDefault = bool(true)
	)
	return &PostInvoiceParams{
		SendToCustomer: &sendToCustomerDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostInvoiceParamsWithTimeout creates a new PostInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostInvoiceParamsWithTimeout(timeout time.Duration) *PostInvoiceParams {
	var (
		sendToCustomerDefault = bool(true)
	)
	return &PostInvoiceParams{
		SendToCustomer: &sendToCustomerDefault,

		timeout: timeout,
	}
}

// NewPostInvoiceParamsWithContext creates a new PostInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostInvoiceParamsWithContext(ctx context.Context) *PostInvoiceParams {
	var (
		sendToCustomerDefault = bool(true)
	)
	return &PostInvoiceParams{
		SendToCustomer: &sendToCustomerDefault,

		Context: ctx,
	}
}

// NewPostInvoiceParamsWithHTTPClient creates a new PostInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostInvoiceParamsWithHTTPClient(client *http.Client) *PostInvoiceParams {
	var (
		sendToCustomerDefault = bool(true)
	)
	return &PostInvoiceParams{
		SendToCustomer: &sendToCustomerDefault,
		HTTPClient:     client,
	}
}

/*PostInvoiceParams contains all the parameters to send to the API endpoint
for the post invoice operation typically these are written to a http.Request
*/
type PostInvoiceParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.Invoice
	/*PaidAmount
	  Paid amount to register prepayment of the invoice, in invoice currency. paymentTypeId and paidAmount are optional, but both must be provided if the invoice has already been paid.

	*/
	PaidAmount *float64
	/*PaymentTypeID
	  Payment type to register prepayment of the invoice. paymentTypeId and paidAmount are optional, but both must be provided if the invoice has already been paid.

	*/
	PaymentTypeID *int32
	/*SendToCustomer
	  Equals

	*/
	SendToCustomer *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post invoice params
func (o *PostInvoiceParams) WithTimeout(timeout time.Duration) *PostInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post invoice params
func (o *PostInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post invoice params
func (o *PostInvoiceParams) WithContext(ctx context.Context) *PostInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post invoice params
func (o *PostInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post invoice params
func (o *PostInvoiceParams) WithHTTPClient(client *http.Client) *PostInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post invoice params
func (o *PostInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post invoice params
func (o *PostInvoiceParams) WithBody(body *models.Invoice) *PostInvoiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post invoice params
func (o *PostInvoiceParams) SetBody(body *models.Invoice) {
	o.Body = body
}

// WithPaidAmount adds the paidAmount to the post invoice params
func (o *PostInvoiceParams) WithPaidAmount(paidAmount *float64) *PostInvoiceParams {
	o.SetPaidAmount(paidAmount)
	return o
}

// SetPaidAmount adds the paidAmount to the post invoice params
func (o *PostInvoiceParams) SetPaidAmount(paidAmount *float64) {
	o.PaidAmount = paidAmount
}

// WithPaymentTypeID adds the paymentTypeID to the post invoice params
func (o *PostInvoiceParams) WithPaymentTypeID(paymentTypeID *int32) *PostInvoiceParams {
	o.SetPaymentTypeID(paymentTypeID)
	return o
}

// SetPaymentTypeID adds the paymentTypeId to the post invoice params
func (o *PostInvoiceParams) SetPaymentTypeID(paymentTypeID *int32) {
	o.PaymentTypeID = paymentTypeID
}

// WithSendToCustomer adds the sendToCustomer to the post invoice params
func (o *PostInvoiceParams) WithSendToCustomer(sendToCustomer *bool) *PostInvoiceParams {
	o.SetSendToCustomer(sendToCustomer)
	return o
}

// SetSendToCustomer adds the sendToCustomer to the post invoice params
func (o *PostInvoiceParams) SetSendToCustomer(sendToCustomer *bool) {
	o.SendToCustomer = sendToCustomer
}

// WriteToRequest writes these params to a swagger request
func (o *PostInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.PaidAmount != nil {

		// query param paidAmount
		var qrPaidAmount float64
		if o.PaidAmount != nil {
			qrPaidAmount = *o.PaidAmount
		}
		qPaidAmount := swag.FormatFloat64(qrPaidAmount)
		if qPaidAmount != "" {
			if err := r.SetQueryParam("paidAmount", qPaidAmount); err != nil {
				return err
			}
		}

	}

	if o.PaymentTypeID != nil {

		// query param paymentTypeId
		var qrPaymentTypeID int32
		if o.PaymentTypeID != nil {
			qrPaymentTypeID = *o.PaymentTypeID
		}
		qPaymentTypeID := swag.FormatInt32(qrPaymentTypeID)
		if qPaymentTypeID != "" {
			if err := r.SetQueryParam("paymentTypeId", qPaymentTypeID); err != nil {
				return err
			}
		}

	}

	if o.SendToCustomer != nil {

		// query param sendToCustomer
		var qrSendToCustomer bool
		if o.SendToCustomer != nil {
			qrSendToCustomer = *o.SendToCustomer
		}
		qSendToCustomer := swag.FormatBool(qrSendToCustomer)
		if qSendToCustomer != "" {
			if err := r.SetQueryParam("sendToCustomer", qSendToCustomer); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
