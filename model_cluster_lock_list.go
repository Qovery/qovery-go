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

// checks if the ClusterLockList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterLockList{}

// ClusterLockList struct for ClusterLockList
type ClusterLockList struct {
	Results              []ClusterLock `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClusterLockList ClusterLockList

// NewClusterLockList instantiates a new ClusterLockList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterLockList() *ClusterLockList {
	this := ClusterLockList{}
	return &this
}

// NewClusterLockListWithDefaults instantiates a new ClusterLockList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterLockListWithDefaults() *ClusterLockList {
	this := ClusterLockList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ClusterLockList) GetResults() []ClusterLock {
	if o == nil || IsNil(o.Results) {
		var ret []ClusterLock
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLockList) GetResultsOk() ([]ClusterLock, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ClusterLockList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ClusterLock and assigns it to the Results field.
func (o *ClusterLockList) SetResults(v []ClusterLock) {
	o.Results = v
}

func (o ClusterLockList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterLockList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClusterLockList) UnmarshalJSON(data []byte) (err error) {
	varClusterLockList := _ClusterLockList{}

	err = json.Unmarshal(data, &varClusterLockList)

	if err != nil {
		return err
	}

	*o = ClusterLockList(varClusterLockList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClusterLockList struct {
	value *ClusterLockList
	isSet bool
}

func (v NullableClusterLockList) Get() *ClusterLockList {
	return v.value
}

func (v *NullableClusterLockList) Set(val *ClusterLockList) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterLockList) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterLockList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterLockList(val *ClusterLockList) *NullableClusterLockList {
	return &NullableClusterLockList{value: val, isSet: true}
}

func (v NullableClusterLockList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterLockList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
