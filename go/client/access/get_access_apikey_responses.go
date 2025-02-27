// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loopholelabs/scale/go/client/models"
)

// GetAccessApikeyReader is a Reader for the GetAccessApikey structure.
type GetAccessApikeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessApikeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccessApikeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAccessApikeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccessApikeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccessApikeyOK creates a GetAccessApikeyOK with default headers values
func NewGetAccessApikeyOK() *GetAccessApikeyOK {
	return &GetAccessApikeyOK{}
}

/*
GetAccessApikeyOK describes a response with status code 200, with default header values.

OK
*/
type GetAccessApikeyOK struct {
	Payload []*models.ModelsGetAPIKeyResponse
}

// IsSuccess returns true when this get access apikey o k response has a 2xx status code
func (o *GetAccessApikeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get access apikey o k response has a 3xx status code
func (o *GetAccessApikeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access apikey o k response has a 4xx status code
func (o *GetAccessApikeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get access apikey o k response has a 5xx status code
func (o *GetAccessApikeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get access apikey o k response a status code equal to that given
func (o *GetAccessApikeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get access apikey o k response
func (o *GetAccessApikeyOK) Code() int {
	return 200
}

func (o *GetAccessApikeyOK) Error() string {
	return fmt.Sprintf("[GET /access/apikey][%d] getAccessApikeyOK  %+v", 200, o.Payload)
}

func (o *GetAccessApikeyOK) String() string {
	return fmt.Sprintf("[GET /access/apikey][%d] getAccessApikeyOK  %+v", 200, o.Payload)
}

func (o *GetAccessApikeyOK) GetPayload() []*models.ModelsGetAPIKeyResponse {
	return o.Payload
}

func (o *GetAccessApikeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessApikeyUnauthorized creates a GetAccessApikeyUnauthorized with default headers values
func NewGetAccessApikeyUnauthorized() *GetAccessApikeyUnauthorized {
	return &GetAccessApikeyUnauthorized{}
}

/*
GetAccessApikeyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAccessApikeyUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this get access apikey unauthorized response has a 2xx status code
func (o *GetAccessApikeyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get access apikey unauthorized response has a 3xx status code
func (o *GetAccessApikeyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access apikey unauthorized response has a 4xx status code
func (o *GetAccessApikeyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get access apikey unauthorized response has a 5xx status code
func (o *GetAccessApikeyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get access apikey unauthorized response a status code equal to that given
func (o *GetAccessApikeyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get access apikey unauthorized response
func (o *GetAccessApikeyUnauthorized) Code() int {
	return 401
}

func (o *GetAccessApikeyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /access/apikey][%d] getAccessApikeyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAccessApikeyUnauthorized) String() string {
	return fmt.Sprintf("[GET /access/apikey][%d] getAccessApikeyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAccessApikeyUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *GetAccessApikeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessApikeyInternalServerError creates a GetAccessApikeyInternalServerError with default headers values
func NewGetAccessApikeyInternalServerError() *GetAccessApikeyInternalServerError {
	return &GetAccessApikeyInternalServerError{}
}

/*
GetAccessApikeyInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAccessApikeyInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get access apikey internal server error response has a 2xx status code
func (o *GetAccessApikeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get access apikey internal server error response has a 3xx status code
func (o *GetAccessApikeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access apikey internal server error response has a 4xx status code
func (o *GetAccessApikeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get access apikey internal server error response has a 5xx status code
func (o *GetAccessApikeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get access apikey internal server error response a status code equal to that given
func (o *GetAccessApikeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get access apikey internal server error response
func (o *GetAccessApikeyInternalServerError) Code() int {
	return 500
}

func (o *GetAccessApikeyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /access/apikey][%d] getAccessApikeyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccessApikeyInternalServerError) String() string {
	return fmt.Sprintf("[GET /access/apikey][%d] getAccessApikeyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccessApikeyInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetAccessApikeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
