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

// UpdateTraceMetricsRuleReader is a Reader for the UpdateTraceMetricsRule structure.
type UpdateTraceMetricsRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTraceMetricsRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTraceMetricsRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTraceMetricsRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTraceMetricsRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateTraceMetricsRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTraceMetricsRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateTraceMetricsRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTraceMetricsRuleOK creates a UpdateTraceMetricsRuleOK with default headers values
func NewUpdateTraceMetricsRuleOK() *UpdateTraceMetricsRuleOK {
	return &UpdateTraceMetricsRuleOK{}
}

/*
UpdateTraceMetricsRuleOK describes a response with status code 200, with default header values.

A successful response containing the updated TraceMetricsRule.
*/
type UpdateTraceMetricsRuleOK struct {
	Payload *models.Configv1UpdateTraceMetricsRuleResponse
}

// IsSuccess returns true when this update trace metrics rule o k response has a 2xx status code
func (o *UpdateTraceMetricsRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update trace metrics rule o k response has a 3xx status code
func (o *UpdateTraceMetricsRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace metrics rule o k response has a 4xx status code
func (o *UpdateTraceMetricsRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trace metrics rule o k response has a 5xx status code
func (o *UpdateTraceMetricsRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace metrics rule o k response a status code equal to that given
func (o *UpdateTraceMetricsRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update trace metrics rule o k response
func (o *UpdateTraceMetricsRuleOK) Code() int {
	return 200
}

func (o *UpdateTraceMetricsRuleOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/trace-metrics-rules/{slug}][%d] updateTraceMetricsRuleOK  %+v", 200, o.Payload)
}

func (o *UpdateTraceMetricsRuleOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/trace-metrics-rules/{slug}][%d] updateTraceMetricsRuleOK  %+v", 200, o.Payload)
}

func (o *UpdateTraceMetricsRuleOK) GetPayload() *models.Configv1UpdateTraceMetricsRuleResponse {
	return o.Payload
}

func (o *UpdateTraceMetricsRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1UpdateTraceMetricsRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceMetricsRuleBadRequest creates a UpdateTraceMetricsRuleBadRequest with default headers values
func NewUpdateTraceMetricsRuleBadRequest() *UpdateTraceMetricsRuleBadRequest {
	return &UpdateTraceMetricsRuleBadRequest{}
}

/*
UpdateTraceMetricsRuleBadRequest describes a response with status code 400, with default header values.

Cannot update the TraceMetricsRule because the request is invalid.
*/
type UpdateTraceMetricsRuleBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace metrics rule bad request response has a 2xx status code
func (o *UpdateTraceMetricsRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace metrics rule bad request response has a 3xx status code
func (o *UpdateTraceMetricsRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace metrics rule bad request response has a 4xx status code
func (o *UpdateTraceMetricsRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trace metrics rule bad request response has a 5xx status code
func (o *UpdateTraceMetricsRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace metrics rule bad request response a status code equal to that given
func (o *UpdateTraceMetricsRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update trace metrics rule bad request response
func (o *UpdateTraceMetricsRuleBadRequest) Code() int {
	return 400
}

func (o *UpdateTraceMetricsRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/trace-metrics-rules/{slug}][%d] updateTraceMetricsRuleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTraceMetricsRuleBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/trace-metrics-rules/{slug}][%d] updateTraceMetricsRuleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTraceMetricsRuleBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceMetricsRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceMetricsRuleNotFound creates a UpdateTraceMetricsRuleNotFound with default headers values
func NewUpdateTraceMetricsRuleNotFound() *UpdateTraceMetricsRuleNotFound {
	return &UpdateTraceMetricsRuleNotFound{}
}

/*
UpdateTraceMetricsRuleNotFound describes a response with status code 404, with default header values.

Cannot update the TraceMetricsRule because the slug does not exist.
*/
type UpdateTraceMetricsRuleNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace metrics rule not found response has a 2xx status code
func (o *UpdateTraceMetricsRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace metrics rule not found response has a 3xx status code
func (o *UpdateTraceMetricsRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace metrics rule not found response has a 4xx status code
func (o *UpdateTraceMetricsRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trace metrics rule not found response has a 5xx status code
func (o *UpdateTraceMetricsRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace metrics rule not found response a status code equal to that given
func (o *UpdateTraceMetricsRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update trace metrics rule not found response
func (o *UpdateTraceMetricsRuleNotFound) Code() int {
	return 404
}

func (o *UpdateTraceMetricsRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/trace-metrics-rules/{slug}][%d] updateTraceMetricsRuleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTraceMetricsRuleNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/trace-metrics-rules/{slug}][%d] updateTraceMetricsRuleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTraceMetricsRuleNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceMetricsRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceMetricsRuleConflict creates a UpdateTraceMetricsRuleConflict with default headers values
func NewUpdateTraceMetricsRuleConflict() *UpdateTraceMetricsRuleConflict {
	return &UpdateTraceMetricsRuleConflict{}
}

/*
UpdateTraceMetricsRuleConflict describes a response with status code 409, with default header values.

Cannot update the TraceMetricsRule because there is a conflict with an existing TraceMetricsRule.
*/
type UpdateTraceMetricsRuleConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace metrics rule conflict response has a 2xx status code
func (o *UpdateTraceMetricsRuleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace metrics rule conflict response has a 3xx status code
func (o *UpdateTraceMetricsRuleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace metrics rule conflict response has a 4xx status code
func (o *UpdateTraceMetricsRuleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trace metrics rule conflict response has a 5xx status code
func (o *UpdateTraceMetricsRuleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace metrics rule conflict response a status code equal to that given
func (o *UpdateTraceMetricsRuleConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update trace metrics rule conflict response
func (o *UpdateTraceMetricsRuleConflict) Code() int {
	return 409
}

func (o *UpdateTraceMetricsRuleConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/trace-metrics-rules/{slug}][%d] updateTraceMetricsRuleConflict  %+v", 409, o.Payload)
}

func (o *UpdateTraceMetricsRuleConflict) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/trace-metrics-rules/{slug}][%d] updateTraceMetricsRuleConflict  %+v", 409, o.Payload)
}

func (o *UpdateTraceMetricsRuleConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceMetricsRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceMetricsRuleInternalServerError creates a UpdateTraceMetricsRuleInternalServerError with default headers values
func NewUpdateTraceMetricsRuleInternalServerError() *UpdateTraceMetricsRuleInternalServerError {
	return &UpdateTraceMetricsRuleInternalServerError{}
}

/*
UpdateTraceMetricsRuleInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateTraceMetricsRuleInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace metrics rule internal server error response has a 2xx status code
func (o *UpdateTraceMetricsRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace metrics rule internal server error response has a 3xx status code
func (o *UpdateTraceMetricsRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace metrics rule internal server error response has a 4xx status code
func (o *UpdateTraceMetricsRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trace metrics rule internal server error response has a 5xx status code
func (o *UpdateTraceMetricsRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update trace metrics rule internal server error response a status code equal to that given
func (o *UpdateTraceMetricsRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update trace metrics rule internal server error response
func (o *UpdateTraceMetricsRuleInternalServerError) Code() int {
	return 500
}

func (o *UpdateTraceMetricsRuleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/trace-metrics-rules/{slug}][%d] updateTraceMetricsRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTraceMetricsRuleInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/trace-metrics-rules/{slug}][%d] updateTraceMetricsRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTraceMetricsRuleInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceMetricsRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceMetricsRuleDefault creates a UpdateTraceMetricsRuleDefault with default headers values
func NewUpdateTraceMetricsRuleDefault(code int) *UpdateTraceMetricsRuleDefault {
	return &UpdateTraceMetricsRuleDefault{
		_statusCode: code,
	}
}

/*
UpdateTraceMetricsRuleDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateTraceMetricsRuleDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update trace metrics rule default response has a 2xx status code
func (o *UpdateTraceMetricsRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update trace metrics rule default response has a 3xx status code
func (o *UpdateTraceMetricsRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update trace metrics rule default response has a 4xx status code
func (o *UpdateTraceMetricsRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update trace metrics rule default response has a 5xx status code
func (o *UpdateTraceMetricsRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update trace metrics rule default response a status code equal to that given
func (o *UpdateTraceMetricsRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update trace metrics rule default response
func (o *UpdateTraceMetricsRuleDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTraceMetricsRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/trace-metrics-rules/{slug}][%d] UpdateTraceMetricsRule default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTraceMetricsRuleDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/trace-metrics-rules/{slug}][%d] UpdateTraceMetricsRule default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTraceMetricsRuleDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateTraceMetricsRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
