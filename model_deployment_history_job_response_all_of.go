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

// DeploymentHistoryJobResponseAllOf struct for DeploymentHistoryJobResponseAllOf
type DeploymentHistoryJobResponseAllOf struct {
	// name of the job
	Name       *string                                    `json:"name,omitempty"`
	Status     *DeploymentHistoryStatusEnum               `json:"status,omitempty"`
	ImageName  *string                                    `json:"image_name,omitempty"`
	Tag        *string                                    `json:"tag,omitempty"`
	Commit     *Commit                                    `json:"commit,omitempty"`
	Schedule   *DeploymentHistoryJobResponseAllOfSchedule `json:"schedule,omitempty"`
	Arguments  []string                                   `json:"arguments,omitempty"`
	Entrypoint *string                                    `json:"entrypoint,omitempty"`
}

// NewDeploymentHistoryJobResponseAllOf instantiates a new DeploymentHistoryJobResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryJobResponseAllOf() *DeploymentHistoryJobResponseAllOf {
	this := DeploymentHistoryJobResponseAllOf{}
	return &this
}

// NewDeploymentHistoryJobResponseAllOfWithDefaults instantiates a new DeploymentHistoryJobResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryJobResponseAllOfWithDefaults() *DeploymentHistoryJobResponseAllOf {
	this := DeploymentHistoryJobResponseAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponseAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponseAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponseAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentHistoryJobResponseAllOf) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponseAllOf) GetStatus() DeploymentHistoryStatusEnum {
	if o == nil || o.Status == nil {
		var ret DeploymentHistoryStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponseAllOf) GetStatusOk() (*DeploymentHistoryStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponseAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeploymentHistoryStatusEnum and assigns it to the Status field.
func (o *DeploymentHistoryJobResponseAllOf) SetStatus(v DeploymentHistoryStatusEnum) {
	o.Status = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponseAllOf) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponseAllOf) GetImageNameOk() (*string, bool) {
	if o == nil || o.ImageName == nil {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponseAllOf) HasImageName() bool {
	if o != nil && o.ImageName != nil {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *DeploymentHistoryJobResponseAllOf) SetImageName(v string) {
	o.ImageName = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponseAllOf) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponseAllOf) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponseAllOf) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *DeploymentHistoryJobResponseAllOf) SetTag(v string) {
	o.Tag = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponseAllOf) GetCommit() Commit {
	if o == nil || o.Commit == nil {
		var ret Commit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponseAllOf) GetCommitOk() (*Commit, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponseAllOf) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given Commit and assigns it to the Commit field.
func (o *DeploymentHistoryJobResponseAllOf) SetCommit(v Commit) {
	o.Commit = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponseAllOf) GetSchedule() DeploymentHistoryJobResponseAllOfSchedule {
	if o == nil || o.Schedule == nil {
		var ret DeploymentHistoryJobResponseAllOfSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponseAllOf) GetScheduleOk() (*DeploymentHistoryJobResponseAllOfSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponseAllOf) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given DeploymentHistoryJobResponseAllOfSchedule and assigns it to the Schedule field.
func (o *DeploymentHistoryJobResponseAllOf) SetSchedule(v DeploymentHistoryJobResponseAllOfSchedule) {
	o.Schedule = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponseAllOf) GetArguments() []string {
	if o == nil || o.Arguments == nil {
		var ret []string
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponseAllOf) GetArgumentsOk() ([]string, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponseAllOf) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []string and assigns it to the Arguments field.
func (o *DeploymentHistoryJobResponseAllOf) SetArguments(v []string) {
	o.Arguments = v
}

// GetEntrypoint returns the Entrypoint field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponseAllOf) GetEntrypoint() string {
	if o == nil || o.Entrypoint == nil {
		var ret string
		return ret
	}
	return *o.Entrypoint
}

// GetEntrypointOk returns a tuple with the Entrypoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponseAllOf) GetEntrypointOk() (*string, bool) {
	if o == nil || o.Entrypoint == nil {
		return nil, false
	}
	return o.Entrypoint, true
}

// HasEntrypoint returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponseAllOf) HasEntrypoint() bool {
	if o != nil && o.Entrypoint != nil {
		return true
	}

	return false
}

// SetEntrypoint gets a reference to the given string and assigns it to the Entrypoint field.
func (o *DeploymentHistoryJobResponseAllOf) SetEntrypoint(v string) {
	o.Entrypoint = &v
}

func (o DeploymentHistoryJobResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ImageName != nil {
		toSerialize["image_name"] = o.ImageName
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Entrypoint != nil {
		toSerialize["entrypoint"] = o.Entrypoint
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentHistoryJobResponseAllOf struct {
	value *DeploymentHistoryJobResponseAllOf
	isSet bool
}

func (v NullableDeploymentHistoryJobResponseAllOf) Get() *DeploymentHistoryJobResponseAllOf {
	return v.value
}

func (v *NullableDeploymentHistoryJobResponseAllOf) Set(val *DeploymentHistoryJobResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryJobResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryJobResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryJobResponseAllOf(val *DeploymentHistoryJobResponseAllOf) *NullableDeploymentHistoryJobResponseAllOf {
	return &NullableDeploymentHistoryJobResponseAllOf{value: val, isSet: true}
}

func (v NullableDeploymentHistoryJobResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryJobResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
