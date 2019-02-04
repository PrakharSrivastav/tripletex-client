// Code generated by go-swagger; DO NOT EDIT.

package working_hours_scheme

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new working hours scheme API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for working hours scheme API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetEmployeeEmploymentWorkingHoursScheme bs e t a find working hours scheme ID
*/
func (a *Client) GetEmployeeEmploymentWorkingHoursScheme(params *GetEmployeeEmploymentWorkingHoursSchemeParams, authInfo runtime.ClientAuthInfoWriter) (*GetEmployeeEmploymentWorkingHoursSchemeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmployeeEmploymentWorkingHoursSchemeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEmployeeEmploymentWorkingHoursScheme",
		Method:             "GET",
		PathPattern:        "/employee/employment/workingHoursScheme",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEmployeeEmploymentWorkingHoursSchemeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEmployeeEmploymentWorkingHoursSchemeOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}