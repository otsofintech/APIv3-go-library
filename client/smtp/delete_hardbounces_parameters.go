// Code generated by go-swagger; DO NOT EDIT.

package smtp

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

// NewDeleteHardbouncesParams creates a new DeleteHardbouncesParams object
// with the default values initialized.
func NewDeleteHardbouncesParams() *DeleteHardbouncesParams {
	var ()
	return &DeleteHardbouncesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteHardbouncesParamsWithTimeout creates a new DeleteHardbouncesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteHardbouncesParamsWithTimeout(timeout time.Duration) *DeleteHardbouncesParams {
	var ()
	return &DeleteHardbouncesParams{

		timeout: timeout,
	}
}

// NewDeleteHardbouncesParamsWithContext creates a new DeleteHardbouncesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteHardbouncesParamsWithContext(ctx context.Context) *DeleteHardbouncesParams {
	var ()
	return &DeleteHardbouncesParams{

		Context: ctx,
	}
}

// NewDeleteHardbouncesParamsWithHTTPClient creates a new DeleteHardbouncesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteHardbouncesParamsWithHTTPClient(client *http.Client) *DeleteHardbouncesParams {
	var ()
	return &DeleteHardbouncesParams{
		HTTPClient: client,
	}
}

/*DeleteHardbouncesParams contains all the parameters to send to the API endpoint
for the delete hardbounces operation typically these are written to a http.Request
*/
type DeleteHardbouncesParams struct {

	/*DeleteHardbounces
	  values to delete hardbounces

	*/
	DeleteHardbounces *models.DeleteHardbounces

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete hardbounces params
func (o *DeleteHardbouncesParams) WithTimeout(timeout time.Duration) *DeleteHardbouncesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete hardbounces params
func (o *DeleteHardbouncesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete hardbounces params
func (o *DeleteHardbouncesParams) WithContext(ctx context.Context) *DeleteHardbouncesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete hardbounces params
func (o *DeleteHardbouncesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete hardbounces params
func (o *DeleteHardbouncesParams) WithHTTPClient(client *http.Client) *DeleteHardbouncesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete hardbounces params
func (o *DeleteHardbouncesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteHardbounces adds the deleteHardbounces to the delete hardbounces params
func (o *DeleteHardbouncesParams) WithDeleteHardbounces(deleteHardbounces *models.DeleteHardbounces) *DeleteHardbouncesParams {
	o.SetDeleteHardbounces(deleteHardbounces)
	return o
}

// SetDeleteHardbounces adds the deleteHardbounces to the delete hardbounces params
func (o *DeleteHardbouncesParams) SetDeleteHardbounces(deleteHardbounces *models.DeleteHardbounces) {
	o.DeleteHardbounces = deleteHardbounces
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteHardbouncesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteHardbounces != nil {
		if err := r.SetBodyParam(o.DeleteHardbounces); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
