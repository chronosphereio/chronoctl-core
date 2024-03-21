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

// DeleteResourcePoolsReader is a Reader for the DeleteResourcePools structure.
type DeleteResourcePoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteResourcePoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteResourcePoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteResourcePoolsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteResourcePoolsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteResourcePoolsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteResourcePoolsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteResourcePoolsOK creates a DeleteResourcePoolsOK with default headers values
func NewDeleteResourcePoolsOK() *DeleteResourcePoolsOK {
	return &DeleteResourcePoolsOK{}
}

/*
DeleteResourcePoolsOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteResourcePoolsOK struct {
	Payload models.Configv1DeleteResourcePoolsResponse
}

// IsSuccess returns true when this delete resource pools o k response has a 2xx status code
func (o *DeleteResourcePoolsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete resource pools o k response has a 3xx status code
func (o *DeleteResourcePoolsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resource pools o k response has a 4xx status code
func (o *DeleteResourcePoolsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete resource pools o k response has a 5xx status code
func (o *DeleteResourcePoolsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resource pools o k response a status code equal to that given
func (o *DeleteResourcePoolsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete resource pools o k response
func (o *DeleteResourcePoolsOK) Code() int {
	return 200
}

func (o *DeleteResourcePoolsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/resource-pools][%d] deleteResourcePoolsOK  %+v", 200, o.Payload)
}

func (o *DeleteResourcePoolsOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/resource-pools][%d] deleteResourcePoolsOK  %+v", 200, o.Payload)
}

func (o *DeleteResourcePoolsOK) GetPayload() models.Configv1DeleteResourcePoolsResponse {
	return o.Payload
}

func (o *DeleteResourcePoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteResourcePoolsBadRequest creates a DeleteResourcePoolsBadRequest with default headers values
func NewDeleteResourcePoolsBadRequest() *DeleteResourcePoolsBadRequest {
	return &DeleteResourcePoolsBadRequest{}
}

/*
DeleteResourcePoolsBadRequest describes a response with status code 400, with default header values.

Cannot delete the ResourcePools because it is in use.
*/
type DeleteResourcePoolsBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete resource pools bad request response has a 2xx status code
func (o *DeleteResourcePoolsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete resource pools bad request response has a 3xx status code
func (o *DeleteResourcePoolsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resource pools bad request response has a 4xx status code
func (o *DeleteResourcePoolsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete resource pools bad request response has a 5xx status code
func (o *DeleteResourcePoolsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resource pools bad request response a status code equal to that given
func (o *DeleteResourcePoolsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete resource pools bad request response
func (o *DeleteResourcePoolsBadRequest) Code() int {
	return 400
}

func (o *DeleteResourcePoolsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/resource-pools][%d] deleteResourcePoolsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteResourcePoolsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/resource-pools][%d] deleteResourcePoolsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteResourcePoolsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteResourcePoolsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteResourcePoolsNotFound creates a DeleteResourcePoolsNotFound with default headers values
func NewDeleteResourcePoolsNotFound() *DeleteResourcePoolsNotFound {
	return &DeleteResourcePoolsNotFound{}
}

/*
DeleteResourcePoolsNotFound describes a response with status code 404, with default header values.

Cannot delete the ResourcePools because the slug does not exist.
*/
type DeleteResourcePoolsNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete resource pools not found response has a 2xx status code
func (o *DeleteResourcePoolsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete resource pools not found response has a 3xx status code
func (o *DeleteResourcePoolsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resource pools not found response has a 4xx status code
func (o *DeleteResourcePoolsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete resource pools not found response has a 5xx status code
func (o *DeleteResourcePoolsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resource pools not found response a status code equal to that given
func (o *DeleteResourcePoolsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete resource pools not found response
func (o *DeleteResourcePoolsNotFound) Code() int {
	return 404
}

func (o *DeleteResourcePoolsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/resource-pools][%d] deleteResourcePoolsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteResourcePoolsNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/resource-pools][%d] deleteResourcePoolsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteResourcePoolsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteResourcePoolsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteResourcePoolsInternalServerError creates a DeleteResourcePoolsInternalServerError with default headers values
func NewDeleteResourcePoolsInternalServerError() *DeleteResourcePoolsInternalServerError {
	return &DeleteResourcePoolsInternalServerError{}
}

/*
DeleteResourcePoolsInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteResourcePoolsInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete resource pools internal server error response has a 2xx status code
func (o *DeleteResourcePoolsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete resource pools internal server error response has a 3xx status code
func (o *DeleteResourcePoolsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resource pools internal server error response has a 4xx status code
func (o *DeleteResourcePoolsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete resource pools internal server error response has a 5xx status code
func (o *DeleteResourcePoolsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete resource pools internal server error response a status code equal to that given
func (o *DeleteResourcePoolsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete resource pools internal server error response
func (o *DeleteResourcePoolsInternalServerError) Code() int {
	return 500
}

func (o *DeleteResourcePoolsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/resource-pools][%d] deleteResourcePoolsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteResourcePoolsInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/resource-pools][%d] deleteResourcePoolsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteResourcePoolsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteResourcePoolsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteResourcePoolsDefault creates a DeleteResourcePoolsDefault with default headers values
func NewDeleteResourcePoolsDefault(code int) *DeleteResourcePoolsDefault {
	return &DeleteResourcePoolsDefault{
		_statusCode: code,
	}
}

/*
DeleteResourcePoolsDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteResourcePoolsDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete resource pools default response has a 2xx status code
func (o *DeleteResourcePoolsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete resource pools default response has a 3xx status code
func (o *DeleteResourcePoolsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete resource pools default response has a 4xx status code
func (o *DeleteResourcePoolsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete resource pools default response has a 5xx status code
func (o *DeleteResourcePoolsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete resource pools default response a status code equal to that given
func (o *DeleteResourcePoolsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete resource pools default response
func (o *DeleteResourcePoolsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteResourcePoolsDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/resource-pools][%d] DeleteResourcePools default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteResourcePoolsDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/resource-pools][%d] DeleteResourcePools default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteResourcePoolsDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteResourcePoolsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
