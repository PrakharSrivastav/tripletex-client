// Code generated by go-swagger; DO NOT EDIT.

package passenger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteTravelExpensePassengerIDReader is a Reader for the DeleteTravelExpensePassengerID structure.
type DeleteTravelExpensePassengerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTravelExpensePassengerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteTravelExpensePassengerIDDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteTravelExpensePassengerIDDefault creates a DeleteTravelExpensePassengerIDDefault with default headers values
func NewDeleteTravelExpensePassengerIDDefault(code int) *DeleteTravelExpensePassengerIDDefault {
	return &DeleteTravelExpensePassengerIDDefault{
		_statusCode: code,
	}
}

/*DeleteTravelExpensePassengerIDDefault handles this case with default header values.

successful operation
*/
type DeleteTravelExpensePassengerIDDefault struct {
	_statusCode int
}

// Code gets the status code for the delete travel expense passenger ID default response
func (o *DeleteTravelExpensePassengerIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTravelExpensePassengerIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /travelExpense/passenger/{id}][%d] DeleteTravelExpensePassengerID default ", o._statusCode)
}

func (o *DeleteTravelExpensePassengerIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
