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

// checks if the ProjectStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectStats{}

// ProjectStats struct for ProjectStats
type ProjectStats struct {
	Id                     string   `json:"id"`
	ServiceTotalNumber     *float32 `json:"service_total_number,omitempty"`
	EnvironmentTotalNumber *float32 `json:"environment_total_number,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _ProjectStats ProjectStats

// NewProjectStats instantiates a new ProjectStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectStats(id string) *ProjectStats {
	this := ProjectStats{}
	this.Id = id
	return &this
}

// NewProjectStatsWithDefaults instantiates a new ProjectStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectStatsWithDefaults() *ProjectStats {
	this := ProjectStats{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectStats) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectStats) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectStats) SetId(v string) {
	o.Id = v
}

// GetServiceTotalNumber returns the ServiceTotalNumber field value if set, zero value otherwise.
func (o *ProjectStats) GetServiceTotalNumber() float32 {
	if o == nil || IsNil(o.ServiceTotalNumber) {
		var ret float32
		return ret
	}
	return *o.ServiceTotalNumber
}

// GetServiceTotalNumberOk returns a tuple with the ServiceTotalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectStats) GetServiceTotalNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.ServiceTotalNumber) {
		return nil, false
	}
	return o.ServiceTotalNumber, true
}

// HasServiceTotalNumber returns a boolean if a field has been set.
func (o *ProjectStats) HasServiceTotalNumber() bool {
	if o != nil && !IsNil(o.ServiceTotalNumber) {
		return true
	}

	return false
}

// SetServiceTotalNumber gets a reference to the given float32 and assigns it to the ServiceTotalNumber field.
func (o *ProjectStats) SetServiceTotalNumber(v float32) {
	o.ServiceTotalNumber = &v
}

// GetEnvironmentTotalNumber returns the EnvironmentTotalNumber field value if set, zero value otherwise.
func (o *ProjectStats) GetEnvironmentTotalNumber() float32 {
	if o == nil || IsNil(o.EnvironmentTotalNumber) {
		var ret float32
		return ret
	}
	return *o.EnvironmentTotalNumber
}

// GetEnvironmentTotalNumberOk returns a tuple with the EnvironmentTotalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectStats) GetEnvironmentTotalNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.EnvironmentTotalNumber) {
		return nil, false
	}
	return o.EnvironmentTotalNumber, true
}

// HasEnvironmentTotalNumber returns a boolean if a field has been set.
func (o *ProjectStats) HasEnvironmentTotalNumber() bool {
	if o != nil && !IsNil(o.EnvironmentTotalNumber) {
		return true
	}

	return false
}

// SetEnvironmentTotalNumber gets a reference to the given float32 and assigns it to the EnvironmentTotalNumber field.
func (o *ProjectStats) SetEnvironmentTotalNumber(v float32) {
	o.EnvironmentTotalNumber = &v
}

func (o ProjectStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.ServiceTotalNumber) {
		toSerialize["service_total_number"] = o.ServiceTotalNumber
	}
	if !IsNil(o.EnvironmentTotalNumber) {
		toSerialize["environment_total_number"] = o.EnvironmentTotalNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varProjectStats := _ProjectStats{}

	err = json.Unmarshal(data, &varProjectStats)

	if err != nil {
		return err
	}

	*o = ProjectStats(varProjectStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "service_total_number")
		delete(additionalProperties, "environment_total_number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectStats struct {
	value *ProjectStats
	isSet bool
}

func (v NullableProjectStats) Get() *ProjectStats {
	return v.value
}

func (v *NullableProjectStats) Set(val *ProjectStats) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectStats) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectStats(val *ProjectStats) *NullableProjectStats {
	return &NullableProjectStats{value: val, isSet: true}
}

func (v NullableProjectStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
