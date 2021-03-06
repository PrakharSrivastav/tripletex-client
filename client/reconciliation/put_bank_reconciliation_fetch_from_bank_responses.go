// Code generated by go-swagger; DO NOT EDIT.

package reconciliation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PutBankReconciliationFetchFromBankReader is a Reader for the PutBankReconciliationFetchFromBank structure.
type PutBankReconciliationFetchFromBankReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutBankReconciliationFetchFromBankReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutBankReconciliationFetchFromBankOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutBankReconciliationFetchFromBankOK creates a PutBankReconciliationFetchFromBankOK with default headers values
func NewPutBankReconciliationFetchFromBankOK() *PutBankReconciliationFetchFromBankOK {
	return &PutBankReconciliationFetchFromBankOK{}
}

/*PutBankReconciliationFetchFromBankOK handles this case with default header values.

successful operation
*/
type PutBankReconciliationFetchFromBankOK struct {
	Payload *models.ResponseWrapperBankReconciliation
}

func (o *PutBankReconciliationFetchFromBankOK) Error() string {
	return fmt.Sprintf("[PUT /bank/reconciliation/:fetchFromBank][%d] putBankReconciliationFetchFromBankOK  %+v", 200, o.Payload)
}

func (o *PutBankReconciliationFetchFromBankOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBankReconciliation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
