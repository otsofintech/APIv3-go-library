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

	models "github.com/otsofintech/APIv3-go-library/models"
)

// NewSendSMSReportParams creates a new SendSMSReportParams object
// with the default values initialized.
func NewSendSMSReportParams() *SendSMSReportParams {
	var ()
	return &SendSMSReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendSMSReportParamsWithTimeout creates a new SendSMSReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendSMSReportParamsWithTimeout(timeout time.Duration) *SendSMSReportParams {
	var ()
	return &SendSMSReportParams{

		timeout: timeout,
	}
}

// NewSendSMSReportParamsWithContext creates a new SendSMSReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendSMSReportParamsWithContext(ctx context.Context) *SendSMSReportParams {
	var ()
	return &SendSMSReportParams{

		Context: ctx,
	}
}

// NewSendSMSReportParamsWithHTTPClient creates a new SendSMSReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendSMSReportParamsWithHTTPClient(client *http.Client) *SendSMSReportParams {
	var ()
	return &SendSMSReportParams{
		HTTPClient: client,
	}
}

/*SendSMSReportParams contains all the parameters to send to the API endpoint
for the send SMS report operation typically these are written to a http.Request
*/
type SendSMSReportParams struct {

	/*CampaignID
	  id of the campaign

	*/
	CampaignID int64
	/*SendReport
	  Values for send a report

	*/
	SendReport *models.SendReport

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send SMS report params
func (o *SendSMSReportParams) WithTimeout(timeout time.Duration) *SendSMSReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send SMS report params
func (o *SendSMSReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send SMS report params
func (o *SendSMSReportParams) WithContext(ctx context.Context) *SendSMSReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send SMS report params
func (o *SendSMSReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send SMS report params
func (o *SendSMSReportParams) WithHTTPClient(client *http.Client) *SendSMSReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send SMS report params
func (o *SendSMSReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the send SMS report params
func (o *SendSMSReportParams) WithCampaignID(campaignID int64) *SendSMSReportParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the send SMS report params
func (o *SendSMSReportParams) SetCampaignID(campaignID int64) {
	o.CampaignID = campaignID
}

// WithSendReport adds the sendReport to the send SMS report params
func (o *SendSMSReportParams) WithSendReport(sendReport *models.SendReport) *SendSMSReportParams {
	o.SetSendReport(sendReport)
	return o
}

// SetSendReport adds the sendReport to the send SMS report params
func (o *SendSMSReportParams) SetSendReport(sendReport *models.SendReport) {
	o.SendReport = sendReport
}

// WriteToRequest writes these params to a swagger request
func (o *SendSMSReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaignId
	if err := r.SetPathParam("campaignId", swag.FormatInt64(o.CampaignID)); err != nil {
		return err
	}

	if o.SendReport != nil {
		if err := r.SetBodyParam(o.SendReport); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
