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

// AccountInfoEditRequest struct for AccountInfoEditRequest
type AccountInfoEditRequest struct {
	FirstName         *string `json:"first_name,omitempty"`
	LastName          *string `json:"last_name,omitempty"`
	ProfilePictureUrl *string `json:"profile_picture_url,omitempty"`
}

// NewAccountInfoEditRequest instantiates a new AccountInfoEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInfoEditRequest() *AccountInfoEditRequest {
	this := AccountInfoEditRequest{}
	return &this
}

// NewAccountInfoEditRequestWithDefaults instantiates a new AccountInfoEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInfoEditRequestWithDefaults() *AccountInfoEditRequest {
	this := AccountInfoEditRequest{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *AccountInfoEditRequest) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoEditRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *AccountInfoEditRequest) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *AccountInfoEditRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *AccountInfoEditRequest) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoEditRequest) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *AccountInfoEditRequest) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *AccountInfoEditRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetProfilePictureUrl returns the ProfilePictureUrl field value if set, zero value otherwise.
func (o *AccountInfoEditRequest) GetProfilePictureUrl() string {
	if o == nil || o.ProfilePictureUrl == nil {
		var ret string
		return ret
	}
	return *o.ProfilePictureUrl
}

// GetProfilePictureUrlOk returns a tuple with the ProfilePictureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoEditRequest) GetProfilePictureUrlOk() (*string, bool) {
	if o == nil || o.ProfilePictureUrl == nil {
		return nil, false
	}
	return o.ProfilePictureUrl, true
}

// HasProfilePictureUrl returns a boolean if a field has been set.
func (o *AccountInfoEditRequest) HasProfilePictureUrl() bool {
	if o != nil && o.ProfilePictureUrl != nil {
		return true
	}

	return false
}

// SetProfilePictureUrl gets a reference to the given string and assigns it to the ProfilePictureUrl field.
func (o *AccountInfoEditRequest) SetProfilePictureUrl(v string) {
	o.ProfilePictureUrl = &v
}

func (o AccountInfoEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.ProfilePictureUrl != nil {
		toSerialize["profile_picture_url"] = o.ProfilePictureUrl
	}
	return json.Marshal(toSerialize)
}

type NullableAccountInfoEditRequest struct {
	value *AccountInfoEditRequest
	isSet bool
}

func (v NullableAccountInfoEditRequest) Get() *AccountInfoEditRequest {
	return v.value
}

func (v *NullableAccountInfoEditRequest) Set(val *AccountInfoEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInfoEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInfoEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInfoEditRequest(val *AccountInfoEditRequest) *NullableAccountInfoEditRequest {
	return &NullableAccountInfoEditRequest{value: val, isSet: true}
}

func (v NullableAccountInfoEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInfoEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
