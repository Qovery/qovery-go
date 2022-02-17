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

// CustomDomainRequest struct for CustomDomainRequest
type CustomDomainRequest struct {
	// your custom domain
	Domain string `json:"domain"`
}

// NewCustomDomainRequest instantiates a new CustomDomainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDomainRequest(domain string) *CustomDomainRequest {
	this := CustomDomainRequest{}
	this.Domain = domain
	return &this
}

// NewCustomDomainRequestWithDefaults instantiates a new CustomDomainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDomainRequestWithDefaults() *CustomDomainRequest {
	this := CustomDomainRequest{}
	return &this
}

// GetDomain returns the Domain field value
func (o *CustomDomainRequest) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *CustomDomainRequest) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *CustomDomainRequest) SetDomain(v string) {
	o.Domain = v
}

func (o CustomDomainRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["domain"] = o.Domain
	}
	return json.Marshal(toSerialize)
}

type NullableCustomDomainRequest struct {
	value *CustomDomainRequest
	isSet bool
}

func (v NullableCustomDomainRequest) Get() *CustomDomainRequest {
	return v.value
}

func (v *NullableCustomDomainRequest) Set(val *CustomDomainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDomainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDomainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDomainRequest(val *CustomDomainRequest) *NullableCustomDomainRequest {
	return &NullableCustomDomainRequest{value: val, isSet: true}
}

func (v NullableCustomDomainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDomainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
