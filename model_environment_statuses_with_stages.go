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
)

// checks if the EnvironmentStatusesWithStages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentStatusesWithStages{}

// EnvironmentStatusesWithStages struct for EnvironmentStatusesWithStages
type EnvironmentStatusesWithStages struct {
	Environment *EnvironmentStatus                    `json:"environment,omitempty"`
	Stages      []DeploymentStageWithServicesStatuses `json:"stages,omitempty"`
}

// NewEnvironmentStatusesWithStages instantiates a new EnvironmentStatusesWithStages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentStatusesWithStages() *EnvironmentStatusesWithStages {
	this := EnvironmentStatusesWithStages{}
	return &this
}

// NewEnvironmentStatusesWithStagesWithDefaults instantiates a new EnvironmentStatusesWithStages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentStatusesWithStagesWithDefaults() *EnvironmentStatusesWithStages {
	this := EnvironmentStatusesWithStages{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *EnvironmentStatusesWithStages) GetEnvironment() EnvironmentStatus {
	if o == nil || IsNil(o.Environment) {
		var ret EnvironmentStatus
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusesWithStages) GetEnvironmentOk() (*EnvironmentStatus, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *EnvironmentStatusesWithStages) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given EnvironmentStatus and assigns it to the Environment field.
func (o *EnvironmentStatusesWithStages) SetEnvironment(v EnvironmentStatus) {
	o.Environment = &v
}

// GetStages returns the Stages field value if set, zero value otherwise.
func (o *EnvironmentStatusesWithStages) GetStages() []DeploymentStageWithServicesStatuses {
	if o == nil || IsNil(o.Stages) {
		var ret []DeploymentStageWithServicesStatuses
		return ret
	}
	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusesWithStages) GetStagesOk() ([]DeploymentStageWithServicesStatuses, bool) {
	if o == nil || IsNil(o.Stages) {
		return nil, false
	}
	return o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *EnvironmentStatusesWithStages) HasStages() bool {
	if o != nil && !IsNil(o.Stages) {
		return true
	}

	return false
}

// SetStages gets a reference to the given []DeploymentStageWithServicesStatuses and assigns it to the Stages field.
func (o *EnvironmentStatusesWithStages) SetStages(v []DeploymentStageWithServicesStatuses) {
	o.Stages = v
}

func (o EnvironmentStatusesWithStages) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentStatusesWithStages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Stages) {
		toSerialize["stages"] = o.Stages
	}
	return toSerialize, nil
}

type NullableEnvironmentStatusesWithStages struct {
	value *EnvironmentStatusesWithStages
	isSet bool
}

func (v NullableEnvironmentStatusesWithStages) Get() *EnvironmentStatusesWithStages {
	return v.value
}

func (v *NullableEnvironmentStatusesWithStages) Set(val *EnvironmentStatusesWithStages) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentStatusesWithStages) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentStatusesWithStages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentStatusesWithStages(val *EnvironmentStatusesWithStages) *NullableEnvironmentStatusesWithStages {
	return &NullableEnvironmentStatusesWithStages{value: val, isSet: true}
}

func (v NullableEnvironmentStatusesWithStages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentStatusesWithStages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
