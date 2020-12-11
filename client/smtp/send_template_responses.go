// Code generated by go-swagger; DO NOT EDIT.

package smtp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/otsofintech/APIv3-go-library/models"
)

// SendTemplateReader is a Reader for the SendTemplate structure.
type SendTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewSendTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendTemplateCreated creates a SendTemplateCreated with default headers values
func NewSendTemplateCreated() *SendTemplateCreated {
	return &SendTemplateCreated{}
}

/*SendTemplateCreated handles this case with default header values.

Email has been sent successfully to all recipients
*/
type SendTemplateCreated struct {
	Payload *models.SendTemplateEmail
}

func (o *SendTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /smtp/templates/{templateId}/send][%d] sendTemplateCreated  %+v", 201, o.Payload)
}

func (o *SendTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SendTemplateEmail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendTemplateBadRequest creates a SendTemplateBadRequest with default headers values
func NewSendTemplateBadRequest() *SendTemplateBadRequest {
	return &SendTemplateBadRequest{}
}

/*SendTemplateBadRequest handles this case with default header values.

Email could not be sent to the following email addresses
*/
type SendTemplateBadRequest struct {
	Payload *models.PostSendFailed
}

func (o *SendTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /smtp/templates/{templateId}/send][%d] sendTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *SendTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostSendFailed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendTemplateNotFound creates a SendTemplateNotFound with default headers values
func NewSendTemplateNotFound() *SendTemplateNotFound {
	return &SendTemplateNotFound{}
}

/*SendTemplateNotFound handles this case with default header values.

Template ID not found
*/
type SendTemplateNotFound struct {
	Payload *models.ErrorModel
}

func (o *SendTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /smtp/templates/{templateId}/send][%d] sendTemplateNotFound  %+v", 404, o.Payload)
}

func (o *SendTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
