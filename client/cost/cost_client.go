// Code generated by go-swagger; DO NOT EDIT.

package cost

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cost API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cost API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteTravelExpenseCostID bs e t a delete cost
*/
func (a *Client) DeleteTravelExpenseCostID(params *DeleteTravelExpenseCostIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTravelExpenseCostIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteTravelExpenseCostID",
		Method:             "DELETE",
		PathPattern:        "/travelExpense/cost/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTravelExpenseCostIDReader{formats: a.formats},
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
GetTravelExpenseCost bs e t a find costs corresponding with sent data
*/
func (a *Client) GetTravelExpenseCost(params *GetTravelExpenseCostParams, authInfo runtime.ClientAuthInfoWriter) (*GetTravelExpenseCostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTravelExpenseCostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTravelExpenseCost",
		Method:             "GET",
		PathPattern:        "/travelExpense/cost",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTravelExpenseCostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTravelExpenseCostOK), nil

}

/*
GetTravelExpenseCostID bs e t a get cost by ID
*/
func (a *Client) GetTravelExpenseCostID(params *GetTravelExpenseCostIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetTravelExpenseCostIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTravelExpenseCostIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTravelExpenseCostID",
		Method:             "GET",
		PathPattern:        "/travelExpense/cost/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTravelExpenseCostIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTravelExpenseCostIDOK), nil

}

/*
PostTravelExpenseCost bs e t a create cost
*/
func (a *Client) PostTravelExpenseCost(params *PostTravelExpenseCostParams, authInfo runtime.ClientAuthInfoWriter) (*PostTravelExpenseCostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTravelExpenseCostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTravelExpenseCost",
		Method:             "POST",
		PathPattern:        "/travelExpense/cost",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostTravelExpenseCostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTravelExpenseCostCreated), nil

}

/*
PutTravelExpenseCostID bs e t a update cost
*/
func (a *Client) PutTravelExpenseCostID(params *PutTravelExpenseCostIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutTravelExpenseCostIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutTravelExpenseCostIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutTravelExpenseCostID",
		Method:             "PUT",
		PathPattern:        "/travelExpense/cost/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutTravelExpenseCostIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutTravelExpenseCostIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
