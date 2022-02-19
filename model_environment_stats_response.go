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

// EnvironmentStatsResponse struct for EnvironmentStatsResponse
type EnvironmentStatsResponse struct {
	Id                 string   `json:"id"`
	ServiceTotalNumber *float32 `json:"service_total_number,omitempty"`
}

// NewEnvironmentStatsResponse instantiates a new EnvironmentStatsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentStatsResponse(id string) *EnvironmentStatsResponse {
	this := EnvironmentStatsResponse{}
	this.Id = id
	return &this
}

// NewEnvironmentStatsResponseWithDefaults instantiates a new EnvironmentStatsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentStatsResponseWithDefaults() *EnvironmentStatsResponse {
	this := EnvironmentStatsResponse{}
	return &this
}

// GetId returns the Id field value
func (o *EnvironmentStatsResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatsResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentStatsResponse) SetId(v string) {
	o.Id = v
}

// GetServiceTotalNumber returns the ServiceTotalNumber field value if set, zero value otherwise.
func (o *EnvironmentStatsResponse) GetServiceTotalNumber() float32 {
	if o == nil || o.ServiceTotalNumber == nil {
		var ret float32
		return ret
	}
	return *o.ServiceTotalNumber
}

// GetServiceTotalNumberOk returns a tuple with the ServiceTotalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentStatsResponse) GetServiceTotalNumberOk() (*float32, bool) {
	if o == nil || o.ServiceTotalNumber == nil {
		return nil, false
	}
	return o.ServiceTotalNumber, true
}

// HasServiceTotalNumber returns a boolean if a field has been set.
func (o *EnvironmentStatsResponse) HasServiceTotalNumber() bool {
	if o != nil && o.ServiceTotalNumber != nil {
		return true
	}

	return false
}

// SetServiceTotalNumber gets a reference to the given float32 and assigns it to the ServiceTotalNumber field.
func (o *EnvironmentStatsResponse) SetServiceTotalNumber(v float32) {
	o.ServiceTotalNumber = &v
}

func (o EnvironmentStatsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.ServiceTotalNumber != nil {
		toSerialize["service_total_number"] = o.ServiceTotalNumber
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentStatsResponse struct {
	value *EnvironmentStatsResponse
	isSet bool
}

func (v NullableEnvironmentStatsResponse) Get() *EnvironmentStatsResponse {
	return v.value
}

func (v *NullableEnvironmentStatsResponse) Set(val *EnvironmentStatsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentStatsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentStatsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentStatsResponse(val *EnvironmentStatsResponse) *NullableEnvironmentStatsResponse {
	return &NullableEnvironmentStatsResponse{value: val, isSet: true}
}

func (v NullableEnvironmentStatsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentStatsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
