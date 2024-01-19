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

// CreateLinkTemplateReader is a Reader for the CreateLinkTemplate structure.
type CreateLinkTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLinkTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateLinkTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateLinkTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateLinkTemplateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateLinkTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateLinkTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateLinkTemplateOK creates a CreateLinkTemplateOK with default headers values
func NewCreateLinkTemplateOK() *CreateLinkTemplateOK {
	return &CreateLinkTemplateOK{}
}

/*
CreateLinkTemplateOK describes a response with status code 200, with default header values.

A successful response containing the created LinkTemplate.
*/
type CreateLinkTemplateOK struct {
	Payload *models.ConfigunstableCreateLinkTemplateResponse
}

// IsSuccess returns true when this create link template o k response has a 2xx status code
func (o *CreateLinkTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create link template o k response has a 3xx status code
func (o *CreateLinkTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create link template o k response has a 4xx status code
func (o *CreateLinkTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create link template o k response has a 5xx status code
func (o *CreateLinkTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create link template o k response a status code equal to that given
func (o *CreateLinkTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create link template o k response
func (o *CreateLinkTemplateOK) Code() int {
	return 200
}

func (o *CreateLinkTemplateOK) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/link-templates][%d] createLinkTemplateOK  %+v", 200, o.Payload)
}

func (o *CreateLinkTemplateOK) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/link-templates][%d] createLinkTemplateOK  %+v", 200, o.Payload)
}

func (o *CreateLinkTemplateOK) GetPayload() *models.ConfigunstableCreateLinkTemplateResponse {
	return o.Payload
}

func (o *CreateLinkTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableCreateLinkTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLinkTemplateBadRequest creates a CreateLinkTemplateBadRequest with default headers values
func NewCreateLinkTemplateBadRequest() *CreateLinkTemplateBadRequest {
	return &CreateLinkTemplateBadRequest{}
}

/*
CreateLinkTemplateBadRequest describes a response with status code 400, with default header values.

Cannot create the LinkTemplate because the request is invalid.
*/
type CreateLinkTemplateBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create link template bad request response has a 2xx status code
func (o *CreateLinkTemplateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create link template bad request response has a 3xx status code
func (o *CreateLinkTemplateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create link template bad request response has a 4xx status code
func (o *CreateLinkTemplateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create link template bad request response has a 5xx status code
func (o *CreateLinkTemplateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create link template bad request response a status code equal to that given
func (o *CreateLinkTemplateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create link template bad request response
func (o *CreateLinkTemplateBadRequest) Code() int {
	return 400
}

func (o *CreateLinkTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/link-templates][%d] createLinkTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLinkTemplateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/link-templates][%d] createLinkTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLinkTemplateBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateLinkTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLinkTemplateConflict creates a CreateLinkTemplateConflict with default headers values
func NewCreateLinkTemplateConflict() *CreateLinkTemplateConflict {
	return &CreateLinkTemplateConflict{}
}

/*
CreateLinkTemplateConflict describes a response with status code 409, with default header values.

Cannot create the LinkTemplate because there is a conflict with an existing LinkTemplate.
*/
type CreateLinkTemplateConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create link template conflict response has a 2xx status code
func (o *CreateLinkTemplateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create link template conflict response has a 3xx status code
func (o *CreateLinkTemplateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create link template conflict response has a 4xx status code
func (o *CreateLinkTemplateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create link template conflict response has a 5xx status code
func (o *CreateLinkTemplateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create link template conflict response a status code equal to that given
func (o *CreateLinkTemplateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create link template conflict response
func (o *CreateLinkTemplateConflict) Code() int {
	return 409
}

func (o *CreateLinkTemplateConflict) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/link-templates][%d] createLinkTemplateConflict  %+v", 409, o.Payload)
}

func (o *CreateLinkTemplateConflict) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/link-templates][%d] createLinkTemplateConflict  %+v", 409, o.Payload)
}

func (o *CreateLinkTemplateConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateLinkTemplateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLinkTemplateInternalServerError creates a CreateLinkTemplateInternalServerError with default headers values
func NewCreateLinkTemplateInternalServerError() *CreateLinkTemplateInternalServerError {
	return &CreateLinkTemplateInternalServerError{}
}

/*
CreateLinkTemplateInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateLinkTemplateInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create link template internal server error response has a 2xx status code
func (o *CreateLinkTemplateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create link template internal server error response has a 3xx status code
func (o *CreateLinkTemplateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create link template internal server error response has a 4xx status code
func (o *CreateLinkTemplateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create link template internal server error response has a 5xx status code
func (o *CreateLinkTemplateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create link template internal server error response a status code equal to that given
func (o *CreateLinkTemplateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create link template internal server error response
func (o *CreateLinkTemplateInternalServerError) Code() int {
	return 500
}

func (o *CreateLinkTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/link-templates][%d] createLinkTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateLinkTemplateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/link-templates][%d] createLinkTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateLinkTemplateInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateLinkTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLinkTemplateDefault creates a CreateLinkTemplateDefault with default headers values
func NewCreateLinkTemplateDefault(code int) *CreateLinkTemplateDefault {
	return &CreateLinkTemplateDefault{
		_statusCode: code,
	}
}

/*
CreateLinkTemplateDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateLinkTemplateDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create link template default response has a 2xx status code
func (o *CreateLinkTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create link template default response has a 3xx status code
func (o *CreateLinkTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create link template default response has a 4xx status code
func (o *CreateLinkTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create link template default response has a 5xx status code
func (o *CreateLinkTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create link template default response a status code equal to that given
func (o *CreateLinkTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create link template default response
func (o *CreateLinkTemplateDefault) Code() int {
	return o._statusCode
}

func (o *CreateLinkTemplateDefault) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/link-templates][%d] CreateLinkTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *CreateLinkTemplateDefault) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/link-templates][%d] CreateLinkTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *CreateLinkTemplateDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateLinkTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}