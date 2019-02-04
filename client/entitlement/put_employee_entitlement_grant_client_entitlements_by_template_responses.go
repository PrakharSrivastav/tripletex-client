// Code generated by go-swagger; DO NOT EDIT.

package entitlement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutEmployeeEntitlementGrantClientEntitlementsByTemplateReader is a Reader for the PutEmployeeEntitlementGrantClientEntitlementsByTemplate structure.
type PutEmployeeEntitlementGrantClientEntitlementsByTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutEmployeeEntitlementGrantClientEntitlementsByTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPutEmployeeEntitlementGrantClientEntitlementsByTemplateDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPutEmployeeEntitlementGrantClientEntitlementsByTemplateDefault creates a PutEmployeeEntitlementGrantClientEntitlementsByTemplateDefault with default headers values
func NewPutEmployeeEntitlementGrantClientEntitlementsByTemplateDefault(code int) *PutEmployeeEntitlementGrantClientEntitlementsByTemplateDefault {
	return &PutEmployeeEntitlementGrantClientEntitlementsByTemplateDefault{
		_statusCode: code,
	}
}

/*PutEmployeeEntitlementGrantClientEntitlementsByTemplateDefault handles this case with default header values.

successful operation
*/
type PutEmployeeEntitlementGrantClientEntitlementsByTemplateDefault struct {
	_statusCode int
}

// Code gets the status code for the put employee entitlement grant client entitlements by template default response
func (o *PutEmployeeEntitlementGrantClientEntitlementsByTemplateDefault) Code() int {
	return o._statusCode
}

func (o *PutEmployeeEntitlementGrantClientEntitlementsByTemplateDefault) Error() string {
	return fmt.Sprintf("[PUT /employee/entitlement/:grantClientEntitlementsByTemplate][%d] PutEmployeeEntitlementGrantClientEntitlementsByTemplate default ", o._statusCode)
}

func (o *PutEmployeeEntitlementGrantClientEntitlementsByTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}