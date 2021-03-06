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

// NewGame new game
// swagger:model NewGame
type NewGame struct {

	// columns
	// Required: true
	Columns *int64 `json:"columns"`

	// players
	// Required: true
	Players []string `json:"players"`

	// rows
	// Required: true
	Rows *int64 `json:"rows"`
}

// Validate validates this new game
func (m *NewGame) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlayers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewGame) validateColumns(formats strfmt.Registry) error {

	if err := validate.Required("columns", "body", m.Columns); err != nil {
		return err
	}

	return nil
}

func (m *NewGame) validatePlayers(formats strfmt.Registry) error {

	if err := validate.Required("players", "body", m.Players); err != nil {
		return err
	}

	for i := 0; i < len(m.Players); i++ {

	}

	return nil
}

func (m *NewGame) validateRows(formats strfmt.Registry) error {

	if err := validate.Required("rows", "body", m.Rows); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewGame) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewGame) UnmarshalBinary(b []byte) error {
	var res NewGame
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
