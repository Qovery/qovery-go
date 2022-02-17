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
	"time"
)

// EnvironmentDeploymentRuleResponse struct for EnvironmentDeploymentRuleResponse
type EnvironmentDeploymentRuleResponse struct {
	AutoDeploy *bool `json:"auto_deploy,omitempty"`
	AutoStop *bool `json:"auto_stop,omitempty"`
	// specify value only if auto_stop = false
	Timezone *string `json:"timezone,omitempty"`
	// specify value only if auto_stop = false
	StartTime NullableTime `json:"start_time,omitempty"`
	// specify value only if auto_stop = false
	StopTime NullableTime `json:"stop_time,omitempty"`
	// specify value only if auto_stop = false
	Weekdays []string `json:"weekdays,omitempty"`
	Id string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewEnvironmentDeploymentRuleResponse instantiates a new EnvironmentDeploymentRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentDeploymentRuleResponse(id string, createdAt time.Time) *EnvironmentDeploymentRuleResponse {
	this := EnvironmentDeploymentRuleResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewEnvironmentDeploymentRuleResponseWithDefaults instantiates a new EnvironmentDeploymentRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentDeploymentRuleResponseWithDefaults() *EnvironmentDeploymentRuleResponse {
	this := EnvironmentDeploymentRuleResponse{}
	var autoDeploy bool = true
	this.AutoDeploy = &autoDeploy
	var autoStop bool = false
	this.AutoStop = &autoStop
	var timezone string = "Europe/London"
	this.Timezone = &timezone
	return &this
}

// GetAutoDeploy returns the AutoDeploy field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleResponse) GetAutoDeploy() bool {
	if o == nil || o.AutoDeploy == nil {
		var ret bool
		return ret
	}
	return *o.AutoDeploy
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleResponse) GetAutoDeployOk() (*bool, bool) {
	if o == nil || o.AutoDeploy == nil {
		return nil, false
	}
	return o.AutoDeploy, true
}

// HasAutoDeploy returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleResponse) HasAutoDeploy() bool {
	if o != nil && o.AutoDeploy != nil {
		return true
	}

	return false
}

// SetAutoDeploy gets a reference to the given bool and assigns it to the AutoDeploy field.
func (o *EnvironmentDeploymentRuleResponse) SetAutoDeploy(v bool) {
	o.AutoDeploy = &v
}

// GetAutoStop returns the AutoStop field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleResponse) GetAutoStop() bool {
	if o == nil || o.AutoStop == nil {
		var ret bool
		return ret
	}
	return *o.AutoStop
}

// GetAutoStopOk returns a tuple with the AutoStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleResponse) GetAutoStopOk() (*bool, bool) {
	if o == nil || o.AutoStop == nil {
		return nil, false
	}
	return o.AutoStop, true
}

// HasAutoStop returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleResponse) HasAutoStop() bool {
	if o != nil && o.AutoStop != nil {
		return true
	}

	return false
}

// SetAutoStop gets a reference to the given bool and assigns it to the AutoStop field.
func (o *EnvironmentDeploymentRuleResponse) SetAutoStop(v bool) {
	o.AutoStop = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleResponse) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleResponse) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleResponse) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *EnvironmentDeploymentRuleResponse) SetTimezone(v string) {
	o.Timezone = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentDeploymentRuleResponse) GetStartTime() time.Time {
	if o == nil || o.StartTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentDeploymentRuleResponse) GetStartTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleResponse) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableTime and assigns it to the StartTime field.
func (o *EnvironmentDeploymentRuleResponse) SetStartTime(v time.Time) {
	o.StartTime.Set(&v)
}
// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *EnvironmentDeploymentRuleResponse) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *EnvironmentDeploymentRuleResponse) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetStopTime returns the StopTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentDeploymentRuleResponse) GetStopTime() time.Time {
	if o == nil || o.StopTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StopTime.Get()
}

// GetStopTimeOk returns a tuple with the StopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentDeploymentRuleResponse) GetStopTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StopTime.Get(), o.StopTime.IsSet()
}

// HasStopTime returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleResponse) HasStopTime() bool {
	if o != nil && o.StopTime.IsSet() {
		return true
	}

	return false
}

// SetStopTime gets a reference to the given NullableTime and assigns it to the StopTime field.
func (o *EnvironmentDeploymentRuleResponse) SetStopTime(v time.Time) {
	o.StopTime.Set(&v)
}
// SetStopTimeNil sets the value for StopTime to be an explicit nil
func (o *EnvironmentDeploymentRuleResponse) SetStopTimeNil() {
	o.StopTime.Set(nil)
}

// UnsetStopTime ensures that no value is present for StopTime, not even an explicit nil
func (o *EnvironmentDeploymentRuleResponse) UnsetStopTime() {
	o.StopTime.Unset()
}

// GetWeekdays returns the Weekdays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentDeploymentRuleResponse) GetWeekdays() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentDeploymentRuleResponse) GetWeekdaysOk() ([]string, bool) {
	if o == nil || o.Weekdays == nil {
		return nil, false
	}
	return o.Weekdays, true
}

// HasWeekdays returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleResponse) HasWeekdays() bool {
	if o != nil && o.Weekdays != nil {
		return true
	}

	return false
}

// SetWeekdays gets a reference to the given []string and assigns it to the Weekdays field.
func (o *EnvironmentDeploymentRuleResponse) SetWeekdays(v []string) {
	o.Weekdays = v
}

// GetId returns the Id field value
func (o *EnvironmentDeploymentRuleResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentDeploymentRuleResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EnvironmentDeploymentRuleResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EnvironmentDeploymentRuleResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentDeploymentRuleResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o EnvironmentDeploymentRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoDeploy != nil {
		toSerialize["auto_deploy"] = o.AutoDeploy
	}
	if o.AutoStop != nil {
		toSerialize["auto_stop"] = o.AutoStop
	}
	if o.Timezone != nil {
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
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentDeploymentRuleResponse struct {
	value *EnvironmentDeploymentRuleResponse
	isSet bool
}

func (v NullableEnvironmentDeploymentRuleResponse) Get() *EnvironmentDeploymentRuleResponse {
	return v.value
}

func (v *NullableEnvironmentDeploymentRuleResponse) Set(val *EnvironmentDeploymentRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentDeploymentRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentDeploymentRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentDeploymentRuleResponse(val *EnvironmentDeploymentRuleResponse) *NullableEnvironmentDeploymentRuleResponse {
	return &NullableEnvironmentDeploymentRuleResponse{value: val, isSet: true}
}

func (v NullableEnvironmentDeploymentRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentDeploymentRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


