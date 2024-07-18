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

// CreateOtelMetricsIngestionReader is a Reader for the CreateOtelMetricsIngestion structure.
type CreateOtelMetricsIngestionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOtelMetricsIngestionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOtelMetricsIngestionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOtelMetricsIngestionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateOtelMetricsIngestionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateOtelMetricsIngestionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateOtelMetricsIngestionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOtelMetricsIngestionOK creates a CreateOtelMetricsIngestionOK with default headers values
func NewCreateOtelMetricsIngestionOK() *CreateOtelMetricsIngestionOK {
	return &CreateOtelMetricsIngestionOK{}
}

/*
CreateOtelMetricsIngestionOK describes a response with status code 200, with default header values.

A successful response containing the created OtelMetricsIngestion.
*/
type CreateOtelMetricsIngestionOK struct {
	Payload *models.Configv1CreateOtelMetricsIngestionResponse
}

// IsSuccess returns true when this create otel metrics ingestion o k response has a 2xx status code
func (o *CreateOtelMetricsIngestionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create otel metrics ingestion o k response has a 3xx status code
func (o *CreateOtelMetricsIngestionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create otel metrics ingestion o k response has a 4xx status code
func (o *CreateOtelMetricsIngestionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create otel metrics ingestion o k response has a 5xx status code
func (o *CreateOtelMetricsIngestionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create otel metrics ingestion o k response a status code equal to that given
func (o *CreateOtelMetricsIngestionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create otel metrics ingestion o k response
func (o *CreateOtelMetricsIngestionOK) Code() int {
	return 200
}

func (o *CreateOtelMetricsIngestionOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/otel-metrics-ingestion][%d] createOtelMetricsIngestionOK  %+v", 200, o.Payload)
}

func (o *CreateOtelMetricsIngestionOK) String() string {
	return fmt.Sprintf("[POST /api/v1/config/otel-metrics-ingestion][%d] createOtelMetricsIngestionOK  %+v", 200, o.Payload)
}

func (o *CreateOtelMetricsIngestionOK) GetPayload() *models.Configv1CreateOtelMetricsIngestionResponse {
	return o.Payload
}

func (o *CreateOtelMetricsIngestionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1CreateOtelMetricsIngestionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOtelMetricsIngestionBadRequest creates a CreateOtelMetricsIngestionBadRequest with default headers values
func NewCreateOtelMetricsIngestionBadRequest() *CreateOtelMetricsIngestionBadRequest {
	return &CreateOtelMetricsIngestionBadRequest{}
}

/*
CreateOtelMetricsIngestionBadRequest describes a response with status code 400, with default header values.

Cannot create the OtelMetricsIngestion because the request is invalid.
*/
type CreateOtelMetricsIngestionBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create otel metrics ingestion bad request response has a 2xx status code
func (o *CreateOtelMetricsIngestionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create otel metrics ingestion bad request response has a 3xx status code
func (o *CreateOtelMetricsIngestionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create otel metrics ingestion bad request response has a 4xx status code
func (o *CreateOtelMetricsIngestionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create otel metrics ingestion bad request response has a 5xx status code
func (o *CreateOtelMetricsIngestionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create otel metrics ingestion bad request response a status code equal to that given
func (o *CreateOtelMetricsIngestionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create otel metrics ingestion bad request response
func (o *CreateOtelMetricsIngestionBadRequest) Code() int {
	return 400
}

func (o *CreateOtelMetricsIngestionBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/otel-metrics-ingestion][%d] createOtelMetricsIngestionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOtelMetricsIngestionBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/config/otel-metrics-ingestion][%d] createOtelMetricsIngestionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOtelMetricsIngestionBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateOtelMetricsIngestionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOtelMetricsIngestionConflict creates a CreateOtelMetricsIngestionConflict with default headers values
func NewCreateOtelMetricsIngestionConflict() *CreateOtelMetricsIngestionConflict {
	return &CreateOtelMetricsIngestionConflict{}
}

/*
CreateOtelMetricsIngestionConflict describes a response with status code 409, with default header values.

Cannot create the OtelMetricsIngestion because there is a conflict with an existing OtelMetricsIngestion.
*/
type CreateOtelMetricsIngestionConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create otel metrics ingestion conflict response has a 2xx status code
func (o *CreateOtelMetricsIngestionConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create otel metrics ingestion conflict response has a 3xx status code
func (o *CreateOtelMetricsIngestionConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create otel metrics ingestion conflict response has a 4xx status code
func (o *CreateOtelMetricsIngestionConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create otel metrics ingestion conflict response has a 5xx status code
func (o *CreateOtelMetricsIngestionConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create otel metrics ingestion conflict response a status code equal to that given
func (o *CreateOtelMetricsIngestionConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create otel metrics ingestion conflict response
func (o *CreateOtelMetricsIngestionConflict) Code() int {
	return 409
}

func (o *CreateOtelMetricsIngestionConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/otel-metrics-ingestion][%d] createOtelMetricsIngestionConflict  %+v", 409, o.Payload)
}

func (o *CreateOtelMetricsIngestionConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/config/otel-metrics-ingestion][%d] createOtelMetricsIngestionConflict  %+v", 409, o.Payload)
}

func (o *CreateOtelMetricsIngestionConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateOtelMetricsIngestionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOtelMetricsIngestionInternalServerError creates a CreateOtelMetricsIngestionInternalServerError with default headers values
func NewCreateOtelMetricsIngestionInternalServerError() *CreateOtelMetricsIngestionInternalServerError {
	return &CreateOtelMetricsIngestionInternalServerError{}
}

/*
CreateOtelMetricsIngestionInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateOtelMetricsIngestionInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create otel metrics ingestion internal server error response has a 2xx status code
func (o *CreateOtelMetricsIngestionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create otel metrics ingestion internal server error response has a 3xx status code
func (o *CreateOtelMetricsIngestionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create otel metrics ingestion internal server error response has a 4xx status code
func (o *CreateOtelMetricsIngestionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create otel metrics ingestion internal server error response has a 5xx status code
func (o *CreateOtelMetricsIngestionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create otel metrics ingestion internal server error response a status code equal to that given
func (o *CreateOtelMetricsIngestionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create otel metrics ingestion internal server error response
func (o *CreateOtelMetricsIngestionInternalServerError) Code() int {
	return 500
}

func (o *CreateOtelMetricsIngestionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/otel-metrics-ingestion][%d] createOtelMetricsIngestionInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateOtelMetricsIngestionInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/config/otel-metrics-ingestion][%d] createOtelMetricsIngestionInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateOtelMetricsIngestionInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateOtelMetricsIngestionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOtelMetricsIngestionDefault creates a CreateOtelMetricsIngestionDefault with default headers values
func NewCreateOtelMetricsIngestionDefault(code int) *CreateOtelMetricsIngestionDefault {
	return &CreateOtelMetricsIngestionDefault{
		_statusCode: code,
	}
}

/*
CreateOtelMetricsIngestionDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateOtelMetricsIngestionDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create otel metrics ingestion default response has a 2xx status code
func (o *CreateOtelMetricsIngestionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create otel metrics ingestion default response has a 3xx status code
func (o *CreateOtelMetricsIngestionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create otel metrics ingestion default response has a 4xx status code
func (o *CreateOtelMetricsIngestionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create otel metrics ingestion default response has a 5xx status code
func (o *CreateOtelMetricsIngestionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create otel metrics ingestion default response a status code equal to that given
func (o *CreateOtelMetricsIngestionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create otel metrics ingestion default response
func (o *CreateOtelMetricsIngestionDefault) Code() int {
	return o._statusCode
}

func (o *CreateOtelMetricsIngestionDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/otel-metrics-ingestion][%d] CreateOtelMetricsIngestion default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOtelMetricsIngestionDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/config/otel-metrics-ingestion][%d] CreateOtelMetricsIngestion default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOtelMetricsIngestionDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateOtelMetricsIngestionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}