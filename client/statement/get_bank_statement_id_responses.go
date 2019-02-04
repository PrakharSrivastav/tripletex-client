// Code generated by go-swagger; DO NOT EDIT.

package statement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetBankStatementIDReader is a Reader for the GetBankStatementID structure.
type GetBankStatementIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBankStatementIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBankStatementIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBankStatementIDOK creates a GetBankStatementIDOK with default headers values
func NewGetBankStatementIDOK() *GetBankStatementIDOK {
	return &GetBankStatementIDOK{}
}

/*GetBankStatementIDOK handles this case with default header values.

successful operation
*/
type GetBankStatementIDOK struct {
	Payload *models.ResponseWrapperBankStatement
}

func (o *GetBankStatementIDOK) Error() string {
	return fmt.Sprintf("[GET /bank/statement/{id}][%d] getBankStatementIdOK  %+v", 200, o.Payload)
}

func (o *GetBankStatementIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBankStatement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}