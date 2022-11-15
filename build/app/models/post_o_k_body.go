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

// PostOKBody post o k body
//
// swagger:model postOKBody
type PostOKBody struct {

	// memlab
	Memlab []*PostOKBodyMemlabItems `json:"memlab"`
}

// Validate validates this post o k body
func (m *PostOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMemlab(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostOKBody) validateMemlab(formats strfmt.Registry) error {
	if swag.IsZero(m.Memlab) { // not required
		return nil
	}

	for i := 0; i < len(m.Memlab); i++ {
		if swag.IsZero(m.Memlab[i]) { // not required
			continue
		}

		if m.Memlab[i] != nil {
			if err := m.Memlab[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memlab" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("memlab" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post o k body based on the context it is used
func (m *PostOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMemlab(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostOKBody) contextValidateMemlab(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Memlab); i++ {

		if m.Memlab[i] != nil {
			if err := m.Memlab[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memlab" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("memlab" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostOKBody) UnmarshalBinary(b []byte) error {
	var res PostOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}