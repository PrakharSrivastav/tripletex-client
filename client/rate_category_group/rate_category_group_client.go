// Code generated by go-swagger; DO NOT EDIT.

package rate_category_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new rate category group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rate category group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetTravelExpenseRateCategoryGroup bs e t a find rate categoriy groups corresponding with sent data
*/
func (a *Client) GetTravelExpenseRateCategoryGroup(params *GetTravelExpenseRateCategoryGroupParams, authInfo runtime.ClientAuthInfoWriter) (*GetTravelExpenseRateCategoryGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTravelExpenseRateCategoryGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTravelExpenseRateCategoryGroup",
		Method:             "GET",
		PathPattern:        "/travelExpense/rateCategoryGroup",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTravelExpenseRateCategoryGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTravelExpenseRateCategoryGroupOK), nil

}

/*
GetTravelExpenseRateCategoryGroupID bs e t a get travel report rate category group by ID
*/
func (a *Client) GetTravelExpenseRateCategoryGroupID(params *GetTravelExpenseRateCategoryGroupIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetTravelExpenseRateCategoryGroupIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTravelExpenseRateCategoryGroupIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTravelExpenseRateCategoryGroupID",
		Method:             "GET",
		PathPattern:        "/travelExpense/rateCategoryGroup/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTravelExpenseRateCategoryGroupIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTravelExpenseRateCategoryGroupIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
