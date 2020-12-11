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

// DeleteSMTPTemplateReader is a Reader for the DeleteSMTPTemplate structure.
type DeleteSMTPTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSMTPTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSMTPTemplateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteSMTPTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteSMTPTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSMTPTemplateNoContent creates a DeleteSMTPTemplateNoContent with default headers values
func NewDeleteSMTPTemplateNoContent() *DeleteSMTPTemplateNoContent {
	return &DeleteSMTPTemplateNoContent{}
}

/*DeleteSMTPTemplateNoContent handles this case with default header values.

Inactive smtp template has been deleted
*/
type DeleteSMTPTemplateNoContent struct {
}

func (o *DeleteSMTPTemplateNoContent) Error() string {
	return fmt.Sprintf("[DELETE /smtp/templates/{templateId}][%d] deleteSmtpTemplateNoContent ", 204)
}

func (o *DeleteSMTPTemplateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSMTPTemplateBadRequest creates a DeleteSMTPTemplateBadRequest with default headers values
func NewDeleteSMTPTemplateBadRequest() *DeleteSMTPTemplateBadRequest {
	return &DeleteSMTPTemplateBadRequest{}
}

/*DeleteSMTPTemplateBadRequest handles this case with default header values.

bad request
*/
type DeleteSMTPTemplateBadRequest struct {
	Payload *models.ErrorModel
}

func (o *DeleteSMTPTemplateBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /smtp/templates/{templateId}][%d] deleteSmtpTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSMTPTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSMTPTemplateNotFound creates a DeleteSMTPTemplateNotFound with default headers values
func NewDeleteSMTPTemplateNotFound() *DeleteSMTPTemplateNotFound {
	return &DeleteSMTPTemplateNotFound{}
}

/*DeleteSMTPTemplateNotFound handles this case with default header values.

Template ID not found
*/
type DeleteSMTPTemplateNotFound struct {
	Payload *models.ErrorModel
}

func (o *DeleteSMTPTemplateNotFound) Error() string {
	return fmt.Sprintf("[DELETE /smtp/templates/{templateId}][%d] deleteSmtpTemplateNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSMTPTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
