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

// checks if the ApplicationGitRepositoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationGitRepositoryRequest{}

// ApplicationGitRepositoryRequest struct for ApplicationGitRepositoryRequest
type ApplicationGitRepositoryRequest struct {
	// application git repository URL
	Url string `json:"url"`
	// Name of the branch to use. This is optional If not specified, then the branch used is the `main` or `master` one
	Branch *string `json:"branch,omitempty"`
	// indicates the root path of the application.
	RootPath *string `json:"root_path,omitempty"`
	// The git token id on Qovery side
	GitTokenId           NullableString `json:"git_token_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationGitRepositoryRequest ApplicationGitRepositoryRequest

// NewApplicationGitRepositoryRequest instantiates a new ApplicationGitRepositoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationGitRepositoryRequest(url string) *ApplicationGitRepositoryRequest {
	this := ApplicationGitRepositoryRequest{}
	this.Url = url
	var rootPath string = "/"
	this.RootPath = &rootPath
	return &this
}

// NewApplicationGitRepositoryRequestWithDefaults instantiates a new ApplicationGitRepositoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationGitRepositoryRequestWithDefaults() *ApplicationGitRepositoryRequest {
	this := ApplicationGitRepositoryRequest{}
	var rootPath string = "/"
	this.RootPath = &rootPath
	return &this
}

// GetUrl returns the Url field value
func (o *ApplicationGitRepositoryRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ApplicationGitRepositoryRequest) SetUrl(v string) {
	o.Url = v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *ApplicationGitRepositoryRequest) GetBranch() string {
	if o == nil || IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryRequest) GetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *ApplicationGitRepositoryRequest) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *ApplicationGitRepositoryRequest) SetBranch(v string) {
	o.Branch = &v
}

// GetRootPath returns the RootPath field value if set, zero value otherwise.
func (o *ApplicationGitRepositoryRequest) GetRootPath() string {
	if o == nil || IsNil(o.RootPath) {
		var ret string
		return ret
	}
	return *o.RootPath
}

// GetRootPathOk returns a tuple with the RootPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryRequest) GetRootPathOk() (*string, bool) {
	if o == nil || IsNil(o.RootPath) {
		return nil, false
	}
	return o.RootPath, true
}

// HasRootPath returns a boolean if a field has been set.
func (o *ApplicationGitRepositoryRequest) HasRootPath() bool {
	if o != nil && !IsNil(o.RootPath) {
		return true
	}

	return false
}

// SetRootPath gets a reference to the given string and assigns it to the RootPath field.
func (o *ApplicationGitRepositoryRequest) SetRootPath(v string) {
	o.RootPath = &v
}

// GetGitTokenId returns the GitTokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationGitRepositoryRequest) GetGitTokenId() string {
	if o == nil || IsNil(o.GitTokenId.Get()) {
		var ret string
		return ret
	}
	return *o.GitTokenId.Get()
}

// GetGitTokenIdOk returns a tuple with the GitTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationGitRepositoryRequest) GetGitTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GitTokenId.Get(), o.GitTokenId.IsSet()
}

// HasGitTokenId returns a boolean if a field has been set.
func (o *ApplicationGitRepositoryRequest) HasGitTokenId() bool {
	if o != nil && o.GitTokenId.IsSet() {
		return true
	}

	return false
}

// SetGitTokenId gets a reference to the given NullableString and assigns it to the GitTokenId field.
func (o *ApplicationGitRepositoryRequest) SetGitTokenId(v string) {
	o.GitTokenId.Set(&v)
}

// SetGitTokenIdNil sets the value for GitTokenId to be an explicit nil
func (o *ApplicationGitRepositoryRequest) SetGitTokenIdNil() {
	o.GitTokenId.Set(nil)
}

// UnsetGitTokenId ensures that no value is present for GitTokenId, not even an explicit nil
func (o *ApplicationGitRepositoryRequest) UnsetGitTokenId() {
	o.GitTokenId.Unset()
}

func (o ApplicationGitRepositoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationGitRepositoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !IsNil(o.RootPath) {
		toSerialize["root_path"] = o.RootPath
	}
	if o.GitTokenId.IsSet() {
		toSerialize["git_token_id"] = o.GitTokenId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationGitRepositoryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varApplicationGitRepositoryRequest := _ApplicationGitRepositoryRequest{}

	err = json.Unmarshal(data, &varApplicationGitRepositoryRequest)

	if err != nil {
		return err
	}

	*o = ApplicationGitRepositoryRequest(varApplicationGitRepositoryRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "branch")
		delete(additionalProperties, "root_path")
		delete(additionalProperties, "git_token_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationGitRepositoryRequest struct {
	value *ApplicationGitRepositoryRequest
	isSet bool
}

func (v NullableApplicationGitRepositoryRequest) Get() *ApplicationGitRepositoryRequest {
	return v.value
}

func (v *NullableApplicationGitRepositoryRequest) Set(val *ApplicationGitRepositoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationGitRepositoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationGitRepositoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationGitRepositoryRequest(val *ApplicationGitRepositoryRequest) *NullableApplicationGitRepositoryRequest {
	return &NullableApplicationGitRepositoryRequest{value: val, isSet: true}
}

func (v NullableApplicationGitRepositoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationGitRepositoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
