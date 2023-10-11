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

// checks if the EnvironmentContainersStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentContainersStorage{}

// EnvironmentContainersStorage struct for EnvironmentContainersStorage
type EnvironmentContainersStorage struct {
	Container string        `json:"container"`
	Disks     []StorageDisk `json:"disks,omitempty"`
}

// NewEnvironmentContainersStorage instantiates a new EnvironmentContainersStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentContainersStorage(container string) *EnvironmentContainersStorage {
	this := EnvironmentContainersStorage{}
	this.Container = container
	return &this
}

// NewEnvironmentContainersStorageWithDefaults instantiates a new EnvironmentContainersStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentContainersStorageWithDefaults() *EnvironmentContainersStorage {
	this := EnvironmentContainersStorage{}
	return &this
}

// GetContainer returns the Container field value
func (o *EnvironmentContainersStorage) GetContainer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *EnvironmentContainersStorage) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *EnvironmentContainersStorage) SetContainer(v string) {
	o.Container = v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *EnvironmentContainersStorage) GetDisks() []StorageDisk {
	if o == nil || IsNil(o.Disks) {
		var ret []StorageDisk
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentContainersStorage) GetDisksOk() ([]StorageDisk, bool) {
	if o == nil || IsNil(o.Disks) {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *EnvironmentContainersStorage) HasDisks() bool {
	if o != nil && !IsNil(o.Disks) {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []StorageDisk and assigns it to the Disks field.
func (o *EnvironmentContainersStorage) SetDisks(v []StorageDisk) {
	o.Disks = v
}

func (o EnvironmentContainersStorage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentContainersStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["container"] = o.Container
	if !IsNil(o.Disks) {
		toSerialize["disks"] = o.Disks
	}
	return toSerialize, nil
}

type NullableEnvironmentContainersStorage struct {
	value *EnvironmentContainersStorage
	isSet bool
}

func (v NullableEnvironmentContainersStorage) Get() *EnvironmentContainersStorage {
	return v.value
}

func (v *NullableEnvironmentContainersStorage) Set(val *EnvironmentContainersStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentContainersStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentContainersStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentContainersStorage(val *EnvironmentContainersStorage) *NullableEnvironmentContainersStorage {
	return &NullableEnvironmentContainersStorage{value: val, isSet: true}
}

func (v NullableEnvironmentContainersStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentContainersStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
