// Code generated by go-swagger; DO NOT EDIT.

package reminder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new reminder API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reminder API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetReminder finds reminders corresponding with sent data
*/
func (a *Client) GetReminder(params *GetReminderParams, authInfo runtime.ClientAuthInfoWriter) (*GetReminderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReminderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetReminder",
		Method:             "GET",
		PathPattern:        "/reminder",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReminderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReminderOK), nil

}

/*
GetReminderID gets reminder by ID
*/
func (a *Client) GetReminderID(params *GetReminderIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetReminderIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReminderIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetReminderID",
		Method:             "GET",
		PathPattern:        "/reminder/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReminderIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReminderIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
