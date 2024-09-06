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

// checks if the KubernetesServiceResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesServiceResponseList{}

// KubernetesServiceResponseList struct for KubernetesServiceResponseList
type KubernetesServiceResponseList struct {
	Results              []KubernetesService `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesServiceResponseList KubernetesServiceResponseList

// NewKubernetesServiceResponseList instantiates a new KubernetesServiceResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesServiceResponseList() *KubernetesServiceResponseList {
	this := KubernetesServiceResponseList{}
	return &this
}

// NewKubernetesServiceResponseListWithDefaults instantiates a new KubernetesServiceResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesServiceResponseListWithDefaults() *KubernetesServiceResponseList {
	this := KubernetesServiceResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *KubernetesServiceResponseList) GetResults() []KubernetesService {
	if o == nil || IsNil(o.Results) {
		var ret []KubernetesService
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesServiceResponseList) GetResultsOk() ([]KubernetesService, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *KubernetesServiceResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []KubernetesService and assigns it to the Results field.
func (o *KubernetesServiceResponseList) SetResults(v []KubernetesService) {
	o.Results = v
}

func (o KubernetesServiceResponseList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesServiceResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesServiceResponseList) UnmarshalJSON(data []byte) (err error) {
	varKubernetesServiceResponseList := _KubernetesServiceResponseList{}

	err = json.Unmarshal(data, &varKubernetesServiceResponseList)

	if err != nil {
		return err
	}

	*o = KubernetesServiceResponseList(varKubernetesServiceResponseList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesServiceResponseList struct {
	value *KubernetesServiceResponseList
	isSet bool
}

func (v NullableKubernetesServiceResponseList) Get() *KubernetesServiceResponseList {
	return v.value
}

func (v *NullableKubernetesServiceResponseList) Set(val *KubernetesServiceResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesServiceResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesServiceResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesServiceResponseList(val *KubernetesServiceResponseList) *NullableKubernetesServiceResponseList {
	return &NullableKubernetesServiceResponseList{value: val, isSet: true}
}

func (v NullableKubernetesServiceResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesServiceResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
