// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// PostProjectCategoryReader is a Reader for the PostProjectCategory structure.
type PostProjectCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProjectCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostProjectCategoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostProjectCategoryCreated creates a PostProjectCategoryCreated with default headers values
func NewPostProjectCategoryCreated() *PostProjectCategoryCreated {
	return &PostProjectCategoryCreated{}
}

/*PostProjectCategoryCreated handles this case with default header values.

successfully created
*/
type PostProjectCategoryCreated struct {
	Payload *models.ResponseWrapperProjectCategory
}

func (o *PostProjectCategoryCreated) Error() string {
	return fmt.Sprintf("[POST /project/category][%d] postProjectCategoryCreated  %+v", 201, o.Payload)
}

func (o *PostProjectCategoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProjectCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
