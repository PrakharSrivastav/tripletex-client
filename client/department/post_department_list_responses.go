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

// PostDepartmentListReader is a Reader for the PostDepartmentList structure.
type PostDepartmentListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDepartmentListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostDepartmentListCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDepartmentListCreated creates a PostDepartmentListCreated with default headers values
func NewPostDepartmentListCreated() *PostDepartmentListCreated {
	return &PostDepartmentListCreated{}
}

/*PostDepartmentListCreated handles this case with default header values.

successfully created
*/
type PostDepartmentListCreated struct {
	Payload *models.ListResponseDepartment
}

func (o *PostDepartmentListCreated) Error() string {
	return fmt.Sprintf("[POST /department/list][%d] postDepartmentListCreated  %+v", 201, o.Payload)
}

func (o *PostDepartmentListCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseDepartment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
