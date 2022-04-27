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

// ServiceTypeEnum type of the service (application, database, job, gateway...)
type ServiceTypeEnum string

// List of ServiceTypeEnum
const (
	SERVICETYPEENUM_APPLICATION ServiceTypeEnum = "APPLICATION"
	SERVICETYPEENUM_DATABASE    ServiceTypeEnum = "DATABASE"
)

// All allowed values of ServiceTypeEnum enum
var AllowedServiceTypeEnumEnumValues = []ServiceTypeEnum{
	"APPLICATION",
	"DATABASE",
}

func (v *ServiceTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceTypeEnum(value)
	for _, existing := range AllowedServiceTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceTypeEnum", value)
}

// NewServiceTypeEnumFromValue returns a pointer to a valid ServiceTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceTypeEnumFromValue(v string) (*ServiceTypeEnum, error) {
	ev := ServiceTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceTypeEnum: valid values are %v", v, AllowedServiceTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceTypeEnum) IsValid() bool {
	for _, existing := range AllowedServiceTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceTypeEnum value
func (v ServiceTypeEnum) Ptr() *ServiceTypeEnum {
	return &v
}

type NullableServiceTypeEnum struct {
	value *ServiceTypeEnum
	isSet bool
}

func (v NullableServiceTypeEnum) Get() *ServiceTypeEnum {
	return v.value
}

func (v *NullableServiceTypeEnum) Set(val *ServiceTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTypeEnum(val *ServiceTypeEnum) *NullableServiceTypeEnum {
	return &NullableServiceTypeEnum{value: val, isSet: true}
}

func (v NullableServiceTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
