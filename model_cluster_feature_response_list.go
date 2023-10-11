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

// checks if the ClusterFeatureResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterFeatureResponseList{}

// ClusterFeatureResponseList struct for ClusterFeatureResponseList
type ClusterFeatureResponseList struct {
	Results []ClusterFeature `json:"results,omitempty"`
}

// NewClusterFeatureResponseList instantiates a new ClusterFeatureResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterFeatureResponseList() *ClusterFeatureResponseList {
	this := ClusterFeatureResponseList{}
	return &this
}

// NewClusterFeatureResponseListWithDefaults instantiates a new ClusterFeatureResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterFeatureResponseListWithDefaults() *ClusterFeatureResponseList {
	this := ClusterFeatureResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ClusterFeatureResponseList) GetResults() []ClusterFeature {
	if o == nil || IsNil(o.Results) {
		var ret []ClusterFeature
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeatureResponseList) GetResultsOk() ([]ClusterFeature, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ClusterFeatureResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ClusterFeature and assigns it to the Results field.
func (o *ClusterFeatureResponseList) SetResults(v []ClusterFeature) {
	o.Results = v
}

func (o ClusterFeatureResponseList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterFeatureResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableClusterFeatureResponseList struct {
	value *ClusterFeatureResponseList
	isSet bool
}

func (v NullableClusterFeatureResponseList) Get() *ClusterFeatureResponseList {
	return v.value
}

func (v *NullableClusterFeatureResponseList) Set(val *ClusterFeatureResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterFeatureResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterFeatureResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterFeatureResponseList(val *ClusterFeatureResponseList) *NullableClusterFeatureResponseList {
	return &NullableClusterFeatureResponseList{value: val, isSet: true}
}

func (v NullableClusterFeatureResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterFeatureResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
