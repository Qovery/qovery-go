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

// ClusterStatusResponseList struct for ClusterStatusResponseList
type ClusterStatusResponseList struct {
	Results []ClusterStatusResponse `json:"results,omitempty"`
}

// NewClusterStatusResponseList instantiates a new ClusterStatusResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterStatusResponseList() *ClusterStatusResponseList {
	this := ClusterStatusResponseList{}
	return &this
}

// NewClusterStatusResponseListWithDefaults instantiates a new ClusterStatusResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterStatusResponseListWithDefaults() *ClusterStatusResponseList {
	this := ClusterStatusResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ClusterStatusResponseList) GetResults() []ClusterStatusResponse {
	if o == nil || o.Results == nil {
		var ret []ClusterStatusResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatusResponseList) GetResultsOk() ([]ClusterStatusResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ClusterStatusResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ClusterStatusResponse and assigns it to the Results field.
func (o *ClusterStatusResponseList) SetResults(v []ClusterStatusResponse) {
	o.Results = v
}

func (o ClusterStatusResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableClusterStatusResponseList struct {
	value *ClusterStatusResponseList
	isSet bool
}

func (v NullableClusterStatusResponseList) Get() *ClusterStatusResponseList {
	return v.value
}

func (v *NullableClusterStatusResponseList) Set(val *ClusterStatusResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStatusResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStatusResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStatusResponseList(val *ClusterStatusResponseList) *NullableClusterStatusResponseList {
	return &NullableClusterStatusResponseList{value: val, isSet: true}
}

func (v NullableClusterStatusResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStatusResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
