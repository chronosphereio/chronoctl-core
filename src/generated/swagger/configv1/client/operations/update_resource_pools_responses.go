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

// UpdateResourcePoolsReader is a Reader for the UpdateResourcePools structure.
type UpdateResourcePoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateResourcePoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateResourcePoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateResourcePoolsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateResourcePoolsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateResourcePoolsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateResourcePoolsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateResourcePoolsOK creates a UpdateResourcePoolsOK with default headers values
func NewUpdateResourcePoolsOK() *UpdateResourcePoolsOK {
	return &UpdateResourcePoolsOK{}
}

/*
UpdateResourcePoolsOK describes a response with status code 200, with default header values.

A successful response containing the updated ResourcePools.
*/
type UpdateResourcePoolsOK struct {
	Payload *models.Configv1UpdateResourcePoolsResponse
}

// IsSuccess returns true when this update resource pools o k response has a 2xx status code
func (o *UpdateResourcePoolsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update resource pools o k response has a 3xx status code
func (o *UpdateResourcePoolsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource pools o k response has a 4xx status code
func (o *UpdateResourcePoolsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update resource pools o k response has a 5xx status code
func (o *UpdateResourcePoolsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update resource pools o k response a status code equal to that given
func (o *UpdateResourcePoolsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update resource pools o k response
func (o *UpdateResourcePoolsOK) Code() int {
	return 200
}

func (o *UpdateResourcePoolsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/resource-pools][%d] updateResourcePoolsOK  %+v", 200, o.Payload)
}

func (o *UpdateResourcePoolsOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/resource-pools][%d] updateResourcePoolsOK  %+v", 200, o.Payload)
}

func (o *UpdateResourcePoolsOK) GetPayload() *models.Configv1UpdateResourcePoolsResponse {
	return o.Payload
}

func (o *UpdateResourcePoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1UpdateResourcePoolsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourcePoolsBadRequest creates a UpdateResourcePoolsBadRequest with default headers values
func NewUpdateResourcePoolsBadRequest() *UpdateResourcePoolsBadRequest {
	return &UpdateResourcePoolsBadRequest{}
}

/*
UpdateResourcePoolsBadRequest describes a response with status code 400, with default header values.

Cannot update the ResourcePools because the request is invalid.
*/
type UpdateResourcePoolsBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update resource pools bad request response has a 2xx status code
func (o *UpdateResourcePoolsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update resource pools bad request response has a 3xx status code
func (o *UpdateResourcePoolsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource pools bad request response has a 4xx status code
func (o *UpdateResourcePoolsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update resource pools bad request response has a 5xx status code
func (o *UpdateResourcePoolsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update resource pools bad request response a status code equal to that given
func (o *UpdateResourcePoolsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update resource pools bad request response
func (o *UpdateResourcePoolsBadRequest) Code() int {
	return 400
}

func (o *UpdateResourcePoolsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/resource-pools][%d] updateResourcePoolsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateResourcePoolsBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/resource-pools][%d] updateResourcePoolsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateResourcePoolsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateResourcePoolsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourcePoolsNotFound creates a UpdateResourcePoolsNotFound with default headers values
func NewUpdateResourcePoolsNotFound() *UpdateResourcePoolsNotFound {
	return &UpdateResourcePoolsNotFound{}
}

/*
UpdateResourcePoolsNotFound describes a response with status code 404, with default header values.

Cannot update the ResourcePools because ResourcePools has not been created.
*/
type UpdateResourcePoolsNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update resource pools not found response has a 2xx status code
func (o *UpdateResourcePoolsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update resource pools not found response has a 3xx status code
func (o *UpdateResourcePoolsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource pools not found response has a 4xx status code
func (o *UpdateResourcePoolsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update resource pools not found response has a 5xx status code
func (o *UpdateResourcePoolsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update resource pools not found response a status code equal to that given
func (o *UpdateResourcePoolsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update resource pools not found response
func (o *UpdateResourcePoolsNotFound) Code() int {
	return 404
}

func (o *UpdateResourcePoolsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/resource-pools][%d] updateResourcePoolsNotFound  %+v", 404, o.Payload)
}

func (o *UpdateResourcePoolsNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/resource-pools][%d] updateResourcePoolsNotFound  %+v", 404, o.Payload)
}

func (o *UpdateResourcePoolsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateResourcePoolsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourcePoolsInternalServerError creates a UpdateResourcePoolsInternalServerError with default headers values
func NewUpdateResourcePoolsInternalServerError() *UpdateResourcePoolsInternalServerError {
	return &UpdateResourcePoolsInternalServerError{}
}

/*
UpdateResourcePoolsInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateResourcePoolsInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update resource pools internal server error response has a 2xx status code
func (o *UpdateResourcePoolsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update resource pools internal server error response has a 3xx status code
func (o *UpdateResourcePoolsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource pools internal server error response has a 4xx status code
func (o *UpdateResourcePoolsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update resource pools internal server error response has a 5xx status code
func (o *UpdateResourcePoolsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update resource pools internal server error response a status code equal to that given
func (o *UpdateResourcePoolsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update resource pools internal server error response
func (o *UpdateResourcePoolsInternalServerError) Code() int {
	return 500
}

func (o *UpdateResourcePoolsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/resource-pools][%d] updateResourcePoolsInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateResourcePoolsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/resource-pools][%d] updateResourcePoolsInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateResourcePoolsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateResourcePoolsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourcePoolsDefault creates a UpdateResourcePoolsDefault with default headers values
func NewUpdateResourcePoolsDefault(code int) *UpdateResourcePoolsDefault {
	return &UpdateResourcePoolsDefault{
		_statusCode: code,
	}
}

/*
UpdateResourcePoolsDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateResourcePoolsDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update resource pools default response has a 2xx status code
func (o *UpdateResourcePoolsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update resource pools default response has a 3xx status code
func (o *UpdateResourcePoolsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update resource pools default response has a 4xx status code
func (o *UpdateResourcePoolsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update resource pools default response has a 5xx status code
func (o *UpdateResourcePoolsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update resource pools default response a status code equal to that given
func (o *UpdateResourcePoolsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update resource pools default response
func (o *UpdateResourcePoolsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateResourcePoolsDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/resource-pools][%d] UpdateResourcePools default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateResourcePoolsDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/resource-pools][%d] UpdateResourcePools default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateResourcePoolsDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateResourcePoolsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
