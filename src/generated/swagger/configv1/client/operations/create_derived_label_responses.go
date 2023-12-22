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

// CreateDerivedLabelReader is a Reader for the CreateDerivedLabel structure.
type CreateDerivedLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDerivedLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDerivedLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDerivedLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDerivedLabelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDerivedLabelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateDerivedLabelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDerivedLabelOK creates a CreateDerivedLabelOK with default headers values
func NewCreateDerivedLabelOK() *CreateDerivedLabelOK {
	return &CreateDerivedLabelOK{}
}

/*
CreateDerivedLabelOK describes a response with status code 200, with default header values.

A successful response containing the created DerivedLabel.
*/
type CreateDerivedLabelOK struct {
	Payload *models.Configv1CreateDerivedLabelResponse
}

// IsSuccess returns true when this create derived label o k response has a 2xx status code
func (o *CreateDerivedLabelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create derived label o k response has a 3xx status code
func (o *CreateDerivedLabelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create derived label o k response has a 4xx status code
func (o *CreateDerivedLabelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create derived label o k response has a 5xx status code
func (o *CreateDerivedLabelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create derived label o k response a status code equal to that given
func (o *CreateDerivedLabelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create derived label o k response
func (o *CreateDerivedLabelOK) Code() int {
	return 200
}

func (o *CreateDerivedLabelOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-labels][%d] createDerivedLabelOK  %+v", 200, o.Payload)
}

func (o *CreateDerivedLabelOK) String() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-labels][%d] createDerivedLabelOK  %+v", 200, o.Payload)
}

func (o *CreateDerivedLabelOK) GetPayload() *models.Configv1CreateDerivedLabelResponse {
	return o.Payload
}

func (o *CreateDerivedLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1CreateDerivedLabelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDerivedLabelBadRequest creates a CreateDerivedLabelBadRequest with default headers values
func NewCreateDerivedLabelBadRequest() *CreateDerivedLabelBadRequest {
	return &CreateDerivedLabelBadRequest{}
}

/*
CreateDerivedLabelBadRequest describes a response with status code 400, with default header values.

Cannot create the DerivedLabel because the request is invalid.
*/
type CreateDerivedLabelBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create derived label bad request response has a 2xx status code
func (o *CreateDerivedLabelBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create derived label bad request response has a 3xx status code
func (o *CreateDerivedLabelBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create derived label bad request response has a 4xx status code
func (o *CreateDerivedLabelBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create derived label bad request response has a 5xx status code
func (o *CreateDerivedLabelBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create derived label bad request response a status code equal to that given
func (o *CreateDerivedLabelBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create derived label bad request response
func (o *CreateDerivedLabelBadRequest) Code() int {
	return 400
}

func (o *CreateDerivedLabelBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-labels][%d] createDerivedLabelBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDerivedLabelBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-labels][%d] createDerivedLabelBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDerivedLabelBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateDerivedLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDerivedLabelConflict creates a CreateDerivedLabelConflict with default headers values
func NewCreateDerivedLabelConflict() *CreateDerivedLabelConflict {
	return &CreateDerivedLabelConflict{}
}

/*
CreateDerivedLabelConflict describes a response with status code 409, with default header values.

Cannot create the DerivedLabel because there is a conflict with an existing DerivedLabel.
*/
type CreateDerivedLabelConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create derived label conflict response has a 2xx status code
func (o *CreateDerivedLabelConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create derived label conflict response has a 3xx status code
func (o *CreateDerivedLabelConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create derived label conflict response has a 4xx status code
func (o *CreateDerivedLabelConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create derived label conflict response has a 5xx status code
func (o *CreateDerivedLabelConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create derived label conflict response a status code equal to that given
func (o *CreateDerivedLabelConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create derived label conflict response
func (o *CreateDerivedLabelConflict) Code() int {
	return 409
}

func (o *CreateDerivedLabelConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-labels][%d] createDerivedLabelConflict  %+v", 409, o.Payload)
}

func (o *CreateDerivedLabelConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-labels][%d] createDerivedLabelConflict  %+v", 409, o.Payload)
}

func (o *CreateDerivedLabelConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateDerivedLabelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDerivedLabelInternalServerError creates a CreateDerivedLabelInternalServerError with default headers values
func NewCreateDerivedLabelInternalServerError() *CreateDerivedLabelInternalServerError {
	return &CreateDerivedLabelInternalServerError{}
}

/*
CreateDerivedLabelInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateDerivedLabelInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create derived label internal server error response has a 2xx status code
func (o *CreateDerivedLabelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create derived label internal server error response has a 3xx status code
func (o *CreateDerivedLabelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create derived label internal server error response has a 4xx status code
func (o *CreateDerivedLabelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create derived label internal server error response has a 5xx status code
func (o *CreateDerivedLabelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create derived label internal server error response a status code equal to that given
func (o *CreateDerivedLabelInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create derived label internal server error response
func (o *CreateDerivedLabelInternalServerError) Code() int {
	return 500
}

func (o *CreateDerivedLabelInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-labels][%d] createDerivedLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDerivedLabelInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-labels][%d] createDerivedLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDerivedLabelInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateDerivedLabelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDerivedLabelDefault creates a CreateDerivedLabelDefault with default headers values
func NewCreateDerivedLabelDefault(code int) *CreateDerivedLabelDefault {
	return &CreateDerivedLabelDefault{
		_statusCode: code,
	}
}

/*
CreateDerivedLabelDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateDerivedLabelDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create derived label default response has a 2xx status code
func (o *CreateDerivedLabelDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create derived label default response has a 3xx status code
func (o *CreateDerivedLabelDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create derived label default response has a 4xx status code
func (o *CreateDerivedLabelDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create derived label default response has a 5xx status code
func (o *CreateDerivedLabelDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create derived label default response a status code equal to that given
func (o *CreateDerivedLabelDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create derived label default response
func (o *CreateDerivedLabelDefault) Code() int {
	return o._statusCode
}

func (o *CreateDerivedLabelDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-labels][%d] CreateDerivedLabel default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDerivedLabelDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/config/derived-labels][%d] CreateDerivedLabel default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDerivedLabelDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateDerivedLabelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
