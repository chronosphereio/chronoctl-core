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

// UpdateTeamReader is a Reader for the UpdateTeam structure.
type UpdateTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTeamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateTeamConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTeamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTeamOK creates a UpdateTeamOK with default headers values
func NewUpdateTeamOK() *UpdateTeamOK {
	return &UpdateTeamOK{}
}

/*
UpdateTeamOK describes a response with status code 200, with default header values.

A successful response containing the updated Team.
*/
type UpdateTeamOK struct {
	Payload *models.Configv1UpdateTeamResponse
}

// IsSuccess returns true when this update team o k response has a 2xx status code
func (o *UpdateTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update team o k response has a 3xx status code
func (o *UpdateTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team o k response has a 4xx status code
func (o *UpdateTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update team o k response has a 5xx status code
func (o *UpdateTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update team o k response a status code equal to that given
func (o *UpdateTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update team o k response
func (o *UpdateTeamOK) Code() int {
	return 200
}

func (o *UpdateTeamOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/teams/{slug}][%d] updateTeamOK  %+v", 200, o.Payload)
}

func (o *UpdateTeamOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/teams/{slug}][%d] updateTeamOK  %+v", 200, o.Payload)
}

func (o *UpdateTeamOK) GetPayload() *models.Configv1UpdateTeamResponse {
	return o.Payload
}

func (o *UpdateTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1UpdateTeamResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamBadRequest creates a UpdateTeamBadRequest with default headers values
func NewUpdateTeamBadRequest() *UpdateTeamBadRequest {
	return &UpdateTeamBadRequest{}
}

/*
UpdateTeamBadRequest describes a response with status code 400, with default header values.

Cannot update the Team because the request is invalid.
*/
type UpdateTeamBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update team bad request response has a 2xx status code
func (o *UpdateTeamBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update team bad request response has a 3xx status code
func (o *UpdateTeamBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team bad request response has a 4xx status code
func (o *UpdateTeamBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update team bad request response has a 5xx status code
func (o *UpdateTeamBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update team bad request response a status code equal to that given
func (o *UpdateTeamBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update team bad request response
func (o *UpdateTeamBadRequest) Code() int {
	return 400
}

func (o *UpdateTeamBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/teams/{slug}][%d] updateTeamBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTeamBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/teams/{slug}][%d] updateTeamBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTeamBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTeamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamNotFound creates a UpdateTeamNotFound with default headers values
func NewUpdateTeamNotFound() *UpdateTeamNotFound {
	return &UpdateTeamNotFound{}
}

/*
UpdateTeamNotFound describes a response with status code 404, with default header values.

Cannot update the Team because the slug does not exist.
*/
type UpdateTeamNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update team not found response has a 2xx status code
func (o *UpdateTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update team not found response has a 3xx status code
func (o *UpdateTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team not found response has a 4xx status code
func (o *UpdateTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update team not found response has a 5xx status code
func (o *UpdateTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update team not found response a status code equal to that given
func (o *UpdateTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update team not found response
func (o *UpdateTeamNotFound) Code() int {
	return 404
}

func (o *UpdateTeamNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/teams/{slug}][%d] updateTeamNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTeamNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/teams/{slug}][%d] updateTeamNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTeamNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamConflict creates a UpdateTeamConflict with default headers values
func NewUpdateTeamConflict() *UpdateTeamConflict {
	return &UpdateTeamConflict{}
}

/*
UpdateTeamConflict describes a response with status code 409, with default header values.

Cannot update the Team because there is a conflict with an existing Team.
*/
type UpdateTeamConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update team conflict response has a 2xx status code
func (o *UpdateTeamConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update team conflict response has a 3xx status code
func (o *UpdateTeamConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team conflict response has a 4xx status code
func (o *UpdateTeamConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update team conflict response has a 5xx status code
func (o *UpdateTeamConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update team conflict response a status code equal to that given
func (o *UpdateTeamConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update team conflict response
func (o *UpdateTeamConflict) Code() int {
	return 409
}

func (o *UpdateTeamConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/teams/{slug}][%d] updateTeamConflict  %+v", 409, o.Payload)
}

func (o *UpdateTeamConflict) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/teams/{slug}][%d] updateTeamConflict  %+v", 409, o.Payload)
}

func (o *UpdateTeamConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTeamConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamInternalServerError creates a UpdateTeamInternalServerError with default headers values
func NewUpdateTeamInternalServerError() *UpdateTeamInternalServerError {
	return &UpdateTeamInternalServerError{}
}

/*
UpdateTeamInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateTeamInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update team internal server error response has a 2xx status code
func (o *UpdateTeamInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update team internal server error response has a 3xx status code
func (o *UpdateTeamInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team internal server error response has a 4xx status code
func (o *UpdateTeamInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update team internal server error response has a 5xx status code
func (o *UpdateTeamInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update team internal server error response a status code equal to that given
func (o *UpdateTeamInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update team internal server error response
func (o *UpdateTeamInternalServerError) Code() int {
	return 500
}

func (o *UpdateTeamInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/teams/{slug}][%d] updateTeamInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTeamInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/teams/{slug}][%d] updateTeamInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTeamInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTeamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamDefault creates a UpdateTeamDefault with default headers values
func NewUpdateTeamDefault(code int) *UpdateTeamDefault {
	return &UpdateTeamDefault{
		_statusCode: code,
	}
}

/*
UpdateTeamDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateTeamDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update team default response has a 2xx status code
func (o *UpdateTeamDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update team default response has a 3xx status code
func (o *UpdateTeamDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update team default response has a 4xx status code
func (o *UpdateTeamDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update team default response has a 5xx status code
func (o *UpdateTeamDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update team default response a status code equal to that given
func (o *UpdateTeamDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update team default response
func (o *UpdateTeamDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTeamDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/teams/{slug}][%d] UpdateTeam default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTeamDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/teams/{slug}][%d] UpdateTeam default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTeamDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
