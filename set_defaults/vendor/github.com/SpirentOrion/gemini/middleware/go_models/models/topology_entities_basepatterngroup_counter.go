// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TopologyEntitiesBasepatterngroupCounter topology entities basepatterngroup counter
//
// swagger:model topology.entities.basepatterngroup.Counter
type TopologyEntitiesBasepatterngroupCounter struct {

	// nested steps
	NestedSteps []*TopologyEntitiesStepGroup `json:"nested-steps"`

	// The start value of the increment pattern
	Start string `json:"start,omitempty"`

	// The step value of the increment pattern
	Step *string `json:"step,omitempty"`
}

// Validate validates this topology entities basepatterngroup counter
func (m *TopologyEntitiesBasepatterngroupCounter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNestedSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologyEntitiesBasepatterngroupCounter) validateNestedSteps(formats strfmt.Registry) error {
	if swag.IsZero(m.NestedSteps) { // not required
		return nil
	}

	for i := 0; i < len(m.NestedSteps); i++ {
		if swag.IsZero(m.NestedSteps[i]) { // not required
			continue
		}

		if m.NestedSteps[i] != nil {
			if err := m.NestedSteps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nested-steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this topology entities basepatterngroup counter based on the context it is used
func (m *TopologyEntitiesBasepatterngroupCounter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNestedSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologyEntitiesBasepatterngroupCounter) contextValidateNestedSteps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NestedSteps); i++ {

		if m.NestedSteps[i] != nil {
			if err := m.NestedSteps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nested-steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopologyEntitiesBasepatterngroupCounter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopologyEntitiesBasepatterngroupCounter) UnmarshalBinary(b []byte) error {
	var res TopologyEntitiesBasepatterngroupCounter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
