// Code generated by go-swagger; DO NOT EDIT.

package supplier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new supplier API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for supplier API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSupplier finds suppliers corresponding with sent data
*/
func (a *Client) GetSupplier(params *GetSupplierParams, authInfo runtime.ClientAuthInfoWriter) (*GetSupplierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSupplierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSupplier",
		Method:             "GET",
		PathPattern:        "/supplier",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSupplierReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSupplierOK), nil

}

/*
GetSupplierID gets supplier by ID
*/
func (a *Client) GetSupplierID(params *GetSupplierIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetSupplierIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSupplierIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSupplierID",
		Method:             "GET",
		PathPattern:        "/supplier/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSupplierIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSupplierIDOK), nil

}

/*
PostSupplier creates supplier related supplier addresses may also be created
*/
func (a *Client) PostSupplier(params *PostSupplierParams, authInfo runtime.ClientAuthInfoWriter) (*PostSupplierCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSupplierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostSupplier",
		Method:             "POST",
		PathPattern:        "/supplier",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostSupplierReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSupplierCreated), nil

}

/*
PostSupplierList bs e t a create multiple suppliers related supplier addresses may also be created
*/
func (a *Client) PostSupplierList(params *PostSupplierListParams, authInfo runtime.ClientAuthInfoWriter) (*PostSupplierListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSupplierListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostSupplierList",
		Method:             "POST",
		PathPattern:        "/supplier/list",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostSupplierListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSupplierListCreated), nil

}

/*
PutSupplierID updates supplier
*/
func (a *Client) PutSupplierID(params *PutSupplierIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutSupplierIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutSupplierIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutSupplierID",
		Method:             "PUT",
		PathPattern:        "/supplier/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutSupplierIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutSupplierIDOK), nil

}

/*
PutSupplierList bs e t a update multiple suppliers addresses can also be updated
*/
func (a *Client) PutSupplierList(params *PutSupplierListParams, authInfo runtime.ClientAuthInfoWriter) (*PutSupplierListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutSupplierListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutSupplierList",
		Method:             "PUT",
		PathPattern:        "/supplier/list",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutSupplierListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutSupplierListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
