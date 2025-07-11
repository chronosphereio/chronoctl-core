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

// CreateLogBudgetReader is a Reader for the CreateLogBudget structure.
type CreateLogBudgetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLogBudgetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateLogBudgetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateLogBudgetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateLogBudgetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateLogBudgetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateLogBudgetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateLogBudgetOK creates a CreateLogBudgetOK with default headers values
func NewCreateLogBudgetOK() *CreateLogBudgetOK {
	return &CreateLogBudgetOK{}
}

/*
CreateLogBudgetOK describes a response with status code 200, with default header values.

A successful response containing the created LogBudget.
*/
type CreateLogBudgetOK struct {
	Payload *models.ConfigunstableCreateLogBudgetResponse
}

// IsSuccess returns true when this create log budget o k response has a 2xx status code
func (o *CreateLogBudgetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create log budget o k response has a 3xx status code
func (o *CreateLogBudgetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create log budget o k response has a 4xx status code
func (o *CreateLogBudgetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create log budget o k response has a 5xx status code
func (o *CreateLogBudgetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create log budget o k response a status code equal to that given
func (o *CreateLogBudgetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create log budget o k response
func (o *CreateLogBudgetOK) Code() int {
	return 200
}

func (o *CreateLogBudgetOK) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/log-budget][%d] createLogBudgetOK  %+v", 200, o.Payload)
}

func (o *CreateLogBudgetOK) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/log-budget][%d] createLogBudgetOK  %+v", 200, o.Payload)
}

func (o *CreateLogBudgetOK) GetPayload() *models.ConfigunstableCreateLogBudgetResponse {
	return o.Payload
}

func (o *CreateLogBudgetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableCreateLogBudgetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLogBudgetBadRequest creates a CreateLogBudgetBadRequest with default headers values
func NewCreateLogBudgetBadRequest() *CreateLogBudgetBadRequest {
	return &CreateLogBudgetBadRequest{}
}

/*
CreateLogBudgetBadRequest describes a response with status code 400, with default header values.

Cannot create the LogBudget because the request is invalid.
*/
type CreateLogBudgetBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create log budget bad request response has a 2xx status code
func (o *CreateLogBudgetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create log budget bad request response has a 3xx status code
func (o *CreateLogBudgetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create log budget bad request response has a 4xx status code
func (o *CreateLogBudgetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create log budget bad request response has a 5xx status code
func (o *CreateLogBudgetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create log budget bad request response a status code equal to that given
func (o *CreateLogBudgetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create log budget bad request response
func (o *CreateLogBudgetBadRequest) Code() int {
	return 400
}

func (o *CreateLogBudgetBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/log-budget][%d] createLogBudgetBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLogBudgetBadRequest) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/log-budget][%d] createLogBudgetBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLogBudgetBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateLogBudgetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLogBudgetConflict creates a CreateLogBudgetConflict with default headers values
func NewCreateLogBudgetConflict() *CreateLogBudgetConflict {
	return &CreateLogBudgetConflict{}
}

/*
CreateLogBudgetConflict describes a response with status code 409, with default header values.

Cannot create the LogBudget because there is a conflict with an existing LogBudget.
*/
type CreateLogBudgetConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create log budget conflict response has a 2xx status code
func (o *CreateLogBudgetConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create log budget conflict response has a 3xx status code
func (o *CreateLogBudgetConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create log budget conflict response has a 4xx status code
func (o *CreateLogBudgetConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create log budget conflict response has a 5xx status code
func (o *CreateLogBudgetConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create log budget conflict response a status code equal to that given
func (o *CreateLogBudgetConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create log budget conflict response
func (o *CreateLogBudgetConflict) Code() int {
	return 409
}

func (o *CreateLogBudgetConflict) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/log-budget][%d] createLogBudgetConflict  %+v", 409, o.Payload)
}

func (o *CreateLogBudgetConflict) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/log-budget][%d] createLogBudgetConflict  %+v", 409, o.Payload)
}

func (o *CreateLogBudgetConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateLogBudgetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLogBudgetInternalServerError creates a CreateLogBudgetInternalServerError with default headers values
func NewCreateLogBudgetInternalServerError() *CreateLogBudgetInternalServerError {
	return &CreateLogBudgetInternalServerError{}
}

/*
CreateLogBudgetInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateLogBudgetInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create log budget internal server error response has a 2xx status code
func (o *CreateLogBudgetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create log budget internal server error response has a 3xx status code
func (o *CreateLogBudgetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create log budget internal server error response has a 4xx status code
func (o *CreateLogBudgetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create log budget internal server error response has a 5xx status code
func (o *CreateLogBudgetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create log budget internal server error response a status code equal to that given
func (o *CreateLogBudgetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create log budget internal server error response
func (o *CreateLogBudgetInternalServerError) Code() int {
	return 500
}

func (o *CreateLogBudgetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/log-budget][%d] createLogBudgetInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateLogBudgetInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/log-budget][%d] createLogBudgetInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateLogBudgetInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateLogBudgetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLogBudgetDefault creates a CreateLogBudgetDefault with default headers values
func NewCreateLogBudgetDefault(code int) *CreateLogBudgetDefault {
	return &CreateLogBudgetDefault{
		_statusCode: code,
	}
}

/*
CreateLogBudgetDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateLogBudgetDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create log budget default response has a 2xx status code
func (o *CreateLogBudgetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create log budget default response has a 3xx status code
func (o *CreateLogBudgetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create log budget default response has a 4xx status code
func (o *CreateLogBudgetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create log budget default response has a 5xx status code
func (o *CreateLogBudgetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create log budget default response a status code equal to that given
func (o *CreateLogBudgetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create log budget default response
func (o *CreateLogBudgetDefault) Code() int {
	return o._statusCode
}

func (o *CreateLogBudgetDefault) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/log-budget][%d] CreateLogBudget default  %+v", o._statusCode, o.Payload)
}

func (o *CreateLogBudgetDefault) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/log-budget][%d] CreateLogBudget default  %+v", o._statusCode, o.Payload)
}

func (o *CreateLogBudgetDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateLogBudgetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
