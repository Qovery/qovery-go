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
	"time"
)

// MetricStorageDatapointResponse struct for MetricStorageDatapointResponse
type MetricStorageDatapointResponse struct {
	CreatedAt         time.Time `json:"created_at"`
	RequestedInGb     *int32    `json:"requested_in_gb,omitempty"`
	ConsumedInGb      *float32  `json:"consumed_in_gb,omitempty"`
	ConsumedInPercent float32   `json:"consumed_in_percent"`
}

// NewMetricStorageDatapointResponse instantiates a new MetricStorageDatapointResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricStorageDatapointResponse(createdAt time.Time, consumedInPercent float32) *MetricStorageDatapointResponse {
	this := MetricStorageDatapointResponse{}
	this.CreatedAt = createdAt
	this.ConsumedInPercent = consumedInPercent
	return &this
}

// NewMetricStorageDatapointResponseWithDefaults instantiates a new MetricStorageDatapointResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricStorageDatapointResponseWithDefaults() *MetricStorageDatapointResponse {
	this := MetricStorageDatapointResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *MetricStorageDatapointResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MetricStorageDatapointResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MetricStorageDatapointResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetRequestedInGb returns the RequestedInGb field value if set, zero value otherwise.
func (o *MetricStorageDatapointResponse) GetRequestedInGb() int32 {
	if o == nil || o.RequestedInGb == nil {
		var ret int32
		return ret
	}
	return *o.RequestedInGb
}

// GetRequestedInGbOk returns a tuple with the RequestedInGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricStorageDatapointResponse) GetRequestedInGbOk() (*int32, bool) {
	if o == nil || o.RequestedInGb == nil {
		return nil, false
	}
	return o.RequestedInGb, true
}

// HasRequestedInGb returns a boolean if a field has been set.
func (o *MetricStorageDatapointResponse) HasRequestedInGb() bool {
	if o != nil && o.RequestedInGb != nil {
		return true
	}

	return false
}

// SetRequestedInGb gets a reference to the given int32 and assigns it to the RequestedInGb field.
func (o *MetricStorageDatapointResponse) SetRequestedInGb(v int32) {
	o.RequestedInGb = &v
}

// GetConsumedInGb returns the ConsumedInGb field value if set, zero value otherwise.
func (o *MetricStorageDatapointResponse) GetConsumedInGb() float32 {
	if o == nil || o.ConsumedInGb == nil {
		var ret float32
		return ret
	}
	return *o.ConsumedInGb
}

// GetConsumedInGbOk returns a tuple with the ConsumedInGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricStorageDatapointResponse) GetConsumedInGbOk() (*float32, bool) {
	if o == nil || o.ConsumedInGb == nil {
		return nil, false
	}
	return o.ConsumedInGb, true
}

// HasConsumedInGb returns a boolean if a field has been set.
func (o *MetricStorageDatapointResponse) HasConsumedInGb() bool {
	if o != nil && o.ConsumedInGb != nil {
		return true
	}

	return false
}

// SetConsumedInGb gets a reference to the given float32 and assigns it to the ConsumedInGb field.
func (o *MetricStorageDatapointResponse) SetConsumedInGb(v float32) {
	o.ConsumedInGb = &v
}

// GetConsumedInPercent returns the ConsumedInPercent field value
func (o *MetricStorageDatapointResponse) GetConsumedInPercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ConsumedInPercent
}

// GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field value
// and a boolean to check if the value has been set.
func (o *MetricStorageDatapointResponse) GetConsumedInPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumedInPercent, true
}

// SetConsumedInPercent sets field value
func (o *MetricStorageDatapointResponse) SetConsumedInPercent(v float32) {
	o.ConsumedInPercent = v
}

func (o MetricStorageDatapointResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.RequestedInGb != nil {
		toSerialize["requested_in_gb"] = o.RequestedInGb
	}
	if o.ConsumedInGb != nil {
		toSerialize["consumed_in_gb"] = o.ConsumedInGb
	}
	if true {
		toSerialize["consumed_in_percent"] = o.ConsumedInPercent
	}
	return json.Marshal(toSerialize)
}

type NullableMetricStorageDatapointResponse struct {
	value *MetricStorageDatapointResponse
	isSet bool
}

func (v NullableMetricStorageDatapointResponse) Get() *MetricStorageDatapointResponse {
	return v.value
}

func (v *NullableMetricStorageDatapointResponse) Set(val *MetricStorageDatapointResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricStorageDatapointResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricStorageDatapointResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricStorageDatapointResponse(val *MetricStorageDatapointResponse) *NullableMetricStorageDatapointResponse {
	return &NullableMetricStorageDatapointResponse{value: val, isSet: true}
}

func (v NullableMetricStorageDatapointResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricStorageDatapointResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
