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

// GetDepartmentIDReader is a Reader for the GetDepartmentID structure.
type GetDepartmentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDepartmentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDepartmentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDepartmentIDOK creates a GetDepartmentIDOK with default headers values
func NewGetDepartmentIDOK() *GetDepartmentIDOK {
	return &GetDepartmentIDOK{}
}

/*GetDepartmentIDOK handles this case with default header values.

successful operation
*/
type GetDepartmentIDOK struct {
	Payload *models.ResponseWrapperDepartment
}

func (o *GetDepartmentIDOK) Error() string {
	return fmt.Sprintf("[GET /department/{id}][%d] getDepartmentIdOK  %+v", 200, o.Payload)
}

func (o *GetDepartmentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperDepartment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
