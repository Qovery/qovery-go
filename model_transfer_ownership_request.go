/*
 * [BETA] Qovery API
 *
 * - Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.
 *
 * API version: 1.0.0
 * Contact: support+api+documentation@qovery.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// TransferOwnershipRequest struct for TransferOwnershipRequest
type TransferOwnershipRequest struct {
	UserId string `json:"user_id"`
}

// NewTransferOwnershipRequest instantiates a new TransferOwnershipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferOwnershipRequest(userId string) *TransferOwnershipRequest {
	this := TransferOwnershipRequest{}
	this.UserId = userId
	return &this
}

// NewTransferOwnershipRequestWithDefaults instantiates a new TransferOwnershipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferOwnershipRequestWithDefaults() *TransferOwnershipRequest {
	this := TransferOwnershipRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *TransferOwnershipRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *TransferOwnershipRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *TransferOwnershipRequest) SetUserId(v string) {
	o.UserId = v
}

func (o TransferOwnershipRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableTransferOwnershipRequest struct {
	value *TransferOwnershipRequest
	isSet bool
}

func (v NullableTransferOwnershipRequest) Get() *TransferOwnershipRequest {
	return v.value
}

func (v *NullableTransferOwnershipRequest) Set(val *TransferOwnershipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferOwnershipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferOwnershipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferOwnershipRequest(val *TransferOwnershipRequest) *NullableTransferOwnershipRequest {
	return &NullableTransferOwnershipRequest{value: val, isSet: true}
}

func (v NullableTransferOwnershipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferOwnershipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
