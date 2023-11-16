/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"fmt"
)

// ServiceTypeForVariableEnum type of the service
type ServiceTypeForVariableEnum string

// List of ServiceTypeForVariableEnum
const (
	SERVICETYPEFORVARIABLEENUM_APPLICATION ServiceTypeForVariableEnum = "APPLICATION"
	SERVICETYPEFORVARIABLEENUM_CONTAINER   ServiceTypeForVariableEnum = "CONTAINER"
	SERVICETYPEFORVARIABLEENUM_JOB         ServiceTypeForVariableEnum = "JOB"
	SERVICETYPEFORVARIABLEENUM_HELM        ServiceTypeForVariableEnum = "HELM"
)

// All allowed values of ServiceTypeForVariableEnum enum
var AllowedServiceTypeForVariableEnumEnumValues = []ServiceTypeForVariableEnum{
	"APPLICATION",
	"CONTAINER",
	"JOB",
	"HELM",
}

func (v *ServiceTypeForVariableEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceTypeForVariableEnum(value)
	for _, existing := range AllowedServiceTypeForVariableEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceTypeForVariableEnum", value)
}

// NewServiceTypeForVariableEnumFromValue returns a pointer to a valid ServiceTypeForVariableEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceTypeForVariableEnumFromValue(v string) (*ServiceTypeForVariableEnum, error) {
	ev := ServiceTypeForVariableEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceTypeForVariableEnum: valid values are %v", v, AllowedServiceTypeForVariableEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceTypeForVariableEnum) IsValid() bool {
	for _, existing := range AllowedServiceTypeForVariableEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceTypeForVariableEnum value
func (v ServiceTypeForVariableEnum) Ptr() *ServiceTypeForVariableEnum {
	return &v
}

type NullableServiceTypeForVariableEnum struct {
	value *ServiceTypeForVariableEnum
	isSet bool
}

func (v NullableServiceTypeForVariableEnum) Get() *ServiceTypeForVariableEnum {
	return v.value
}

func (v *NullableServiceTypeForVariableEnum) Set(val *ServiceTypeForVariableEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTypeForVariableEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTypeForVariableEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTypeForVariableEnum(val *ServiceTypeForVariableEnum) *NullableServiceTypeForVariableEnum {
	return &NullableServiceTypeForVariableEnum{value: val, isSet: true}
}

func (v NullableServiceTypeForVariableEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTypeForVariableEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
