// Code generated by go-swagger; DO NOT EDIT.

package leave_of_absence_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetEmployeeEmploymentLeaveOfAbsenceTypeReader is a Reader for the GetEmployeeEmploymentLeaveOfAbsenceType structure.
type GetEmployeeEmploymentLeaveOfAbsenceTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEmployeeEmploymentLeaveOfAbsenceTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmployeeEmploymentLeaveOfAbsenceTypeOK creates a GetEmployeeEmploymentLeaveOfAbsenceTypeOK with default headers values
func NewGetEmployeeEmploymentLeaveOfAbsenceTypeOK() *GetEmployeeEmploymentLeaveOfAbsenceTypeOK {
	return &GetEmployeeEmploymentLeaveOfAbsenceTypeOK{}
}

/*GetEmployeeEmploymentLeaveOfAbsenceTypeOK handles this case with default header values.

successful operation
*/
type GetEmployeeEmploymentLeaveOfAbsenceTypeOK struct {
	Payload *models.ListResponseLeaveOfAbsenceType
}

func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeOK) Error() string {
	return fmt.Sprintf("[GET /employee/employment/leaveOfAbsenceType][%d] getEmployeeEmploymentLeaveOfAbsenceTypeOK  %+v", 200, o.Payload)
}

func (o *GetEmployeeEmploymentLeaveOfAbsenceTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseLeaveOfAbsenceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
