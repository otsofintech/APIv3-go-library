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

// GetListsListsItemsAllOf1 get lists lists items all of1
// swagger:model getListsListsItemsAllOf1
type GetListsListsItemsAllOf1 struct {

	// ID of the folder
	// Required: true
	FolderID *int64 `json:"folderId"`
}

// Validate validates this get lists lists items all of1
func (m *GetListsListsItemsAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFolderID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetListsListsItemsAllOf1) validateFolderID(formats strfmt.Registry) error {

	if err := validate.Required("folderId", "body", m.FolderID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetListsListsItemsAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetListsListsItemsAllOf1) UnmarshalBinary(b []byte) error {
	var res GetListsListsItemsAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
