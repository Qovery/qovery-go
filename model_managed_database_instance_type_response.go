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

// checks if the ManagedDatabaseInstanceTypeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedDatabaseInstanceTypeResponse{}

// ManagedDatabaseInstanceTypeResponse struct for ManagedDatabaseInstanceTypeResponse
type ManagedDatabaseInstanceTypeResponse struct {
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ManagedDatabaseInstanceTypeResponse ManagedDatabaseInstanceTypeResponse

// NewManagedDatabaseInstanceTypeResponse instantiates a new ManagedDatabaseInstanceTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedDatabaseInstanceTypeResponse(name string) *ManagedDatabaseInstanceTypeResponse {
	this := ManagedDatabaseInstanceTypeResponse{}
	this.Name = name
	return &this
}

// NewManagedDatabaseInstanceTypeResponseWithDefaults instantiates a new ManagedDatabaseInstanceTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedDatabaseInstanceTypeResponseWithDefaults() *ManagedDatabaseInstanceTypeResponse {
	this := ManagedDatabaseInstanceTypeResponse{}
	return &this
}

// GetName returns the Name field value
func (o *ManagedDatabaseInstanceTypeResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ManagedDatabaseInstanceTypeResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ManagedDatabaseInstanceTypeResponse) SetName(v string) {
	o.Name = v
}

func (o ManagedDatabaseInstanceTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedDatabaseInstanceTypeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagedDatabaseInstanceTypeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varManagedDatabaseInstanceTypeResponse := _ManagedDatabaseInstanceTypeResponse{}

	err = json.Unmarshal(data, &varManagedDatabaseInstanceTypeResponse)

	if err != nil {
		return err
	}

	*o = ManagedDatabaseInstanceTypeResponse(varManagedDatabaseInstanceTypeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagedDatabaseInstanceTypeResponse struct {
	value *ManagedDatabaseInstanceTypeResponse
	isSet bool
}

func (v NullableManagedDatabaseInstanceTypeResponse) Get() *ManagedDatabaseInstanceTypeResponse {
	return v.value
}

func (v *NullableManagedDatabaseInstanceTypeResponse) Set(val *ManagedDatabaseInstanceTypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedDatabaseInstanceTypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedDatabaseInstanceTypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedDatabaseInstanceTypeResponse(val *ManagedDatabaseInstanceTypeResponse) *NullableManagedDatabaseInstanceTypeResponse {
	return &NullableManagedDatabaseInstanceTypeResponse{value: val, isSet: true}
}

func (v NullableManagedDatabaseInstanceTypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedDatabaseInstanceTypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
