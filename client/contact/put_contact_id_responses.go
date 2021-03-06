// Code generated by go-swagger; DO NOT EDIT.

package contact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PutContactIDReader is a Reader for the PutContactID structure.
type PutContactIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutContactIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutContactIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutContactIDOK creates a PutContactIDOK with default headers values
func NewPutContactIDOK() *PutContactIDOK {
	return &PutContactIDOK{}
}

/*PutContactIDOK handles this case with default header values.

successful operation
*/
type PutContactIDOK struct {
	Payload *models.ResponseWrapperContact
}

func (o *PutContactIDOK) Error() string {
	return fmt.Sprintf("[PUT /contact/{id}][%d] putContactIdOK  %+v", 200, o.Payload)
}

func (o *PutContactIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
