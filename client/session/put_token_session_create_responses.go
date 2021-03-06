// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PutTokenSessionCreateReader is a Reader for the PutTokenSessionCreate structure.
type PutTokenSessionCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTokenSessionCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutTokenSessionCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutTokenSessionCreateOK creates a PutTokenSessionCreateOK with default headers values
func NewPutTokenSessionCreateOK() *PutTokenSessionCreateOK {
	return &PutTokenSessionCreateOK{}
}

/*PutTokenSessionCreateOK handles this case with default header values.

successful operation
*/
type PutTokenSessionCreateOK struct {
	Payload *models.ResponseWrapperSessionToken
}

func (o *PutTokenSessionCreateOK) Error() string {
	return fmt.Sprintf("[PUT /token/session/:create][%d] putTokenSessionCreateOK  %+v", 200, o.Payload)
}

func (o *PutTokenSessionCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperSessionToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
