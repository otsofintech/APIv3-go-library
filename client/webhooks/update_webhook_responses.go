// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/otsofintech/APIv3-go-library/models"
)

// UpdateWebhookReader is a Reader for the UpdateWebhook structure.
type UpdateWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateWebhookNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateWebhookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateWebhookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateWebhookNoContent creates a UpdateWebhookNoContent with default headers values
func NewUpdateWebhookNoContent() *UpdateWebhookNoContent {
	return &UpdateWebhookNoContent{}
}

/*UpdateWebhookNoContent handles this case with default header values.

Webhook updated
*/
type UpdateWebhookNoContent struct {
}

func (o *UpdateWebhookNoContent) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{webhookId}][%d] updateWebhookNoContent ", 204)
}

func (o *UpdateWebhookNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateWebhookBadRequest creates a UpdateWebhookBadRequest with default headers values
func NewUpdateWebhookBadRequest() *UpdateWebhookBadRequest {
	return &UpdateWebhookBadRequest{}
}

/*UpdateWebhookBadRequest handles this case with default header values.

bad request
*/
type UpdateWebhookBadRequest struct {
	Payload *models.ErrorModel
}

func (o *UpdateWebhookBadRequest) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{webhookId}][%d] updateWebhookBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateWebhookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateWebhookNotFound creates a UpdateWebhookNotFound with default headers values
func NewUpdateWebhookNotFound() *UpdateWebhookNotFound {
	return &UpdateWebhookNotFound{}
}

/*UpdateWebhookNotFound handles this case with default header values.

Webhook ID not found
*/
type UpdateWebhookNotFound struct {
	Payload *models.ErrorModel
}

func (o *UpdateWebhookNotFound) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{webhookId}][%d] updateWebhookNotFound  %+v", 404, o.Payload)
}

func (o *UpdateWebhookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
