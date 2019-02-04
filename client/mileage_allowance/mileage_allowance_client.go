// Code generated by go-swagger; DO NOT EDIT.

package mileage_allowance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new mileage allowance API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mileage allowance API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteTravelExpenseMileageAllowanceID bs e t a delete mileage allowance
*/
func (a *Client) DeleteTravelExpenseMileageAllowanceID(params *DeleteTravelExpenseMileageAllowanceIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTravelExpenseMileageAllowanceIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteTravelExpenseMileageAllowanceID",
		Method:             "DELETE",
		PathPattern:        "/travelExpense/mileageAllowance/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTravelExpenseMileageAllowanceIDReader{formats: a.formats},
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
GetTravelExpenseMileageAllowance bs e t a find mileage allowances corresponding with sent data
*/
func (a *Client) GetTravelExpenseMileageAllowance(params *GetTravelExpenseMileageAllowanceParams, authInfo runtime.ClientAuthInfoWriter) (*GetTravelExpenseMileageAllowanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTravelExpenseMileageAllowanceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTravelExpenseMileageAllowance",
		Method:             "GET",
		PathPattern:        "/travelExpense/mileageAllowance",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTravelExpenseMileageAllowanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTravelExpenseMileageAllowanceOK), nil

}

/*
GetTravelExpenseMileageAllowanceID bs e t a get mileage allowance by ID
*/
func (a *Client) GetTravelExpenseMileageAllowanceID(params *GetTravelExpenseMileageAllowanceIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetTravelExpenseMileageAllowanceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTravelExpenseMileageAllowanceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTravelExpenseMileageAllowanceID",
		Method:             "GET",
		PathPattern:        "/travelExpense/mileageAllowance/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTravelExpenseMileageAllowanceIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTravelExpenseMileageAllowanceIDOK), nil

}

/*
PostTravelExpenseMileageAllowance bs e t a create mileage allowance
*/
func (a *Client) PostTravelExpenseMileageAllowance(params *PostTravelExpenseMileageAllowanceParams, authInfo runtime.ClientAuthInfoWriter) (*PostTravelExpenseMileageAllowanceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTravelExpenseMileageAllowanceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTravelExpenseMileageAllowance",
		Method:             "POST",
		PathPattern:        "/travelExpense/mileageAllowance",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostTravelExpenseMileageAllowanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTravelExpenseMileageAllowanceCreated), nil

}

/*
PutTravelExpenseMileageAllowanceID bs e t a update mileage allowance
*/
func (a *Client) PutTravelExpenseMileageAllowanceID(params *PutTravelExpenseMileageAllowanceIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutTravelExpenseMileageAllowanceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutTravelExpenseMileageAllowanceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutTravelExpenseMileageAllowanceID",
		Method:             "PUT",
		PathPattern:        "/travelExpense/mileageAllowance/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutTravelExpenseMileageAllowanceIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutTravelExpenseMileageAllowanceIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}