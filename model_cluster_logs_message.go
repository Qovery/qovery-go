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

// checks if the ClusterLogsMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterLogsMessage{}

// ClusterLogsMessage struct for ClusterLogsMessage
type ClusterLogsMessage struct {
	// log global message
	SafeMessage *string `json:"safe_message,omitempty"`
}

// NewClusterLogsMessage instantiates a new ClusterLogsMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterLogsMessage() *ClusterLogsMessage {
	this := ClusterLogsMessage{}
	return &this
}

// NewClusterLogsMessageWithDefaults instantiates a new ClusterLogsMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterLogsMessageWithDefaults() *ClusterLogsMessage {
	this := ClusterLogsMessage{}
	return &this
}

// GetSafeMessage returns the SafeMessage field value if set, zero value otherwise.
func (o *ClusterLogsMessage) GetSafeMessage() string {
	if o == nil || IsNil(o.SafeMessage) {
		var ret string
		return ret
	}
	return *o.SafeMessage
}

// GetSafeMessageOk returns a tuple with the SafeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsMessage) GetSafeMessageOk() (*string, bool) {
	if o == nil || IsNil(o.SafeMessage) {
		return nil, false
	}
	return o.SafeMessage, true
}

// HasSafeMessage returns a boolean if a field has been set.
func (o *ClusterLogsMessage) HasSafeMessage() bool {
	if o != nil && !IsNil(o.SafeMessage) {
		return true
	}

	return false
}

// SetSafeMessage gets a reference to the given string and assigns it to the SafeMessage field.
func (o *ClusterLogsMessage) SetSafeMessage(v string) {
	o.SafeMessage = &v
}

func (o ClusterLogsMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterLogsMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SafeMessage) {
		toSerialize["safe_message"] = o.SafeMessage
	}
	return toSerialize, nil
}

type NullableClusterLogsMessage struct {
	value *ClusterLogsMessage
	isSet bool
}

func (v NullableClusterLogsMessage) Get() *ClusterLogsMessage {
	return v.value
}

func (v *NullableClusterLogsMessage) Set(val *ClusterLogsMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterLogsMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterLogsMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterLogsMessage(val *ClusterLogsMessage) *NullableClusterLogsMessage {
	return &NullableClusterLogsMessage{value: val, isSet: true}
}

func (v NullableClusterLogsMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterLogsMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
