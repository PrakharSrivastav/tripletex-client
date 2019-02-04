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

// GetInvoiceReader is a Reader for the GetInvoice structure.
type GetInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInvoiceOK creates a GetInvoiceOK with default headers values
func NewGetInvoiceOK() *GetInvoiceOK {
	return &GetInvoiceOK{}
}

/*GetInvoiceOK handles this case with default header values.

successful operation
*/
type GetInvoiceOK struct {
	Payload *models.ListResponseInvoice
}

func (o *GetInvoiceOK) Error() string {
	return fmt.Sprintf("[GET /invoice][%d] getInvoiceOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseInvoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}