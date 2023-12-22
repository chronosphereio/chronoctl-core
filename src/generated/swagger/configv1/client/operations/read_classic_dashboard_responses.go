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

// ReadClassicDashboardReader is a Reader for the ReadClassicDashboard structure.
type ReadClassicDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadClassicDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadClassicDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadClassicDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadClassicDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadClassicDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadClassicDashboardOK creates a ReadClassicDashboardOK with default headers values
func NewReadClassicDashboardOK() *ReadClassicDashboardOK {
	return &ReadClassicDashboardOK{}
}

/*
ReadClassicDashboardOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadClassicDashboardOK struct {
	Payload *models.Configv1ReadClassicDashboardResponse
}

// IsSuccess returns true when this read classic dashboard o k response has a 2xx status code
func (o *ReadClassicDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read classic dashboard o k response has a 3xx status code
func (o *ReadClassicDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read classic dashboard o k response has a 4xx status code
func (o *ReadClassicDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read classic dashboard o k response has a 5xx status code
func (o *ReadClassicDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read classic dashboard o k response a status code equal to that given
func (o *ReadClassicDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read classic dashboard o k response
func (o *ReadClassicDashboardOK) Code() int {
	return 200
}

func (o *ReadClassicDashboardOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/classic-dashboards/{slug}][%d] readClassicDashboardOK  %+v", 200, o.Payload)
}

func (o *ReadClassicDashboardOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/classic-dashboards/{slug}][%d] readClassicDashboardOK  %+v", 200, o.Payload)
}

func (o *ReadClassicDashboardOK) GetPayload() *models.Configv1ReadClassicDashboardResponse {
	return o.Payload
}

func (o *ReadClassicDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ReadClassicDashboardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadClassicDashboardNotFound creates a ReadClassicDashboardNotFound with default headers values
func NewReadClassicDashboardNotFound() *ReadClassicDashboardNotFound {
	return &ReadClassicDashboardNotFound{}
}

/*
ReadClassicDashboardNotFound describes a response with status code 404, with default header values.

Cannot read the GrafanaDashboard because the slug does not exist.
*/
type ReadClassicDashboardNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read classic dashboard not found response has a 2xx status code
func (o *ReadClassicDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read classic dashboard not found response has a 3xx status code
func (o *ReadClassicDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read classic dashboard not found response has a 4xx status code
func (o *ReadClassicDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read classic dashboard not found response has a 5xx status code
func (o *ReadClassicDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read classic dashboard not found response a status code equal to that given
func (o *ReadClassicDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read classic dashboard not found response
func (o *ReadClassicDashboardNotFound) Code() int {
	return 404
}

func (o *ReadClassicDashboardNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/classic-dashboards/{slug}][%d] readClassicDashboardNotFound  %+v", 404, o.Payload)
}

func (o *ReadClassicDashboardNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/config/classic-dashboards/{slug}][%d] readClassicDashboardNotFound  %+v", 404, o.Payload)
}

func (o *ReadClassicDashboardNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadClassicDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadClassicDashboardInternalServerError creates a ReadClassicDashboardInternalServerError with default headers values
func NewReadClassicDashboardInternalServerError() *ReadClassicDashboardInternalServerError {
	return &ReadClassicDashboardInternalServerError{}
}

/*
ReadClassicDashboardInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadClassicDashboardInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read classic dashboard internal server error response has a 2xx status code
func (o *ReadClassicDashboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read classic dashboard internal server error response has a 3xx status code
func (o *ReadClassicDashboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read classic dashboard internal server error response has a 4xx status code
func (o *ReadClassicDashboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read classic dashboard internal server error response has a 5xx status code
func (o *ReadClassicDashboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read classic dashboard internal server error response a status code equal to that given
func (o *ReadClassicDashboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read classic dashboard internal server error response
func (o *ReadClassicDashboardInternalServerError) Code() int {
	return 500
}

func (o *ReadClassicDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/classic-dashboards/{slug}][%d] readClassicDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadClassicDashboardInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/classic-dashboards/{slug}][%d] readClassicDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadClassicDashboardInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadClassicDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadClassicDashboardDefault creates a ReadClassicDashboardDefault with default headers values
func NewReadClassicDashboardDefault(code int) *ReadClassicDashboardDefault {
	return &ReadClassicDashboardDefault{
		_statusCode: code,
	}
}

/*
ReadClassicDashboardDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadClassicDashboardDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read classic dashboard default response has a 2xx status code
func (o *ReadClassicDashboardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read classic dashboard default response has a 3xx status code
func (o *ReadClassicDashboardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read classic dashboard default response has a 4xx status code
func (o *ReadClassicDashboardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read classic dashboard default response has a 5xx status code
func (o *ReadClassicDashboardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read classic dashboard default response a status code equal to that given
func (o *ReadClassicDashboardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read classic dashboard default response
func (o *ReadClassicDashboardDefault) Code() int {
	return o._statusCode
}

func (o *ReadClassicDashboardDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/classic-dashboards/{slug}][%d] ReadClassicDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *ReadClassicDashboardDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/classic-dashboards/{slug}][%d] ReadClassicDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *ReadClassicDashboardDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadClassicDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
