// Code generated by go-swagger; DO NOT EDIT.

package time_clock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetTimesheetTimeClockPresentReader is a Reader for the GetTimesheetTimeClockPresent structure.
type GetTimesheetTimeClockPresentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTimesheetTimeClockPresentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTimesheetTimeClockPresentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTimesheetTimeClockPresentOK creates a GetTimesheetTimeClockPresentOK with default headers values
func NewGetTimesheetTimeClockPresentOK() *GetTimesheetTimeClockPresentOK {
	return &GetTimesheetTimeClockPresentOK{}
}

/*GetTimesheetTimeClockPresentOK handles this case with default header values.

successful operation
*/
type GetTimesheetTimeClockPresentOK struct {
	Payload *models.ResponseWrapperTimeClock
}

func (o *GetTimesheetTimeClockPresentOK) Error() string {
	return fmt.Sprintf("[GET /timesheet/timeClock/present][%d] getTimesheetTimeClockPresentOK  %+v", 200, o.Payload)
}

func (o *GetTimesheetTimeClockPresentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperTimeClock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
