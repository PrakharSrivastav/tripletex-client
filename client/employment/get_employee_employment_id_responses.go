// Code generated by go-swagger; DO NOT EDIT.

package employment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetEmployeeEmploymentIDReader is a Reader for the GetEmployeeEmploymentID structure.
type GetEmployeeEmploymentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmployeeEmploymentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEmployeeEmploymentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmployeeEmploymentIDOK creates a GetEmployeeEmploymentIDOK with default headers values
func NewGetEmployeeEmploymentIDOK() *GetEmployeeEmploymentIDOK {
	return &GetEmployeeEmploymentIDOK{}
}

/*GetEmployeeEmploymentIDOK handles this case with default header values.

successful operation
*/
type GetEmployeeEmploymentIDOK struct {
	Payload *models.ResponseWrapperEmployment
}

func (o *GetEmployeeEmploymentIDOK) Error() string {
	return fmt.Sprintf("[GET /employee/employment/{id}][%d] getEmployeeEmploymentIdOK  %+v", 200, o.Payload)
}

func (o *GetEmployeeEmploymentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperEmployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
