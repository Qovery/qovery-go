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

// MetricStorageResponseList struct for MetricStorageResponseList
type MetricStorageResponseList struct {
	Results []MetricStorageResponse `json:"results,omitempty"`
}

// NewMetricStorageResponseList instantiates a new MetricStorageResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricStorageResponseList() *MetricStorageResponseList {
	this := MetricStorageResponseList{}
	return &this
}

// NewMetricStorageResponseListWithDefaults instantiates a new MetricStorageResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricStorageResponseListWithDefaults() *MetricStorageResponseList {
	this := MetricStorageResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MetricStorageResponseList) GetResults() []MetricStorageResponse {
	if o == nil || o.Results == nil {
		var ret []MetricStorageResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricStorageResponseList) GetResultsOk() ([]MetricStorageResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MetricStorageResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MetricStorageResponse and assigns it to the Results field.
func (o *MetricStorageResponseList) SetResults(v []MetricStorageResponse) {
	o.Results = v
}

func (o MetricStorageResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableMetricStorageResponseList struct {
	value *MetricStorageResponseList
	isSet bool
}

func (v NullableMetricStorageResponseList) Get() *MetricStorageResponseList {
	return v.value
}

func (v *NullableMetricStorageResponseList) Set(val *MetricStorageResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricStorageResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricStorageResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricStorageResponseList(val *MetricStorageResponseList) *NullableMetricStorageResponseList {
	return &NullableMetricStorageResponseList{value: val, isSet: true}
}

func (v NullableMetricStorageResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricStorageResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
