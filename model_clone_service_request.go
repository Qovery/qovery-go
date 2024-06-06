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

// checks if the CloneServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloneServiceRequest{}

// CloneServiceRequest struct for CloneServiceRequest
type CloneServiceRequest struct {
	Name                 string `json:"name"`
	EnvironmentId        string `json:"environment_id"`
	AdditionalProperties map[string]interface{}
}

type _CloneServiceRequest CloneServiceRequest

// NewCloneServiceRequest instantiates a new CloneServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloneServiceRequest(name string, environmentId string) *CloneServiceRequest {
	this := CloneServiceRequest{}
	this.Name = name
	this.EnvironmentId = environmentId
	return &this
}

// NewCloneServiceRequestWithDefaults instantiates a new CloneServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloneServiceRequestWithDefaults() *CloneServiceRequest {
	this := CloneServiceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CloneServiceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloneServiceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloneServiceRequest) SetName(v string) {
	o.Name = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *CloneServiceRequest) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *CloneServiceRequest) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *CloneServiceRequest) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

func (o CloneServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloneServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["environment_id"] = o.EnvironmentId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloneServiceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"environment_id",
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

	varCloneServiceRequest := _CloneServiceRequest{}

	err = json.Unmarshal(data, &varCloneServiceRequest)

	if err != nil {
		return err
	}

	*o = CloneServiceRequest(varCloneServiceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "environment_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloneServiceRequest struct {
	value *CloneServiceRequest
	isSet bool
}

func (v NullableCloneServiceRequest) Get() *CloneServiceRequest {
	return v.value
}

func (v *NullableCloneServiceRequest) Set(val *CloneServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneServiceRequest(val *CloneServiceRequest) *NullableCloneServiceRequest {
	return &NullableCloneServiceRequest{value: val, isSet: true}
}

func (v NullableCloneServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
