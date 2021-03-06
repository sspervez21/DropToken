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

// MoveOutput move output
// swagger:model MoveOutput
type MoveOutput struct {

	// move
	// Required: true
	Move *string `json:"move"`
}

// Validate validates this move output
func (m *MoveOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMove(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveOutput) validateMove(formats strfmt.Registry) error {

	if err := validate.Required("move", "body", m.Move); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveOutput) UnmarshalBinary(b []byte) error {
	var res MoveOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
