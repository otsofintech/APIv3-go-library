// Code generated by go-swagger; DO NOT EDIT.

package reseller

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

// NewAssociateIPToChildParams creates a new AssociateIPToChildParams object
// with the default values initialized.
func NewAssociateIPToChildParams() *AssociateIPToChildParams {
	var ()
	return &AssociateIPToChildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAssociateIPToChildParamsWithTimeout creates a new AssociateIPToChildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAssociateIPToChildParamsWithTimeout(timeout time.Duration) *AssociateIPToChildParams {
	var ()
	return &AssociateIPToChildParams{

		timeout: timeout,
	}
}

// NewAssociateIPToChildParamsWithContext creates a new AssociateIPToChildParams object
// with the default values initialized, and the ability to set a context for a request
func NewAssociateIPToChildParamsWithContext(ctx context.Context) *AssociateIPToChildParams {
	var ()
	return &AssociateIPToChildParams{

		Context: ctx,
	}
}

// NewAssociateIPToChildParamsWithHTTPClient creates a new AssociateIPToChildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAssociateIPToChildParamsWithHTTPClient(client *http.Client) *AssociateIPToChildParams {
	var ()
	return &AssociateIPToChildParams{
		HTTPClient: client,
	}
}

/*AssociateIPToChildParams contains all the parameters to send to the API endpoint
for the associate Ip to child operation typically these are written to a http.Request
*/
type AssociateIPToChildParams struct {

	/*ChildID
	  id of reseller's child

	*/
	ChildID int64
	/*IPID
	  IP's id

	*/
	IPID *models.ManageIP

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the associate Ip to child params
func (o *AssociateIPToChildParams) WithTimeout(timeout time.Duration) *AssociateIPToChildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the associate Ip to child params
func (o *AssociateIPToChildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the associate Ip to child params
func (o *AssociateIPToChildParams) WithContext(ctx context.Context) *AssociateIPToChildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the associate Ip to child params
func (o *AssociateIPToChildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the associate Ip to child params
func (o *AssociateIPToChildParams) WithHTTPClient(client *http.Client) *AssociateIPToChildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the associate Ip to child params
func (o *AssociateIPToChildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChildID adds the childID to the associate Ip to child params
func (o *AssociateIPToChildParams) WithChildID(childID int64) *AssociateIPToChildParams {
	o.SetChildID(childID)
	return o
}

// SetChildID adds the childId to the associate Ip to child params
func (o *AssociateIPToChildParams) SetChildID(childID int64) {
	o.ChildID = childID
}

// WithIPID adds the iPID to the associate Ip to child params
func (o *AssociateIPToChildParams) WithIPID(iPID *models.ManageIP) *AssociateIPToChildParams {
	o.SetIPID(iPID)
	return o
}

// SetIPID adds the ipId to the associate Ip to child params
func (o *AssociateIPToChildParams) SetIPID(iPID *models.ManageIP) {
	o.IPID = iPID
}

// WriteToRequest writes these params to a swagger request
func (o *AssociateIPToChildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param childId
	if err := r.SetPathParam("childId", swag.FormatInt64(o.ChildID)); err != nil {
		return err
	}

	if o.IPID != nil {
		if err := r.SetBodyParam(o.IPID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
