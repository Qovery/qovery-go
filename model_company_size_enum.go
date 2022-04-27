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

// CompanySizeEnum the model 'CompanySizeEnum'
type CompanySizeEnum string

// List of CompanySizeEnum
const (
	COMPANYSIZEENUM__1_10    CompanySizeEnum = "1-10"
	COMPANYSIZEENUM__11_50   CompanySizeEnum = "11-50"
	COMPANYSIZEENUM__51_200  CompanySizeEnum = "51-200"
	COMPANYSIZEENUM__201_500 CompanySizeEnum = "201-500"
	COMPANYSIZEENUM__500     CompanySizeEnum = "500+"
)

// All allowed values of CompanySizeEnum enum
var AllowedCompanySizeEnumEnumValues = []CompanySizeEnum{
	"1-10",
	"11-50",
	"51-200",
	"201-500",
	"500+",
}

func (v *CompanySizeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CompanySizeEnum(value)
	for _, existing := range AllowedCompanySizeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CompanySizeEnum", value)
}

// NewCompanySizeEnumFromValue returns a pointer to a valid CompanySizeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCompanySizeEnumFromValue(v string) (*CompanySizeEnum, error) {
	ev := CompanySizeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CompanySizeEnum: valid values are %v", v, AllowedCompanySizeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CompanySizeEnum) IsValid() bool {
	for _, existing := range AllowedCompanySizeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompanySizeEnum value
func (v CompanySizeEnum) Ptr() *CompanySizeEnum {
	return &v
}

type NullableCompanySizeEnum struct {
	value *CompanySizeEnum
	isSet bool
}

func (v NullableCompanySizeEnum) Get() *CompanySizeEnum {
	return v.value
}

func (v *NullableCompanySizeEnum) Set(val *CompanySizeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanySizeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanySizeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanySizeEnum(val *CompanySizeEnum) *NullableCompanySizeEnum {
	return &NullableCompanySizeEnum{value: val, isSet: true}
}

func (v NullableCompanySizeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanySizeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
