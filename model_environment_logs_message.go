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

// checks if the EnvironmentLogsMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentLogsMessage{}

// EnvironmentLogsMessage struct for EnvironmentLogsMessage
type EnvironmentLogsMessage struct {
	SafeMessage          *string `json:"safe_message,omitempty"`
	FullDetails          *string `json:"full_details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnvironmentLogsMessage EnvironmentLogsMessage

// NewEnvironmentLogsMessage instantiates a new EnvironmentLogsMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentLogsMessage() *EnvironmentLogsMessage {
	this := EnvironmentLogsMessage{}
	return &this
}

// NewEnvironmentLogsMessageWithDefaults instantiates a new EnvironmentLogsMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentLogsMessageWithDefaults() *EnvironmentLogsMessage {
	this := EnvironmentLogsMessage{}
	return &this
}

// GetSafeMessage returns the SafeMessage field value if set, zero value otherwise.
func (o *EnvironmentLogsMessage) GetSafeMessage() string {
	if o == nil || IsNil(o.SafeMessage) {
		var ret string
		return ret
	}
	return *o.SafeMessage
}

// GetSafeMessageOk returns a tuple with the SafeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogsMessage) GetSafeMessageOk() (*string, bool) {
	if o == nil || IsNil(o.SafeMessage) {
		return nil, false
	}
	return o.SafeMessage, true
}

// HasSafeMessage returns a boolean if a field has been set.
func (o *EnvironmentLogsMessage) HasSafeMessage() bool {
	if o != nil && !IsNil(o.SafeMessage) {
		return true
	}

	return false
}

// SetSafeMessage gets a reference to the given string and assigns it to the SafeMessage field.
func (o *EnvironmentLogsMessage) SetSafeMessage(v string) {
	o.SafeMessage = &v
}

// GetFullDetails returns the FullDetails field value if set, zero value otherwise.
func (o *EnvironmentLogsMessage) GetFullDetails() string {
	if o == nil || IsNil(o.FullDetails) {
		var ret string
		return ret
	}
	return *o.FullDetails
}

// GetFullDetailsOk returns a tuple with the FullDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogsMessage) GetFullDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.FullDetails) {
		return nil, false
	}
	return o.FullDetails, true
}

// HasFullDetails returns a boolean if a field has been set.
func (o *EnvironmentLogsMessage) HasFullDetails() bool {
	if o != nil && !IsNil(o.FullDetails) {
		return true
	}

	return false
}

// SetFullDetails gets a reference to the given string and assigns it to the FullDetails field.
func (o *EnvironmentLogsMessage) SetFullDetails(v string) {
	o.FullDetails = &v
}

func (o EnvironmentLogsMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentLogsMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SafeMessage) {
		toSerialize["safe_message"] = o.SafeMessage
	}
	if !IsNil(o.FullDetails) {
		toSerialize["full_details"] = o.FullDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvironmentLogsMessage) UnmarshalJSON(data []byte) (err error) {
	varEnvironmentLogsMessage := _EnvironmentLogsMessage{}

	err = json.Unmarshal(data, &varEnvironmentLogsMessage)

	if err != nil {
		return err
	}

	*o = EnvironmentLogsMessage(varEnvironmentLogsMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "safe_message")
		delete(additionalProperties, "full_details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnvironmentLogsMessage struct {
	value *EnvironmentLogsMessage
	isSet bool
}

func (v NullableEnvironmentLogsMessage) Get() *EnvironmentLogsMessage {
	return v.value
}

func (v *NullableEnvironmentLogsMessage) Set(val *EnvironmentLogsMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentLogsMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentLogsMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentLogsMessage(val *EnvironmentLogsMessage) *NullableEnvironmentLogsMessage {
	return &NullableEnvironmentLogsMessage{value: val, isSet: true}
}

func (v NullableEnvironmentLogsMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentLogsMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
