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

// GetIpsFromSenderReader is a Reader for the GetIpsFromSender structure.
type GetIpsFromSenderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIpsFromSenderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIpsFromSenderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetIpsFromSenderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetIpsFromSenderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIpsFromSenderOK creates a GetIpsFromSenderOK with default headers values
func NewGetIpsFromSenderOK() *GetIpsFromSenderOK {
	return &GetIpsFromSenderOK{}
}

/*GetIpsFromSenderOK handles this case with default header values.

list of dedicated IPs
*/
type GetIpsFromSenderOK struct {
	Payload *models.GetIpsFromSender
}

func (o *GetIpsFromSenderOK) Error() string {
	return fmt.Sprintf("[GET /senders/{senderId}/ips][%d] getIpsFromSenderOK  %+v", 200, o.Payload)
}

func (o *GetIpsFromSenderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIpsFromSender)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIpsFromSenderBadRequest creates a GetIpsFromSenderBadRequest with default headers values
func NewGetIpsFromSenderBadRequest() *GetIpsFromSenderBadRequest {
	return &GetIpsFromSenderBadRequest{}
}

/*GetIpsFromSenderBadRequest handles this case with default header values.

bad request
*/
type GetIpsFromSenderBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetIpsFromSenderBadRequest) Error() string {
	return fmt.Sprintf("[GET /senders/{senderId}/ips][%d] getIpsFromSenderBadRequest  %+v", 400, o.Payload)
}

func (o *GetIpsFromSenderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIpsFromSenderNotFound creates a GetIpsFromSenderNotFound with default headers values
func NewGetIpsFromSenderNotFound() *GetIpsFromSenderNotFound {
	return &GetIpsFromSenderNotFound{}
}

/*GetIpsFromSenderNotFound handles this case with default header values.

Sender ID not found
*/
type GetIpsFromSenderNotFound struct {
	Payload *models.ErrorModel
}

func (o *GetIpsFromSenderNotFound) Error() string {
	return fmt.Sprintf("[GET /senders/{senderId}/ips][%d] getIpsFromSenderNotFound  %+v", 404, o.Payload)
}

func (o *GetIpsFromSenderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
