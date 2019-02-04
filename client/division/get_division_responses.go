// Code generated by go-swagger; DO NOT EDIT.

package division

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetDivisionReader is a Reader for the GetDivision structure.
type GetDivisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDivisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDivisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDivisionOK creates a GetDivisionOK with default headers values
func NewGetDivisionOK() *GetDivisionOK {
	return &GetDivisionOK{}
}

/*GetDivisionOK handles this case with default header values.

successful operation
*/
type GetDivisionOK struct {
	Payload *models.ListResponseDivision
}

func (o *GetDivisionOK) Error() string {
	return fmt.Sprintf("[GET /division][%d] getDivisionOK  %+v", 200, o.Payload)
}

func (o *GetDivisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseDivision)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}