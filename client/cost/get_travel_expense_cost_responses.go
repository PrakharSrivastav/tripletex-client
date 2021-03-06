// Code generated by go-swagger; DO NOT EDIT.

package cost

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetTravelExpenseCostReader is a Reader for the GetTravelExpenseCost structure.
type GetTravelExpenseCostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTravelExpenseCostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTravelExpenseCostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTravelExpenseCostOK creates a GetTravelExpenseCostOK with default headers values
func NewGetTravelExpenseCostOK() *GetTravelExpenseCostOK {
	return &GetTravelExpenseCostOK{}
}

/*GetTravelExpenseCostOK handles this case with default header values.

successful operation
*/
type GetTravelExpenseCostOK struct {
	Payload *models.ListResponseCost
}

func (o *GetTravelExpenseCostOK) Error() string {
	return fmt.Sprintf("[GET /travelExpense/cost][%d] getTravelExpenseCostOK  %+v", 200, o.Payload)
}

func (o *GetTravelExpenseCostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseCost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
