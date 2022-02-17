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

// ProjectDeploymentRuleRequest struct for ProjectDeploymentRuleRequest
type ProjectDeploymentRuleRequest struct {
	// specify here a wildcard based on environment name that will target new environments after their creation
	EnvironmentTarget string `json:"environment_target"`
	// name is case insensitive
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Mode        string  `json:"mode"`
	Cluster     string  `json:"cluster"`
	AutoDeploy  *bool   `json:"auto_deploy,omitempty"`
	AutoStop    bool    `json:"auto_stop"`
	// specify value only if auto_stop = false
	Timezone *string `json:"timezone,omitempty"`
	// specify value only if auto_stop = false
	StartTime NullableTime `json:"start_time,omitempty"`
	// specify value only if auto_stop = false
	StopTime NullableTime `json:"stop_time,omitempty"`
	// specify value only if auto_stop = false
	Weekdays []string `json:"weekdays,omitempty"`
}

// NewProjectDeploymentRuleRequest instantiates a new ProjectDeploymentRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDeploymentRuleRequest(environmentTarget string, name string, mode string, cluster string, autoStop bool) *ProjectDeploymentRuleRequest {
	this := ProjectDeploymentRuleRequest{}
	this.Name = name
	this.Mode = mode
	this.Cluster = cluster
	var autoDeploy bool = true
	this.AutoDeploy = &autoDeploy
	this.AutoStop = autoStop
	var timezone string = "Europe/London"
	this.Timezone = &timezone
	return &this
}

// NewProjectDeploymentRuleRequestWithDefaults instantiates a new ProjectDeploymentRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDeploymentRuleRequestWithDefaults() *ProjectDeploymentRuleRequest {
	this := ProjectDeploymentRuleRequest{}
	var autoDeploy bool = true
	this.AutoDeploy = &autoDeploy
	var autoStop bool = false
	this.AutoStop = autoStop
	var timezone string = "Europe/London"
	this.Timezone = &timezone
	return &this
}

// GetEnvironmentTarget returns the EnvironmentTarget field value
func (o *ProjectDeploymentRuleRequest) GetEnvironmentTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentTarget
}

// GetEnvironmentTargetOk returns a tuple with the EnvironmentTarget field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleRequest) GetEnvironmentTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentTarget, true
}

// SetEnvironmentTarget sets field value
func (o *ProjectDeploymentRuleRequest) SetEnvironmentTarget(v string) {
	o.EnvironmentTarget = v
}

// GetName returns the Name field value
func (o *ProjectDeploymentRuleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectDeploymentRuleRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectDeploymentRuleRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectDeploymentRuleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value
func (o *ProjectDeploymentRuleRequest) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleRequest) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *ProjectDeploymentRuleRequest) SetMode(v string) {
	o.Mode = v
}

// GetCluster returns the Cluster field value
func (o *ProjectDeploymentRuleRequest) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleRequest) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *ProjectDeploymentRuleRequest) SetCluster(v string) {
	o.Cluster = v
}

// GetAutoDeploy returns the AutoDeploy field value if set, zero value otherwise.
func (o *ProjectDeploymentRuleRequest) GetAutoDeploy() bool {
	if o == nil || o.AutoDeploy == nil {
		var ret bool
		return ret
	}
	return *o.AutoDeploy
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleRequest) GetAutoDeployOk() (*bool, bool) {
	if o == nil || o.AutoDeploy == nil {
		return nil, false
	}
	return o.AutoDeploy, true
}

// HasAutoDeploy returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleRequest) HasAutoDeploy() bool {
	if o != nil && o.AutoDeploy != nil {
		return true
	}

	return false
}

// SetAutoDeploy gets a reference to the given bool and assigns it to the AutoDeploy field.
func (o *ProjectDeploymentRuleRequest) SetAutoDeploy(v bool) {
	o.AutoDeploy = &v
}

// GetAutoStop returns the AutoStop field value
func (o *ProjectDeploymentRuleRequest) GetAutoStop() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoStop
}

// GetAutoStopOk returns a tuple with the AutoStop field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleRequest) GetAutoStopOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoStop, true
}

// SetAutoStop sets field value
func (o *ProjectDeploymentRuleRequest) SetAutoStop(v bool) {
	o.AutoStop = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ProjectDeploymentRuleRequest) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleRequest) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleRequest) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ProjectDeploymentRuleRequest) SetTimezone(v string) {
	o.Timezone = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectDeploymentRuleRequest) GetStartTime() time.Time {
	if o == nil || o.StartTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectDeploymentRuleRequest) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleRequest) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableTime and assigns it to the StartTime field.
func (o *ProjectDeploymentRuleRequest) SetStartTime(v time.Time) {
	o.StartTime.Set(&v)
}

// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *ProjectDeploymentRuleRequest) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *ProjectDeploymentRuleRequest) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetStopTime returns the StopTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectDeploymentRuleRequest) GetStopTime() time.Time {
	if o == nil || o.StopTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StopTime.Get()
}

// GetStopTimeOk returns a tuple with the StopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectDeploymentRuleRequest) GetStopTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StopTime.Get(), o.StopTime.IsSet()
}

// HasStopTime returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleRequest) HasStopTime() bool {
	if o != nil && o.StopTime.IsSet() {
		return true
	}

	return false
}

// SetStopTime gets a reference to the given NullableTime and assigns it to the StopTime field.
func (o *ProjectDeploymentRuleRequest) SetStopTime(v time.Time) {
	o.StopTime.Set(&v)
}

// SetStopTimeNil sets the value for StopTime to be an explicit nil
func (o *ProjectDeploymentRuleRequest) SetStopTimeNil() {
	o.StopTime.Set(nil)
}

// UnsetStopTime ensures that no value is present for StopTime, not even an explicit nil
func (o *ProjectDeploymentRuleRequest) UnsetStopTime() {
	o.StopTime.Unset()
}

// GetWeekdays returns the Weekdays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectDeploymentRuleRequest) GetWeekdays() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectDeploymentRuleRequest) GetWeekdaysOk() (*[]string, bool) {
	if o == nil || o.Weekdays == nil {
		return nil, false
	}
	return &o.Weekdays, true
}

// HasWeekdays returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleRequest) HasWeekdays() bool {
	if o != nil && o.Weekdays != nil {
		return true
	}

	return false
}

// SetWeekdays gets a reference to the given []string and assigns it to the Weekdays field.
func (o *ProjectDeploymentRuleRequest) SetWeekdays(v []string) {
	o.Weekdays = v
}

func (o ProjectDeploymentRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["environment_target"] = o.EnvironmentTarget
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["cluster"] = o.Cluster
	}
	if o.AutoDeploy != nil {
		toSerialize["auto_deploy"] = o.AutoDeploy
	}
	if true {
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
	return json.Marshal(toSerialize)
}

type NullableProjectDeploymentRuleRequest struct {
	value *ProjectDeploymentRuleRequest
	isSet bool
}

func (v NullableProjectDeploymentRuleRequest) Get() *ProjectDeploymentRuleRequest {
	return v.value
}

func (v *NullableProjectDeploymentRuleRequest) Set(val *ProjectDeploymentRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDeploymentRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDeploymentRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDeploymentRuleRequest(val *ProjectDeploymentRuleRequest) *NullableProjectDeploymentRuleRequest {
	return &NullableProjectDeploymentRuleRequest{value: val, isSet: true}
}

func (v NullableProjectDeploymentRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDeploymentRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
