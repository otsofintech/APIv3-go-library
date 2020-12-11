// Code generated by go-swagger; DO NOT EDIT.

package email_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/otsofintech/APIv3-go-library/models"
)

// NewCreateEmailCampaignParams creates a new CreateEmailCampaignParams object
// with the default values initialized.
func NewCreateEmailCampaignParams() *CreateEmailCampaignParams {
	var ()
	return &CreateEmailCampaignParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEmailCampaignParamsWithTimeout creates a new CreateEmailCampaignParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateEmailCampaignParamsWithTimeout(timeout time.Duration) *CreateEmailCampaignParams {
	var ()
	return &CreateEmailCampaignParams{

		timeout: timeout,
	}
}

// NewCreateEmailCampaignParamsWithContext creates a new CreateEmailCampaignParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateEmailCampaignParamsWithContext(ctx context.Context) *CreateEmailCampaignParams {
	var ()
	return &CreateEmailCampaignParams{

		Context: ctx,
	}
}

// NewCreateEmailCampaignParamsWithHTTPClient creates a new CreateEmailCampaignParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateEmailCampaignParamsWithHTTPClient(client *http.Client) *CreateEmailCampaignParams {
	var ()
	return &CreateEmailCampaignParams{
		HTTPClient: client,
	}
}

/*CreateEmailCampaignParams contains all the parameters to send to the API endpoint
for the create email campaign operation typically these are written to a http.Request
*/
type CreateEmailCampaignParams struct {

	/*EmailCampaigns
	  Values to create a campaign

	*/
	EmailCampaigns *models.CreateEmailCampaign

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create email campaign params
func (o *CreateEmailCampaignParams) WithTimeout(timeout time.Duration) *CreateEmailCampaignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create email campaign params
func (o *CreateEmailCampaignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create email campaign params
func (o *CreateEmailCampaignParams) WithContext(ctx context.Context) *CreateEmailCampaignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create email campaign params
func (o *CreateEmailCampaignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create email campaign params
func (o *CreateEmailCampaignParams) WithHTTPClient(client *http.Client) *CreateEmailCampaignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create email campaign params
func (o *CreateEmailCampaignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailCampaigns adds the emailCampaigns to the create email campaign params
func (o *CreateEmailCampaignParams) WithEmailCampaigns(emailCampaigns *models.CreateEmailCampaign) *CreateEmailCampaignParams {
	o.SetEmailCampaigns(emailCampaigns)
	return o
}

// SetEmailCampaigns adds the emailCampaigns to the create email campaign params
func (o *CreateEmailCampaignParams) SetEmailCampaigns(emailCampaigns *models.CreateEmailCampaign) {
	o.EmailCampaigns = emailCampaigns
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEmailCampaignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EmailCampaigns != nil {
		if err := r.SetBodyParam(o.EmailCampaigns); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
