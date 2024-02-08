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

// checks if the RewardClaim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RewardClaim{}

// RewardClaim struct for RewardClaim
type RewardClaim struct {
	Type *string `json:"type,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewRewardClaim instantiates a new RewardClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRewardClaim() *RewardClaim {
	this := RewardClaim{}
	return &this
}

// NewRewardClaimWithDefaults instantiates a new RewardClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRewardClaimWithDefaults() *RewardClaim {
	this := RewardClaim{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RewardClaim) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardClaim) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RewardClaim) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RewardClaim) SetType(v string) {
	o.Type = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RewardClaim) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardClaim) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RewardClaim) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *RewardClaim) SetCode(v string) {
	o.Code = &v
}

func (o RewardClaim) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RewardClaim) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableRewardClaim struct {
	value *RewardClaim
	isSet bool
}

func (v NullableRewardClaim) Get() *RewardClaim {
	return v.value
}

func (v *NullableRewardClaim) Set(val *RewardClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableRewardClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableRewardClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRewardClaim(val *RewardClaim) *NullableRewardClaim {
	return &NullableRewardClaim{value: val, isSet: true}
}

func (v NullableRewardClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRewardClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
