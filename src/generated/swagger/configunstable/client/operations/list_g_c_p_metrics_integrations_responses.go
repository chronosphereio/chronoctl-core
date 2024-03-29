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

// ListGCPMetricsIntegrationsReader is a Reader for the ListGCPMetricsIntegrations structure.
type ListGCPMetricsIntegrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPMetricsIntegrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGCPMetricsIntegrationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListGCPMetricsIntegrationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListGCPMetricsIntegrationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPMetricsIntegrationsOK creates a ListGCPMetricsIntegrationsOK with default headers values
func NewListGCPMetricsIntegrationsOK() *ListGCPMetricsIntegrationsOK {
	return &ListGCPMetricsIntegrationsOK{}
}

/*
ListGCPMetricsIntegrationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListGCPMetricsIntegrationsOK struct {
	Payload *models.ConfigunstableListGCPMetricsIntegrationsResponse
}

// IsSuccess returns true when this list g c p metrics integrations o k response has a 2xx status code
func (o *ListGCPMetricsIntegrationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list g c p metrics integrations o k response has a 3xx status code
func (o *ListGCPMetricsIntegrationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g c p metrics integrations o k response has a 4xx status code
func (o *ListGCPMetricsIntegrationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list g c p metrics integrations o k response has a 5xx status code
func (o *ListGCPMetricsIntegrationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list g c p metrics integrations o k response a status code equal to that given
func (o *ListGCPMetricsIntegrationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list g c p metrics integrations o k response
func (o *ListGCPMetricsIntegrationsOK) Code() int {
	return 200
}

func (o *ListGCPMetricsIntegrationsOK) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/gcp-metrics-integrations][%d] listGCPMetricsIntegrationsOK  %+v", 200, o.Payload)
}

func (o *ListGCPMetricsIntegrationsOK) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/gcp-metrics-integrations][%d] listGCPMetricsIntegrationsOK  %+v", 200, o.Payload)
}

func (o *ListGCPMetricsIntegrationsOK) GetPayload() *models.ConfigunstableListGCPMetricsIntegrationsResponse {
	return o.Payload
}

func (o *ListGCPMetricsIntegrationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableListGCPMetricsIntegrationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPMetricsIntegrationsInternalServerError creates a ListGCPMetricsIntegrationsInternalServerError with default headers values
func NewListGCPMetricsIntegrationsInternalServerError() *ListGCPMetricsIntegrationsInternalServerError {
	return &ListGCPMetricsIntegrationsInternalServerError{}
}

/*
ListGCPMetricsIntegrationsInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ListGCPMetricsIntegrationsInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this list g c p metrics integrations internal server error response has a 2xx status code
func (o *ListGCPMetricsIntegrationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list g c p metrics integrations internal server error response has a 3xx status code
func (o *ListGCPMetricsIntegrationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g c p metrics integrations internal server error response has a 4xx status code
func (o *ListGCPMetricsIntegrationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list g c p metrics integrations internal server error response has a 5xx status code
func (o *ListGCPMetricsIntegrationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list g c p metrics integrations internal server error response a status code equal to that given
func (o *ListGCPMetricsIntegrationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list g c p metrics integrations internal server error response
func (o *ListGCPMetricsIntegrationsInternalServerError) Code() int {
	return 500
}

func (o *ListGCPMetricsIntegrationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/gcp-metrics-integrations][%d] listGCPMetricsIntegrationsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListGCPMetricsIntegrationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/gcp-metrics-integrations][%d] listGCPMetricsIntegrationsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListGCPMetricsIntegrationsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListGCPMetricsIntegrationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPMetricsIntegrationsDefault creates a ListGCPMetricsIntegrationsDefault with default headers values
func NewListGCPMetricsIntegrationsDefault(code int) *ListGCPMetricsIntegrationsDefault {
	return &ListGCPMetricsIntegrationsDefault{
		_statusCode: code,
	}
}

/*
ListGCPMetricsIntegrationsDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ListGCPMetricsIntegrationsDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this list g c p metrics integrations default response has a 2xx status code
func (o *ListGCPMetricsIntegrationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list g c p metrics integrations default response has a 3xx status code
func (o *ListGCPMetricsIntegrationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list g c p metrics integrations default response has a 4xx status code
func (o *ListGCPMetricsIntegrationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list g c p metrics integrations default response has a 5xx status code
func (o *ListGCPMetricsIntegrationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list g c p metrics integrations default response a status code equal to that given
func (o *ListGCPMetricsIntegrationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list g c p metrics integrations default response
func (o *ListGCPMetricsIntegrationsDefault) Code() int {
	return o._statusCode
}

func (o *ListGCPMetricsIntegrationsDefault) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/gcp-metrics-integrations][%d] ListGCPMetricsIntegrations default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPMetricsIntegrationsDefault) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/gcp-metrics-integrations][%d] ListGCPMetricsIntegrations default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPMetricsIntegrationsDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ListGCPMetricsIntegrationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
