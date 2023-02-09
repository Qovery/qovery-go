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
	"time"
)

// DeploymentStageServiceResponse struct for DeploymentStageServiceResponse
type DeploymentStageServiceResponse struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// id of the service attached to the stage
	ServiceId *string `json:"service_id,omitempty"`
	// type of the service (i.e APPLICATION, JOB, DATABASE, ...)
	ServiceType *string `json:"service_type,omitempty"`
}

// NewDeploymentStageServiceResponse instantiates a new DeploymentStageServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStageServiceResponse(id string, createdAt time.Time) *DeploymentStageServiceResponse {
	this := DeploymentStageServiceResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewDeploymentStageServiceResponseWithDefaults instantiates a new DeploymentStageServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStageServiceResponseWithDefaults() *DeploymentStageServiceResponse {
	this := DeploymentStageServiceResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DeploymentStageServiceResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploymentStageServiceResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeploymentStageServiceResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeploymentStageServiceResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentStageServiceResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeploymentStageServiceResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeploymentStageServiceResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageServiceResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeploymentStageServiceResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeploymentStageServiceResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *DeploymentStageServiceResponse) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageServiceResponse) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *DeploymentStageServiceResponse) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *DeploymentStageServiceResponse) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *DeploymentStageServiceResponse) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageServiceResponse) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *DeploymentStageServiceResponse) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *DeploymentStageServiceResponse) SetServiceType(v string) {
	o.ServiceType = &v
}

func (o DeploymentStageServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.ServiceType != nil {
		toSerialize["service_type"] = o.ServiceType
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentStageServiceResponse struct {
	value *DeploymentStageServiceResponse
	isSet bool
}

func (v NullableDeploymentStageServiceResponse) Get() *DeploymentStageServiceResponse {
	return v.value
}

func (v *NullableDeploymentStageServiceResponse) Set(val *DeploymentStageServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStageServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStageServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStageServiceResponse(val *DeploymentStageServiceResponse) *NullableDeploymentStageServiceResponse {
	return &NullableDeploymentStageServiceResponse{value: val, isSet: true}
}

func (v NullableDeploymentStageServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStageServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
