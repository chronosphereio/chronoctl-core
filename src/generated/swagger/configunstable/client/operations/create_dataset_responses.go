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

// CreateDatasetReader is a Reader for the CreateDataset structure.
type CreateDatasetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDatasetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDatasetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDatasetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDatasetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDatasetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateDatasetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDatasetOK creates a CreateDatasetOK with default headers values
func NewCreateDatasetOK() *CreateDatasetOK {
	return &CreateDatasetOK{}
}

/*
CreateDatasetOK describes a response with status code 200, with default header values.

A successful response containing the created Dataset.
*/
type CreateDatasetOK struct {
	Payload *models.ConfigunstableCreateDatasetResponse
}

// IsSuccess returns true when this create dataset o k response has a 2xx status code
func (o *CreateDatasetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create dataset o k response has a 3xx status code
func (o *CreateDatasetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create dataset o k response has a 4xx status code
func (o *CreateDatasetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create dataset o k response has a 5xx status code
func (o *CreateDatasetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create dataset o k response a status code equal to that given
func (o *CreateDatasetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create dataset o k response
func (o *CreateDatasetOK) Code() int {
	return 200
}

func (o *CreateDatasetOK) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/datasets][%d] createDatasetOK  %+v", 200, o.Payload)
}

func (o *CreateDatasetOK) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/datasets][%d] createDatasetOK  %+v", 200, o.Payload)
}

func (o *CreateDatasetOK) GetPayload() *models.ConfigunstableCreateDatasetResponse {
	return o.Payload
}

func (o *CreateDatasetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableCreateDatasetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDatasetBadRequest creates a CreateDatasetBadRequest with default headers values
func NewCreateDatasetBadRequest() *CreateDatasetBadRequest {
	return &CreateDatasetBadRequest{}
}

/*
CreateDatasetBadRequest describes a response with status code 400, with default header values.

Cannot create the Dataset because the request is invalid.
*/
type CreateDatasetBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create dataset bad request response has a 2xx status code
func (o *CreateDatasetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create dataset bad request response has a 3xx status code
func (o *CreateDatasetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create dataset bad request response has a 4xx status code
func (o *CreateDatasetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create dataset bad request response has a 5xx status code
func (o *CreateDatasetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create dataset bad request response a status code equal to that given
func (o *CreateDatasetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create dataset bad request response
func (o *CreateDatasetBadRequest) Code() int {
	return 400
}

func (o *CreateDatasetBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/datasets][%d] createDatasetBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDatasetBadRequest) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/datasets][%d] createDatasetBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDatasetBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateDatasetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDatasetConflict creates a CreateDatasetConflict with default headers values
func NewCreateDatasetConflict() *CreateDatasetConflict {
	return &CreateDatasetConflict{}
}

/*
CreateDatasetConflict describes a response with status code 409, with default header values.

Cannot create the Dataset because there is a conflict with an existing Dataset.
*/
type CreateDatasetConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create dataset conflict response has a 2xx status code
func (o *CreateDatasetConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create dataset conflict response has a 3xx status code
func (o *CreateDatasetConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create dataset conflict response has a 4xx status code
func (o *CreateDatasetConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create dataset conflict response has a 5xx status code
func (o *CreateDatasetConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create dataset conflict response a status code equal to that given
func (o *CreateDatasetConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create dataset conflict response
func (o *CreateDatasetConflict) Code() int {
	return 409
}

func (o *CreateDatasetConflict) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/datasets][%d] createDatasetConflict  %+v", 409, o.Payload)
}

func (o *CreateDatasetConflict) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/datasets][%d] createDatasetConflict  %+v", 409, o.Payload)
}

func (o *CreateDatasetConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateDatasetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDatasetInternalServerError creates a CreateDatasetInternalServerError with default headers values
func NewCreateDatasetInternalServerError() *CreateDatasetInternalServerError {
	return &CreateDatasetInternalServerError{}
}

/*
CreateDatasetInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateDatasetInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create dataset internal server error response has a 2xx status code
func (o *CreateDatasetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create dataset internal server error response has a 3xx status code
func (o *CreateDatasetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create dataset internal server error response has a 4xx status code
func (o *CreateDatasetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create dataset internal server error response has a 5xx status code
func (o *CreateDatasetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create dataset internal server error response a status code equal to that given
func (o *CreateDatasetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create dataset internal server error response
func (o *CreateDatasetInternalServerError) Code() int {
	return 500
}

func (o *CreateDatasetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/datasets][%d] createDatasetInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDatasetInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/datasets][%d] createDatasetInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDatasetInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateDatasetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDatasetDefault creates a CreateDatasetDefault with default headers values
func NewCreateDatasetDefault(code int) *CreateDatasetDefault {
	return &CreateDatasetDefault{
		_statusCode: code,
	}
}

/*
CreateDatasetDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateDatasetDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create dataset default response has a 2xx status code
func (o *CreateDatasetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create dataset default response has a 3xx status code
func (o *CreateDatasetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create dataset default response has a 4xx status code
func (o *CreateDatasetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create dataset default response has a 5xx status code
func (o *CreateDatasetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create dataset default response a status code equal to that given
func (o *CreateDatasetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create dataset default response
func (o *CreateDatasetDefault) Code() int {
	return o._statusCode
}

func (o *CreateDatasetDefault) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/datasets][%d] CreateDataset default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDatasetDefault) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/datasets][%d] CreateDataset default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDatasetDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateDatasetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
