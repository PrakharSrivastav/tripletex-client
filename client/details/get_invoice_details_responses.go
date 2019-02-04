// Code generated by go-swagger; DO NOT EDIT.

package details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetInvoiceDetailsReader is a Reader for the GetInvoiceDetails structure.
type GetInvoiceDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInvoiceDetailsOK creates a GetInvoiceDetailsOK with default headers values
func NewGetInvoiceDetailsOK() *GetInvoiceDetailsOK {
	return &GetInvoiceDetailsOK{}
}

/*GetInvoiceDetailsOK handles this case with default header values.

successful operation
*/
type GetInvoiceDetailsOK struct {
	Payload *models.ListResponseProjectInvoiceDetails
}

func (o *GetInvoiceDetailsOK) Error() string {
	return fmt.Sprintf("[GET /invoice/details][%d] getInvoiceDetailsOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseProjectInvoiceDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
