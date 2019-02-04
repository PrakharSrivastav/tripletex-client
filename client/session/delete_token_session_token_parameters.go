// Code generated by go-swagger; DO NOT EDIT.

package session

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

// NewDeleteTokenSessionTokenParams creates a new DeleteTokenSessionTokenParams object
// with the default values initialized.
func NewDeleteTokenSessionTokenParams() *DeleteTokenSessionTokenParams {
	var ()
	return &DeleteTokenSessionTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTokenSessionTokenParamsWithTimeout creates a new DeleteTokenSessionTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTokenSessionTokenParamsWithTimeout(timeout time.Duration) *DeleteTokenSessionTokenParams {
	var ()
	return &DeleteTokenSessionTokenParams{

		timeout: timeout,
	}
}

// NewDeleteTokenSessionTokenParamsWithContext creates a new DeleteTokenSessionTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTokenSessionTokenParamsWithContext(ctx context.Context) *DeleteTokenSessionTokenParams {
	var ()
	return &DeleteTokenSessionTokenParams{

		Context: ctx,
	}
}

// NewDeleteTokenSessionTokenParamsWithHTTPClient creates a new DeleteTokenSessionTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTokenSessionTokenParamsWithHTTPClient(client *http.Client) *DeleteTokenSessionTokenParams {
	var ()
	return &DeleteTokenSessionTokenParams{
		HTTPClient: client,
	}
}

/*DeleteTokenSessionTokenParams contains all the parameters to send to the API endpoint
for the delete token session token operation typically these are written to a http.Request
*/
type DeleteTokenSessionTokenParams struct {

	/*Token
	  The login token string to delete

	*/
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete token session token params
func (o *DeleteTokenSessionTokenParams) WithTimeout(timeout time.Duration) *DeleteTokenSessionTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete token session token params
func (o *DeleteTokenSessionTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete token session token params
func (o *DeleteTokenSessionTokenParams) WithContext(ctx context.Context) *DeleteTokenSessionTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete token session token params
func (o *DeleteTokenSessionTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete token session token params
func (o *DeleteTokenSessionTokenParams) WithHTTPClient(client *http.Client) *DeleteTokenSessionTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete token session token params
func (o *DeleteTokenSessionTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the delete token session token params
func (o *DeleteTokenSessionTokenParams) WithToken(token string) *DeleteTokenSessionTokenParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the delete token session token params
func (o *DeleteTokenSessionTokenParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTokenSessionTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param token
	if err := r.SetPathParam("token", o.Token); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}