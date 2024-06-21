// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
)

// UpdateDashboardReader is a Reader for the UpdateDashboard structure.
type UpdateDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateDashboardConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDashboardOK creates a UpdateDashboardOK with default headers values
func NewUpdateDashboardOK() *UpdateDashboardOK {
	return &UpdateDashboardOK{}
}

/*
UpdateDashboardOK describes a response with status code 200, with default header values.

A successful response containing the updated Dashboard.
*/
type UpdateDashboardOK struct {
	Payload *models.Configv1UpdateDashboardResponse
}

// IsSuccess returns true when this update dashboard o k response has a 2xx status code
func (o *UpdateDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update dashboard o k response has a 3xx status code
func (o *UpdateDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard o k response has a 4xx status code
func (o *UpdateDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update dashboard o k response has a 5xx status code
func (o *UpdateDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update dashboard o k response a status code equal to that given
func (o *UpdateDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update dashboard o k response
func (o *UpdateDashboardOK) Code() int {
	return 200
}

func (o *UpdateDashboardOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/dashboards/{slug}][%d] updateDashboardOK  %+v", 200, o.Payload)
}

func (o *UpdateDashboardOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/dashboards/{slug}][%d] updateDashboardOK  %+v", 200, o.Payload)
}

func (o *UpdateDashboardOK) GetPayload() *models.Configv1UpdateDashboardResponse {
	return o.Payload
}

func (o *UpdateDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1UpdateDashboardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardBadRequest creates a UpdateDashboardBadRequest with default headers values
func NewUpdateDashboardBadRequest() *UpdateDashboardBadRequest {
	return &UpdateDashboardBadRequest{}
}

/*
UpdateDashboardBadRequest describes a response with status code 400, with default header values.

Cannot update the Dashboard because the request is invalid.
*/
type UpdateDashboardBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update dashboard bad request response has a 2xx status code
func (o *UpdateDashboardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dashboard bad request response has a 3xx status code
func (o *UpdateDashboardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard bad request response has a 4xx status code
func (o *UpdateDashboardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update dashboard bad request response has a 5xx status code
func (o *UpdateDashboardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update dashboard bad request response a status code equal to that given
func (o *UpdateDashboardBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update dashboard bad request response
func (o *UpdateDashboardBadRequest) Code() int {
	return 400
}

func (o *UpdateDashboardBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/dashboards/{slug}][%d] updateDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDashboardBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/dashboards/{slug}][%d] updateDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDashboardBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardNotFound creates a UpdateDashboardNotFound with default headers values
func NewUpdateDashboardNotFound() *UpdateDashboardNotFound {
	return &UpdateDashboardNotFound{}
}

/*
UpdateDashboardNotFound describes a response with status code 404, with default header values.

Cannot update the Dashboard because the slug does not exist.
*/
type UpdateDashboardNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update dashboard not found response has a 2xx status code
func (o *UpdateDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dashboard not found response has a 3xx status code
func (o *UpdateDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard not found response has a 4xx status code
func (o *UpdateDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update dashboard not found response has a 5xx status code
func (o *UpdateDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update dashboard not found response a status code equal to that given
func (o *UpdateDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update dashboard not found response
func (o *UpdateDashboardNotFound) Code() int {
	return 404
}

func (o *UpdateDashboardNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/dashboards/{slug}][%d] updateDashboardNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDashboardNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/dashboards/{slug}][%d] updateDashboardNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDashboardNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardConflict creates a UpdateDashboardConflict with default headers values
func NewUpdateDashboardConflict() *UpdateDashboardConflict {
	return &UpdateDashboardConflict{}
}

/*
UpdateDashboardConflict describes a response with status code 409, with default header values.

Cannot update the Dashboard because there is a conflict with an existing Dashboard.
*/
type UpdateDashboardConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update dashboard conflict response has a 2xx status code
func (o *UpdateDashboardConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dashboard conflict response has a 3xx status code
func (o *UpdateDashboardConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard conflict response has a 4xx status code
func (o *UpdateDashboardConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update dashboard conflict response has a 5xx status code
func (o *UpdateDashboardConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update dashboard conflict response a status code equal to that given
func (o *UpdateDashboardConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update dashboard conflict response
func (o *UpdateDashboardConflict) Code() int {
	return 409
}

func (o *UpdateDashboardConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/dashboards/{slug}][%d] updateDashboardConflict  %+v", 409, o.Payload)
}

func (o *UpdateDashboardConflict) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/dashboards/{slug}][%d] updateDashboardConflict  %+v", 409, o.Payload)
}

func (o *UpdateDashboardConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateDashboardConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardInternalServerError creates a UpdateDashboardInternalServerError with default headers values
func NewUpdateDashboardInternalServerError() *UpdateDashboardInternalServerError {
	return &UpdateDashboardInternalServerError{}
}

/*
UpdateDashboardInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateDashboardInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update dashboard internal server error response has a 2xx status code
func (o *UpdateDashboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dashboard internal server error response has a 3xx status code
func (o *UpdateDashboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard internal server error response has a 4xx status code
func (o *UpdateDashboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update dashboard internal server error response has a 5xx status code
func (o *UpdateDashboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update dashboard internal server error response a status code equal to that given
func (o *UpdateDashboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update dashboard internal server error response
func (o *UpdateDashboardInternalServerError) Code() int {
	return 500
}

func (o *UpdateDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/dashboards/{slug}][%d] updateDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDashboardInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/dashboards/{slug}][%d] updateDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDashboardInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardDefault creates a UpdateDashboardDefault with default headers values
func NewUpdateDashboardDefault(code int) *UpdateDashboardDefault {
	return &UpdateDashboardDefault{
		_statusCode: code,
	}
}

/*
UpdateDashboardDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateDashboardDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update dashboard default response has a 2xx status code
func (o *UpdateDashboardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update dashboard default response has a 3xx status code
func (o *UpdateDashboardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update dashboard default response has a 4xx status code
func (o *UpdateDashboardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update dashboard default response has a 5xx status code
func (o *UpdateDashboardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update dashboard default response a status code equal to that given
func (o *UpdateDashboardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update dashboard default response
func (o *UpdateDashboardDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDashboardDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/dashboards/{slug}][%d] UpdateDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDashboardDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/dashboards/{slug}][%d] UpdateDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDashboardDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
