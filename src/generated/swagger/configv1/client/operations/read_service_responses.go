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

// ReadServiceReader is a Reader for the ReadService structure.
type ReadServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadServiceOK creates a ReadServiceOK with default headers values
func NewReadServiceOK() *ReadServiceOK {
	return &ReadServiceOK{}
}

/*
ReadServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadServiceOK struct {
	Payload *models.Configv1ReadServiceResponse
}

// IsSuccess returns true when this read service o k response has a 2xx status code
func (o *ReadServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read service o k response has a 3xx status code
func (o *ReadServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read service o k response has a 4xx status code
func (o *ReadServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read service o k response has a 5xx status code
func (o *ReadServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read service o k response a status code equal to that given
func (o *ReadServiceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read service o k response
func (o *ReadServiceOK) Code() int {
	return 200
}

func (o *ReadServiceOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/services/{slug}][%d] readServiceOK  %+v", 200, o.Payload)
}

func (o *ReadServiceOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/services/{slug}][%d] readServiceOK  %+v", 200, o.Payload)
}

func (o *ReadServiceOK) GetPayload() *models.Configv1ReadServiceResponse {
	return o.Payload
}

func (o *ReadServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ReadServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadServiceNotFound creates a ReadServiceNotFound with default headers values
func NewReadServiceNotFound() *ReadServiceNotFound {
	return &ReadServiceNotFound{}
}

/*
ReadServiceNotFound describes a response with status code 404, with default header values.

Cannot read the Service because the slug does not exist.
*/
type ReadServiceNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read service not found response has a 2xx status code
func (o *ReadServiceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read service not found response has a 3xx status code
func (o *ReadServiceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read service not found response has a 4xx status code
func (o *ReadServiceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read service not found response has a 5xx status code
func (o *ReadServiceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read service not found response a status code equal to that given
func (o *ReadServiceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read service not found response
func (o *ReadServiceNotFound) Code() int {
	return 404
}

func (o *ReadServiceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/services/{slug}][%d] readServiceNotFound  %+v", 404, o.Payload)
}

func (o *ReadServiceNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/config/services/{slug}][%d] readServiceNotFound  %+v", 404, o.Payload)
}

func (o *ReadServiceNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadServiceInternalServerError creates a ReadServiceInternalServerError with default headers values
func NewReadServiceInternalServerError() *ReadServiceInternalServerError {
	return &ReadServiceInternalServerError{}
}

/*
ReadServiceInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadServiceInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read service internal server error response has a 2xx status code
func (o *ReadServiceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read service internal server error response has a 3xx status code
func (o *ReadServiceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read service internal server error response has a 4xx status code
func (o *ReadServiceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read service internal server error response has a 5xx status code
func (o *ReadServiceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read service internal server error response a status code equal to that given
func (o *ReadServiceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read service internal server error response
func (o *ReadServiceInternalServerError) Code() int {
	return 500
}

func (o *ReadServiceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/services/{slug}][%d] readServiceInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadServiceInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/services/{slug}][%d] readServiceInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadServiceInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadServiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadServiceDefault creates a ReadServiceDefault with default headers values
func NewReadServiceDefault(code int) *ReadServiceDefault {
	return &ReadServiceDefault{
		_statusCode: code,
	}
}

/*
ReadServiceDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadServiceDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read service default response has a 2xx status code
func (o *ReadServiceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read service default response has a 3xx status code
func (o *ReadServiceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read service default response has a 4xx status code
func (o *ReadServiceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read service default response has a 5xx status code
func (o *ReadServiceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read service default response a status code equal to that given
func (o *ReadServiceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read service default response
func (o *ReadServiceDefault) Code() int {
	return o._statusCode
}

func (o *ReadServiceDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/services/{slug}][%d] ReadService default  %+v", o._statusCode, o.Payload)
}

func (o *ReadServiceDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/services/{slug}][%d] ReadService default  %+v", o._statusCode, o.Payload)
}

func (o *ReadServiceDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
