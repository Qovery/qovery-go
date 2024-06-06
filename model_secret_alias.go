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
	"fmt"
)

// checks if the SecretAlias type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretAlias{}

// SecretAlias struct for SecretAlias
type SecretAlias struct {
	Id                   string               `json:"id"`
	Key                  string               `json:"key"`
	MountPath            string               `json:"mount_path"`
	Scope                APIVariableScopeEnum `json:"scope"`
	VariableType         APIVariableTypeEnum  `json:"variable_type"`
	AdditionalProperties map[string]interface{}
}

type _SecretAlias SecretAlias

// NewSecretAlias instantiates a new SecretAlias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretAlias(id string, key string, mountPath string, scope APIVariableScopeEnum, variableType APIVariableTypeEnum) *SecretAlias {
	this := SecretAlias{}
	this.Id = id
	this.Key = key
	this.MountPath = mountPath
	this.Scope = scope
	this.VariableType = variableType
	return &this
}

// NewSecretAliasWithDefaults instantiates a new SecretAlias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretAliasWithDefaults() *SecretAlias {
	this := SecretAlias{}
	return &this
}

// GetId returns the Id field value
func (o *SecretAlias) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecretAlias) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecretAlias) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *SecretAlias) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SecretAlias) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SecretAlias) SetKey(v string) {
	o.Key = v
}

// GetMountPath returns the MountPath field value
func (o *SecretAlias) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *SecretAlias) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountPath, true
}

// SetMountPath sets field value
func (o *SecretAlias) SetMountPath(v string) {
	o.MountPath = v
}

// GetScope returns the Scope field value
func (o *SecretAlias) GetScope() APIVariableScopeEnum {
	if o == nil {
		var ret APIVariableScopeEnum
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *SecretAlias) GetScopeOk() (*APIVariableScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *SecretAlias) SetScope(v APIVariableScopeEnum) {
	o.Scope = v
}

// GetVariableType returns the VariableType field value
func (o *SecretAlias) GetVariableType() APIVariableTypeEnum {
	if o == nil {
		var ret APIVariableTypeEnum
		return ret
	}

	return o.VariableType
}

// GetVariableTypeOk returns a tuple with the VariableType field value
// and a boolean to check if the value has been set.
func (o *SecretAlias) GetVariableTypeOk() (*APIVariableTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariableType, true
}

// SetVariableType sets field value
func (o *SecretAlias) SetVariableType(v APIVariableTypeEnum) {
	o.VariableType = v
}

func (o SecretAlias) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretAlias) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["mount_path"] = o.MountPath
	toSerialize["scope"] = o.Scope
	toSerialize["variable_type"] = o.VariableType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecretAlias) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"key",
		"mount_path",
		"scope",
		"variable_type",
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

	varSecretAlias := _SecretAlias{}

	err = json.Unmarshal(data, &varSecretAlias)

	if err != nil {
		return err
	}

	*o = SecretAlias(varSecretAlias)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "mount_path")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "variable_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecretAlias struct {
	value *SecretAlias
	isSet bool
}

func (v NullableSecretAlias) Get() *SecretAlias {
	return v.value
}

func (v *NullableSecretAlias) Set(val *SecretAlias) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretAlias(val *SecretAlias) *NullableSecretAlias {
	return &NullableSecretAlias{value: val, isSet: true}
}

func (v NullableSecretAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
