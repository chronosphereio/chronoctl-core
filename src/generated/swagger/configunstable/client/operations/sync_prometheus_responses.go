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

// SyncPrometheusReader is a Reader for the SyncPrometheus structure.
type SyncPrometheusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncPrometheusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncPrometheusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSyncPrometheusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSyncPrometheusOK creates a SyncPrometheusOK with default headers values
func NewSyncPrometheusOK() *SyncPrometheusOK {
	return &SyncPrometheusOK{}
}

/*
SyncPrometheusOK describes a response with status code 200, with default header values.

A successful response.
*/
type SyncPrometheusOK struct {
	Payload *models.ConfigunstableSyncPrometheusResponse
}

// IsSuccess returns true when this sync prometheus o k response has a 2xx status code
func (o *SyncPrometheusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sync prometheus o k response has a 3xx status code
func (o *SyncPrometheusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync prometheus o k response has a 4xx status code
func (o *SyncPrometheusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sync prometheus o k response has a 5xx status code
func (o *SyncPrometheusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sync prometheus o k response a status code equal to that given
func (o *SyncPrometheusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sync prometheus o k response
func (o *SyncPrometheusOK) Code() int {
	return 200
}

func (o *SyncPrometheusOK) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/sync-prometheus][%d] syncPrometheusOK  %+v", 200, o.Payload)
}

func (o *SyncPrometheusOK) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/sync-prometheus][%d] syncPrometheusOK  %+v", 200, o.Payload)
}

func (o *SyncPrometheusOK) GetPayload() *models.ConfigunstableSyncPrometheusResponse {
	return o.Payload
}

func (o *SyncPrometheusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableSyncPrometheusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncPrometheusDefault creates a SyncPrometheusDefault with default headers values
func NewSyncPrometheusDefault(code int) *SyncPrometheusDefault {
	return &SyncPrometheusDefault{
		_statusCode: code,
	}
}

/*
SyncPrometheusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SyncPrometheusDefault struct {
	_statusCode int

	Payload *models.APIError
}

// IsSuccess returns true when this sync prometheus default response has a 2xx status code
func (o *SyncPrometheusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this sync prometheus default response has a 3xx status code
func (o *SyncPrometheusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this sync prometheus default response has a 4xx status code
func (o *SyncPrometheusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this sync prometheus default response has a 5xx status code
func (o *SyncPrometheusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this sync prometheus default response a status code equal to that given
func (o *SyncPrometheusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the sync prometheus default response
func (o *SyncPrometheusDefault) Code() int {
	return o._statusCode
}

func (o *SyncPrometheusDefault) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/sync-prometheus][%d] SyncPrometheus default  %+v", o._statusCode, o.Payload)
}

func (o *SyncPrometheusDefault) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/sync-prometheus][%d] SyncPrometheus default  %+v", o._statusCode, o.Payload)
}

func (o *SyncPrometheusDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *SyncPrometheusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
