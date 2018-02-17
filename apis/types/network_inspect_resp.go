// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NetworkInspectResp is the expected body of the 'GET networks/{id}'' http request message
// swagger:model NetworkInspectResp

type NetworkInspectResp struct {

	// Driver means the network's driver.
	Driver string `json:"Driver,omitempty"`

	// EnableIPv6 represents whether to enable IPv6.
	EnableIPV6 bool `json:"EnableIPv6,omitempty"`

	// ID uniquely identifies a network on a single machine
	ID string `json:"ID,omitempty"`

	// IPAM is the network's IP Address Management.
	IPAM *IPAM `json:"IPAM,omitempty"`

	// Internal checks the network is internal network or not.
	Internal bool `json:"Internal,omitempty"`

	// Labels holds metadata specific to the network being created.
	Labels map[string]string `json:"Labels,omitempty"`

	// Name is the requested name of the network
	Name string `json:"Name,omitempty"`

	// Options holds the network specific options to use for when creating the network.
	Options map[string]string `json:"Options,omitempty"`

	// Scope describes the level at which the network exists.
	Scope string `json:"Scope,omitempty"`
}

/* polymorph NetworkInspectResp Driver false */

/* polymorph NetworkInspectResp EnableIPv6 false */

/* polymorph NetworkInspectResp ID false */

/* polymorph NetworkInspectResp IPAM false */

/* polymorph NetworkInspectResp Internal false */

/* polymorph NetworkInspectResp Labels false */

/* polymorph NetworkInspectResp Name false */

/* polymorph NetworkInspectResp Options false */

/* polymorph NetworkInspectResp Scope false */

// Validate validates this network inspect resp
func (m *NetworkInspectResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAM(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkInspectResp) validateIPAM(formats strfmt.Registry) error {

	if swag.IsZero(m.IPAM) { // not required
		return nil
	}

	if m.IPAM != nil {

		if err := m.IPAM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IPAM")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInspectResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInspectResp) UnmarshalBinary(b []byte) error {
	var res NetworkInspectResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
