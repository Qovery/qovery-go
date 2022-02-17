/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet. 

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// GenericObjectCurrentCostResponse struct for GenericObjectCurrentCostResponse
type GenericObjectCurrentCostResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ConsumedTimeInSeconds int32 `json:"consumed_time_in_seconds"`
	Cost CostResponse `json:"cost"`
}

// NewGenericObjectCurrentCostResponse instantiates a new GenericObjectCurrentCostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericObjectCurrentCostResponse(id string, name string, consumedTimeInSeconds int32, cost CostResponse) *GenericObjectCurrentCostResponse {
	this := GenericObjectCurrentCostResponse{}
	this.Id = id
	this.Name = name
	this.ConsumedTimeInSeconds = consumedTimeInSeconds
	this.Cost = cost
	return &this
}

// NewGenericObjectCurrentCostResponseWithDefaults instantiates a new GenericObjectCurrentCostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericObjectCurrentCostResponseWithDefaults() *GenericObjectCurrentCostResponse {
	this := GenericObjectCurrentCostResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GenericObjectCurrentCostResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GenericObjectCurrentCostResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GenericObjectCurrentCostResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GenericObjectCurrentCostResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GenericObjectCurrentCostResponse) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GenericObjectCurrentCostResponse) SetName(v string) {
	o.Name = v
}

// GetConsumedTimeInSeconds returns the ConsumedTimeInSeconds field value
func (o *GenericObjectCurrentCostResponse) GetConsumedTimeInSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConsumedTimeInSeconds
}

// GetConsumedTimeInSecondsOk returns a tuple with the ConsumedTimeInSeconds field value
// and a boolean to check if the value has been set.
func (o *GenericObjectCurrentCostResponse) GetConsumedTimeInSecondsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConsumedTimeInSeconds, true
}

// SetConsumedTimeInSeconds sets field value
func (o *GenericObjectCurrentCostResponse) SetConsumedTimeInSeconds(v int32) {
	o.ConsumedTimeInSeconds = v
}

// GetCost returns the Cost field value
func (o *GenericObjectCurrentCostResponse) GetCost() CostResponse {
	if o == nil {
		var ret CostResponse
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *GenericObjectCurrentCostResponse) GetCostOk() (*CostResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value
func (o *GenericObjectCurrentCostResponse) SetCost(v CostResponse) {
	o.Cost = v
}

func (o GenericObjectCurrentCostResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["consumed_time_in_seconds"] = o.ConsumedTimeInSeconds
	}
	if true {
		toSerialize["cost"] = o.Cost
	}
	return json.Marshal(toSerialize)
}

type NullableGenericObjectCurrentCostResponse struct {
	value *GenericObjectCurrentCostResponse
	isSet bool
}

func (v NullableGenericObjectCurrentCostResponse) Get() *GenericObjectCurrentCostResponse {
	return v.value
}

func (v *NullableGenericObjectCurrentCostResponse) Set(val *GenericObjectCurrentCostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericObjectCurrentCostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericObjectCurrentCostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericObjectCurrentCostResponse(val *GenericObjectCurrentCostResponse) *NullableGenericObjectCurrentCostResponse {
	return &NullableGenericObjectCurrentCostResponse{value: val, isSet: true}
}

func (v NullableGenericObjectCurrentCostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericObjectCurrentCostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


