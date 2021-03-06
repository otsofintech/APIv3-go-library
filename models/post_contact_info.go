// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostContactInfo post contact info
// swagger:model postContactInfo
type PostContactInfo struct {

	// contacts
	// Required: true
	Contacts *PostContactInfoContacts `json:"contacts"`
}

// Validate validates this post contact info
func (m *PostContactInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContacts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostContactInfo) validateContacts(formats strfmt.Registry) error {

	if err := validate.Required("contacts", "body", m.Contacts); err != nil {
		return err
	}

	if m.Contacts != nil {

		if err := m.Contacts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contacts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostContactInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostContactInfo) UnmarshalBinary(b []byte) error {
	var res PostContactInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
