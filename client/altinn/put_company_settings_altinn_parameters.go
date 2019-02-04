// Code generated by go-swagger; DO NOT EDIT.

package altinn

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

// NewPutCompanySettingsAltinnParams creates a new PutCompanySettingsAltinnParams object
// with the default values initialized.
func NewPutCompanySettingsAltinnParams() *PutCompanySettingsAltinnParams {
	var ()
	return &PutCompanySettingsAltinnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCompanySettingsAltinnParamsWithTimeout creates a new PutCompanySettingsAltinnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCompanySettingsAltinnParamsWithTimeout(timeout time.Duration) *PutCompanySettingsAltinnParams {
	var ()
	return &PutCompanySettingsAltinnParams{

		timeout: timeout,
	}
}

// NewPutCompanySettingsAltinnParamsWithContext creates a new PutCompanySettingsAltinnParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCompanySettingsAltinnParamsWithContext(ctx context.Context) *PutCompanySettingsAltinnParams {
	var ()
	return &PutCompanySettingsAltinnParams{

		Context: ctx,
	}
}

// NewPutCompanySettingsAltinnParamsWithHTTPClient creates a new PutCompanySettingsAltinnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCompanySettingsAltinnParamsWithHTTPClient(client *http.Client) *PutCompanySettingsAltinnParams {
	var ()
	return &PutCompanySettingsAltinnParams{
		HTTPClient: client,
	}
}

/*PutCompanySettingsAltinnParams contains all the parameters to send to the API endpoint
for the put company settings altinn operation typically these are written to a http.Request
*/
type PutCompanySettingsAltinnParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.AltinnCompanyModule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put company settings altinn params
func (o *PutCompanySettingsAltinnParams) WithTimeout(timeout time.Duration) *PutCompanySettingsAltinnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put company settings altinn params
func (o *PutCompanySettingsAltinnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put company settings altinn params
func (o *PutCompanySettingsAltinnParams) WithContext(ctx context.Context) *PutCompanySettingsAltinnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put company settings altinn params
func (o *PutCompanySettingsAltinnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put company settings altinn params
func (o *PutCompanySettingsAltinnParams) WithHTTPClient(client *http.Client) *PutCompanySettingsAltinnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put company settings altinn params
func (o *PutCompanySettingsAltinnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put company settings altinn params
func (o *PutCompanySettingsAltinnParams) WithBody(body *models.AltinnCompanyModule) *PutCompanySettingsAltinnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put company settings altinn params
func (o *PutCompanySettingsAltinnParams) SetBody(body *models.AltinnCompanyModule) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutCompanySettingsAltinnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
