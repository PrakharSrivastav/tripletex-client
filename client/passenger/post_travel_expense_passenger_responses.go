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

// PostTravelExpensePassengerReader is a Reader for the PostTravelExpensePassenger structure.
type PostTravelExpensePassengerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTravelExpensePassengerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostTravelExpensePassengerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTravelExpensePassengerCreated creates a PostTravelExpensePassengerCreated with default headers values
func NewPostTravelExpensePassengerCreated() *PostTravelExpensePassengerCreated {
	return &PostTravelExpensePassengerCreated{}
}

/*PostTravelExpensePassengerCreated handles this case with default header values.

successfully created
*/
type PostTravelExpensePassengerCreated struct {
	Payload *models.ResponseWrapperPassenger
}

func (o *PostTravelExpensePassengerCreated) Error() string {
	return fmt.Sprintf("[POST /travelExpense/passenger][%d] postTravelExpensePassengerCreated  %+v", 201, o.Payload)
}

func (o *PostTravelExpensePassengerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperPassenger)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
