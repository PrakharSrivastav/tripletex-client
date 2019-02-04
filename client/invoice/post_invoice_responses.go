// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PostInvoiceReader is a Reader for the PostInvoice structure.
type PostInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostInvoiceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostInvoiceCreated creates a PostInvoiceCreated with default headers values
func NewPostInvoiceCreated() *PostInvoiceCreated {
	return &PostInvoiceCreated{}
}

/*PostInvoiceCreated handles this case with default header values.

successfully created
*/
type PostInvoiceCreated struct {
	Payload *models.ResponseWrapperInvoice
}

func (o *PostInvoiceCreated) Error() string {
	return fmt.Sprintf("[POST /invoice][%d] postInvoiceCreated  %+v", 201, o.Payload)
}

func (o *PostInvoiceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperInvoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
