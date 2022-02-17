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

// EnvironmentEditRequest struct for EnvironmentEditRequest
type EnvironmentEditRequest struct {
	Name *string `json:"name,omitempty"`
}

// NewEnvironmentEditRequest instantiates a new EnvironmentEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentEditRequest() *EnvironmentEditRequest {
	this := EnvironmentEditRequest{}
	return &this
}

// NewEnvironmentEditRequestWithDefaults instantiates a new EnvironmentEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentEditRequestWithDefaults() *EnvironmentEditRequest {
	this := EnvironmentEditRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentEditRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentEditRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentEditRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentEditRequest) SetName(v string) {
	o.Name = &v
}

func (o EnvironmentEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentEditRequest struct {
	value *EnvironmentEditRequest
	isSet bool
}

func (v NullableEnvironmentEditRequest) Get() *EnvironmentEditRequest {
	return v.value
}

func (v *NullableEnvironmentEditRequest) Set(val *EnvironmentEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentEditRequest(val *EnvironmentEditRequest) *NullableEnvironmentEditRequest {
	return &NullableEnvironmentEditRequest{value: val, isSet: true}
}

func (v NullableEnvironmentEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


