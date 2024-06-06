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

// checks if the HelmDeploymentRestrictionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmDeploymentRestrictionRequest{}

// HelmDeploymentRestrictionRequest struct for HelmDeploymentRestrictionRequest
type HelmDeploymentRestrictionRequest struct {
	Mode DeploymentRestrictionModeEnum `json:"mode"`
	Type DeploymentRestrictionTypeEnum `json:"type"`
	// For `PATH` restrictions, the value must not start with `/`
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _HelmDeploymentRestrictionRequest HelmDeploymentRestrictionRequest

// NewHelmDeploymentRestrictionRequest instantiates a new HelmDeploymentRestrictionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmDeploymentRestrictionRequest(mode DeploymentRestrictionModeEnum, type_ DeploymentRestrictionTypeEnum, value string) *HelmDeploymentRestrictionRequest {
	this := HelmDeploymentRestrictionRequest{}
	this.Mode = mode
	this.Type = type_
	this.Value = value
	return &this
}

// NewHelmDeploymentRestrictionRequestWithDefaults instantiates a new HelmDeploymentRestrictionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmDeploymentRestrictionRequestWithDefaults() *HelmDeploymentRestrictionRequest {
	this := HelmDeploymentRestrictionRequest{}
	return &this
}

// GetMode returns the Mode field value
func (o *HelmDeploymentRestrictionRequest) GetMode() DeploymentRestrictionModeEnum {
	if o == nil {
		var ret DeploymentRestrictionModeEnum
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *HelmDeploymentRestrictionRequest) GetModeOk() (*DeploymentRestrictionModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *HelmDeploymentRestrictionRequest) SetMode(v DeploymentRestrictionModeEnum) {
	o.Mode = v
}

// GetType returns the Type field value
func (o *HelmDeploymentRestrictionRequest) GetType() DeploymentRestrictionTypeEnum {
	if o == nil {
		var ret DeploymentRestrictionTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HelmDeploymentRestrictionRequest) GetTypeOk() (*DeploymentRestrictionTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HelmDeploymentRestrictionRequest) SetType(v DeploymentRestrictionTypeEnum) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *HelmDeploymentRestrictionRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *HelmDeploymentRestrictionRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *HelmDeploymentRestrictionRequest) SetValue(v string) {
	o.Value = v
}

func (o HelmDeploymentRestrictionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmDeploymentRestrictionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mode"] = o.Mode
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmDeploymentRestrictionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mode",
		"type",
		"value",
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

	varHelmDeploymentRestrictionRequest := _HelmDeploymentRestrictionRequest{}

	err = json.Unmarshal(data, &varHelmDeploymentRestrictionRequest)

	if err != nil {
		return err
	}

	*o = HelmDeploymentRestrictionRequest(varHelmDeploymentRestrictionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mode")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmDeploymentRestrictionRequest struct {
	value *HelmDeploymentRestrictionRequest
	isSet bool
}

func (v NullableHelmDeploymentRestrictionRequest) Get() *HelmDeploymentRestrictionRequest {
	return v.value
}

func (v *NullableHelmDeploymentRestrictionRequest) Set(val *HelmDeploymentRestrictionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmDeploymentRestrictionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmDeploymentRestrictionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmDeploymentRestrictionRequest(val *HelmDeploymentRestrictionRequest) *NullableHelmDeploymentRestrictionRequest {
	return &NullableHelmDeploymentRestrictionRequest{value: val, isSet: true}
}

func (v NullableHelmDeploymentRestrictionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmDeploymentRestrictionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
