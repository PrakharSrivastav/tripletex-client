// Code generated by go-swagger; DO NOT EDIT.

package department

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PostDepartmentReader is a Reader for the PostDepartment structure.
type PostDepartmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDepartmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostDepartmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDepartmentCreated creates a PostDepartmentCreated with default headers values
func NewPostDepartmentCreated() *PostDepartmentCreated {
	return &PostDepartmentCreated{}
}

/*PostDepartmentCreated handles this case with default header values.

successfully created
*/
type PostDepartmentCreated struct {
	Payload *models.ResponseWrapperDepartment
}

func (o *PostDepartmentCreated) Error() string {
	return fmt.Sprintf("[POST /department][%d] postDepartmentCreated  %+v", 201, o.Payload)
}

func (o *PostDepartmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperDepartment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
