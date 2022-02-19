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

// SecretResponseList struct for SecretResponseList
type SecretResponseList struct {
	Results []SecretResponse `json:"results,omitempty"`
}

// NewSecretResponseList instantiates a new SecretResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretResponseList() *SecretResponseList {
	this := SecretResponseList{}
	return &this
}

// NewSecretResponseListWithDefaults instantiates a new SecretResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretResponseListWithDefaults() *SecretResponseList {
	this := SecretResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SecretResponseList) GetResults() []SecretResponse {
	if o == nil || o.Results == nil {
		var ret []SecretResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretResponseList) GetResultsOk() ([]SecretResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SecretResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []SecretResponse and assigns it to the Results field.
func (o *SecretResponseList) SetResults(v []SecretResponse) {
	o.Results = v
}

func (o SecretResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableSecretResponseList struct {
	value *SecretResponseList
	isSet bool
}

func (v NullableSecretResponseList) Get() *SecretResponseList {
	return v.value
}

func (v *NullableSecretResponseList) Set(val *SecretResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretResponseList(val *SecretResponseList) *NullableSecretResponseList {
	return &NullableSecretResponseList{value: val, isSet: true}
}

func (v NullableSecretResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
