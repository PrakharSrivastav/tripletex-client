// Code generated by go-swagger; DO NOT EDIT.

package time_clock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutTimesheetTimeClockIDStopReader is a Reader for the PutTimesheetTimeClockIDStop structure.
type PutTimesheetTimeClockIDStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTimesheetTimeClockIDStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPutTimesheetTimeClockIDStopDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPutTimesheetTimeClockIDStopDefault creates a PutTimesheetTimeClockIDStopDefault with default headers values
func NewPutTimesheetTimeClockIDStopDefault(code int) *PutTimesheetTimeClockIDStopDefault {
	return &PutTimesheetTimeClockIDStopDefault{
		_statusCode: code,
	}
}

/*PutTimesheetTimeClockIDStopDefault handles this case with default header values.

successful operation
*/
type PutTimesheetTimeClockIDStopDefault struct {
	_statusCode int
}

// Code gets the status code for the put timesheet time clock ID stop default response
func (o *PutTimesheetTimeClockIDStopDefault) Code() int {
	return o._statusCode
}

func (o *PutTimesheetTimeClockIDStopDefault) Error() string {
	return fmt.Sprintf("[PUT /timesheet/timeClock/{id}/:stop][%d] PutTimesheetTimeClockIDStop default ", o._statusCode)
}

func (o *PutTimesheetTimeClockIDStopDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
