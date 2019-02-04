// Code generated by go-swagger; DO NOT EDIT.

package entitlement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new entitlement API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for entitlement API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetEmployeeEntitlement finds all entitlements for user
*/
func (a *Client) GetEmployeeEntitlement(params *GetEmployeeEntitlementParams, authInfo runtime.ClientAuthInfoWriter) (*GetEmployeeEntitlementOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmployeeEntitlementParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEmployeeEntitlement",
		Method:             "GET",
		PathPattern:        "/employee/entitlement",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEmployeeEntitlementReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEmployeeEntitlementOK), nil

}

/*
GetEmployeeEntitlementClient bs e t a find all entitlements at client for user
*/
func (a *Client) GetEmployeeEntitlementClient(params *GetEmployeeEntitlementClientParams, authInfo runtime.ClientAuthInfoWriter) (*GetEmployeeEntitlementClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmployeeEntitlementClientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEmployeeEntitlementClient",
		Method:             "GET",
		PathPattern:        "/employee/entitlement/client",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEmployeeEntitlementClientReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEmployeeEntitlementClientOK), nil

}

/*
GetEmployeeEntitlementID gets entitlement by ID
*/
func (a *Client) GetEmployeeEntitlementID(params *GetEmployeeEntitlementIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetEmployeeEntitlementIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmployeeEntitlementIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEmployeeEntitlementID",
		Method:             "GET",
		PathPattern:        "/employee/entitlement/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEmployeeEntitlementIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEmployeeEntitlementIDOK), nil

}

/*
PutEmployeeEntitlementGrantClientEntitlementsByTemplate bs e t a update employee entitlements in client account
*/
func (a *Client) PutEmployeeEntitlementGrantClientEntitlementsByTemplate(params *PutEmployeeEntitlementGrantClientEntitlementsByTemplateParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEmployeeEntitlementGrantClientEntitlementsByTemplateParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutEmployeeEntitlementGrantClientEntitlementsByTemplate",
		Method:             "PUT",
		PathPattern:        "/employee/entitlement/:grantClientEntitlementsByTemplate",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutEmployeeEntitlementGrantClientEntitlementsByTemplateReader{formats: a.formats},
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
PutEmployeeEntitlementGrantEntitlementsByTemplate bs e t a update employee entitlements

The user will only receive the entitlements which are possible with the registered user type
*/
func (a *Client) PutEmployeeEntitlementGrantEntitlementsByTemplate(params *PutEmployeeEntitlementGrantEntitlementsByTemplateParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEmployeeEntitlementGrantEntitlementsByTemplateParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutEmployeeEntitlementGrantEntitlementsByTemplate",
		Method:             "PUT",
		PathPattern:        "/employee/entitlement/:grantEntitlementsByTemplate",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutEmployeeEntitlementGrantEntitlementsByTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
