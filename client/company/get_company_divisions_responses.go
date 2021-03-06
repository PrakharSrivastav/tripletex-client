// Code generated by go-swagger; DO NOT EDIT.

package company

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetCompanyDivisionsReader is a Reader for the GetCompanyDivisions structure.
type GetCompanyDivisionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCompanyDivisionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCompanyDivisionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCompanyDivisionsOK creates a GetCompanyDivisionsOK with default headers values
func NewGetCompanyDivisionsOK() *GetCompanyDivisionsOK {
	return &GetCompanyDivisionsOK{}
}

/*GetCompanyDivisionsOK handles this case with default header values.

successful operation
*/
type GetCompanyDivisionsOK struct {
	Payload *models.ListResponseCompany
}

func (o *GetCompanyDivisionsOK) Error() string {
	return fmt.Sprintf("[GET /company/divisions][%d] getCompanyDivisionsOK  %+v", 200, o.Payload)
}

func (o *GetCompanyDivisionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseCompany)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
