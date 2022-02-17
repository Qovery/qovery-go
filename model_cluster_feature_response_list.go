/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet. 

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// ClusterFeatureResponseList struct for ClusterFeatureResponseList
type ClusterFeatureResponseList struct {
	Results []ClusterFeatureResponse `json:"results,omitempty"`
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
func (o *ClusterFeatureResponseList) GetResults() []ClusterFeatureResponse {
	if o == nil || o.Results == nil {
		var ret []ClusterFeatureResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeatureResponseList) GetResultsOk() ([]ClusterFeatureResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ClusterFeatureResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ClusterFeatureResponse and assigns it to the Results field.
func (o *ClusterFeatureResponseList) SetResults(v []ClusterFeatureResponse) {
	o.Results = v
}

func (o ClusterFeatureResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
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


