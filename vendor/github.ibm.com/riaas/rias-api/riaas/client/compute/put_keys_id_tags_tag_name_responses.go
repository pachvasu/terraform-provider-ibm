// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// PutKeysIDTagsTagNameReader is a Reader for the PutKeysIDTagsTagName structure.
type PutKeysIDTagsTagNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutKeysIDTagsTagNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutKeysIDTagsTagNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutKeysIDTagsTagNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutKeysIDTagsTagNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutKeysIDTagsTagNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewPutKeysIDTagsTagNameNoContent creates a PutKeysIDTagsTagNameNoContent with default headers values
func NewPutKeysIDTagsTagNameNoContent() *PutKeysIDTagsTagNameNoContent {
	return &PutKeysIDTagsTagNameNoContent{}
}

/*PutKeysIDTagsTagNameNoContent handles this case with default header values.

error
*/
type PutKeysIDTagsTagNameNoContent struct {
	Payload *models.Riaaserror
}

func (o *PutKeysIDTagsTagNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /keys/{id}/tags/{tag_name}][%d] putKeysIdTagsTagNameNoContent  %+v", 204, o.Payload)
}

func (o *PutKeysIDTagsTagNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutKeysIDTagsTagNameBadRequest creates a PutKeysIDTagsTagNameBadRequest with default headers values
func NewPutKeysIDTagsTagNameBadRequest() *PutKeysIDTagsTagNameBadRequest {
	return &PutKeysIDTagsTagNameBadRequest{}
}

/*PutKeysIDTagsTagNameBadRequest handles this case with default header values.

error
*/
type PutKeysIDTagsTagNameBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PutKeysIDTagsTagNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /keys/{id}/tags/{tag_name}][%d] putKeysIdTagsTagNameBadRequest  %+v", 400, o.Payload)
}

func (o *PutKeysIDTagsTagNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutKeysIDTagsTagNameNotFound creates a PutKeysIDTagsTagNameNotFound with default headers values
func NewPutKeysIDTagsTagNameNotFound() *PutKeysIDTagsTagNameNotFound {
	return &PutKeysIDTagsTagNameNotFound{}
}

/*PutKeysIDTagsTagNameNotFound handles this case with default header values.

error
*/
type PutKeysIDTagsTagNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *PutKeysIDTagsTagNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /keys/{id}/tags/{tag_name}][%d] putKeysIdTagsTagNameNotFound  %+v", 404, o.Payload)
}

func (o *PutKeysIDTagsTagNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutKeysIDTagsTagNameDefault creates a PutKeysIDTagsTagNameDefault with default headers values
func NewPutKeysIDTagsTagNameDefault(code int) *PutKeysIDTagsTagNameDefault {
	return &PutKeysIDTagsTagNameDefault{
		_statusCode: code,
	}
}

/*PutKeysIDTagsTagNameDefault handles this case with default header values.

unexpectederror
*/
type PutKeysIDTagsTagNameDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the put keys ID tags tag name default response
func (o *PutKeysIDTagsTagNameDefault) Code() int {
	return o._statusCode
}

func (o *PutKeysIDTagsTagNameDefault) Error() string {
	return fmt.Sprintf("[PUT /keys/{id}/tags/{tag_name}][%d] PutKeysIDTagsTagName default  %+v", o._statusCode, o.Payload)
}

func (o *PutKeysIDTagsTagNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}