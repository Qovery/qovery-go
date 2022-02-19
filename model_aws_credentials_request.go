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

// AwsCredentialsRequest struct for AwsCredentialsRequest
type AwsCredentialsRequest struct {
	Name            string  `json:"name"`
	AccessKeyId     *string `json:"access_key_id,omitempty"`
	SecretAccessKey *string `json:"secret_access_key,omitempty"`
}

// NewAwsCredentialsRequest instantiates a new AwsCredentialsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsCredentialsRequest(name string) *AwsCredentialsRequest {
	this := AwsCredentialsRequest{}
	this.Name = name
	return &this
}

// NewAwsCredentialsRequestWithDefaults instantiates a new AwsCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsCredentialsRequestWithDefaults() *AwsCredentialsRequest {
	this := AwsCredentialsRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AwsCredentialsRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AwsCredentialsRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AwsCredentialsRequest) SetName(v string) {
	o.Name = v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *AwsCredentialsRequest) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCredentialsRequest) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *AwsCredentialsRequest) HasAccessKeyId() bool {
	if o != nil && o.AccessKeyId != nil {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *AwsCredentialsRequest) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *AwsCredentialsRequest) GetSecretAccessKey() string {
	if o == nil || o.SecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCredentialsRequest) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.SecretAccessKey == nil {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *AwsCredentialsRequest) HasSecretAccessKey() bool {
	if o != nil && o.SecretAccessKey != nil {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *AwsCredentialsRequest) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

func (o AwsCredentialsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.AccessKeyId != nil {
		toSerialize["access_key_id"] = o.AccessKeyId
	}
	if o.SecretAccessKey != nil {
		toSerialize["secret_access_key"] = o.SecretAccessKey
	}
	return json.Marshal(toSerialize)
}

type NullableAwsCredentialsRequest struct {
	value *AwsCredentialsRequest
	isSet bool
}

func (v NullableAwsCredentialsRequest) Get() *AwsCredentialsRequest {
	return v.value
}

func (v *NullableAwsCredentialsRequest) Set(val *AwsCredentialsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsCredentialsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsCredentialsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsCredentialsRequest(val *AwsCredentialsRequest) *NullableAwsCredentialsRequest {
	return &NullableAwsCredentialsRequest{value: val, isSet: true}
}

func (v NullableAwsCredentialsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsCredentialsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
