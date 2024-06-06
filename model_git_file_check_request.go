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
	"fmt"
)

// checks if the GitFileCheckRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitFileCheckRequest{}

// GitFileCheckRequest struct for GitFileCheckRequest
type GitFileCheckRequest struct {
	GitRepository        HelmGitRepositoryRequest `json:"git_repository"`
	Files                []string                 `json:"files"`
	AdditionalProperties map[string]interface{}
}

type _GitFileCheckRequest GitFileCheckRequest

// NewGitFileCheckRequest instantiates a new GitFileCheckRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitFileCheckRequest(gitRepository HelmGitRepositoryRequest, files []string) *GitFileCheckRequest {
	this := GitFileCheckRequest{}
	this.GitRepository = gitRepository
	this.Files = files
	return &this
}

// NewGitFileCheckRequestWithDefaults instantiates a new GitFileCheckRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitFileCheckRequestWithDefaults() *GitFileCheckRequest {
	this := GitFileCheckRequest{}
	return &this
}

// GetGitRepository returns the GitRepository field value
func (o *GitFileCheckRequest) GetGitRepository() HelmGitRepositoryRequest {
	if o == nil {
		var ret HelmGitRepositoryRequest
		return ret
	}

	return o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value
// and a boolean to check if the value has been set.
func (o *GitFileCheckRequest) GetGitRepositoryOk() (*HelmGitRepositoryRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitRepository, true
}

// SetGitRepository sets field value
func (o *GitFileCheckRequest) SetGitRepository(v HelmGitRepositoryRequest) {
	o.GitRepository = v
}

// GetFiles returns the Files field value
func (o *GitFileCheckRequest) GetFiles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *GitFileCheckRequest) GetFilesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *GitFileCheckRequest) SetFiles(v []string) {
	o.Files = v
}

func (o GitFileCheckRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitFileCheckRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["git_repository"] = o.GitRepository
	toSerialize["files"] = o.Files

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GitFileCheckRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"git_repository",
		"files",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGitFileCheckRequest := _GitFileCheckRequest{}

	err = json.Unmarshal(data, &varGitFileCheckRequest)

	if err != nil {
		return err
	}

	*o = GitFileCheckRequest(varGitFileCheckRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "git_repository")
		delete(additionalProperties, "files")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGitFileCheckRequest struct {
	value *GitFileCheckRequest
	isSet bool
}

func (v NullableGitFileCheckRequest) Get() *GitFileCheckRequest {
	return v.value
}

func (v *NullableGitFileCheckRequest) Set(val *GitFileCheckRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGitFileCheckRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGitFileCheckRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitFileCheckRequest(val *GitFileCheckRequest) *NullableGitFileCheckRequest {
	return &NullableGitFileCheckRequest{value: val, isSet: true}
}

func (v NullableGitFileCheckRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitFileCheckRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
