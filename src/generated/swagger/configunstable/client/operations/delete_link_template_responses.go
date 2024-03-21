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

// DeleteLinkTemplateReader is a Reader for the DeleteLinkTemplate structure.
type DeleteLinkTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLinkTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLinkTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLinkTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLinkTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLinkTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteLinkTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteLinkTemplateOK creates a DeleteLinkTemplateOK with default headers values
func NewDeleteLinkTemplateOK() *DeleteLinkTemplateOK {
	return &DeleteLinkTemplateOK{}
}

/*
DeleteLinkTemplateOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteLinkTemplateOK struct {
	Payload models.ConfigunstableDeleteLinkTemplateResponse
}

// IsSuccess returns true when this delete link template o k response has a 2xx status code
func (o *DeleteLinkTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete link template o k response has a 3xx status code
func (o *DeleteLinkTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete link template o k response has a 4xx status code
func (o *DeleteLinkTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete link template o k response has a 5xx status code
func (o *DeleteLinkTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete link template o k response a status code equal to that given
func (o *DeleteLinkTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete link template o k response
func (o *DeleteLinkTemplateOK) Code() int {
	return 200
}

func (o *DeleteLinkTemplateOK) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/link-templates/{slug}][%d] deleteLinkTemplateOK  %+v", 200, o.Payload)
}

func (o *DeleteLinkTemplateOK) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/link-templates/{slug}][%d] deleteLinkTemplateOK  %+v", 200, o.Payload)
}

func (o *DeleteLinkTemplateOK) GetPayload() models.ConfigunstableDeleteLinkTemplateResponse {
	return o.Payload
}

func (o *DeleteLinkTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLinkTemplateBadRequest creates a DeleteLinkTemplateBadRequest with default headers values
func NewDeleteLinkTemplateBadRequest() *DeleteLinkTemplateBadRequest {
	return &DeleteLinkTemplateBadRequest{}
}

/*
DeleteLinkTemplateBadRequest describes a response with status code 400, with default header values.

Cannot delete the LinkTemplate because it is in use.
*/
type DeleteLinkTemplateBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete link template bad request response has a 2xx status code
func (o *DeleteLinkTemplateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete link template bad request response has a 3xx status code
func (o *DeleteLinkTemplateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete link template bad request response has a 4xx status code
func (o *DeleteLinkTemplateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete link template bad request response has a 5xx status code
func (o *DeleteLinkTemplateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete link template bad request response a status code equal to that given
func (o *DeleteLinkTemplateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete link template bad request response
func (o *DeleteLinkTemplateBadRequest) Code() int {
	return 400
}

func (o *DeleteLinkTemplateBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/link-templates/{slug}][%d] deleteLinkTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLinkTemplateBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/link-templates/{slug}][%d] deleteLinkTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLinkTemplateBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteLinkTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLinkTemplateNotFound creates a DeleteLinkTemplateNotFound with default headers values
func NewDeleteLinkTemplateNotFound() *DeleteLinkTemplateNotFound {
	return &DeleteLinkTemplateNotFound{}
}

/*
DeleteLinkTemplateNotFound describes a response with status code 404, with default header values.

Cannot delete the LinkTemplate because the slug does not exist.
*/
type DeleteLinkTemplateNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete link template not found response has a 2xx status code
func (o *DeleteLinkTemplateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete link template not found response has a 3xx status code
func (o *DeleteLinkTemplateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete link template not found response has a 4xx status code
func (o *DeleteLinkTemplateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete link template not found response has a 5xx status code
func (o *DeleteLinkTemplateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete link template not found response a status code equal to that given
func (o *DeleteLinkTemplateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete link template not found response
func (o *DeleteLinkTemplateNotFound) Code() int {
	return 404
}

func (o *DeleteLinkTemplateNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/link-templates/{slug}][%d] deleteLinkTemplateNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLinkTemplateNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/link-templates/{slug}][%d] deleteLinkTemplateNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLinkTemplateNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteLinkTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLinkTemplateInternalServerError creates a DeleteLinkTemplateInternalServerError with default headers values
func NewDeleteLinkTemplateInternalServerError() *DeleteLinkTemplateInternalServerError {
	return &DeleteLinkTemplateInternalServerError{}
}

/*
DeleteLinkTemplateInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteLinkTemplateInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete link template internal server error response has a 2xx status code
func (o *DeleteLinkTemplateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete link template internal server error response has a 3xx status code
func (o *DeleteLinkTemplateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete link template internal server error response has a 4xx status code
func (o *DeleteLinkTemplateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete link template internal server error response has a 5xx status code
func (o *DeleteLinkTemplateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete link template internal server error response a status code equal to that given
func (o *DeleteLinkTemplateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete link template internal server error response
func (o *DeleteLinkTemplateInternalServerError) Code() int {
	return 500
}

func (o *DeleteLinkTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/link-templates/{slug}][%d] deleteLinkTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLinkTemplateInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/link-templates/{slug}][%d] deleteLinkTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLinkTemplateInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteLinkTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLinkTemplateDefault creates a DeleteLinkTemplateDefault with default headers values
func NewDeleteLinkTemplateDefault(code int) *DeleteLinkTemplateDefault {
	return &DeleteLinkTemplateDefault{
		_statusCode: code,
	}
}

/*
DeleteLinkTemplateDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteLinkTemplateDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete link template default response has a 2xx status code
func (o *DeleteLinkTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete link template default response has a 3xx status code
func (o *DeleteLinkTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete link template default response has a 4xx status code
func (o *DeleteLinkTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete link template default response has a 5xx status code
func (o *DeleteLinkTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete link template default response a status code equal to that given
func (o *DeleteLinkTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete link template default response
func (o *DeleteLinkTemplateDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLinkTemplateDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/link-templates/{slug}][%d] DeleteLinkTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteLinkTemplateDefault) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/link-templates/{slug}][%d] DeleteLinkTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteLinkTemplateDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteLinkTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
