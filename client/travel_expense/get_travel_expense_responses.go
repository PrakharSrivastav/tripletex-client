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

// GetTravelExpenseReader is a Reader for the GetTravelExpense structure.
type GetTravelExpenseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTravelExpenseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTravelExpenseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTravelExpenseOK creates a GetTravelExpenseOK with default headers values
func NewGetTravelExpenseOK() *GetTravelExpenseOK {
	return &GetTravelExpenseOK{}
}

/*GetTravelExpenseOK handles this case with default header values.

successful operation
*/
type GetTravelExpenseOK struct {
	Payload *models.ListResponseTravelExpense
}

func (o *GetTravelExpenseOK) Error() string {
	return fmt.Sprintf("[GET /travelExpense][%d] getTravelExpenseOK  %+v", 200, o.Payload)
}

func (o *GetTravelExpenseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseTravelExpense)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
