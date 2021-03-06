// Code generated by go-swagger; DO NOT EDIT.

package holiday

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

// NewPutSalarySettingsHolidayListParams creates a new PutSalarySettingsHolidayListParams object
// with the default values initialized.
func NewPutSalarySettingsHolidayListParams() *PutSalarySettingsHolidayListParams {
	var ()
	return &PutSalarySettingsHolidayListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSalarySettingsHolidayListParamsWithTimeout creates a new PutSalarySettingsHolidayListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSalarySettingsHolidayListParamsWithTimeout(timeout time.Duration) *PutSalarySettingsHolidayListParams {
	var ()
	return &PutSalarySettingsHolidayListParams{

		timeout: timeout,
	}
}

// NewPutSalarySettingsHolidayListParamsWithContext creates a new PutSalarySettingsHolidayListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSalarySettingsHolidayListParamsWithContext(ctx context.Context) *PutSalarySettingsHolidayListParams {
	var ()
	return &PutSalarySettingsHolidayListParams{

		Context: ctx,
	}
}

// NewPutSalarySettingsHolidayListParamsWithHTTPClient creates a new PutSalarySettingsHolidayListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSalarySettingsHolidayListParamsWithHTTPClient(client *http.Client) *PutSalarySettingsHolidayListParams {
	var ()
	return &PutSalarySettingsHolidayListParams{
		HTTPClient: client,
	}
}

/*PutSalarySettingsHolidayListParams contains all the parameters to send to the API endpoint
for the put salary settings holiday list operation typically these are written to a http.Request
*/
type PutSalarySettingsHolidayListParams struct {

	/*Body
	  JSON representing updates to object. Should have ID and version set.

	*/
	Body []*models.CompanyHoliday

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put salary settings holiday list params
func (o *PutSalarySettingsHolidayListParams) WithTimeout(timeout time.Duration) *PutSalarySettingsHolidayListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put salary settings holiday list params
func (o *PutSalarySettingsHolidayListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put salary settings holiday list params
func (o *PutSalarySettingsHolidayListParams) WithContext(ctx context.Context) *PutSalarySettingsHolidayListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put salary settings holiday list params
func (o *PutSalarySettingsHolidayListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put salary settings holiday list params
func (o *PutSalarySettingsHolidayListParams) WithHTTPClient(client *http.Client) *PutSalarySettingsHolidayListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put salary settings holiday list params
func (o *PutSalarySettingsHolidayListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put salary settings holiday list params
func (o *PutSalarySettingsHolidayListParams) WithBody(body []*models.CompanyHoliday) *PutSalarySettingsHolidayListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put salary settings holiday list params
func (o *PutSalarySettingsHolidayListParams) SetBody(body []*models.CompanyHoliday) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutSalarySettingsHolidayListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
