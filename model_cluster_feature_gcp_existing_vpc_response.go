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

// checks if the ClusterFeatureGcpExistingVpcResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterFeatureGcpExistingVpcResponse{}

// ClusterFeatureGcpExistingVpcResponse struct for ClusterFeatureGcpExistingVpcResponse
type ClusterFeatureGcpExistingVpcResponse struct {
	Type                 ClusterFeatureResponseTypeEnum `json:"type"`
	Value                ClusterFeatureGcpExistingVpc   `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _ClusterFeatureGcpExistingVpcResponse ClusterFeatureGcpExistingVpcResponse

// NewClusterFeatureGcpExistingVpcResponse instantiates a new ClusterFeatureGcpExistingVpcResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterFeatureGcpExistingVpcResponse(type_ ClusterFeatureResponseTypeEnum, value ClusterFeatureGcpExistingVpc) *ClusterFeatureGcpExistingVpcResponse {
	this := ClusterFeatureGcpExistingVpcResponse{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewClusterFeatureGcpExistingVpcResponseWithDefaults instantiates a new ClusterFeatureGcpExistingVpcResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterFeatureGcpExistingVpcResponseWithDefaults() *ClusterFeatureGcpExistingVpcResponse {
	this := ClusterFeatureGcpExistingVpcResponse{}
	return &this
}

// GetType returns the Type field value
func (o *ClusterFeatureGcpExistingVpcResponse) GetType() ClusterFeatureResponseTypeEnum {
	if o == nil {
		var ret ClusterFeatureResponseTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ClusterFeatureGcpExistingVpcResponse) GetTypeOk() (*ClusterFeatureResponseTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ClusterFeatureGcpExistingVpcResponse) SetType(v ClusterFeatureResponseTypeEnum) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *ClusterFeatureGcpExistingVpcResponse) GetValue() ClusterFeatureGcpExistingVpc {
	if o == nil {
		var ret ClusterFeatureGcpExistingVpc
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ClusterFeatureGcpExistingVpcResponse) GetValueOk() (*ClusterFeatureGcpExistingVpc, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ClusterFeatureGcpExistingVpcResponse) SetValue(v ClusterFeatureGcpExistingVpc) {
	o.Value = v
}

func (o ClusterFeatureGcpExistingVpcResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterFeatureGcpExistingVpcResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClusterFeatureGcpExistingVpcResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varClusterFeatureGcpExistingVpcResponse := _ClusterFeatureGcpExistingVpcResponse{}

	err = json.Unmarshal(data, &varClusterFeatureGcpExistingVpcResponse)

	if err != nil {
		return err
	}

	*o = ClusterFeatureGcpExistingVpcResponse(varClusterFeatureGcpExistingVpcResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClusterFeatureGcpExistingVpcResponse struct {
	value *ClusterFeatureGcpExistingVpcResponse
	isSet bool
}

func (v NullableClusterFeatureGcpExistingVpcResponse) Get() *ClusterFeatureGcpExistingVpcResponse {
	return v.value
}

func (v *NullableClusterFeatureGcpExistingVpcResponse) Set(val *ClusterFeatureGcpExistingVpcResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterFeatureGcpExistingVpcResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterFeatureGcpExistingVpcResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterFeatureGcpExistingVpcResponse(val *ClusterFeatureGcpExistingVpcResponse) *NullableClusterFeatureGcpExistingVpcResponse {
	return &NullableClusterFeatureGcpExistingVpcResponse{value: val, isSet: true}
}

func (v NullableClusterFeatureGcpExistingVpcResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterFeatureGcpExistingVpcResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
