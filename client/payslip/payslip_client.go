// Code generated by go-swagger; DO NOT EDIT.

package payslip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new payslip API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payslip API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSalaryPayslip bs e t a find payslips corresponding with sent data
*/
func (a *Client) GetSalaryPayslip(params *GetSalaryPayslipParams, authInfo runtime.ClientAuthInfoWriter) (*GetSalaryPayslipOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSalaryPayslipParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSalaryPayslip",
		Method:             "GET",
		PathPattern:        "/salary/payslip",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSalaryPayslipReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSalaryPayslipOK), nil

}

/*
GetSalaryPayslipID bs e t a find payslip by ID
*/
func (a *Client) GetSalaryPayslipID(params *GetSalaryPayslipIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetSalaryPayslipIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSalaryPayslipIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSalaryPayslipID",
		Method:             "GET",
		PathPattern:        "/salary/payslip/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSalaryPayslipIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSalaryPayslipIDOK), nil

}

/*
GetSalaryPayslipIDPdf bs e t a find payslip p d f document by ID
*/
func (a *Client) GetSalaryPayslipIDPdf(params *GetSalaryPayslipIDPdfParams, authInfo runtime.ClientAuthInfoWriter) (*GetSalaryPayslipIDPdfOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSalaryPayslipIDPdfParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSalaryPayslipIDPdf",
		Method:             "GET",
		PathPattern:        "/salary/payslip/{id}/pdf",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSalaryPayslipIDPdfReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSalaryPayslipIDPdfOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}