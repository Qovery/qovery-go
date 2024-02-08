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
	"time"
)

// checks if the DeploymentRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentRuleRequest{}

// DeploymentRuleRequest struct for DeploymentRuleRequest
type DeploymentRuleRequest struct {
	// name is case insensitive
	Name        string              `json:"name"`
	Description *string             `json:"description,omitempty"`
	Mode        EnvironmentModeEnum `json:"mode"`
	Cluster     string              `json:"cluster"`
	AutoStop    bool                `json:"auto_stop"`
	// specify value only if auto_stop = false
	Timezone *string `json:"timezone,omitempty"`
	// specify value only if auto_stop = false
	StartTime NullableTime `json:"start_time,omitempty"`
	// specify value only if auto_stop = false
	StopTime NullableTime `json:"stop_time,omitempty"`
	// specify value only if auto_stop = false
	Weekdays             []WeekdayEnum `json:"weekdays,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentRuleRequest DeploymentRuleRequest

// NewDeploymentRuleRequest instantiates a new DeploymentRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentRuleRequest(name string, mode EnvironmentModeEnum, cluster string, autoStop bool) *DeploymentRuleRequest {
	this := DeploymentRuleRequest{}
	this.Name = name
	this.Mode = mode
	this.Cluster = cluster
	this.AutoStop = autoStop
	var timezone string = "Europe/London"
	this.Timezone = &timezone
	return &this
}

// NewDeploymentRuleRequestWithDefaults instantiates a new DeploymentRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentRuleRequestWithDefaults() *DeploymentRuleRequest {
	this := DeploymentRuleRequest{}
	var autoStop bool = false
	this.AutoStop = autoStop
	var timezone string = "Europe/London"
	this.Timezone = &timezone
	return &this
}

// GetName returns the Name field value
func (o *DeploymentRuleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeploymentRuleRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeploymentRuleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeploymentRuleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeploymentRuleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value
func (o *DeploymentRuleRequest) GetMode() EnvironmentModeEnum {
	if o == nil {
		var ret EnvironmentModeEnum
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleRequest) GetModeOk() (*EnvironmentModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *DeploymentRuleRequest) SetMode(v EnvironmentModeEnum) {
	o.Mode = v
}

// GetCluster returns the Cluster field value
func (o *DeploymentRuleRequest) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleRequest) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *DeploymentRuleRequest) SetCluster(v string) {
	o.Cluster = v
}

// GetAutoStop returns the AutoStop field value
func (o *DeploymentRuleRequest) GetAutoStop() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoStop
}

// GetAutoStopOk returns a tuple with the AutoStop field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleRequest) GetAutoStopOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoStop, true
}

// SetAutoStop sets field value
func (o *DeploymentRuleRequest) SetAutoStop(v bool) {
	o.AutoStop = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *DeploymentRuleRequest) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleRequest) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *DeploymentRuleRequest) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *DeploymentRuleRequest) SetTimezone(v string) {
	o.Timezone = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentRuleRequest) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentRuleRequest) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *DeploymentRuleRequest) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableTime and assigns it to the StartTime field.
func (o *DeploymentRuleRequest) SetStartTime(v time.Time) {
	o.StartTime.Set(&v)
}

// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *DeploymentRuleRequest) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *DeploymentRuleRequest) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetStopTime returns the StopTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentRuleRequest) GetStopTime() time.Time {
	if o == nil || IsNil(o.StopTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StopTime.Get()
}

// GetStopTimeOk returns a tuple with the StopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentRuleRequest) GetStopTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StopTime.Get(), o.StopTime.IsSet()
}

// HasStopTime returns a boolean if a field has been set.
func (o *DeploymentRuleRequest) HasStopTime() bool {
	if o != nil && o.StopTime.IsSet() {
		return true
	}

	return false
}

// SetStopTime gets a reference to the given NullableTime and assigns it to the StopTime field.
func (o *DeploymentRuleRequest) SetStopTime(v time.Time) {
	o.StopTime.Set(&v)
}

// SetStopTimeNil sets the value for StopTime to be an explicit nil
func (o *DeploymentRuleRequest) SetStopTimeNil() {
	o.StopTime.Set(nil)
}

// UnsetStopTime ensures that no value is present for StopTime, not even an explicit nil
func (o *DeploymentRuleRequest) UnsetStopTime() {
	o.StopTime.Unset()
}

// GetWeekdays returns the Weekdays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentRuleRequest) GetWeekdays() []WeekdayEnum {
	if o == nil {
		var ret []WeekdayEnum
		return ret
	}
	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentRuleRequest) GetWeekdaysOk() ([]WeekdayEnum, bool) {
	if o == nil || IsNil(o.Weekdays) {
		return nil, false
	}
	return o.Weekdays, true
}

// HasWeekdays returns a boolean if a field has been set.
func (o *DeploymentRuleRequest) HasWeekdays() bool {
	if o != nil && IsNil(o.Weekdays) {
		return true
	}

	return false
}

// SetWeekdays gets a reference to the given []WeekdayEnum and assigns it to the Weekdays field.
func (o *DeploymentRuleRequest) SetWeekdays(v []WeekdayEnum) {
	o.Weekdays = v
}

func (o DeploymentRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["mode"] = o.Mode
	toSerialize["cluster"] = o.Cluster
	toSerialize["auto_stop"] = o.AutoStop
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if o.StartTime.IsSet() {
		toSerialize["start_time"] = o.StartTime.Get()
	}
	if o.StopTime.IsSet() {
		toSerialize["stop_time"] = o.StopTime.Get()
	}
	if o.Weekdays != nil {
		toSerialize["weekdays"] = o.Weekdays
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploymentRuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"mode",
		"cluster",
		"auto_stop",
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

	varDeploymentRuleRequest := _DeploymentRuleRequest{}

	err = json.Unmarshal(data, &varDeploymentRuleRequest)

	if err != nil {
		return err
	}

	*o = DeploymentRuleRequest(varDeploymentRuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "cluster")
		delete(additionalProperties, "auto_stop")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "stop_time")
		delete(additionalProperties, "weekdays")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentRuleRequest struct {
	value *DeploymentRuleRequest
	isSet bool
}

func (v NullableDeploymentRuleRequest) Get() *DeploymentRuleRequest {
	return v.value
}

func (v *NullableDeploymentRuleRequest) Set(val *DeploymentRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentRuleRequest(val *DeploymentRuleRequest) *NullableDeploymentRuleRequest {
	return &NullableDeploymentRuleRequest{value: val, isSet: true}
}

func (v NullableDeploymentRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
