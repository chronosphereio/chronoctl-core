// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
)

// UpdateGrafanaDashboardReader is a Reader for the UpdateGrafanaDashboard structure.
type UpdateGrafanaDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGrafanaDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGrafanaDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateGrafanaDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateGrafanaDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateGrafanaDashboardConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateGrafanaDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateGrafanaDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateGrafanaDashboardOK creates a UpdateGrafanaDashboardOK with default headers values
func NewUpdateGrafanaDashboardOK() *UpdateGrafanaDashboardOK {
	return &UpdateGrafanaDashboardOK{}
}

/*
UpdateGrafanaDashboardOK describes a response with status code 200, with default header values.

A successful response containing the updated GrafanaDashboard.
*/
type UpdateGrafanaDashboardOK struct {
	Payload *models.Configv1UpdateGrafanaDashboardResponse
}

// IsSuccess returns true when this update grafana dashboard o k response has a 2xx status code
func (o *UpdateGrafanaDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update grafana dashboard o k response has a 3xx status code
func (o *UpdateGrafanaDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update grafana dashboard o k response has a 4xx status code
func (o *UpdateGrafanaDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update grafana dashboard o k response has a 5xx status code
func (o *UpdateGrafanaDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update grafana dashboard o k response a status code equal to that given
func (o *UpdateGrafanaDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update grafana dashboard o k response
func (o *UpdateGrafanaDashboardOK) Code() int {
	return 200
}

func (o *UpdateGrafanaDashboardOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/grafana-dashboards/{slug}][%d] updateGrafanaDashboardOK  %+v", 200, o.Payload)
}

func (o *UpdateGrafanaDashboardOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/grafana-dashboards/{slug}][%d] updateGrafanaDashboardOK  %+v", 200, o.Payload)
}

func (o *UpdateGrafanaDashboardOK) GetPayload() *models.Configv1UpdateGrafanaDashboardResponse {
	return o.Payload
}

func (o *UpdateGrafanaDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1UpdateGrafanaDashboardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGrafanaDashboardBadRequest creates a UpdateGrafanaDashboardBadRequest with default headers values
func NewUpdateGrafanaDashboardBadRequest() *UpdateGrafanaDashboardBadRequest {
	return &UpdateGrafanaDashboardBadRequest{}
}

/*
UpdateGrafanaDashboardBadRequest describes a response with status code 400, with default header values.

Cannot update the GrafanaDashboard because the request is invalid.
*/
type UpdateGrafanaDashboardBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update grafana dashboard bad request response has a 2xx status code
func (o *UpdateGrafanaDashboardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update grafana dashboard bad request response has a 3xx status code
func (o *UpdateGrafanaDashboardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update grafana dashboard bad request response has a 4xx status code
func (o *UpdateGrafanaDashboardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update grafana dashboard bad request response has a 5xx status code
func (o *UpdateGrafanaDashboardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update grafana dashboard bad request response a status code equal to that given
func (o *UpdateGrafanaDashboardBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update grafana dashboard bad request response
func (o *UpdateGrafanaDashboardBadRequest) Code() int {
	return 400
}

func (o *UpdateGrafanaDashboardBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/grafana-dashboards/{slug}][%d] updateGrafanaDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateGrafanaDashboardBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/grafana-dashboards/{slug}][%d] updateGrafanaDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateGrafanaDashboardBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateGrafanaDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGrafanaDashboardNotFound creates a UpdateGrafanaDashboardNotFound with default headers values
func NewUpdateGrafanaDashboardNotFound() *UpdateGrafanaDashboardNotFound {
	return &UpdateGrafanaDashboardNotFound{}
}

/*
UpdateGrafanaDashboardNotFound describes a response with status code 404, with default header values.

Cannot update the GrafanaDashboard because the slug does not exist.
*/
type UpdateGrafanaDashboardNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update grafana dashboard not found response has a 2xx status code
func (o *UpdateGrafanaDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update grafana dashboard not found response has a 3xx status code
func (o *UpdateGrafanaDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update grafana dashboard not found response has a 4xx status code
func (o *UpdateGrafanaDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update grafana dashboard not found response has a 5xx status code
func (o *UpdateGrafanaDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update grafana dashboard not found response a status code equal to that given
func (o *UpdateGrafanaDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update grafana dashboard not found response
func (o *UpdateGrafanaDashboardNotFound) Code() int {
	return 404
}

func (o *UpdateGrafanaDashboardNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/grafana-dashboards/{slug}][%d] updateGrafanaDashboardNotFound  %+v", 404, o.Payload)
}

func (o *UpdateGrafanaDashboardNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/grafana-dashboards/{slug}][%d] updateGrafanaDashboardNotFound  %+v", 404, o.Payload)
}

func (o *UpdateGrafanaDashboardNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateGrafanaDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGrafanaDashboardConflict creates a UpdateGrafanaDashboardConflict with default headers values
func NewUpdateGrafanaDashboardConflict() *UpdateGrafanaDashboardConflict {
	return &UpdateGrafanaDashboardConflict{}
}

/*
UpdateGrafanaDashboardConflict describes a response with status code 409, with default header values.

Cannot update the GrafanaDashboard because there is a conflict with an existing GrafanaDashboard.
*/
type UpdateGrafanaDashboardConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update grafana dashboard conflict response has a 2xx status code
func (o *UpdateGrafanaDashboardConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update grafana dashboard conflict response has a 3xx status code
func (o *UpdateGrafanaDashboardConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update grafana dashboard conflict response has a 4xx status code
func (o *UpdateGrafanaDashboardConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update grafana dashboard conflict response has a 5xx status code
func (o *UpdateGrafanaDashboardConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update grafana dashboard conflict response a status code equal to that given
func (o *UpdateGrafanaDashboardConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update grafana dashboard conflict response
func (o *UpdateGrafanaDashboardConflict) Code() int {
	return 409
}

func (o *UpdateGrafanaDashboardConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/grafana-dashboards/{slug}][%d] updateGrafanaDashboardConflict  %+v", 409, o.Payload)
}

func (o *UpdateGrafanaDashboardConflict) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/grafana-dashboards/{slug}][%d] updateGrafanaDashboardConflict  %+v", 409, o.Payload)
}

func (o *UpdateGrafanaDashboardConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateGrafanaDashboardConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGrafanaDashboardInternalServerError creates a UpdateGrafanaDashboardInternalServerError with default headers values
func NewUpdateGrafanaDashboardInternalServerError() *UpdateGrafanaDashboardInternalServerError {
	return &UpdateGrafanaDashboardInternalServerError{}
}

/*
UpdateGrafanaDashboardInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateGrafanaDashboardInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update grafana dashboard internal server error response has a 2xx status code
func (o *UpdateGrafanaDashboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update grafana dashboard internal server error response has a 3xx status code
func (o *UpdateGrafanaDashboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update grafana dashboard internal server error response has a 4xx status code
func (o *UpdateGrafanaDashboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update grafana dashboard internal server error response has a 5xx status code
func (o *UpdateGrafanaDashboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update grafana dashboard internal server error response a status code equal to that given
func (o *UpdateGrafanaDashboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update grafana dashboard internal server error response
func (o *UpdateGrafanaDashboardInternalServerError) Code() int {
	return 500
}

func (o *UpdateGrafanaDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/grafana-dashboards/{slug}][%d] updateGrafanaDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateGrafanaDashboardInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/grafana-dashboards/{slug}][%d] updateGrafanaDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateGrafanaDashboardInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateGrafanaDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGrafanaDashboardDefault creates a UpdateGrafanaDashboardDefault with default headers values
func NewUpdateGrafanaDashboardDefault(code int) *UpdateGrafanaDashboardDefault {
	return &UpdateGrafanaDashboardDefault{
		_statusCode: code,
	}
}

/*
UpdateGrafanaDashboardDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateGrafanaDashboardDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update grafana dashboard default response has a 2xx status code
func (o *UpdateGrafanaDashboardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update grafana dashboard default response has a 3xx status code
func (o *UpdateGrafanaDashboardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update grafana dashboard default response has a 4xx status code
func (o *UpdateGrafanaDashboardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update grafana dashboard default response has a 5xx status code
func (o *UpdateGrafanaDashboardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update grafana dashboard default response a status code equal to that given
func (o *UpdateGrafanaDashboardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update grafana dashboard default response
func (o *UpdateGrafanaDashboardDefault) Code() int {
	return o._statusCode
}

func (o *UpdateGrafanaDashboardDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/grafana-dashboards/{slug}][%d] UpdateGrafanaDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateGrafanaDashboardDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/grafana-dashboards/{slug}][%d] UpdateGrafanaDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateGrafanaDashboardDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateGrafanaDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateGrafanaDashboardBody update grafana dashboard body
swagger:model UpdateGrafanaDashboardBody
*/
type UpdateGrafanaDashboardBody struct {

	// If true, the GrafanaDashboard will be created if it does not already exist, identified by slug. If false, an error will be returned if the GrafanaDashboard does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the GrafanaDashboard will not be created nor updated, and no response GrafanaDashboard will be returned. The response will return an error if the given GrafanaDashboard is invalid.
	DryRun bool `json:"dry_run,omitempty"`

	// grafana dashboard
	GrafanaDashboard *models.Configv1GrafanaDashboard `json:"grafana_dashboard,omitempty"`
}

// Validate validates this update grafana dashboard body
func (o *UpdateGrafanaDashboardBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGrafanaDashboard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateGrafanaDashboardBody) validateGrafanaDashboard(formats strfmt.Registry) error {
	if swag.IsZero(o.GrafanaDashboard) { // not required
		return nil
	}

	if o.GrafanaDashboard != nil {
		if err := o.GrafanaDashboard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "grafana_dashboard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "grafana_dashboard")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update grafana dashboard body based on the context it is used
func (o *UpdateGrafanaDashboardBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateGrafanaDashboard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateGrafanaDashboardBody) contextValidateGrafanaDashboard(ctx context.Context, formats strfmt.Registry) error {

	if o.GrafanaDashboard != nil {

		if swag.IsZero(o.GrafanaDashboard) { // not required
			return nil
		}

		if err := o.GrafanaDashboard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "grafana_dashboard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "grafana_dashboard")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateGrafanaDashboardBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateGrafanaDashboardBody) UnmarshalBinary(b []byte) error {
	var res UpdateGrafanaDashboardBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}