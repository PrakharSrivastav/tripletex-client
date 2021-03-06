// Code generated by go-swagger; DO NOT EDIT.

package remuneration_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetEmployeeEmploymentRemunerationTypeReader is a Reader for the GetEmployeeEmploymentRemunerationType structure.
type GetEmployeeEmploymentRemunerationTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmployeeEmploymentRemunerationTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEmployeeEmploymentRemunerationTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmployeeEmploymentRemunerationTypeOK creates a GetEmployeeEmploymentRemunerationTypeOK with default headers values
func NewGetEmployeeEmploymentRemunerationTypeOK() *GetEmployeeEmploymentRemunerationTypeOK {
	return &GetEmployeeEmploymentRemunerationTypeOK{}
}

/*GetEmployeeEmploymentRemunerationTypeOK handles this case with default header values.

successful operation
*/
type GetEmployeeEmploymentRemunerationTypeOK struct {
	Payload *models.ListResponseRemunerationType
}

func (o *GetEmployeeEmploymentRemunerationTypeOK) Error() string {
	return fmt.Sprintf("[GET /employee/employment/remunerationType][%d] getEmployeeEmploymentRemunerationTypeOK  %+v", 200, o.Payload)
}

func (o *GetEmployeeEmploymentRemunerationTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseRemunerationType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
