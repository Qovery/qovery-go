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

// BackupPaginatedResponseList struct for BackupPaginatedResponseList
type BackupPaginatedResponseList struct {
	Page     float32          `json:"page"`
	PageSize float32          `json:"page_size"`
	Results  []BackupResponse `json:"results,omitempty"`
}

// NewBackupPaginatedResponseList instantiates a new BackupPaginatedResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupPaginatedResponseList(page float32, pageSize float32) *BackupPaginatedResponseList {
	this := BackupPaginatedResponseList{}
	this.Page = page
	this.PageSize = pageSize
	return &this
}

// NewBackupPaginatedResponseListWithDefaults instantiates a new BackupPaginatedResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupPaginatedResponseListWithDefaults() *BackupPaginatedResponseList {
	this := BackupPaginatedResponseList{}
	return &this
}

// GetPage returns the Page field value
func (o *BackupPaginatedResponseList) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *BackupPaginatedResponseList) GetPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *BackupPaginatedResponseList) SetPage(v float32) {
	o.Page = v
}

// GetPageSize returns the PageSize field value
func (o *BackupPaginatedResponseList) GetPageSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *BackupPaginatedResponseList) GetPageSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *BackupPaginatedResponseList) SetPageSize(v float32) {
	o.PageSize = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *BackupPaginatedResponseList) GetResults() []BackupResponse {
	if o == nil || o.Results == nil {
		var ret []BackupResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPaginatedResponseList) GetResultsOk() ([]BackupResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BackupPaginatedResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BackupResponse and assigns it to the Results field.
func (o *BackupPaginatedResponseList) SetResults(v []BackupResponse) {
	o.Results = v
}

func (o BackupPaginatedResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["page_size"] = o.PageSize
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableBackupPaginatedResponseList struct {
	value *BackupPaginatedResponseList
	isSet bool
}

func (v NullableBackupPaginatedResponseList) Get() *BackupPaginatedResponseList {
	return v.value
}

func (v *NullableBackupPaginatedResponseList) Set(val *BackupPaginatedResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupPaginatedResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupPaginatedResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupPaginatedResponseList(val *BackupPaginatedResponseList) *NullableBackupPaginatedResponseList {
	return &NullableBackupPaginatedResponseList{value: val, isSet: true}
}

func (v NullableBackupPaginatedResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupPaginatedResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
