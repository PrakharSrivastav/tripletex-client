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

// GetLedgerAccountReader is a Reader for the GetLedgerAccount structure.
type GetLedgerAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLedgerAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLedgerAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLedgerAccountOK creates a GetLedgerAccountOK with default headers values
func NewGetLedgerAccountOK() *GetLedgerAccountOK {
	return &GetLedgerAccountOK{}
}

/*GetLedgerAccountOK handles this case with default header values.

successful operation
*/
type GetLedgerAccountOK struct {
	Payload *models.ListResponseAccount
}

func (o *GetLedgerAccountOK) Error() string {
	return fmt.Sprintf("[GET /ledger/account][%d] getLedgerAccountOK  %+v", 200, o.Payload)
}

func (o *GetLedgerAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}