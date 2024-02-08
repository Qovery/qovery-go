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
	"time"
)

// checks if the EnvironmentVariable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentVariable{}

// EnvironmentVariable struct for EnvironmentVariable
type EnvironmentVariable struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// key is case sensitive.
	Key string `json:"key"`
	// value of the env variable.
	Value *string `json:"value,omitempty"`
	// should be set for file only. variable mount path makes variable a file (where file should be mounted).
	MountPath          NullableString               `json:"mount_path,omitempty"`
	OverriddenVariable *EnvironmentVariableOverride `json:"overridden_variable,omitempty"`
	AliasedVariable    *EnvironmentVariableAlias    `json:"aliased_variable,omitempty"`
	Scope              APIVariableScopeEnum         `json:"scope"`
	VariableType       *APIVariableTypeEnum         `json:"variable_type,omitempty"`
	ServiceId          *string                      `json:"service_id,omitempty"`
	ServiceName        *string                      `json:"service_name,omitempty"`
	ServiceType        *LinkedServiceTypeEnum       `json:"service_type,omitempty"`
	// Entity that created/own the variable (i.e: Qovery, Doppler)
	OwnedBy              *string `json:"owned_by,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnvironmentVariable EnvironmentVariable

// NewEnvironmentVariable instantiates a new EnvironmentVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariable(id string, createdAt time.Time, key string, scope APIVariableScopeEnum) *EnvironmentVariable {
	this := EnvironmentVariable{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Key = key
	this.Scope = scope
	return &this
}

// NewEnvironmentVariableWithDefaults instantiates a new EnvironmentVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableWithDefaults() *EnvironmentVariable {
	this := EnvironmentVariable{}
	return &this
}

// GetId returns the Id field value
func (o *EnvironmentVariable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentVariable) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EnvironmentVariable) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EnvironmentVariable) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentVariable) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetKey returns the Key field value
func (o *EnvironmentVariable) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EnvironmentVariable) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EnvironmentVariable) SetValue(v string) {
	o.Value = &v
}

// GetMountPath returns the MountPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentVariable) GetMountPath() string {
	if o == nil || IsNil(o.MountPath.Get()) {
		var ret string
		return ret
	}
	return *o.MountPath.Get()
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentVariable) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountPath.Get(), o.MountPath.IsSet()
}

// HasMountPath returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasMountPath() bool {
	if o != nil && o.MountPath.IsSet() {
		return true
	}

	return false
}

// SetMountPath gets a reference to the given NullableString and assigns it to the MountPath field.
func (o *EnvironmentVariable) SetMountPath(v string) {
	o.MountPath.Set(&v)
}

// SetMountPathNil sets the value for MountPath to be an explicit nil
func (o *EnvironmentVariable) SetMountPathNil() {
	o.MountPath.Set(nil)
}

// UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
func (o *EnvironmentVariable) UnsetMountPath() {
	o.MountPath.Unset()
}

// GetOverriddenVariable returns the OverriddenVariable field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetOverriddenVariable() EnvironmentVariableOverride {
	if o == nil || IsNil(o.OverriddenVariable) {
		var ret EnvironmentVariableOverride
		return ret
	}
	return *o.OverriddenVariable
}

// GetOverriddenVariableOk returns a tuple with the OverriddenVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetOverriddenVariableOk() (*EnvironmentVariableOverride, bool) {
	if o == nil || IsNil(o.OverriddenVariable) {
		return nil, false
	}
	return o.OverriddenVariable, true
}

// HasOverriddenVariable returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasOverriddenVariable() bool {
	if o != nil && !IsNil(o.OverriddenVariable) {
		return true
	}

	return false
}

// SetOverriddenVariable gets a reference to the given EnvironmentVariableOverride and assigns it to the OverriddenVariable field.
func (o *EnvironmentVariable) SetOverriddenVariable(v EnvironmentVariableOverride) {
	o.OverriddenVariable = &v
}

// GetAliasedVariable returns the AliasedVariable field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetAliasedVariable() EnvironmentVariableAlias {
	if o == nil || IsNil(o.AliasedVariable) {
		var ret EnvironmentVariableAlias
		return ret
	}
	return *o.AliasedVariable
}

// GetAliasedVariableOk returns a tuple with the AliasedVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetAliasedVariableOk() (*EnvironmentVariableAlias, bool) {
	if o == nil || IsNil(o.AliasedVariable) {
		return nil, false
	}
	return o.AliasedVariable, true
}

// HasAliasedVariable returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasAliasedVariable() bool {
	if o != nil && !IsNil(o.AliasedVariable) {
		return true
	}

	return false
}

// SetAliasedVariable gets a reference to the given EnvironmentVariableAlias and assigns it to the AliasedVariable field.
func (o *EnvironmentVariable) SetAliasedVariable(v EnvironmentVariableAlias) {
	o.AliasedVariable = &v
}

// GetScope returns the Scope field value
func (o *EnvironmentVariable) GetScope() APIVariableScopeEnum {
	if o == nil {
		var ret APIVariableScopeEnum
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetScopeOk() (*APIVariableScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *EnvironmentVariable) SetScope(v APIVariableScopeEnum) {
	o.Scope = v
}

// GetVariableType returns the VariableType field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetVariableType() APIVariableTypeEnum {
	if o == nil || IsNil(o.VariableType) {
		var ret APIVariableTypeEnum
		return ret
	}
	return *o.VariableType
}

// GetVariableTypeOk returns a tuple with the VariableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetVariableTypeOk() (*APIVariableTypeEnum, bool) {
	if o == nil || IsNil(o.VariableType) {
		return nil, false
	}
	return o.VariableType, true
}

// HasVariableType returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasVariableType() bool {
	if o != nil && !IsNil(o.VariableType) {
		return true
	}

	return false
}

// SetVariableType gets a reference to the given APIVariableTypeEnum and assigns it to the VariableType field.
func (o *EnvironmentVariable) SetVariableType(v APIVariableTypeEnum) {
	o.VariableType = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *EnvironmentVariable) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *EnvironmentVariable) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetServiceType() LinkedServiceTypeEnum {
	if o == nil || IsNil(o.ServiceType) {
		var ret LinkedServiceTypeEnum
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetServiceTypeOk() (*LinkedServiceTypeEnum, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given LinkedServiceTypeEnum and assigns it to the ServiceType field.
func (o *EnvironmentVariable) SetServiceType(v LinkedServiceTypeEnum) {
	o.ServiceType = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetOwnedBy() string {
	if o == nil || IsNil(o.OwnedBy) {
		var ret string
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetOwnedByOk() (*string, bool) {
	if o == nil || IsNil(o.OwnedBy) {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasOwnedBy() bool {
	if o != nil && !IsNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given string and assigns it to the OwnedBy field.
func (o *EnvironmentVariable) SetOwnedBy(v string) {
	o.OwnedBy = &v
}

func (o EnvironmentVariable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentVariable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["key"] = o.Key
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if o.MountPath.IsSet() {
		toSerialize["mount_path"] = o.MountPath.Get()
	}
	if !IsNil(o.OverriddenVariable) {
		toSerialize["overridden_variable"] = o.OverriddenVariable
	}
	if !IsNil(o.AliasedVariable) {
		toSerialize["aliased_variable"] = o.AliasedVariable
	}
	toSerialize["scope"] = o.Scope
	if !IsNil(o.VariableType) {
		toSerialize["variable_type"] = o.VariableType
	}
	if !IsNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !IsNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}
	if !IsNil(o.ServiceType) {
		toSerialize["service_type"] = o.ServiceType
	}
	if !IsNil(o.OwnedBy) {
		toSerialize["owned_by"] = o.OwnedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"key",
		"scope",
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

	varEnvironmentVariable := _EnvironmentVariable{}

	err = json.Unmarshal(data, &varEnvironmentVariable)

	if err != nil {
		return err
	}

	*o = EnvironmentVariable(varEnvironmentVariable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "key")
		delete(additionalProperties, "value")
		delete(additionalProperties, "mount_path")
		delete(additionalProperties, "overridden_variable")
		delete(additionalProperties, "aliased_variable")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "variable_type")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "service_name")
		delete(additionalProperties, "service_type")
		delete(additionalProperties, "owned_by")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnvironmentVariable struct {
	value *EnvironmentVariable
	isSet bool
}

func (v NullableEnvironmentVariable) Get() *EnvironmentVariable {
	return v.value
}

func (v *NullableEnvironmentVariable) Set(val *EnvironmentVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariable(val *EnvironmentVariable) *NullableEnvironmentVariable {
	return &NullableEnvironmentVariable{value: val, isSet: true}
}

func (v NullableEnvironmentVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
