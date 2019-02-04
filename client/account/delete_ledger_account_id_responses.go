// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteLedgerAccountIDReader is a Reader for the DeleteLedgerAccountID structure.
type DeleteLedgerAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLedgerAccountIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteLedgerAccountIDDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteLedgerAccountIDDefault creates a DeleteLedgerAccountIDDefault with default headers values
func NewDeleteLedgerAccountIDDefault(code int) *DeleteLedgerAccountIDDefault {
	return &DeleteLedgerAccountIDDefault{
		_statusCode: code,
	}
}

/*DeleteLedgerAccountIDDefault handles this case with default header values.

successful operation
*/
type DeleteLedgerAccountIDDefault struct {
	_statusCode int
}

// Code gets the status code for the delete ledger account ID default response
func (o *DeleteLedgerAccountIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLedgerAccountIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /ledger/account/{id}][%d] DeleteLedgerAccountID default ", o._statusCode)
}

func (o *DeleteLedgerAccountIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
