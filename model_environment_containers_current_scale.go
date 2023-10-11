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

// checks if the EnvironmentContainersCurrentScale type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentContainersCurrentScale{}

// EnvironmentContainersCurrentScale struct for EnvironmentContainersCurrentScale
type EnvironmentContainersCurrentScale struct {
	Container                 *string                    `json:"container,omitempty"`
	Min                       *int32                     `json:"min,omitempty"`
	Max                       *int32                     `json:"max,omitempty"`
	Running                   *int32                     `json:"running,omitempty"`
	RunningInPercent          *float32                   `json:"running_in_percent,omitempty"`
	WarningThresholdInPercent *float32                   `json:"warning_threshold_in_percent,omitempty"`
	AlertThresholdInPercent   *float32                   `json:"alert_threshold_in_percent,omitempty"`
	Status                    *ThresholdMetricStatusEnum `json:"status,omitempty"`
	UpdatedAt                 *time.Time                 `json:"updated_at,omitempty"`
}

// NewEnvironmentContainersCurrentScale instantiates a new EnvironmentContainersCurrentScale object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentContainersCurrentScale() *EnvironmentContainersCurrentScale {
	this := EnvironmentContainersCurrentScale{}
	return &this
}

// NewEnvironmentContainersCurrentScaleWithDefaults instantiates a new EnvironmentContainersCurrentScale object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentContainersCurrentScaleWithDefaults() *EnvironmentContainersCurrentScale {
	this := EnvironmentContainersCurrentScale{}
	return &this
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *EnvironmentContainersCurrentScale) GetContainer() string {
	if o == nil || IsNil(o.Container) {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentContainersCurrentScale) GetContainerOk() (*string, bool) {
	if o == nil || IsNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *EnvironmentContainersCurrentScale) HasContainer() bool {
	if o != nil && !IsNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *EnvironmentContainersCurrentScale) SetContainer(v string) {
	o.Container = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *EnvironmentContainersCurrentScale) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentContainersCurrentScale) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *EnvironmentContainersCurrentScale) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *EnvironmentContainersCurrentScale) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *EnvironmentContainersCurrentScale) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentContainersCurrentScale) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *EnvironmentContainersCurrentScale) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *EnvironmentContainersCurrentScale) SetMax(v int32) {
	o.Max = &v
}

// GetRunning returns the Running field value if set, zero value otherwise.
func (o *EnvironmentContainersCurrentScale) GetRunning() int32 {
	if o == nil || IsNil(o.Running) {
		var ret int32
		return ret
	}
	return *o.Running
}

// GetRunningOk returns a tuple with the Running field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentContainersCurrentScale) GetRunningOk() (*int32, bool) {
	if o == nil || IsNil(o.Running) {
		return nil, false
	}
	return o.Running, true
}

// HasRunning returns a boolean if a field has been set.
func (o *EnvironmentContainersCurrentScale) HasRunning() bool {
	if o != nil && !IsNil(o.Running) {
		return true
	}

	return false
}

// SetRunning gets a reference to the given int32 and assigns it to the Running field.
func (o *EnvironmentContainersCurrentScale) SetRunning(v int32) {
	o.Running = &v
}

// GetRunningInPercent returns the RunningInPercent field value if set, zero value otherwise.
func (o *EnvironmentContainersCurrentScale) GetRunningInPercent() float32 {
	if o == nil || IsNil(o.RunningInPercent) {
		var ret float32
		return ret
	}
	return *o.RunningInPercent
}

// GetRunningInPercentOk returns a tuple with the RunningInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentContainersCurrentScale) GetRunningInPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.RunningInPercent) {
		return nil, false
	}
	return o.RunningInPercent, true
}

// HasRunningInPercent returns a boolean if a field has been set.
func (o *EnvironmentContainersCurrentScale) HasRunningInPercent() bool {
	if o != nil && !IsNil(o.RunningInPercent) {
		return true
	}

	return false
}

// SetRunningInPercent gets a reference to the given float32 and assigns it to the RunningInPercent field.
func (o *EnvironmentContainersCurrentScale) SetRunningInPercent(v float32) {
	o.RunningInPercent = &v
}

// GetWarningThresholdInPercent returns the WarningThresholdInPercent field value if set, zero value otherwise.
func (o *EnvironmentContainersCurrentScale) GetWarningThresholdInPercent() float32 {
	if o == nil || IsNil(o.WarningThresholdInPercent) {
		var ret float32
		return ret
	}
	return *o.WarningThresholdInPercent
}

// GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentContainersCurrentScale) GetWarningThresholdInPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.WarningThresholdInPercent) {
		return nil, false
	}
	return o.WarningThresholdInPercent, true
}

// HasWarningThresholdInPercent returns a boolean if a field has been set.
func (o *EnvironmentContainersCurrentScale) HasWarningThresholdInPercent() bool {
	if o != nil && !IsNil(o.WarningThresholdInPercent) {
		return true
	}

	return false
}

// SetWarningThresholdInPercent gets a reference to the given float32 and assigns it to the WarningThresholdInPercent field.
func (o *EnvironmentContainersCurrentScale) SetWarningThresholdInPercent(v float32) {
	o.WarningThresholdInPercent = &v
}

// GetAlertThresholdInPercent returns the AlertThresholdInPercent field value if set, zero value otherwise.
func (o *EnvironmentContainersCurrentScale) GetAlertThresholdInPercent() float32 {
	if o == nil || IsNil(o.AlertThresholdInPercent) {
		var ret float32
		return ret
	}
	return *o.AlertThresholdInPercent
}

// GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentContainersCurrentScale) GetAlertThresholdInPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.AlertThresholdInPercent) {
		return nil, false
	}
	return o.AlertThresholdInPercent, true
}

// HasAlertThresholdInPercent returns a boolean if a field has been set.
func (o *EnvironmentContainersCurrentScale) HasAlertThresholdInPercent() bool {
	if o != nil && !IsNil(o.AlertThresholdInPercent) {
		return true
	}

	return false
}

// SetAlertThresholdInPercent gets a reference to the given float32 and assigns it to the AlertThresholdInPercent field.
func (o *EnvironmentContainersCurrentScale) SetAlertThresholdInPercent(v float32) {
	o.AlertThresholdInPercent = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EnvironmentContainersCurrentScale) GetStatus() ThresholdMetricStatusEnum {
	if o == nil || IsNil(o.Status) {
		var ret ThresholdMetricStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentContainersCurrentScale) GetStatusOk() (*ThresholdMetricStatusEnum, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EnvironmentContainersCurrentScale) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ThresholdMetricStatusEnum and assigns it to the Status field.
func (o *EnvironmentContainersCurrentScale) SetStatus(v ThresholdMetricStatusEnum) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentContainersCurrentScale) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentContainersCurrentScale) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentContainersCurrentScale) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentContainersCurrentScale) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o EnvironmentContainersCurrentScale) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentContainersCurrentScale) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Container) {
		toSerialize["container"] = o.Container
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Running) {
		toSerialize["running"] = o.Running
	}
	if !IsNil(o.RunningInPercent) {
		toSerialize["running_in_percent"] = o.RunningInPercent
	}
	if !IsNil(o.WarningThresholdInPercent) {
		toSerialize["warning_threshold_in_percent"] = o.WarningThresholdInPercent
	}
	if !IsNil(o.AlertThresholdInPercent) {
		toSerialize["alert_threshold_in_percent"] = o.AlertThresholdInPercent
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableEnvironmentContainersCurrentScale struct {
	value *EnvironmentContainersCurrentScale
	isSet bool
}

func (v NullableEnvironmentContainersCurrentScale) Get() *EnvironmentContainersCurrentScale {
	return v.value
}

func (v *NullableEnvironmentContainersCurrentScale) Set(val *EnvironmentContainersCurrentScale) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentContainersCurrentScale) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentContainersCurrentScale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentContainersCurrentScale(val *EnvironmentContainersCurrentScale) *NullableEnvironmentContainersCurrentScale {
	return &NullableEnvironmentContainersCurrentScale{value: val, isSet: true}
}

func (v NullableEnvironmentContainersCurrentScale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentContainersCurrentScale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
