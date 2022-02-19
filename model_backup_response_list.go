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

// BackupResponseList struct for BackupResponseList
type BackupResponseList struct {
	Results []BackupResponse `json:"results,omitempty"`
}

// NewBackupResponseList instantiates a new BackupResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupResponseList() *BackupResponseList {
	this := BackupResponseList{}
	return &this
}

// NewBackupResponseListWithDefaults instantiates a new BackupResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupResponseListWithDefaults() *BackupResponseList {
	this := BackupResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *BackupResponseList) GetResults() []BackupResponse {
	if o == nil || o.Results == nil {
		var ret []BackupResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponseList) GetResultsOk() ([]BackupResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BackupResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BackupResponse and assigns it to the Results field.
func (o *BackupResponseList) SetResults(v []BackupResponse) {
	o.Results = v
}

func (o BackupResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableBackupResponseList struct {
	value *BackupResponseList
	isSet bool
}

func (v NullableBackupResponseList) Get() *BackupResponseList {
	return v.value
}

func (v *NullableBackupResponseList) Set(val *BackupResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupResponseList(val *BackupResponseList) *NullableBackupResponseList {
	return &NullableBackupResponseList{value: val, isSet: true}
}

func (v NullableBackupResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
