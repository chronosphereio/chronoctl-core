// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/stateunstable/models"
)

// ListMetricUsagesByMetricNameReader is a Reader for the ListMetricUsagesByMetricName structure.
type ListMetricUsagesByMetricNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMetricUsagesByMetricNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMetricUsagesByMetricNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListMetricUsagesByMetricNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListMetricUsagesByMetricNameOK creates a ListMetricUsagesByMetricNameOK with default headers values
func NewListMetricUsagesByMetricNameOK() *ListMetricUsagesByMetricNameOK {
	return &ListMetricUsagesByMetricNameOK{}
}

/*
ListMetricUsagesByMetricNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListMetricUsagesByMetricNameOK struct {
	Payload *models.StateunstableListMetricUsagesByMetricNameResponse
}

// IsSuccess returns true when this list metric usages by metric name o k response has a 2xx status code
func (o *ListMetricUsagesByMetricNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list metric usages by metric name o k response has a 3xx status code
func (o *ListMetricUsagesByMetricNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list metric usages by metric name o k response has a 4xx status code
func (o *ListMetricUsagesByMetricNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list metric usages by metric name o k response has a 5xx status code
func (o *ListMetricUsagesByMetricNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list metric usages by metric name o k response a status code equal to that given
func (o *ListMetricUsagesByMetricNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list metric usages by metric name o k response
func (o *ListMetricUsagesByMetricNameOK) Code() int {
	return 200
}

func (o *ListMetricUsagesByMetricNameOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/state/metric-usages-by-metric-name][%d] listMetricUsagesByMetricNameOK  %+v", 200, o.Payload)
}

func (o *ListMetricUsagesByMetricNameOK) String() string {
	return fmt.Sprintf("[GET /api/v1/state/metric-usages-by-metric-name][%d] listMetricUsagesByMetricNameOK  %+v", 200, o.Payload)
}

func (o *ListMetricUsagesByMetricNameOK) GetPayload() *models.StateunstableListMetricUsagesByMetricNameResponse {
	return o.Payload
}

func (o *ListMetricUsagesByMetricNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StateunstableListMetricUsagesByMetricNameResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMetricUsagesByMetricNameDefault creates a ListMetricUsagesByMetricNameDefault with default headers values
func NewListMetricUsagesByMetricNameDefault(code int) *ListMetricUsagesByMetricNameDefault {
	return &ListMetricUsagesByMetricNameDefault{
		_statusCode: code,
	}
}

/*
ListMetricUsagesByMetricNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListMetricUsagesByMetricNameDefault struct {
	_statusCode int

	Payload *models.APIError
}

// IsSuccess returns true when this list metric usages by metric name default response has a 2xx status code
func (o *ListMetricUsagesByMetricNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list metric usages by metric name default response has a 3xx status code
func (o *ListMetricUsagesByMetricNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list metric usages by metric name default response has a 4xx status code
func (o *ListMetricUsagesByMetricNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list metric usages by metric name default response has a 5xx status code
func (o *ListMetricUsagesByMetricNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list metric usages by metric name default response a status code equal to that given
func (o *ListMetricUsagesByMetricNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list metric usages by metric name default response
func (o *ListMetricUsagesByMetricNameDefault) Code() int {
	return o._statusCode
}

func (o *ListMetricUsagesByMetricNameDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/state/metric-usages-by-metric-name][%d] ListMetricUsagesByMetricName default  %+v", o._statusCode, o.Payload)
}

func (o *ListMetricUsagesByMetricNameDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/state/metric-usages-by-metric-name][%d] ListMetricUsagesByMetricName default  %+v", o._statusCode, o.Payload)
}

func (o *ListMetricUsagesByMetricNameDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListMetricUsagesByMetricNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
