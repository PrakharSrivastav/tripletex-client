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

// PutTravelExpenseCopyReader is a Reader for the PutTravelExpenseCopy structure.
type PutTravelExpenseCopyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTravelExpenseCopyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutTravelExpenseCopyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutTravelExpenseCopyOK creates a PutTravelExpenseCopyOK with default headers values
func NewPutTravelExpenseCopyOK() *PutTravelExpenseCopyOK {
	return &PutTravelExpenseCopyOK{}
}

/*PutTravelExpenseCopyOK handles this case with default header values.

successful operation
*/
type PutTravelExpenseCopyOK struct {
	Payload *models.ResponseWrapperTravelExpense
}

func (o *PutTravelExpenseCopyOK) Error() string {
	return fmt.Sprintf("[PUT /travelExpense/:copy][%d] putTravelExpenseCopyOK  %+v", 200, o.Payload)
}

func (o *PutTravelExpenseCopyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperTravelExpense)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
