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

// PutDivisionListReader is a Reader for the PutDivisionList structure.
type PutDivisionListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDivisionListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutDivisionListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDivisionListOK creates a PutDivisionListOK with default headers values
func NewPutDivisionListOK() *PutDivisionListOK {
	return &PutDivisionListOK{}
}

/*PutDivisionListOK handles this case with default header values.

successful operation
*/
type PutDivisionListOK struct {
	Payload *models.ListResponseDivision
}

func (o *PutDivisionListOK) Error() string {
	return fmt.Sprintf("[PUT /division/list][%d] putDivisionListOK  %+v", 200, o.Payload)
}

func (o *PutDivisionListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseDivision)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}