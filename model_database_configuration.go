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

// checks if the DatabaseConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseConfiguration{}

// DatabaseConfiguration struct for DatabaseConfiguration
type DatabaseConfiguration struct {
	DatabaseType *DatabaseTypeEnum     `json:"database_type,omitempty"`
	Version      []DatabaseVersionMode `json:"version,omitempty"`
}

// NewDatabaseConfiguration instantiates a new DatabaseConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseConfiguration() *DatabaseConfiguration {
	this := DatabaseConfiguration{}
	return &this
}

// NewDatabaseConfigurationWithDefaults instantiates a new DatabaseConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseConfigurationWithDefaults() *DatabaseConfiguration {
	this := DatabaseConfiguration{}
	return &this
}

// GetDatabaseType returns the DatabaseType field value if set, zero value otherwise.
func (o *DatabaseConfiguration) GetDatabaseType() DatabaseTypeEnum {
	if o == nil || IsNil(o.DatabaseType) {
		var ret DatabaseTypeEnum
		return ret
	}
	return *o.DatabaseType
}

// GetDatabaseTypeOk returns a tuple with the DatabaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseConfiguration) GetDatabaseTypeOk() (*DatabaseTypeEnum, bool) {
	if o == nil || IsNil(o.DatabaseType) {
		return nil, false
	}
	return o.DatabaseType, true
}

// HasDatabaseType returns a boolean if a field has been set.
func (o *DatabaseConfiguration) HasDatabaseType() bool {
	if o != nil && !IsNil(o.DatabaseType) {
		return true
	}

	return false
}

// SetDatabaseType gets a reference to the given DatabaseTypeEnum and assigns it to the DatabaseType field.
func (o *DatabaseConfiguration) SetDatabaseType(v DatabaseTypeEnum) {
	o.DatabaseType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DatabaseConfiguration) GetVersion() []DatabaseVersionMode {
	if o == nil || IsNil(o.Version) {
		var ret []DatabaseVersionMode
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseConfiguration) GetVersionOk() ([]DatabaseVersionMode, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DatabaseConfiguration) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given []DatabaseVersionMode and assigns it to the Version field.
func (o *DatabaseConfiguration) SetVersion(v []DatabaseVersionMode) {
	o.Version = v
}

func (o DatabaseConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatabaseType) {
		toSerialize["database_type"] = o.DatabaseType
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableDatabaseConfiguration struct {
	value *DatabaseConfiguration
	isSet bool
}

func (v NullableDatabaseConfiguration) Get() *DatabaseConfiguration {
	return v.value
}

func (v *NullableDatabaseConfiguration) Set(val *DatabaseConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseConfiguration(val *DatabaseConfiguration) *NullableDatabaseConfiguration {
	return &NullableDatabaseConfiguration{value: val, isSet: true}
}

func (v NullableDatabaseConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
