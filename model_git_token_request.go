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

// checks if the GitTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitTokenRequest{}

// GitTokenRequest struct for GitTokenRequest
type GitTokenRequest struct {
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Type        GitProviderEnum `json:"type"`
	// The token from your git provider side
	Token string `json:"token"`
	// Mandatory only for BITBUCKET git provider, to allow us to fetch repositories at creation/edition of a service
	Workspace *string `json:"workspace,omitempty"`
}

// NewGitTokenRequest instantiates a new GitTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitTokenRequest(name string, type_ GitProviderEnum, token string) *GitTokenRequest {
	this := GitTokenRequest{}
	this.Name = name
	this.Type = type_
	this.Token = token
	return &this
}

// NewGitTokenRequestWithDefaults instantiates a new GitTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitTokenRequestWithDefaults() *GitTokenRequest {
	this := GitTokenRequest{}
	return &this
}

// GetName returns the Name field value
func (o *GitTokenRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GitTokenRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GitTokenRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GitTokenRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTokenRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GitTokenRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GitTokenRequest) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *GitTokenRequest) GetType() GitProviderEnum {
	if o == nil {
		var ret GitProviderEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GitTokenRequest) GetTypeOk() (*GitProviderEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GitTokenRequest) SetType(v GitProviderEnum) {
	o.Type = v
}

// GetToken returns the Token field value
func (o *GitTokenRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *GitTokenRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *GitTokenRequest) SetToken(v string) {
	o.Token = v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *GitTokenRequest) GetWorkspace() string {
	if o == nil || IsNil(o.Workspace) {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTokenRequest) GetWorkspaceOk() (*string, bool) {
	if o == nil || IsNil(o.Workspace) {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *GitTokenRequest) HasWorkspace() bool {
	if o != nil && !IsNil(o.Workspace) {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *GitTokenRequest) SetWorkspace(v string) {
	o.Workspace = &v
}

func (o GitTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	toSerialize["token"] = o.Token
	if !IsNil(o.Workspace) {
		toSerialize["workspace"] = o.Workspace
	}
	return toSerialize, nil
}

type NullableGitTokenRequest struct {
	value *GitTokenRequest
	isSet bool
}

func (v NullableGitTokenRequest) Get() *GitTokenRequest {
	return v.value
}

func (v *NullableGitTokenRequest) Set(val *GitTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGitTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGitTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitTokenRequest(val *GitTokenRequest) *NullableGitTokenRequest {
	return &NullableGitTokenRequest{value: val, isSet: true}
}

func (v NullableGitTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
