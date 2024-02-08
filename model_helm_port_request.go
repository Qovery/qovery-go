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

// checks if the HelmPortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmPortRequest{}

// HelmPortRequest struct for HelmPortRequest
type HelmPortRequest struct {
	Ports                []HelmPortRequestPortsInner `json:"ports,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HelmPortRequest HelmPortRequest

// NewHelmPortRequest instantiates a new HelmPortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmPortRequest() *HelmPortRequest {
	this := HelmPortRequest{}
	return &this
}

// NewHelmPortRequestWithDefaults instantiates a new HelmPortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmPortRequestWithDefaults() *HelmPortRequest {
	this := HelmPortRequest{}
	return &this
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *HelmPortRequest) GetPorts() []HelmPortRequestPortsInner {
	if o == nil || IsNil(o.Ports) {
		var ret []HelmPortRequestPortsInner
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmPortRequest) GetPortsOk() ([]HelmPortRequestPortsInner, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *HelmPortRequest) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []HelmPortRequestPortsInner and assigns it to the Ports field.
func (o *HelmPortRequest) SetPorts(v []HelmPortRequestPortsInner) {
	o.Ports = v
}

func (o HelmPortRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmPortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmPortRequest) UnmarshalJSON(data []byte) (err error) {
	varHelmPortRequest := _HelmPortRequest{}

	err = json.Unmarshal(data, &varHelmPortRequest)

	if err != nil {
		return err
	}

	*o = HelmPortRequest(varHelmPortRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmPortRequest struct {
	value *HelmPortRequest
	isSet bool
}

func (v NullableHelmPortRequest) Get() *HelmPortRequest {
	return v.value
}

func (v *NullableHelmPortRequest) Set(val *HelmPortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmPortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmPortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmPortRequest(val *HelmPortRequest) *NullableHelmPortRequest {
	return &NullableHelmPortRequest{value: val, isSet: true}
}

func (v NullableHelmPortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmPortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
