// Code generated by go-swagger; DO NOT EDIT.

package webhooks

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

// NewUpdateWebhookParams creates a new UpdateWebhookParams object
// with the default values initialized.
func NewUpdateWebhookParams() *UpdateWebhookParams {
	var ()
	return &UpdateWebhookParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWebhookParamsWithTimeout creates a new UpdateWebhookParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateWebhookParamsWithTimeout(timeout time.Duration) *UpdateWebhookParams {
	var ()
	return &UpdateWebhookParams{

		timeout: timeout,
	}
}

// NewUpdateWebhookParamsWithContext creates a new UpdateWebhookParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateWebhookParamsWithContext(ctx context.Context) *UpdateWebhookParams {
	var ()
	return &UpdateWebhookParams{

		Context: ctx,
	}
}

// NewUpdateWebhookParamsWithHTTPClient creates a new UpdateWebhookParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateWebhookParamsWithHTTPClient(client *http.Client) *UpdateWebhookParams {
	var ()
	return &UpdateWebhookParams{
		HTTPClient: client,
	}
}

/*UpdateWebhookParams contains all the parameters to send to the API endpoint
for the update webhook operation typically these are written to a http.Request
*/
type UpdateWebhookParams struct {

	/*UpdateWebhook
	  Values to update a webhook

	*/
	UpdateWebhook *models.UpdateWebhook
	/*WebhookID
	  Id of the webhook

	*/
	WebhookID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update webhook params
func (o *UpdateWebhookParams) WithTimeout(timeout time.Duration) *UpdateWebhookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update webhook params
func (o *UpdateWebhookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update webhook params
func (o *UpdateWebhookParams) WithContext(ctx context.Context) *UpdateWebhookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update webhook params
func (o *UpdateWebhookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update webhook params
func (o *UpdateWebhookParams) WithHTTPClient(client *http.Client) *UpdateWebhookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update webhook params
func (o *UpdateWebhookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateWebhook adds the updateWebhook to the update webhook params
func (o *UpdateWebhookParams) WithUpdateWebhook(updateWebhook *models.UpdateWebhook) *UpdateWebhookParams {
	o.SetUpdateWebhook(updateWebhook)
	return o
}

// SetUpdateWebhook adds the updateWebhook to the update webhook params
func (o *UpdateWebhookParams) SetUpdateWebhook(updateWebhook *models.UpdateWebhook) {
	o.UpdateWebhook = updateWebhook
}

// WithWebhookID adds the webhookID to the update webhook params
func (o *UpdateWebhookParams) WithWebhookID(webhookID int64) *UpdateWebhookParams {
	o.SetWebhookID(webhookID)
	return o
}

// SetWebhookID adds the webhookId to the update webhook params
func (o *UpdateWebhookParams) SetWebhookID(webhookID int64) {
	o.WebhookID = webhookID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWebhookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdateWebhook != nil {
		if err := r.SetBodyParam(o.UpdateWebhook); err != nil {
			return err
		}
	}

	// path param webhookId
	if err := r.SetPathParam("webhookId", swag.FormatInt64(o.WebhookID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
