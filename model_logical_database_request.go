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

// LogicalDatabaseRequest struct for LogicalDatabaseRequest
type LogicalDatabaseRequest struct {
	// name is case insensitive
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
}

// NewLogicalDatabaseRequest instantiates a new LogicalDatabaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogicalDatabaseRequest(name string) *LogicalDatabaseRequest {
	this := LogicalDatabaseRequest{}
	this.Name = name
	return &this
}

// NewLogicalDatabaseRequestWithDefaults instantiates a new LogicalDatabaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogicalDatabaseRequestWithDefaults() *LogicalDatabaseRequest {
	this := LogicalDatabaseRequest{}
	return &this
}

// GetName returns the Name field value
func (o *LogicalDatabaseRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogicalDatabaseRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LogicalDatabaseRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LogicalDatabaseRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogicalDatabaseRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LogicalDatabaseRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LogicalDatabaseRequest) SetDescription(v string) {
	o.Description = &v
}

func (o LogicalDatabaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableLogicalDatabaseRequest struct {
	value *LogicalDatabaseRequest
	isSet bool
}

func (v NullableLogicalDatabaseRequest) Get() *LogicalDatabaseRequest {
	return v.value
}

func (v *NullableLogicalDatabaseRequest) Set(val *LogicalDatabaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLogicalDatabaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLogicalDatabaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogicalDatabaseRequest(val *LogicalDatabaseRequest) *NullableLogicalDatabaseRequest {
	return &NullableLogicalDatabaseRequest{value: val, isSet: true}
}

func (v NullableLogicalDatabaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogicalDatabaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


