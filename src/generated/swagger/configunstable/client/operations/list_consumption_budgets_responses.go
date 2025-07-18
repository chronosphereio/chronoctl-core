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

// ListConsumptionBudgetsReader is a Reader for the ListConsumptionBudgets structure.
type ListConsumptionBudgetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListConsumptionBudgetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListConsumptionBudgetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListConsumptionBudgetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListConsumptionBudgetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListConsumptionBudgetsOK creates a ListConsumptionBudgetsOK with default headers values
func NewListConsumptionBudgetsOK() *ListConsumptionBudgetsOK {
	return &ListConsumptionBudgetsOK{}
}

/*
ListConsumptionBudgetsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListConsumptionBudgetsOK struct {
	Payload *models.ConfigunstableListConsumptionBudgetsResponse
}

// IsSuccess returns true when this list consumption budgets o k response has a 2xx status code
func (o *ListConsumptionBudgetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list consumption budgets o k response has a 3xx status code
func (o *ListConsumptionBudgetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list consumption budgets o k response has a 4xx status code
func (o *ListConsumptionBudgetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list consumption budgets o k response has a 5xx status code
func (o *ListConsumptionBudgetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list consumption budgets o k response a status code equal to that given
func (o *ListConsumptionBudgetsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list consumption budgets o k response
func (o *ListConsumptionBudgetsOK) Code() int {
	return 200
}

func (o *ListConsumptionBudgetsOK) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-budgets][%d] listConsumptionBudgetsOK  %+v", 200, o.Payload)
}

func (o *ListConsumptionBudgetsOK) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-budgets][%d] listConsumptionBudgetsOK  %+v", 200, o.Payload)
}

func (o *ListConsumptionBudgetsOK) GetPayload() *models.ConfigunstableListConsumptionBudgetsResponse {
	return o.Payload
}

func (o *ListConsumptionBudgetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableListConsumptionBudgetsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConsumptionBudgetsInternalServerError creates a ListConsumptionBudgetsInternalServerError with default headers values
func NewListConsumptionBudgetsInternalServerError() *ListConsumptionBudgetsInternalServerError {
	return &ListConsumptionBudgetsInternalServerError{}
}

/*
ListConsumptionBudgetsInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ListConsumptionBudgetsInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this list consumption budgets internal server error response has a 2xx status code
func (o *ListConsumptionBudgetsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list consumption budgets internal server error response has a 3xx status code
func (o *ListConsumptionBudgetsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list consumption budgets internal server error response has a 4xx status code
func (o *ListConsumptionBudgetsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list consumption budgets internal server error response has a 5xx status code
func (o *ListConsumptionBudgetsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list consumption budgets internal server error response a status code equal to that given
func (o *ListConsumptionBudgetsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list consumption budgets internal server error response
func (o *ListConsumptionBudgetsInternalServerError) Code() int {
	return 500
}

func (o *ListConsumptionBudgetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-budgets][%d] listConsumptionBudgetsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListConsumptionBudgetsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-budgets][%d] listConsumptionBudgetsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListConsumptionBudgetsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListConsumptionBudgetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConsumptionBudgetsDefault creates a ListConsumptionBudgetsDefault with default headers values
func NewListConsumptionBudgetsDefault(code int) *ListConsumptionBudgetsDefault {
	return &ListConsumptionBudgetsDefault{
		_statusCode: code,
	}
}

/*
ListConsumptionBudgetsDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ListConsumptionBudgetsDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this list consumption budgets default response has a 2xx status code
func (o *ListConsumptionBudgetsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list consumption budgets default response has a 3xx status code
func (o *ListConsumptionBudgetsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list consumption budgets default response has a 4xx status code
func (o *ListConsumptionBudgetsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list consumption budgets default response has a 5xx status code
func (o *ListConsumptionBudgetsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list consumption budgets default response a status code equal to that given
func (o *ListConsumptionBudgetsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list consumption budgets default response
func (o *ListConsumptionBudgetsDefault) Code() int {
	return o._statusCode
}

func (o *ListConsumptionBudgetsDefault) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-budgets][%d] ListConsumptionBudgets default  %+v", o._statusCode, o.Payload)
}

func (o *ListConsumptionBudgetsDefault) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-budgets][%d] ListConsumptionBudgets default  %+v", o._statusCode, o.Payload)
}

func (o *ListConsumptionBudgetsDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ListConsumptionBudgetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
