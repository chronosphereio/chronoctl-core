// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/models"
)

// UpdateSLOReader is a Reader for the UpdateSLO structure.
type UpdateSLOReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSLOReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSLOOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSLOBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSLONotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateSLOConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSLOInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateSLODefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSLOOK creates a UpdateSLOOK with default headers values
func NewUpdateSLOOK() *UpdateSLOOK {
	return &UpdateSLOOK{}
}

/*
UpdateSLOOK describes a response with status code 200, with default header values.

A successful response containing the updated SLO.
*/
type UpdateSLOOK struct {
	Payload *models.ConfigunstableUpdateSLOResponse
}

// IsSuccess returns true when this update Slo o k response has a 2xx status code
func (o *UpdateSLOOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update Slo o k response has a 3xx status code
func (o *UpdateSLOOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Slo o k response has a 4xx status code
func (o *UpdateSLOOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update Slo o k response has a 5xx status code
func (o *UpdateSLOOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update Slo o k response a status code equal to that given
func (o *UpdateSLOOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update Slo o k response
func (o *UpdateSLOOK) Code() int {
	return 200
}

func (o *UpdateSLOOK) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/slos/{slug}][%d] updateSloOK  %+v", 200, o.Payload)
}

func (o *UpdateSLOOK) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/slos/{slug}][%d] updateSloOK  %+v", 200, o.Payload)
}

func (o *UpdateSLOOK) GetPayload() *models.ConfigunstableUpdateSLOResponse {
	return o.Payload
}

func (o *UpdateSLOOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableUpdateSLOResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSLOBadRequest creates a UpdateSLOBadRequest with default headers values
func NewUpdateSLOBadRequest() *UpdateSLOBadRequest {
	return &UpdateSLOBadRequest{}
}

/*
UpdateSLOBadRequest describes a response with status code 400, with default header values.

Cannot update the SLO because the request is invalid.
*/
type UpdateSLOBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update Slo bad request response has a 2xx status code
func (o *UpdateSLOBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Slo bad request response has a 3xx status code
func (o *UpdateSLOBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Slo bad request response has a 4xx status code
func (o *UpdateSLOBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update Slo bad request response has a 5xx status code
func (o *UpdateSLOBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update Slo bad request response a status code equal to that given
func (o *UpdateSLOBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update Slo bad request response
func (o *UpdateSLOBadRequest) Code() int {
	return 400
}

func (o *UpdateSLOBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/slos/{slug}][%d] updateSloBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSLOBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/slos/{slug}][%d] updateSloBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSLOBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateSLOBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSLONotFound creates a UpdateSLONotFound with default headers values
func NewUpdateSLONotFound() *UpdateSLONotFound {
	return &UpdateSLONotFound{}
}

/*
UpdateSLONotFound describes a response with status code 404, with default header values.

Cannot update the SLO because the slug does not exist.
*/
type UpdateSLONotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update Slo not found response has a 2xx status code
func (o *UpdateSLONotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Slo not found response has a 3xx status code
func (o *UpdateSLONotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Slo not found response has a 4xx status code
func (o *UpdateSLONotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update Slo not found response has a 5xx status code
func (o *UpdateSLONotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update Slo not found response a status code equal to that given
func (o *UpdateSLONotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update Slo not found response
func (o *UpdateSLONotFound) Code() int {
	return 404
}

func (o *UpdateSLONotFound) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/slos/{slug}][%d] updateSloNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSLONotFound) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/slos/{slug}][%d] updateSloNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSLONotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateSLONotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSLOConflict creates a UpdateSLOConflict with default headers values
func NewUpdateSLOConflict() *UpdateSLOConflict {
	return &UpdateSLOConflict{}
}

/*
UpdateSLOConflict describes a response with status code 409, with default header values.

Cannot update the SLO because there is a conflict with an existing SLO.
*/
type UpdateSLOConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update Slo conflict response has a 2xx status code
func (o *UpdateSLOConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Slo conflict response has a 3xx status code
func (o *UpdateSLOConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Slo conflict response has a 4xx status code
func (o *UpdateSLOConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update Slo conflict response has a 5xx status code
func (o *UpdateSLOConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update Slo conflict response a status code equal to that given
func (o *UpdateSLOConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update Slo conflict response
func (o *UpdateSLOConflict) Code() int {
	return 409
}

func (o *UpdateSLOConflict) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/slos/{slug}][%d] updateSloConflict  %+v", 409, o.Payload)
}

func (o *UpdateSLOConflict) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/slos/{slug}][%d] updateSloConflict  %+v", 409, o.Payload)
}

func (o *UpdateSLOConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateSLOConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSLOInternalServerError creates a UpdateSLOInternalServerError with default headers values
func NewUpdateSLOInternalServerError() *UpdateSLOInternalServerError {
	return &UpdateSLOInternalServerError{}
}

/*
UpdateSLOInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateSLOInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update Slo internal server error response has a 2xx status code
func (o *UpdateSLOInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Slo internal server error response has a 3xx status code
func (o *UpdateSLOInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Slo internal server error response has a 4xx status code
func (o *UpdateSLOInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update Slo internal server error response has a 5xx status code
func (o *UpdateSLOInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update Slo internal server error response a status code equal to that given
func (o *UpdateSLOInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update Slo internal server error response
func (o *UpdateSLOInternalServerError) Code() int {
	return 500
}

func (o *UpdateSLOInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/slos/{slug}][%d] updateSloInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSLOInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/slos/{slug}][%d] updateSloInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSLOInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateSLOInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSLODefault creates a UpdateSLODefault with default headers values
func NewUpdateSLODefault(code int) *UpdateSLODefault {
	return &UpdateSLODefault{
		_statusCode: code,
	}
}

/*
UpdateSLODefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateSLODefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update SLO default response has a 2xx status code
func (o *UpdateSLODefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update SLO default response has a 3xx status code
func (o *UpdateSLODefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update SLO default response has a 4xx status code
func (o *UpdateSLODefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update SLO default response has a 5xx status code
func (o *UpdateSLODefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update SLO default response a status code equal to that given
func (o *UpdateSLODefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update SLO default response
func (o *UpdateSLODefault) Code() int {
	return o._statusCode
}

func (o *UpdateSLODefault) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/slos/{slug}][%d] UpdateSLO default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSLODefault) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/slos/{slug}][%d] UpdateSLO default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSLODefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateSLODefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
