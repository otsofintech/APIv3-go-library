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

// GetChildInfoReader is a Reader for the GetChildInfo structure.
type GetChildInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChildInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetChildInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetChildInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetChildInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetChildInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetChildInfoOK creates a GetChildInfoOK with default headers values
func NewGetChildInfoOK() *GetChildInfoOK {
	return &GetChildInfoOK{}
}

/*GetChildInfoOK handles this case with default header values.

Information for the child
*/
type GetChildInfoOK struct {
	Payload *models.GetChildInfo
}

func (o *GetChildInfoOK) Error() string {
	return fmt.Sprintf("[GET /reseller/children/{childId}][%d] getChildInfoOK  %+v", 200, o.Payload)
}

func (o *GetChildInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetChildInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChildInfoBadRequest creates a GetChildInfoBadRequest with default headers values
func NewGetChildInfoBadRequest() *GetChildInfoBadRequest {
	return &GetChildInfoBadRequest{}
}

/*GetChildInfoBadRequest handles this case with default header values.

bad request
*/
type GetChildInfoBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetChildInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /reseller/children/{childId}][%d] getChildInfoBadRequest  %+v", 400, o.Payload)
}

func (o *GetChildInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChildInfoForbidden creates a GetChildInfoForbidden with default headers values
func NewGetChildInfoForbidden() *GetChildInfoForbidden {
	return &GetChildInfoForbidden{}
}

/*GetChildInfoForbidden handles this case with default header values.

Current account is not a reseller
*/
type GetChildInfoForbidden struct {
	Payload *models.ErrorModel
}

func (o *GetChildInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /reseller/children/{childId}][%d] getChildInfoForbidden  %+v", 403, o.Payload)
}

func (o *GetChildInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChildInfoNotFound creates a GetChildInfoNotFound with default headers values
func NewGetChildInfoNotFound() *GetChildInfoNotFound {
	return &GetChildInfoNotFound{}
}

/*GetChildInfoNotFound handles this case with default header values.

Child ID not found
*/
type GetChildInfoNotFound struct {
	Payload *models.ErrorModel
}

func (o *GetChildInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /reseller/children/{childId}][%d] getChildInfoNotFound  %+v", 404, o.Payload)
}

func (o *GetChildInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
