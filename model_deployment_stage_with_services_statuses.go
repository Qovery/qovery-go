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

// checks if the DeploymentStageWithServicesStatuses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentStageWithServicesStatuses{}

// DeploymentStageWithServicesStatuses struct for DeploymentStageWithServicesStatuses
type DeploymentStageWithServicesStatuses struct {
	Applications         []Status `json:"applications,omitempty"`
	Containers           []Status `json:"containers,omitempty"`
	Jobs                 []Status `json:"jobs,omitempty"`
	Databases            []Status `json:"databases,omitempty"`
	Helms                []Status `json:"helms,omitempty"`
	Stage                *Stage   `json:"stage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentStageWithServicesStatuses DeploymentStageWithServicesStatuses

// NewDeploymentStageWithServicesStatuses instantiates a new DeploymentStageWithServicesStatuses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStageWithServicesStatuses() *DeploymentStageWithServicesStatuses {
	this := DeploymentStageWithServicesStatuses{}
	return &this
}

// NewDeploymentStageWithServicesStatusesWithDefaults instantiates a new DeploymentStageWithServicesStatuses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStageWithServicesStatusesWithDefaults() *DeploymentStageWithServicesStatuses {
	this := DeploymentStageWithServicesStatuses{}
	return &this
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *DeploymentStageWithServicesStatuses) GetApplications() []Status {
	if o == nil || IsNil(o.Applications) {
		var ret []Status
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageWithServicesStatuses) GetApplicationsOk() ([]Status, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *DeploymentStageWithServicesStatuses) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []Status and assigns it to the Applications field.
func (o *DeploymentStageWithServicesStatuses) SetApplications(v []Status) {
	o.Applications = v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *DeploymentStageWithServicesStatuses) GetContainers() []Status {
	if o == nil || IsNil(o.Containers) {
		var ret []Status
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageWithServicesStatuses) GetContainersOk() ([]Status, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *DeploymentStageWithServicesStatuses) HasContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []Status and assigns it to the Containers field.
func (o *DeploymentStageWithServicesStatuses) SetContainers(v []Status) {
	o.Containers = v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *DeploymentStageWithServicesStatuses) GetJobs() []Status {
	if o == nil || IsNil(o.Jobs) {
		var ret []Status
		return ret
	}
	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageWithServicesStatuses) GetJobsOk() ([]Status, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *DeploymentStageWithServicesStatuses) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []Status and assigns it to the Jobs field.
func (o *DeploymentStageWithServicesStatuses) SetJobs(v []Status) {
	o.Jobs = v
}

// GetDatabases returns the Databases field value if set, zero value otherwise.
func (o *DeploymentStageWithServicesStatuses) GetDatabases() []Status {
	if o == nil || IsNil(o.Databases) {
		var ret []Status
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageWithServicesStatuses) GetDatabasesOk() ([]Status, bool) {
	if o == nil || IsNil(o.Databases) {
		return nil, false
	}
	return o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *DeploymentStageWithServicesStatuses) HasDatabases() bool {
	if o != nil && !IsNil(o.Databases) {
		return true
	}

	return false
}

// SetDatabases gets a reference to the given []Status and assigns it to the Databases field.
func (o *DeploymentStageWithServicesStatuses) SetDatabases(v []Status) {
	o.Databases = v
}

// GetHelms returns the Helms field value if set, zero value otherwise.
func (o *DeploymentStageWithServicesStatuses) GetHelms() []Status {
	if o == nil || IsNil(o.Helms) {
		var ret []Status
		return ret
	}
	return o.Helms
}

// GetHelmsOk returns a tuple with the Helms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageWithServicesStatuses) GetHelmsOk() ([]Status, bool) {
	if o == nil || IsNil(o.Helms) {
		return nil, false
	}
	return o.Helms, true
}

// HasHelms returns a boolean if a field has been set.
func (o *DeploymentStageWithServicesStatuses) HasHelms() bool {
	if o != nil && !IsNil(o.Helms) {
		return true
	}

	return false
}

// SetHelms gets a reference to the given []Status and assigns it to the Helms field.
func (o *DeploymentStageWithServicesStatuses) SetHelms(v []Status) {
	o.Helms = v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *DeploymentStageWithServicesStatuses) GetStage() Stage {
	if o == nil || IsNil(o.Stage) {
		var ret Stage
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageWithServicesStatuses) GetStageOk() (*Stage, bool) {
	if o == nil || IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *DeploymentStageWithServicesStatuses) HasStage() bool {
	if o != nil && !IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given Stage and assigns it to the Stage field.
func (o *DeploymentStageWithServicesStatuses) SetStage(v Stage) {
	o.Stage = &v
}

func (o DeploymentStageWithServicesStatuses) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentStageWithServicesStatuses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	if !IsNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	if !IsNil(o.Jobs) {
		toSerialize["jobs"] = o.Jobs
	}
	if !IsNil(o.Databases) {
		toSerialize["databases"] = o.Databases
	}
	if !IsNil(o.Helms) {
		toSerialize["helms"] = o.Helms
	}
	if !IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploymentStageWithServicesStatuses) UnmarshalJSON(data []byte) (err error) {
	varDeploymentStageWithServicesStatuses := _DeploymentStageWithServicesStatuses{}

	err = json.Unmarshal(data, &varDeploymentStageWithServicesStatuses)

	if err != nil {
		return err
	}

	*o = DeploymentStageWithServicesStatuses(varDeploymentStageWithServicesStatuses)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "applications")
		delete(additionalProperties, "containers")
		delete(additionalProperties, "jobs")
		delete(additionalProperties, "databases")
		delete(additionalProperties, "helms")
		delete(additionalProperties, "stage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentStageWithServicesStatuses struct {
	value *DeploymentStageWithServicesStatuses
	isSet bool
}

func (v NullableDeploymentStageWithServicesStatuses) Get() *DeploymentStageWithServicesStatuses {
	return v.value
}

func (v *NullableDeploymentStageWithServicesStatuses) Set(val *DeploymentStageWithServicesStatuses) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStageWithServicesStatuses) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStageWithServicesStatuses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStageWithServicesStatuses(val *DeploymentStageWithServicesStatuses) *NullableDeploymentStageWithServicesStatuses {
	return &NullableDeploymentStageWithServicesStatuses{value: val, isSet: true}
}

func (v NullableDeploymentStageWithServicesStatuses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStageWithServicesStatuses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
