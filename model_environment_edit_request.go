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

// checks if the EnvironmentEditRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentEditRequest{}

// EnvironmentEditRequest struct for EnvironmentEditRequest
type EnvironmentEditRequest struct {
	Name                 *string                    `json:"name,omitempty"`
	Mode                 *CreateEnvironmentModeEnum `json:"mode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnvironmentEditRequest EnvironmentEditRequest

// NewEnvironmentEditRequest instantiates a new EnvironmentEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentEditRequest() *EnvironmentEditRequest {
	this := EnvironmentEditRequest{}
	return &this
}

// NewEnvironmentEditRequestWithDefaults instantiates a new EnvironmentEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentEditRequestWithDefaults() *EnvironmentEditRequest {
	this := EnvironmentEditRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentEditRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentEditRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentEditRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentEditRequest) SetName(v string) {
	o.Name = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *EnvironmentEditRequest) GetMode() CreateEnvironmentModeEnum {
	if o == nil || IsNil(o.Mode) {
		var ret CreateEnvironmentModeEnum
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentEditRequest) GetModeOk() (*CreateEnvironmentModeEnum, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *EnvironmentEditRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given CreateEnvironmentModeEnum and assigns it to the Mode field.
func (o *EnvironmentEditRequest) SetMode(v CreateEnvironmentModeEnum) {
	o.Mode = &v
}

func (o EnvironmentEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentEditRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvironmentEditRequest) UnmarshalJSON(data []byte) (err error) {
	varEnvironmentEditRequest := _EnvironmentEditRequest{}

	err = json.Unmarshal(data, &varEnvironmentEditRequest)

	if err != nil {
		return err
	}

	*o = EnvironmentEditRequest(varEnvironmentEditRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "mode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnvironmentEditRequest struct {
	value *EnvironmentEditRequest
	isSet bool
}

func (v NullableEnvironmentEditRequest) Get() *EnvironmentEditRequest {
	return v.value
}

func (v *NullableEnvironmentEditRequest) Set(val *EnvironmentEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentEditRequest(val *EnvironmentEditRequest) *NullableEnvironmentEditRequest {
	return &NullableEnvironmentEditRequest{value: val, isSet: true}
}

func (v NullableEnvironmentEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
