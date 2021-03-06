// Code generated by go-swagger; DO NOT EDIT.

package annual_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetLedgerAnnualAccountReader is a Reader for the GetLedgerAnnualAccount structure.
type GetLedgerAnnualAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLedgerAnnualAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLedgerAnnualAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLedgerAnnualAccountOK creates a GetLedgerAnnualAccountOK with default headers values
func NewGetLedgerAnnualAccountOK() *GetLedgerAnnualAccountOK {
	return &GetLedgerAnnualAccountOK{}
}

/*GetLedgerAnnualAccountOK handles this case with default header values.

successful operation
*/
type GetLedgerAnnualAccountOK struct {
	Payload *models.ListResponseAnnualAccount
}

func (o *GetLedgerAnnualAccountOK) Error() string {
	return fmt.Sprintf("[GET /ledger/annualAccount][%d] getLedgerAnnualAccountOK  %+v", 200, o.Payload)
}

func (o *GetLedgerAnnualAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseAnnualAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
