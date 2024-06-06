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

// checks if the ClusterRoutingTableResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterRoutingTableResultsInner{}

// ClusterRoutingTableResultsInner struct for ClusterRoutingTableResultsInner
type ClusterRoutingTableResultsInner struct {
	Destination          string `json:"destination"`
	Target               string `json:"target"`
	Description          string `json:"description"`
	AdditionalProperties map[string]interface{}
}

type _ClusterRoutingTableResultsInner ClusterRoutingTableResultsInner

// NewClusterRoutingTableResultsInner instantiates a new ClusterRoutingTableResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRoutingTableResultsInner(destination string, target string, description string) *ClusterRoutingTableResultsInner {
	this := ClusterRoutingTableResultsInner{}
	this.Destination = destination
	this.Target = target
	this.Description = description
	return &this
}

// NewClusterRoutingTableResultsInnerWithDefaults instantiates a new ClusterRoutingTableResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterRoutingTableResultsInnerWithDefaults() *ClusterRoutingTableResultsInner {
	this := ClusterRoutingTableResultsInner{}
	return &this
}

// GetDestination returns the Destination field value
func (o *ClusterRoutingTableResultsInner) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *ClusterRoutingTableResultsInner) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *ClusterRoutingTableResultsInner) SetDestination(v string) {
	o.Destination = v
}

// GetTarget returns the Target field value
func (o *ClusterRoutingTableResultsInner) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ClusterRoutingTableResultsInner) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *ClusterRoutingTableResultsInner) SetTarget(v string) {
	o.Target = v
}

// GetDescription returns the Description field value
func (o *ClusterRoutingTableResultsInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ClusterRoutingTableResultsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ClusterRoutingTableResultsInner) SetDescription(v string) {
	o.Description = v
}

func (o ClusterRoutingTableResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterRoutingTableResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination"] = o.Destination
	toSerialize["target"] = o.Target
	toSerialize["description"] = o.Description

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClusterRoutingTableResultsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination",
		"target",
		"description",
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

	varClusterRoutingTableResultsInner := _ClusterRoutingTableResultsInner{}

	err = json.Unmarshal(data, &varClusterRoutingTableResultsInner)

	if err != nil {
		return err
	}

	*o = ClusterRoutingTableResultsInner(varClusterRoutingTableResultsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "destination")
		delete(additionalProperties, "target")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClusterRoutingTableResultsInner struct {
	value *ClusterRoutingTableResultsInner
	isSet bool
}

func (v NullableClusterRoutingTableResultsInner) Get() *ClusterRoutingTableResultsInner {
	return v.value
}

func (v *NullableClusterRoutingTableResultsInner) Set(val *ClusterRoutingTableResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRoutingTableResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRoutingTableResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRoutingTableResultsInner(val *ClusterRoutingTableResultsInner) *NullableClusterRoutingTableResultsInner {
	return &NullableClusterRoutingTableResultsInner{value: val, isSet: true}
}

func (v NullableClusterRoutingTableResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRoutingTableResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
