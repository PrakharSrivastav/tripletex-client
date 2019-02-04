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

// PutBankReconciliationMatchSuggestReader is a Reader for the PutBankReconciliationMatchSuggest structure.
type PutBankReconciliationMatchSuggestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutBankReconciliationMatchSuggestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutBankReconciliationMatchSuggestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutBankReconciliationMatchSuggestOK creates a PutBankReconciliationMatchSuggestOK with default headers values
func NewPutBankReconciliationMatchSuggestOK() *PutBankReconciliationMatchSuggestOK {
	return &PutBankReconciliationMatchSuggestOK{}
}

/*PutBankReconciliationMatchSuggestOK handles this case with default header values.

successful operation
*/
type PutBankReconciliationMatchSuggestOK struct {
	Payload *models.ListResponseBankReconciliationMatch
}

func (o *PutBankReconciliationMatchSuggestOK) Error() string {
	return fmt.Sprintf("[PUT /bank/reconciliation/match/:suggest][%d] putBankReconciliationMatchSuggestOK  %+v", 200, o.Payload)
}

func (o *PutBankReconciliationMatchSuggestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseBankReconciliationMatch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
