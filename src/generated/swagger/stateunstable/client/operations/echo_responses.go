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

// EchoReader is a Reader for the Echo structure.
type EchoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EchoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEchoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEchoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEchoOK creates a EchoOK with default headers values
func NewEchoOK() *EchoOK {
	return &EchoOK{}
}

/*
EchoOK describes a response with status code 200, with default header values.

A successful response.
*/
type EchoOK struct {
	Payload *models.StateunstableEchoResponse
}

// IsSuccess returns true when this echo o k response has a 2xx status code
func (o *EchoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this echo o k response has a 3xx status code
func (o *EchoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this echo o k response has a 4xx status code
func (o *EchoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this echo o k response has a 5xx status code
func (o *EchoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this echo o k response a status code equal to that given
func (o *EchoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the echo o k response
func (o *EchoOK) Code() int {
	return 200
}

func (o *EchoOK) Error() string {
	return fmt.Sprintf("[GET /api/unstable/data/echo][%d] echoOK  %+v", 200, o.Payload)
}

func (o *EchoOK) String() string {
	return fmt.Sprintf("[GET /api/unstable/data/echo][%d] echoOK  %+v", 200, o.Payload)
}

func (o *EchoOK) GetPayload() *models.StateunstableEchoResponse {
	return o.Payload
}

func (o *EchoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StateunstableEchoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEchoDefault creates a EchoDefault with default headers values
func NewEchoDefault(code int) *EchoDefault {
	return &EchoDefault{
		_statusCode: code,
	}
}

/*
EchoDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EchoDefault struct {
	_statusCode int

	Payload *models.APIError
}

// IsSuccess returns true when this echo default response has a 2xx status code
func (o *EchoDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this echo default response has a 3xx status code
func (o *EchoDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this echo default response has a 4xx status code
func (o *EchoDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this echo default response has a 5xx status code
func (o *EchoDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this echo default response a status code equal to that given
func (o *EchoDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the echo default response
func (o *EchoDefault) Code() int {
	return o._statusCode
}

func (o *EchoDefault) Error() string {
	return fmt.Sprintf("[GET /api/unstable/data/echo][%d] Echo default  %+v", o._statusCode, o.Payload)
}

func (o *EchoDefault) String() string {
	return fmt.Sprintf("[GET /api/unstable/data/echo][%d] Echo default  %+v", o._statusCode, o.Payload)
}

func (o *EchoDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *EchoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
