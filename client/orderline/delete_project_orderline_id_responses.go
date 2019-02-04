// Code generated by go-swagger; DO NOT EDIT.

package orderline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteProjectOrderlineIDReader is a Reader for the DeleteProjectOrderlineID structure.
type DeleteProjectOrderlineIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectOrderlineIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteProjectOrderlineIDDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteProjectOrderlineIDDefault creates a DeleteProjectOrderlineIDDefault with default headers values
func NewDeleteProjectOrderlineIDDefault(code int) *DeleteProjectOrderlineIDDefault {
	return &DeleteProjectOrderlineIDDefault{
		_statusCode: code,
	}
}

/*DeleteProjectOrderlineIDDefault handles this case with default header values.

successful operation
*/
type DeleteProjectOrderlineIDDefault struct {
	_statusCode int
}

// Code gets the status code for the delete project orderline ID default response
func (o *DeleteProjectOrderlineIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteProjectOrderlineIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /project/orderline/{id}][%d] DeleteProjectOrderlineID default ", o._statusCode)
}

func (o *DeleteProjectOrderlineIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
