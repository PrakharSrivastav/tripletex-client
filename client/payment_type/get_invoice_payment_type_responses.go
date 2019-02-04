// Code generated by go-swagger; DO NOT EDIT.

package payment_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetInvoicePaymentTypeReader is a Reader for the GetInvoicePaymentType structure.
type GetInvoicePaymentTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoicePaymentTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoicePaymentTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInvoicePaymentTypeOK creates a GetInvoicePaymentTypeOK with default headers values
func NewGetInvoicePaymentTypeOK() *GetInvoicePaymentTypeOK {
	return &GetInvoicePaymentTypeOK{}
}

/*GetInvoicePaymentTypeOK handles this case with default header values.

successful operation
*/
type GetInvoicePaymentTypeOK struct {
	Payload *models.ListResponsePaymentType
}

func (o *GetInvoicePaymentTypeOK) Error() string {
	return fmt.Sprintf("[GET /invoice/paymentType][%d] getInvoicePaymentTypeOK  %+v", 200, o.Payload)
}

func (o *GetInvoicePaymentTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponsePaymentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
