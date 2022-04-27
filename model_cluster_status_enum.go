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

// ClusterStatusEnum the model 'ClusterStatusEnum'
type ClusterStatusEnum string

// List of ClusterStatusEnum
const (
	CLUSTERSTATUSENUM_READY            ClusterStatusEnum = "READY"
	CLUSTERSTATUSENUM_QUEUED           ClusterStatusEnum = "QUEUED"
	CLUSTERSTATUSENUM_STOP_QUEUED      ClusterStatusEnum = "STOP_QUEUED"
	CLUSTERSTATUSENUM_DELETE_QUEUED    ClusterStatusEnum = "DELETE_QUEUED"
	CLUSTERSTATUSENUM_BUILDING         ClusterStatusEnum = "BUILDING"
	CLUSTERSTATUSENUM_BUILD_ERROR      ClusterStatusEnum = "BUILD_ERROR"
	CLUSTERSTATUSENUM_BUILT            ClusterStatusEnum = "BUILT"
	CLUSTERSTATUSENUM_DEPLOYING        ClusterStatusEnum = "DEPLOYING"
	CLUSTERSTATUSENUM_DEPLOYMENT_ERROR ClusterStatusEnum = "DEPLOYMENT_ERROR"
	CLUSTERSTATUSENUM_DEPLOYED         ClusterStatusEnum = "DEPLOYED"
	CLUSTERSTATUSENUM_STOPPING         ClusterStatusEnum = "STOPPING"
	CLUSTERSTATUSENUM_STOP_ERROR       ClusterStatusEnum = "STOP_ERROR"
	CLUSTERSTATUSENUM_STOPPED          ClusterStatusEnum = "STOPPED"
	CLUSTERSTATUSENUM_DELETING         ClusterStatusEnum = "DELETING"
	CLUSTERSTATUSENUM_DELETE_ERROR     ClusterStatusEnum = "DELETE_ERROR"
	CLUSTERSTATUSENUM_DELETED          ClusterStatusEnum = "DELETED"
	CLUSTERSTATUSENUM_RUNNING          ClusterStatusEnum = "RUNNING"
	CLUSTERSTATUSENUM_RUNNING_ERROR    ClusterStatusEnum = "RUNNING_ERROR"
)

// All allowed values of ClusterStatusEnum enum
var AllowedClusterStatusEnumEnumValues = []ClusterStatusEnum{
	"READY",
	"QUEUED",
	"STOP_QUEUED",
	"DELETE_QUEUED",
	"BUILDING",
	"BUILD_ERROR",
	"BUILT",
	"DEPLOYING",
	"DEPLOYMENT_ERROR",
	"DEPLOYED",
	"STOPPING",
	"STOP_ERROR",
	"STOPPED",
	"DELETING",
	"DELETE_ERROR",
	"DELETED",
	"RUNNING",
	"RUNNING_ERROR",
}

func (v *ClusterStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterStatusEnum(value)
	for _, existing := range AllowedClusterStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterStatusEnum", value)
}

// NewClusterStatusEnumFromValue returns a pointer to a valid ClusterStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterStatusEnumFromValue(v string) (*ClusterStatusEnum, error) {
	ev := ClusterStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterStatusEnum: valid values are %v", v, AllowedClusterStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterStatusEnum) IsValid() bool {
	for _, existing := range AllowedClusterStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterStatusEnum value
func (v ClusterStatusEnum) Ptr() *ClusterStatusEnum {
	return &v
}

type NullableClusterStatusEnum struct {
	value *ClusterStatusEnum
	isSet bool
}

func (v NullableClusterStatusEnum) Get() *ClusterStatusEnum {
	return v.value
}

func (v *NullableClusterStatusEnum) Set(val *ClusterStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStatusEnum(val *ClusterStatusEnum) *NullableClusterStatusEnum {
	return &NullableClusterStatusEnum{value: val, isSet: true}
}

func (v NullableClusterStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
