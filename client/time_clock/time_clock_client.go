// Code generated by go-swagger; DO NOT EDIT.

package time_clock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new time clock API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for time clock API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetTimesheetTimeClock finds time clock entries corresponding with sent data
*/
func (a *Client) GetTimesheetTimeClock(params *GetTimesheetTimeClockParams, authInfo runtime.ClientAuthInfoWriter) (*GetTimesheetTimeClockOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTimesheetTimeClockParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTimesheetTimeClock",
		Method:             "GET",
		PathPattern:        "/timesheet/timeClock",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTimesheetTimeClockReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTimesheetTimeClockOK), nil

}

/*
GetTimesheetTimeClockID finds time clock entry by ID
*/
func (a *Client) GetTimesheetTimeClockID(params *GetTimesheetTimeClockIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetTimesheetTimeClockIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTimesheetTimeClockIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTimesheetTimeClockID",
		Method:             "GET",
		PathPattern:        "/timesheet/timeClock/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTimesheetTimeClockIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTimesheetTimeClockIDOK), nil

}

/*
GetTimesheetTimeClockPresent finds a user s present running time clock
*/
func (a *Client) GetTimesheetTimeClockPresent(params *GetTimesheetTimeClockPresentParams, authInfo runtime.ClientAuthInfoWriter) (*GetTimesheetTimeClockPresentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTimesheetTimeClockPresentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTimesheetTimeClockPresent",
		Method:             "GET",
		PathPattern:        "/timesheet/timeClock/present",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTimesheetTimeClockPresentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTimesheetTimeClockPresentOK), nil

}

/*
PutTimesheetTimeClockID updates time clock by ID
*/
func (a *Client) PutTimesheetTimeClockID(params *PutTimesheetTimeClockIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutTimesheetTimeClockIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutTimesheetTimeClockIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutTimesheetTimeClockID",
		Method:             "PUT",
		PathPattern:        "/timesheet/timeClock/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutTimesheetTimeClockIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutTimesheetTimeClockIDOK), nil

}

/*
PutTimesheetTimeClockIDStop stops time clock
*/
func (a *Client) PutTimesheetTimeClockIDStop(params *PutTimesheetTimeClockIDStopParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutTimesheetTimeClockIDStopParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutTimesheetTimeClockIDStop",
		Method:             "PUT",
		PathPattern:        "/timesheet/timeClock/{id}/:stop",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutTimesheetTimeClockIDStopReader{formats: a.formats},
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
PutTimesheetTimeClockStart starts time clock
*/
func (a *Client) PutTimesheetTimeClockStart(params *PutTimesheetTimeClockStartParams, authInfo runtime.ClientAuthInfoWriter) (*PutTimesheetTimeClockStartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutTimesheetTimeClockStartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutTimesheetTimeClockStart",
		Method:             "PUT",
		PathPattern:        "/timesheet/timeClock/:start",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutTimesheetTimeClockStartReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutTimesheetTimeClockStartOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}