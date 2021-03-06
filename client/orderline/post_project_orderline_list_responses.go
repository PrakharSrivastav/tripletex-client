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

// PostProjectOrderlineListReader is a Reader for the PostProjectOrderlineList structure.
type PostProjectOrderlineListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProjectOrderlineListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostProjectOrderlineListCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostProjectOrderlineListCreated creates a PostProjectOrderlineListCreated with default headers values
func NewPostProjectOrderlineListCreated() *PostProjectOrderlineListCreated {
	return &PostProjectOrderlineListCreated{}
}

/*PostProjectOrderlineListCreated handles this case with default header values.

successfully created
*/
type PostProjectOrderlineListCreated struct {
	Payload *models.ListResponseProjectOrderLine
}

func (o *PostProjectOrderlineListCreated) Error() string {
	return fmt.Sprintf("[POST /project/orderline/list][%d] postProjectOrderlineListCreated  %+v", 201, o.Payload)
}

func (o *PostProjectOrderlineListCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseProjectOrderLine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
