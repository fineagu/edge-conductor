// Code generated by go-swagger; DO NOT EDIT.

//
//   Copyright (c) 2022 Intel Corporation.
//
//   SPDX-License-Identifier: Apache-2.0
//
//
//

package workflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Workflow workflow
//
// swagger:model workflow
type Workflow struct {

	// api version
	APIVersion string `json:"apiVersion,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// metadata
	Metadata *WorkflowMetadata `json:"metadata,omitempty"`

	// spec
	Spec *WorkflowSpec `json:"spec,omitempty"`
}

// Validate validates this workflow
func (m *Workflow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workflow) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *Workflow) validateSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this workflow based on the context it is used
func (m *Workflow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workflow) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *Workflow) contextValidateSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.Spec != nil {
		if err := m.Spec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Workflow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Workflow) UnmarshalBinary(b []byte) error {
	var res Workflow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowMetadata workflow metadata
//
// swagger:model WorkflowMetadata
type WorkflowMetadata struct {

	// name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Name string `json:"name,omitempty"`

	// namespace
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Namespace string `json:"namespace,omitempty"`
}

// Validate validates this workflow metadata
func (m *WorkflowMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowMetadata) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("metadata"+"."+"name", "body", m.Name, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowMetadata) validateNamespace(formats strfmt.Registry) error {
	if swag.IsZero(m.Namespace) { // not required
		return nil
	}

	if err := validate.Pattern("metadata"+"."+"namespace", "body", m.Namespace, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this workflow metadata based on context it is used
func (m *WorkflowMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowMetadata) UnmarshalBinary(b []byte) error {
	var res WorkflowMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowSpec workflow spec
//
// swagger:model WorkflowSpec
type WorkflowSpec struct {

	// containers
	Containers Containers `json:"containers,omitempty"`

	// data
	Data []*WorkflowSpecDataItems0 `json:"data"`

	// plugins
	Plugins []*WorkflowSpecPluginsItems0 `json:"plugins"`

	// workflows
	Workflows []*WorkflowSpecWorkflowsItems0 `json:"workflows"`
}

// Validate validates this workflow spec
func (m *WorkflowSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlugins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowSpec) validateContainers(formats strfmt.Registry) error {
	if swag.IsZero(m.Containers) { // not required
		return nil
	}

	if err := m.Containers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("spec" + "." + "containers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("spec" + "." + "containers")
		}
		return err
	}

	return nil
}

func (m *WorkflowSpec) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spec" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowSpec) validatePlugins(formats strfmt.Registry) error {
	if swag.IsZero(m.Plugins) { // not required
		return nil
	}

	for i := 0; i < len(m.Plugins); i++ {
		if swag.IsZero(m.Plugins[i]) { // not required
			continue
		}

		if m.Plugins[i] != nil {
			if err := m.Plugins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "plugins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spec" + "." + "plugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowSpec) validateWorkflows(formats strfmt.Registry) error {
	if swag.IsZero(m.Workflows) { // not required
		return nil
	}

	for i := 0; i < len(m.Workflows); i++ {
		if swag.IsZero(m.Workflows[i]) { // not required
			continue
		}

		if m.Workflows[i] != nil {
			if err := m.Workflows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "workflows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spec" + "." + "workflows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this workflow spec based on the context it is used
func (m *WorkflowSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlugins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkflows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowSpec) contextValidateContainers(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Containers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("spec" + "." + "containers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("spec" + "." + "containers")
		}
		return err
	}

	return nil
}

func (m *WorkflowSpec) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {
			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spec" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowSpec) contextValidatePlugins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Plugins); i++ {

		if m.Plugins[i] != nil {
			if err := m.Plugins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "plugins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spec" + "." + "plugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowSpec) contextValidateWorkflows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Workflows); i++ {

		if m.Workflows[i] != nil {
			if err := m.Workflows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "workflows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spec" + "." + "workflows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowSpec) UnmarshalBinary(b []byte) error {
	var res WorkflowSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowSpecDataItems0 workflow spec data items0
//
// swagger:model WorkflowSpecDataItems0
type WorkflowSpecDataItems0 struct {

	// confidential
	Confidential bool `json:"confidential,omitempty"`

	// name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Name string `json:"name,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this workflow spec data items0
func (m *WorkflowSpecDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowSpecDataItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this workflow spec data items0 based on context it is used
func (m *WorkflowSpecDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowSpecDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowSpecDataItems0) UnmarshalBinary(b []byte) error {
	var res WorkflowSpecDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowSpecPluginsItems0 workflow spec plugins items0
//
// swagger:model WorkflowSpecPluginsItems0
type WorkflowSpecPluginsItems0 struct {

	// container
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Container string `json:"container,omitempty"`

	// name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this workflow spec plugins items0
func (m *WorkflowSpecPluginsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowSpecPluginsItems0) validateContainer(formats strfmt.Registry) error {
	if swag.IsZero(m.Container) { // not required
		return nil
	}

	if err := validate.Pattern("container", "body", m.Container, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowSpecPluginsItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this workflow spec plugins items0 based on context it is used
func (m *WorkflowSpecPluginsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowSpecPluginsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowSpecPluginsItems0) UnmarshalBinary(b []byte) error {
	var res WorkflowSpecPluginsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowSpecWorkflowsItems0 workflow spec workflows items0
//
// swagger:model WorkflowSpecWorkflowsItems0
type WorkflowSpecWorkflowsItems0 struct {

	// name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Name string `json:"name,omitempty"`

	// steps
	Steps []*WorkflowSpecWorkflowsItems0StepsItems0 `json:"steps"`
}

// Validate validates this workflow spec workflows items0
func (m *WorkflowSpecWorkflowsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowSpecWorkflowsItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowSpecWorkflowsItems0) validateSteps(formats strfmt.Registry) error {
	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this workflow spec workflows items0 based on the context it is used
func (m *WorkflowSpecWorkflowsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowSpecWorkflowsItems0) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Steps); i++ {

		if m.Steps[i] != nil {
			if err := m.Steps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowSpecWorkflowsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowSpecWorkflowsItems0) UnmarshalBinary(b []byte) error {
	var res WorkflowSpecWorkflowsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowSpecWorkflowsItems0StepsItems0 workflow spec workflows items0 steps items0
//
// swagger:model WorkflowSpecWorkflowsItems0StepsItems0
type WorkflowSpecWorkflowsItems0StepsItems0 struct {

	// input
	Input []*WorkflowSpecWorkflowsItems0StepsItems0InputItems0 `json:"input"`

	// name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Name string `json:"name,omitempty"`

	// output
	Output []*WorkflowSpecWorkflowsItems0StepsItems0OutputItems0 `json:"output"`
}

// Validate validates this workflow spec workflows items0 steps items0
func (m *WorkflowSpecWorkflowsItems0StepsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowSpecWorkflowsItems0StepsItems0) validateInput(formats strfmt.Registry) error {
	if swag.IsZero(m.Input) { // not required
		return nil
	}

	for i := 0; i < len(m.Input); i++ {
		if swag.IsZero(m.Input[i]) { // not required
			continue
		}

		if m.Input[i] != nil {
			if err := m.Input[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("input" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowSpecWorkflowsItems0StepsItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowSpecWorkflowsItems0StepsItems0) validateOutput(formats strfmt.Registry) error {
	if swag.IsZero(m.Output) { // not required
		return nil
	}

	for i := 0; i < len(m.Output); i++ {
		if swag.IsZero(m.Output[i]) { // not required
			continue
		}

		if m.Output[i] != nil {
			if err := m.Output[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("output" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("output" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this workflow spec workflows items0 steps items0 based on the context it is used
func (m *WorkflowSpecWorkflowsItems0StepsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowSpecWorkflowsItems0StepsItems0) contextValidateInput(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Input); i++ {

		if m.Input[i] != nil {
			if err := m.Input[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("input" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowSpecWorkflowsItems0StepsItems0) contextValidateOutput(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Output); i++ {

		if m.Output[i] != nil {
			if err := m.Output[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("output" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("output" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowSpecWorkflowsItems0StepsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowSpecWorkflowsItems0StepsItems0) UnmarshalBinary(b []byte) error {
	var res WorkflowSpecWorkflowsItems0StepsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowSpecWorkflowsItems0StepsItems0InputItems0 workflow spec workflows items0 steps items0 input items0
//
// swagger:model WorkflowSpecWorkflowsItems0StepsItems0InputItems0
type WorkflowSpecWorkflowsItems0StepsItems0InputItems0 struct {

	// name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Name string `json:"name,omitempty"`

	// schema
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Schema string `json:"schema,omitempty"`
}

// Validate validates this workflow spec workflows items0 steps items0 input items0
func (m *WorkflowSpecWorkflowsItems0StepsItems0InputItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowSpecWorkflowsItems0StepsItems0InputItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowSpecWorkflowsItems0StepsItems0InputItems0) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if err := validate.Pattern("schema", "body", m.Schema, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this workflow spec workflows items0 steps items0 input items0 based on context it is used
func (m *WorkflowSpecWorkflowsItems0StepsItems0InputItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowSpecWorkflowsItems0StepsItems0InputItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowSpecWorkflowsItems0StepsItems0InputItems0) UnmarshalBinary(b []byte) error {
	var res WorkflowSpecWorkflowsItems0StepsItems0InputItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowSpecWorkflowsItems0StepsItems0OutputItems0 workflow spec workflows items0 steps items0 output items0
//
// swagger:model WorkflowSpecWorkflowsItems0StepsItems0OutputItems0
type WorkflowSpecWorkflowsItems0StepsItems0OutputItems0 struct {

	// name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Name string `json:"name,omitempty"`

	// schema
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Schema string `json:"schema,omitempty"`
}

// Validate validates this workflow spec workflows items0 steps items0 output items0
func (m *WorkflowSpecWorkflowsItems0StepsItems0OutputItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowSpecWorkflowsItems0StepsItems0OutputItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowSpecWorkflowsItems0StepsItems0OutputItems0) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if err := validate.Pattern("schema", "body", m.Schema, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this workflow spec workflows items0 steps items0 output items0 based on context it is used
func (m *WorkflowSpecWorkflowsItems0StepsItems0OutputItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowSpecWorkflowsItems0StepsItems0OutputItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowSpecWorkflowsItems0StepsItems0OutputItems0) UnmarshalBinary(b []byte) error {
	var res WorkflowSpecWorkflowsItems0StepsItems0OutputItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}