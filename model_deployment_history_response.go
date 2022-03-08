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

// DeploymentHistoryResponse struct for DeploymentHistoryResponse
type DeploymentHistoryResponse struct {
	Id        string                       `json:"id"`
	CreatedAt time.Time                    `json:"created_at"`
	UpdatedAt *time.Time                   `json:"updated_at,omitempty"`
	Commit    *CommitResponse              `json:"commit,omitempty"`
	Status    *DeploymentHistoryStatusEnum `json:"status,omitempty"`
}

// NewDeploymentHistoryResponse instantiates a new DeploymentHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryResponse(id string, createdAt time.Time) *DeploymentHistoryResponse {
	this := DeploymentHistoryResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewDeploymentHistoryResponseWithDefaults instantiates a new DeploymentHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryResponseWithDefaults() *DeploymentHistoryResponse {
	this := DeploymentHistoryResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DeploymentHistoryResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeploymentHistoryResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeploymentHistoryResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeploymentHistoryResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeploymentHistoryResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeploymentHistoryResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeploymentHistoryResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *DeploymentHistoryResponse) GetCommit() CommitResponse {
	if o == nil || o.Commit == nil {
		var ret CommitResponse
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryResponse) GetCommitOk() (*CommitResponse, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *DeploymentHistoryResponse) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given CommitResponse and assigns it to the Commit field.
func (o *DeploymentHistoryResponse) SetCommit(v CommitResponse) {
	o.Commit = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentHistoryResponse) GetStatus() DeploymentHistoryStatusEnum {
	if o == nil || o.Status == nil {
		var ret DeploymentHistoryStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryResponse) GetStatusOk() (*DeploymentHistoryStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentHistoryResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeploymentHistoryStatusEnum and assigns it to the Status field.
func (o *DeploymentHistoryResponse) SetStatus(v DeploymentHistoryStatusEnum) {
	o.Status = &v
}

func (o DeploymentHistoryResponse) MarshalJSON() ([]byte, error) {
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
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentHistoryResponse struct {
	value *DeploymentHistoryResponse
	isSet bool
}

func (v NullableDeploymentHistoryResponse) Get() *DeploymentHistoryResponse {
	return v.value
}

func (v *NullableDeploymentHistoryResponse) Set(val *DeploymentHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryResponse(val *DeploymentHistoryResponse) *NullableDeploymentHistoryResponse {
	return &NullableDeploymentHistoryResponse{value: val, isSet: true}
}

func (v NullableDeploymentHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
