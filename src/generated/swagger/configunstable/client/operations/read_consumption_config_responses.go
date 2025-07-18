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

// ReadConsumptionConfigReader is a Reader for the ReadConsumptionConfig structure.
type ReadConsumptionConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadConsumptionConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadConsumptionConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadConsumptionConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadConsumptionConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadConsumptionConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadConsumptionConfigOK creates a ReadConsumptionConfigOK with default headers values
func NewReadConsumptionConfigOK() *ReadConsumptionConfigOK {
	return &ReadConsumptionConfigOK{}
}

/*
ReadConsumptionConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadConsumptionConfigOK struct {
	Payload *models.ConfigunstableReadConsumptionConfigResponse
}

// IsSuccess returns true when this read consumption config o k response has a 2xx status code
func (o *ReadConsumptionConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read consumption config o k response has a 3xx status code
func (o *ReadConsumptionConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read consumption config o k response has a 4xx status code
func (o *ReadConsumptionConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read consumption config o k response has a 5xx status code
func (o *ReadConsumptionConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read consumption config o k response a status code equal to that given
func (o *ReadConsumptionConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read consumption config o k response
func (o *ReadConsumptionConfigOK) Code() int {
	return 200
}

func (o *ReadConsumptionConfigOK) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-config][%d] readConsumptionConfigOK  %+v", 200, o.Payload)
}

func (o *ReadConsumptionConfigOK) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-config][%d] readConsumptionConfigOK  %+v", 200, o.Payload)
}

func (o *ReadConsumptionConfigOK) GetPayload() *models.ConfigunstableReadConsumptionConfigResponse {
	return o.Payload
}

func (o *ReadConsumptionConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableReadConsumptionConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadConsumptionConfigNotFound creates a ReadConsumptionConfigNotFound with default headers values
func NewReadConsumptionConfigNotFound() *ReadConsumptionConfigNotFound {
	return &ReadConsumptionConfigNotFound{}
}

/*
ReadConsumptionConfigNotFound describes a response with status code 404, with default header values.

Cannot read the ConsumptionConfig because ConsumptionConfig has not been created.
*/
type ReadConsumptionConfigNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read consumption config not found response has a 2xx status code
func (o *ReadConsumptionConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read consumption config not found response has a 3xx status code
func (o *ReadConsumptionConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read consumption config not found response has a 4xx status code
func (o *ReadConsumptionConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read consumption config not found response has a 5xx status code
func (o *ReadConsumptionConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read consumption config not found response a status code equal to that given
func (o *ReadConsumptionConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read consumption config not found response
func (o *ReadConsumptionConfigNotFound) Code() int {
	return 404
}

func (o *ReadConsumptionConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-config][%d] readConsumptionConfigNotFound  %+v", 404, o.Payload)
}

func (o *ReadConsumptionConfigNotFound) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-config][%d] readConsumptionConfigNotFound  %+v", 404, o.Payload)
}

func (o *ReadConsumptionConfigNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadConsumptionConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadConsumptionConfigInternalServerError creates a ReadConsumptionConfigInternalServerError with default headers values
func NewReadConsumptionConfigInternalServerError() *ReadConsumptionConfigInternalServerError {
	return &ReadConsumptionConfigInternalServerError{}
}

/*
ReadConsumptionConfigInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadConsumptionConfigInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read consumption config internal server error response has a 2xx status code
func (o *ReadConsumptionConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read consumption config internal server error response has a 3xx status code
func (o *ReadConsumptionConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read consumption config internal server error response has a 4xx status code
func (o *ReadConsumptionConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read consumption config internal server error response has a 5xx status code
func (o *ReadConsumptionConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read consumption config internal server error response a status code equal to that given
func (o *ReadConsumptionConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read consumption config internal server error response
func (o *ReadConsumptionConfigInternalServerError) Code() int {
	return 500
}

func (o *ReadConsumptionConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-config][%d] readConsumptionConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadConsumptionConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-config][%d] readConsumptionConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadConsumptionConfigInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadConsumptionConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadConsumptionConfigDefault creates a ReadConsumptionConfigDefault with default headers values
func NewReadConsumptionConfigDefault(code int) *ReadConsumptionConfigDefault {
	return &ReadConsumptionConfigDefault{
		_statusCode: code,
	}
}

/*
ReadConsumptionConfigDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadConsumptionConfigDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read consumption config default response has a 2xx status code
func (o *ReadConsumptionConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read consumption config default response has a 3xx status code
func (o *ReadConsumptionConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read consumption config default response has a 4xx status code
func (o *ReadConsumptionConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read consumption config default response has a 5xx status code
func (o *ReadConsumptionConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read consumption config default response a status code equal to that given
func (o *ReadConsumptionConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read consumption config default response
func (o *ReadConsumptionConfigDefault) Code() int {
	return o._statusCode
}

func (o *ReadConsumptionConfigDefault) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-config][%d] ReadConsumptionConfig default  %+v", o._statusCode, o.Payload)
}

func (o *ReadConsumptionConfigDefault) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/consumption-config][%d] ReadConsumptionConfig default  %+v", o._statusCode, o.Payload)
}

func (o *ReadConsumptionConfigDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadConsumptionConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
