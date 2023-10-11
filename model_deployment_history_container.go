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

// checks if the DeploymentHistoryContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentHistoryContainer{}

// DeploymentHistoryContainer struct for DeploymentHistoryContainer
type DeploymentHistoryContainer struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// name of the container
	Name       *string    `json:"name,omitempty"`
	Status     *StateEnum `json:"status,omitempty"`
	ImageName  *string    `json:"image_name,omitempty"`
	Tag        *string    `json:"tag,omitempty"`
	Arguments  []string   `json:"arguments,omitempty"`
	Entrypoint *string    `json:"entrypoint,omitempty"`
}

// NewDeploymentHistoryContainer instantiates a new DeploymentHistoryContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryContainer(id string, createdAt time.Time) *DeploymentHistoryContainer {
	this := DeploymentHistoryContainer{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewDeploymentHistoryContainerWithDefaults instantiates a new DeploymentHistoryContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryContainerWithDefaults() *DeploymentHistoryContainer {
	this := DeploymentHistoryContainer{}
	return &this
}

// GetId returns the Id field value
func (o *DeploymentHistoryContainer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryContainer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeploymentHistoryContainer) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeploymentHistoryContainer) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryContainer) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeploymentHistoryContainer) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeploymentHistoryContainer) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryContainer) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeploymentHistoryContainer) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeploymentHistoryContainer) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentHistoryContainer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryContainer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentHistoryContainer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentHistoryContainer) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentHistoryContainer) GetStatus() StateEnum {
	if o == nil || IsNil(o.Status) {
		var ret StateEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryContainer) GetStatusOk() (*StateEnum, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentHistoryContainer) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StateEnum and assigns it to the Status field.
func (o *DeploymentHistoryContainer) SetStatus(v StateEnum) {
	o.Status = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *DeploymentHistoryContainer) GetImageName() string {
	if o == nil || IsNil(o.ImageName) {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryContainer) GetImageNameOk() (*string, bool) {
	if o == nil || IsNil(o.ImageName) {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *DeploymentHistoryContainer) HasImageName() bool {
	if o != nil && !IsNil(o.ImageName) {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *DeploymentHistoryContainer) SetImageName(v string) {
	o.ImageName = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *DeploymentHistoryContainer) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryContainer) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *DeploymentHistoryContainer) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *DeploymentHistoryContainer) SetTag(v string) {
	o.Tag = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *DeploymentHistoryContainer) GetArguments() []string {
	if o == nil || IsNil(o.Arguments) {
		var ret []string
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryContainer) GetArgumentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Arguments) {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *DeploymentHistoryContainer) HasArguments() bool {
	if o != nil && !IsNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []string and assigns it to the Arguments field.
func (o *DeploymentHistoryContainer) SetArguments(v []string) {
	o.Arguments = v
}

// GetEntrypoint returns the Entrypoint field value if set, zero value otherwise.
func (o *DeploymentHistoryContainer) GetEntrypoint() string {
	if o == nil || IsNil(o.Entrypoint) {
		var ret string
		return ret
	}
	return *o.Entrypoint
}

// GetEntrypointOk returns a tuple with the Entrypoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryContainer) GetEntrypointOk() (*string, bool) {
	if o == nil || IsNil(o.Entrypoint) {
		return nil, false
	}
	return o.Entrypoint, true
}

// HasEntrypoint returns a boolean if a field has been set.
func (o *DeploymentHistoryContainer) HasEntrypoint() bool {
	if o != nil && !IsNil(o.Entrypoint) {
		return true
	}

	return false
}

// SetEntrypoint gets a reference to the given string and assigns it to the Entrypoint field.
func (o *DeploymentHistoryContainer) SetEntrypoint(v string) {
	o.Entrypoint = &v
}

func (o DeploymentHistoryContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentHistoryContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ImageName) {
		toSerialize["image_name"] = o.ImageName
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}
	if !IsNil(o.Entrypoint) {
		toSerialize["entrypoint"] = o.Entrypoint
	}
	return toSerialize, nil
}

type NullableDeploymentHistoryContainer struct {
	value *DeploymentHistoryContainer
	isSet bool
}

func (v NullableDeploymentHistoryContainer) Get() *DeploymentHistoryContainer {
	return v.value
}

func (v *NullableDeploymentHistoryContainer) Set(val *DeploymentHistoryContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryContainer(val *DeploymentHistoryContainer) *NullableDeploymentHistoryContainer {
	return &NullableDeploymentHistoryContainer{value: val, isSet: true}
}

func (v NullableDeploymentHistoryContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
