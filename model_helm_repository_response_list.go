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

// checks if the HelmRepositoryResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmRepositoryResponseList{}

// HelmRepositoryResponseList struct for HelmRepositoryResponseList
type HelmRepositoryResponseList struct {
	Results              []HelmRepositoryResponse `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HelmRepositoryResponseList HelmRepositoryResponseList

// NewHelmRepositoryResponseList instantiates a new HelmRepositoryResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmRepositoryResponseList() *HelmRepositoryResponseList {
	this := HelmRepositoryResponseList{}
	return &this
}

// NewHelmRepositoryResponseListWithDefaults instantiates a new HelmRepositoryResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmRepositoryResponseListWithDefaults() *HelmRepositoryResponseList {
	this := HelmRepositoryResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *HelmRepositoryResponseList) GetResults() []HelmRepositoryResponse {
	if o == nil || IsNil(o.Results) {
		var ret []HelmRepositoryResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRepositoryResponseList) GetResultsOk() ([]HelmRepositoryResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *HelmRepositoryResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []HelmRepositoryResponse and assigns it to the Results field.
func (o *HelmRepositoryResponseList) SetResults(v []HelmRepositoryResponse) {
	o.Results = v
}

func (o HelmRepositoryResponseList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmRepositoryResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmRepositoryResponseList) UnmarshalJSON(data []byte) (err error) {
	varHelmRepositoryResponseList := _HelmRepositoryResponseList{}

	err = json.Unmarshal(data, &varHelmRepositoryResponseList)

	if err != nil {
		return err
	}

	*o = HelmRepositoryResponseList(varHelmRepositoryResponseList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmRepositoryResponseList struct {
	value *HelmRepositoryResponseList
	isSet bool
}

func (v NullableHelmRepositoryResponseList) Get() *HelmRepositoryResponseList {
	return v.value
}

func (v *NullableHelmRepositoryResponseList) Set(val *HelmRepositoryResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmRepositoryResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmRepositoryResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmRepositoryResponseList(val *HelmRepositoryResponseList) *NullableHelmRepositoryResponseList {
	return &NullableHelmRepositoryResponseList{value: val, isSet: true}
}

func (v NullableHelmRepositoryResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmRepositoryResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
