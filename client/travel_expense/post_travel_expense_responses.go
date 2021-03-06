// Code generated by go-swagger; DO NOT EDIT.

package travel_expense

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PostTravelExpenseReader is a Reader for the PostTravelExpense structure.
type PostTravelExpenseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTravelExpenseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostTravelExpenseCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTravelExpenseCreated creates a PostTravelExpenseCreated with default headers values
func NewPostTravelExpenseCreated() *PostTravelExpenseCreated {
	return &PostTravelExpenseCreated{}
}

/*PostTravelExpenseCreated handles this case with default header values.

successfully created
*/
type PostTravelExpenseCreated struct {
	Payload *models.ResponseWrapperTravelExpense
}

func (o *PostTravelExpenseCreated) Error() string {
	return fmt.Sprintf("[POST /travelExpense][%d] postTravelExpenseCreated  %+v", 201, o.Payload)
}

func (o *PostTravelExpenseCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperTravelExpense)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
