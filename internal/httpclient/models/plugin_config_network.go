// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PluginConfigNetwork PluginConfigNetwork PluginConfigNetwork PluginConfigNetwork PluginConfigNetwork PluginConfigNetwork PluginConfigNetwork PluginConfigNetwork PluginConfigNetwork PluginConfigNetwork PluginConfigNetwork PluginConfigNetwork PluginConfigNetwork plugin config network
//
// swagger:model PluginConfigNetwork
type PluginConfigNetwork struct {

	// type
	// Required: true
	Type *string `json:"Type"`
}

// Validate validates this plugin config network
func (m *PluginConfigNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginConfigNetwork) validateType(formats strfmt.Registry) error {

	if err := validate.Required("Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plugin config network based on context it is used
func (m *PluginConfigNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfigNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfigNetwork) UnmarshalBinary(b []byte) error {
	var res PluginConfigNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
