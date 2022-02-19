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

// InviteMemberRequest struct for InviteMemberRequest
type InviteMemberRequest struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

// NewInviteMemberRequest instantiates a new InviteMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteMemberRequest(email string, role string) *InviteMemberRequest {
	this := InviteMemberRequest{}
	this.Email = email
	this.Role = role
	return &this
}

// NewInviteMemberRequestWithDefaults instantiates a new InviteMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteMemberRequestWithDefaults() *InviteMemberRequest {
	this := InviteMemberRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *InviteMemberRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InviteMemberRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InviteMemberRequest) SetEmail(v string) {
	o.Email = v
}

// GetRole returns the Role field value
func (o *InviteMemberRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *InviteMemberRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *InviteMemberRequest) SetRole(v string) {
	o.Role = v
}

func (o InviteMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableInviteMemberRequest struct {
	value *InviteMemberRequest
	isSet bool
}

func (v NullableInviteMemberRequest) Get() *InviteMemberRequest {
	return v.value
}

func (v *NullableInviteMemberRequest) Set(val *InviteMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteMemberRequest(val *InviteMemberRequest) *NullableInviteMemberRequest {
	return &NullableInviteMemberRequest{value: val, isSet: true}
}

func (v NullableInviteMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
