/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// EventResponseList struct for EventResponseList
type EventResponseList struct {
	Results *[]EventResponse `json:"results,omitempty"`
}

// NewEventResponseList instantiates a new EventResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResponseList() *EventResponseList {
	this := EventResponseList{}
	return &this
}

// NewEventResponseListWithDefaults instantiates a new EventResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResponseListWithDefaults() *EventResponseList {
	this := EventResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *EventResponseList) GetResults() []EventResponse {
	if o == nil || o.Results == nil {
		var ret []EventResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseList) GetResultsOk() (*[]EventResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EventResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EventResponse and assigns it to the Results field.
func (o *EventResponseList) SetResults(v []EventResponse) {
	o.Results = &v
}

func (o EventResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableEventResponseList struct {
	value *EventResponseList
	isSet bool
}

func (v NullableEventResponseList) Get() *EventResponseList {
	return v.value
}

func (v *NullableEventResponseList) Set(val *EventResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResponseList(val *EventResponseList) *NullableEventResponseList {
	return &NullableEventResponseList{value: val, isSet: true}
}

func (v NullableEventResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
