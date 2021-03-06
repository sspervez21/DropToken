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

// Move move
// swagger:model Move
type Move struct {

	// column
	Column int64 `json:"column,omitempty"`

	// player
	// Required: true
	Player *string `json:"player"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this move
func (m *Move) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlayer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Move) validatePlayer(formats strfmt.Registry) error {

	if err := validate.Required("player", "body", m.Player); err != nil {
		return err
	}

	return nil
}

func (m *Move) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Move) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Move) UnmarshalBinary(b []byte) error {
	var res Move
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
