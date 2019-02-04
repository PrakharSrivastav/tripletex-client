// Code generated by go-swagger; DO NOT EDIT.

package rate_category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetTravelExpenseRateCategoryReader is a Reader for the GetTravelExpenseRateCategory structure.
type GetTravelExpenseRateCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTravelExpenseRateCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTravelExpenseRateCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTravelExpenseRateCategoryOK creates a GetTravelExpenseRateCategoryOK with default headers values
func NewGetTravelExpenseRateCategoryOK() *GetTravelExpenseRateCategoryOK {
	return &GetTravelExpenseRateCategoryOK{}
}

/*GetTravelExpenseRateCategoryOK handles this case with default header values.

successful operation
*/
type GetTravelExpenseRateCategoryOK struct {
	Payload *models.ListResponseTravelExpenseRateCategory
}

func (o *GetTravelExpenseRateCategoryOK) Error() string {
	return fmt.Sprintf("[GET /travelExpense/rateCategory][%d] getTravelExpenseRateCategoryOK  %+v", 200, o.Payload)
}

func (o *GetTravelExpenseRateCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseTravelExpenseRateCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}