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

// ClusterDeploymentStatusEnum the model 'ClusterDeploymentStatusEnum'
type ClusterDeploymentStatusEnum string

// List of ClusterDeploymentStatusEnum
const (
	CLUSTERDEPLOYMENTSTATUSENUM_NEVER_DEPLOYED ClusterDeploymentStatusEnum = "NEVER_DEPLOYED"
	CLUSTERDEPLOYMENTSTATUSENUM_OUT_OF_DATE    ClusterDeploymentStatusEnum = "OUT_OF_DATE"
	CLUSTERDEPLOYMENTSTATUSENUM_UP_TO_DATE     ClusterDeploymentStatusEnum = "UP_TO_DATE"
)

// All allowed values of ClusterDeploymentStatusEnum enum
var AllowedClusterDeploymentStatusEnumEnumValues = []ClusterDeploymentStatusEnum{
	"NEVER_DEPLOYED",
	"OUT_OF_DATE",
	"UP_TO_DATE",
}

func (v *ClusterDeploymentStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterDeploymentStatusEnum(value)
	for _, existing := range AllowedClusterDeploymentStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterDeploymentStatusEnum", value)
}

// NewClusterDeploymentStatusEnumFromValue returns a pointer to a valid ClusterDeploymentStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterDeploymentStatusEnumFromValue(v string) (*ClusterDeploymentStatusEnum, error) {
	ev := ClusterDeploymentStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterDeploymentStatusEnum: valid values are %v", v, AllowedClusterDeploymentStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterDeploymentStatusEnum) IsValid() bool {
	for _, existing := range AllowedClusterDeploymentStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterDeploymentStatusEnum value
func (v ClusterDeploymentStatusEnum) Ptr() *ClusterDeploymentStatusEnum {
	return &v
}

type NullableClusterDeploymentStatusEnum struct {
	value *ClusterDeploymentStatusEnum
	isSet bool
}

func (v NullableClusterDeploymentStatusEnum) Get() *ClusterDeploymentStatusEnum {
	return v.value
}

func (v *NullableClusterDeploymentStatusEnum) Set(val *ClusterDeploymentStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDeploymentStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDeploymentStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDeploymentStatusEnum(val *ClusterDeploymentStatusEnum) *NullableClusterDeploymentStatusEnum {
	return &NullableClusterDeploymentStatusEnum{value: val, isSet: true}
}

func (v NullableClusterDeploymentStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterDeploymentStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
