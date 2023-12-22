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

// ListDropRulesReader is a Reader for the ListDropRules structure.
type ListDropRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDropRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDropRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListDropRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListDropRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDropRulesOK creates a ListDropRulesOK with default headers values
func NewListDropRulesOK() *ListDropRulesOK {
	return &ListDropRulesOK{}
}

/*
ListDropRulesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListDropRulesOK struct {
	Payload *models.Configv1ListDropRulesResponse
}

// IsSuccess returns true when this list drop rules o k response has a 2xx status code
func (o *ListDropRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list drop rules o k response has a 3xx status code
func (o *ListDropRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list drop rules o k response has a 4xx status code
func (o *ListDropRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list drop rules o k response has a 5xx status code
func (o *ListDropRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list drop rules o k response a status code equal to that given
func (o *ListDropRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list drop rules o k response
func (o *ListDropRulesOK) Code() int {
	return 200
}

func (o *ListDropRulesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/drop-rules][%d] listDropRulesOK  %+v", 200, o.Payload)
}

func (o *ListDropRulesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/drop-rules][%d] listDropRulesOK  %+v", 200, o.Payload)
}

func (o *ListDropRulesOK) GetPayload() *models.Configv1ListDropRulesResponse {
	return o.Payload
}

func (o *ListDropRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ListDropRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDropRulesInternalServerError creates a ListDropRulesInternalServerError with default headers values
func NewListDropRulesInternalServerError() *ListDropRulesInternalServerError {
	return &ListDropRulesInternalServerError{}
}

/*
ListDropRulesInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ListDropRulesInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this list drop rules internal server error response has a 2xx status code
func (o *ListDropRulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list drop rules internal server error response has a 3xx status code
func (o *ListDropRulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list drop rules internal server error response has a 4xx status code
func (o *ListDropRulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list drop rules internal server error response has a 5xx status code
func (o *ListDropRulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list drop rules internal server error response a status code equal to that given
func (o *ListDropRulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list drop rules internal server error response
func (o *ListDropRulesInternalServerError) Code() int {
	return 500
}

func (o *ListDropRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/drop-rules][%d] listDropRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListDropRulesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/drop-rules][%d] listDropRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListDropRulesInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListDropRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDropRulesDefault creates a ListDropRulesDefault with default headers values
func NewListDropRulesDefault(code int) *ListDropRulesDefault {
	return &ListDropRulesDefault{
		_statusCode: code,
	}
}

/*
ListDropRulesDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ListDropRulesDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this list drop rules default response has a 2xx status code
func (o *ListDropRulesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list drop rules default response has a 3xx status code
func (o *ListDropRulesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list drop rules default response has a 4xx status code
func (o *ListDropRulesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list drop rules default response has a 5xx status code
func (o *ListDropRulesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list drop rules default response a status code equal to that given
func (o *ListDropRulesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list drop rules default response
func (o *ListDropRulesDefault) Code() int {
	return o._statusCode
}

func (o *ListDropRulesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/drop-rules][%d] ListDropRules default  %+v", o._statusCode, o.Payload)
}

func (o *ListDropRulesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/drop-rules][%d] ListDropRules default  %+v", o._statusCode, o.Payload)
}

func (o *ListDropRulesDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ListDropRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
