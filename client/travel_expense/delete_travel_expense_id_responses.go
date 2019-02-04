// Code generated by go-swagger; DO NOT EDIT.

package travel_expense

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteTravelExpenseIDReader is a Reader for the DeleteTravelExpenseID structure.
type DeleteTravelExpenseIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTravelExpenseIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteTravelExpenseIDDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteTravelExpenseIDDefault creates a DeleteTravelExpenseIDDefault with default headers values
func NewDeleteTravelExpenseIDDefault(code int) *DeleteTravelExpenseIDDefault {
	return &DeleteTravelExpenseIDDefault{
		_statusCode: code,
	}
}

/*DeleteTravelExpenseIDDefault handles this case with default header values.

successful operation
*/
type DeleteTravelExpenseIDDefault struct {
	_statusCode int
}

// Code gets the status code for the delete travel expense ID default response
func (o *DeleteTravelExpenseIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTravelExpenseIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /travelExpense/{id}][%d] DeleteTravelExpenseID default ", o._statusCode)
}

func (o *DeleteTravelExpenseIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
