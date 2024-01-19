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

// CreateTraceJaegerRemoteSamplingStrategyReader is a Reader for the CreateTraceJaegerRemoteSamplingStrategy structure.
type CreateTraceJaegerRemoteSamplingStrategyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTraceJaegerRemoteSamplingStrategyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTraceJaegerRemoteSamplingStrategyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTraceJaegerRemoteSamplingStrategyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateTraceJaegerRemoteSamplingStrategyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTraceJaegerRemoteSamplingStrategyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateTraceJaegerRemoteSamplingStrategyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTraceJaegerRemoteSamplingStrategyOK creates a CreateTraceJaegerRemoteSamplingStrategyOK with default headers values
func NewCreateTraceJaegerRemoteSamplingStrategyOK() *CreateTraceJaegerRemoteSamplingStrategyOK {
	return &CreateTraceJaegerRemoteSamplingStrategyOK{}
}

/*
CreateTraceJaegerRemoteSamplingStrategyOK describes a response with status code 200, with default header values.

A successful response containing the created TraceJaegerRemoteSamplingStrategy.
*/
type CreateTraceJaegerRemoteSamplingStrategyOK struct {
	Payload *models.Configv1CreateTraceJaegerRemoteSamplingStrategyResponse
}

// IsSuccess returns true when this create trace jaeger remote sampling strategy o k response has a 2xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create trace jaeger remote sampling strategy o k response has a 3xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create trace jaeger remote sampling strategy o k response has a 4xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create trace jaeger remote sampling strategy o k response has a 5xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create trace jaeger remote sampling strategy o k response a status code equal to that given
func (o *CreateTraceJaegerRemoteSamplingStrategyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create trace jaeger remote sampling strategy o k response
func (o *CreateTraceJaegerRemoteSamplingStrategyOK) Code() int {
	return 200
}

func (o *CreateTraceJaegerRemoteSamplingStrategyOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-jaeger-remote-sampling-strategies][%d] createTraceJaegerRemoteSamplingStrategyOK  %+v", 200, o.Payload)
}

func (o *CreateTraceJaegerRemoteSamplingStrategyOK) String() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-jaeger-remote-sampling-strategies][%d] createTraceJaegerRemoteSamplingStrategyOK  %+v", 200, o.Payload)
}

func (o *CreateTraceJaegerRemoteSamplingStrategyOK) GetPayload() *models.Configv1CreateTraceJaegerRemoteSamplingStrategyResponse {
	return o.Payload
}

func (o *CreateTraceJaegerRemoteSamplingStrategyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1CreateTraceJaegerRemoteSamplingStrategyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTraceJaegerRemoteSamplingStrategyBadRequest creates a CreateTraceJaegerRemoteSamplingStrategyBadRequest with default headers values
func NewCreateTraceJaegerRemoteSamplingStrategyBadRequest() *CreateTraceJaegerRemoteSamplingStrategyBadRequest {
	return &CreateTraceJaegerRemoteSamplingStrategyBadRequest{}
}

/*
CreateTraceJaegerRemoteSamplingStrategyBadRequest describes a response with status code 400, with default header values.

Cannot create the TraceJaegerRemoteSamplingStrategy because the request is invalid.
*/
type CreateTraceJaegerRemoteSamplingStrategyBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create trace jaeger remote sampling strategy bad request response has a 2xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create trace jaeger remote sampling strategy bad request response has a 3xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create trace jaeger remote sampling strategy bad request response has a 4xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create trace jaeger remote sampling strategy bad request response has a 5xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create trace jaeger remote sampling strategy bad request response a status code equal to that given
func (o *CreateTraceJaegerRemoteSamplingStrategyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create trace jaeger remote sampling strategy bad request response
func (o *CreateTraceJaegerRemoteSamplingStrategyBadRequest) Code() int {
	return 400
}

func (o *CreateTraceJaegerRemoteSamplingStrategyBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-jaeger-remote-sampling-strategies][%d] createTraceJaegerRemoteSamplingStrategyBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTraceJaegerRemoteSamplingStrategyBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-jaeger-remote-sampling-strategies][%d] createTraceJaegerRemoteSamplingStrategyBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTraceJaegerRemoteSamplingStrategyBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateTraceJaegerRemoteSamplingStrategyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTraceJaegerRemoteSamplingStrategyConflict creates a CreateTraceJaegerRemoteSamplingStrategyConflict with default headers values
func NewCreateTraceJaegerRemoteSamplingStrategyConflict() *CreateTraceJaegerRemoteSamplingStrategyConflict {
	return &CreateTraceJaegerRemoteSamplingStrategyConflict{}
}

/*
CreateTraceJaegerRemoteSamplingStrategyConflict describes a response with status code 409, with default header values.

Cannot create the TraceJaegerRemoteSamplingStrategy because there is a conflict with an existing TraceJaegerRemoteSamplingStrategy.
*/
type CreateTraceJaegerRemoteSamplingStrategyConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create trace jaeger remote sampling strategy conflict response has a 2xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create trace jaeger remote sampling strategy conflict response has a 3xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create trace jaeger remote sampling strategy conflict response has a 4xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create trace jaeger remote sampling strategy conflict response has a 5xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create trace jaeger remote sampling strategy conflict response a status code equal to that given
func (o *CreateTraceJaegerRemoteSamplingStrategyConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create trace jaeger remote sampling strategy conflict response
func (o *CreateTraceJaegerRemoteSamplingStrategyConflict) Code() int {
	return 409
}

func (o *CreateTraceJaegerRemoteSamplingStrategyConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-jaeger-remote-sampling-strategies][%d] createTraceJaegerRemoteSamplingStrategyConflict  %+v", 409, o.Payload)
}

func (o *CreateTraceJaegerRemoteSamplingStrategyConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-jaeger-remote-sampling-strategies][%d] createTraceJaegerRemoteSamplingStrategyConflict  %+v", 409, o.Payload)
}

func (o *CreateTraceJaegerRemoteSamplingStrategyConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateTraceJaegerRemoteSamplingStrategyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTraceJaegerRemoteSamplingStrategyInternalServerError creates a CreateTraceJaegerRemoteSamplingStrategyInternalServerError with default headers values
func NewCreateTraceJaegerRemoteSamplingStrategyInternalServerError() *CreateTraceJaegerRemoteSamplingStrategyInternalServerError {
	return &CreateTraceJaegerRemoteSamplingStrategyInternalServerError{}
}

/*
CreateTraceJaegerRemoteSamplingStrategyInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateTraceJaegerRemoteSamplingStrategyInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create trace jaeger remote sampling strategy internal server error response has a 2xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create trace jaeger remote sampling strategy internal server error response has a 3xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create trace jaeger remote sampling strategy internal server error response has a 4xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create trace jaeger remote sampling strategy internal server error response has a 5xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create trace jaeger remote sampling strategy internal server error response a status code equal to that given
func (o *CreateTraceJaegerRemoteSamplingStrategyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create trace jaeger remote sampling strategy internal server error response
func (o *CreateTraceJaegerRemoteSamplingStrategyInternalServerError) Code() int {
	return 500
}

func (o *CreateTraceJaegerRemoteSamplingStrategyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-jaeger-remote-sampling-strategies][%d] createTraceJaegerRemoteSamplingStrategyInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTraceJaegerRemoteSamplingStrategyInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-jaeger-remote-sampling-strategies][%d] createTraceJaegerRemoteSamplingStrategyInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTraceJaegerRemoteSamplingStrategyInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateTraceJaegerRemoteSamplingStrategyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTraceJaegerRemoteSamplingStrategyDefault creates a CreateTraceJaegerRemoteSamplingStrategyDefault with default headers values
func NewCreateTraceJaegerRemoteSamplingStrategyDefault(code int) *CreateTraceJaegerRemoteSamplingStrategyDefault {
	return &CreateTraceJaegerRemoteSamplingStrategyDefault{
		_statusCode: code,
	}
}

/*
CreateTraceJaegerRemoteSamplingStrategyDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateTraceJaegerRemoteSamplingStrategyDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create trace jaeger remote sampling strategy default response has a 2xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create trace jaeger remote sampling strategy default response has a 3xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create trace jaeger remote sampling strategy default response has a 4xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create trace jaeger remote sampling strategy default response has a 5xx status code
func (o *CreateTraceJaegerRemoteSamplingStrategyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create trace jaeger remote sampling strategy default response a status code equal to that given
func (o *CreateTraceJaegerRemoteSamplingStrategyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create trace jaeger remote sampling strategy default response
func (o *CreateTraceJaegerRemoteSamplingStrategyDefault) Code() int {
	return o._statusCode
}

func (o *CreateTraceJaegerRemoteSamplingStrategyDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-jaeger-remote-sampling-strategies][%d] CreateTraceJaegerRemoteSamplingStrategy default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTraceJaegerRemoteSamplingStrategyDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-jaeger-remote-sampling-strategies][%d] CreateTraceJaegerRemoteSamplingStrategy default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTraceJaegerRemoteSamplingStrategyDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateTraceJaegerRemoteSamplingStrategyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
