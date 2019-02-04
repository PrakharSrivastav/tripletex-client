// Code generated by go-swagger; DO NOT EDIT.

package type_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new type operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for type operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSalaryType bs e t a find salary type corresponding with sent data
*/
func (a *Client) GetSalaryType(params *GetSalaryTypeParams, authInfo runtime.ClientAuthInfoWriter) (*GetSalaryTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSalaryTypeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSalaryType",
		Method:             "GET",
		PathPattern:        "/salary/type",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSalaryTypeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSalaryTypeOK), nil

}

/*
GetSalaryTypeID bs e t a find salary type by ID
*/
func (a *Client) GetSalaryTypeID(params *GetSalaryTypeIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetSalaryTypeIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSalaryTypeIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSalaryTypeID",
		Method:             "GET",
		PathPattern:        "/salary/type/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSalaryTypeIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSalaryTypeIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}