/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// EnvironmentApplicationsInstanceResponseList struct for EnvironmentApplicationsInstanceResponseList
type EnvironmentApplicationsInstanceResponseList struct {
	Results []EnvironmentApplicationsInstanceResponseListResults `json:"results,omitempty"`
}

// NewEnvironmentApplicationsInstanceResponseList instantiates a new EnvironmentApplicationsInstanceResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentApplicationsInstanceResponseList() *EnvironmentApplicationsInstanceResponseList {
	this := EnvironmentApplicationsInstanceResponseList{}
	return &this
}

// NewEnvironmentApplicationsInstanceResponseListWithDefaults instantiates a new EnvironmentApplicationsInstanceResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentApplicationsInstanceResponseListWithDefaults() *EnvironmentApplicationsInstanceResponseList {
	this := EnvironmentApplicationsInstanceResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *EnvironmentApplicationsInstanceResponseList) GetResults() []EnvironmentApplicationsInstanceResponseListResults {
	if o == nil || o.Results == nil {
		var ret []EnvironmentApplicationsInstanceResponseListResults
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsInstanceResponseList) GetResultsOk() ([]EnvironmentApplicationsInstanceResponseListResults, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EnvironmentApplicationsInstanceResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EnvironmentApplicationsInstanceResponseListResults and assigns it to the Results field.
func (o *EnvironmentApplicationsInstanceResponseList) SetResults(v []EnvironmentApplicationsInstanceResponseListResults) {
	o.Results = v
}

func (o EnvironmentApplicationsInstanceResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentApplicationsInstanceResponseList struct {
	value *EnvironmentApplicationsInstanceResponseList
	isSet bool
}

func (v NullableEnvironmentApplicationsInstanceResponseList) Get() *EnvironmentApplicationsInstanceResponseList {
	return v.value
}

func (v *NullableEnvironmentApplicationsInstanceResponseList) Set(val *EnvironmentApplicationsInstanceResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentApplicationsInstanceResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentApplicationsInstanceResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentApplicationsInstanceResponseList(val *EnvironmentApplicationsInstanceResponseList) *NullableEnvironmentApplicationsInstanceResponseList {
	return &NullableEnvironmentApplicationsInstanceResponseList{value: val, isSet: true}
}

func (v NullableEnvironmentApplicationsInstanceResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentApplicationsInstanceResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
