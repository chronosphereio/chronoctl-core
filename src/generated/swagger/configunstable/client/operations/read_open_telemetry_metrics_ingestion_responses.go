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

// ReadOpenTelemetryMetricsIngestionReader is a Reader for the ReadOpenTelemetryMetricsIngestion structure.
type ReadOpenTelemetryMetricsIngestionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadOpenTelemetryMetricsIngestionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadOpenTelemetryMetricsIngestionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadOpenTelemetryMetricsIngestionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadOpenTelemetryMetricsIngestionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadOpenTelemetryMetricsIngestionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadOpenTelemetryMetricsIngestionOK creates a ReadOpenTelemetryMetricsIngestionOK with default headers values
func NewReadOpenTelemetryMetricsIngestionOK() *ReadOpenTelemetryMetricsIngestionOK {
	return &ReadOpenTelemetryMetricsIngestionOK{}
}

/*
ReadOpenTelemetryMetricsIngestionOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadOpenTelemetryMetricsIngestionOK struct {
	Payload *models.ConfigunstableReadOpenTelemetryMetricsIngestionResponse
}

// IsSuccess returns true when this read open telemetry metrics ingestion o k response has a 2xx status code
func (o *ReadOpenTelemetryMetricsIngestionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read open telemetry metrics ingestion o k response has a 3xx status code
func (o *ReadOpenTelemetryMetricsIngestionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read open telemetry metrics ingestion o k response has a 4xx status code
func (o *ReadOpenTelemetryMetricsIngestionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read open telemetry metrics ingestion o k response has a 5xx status code
func (o *ReadOpenTelemetryMetricsIngestionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read open telemetry metrics ingestion o k response a status code equal to that given
func (o *ReadOpenTelemetryMetricsIngestionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read open telemetry metrics ingestion o k response
func (o *ReadOpenTelemetryMetricsIngestionOK) Code() int {
	return 200
}

func (o *ReadOpenTelemetryMetricsIngestionOK) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/open-telemetry-metrics-ingestion][%d] readOpenTelemetryMetricsIngestionOK  %+v", 200, o.Payload)
}

func (o *ReadOpenTelemetryMetricsIngestionOK) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/open-telemetry-metrics-ingestion][%d] readOpenTelemetryMetricsIngestionOK  %+v", 200, o.Payload)
}

func (o *ReadOpenTelemetryMetricsIngestionOK) GetPayload() *models.ConfigunstableReadOpenTelemetryMetricsIngestionResponse {
	return o.Payload
}

func (o *ReadOpenTelemetryMetricsIngestionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableReadOpenTelemetryMetricsIngestionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadOpenTelemetryMetricsIngestionNotFound creates a ReadOpenTelemetryMetricsIngestionNotFound with default headers values
func NewReadOpenTelemetryMetricsIngestionNotFound() *ReadOpenTelemetryMetricsIngestionNotFound {
	return &ReadOpenTelemetryMetricsIngestionNotFound{}
}

/*
ReadOpenTelemetryMetricsIngestionNotFound describes a response with status code 404, with default header values.

Cannot read the OpenTelemetryMetricsIngestion because OpenTelemetryMetricsIngestion has not been created.
*/
type ReadOpenTelemetryMetricsIngestionNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read open telemetry metrics ingestion not found response has a 2xx status code
func (o *ReadOpenTelemetryMetricsIngestionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read open telemetry metrics ingestion not found response has a 3xx status code
func (o *ReadOpenTelemetryMetricsIngestionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read open telemetry metrics ingestion not found response has a 4xx status code
func (o *ReadOpenTelemetryMetricsIngestionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read open telemetry metrics ingestion not found response has a 5xx status code
func (o *ReadOpenTelemetryMetricsIngestionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read open telemetry metrics ingestion not found response a status code equal to that given
func (o *ReadOpenTelemetryMetricsIngestionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read open telemetry metrics ingestion not found response
func (o *ReadOpenTelemetryMetricsIngestionNotFound) Code() int {
	return 404
}

func (o *ReadOpenTelemetryMetricsIngestionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/open-telemetry-metrics-ingestion][%d] readOpenTelemetryMetricsIngestionNotFound  %+v", 404, o.Payload)
}

func (o *ReadOpenTelemetryMetricsIngestionNotFound) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/open-telemetry-metrics-ingestion][%d] readOpenTelemetryMetricsIngestionNotFound  %+v", 404, o.Payload)
}

func (o *ReadOpenTelemetryMetricsIngestionNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadOpenTelemetryMetricsIngestionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadOpenTelemetryMetricsIngestionInternalServerError creates a ReadOpenTelemetryMetricsIngestionInternalServerError with default headers values
func NewReadOpenTelemetryMetricsIngestionInternalServerError() *ReadOpenTelemetryMetricsIngestionInternalServerError {
	return &ReadOpenTelemetryMetricsIngestionInternalServerError{}
}

/*
ReadOpenTelemetryMetricsIngestionInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadOpenTelemetryMetricsIngestionInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read open telemetry metrics ingestion internal server error response has a 2xx status code
func (o *ReadOpenTelemetryMetricsIngestionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read open telemetry metrics ingestion internal server error response has a 3xx status code
func (o *ReadOpenTelemetryMetricsIngestionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read open telemetry metrics ingestion internal server error response has a 4xx status code
func (o *ReadOpenTelemetryMetricsIngestionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read open telemetry metrics ingestion internal server error response has a 5xx status code
func (o *ReadOpenTelemetryMetricsIngestionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read open telemetry metrics ingestion internal server error response a status code equal to that given
func (o *ReadOpenTelemetryMetricsIngestionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read open telemetry metrics ingestion internal server error response
func (o *ReadOpenTelemetryMetricsIngestionInternalServerError) Code() int {
	return 500
}

func (o *ReadOpenTelemetryMetricsIngestionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/open-telemetry-metrics-ingestion][%d] readOpenTelemetryMetricsIngestionInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadOpenTelemetryMetricsIngestionInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/open-telemetry-metrics-ingestion][%d] readOpenTelemetryMetricsIngestionInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadOpenTelemetryMetricsIngestionInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadOpenTelemetryMetricsIngestionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadOpenTelemetryMetricsIngestionDefault creates a ReadOpenTelemetryMetricsIngestionDefault with default headers values
func NewReadOpenTelemetryMetricsIngestionDefault(code int) *ReadOpenTelemetryMetricsIngestionDefault {
	return &ReadOpenTelemetryMetricsIngestionDefault{
		_statusCode: code,
	}
}

/*
ReadOpenTelemetryMetricsIngestionDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadOpenTelemetryMetricsIngestionDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read open telemetry metrics ingestion default response has a 2xx status code
func (o *ReadOpenTelemetryMetricsIngestionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read open telemetry metrics ingestion default response has a 3xx status code
func (o *ReadOpenTelemetryMetricsIngestionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read open telemetry metrics ingestion default response has a 4xx status code
func (o *ReadOpenTelemetryMetricsIngestionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read open telemetry metrics ingestion default response has a 5xx status code
func (o *ReadOpenTelemetryMetricsIngestionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read open telemetry metrics ingestion default response a status code equal to that given
func (o *ReadOpenTelemetryMetricsIngestionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read open telemetry metrics ingestion default response
func (o *ReadOpenTelemetryMetricsIngestionDefault) Code() int {
	return o._statusCode
}

func (o *ReadOpenTelemetryMetricsIngestionDefault) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/open-telemetry-metrics-ingestion][%d] ReadOpenTelemetryMetricsIngestion default  %+v", o._statusCode, o.Payload)
}

func (o *ReadOpenTelemetryMetricsIngestionDefault) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/open-telemetry-metrics-ingestion][%d] ReadOpenTelemetryMetricsIngestion default  %+v", o._statusCode, o.Payload)
}

func (o *ReadOpenTelemetryMetricsIngestionDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadOpenTelemetryMetricsIngestionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
