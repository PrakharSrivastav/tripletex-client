// Code generated by go-swagger; DO NOT EDIT.

package prospect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetCrmProspectIDReader is a Reader for the GetCrmProspectID structure.
type GetCrmProspectIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCrmProspectIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCrmProspectIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCrmProspectIDOK creates a GetCrmProspectIDOK with default headers values
func NewGetCrmProspectIDOK() *GetCrmProspectIDOK {
	return &GetCrmProspectIDOK{}
}

/*GetCrmProspectIDOK handles this case with default header values.

successful operation
*/
type GetCrmProspectIDOK struct {
	Payload *models.ResponseWrapperProspect
}

func (o *GetCrmProspectIDOK) Error() string {
	return fmt.Sprintf("[GET /crm/prospect/{id}][%d] getCrmProspectIdOK  %+v", 200, o.Payload)
}

func (o *GetCrmProspectIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProspect)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
