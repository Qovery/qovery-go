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

// checks if the ServiceStorageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceStorageRequest{}

// ServiceStorageRequest struct for ServiceStorageRequest
type ServiceStorageRequest struct {
	Storage []ServiceStorageRequestStorageInner `json:"storage,omitempty"`
}

// NewServiceStorageRequest instantiates a new ServiceStorageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStorageRequest() *ServiceStorageRequest {
	this := ServiceStorageRequest{}
	return &this
}

// NewServiceStorageRequestWithDefaults instantiates a new ServiceStorageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStorageRequestWithDefaults() *ServiceStorageRequest {
	this := ServiceStorageRequest{}
	return &this
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ServiceStorageRequest) GetStorage() []ServiceStorageRequestStorageInner {
	if o == nil || IsNil(o.Storage) {
		var ret []ServiceStorageRequestStorageInner
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStorageRequest) GetStorageOk() ([]ServiceStorageRequestStorageInner, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ServiceStorageRequest) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []ServiceStorageRequestStorageInner and assigns it to the Storage field.
func (o *ServiceStorageRequest) SetStorage(v []ServiceStorageRequestStorageInner) {
	o.Storage = v
}

func (o ServiceStorageRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceStorageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	return toSerialize, nil
}

type NullableServiceStorageRequest struct {
	value *ServiceStorageRequest
	isSet bool
}

func (v NullableServiceStorageRequest) Get() *ServiceStorageRequest {
	return v.value
}

func (v *NullableServiceStorageRequest) Set(val *ServiceStorageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStorageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStorageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStorageRequest(val *ServiceStorageRequest) *NullableServiceStorageRequest {
	return &NullableServiceStorageRequest{value: val, isSet: true}
}

func (v NullableServiceStorageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStorageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
