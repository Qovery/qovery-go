/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GitAuthProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitAuthProvider{}

// GitAuthProvider struct for GitAuthProvider
type GitAuthProvider struct {
	Id     *string `json:"id,omitempty"`
	Name   string  `json:"name"`
	Owner  string  `json:"owner"`
	UseBot *bool   `json:"use_bot,omitempty"`
}

type _GitAuthProvider GitAuthProvider

// NewGitAuthProvider instantiates a new GitAuthProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitAuthProvider(name string, owner string) *GitAuthProvider {
	this := GitAuthProvider{}
	this.Name = name
	this.Owner = owner
	return &this
}

// NewGitAuthProviderWithDefaults instantiates a new GitAuthProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitAuthProviderWithDefaults() *GitAuthProvider {
	this := GitAuthProvider{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GitAuthProvider) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitAuthProvider) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GitAuthProvider) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GitAuthProvider) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *GitAuthProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GitAuthProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GitAuthProvider) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value
func (o *GitAuthProvider) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *GitAuthProvider) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *GitAuthProvider) SetOwner(v string) {
	o.Owner = v
}

// GetUseBot returns the UseBot field value if set, zero value otherwise.
func (o *GitAuthProvider) GetUseBot() bool {
	if o == nil || IsNil(o.UseBot) {
		var ret bool
		return ret
	}
	return *o.UseBot
}

// GetUseBotOk returns a tuple with the UseBot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitAuthProvider) GetUseBotOk() (*bool, bool) {
	if o == nil || IsNil(o.UseBot) {
		return nil, false
	}
	return o.UseBot, true
}

// HasUseBot returns a boolean if a field has been set.
func (o *GitAuthProvider) HasUseBot() bool {
	if o != nil && !IsNil(o.UseBot) {
		return true
	}

	return false
}

// SetUseBot gets a reference to the given bool and assigns it to the UseBot field.
func (o *GitAuthProvider) SetUseBot(v bool) {
	o.UseBot = &v
}

func (o GitAuthProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitAuthProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["owner"] = o.Owner
	if !IsNil(o.UseBot) {
		toSerialize["use_bot"] = o.UseBot
	}
	return toSerialize, nil
}

func (o *GitAuthProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"owner",
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

	varGitAuthProvider := _GitAuthProvider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGitAuthProvider)

	if err != nil {
		return err
	}

	*o = GitAuthProvider(varGitAuthProvider)

	return err
}

type NullableGitAuthProvider struct {
	value *GitAuthProvider
	isSet bool
}

func (v NullableGitAuthProvider) Get() *GitAuthProvider {
	return v.value
}

func (v *NullableGitAuthProvider) Set(val *GitAuthProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableGitAuthProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableGitAuthProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitAuthProvider(val *GitAuthProvider) *NullableGitAuthProvider {
	return &NullableGitAuthProvider{value: val, isSet: true}
}

func (v NullableGitAuthProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitAuthProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
