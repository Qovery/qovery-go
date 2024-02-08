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

// checks if the OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner{}

// OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner struct for OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner
type OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner struct {
	EnvironmentType      *EnvironmentModeEnum                     `json:"environment_type,omitempty"`
	Permission           *OrganizationCustomRoleProjectPermission `json:"permission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner

// NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner instantiates a new OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner() *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner {
	this := OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner{}
	return &this
}

// NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInnerWithDefaults instantiates a new OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInnerWithDefaults() *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner {
	this := OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner{}
	return &this
}

// GetEnvironmentType returns the EnvironmentType field value if set, zero value otherwise.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) GetEnvironmentType() EnvironmentModeEnum {
	if o == nil || IsNil(o.EnvironmentType) {
		var ret EnvironmentModeEnum
		return ret
	}
	return *o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) GetEnvironmentTypeOk() (*EnvironmentModeEnum, bool) {
	if o == nil || IsNil(o.EnvironmentType) {
		return nil, false
	}
	return o.EnvironmentType, true
}

// HasEnvironmentType returns a boolean if a field has been set.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) HasEnvironmentType() bool {
	if o != nil && !IsNil(o.EnvironmentType) {
		return true
	}

	return false
}

// SetEnvironmentType gets a reference to the given EnvironmentModeEnum and assigns it to the EnvironmentType field.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) SetEnvironmentType(v EnvironmentModeEnum) {
	o.EnvironmentType = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) GetPermission() OrganizationCustomRoleProjectPermission {
	if o == nil || IsNil(o.Permission) {
		var ret OrganizationCustomRoleProjectPermission
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) GetPermissionOk() (*OrganizationCustomRoleProjectPermission, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given OrganizationCustomRoleProjectPermission and assigns it to the Permission field.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) SetPermission(v OrganizationCustomRoleProjectPermission) {
	o.Permission = &v
}

func (o OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvironmentType) {
		toSerialize["environment_type"] = o.EnvironmentType
	}
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) UnmarshalJSON(data []byte) (err error) {
	varOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner := _OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner{}

	err = json.Unmarshal(data, &varOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner)

	if err != nil {
		return err
	}

	*o = OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner(varOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environment_type")
		delete(additionalProperties, "permission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner struct {
	value *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner
	isSet bool
}

func (v NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) Get() *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner {
	return v.value
}

func (v *NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) Set(val *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner(val *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) *NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner {
	return &NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner{value: val, isSet: true}
}

func (v NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
