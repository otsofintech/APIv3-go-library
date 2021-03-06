// Code generated by go-swagger; DO NOT EDIT.

package email_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/otsofintech/APIv3-go-library/models"
)

// GetEmailCampaignsReader is a Reader for the GetEmailCampaigns structure.
type GetEmailCampaignsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmailCampaignsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEmailCampaignsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetEmailCampaignsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmailCampaignsOK creates a GetEmailCampaignsOK with default headers values
func NewGetEmailCampaignsOK() *GetEmailCampaignsOK {
	return &GetEmailCampaignsOK{}
}

/*GetEmailCampaignsOK handles this case with default header values.

Email campaigns informations
*/
type GetEmailCampaignsOK struct {
	Payload *models.GetEmailCampaigns
}

func (o *GetEmailCampaignsOK) Error() string {
	return fmt.Sprintf("[GET /emailCampaigns][%d] getEmailCampaignsOK  %+v", 200, o.Payload)
}

func (o *GetEmailCampaignsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetEmailCampaigns)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmailCampaignsBadRequest creates a GetEmailCampaignsBadRequest with default headers values
func NewGetEmailCampaignsBadRequest() *GetEmailCampaignsBadRequest {
	return &GetEmailCampaignsBadRequest{}
}

/*GetEmailCampaignsBadRequest handles this case with default header values.

bad request
*/
type GetEmailCampaignsBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetEmailCampaignsBadRequest) Error() string {
	return fmt.Sprintf("[GET /emailCampaigns][%d] getEmailCampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *GetEmailCampaignsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
