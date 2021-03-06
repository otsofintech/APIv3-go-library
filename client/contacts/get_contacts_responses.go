// Code generated by go-swagger; DO NOT EDIT.

package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/otsofintech/APIv3-go-library/models"
)

// GetContactsReader is a Reader for the GetContacts structure.
type GetContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetContactsOK creates a GetContactsOK with default headers values
func NewGetContactsOK() *GetContactsOK {
	return &GetContactsOK{}
}

/*GetContactsOK handles this case with default header values.

All contacts listed
*/
type GetContactsOK struct {
	Payload *models.GetContacts
}

func (o *GetContactsOK) Error() string {
	return fmt.Sprintf("[GET /contacts][%d] getContactsOK  %+v", 200, o.Payload)
}

func (o *GetContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetContacts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContactsBadRequest creates a GetContactsBadRequest with default headers values
func NewGetContactsBadRequest() *GetContactsBadRequest {
	return &GetContactsBadRequest{}
}

/*GetContactsBadRequest handles this case with default header values.

bad request
*/
type GetContactsBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetContactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /contacts][%d] getContactsBadRequest  %+v", 400, o.Payload)
}

func (o *GetContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
