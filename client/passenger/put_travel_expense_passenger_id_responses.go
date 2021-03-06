// Code generated by go-swagger; DO NOT EDIT.

package passenger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PutTravelExpensePassengerIDReader is a Reader for the PutTravelExpensePassengerID structure.
type PutTravelExpensePassengerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTravelExpensePassengerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutTravelExpensePassengerIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutTravelExpensePassengerIDOK creates a PutTravelExpensePassengerIDOK with default headers values
func NewPutTravelExpensePassengerIDOK() *PutTravelExpensePassengerIDOK {
	return &PutTravelExpensePassengerIDOK{}
}

/*PutTravelExpensePassengerIDOK handles this case with default header values.

successful operation
*/
type PutTravelExpensePassengerIDOK struct {
	Payload *models.ResponseWrapperPassenger
}

func (o *PutTravelExpensePassengerIDOK) Error() string {
	return fmt.Sprintf("[PUT /travelExpense/passenger/{id}][%d] putTravelExpensePassengerIdOK  %+v", 200, o.Payload)
}

func (o *PutTravelExpensePassengerIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperPassenger)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
