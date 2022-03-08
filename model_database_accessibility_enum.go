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

// DatabaseAccessibilityEnum the model 'DatabaseAccessibilityEnum'
type DatabaseAccessibilityEnum string

// List of DatabaseAccessibilityEnum
const (
	PUBLIC  DatabaseAccessibilityEnum = "PUBLIC"
	PRIVATE DatabaseAccessibilityEnum = "PRIVATE"
)

// All allowed values of DatabaseAccessibilityEnum enum
var AllowedDatabaseAccessibilityEnumEnumValues = []DatabaseAccessibilityEnum{
	"PUBLIC",
	"PRIVATE",
}

func (v *DatabaseAccessibilityEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DatabaseAccessibilityEnum(value)
	for _, existing := range AllowedDatabaseAccessibilityEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DatabaseAccessibilityEnum", value)
}

// NewDatabaseAccessibilityEnumFromValue returns a pointer to a valid DatabaseAccessibilityEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDatabaseAccessibilityEnumFromValue(v string) (*DatabaseAccessibilityEnum, error) {
	ev := DatabaseAccessibilityEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DatabaseAccessibilityEnum: valid values are %v", v, AllowedDatabaseAccessibilityEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DatabaseAccessibilityEnum) IsValid() bool {
	for _, existing := range AllowedDatabaseAccessibilityEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatabaseAccessibilityEnum value
func (v DatabaseAccessibilityEnum) Ptr() *DatabaseAccessibilityEnum {
	return &v
}

type NullableDatabaseAccessibilityEnum struct {
	value *DatabaseAccessibilityEnum
	isSet bool
}

func (v NullableDatabaseAccessibilityEnum) Get() *DatabaseAccessibilityEnum {
	return v.value
}

func (v *NullableDatabaseAccessibilityEnum) Set(val *DatabaseAccessibilityEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseAccessibilityEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseAccessibilityEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseAccessibilityEnum(val *DatabaseAccessibilityEnum) *NullableDatabaseAccessibilityEnum {
	return &NullableDatabaseAccessibilityEnum{value: val, isSet: true}
}

func (v NullableDatabaseAccessibilityEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseAccessibilityEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
