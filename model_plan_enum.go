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

// PlanEnum the model 'PlanEnum'
type PlanEnum string

// List of PlanEnum
const (
	PLANENUM_COMMUNITY    PlanEnum = "COMMUNITY"
	PLANENUM_FREE         PlanEnum = "FREE"
	PLANENUM_PROFESSIONAL PlanEnum = "PROFESSIONAL"
	PLANENUM_BUSINESS     PlanEnum = "BUSINESS"
	PLANENUM_ENTERPRISE   PlanEnum = "ENTERPRISE"
)

// All allowed values of PlanEnum enum
var AllowedPlanEnumEnumValues = []PlanEnum{
	"COMMUNITY",
	"FREE",
	"PROFESSIONAL",
	"BUSINESS",
	"ENTERPRISE",
}

func (v *PlanEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlanEnum(value)
	for _, existing := range AllowedPlanEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlanEnum", value)
}

// NewPlanEnumFromValue returns a pointer to a valid PlanEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlanEnumFromValue(v string) (*PlanEnum, error) {
	ev := PlanEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlanEnum: valid values are %v", v, AllowedPlanEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlanEnum) IsValid() bool {
	for _, existing := range AllowedPlanEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlanEnum value
func (v PlanEnum) Ptr() *PlanEnum {
	return &v
}

type NullablePlanEnum struct {
	value *PlanEnum
	isSet bool
}

func (v NullablePlanEnum) Get() *PlanEnum {
	return v.value
}

func (v *NullablePlanEnum) Set(val *PlanEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanEnum(val *PlanEnum) *NullablePlanEnum {
	return &NullablePlanEnum{value: val, isSet: true}
}

func (v NullablePlanEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
