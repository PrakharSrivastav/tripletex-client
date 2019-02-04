// Code generated by go-swagger; DO NOT EDIT.

package vat_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new vat type API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vat type API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetLedgerVatType finds vat types corresponding with sent data
*/
func (a *Client) GetLedgerVatType(params *GetLedgerVatTypeParams, authInfo runtime.ClientAuthInfoWriter) (*GetLedgerVatTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLedgerVatTypeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLedgerVatType",
		Method:             "GET",
		PathPattern:        "/ledger/vatType",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLedgerVatTypeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLedgerVatTypeOK), nil

}

/*
GetLedgerVatTypeID gets vat type by ID
*/
func (a *Client) GetLedgerVatTypeID(params *GetLedgerVatTypeIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLedgerVatTypeIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLedgerVatTypeIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLedgerVatTypeID",
		Method:             "GET",
		PathPattern:        "/ledger/vatType/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLedgerVatTypeIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLedgerVatTypeIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
