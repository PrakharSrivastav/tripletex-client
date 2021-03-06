// Code generated by go-swagger; DO NOT EDIT.

package cost_category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cost category API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cost category API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetTravelExpenseCostCategory bs e t a find cost category corresponding with sent data
*/
func (a *Client) GetTravelExpenseCostCategory(params *GetTravelExpenseCostCategoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetTravelExpenseCostCategoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTravelExpenseCostCategoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTravelExpenseCostCategory",
		Method:             "GET",
		PathPattern:        "/travelExpense/costCategory",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTravelExpenseCostCategoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTravelExpenseCostCategoryOK), nil

}

/*
GetTravelExpenseCostCategoryID bs e t a get cost category by ID
*/
func (a *Client) GetTravelExpenseCostCategoryID(params *GetTravelExpenseCostCategoryIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetTravelExpenseCostCategoryIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTravelExpenseCostCategoryIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTravelExpenseCostCategoryID",
		Method:             "GET",
		PathPattern:        "/travelExpense/costCategory/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTravelExpenseCostCategoryIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTravelExpenseCostCategoryIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
