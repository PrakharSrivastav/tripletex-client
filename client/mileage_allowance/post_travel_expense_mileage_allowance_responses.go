// Code generated by go-swagger; DO NOT EDIT.

package mileage_allowance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PostTravelExpenseMileageAllowanceReader is a Reader for the PostTravelExpenseMileageAllowance structure.
type PostTravelExpenseMileageAllowanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTravelExpenseMileageAllowanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostTravelExpenseMileageAllowanceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTravelExpenseMileageAllowanceCreated creates a PostTravelExpenseMileageAllowanceCreated with default headers values
func NewPostTravelExpenseMileageAllowanceCreated() *PostTravelExpenseMileageAllowanceCreated {
	return &PostTravelExpenseMileageAllowanceCreated{}
}

/*PostTravelExpenseMileageAllowanceCreated handles this case with default header values.

successfully created
*/
type PostTravelExpenseMileageAllowanceCreated struct {
	Payload *models.ResponseWrapperMileageAllowance
}

func (o *PostTravelExpenseMileageAllowanceCreated) Error() string {
	return fmt.Sprintf("[POST /travelExpense/mileageAllowance][%d] postTravelExpenseMileageAllowanceCreated  %+v", 201, o.Payload)
}

func (o *PostTravelExpenseMileageAllowanceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperMileageAllowance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
