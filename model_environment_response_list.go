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

// EnvironmentResponseList struct for EnvironmentResponseList
type EnvironmentResponseList struct {
	Results []EnvironmentResponse `json:"results,omitempty"`
}

// NewEnvironmentResponseList instantiates a new EnvironmentResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentResponseList() *EnvironmentResponseList {
	this := EnvironmentResponseList{}
	return &this
}

// NewEnvironmentResponseListWithDefaults instantiates a new EnvironmentResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentResponseListWithDefaults() *EnvironmentResponseList {
	this := EnvironmentResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *EnvironmentResponseList) GetResults() []EnvironmentResponse {
	if o == nil || o.Results == nil {
		var ret []EnvironmentResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentResponseList) GetResultsOk() ([]EnvironmentResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EnvironmentResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EnvironmentResponse and assigns it to the Results field.
func (o *EnvironmentResponseList) SetResults(v []EnvironmentResponse) {
	o.Results = v
}

func (o EnvironmentResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentResponseList struct {
	value *EnvironmentResponseList
	isSet bool
}

func (v NullableEnvironmentResponseList) Get() *EnvironmentResponseList {
	return v.value
}

func (v *NullableEnvironmentResponseList) Set(val *EnvironmentResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentResponseList(val *EnvironmentResponseList) *NullableEnvironmentResponseList {
	return &NullableEnvironmentResponseList{value: val, isSet: true}
}

func (v NullableEnvironmentResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
