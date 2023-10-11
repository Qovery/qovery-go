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

// checks if the EnvironmentVariableRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentVariableRequest{}

// EnvironmentVariableRequest struct for EnvironmentVariableRequest
type EnvironmentVariableRequest struct {
	// key is case sensitive.
	Key string `json:"key"`
	// value of the env variable.
	Value *string `json:"value,omitempty"`
	// should be set for file only. variable mount path makes variable a file (where file should be mounted).
	MountPath NullableString `json:"mount_path,omitempty"`
}

// NewEnvironmentVariableRequest instantiates a new EnvironmentVariableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariableRequest(key string) *EnvironmentVariableRequest {
	this := EnvironmentVariableRequest{}
	this.Key = key
	return &this
}

// NewEnvironmentVariableRequestWithDefaults instantiates a new EnvironmentVariableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableRequestWithDefaults() *EnvironmentVariableRequest {
	this := EnvironmentVariableRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *EnvironmentVariableRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EnvironmentVariableRequest) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EnvironmentVariableRequest) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableRequest) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EnvironmentVariableRequest) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EnvironmentVariableRequest) SetValue(v string) {
	o.Value = &v
}

// GetMountPath returns the MountPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentVariableRequest) GetMountPath() string {
	if o == nil || IsNil(o.MountPath.Get()) {
		var ret string
		return ret
	}
	return *o.MountPath.Get()
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentVariableRequest) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountPath.Get(), o.MountPath.IsSet()
}

// HasMountPath returns a boolean if a field has been set.
func (o *EnvironmentVariableRequest) HasMountPath() bool {
	if o != nil && o.MountPath.IsSet() {
		return true
	}

	return false
}

// SetMountPath gets a reference to the given NullableString and assigns it to the MountPath field.
func (o *EnvironmentVariableRequest) SetMountPath(v string) {
	o.MountPath.Set(&v)
}

// SetMountPathNil sets the value for MountPath to be an explicit nil
func (o *EnvironmentVariableRequest) SetMountPathNil() {
	o.MountPath.Set(nil)
}

// UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
func (o *EnvironmentVariableRequest) UnsetMountPath() {
	o.MountPath.Unset()
}

func (o EnvironmentVariableRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentVariableRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if o.MountPath.IsSet() {
		toSerialize["mount_path"] = o.MountPath.Get()
	}
	return toSerialize, nil
}

type NullableEnvironmentVariableRequest struct {
	value *EnvironmentVariableRequest
	isSet bool
}

func (v NullableEnvironmentVariableRequest) Get() *EnvironmentVariableRequest {
	return v.value
}

func (v *NullableEnvironmentVariableRequest) Set(val *EnvironmentVariableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariableRequest(val *EnvironmentVariableRequest) *NullableEnvironmentVariableRequest {
	return &NullableEnvironmentVariableRequest{value: val, isSet: true}
}

func (v NullableEnvironmentVariableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
