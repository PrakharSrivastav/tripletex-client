// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSalarySettings bs e t a get salary settings of logged in company
*/
func (a *Client) GetSalarySettings(params *GetSalarySettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSalarySettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSalarySettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSalarySettings",
		Method:             "GET",
		PathPattern:        "/salary/settings",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSalarySettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSalarySettingsOK), nil

}

/*
GetTimesheetSettings bs e t a get timesheet settings of logged in company
*/
func (a *Client) GetTimesheetSettings(params *GetTimesheetSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTimesheetSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTimesheetSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTimesheetSettings",
		Method:             "GET",
		PathPattern:        "/timesheet/settings",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTimesheetSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTimesheetSettingsOK), nil

}

/*
PutSalarySettings bs e t a update settings of logged in company
*/
func (a *Client) PutSalarySettings(params *PutSalarySettingsParams, authInfo runtime.ClientAuthInfoWriter) (*PutSalarySettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutSalarySettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutSalarySettings",
		Method:             "PUT",
		PathPattern:        "/salary/settings",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutSalarySettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutSalarySettingsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
