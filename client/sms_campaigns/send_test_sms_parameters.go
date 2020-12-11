// Code generated by go-swagger; DO NOT EDIT.

package sms_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// NewSendTestSMSParams creates a new SendTestSMSParams object
// with the default values initialized.
func NewSendTestSMSParams() *SendTestSMSParams {
	var ()
	return &SendTestSMSParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendTestSMSParamsWithTimeout creates a new SendTestSMSParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendTestSMSParamsWithTimeout(timeout time.Duration) *SendTestSMSParams {
	var ()
	return &SendTestSMSParams{

		timeout: timeout,
	}
}

// NewSendTestSMSParamsWithContext creates a new SendTestSMSParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendTestSMSParamsWithContext(ctx context.Context) *SendTestSMSParams {
	var ()
	return &SendTestSMSParams{

		Context: ctx,
	}
}

// NewSendTestSMSParamsWithHTTPClient creates a new SendTestSMSParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendTestSMSParamsWithHTTPClient(client *http.Client) *SendTestSMSParams {
	var ()
	return &SendTestSMSParams{
		HTTPClient: client,
	}
}

/*SendTestSMSParams contains all the parameters to send to the API endpoint
for the send test Sms operation typically these are written to a http.Request
*/
type SendTestSMSParams struct {

	/*CampaignID
	  Id of the SMS campaign

	*/
	CampaignID int64
	/*SendTestSMS
	  Mobile number to which send the test

	*/
	SendTestSMS *models.SendTestSMS

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send test Sms params
func (o *SendTestSMSParams) WithTimeout(timeout time.Duration) *SendTestSMSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send test Sms params
func (o *SendTestSMSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send test Sms params
func (o *SendTestSMSParams) WithContext(ctx context.Context) *SendTestSMSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send test Sms params
func (o *SendTestSMSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send test Sms params
func (o *SendTestSMSParams) WithHTTPClient(client *http.Client) *SendTestSMSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send test Sms params
func (o *SendTestSMSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the send test Sms params
func (o *SendTestSMSParams) WithCampaignID(campaignID int64) *SendTestSMSParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the send test Sms params
func (o *SendTestSMSParams) SetCampaignID(campaignID int64) {
	o.CampaignID = campaignID
}

// WithSendTestSMS adds the sendTestSMS to the send test Sms params
func (o *SendTestSMSParams) WithSendTestSMS(sendTestSMS *models.SendTestSMS) *SendTestSMSParams {
	o.SetSendTestSMS(sendTestSMS)
	return o
}

// SetSendTestSMS adds the sendTestSms to the send test Sms params
func (o *SendTestSMSParams) SetSendTestSMS(sendTestSMS *models.SendTestSMS) {
	o.SendTestSMS = sendTestSMS
}

// WriteToRequest writes these params to a swagger request
func (o *SendTestSMSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaignId
	if err := r.SetPathParam("campaignId", swag.FormatInt64(o.CampaignID)); err != nil {
		return err
	}

	if o.SendTestSMS != nil {
		if err := r.SetBodyParam(o.SendTestSMS); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}