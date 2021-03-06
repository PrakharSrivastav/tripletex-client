// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetProductIDReader is a Reader for the GetProductID structure.
type GetProductIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProductIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProductIDOK creates a GetProductIDOK with default headers values
func NewGetProductIDOK() *GetProductIDOK {
	return &GetProductIDOK{}
}

/*GetProductIDOK handles this case with default header values.

successful operation
*/
type GetProductIDOK struct {
	Payload *models.ResponseWrapperProduct
}

func (o *GetProductIDOK) Error() string {
	return fmt.Sprintf("[GET /product/{id}][%d] getProductIdOK  %+v", 200, o.Payload)
}

func (o *GetProductIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProduct)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
