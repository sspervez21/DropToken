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

// GameState game state
// swagger:model GameState
type GameState struct {

	// players
	// Required: true
	Players []string `json:"players"`

	// state
	// Required: true
	State *string `json:"state"`

	// winner
	Winner string `json:"winner,omitempty"`
}

// Validate validates this game state
func (m *GameState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlayers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GameState) validatePlayers(formats strfmt.Registry) error {

	if err := validate.Required("players", "body", m.Players); err != nil {
		return err
	}

	for i := 0; i < len(m.Players); i++ {

	}

	return nil
}

func (m *GameState) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GameState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GameState) UnmarshalBinary(b []byte) error {
	var res GameState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
