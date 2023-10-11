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

// checks if the JobRequestAllOfSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobRequestAllOfSchedule{}

// JobRequestAllOfSchedule If you want to define a Cron job, only the `cronjob` property must be filled   A Lifecycle job should contain at least one property `on_XXX` among the 3 properties: `on_start`, `on_stop`, `on_delete`
type JobRequestAllOfSchedule struct {
	OnStart  *JobRequestAllOfScheduleOnStart `json:"on_start,omitempty"`
	OnStop   *JobRequestAllOfScheduleOnStart `json:"on_stop,omitempty"`
	OnDelete *JobRequestAllOfScheduleOnStart `json:"on_delete,omitempty"`
	Cronjob  *JobRequestAllOfScheduleCronjob `json:"cronjob,omitempty"`
}

// NewJobRequestAllOfSchedule instantiates a new JobRequestAllOfSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobRequestAllOfSchedule() *JobRequestAllOfSchedule {
	this := JobRequestAllOfSchedule{}
	return &this
}

// NewJobRequestAllOfScheduleWithDefaults instantiates a new JobRequestAllOfSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobRequestAllOfScheduleWithDefaults() *JobRequestAllOfSchedule {
	this := JobRequestAllOfSchedule{}
	return &this
}

// GetOnStart returns the OnStart field value if set, zero value otherwise.
func (o *JobRequestAllOfSchedule) GetOnStart() JobRequestAllOfScheduleOnStart {
	if o == nil || IsNil(o.OnStart) {
		var ret JobRequestAllOfScheduleOnStart
		return ret
	}
	return *o.OnStart
}

// GetOnStartOk returns a tuple with the OnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequestAllOfSchedule) GetOnStartOk() (*JobRequestAllOfScheduleOnStart, bool) {
	if o == nil || IsNil(o.OnStart) {
		return nil, false
	}
	return o.OnStart, true
}

// HasOnStart returns a boolean if a field has been set.
func (o *JobRequestAllOfSchedule) HasOnStart() bool {
	if o != nil && !IsNil(o.OnStart) {
		return true
	}

	return false
}

// SetOnStart gets a reference to the given JobRequestAllOfScheduleOnStart and assigns it to the OnStart field.
func (o *JobRequestAllOfSchedule) SetOnStart(v JobRequestAllOfScheduleOnStart) {
	o.OnStart = &v
}

// GetOnStop returns the OnStop field value if set, zero value otherwise.
func (o *JobRequestAllOfSchedule) GetOnStop() JobRequestAllOfScheduleOnStart {
	if o == nil || IsNil(o.OnStop) {
		var ret JobRequestAllOfScheduleOnStart
		return ret
	}
	return *o.OnStop
}

// GetOnStopOk returns a tuple with the OnStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequestAllOfSchedule) GetOnStopOk() (*JobRequestAllOfScheduleOnStart, bool) {
	if o == nil || IsNil(o.OnStop) {
		return nil, false
	}
	return o.OnStop, true
}

// HasOnStop returns a boolean if a field has been set.
func (o *JobRequestAllOfSchedule) HasOnStop() bool {
	if o != nil && !IsNil(o.OnStop) {
		return true
	}

	return false
}

// SetOnStop gets a reference to the given JobRequestAllOfScheduleOnStart and assigns it to the OnStop field.
func (o *JobRequestAllOfSchedule) SetOnStop(v JobRequestAllOfScheduleOnStart) {
	o.OnStop = &v
}

// GetOnDelete returns the OnDelete field value if set, zero value otherwise.
func (o *JobRequestAllOfSchedule) GetOnDelete() JobRequestAllOfScheduleOnStart {
	if o == nil || IsNil(o.OnDelete) {
		var ret JobRequestAllOfScheduleOnStart
		return ret
	}
	return *o.OnDelete
}

// GetOnDeleteOk returns a tuple with the OnDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequestAllOfSchedule) GetOnDeleteOk() (*JobRequestAllOfScheduleOnStart, bool) {
	if o == nil || IsNil(o.OnDelete) {
		return nil, false
	}
	return o.OnDelete, true
}

// HasOnDelete returns a boolean if a field has been set.
func (o *JobRequestAllOfSchedule) HasOnDelete() bool {
	if o != nil && !IsNil(o.OnDelete) {
		return true
	}

	return false
}

// SetOnDelete gets a reference to the given JobRequestAllOfScheduleOnStart and assigns it to the OnDelete field.
func (o *JobRequestAllOfSchedule) SetOnDelete(v JobRequestAllOfScheduleOnStart) {
	o.OnDelete = &v
}

// GetCronjob returns the Cronjob field value if set, zero value otherwise.
func (o *JobRequestAllOfSchedule) GetCronjob() JobRequestAllOfScheduleCronjob {
	if o == nil || IsNil(o.Cronjob) {
		var ret JobRequestAllOfScheduleCronjob
		return ret
	}
	return *o.Cronjob
}

// GetCronjobOk returns a tuple with the Cronjob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequestAllOfSchedule) GetCronjobOk() (*JobRequestAllOfScheduleCronjob, bool) {
	if o == nil || IsNil(o.Cronjob) {
		return nil, false
	}
	return o.Cronjob, true
}

// HasCronjob returns a boolean if a field has been set.
func (o *JobRequestAllOfSchedule) HasCronjob() bool {
	if o != nil && !IsNil(o.Cronjob) {
		return true
	}

	return false
}

// SetCronjob gets a reference to the given JobRequestAllOfScheduleCronjob and assigns it to the Cronjob field.
func (o *JobRequestAllOfSchedule) SetCronjob(v JobRequestAllOfScheduleCronjob) {
	o.Cronjob = &v
}

func (o JobRequestAllOfSchedule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobRequestAllOfSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OnStart) {
		toSerialize["on_start"] = o.OnStart
	}
	if !IsNil(o.OnStop) {
		toSerialize["on_stop"] = o.OnStop
	}
	if !IsNil(o.OnDelete) {
		toSerialize["on_delete"] = o.OnDelete
	}
	if !IsNil(o.Cronjob) {
		toSerialize["cronjob"] = o.Cronjob
	}
	return toSerialize, nil
}

type NullableJobRequestAllOfSchedule struct {
	value *JobRequestAllOfSchedule
	isSet bool
}

func (v NullableJobRequestAllOfSchedule) Get() *JobRequestAllOfSchedule {
	return v.value
}

func (v *NullableJobRequestAllOfSchedule) Set(val *JobRequestAllOfSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableJobRequestAllOfSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableJobRequestAllOfSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobRequestAllOfSchedule(val *JobRequestAllOfSchedule) *NullableJobRequestAllOfSchedule {
	return &NullableJobRequestAllOfSchedule{value: val, isSet: true}
}

func (v NullableJobRequestAllOfSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobRequestAllOfSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
