// Code generated by go-swagger; DO NOT EDIT.

package accounting_period

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetLedgerAccountingPeriodIDReader is a Reader for the GetLedgerAccountingPeriodID structure.
type GetLedgerAccountingPeriodIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLedgerAccountingPeriodIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLedgerAccountingPeriodIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLedgerAccountingPeriodIDOK creates a GetLedgerAccountingPeriodIDOK with default headers values
func NewGetLedgerAccountingPeriodIDOK() *GetLedgerAccountingPeriodIDOK {
	return &GetLedgerAccountingPeriodIDOK{}
}

/*GetLedgerAccountingPeriodIDOK handles this case with default header values.

successful operation
*/
type GetLedgerAccountingPeriodIDOK struct {
	Payload *models.ResponseWrapperAccountingPeriod
}

func (o *GetLedgerAccountingPeriodIDOK) Error() string {
	return fmt.Sprintf("[GET /ledger/accountingPeriod/{id}][%d] getLedgerAccountingPeriodIdOK  %+v", 200, o.Payload)
}

func (o *GetLedgerAccountingPeriodIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperAccountingPeriod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
