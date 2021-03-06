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

// GetFolderLists get folder lists
// swagger:model getFolderLists
type GetFolderLists struct {

	// Number of lists in the folder
	// Required: true
	Count *int64 `json:"count"`

	// lists
	// Required: true
	Lists GetFolderListsLists `json:"lists"`
}

// Validate validates this get folder lists
func (m *GetFolderLists) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLists(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetFolderLists) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *GetFolderLists) validateLists(formats strfmt.Registry) error {

	if err := validate.Required("lists", "body", m.Lists); err != nil {
		return err
	}

	if err := m.Lists.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lists")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetFolderLists) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFolderLists) UnmarshalBinary(b []byte) error {
	var res GetFolderLists
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
