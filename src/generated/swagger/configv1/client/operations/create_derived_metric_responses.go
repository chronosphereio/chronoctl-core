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

// CreateDerivedMetricReader is a Reader for the CreateDerivedMetric structure.
type CreateDerivedMetricReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDerivedMetricReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDerivedMetricOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDerivedMetricBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDerivedMetricConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDerivedMetricInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateDerivedMetricDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDerivedMetricOK creates a CreateDerivedMetricOK with default headers values
func NewCreateDerivedMetricOK() *CreateDerivedMetricOK {
	return &CreateDerivedMetricOK{}
}

/*
CreateDerivedMetricOK describes a response with status code 200, with default header values.

A successful response containing the created DerivedMetric.
*/
type CreateDerivedMetricOK struct {
	Payload *models.Configv1CreateDerivedMetricResponse
}

// IsSuccess returns true when this create derived metric o k response has a 2xx status code
func (o *CreateDerivedMetricOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create derived metric o k response has a 3xx status code
func (o *CreateDerivedMetricOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create derived metric o k response has a 4xx status code
func (o *CreateDerivedMetricOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create derived metric o k response has a 5xx status code
func (o *CreateDerivedMetricOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create derived metric o k response a status code equal to that given
func (o *CreateDerivedMetricOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create derived metric o k response
func (o *CreateDerivedMetricOK) Code() int {
	return 200
}

func (o *CreateDerivedMetricOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-metrics][%d] createDerivedMetricOK  %+v", 200, o.Payload)
}

func (o *CreateDerivedMetricOK) String() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-metrics][%d] createDerivedMetricOK  %+v", 200, o.Payload)
}

func (o *CreateDerivedMetricOK) GetPayload() *models.Configv1CreateDerivedMetricResponse {
	return o.Payload
}

func (o *CreateDerivedMetricOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1CreateDerivedMetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDerivedMetricBadRequest creates a CreateDerivedMetricBadRequest with default headers values
func NewCreateDerivedMetricBadRequest() *CreateDerivedMetricBadRequest {
	return &CreateDerivedMetricBadRequest{}
}

/*
CreateDerivedMetricBadRequest describes a response with status code 400, with default header values.

Cannot create the DerivedMetric because the request is invalid.
*/
type CreateDerivedMetricBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create derived metric bad request response has a 2xx status code
func (o *CreateDerivedMetricBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create derived metric bad request response has a 3xx status code
func (o *CreateDerivedMetricBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create derived metric bad request response has a 4xx status code
func (o *CreateDerivedMetricBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create derived metric bad request response has a 5xx status code
func (o *CreateDerivedMetricBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create derived metric bad request response a status code equal to that given
func (o *CreateDerivedMetricBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create derived metric bad request response
func (o *CreateDerivedMetricBadRequest) Code() int {
	return 400
}

func (o *CreateDerivedMetricBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-metrics][%d] createDerivedMetricBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDerivedMetricBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-metrics][%d] createDerivedMetricBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDerivedMetricBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateDerivedMetricBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDerivedMetricConflict creates a CreateDerivedMetricConflict with default headers values
func NewCreateDerivedMetricConflict() *CreateDerivedMetricConflict {
	return &CreateDerivedMetricConflict{}
}

/*
CreateDerivedMetricConflict describes a response with status code 409, with default header values.

Cannot create the DerivedMetric because there is a conflict with an existing DerivedMetric.
*/
type CreateDerivedMetricConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create derived metric conflict response has a 2xx status code
func (o *CreateDerivedMetricConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create derived metric conflict response has a 3xx status code
func (o *CreateDerivedMetricConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create derived metric conflict response has a 4xx status code
func (o *CreateDerivedMetricConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create derived metric conflict response has a 5xx status code
func (o *CreateDerivedMetricConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create derived metric conflict response a status code equal to that given
func (o *CreateDerivedMetricConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create derived metric conflict response
func (o *CreateDerivedMetricConflict) Code() int {
	return 409
}

func (o *CreateDerivedMetricConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-metrics][%d] createDerivedMetricConflict  %+v", 409, o.Payload)
}

func (o *CreateDerivedMetricConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-metrics][%d] createDerivedMetricConflict  %+v", 409, o.Payload)
}

func (o *CreateDerivedMetricConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateDerivedMetricConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDerivedMetricInternalServerError creates a CreateDerivedMetricInternalServerError with default headers values
func NewCreateDerivedMetricInternalServerError() *CreateDerivedMetricInternalServerError {
	return &CreateDerivedMetricInternalServerError{}
}

/*
CreateDerivedMetricInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateDerivedMetricInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create derived metric internal server error response has a 2xx status code
func (o *CreateDerivedMetricInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create derived metric internal server error response has a 3xx status code
func (o *CreateDerivedMetricInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create derived metric internal server error response has a 4xx status code
func (o *CreateDerivedMetricInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create derived metric internal server error response has a 5xx status code
func (o *CreateDerivedMetricInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create derived metric internal server error response a status code equal to that given
func (o *CreateDerivedMetricInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create derived metric internal server error response
func (o *CreateDerivedMetricInternalServerError) Code() int {
	return 500
}

func (o *CreateDerivedMetricInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-metrics][%d] createDerivedMetricInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDerivedMetricInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-metrics][%d] createDerivedMetricInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDerivedMetricInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateDerivedMetricInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDerivedMetricDefault creates a CreateDerivedMetricDefault with default headers values
func NewCreateDerivedMetricDefault(code int) *CreateDerivedMetricDefault {
	return &CreateDerivedMetricDefault{
		_statusCode: code,
	}
}

/*
CreateDerivedMetricDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateDerivedMetricDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create derived metric default response has a 2xx status code
func (o *CreateDerivedMetricDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create derived metric default response has a 3xx status code
func (o *CreateDerivedMetricDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create derived metric default response has a 4xx status code
func (o *CreateDerivedMetricDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create derived metric default response has a 5xx status code
func (o *CreateDerivedMetricDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create derived metric default response a status code equal to that given
func (o *CreateDerivedMetricDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create derived metric default response
func (o *CreateDerivedMetricDefault) Code() int {
	return o._statusCode
}

func (o *CreateDerivedMetricDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-metrics][%d] CreateDerivedMetric default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDerivedMetricDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-metrics][%d] CreateDerivedMetric default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDerivedMetricDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateDerivedMetricDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
