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

// EnvironmentApplicationsStorageResponse struct for EnvironmentApplicationsStorageResponse
type EnvironmentApplicationsStorageResponse struct {
	Application string `json:"application"`
	Disks []StorageDiskResponse `json:"disks,omitempty"`
}

// NewEnvironmentApplicationsStorageResponse instantiates a new EnvironmentApplicationsStorageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentApplicationsStorageResponse(application string) *EnvironmentApplicationsStorageResponse {
	this := EnvironmentApplicationsStorageResponse{}
	this.Application = application
	return &this
}

// NewEnvironmentApplicationsStorageResponseWithDefaults instantiates a new EnvironmentApplicationsStorageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentApplicationsStorageResponseWithDefaults() *EnvironmentApplicationsStorageResponse {
	this := EnvironmentApplicationsStorageResponse{}
	return &this
}

// GetApplication returns the Application field value
func (o *EnvironmentApplicationsStorageResponse) GetApplication() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsStorageResponse) GetApplicationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *EnvironmentApplicationsStorageResponse) SetApplication(v string) {
	o.Application = v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *EnvironmentApplicationsStorageResponse) GetDisks() []StorageDiskResponse {
	if o == nil || o.Disks == nil {
		var ret []StorageDiskResponse
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsStorageResponse) GetDisksOk() ([]StorageDiskResponse, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *EnvironmentApplicationsStorageResponse) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []StorageDiskResponse and assigns it to the Disks field.
func (o *EnvironmentApplicationsStorageResponse) SetDisks(v []StorageDiskResponse) {
	o.Disks = v
}

func (o EnvironmentApplicationsStorageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["application"] = o.Application
	}
	if o.Disks != nil {
		toSerialize["disks"] = o.Disks
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentApplicationsStorageResponse struct {
	value *EnvironmentApplicationsStorageResponse
	isSet bool
}

func (v NullableEnvironmentApplicationsStorageResponse) Get() *EnvironmentApplicationsStorageResponse {
	return v.value
}

func (v *NullableEnvironmentApplicationsStorageResponse) Set(val *EnvironmentApplicationsStorageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentApplicationsStorageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentApplicationsStorageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentApplicationsStorageResponse(val *EnvironmentApplicationsStorageResponse) *NullableEnvironmentApplicationsStorageResponse {
	return &NullableEnvironmentApplicationsStorageResponse{value: val, isSet: true}
}

func (v NullableEnvironmentApplicationsStorageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentApplicationsStorageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


