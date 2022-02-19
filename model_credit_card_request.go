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

// CreditCardRequest struct for CreditCardRequest
type CreditCardRequest struct {
	Number      string `json:"number"`
	Cvv         string `json:"cvv"`
	ExpiryMonth int32  `json:"expiry_month"`
	ExpiryYear  int32  `json:"expiry_year"`
}

// NewCreditCardRequest instantiates a new CreditCardRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditCardRequest(number string, cvv string, expiryMonth int32, expiryYear int32) *CreditCardRequest {
	this := CreditCardRequest{}
	this.Number = number
	this.Cvv = cvv
	this.ExpiryMonth = expiryMonth
	this.ExpiryYear = expiryYear
	return &this
}

// NewCreditCardRequestWithDefaults instantiates a new CreditCardRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditCardRequestWithDefaults() *CreditCardRequest {
	this := CreditCardRequest{}
	return &this
}

// GetNumber returns the Number field value
func (o *CreditCardRequest) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *CreditCardRequest) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *CreditCardRequest) SetNumber(v string) {
	o.Number = v
}

// GetCvv returns the Cvv field value
func (o *CreditCardRequest) GetCvv() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cvv
}

// GetCvvOk returns a tuple with the Cvv field value
// and a boolean to check if the value has been set.
func (o *CreditCardRequest) GetCvvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cvv, true
}

// SetCvv sets field value
func (o *CreditCardRequest) SetCvv(v string) {
	o.Cvv = v
}

// GetExpiryMonth returns the ExpiryMonth field value
func (o *CreditCardRequest) GetExpiryMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiryMonth
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value
// and a boolean to check if the value has been set.
func (o *CreditCardRequest) GetExpiryMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryMonth, true
}

// SetExpiryMonth sets field value
func (o *CreditCardRequest) SetExpiryMonth(v int32) {
	o.ExpiryMonth = v
}

// GetExpiryYear returns the ExpiryYear field value
func (o *CreditCardRequest) GetExpiryYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiryYear
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value
// and a boolean to check if the value has been set.
func (o *CreditCardRequest) GetExpiryYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryYear, true
}

// SetExpiryYear sets field value
func (o *CreditCardRequest) SetExpiryYear(v int32) {
	o.ExpiryYear = v
}

func (o CreditCardRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["number"] = o.Number
	}
	if true {
		toSerialize["cvv"] = o.Cvv
	}
	if true {
		toSerialize["expiry_month"] = o.ExpiryMonth
	}
	if true {
		toSerialize["expiry_year"] = o.ExpiryYear
	}
	return json.Marshal(toSerialize)
}

type NullableCreditCardRequest struct {
	value *CreditCardRequest
	isSet bool
}

func (v NullableCreditCardRequest) Get() *CreditCardRequest {
	return v.value
}

func (v *NullableCreditCardRequest) Set(val *CreditCardRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditCardRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditCardRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditCardRequest(val *CreditCardRequest) *NullableCreditCardRequest {
	return &NullableCreditCardRequest{value: val, isSet: true}
}

func (v NullableCreditCardRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditCardRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
