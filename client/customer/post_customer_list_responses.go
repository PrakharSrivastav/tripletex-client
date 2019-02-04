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

// PostCustomerListReader is a Reader for the PostCustomerList structure.
type PostCustomerListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomerListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostCustomerListCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCustomerListCreated creates a PostCustomerListCreated with default headers values
func NewPostCustomerListCreated() *PostCustomerListCreated {
	return &PostCustomerListCreated{}
}

/*PostCustomerListCreated handles this case with default header values.

successfully created
*/
type PostCustomerListCreated struct {
	Payload *models.ListResponseCustomer
}

func (o *PostCustomerListCreated) Error() string {
	return fmt.Sprintf("[POST /customer/list][%d] postCustomerListCreated  %+v", 201, o.Payload)
}

func (o *PostCustomerListCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseCustomer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
