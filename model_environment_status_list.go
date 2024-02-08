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

// checks if the EnvironmentStatusList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentStatusList{}

// EnvironmentStatusList struct for EnvironmentStatusList
type EnvironmentStatusList struct {
	Results              []EnvironmentStatus `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnvironmentStatusList EnvironmentStatusList

// NewEnvironmentStatusList instantiates a new EnvironmentStatusList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentStatusList() *EnvironmentStatusList {
	this := EnvironmentStatusList{}
	return &this
}

// NewEnvironmentStatusListWithDefaults instantiates a new EnvironmentStatusList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentStatusListWithDefaults() *EnvironmentStatusList {
	this := EnvironmentStatusList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *EnvironmentStatusList) GetResults() []EnvironmentStatus {
	if o == nil || IsNil(o.Results) {
		var ret []EnvironmentStatus
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusList) GetResultsOk() ([]EnvironmentStatus, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EnvironmentStatusList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EnvironmentStatus and assigns it to the Results field.
func (o *EnvironmentStatusList) SetResults(v []EnvironmentStatus) {
	o.Results = v
}

func (o EnvironmentStatusList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentStatusList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvironmentStatusList) UnmarshalJSON(data []byte) (err error) {
	varEnvironmentStatusList := _EnvironmentStatusList{}

	err = json.Unmarshal(data, &varEnvironmentStatusList)

	if err != nil {
		return err
	}

	*o = EnvironmentStatusList(varEnvironmentStatusList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnvironmentStatusList struct {
	value *EnvironmentStatusList
	isSet bool
}

func (v NullableEnvironmentStatusList) Get() *EnvironmentStatusList {
	return v.value
}

func (v *NullableEnvironmentStatusList) Set(val *EnvironmentStatusList) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentStatusList) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentStatusList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentStatusList(val *EnvironmentStatusList) *NullableEnvironmentStatusList {
	return &NullableEnvironmentStatusList{value: val, isSet: true}
}

func (v NullableEnvironmentStatusList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentStatusList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
