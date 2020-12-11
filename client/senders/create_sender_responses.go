// Code generated by go-swagger; DO NOT EDIT.

package senders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/otsofintech/APIv3-go-library/models"
)

// CreateSenderReader is a Reader for the CreateSender structure.
type CreateSenderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSenderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSenderCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSenderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSenderCreated creates a CreateSenderCreated with default headers values
func NewCreateSenderCreated() *CreateSenderCreated {
	return &CreateSenderCreated{}
}

/*CreateSenderCreated handles this case with default header values.

sender created
*/
type CreateSenderCreated struct {
	Payload *models.CreateSenderModel
}

func (o *CreateSenderCreated) Error() string {
	return fmt.Sprintf("[POST /senders][%d] createSenderCreated  %+v", 201, o.Payload)
}

func (o *CreateSenderCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateSenderModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSenderBadRequest creates a CreateSenderBadRequest with default headers values
func NewCreateSenderBadRequest() *CreateSenderBadRequest {
	return &CreateSenderBadRequest{}
}

/*CreateSenderBadRequest handles this case with default header values.

bad request
*/
type CreateSenderBadRequest struct {
	Payload *models.ErrorModel
}

func (o *CreateSenderBadRequest) Error() string {
	return fmt.Sprintf("[POST /senders][%d] createSenderBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSenderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
