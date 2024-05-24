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

// checks if the OrganizationLabelsGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationLabelsGroupResponse{}

// OrganizationLabelsGroupResponse struct for OrganizationLabelsGroupResponse
type OrganizationLabelsGroupResponse struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Name      string     `json:"name"`
	Labels    []Label    `json:"labels"`
}

// NewOrganizationLabelsGroupResponse instantiates a new OrganizationLabelsGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationLabelsGroupResponse(id string, createdAt time.Time, name string, labels []Label) *OrganizationLabelsGroupResponse {
	this := OrganizationLabelsGroupResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Name = name
	this.Labels = labels
	return &this
}

// NewOrganizationLabelsGroupResponseWithDefaults instantiates a new OrganizationLabelsGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationLabelsGroupResponseWithDefaults() *OrganizationLabelsGroupResponse {
	this := OrganizationLabelsGroupResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationLabelsGroupResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationLabelsGroupResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationLabelsGroupResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationLabelsGroupResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrganizationLabelsGroupResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrganizationLabelsGroupResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OrganizationLabelsGroupResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value
func (o *OrganizationLabelsGroupResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationLabelsGroupResponse) SetName(v string) {
	o.Name = v
}

// GetLabels returns the Labels field value
func (o *OrganizationLabelsGroupResponse) GetLabels() []Label {
	if o == nil {
		var ret []Label
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupResponse) GetLabelsOk() ([]Label, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *OrganizationLabelsGroupResponse) SetLabels(v []Label) {
	o.Labels = v
}

func (o OrganizationLabelsGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationLabelsGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["name"] = o.Name
	toSerialize["labels"] = o.Labels
	return toSerialize, nil
}

type NullableOrganizationLabelsGroupResponse struct {
	value *OrganizationLabelsGroupResponse
	isSet bool
}

func (v NullableOrganizationLabelsGroupResponse) Get() *OrganizationLabelsGroupResponse {
	return v.value
}

func (v *NullableOrganizationLabelsGroupResponse) Set(val *OrganizationLabelsGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationLabelsGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationLabelsGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationLabelsGroupResponse(val *OrganizationLabelsGroupResponse) *NullableOrganizationLabelsGroupResponse {
	return &NullableOrganizationLabelsGroupResponse{value: val, isSet: true}
}

func (v NullableOrganizationLabelsGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationLabelsGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
