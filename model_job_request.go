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

// JobRequest struct for JobRequest
type JobRequest struct {
	// name is case insensitive
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *int32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *int32 `json:"memory,omitempty"`
	// Maximum number of restart allowed before the job is considered as failed 0 means that no restart/crash of the job is allowed
	MaxNbRestart *int32 `json:"max_nb_restart,omitempty"`
	// Maximum number of seconds allowed for the job to run before killing it and mark it as failed
	MaxDurationSeconds *int32 `json:"max_duration_seconds,omitempty"`
	// Indicates if the 'environment preview option' is enabled for this container.   If enabled, a preview environment will be automatically cloned when `/preview` endpoint is called.   If not specified, it takes the value of the `auto_preview` property from the associated environment.
	AutoPreview *bool `json:"auto_preview,omitempty"`
	// Port where to run readiness and liveliness probes checks. The port will not be exposed externally
	Port         NullableInt32            `json:"port,omitempty"`
	Source       *JobRequestAllOfSource   `json:"source,omitempty"`
	Healthchecks *Healthcheck             `json:"healthchecks,omitempty"`
	Schedule     *JobRequestAllOfSchedule `json:"schedule,omitempty"`
}

// NewJobRequest instantiates a new JobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobRequest(name string) *JobRequest {
	this := JobRequest{}
	this.Name = name
	var cpu int32 = 500
	this.Cpu = &cpu
	var memory int32 = 512
	this.Memory = &memory
	var maxNbRestart int32 = 0
	this.MaxNbRestart = &maxNbRestart
	return &this
}

// NewJobRequestWithDefaults instantiates a new JobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobRequestWithDefaults() *JobRequest {
	this := JobRequest{}
	var cpu int32 = 500
	this.Cpu = &cpu
	var memory int32 = 512
	this.Memory = &memory
	var maxNbRestart int32 = 0
	this.MaxNbRestart = &maxNbRestart
	return &this
}

// GetName returns the Name field value
func (o *JobRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JobRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JobRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JobRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *JobRequest) GetCpu() int32 {
	if o == nil || o.Cpu == nil {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequest) GetCpuOk() (*int32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *JobRequest) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *JobRequest) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *JobRequest) GetMemory() int32 {
	if o == nil || o.Memory == nil {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequest) GetMemoryOk() (*int32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *JobRequest) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *JobRequest) SetMemory(v int32) {
	o.Memory = &v
}

// GetMaxNbRestart returns the MaxNbRestart field value if set, zero value otherwise.
func (o *JobRequest) GetMaxNbRestart() int32 {
	if o == nil || o.MaxNbRestart == nil {
		var ret int32
		return ret
	}
	return *o.MaxNbRestart
}

// GetMaxNbRestartOk returns a tuple with the MaxNbRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequest) GetMaxNbRestartOk() (*int32, bool) {
	if o == nil || o.MaxNbRestart == nil {
		return nil, false
	}
	return o.MaxNbRestart, true
}

// HasMaxNbRestart returns a boolean if a field has been set.
func (o *JobRequest) HasMaxNbRestart() bool {
	if o != nil && o.MaxNbRestart != nil {
		return true
	}

	return false
}

// SetMaxNbRestart gets a reference to the given int32 and assigns it to the MaxNbRestart field.
func (o *JobRequest) SetMaxNbRestart(v int32) {
	o.MaxNbRestart = &v
}

// GetMaxDurationSeconds returns the MaxDurationSeconds field value if set, zero value otherwise.
func (o *JobRequest) GetMaxDurationSeconds() int32 {
	if o == nil || o.MaxDurationSeconds == nil {
		var ret int32
		return ret
	}
	return *o.MaxDurationSeconds
}

// GetMaxDurationSecondsOk returns a tuple with the MaxDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequest) GetMaxDurationSecondsOk() (*int32, bool) {
	if o == nil || o.MaxDurationSeconds == nil {
		return nil, false
	}
	return o.MaxDurationSeconds, true
}

// HasMaxDurationSeconds returns a boolean if a field has been set.
func (o *JobRequest) HasMaxDurationSeconds() bool {
	if o != nil && o.MaxDurationSeconds != nil {
		return true
	}

	return false
}

// SetMaxDurationSeconds gets a reference to the given int32 and assigns it to the MaxDurationSeconds field.
func (o *JobRequest) SetMaxDurationSeconds(v int32) {
	o.MaxDurationSeconds = &v
}

// GetAutoPreview returns the AutoPreview field value if set, zero value otherwise.
func (o *JobRequest) GetAutoPreview() bool {
	if o == nil || o.AutoPreview == nil {
		var ret bool
		return ret
	}
	return *o.AutoPreview
}

// GetAutoPreviewOk returns a tuple with the AutoPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequest) GetAutoPreviewOk() (*bool, bool) {
	if o == nil || o.AutoPreview == nil {
		return nil, false
	}
	return o.AutoPreview, true
}

// HasAutoPreview returns a boolean if a field has been set.
func (o *JobRequest) HasAutoPreview() bool {
	if o != nil && o.AutoPreview != nil {
		return true
	}

	return false
}

// SetAutoPreview gets a reference to the given bool and assigns it to the AutoPreview field.
func (o *JobRequest) SetAutoPreview(v bool) {
	o.AutoPreview = &v
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobRequest) GetPort() int32 {
	if o == nil || o.Port.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobRequest) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *JobRequest) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *JobRequest) SetPort(v int32) {
	o.Port.Set(&v)
}

// SetPortNil sets the value for Port to be an explicit nil
func (o *JobRequest) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *JobRequest) UnsetPort() {
	o.Port.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *JobRequest) GetSource() JobRequestAllOfSource {
	if o == nil || o.Source == nil {
		var ret JobRequestAllOfSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequest) GetSourceOk() (*JobRequestAllOfSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *JobRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given JobRequestAllOfSource and assigns it to the Source field.
func (o *JobRequest) SetSource(v JobRequestAllOfSource) {
	o.Source = &v
}

// GetHealthchecks returns the Healthchecks field value if set, zero value otherwise.
func (o *JobRequest) GetHealthchecks() Healthcheck {
	if o == nil || o.Healthchecks == nil {
		var ret Healthcheck
		return ret
	}
	return *o.Healthchecks
}

// GetHealthchecksOk returns a tuple with the Healthchecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequest) GetHealthchecksOk() (*Healthcheck, bool) {
	if o == nil || o.Healthchecks == nil {
		return nil, false
	}
	return o.Healthchecks, true
}

// HasHealthchecks returns a boolean if a field has been set.
func (o *JobRequest) HasHealthchecks() bool {
	if o != nil && o.Healthchecks != nil {
		return true
	}

	return false
}

// SetHealthchecks gets a reference to the given Healthcheck and assigns it to the Healthchecks field.
func (o *JobRequest) SetHealthchecks(v Healthcheck) {
	o.Healthchecks = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *JobRequest) GetSchedule() JobRequestAllOfSchedule {
	if o == nil || o.Schedule == nil {
		var ret JobRequestAllOfSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequest) GetScheduleOk() (*JobRequestAllOfSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *JobRequest) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given JobRequestAllOfSchedule and assigns it to the Schedule field.
func (o *JobRequest) SetSchedule(v JobRequestAllOfSchedule) {
	o.Schedule = &v
}

func (o JobRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.MaxNbRestart != nil {
		toSerialize["max_nb_restart"] = o.MaxNbRestart
	}
	if o.MaxDurationSeconds != nil {
		toSerialize["max_duration_seconds"] = o.MaxDurationSeconds
	}
	if o.AutoPreview != nil {
		toSerialize["auto_preview"] = o.AutoPreview
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Healthchecks != nil {
		toSerialize["healthchecks"] = o.Healthchecks
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableJobRequest struct {
	value *JobRequest
	isSet bool
}

func (v NullableJobRequest) Get() *JobRequest {
	return v.value
}

func (v *NullableJobRequest) Set(val *JobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobRequest(val *JobRequest) *NullableJobRequest {
	return &NullableJobRequest{value: val, isSet: true}
}

func (v NullableJobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
