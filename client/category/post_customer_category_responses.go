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

// PostCustomerCategoryReader is a Reader for the PostCustomerCategory structure.
type PostCustomerCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomerCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostCustomerCategoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCustomerCategoryCreated creates a PostCustomerCategoryCreated with default headers values
func NewPostCustomerCategoryCreated() *PostCustomerCategoryCreated {
	return &PostCustomerCategoryCreated{}
}

/*PostCustomerCategoryCreated handles this case with default header values.

successfully created
*/
type PostCustomerCategoryCreated struct {
	Payload *models.ResponseWrapperCustomerCategory
}

func (o *PostCustomerCategoryCreated) Error() string {
	return fmt.Sprintf("[POST /customer/category][%d] postCustomerCategoryCreated  %+v", 201, o.Payload)
}

func (o *PostCustomerCategoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperCustomerCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
