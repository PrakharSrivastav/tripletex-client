// Code generated by go-swagger; DO NOT EDIT.

package next_of_kin

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

// NewPostEmployeeNextOfKinParams creates a new PostEmployeeNextOfKinParams object
// with the default values initialized.
func NewPostEmployeeNextOfKinParams() *PostEmployeeNextOfKinParams {
	var ()
	return &PostEmployeeNextOfKinParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostEmployeeNextOfKinParamsWithTimeout creates a new PostEmployeeNextOfKinParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostEmployeeNextOfKinParamsWithTimeout(timeout time.Duration) *PostEmployeeNextOfKinParams {
	var ()
	return &PostEmployeeNextOfKinParams{

		timeout: timeout,
	}
}

// NewPostEmployeeNextOfKinParamsWithContext creates a new PostEmployeeNextOfKinParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostEmployeeNextOfKinParamsWithContext(ctx context.Context) *PostEmployeeNextOfKinParams {
	var ()
	return &PostEmployeeNextOfKinParams{

		Context: ctx,
	}
}

// NewPostEmployeeNextOfKinParamsWithHTTPClient creates a new PostEmployeeNextOfKinParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostEmployeeNextOfKinParamsWithHTTPClient(client *http.Client) *PostEmployeeNextOfKinParams {
	var ()
	return &PostEmployeeNextOfKinParams{
		HTTPClient: client,
	}
}

/*PostEmployeeNextOfKinParams contains all the parameters to send to the API endpoint
for the post employee next of kin operation typically these are written to a http.Request
*/
type PostEmployeeNextOfKinParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.NextOfKin

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post employee next of kin params
func (o *PostEmployeeNextOfKinParams) WithTimeout(timeout time.Duration) *PostEmployeeNextOfKinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post employee next of kin params
func (o *PostEmployeeNextOfKinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post employee next of kin params
func (o *PostEmployeeNextOfKinParams) WithContext(ctx context.Context) *PostEmployeeNextOfKinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post employee next of kin params
func (o *PostEmployeeNextOfKinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post employee next of kin params
func (o *PostEmployeeNextOfKinParams) WithHTTPClient(client *http.Client) *PostEmployeeNextOfKinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post employee next of kin params
func (o *PostEmployeeNextOfKinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post employee next of kin params
func (o *PostEmployeeNextOfKinParams) WithBody(body *models.NextOfKin) *PostEmployeeNextOfKinParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post employee next of kin params
func (o *PostEmployeeNextOfKinParams) SetBody(body *models.NextOfKin) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostEmployeeNextOfKinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
