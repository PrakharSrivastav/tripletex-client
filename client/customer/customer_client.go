// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new customer API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCustomer finds customers corresponding with sent data
*/
func (a *Client) GetCustomer(params *GetCustomerParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCustomer",
		Method:             "GET",
		PathPattern:        "/customer",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCustomerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCustomerOK), nil

}

/*
GetCustomerID gets customer by ID
*/
func (a *Client) GetCustomerID(params *GetCustomerIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomerIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCustomerID",
		Method:             "GET",
		PathPattern:        "/customer/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCustomerIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCustomerIDOK), nil

}

/*
PostCustomer creates customer related customer addresses may also be created
*/
func (a *Client) PostCustomer(params *PostCustomerParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomerCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCustomer",
		Method:             "POST",
		PathPattern:        "/customer",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostCustomerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCustomerCreated), nil

}

/*
PostCustomerList bs e t a create multiple customers related supplier addresses may also be created
*/
func (a *Client) PostCustomerList(params *PostCustomerListParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomerListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomerListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCustomerList",
		Method:             "POST",
		PathPattern:        "/customer/list",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostCustomerListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCustomerListCreated), nil

}

/*
PutCustomerID updates customer
*/
func (a *Client) PutCustomerID(params *PutCustomerIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutCustomerIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCustomerIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCustomerID",
		Method:             "PUT",
		PathPattern:        "/customer/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutCustomerIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutCustomerIDOK), nil

}

/*
PutCustomerList bs e t a update multiple customers addresses can also be updated
*/
func (a *Client) PutCustomerList(params *PutCustomerListParams, authInfo runtime.ClientAuthInfoWriter) (*PutCustomerListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCustomerListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCustomerList",
		Method:             "PUT",
		PathPattern:        "/customer/list",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutCustomerListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutCustomerListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}