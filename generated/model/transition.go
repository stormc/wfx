// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: 2023 Siemens AG
//
// SPDX-License-Identifier: Apache-2.0
//

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Transition transition
//
// swagger:model Transition
type Transition struct {

	// The transition execution action (default: WAIT)
	// Example: WAIT
	Action ActionEnum `json:"action,omitempty"`

	// description
	// Example: Description of the transition
	Description string `json:"description,omitempty"`

	// The entity that may execute the transition
	// Example: CLIENT
	// Required: true
	Eligible EligibleEnum `json:"eligible"`

	// from
	// Example: START
	// Required: true
	From string `json:"from"`

	// to
	// Example: END
	// Required: true
	To string `json:"to"`
}

// Validate validates this transition
func (m *Transition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEligible(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transition) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := m.Action.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("action")
		}
		return err
	}

	return nil
}

func (m *Transition) validateEligible(formats strfmt.Registry) error {

	if err := validate.Required("eligible", "body", EligibleEnum(m.Eligible)); err != nil {
		return err
	}

	if err := m.Eligible.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("eligible")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("eligible")
		}
		return err
	}

	return nil
}

func (m *Transition) validateFrom(formats strfmt.Registry) error {

	if err := validate.RequiredString("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

func (m *Transition) validateTo(formats strfmt.Registry) error {

	if err := validate.RequiredString("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this transition based on the context it is used
func (m *Transition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEligible(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transition) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := m.Action.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("action")
		}
		return err
	}

	return nil
}

func (m *Transition) contextValidateEligible(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Eligible.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("eligible")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("eligible")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Transition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transition) UnmarshalBinary(b []byte) error {
	var res Transition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
