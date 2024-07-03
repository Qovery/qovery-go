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

// checks if the VariableAliasRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableAliasRequest{}

// VariableAliasRequest struct for VariableAliasRequest
type VariableAliasRequest struct {
	// the value to be used as Alias of the targeted environment variable.
	Key        string               `json:"key"`
	AliasScope APIVariableScopeEnum `json:"alias_scope"`
	// the id of the variable that is aliased.
	AliasParentId string `json:"alias_parent_id"`
	// optional variable description (255 characters maximum)
	Description          NullableString `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VariableAliasRequest VariableAliasRequest

// NewVariableAliasRequest instantiates a new VariableAliasRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableAliasRequest(key string, aliasScope APIVariableScopeEnum, aliasParentId string) *VariableAliasRequest {
	this := VariableAliasRequest{}
	this.Key = key
	this.AliasScope = aliasScope
	this.AliasParentId = aliasParentId
	return &this
}

// NewVariableAliasRequestWithDefaults instantiates a new VariableAliasRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableAliasRequestWithDefaults() *VariableAliasRequest {
	this := VariableAliasRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *VariableAliasRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *VariableAliasRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *VariableAliasRequest) SetKey(v string) {
	o.Key = v
}

// GetAliasScope returns the AliasScope field value
func (o *VariableAliasRequest) GetAliasScope() APIVariableScopeEnum {
	if o == nil {
		var ret APIVariableScopeEnum
		return ret
	}

	return o.AliasScope
}

// GetAliasScopeOk returns a tuple with the AliasScope field value
// and a boolean to check if the value has been set.
func (o *VariableAliasRequest) GetAliasScopeOk() (*APIVariableScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AliasScope, true
}

// SetAliasScope sets field value
func (o *VariableAliasRequest) SetAliasScope(v APIVariableScopeEnum) {
	o.AliasScope = v
}

// GetAliasParentId returns the AliasParentId field value
func (o *VariableAliasRequest) GetAliasParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AliasParentId
}

// GetAliasParentIdOk returns a tuple with the AliasParentId field value
// and a boolean to check if the value has been set.
func (o *VariableAliasRequest) GetAliasParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AliasParentId, true
}

// SetAliasParentId sets field value
func (o *VariableAliasRequest) SetAliasParentId(v string) {
	o.AliasParentId = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VariableAliasRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariableAliasRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *VariableAliasRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *VariableAliasRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *VariableAliasRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *VariableAliasRequest) UnsetDescription() {
	o.Description.Unset()
}

func (o VariableAliasRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableAliasRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["alias_scope"] = o.AliasScope
	toSerialize["alias_parent_id"] = o.AliasParentId
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VariableAliasRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"alias_scope",
		"alias_parent_id",
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

	varVariableAliasRequest := _VariableAliasRequest{}

	err = json.Unmarshal(data, &varVariableAliasRequest)

	if err != nil {
		return err
	}

	*o = VariableAliasRequest(varVariableAliasRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "alias_scope")
		delete(additionalProperties, "alias_parent_id")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVariableAliasRequest struct {
	value *VariableAliasRequest
	isSet bool
}

func (v NullableVariableAliasRequest) Get() *VariableAliasRequest {
	return v.value
}

func (v *NullableVariableAliasRequest) Set(val *VariableAliasRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableAliasRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableAliasRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableAliasRequest(val *VariableAliasRequest) *NullableVariableAliasRequest {
	return &NullableVariableAliasRequest{value: val, isSet: true}
}

func (v NullableVariableAliasRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableAliasRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
