// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new inventory API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for inventory API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetInventory finds inventory corresponding with sent data
*/
func (a *Client) GetInventory(params *GetInventoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetInventoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInventory",
		Method:             "GET",
		PathPattern:        "/inventory",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInventoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInventoryOK), nil

}

/*
GetInventoryID gets inventory by ID
*/
func (a *Client) GetInventoryID(params *GetInventoryIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetInventoryIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInventoryID",
		Method:             "GET",
		PathPattern:        "/inventory/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInventoryIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInventoryIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
