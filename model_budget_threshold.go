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

// checks if the BudgetThreshold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetThreshold{}

// BudgetThreshold struct for BudgetThreshold
type BudgetThreshold struct {
	TotalInCents *int32   `json:"total_in_cents,omitempty"`
	Total        *float32 `json:"total,omitempty"`
	CurrencyCode *string  `json:"currency_code,omitempty"`
}

// NewBudgetThreshold instantiates a new BudgetThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetThreshold() *BudgetThreshold {
	this := BudgetThreshold{}
	return &this
}

// NewBudgetThresholdWithDefaults instantiates a new BudgetThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetThresholdWithDefaults() *BudgetThreshold {
	this := BudgetThreshold{}
	return &this
}

// GetTotalInCents returns the TotalInCents field value if set, zero value otherwise.
func (o *BudgetThreshold) GetTotalInCents() int32 {
	if o == nil || IsNil(o.TotalInCents) {
		var ret int32
		return ret
	}
	return *o.TotalInCents
}

// GetTotalInCentsOk returns a tuple with the TotalInCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetThreshold) GetTotalInCentsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalInCents) {
		return nil, false
	}
	return o.TotalInCents, true
}

// HasTotalInCents returns a boolean if a field has been set.
func (o *BudgetThreshold) HasTotalInCents() bool {
	if o != nil && !IsNil(o.TotalInCents) {
		return true
	}

	return false
}

// SetTotalInCents gets a reference to the given int32 and assigns it to the TotalInCents field.
func (o *BudgetThreshold) SetTotalInCents(v int32) {
	o.TotalInCents = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *BudgetThreshold) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetThreshold) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *BudgetThreshold) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *BudgetThreshold) SetTotal(v float32) {
	o.Total = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *BudgetThreshold) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetThreshold) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *BudgetThreshold) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *BudgetThreshold) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

func (o BudgetThreshold) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BudgetThreshold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalInCents) {
		toSerialize["total_in_cents"] = o.TotalInCents
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	return toSerialize, nil
}

type NullableBudgetThreshold struct {
	value *BudgetThreshold
	isSet bool
}

func (v NullableBudgetThreshold) Get() *BudgetThreshold {
	return v.value
}

func (v *NullableBudgetThreshold) Set(val *BudgetThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetThreshold(val *BudgetThreshold) *NullableBudgetThreshold {
	return &NullableBudgetThreshold{value: val, isSet: true}
}

func (v NullableBudgetThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
