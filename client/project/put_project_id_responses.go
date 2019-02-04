// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PutProjectIDReader is a Reader for the PutProjectID structure.
type PutProjectIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutProjectIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutProjectIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutProjectIDOK creates a PutProjectIDOK with default headers values
func NewPutProjectIDOK() *PutProjectIDOK {
	return &PutProjectIDOK{}
}

/*PutProjectIDOK handles this case with default header values.

successful operation
*/
type PutProjectIDOK struct {
	Payload *models.ResponseWrapperProject
}

func (o *PutProjectIDOK) Error() string {
	return fmt.Sprintf("[PUT /project/{id}][%d] putProjectIdOK  %+v", 200, o.Payload)
}

func (o *PutProjectIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
