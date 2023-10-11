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

// checks if the EnvironmentDatabasesCurrentMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentDatabasesCurrentMetric{}

// EnvironmentDatabasesCurrentMetric struct for EnvironmentDatabasesCurrentMetric
type EnvironmentDatabasesCurrentMetric struct {
	Database *string                                   `json:"database,omitempty"`
	Cpu      *EnvironmentDatabasesCurrentMetricCpu     `json:"cpu,omitempty"`
	Memory   *EnvironmentDatabasesCurrentMetricMemory  `json:"memory,omitempty"`
	Storage  *EnvironmentDatabasesCurrentMetricStorage `json:"storage,omitempty"`
}

// NewEnvironmentDatabasesCurrentMetric instantiates a new EnvironmentDatabasesCurrentMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentDatabasesCurrentMetric() *EnvironmentDatabasesCurrentMetric {
	this := EnvironmentDatabasesCurrentMetric{}
	return &this
}

// NewEnvironmentDatabasesCurrentMetricWithDefaults instantiates a new EnvironmentDatabasesCurrentMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentDatabasesCurrentMetricWithDefaults() *EnvironmentDatabasesCurrentMetric {
	this := EnvironmentDatabasesCurrentMetric{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *EnvironmentDatabasesCurrentMetric) GetDatabase() string {
	if o == nil || IsNil(o.Database) {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDatabasesCurrentMetric) GetDatabaseOk() (*string, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *EnvironmentDatabasesCurrentMetric) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *EnvironmentDatabasesCurrentMetric) SetDatabase(v string) {
	o.Database = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *EnvironmentDatabasesCurrentMetric) GetCpu() EnvironmentDatabasesCurrentMetricCpu {
	if o == nil || IsNil(o.Cpu) {
		var ret EnvironmentDatabasesCurrentMetricCpu
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDatabasesCurrentMetric) GetCpuOk() (*EnvironmentDatabasesCurrentMetricCpu, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *EnvironmentDatabasesCurrentMetric) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given EnvironmentDatabasesCurrentMetricCpu and assigns it to the Cpu field.
func (o *EnvironmentDatabasesCurrentMetric) SetCpu(v EnvironmentDatabasesCurrentMetricCpu) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *EnvironmentDatabasesCurrentMetric) GetMemory() EnvironmentDatabasesCurrentMetricMemory {
	if o == nil || IsNil(o.Memory) {
		var ret EnvironmentDatabasesCurrentMetricMemory
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDatabasesCurrentMetric) GetMemoryOk() (*EnvironmentDatabasesCurrentMetricMemory, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *EnvironmentDatabasesCurrentMetric) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given EnvironmentDatabasesCurrentMetricMemory and assigns it to the Memory field.
func (o *EnvironmentDatabasesCurrentMetric) SetMemory(v EnvironmentDatabasesCurrentMetricMemory) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *EnvironmentDatabasesCurrentMetric) GetStorage() EnvironmentDatabasesCurrentMetricStorage {
	if o == nil || IsNil(o.Storage) {
		var ret EnvironmentDatabasesCurrentMetricStorage
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDatabasesCurrentMetric) GetStorageOk() (*EnvironmentDatabasesCurrentMetricStorage, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *EnvironmentDatabasesCurrentMetric) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given EnvironmentDatabasesCurrentMetricStorage and assigns it to the Storage field.
func (o *EnvironmentDatabasesCurrentMetric) SetStorage(v EnvironmentDatabasesCurrentMetricStorage) {
	o.Storage = &v
}

func (o EnvironmentDatabasesCurrentMetric) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentDatabasesCurrentMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	return toSerialize, nil
}

type NullableEnvironmentDatabasesCurrentMetric struct {
	value *EnvironmentDatabasesCurrentMetric
	isSet bool
}

func (v NullableEnvironmentDatabasesCurrentMetric) Get() *EnvironmentDatabasesCurrentMetric {
	return v.value
}

func (v *NullableEnvironmentDatabasesCurrentMetric) Set(val *EnvironmentDatabasesCurrentMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentDatabasesCurrentMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentDatabasesCurrentMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentDatabasesCurrentMetric(val *EnvironmentDatabasesCurrentMetric) *NullableEnvironmentDatabasesCurrentMetric {
	return &NullableEnvironmentDatabasesCurrentMetric{value: val, isSet: true}
}

func (v NullableEnvironmentDatabasesCurrentMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentDatabasesCurrentMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
