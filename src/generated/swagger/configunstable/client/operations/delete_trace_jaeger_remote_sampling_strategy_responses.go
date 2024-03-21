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

// DeleteTraceJaegerRemoteSamplingStrategyReader is a Reader for the DeleteTraceJaegerRemoteSamplingStrategy structure.
type DeleteTraceJaegerRemoteSamplingStrategyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTraceJaegerRemoteSamplingStrategyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTraceJaegerRemoteSamplingStrategyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTraceJaegerRemoteSamplingStrategyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTraceJaegerRemoteSamplingStrategyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTraceJaegerRemoteSamplingStrategyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTraceJaegerRemoteSamplingStrategyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTraceJaegerRemoteSamplingStrategyOK creates a DeleteTraceJaegerRemoteSamplingStrategyOK with default headers values
func NewDeleteTraceJaegerRemoteSamplingStrategyOK() *DeleteTraceJaegerRemoteSamplingStrategyOK {
	return &DeleteTraceJaegerRemoteSamplingStrategyOK{}
}

/*
DeleteTraceJaegerRemoteSamplingStrategyOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteTraceJaegerRemoteSamplingStrategyOK struct {
	Payload models.ConfigunstableDeleteTraceJaegerRemoteSamplingStrategyResponse
}

// IsSuccess returns true when this delete trace jaeger remote sampling strategy o k response has a 2xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete trace jaeger remote sampling strategy o k response has a 3xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace jaeger remote sampling strategy o k response has a 4xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete trace jaeger remote sampling strategy o k response has a 5xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete trace jaeger remote sampling strategy o k response a status code equal to that given
func (o *DeleteTraceJaegerRemoteSamplingStrategyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete trace jaeger remote sampling strategy o k response
func (o *DeleteTraceJaegerRemoteSamplingStrategyOK) Code() int {
	return 200
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyOK) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] deleteTraceJaegerRemoteSamplingStrategyOK  %+v", 200, o.Payload)
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyOK) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] deleteTraceJaegerRemoteSamplingStrategyOK  %+v", 200, o.Payload)
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyOK) GetPayload() models.ConfigunstableDeleteTraceJaegerRemoteSamplingStrategyResponse {
	return o.Payload
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceJaegerRemoteSamplingStrategyBadRequest creates a DeleteTraceJaegerRemoteSamplingStrategyBadRequest with default headers values
func NewDeleteTraceJaegerRemoteSamplingStrategyBadRequest() *DeleteTraceJaegerRemoteSamplingStrategyBadRequest {
	return &DeleteTraceJaegerRemoteSamplingStrategyBadRequest{}
}

/*
DeleteTraceJaegerRemoteSamplingStrategyBadRequest describes a response with status code 400, with default header values.

Cannot delete the TraceJaegerRemoteSamplingStrategy because it is in use.
*/
type DeleteTraceJaegerRemoteSamplingStrategyBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete trace jaeger remote sampling strategy bad request response has a 2xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete trace jaeger remote sampling strategy bad request response has a 3xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace jaeger remote sampling strategy bad request response has a 4xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete trace jaeger remote sampling strategy bad request response has a 5xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete trace jaeger remote sampling strategy bad request response a status code equal to that given
func (o *DeleteTraceJaegerRemoteSamplingStrategyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete trace jaeger remote sampling strategy bad request response
func (o *DeleteTraceJaegerRemoteSamplingStrategyBadRequest) Code() int {
	return 400
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] deleteTraceJaegerRemoteSamplingStrategyBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] deleteTraceJaegerRemoteSamplingStrategyBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceJaegerRemoteSamplingStrategyNotFound creates a DeleteTraceJaegerRemoteSamplingStrategyNotFound with default headers values
func NewDeleteTraceJaegerRemoteSamplingStrategyNotFound() *DeleteTraceJaegerRemoteSamplingStrategyNotFound {
	return &DeleteTraceJaegerRemoteSamplingStrategyNotFound{}
}

/*
DeleteTraceJaegerRemoteSamplingStrategyNotFound describes a response with status code 404, with default header values.

Cannot delete the TraceJaegerRemoteSamplingStrategy because the slug does not exist.
*/
type DeleteTraceJaegerRemoteSamplingStrategyNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete trace jaeger remote sampling strategy not found response has a 2xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete trace jaeger remote sampling strategy not found response has a 3xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace jaeger remote sampling strategy not found response has a 4xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete trace jaeger remote sampling strategy not found response has a 5xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete trace jaeger remote sampling strategy not found response a status code equal to that given
func (o *DeleteTraceJaegerRemoteSamplingStrategyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete trace jaeger remote sampling strategy not found response
func (o *DeleteTraceJaegerRemoteSamplingStrategyNotFound) Code() int {
	return 404
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] deleteTraceJaegerRemoteSamplingStrategyNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] deleteTraceJaegerRemoteSamplingStrategyNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceJaegerRemoteSamplingStrategyInternalServerError creates a DeleteTraceJaegerRemoteSamplingStrategyInternalServerError with default headers values
func NewDeleteTraceJaegerRemoteSamplingStrategyInternalServerError() *DeleteTraceJaegerRemoteSamplingStrategyInternalServerError {
	return &DeleteTraceJaegerRemoteSamplingStrategyInternalServerError{}
}

/*
DeleteTraceJaegerRemoteSamplingStrategyInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteTraceJaegerRemoteSamplingStrategyInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete trace jaeger remote sampling strategy internal server error response has a 2xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete trace jaeger remote sampling strategy internal server error response has a 3xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace jaeger remote sampling strategy internal server error response has a 4xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete trace jaeger remote sampling strategy internal server error response has a 5xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete trace jaeger remote sampling strategy internal server error response a status code equal to that given
func (o *DeleteTraceJaegerRemoteSamplingStrategyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete trace jaeger remote sampling strategy internal server error response
func (o *DeleteTraceJaegerRemoteSamplingStrategyInternalServerError) Code() int {
	return 500
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] deleteTraceJaegerRemoteSamplingStrategyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] deleteTraceJaegerRemoteSamplingStrategyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceJaegerRemoteSamplingStrategyDefault creates a DeleteTraceJaegerRemoteSamplingStrategyDefault with default headers values
func NewDeleteTraceJaegerRemoteSamplingStrategyDefault(code int) *DeleteTraceJaegerRemoteSamplingStrategyDefault {
	return &DeleteTraceJaegerRemoteSamplingStrategyDefault{
		_statusCode: code,
	}
}

/*
DeleteTraceJaegerRemoteSamplingStrategyDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteTraceJaegerRemoteSamplingStrategyDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete trace jaeger remote sampling strategy default response has a 2xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete trace jaeger remote sampling strategy default response has a 3xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete trace jaeger remote sampling strategy default response has a 4xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete trace jaeger remote sampling strategy default response has a 5xx status code
func (o *DeleteTraceJaegerRemoteSamplingStrategyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete trace jaeger remote sampling strategy default response a status code equal to that given
func (o *DeleteTraceJaegerRemoteSamplingStrategyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete trace jaeger remote sampling strategy default response
func (o *DeleteTraceJaegerRemoteSamplingStrategyDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] DeleteTraceJaegerRemoteSamplingStrategy default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyDefault) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] DeleteTraceJaegerRemoteSamplingStrategy default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteTraceJaegerRemoteSamplingStrategyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
