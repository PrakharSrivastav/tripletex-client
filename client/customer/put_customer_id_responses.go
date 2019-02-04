// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PutCustomerIDReader is a Reader for the PutCustomerID structure.
type PutCustomerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomerIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomerIDOK creates a PutCustomerIDOK with default headers values
func NewPutCustomerIDOK() *PutCustomerIDOK {
	return &PutCustomerIDOK{}
}

/*PutCustomerIDOK handles this case with default header values.

successful operation
*/
type PutCustomerIDOK struct {
	Payload *models.ResponseWrapperCustomer
}

func (o *PutCustomerIDOK) Error() string {
	return fmt.Sprintf("[PUT /customer/{id}][%d] putCustomerIdOK  %+v", 200, o.Payload)
}

func (o *PutCustomerIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperCustomer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}