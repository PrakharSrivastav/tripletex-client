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

// GetLedgerAccountIDReader is a Reader for the GetLedgerAccountID structure.
type GetLedgerAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLedgerAccountIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLedgerAccountIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLedgerAccountIDOK creates a GetLedgerAccountIDOK with default headers values
func NewGetLedgerAccountIDOK() *GetLedgerAccountIDOK {
	return &GetLedgerAccountIDOK{}
}

/*GetLedgerAccountIDOK handles this case with default header values.

successful operation
*/
type GetLedgerAccountIDOK struct {
	Payload *models.ResponseWrapperAccount
}

func (o *GetLedgerAccountIDOK) Error() string {
	return fmt.Sprintf("[GET /ledger/account/{id}][%d] getLedgerAccountIdOK  %+v", 200, o.Payload)
}

func (o *GetLedgerAccountIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}