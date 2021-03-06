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

// PutEventSubscriptionIDReader is a Reader for the PutEventSubscriptionID structure.
type PutEventSubscriptionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutEventSubscriptionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutEventSubscriptionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutEventSubscriptionIDOK creates a PutEventSubscriptionIDOK with default headers values
func NewPutEventSubscriptionIDOK() *PutEventSubscriptionIDOK {
	return &PutEventSubscriptionIDOK{}
}

/*PutEventSubscriptionIDOK handles this case with default header values.

successful operation
*/
type PutEventSubscriptionIDOK struct {
	Payload *models.ResponseWrapperSubscription
}

func (o *PutEventSubscriptionIDOK) Error() string {
	return fmt.Sprintf("[PUT /event/subscription/{id}][%d] putEventSubscriptionIdOK  %+v", 200, o.Payload)
}

func (o *PutEventSubscriptionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperSubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
