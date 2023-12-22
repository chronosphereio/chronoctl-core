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

// DeleteTraceMetricsRuleReader is a Reader for the DeleteTraceMetricsRule structure.
type DeleteTraceMetricsRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTraceMetricsRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTraceMetricsRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTraceMetricsRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTraceMetricsRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTraceMetricsRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTraceMetricsRuleOK creates a DeleteTraceMetricsRuleOK with default headers values
func NewDeleteTraceMetricsRuleOK() *DeleteTraceMetricsRuleOK {
	return &DeleteTraceMetricsRuleOK{}
}

/*
DeleteTraceMetricsRuleOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteTraceMetricsRuleOK struct {
	Payload models.ConfigunstableDeleteTraceMetricsRuleResponse
}

// IsSuccess returns true when this delete trace metrics rule o k response has a 2xx status code
func (o *DeleteTraceMetricsRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete trace metrics rule o k response has a 3xx status code
func (o *DeleteTraceMetricsRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace metrics rule o k response has a 4xx status code
func (o *DeleteTraceMetricsRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete trace metrics rule o k response has a 5xx status code
func (o *DeleteTraceMetricsRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete trace metrics rule o k response a status code equal to that given
func (o *DeleteTraceMetricsRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete trace metrics rule o k response
func (o *DeleteTraceMetricsRuleOK) Code() int {
	return 200
}

func (o *DeleteTraceMetricsRuleOK) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-metrics-rules/{slug}][%d] deleteTraceMetricsRuleOK  %+v", 200, o.Payload)
}

func (o *DeleteTraceMetricsRuleOK) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-metrics-rules/{slug}][%d] deleteTraceMetricsRuleOK  %+v", 200, o.Payload)
}

func (o *DeleteTraceMetricsRuleOK) GetPayload() models.ConfigunstableDeleteTraceMetricsRuleResponse {
	return o.Payload
}

func (o *DeleteTraceMetricsRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceMetricsRuleNotFound creates a DeleteTraceMetricsRuleNotFound with default headers values
func NewDeleteTraceMetricsRuleNotFound() *DeleteTraceMetricsRuleNotFound {
	return &DeleteTraceMetricsRuleNotFound{}
}

/*
DeleteTraceMetricsRuleNotFound describes a response with status code 404, with default header values.

Cannot delete the TraceMetricsRule because the slug does not exist.
*/
type DeleteTraceMetricsRuleNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete trace metrics rule not found response has a 2xx status code
func (o *DeleteTraceMetricsRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete trace metrics rule not found response has a 3xx status code
func (o *DeleteTraceMetricsRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace metrics rule not found response has a 4xx status code
func (o *DeleteTraceMetricsRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete trace metrics rule not found response has a 5xx status code
func (o *DeleteTraceMetricsRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete trace metrics rule not found response a status code equal to that given
func (o *DeleteTraceMetricsRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete trace metrics rule not found response
func (o *DeleteTraceMetricsRuleNotFound) Code() int {
	return 404
}

func (o *DeleteTraceMetricsRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-metrics-rules/{slug}][%d] deleteTraceMetricsRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTraceMetricsRuleNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-metrics-rules/{slug}][%d] deleteTraceMetricsRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTraceMetricsRuleNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteTraceMetricsRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceMetricsRuleInternalServerError creates a DeleteTraceMetricsRuleInternalServerError with default headers values
func NewDeleteTraceMetricsRuleInternalServerError() *DeleteTraceMetricsRuleInternalServerError {
	return &DeleteTraceMetricsRuleInternalServerError{}
}

/*
DeleteTraceMetricsRuleInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteTraceMetricsRuleInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete trace metrics rule internal server error response has a 2xx status code
func (o *DeleteTraceMetricsRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete trace metrics rule internal server error response has a 3xx status code
func (o *DeleteTraceMetricsRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace metrics rule internal server error response has a 4xx status code
func (o *DeleteTraceMetricsRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete trace metrics rule internal server error response has a 5xx status code
func (o *DeleteTraceMetricsRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete trace metrics rule internal server error response a status code equal to that given
func (o *DeleteTraceMetricsRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete trace metrics rule internal server error response
func (o *DeleteTraceMetricsRuleInternalServerError) Code() int {
	return 500
}

func (o *DeleteTraceMetricsRuleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-metrics-rules/{slug}][%d] deleteTraceMetricsRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTraceMetricsRuleInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-metrics-rules/{slug}][%d] deleteTraceMetricsRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTraceMetricsRuleInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteTraceMetricsRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceMetricsRuleDefault creates a DeleteTraceMetricsRuleDefault with default headers values
func NewDeleteTraceMetricsRuleDefault(code int) *DeleteTraceMetricsRuleDefault {
	return &DeleteTraceMetricsRuleDefault{
		_statusCode: code,
	}
}

/*
DeleteTraceMetricsRuleDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteTraceMetricsRuleDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete trace metrics rule default response has a 2xx status code
func (o *DeleteTraceMetricsRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete trace metrics rule default response has a 3xx status code
func (o *DeleteTraceMetricsRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete trace metrics rule default response has a 4xx status code
func (o *DeleteTraceMetricsRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete trace metrics rule default response has a 5xx status code
func (o *DeleteTraceMetricsRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete trace metrics rule default response a status code equal to that given
func (o *DeleteTraceMetricsRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete trace metrics rule default response
func (o *DeleteTraceMetricsRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTraceMetricsRuleDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-metrics-rules/{slug}][%d] DeleteTraceMetricsRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTraceMetricsRuleDefault) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-metrics-rules/{slug}][%d] DeleteTraceMetricsRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTraceMetricsRuleDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteTraceMetricsRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
