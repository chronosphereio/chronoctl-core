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

// UpdateNotificationPolicyReader is a Reader for the UpdateNotificationPolicy structure.
type UpdateNotificationPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNotificationPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNotificationPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateNotificationPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateNotificationPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateNotificationPolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateNotificationPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateNotificationPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateNotificationPolicyOK creates a UpdateNotificationPolicyOK with default headers values
func NewUpdateNotificationPolicyOK() *UpdateNotificationPolicyOK {
	return &UpdateNotificationPolicyOK{}
}

/*
UpdateNotificationPolicyOK describes a response with status code 200, with default header values.

A successful response containing the updated NotificationPolicy.
*/
type UpdateNotificationPolicyOK struct {
	Payload *models.Configv1UpdateNotificationPolicyResponse
}

// IsSuccess returns true when this update notification policy o k response has a 2xx status code
func (o *UpdateNotificationPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update notification policy o k response has a 3xx status code
func (o *UpdateNotificationPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update notification policy o k response has a 4xx status code
func (o *UpdateNotificationPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update notification policy o k response has a 5xx status code
func (o *UpdateNotificationPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update notification policy o k response a status code equal to that given
func (o *UpdateNotificationPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update notification policy o k response
func (o *UpdateNotificationPolicyOK) Code() int {
	return 200
}

func (o *UpdateNotificationPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/notification-policies/{slug}][%d] updateNotificationPolicyOK  %+v", 200, o.Payload)
}

func (o *UpdateNotificationPolicyOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/notification-policies/{slug}][%d] updateNotificationPolicyOK  %+v", 200, o.Payload)
}

func (o *UpdateNotificationPolicyOK) GetPayload() *models.Configv1UpdateNotificationPolicyResponse {
	return o.Payload
}

func (o *UpdateNotificationPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1UpdateNotificationPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationPolicyBadRequest creates a UpdateNotificationPolicyBadRequest with default headers values
func NewUpdateNotificationPolicyBadRequest() *UpdateNotificationPolicyBadRequest {
	return &UpdateNotificationPolicyBadRequest{}
}

/*
UpdateNotificationPolicyBadRequest describes a response with status code 400, with default header values.

Cannot update the NotificationPolicy because the request is invalid.
*/
type UpdateNotificationPolicyBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update notification policy bad request response has a 2xx status code
func (o *UpdateNotificationPolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update notification policy bad request response has a 3xx status code
func (o *UpdateNotificationPolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update notification policy bad request response has a 4xx status code
func (o *UpdateNotificationPolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update notification policy bad request response has a 5xx status code
func (o *UpdateNotificationPolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update notification policy bad request response a status code equal to that given
func (o *UpdateNotificationPolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update notification policy bad request response
func (o *UpdateNotificationPolicyBadRequest) Code() int {
	return 400
}

func (o *UpdateNotificationPolicyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/notification-policies/{slug}][%d] updateNotificationPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateNotificationPolicyBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/notification-policies/{slug}][%d] updateNotificationPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateNotificationPolicyBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateNotificationPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationPolicyNotFound creates a UpdateNotificationPolicyNotFound with default headers values
func NewUpdateNotificationPolicyNotFound() *UpdateNotificationPolicyNotFound {
	return &UpdateNotificationPolicyNotFound{}
}

/*
UpdateNotificationPolicyNotFound describes a response with status code 404, with default header values.

Cannot update the NotificationPolicy because the slug does not exist.
*/
type UpdateNotificationPolicyNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update notification policy not found response has a 2xx status code
func (o *UpdateNotificationPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update notification policy not found response has a 3xx status code
func (o *UpdateNotificationPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update notification policy not found response has a 4xx status code
func (o *UpdateNotificationPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update notification policy not found response has a 5xx status code
func (o *UpdateNotificationPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update notification policy not found response a status code equal to that given
func (o *UpdateNotificationPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update notification policy not found response
func (o *UpdateNotificationPolicyNotFound) Code() int {
	return 404
}

func (o *UpdateNotificationPolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/notification-policies/{slug}][%d] updateNotificationPolicyNotFound  %+v", 404, o.Payload)
}

func (o *UpdateNotificationPolicyNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/notification-policies/{slug}][%d] updateNotificationPolicyNotFound  %+v", 404, o.Payload)
}

func (o *UpdateNotificationPolicyNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateNotificationPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationPolicyConflict creates a UpdateNotificationPolicyConflict with default headers values
func NewUpdateNotificationPolicyConflict() *UpdateNotificationPolicyConflict {
	return &UpdateNotificationPolicyConflict{}
}

/*
UpdateNotificationPolicyConflict describes a response with status code 409, with default header values.

Cannot update the NotificationPolicy because there is a conflict with an existing NotificationPolicy.
*/
type UpdateNotificationPolicyConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update notification policy conflict response has a 2xx status code
func (o *UpdateNotificationPolicyConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update notification policy conflict response has a 3xx status code
func (o *UpdateNotificationPolicyConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update notification policy conflict response has a 4xx status code
func (o *UpdateNotificationPolicyConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update notification policy conflict response has a 5xx status code
func (o *UpdateNotificationPolicyConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update notification policy conflict response a status code equal to that given
func (o *UpdateNotificationPolicyConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update notification policy conflict response
func (o *UpdateNotificationPolicyConflict) Code() int {
	return 409
}

func (o *UpdateNotificationPolicyConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/notification-policies/{slug}][%d] updateNotificationPolicyConflict  %+v", 409, o.Payload)
}

func (o *UpdateNotificationPolicyConflict) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/notification-policies/{slug}][%d] updateNotificationPolicyConflict  %+v", 409, o.Payload)
}

func (o *UpdateNotificationPolicyConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateNotificationPolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationPolicyInternalServerError creates a UpdateNotificationPolicyInternalServerError with default headers values
func NewUpdateNotificationPolicyInternalServerError() *UpdateNotificationPolicyInternalServerError {
	return &UpdateNotificationPolicyInternalServerError{}
}

/*
UpdateNotificationPolicyInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateNotificationPolicyInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update notification policy internal server error response has a 2xx status code
func (o *UpdateNotificationPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update notification policy internal server error response has a 3xx status code
func (o *UpdateNotificationPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update notification policy internal server error response has a 4xx status code
func (o *UpdateNotificationPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update notification policy internal server error response has a 5xx status code
func (o *UpdateNotificationPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update notification policy internal server error response a status code equal to that given
func (o *UpdateNotificationPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update notification policy internal server error response
func (o *UpdateNotificationPolicyInternalServerError) Code() int {
	return 500
}

func (o *UpdateNotificationPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/notification-policies/{slug}][%d] updateNotificationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateNotificationPolicyInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/notification-policies/{slug}][%d] updateNotificationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateNotificationPolicyInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateNotificationPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationPolicyDefault creates a UpdateNotificationPolicyDefault with default headers values
func NewUpdateNotificationPolicyDefault(code int) *UpdateNotificationPolicyDefault {
	return &UpdateNotificationPolicyDefault{
		_statusCode: code,
	}
}

/*
UpdateNotificationPolicyDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateNotificationPolicyDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update notification policy default response has a 2xx status code
func (o *UpdateNotificationPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update notification policy default response has a 3xx status code
func (o *UpdateNotificationPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update notification policy default response has a 4xx status code
func (o *UpdateNotificationPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update notification policy default response has a 5xx status code
func (o *UpdateNotificationPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update notification policy default response a status code equal to that given
func (o *UpdateNotificationPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update notification policy default response
func (o *UpdateNotificationPolicyDefault) Code() int {
	return o._statusCode
}

func (o *UpdateNotificationPolicyDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/notification-policies/{slug}][%d] UpdateNotificationPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateNotificationPolicyDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/notification-policies/{slug}][%d] UpdateNotificationPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateNotificationPolicyDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateNotificationPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNotificationPolicyBody update notification policy body
swagger:model UpdateNotificationPolicyBody
*/
type UpdateNotificationPolicyBody struct {

	// If true, the NotificationPolicy will be created if it does not already exist, identified by slug. If false, an error will be returned if the NotificationPolicy does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the NotificationPolicy will not be created nor updated, and no response NotificationPolicy will be returned. The response will return an error if the given NotificationPolicy is invalid.
	DryRun bool `json:"dry_run,omitempty"`

	// notification policy
	NotificationPolicy *models.Configv1NotificationPolicy `json:"notification_policy,omitempty"`
}

// Validate validates this update notification policy body
func (o *UpdateNotificationPolicyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNotificationPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNotificationPolicyBody) validateNotificationPolicy(formats strfmt.Registry) error {
	if swag.IsZero(o.NotificationPolicy) { // not required
		return nil
	}

	if o.NotificationPolicy != nil {
		if err := o.NotificationPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "notification_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "notification_policy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update notification policy body based on the context it is used
func (o *UpdateNotificationPolicyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateNotificationPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNotificationPolicyBody) contextValidateNotificationPolicy(ctx context.Context, formats strfmt.Registry) error {

	if o.NotificationPolicy != nil {

		if swag.IsZero(o.NotificationPolicy) { // not required
			return nil
		}

		if err := o.NotificationPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "notification_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "notification_policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNotificationPolicyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNotificationPolicyBody) UnmarshalBinary(b []byte) error {
	var res UpdateNotificationPolicyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}