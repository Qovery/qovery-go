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

// checks if the VariableResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableResponse{}

// VariableResponse struct for VariableResponse
type VariableResponse struct {
	Id                 string               `json:"id"`
	CreatedAt          time.Time            `json:"created_at"`
	UpdatedAt          *time.Time           `json:"updated_at,omitempty"`
	Key                string               `json:"key"`
	Value              NullableString       `json:"value"`
	MountPath          NullableString       `json:"mount_path,omitempty"`
	OverriddenVariable *VariableOverride    `json:"overridden_variable,omitempty"`
	AliasedVariable    *VariableAlias       `json:"aliased_variable,omitempty"`
	Scope              APIVariableScopeEnum `json:"scope"`
	VariableType       APIVariableTypeEnum  `json:"variable_type"`
	// The id of the service referenced by this variable.
	ServiceId *string `json:"service_id,omitempty"`
	// The name of the service referenced by this variable.
	ServiceName *string                `json:"service_name,omitempty"`
	ServiceType *LinkedServiceTypeEnum `json:"service_type,omitempty"`
	// Entity that created/own the variable (i.e: Qovery, Doppler)
	OwnedBy                   *string `json:"owned_by,omitempty"`
	IsSecret                  bool    `json:"is_secret"`
	Description               *string `json:"description,omitempty"`
	EnableInterpolationInFile *bool   `json:"enable_interpolation_in_file,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _VariableResponse VariableResponse

// NewVariableResponse instantiates a new VariableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableResponse(id string, createdAt time.Time, key string, value NullableString, scope APIVariableScopeEnum, variableType APIVariableTypeEnum, isSecret bool) *VariableResponse {
	this := VariableResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Key = key
	this.Value = value
	this.Scope = scope
	this.VariableType = variableType
	this.IsSecret = isSecret
	return &this
}

// NewVariableResponseWithDefaults instantiates a new VariableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableResponseWithDefaults() *VariableResponse {
	this := VariableResponse{}
	return &this
}

// GetId returns the Id field value
func (o *VariableResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VariableResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *VariableResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *VariableResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VariableResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VariableResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VariableResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetKey returns the Key field value
func (o *VariableResponse) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *VariableResponse) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VariableResponse) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariableResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *VariableResponse) SetValue(v string) {
	o.Value.Set(&v)
}

// GetMountPath returns the MountPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VariableResponse) GetMountPath() string {
	if o == nil || IsNil(o.MountPath.Get()) {
		var ret string
		return ret
	}
	return *o.MountPath.Get()
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariableResponse) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountPath.Get(), o.MountPath.IsSet()
}

// HasMountPath returns a boolean if a field has been set.
func (o *VariableResponse) HasMountPath() bool {
	if o != nil && o.MountPath.IsSet() {
		return true
	}

	return false
}

// SetMountPath gets a reference to the given NullableString and assigns it to the MountPath field.
func (o *VariableResponse) SetMountPath(v string) {
	o.MountPath.Set(&v)
}

// SetMountPathNil sets the value for MountPath to be an explicit nil
func (o *VariableResponse) SetMountPathNil() {
	o.MountPath.Set(nil)
}

// UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
func (o *VariableResponse) UnsetMountPath() {
	o.MountPath.Unset()
}

// GetOverriddenVariable returns the OverriddenVariable field value if set, zero value otherwise.
func (o *VariableResponse) GetOverriddenVariable() VariableOverride {
	if o == nil || IsNil(o.OverriddenVariable) {
		var ret VariableOverride
		return ret
	}
	return *o.OverriddenVariable
}

// GetOverriddenVariableOk returns a tuple with the OverriddenVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetOverriddenVariableOk() (*VariableOverride, bool) {
	if o == nil || IsNil(o.OverriddenVariable) {
		return nil, false
	}
	return o.OverriddenVariable, true
}

// HasOverriddenVariable returns a boolean if a field has been set.
func (o *VariableResponse) HasOverriddenVariable() bool {
	if o != nil && !IsNil(o.OverriddenVariable) {
		return true
	}

	return false
}

// SetOverriddenVariable gets a reference to the given VariableOverride and assigns it to the OverriddenVariable field.
func (o *VariableResponse) SetOverriddenVariable(v VariableOverride) {
	o.OverriddenVariable = &v
}

// GetAliasedVariable returns the AliasedVariable field value if set, zero value otherwise.
func (o *VariableResponse) GetAliasedVariable() VariableAlias {
	if o == nil || IsNil(o.AliasedVariable) {
		var ret VariableAlias
		return ret
	}
	return *o.AliasedVariable
}

// GetAliasedVariableOk returns a tuple with the AliasedVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetAliasedVariableOk() (*VariableAlias, bool) {
	if o == nil || IsNil(o.AliasedVariable) {
		return nil, false
	}
	return o.AliasedVariable, true
}

// HasAliasedVariable returns a boolean if a field has been set.
func (o *VariableResponse) HasAliasedVariable() bool {
	if o != nil && !IsNil(o.AliasedVariable) {
		return true
	}

	return false
}

// SetAliasedVariable gets a reference to the given VariableAlias and assigns it to the AliasedVariable field.
func (o *VariableResponse) SetAliasedVariable(v VariableAlias) {
	o.AliasedVariable = &v
}

// GetScope returns the Scope field value
func (o *VariableResponse) GetScope() APIVariableScopeEnum {
	if o == nil {
		var ret APIVariableScopeEnum
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetScopeOk() (*APIVariableScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *VariableResponse) SetScope(v APIVariableScopeEnum) {
	o.Scope = v
}

// GetVariableType returns the VariableType field value
func (o *VariableResponse) GetVariableType() APIVariableTypeEnum {
	if o == nil {
		var ret APIVariableTypeEnum
		return ret
	}

	return o.VariableType
}

// GetVariableTypeOk returns a tuple with the VariableType field value
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetVariableTypeOk() (*APIVariableTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariableType, true
}

// SetVariableType sets field value
func (o *VariableResponse) SetVariableType(v APIVariableTypeEnum) {
	o.VariableType = v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *VariableResponse) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *VariableResponse) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *VariableResponse) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *VariableResponse) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *VariableResponse) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *VariableResponse) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *VariableResponse) GetServiceType() LinkedServiceTypeEnum {
	if o == nil || IsNil(o.ServiceType) {
		var ret LinkedServiceTypeEnum
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetServiceTypeOk() (*LinkedServiceTypeEnum, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *VariableResponse) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given LinkedServiceTypeEnum and assigns it to the ServiceType field.
func (o *VariableResponse) SetServiceType(v LinkedServiceTypeEnum) {
	o.ServiceType = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *VariableResponse) GetOwnedBy() string {
	if o == nil || IsNil(o.OwnedBy) {
		var ret string
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetOwnedByOk() (*string, bool) {
	if o == nil || IsNil(o.OwnedBy) {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *VariableResponse) HasOwnedBy() bool {
	if o != nil && !IsNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given string and assigns it to the OwnedBy field.
func (o *VariableResponse) SetOwnedBy(v string) {
	o.OwnedBy = &v
}

// GetIsSecret returns the IsSecret field value
func (o *VariableResponse) GetIsSecret() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSecret
}

// GetIsSecretOk returns a tuple with the IsSecret field value
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetIsSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSecret, true
}

// SetIsSecret sets field value
func (o *VariableResponse) SetIsSecret(v bool) {
	o.IsSecret = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VariableResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VariableResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VariableResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnableInterpolationInFile returns the EnableInterpolationInFile field value if set, zero value otherwise.
func (o *VariableResponse) GetEnableInterpolationInFile() bool {
	if o == nil || IsNil(o.EnableInterpolationInFile) {
		var ret bool
		return ret
	}
	return *o.EnableInterpolationInFile
}

// GetEnableInterpolationInFileOk returns a tuple with the EnableInterpolationInFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponse) GetEnableInterpolationInFileOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableInterpolationInFile) {
		return nil, false
	}
	return o.EnableInterpolationInFile, true
}

// HasEnableInterpolationInFile returns a boolean if a field has been set.
func (o *VariableResponse) HasEnableInterpolationInFile() bool {
	if o != nil && !IsNil(o.EnableInterpolationInFile) {
		return true
	}

	return false
}

// SetEnableInterpolationInFile gets a reference to the given bool and assigns it to the EnableInterpolationInFile field.
func (o *VariableResponse) SetEnableInterpolationInFile(v bool) {
	o.EnableInterpolationInFile = &v
}

func (o VariableResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value.Get()
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
	toSerialize["variable_type"] = o.VariableType
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
	toSerialize["is_secret"] = o.IsSecret
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EnableInterpolationInFile) {
		toSerialize["enable_interpolation_in_file"] = o.EnableInterpolationInFile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VariableResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"key",
		"value",
		"scope",
		"variable_type",
		"is_secret",
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

	varVariableResponse := _VariableResponse{}

	err = json.Unmarshal(data, &varVariableResponse)

	if err != nil {
		return err
	}

	*o = VariableResponse(varVariableResponse)

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
		delete(additionalProperties, "is_secret")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enable_interpolation_in_file")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVariableResponse struct {
	value *VariableResponse
	isSet bool
}

func (v NullableVariableResponse) Get() *VariableResponse {
	return v.value
}

func (v *NullableVariableResponse) Set(val *VariableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableResponse(val *VariableResponse) *NullableVariableResponse {
	return &NullableVariableResponse{value: val, isSet: true}
}

func (v NullableVariableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
