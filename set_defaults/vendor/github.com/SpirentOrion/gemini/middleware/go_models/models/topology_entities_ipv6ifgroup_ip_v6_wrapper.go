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

// TopologyEntitiesIpv6ifgroupIPV6Wrapper topology entities ipv6ifgroup Ipv6 wrapper
//
// swagger:model topology.entities.ipv6ifgroup.Ipv6Wrapper
type TopologyEntitiesIpv6ifgroupIPV6Wrapper struct {

	// ipv6
	IPV6 *TopologyEntitiesIpv6ifgroupIPV6 `json:"ipv6,omitempty"`
}

// Validate validates this topology entities ipv6ifgroup Ipv6 wrapper
func (m *TopologyEntitiesIpv6ifgroupIPV6Wrapper) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPV6(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologyEntitiesIpv6ifgroupIPV6Wrapper) validateIPV6(formats strfmt.Registry) error {
	if swag.IsZero(m.IPV6) { // not required
		return nil
	}

	if m.IPV6 != nil {
		if err := m.IPV6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv6")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this topology entities ipv6ifgroup Ipv6 wrapper based on the context it is used
func (m *TopologyEntitiesIpv6ifgroupIPV6Wrapper) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPV6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologyEntitiesIpv6ifgroupIPV6Wrapper) contextValidateIPV6(ctx context.Context, formats strfmt.Registry) error {

	if m.IPV6 != nil {
		if err := m.IPV6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv6")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopologyEntitiesIpv6ifgroupIPV6Wrapper) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopologyEntitiesIpv6ifgroupIPV6Wrapper) UnmarshalBinary(b []byte) error {
	var res TopologyEntitiesIpv6ifgroupIPV6Wrapper
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
