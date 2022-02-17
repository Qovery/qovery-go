/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// MetricMemoryDatapointResponseList struct for MetricMemoryDatapointResponseList
type MetricMemoryDatapointResponseList struct {
	Results []MetricMemoryDatapointResponse `json:"results,omitempty"`
}

// NewMetricMemoryDatapointResponseList instantiates a new MetricMemoryDatapointResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricMemoryDatapointResponseList() *MetricMemoryDatapointResponseList {
	this := MetricMemoryDatapointResponseList{}
	return &this
}

// NewMetricMemoryDatapointResponseListWithDefaults instantiates a new MetricMemoryDatapointResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricMemoryDatapointResponseListWithDefaults() *MetricMemoryDatapointResponseList {
	this := MetricMemoryDatapointResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MetricMemoryDatapointResponseList) GetResults() []MetricMemoryDatapointResponse {
	if o == nil || o.Results == nil {
		var ret []MetricMemoryDatapointResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMemoryDatapointResponseList) GetResultsOk() ([]MetricMemoryDatapointResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MetricMemoryDatapointResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MetricMemoryDatapointResponse and assigns it to the Results field.
func (o *MetricMemoryDatapointResponseList) SetResults(v []MetricMemoryDatapointResponse) {
	o.Results = v
}

func (o MetricMemoryDatapointResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableMetricMemoryDatapointResponseList struct {
	value *MetricMemoryDatapointResponseList
	isSet bool
}

func (v NullableMetricMemoryDatapointResponseList) Get() *MetricMemoryDatapointResponseList {
	return v.value
}

func (v *NullableMetricMemoryDatapointResponseList) Set(val *MetricMemoryDatapointResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricMemoryDatapointResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricMemoryDatapointResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricMemoryDatapointResponseList(val *MetricMemoryDatapointResponseList) *NullableMetricMemoryDatapointResponseList {
	return &NullableMetricMemoryDatapointResponseList{value: val, isSet: true}
}

func (v NullableMetricMemoryDatapointResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricMemoryDatapointResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
