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

// EnvironmentLogsDetailsStage struct for EnvironmentLogsDetailsStage
type EnvironmentLogsDetailsStage struct {
	Step *string `json:"step,omitempty"`
}

// NewEnvironmentLogsDetailsStage instantiates a new EnvironmentLogsDetailsStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentLogsDetailsStage() *EnvironmentLogsDetailsStage {
	this := EnvironmentLogsDetailsStage{}
	return &this
}

// NewEnvironmentLogsDetailsStageWithDefaults instantiates a new EnvironmentLogsDetailsStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentLogsDetailsStageWithDefaults() *EnvironmentLogsDetailsStage {
	this := EnvironmentLogsDetailsStage{}
	return &this
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *EnvironmentLogsDetailsStage) GetStep() string {
	if o == nil || o.Step == nil {
		var ret string
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogsDetailsStage) GetStepOk() (*string, bool) {
	if o == nil || o.Step == nil {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *EnvironmentLogsDetailsStage) HasStep() bool {
	if o != nil && o.Step != nil {
		return true
	}

	return false
}

// SetStep gets a reference to the given string and assigns it to the Step field.
func (o *EnvironmentLogsDetailsStage) SetStep(v string) {
	o.Step = &v
}

func (o EnvironmentLogsDetailsStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Step != nil {
		toSerialize["step"] = o.Step
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentLogsDetailsStage struct {
	value *EnvironmentLogsDetailsStage
	isSet bool
}

func (v NullableEnvironmentLogsDetailsStage) Get() *EnvironmentLogsDetailsStage {
	return v.value
}

func (v *NullableEnvironmentLogsDetailsStage) Set(val *EnvironmentLogsDetailsStage) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentLogsDetailsStage) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentLogsDetailsStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentLogsDetailsStage(val *EnvironmentLogsDetailsStage) *NullableEnvironmentLogsDetailsStage {
	return &NullableEnvironmentLogsDetailsStage{value: val, isSet: true}
}

func (v NullableEnvironmentLogsDetailsStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentLogsDetailsStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
