// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp/internal/pkg/api/releasesapi/models"
)

// ListReleasesV1Reader is a Reader for the ListReleasesV1 structure.
type ListReleasesV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListReleasesV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListReleasesV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListReleasesV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListReleasesV1OK creates a ListReleasesV1OK with default headers values
func NewListReleasesV1OK() *ListReleasesV1OK {
	return &ListReleasesV1OK{}
}

/*
ListReleasesV1OK describes a response with status code 200, with default header values.

OK (releases retrieved)
*/
type ListReleasesV1OK struct {

	/* The ETag for this resource
	 */
	ETag string

	Payload models.ProductReleasesResponseV1
}

// IsSuccess returns true when this list releases v1 o k response has a 2xx status code
func (o *ListReleasesV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list releases v1 o k response has a 3xx status code
func (o *ListReleasesV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list releases v1 o k response has a 4xx status code
func (o *ListReleasesV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list releases v1 o k response has a 5xx status code
func (o *ListReleasesV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list releases v1 o k response a status code equal to that given
func (o *ListReleasesV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list releases v1 o k response
func (o *ListReleasesV1OK) Code() int {
	return 200
}

func (o *ListReleasesV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/releases/{product}][%d] listReleasesV1OK  %+v", 200, o.Payload)
}

func (o *ListReleasesV1OK) String() string {
	return fmt.Sprintf("[GET /v1/releases/{product}][%d] listReleasesV1OK  %+v", 200, o.Payload)
}

func (o *ListReleasesV1OK) GetPayload() models.ProductReleasesResponseV1 {
	return o.Payload
}

func (o *ListReleasesV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header ETag
	hdrETag := response.GetHeader("ETag")

	if hdrETag != "" {
		o.ETag = hdrETag
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReleasesV1Default creates a ListReleasesV1Default with default headers values
func NewListReleasesV1Default(code int) *ListReleasesV1Default {
	return &ListReleasesV1Default{
		_statusCode: code,
	}
}

/*
ListReleasesV1Default describes a response with status code -1, with default header values.

an error condition
*/
type ListReleasesV1Default struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list releases v1 default response has a 2xx status code
func (o *ListReleasesV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list releases v1 default response has a 3xx status code
func (o *ListReleasesV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list releases v1 default response has a 4xx status code
func (o *ListReleasesV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list releases v1 default response has a 5xx status code
func (o *ListReleasesV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list releases v1 default response a status code equal to that given
func (o *ListReleasesV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list releases v1 default response
func (o *ListReleasesV1Default) Code() int {
	return o._statusCode
}

func (o *ListReleasesV1Default) Error() string {
	return fmt.Sprintf("[GET /v1/releases/{product}][%d] listReleasesV1 default  %+v", o._statusCode, o.Payload)
}

func (o *ListReleasesV1Default) String() string {
	return fmt.Sprintf("[GET /v1/releases/{product}][%d] listReleasesV1 default  %+v", o._statusCode, o.Payload)
}

func (o *ListReleasesV1Default) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListReleasesV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
