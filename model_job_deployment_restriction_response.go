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
	"fmt"
	"time"
)

// checks if the JobDeploymentRestrictionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobDeploymentRestrictionResponse{}

// JobDeploymentRestrictionResponse struct for JobDeploymentRestrictionResponse
type JobDeploymentRestrictionResponse struct {
	Id        string                        `json:"id"`
	CreatedAt time.Time                     `json:"created_at"`
	UpdatedAt *time.Time                    `json:"updated_at,omitempty"`
	Mode      DeploymentRestrictionModeEnum `json:"mode"`
	Type      DeploymentRestrictionTypeEnum `json:"type"`
	// For `PATH` restrictions, the value must not start with `/`
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _JobDeploymentRestrictionResponse JobDeploymentRestrictionResponse

// NewJobDeploymentRestrictionResponse instantiates a new JobDeploymentRestrictionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDeploymentRestrictionResponse(id string, createdAt time.Time, mode DeploymentRestrictionModeEnum, type_ DeploymentRestrictionTypeEnum, value string) *JobDeploymentRestrictionResponse {
	this := JobDeploymentRestrictionResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Mode = mode
	this.Type = type_
	this.Value = value
	return &this
}

// NewJobDeploymentRestrictionResponseWithDefaults instantiates a new JobDeploymentRestrictionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDeploymentRestrictionResponseWithDefaults() *JobDeploymentRestrictionResponse {
	this := JobDeploymentRestrictionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *JobDeploymentRestrictionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobDeploymentRestrictionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobDeploymentRestrictionResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *JobDeploymentRestrictionResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *JobDeploymentRestrictionResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *JobDeploymentRestrictionResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *JobDeploymentRestrictionResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDeploymentRestrictionResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *JobDeploymentRestrictionResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *JobDeploymentRestrictionResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetMode returns the Mode field value
func (o *JobDeploymentRestrictionResponse) GetMode() DeploymentRestrictionModeEnum {
	if o == nil {
		var ret DeploymentRestrictionModeEnum
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *JobDeploymentRestrictionResponse) GetModeOk() (*DeploymentRestrictionModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *JobDeploymentRestrictionResponse) SetMode(v DeploymentRestrictionModeEnum) {
	o.Mode = v
}

// GetType returns the Type field value
func (o *JobDeploymentRestrictionResponse) GetType() DeploymentRestrictionTypeEnum {
	if o == nil {
		var ret DeploymentRestrictionTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JobDeploymentRestrictionResponse) GetTypeOk() (*DeploymentRestrictionTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *JobDeploymentRestrictionResponse) SetType(v DeploymentRestrictionTypeEnum) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *JobDeploymentRestrictionResponse) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *JobDeploymentRestrictionResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *JobDeploymentRestrictionResponse) SetValue(v string) {
	o.Value = v
}

func (o JobDeploymentRestrictionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobDeploymentRestrictionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["mode"] = o.Mode
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JobDeploymentRestrictionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"mode",
		"type",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varJobDeploymentRestrictionResponse := _JobDeploymentRestrictionResponse{}

	err = json.Unmarshal(data, &varJobDeploymentRestrictionResponse)

	if err != nil {
		return err
	}

	*o = JobDeploymentRestrictionResponse(varJobDeploymentRestrictionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJobDeploymentRestrictionResponse struct {
	value *JobDeploymentRestrictionResponse
	isSet bool
}

func (v NullableJobDeploymentRestrictionResponse) Get() *JobDeploymentRestrictionResponse {
	return v.value
}

func (v *NullableJobDeploymentRestrictionResponse) Set(val *JobDeploymentRestrictionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDeploymentRestrictionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDeploymentRestrictionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDeploymentRestrictionResponse(val *JobDeploymentRestrictionResponse) *NullableJobDeploymentRestrictionResponse {
	return &NullableJobDeploymentRestrictionResponse{value: val, isSet: true}
}

func (v NullableJobDeploymentRestrictionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDeploymentRestrictionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
