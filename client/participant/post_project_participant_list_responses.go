// Code generated by go-swagger; DO NOT EDIT.

package participant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PostProjectParticipantListReader is a Reader for the PostProjectParticipantList structure.
type PostProjectParticipantListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProjectParticipantListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostProjectParticipantListCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostProjectParticipantListCreated creates a PostProjectParticipantListCreated with default headers values
func NewPostProjectParticipantListCreated() *PostProjectParticipantListCreated {
	return &PostProjectParticipantListCreated{}
}

/*PostProjectParticipantListCreated handles this case with default header values.

successfully created
*/
type PostProjectParticipantListCreated struct {
	Payload *models.ListResponseProjectParticipant
}

func (o *PostProjectParticipantListCreated) Error() string {
	return fmt.Sprintf("[POST /project/participant/list][%d] postProjectParticipantListCreated  %+v", 201, o.Payload)
}

func (o *PostProjectParticipantListCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseProjectParticipant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
