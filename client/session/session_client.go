// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new session API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for session API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteTokenSessionToken deletes session token
*/
func (a *Client) DeleteTokenSessionToken(params *DeleteTokenSessionTokenParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTokenSessionTokenParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteTokenSessionToken",
		Method:             "DELETE",
		PathPattern:        "/token/session/{token}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTokenSessionTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetTokenSessionWhoAmI finds information about the current user
*/
func (a *Client) GetTokenSessionWhoAmI(params *GetTokenSessionWhoAmIParams, authInfo runtime.ClientAuthInfoWriter) (*GetTokenSessionWhoAmIOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTokenSessionWhoAmIParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTokenSessionWhoAmI",
		Method:             "GET",
		PathPattern:        "/token/session/whoAmI",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTokenSessionWhoAmIReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTokenSessionWhoAmIOK), nil

}

/*
PutTokenSessionCreate creates session token
*/
func (a *Client) PutTokenSessionCreate(params *PutTokenSessionCreateParams) (*PutTokenSessionCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutTokenSessionCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutTokenSessionCreate",
		Method:             "PUT",
		PathPattern:        "/token/session/:create",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutTokenSessionCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutTokenSessionCreateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
