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

// checks if the VariableImportRequestVarsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableImportRequestVarsInner{}

// VariableImportRequestVarsInner struct for VariableImportRequestVarsInner
type VariableImportRequestVarsInner struct {
	Name     string               `json:"name"`
	Value    string               `json:"value"`
	Scope    APIVariableScopeEnum `json:"scope"`
	IsSecret bool                 `json:"is_secret"`
}

// NewVariableImportRequestVarsInner instantiates a new VariableImportRequestVarsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableImportRequestVarsInner(name string, value string, scope APIVariableScopeEnum, isSecret bool) *VariableImportRequestVarsInner {
	this := VariableImportRequestVarsInner{}
	this.Name = name
	this.Value = value
	this.Scope = scope
	this.IsSecret = isSecret
	return &this
}

// NewVariableImportRequestVarsInnerWithDefaults instantiates a new VariableImportRequestVarsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableImportRequestVarsInnerWithDefaults() *VariableImportRequestVarsInner {
	this := VariableImportRequestVarsInner{}
	return &this
}

// GetName returns the Name field value
func (o *VariableImportRequestVarsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VariableImportRequestVarsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VariableImportRequestVarsInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *VariableImportRequestVarsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VariableImportRequestVarsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VariableImportRequestVarsInner) SetValue(v string) {
	o.Value = v
}

// GetScope returns the Scope field value
func (o *VariableImportRequestVarsInner) GetScope() APIVariableScopeEnum {
	if o == nil {
		var ret APIVariableScopeEnum
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *VariableImportRequestVarsInner) GetScopeOk() (*APIVariableScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *VariableImportRequestVarsInner) SetScope(v APIVariableScopeEnum) {
	o.Scope = v
}

// GetIsSecret returns the IsSecret field value
func (o *VariableImportRequestVarsInner) GetIsSecret() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSecret
}

// GetIsSecretOk returns a tuple with the IsSecret field value
// and a boolean to check if the value has been set.
func (o *VariableImportRequestVarsInner) GetIsSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSecret, true
}

// SetIsSecret sets field value
func (o *VariableImportRequestVarsInner) SetIsSecret(v bool) {
	o.IsSecret = v
}

func (o VariableImportRequestVarsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableImportRequestVarsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	toSerialize["scope"] = o.Scope
	toSerialize["is_secret"] = o.IsSecret
	return toSerialize, nil
}

type NullableVariableImportRequestVarsInner struct {
	value *VariableImportRequestVarsInner
	isSet bool
}

func (v NullableVariableImportRequestVarsInner) Get() *VariableImportRequestVarsInner {
	return v.value
}

func (v *NullableVariableImportRequestVarsInner) Set(val *VariableImportRequestVarsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableImportRequestVarsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableImportRequestVarsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableImportRequestVarsInner(val *VariableImportRequestVarsInner) *NullableVariableImportRequestVarsInner {
	return &NullableVariableImportRequestVarsInner{value: val, isSet: true}
}

func (v NullableVariableImportRequestVarsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableImportRequestVarsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
