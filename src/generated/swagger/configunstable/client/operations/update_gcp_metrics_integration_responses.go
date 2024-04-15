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

	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/models"
)

// UpdateGcpMetricsIntegrationReader is a Reader for the UpdateGcpMetricsIntegration structure.
type UpdateGcpMetricsIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGcpMetricsIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGcpMetricsIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateGcpMetricsIntegrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateGcpMetricsIntegrationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateGcpMetricsIntegrationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateGcpMetricsIntegrationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateGcpMetricsIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateGcpMetricsIntegrationOK creates a UpdateGcpMetricsIntegrationOK with default headers values
func NewUpdateGcpMetricsIntegrationOK() *UpdateGcpMetricsIntegrationOK {
	return &UpdateGcpMetricsIntegrationOK{}
}

/*
UpdateGcpMetricsIntegrationOK describes a response with status code 200, with default header values.

A successful response containing the updated GcpMetricsIntegration.
*/
type UpdateGcpMetricsIntegrationOK struct {
	Payload *models.ConfigunstableUpdateGcpMetricsIntegrationResponse
}

// IsSuccess returns true when this update gcp metrics integration o k response has a 2xx status code
func (o *UpdateGcpMetricsIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update gcp metrics integration o k response has a 3xx status code
func (o *UpdateGcpMetricsIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update gcp metrics integration o k response has a 4xx status code
func (o *UpdateGcpMetricsIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update gcp metrics integration o k response has a 5xx status code
func (o *UpdateGcpMetricsIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update gcp metrics integration o k response a status code equal to that given
func (o *UpdateGcpMetricsIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update gcp metrics integration o k response
func (o *UpdateGcpMetricsIntegrationOK) Code() int {
	return 200
}

func (o *UpdateGcpMetricsIntegrationOK) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/gcp-metrics-integrations/{slug}][%d] updateGcpMetricsIntegrationOK  %+v", 200, o.Payload)
}

func (o *UpdateGcpMetricsIntegrationOK) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/gcp-metrics-integrations/{slug}][%d] updateGcpMetricsIntegrationOK  %+v", 200, o.Payload)
}

func (o *UpdateGcpMetricsIntegrationOK) GetPayload() *models.ConfigunstableUpdateGcpMetricsIntegrationResponse {
	return o.Payload
}

func (o *UpdateGcpMetricsIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableUpdateGcpMetricsIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGcpMetricsIntegrationBadRequest creates a UpdateGcpMetricsIntegrationBadRequest with default headers values
func NewUpdateGcpMetricsIntegrationBadRequest() *UpdateGcpMetricsIntegrationBadRequest {
	return &UpdateGcpMetricsIntegrationBadRequest{}
}

/*
UpdateGcpMetricsIntegrationBadRequest describes a response with status code 400, with default header values.

Cannot update the GcpMetricsIntegration because the request is invalid.
*/
type UpdateGcpMetricsIntegrationBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update gcp metrics integration bad request response has a 2xx status code
func (o *UpdateGcpMetricsIntegrationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update gcp metrics integration bad request response has a 3xx status code
func (o *UpdateGcpMetricsIntegrationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update gcp metrics integration bad request response has a 4xx status code
func (o *UpdateGcpMetricsIntegrationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update gcp metrics integration bad request response has a 5xx status code
func (o *UpdateGcpMetricsIntegrationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update gcp metrics integration bad request response a status code equal to that given
func (o *UpdateGcpMetricsIntegrationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update gcp metrics integration bad request response
func (o *UpdateGcpMetricsIntegrationBadRequest) Code() int {
	return 400
}

func (o *UpdateGcpMetricsIntegrationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/gcp-metrics-integrations/{slug}][%d] updateGcpMetricsIntegrationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateGcpMetricsIntegrationBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/gcp-metrics-integrations/{slug}][%d] updateGcpMetricsIntegrationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateGcpMetricsIntegrationBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateGcpMetricsIntegrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGcpMetricsIntegrationNotFound creates a UpdateGcpMetricsIntegrationNotFound with default headers values
func NewUpdateGcpMetricsIntegrationNotFound() *UpdateGcpMetricsIntegrationNotFound {
	return &UpdateGcpMetricsIntegrationNotFound{}
}

/*
UpdateGcpMetricsIntegrationNotFound describes a response with status code 404, with default header values.

Cannot update the GcpMetricsIntegration because the slug does not exist.
*/
type UpdateGcpMetricsIntegrationNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update gcp metrics integration not found response has a 2xx status code
func (o *UpdateGcpMetricsIntegrationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update gcp metrics integration not found response has a 3xx status code
func (o *UpdateGcpMetricsIntegrationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update gcp metrics integration not found response has a 4xx status code
func (o *UpdateGcpMetricsIntegrationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update gcp metrics integration not found response has a 5xx status code
func (o *UpdateGcpMetricsIntegrationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update gcp metrics integration not found response a status code equal to that given
func (o *UpdateGcpMetricsIntegrationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update gcp metrics integration not found response
func (o *UpdateGcpMetricsIntegrationNotFound) Code() int {
	return 404
}

func (o *UpdateGcpMetricsIntegrationNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/gcp-metrics-integrations/{slug}][%d] updateGcpMetricsIntegrationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateGcpMetricsIntegrationNotFound) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/gcp-metrics-integrations/{slug}][%d] updateGcpMetricsIntegrationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateGcpMetricsIntegrationNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateGcpMetricsIntegrationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGcpMetricsIntegrationConflict creates a UpdateGcpMetricsIntegrationConflict with default headers values
func NewUpdateGcpMetricsIntegrationConflict() *UpdateGcpMetricsIntegrationConflict {
	return &UpdateGcpMetricsIntegrationConflict{}
}

/*
UpdateGcpMetricsIntegrationConflict describes a response with status code 409, with default header values.

Cannot update the GcpMetricsIntegration because there is a conflict with an existing GcpMetricsIntegration.
*/
type UpdateGcpMetricsIntegrationConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update gcp metrics integration conflict response has a 2xx status code
func (o *UpdateGcpMetricsIntegrationConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update gcp metrics integration conflict response has a 3xx status code
func (o *UpdateGcpMetricsIntegrationConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update gcp metrics integration conflict response has a 4xx status code
func (o *UpdateGcpMetricsIntegrationConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update gcp metrics integration conflict response has a 5xx status code
func (o *UpdateGcpMetricsIntegrationConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update gcp metrics integration conflict response a status code equal to that given
func (o *UpdateGcpMetricsIntegrationConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update gcp metrics integration conflict response
func (o *UpdateGcpMetricsIntegrationConflict) Code() int {
	return 409
}

func (o *UpdateGcpMetricsIntegrationConflict) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/gcp-metrics-integrations/{slug}][%d] updateGcpMetricsIntegrationConflict  %+v", 409, o.Payload)
}

func (o *UpdateGcpMetricsIntegrationConflict) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/gcp-metrics-integrations/{slug}][%d] updateGcpMetricsIntegrationConflict  %+v", 409, o.Payload)
}

func (o *UpdateGcpMetricsIntegrationConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateGcpMetricsIntegrationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGcpMetricsIntegrationInternalServerError creates a UpdateGcpMetricsIntegrationInternalServerError with default headers values
func NewUpdateGcpMetricsIntegrationInternalServerError() *UpdateGcpMetricsIntegrationInternalServerError {
	return &UpdateGcpMetricsIntegrationInternalServerError{}
}

/*
UpdateGcpMetricsIntegrationInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateGcpMetricsIntegrationInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update gcp metrics integration internal server error response has a 2xx status code
func (o *UpdateGcpMetricsIntegrationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update gcp metrics integration internal server error response has a 3xx status code
func (o *UpdateGcpMetricsIntegrationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update gcp metrics integration internal server error response has a 4xx status code
func (o *UpdateGcpMetricsIntegrationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update gcp metrics integration internal server error response has a 5xx status code
func (o *UpdateGcpMetricsIntegrationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update gcp metrics integration internal server error response a status code equal to that given
func (o *UpdateGcpMetricsIntegrationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update gcp metrics integration internal server error response
func (o *UpdateGcpMetricsIntegrationInternalServerError) Code() int {
	return 500
}

func (o *UpdateGcpMetricsIntegrationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/gcp-metrics-integrations/{slug}][%d] updateGcpMetricsIntegrationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateGcpMetricsIntegrationInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/gcp-metrics-integrations/{slug}][%d] updateGcpMetricsIntegrationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateGcpMetricsIntegrationInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateGcpMetricsIntegrationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGcpMetricsIntegrationDefault creates a UpdateGcpMetricsIntegrationDefault with default headers values
func NewUpdateGcpMetricsIntegrationDefault(code int) *UpdateGcpMetricsIntegrationDefault {
	return &UpdateGcpMetricsIntegrationDefault{
		_statusCode: code,
	}
}

/*
UpdateGcpMetricsIntegrationDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateGcpMetricsIntegrationDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update gcp metrics integration default response has a 2xx status code
func (o *UpdateGcpMetricsIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update gcp metrics integration default response has a 3xx status code
func (o *UpdateGcpMetricsIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update gcp metrics integration default response has a 4xx status code
func (o *UpdateGcpMetricsIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update gcp metrics integration default response has a 5xx status code
func (o *UpdateGcpMetricsIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update gcp metrics integration default response a status code equal to that given
func (o *UpdateGcpMetricsIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update gcp metrics integration default response
func (o *UpdateGcpMetricsIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateGcpMetricsIntegrationDefault) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/gcp-metrics-integrations/{slug}][%d] UpdateGcpMetricsIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateGcpMetricsIntegrationDefault) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/gcp-metrics-integrations/{slug}][%d] UpdateGcpMetricsIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateGcpMetricsIntegrationDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateGcpMetricsIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateGcpMetricsIntegrationBody update gcp metrics integration body
swagger:model UpdateGcpMetricsIntegrationBody
*/
type UpdateGcpMetricsIntegrationBody struct {

	// If true, the GcpMetricsIntegration will be created if it does not already exist, identified by slug. If false, an error will be returned if the GcpMetricsIntegration does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the GcpMetricsIntegration will not be created nor updated, and no response GcpMetricsIntegration will be returned. The response will return an error if the given GcpMetricsIntegration is invalid.
	DryRun bool `json:"dry_run,omitempty"`

	// gcp metrics integration
	GcpMetricsIntegration *models.ConfigunstableGcpMetricsIntegration `json:"gcp_metrics_integration,omitempty"`
}

// Validate validates this update gcp metrics integration body
func (o *UpdateGcpMetricsIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGcpMetricsIntegration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateGcpMetricsIntegrationBody) validateGcpMetricsIntegration(formats strfmt.Registry) error {
	if swag.IsZero(o.GcpMetricsIntegration) { // not required
		return nil
	}

	if o.GcpMetricsIntegration != nil {
		if err := o.GcpMetricsIntegration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "gcp_metrics_integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "gcp_metrics_integration")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update gcp metrics integration body based on the context it is used
func (o *UpdateGcpMetricsIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateGcpMetricsIntegration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateGcpMetricsIntegrationBody) contextValidateGcpMetricsIntegration(ctx context.Context, formats strfmt.Registry) error {

	if o.GcpMetricsIntegration != nil {

		if swag.IsZero(o.GcpMetricsIntegration) { // not required
			return nil
		}

		if err := o.GcpMetricsIntegration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "gcp_metrics_integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "gcp_metrics_integration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateGcpMetricsIntegrationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateGcpMetricsIntegrationBody) UnmarshalBinary(b []byte) error {
	var res UpdateGcpMetricsIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
