// Code generated by go-swagger; DO NOT EDIT.

package leave_of_absence

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

// NewPostEmployeeEmploymentLeaveOfAbsenceListParams creates a new PostEmployeeEmploymentLeaveOfAbsenceListParams object
// with the default values initialized.
func NewPostEmployeeEmploymentLeaveOfAbsenceListParams() *PostEmployeeEmploymentLeaveOfAbsenceListParams {
	var ()
	return &PostEmployeeEmploymentLeaveOfAbsenceListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostEmployeeEmploymentLeaveOfAbsenceListParamsWithTimeout creates a new PostEmployeeEmploymentLeaveOfAbsenceListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostEmployeeEmploymentLeaveOfAbsenceListParamsWithTimeout(timeout time.Duration) *PostEmployeeEmploymentLeaveOfAbsenceListParams {
	var ()
	return &PostEmployeeEmploymentLeaveOfAbsenceListParams{

		timeout: timeout,
	}
}

// NewPostEmployeeEmploymentLeaveOfAbsenceListParamsWithContext creates a new PostEmployeeEmploymentLeaveOfAbsenceListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostEmployeeEmploymentLeaveOfAbsenceListParamsWithContext(ctx context.Context) *PostEmployeeEmploymentLeaveOfAbsenceListParams {
	var ()
	return &PostEmployeeEmploymentLeaveOfAbsenceListParams{

		Context: ctx,
	}
}

// NewPostEmployeeEmploymentLeaveOfAbsenceListParamsWithHTTPClient creates a new PostEmployeeEmploymentLeaveOfAbsenceListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostEmployeeEmploymentLeaveOfAbsenceListParamsWithHTTPClient(client *http.Client) *PostEmployeeEmploymentLeaveOfAbsenceListParams {
	var ()
	return &PostEmployeeEmploymentLeaveOfAbsenceListParams{
		HTTPClient: client,
	}
}

/*PostEmployeeEmploymentLeaveOfAbsenceListParams contains all the parameters to send to the API endpoint
for the post employee employment leave of absence list operation typically these are written to a http.Request
*/
type PostEmployeeEmploymentLeaveOfAbsenceListParams struct {

	/*Body
	  JSON representing a list of new object to be created. Should not have ID and version set.

	*/
	Body []*models.LeaveOfAbsence

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post employee employment leave of absence list params
func (o *PostEmployeeEmploymentLeaveOfAbsenceListParams) WithTimeout(timeout time.Duration) *PostEmployeeEmploymentLeaveOfAbsenceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post employee employment leave of absence list params
func (o *PostEmployeeEmploymentLeaveOfAbsenceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post employee employment leave of absence list params
func (o *PostEmployeeEmploymentLeaveOfAbsenceListParams) WithContext(ctx context.Context) *PostEmployeeEmploymentLeaveOfAbsenceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post employee employment leave of absence list params
func (o *PostEmployeeEmploymentLeaveOfAbsenceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post employee employment leave of absence list params
func (o *PostEmployeeEmploymentLeaveOfAbsenceListParams) WithHTTPClient(client *http.Client) *PostEmployeeEmploymentLeaveOfAbsenceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post employee employment leave of absence list params
func (o *PostEmployeeEmploymentLeaveOfAbsenceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post employee employment leave of absence list params
func (o *PostEmployeeEmploymentLeaveOfAbsenceListParams) WithBody(body []*models.LeaveOfAbsence) *PostEmployeeEmploymentLeaveOfAbsenceListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post employee employment leave of absence list params
func (o *PostEmployeeEmploymentLeaveOfAbsenceListParams) SetBody(body []*models.LeaveOfAbsence) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostEmployeeEmploymentLeaveOfAbsenceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
