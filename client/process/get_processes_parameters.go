// Code generated by go-swagger; DO NOT EDIT.

package process

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

// NewGetProcessesParams creates a new GetProcessesParams object
// with the default values initialized.
func NewGetProcessesParams() *GetProcessesParams {
	var (
		limitDefault  = int64(10)
		offsetDefault = int64(0)
	)
	return &GetProcessesParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProcessesParamsWithTimeout creates a new GetProcessesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProcessesParamsWithTimeout(timeout time.Duration) *GetProcessesParams {
	var (
		limitDefault  = int64(10)
		offsetDefault = int64(0)
	)
	return &GetProcessesParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetProcessesParamsWithContext creates a new GetProcessesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProcessesParamsWithContext(ctx context.Context) *GetProcessesParams {
	var (
		limitDefault  = int64(10)
		offsetDefault = int64(0)
	)
	return &GetProcessesParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetProcessesParamsWithHTTPClient creates a new GetProcessesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProcessesParamsWithHTTPClient(client *http.Client) *GetProcessesParams {
	var (
		limitDefault  = int64(10)
		offsetDefault = int64(0)
	)
	return &GetProcessesParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetProcessesParams contains all the parameters to send to the API endpoint
for the get processes operation typically these are written to a http.Request
*/
type GetProcessesParams struct {

	/*Limit
	  Number limitation for the result returned

	*/
	Limit *int64
	/*Offset
	  Beginning point in the list to retrieve from.

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get processes params
func (o *GetProcessesParams) WithTimeout(timeout time.Duration) *GetProcessesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get processes params
func (o *GetProcessesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get processes params
func (o *GetProcessesParams) WithContext(ctx context.Context) *GetProcessesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get processes params
func (o *GetProcessesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get processes params
func (o *GetProcessesParams) WithHTTPClient(client *http.Client) *GetProcessesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get processes params
func (o *GetProcessesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get processes params
func (o *GetProcessesParams) WithLimit(limit *int64) *GetProcessesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get processes params
func (o *GetProcessesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get processes params
func (o *GetProcessesParams) WithOffset(offset *int64) *GetProcessesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get processes params
func (o *GetProcessesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetProcessesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}