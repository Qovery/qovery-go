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

// DeploymentHistoryStatusEnum the model 'DeploymentHistoryStatusEnum'
type DeploymentHistoryStatusEnum string

// List of DeploymentHistoryStatusEnum
const (
	DEPLOYMENTHISTORYSTATUSENUM_SUCCESS DeploymentHistoryStatusEnum = "SUCCESS"
	DEPLOYMENTHISTORYSTATUSENUM_FAILED  DeploymentHistoryStatusEnum = "FAILED"
)

// All allowed values of DeploymentHistoryStatusEnum enum
var AllowedDeploymentHistoryStatusEnumEnumValues = []DeploymentHistoryStatusEnum{
	"SUCCESS",
	"FAILED",
}

func (v *DeploymentHistoryStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeploymentHistoryStatusEnum(value)
	for _, existing := range AllowedDeploymentHistoryStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeploymentHistoryStatusEnum", value)
}

// NewDeploymentHistoryStatusEnumFromValue returns a pointer to a valid DeploymentHistoryStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeploymentHistoryStatusEnumFromValue(v string) (*DeploymentHistoryStatusEnum, error) {
	ev := DeploymentHistoryStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeploymentHistoryStatusEnum: valid values are %v", v, AllowedDeploymentHistoryStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeploymentHistoryStatusEnum) IsValid() bool {
	for _, existing := range AllowedDeploymentHistoryStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentHistoryStatusEnum value
func (v DeploymentHistoryStatusEnum) Ptr() *DeploymentHistoryStatusEnum {
	return &v
}

type NullableDeploymentHistoryStatusEnum struct {
	value *DeploymentHistoryStatusEnum
	isSet bool
}

func (v NullableDeploymentHistoryStatusEnum) Get() *DeploymentHistoryStatusEnum {
	return v.value
}

func (v *NullableDeploymentHistoryStatusEnum) Set(val *DeploymentHistoryStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryStatusEnum(val *DeploymentHistoryStatusEnum) *NullableDeploymentHistoryStatusEnum {
	return &NullableDeploymentHistoryStatusEnum{value: val, isSet: true}
}

func (v NullableDeploymentHistoryStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
