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

// CreditCardResponseList struct for CreditCardResponseList
type CreditCardResponseList struct {
	Results []CreditCardResponse `json:"results,omitempty"`
}

// NewCreditCardResponseList instantiates a new CreditCardResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditCardResponseList() *CreditCardResponseList {
	this := CreditCardResponseList{}
	return &this
}

// NewCreditCardResponseListWithDefaults instantiates a new CreditCardResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditCardResponseListWithDefaults() *CreditCardResponseList {
	this := CreditCardResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CreditCardResponseList) GetResults() []CreditCardResponse {
	if o == nil || o.Results == nil {
		var ret []CreditCardResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardResponseList) GetResultsOk() ([]CreditCardResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CreditCardResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CreditCardResponse and assigns it to the Results field.
func (o *CreditCardResponseList) SetResults(v []CreditCardResponse) {
	o.Results = v
}

func (o CreditCardResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableCreditCardResponseList struct {
	value *CreditCardResponseList
	isSet bool
}

func (v NullableCreditCardResponseList) Get() *CreditCardResponseList {
	return v.value
}

func (v *NullableCreditCardResponseList) Set(val *CreditCardResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditCardResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditCardResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditCardResponseList(val *CreditCardResponseList) *NullableCreditCardResponseList {
	return &NullableCreditCardResponseList{value: val, isSet: true}
}

func (v NullableCreditCardResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditCardResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
