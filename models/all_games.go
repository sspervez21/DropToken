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

// AllGames all games
// swagger:model AllGames
type AllGames struct {

	// games
	// Required: true
	Games []string `json:"games"`
}

// Validate validates this all games
func (m *AllGames) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllGames) validateGames(formats strfmt.Registry) error {

	if err := validate.Required("games", "body", m.Games); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllGames) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllGames) UnmarshalBinary(b []byte) error {
	var res AllGames
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
