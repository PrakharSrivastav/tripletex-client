// Code generated by go-swagger; DO NOT EDIT.

package employee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PostEmployeeListReader is a Reader for the PostEmployeeList structure.
type PostEmployeeListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEmployeeListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostEmployeeListCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostEmployeeListCreated creates a PostEmployeeListCreated with default headers values
func NewPostEmployeeListCreated() *PostEmployeeListCreated {
	return &PostEmployeeListCreated{}
}

/*PostEmployeeListCreated handles this case with default header values.

successfully created
*/
type PostEmployeeListCreated struct {
	Payload *models.ListResponseEmployee
}

func (o *PostEmployeeListCreated) Error() string {
	return fmt.Sprintf("[POST /employee/list][%d] postEmployeeListCreated  %+v", 201, o.Payload)
}

func (o *PostEmployeeListCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseEmployee)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
