// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IpamsvcHAGroupHeartbeats ipamsvc h a group heartbeats
//
// swagger:model ipamsvcHAGroupHeartbeats
type IpamsvcHAGroupHeartbeats struct {

	// The name of the peer.
	Peer string `json:"peer,omitempty"`

	// The timestamp as a string of the last successful heartbeat received from the peer above.
	SuccessfulHeartbeat string `json:"successful_heartbeat,omitempty"`
}

// Validate validates this ipamsvc h a group heartbeats
func (m *IpamsvcHAGroupHeartbeats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ipamsvc h a group heartbeats based on context it is used
func (m *IpamsvcHAGroupHeartbeats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcHAGroupHeartbeats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcHAGroupHeartbeats) UnmarshalBinary(b []byte) error {
	var res IpamsvcHAGroupHeartbeats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}