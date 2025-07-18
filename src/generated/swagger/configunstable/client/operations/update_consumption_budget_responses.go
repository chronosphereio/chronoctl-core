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

// UpdateConsumptionBudgetReader is a Reader for the UpdateConsumptionBudget structure.
type UpdateConsumptionBudgetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConsumptionBudgetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateConsumptionBudgetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateConsumptionBudgetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateConsumptionBudgetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateConsumptionBudgetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateConsumptionBudgetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateConsumptionBudgetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateConsumptionBudgetOK creates a UpdateConsumptionBudgetOK with default headers values
func NewUpdateConsumptionBudgetOK() *UpdateConsumptionBudgetOK {
	return &UpdateConsumptionBudgetOK{}
}

/*
UpdateConsumptionBudgetOK describes a response with status code 200, with default header values.

A successful response containing the updated ConsumptionBudget.
*/
type UpdateConsumptionBudgetOK struct {
	Payload *models.ConfigunstableUpdateConsumptionBudgetResponse
}

// IsSuccess returns true when this update consumption budget o k response has a 2xx status code
func (o *UpdateConsumptionBudgetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update consumption budget o k response has a 3xx status code
func (o *UpdateConsumptionBudgetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update consumption budget o k response has a 4xx status code
func (o *UpdateConsumptionBudgetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update consumption budget o k response has a 5xx status code
func (o *UpdateConsumptionBudgetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update consumption budget o k response a status code equal to that given
func (o *UpdateConsumptionBudgetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update consumption budget o k response
func (o *UpdateConsumptionBudgetOK) Code() int {
	return 200
}

func (o *UpdateConsumptionBudgetOK) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/consumption-budgets/{slug}][%d] updateConsumptionBudgetOK  %+v", 200, o.Payload)
}

func (o *UpdateConsumptionBudgetOK) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/consumption-budgets/{slug}][%d] updateConsumptionBudgetOK  %+v", 200, o.Payload)
}

func (o *UpdateConsumptionBudgetOK) GetPayload() *models.ConfigunstableUpdateConsumptionBudgetResponse {
	return o.Payload
}

func (o *UpdateConsumptionBudgetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableUpdateConsumptionBudgetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsumptionBudgetBadRequest creates a UpdateConsumptionBudgetBadRequest with default headers values
func NewUpdateConsumptionBudgetBadRequest() *UpdateConsumptionBudgetBadRequest {
	return &UpdateConsumptionBudgetBadRequest{}
}

/*
UpdateConsumptionBudgetBadRequest describes a response with status code 400, with default header values.

Cannot update the ConsumptionBudget because the request is invalid.
*/
type UpdateConsumptionBudgetBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update consumption budget bad request response has a 2xx status code
func (o *UpdateConsumptionBudgetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update consumption budget bad request response has a 3xx status code
func (o *UpdateConsumptionBudgetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update consumption budget bad request response has a 4xx status code
func (o *UpdateConsumptionBudgetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update consumption budget bad request response has a 5xx status code
func (o *UpdateConsumptionBudgetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update consumption budget bad request response a status code equal to that given
func (o *UpdateConsumptionBudgetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update consumption budget bad request response
func (o *UpdateConsumptionBudgetBadRequest) Code() int {
	return 400
}

func (o *UpdateConsumptionBudgetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/consumption-budgets/{slug}][%d] updateConsumptionBudgetBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateConsumptionBudgetBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/consumption-budgets/{slug}][%d] updateConsumptionBudgetBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateConsumptionBudgetBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateConsumptionBudgetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsumptionBudgetNotFound creates a UpdateConsumptionBudgetNotFound with default headers values
func NewUpdateConsumptionBudgetNotFound() *UpdateConsumptionBudgetNotFound {
	return &UpdateConsumptionBudgetNotFound{}
}

/*
UpdateConsumptionBudgetNotFound describes a response with status code 404, with default header values.

Cannot update the ConsumptionBudget because the slug does not exist.
*/
type UpdateConsumptionBudgetNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update consumption budget not found response has a 2xx status code
func (o *UpdateConsumptionBudgetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update consumption budget not found response has a 3xx status code
func (o *UpdateConsumptionBudgetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update consumption budget not found response has a 4xx status code
func (o *UpdateConsumptionBudgetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update consumption budget not found response has a 5xx status code
func (o *UpdateConsumptionBudgetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update consumption budget not found response a status code equal to that given
func (o *UpdateConsumptionBudgetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update consumption budget not found response
func (o *UpdateConsumptionBudgetNotFound) Code() int {
	return 404
}

func (o *UpdateConsumptionBudgetNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/consumption-budgets/{slug}][%d] updateConsumptionBudgetNotFound  %+v", 404, o.Payload)
}

func (o *UpdateConsumptionBudgetNotFound) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/consumption-budgets/{slug}][%d] updateConsumptionBudgetNotFound  %+v", 404, o.Payload)
}

func (o *UpdateConsumptionBudgetNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateConsumptionBudgetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsumptionBudgetConflict creates a UpdateConsumptionBudgetConflict with default headers values
func NewUpdateConsumptionBudgetConflict() *UpdateConsumptionBudgetConflict {
	return &UpdateConsumptionBudgetConflict{}
}

/*
UpdateConsumptionBudgetConflict describes a response with status code 409, with default header values.

Cannot update the ConsumptionBudget because there is a conflict with an existing ConsumptionBudget.
*/
type UpdateConsumptionBudgetConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update consumption budget conflict response has a 2xx status code
func (o *UpdateConsumptionBudgetConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update consumption budget conflict response has a 3xx status code
func (o *UpdateConsumptionBudgetConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update consumption budget conflict response has a 4xx status code
func (o *UpdateConsumptionBudgetConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update consumption budget conflict response has a 5xx status code
func (o *UpdateConsumptionBudgetConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update consumption budget conflict response a status code equal to that given
func (o *UpdateConsumptionBudgetConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update consumption budget conflict response
func (o *UpdateConsumptionBudgetConflict) Code() int {
	return 409
}

func (o *UpdateConsumptionBudgetConflict) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/consumption-budgets/{slug}][%d] updateConsumptionBudgetConflict  %+v", 409, o.Payload)
}

func (o *UpdateConsumptionBudgetConflict) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/consumption-budgets/{slug}][%d] updateConsumptionBudgetConflict  %+v", 409, o.Payload)
}

func (o *UpdateConsumptionBudgetConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateConsumptionBudgetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsumptionBudgetInternalServerError creates a UpdateConsumptionBudgetInternalServerError with default headers values
func NewUpdateConsumptionBudgetInternalServerError() *UpdateConsumptionBudgetInternalServerError {
	return &UpdateConsumptionBudgetInternalServerError{}
}

/*
UpdateConsumptionBudgetInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateConsumptionBudgetInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update consumption budget internal server error response has a 2xx status code
func (o *UpdateConsumptionBudgetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update consumption budget internal server error response has a 3xx status code
func (o *UpdateConsumptionBudgetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update consumption budget internal server error response has a 4xx status code
func (o *UpdateConsumptionBudgetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update consumption budget internal server error response has a 5xx status code
func (o *UpdateConsumptionBudgetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update consumption budget internal server error response a status code equal to that given
func (o *UpdateConsumptionBudgetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update consumption budget internal server error response
func (o *UpdateConsumptionBudgetInternalServerError) Code() int {
	return 500
}

func (o *UpdateConsumptionBudgetInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/consumption-budgets/{slug}][%d] updateConsumptionBudgetInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateConsumptionBudgetInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/consumption-budgets/{slug}][%d] updateConsumptionBudgetInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateConsumptionBudgetInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateConsumptionBudgetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsumptionBudgetDefault creates a UpdateConsumptionBudgetDefault with default headers values
func NewUpdateConsumptionBudgetDefault(code int) *UpdateConsumptionBudgetDefault {
	return &UpdateConsumptionBudgetDefault{
		_statusCode: code,
	}
}

/*
UpdateConsumptionBudgetDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateConsumptionBudgetDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update consumption budget default response has a 2xx status code
func (o *UpdateConsumptionBudgetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update consumption budget default response has a 3xx status code
func (o *UpdateConsumptionBudgetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update consumption budget default response has a 4xx status code
func (o *UpdateConsumptionBudgetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update consumption budget default response has a 5xx status code
func (o *UpdateConsumptionBudgetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update consumption budget default response a status code equal to that given
func (o *UpdateConsumptionBudgetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update consumption budget default response
func (o *UpdateConsumptionBudgetDefault) Code() int {
	return o._statusCode
}

func (o *UpdateConsumptionBudgetDefault) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/consumption-budgets/{slug}][%d] UpdateConsumptionBudget default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateConsumptionBudgetDefault) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/consumption-budgets/{slug}][%d] UpdateConsumptionBudget default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateConsumptionBudgetDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateConsumptionBudgetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
