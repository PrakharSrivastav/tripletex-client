// Code generated by go-swagger; DO NOT EDIT.

package transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new transaction API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for transaction API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteSalaryTransactionID bs e t a delete salary transaction by ID
*/
func (a *Client) DeleteSalaryTransactionID(params *DeleteSalaryTransactionIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSalaryTransactionIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteSalaryTransactionID",
		Method:             "DELETE",
		PathPattern:        "/salary/transaction/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSalaryTransactionIDReader{formats: a.formats},
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
GetBankStatementTransaction bs e t a find bank transaction corresponding with sent data
*/
func (a *Client) GetBankStatementTransaction(params *GetBankStatementTransactionParams, authInfo runtime.ClientAuthInfoWriter) (*GetBankStatementTransactionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBankStatementTransactionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBankStatementTransaction",
		Method:             "GET",
		PathPattern:        "/bank/statement/transaction",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBankStatementTransactionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBankStatementTransactionOK), nil

}

/*
GetBankStatementTransactionID bs e t a get bank transaction by ID
*/
func (a *Client) GetBankStatementTransactionID(params *GetBankStatementTransactionIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetBankStatementTransactionIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBankStatementTransactionIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBankStatementTransactionID",
		Method:             "GET",
		PathPattern:        "/bank/statement/transaction/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBankStatementTransactionIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBankStatementTransactionIDOK), nil

}

/*
GetBankStatementTransactionIDDetails bs e t a get additional details about transaction by ID
*/
func (a *Client) GetBankStatementTransactionIDDetails(params *GetBankStatementTransactionIDDetailsParams, authInfo runtime.ClientAuthInfoWriter) (*GetBankStatementTransactionIDDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBankStatementTransactionIDDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBankStatementTransactionIDDetails",
		Method:             "GET",
		PathPattern:        "/bank/statement/transaction/{id}/details",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBankStatementTransactionIDDetailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBankStatementTransactionIDDetailsOK), nil

}

/*
GetSalaryTransactionID bs e t a find salary transaction by ID
*/
func (a *Client) GetSalaryTransactionID(params *GetSalaryTransactionIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetSalaryTransactionIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSalaryTransactionIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSalaryTransactionID",
		Method:             "GET",
		PathPattern:        "/salary/transaction/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSalaryTransactionIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSalaryTransactionIDOK), nil

}

/*
PostSalaryTransaction bs e t a create a new salary transaction
*/
func (a *Client) PostSalaryTransaction(params *PostSalaryTransactionParams, authInfo runtime.ClientAuthInfoWriter) (*PostSalaryTransactionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSalaryTransactionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostSalaryTransaction",
		Method:             "POST",
		PathPattern:        "/salary/transaction",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostSalaryTransactionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSalaryTransactionCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
