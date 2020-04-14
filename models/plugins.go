// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Plugins plugins
// swagger:model plugins
type Plugins struct {

	// swagger
	Swagger *bool `json:"swagger,omitempty"`
}

func (m *Plugins) UnmarshalJSON(b []byte) error {
	type PluginsAlias Plugins
	var t PluginsAlias
	if err := json.Unmarshal([]byte("{\"swagger\":true}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = Plugins(t)
	return nil
}

// Validate validates this plugins
func (m *Plugins) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Plugins) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plugins) UnmarshalBinary(b []byte) error {
	var res Plugins
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
