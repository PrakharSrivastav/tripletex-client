// Code generated by go-swagger; DO NOT EDIT.

package travel_expense

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
)

// NewPutTravelExpenseUndeliverParams creates a new PutTravelExpenseUndeliverParams object
// with the default values initialized.
func NewPutTravelExpenseUndeliverParams() *PutTravelExpenseUndeliverParams {
	var ()
	return &PutTravelExpenseUndeliverParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutTravelExpenseUndeliverParamsWithTimeout creates a new PutTravelExpenseUndeliverParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutTravelExpenseUndeliverParamsWithTimeout(timeout time.Duration) *PutTravelExpenseUndeliverParams {
	var ()
	return &PutTravelExpenseUndeliverParams{

		timeout: timeout,
	}
}

// NewPutTravelExpenseUndeliverParamsWithContext creates a new PutTravelExpenseUndeliverParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutTravelExpenseUndeliverParamsWithContext(ctx context.Context) *PutTravelExpenseUndeliverParams {
	var ()
	return &PutTravelExpenseUndeliverParams{

		Context: ctx,
	}
}

// NewPutTravelExpenseUndeliverParamsWithHTTPClient creates a new PutTravelExpenseUndeliverParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutTravelExpenseUndeliverParamsWithHTTPClient(client *http.Client) *PutTravelExpenseUndeliverParams {
	var ()
	return &PutTravelExpenseUndeliverParams{
		HTTPClient: client,
	}
}

/*PutTravelExpenseUndeliverParams contains all the parameters to send to the API endpoint
for the put travel expense undeliver operation typically these are written to a http.Request
*/
type PutTravelExpenseUndeliverParams struct {

	/*ID
	  ID of the elements

	*/
	ID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put travel expense undeliver params
func (o *PutTravelExpenseUndeliverParams) WithTimeout(timeout time.Duration) *PutTravelExpenseUndeliverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put travel expense undeliver params
func (o *PutTravelExpenseUndeliverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put travel expense undeliver params
func (o *PutTravelExpenseUndeliverParams) WithContext(ctx context.Context) *PutTravelExpenseUndeliverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put travel expense undeliver params
func (o *PutTravelExpenseUndeliverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put travel expense undeliver params
func (o *PutTravelExpenseUndeliverParams) WithHTTPClient(client *http.Client) *PutTravelExpenseUndeliverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put travel expense undeliver params
func (o *PutTravelExpenseUndeliverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put travel expense undeliver params
func (o *PutTravelExpenseUndeliverParams) WithID(id *string) *PutTravelExpenseUndeliverParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put travel expense undeliver params
func (o *PutTravelExpenseUndeliverParams) SetID(id *string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutTravelExpenseUndeliverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
