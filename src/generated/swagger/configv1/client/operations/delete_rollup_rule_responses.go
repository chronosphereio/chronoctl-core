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

// DeleteRollupRuleReader is a Reader for the DeleteRollupRule structure.
type DeleteRollupRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRollupRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRollupRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteRollupRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRollupRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteRollupRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRollupRuleOK creates a DeleteRollupRuleOK with default headers values
func NewDeleteRollupRuleOK() *DeleteRollupRuleOK {
	return &DeleteRollupRuleOK{}
}

/*
DeleteRollupRuleOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteRollupRuleOK struct {
	Payload models.Configv1DeleteRollupRuleResponse
}

// IsSuccess returns true when this delete rollup rule o k response has a 2xx status code
func (o *DeleteRollupRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete rollup rule o k response has a 3xx status code
func (o *DeleteRollupRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rollup rule o k response has a 4xx status code
func (o *DeleteRollupRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete rollup rule o k response has a 5xx status code
func (o *DeleteRollupRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rollup rule o k response a status code equal to that given
func (o *DeleteRollupRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete rollup rule o k response
func (o *DeleteRollupRuleOK) Code() int {
	return 200
}

func (o *DeleteRollupRuleOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/rollup-rules/{slug}][%d] deleteRollupRuleOK  %+v", 200, o.Payload)
}

func (o *DeleteRollupRuleOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/rollup-rules/{slug}][%d] deleteRollupRuleOK  %+v", 200, o.Payload)
}

func (o *DeleteRollupRuleOK) GetPayload() models.Configv1DeleteRollupRuleResponse {
	return o.Payload
}

func (o *DeleteRollupRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRollupRuleNotFound creates a DeleteRollupRuleNotFound with default headers values
func NewDeleteRollupRuleNotFound() *DeleteRollupRuleNotFound {
	return &DeleteRollupRuleNotFound{}
}

/*
DeleteRollupRuleNotFound describes a response with status code 404, with default header values.

Cannot delete the RollupRule because the slug does not exist.
*/
type DeleteRollupRuleNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete rollup rule not found response has a 2xx status code
func (o *DeleteRollupRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rollup rule not found response has a 3xx status code
func (o *DeleteRollupRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rollup rule not found response has a 4xx status code
func (o *DeleteRollupRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rollup rule not found response has a 5xx status code
func (o *DeleteRollupRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rollup rule not found response a status code equal to that given
func (o *DeleteRollupRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete rollup rule not found response
func (o *DeleteRollupRuleNotFound) Code() int {
	return 404
}

func (o *DeleteRollupRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/rollup-rules/{slug}][%d] deleteRollupRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRollupRuleNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/rollup-rules/{slug}][%d] deleteRollupRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRollupRuleNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteRollupRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRollupRuleInternalServerError creates a DeleteRollupRuleInternalServerError with default headers values
func NewDeleteRollupRuleInternalServerError() *DeleteRollupRuleInternalServerError {
	return &DeleteRollupRuleInternalServerError{}
}

/*
DeleteRollupRuleInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteRollupRuleInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete rollup rule internal server error response has a 2xx status code
func (o *DeleteRollupRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rollup rule internal server error response has a 3xx status code
func (o *DeleteRollupRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rollup rule internal server error response has a 4xx status code
func (o *DeleteRollupRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete rollup rule internal server error response has a 5xx status code
func (o *DeleteRollupRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete rollup rule internal server error response a status code equal to that given
func (o *DeleteRollupRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete rollup rule internal server error response
func (o *DeleteRollupRuleInternalServerError) Code() int {
	return 500
}

func (o *DeleteRollupRuleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/rollup-rules/{slug}][%d] deleteRollupRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRollupRuleInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/rollup-rules/{slug}][%d] deleteRollupRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRollupRuleInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteRollupRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRollupRuleDefault creates a DeleteRollupRuleDefault with default headers values
func NewDeleteRollupRuleDefault(code int) *DeleteRollupRuleDefault {
	return &DeleteRollupRuleDefault{
		_statusCode: code,
	}
}

/*
DeleteRollupRuleDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteRollupRuleDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete rollup rule default response has a 2xx status code
func (o *DeleteRollupRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete rollup rule default response has a 3xx status code
func (o *DeleteRollupRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete rollup rule default response has a 4xx status code
func (o *DeleteRollupRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete rollup rule default response has a 5xx status code
func (o *DeleteRollupRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete rollup rule default response a status code equal to that given
func (o *DeleteRollupRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete rollup rule default response
func (o *DeleteRollupRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRollupRuleDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/rollup-rules/{slug}][%d] DeleteRollupRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRollupRuleDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/rollup-rules/{slug}][%d] DeleteRollupRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRollupRuleDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteRollupRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
