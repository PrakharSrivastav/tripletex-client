// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetEventSubscriptionIDReader is a Reader for the GetEventSubscriptionID structure.
type GetEventSubscriptionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventSubscriptionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventSubscriptionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventSubscriptionIDOK creates a GetEventSubscriptionIDOK with default headers values
func NewGetEventSubscriptionIDOK() *GetEventSubscriptionIDOK {
	return &GetEventSubscriptionIDOK{}
}

/*GetEventSubscriptionIDOK handles this case with default header values.

successful operation
*/
type GetEventSubscriptionIDOK struct {
	Payload *models.ResponseWrapperSubscription
}

func (o *GetEventSubscriptionIDOK) Error() string {
	return fmt.Sprintf("[GET /event/subscription/{id}][%d] getEventSubscriptionIdOK  %+v", 200, o.Payload)
}

func (o *GetEventSubscriptionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperSubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
