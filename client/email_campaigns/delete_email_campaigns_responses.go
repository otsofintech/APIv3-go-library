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

// DeleteEmailCampaignsReader is a Reader for the DeleteEmailCampaigns structure.
type DeleteEmailCampaignsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEmailCampaignsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteEmailCampaignsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteEmailCampaignsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteEmailCampaignsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteEmailCampaignsNoContent creates a DeleteEmailCampaignsNoContent with default headers values
func NewDeleteEmailCampaignsNoContent() *DeleteEmailCampaignsNoContent {
	return &DeleteEmailCampaignsNoContent{}
}

/*DeleteEmailCampaignsNoContent handles this case with default header values.

Email campaign has been deleted
*/
type DeleteEmailCampaignsNoContent struct {
}

func (o *DeleteEmailCampaignsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /emailCampaigns/{campaignId}][%d] deleteEmailCampaignsNoContent ", 204)
}

func (o *DeleteEmailCampaignsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEmailCampaignsBadRequest creates a DeleteEmailCampaignsBadRequest with default headers values
func NewDeleteEmailCampaignsBadRequest() *DeleteEmailCampaignsBadRequest {
	return &DeleteEmailCampaignsBadRequest{}
}

/*DeleteEmailCampaignsBadRequest handles this case with default header values.

bad request
*/
type DeleteEmailCampaignsBadRequest struct {
	Payload *models.ErrorModel
}

func (o *DeleteEmailCampaignsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /emailCampaigns/{campaignId}][%d] deleteEmailCampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteEmailCampaignsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEmailCampaignsNotFound creates a DeleteEmailCampaignsNotFound with default headers values
func NewDeleteEmailCampaignsNotFound() *DeleteEmailCampaignsNotFound {
	return &DeleteEmailCampaignsNotFound{}
}

/*DeleteEmailCampaignsNotFound handles this case with default header values.

Campaign ID not found
*/
type DeleteEmailCampaignsNotFound struct {
	Payload *models.ErrorModel
}

func (o *DeleteEmailCampaignsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /emailCampaigns/{campaignId}][%d] deleteEmailCampaignsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteEmailCampaignsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
