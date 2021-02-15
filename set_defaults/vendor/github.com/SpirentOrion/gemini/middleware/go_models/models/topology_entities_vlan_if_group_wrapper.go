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

// TopologyEntitiesVlanIfGroupWrapper topology entities vlan if group wrapper
//
// swagger:model topology.entities.VlanIfGroupWrapper
type TopologyEntitiesVlanIfGroupWrapper struct {

	// vlans
	Vlans *TopologyEntitiesVlanIfGroup `json:"vlans,omitempty"`
}

// Validate validates this topology entities vlan if group wrapper
func (m *TopologyEntitiesVlanIfGroupWrapper) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVlans(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologyEntitiesVlanIfGroupWrapper) validateVlans(formats strfmt.Registry) error {
	if swag.IsZero(m.Vlans) { // not required
		return nil
	}

	if m.Vlans != nil {
		if err := m.Vlans.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlans")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this topology entities vlan if group wrapper based on the context it is used
func (m *TopologyEntitiesVlanIfGroupWrapper) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVlans(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologyEntitiesVlanIfGroupWrapper) contextValidateVlans(ctx context.Context, formats strfmt.Registry) error {

	if m.Vlans != nil {
		if err := m.Vlans.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlans")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopologyEntitiesVlanIfGroupWrapper) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopologyEntitiesVlanIfGroupWrapper) UnmarshalBinary(b []byte) error {
	var res TopologyEntitiesVlanIfGroupWrapper
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}