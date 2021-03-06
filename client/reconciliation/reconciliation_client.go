// Code generated by go-swagger; DO NOT EDIT.

package reconciliation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new reconciliation API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reconciliation API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteBankReconciliationID bs e t a delete bank reconciliation by ID
*/
func (a *Client) DeleteBankReconciliationID(params *DeleteBankReconciliationIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBankReconciliationIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteBankReconciliationID",
		Method:             "DELETE",
		PathPattern:        "/bank/reconciliation/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBankReconciliationIDReader{formats: a.formats},
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
GetBankReconciliation bs e t a find bank reconciliation corresponding with sent data
*/
func (a *Client) GetBankReconciliation(params *GetBankReconciliationParams, authInfo runtime.ClientAuthInfoWriter) (*GetBankReconciliationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBankReconciliationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBankReconciliation",
		Method:             "GET",
		PathPattern:        "/bank/reconciliation",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBankReconciliationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBankReconciliationOK), nil

}

/*
GetBankReconciliationID bs e t a get bank reconciliation
*/
func (a *Client) GetBankReconciliationID(params *GetBankReconciliationIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetBankReconciliationIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBankReconciliationIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBankReconciliationID",
		Method:             "GET",
		PathPattern:        "/bank/reconciliation/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBankReconciliationIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBankReconciliationIDOK), nil

}

/*
GetBankReconciliationLastClosed bs e t a get last closed reconciliation by account ID
*/
func (a *Client) GetBankReconciliationLastClosed(params *GetBankReconciliationLastClosedParams, authInfo runtime.ClientAuthInfoWriter) (*GetBankReconciliationLastClosedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBankReconciliationLastClosedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBankReconciliationLastClosed",
		Method:             "GET",
		PathPattern:        "/bank/reconciliation/lastClosed",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBankReconciliationLastClosedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBankReconciliationLastClosedOK), nil

}

/*
PostBankReconciliation bs e t a post a bank reconciliation
*/
func (a *Client) PostBankReconciliation(params *PostBankReconciliationParams, authInfo runtime.ClientAuthInfoWriter) (*PostBankReconciliationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostBankReconciliationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostBankReconciliation",
		Method:             "POST",
		PathPattern:        "/bank/reconciliation",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostBankReconciliationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostBankReconciliationCreated), nil

}

/*
PutBankReconciliationFetchFromBank bs e t a create a bank reconciliation by fetching bank statement from the bank
*/
func (a *Client) PutBankReconciliationFetchFromBank(params *PutBankReconciliationFetchFromBankParams, authInfo runtime.ClientAuthInfoWriter) (*PutBankReconciliationFetchFromBankOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutBankReconciliationFetchFromBankParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutBankReconciliationFetchFromBank",
		Method:             "PUT",
		PathPattern:        "/bank/reconciliation/:fetchFromBank",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutBankReconciliationFetchFromBankReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutBankReconciliationFetchFromBankOK), nil

}

/*
PutBankReconciliationID bs e t a update a bank reconciliation
*/
func (a *Client) PutBankReconciliationID(params *PutBankReconciliationIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutBankReconciliationIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutBankReconciliationIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutBankReconciliationID",
		Method:             "PUT",
		PathPattern:        "/bank/reconciliation/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutBankReconciliationIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutBankReconciliationIDOK), nil

}

/*
PutBankReconciliationIDAdjustment bs e t a add an adjustment to reconciliation by ID
*/
func (a *Client) PutBankReconciliationIDAdjustment(params *PutBankReconciliationIDAdjustmentParams, authInfo runtime.ClientAuthInfoWriter) (*PutBankReconciliationIDAdjustmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutBankReconciliationIDAdjustmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutBankReconciliationIDAdjustment",
		Method:             "PUT",
		PathPattern:        "/bank/reconciliation/{id}/:adjustment",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutBankReconciliationIDAdjustmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutBankReconciliationIDAdjustmentOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
