// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PutLedgerAccountListReader is a Reader for the PutLedgerAccountList structure.
type PutLedgerAccountListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLedgerAccountListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutLedgerAccountListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutLedgerAccountListOK creates a PutLedgerAccountListOK with default headers values
func NewPutLedgerAccountListOK() *PutLedgerAccountListOK {
	return &PutLedgerAccountListOK{}
}

/*PutLedgerAccountListOK handles this case with default header values.

successful operation
*/
type PutLedgerAccountListOK struct {
	Payload *models.ListResponseAccount
}

func (o *PutLedgerAccountListOK) Error() string {
	return fmt.Sprintf("[PUT /ledger/account/list][%d] putLedgerAccountListOK  %+v", 200, o.Payload)
}

func (o *PutLedgerAccountListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}