// Code generated by go-swagger; DO NOT EDIT.

package address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetAddressIDReader is a Reader for the GetAddressID structure.
type GetAddressIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAddressIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAddressIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAddressIDOK creates a GetAddressIDOK with default headers values
func NewGetAddressIDOK() *GetAddressIDOK {
	return &GetAddressIDOK{}
}

/*GetAddressIDOK handles this case with default header values.

successful operation
*/
type GetAddressIDOK struct {
	Payload *models.ResponseWrapperAddress
}

func (o *GetAddressIDOK) Error() string {
	return fmt.Sprintf("[GET /address/{id}][%d] getAddressIdOK  %+v", 200, o.Payload)
}

func (o *GetAddressIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}