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

// DeleteDashboardReader is a Reader for the DeleteDashboard structure.
type DeleteDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDashboardOK creates a DeleteDashboardOK with default headers values
func NewDeleteDashboardOK() *DeleteDashboardOK {
	return &DeleteDashboardOK{}
}

/*
DeleteDashboardOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteDashboardOK struct {
	Payload models.Configv1DeleteDashboardResponse
}

// IsSuccess returns true when this delete dashboard o k response has a 2xx status code
func (o *DeleteDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete dashboard o k response has a 3xx status code
func (o *DeleteDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard o k response has a 4xx status code
func (o *DeleteDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete dashboard o k response has a 5xx status code
func (o *DeleteDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete dashboard o k response a status code equal to that given
func (o *DeleteDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete dashboard o k response
func (o *DeleteDashboardOK) Code() int {
	return 200
}

func (o *DeleteDashboardOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/dashboards/{slug}][%d] deleteDashboardOK  %+v", 200, o.Payload)
}

func (o *DeleteDashboardOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/dashboards/{slug}][%d] deleteDashboardOK  %+v", 200, o.Payload)
}

func (o *DeleteDashboardOK) GetPayload() models.Configv1DeleteDashboardResponse {
	return o.Payload
}

func (o *DeleteDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardBadRequest creates a DeleteDashboardBadRequest with default headers values
func NewDeleteDashboardBadRequest() *DeleteDashboardBadRequest {
	return &DeleteDashboardBadRequest{}
}

/*
DeleteDashboardBadRequest describes a response with status code 400, with default header values.

Cannot delete the Dashboard because it is in use.
*/
type DeleteDashboardBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete dashboard bad request response has a 2xx status code
func (o *DeleteDashboardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete dashboard bad request response has a 3xx status code
func (o *DeleteDashboardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard bad request response has a 4xx status code
func (o *DeleteDashboardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete dashboard bad request response has a 5xx status code
func (o *DeleteDashboardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete dashboard bad request response a status code equal to that given
func (o *DeleteDashboardBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete dashboard bad request response
func (o *DeleteDashboardBadRequest) Code() int {
	return 400
}

func (o *DeleteDashboardBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/dashboards/{slug}][%d] deleteDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDashboardBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/dashboards/{slug}][%d] deleteDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDashboardBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardNotFound creates a DeleteDashboardNotFound with default headers values
func NewDeleteDashboardNotFound() *DeleteDashboardNotFound {
	return &DeleteDashboardNotFound{}
}

/*
DeleteDashboardNotFound describes a response with status code 404, with default header values.

Cannot delete the Dashboard because the slug does not exist.
*/
type DeleteDashboardNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete dashboard not found response has a 2xx status code
func (o *DeleteDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete dashboard not found response has a 3xx status code
func (o *DeleteDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard not found response has a 4xx status code
func (o *DeleteDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete dashboard not found response has a 5xx status code
func (o *DeleteDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete dashboard not found response a status code equal to that given
func (o *DeleteDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete dashboard not found response
func (o *DeleteDashboardNotFound) Code() int {
	return 404
}

func (o *DeleteDashboardNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/dashboards/{slug}][%d] deleteDashboardNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDashboardNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/dashboards/{slug}][%d] deleteDashboardNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDashboardNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardInternalServerError creates a DeleteDashboardInternalServerError with default headers values
func NewDeleteDashboardInternalServerError() *DeleteDashboardInternalServerError {
	return &DeleteDashboardInternalServerError{}
}

/*
DeleteDashboardInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteDashboardInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete dashboard internal server error response has a 2xx status code
func (o *DeleteDashboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete dashboard internal server error response has a 3xx status code
func (o *DeleteDashboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard internal server error response has a 4xx status code
func (o *DeleteDashboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete dashboard internal server error response has a 5xx status code
func (o *DeleteDashboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete dashboard internal server error response a status code equal to that given
func (o *DeleteDashboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete dashboard internal server error response
func (o *DeleteDashboardInternalServerError) Code() int {
	return 500
}

func (o *DeleteDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/dashboards/{slug}][%d] deleteDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDashboardInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/dashboards/{slug}][%d] deleteDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDashboardInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardDefault creates a DeleteDashboardDefault with default headers values
func NewDeleteDashboardDefault(code int) *DeleteDashboardDefault {
	return &DeleteDashboardDefault{
		_statusCode: code,
	}
}

/*
DeleteDashboardDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteDashboardDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete dashboard default response has a 2xx status code
func (o *DeleteDashboardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete dashboard default response has a 3xx status code
func (o *DeleteDashboardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete dashboard default response has a 4xx status code
func (o *DeleteDashboardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete dashboard default response has a 5xx status code
func (o *DeleteDashboardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete dashboard default response a status code equal to that given
func (o *DeleteDashboardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete dashboard default response
func (o *DeleteDashboardDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDashboardDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/dashboards/{slug}][%d] DeleteDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDashboardDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/dashboards/{slug}][%d] DeleteDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDashboardDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
