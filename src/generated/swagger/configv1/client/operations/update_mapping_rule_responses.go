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

// UpdateMappingRuleReader is a Reader for the UpdateMappingRule structure.
type UpdateMappingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMappingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMappingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateMappingRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMappingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateMappingRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMappingRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateMappingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMappingRuleOK creates a UpdateMappingRuleOK with default headers values
func NewUpdateMappingRuleOK() *UpdateMappingRuleOK {
	return &UpdateMappingRuleOK{}
}

/*
UpdateMappingRuleOK describes a response with status code 200, with default header values.

A successful response containing the updated MappingRule.
*/
type UpdateMappingRuleOK struct {
	Payload *models.Configv1UpdateMappingRuleResponse
}

// IsSuccess returns true when this update mapping rule o k response has a 2xx status code
func (o *UpdateMappingRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update mapping rule o k response has a 3xx status code
func (o *UpdateMappingRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update mapping rule o k response has a 4xx status code
func (o *UpdateMappingRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update mapping rule o k response has a 5xx status code
func (o *UpdateMappingRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update mapping rule o k response a status code equal to that given
func (o *UpdateMappingRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update mapping rule o k response
func (o *UpdateMappingRuleOK) Code() int {
	return 200
}

func (o *UpdateMappingRuleOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/mapping-rules/{slug}][%d] updateMappingRuleOK  %+v", 200, o.Payload)
}

func (o *UpdateMappingRuleOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/mapping-rules/{slug}][%d] updateMappingRuleOK  %+v", 200, o.Payload)
}

func (o *UpdateMappingRuleOK) GetPayload() *models.Configv1UpdateMappingRuleResponse {
	return o.Payload
}

func (o *UpdateMappingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1UpdateMappingRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMappingRuleBadRequest creates a UpdateMappingRuleBadRequest with default headers values
func NewUpdateMappingRuleBadRequest() *UpdateMappingRuleBadRequest {
	return &UpdateMappingRuleBadRequest{}
}

/*
UpdateMappingRuleBadRequest describes a response with status code 400, with default header values.

Cannot update the MappingRule because the request is invalid.
*/
type UpdateMappingRuleBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update mapping rule bad request response has a 2xx status code
func (o *UpdateMappingRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update mapping rule bad request response has a 3xx status code
func (o *UpdateMappingRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update mapping rule bad request response has a 4xx status code
func (o *UpdateMappingRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update mapping rule bad request response has a 5xx status code
func (o *UpdateMappingRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update mapping rule bad request response a status code equal to that given
func (o *UpdateMappingRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update mapping rule bad request response
func (o *UpdateMappingRuleBadRequest) Code() int {
	return 400
}

func (o *UpdateMappingRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/mapping-rules/{slug}][%d] updateMappingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateMappingRuleBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/mapping-rules/{slug}][%d] updateMappingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateMappingRuleBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateMappingRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMappingRuleNotFound creates a UpdateMappingRuleNotFound with default headers values
func NewUpdateMappingRuleNotFound() *UpdateMappingRuleNotFound {
	return &UpdateMappingRuleNotFound{}
}

/*
UpdateMappingRuleNotFound describes a response with status code 404, with default header values.

Cannot update the MappingRule because the slug does not exist.
*/
type UpdateMappingRuleNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update mapping rule not found response has a 2xx status code
func (o *UpdateMappingRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update mapping rule not found response has a 3xx status code
func (o *UpdateMappingRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update mapping rule not found response has a 4xx status code
func (o *UpdateMappingRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update mapping rule not found response has a 5xx status code
func (o *UpdateMappingRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update mapping rule not found response a status code equal to that given
func (o *UpdateMappingRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update mapping rule not found response
func (o *UpdateMappingRuleNotFound) Code() int {
	return 404
}

func (o *UpdateMappingRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/mapping-rules/{slug}][%d] updateMappingRuleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateMappingRuleNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/mapping-rules/{slug}][%d] updateMappingRuleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateMappingRuleNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateMappingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMappingRuleConflict creates a UpdateMappingRuleConflict with default headers values
func NewUpdateMappingRuleConflict() *UpdateMappingRuleConflict {
	return &UpdateMappingRuleConflict{}
}

/*
UpdateMappingRuleConflict describes a response with status code 409, with default header values.

Cannot update the MappingRule because there is a conflict with an existing MappingRule.
*/
type UpdateMappingRuleConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update mapping rule conflict response has a 2xx status code
func (o *UpdateMappingRuleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update mapping rule conflict response has a 3xx status code
func (o *UpdateMappingRuleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update mapping rule conflict response has a 4xx status code
func (o *UpdateMappingRuleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update mapping rule conflict response has a 5xx status code
func (o *UpdateMappingRuleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update mapping rule conflict response a status code equal to that given
func (o *UpdateMappingRuleConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update mapping rule conflict response
func (o *UpdateMappingRuleConflict) Code() int {
	return 409
}

func (o *UpdateMappingRuleConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/mapping-rules/{slug}][%d] updateMappingRuleConflict  %+v", 409, o.Payload)
}

func (o *UpdateMappingRuleConflict) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/mapping-rules/{slug}][%d] updateMappingRuleConflict  %+v", 409, o.Payload)
}

func (o *UpdateMappingRuleConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateMappingRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMappingRuleInternalServerError creates a UpdateMappingRuleInternalServerError with default headers values
func NewUpdateMappingRuleInternalServerError() *UpdateMappingRuleInternalServerError {
	return &UpdateMappingRuleInternalServerError{}
}

/*
UpdateMappingRuleInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateMappingRuleInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update mapping rule internal server error response has a 2xx status code
func (o *UpdateMappingRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update mapping rule internal server error response has a 3xx status code
func (o *UpdateMappingRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update mapping rule internal server error response has a 4xx status code
func (o *UpdateMappingRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update mapping rule internal server error response has a 5xx status code
func (o *UpdateMappingRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update mapping rule internal server error response a status code equal to that given
func (o *UpdateMappingRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update mapping rule internal server error response
func (o *UpdateMappingRuleInternalServerError) Code() int {
	return 500
}

func (o *UpdateMappingRuleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/mapping-rules/{slug}][%d] updateMappingRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateMappingRuleInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/mapping-rules/{slug}][%d] updateMappingRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateMappingRuleInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateMappingRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMappingRuleDefault creates a UpdateMappingRuleDefault with default headers values
func NewUpdateMappingRuleDefault(code int) *UpdateMappingRuleDefault {
	return &UpdateMappingRuleDefault{
		_statusCode: code,
	}
}

/*
UpdateMappingRuleDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateMappingRuleDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update mapping rule default response has a 2xx status code
func (o *UpdateMappingRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update mapping rule default response has a 3xx status code
func (o *UpdateMappingRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update mapping rule default response has a 4xx status code
func (o *UpdateMappingRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update mapping rule default response has a 5xx status code
func (o *UpdateMappingRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update mapping rule default response a status code equal to that given
func (o *UpdateMappingRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update mapping rule default response
func (o *UpdateMappingRuleDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMappingRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/mapping-rules/{slug}][%d] UpdateMappingRule default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMappingRuleDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/mapping-rules/{slug}][%d] UpdateMappingRule default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMappingRuleDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateMappingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateMappingRuleBody update mapping rule body
swagger:model UpdateMappingRuleBody
*/
type UpdateMappingRuleBody struct {

	// If true, the MappingRule will be created if it does not already exist, identified by slug. If false, an error will be returned if the MappingRule does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the MappingRule will not be created nor updated, and no response MappingRule will be returned. The response will return an error if the given MappingRule is invalid.
	DryRun bool `json:"dry_run,omitempty"`

	// mapping rule
	MappingRule *models.Configv1MappingRule `json:"mapping_rule,omitempty"`
}

// Validate validates this update mapping rule body
func (o *UpdateMappingRuleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMappingRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateMappingRuleBody) validateMappingRule(formats strfmt.Registry) error {
	if swag.IsZero(o.MappingRule) { // not required
		return nil
	}

	if o.MappingRule != nil {
		if err := o.MappingRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "mapping_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "mapping_rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update mapping rule body based on the context it is used
func (o *UpdateMappingRuleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMappingRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateMappingRuleBody) contextValidateMappingRule(ctx context.Context, formats strfmt.Registry) error {

	if o.MappingRule != nil {

		if swag.IsZero(o.MappingRule) { // not required
			return nil
		}

		if err := o.MappingRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "mapping_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "mapping_rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateMappingRuleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateMappingRuleBody) UnmarshalBinary(b []byte) error {
	var res UpdateMappingRuleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
