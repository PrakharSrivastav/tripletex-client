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

// GetTravelExpensePerDiemCompensationReader is a Reader for the GetTravelExpensePerDiemCompensation structure.
type GetTravelExpensePerDiemCompensationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTravelExpensePerDiemCompensationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTravelExpensePerDiemCompensationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTravelExpensePerDiemCompensationOK creates a GetTravelExpensePerDiemCompensationOK with default headers values
func NewGetTravelExpensePerDiemCompensationOK() *GetTravelExpensePerDiemCompensationOK {
	return &GetTravelExpensePerDiemCompensationOK{}
}

/*GetTravelExpensePerDiemCompensationOK handles this case with default header values.

successful operation
*/
type GetTravelExpensePerDiemCompensationOK struct {
	Payload *models.ListResponsePerDiemCompensation
}

func (o *GetTravelExpensePerDiemCompensationOK) Error() string {
	return fmt.Sprintf("[GET /travelExpense/perDiemCompensation][%d] getTravelExpensePerDiemCompensationOK  %+v", 200, o.Payload)
}

func (o *GetTravelExpensePerDiemCompensationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponsePerDiemCompensation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}