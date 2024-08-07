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

// ReadTraceTopTagConfigReader is a Reader for the ReadTraceTopTagConfig structure.
type ReadTraceTopTagConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadTraceTopTagConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadTraceTopTagConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadTraceTopTagConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadTraceTopTagConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadTraceTopTagConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadTraceTopTagConfigOK creates a ReadTraceTopTagConfigOK with default headers values
func NewReadTraceTopTagConfigOK() *ReadTraceTopTagConfigOK {
	return &ReadTraceTopTagConfigOK{}
}

/*
ReadTraceTopTagConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadTraceTopTagConfigOK struct {
	Payload *models.ConfigunstableReadTraceTopTagConfigResponse
}

// IsSuccess returns true when this read trace top tag config o k response has a 2xx status code
func (o *ReadTraceTopTagConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read trace top tag config o k response has a 3xx status code
func (o *ReadTraceTopTagConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read trace top tag config o k response has a 4xx status code
func (o *ReadTraceTopTagConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read trace top tag config o k response has a 5xx status code
func (o *ReadTraceTopTagConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read trace top tag config o k response a status code equal to that given
func (o *ReadTraceTopTagConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read trace top tag config o k response
func (o *ReadTraceTopTagConfigOK) Code() int {
	return 200
}

func (o *ReadTraceTopTagConfigOK) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-top-tag-config][%d] readTraceTopTagConfigOK  %+v", 200, o.Payload)
}

func (o *ReadTraceTopTagConfigOK) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-top-tag-config][%d] readTraceTopTagConfigOK  %+v", 200, o.Payload)
}

func (o *ReadTraceTopTagConfigOK) GetPayload() *models.ConfigunstableReadTraceTopTagConfigResponse {
	return o.Payload
}

func (o *ReadTraceTopTagConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableReadTraceTopTagConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTraceTopTagConfigNotFound creates a ReadTraceTopTagConfigNotFound with default headers values
func NewReadTraceTopTagConfigNotFound() *ReadTraceTopTagConfigNotFound {
	return &ReadTraceTopTagConfigNotFound{}
}

/*
ReadTraceTopTagConfigNotFound describes a response with status code 404, with default header values.

Cannot read the TraceTopTagConfig because TraceTopTagConfig has not been created.
*/
type ReadTraceTopTagConfigNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read trace top tag config not found response has a 2xx status code
func (o *ReadTraceTopTagConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read trace top tag config not found response has a 3xx status code
func (o *ReadTraceTopTagConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read trace top tag config not found response has a 4xx status code
func (o *ReadTraceTopTagConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read trace top tag config not found response has a 5xx status code
func (o *ReadTraceTopTagConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read trace top tag config not found response a status code equal to that given
func (o *ReadTraceTopTagConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read trace top tag config not found response
func (o *ReadTraceTopTagConfigNotFound) Code() int {
	return 404
}

func (o *ReadTraceTopTagConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-top-tag-config][%d] readTraceTopTagConfigNotFound  %+v", 404, o.Payload)
}

func (o *ReadTraceTopTagConfigNotFound) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-top-tag-config][%d] readTraceTopTagConfigNotFound  %+v", 404, o.Payload)
}

func (o *ReadTraceTopTagConfigNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadTraceTopTagConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTraceTopTagConfigInternalServerError creates a ReadTraceTopTagConfigInternalServerError with default headers values
func NewReadTraceTopTagConfigInternalServerError() *ReadTraceTopTagConfigInternalServerError {
	return &ReadTraceTopTagConfigInternalServerError{}
}

/*
ReadTraceTopTagConfigInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadTraceTopTagConfigInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read trace top tag config internal server error response has a 2xx status code
func (o *ReadTraceTopTagConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read trace top tag config internal server error response has a 3xx status code
func (o *ReadTraceTopTagConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read trace top tag config internal server error response has a 4xx status code
func (o *ReadTraceTopTagConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read trace top tag config internal server error response has a 5xx status code
func (o *ReadTraceTopTagConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read trace top tag config internal server error response a status code equal to that given
func (o *ReadTraceTopTagConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read trace top tag config internal server error response
func (o *ReadTraceTopTagConfigInternalServerError) Code() int {
	return 500
}

func (o *ReadTraceTopTagConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-top-tag-config][%d] readTraceTopTagConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadTraceTopTagConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-top-tag-config][%d] readTraceTopTagConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadTraceTopTagConfigInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadTraceTopTagConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTraceTopTagConfigDefault creates a ReadTraceTopTagConfigDefault with default headers values
func NewReadTraceTopTagConfigDefault(code int) *ReadTraceTopTagConfigDefault {
	return &ReadTraceTopTagConfigDefault{
		_statusCode: code,
	}
}

/*
ReadTraceTopTagConfigDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadTraceTopTagConfigDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read trace top tag config default response has a 2xx status code
func (o *ReadTraceTopTagConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read trace top tag config default response has a 3xx status code
func (o *ReadTraceTopTagConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read trace top tag config default response has a 4xx status code
func (o *ReadTraceTopTagConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read trace top tag config default response has a 5xx status code
func (o *ReadTraceTopTagConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read trace top tag config default response a status code equal to that given
func (o *ReadTraceTopTagConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read trace top tag config default response
func (o *ReadTraceTopTagConfigDefault) Code() int {
	return o._statusCode
}

func (o *ReadTraceTopTagConfigDefault) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-top-tag-config][%d] ReadTraceTopTagConfig default  %+v", o._statusCode, o.Payload)
}

func (o *ReadTraceTopTagConfigDefault) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-top-tag-config][%d] ReadTraceTopTagConfig default  %+v", o._statusCode, o.Payload)
}

func (o *ReadTraceTopTagConfigDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadTraceTopTagConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
