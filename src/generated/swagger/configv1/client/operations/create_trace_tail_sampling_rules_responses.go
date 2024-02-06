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

// CreateTraceTailSamplingRulesReader is a Reader for the CreateTraceTailSamplingRules structure.
type CreateTraceTailSamplingRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTraceTailSamplingRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTraceTailSamplingRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTraceTailSamplingRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateTraceTailSamplingRulesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTraceTailSamplingRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateTraceTailSamplingRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTraceTailSamplingRulesOK creates a CreateTraceTailSamplingRulesOK with default headers values
func NewCreateTraceTailSamplingRulesOK() *CreateTraceTailSamplingRulesOK {
	return &CreateTraceTailSamplingRulesOK{}
}

/*
CreateTraceTailSamplingRulesOK describes a response with status code 200, with default header values.

A successful response containing the created TraceTailSamplingRules.
*/
type CreateTraceTailSamplingRulesOK struct {
	Payload *models.Configv1CreateTraceTailSamplingRulesResponse
}

// IsSuccess returns true when this create trace tail sampling rules o k response has a 2xx status code
func (o *CreateTraceTailSamplingRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create trace tail sampling rules o k response has a 3xx status code
func (o *CreateTraceTailSamplingRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create trace tail sampling rules o k response has a 4xx status code
func (o *CreateTraceTailSamplingRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create trace tail sampling rules o k response has a 5xx status code
func (o *CreateTraceTailSamplingRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create trace tail sampling rules o k response a status code equal to that given
func (o *CreateTraceTailSamplingRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create trace tail sampling rules o k response
func (o *CreateTraceTailSamplingRulesOK) Code() int {
	return 200
}

func (o *CreateTraceTailSamplingRulesOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-tail-sampling-rules][%d] createTraceTailSamplingRulesOK  %+v", 200, o.Payload)
}

func (o *CreateTraceTailSamplingRulesOK) String() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-tail-sampling-rules][%d] createTraceTailSamplingRulesOK  %+v", 200, o.Payload)
}

func (o *CreateTraceTailSamplingRulesOK) GetPayload() *models.Configv1CreateTraceTailSamplingRulesResponse {
	return o.Payload
}

func (o *CreateTraceTailSamplingRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1CreateTraceTailSamplingRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTraceTailSamplingRulesBadRequest creates a CreateTraceTailSamplingRulesBadRequest with default headers values
func NewCreateTraceTailSamplingRulesBadRequest() *CreateTraceTailSamplingRulesBadRequest {
	return &CreateTraceTailSamplingRulesBadRequest{}
}

/*
CreateTraceTailSamplingRulesBadRequest describes a response with status code 400, with default header values.

Cannot create the TraceTailSamplingRules because the request is invalid.
*/
type CreateTraceTailSamplingRulesBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create trace tail sampling rules bad request response has a 2xx status code
func (o *CreateTraceTailSamplingRulesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create trace tail sampling rules bad request response has a 3xx status code
func (o *CreateTraceTailSamplingRulesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create trace tail sampling rules bad request response has a 4xx status code
func (o *CreateTraceTailSamplingRulesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create trace tail sampling rules bad request response has a 5xx status code
func (o *CreateTraceTailSamplingRulesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create trace tail sampling rules bad request response a status code equal to that given
func (o *CreateTraceTailSamplingRulesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create trace tail sampling rules bad request response
func (o *CreateTraceTailSamplingRulesBadRequest) Code() int {
	return 400
}

func (o *CreateTraceTailSamplingRulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-tail-sampling-rules][%d] createTraceTailSamplingRulesBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTraceTailSamplingRulesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-tail-sampling-rules][%d] createTraceTailSamplingRulesBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTraceTailSamplingRulesBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateTraceTailSamplingRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTraceTailSamplingRulesConflict creates a CreateTraceTailSamplingRulesConflict with default headers values
func NewCreateTraceTailSamplingRulesConflict() *CreateTraceTailSamplingRulesConflict {
	return &CreateTraceTailSamplingRulesConflict{}
}

/*
CreateTraceTailSamplingRulesConflict describes a response with status code 409, with default header values.

Cannot create the TraceTailSamplingRules because there is a conflict with an existing TraceTailSamplingRules.
*/
type CreateTraceTailSamplingRulesConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create trace tail sampling rules conflict response has a 2xx status code
func (o *CreateTraceTailSamplingRulesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create trace tail sampling rules conflict response has a 3xx status code
func (o *CreateTraceTailSamplingRulesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create trace tail sampling rules conflict response has a 4xx status code
func (o *CreateTraceTailSamplingRulesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create trace tail sampling rules conflict response has a 5xx status code
func (o *CreateTraceTailSamplingRulesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create trace tail sampling rules conflict response a status code equal to that given
func (o *CreateTraceTailSamplingRulesConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create trace tail sampling rules conflict response
func (o *CreateTraceTailSamplingRulesConflict) Code() int {
	return 409
}

func (o *CreateTraceTailSamplingRulesConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-tail-sampling-rules][%d] createTraceTailSamplingRulesConflict  %+v", 409, o.Payload)
}

func (o *CreateTraceTailSamplingRulesConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-tail-sampling-rules][%d] createTraceTailSamplingRulesConflict  %+v", 409, o.Payload)
}

func (o *CreateTraceTailSamplingRulesConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateTraceTailSamplingRulesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTraceTailSamplingRulesInternalServerError creates a CreateTraceTailSamplingRulesInternalServerError with default headers values
func NewCreateTraceTailSamplingRulesInternalServerError() *CreateTraceTailSamplingRulesInternalServerError {
	return &CreateTraceTailSamplingRulesInternalServerError{}
}

/*
CreateTraceTailSamplingRulesInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateTraceTailSamplingRulesInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create trace tail sampling rules internal server error response has a 2xx status code
func (o *CreateTraceTailSamplingRulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create trace tail sampling rules internal server error response has a 3xx status code
func (o *CreateTraceTailSamplingRulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create trace tail sampling rules internal server error response has a 4xx status code
func (o *CreateTraceTailSamplingRulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create trace tail sampling rules internal server error response has a 5xx status code
func (o *CreateTraceTailSamplingRulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create trace tail sampling rules internal server error response a status code equal to that given
func (o *CreateTraceTailSamplingRulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create trace tail sampling rules internal server error response
func (o *CreateTraceTailSamplingRulesInternalServerError) Code() int {
	return 500
}

func (o *CreateTraceTailSamplingRulesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-tail-sampling-rules][%d] createTraceTailSamplingRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTraceTailSamplingRulesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-tail-sampling-rules][%d] createTraceTailSamplingRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTraceTailSamplingRulesInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateTraceTailSamplingRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTraceTailSamplingRulesDefault creates a CreateTraceTailSamplingRulesDefault with default headers values
func NewCreateTraceTailSamplingRulesDefault(code int) *CreateTraceTailSamplingRulesDefault {
	return &CreateTraceTailSamplingRulesDefault{
		_statusCode: code,
	}
}

/*
CreateTraceTailSamplingRulesDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateTraceTailSamplingRulesDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create trace tail sampling rules default response has a 2xx status code
func (o *CreateTraceTailSamplingRulesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create trace tail sampling rules default response has a 3xx status code
func (o *CreateTraceTailSamplingRulesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create trace tail sampling rules default response has a 4xx status code
func (o *CreateTraceTailSamplingRulesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create trace tail sampling rules default response has a 5xx status code
func (o *CreateTraceTailSamplingRulesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create trace tail sampling rules default response a status code equal to that given
func (o *CreateTraceTailSamplingRulesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create trace tail sampling rules default response
func (o *CreateTraceTailSamplingRulesDefault) Code() int {
	return o._statusCode
}

func (o *CreateTraceTailSamplingRulesDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-tail-sampling-rules][%d] CreateTraceTailSamplingRules default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTraceTailSamplingRulesDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/config/trace-tail-sampling-rules][%d] CreateTraceTailSamplingRules default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTraceTailSamplingRulesDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateTraceTailSamplingRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
