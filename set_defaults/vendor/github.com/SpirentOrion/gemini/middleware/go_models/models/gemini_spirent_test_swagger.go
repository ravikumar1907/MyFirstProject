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

// GeminiSpirentTest gemini spirent test
//
// swagger:model gemini.SpirentTest
type GeminiSpirentTest struct {

	// The description of the test.
	Description string `json:"description,omitempty"`

	// UUID.
	ID string `json:"id,omitempty"`

	// The name of the test.
	Name string `json:"name,omitempty"`

	// Global end-point collection
	SpirentEndpoints []*TestEndpointEndpointGroup `json:"spirent-endpoints"`

	// spirent topologies
	SpirentTopologies []*TopologyTopology `json:"spirent-topologies"`

	// This field captures UI related metadata as an decoded string.
	//          Backend echos back the values as is.
	UIDescriptor string `json:"ui-descriptor,omitempty"`
}

// Validate validates this gemini spirent test
func (m *GeminiSpirentTest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpirentEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpirentTopologies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeminiSpirentTest) validateSpirentEndpoints(formats strfmt.Registry) error {
	if swag.IsZero(m.SpirentEndpoints) { // not required
		return nil
	}

	for i := 0; i < len(m.SpirentEndpoints); i++ {
		if swag.IsZero(m.SpirentEndpoints[i]) { // not required
			continue
		}

		if m.SpirentEndpoints[i] != nil {
			if err := m.SpirentEndpoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spirent-endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GeminiSpirentTest) validateSpirentTopologies(formats strfmt.Registry) error {
	if swag.IsZero(m.SpirentTopologies) { // not required
		return nil
	}

	for i := 0; i < len(m.SpirentTopologies); i++ {
		if swag.IsZero(m.SpirentTopologies[i]) { // not required
			continue
		}

		if m.SpirentTopologies[i] != nil {
			if err := m.SpirentTopologies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spirent-topologies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this gemini spirent test based on the context it is used
func (m *GeminiSpirentTest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSpirentEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpirentTopologies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeminiSpirentTest) contextValidateSpirentEndpoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SpirentEndpoints); i++ {

		if m.SpirentEndpoints[i] != nil {
			if err := m.SpirentEndpoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spirent-endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GeminiSpirentTest) contextValidateSpirentTopologies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SpirentTopologies); i++ {

		if m.SpirentTopologies[i] != nil {
			if err := m.SpirentTopologies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spirent-topologies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GeminiSpirentTest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeminiSpirentTest) UnmarshalBinary(b []byte) error {
	var res GeminiSpirentTest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
