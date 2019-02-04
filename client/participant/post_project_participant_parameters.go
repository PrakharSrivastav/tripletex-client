// Code generated by go-swagger; DO NOT EDIT.

package participant

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

// NewPostProjectParticipantParams creates a new PostProjectParticipantParams object
// with the default values initialized.
func NewPostProjectParticipantParams() *PostProjectParticipantParams {
	var ()
	return &PostProjectParticipantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostProjectParticipantParamsWithTimeout creates a new PostProjectParticipantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostProjectParticipantParamsWithTimeout(timeout time.Duration) *PostProjectParticipantParams {
	var ()
	return &PostProjectParticipantParams{

		timeout: timeout,
	}
}

// NewPostProjectParticipantParamsWithContext creates a new PostProjectParticipantParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostProjectParticipantParamsWithContext(ctx context.Context) *PostProjectParticipantParams {
	var ()
	return &PostProjectParticipantParams{

		Context: ctx,
	}
}

// NewPostProjectParticipantParamsWithHTTPClient creates a new PostProjectParticipantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostProjectParticipantParamsWithHTTPClient(client *http.Client) *PostProjectParticipantParams {
	var ()
	return &PostProjectParticipantParams{
		HTTPClient: client,
	}
}

/*PostProjectParticipantParams contains all the parameters to send to the API endpoint
for the post project participant operation typically these are written to a http.Request
*/
type PostProjectParticipantParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.ProjectParticipant

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post project participant params
func (o *PostProjectParticipantParams) WithTimeout(timeout time.Duration) *PostProjectParticipantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post project participant params
func (o *PostProjectParticipantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post project participant params
func (o *PostProjectParticipantParams) WithContext(ctx context.Context) *PostProjectParticipantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post project participant params
func (o *PostProjectParticipantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post project participant params
func (o *PostProjectParticipantParams) WithHTTPClient(client *http.Client) *PostProjectParticipantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post project participant params
func (o *PostProjectParticipantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post project participant params
func (o *PostProjectParticipantParams) WithBody(body *models.ProjectParticipant) *PostProjectParticipantParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post project participant params
func (o *PostProjectParticipantParams) SetBody(body *models.ProjectParticipant) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostProjectParticipantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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