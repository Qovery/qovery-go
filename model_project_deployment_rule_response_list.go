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

// checks if the ProjectDeploymentRuleResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectDeploymentRuleResponseList{}

// ProjectDeploymentRuleResponseList struct for ProjectDeploymentRuleResponseList
type ProjectDeploymentRuleResponseList struct {
	Results              []ProjectDeploymentRule `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectDeploymentRuleResponseList ProjectDeploymentRuleResponseList

// NewProjectDeploymentRuleResponseList instantiates a new ProjectDeploymentRuleResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDeploymentRuleResponseList() *ProjectDeploymentRuleResponseList {
	this := ProjectDeploymentRuleResponseList{}
	return &this
}

// NewProjectDeploymentRuleResponseListWithDefaults instantiates a new ProjectDeploymentRuleResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDeploymentRuleResponseListWithDefaults() *ProjectDeploymentRuleResponseList {
	this := ProjectDeploymentRuleResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ProjectDeploymentRuleResponseList) GetResults() []ProjectDeploymentRule {
	if o == nil || IsNil(o.Results) {
		var ret []ProjectDeploymentRule
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponseList) GetResultsOk() ([]ProjectDeploymentRule, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ProjectDeploymentRule and assigns it to the Results field.
func (o *ProjectDeploymentRuleResponseList) SetResults(v []ProjectDeploymentRule) {
	o.Results = v
}

func (o ProjectDeploymentRuleResponseList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectDeploymentRuleResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectDeploymentRuleResponseList) UnmarshalJSON(data []byte) (err error) {
	varProjectDeploymentRuleResponseList := _ProjectDeploymentRuleResponseList{}

	err = json.Unmarshal(data, &varProjectDeploymentRuleResponseList)

	if err != nil {
		return err
	}

	*o = ProjectDeploymentRuleResponseList(varProjectDeploymentRuleResponseList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectDeploymentRuleResponseList struct {
	value *ProjectDeploymentRuleResponseList
	isSet bool
}

func (v NullableProjectDeploymentRuleResponseList) Get() *ProjectDeploymentRuleResponseList {
	return v.value
}

func (v *NullableProjectDeploymentRuleResponseList) Set(val *ProjectDeploymentRuleResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDeploymentRuleResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDeploymentRuleResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDeploymentRuleResponseList(val *ProjectDeploymentRuleResponseList) *NullableProjectDeploymentRuleResponseList {
	return &NullableProjectDeploymentRuleResponseList{value: val, isSet: true}
}

func (v NullableProjectDeploymentRuleResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDeploymentRuleResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
