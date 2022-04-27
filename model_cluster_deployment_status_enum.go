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
	CLUSTERDEPLOYMENTSTATUSENUM_READY            ClusterDeploymentStatusEnum = "READY"
	CLUSTERDEPLOYMENTSTATUSENUM_QUEUED           ClusterDeploymentStatusEnum = "QUEUED"
	CLUSTERDEPLOYMENTSTATUSENUM_STOP_QUEUED      ClusterDeploymentStatusEnum = "STOP_QUEUED"
	CLUSTERDEPLOYMENTSTATUSENUM_DELETE_QUEUED    ClusterDeploymentStatusEnum = "DELETE_QUEUED"
	CLUSTERDEPLOYMENTSTATUSENUM_BUILDING         ClusterDeploymentStatusEnum = "BUILDING"
	CLUSTERDEPLOYMENTSTATUSENUM_BUILD_ERROR      ClusterDeploymentStatusEnum = "BUILD_ERROR"
	CLUSTERDEPLOYMENTSTATUSENUM_BUILT            ClusterDeploymentStatusEnum = "BUILT"
	CLUSTERDEPLOYMENTSTATUSENUM_DEPLOYING        ClusterDeploymentStatusEnum = "DEPLOYING"
	CLUSTERDEPLOYMENTSTATUSENUM_DEPLOYMENT_ERROR ClusterDeploymentStatusEnum = "DEPLOYMENT_ERROR"
	CLUSTERDEPLOYMENTSTATUSENUM_DEPLOYED         ClusterDeploymentStatusEnum = "DEPLOYED"
	CLUSTERDEPLOYMENTSTATUSENUM_STOPPING         ClusterDeploymentStatusEnum = "STOPPING"
	CLUSTERDEPLOYMENTSTATUSENUM_STOP_ERROR       ClusterDeploymentStatusEnum = "STOP_ERROR"
	CLUSTERDEPLOYMENTSTATUSENUM_STOPPED          ClusterDeploymentStatusEnum = "STOPPED"
	CLUSTERDEPLOYMENTSTATUSENUM_DELETING         ClusterDeploymentStatusEnum = "DELETING"
	CLUSTERDEPLOYMENTSTATUSENUM_DELETE_ERROR     ClusterDeploymentStatusEnum = "DELETE_ERROR"
	CLUSTERDEPLOYMENTSTATUSENUM_DELETED          ClusterDeploymentStatusEnum = "DELETED"
	CLUSTERDEPLOYMENTSTATUSENUM_RUNNING          ClusterDeploymentStatusEnum = "RUNNING"
	CLUSTERDEPLOYMENTSTATUSENUM_RUNNING_ERROR    ClusterDeploymentStatusEnum = "RUNNING_ERROR"
)

// All allowed values of ClusterDeploymentStatusEnum enum
var AllowedClusterDeploymentStatusEnumEnumValues = []ClusterDeploymentStatusEnum{
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
