// Code generated by go-swagger; DO NOT EDIT.

package reseller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/otsofintech/APIv3-go-library/models"
)

// AssociateIPToChildReader is a Reader for the AssociateIPToChild structure.
type AssociateIPToChildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssociateIPToChildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewAssociateIPToChildNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAssociateIPToChildBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAssociateIPToChildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAssociateIPToChildNoContent creates a AssociateIPToChildNoContent with default headers values
func NewAssociateIPToChildNoContent() *AssociateIPToChildNoContent {
	return &AssociateIPToChildNoContent{}
}

/*AssociateIPToChildNoContent handles this case with default header values.

Dedicated IP is associated to the child
*/
type AssociateIPToChildNoContent struct {
}

func (o *AssociateIPToChildNoContent) Error() string {
	return fmt.Sprintf("[POST /reseller/children/{childId}/ips/associate][%d] associateIpToChildNoContent ", 204)
}

func (o *AssociateIPToChildNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssociateIPToChildBadRequest creates a AssociateIPToChildBadRequest with default headers values
func NewAssociateIPToChildBadRequest() *AssociateIPToChildBadRequest {
	return &AssociateIPToChildBadRequest{}
}

/*AssociateIPToChildBadRequest handles this case with default header values.

Dedicated IP is not associated to this child
*/
type AssociateIPToChildBadRequest struct {
	Payload *models.ErrorModel
}

func (o *AssociateIPToChildBadRequest) Error() string {
	return fmt.Sprintf("[POST /reseller/children/{childId}/ips/associate][%d] associateIpToChildBadRequest  %+v", 400, o.Payload)
}

func (o *AssociateIPToChildBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssociateIPToChildNotFound creates a AssociateIPToChildNotFound with default headers values
func NewAssociateIPToChildNotFound() *AssociateIPToChildNotFound {
	return &AssociateIPToChildNotFound{}
}

/*AssociateIPToChildNotFound handles this case with default header values.

Child ID not found
*/
type AssociateIPToChildNotFound struct {
	Payload *models.ErrorModel
}

func (o *AssociateIPToChildNotFound) Error() string {
	return fmt.Sprintf("[POST /reseller/children/{childId}/ips/associate][%d] associateIpToChildNotFound  %+v", 404, o.Payload)
}

func (o *AssociateIPToChildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
