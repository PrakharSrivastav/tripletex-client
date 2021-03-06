// Code generated by go-swagger; DO NOT EDIT.

package unit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetProductUnitReader is a Reader for the GetProductUnit structure.
type GetProductUnitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductUnitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProductUnitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProductUnitOK creates a GetProductUnitOK with default headers values
func NewGetProductUnitOK() *GetProductUnitOK {
	return &GetProductUnitOK{}
}

/*GetProductUnitOK handles this case with default header values.

successful operation
*/
type GetProductUnitOK struct {
	Payload *models.ListResponseProductUnit
}

func (o *GetProductUnitOK) Error() string {
	return fmt.Sprintf("[GET /product/unit][%d] getProductUnitOK  %+v", 200, o.Payload)
}

func (o *GetProductUnitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseProductUnit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
