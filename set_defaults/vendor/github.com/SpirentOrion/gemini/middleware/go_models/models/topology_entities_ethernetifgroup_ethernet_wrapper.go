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

// TopologyEntitiesEthernetifgroupEthernetWrapper topology entities ethernetifgroup ethernet wrapper
//
// swagger:model topology.entities.ethernetifgroup.EthernetWrapper
type TopologyEntitiesEthernetifgroupEthernetWrapper struct {

	// ethernet
	Ethernet *TopologyEntitiesEthernetifgroupEthernet `json:"ethernet,omitempty"`
}

// Validate validates this topology entities ethernetifgroup ethernet wrapper
func (m *TopologyEntitiesEthernetifgroupEthernetWrapper) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthernet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologyEntitiesEthernetifgroupEthernetWrapper) validateEthernet(formats strfmt.Registry) error {
	if swag.IsZero(m.Ethernet) { // not required
		return nil
	}

	if m.Ethernet != nil {
		if err := m.Ethernet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ethernet")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this topology entities ethernetifgroup ethernet wrapper based on the context it is used
func (m *TopologyEntitiesEthernetifgroupEthernetWrapper) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEthernet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologyEntitiesEthernetifgroupEthernetWrapper) contextValidateEthernet(ctx context.Context, formats strfmt.Registry) error {

	if m.Ethernet != nil {
		if err := m.Ethernet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ethernet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopologyEntitiesEthernetifgroupEthernetWrapper) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopologyEntitiesEthernetifgroupEthernetWrapper) UnmarshalBinary(b []byte) error {
	var res TopologyEntitiesEthernetifgroupEthernetWrapper
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}