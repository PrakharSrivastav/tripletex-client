// Code generated by go-swagger; DO NOT EDIT.

package holiday

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/Petstore/models"
)

// GetSalarySettingsHolidayReader is a Reader for the GetSalarySettingsHoliday structure.
type GetSalarySettingsHolidayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSalarySettingsHolidayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSalarySettingsHolidayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSalarySettingsHolidayOK creates a GetSalarySettingsHolidayOK with default headers values
func NewGetSalarySettingsHolidayOK() *GetSalarySettingsHolidayOK {
	return &GetSalarySettingsHolidayOK{}
}

/*GetSalarySettingsHolidayOK handles this case with default header values.

successful operation
*/
type GetSalarySettingsHolidayOK struct {
	Payload *models.ListResponseCompanyHoliday
}

func (o *GetSalarySettingsHolidayOK) Error() string {
	return fmt.Sprintf("[GET /salary/settings/holiday][%d] getSalarySettingsHolidayOK  %+v", 200, o.Payload)
}

func (o *GetSalarySettingsHolidayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseCompanyHoliday)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
