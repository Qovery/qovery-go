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

// ClusterCredentialsRequest struct for ClusterCredentialsRequest
type ClusterCredentialsRequest struct {
	Name            string  `json:"name"`
	AccessKeyId     *string `json:"access_key_id,omitempty"`
	SecretAccessKey *string `json:"secret_access_key,omitempty"`
}

// NewClusterCredentialsRequest instantiates a new ClusterCredentialsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterCredentialsRequest(name string) *ClusterCredentialsRequest {
	this := ClusterCredentialsRequest{}
	this.Name = name
	return &this
}

// NewClusterCredentialsRequestWithDefaults instantiates a new ClusterCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterCredentialsRequestWithDefaults() *ClusterCredentialsRequest {
	this := ClusterCredentialsRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ClusterCredentialsRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterCredentialsRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterCredentialsRequest) SetName(v string) {
	o.Name = v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *ClusterCredentialsRequest) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCredentialsRequest) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *ClusterCredentialsRequest) HasAccessKeyId() bool {
	if o != nil && o.AccessKeyId != nil {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *ClusterCredentialsRequest) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *ClusterCredentialsRequest) GetSecretAccessKey() string {
	if o == nil || o.SecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCredentialsRequest) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.SecretAccessKey == nil {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *ClusterCredentialsRequest) HasSecretAccessKey() bool {
	if o != nil && o.SecretAccessKey != nil {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *ClusterCredentialsRequest) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

func (o ClusterCredentialsRequest) MarshalJSON() ([]byte, error) {
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

type NullableClusterCredentialsRequest struct {
	value *ClusterCredentialsRequest
	isSet bool
}

func (v NullableClusterCredentialsRequest) Get() *ClusterCredentialsRequest {
	return v.value
}

func (v *NullableClusterCredentialsRequest) Set(val *ClusterCredentialsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterCredentialsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterCredentialsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterCredentialsRequest(val *ClusterCredentialsRequest) *NullableClusterCredentialsRequest {
	return &NullableClusterCredentialsRequest{value: val, isSet: true}
}

func (v NullableClusterCredentialsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterCredentialsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
