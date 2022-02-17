/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// CredentialsResponse struct for CredentialsResponse
type CredentialsResponse struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// NewCredentialsResponse instantiates a new CredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsResponse(login string, password string) *CredentialsResponse {
	this := CredentialsResponse{}
	this.Login = login
	this.Password = password
	return &this
}

// NewCredentialsResponseWithDefaults instantiates a new CredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsResponseWithDefaults() *CredentialsResponse {
	this := CredentialsResponse{}
	return &this
}

// GetLogin returns the Login field value
func (o *CredentialsResponse) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *CredentialsResponse) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *CredentialsResponse) SetLogin(v string) {
	o.Login = v
}

// GetPassword returns the Password field value
func (o *CredentialsResponse) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CredentialsResponse) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CredentialsResponse) SetPassword(v string) {
	o.Password = v
}

func (o CredentialsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["login"] = o.Login
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialsResponse struct {
	value *CredentialsResponse
	isSet bool
}

func (v NullableCredentialsResponse) Get() *CredentialsResponse {
	return v.value
}

func (v *NullableCredentialsResponse) Set(val *CredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsResponse(val *CredentialsResponse) *NullableCredentialsResponse {
	return &NullableCredentialsResponse{value: val, isSet: true}
}

func (v NullableCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
