// Code generated by go-swagger; DO NOT EDIT.

package employee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new employee API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for employee API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetEmployee finds employees corresponding with sent data
*/
func (a *Client) GetEmployee(params *GetEmployeeParams, authInfo runtime.ClientAuthInfoWriter) (*GetEmployeeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmployeeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEmployee",
		Method:             "GET",
		PathPattern:        "/employee",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEmployeeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEmployeeOK), nil

}

/*
GetEmployeeID gets employee by ID
*/
func (a *Client) GetEmployeeID(params *GetEmployeeIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetEmployeeIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmployeeIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEmployeeID",
		Method:             "GET",
		PathPattern:        "/employee/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEmployeeIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEmployeeIDOK), nil

}

/*
PostEmployee bs e t a create one employee
*/
func (a *Client) PostEmployee(params *PostEmployeeParams, authInfo runtime.ClientAuthInfoWriter) (*PostEmployeeCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEmployeeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostEmployee",
		Method:             "POST",
		PathPattern:        "/employee",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostEmployeeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostEmployeeCreated), nil

}

/*
PostEmployeeList bs e t a create several employees
*/
func (a *Client) PostEmployeeList(params *PostEmployeeListParams, authInfo runtime.ClientAuthInfoWriter) (*PostEmployeeListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEmployeeListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostEmployeeList",
		Method:             "POST",
		PathPattern:        "/employee/list",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostEmployeeListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostEmployeeListCreated), nil

}

/*
PutEmployeeID updates employee
*/
func (a *Client) PutEmployeeID(params *PutEmployeeIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutEmployeeIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEmployeeIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutEmployeeID",
		Method:             "PUT",
		PathPattern:        "/employee/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutEmployeeIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutEmployeeIDOK), nil

}

/*
PutTokenEmployeeCreate creates an employee token only selected consumers are allowed
*/
func (a *Client) PutTokenEmployeeCreate(params *PutTokenEmployeeCreateParams, authInfo runtime.ClientAuthInfoWriter) (*PutTokenEmployeeCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutTokenEmployeeCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutTokenEmployeeCreate",
		Method:             "PUT",
		PathPattern:        "/token/employee/:create",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutTokenEmployeeCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutTokenEmployeeCreateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
