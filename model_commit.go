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
	"time"
)

// checks if the Commit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Commit{}

// Commit struct for Commit
type Commit struct {
	CreatedAt            time.Time `json:"created_at"`
	GitCommitId          string    `json:"git_commit_id"`
	Tag                  string    `json:"tag"`
	Message              string    `json:"message"`
	AuthorName           string    `json:"author_name"`
	AuthorAvatarUrl      *string   `json:"author_avatar_url,omitempty"`
	CommitPageUrl        *string   `json:"commit_page_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Commit Commit

// NewCommit instantiates a new Commit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommit(createdAt time.Time, gitCommitId string, tag string, message string, authorName string) *Commit {
	this := Commit{}
	this.CreatedAt = createdAt
	this.GitCommitId = gitCommitId
	this.Tag = tag
	this.Message = message
	this.AuthorName = authorName
	return &this
}

// NewCommitWithDefaults instantiates a new Commit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitWithDefaults() *Commit {
	this := Commit{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *Commit) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Commit) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Commit) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetGitCommitId returns the GitCommitId field value
func (o *Commit) GetGitCommitId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GitCommitId
}

// GetGitCommitIdOk returns a tuple with the GitCommitId field value
// and a boolean to check if the value has been set.
func (o *Commit) GetGitCommitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitCommitId, true
}

// SetGitCommitId sets field value
func (o *Commit) SetGitCommitId(v string) {
	o.GitCommitId = v
}

// GetTag returns the Tag field value
func (o *Commit) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *Commit) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *Commit) SetTag(v string) {
	o.Tag = v
}

// GetMessage returns the Message field value
func (o *Commit) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Commit) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Commit) SetMessage(v string) {
	o.Message = v
}

// GetAuthorName returns the AuthorName field value
func (o *Commit) GetAuthorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value
// and a boolean to check if the value has been set.
func (o *Commit) GetAuthorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorName, true
}

// SetAuthorName sets field value
func (o *Commit) SetAuthorName(v string) {
	o.AuthorName = v
}

// GetAuthorAvatarUrl returns the AuthorAvatarUrl field value if set, zero value otherwise.
func (o *Commit) GetAuthorAvatarUrl() string {
	if o == nil || IsNil(o.AuthorAvatarUrl) {
		var ret string
		return ret
	}
	return *o.AuthorAvatarUrl
}

// GetAuthorAvatarUrlOk returns a tuple with the AuthorAvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetAuthorAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorAvatarUrl) {
		return nil, false
	}
	return o.AuthorAvatarUrl, true
}

// HasAuthorAvatarUrl returns a boolean if a field has been set.
func (o *Commit) HasAuthorAvatarUrl() bool {
	if o != nil && !IsNil(o.AuthorAvatarUrl) {
		return true
	}

	return false
}

// SetAuthorAvatarUrl gets a reference to the given string and assigns it to the AuthorAvatarUrl field.
func (o *Commit) SetAuthorAvatarUrl(v string) {
	o.AuthorAvatarUrl = &v
}

// GetCommitPageUrl returns the CommitPageUrl field value if set, zero value otherwise.
func (o *Commit) GetCommitPageUrl() string {
	if o == nil || IsNil(o.CommitPageUrl) {
		var ret string
		return ret
	}
	return *o.CommitPageUrl
}

// GetCommitPageUrlOk returns a tuple with the CommitPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetCommitPageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CommitPageUrl) {
		return nil, false
	}
	return o.CommitPageUrl, true
}

// HasCommitPageUrl returns a boolean if a field has been set.
func (o *Commit) HasCommitPageUrl() bool {
	if o != nil && !IsNil(o.CommitPageUrl) {
		return true
	}

	return false
}

// SetCommitPageUrl gets a reference to the given string and assigns it to the CommitPageUrl field.
func (o *Commit) SetCommitPageUrl(v string) {
	o.CommitPageUrl = &v
}

func (o Commit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Commit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["git_commit_id"] = o.GitCommitId
	toSerialize["tag"] = o.Tag
	toSerialize["message"] = o.Message
	toSerialize["author_name"] = o.AuthorName
	if !IsNil(o.AuthorAvatarUrl) {
		toSerialize["author_avatar_url"] = o.AuthorAvatarUrl
	}
	if !IsNil(o.CommitPageUrl) {
		toSerialize["commit_page_url"] = o.CommitPageUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Commit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"git_commit_id",
		"tag",
		"message",
		"author_name",
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

	varCommit := _Commit{}

	err = json.Unmarshal(data, &varCommit)

	if err != nil {
		return err
	}

	*o = Commit(varCommit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "git_commit_id")
		delete(additionalProperties, "tag")
		delete(additionalProperties, "message")
		delete(additionalProperties, "author_name")
		delete(additionalProperties, "author_avatar_url")
		delete(additionalProperties, "commit_page_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommit struct {
	value *Commit
	isSet bool
}

func (v NullableCommit) Get() *Commit {
	return v.value
}

func (v *NullableCommit) Set(val *Commit) {
	v.value = val
	v.isSet = true
}

func (v NullableCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommit(val *Commit) *NullableCommit {
	return &NullableCommit{value: val, isSet: true}
}

func (v NullableCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
