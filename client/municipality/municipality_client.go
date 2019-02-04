// Code generated by go-swagger; DO NOT EDIT.

package municipality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new municipality API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for municipality API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetMunicipality bs e t a get municipalities
*/
func (a *Client) GetMunicipality(params *GetMunicipalityParams, authInfo runtime.ClientAuthInfoWriter) (*GetMunicipalityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMunicipalityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMunicipality",
		Method:             "GET",
		PathPattern:        "/municipality",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMunicipalityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMunicipalityOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}