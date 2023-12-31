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

// ListNoopEntitiesReader is a Reader for the ListNoopEntities structure.
type ListNoopEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNoopEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNoopEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListNoopEntitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListNoopEntitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNoopEntitiesOK creates a ListNoopEntitiesOK with default headers values
func NewListNoopEntitiesOK() *ListNoopEntitiesOK {
	return &ListNoopEntitiesOK{}
}

/*
ListNoopEntitiesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListNoopEntitiesOK struct {
	Payload *models.ConfigunstableListNoopEntitiesResponse
}

// IsSuccess returns true when this list noop entities o k response has a 2xx status code
func (o *ListNoopEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list noop entities o k response has a 3xx status code
func (o *ListNoopEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list noop entities o k response has a 4xx status code
func (o *ListNoopEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list noop entities o k response has a 5xx status code
func (o *ListNoopEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list noop entities o k response a status code equal to that given
func (o *ListNoopEntitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list noop entities o k response
func (o *ListNoopEntitiesOK) Code() int {
	return 200
}

func (o *ListNoopEntitiesOK) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/noop-entities][%d] listNoopEntitiesOK  %+v", 200, o.Payload)
}

func (o *ListNoopEntitiesOK) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/noop-entities][%d] listNoopEntitiesOK  %+v", 200, o.Payload)
}

func (o *ListNoopEntitiesOK) GetPayload() *models.ConfigunstableListNoopEntitiesResponse {
	return o.Payload
}

func (o *ListNoopEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableListNoopEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNoopEntitiesInternalServerError creates a ListNoopEntitiesInternalServerError with default headers values
func NewListNoopEntitiesInternalServerError() *ListNoopEntitiesInternalServerError {
	return &ListNoopEntitiesInternalServerError{}
}

/*
ListNoopEntitiesInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ListNoopEntitiesInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this list noop entities internal server error response has a 2xx status code
func (o *ListNoopEntitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list noop entities internal server error response has a 3xx status code
func (o *ListNoopEntitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list noop entities internal server error response has a 4xx status code
func (o *ListNoopEntitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list noop entities internal server error response has a 5xx status code
func (o *ListNoopEntitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list noop entities internal server error response a status code equal to that given
func (o *ListNoopEntitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list noop entities internal server error response
func (o *ListNoopEntitiesInternalServerError) Code() int {
	return 500
}

func (o *ListNoopEntitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/noop-entities][%d] listNoopEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListNoopEntitiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/noop-entities][%d] listNoopEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListNoopEntitiesInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListNoopEntitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNoopEntitiesDefault creates a ListNoopEntitiesDefault with default headers values
func NewListNoopEntitiesDefault(code int) *ListNoopEntitiesDefault {
	return &ListNoopEntitiesDefault{
		_statusCode: code,
	}
}

/*
ListNoopEntitiesDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ListNoopEntitiesDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this list noop entities default response has a 2xx status code
func (o *ListNoopEntitiesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list noop entities default response has a 3xx status code
func (o *ListNoopEntitiesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list noop entities default response has a 4xx status code
func (o *ListNoopEntitiesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list noop entities default response has a 5xx status code
func (o *ListNoopEntitiesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list noop entities default response a status code equal to that given
func (o *ListNoopEntitiesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list noop entities default response
func (o *ListNoopEntitiesDefault) Code() int {
	return o._statusCode
}

func (o *ListNoopEntitiesDefault) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/noop-entities][%d] ListNoopEntities default  %+v", o._statusCode, o.Payload)
}

func (o *ListNoopEntitiesDefault) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/noop-entities][%d] ListNoopEntities default  %+v", o._statusCode, o.Payload)
}

func (o *ListNoopEntitiesDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ListNoopEntitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
