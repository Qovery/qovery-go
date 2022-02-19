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

// EnvironmentApplicationsStorageResponseList struct for EnvironmentApplicationsStorageResponseList
type EnvironmentApplicationsStorageResponseList struct {
	Results []EnvironmentApplicationsStorageResponse `json:"results,omitempty"`
}

// NewEnvironmentApplicationsStorageResponseList instantiates a new EnvironmentApplicationsStorageResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentApplicationsStorageResponseList() *EnvironmentApplicationsStorageResponseList {
	this := EnvironmentApplicationsStorageResponseList{}
	return &this
}

// NewEnvironmentApplicationsStorageResponseListWithDefaults instantiates a new EnvironmentApplicationsStorageResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentApplicationsStorageResponseListWithDefaults() *EnvironmentApplicationsStorageResponseList {
	this := EnvironmentApplicationsStorageResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *EnvironmentApplicationsStorageResponseList) GetResults() []EnvironmentApplicationsStorageResponse {
	if o == nil || o.Results == nil {
		var ret []EnvironmentApplicationsStorageResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsStorageResponseList) GetResultsOk() ([]EnvironmentApplicationsStorageResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EnvironmentApplicationsStorageResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EnvironmentApplicationsStorageResponse and assigns it to the Results field.
func (o *EnvironmentApplicationsStorageResponseList) SetResults(v []EnvironmentApplicationsStorageResponse) {
	o.Results = v
}

func (o EnvironmentApplicationsStorageResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentApplicationsStorageResponseList struct {
	value *EnvironmentApplicationsStorageResponseList
	isSet bool
}

func (v NullableEnvironmentApplicationsStorageResponseList) Get() *EnvironmentApplicationsStorageResponseList {
	return v.value
}

func (v *NullableEnvironmentApplicationsStorageResponseList) Set(val *EnvironmentApplicationsStorageResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentApplicationsStorageResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentApplicationsStorageResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentApplicationsStorageResponseList(val *EnvironmentApplicationsStorageResponseList) *NullableEnvironmentApplicationsStorageResponseList {
	return &NullableEnvironmentApplicationsStorageResponseList{value: val, isSet: true}
}

func (v NullableEnvironmentApplicationsStorageResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentApplicationsStorageResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
