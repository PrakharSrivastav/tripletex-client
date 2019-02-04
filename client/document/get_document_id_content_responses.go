// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetDocumentIDContentReader is a Reader for the GetDocumentIDContent structure.
type GetDocumentIDContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDocumentIDContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDocumentIDContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDocumentIDContentOK creates a GetDocumentIDContentOK with default headers values
func NewGetDocumentIDContentOK() *GetDocumentIDContentOK {
	return &GetDocumentIDContentOK{}
}

/*GetDocumentIDContentOK handles this case with default header values.

successful operation
*/
type GetDocumentIDContentOK struct {
	Payload strfmt.Base64
}

func (o *GetDocumentIDContentOK) Error() string {
	return fmt.Sprintf("[GET /document/{id}/content][%d] getDocumentIdContentOK  %+v", 200, o.Payload)
}

func (o *GetDocumentIDContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}