// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutInvoiceIDSendReader is a Reader for the PutInvoiceIDSend structure.
type PutInvoiceIDSendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutInvoiceIDSendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPutInvoiceIDSendDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPutInvoiceIDSendDefault creates a PutInvoiceIDSendDefault with default headers values
func NewPutInvoiceIDSendDefault(code int) *PutInvoiceIDSendDefault {
	return &PutInvoiceIDSendDefault{
		_statusCode: code,
	}
}

/*PutInvoiceIDSendDefault handles this case with default header values.

successful operation
*/
type PutInvoiceIDSendDefault struct {
	_statusCode int
}

// Code gets the status code for the put invoice ID send default response
func (o *PutInvoiceIDSendDefault) Code() int {
	return o._statusCode
}

func (o *PutInvoiceIDSendDefault) Error() string {
	return fmt.Sprintf("[PUT /invoice/{id}/:send][%d] PutInvoiceIDSend default ", o._statusCode)
}

func (o *PutInvoiceIDSendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
