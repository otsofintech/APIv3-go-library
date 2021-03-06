// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetWebhooksWebhooksItems get webhooks webhooks items
// swagger:model getWebhooksWebhooksItems
type GetWebhooksWebhooksItems struct {
	GetWebhook
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetWebhooksWebhooksItems) UnmarshalJSON(raw []byte) error {

	var aO0 GetWebhook
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetWebhook = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetWebhooksWebhooksItems) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.GetWebhook)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get webhooks webhooks items
func (m *GetWebhooksWebhooksItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.GetWebhook.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetWebhooksWebhooksItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetWebhooksWebhooksItems) UnmarshalBinary(b []byte) error {
	var res GetWebhooksWebhooksItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
