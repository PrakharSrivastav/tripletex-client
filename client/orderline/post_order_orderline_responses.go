// Code generated by go-swagger; DO NOT EDIT.

package orderline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PostOrderOrderlineReader is a Reader for the PostOrderOrderline structure.
type PostOrderOrderlineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrderOrderlineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostOrderOrderlineCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostOrderOrderlineCreated creates a PostOrderOrderlineCreated with default headers values
func NewPostOrderOrderlineCreated() *PostOrderOrderlineCreated {
	return &PostOrderOrderlineCreated{}
}

/*PostOrderOrderlineCreated handles this case with default header values.

successfully created
*/
type PostOrderOrderlineCreated struct {
	Payload *models.ResponseWrapperOrderLine
}

func (o *PostOrderOrderlineCreated) Error() string {
	return fmt.Sprintf("[POST /order/orderline][%d] postOrderOrderlineCreated  %+v", 201, o.Payload)
}

func (o *PostOrderOrderlineCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperOrderLine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
