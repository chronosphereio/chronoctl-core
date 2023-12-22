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

// ListCollectionsReader is a Reader for the ListCollections structure.
type ListCollectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCollectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCollectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListCollectionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListCollectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListCollectionsOK creates a ListCollectionsOK with default headers values
func NewListCollectionsOK() *ListCollectionsOK {
	return &ListCollectionsOK{}
}

/*
ListCollectionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListCollectionsOK struct {
	Payload *models.Configv1ListCollectionsResponse
}

// IsSuccess returns true when this list collections o k response has a 2xx status code
func (o *ListCollectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list collections o k response has a 3xx status code
func (o *ListCollectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list collections o k response has a 4xx status code
func (o *ListCollectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list collections o k response has a 5xx status code
func (o *ListCollectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list collections o k response a status code equal to that given
func (o *ListCollectionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list collections o k response
func (o *ListCollectionsOK) Code() int {
	return 200
}

func (o *ListCollectionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/collections][%d] listCollectionsOK  %+v", 200, o.Payload)
}

func (o *ListCollectionsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/collections][%d] listCollectionsOK  %+v", 200, o.Payload)
}

func (o *ListCollectionsOK) GetPayload() *models.Configv1ListCollectionsResponse {
	return o.Payload
}

func (o *ListCollectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ListCollectionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCollectionsInternalServerError creates a ListCollectionsInternalServerError with default headers values
func NewListCollectionsInternalServerError() *ListCollectionsInternalServerError {
	return &ListCollectionsInternalServerError{}
}

/*
ListCollectionsInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ListCollectionsInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this list collections internal server error response has a 2xx status code
func (o *ListCollectionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list collections internal server error response has a 3xx status code
func (o *ListCollectionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list collections internal server error response has a 4xx status code
func (o *ListCollectionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list collections internal server error response has a 5xx status code
func (o *ListCollectionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list collections internal server error response a status code equal to that given
func (o *ListCollectionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list collections internal server error response
func (o *ListCollectionsInternalServerError) Code() int {
	return 500
}

func (o *ListCollectionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/collections][%d] listCollectionsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListCollectionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/collections][%d] listCollectionsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListCollectionsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListCollectionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCollectionsDefault creates a ListCollectionsDefault with default headers values
func NewListCollectionsDefault(code int) *ListCollectionsDefault {
	return &ListCollectionsDefault{
		_statusCode: code,
	}
}

/*
ListCollectionsDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ListCollectionsDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this list collections default response has a 2xx status code
func (o *ListCollectionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list collections default response has a 3xx status code
func (o *ListCollectionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list collections default response has a 4xx status code
func (o *ListCollectionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list collections default response has a 5xx status code
func (o *ListCollectionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list collections default response a status code equal to that given
func (o *ListCollectionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list collections default response
func (o *ListCollectionsDefault) Code() int {
	return o._statusCode
}

func (o *ListCollectionsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/collections][%d] ListCollections default  %+v", o._statusCode, o.Payload)
}

func (o *ListCollectionsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/collections][%d] ListCollections default  %+v", o._statusCode, o.Payload)
}

func (o *ListCollectionsDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ListCollectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
