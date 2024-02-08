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

// checks if the ServiceStepMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceStepMetric{}

// ServiceStepMetric struct for ServiceStepMetric
type ServiceStepMetric struct {
	StepName *ServiceStepMetricNameEnum `json:"step_name,omitempty"`
	Status   *StepMetricStatusEnum      `json:"status,omitempty"`
	// The duration of the step in seconds.
	DurationSec          *int32 `json:"duration_sec,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceStepMetric ServiceStepMetric

// NewServiceStepMetric instantiates a new ServiceStepMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStepMetric() *ServiceStepMetric {
	this := ServiceStepMetric{}
	return &this
}

// NewServiceStepMetricWithDefaults instantiates a new ServiceStepMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStepMetricWithDefaults() *ServiceStepMetric {
	this := ServiceStepMetric{}
	return &this
}

// GetStepName returns the StepName field value if set, zero value otherwise.
func (o *ServiceStepMetric) GetStepName() ServiceStepMetricNameEnum {
	if o == nil || IsNil(o.StepName) {
		var ret ServiceStepMetricNameEnum
		return ret
	}
	return *o.StepName
}

// GetStepNameOk returns a tuple with the StepName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStepMetric) GetStepNameOk() (*ServiceStepMetricNameEnum, bool) {
	if o == nil || IsNil(o.StepName) {
		return nil, false
	}
	return o.StepName, true
}

// HasStepName returns a boolean if a field has been set.
func (o *ServiceStepMetric) HasStepName() bool {
	if o != nil && !IsNil(o.StepName) {
		return true
	}

	return false
}

// SetStepName gets a reference to the given ServiceStepMetricNameEnum and assigns it to the StepName field.
func (o *ServiceStepMetric) SetStepName(v ServiceStepMetricNameEnum) {
	o.StepName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServiceStepMetric) GetStatus() StepMetricStatusEnum {
	if o == nil || IsNil(o.Status) {
		var ret StepMetricStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStepMetric) GetStatusOk() (*StepMetricStatusEnum, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServiceStepMetric) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StepMetricStatusEnum and assigns it to the Status field.
func (o *ServiceStepMetric) SetStatus(v StepMetricStatusEnum) {
	o.Status = &v
}

// GetDurationSec returns the DurationSec field value if set, zero value otherwise.
func (o *ServiceStepMetric) GetDurationSec() int32 {
	if o == nil || IsNil(o.DurationSec) {
		var ret int32
		return ret
	}
	return *o.DurationSec
}

// GetDurationSecOk returns a tuple with the DurationSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStepMetric) GetDurationSecOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationSec) {
		return nil, false
	}
	return o.DurationSec, true
}

// HasDurationSec returns a boolean if a field has been set.
func (o *ServiceStepMetric) HasDurationSec() bool {
	if o != nil && !IsNil(o.DurationSec) {
		return true
	}

	return false
}

// SetDurationSec gets a reference to the given int32 and assigns it to the DurationSec field.
func (o *ServiceStepMetric) SetDurationSec(v int32) {
	o.DurationSec = &v
}

func (o ServiceStepMetric) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceStepMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StepName) {
		toSerialize["step_name"] = o.StepName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.DurationSec) {
		toSerialize["duration_sec"] = o.DurationSec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceStepMetric) UnmarshalJSON(data []byte) (err error) {
	varServiceStepMetric := _ServiceStepMetric{}

	err = json.Unmarshal(data, &varServiceStepMetric)

	if err != nil {
		return err
	}

	*o = ServiceStepMetric(varServiceStepMetric)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "step_name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "duration_sec")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceStepMetric struct {
	value *ServiceStepMetric
	isSet bool
}

func (v NullableServiceStepMetric) Get() *ServiceStepMetric {
	return v.value
}

func (v *NullableServiceStepMetric) Set(val *ServiceStepMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStepMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStepMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStepMetric(val *ServiceStepMetric) *NullableServiceStepMetric {
	return &NullableServiceStepMetric{value: val, isSet: true}
}

func (v NullableServiceStepMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStepMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
