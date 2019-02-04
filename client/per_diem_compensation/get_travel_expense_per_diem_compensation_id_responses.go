// Code generated by go-swagger; DO NOT EDIT.

package per_diem_compensation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetTravelExpensePerDiemCompensationIDReader is a Reader for the GetTravelExpensePerDiemCompensationID structure.
type GetTravelExpensePerDiemCompensationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTravelExpensePerDiemCompensationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTravelExpensePerDiemCompensationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTravelExpensePerDiemCompensationIDOK creates a GetTravelExpensePerDiemCompensationIDOK with default headers values
func NewGetTravelExpensePerDiemCompensationIDOK() *GetTravelExpensePerDiemCompensationIDOK {
	return &GetTravelExpensePerDiemCompensationIDOK{}
}

/*GetTravelExpensePerDiemCompensationIDOK handles this case with default header values.

successful operation
*/
type GetTravelExpensePerDiemCompensationIDOK struct {
	Payload *models.ResponseWrapperPerDiemCompensation
}

func (o *GetTravelExpensePerDiemCompensationIDOK) Error() string {
	return fmt.Sprintf("[GET /travelExpense/perDiemCompensation/{id}][%d] getTravelExpensePerDiemCompensationIdOK  %+v", 200, o.Payload)
}

func (o *GetTravelExpensePerDiemCompensationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperPerDiemCompensation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
