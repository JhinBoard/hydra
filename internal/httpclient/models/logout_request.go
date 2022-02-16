// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogoutRequest LogoutRequest LogoutRequest LogoutRequest Contains information about an ongoing logout request.
//
// swagger:model logoutRequest
type LogoutRequest struct {

	// Challenge is the identifier ("logout challenge") of the logout authentication request. It is used to
	// identify the session.
	Challenge string `json:"challenge,omitempty"`

	// client
	Client *OAuth2Client `json:"client,omitempty"`

	// RequestURL is the original Logout URL requested.
	RequestURL string `json:"request_url,omitempty"`

	// RPInitiated is set to true if the request was initiated by a Relying Party (RP), also known as an OAuth 2.0 Client.
	RpInitiated bool `json:"rp_initiated,omitempty"`

	// SessionID is the login session ID that was requested to log out.
	Sid string `json:"sid,omitempty"`

	// Subject is the user for whom the logout was request.
	Subject string `json:"subject,omitempty"`
}

// Validate validates this logout request
func (m *LogoutRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClient(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogoutRequest) validateClient(formats strfmt.Registry) error {
	if swag.IsZero(m.Client) { // not required
		return nil
	}

	if m.Client != nil {
		if err := m.Client.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this logout request based on the context it is used
func (m *LogoutRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogoutRequest) contextValidateClient(ctx context.Context, formats strfmt.Registry) error {

	if m.Client != nil {
		if err := m.Client.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogoutRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogoutRequest) UnmarshalBinary(b []byte) error {
	var res LogoutRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
