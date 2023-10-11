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

// checks if the JobResponseAllOfScheduleCronjob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobResponseAllOfScheduleCronjob{}

// JobResponseAllOfScheduleCronjob struct for JobResponseAllOfScheduleCronjob
type JobResponseAllOfScheduleCronjob struct {
	Arguments []string `json:"arguments,omitempty"`
	// optional entrypoint when launching container
	Entrypoint *string `json:"entrypoint,omitempty"`
	// Can only be set if the event is CRON.   Represent the cron format for the job schedule without seconds.   For example: `* * * * *` represent the cron to launch the job every minute.   See https://crontab.guru/ to WISIWIG interface.   Timezone is UT
	ScheduledAt string `json:"scheduled_at"`
}

// NewJobResponseAllOfScheduleCronjob instantiates a new JobResponseAllOfScheduleCronjob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobResponseAllOfScheduleCronjob(scheduledAt string) *JobResponseAllOfScheduleCronjob {
	this := JobResponseAllOfScheduleCronjob{}
	this.ScheduledAt = scheduledAt
	return &this
}

// NewJobResponseAllOfScheduleCronjobWithDefaults instantiates a new JobResponseAllOfScheduleCronjob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobResponseAllOfScheduleCronjobWithDefaults() *JobResponseAllOfScheduleCronjob {
	this := JobResponseAllOfScheduleCronjob{}
	return &this
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *JobResponseAllOfScheduleCronjob) GetArguments() []string {
	if o == nil || IsNil(o.Arguments) {
		var ret []string
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResponseAllOfScheduleCronjob) GetArgumentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Arguments) {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *JobResponseAllOfScheduleCronjob) HasArguments() bool {
	if o != nil && !IsNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []string and assigns it to the Arguments field.
func (o *JobResponseAllOfScheduleCronjob) SetArguments(v []string) {
	o.Arguments = v
}

// GetEntrypoint returns the Entrypoint field value if set, zero value otherwise.
func (o *JobResponseAllOfScheduleCronjob) GetEntrypoint() string {
	if o == nil || IsNil(o.Entrypoint) {
		var ret string
		return ret
	}
	return *o.Entrypoint
}

// GetEntrypointOk returns a tuple with the Entrypoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResponseAllOfScheduleCronjob) GetEntrypointOk() (*string, bool) {
	if o == nil || IsNil(o.Entrypoint) {
		return nil, false
	}
	return o.Entrypoint, true
}

// HasEntrypoint returns a boolean if a field has been set.
func (o *JobResponseAllOfScheduleCronjob) HasEntrypoint() bool {
	if o != nil && !IsNil(o.Entrypoint) {
		return true
	}

	return false
}

// SetEntrypoint gets a reference to the given string and assigns it to the Entrypoint field.
func (o *JobResponseAllOfScheduleCronjob) SetEntrypoint(v string) {
	o.Entrypoint = &v
}

// GetScheduledAt returns the ScheduledAt field value
func (o *JobResponseAllOfScheduleCronjob) GetScheduledAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduledAt
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value
// and a boolean to check if the value has been set.
func (o *JobResponseAllOfScheduleCronjob) GetScheduledAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledAt, true
}

// SetScheduledAt sets field value
func (o *JobResponseAllOfScheduleCronjob) SetScheduledAt(v string) {
	o.ScheduledAt = v
}

func (o JobResponseAllOfScheduleCronjob) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobResponseAllOfScheduleCronjob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}
	if !IsNil(o.Entrypoint) {
		toSerialize["entrypoint"] = o.Entrypoint
	}
	toSerialize["scheduled_at"] = o.ScheduledAt
	return toSerialize, nil
}

type NullableJobResponseAllOfScheduleCronjob struct {
	value *JobResponseAllOfScheduleCronjob
	isSet bool
}

func (v NullableJobResponseAllOfScheduleCronjob) Get() *JobResponseAllOfScheduleCronjob {
	return v.value
}

func (v *NullableJobResponseAllOfScheduleCronjob) Set(val *JobResponseAllOfScheduleCronjob) {
	v.value = val
	v.isSet = true
}

func (v NullableJobResponseAllOfScheduleCronjob) IsSet() bool {
	return v.isSet
}

func (v *NullableJobResponseAllOfScheduleCronjob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobResponseAllOfScheduleCronjob(val *JobResponseAllOfScheduleCronjob) *NullableJobResponseAllOfScheduleCronjob {
	return &NullableJobResponseAllOfScheduleCronjob{value: val, isSet: true}
}

func (v NullableJobResponseAllOfScheduleCronjob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobResponseAllOfScheduleCronjob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
