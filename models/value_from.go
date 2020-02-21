// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ValueFrom value from
// swagger:model valueFrom
type ValueFrom struct {

	// Name of the environment variable.
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// value ref
	// Required: true
	ValueRef *ValueRef `json:"valueRef"`
}

// Validate validates this value from
func (m *ValueFrom) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ValueFrom) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *ValueFrom) validateValueRef(formats strfmt.Registry) error {

	if err := validate.Required("valueRef", "body", m.ValueRef); err != nil {
		return err
	}

	if m.ValueRef != nil {
		if err := m.ValueRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valueRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ValueFrom) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ValueFrom) UnmarshalBinary(b []byte) error {
	var res ValueFrom
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
