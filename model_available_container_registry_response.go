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

// AvailableContainerRegistryResponse struct for AvailableContainerRegistryResponse
type AvailableContainerRegistryResponse struct {
	Kind           *ContainerRegistryKind `json:"kind,omitempty"`
	RequiredConfig map[string]interface{} `json:"required_config,omitempty"`
}

// NewAvailableContainerRegistryResponse instantiates a new AvailableContainerRegistryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableContainerRegistryResponse() *AvailableContainerRegistryResponse {
	this := AvailableContainerRegistryResponse{}
	var kind ContainerRegistryKind = CONTAINERREGISTRYKIND_ECR
	this.Kind = &kind
	return &this
}

// NewAvailableContainerRegistryResponseWithDefaults instantiates a new AvailableContainerRegistryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableContainerRegistryResponseWithDefaults() *AvailableContainerRegistryResponse {
	this := AvailableContainerRegistryResponse{}
	var kind ContainerRegistryKind = CONTAINERREGISTRYKIND_ECR
	this.Kind = &kind
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *AvailableContainerRegistryResponse) GetKind() ContainerRegistryKind {
	if o == nil || o.Kind == nil {
		var ret ContainerRegistryKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableContainerRegistryResponse) GetKindOk() (*ContainerRegistryKind, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *AvailableContainerRegistryResponse) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given ContainerRegistryKind and assigns it to the Kind field.
func (o *AvailableContainerRegistryResponse) SetKind(v ContainerRegistryKind) {
	o.Kind = &v
}

// GetRequiredConfig returns the RequiredConfig field value if set, zero value otherwise.
func (o *AvailableContainerRegistryResponse) GetRequiredConfig() map[string]interface{} {
	if o == nil || o.RequiredConfig == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.RequiredConfig
}

// GetRequiredConfigOk returns a tuple with the RequiredConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableContainerRegistryResponse) GetRequiredConfigOk() (map[string]interface{}, bool) {
	if o == nil || o.RequiredConfig == nil {
		return nil, false
	}
	return o.RequiredConfig, true
}

// HasRequiredConfig returns a boolean if a field has been set.
func (o *AvailableContainerRegistryResponse) HasRequiredConfig() bool {
	if o != nil && o.RequiredConfig != nil {
		return true
	}

	return false
}

// SetRequiredConfig gets a reference to the given map[string]interface{} and assigns it to the RequiredConfig field.
func (o *AvailableContainerRegistryResponse) SetRequiredConfig(v map[string]interface{}) {
	o.RequiredConfig = v
}

func (o AvailableContainerRegistryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.RequiredConfig != nil {
		toSerialize["required_config"] = o.RequiredConfig
	}
	return json.Marshal(toSerialize)
}

type NullableAvailableContainerRegistryResponse struct {
	value *AvailableContainerRegistryResponse
	isSet bool
}

func (v NullableAvailableContainerRegistryResponse) Get() *AvailableContainerRegistryResponse {
	return v.value
}

func (v *NullableAvailableContainerRegistryResponse) Set(val *AvailableContainerRegistryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableContainerRegistryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableContainerRegistryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableContainerRegistryResponse(val *AvailableContainerRegistryResponse) *NullableAvailableContainerRegistryResponse {
	return &NullableAvailableContainerRegistryResponse{value: val, isSet: true}
}

func (v NullableAvailableContainerRegistryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableContainerRegistryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
