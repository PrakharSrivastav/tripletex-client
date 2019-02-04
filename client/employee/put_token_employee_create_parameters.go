// Code generated by go-swagger; DO NOT EDIT.

package employee

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

// NewPutTokenEmployeeCreateParams creates a new PutTokenEmployeeCreateParams object
// with the default values initialized.
func NewPutTokenEmployeeCreateParams() *PutTokenEmployeeCreateParams {
	var ()
	return &PutTokenEmployeeCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutTokenEmployeeCreateParamsWithTimeout creates a new PutTokenEmployeeCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutTokenEmployeeCreateParamsWithTimeout(timeout time.Duration) *PutTokenEmployeeCreateParams {
	var ()
	return &PutTokenEmployeeCreateParams{

		timeout: timeout,
	}
}

// NewPutTokenEmployeeCreateParamsWithContext creates a new PutTokenEmployeeCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutTokenEmployeeCreateParamsWithContext(ctx context.Context) *PutTokenEmployeeCreateParams {
	var ()
	return &PutTokenEmployeeCreateParams{

		Context: ctx,
	}
}

// NewPutTokenEmployeeCreateParamsWithHTTPClient creates a new PutTokenEmployeeCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutTokenEmployeeCreateParamsWithHTTPClient(client *http.Client) *PutTokenEmployeeCreateParams {
	var ()
	return &PutTokenEmployeeCreateParams{
		HTTPClient: client,
	}
}

/*PutTokenEmployeeCreateParams contains all the parameters to send to the API endpoint
for the put token employee create operation typically these are written to a http.Request
*/
type PutTokenEmployeeCreateParams struct {

	/*CompanyOwned
	  Is the key company owned

	*/
	CompanyOwned bool
	/*ConsumerName
	  The name of the consumer

	*/
	ConsumerName string
	/*EmployeeID
	  The id of the employee

	*/
	EmployeeID int32
	/*ExpirationDate
	  Expiration date for the employeeToken

	*/
	ExpirationDate string
	/*TokenName
	  A user defined name for the new token

	*/
	TokenName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put token employee create params
func (o *PutTokenEmployeeCreateParams) WithTimeout(timeout time.Duration) *PutTokenEmployeeCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put token employee create params
func (o *PutTokenEmployeeCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put token employee create params
func (o *PutTokenEmployeeCreateParams) WithContext(ctx context.Context) *PutTokenEmployeeCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put token employee create params
func (o *PutTokenEmployeeCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put token employee create params
func (o *PutTokenEmployeeCreateParams) WithHTTPClient(client *http.Client) *PutTokenEmployeeCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put token employee create params
func (o *PutTokenEmployeeCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompanyOwned adds the companyOwned to the put token employee create params
func (o *PutTokenEmployeeCreateParams) WithCompanyOwned(companyOwned bool) *PutTokenEmployeeCreateParams {
	o.SetCompanyOwned(companyOwned)
	return o
}

// SetCompanyOwned adds the companyOwned to the put token employee create params
func (o *PutTokenEmployeeCreateParams) SetCompanyOwned(companyOwned bool) {
	o.CompanyOwned = companyOwned
}

// WithConsumerName adds the consumerName to the put token employee create params
func (o *PutTokenEmployeeCreateParams) WithConsumerName(consumerName string) *PutTokenEmployeeCreateParams {
	o.SetConsumerName(consumerName)
	return o
}

// SetConsumerName adds the consumerName to the put token employee create params
func (o *PutTokenEmployeeCreateParams) SetConsumerName(consumerName string) {
	o.ConsumerName = consumerName
}

// WithEmployeeID adds the employeeID to the put token employee create params
func (o *PutTokenEmployeeCreateParams) WithEmployeeID(employeeID int32) *PutTokenEmployeeCreateParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the put token employee create params
func (o *PutTokenEmployeeCreateParams) SetEmployeeID(employeeID int32) {
	o.EmployeeID = employeeID
}

// WithExpirationDate adds the expirationDate to the put token employee create params
func (o *PutTokenEmployeeCreateParams) WithExpirationDate(expirationDate string) *PutTokenEmployeeCreateParams {
	o.SetExpirationDate(expirationDate)
	return o
}

// SetExpirationDate adds the expirationDate to the put token employee create params
func (o *PutTokenEmployeeCreateParams) SetExpirationDate(expirationDate string) {
	o.ExpirationDate = expirationDate
}

// WithTokenName adds the tokenName to the put token employee create params
func (o *PutTokenEmployeeCreateParams) WithTokenName(tokenName string) *PutTokenEmployeeCreateParams {
	o.SetTokenName(tokenName)
	return o
}

// SetTokenName adds the tokenName to the put token employee create params
func (o *PutTokenEmployeeCreateParams) SetTokenName(tokenName string) {
	o.TokenName = tokenName
}

// WriteToRequest writes these params to a swagger request
func (o *PutTokenEmployeeCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param companyOwned
	qrCompanyOwned := o.CompanyOwned
	qCompanyOwned := swag.FormatBool(qrCompanyOwned)
	if qCompanyOwned != "" {
		if err := r.SetQueryParam("companyOwned", qCompanyOwned); err != nil {
			return err
		}
	}

	// query param consumerName
	qrConsumerName := o.ConsumerName
	qConsumerName := qrConsumerName
	if qConsumerName != "" {
		if err := r.SetQueryParam("consumerName", qConsumerName); err != nil {
			return err
		}
	}

	// query param employeeId
	qrEmployeeID := o.EmployeeID
	qEmployeeID := swag.FormatInt32(qrEmployeeID)
	if qEmployeeID != "" {
		if err := r.SetQueryParam("employeeId", qEmployeeID); err != nil {
			return err
		}
	}

	// query param expirationDate
	qrExpirationDate := o.ExpirationDate
	qExpirationDate := qrExpirationDate
	if qExpirationDate != "" {
		if err := r.SetQueryParam("expirationDate", qExpirationDate); err != nil {
			return err
		}
	}

	// query param tokenName
	qrTokenName := o.TokenName
	qTokenName := qrTokenName
	if qTokenName != "" {
		if err := r.SetQueryParam("tokenName", qTokenName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
