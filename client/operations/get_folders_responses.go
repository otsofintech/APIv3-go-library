// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/otsofintech/APIv3-go-library/models"
)

// GetFoldersReader is a Reader for the GetFolders structure.
type GetFoldersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFoldersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFoldersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFoldersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFoldersOK creates a GetFoldersOK with default headers values
func NewGetFoldersOK() *GetFoldersOK {
	return &GetFoldersOK{}
}

/*GetFoldersOK handles this case with default header values.

Folders informations
*/
type GetFoldersOK struct {
	Payload *models.GetFolders
}

func (o *GetFoldersOK) Error() string {
	return fmt.Sprintf("[GET /contacts/folders][%d] getFoldersOK  %+v", 200, o.Payload)
}

func (o *GetFoldersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetFolders)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFoldersBadRequest creates a GetFoldersBadRequest with default headers values
func NewGetFoldersBadRequest() *GetFoldersBadRequest {
	return &GetFoldersBadRequest{}
}

/*GetFoldersBadRequest handles this case with default header values.

bad request
*/
type GetFoldersBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetFoldersBadRequest) Error() string {
	return fmt.Sprintf("[GET /contacts/folders][%d] getFoldersBadRequest  %+v", 400, o.Payload)
}

func (o *GetFoldersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
