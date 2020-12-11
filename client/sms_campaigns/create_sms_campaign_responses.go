// Code generated by go-swagger; DO NOT EDIT.

package sms_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/otsofintech/APIv3-go-library/models"
)

// CreateSMSCampaignReader is a Reader for the CreateSMSCampaign structure.
type CreateSMSCampaignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSMSCampaignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSMSCampaignCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSMSCampaignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSMSCampaignCreated creates a CreateSMSCampaignCreated with default headers values
func NewCreateSMSCampaignCreated() *CreateSMSCampaignCreated {
	return &CreateSMSCampaignCreated{}
}

/*CreateSMSCampaignCreated handles this case with default header values.

SMS campaign created
*/
type CreateSMSCampaignCreated struct {
	Payload *models.CreateModel
}

func (o *CreateSMSCampaignCreated) Error() string {
	return fmt.Sprintf("[POST /smsCampaigns][%d] createSmsCampaignCreated  %+v", 201, o.Payload)
}

func (o *CreateSMSCampaignCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSMSCampaignBadRequest creates a CreateSMSCampaignBadRequest with default headers values
func NewCreateSMSCampaignBadRequest() *CreateSMSCampaignBadRequest {
	return &CreateSMSCampaignBadRequest{}
}

/*CreateSMSCampaignBadRequest handles this case with default header values.

bad request
*/
type CreateSMSCampaignBadRequest struct {
	Payload *models.ErrorModel
}

func (o *CreateSMSCampaignBadRequest) Error() string {
	return fmt.Sprintf("[POST /smsCampaigns][%d] createSmsCampaignBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSMSCampaignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
