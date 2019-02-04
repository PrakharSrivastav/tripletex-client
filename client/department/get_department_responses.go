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

// GetDepartmentReader is a Reader for the GetDepartment structure.
type GetDepartmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDepartmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDepartmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDepartmentOK creates a GetDepartmentOK with default headers values
func NewGetDepartmentOK() *GetDepartmentOK {
	return &GetDepartmentOK{}
}

/*GetDepartmentOK handles this case with default header values.

successful operation
*/
type GetDepartmentOK struct {
	Payload *models.ListResponseDepartment
}

func (o *GetDepartmentOK) Error() string {
	return fmt.Sprintf("[GET /department][%d] getDepartmentOK  %+v", 200, o.Payload)
}

func (o *GetDepartmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseDepartment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}