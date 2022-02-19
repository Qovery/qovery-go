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

// SecretRequest struct for SecretRequest
type SecretRequest struct {
	// key is case sensitive
	Key string `json:"key"`
	// value of the secret. Clear value will never be returned
	Value string `json:"value"`
}

// NewSecretRequest instantiates a new SecretRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretRequest(key string, value string) *SecretRequest {
	this := SecretRequest{}
	this.Key = key
	this.Value = value
	return &this
}

// NewSecretRequestWithDefaults instantiates a new SecretRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretRequestWithDefaults() *SecretRequest {
	this := SecretRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *SecretRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SecretRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SecretRequest) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *SecretRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SecretRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SecretRequest) SetValue(v string) {
	o.Value = v
}

func (o SecretRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSecretRequest struct {
	value *SecretRequest
	isSet bool
}

func (v NullableSecretRequest) Get() *SecretRequest {
	return v.value
}

func (v *NullableSecretRequest) Set(val *SecretRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretRequest(val *SecretRequest) *NullableSecretRequest {
	return &NullableSecretRequest{value: val, isSet: true}
}

func (v NullableSecretRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
