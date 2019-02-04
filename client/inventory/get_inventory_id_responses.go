// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetInventoryIDReader is a Reader for the GetInventoryID structure.
type GetInventoryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInventoryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInventoryIDOK creates a GetInventoryIDOK with default headers values
func NewGetInventoryIDOK() *GetInventoryIDOK {
	return &GetInventoryIDOK{}
}

/*GetInventoryIDOK handles this case with default header values.

successful operation
*/
type GetInventoryIDOK struct {
	Payload *models.ResponseWrapperInventory
}

func (o *GetInventoryIDOK) Error() string {
	return fmt.Sprintf("[GET /inventory/{id}][%d] getInventoryIdOK  %+v", 200, o.Payload)
}

func (o *GetInventoryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperInventory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
