// Code generated by go-swagger; DO NOT EDIT.

package senders

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
)

// NewDeleteSenderParams creates a new DeleteSenderParams object
// with the default values initialized.
func NewDeleteSenderParams() *DeleteSenderParams {
	var ()
	return &DeleteSenderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSenderParamsWithTimeout creates a new DeleteSenderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSenderParamsWithTimeout(timeout time.Duration) *DeleteSenderParams {
	var ()
	return &DeleteSenderParams{

		timeout: timeout,
	}
}

// NewDeleteSenderParamsWithContext creates a new DeleteSenderParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSenderParamsWithContext(ctx context.Context) *DeleteSenderParams {
	var ()
	return &DeleteSenderParams{

		Context: ctx,
	}
}

// NewDeleteSenderParamsWithHTTPClient creates a new DeleteSenderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSenderParamsWithHTTPClient(client *http.Client) *DeleteSenderParams {
	var ()
	return &DeleteSenderParams{
		HTTPClient: client,
	}
}

/*DeleteSenderParams contains all the parameters to send to the API endpoint
for the delete sender operation typically these are written to a http.Request
*/
type DeleteSenderParams struct {

	/*SenderID
	  Id of the sender

	*/
	SenderID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete sender params
func (o *DeleteSenderParams) WithTimeout(timeout time.Duration) *DeleteSenderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sender params
func (o *DeleteSenderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sender params
func (o *DeleteSenderParams) WithContext(ctx context.Context) *DeleteSenderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sender params
func (o *DeleteSenderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sender params
func (o *DeleteSenderParams) WithHTTPClient(client *http.Client) *DeleteSenderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sender params
func (o *DeleteSenderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSenderID adds the senderID to the delete sender params
func (o *DeleteSenderParams) WithSenderID(senderID int64) *DeleteSenderParams {
	o.SetSenderID(senderID)
	return o
}

// SetSenderID adds the senderId to the delete sender params
func (o *DeleteSenderParams) SetSenderID(senderID int64) {
	o.SenderID = senderID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSenderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param senderId
	if err := r.SetPathParam("senderId", swag.FormatInt64(o.SenderID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
