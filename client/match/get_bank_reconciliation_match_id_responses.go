// Code generated by go-swagger; DO NOT EDIT.

package match

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetBankReconciliationMatchIDReader is a Reader for the GetBankReconciliationMatchID structure.
type GetBankReconciliationMatchIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBankReconciliationMatchIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBankReconciliationMatchIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBankReconciliationMatchIDOK creates a GetBankReconciliationMatchIDOK with default headers values
func NewGetBankReconciliationMatchIDOK() *GetBankReconciliationMatchIDOK {
	return &GetBankReconciliationMatchIDOK{}
}

/*GetBankReconciliationMatchIDOK handles this case with default header values.

successful operation
*/
type GetBankReconciliationMatchIDOK struct {
	Payload *models.ResponseWrapperBankReconciliationMatch
}

func (o *GetBankReconciliationMatchIDOK) Error() string {
	return fmt.Sprintf("[GET /bank/reconciliation/match/{id}][%d] getBankReconciliationMatchIdOK  %+v", 200, o.Payload)
}

func (o *GetBankReconciliationMatchIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBankReconciliationMatch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
