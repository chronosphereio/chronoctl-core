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

// ReadDerivedLabelReader is a Reader for the ReadDerivedLabel structure.
type ReadDerivedLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadDerivedLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadDerivedLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadDerivedLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadDerivedLabelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadDerivedLabelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadDerivedLabelOK creates a ReadDerivedLabelOK with default headers values
func NewReadDerivedLabelOK() *ReadDerivedLabelOK {
	return &ReadDerivedLabelOK{}
}

/*
ReadDerivedLabelOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadDerivedLabelOK struct {
	Payload *models.Configv1ReadDerivedLabelResponse
}

// IsSuccess returns true when this read derived label o k response has a 2xx status code
func (o *ReadDerivedLabelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read derived label o k response has a 3xx status code
func (o *ReadDerivedLabelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read derived label o k response has a 4xx status code
func (o *ReadDerivedLabelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read derived label o k response has a 5xx status code
func (o *ReadDerivedLabelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read derived label o k response a status code equal to that given
func (o *ReadDerivedLabelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read derived label o k response
func (o *ReadDerivedLabelOK) Code() int {
	return 200
}

func (o *ReadDerivedLabelOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels/{slug}][%d] readDerivedLabelOK  %+v", 200, o.Payload)
}

func (o *ReadDerivedLabelOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels/{slug}][%d] readDerivedLabelOK  %+v", 200, o.Payload)
}

func (o *ReadDerivedLabelOK) GetPayload() *models.Configv1ReadDerivedLabelResponse {
	return o.Payload
}

func (o *ReadDerivedLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ReadDerivedLabelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDerivedLabelNotFound creates a ReadDerivedLabelNotFound with default headers values
func NewReadDerivedLabelNotFound() *ReadDerivedLabelNotFound {
	return &ReadDerivedLabelNotFound{}
}

/*
ReadDerivedLabelNotFound describes a response with status code 404, with default header values.

Cannot read the DerivedLabel because the slug does not exist.
*/
type ReadDerivedLabelNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read derived label not found response has a 2xx status code
func (o *ReadDerivedLabelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read derived label not found response has a 3xx status code
func (o *ReadDerivedLabelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read derived label not found response has a 4xx status code
func (o *ReadDerivedLabelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read derived label not found response has a 5xx status code
func (o *ReadDerivedLabelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read derived label not found response a status code equal to that given
func (o *ReadDerivedLabelNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read derived label not found response
func (o *ReadDerivedLabelNotFound) Code() int {
	return 404
}

func (o *ReadDerivedLabelNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels/{slug}][%d] readDerivedLabelNotFound  %+v", 404, o.Payload)
}

func (o *ReadDerivedLabelNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels/{slug}][%d] readDerivedLabelNotFound  %+v", 404, o.Payload)
}

func (o *ReadDerivedLabelNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadDerivedLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDerivedLabelInternalServerError creates a ReadDerivedLabelInternalServerError with default headers values
func NewReadDerivedLabelInternalServerError() *ReadDerivedLabelInternalServerError {
	return &ReadDerivedLabelInternalServerError{}
}

/*
ReadDerivedLabelInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadDerivedLabelInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read derived label internal server error response has a 2xx status code
func (o *ReadDerivedLabelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read derived label internal server error response has a 3xx status code
func (o *ReadDerivedLabelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read derived label internal server error response has a 4xx status code
func (o *ReadDerivedLabelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read derived label internal server error response has a 5xx status code
func (o *ReadDerivedLabelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read derived label internal server error response a status code equal to that given
func (o *ReadDerivedLabelInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read derived label internal server error response
func (o *ReadDerivedLabelInternalServerError) Code() int {
	return 500
}

func (o *ReadDerivedLabelInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels/{slug}][%d] readDerivedLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadDerivedLabelInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels/{slug}][%d] readDerivedLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadDerivedLabelInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadDerivedLabelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDerivedLabelDefault creates a ReadDerivedLabelDefault with default headers values
func NewReadDerivedLabelDefault(code int) *ReadDerivedLabelDefault {
	return &ReadDerivedLabelDefault{
		_statusCode: code,
	}
}

/*
ReadDerivedLabelDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadDerivedLabelDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read derived label default response has a 2xx status code
func (o *ReadDerivedLabelDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read derived label default response has a 3xx status code
func (o *ReadDerivedLabelDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read derived label default response has a 4xx status code
func (o *ReadDerivedLabelDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read derived label default response has a 5xx status code
func (o *ReadDerivedLabelDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read derived label default response a status code equal to that given
func (o *ReadDerivedLabelDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read derived label default response
func (o *ReadDerivedLabelDefault) Code() int {
	return o._statusCode
}

func (o *ReadDerivedLabelDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels/{slug}][%d] ReadDerivedLabel default  %+v", o._statusCode, o.Payload)
}

func (o *ReadDerivedLabelDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels/{slug}][%d] ReadDerivedLabel default  %+v", o._statusCode, o.Payload)
}

func (o *ReadDerivedLabelDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadDerivedLabelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
