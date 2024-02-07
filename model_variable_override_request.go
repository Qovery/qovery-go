/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the VariableOverrideRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableOverrideRequest{}

// VariableOverrideRequest struct for VariableOverrideRequest
type VariableOverrideRequest struct {
	// the value to be used as Override of the targeted environment variable.
	Value         string               `json:"value"`
	OverrideScope APIVariableScopeEnum `json:"override_scope"`
	// the id of the variable that is aliased.
	OverrideParentId string `json:"override_parent_id"`
}

type _VariableOverrideRequest VariableOverrideRequest

// NewVariableOverrideRequest instantiates a new VariableOverrideRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableOverrideRequest(value string, overrideScope APIVariableScopeEnum, overrideParentId string) *VariableOverrideRequest {
	this := VariableOverrideRequest{}
	this.Value = value
	this.OverrideScope = overrideScope
	this.OverrideParentId = overrideParentId
	return &this
}

// NewVariableOverrideRequestWithDefaults instantiates a new VariableOverrideRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableOverrideRequestWithDefaults() *VariableOverrideRequest {
	this := VariableOverrideRequest{}
	return &this
}

// GetValue returns the Value field value
func (o *VariableOverrideRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VariableOverrideRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VariableOverrideRequest) SetValue(v string) {
	o.Value = v
}

// GetOverrideScope returns the OverrideScope field value
func (o *VariableOverrideRequest) GetOverrideScope() APIVariableScopeEnum {
	if o == nil {
		var ret APIVariableScopeEnum
		return ret
	}

	return o.OverrideScope
}

// GetOverrideScopeOk returns a tuple with the OverrideScope field value
// and a boolean to check if the value has been set.
func (o *VariableOverrideRequest) GetOverrideScopeOk() (*APIVariableScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverrideScope, true
}

// SetOverrideScope sets field value
func (o *VariableOverrideRequest) SetOverrideScope(v APIVariableScopeEnum) {
	o.OverrideScope = v
}

// GetOverrideParentId returns the OverrideParentId field value
func (o *VariableOverrideRequest) GetOverrideParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OverrideParentId
}

// GetOverrideParentIdOk returns a tuple with the OverrideParentId field value
// and a boolean to check if the value has been set.
func (o *VariableOverrideRequest) GetOverrideParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverrideParentId, true
}

// SetOverrideParentId sets field value
func (o *VariableOverrideRequest) SetOverrideParentId(v string) {
	o.OverrideParentId = v
}

func (o VariableOverrideRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableOverrideRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["override_scope"] = o.OverrideScope
	toSerialize["override_parent_id"] = o.OverrideParentId
	return toSerialize, nil
}

func (o *VariableOverrideRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"override_scope",
		"override_parent_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVariableOverrideRequest := _VariableOverrideRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVariableOverrideRequest)

	if err != nil {
		return err
	}

	*o = VariableOverrideRequest(varVariableOverrideRequest)

	return err
}

type NullableVariableOverrideRequest struct {
	value *VariableOverrideRequest
	isSet bool
}

func (v NullableVariableOverrideRequest) Get() *VariableOverrideRequest {
	return v.value
}

func (v *NullableVariableOverrideRequest) Set(val *VariableOverrideRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableOverrideRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableOverrideRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableOverrideRequest(val *VariableOverrideRequest) *NullableVariableOverrideRequest {
	return &NullableVariableOverrideRequest{value: val, isSet: true}
}

func (v NullableVariableOverrideRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableOverrideRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
