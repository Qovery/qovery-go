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

// Stage struct for Stage
type Stage struct {
	Id string `json:"id"`
	// stage name
	Name    string            `json:"name"`
	Metrics *StageStepMetrics `json:"metrics,omitempty"`
}

// NewStage instantiates a new Stage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStage(id string, name string) *Stage {
	this := Stage{}
	this.Id = id
	this.Name = name
	return &this
}

// NewStageWithDefaults instantiates a new Stage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStageWithDefaults() *Stage {
	this := Stage{}
	return &this
}

// GetId returns the Id field value
func (o *Stage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Stage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Stage) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Stage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Stage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Stage) SetName(v string) {
	o.Name = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *Stage) GetMetrics() StageStepMetrics {
	if o == nil || o.Metrics == nil {
		var ret StageStepMetrics
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stage) GetMetricsOk() (*StageStepMetrics, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *Stage) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given StageStepMetrics and assigns it to the Metrics field.
func (o *Stage) SetMetrics(v StageStepMetrics) {
	o.Metrics = &v
}

func (o Stage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	return json.Marshal(toSerialize)
}

type NullableStage struct {
	value *Stage
	isSet bool
}

func (v NullableStage) Get() *Stage {
	return v.value
}

func (v *NullableStage) Set(val *Stage) {
	v.value = val
	v.isSet = true
}

func (v NullableStage) IsSet() bool {
	return v.isSet
}

func (v *NullableStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStage(val *Stage) *NullableStage {
	return &NullableStage{value: val, isSet: true}
}

func (v NullableStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
