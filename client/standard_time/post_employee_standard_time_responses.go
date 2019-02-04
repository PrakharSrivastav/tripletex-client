// Code generated by go-swagger; DO NOT EDIT.

package standard_time

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PostEmployeeStandardTimeReader is a Reader for the PostEmployeeStandardTime structure.
type PostEmployeeStandardTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEmployeeStandardTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostEmployeeStandardTimeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostEmployeeStandardTimeCreated creates a PostEmployeeStandardTimeCreated with default headers values
func NewPostEmployeeStandardTimeCreated() *PostEmployeeStandardTimeCreated {
	return &PostEmployeeStandardTimeCreated{}
}

/*PostEmployeeStandardTimeCreated handles this case with default header values.

successfully created
*/
type PostEmployeeStandardTimeCreated struct {
	Payload *models.ResponseWrapperStandardTime
}

func (o *PostEmployeeStandardTimeCreated) Error() string {
	return fmt.Sprintf("[POST /employee/standardTime][%d] postEmployeeStandardTimeCreated  %+v", 201, o.Payload)
}

func (o *PostEmployeeStandardTimeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperStandardTime)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
