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

// ReadRecordingRuleReader is a Reader for the ReadRecordingRule structure.
type ReadRecordingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadRecordingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadRecordingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadRecordingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadRecordingRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadRecordingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadRecordingRuleOK creates a ReadRecordingRuleOK with default headers values
func NewReadRecordingRuleOK() *ReadRecordingRuleOK {
	return &ReadRecordingRuleOK{}
}

/*
ReadRecordingRuleOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadRecordingRuleOK struct {
	Payload *models.Configv1ReadRecordingRuleResponse
}

// IsSuccess returns true when this read recording rule o k response has a 2xx status code
func (o *ReadRecordingRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read recording rule o k response has a 3xx status code
func (o *ReadRecordingRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read recording rule o k response has a 4xx status code
func (o *ReadRecordingRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read recording rule o k response has a 5xx status code
func (o *ReadRecordingRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read recording rule o k response a status code equal to that given
func (o *ReadRecordingRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read recording rule o k response
func (o *ReadRecordingRuleOK) Code() int {
	return 200
}

func (o *ReadRecordingRuleOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/recording-rules/{slug}][%d] readRecordingRuleOK  %+v", 200, o.Payload)
}

func (o *ReadRecordingRuleOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/recording-rules/{slug}][%d] readRecordingRuleOK  %+v", 200, o.Payload)
}

func (o *ReadRecordingRuleOK) GetPayload() *models.Configv1ReadRecordingRuleResponse {
	return o.Payload
}

func (o *ReadRecordingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ReadRecordingRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadRecordingRuleNotFound creates a ReadRecordingRuleNotFound with default headers values
func NewReadRecordingRuleNotFound() *ReadRecordingRuleNotFound {
	return &ReadRecordingRuleNotFound{}
}

/*
ReadRecordingRuleNotFound describes a response with status code 404, with default header values.

Cannot read the RecordingRule because the slug does not exist.
*/
type ReadRecordingRuleNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read recording rule not found response has a 2xx status code
func (o *ReadRecordingRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read recording rule not found response has a 3xx status code
func (o *ReadRecordingRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read recording rule not found response has a 4xx status code
func (o *ReadRecordingRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read recording rule not found response has a 5xx status code
func (o *ReadRecordingRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read recording rule not found response a status code equal to that given
func (o *ReadRecordingRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read recording rule not found response
func (o *ReadRecordingRuleNotFound) Code() int {
	return 404
}

func (o *ReadRecordingRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/recording-rules/{slug}][%d] readRecordingRuleNotFound  %+v", 404, o.Payload)
}

func (o *ReadRecordingRuleNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/config/recording-rules/{slug}][%d] readRecordingRuleNotFound  %+v", 404, o.Payload)
}

func (o *ReadRecordingRuleNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadRecordingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadRecordingRuleInternalServerError creates a ReadRecordingRuleInternalServerError with default headers values
func NewReadRecordingRuleInternalServerError() *ReadRecordingRuleInternalServerError {
	return &ReadRecordingRuleInternalServerError{}
}

/*
ReadRecordingRuleInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadRecordingRuleInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read recording rule internal server error response has a 2xx status code
func (o *ReadRecordingRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read recording rule internal server error response has a 3xx status code
func (o *ReadRecordingRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read recording rule internal server error response has a 4xx status code
func (o *ReadRecordingRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read recording rule internal server error response has a 5xx status code
func (o *ReadRecordingRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read recording rule internal server error response a status code equal to that given
func (o *ReadRecordingRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read recording rule internal server error response
func (o *ReadRecordingRuleInternalServerError) Code() int {
	return 500
}

func (o *ReadRecordingRuleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/recording-rules/{slug}][%d] readRecordingRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadRecordingRuleInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/recording-rules/{slug}][%d] readRecordingRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadRecordingRuleInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadRecordingRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadRecordingRuleDefault creates a ReadRecordingRuleDefault with default headers values
func NewReadRecordingRuleDefault(code int) *ReadRecordingRuleDefault {
	return &ReadRecordingRuleDefault{
		_statusCode: code,
	}
}

/*
ReadRecordingRuleDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadRecordingRuleDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read recording rule default response has a 2xx status code
func (o *ReadRecordingRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read recording rule default response has a 3xx status code
func (o *ReadRecordingRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read recording rule default response has a 4xx status code
func (o *ReadRecordingRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read recording rule default response has a 5xx status code
func (o *ReadRecordingRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read recording rule default response a status code equal to that given
func (o *ReadRecordingRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read recording rule default response
func (o *ReadRecordingRuleDefault) Code() int {
	return o._statusCode
}

func (o *ReadRecordingRuleDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/recording-rules/{slug}][%d] ReadRecordingRule default  %+v", o._statusCode, o.Payload)
}

func (o *ReadRecordingRuleDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/recording-rules/{slug}][%d] ReadRecordingRule default  %+v", o._statusCode, o.Payload)
}

func (o *ReadRecordingRuleDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadRecordingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
