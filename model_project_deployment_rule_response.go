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

// ProjectDeploymentRuleResponse struct for ProjectDeploymentRuleResponse
type ProjectDeploymentRuleResponse struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// name is case insensitive
	Name        string              `json:"name"`
	Description NullableString      `json:"description,omitempty"`
	Mode        EnvironmentModeEnum `json:"mode"`
	ClusterId   string              `json:"cluster_id"`
	AutoDeploy  *bool               `json:"auto_deploy,omitempty"`
	AutoStop    *bool               `json:"auto_stop,omitempty"`
	AutoDelete  *bool               `json:"auto_delete,omitempty"`
	Timezone    string              `json:"timezone"`
	StartTime   time.Time           `json:"start_time"`
	StopTime    time.Time           `json:"stop_time"`
	Weekdays    []WeekdayEnum       `json:"weekdays"`
	// wildcard pattern composed of '?' and/or '*' used to target new created environments
	Wildcard string `json:"wildcard"`
	// used to select the first deployment rule to match new created environments
	PriorityIndex *int32 `json:"priority_index,omitempty"`
}

// NewProjectDeploymentRuleResponse instantiates a new ProjectDeploymentRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDeploymentRuleResponse(id string, createdAt time.Time, name string, mode EnvironmentModeEnum, clusterId string, timezone string, startTime time.Time, stopTime time.Time, weekdays []WeekdayEnum, wildcard string) *ProjectDeploymentRuleResponse {
	this := ProjectDeploymentRuleResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Name = name
	this.Mode = mode
	this.ClusterId = clusterId
	var autoDeploy bool = false
	this.AutoDeploy = &autoDeploy
	var autoStop bool = false
	this.AutoStop = &autoStop
	var autoDelete bool = false
	this.AutoDelete = &autoDelete
	this.Timezone = timezone
	this.StartTime = startTime
	this.StopTime = stopTime
	this.Weekdays = weekdays
	this.Wildcard = wildcard
	return &this
}

// NewProjectDeploymentRuleResponseWithDefaults instantiates a new ProjectDeploymentRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDeploymentRuleResponseWithDefaults() *ProjectDeploymentRuleResponse {
	this := ProjectDeploymentRuleResponse{}
	var autoDeploy bool = false
	this.AutoDeploy = &autoDeploy
	var autoStop bool = false
	this.AutoStop = &autoStop
	var autoDelete bool = false
	this.AutoDelete = &autoDelete
	var wildcard string = ""
	this.Wildcard = wildcard
	return &this
}

// GetId returns the Id field value
func (o *ProjectDeploymentRuleResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectDeploymentRuleResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProjectDeploymentRuleResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProjectDeploymentRuleResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProjectDeploymentRuleResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ProjectDeploymentRuleResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value
func (o *ProjectDeploymentRuleResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectDeploymentRuleResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectDeploymentRuleResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectDeploymentRuleResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ProjectDeploymentRuleResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ProjectDeploymentRuleResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ProjectDeploymentRuleResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetMode returns the Mode field value
func (o *ProjectDeploymentRuleResponse) GetMode() EnvironmentModeEnum {
	if o == nil {
		var ret EnvironmentModeEnum
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetModeOk() (*EnvironmentModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *ProjectDeploymentRuleResponse) SetMode(v EnvironmentModeEnum) {
	o.Mode = v
}

// GetClusterId returns the ClusterId field value
func (o *ProjectDeploymentRuleResponse) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *ProjectDeploymentRuleResponse) SetClusterId(v string) {
	o.ClusterId = v
}

// GetAutoDeploy returns the AutoDeploy field value if set, zero value otherwise.
func (o *ProjectDeploymentRuleResponse) GetAutoDeploy() bool {
	if o == nil || o.AutoDeploy == nil {
		var ret bool
		return ret
	}
	return *o.AutoDeploy
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetAutoDeployOk() (*bool, bool) {
	if o == nil || o.AutoDeploy == nil {
		return nil, false
	}
	return o.AutoDeploy, true
}

// HasAutoDeploy returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleResponse) HasAutoDeploy() bool {
	if o != nil && o.AutoDeploy != nil {
		return true
	}

	return false
}

// SetAutoDeploy gets a reference to the given bool and assigns it to the AutoDeploy field.
func (o *ProjectDeploymentRuleResponse) SetAutoDeploy(v bool) {
	o.AutoDeploy = &v
}

// GetAutoStop returns the AutoStop field value if set, zero value otherwise.
func (o *ProjectDeploymentRuleResponse) GetAutoStop() bool {
	if o == nil || o.AutoStop == nil {
		var ret bool
		return ret
	}
	return *o.AutoStop
}

// GetAutoStopOk returns a tuple with the AutoStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetAutoStopOk() (*bool, bool) {
	if o == nil || o.AutoStop == nil {
		return nil, false
	}
	return o.AutoStop, true
}

// HasAutoStop returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleResponse) HasAutoStop() bool {
	if o != nil && o.AutoStop != nil {
		return true
	}

	return false
}

// SetAutoStop gets a reference to the given bool and assigns it to the AutoStop field.
func (o *ProjectDeploymentRuleResponse) SetAutoStop(v bool) {
	o.AutoStop = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *ProjectDeploymentRuleResponse) GetAutoDelete() bool {
	if o == nil || o.AutoDelete == nil {
		var ret bool
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetAutoDeleteOk() (*bool, bool) {
	if o == nil || o.AutoDelete == nil {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleResponse) HasAutoDelete() bool {
	if o != nil && o.AutoDelete != nil {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given bool and assigns it to the AutoDelete field.
func (o *ProjectDeploymentRuleResponse) SetAutoDelete(v bool) {
	o.AutoDelete = &v
}

// GetTimezone returns the Timezone field value
func (o *ProjectDeploymentRuleResponse) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *ProjectDeploymentRuleResponse) SetTimezone(v string) {
	o.Timezone = v
}

// GetStartTime returns the StartTime field value
func (o *ProjectDeploymentRuleResponse) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *ProjectDeploymentRuleResponse) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetStopTime returns the StopTime field value
func (o *ProjectDeploymentRuleResponse) GetStopTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetStopTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopTime, true
}

// SetStopTime sets field value
func (o *ProjectDeploymentRuleResponse) SetStopTime(v time.Time) {
	o.StopTime = v
}

// GetWeekdays returns the Weekdays field value
func (o *ProjectDeploymentRuleResponse) GetWeekdays() []WeekdayEnum {
	if o == nil {
		var ret []WeekdayEnum
		return ret
	}

	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetWeekdaysOk() ([]WeekdayEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weekdays, true
}

// SetWeekdays sets field value
func (o *ProjectDeploymentRuleResponse) SetWeekdays(v []WeekdayEnum) {
	o.Weekdays = v
}

// GetWildcard returns the Wildcard field value
func (o *ProjectDeploymentRuleResponse) GetWildcard() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Wildcard
}

// GetWildcardOk returns a tuple with the Wildcard field value
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetWildcardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wildcard, true
}

// SetWildcard sets field value
func (o *ProjectDeploymentRuleResponse) SetWildcard(v string) {
	o.Wildcard = v
}

// GetPriorityIndex returns the PriorityIndex field value if set, zero value otherwise.
func (o *ProjectDeploymentRuleResponse) GetPriorityIndex() int32 {
	if o == nil || o.PriorityIndex == nil {
		var ret int32
		return ret
	}
	return *o.PriorityIndex
}

// GetPriorityIndexOk returns a tuple with the PriorityIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponse) GetPriorityIndexOk() (*int32, bool) {
	if o == nil || o.PriorityIndex == nil {
		return nil, false
	}
	return o.PriorityIndex, true
}

// HasPriorityIndex returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleResponse) HasPriorityIndex() bool {
	if o != nil && o.PriorityIndex != nil {
		return true
	}

	return false
}

// SetPriorityIndex gets a reference to the given int32 and assigns it to the PriorityIndex field.
func (o *ProjectDeploymentRuleResponse) SetPriorityIndex(v int32) {
	o.PriorityIndex = &v
}

func (o ProjectDeploymentRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.AutoDeploy != nil {
		toSerialize["auto_deploy"] = o.AutoDeploy
	}
	if o.AutoStop != nil {
		toSerialize["auto_stop"] = o.AutoStop
	}
	if o.AutoDelete != nil {
		toSerialize["auto_delete"] = o.AutoDelete
	}
	if true {
		toSerialize["timezone"] = o.Timezone
	}
	if true {
		toSerialize["start_time"] = o.StartTime
	}
	if true {
		toSerialize["stop_time"] = o.StopTime
	}
	if true {
		toSerialize["weekdays"] = o.Weekdays
	}
	if true {
		toSerialize["wildcard"] = o.Wildcard
	}
	if o.PriorityIndex != nil {
		toSerialize["priority_index"] = o.PriorityIndex
	}
	return json.Marshal(toSerialize)
}

type NullableProjectDeploymentRuleResponse struct {
	value *ProjectDeploymentRuleResponse
	isSet bool
}

func (v NullableProjectDeploymentRuleResponse) Get() *ProjectDeploymentRuleResponse {
	return v.value
}

func (v *NullableProjectDeploymentRuleResponse) Set(val *ProjectDeploymentRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDeploymentRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDeploymentRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDeploymentRuleResponse(val *ProjectDeploymentRuleResponse) *NullableProjectDeploymentRuleResponse {
	return &NullableProjectDeploymentRuleResponse{value: val, isSet: true}
}

func (v NullableProjectDeploymentRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDeploymentRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
