// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServicePort service port
// swagger:model servicePort
type ServicePort struct {

	// The name of this port within the service
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// The port that will be exposed by this service
	// Required: true
	Port int64 `json:"port"`

	// The IP protocol for this port. Supports "TCP" and "UDP". Default is TCP
	// Min Length: 1
	// Enum: [TCP UDP]
	Protocol string `json:"protocol,omitempty"`

	// Number or name of the port to access on the pods targeted by the service
	TargetPort int64 `json:"targetPort,omitempty"`
}

// Validate validates this service port
func (m *ServicePort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServicePort) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *ServicePort) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", int64(m.Port)); err != nil {
		return err
	}

	return nil
}

var servicePortTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TCP","UDP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		servicePortTypeProtocolPropEnum = append(servicePortTypeProtocolPropEnum, v)
	}
}

const (

	// ServicePortProtocolTCP captures enum value "TCP"
	ServicePortProtocolTCP string = "TCP"

	// ServicePortProtocolUDP captures enum value "UDP"
	ServicePortProtocolUDP string = "UDP"
)

// prop value enum
func (m *ServicePort) validateProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, servicePortTypeProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServicePort) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	if err := validate.MinLength("protocol", "body", string(m.Protocol), 1); err != nil {
		return err
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServicePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServicePort) UnmarshalBinary(b []byte) error {
	var res ServicePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
