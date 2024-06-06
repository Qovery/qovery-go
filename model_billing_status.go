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

// checks if the BillingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingStatus{}

// BillingStatus struct for BillingStatus
type BillingStatus struct {
	IsValid              *bool   `json:"is_valid,omitempty"`
	Message              *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BillingStatus BillingStatus

// NewBillingStatus instantiates a new BillingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingStatus() *BillingStatus {
	this := BillingStatus{}
	return &this
}

// NewBillingStatusWithDefaults instantiates a new BillingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingStatusWithDefaults() *BillingStatus {
	this := BillingStatus{}
	return &this
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *BillingStatus) GetIsValid() bool {
	if o == nil || IsNil(o.IsValid) {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingStatus) GetIsValidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsValid) {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *BillingStatus) HasIsValid() bool {
	if o != nil && !IsNil(o.IsValid) {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *BillingStatus) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BillingStatus) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingStatus) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BillingStatus) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BillingStatus) SetMessage(v string) {
	o.Message = &v
}

func (o BillingStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsValid) {
		toSerialize["is_valid"] = o.IsValid
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BillingStatus) UnmarshalJSON(data []byte) (err error) {
	varBillingStatus := _BillingStatus{}

	err = json.Unmarshal(data, &varBillingStatus)

	if err != nil {
		return err
	}

	*o = BillingStatus(varBillingStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "is_valid")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBillingStatus struct {
	value *BillingStatus
	isSet bool
}

func (v NullableBillingStatus) Get() *BillingStatus {
	return v.value
}

func (v *NullableBillingStatus) Set(val *BillingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingStatus(val *BillingStatus) *NullableBillingStatus {
	return &NullableBillingStatus{value: val, isSet: true}
}

func (v NullableBillingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
