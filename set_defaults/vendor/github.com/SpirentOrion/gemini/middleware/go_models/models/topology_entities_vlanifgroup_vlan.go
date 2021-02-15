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

// TopologyEntitiesVlanifgroupVlan topology entities vlanifgroup vlan
//
// swagger:model topology.entities.vlanifgroup.Vlan
type TopologyEntitiesVlanifgroupVlan struct {

	// priority
	Priority *int32 `json:"priority,omitempty"`

	// tpid
	Tpid *string `json:"tpid,omitempty"`

	// VLAN id
	VlanID *TopologyEntitiesBasePatternGroup `json:"vlan-id,omitempty"`
}

// Validate validates this topology entities vlanifgroup vlan
func (m *TopologyEntitiesVlanifgroupVlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologyEntitiesVlanifgroupVlan) validateVlanID(formats strfmt.Registry) error {
	if swag.IsZero(m.VlanID) { // not required
		return nil
	}

	if m.VlanID != nil {
		if err := m.VlanID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan-id")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this topology entities vlanifgroup vlan based on the context it is used
func (m *TopologyEntitiesVlanifgroupVlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVlanID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologyEntitiesVlanifgroupVlan) contextValidateVlanID(ctx context.Context, formats strfmt.Registry) error {

	if m.VlanID != nil {
		if err := m.VlanID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan-id")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopologyEntitiesVlanifgroupVlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopologyEntitiesVlanifgroupVlan) UnmarshalBinary(b []byte) error {
	var res TopologyEntitiesVlanifgroupVlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}