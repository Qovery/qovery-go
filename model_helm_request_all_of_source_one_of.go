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

// checks if the HelmRequestAllOfSourceOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmRequestAllOfSourceOneOf{}

// HelmRequestAllOfSourceOneOf struct for HelmRequestAllOfSourceOneOf
type HelmRequestAllOfSourceOneOf struct {
	GitRepository        *HelmGitRepositoryRequest `json:"git_repository,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HelmRequestAllOfSourceOneOf HelmRequestAllOfSourceOneOf

// NewHelmRequestAllOfSourceOneOf instantiates a new HelmRequestAllOfSourceOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmRequestAllOfSourceOneOf() *HelmRequestAllOfSourceOneOf {
	this := HelmRequestAllOfSourceOneOf{}
	return &this
}

// NewHelmRequestAllOfSourceOneOfWithDefaults instantiates a new HelmRequestAllOfSourceOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmRequestAllOfSourceOneOfWithDefaults() *HelmRequestAllOfSourceOneOf {
	this := HelmRequestAllOfSourceOneOf{}
	return &this
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise.
func (o *HelmRequestAllOfSourceOneOf) GetGitRepository() HelmGitRepositoryRequest {
	if o == nil || IsNil(o.GitRepository) {
		var ret HelmGitRepositoryRequest
		return ret
	}
	return *o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRequestAllOfSourceOneOf) GetGitRepositoryOk() (*HelmGitRepositoryRequest, bool) {
	if o == nil || IsNil(o.GitRepository) {
		return nil, false
	}
	return o.GitRepository, true
}

// HasGitRepository returns a boolean if a field has been set.
func (o *HelmRequestAllOfSourceOneOf) HasGitRepository() bool {
	if o != nil && !IsNil(o.GitRepository) {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given HelmGitRepositoryRequest and assigns it to the GitRepository field.
func (o *HelmRequestAllOfSourceOneOf) SetGitRepository(v HelmGitRepositoryRequest) {
	o.GitRepository = &v
}

func (o HelmRequestAllOfSourceOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmRequestAllOfSourceOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GitRepository) {
		toSerialize["git_repository"] = o.GitRepository
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmRequestAllOfSourceOneOf) UnmarshalJSON(data []byte) (err error) {
	varHelmRequestAllOfSourceOneOf := _HelmRequestAllOfSourceOneOf{}

	err = json.Unmarshal(data, &varHelmRequestAllOfSourceOneOf)

	if err != nil {
		return err
	}

	*o = HelmRequestAllOfSourceOneOf(varHelmRequestAllOfSourceOneOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "git_repository")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmRequestAllOfSourceOneOf struct {
	value *HelmRequestAllOfSourceOneOf
	isSet bool
}

func (v NullableHelmRequestAllOfSourceOneOf) Get() *HelmRequestAllOfSourceOneOf {
	return v.value
}

func (v *NullableHelmRequestAllOfSourceOneOf) Set(val *HelmRequestAllOfSourceOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmRequestAllOfSourceOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmRequestAllOfSourceOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmRequestAllOfSourceOneOf(val *HelmRequestAllOfSourceOneOf) *NullableHelmRequestAllOfSourceOneOf {
	return &NullableHelmRequestAllOfSourceOneOf{value: val, isSet: true}
}

func (v NullableHelmRequestAllOfSourceOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmRequestAllOfSourceOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
