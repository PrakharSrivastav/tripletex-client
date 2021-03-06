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

// GetEmployeeReader is a Reader for the GetEmployee structure.
type GetEmployeeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmployeeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEmployeeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmployeeOK creates a GetEmployeeOK with default headers values
func NewGetEmployeeOK() *GetEmployeeOK {
	return &GetEmployeeOK{}
}

/*GetEmployeeOK handles this case with default header values.

successful operation
*/
type GetEmployeeOK struct {
	Payload *models.ListResponseEmployee
}

func (o *GetEmployeeOK) Error() string {
	return fmt.Sprintf("[GET /employee][%d] getEmployeeOK  %+v", 200, o.Payload)
}

func (o *GetEmployeeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseEmployee)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
