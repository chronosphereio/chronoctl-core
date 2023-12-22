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

// DeleteSavedTraceSearchReader is a Reader for the DeleteSavedTraceSearch structure.
type DeleteSavedTraceSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSavedTraceSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSavedTraceSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSavedTraceSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSavedTraceSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSavedTraceSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSavedTraceSearchOK creates a DeleteSavedTraceSearchOK with default headers values
func NewDeleteSavedTraceSearchOK() *DeleteSavedTraceSearchOK {
	return &DeleteSavedTraceSearchOK{}
}

/*
DeleteSavedTraceSearchOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteSavedTraceSearchOK struct {
	Payload models.ConfigunstableDeleteSavedTraceSearchResponse
}

// IsSuccess returns true when this delete saved trace search o k response has a 2xx status code
func (o *DeleteSavedTraceSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete saved trace search o k response has a 3xx status code
func (o *DeleteSavedTraceSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete saved trace search o k response has a 4xx status code
func (o *DeleteSavedTraceSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete saved trace search o k response has a 5xx status code
func (o *DeleteSavedTraceSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete saved trace search o k response a status code equal to that given
func (o *DeleteSavedTraceSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete saved trace search o k response
func (o *DeleteSavedTraceSearchOK) Code() int {
	return 200
}

func (o *DeleteSavedTraceSearchOK) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/saved-trace-searches/{slug}][%d] deleteSavedTraceSearchOK  %+v", 200, o.Payload)
}

func (o *DeleteSavedTraceSearchOK) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/saved-trace-searches/{slug}][%d] deleteSavedTraceSearchOK  %+v", 200, o.Payload)
}

func (o *DeleteSavedTraceSearchOK) GetPayload() models.ConfigunstableDeleteSavedTraceSearchResponse {
	return o.Payload
}

func (o *DeleteSavedTraceSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSavedTraceSearchNotFound creates a DeleteSavedTraceSearchNotFound with default headers values
func NewDeleteSavedTraceSearchNotFound() *DeleteSavedTraceSearchNotFound {
	return &DeleteSavedTraceSearchNotFound{}
}

/*
DeleteSavedTraceSearchNotFound describes a response with status code 404, with default header values.

Cannot delete the SavedTraceSearch because the slug does not exist.
*/
type DeleteSavedTraceSearchNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete saved trace search not found response has a 2xx status code
func (o *DeleteSavedTraceSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete saved trace search not found response has a 3xx status code
func (o *DeleteSavedTraceSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete saved trace search not found response has a 4xx status code
func (o *DeleteSavedTraceSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete saved trace search not found response has a 5xx status code
func (o *DeleteSavedTraceSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete saved trace search not found response a status code equal to that given
func (o *DeleteSavedTraceSearchNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete saved trace search not found response
func (o *DeleteSavedTraceSearchNotFound) Code() int {
	return 404
}

func (o *DeleteSavedTraceSearchNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/saved-trace-searches/{slug}][%d] deleteSavedTraceSearchNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSavedTraceSearchNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/saved-trace-searches/{slug}][%d] deleteSavedTraceSearchNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSavedTraceSearchNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSavedTraceSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSavedTraceSearchInternalServerError creates a DeleteSavedTraceSearchInternalServerError with default headers values
func NewDeleteSavedTraceSearchInternalServerError() *DeleteSavedTraceSearchInternalServerError {
	return &DeleteSavedTraceSearchInternalServerError{}
}

/*
DeleteSavedTraceSearchInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteSavedTraceSearchInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete saved trace search internal server error response has a 2xx status code
func (o *DeleteSavedTraceSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete saved trace search internal server error response has a 3xx status code
func (o *DeleteSavedTraceSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete saved trace search internal server error response has a 4xx status code
func (o *DeleteSavedTraceSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete saved trace search internal server error response has a 5xx status code
func (o *DeleteSavedTraceSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete saved trace search internal server error response a status code equal to that given
func (o *DeleteSavedTraceSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete saved trace search internal server error response
func (o *DeleteSavedTraceSearchInternalServerError) Code() int {
	return 500
}

func (o *DeleteSavedTraceSearchInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/saved-trace-searches/{slug}][%d] deleteSavedTraceSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSavedTraceSearchInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/saved-trace-searches/{slug}][%d] deleteSavedTraceSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSavedTraceSearchInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSavedTraceSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSavedTraceSearchDefault creates a DeleteSavedTraceSearchDefault with default headers values
func NewDeleteSavedTraceSearchDefault(code int) *DeleteSavedTraceSearchDefault {
	return &DeleteSavedTraceSearchDefault{
		_statusCode: code,
	}
}

/*
DeleteSavedTraceSearchDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteSavedTraceSearchDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete saved trace search default response has a 2xx status code
func (o *DeleteSavedTraceSearchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete saved trace search default response has a 3xx status code
func (o *DeleteSavedTraceSearchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete saved trace search default response has a 4xx status code
func (o *DeleteSavedTraceSearchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete saved trace search default response has a 5xx status code
func (o *DeleteSavedTraceSearchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete saved trace search default response a status code equal to that given
func (o *DeleteSavedTraceSearchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete saved trace search default response
func (o *DeleteSavedTraceSearchDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSavedTraceSearchDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/saved-trace-searches/{slug}][%d] DeleteSavedTraceSearch default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSavedTraceSearchDefault) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/saved-trace-searches/{slug}][%d] DeleteSavedTraceSearch default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSavedTraceSearchDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteSavedTraceSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
