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

// checks if the OrganizationApiTokenCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationApiTokenCreate{}

// OrganizationApiTokenCreate struct for OrganizationApiTokenCreate
type OrganizationApiTokenCreate struct {
	Id          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	// the generated token to send in 'Authorization' header prefixed by 'Token '
	Token    *string `json:"token,omitempty"`
	RoleName *string `json:"role_name,omitempty"`
	RoleId   *string `json:"role_id,omitempty"`
}

// NewOrganizationApiTokenCreate instantiates a new OrganizationApiTokenCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationApiTokenCreate(id string, createdAt time.Time) *OrganizationApiTokenCreate {
	this := OrganizationApiTokenCreate{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewOrganizationApiTokenCreateWithDefaults instantiates a new OrganizationApiTokenCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationApiTokenCreateWithDefaults() *OrganizationApiTokenCreate {
	this := OrganizationApiTokenCreate{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationApiTokenCreate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationApiTokenCreate) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationApiTokenCreate) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationApiTokenCreate) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrganizationApiTokenCreate) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreate) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrganizationApiTokenCreate) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OrganizationApiTokenCreate) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationApiTokenCreate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationApiTokenCreate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationApiTokenCreate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationApiTokenCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationApiTokenCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationApiTokenCreate) SetDescription(v string) {
	o.Description = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *OrganizationApiTokenCreate) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreate) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *OrganizationApiTokenCreate) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *OrganizationApiTokenCreate) SetToken(v string) {
	o.Token = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *OrganizationApiTokenCreate) GetRoleName() string {
	if o == nil || IsNil(o.RoleName) {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreate) GetRoleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleName) {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *OrganizationApiTokenCreate) HasRoleName() bool {
	if o != nil && !IsNil(o.RoleName) {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *OrganizationApiTokenCreate) SetRoleName(v string) {
	o.RoleName = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *OrganizationApiTokenCreate) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreate) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *OrganizationApiTokenCreate) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *OrganizationApiTokenCreate) SetRoleId(v string) {
	o.RoleId = &v
}

func (o OrganizationApiTokenCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationApiTokenCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.RoleName) {
		toSerialize["role_name"] = o.RoleName
	}
	if !IsNil(o.RoleId) {
		toSerialize["role_id"] = o.RoleId
	}
	return toSerialize, nil
}

type NullableOrganizationApiTokenCreate struct {
	value *OrganizationApiTokenCreate
	isSet bool
}

func (v NullableOrganizationApiTokenCreate) Get() *OrganizationApiTokenCreate {
	return v.value
}

func (v *NullableOrganizationApiTokenCreate) Set(val *OrganizationApiTokenCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationApiTokenCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationApiTokenCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationApiTokenCreate(val *OrganizationApiTokenCreate) *NullableOrganizationApiTokenCreate {
	return &NullableOrganizationApiTokenCreate{value: val, isSet: true}
}

func (v NullableOrganizationApiTokenCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationApiTokenCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
