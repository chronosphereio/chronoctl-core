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

// DeleteGCPMetricsIntegrationReader is a Reader for the DeleteGCPMetricsIntegration structure.
type DeleteGCPMetricsIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGCPMetricsIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteGCPMetricsIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteGCPMetricsIntegrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteGCPMetricsIntegrationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteGCPMetricsIntegrationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteGCPMetricsIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteGCPMetricsIntegrationOK creates a DeleteGCPMetricsIntegrationOK with default headers values
func NewDeleteGCPMetricsIntegrationOK() *DeleteGCPMetricsIntegrationOK {
	return &DeleteGCPMetricsIntegrationOK{}
}

/*
DeleteGCPMetricsIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteGCPMetricsIntegrationOK struct {
	Payload models.ConfigunstableDeleteGCPMetricsIntegrationResponse
}

// IsSuccess returns true when this delete g c p metrics integration o k response has a 2xx status code
func (o *DeleteGCPMetricsIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete g c p metrics integration o k response has a 3xx status code
func (o *DeleteGCPMetricsIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete g c p metrics integration o k response has a 4xx status code
func (o *DeleteGCPMetricsIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete g c p metrics integration o k response has a 5xx status code
func (o *DeleteGCPMetricsIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete g c p metrics integration o k response a status code equal to that given
func (o *DeleteGCPMetricsIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete g c p metrics integration o k response
func (o *DeleteGCPMetricsIntegrationOK) Code() int {
	return 200
}

func (o *DeleteGCPMetricsIntegrationOK) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/gcp-metrics-integrations/{slug}][%d] deleteGCPMetricsIntegrationOK  %+v", 200, o.Payload)
}

func (o *DeleteGCPMetricsIntegrationOK) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/gcp-metrics-integrations/{slug}][%d] deleteGCPMetricsIntegrationOK  %+v", 200, o.Payload)
}

func (o *DeleteGCPMetricsIntegrationOK) GetPayload() models.ConfigunstableDeleteGCPMetricsIntegrationResponse {
	return o.Payload
}

func (o *DeleteGCPMetricsIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGCPMetricsIntegrationBadRequest creates a DeleteGCPMetricsIntegrationBadRequest with default headers values
func NewDeleteGCPMetricsIntegrationBadRequest() *DeleteGCPMetricsIntegrationBadRequest {
	return &DeleteGCPMetricsIntegrationBadRequest{}
}

/*
DeleteGCPMetricsIntegrationBadRequest describes a response with status code 400, with default header values.

Cannot delete the GCPMetricsIntegration because it is in use.
*/
type DeleteGCPMetricsIntegrationBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete g c p metrics integration bad request response has a 2xx status code
func (o *DeleteGCPMetricsIntegrationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete g c p metrics integration bad request response has a 3xx status code
func (o *DeleteGCPMetricsIntegrationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete g c p metrics integration bad request response has a 4xx status code
func (o *DeleteGCPMetricsIntegrationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete g c p metrics integration bad request response has a 5xx status code
func (o *DeleteGCPMetricsIntegrationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete g c p metrics integration bad request response a status code equal to that given
func (o *DeleteGCPMetricsIntegrationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete g c p metrics integration bad request response
func (o *DeleteGCPMetricsIntegrationBadRequest) Code() int {
	return 400
}

func (o *DeleteGCPMetricsIntegrationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/gcp-metrics-integrations/{slug}][%d] deleteGCPMetricsIntegrationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteGCPMetricsIntegrationBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/gcp-metrics-integrations/{slug}][%d] deleteGCPMetricsIntegrationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteGCPMetricsIntegrationBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteGCPMetricsIntegrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGCPMetricsIntegrationNotFound creates a DeleteGCPMetricsIntegrationNotFound with default headers values
func NewDeleteGCPMetricsIntegrationNotFound() *DeleteGCPMetricsIntegrationNotFound {
	return &DeleteGCPMetricsIntegrationNotFound{}
}

/*
DeleteGCPMetricsIntegrationNotFound describes a response with status code 404, with default header values.

Cannot delete the GCPMetricsIntegration because the slug does not exist.
*/
type DeleteGCPMetricsIntegrationNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete g c p metrics integration not found response has a 2xx status code
func (o *DeleteGCPMetricsIntegrationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete g c p metrics integration not found response has a 3xx status code
func (o *DeleteGCPMetricsIntegrationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete g c p metrics integration not found response has a 4xx status code
func (o *DeleteGCPMetricsIntegrationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete g c p metrics integration not found response has a 5xx status code
func (o *DeleteGCPMetricsIntegrationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete g c p metrics integration not found response a status code equal to that given
func (o *DeleteGCPMetricsIntegrationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete g c p metrics integration not found response
func (o *DeleteGCPMetricsIntegrationNotFound) Code() int {
	return 404
}

func (o *DeleteGCPMetricsIntegrationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/gcp-metrics-integrations/{slug}][%d] deleteGCPMetricsIntegrationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteGCPMetricsIntegrationNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/gcp-metrics-integrations/{slug}][%d] deleteGCPMetricsIntegrationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteGCPMetricsIntegrationNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteGCPMetricsIntegrationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGCPMetricsIntegrationInternalServerError creates a DeleteGCPMetricsIntegrationInternalServerError with default headers values
func NewDeleteGCPMetricsIntegrationInternalServerError() *DeleteGCPMetricsIntegrationInternalServerError {
	return &DeleteGCPMetricsIntegrationInternalServerError{}
}

/*
DeleteGCPMetricsIntegrationInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteGCPMetricsIntegrationInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete g c p metrics integration internal server error response has a 2xx status code
func (o *DeleteGCPMetricsIntegrationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete g c p metrics integration internal server error response has a 3xx status code
func (o *DeleteGCPMetricsIntegrationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete g c p metrics integration internal server error response has a 4xx status code
func (o *DeleteGCPMetricsIntegrationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete g c p metrics integration internal server error response has a 5xx status code
func (o *DeleteGCPMetricsIntegrationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete g c p metrics integration internal server error response a status code equal to that given
func (o *DeleteGCPMetricsIntegrationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete g c p metrics integration internal server error response
func (o *DeleteGCPMetricsIntegrationInternalServerError) Code() int {
	return 500
}

func (o *DeleteGCPMetricsIntegrationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/gcp-metrics-integrations/{slug}][%d] deleteGCPMetricsIntegrationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteGCPMetricsIntegrationInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/gcp-metrics-integrations/{slug}][%d] deleteGCPMetricsIntegrationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteGCPMetricsIntegrationInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteGCPMetricsIntegrationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGCPMetricsIntegrationDefault creates a DeleteGCPMetricsIntegrationDefault with default headers values
func NewDeleteGCPMetricsIntegrationDefault(code int) *DeleteGCPMetricsIntegrationDefault {
	return &DeleteGCPMetricsIntegrationDefault{
		_statusCode: code,
	}
}

/*
DeleteGCPMetricsIntegrationDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteGCPMetricsIntegrationDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete g c p metrics integration default response has a 2xx status code
func (o *DeleteGCPMetricsIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete g c p metrics integration default response has a 3xx status code
func (o *DeleteGCPMetricsIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete g c p metrics integration default response has a 4xx status code
func (o *DeleteGCPMetricsIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete g c p metrics integration default response has a 5xx status code
func (o *DeleteGCPMetricsIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete g c p metrics integration default response a status code equal to that given
func (o *DeleteGCPMetricsIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete g c p metrics integration default response
func (o *DeleteGCPMetricsIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteGCPMetricsIntegrationDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/gcp-metrics-integrations/{slug}][%d] DeleteGCPMetricsIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteGCPMetricsIntegrationDefault) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/gcp-metrics-integrations/{slug}][%d] DeleteGCPMetricsIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteGCPMetricsIntegrationDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteGCPMetricsIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}