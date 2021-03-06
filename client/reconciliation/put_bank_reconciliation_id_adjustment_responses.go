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

// PutBankReconciliationIDAdjustmentReader is a Reader for the PutBankReconciliationIDAdjustment structure.
type PutBankReconciliationIDAdjustmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutBankReconciliationIDAdjustmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutBankReconciliationIDAdjustmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutBankReconciliationIDAdjustmentOK creates a PutBankReconciliationIDAdjustmentOK with default headers values
func NewPutBankReconciliationIDAdjustmentOK() *PutBankReconciliationIDAdjustmentOK {
	return &PutBankReconciliationIDAdjustmentOK{}
}

/*PutBankReconciliationIDAdjustmentOK handles this case with default header values.

successful operation
*/
type PutBankReconciliationIDAdjustmentOK struct {
	Payload *models.ListResponseBankReconciliationAdjustment
}

func (o *PutBankReconciliationIDAdjustmentOK) Error() string {
	return fmt.Sprintf("[PUT /bank/reconciliation/{id}/:adjustment][%d] putBankReconciliationIdAdjustmentOK  %+v", 200, o.Payload)
}

func (o *PutBankReconciliationIDAdjustmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseBankReconciliationAdjustment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
