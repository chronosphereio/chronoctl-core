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

// ReadServiceAccountReader is a Reader for the ReadServiceAccount structure.
type ReadServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadServiceAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadServiceAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadServiceAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadServiceAccountOK creates a ReadServiceAccountOK with default headers values
func NewReadServiceAccountOK() *ReadServiceAccountOK {
	return &ReadServiceAccountOK{}
}

/*
ReadServiceAccountOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadServiceAccountOK struct {
	Payload *models.Configv1ReadServiceAccountResponse
}

// IsSuccess returns true when this read service account o k response has a 2xx status code
func (o *ReadServiceAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read service account o k response has a 3xx status code
func (o *ReadServiceAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read service account o k response has a 4xx status code
func (o *ReadServiceAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read service account o k response has a 5xx status code
func (o *ReadServiceAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read service account o k response a status code equal to that given
func (o *ReadServiceAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read service account o k response
func (o *ReadServiceAccountOK) Code() int {
	return 200
}

func (o *ReadServiceAccountOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/service-accounts/{slug}][%d] readServiceAccountOK  %+v", 200, o.Payload)
}

func (o *ReadServiceAccountOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/service-accounts/{slug}][%d] readServiceAccountOK  %+v", 200, o.Payload)
}

func (o *ReadServiceAccountOK) GetPayload() *models.Configv1ReadServiceAccountResponse {
	return o.Payload
}

func (o *ReadServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ReadServiceAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadServiceAccountNotFound creates a ReadServiceAccountNotFound with default headers values
func NewReadServiceAccountNotFound() *ReadServiceAccountNotFound {
	return &ReadServiceAccountNotFound{}
}

/*
ReadServiceAccountNotFound describes a response with status code 404, with default header values.

Cannot read the ServiceAccount because the slug does not exist.
*/
type ReadServiceAccountNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read service account not found response has a 2xx status code
func (o *ReadServiceAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read service account not found response has a 3xx status code
func (o *ReadServiceAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read service account not found response has a 4xx status code
func (o *ReadServiceAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read service account not found response has a 5xx status code
func (o *ReadServiceAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read service account not found response a status code equal to that given
func (o *ReadServiceAccountNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read service account not found response
func (o *ReadServiceAccountNotFound) Code() int {
	return 404
}

func (o *ReadServiceAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/service-accounts/{slug}][%d] readServiceAccountNotFound  %+v", 404, o.Payload)
}

func (o *ReadServiceAccountNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/config/service-accounts/{slug}][%d] readServiceAccountNotFound  %+v", 404, o.Payload)
}

func (o *ReadServiceAccountNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadServiceAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadServiceAccountInternalServerError creates a ReadServiceAccountInternalServerError with default headers values
func NewReadServiceAccountInternalServerError() *ReadServiceAccountInternalServerError {
	return &ReadServiceAccountInternalServerError{}
}

/*
ReadServiceAccountInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadServiceAccountInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read service account internal server error response has a 2xx status code
func (o *ReadServiceAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read service account internal server error response has a 3xx status code
func (o *ReadServiceAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read service account internal server error response has a 4xx status code
func (o *ReadServiceAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read service account internal server error response has a 5xx status code
func (o *ReadServiceAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read service account internal server error response a status code equal to that given
func (o *ReadServiceAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read service account internal server error response
func (o *ReadServiceAccountInternalServerError) Code() int {
	return 500
}

func (o *ReadServiceAccountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/service-accounts/{slug}][%d] readServiceAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadServiceAccountInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/service-accounts/{slug}][%d] readServiceAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadServiceAccountInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadServiceAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadServiceAccountDefault creates a ReadServiceAccountDefault with default headers values
func NewReadServiceAccountDefault(code int) *ReadServiceAccountDefault {
	return &ReadServiceAccountDefault{
		_statusCode: code,
	}
}

/*
ReadServiceAccountDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadServiceAccountDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read service account default response has a 2xx status code
func (o *ReadServiceAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read service account default response has a 3xx status code
func (o *ReadServiceAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read service account default response has a 4xx status code
func (o *ReadServiceAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read service account default response has a 5xx status code
func (o *ReadServiceAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read service account default response a status code equal to that given
func (o *ReadServiceAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read service account default response
func (o *ReadServiceAccountDefault) Code() int {
	return o._statusCode
}

func (o *ReadServiceAccountDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/service-accounts/{slug}][%d] ReadServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *ReadServiceAccountDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/service-accounts/{slug}][%d] ReadServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *ReadServiceAccountDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadServiceAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}