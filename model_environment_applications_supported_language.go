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

// checks if the EnvironmentApplicationsSupportedLanguage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentApplicationsSupportedLanguage{}

// EnvironmentApplicationsSupportedLanguage struct for EnvironmentApplicationsSupportedLanguage
type EnvironmentApplicationsSupportedLanguage struct {
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _EnvironmentApplicationsSupportedLanguage EnvironmentApplicationsSupportedLanguage

// NewEnvironmentApplicationsSupportedLanguage instantiates a new EnvironmentApplicationsSupportedLanguage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentApplicationsSupportedLanguage(name string) *EnvironmentApplicationsSupportedLanguage {
	this := EnvironmentApplicationsSupportedLanguage{}
	this.Name = name
	return &this
}

// NewEnvironmentApplicationsSupportedLanguageWithDefaults instantiates a new EnvironmentApplicationsSupportedLanguage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentApplicationsSupportedLanguageWithDefaults() *EnvironmentApplicationsSupportedLanguage {
	this := EnvironmentApplicationsSupportedLanguage{}
	return &this
}

// GetName returns the Name field value
func (o *EnvironmentApplicationsSupportedLanguage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsSupportedLanguage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentApplicationsSupportedLanguage) SetName(v string) {
	o.Name = v
}

func (o EnvironmentApplicationsSupportedLanguage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentApplicationsSupportedLanguage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvironmentApplicationsSupportedLanguage) UnmarshalJSON(data []byte) (err error) {
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

	varEnvironmentApplicationsSupportedLanguage := _EnvironmentApplicationsSupportedLanguage{}

	err = json.Unmarshal(data, &varEnvironmentApplicationsSupportedLanguage)

	if err != nil {
		return err
	}

	*o = EnvironmentApplicationsSupportedLanguage(varEnvironmentApplicationsSupportedLanguage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnvironmentApplicationsSupportedLanguage struct {
	value *EnvironmentApplicationsSupportedLanguage
	isSet bool
}

func (v NullableEnvironmentApplicationsSupportedLanguage) Get() *EnvironmentApplicationsSupportedLanguage {
	return v.value
}

func (v *NullableEnvironmentApplicationsSupportedLanguage) Set(val *EnvironmentApplicationsSupportedLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentApplicationsSupportedLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentApplicationsSupportedLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentApplicationsSupportedLanguage(val *EnvironmentApplicationsSupportedLanguage) *NullableEnvironmentApplicationsSupportedLanguage {
	return &NullableEnvironmentApplicationsSupportedLanguage{value: val, isSet: true}
}

func (v NullableEnvironmentApplicationsSupportedLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentApplicationsSupportedLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
