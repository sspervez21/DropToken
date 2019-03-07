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

// NewMove new move
// swagger:model NewMove
type NewMove struct {

	// column
	// Required: true
	Column *int64 `json:"column"`
}

// Validate validates this new move
func (m *NewMove) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColumn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewMove) validateColumn(formats strfmt.Registry) error {

	if err := validate.Required("column", "body", m.Column); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewMove) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewMove) UnmarshalBinary(b []byte) error {
	var res NewMove
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
