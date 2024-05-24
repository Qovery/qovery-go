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

// LabelsGroupAssociatedItemType Labels Group Associated Item Type
type LabelsGroupAssociatedItemType string

// List of LabelsGroupAssociatedItemType
const (
	LABELSGROUPASSOCIATEDITEMTYPE_APPLICATION LabelsGroupAssociatedItemType = "APPLICATION"
	LABELSGROUPASSOCIATEDITEMTYPE_DATABASE    LabelsGroupAssociatedItemType = "DATABASE"
	LABELSGROUPASSOCIATEDITEMTYPE_CONTAINER   LabelsGroupAssociatedItemType = "CONTAINER"
	LABELSGROUPASSOCIATEDITEMTYPE_LIFECYCLE   LabelsGroupAssociatedItemType = "LIFECYCLE"
	LABELSGROUPASSOCIATEDITEMTYPE_ENVIRONMENT LabelsGroupAssociatedItemType = "ENVIRONMENT"
	LABELSGROUPASSOCIATEDITEMTYPE_CLUSTER     LabelsGroupAssociatedItemType = "CLUSTER"
	LABELSGROUPASSOCIATEDITEMTYPE_CRON        LabelsGroupAssociatedItemType = "CRON"
)

// All allowed values of LabelsGroupAssociatedItemType enum
var AllowedLabelsGroupAssociatedItemTypeEnumValues = []LabelsGroupAssociatedItemType{
	"APPLICATION",
	"DATABASE",
	"CONTAINER",
	"LIFECYCLE",
	"ENVIRONMENT",
	"CLUSTER",
	"CRON",
}

func (v *LabelsGroupAssociatedItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LabelsGroupAssociatedItemType(value)
	for _, existing := range AllowedLabelsGroupAssociatedItemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LabelsGroupAssociatedItemType", value)
}

// NewLabelsGroupAssociatedItemTypeFromValue returns a pointer to a valid LabelsGroupAssociatedItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLabelsGroupAssociatedItemTypeFromValue(v string) (*LabelsGroupAssociatedItemType, error) {
	ev := LabelsGroupAssociatedItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LabelsGroupAssociatedItemType: valid values are %v", v, AllowedLabelsGroupAssociatedItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LabelsGroupAssociatedItemType) IsValid() bool {
	for _, existing := range AllowedLabelsGroupAssociatedItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LabelsGroupAssociatedItemType value
func (v LabelsGroupAssociatedItemType) Ptr() *LabelsGroupAssociatedItemType {
	return &v
}

type NullableLabelsGroupAssociatedItemType struct {
	value *LabelsGroupAssociatedItemType
	isSet bool
}

func (v NullableLabelsGroupAssociatedItemType) Get() *LabelsGroupAssociatedItemType {
	return v.value
}

func (v *NullableLabelsGroupAssociatedItemType) Set(val *LabelsGroupAssociatedItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelsGroupAssociatedItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelsGroupAssociatedItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelsGroupAssociatedItemType(val *LabelsGroupAssociatedItemType) *NullableLabelsGroupAssociatedItemType {
	return &NullableLabelsGroupAssociatedItemType{value: val, isSet: true}
}

func (v NullableLabelsGroupAssociatedItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelsGroupAssociatedItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
