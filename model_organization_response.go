/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"time"
)

// OrganizationResponse struct for OrganizationResponse
type OrganizationResponse struct {
	// uuid of the user owning the organization
	Owner     *string    `json:"owner,omitempty"`
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// name is case insensitive
	Name        string         `json:"name"`
	Description *string        `json:"description,omitempty"`
	Plan        string         `json:"plan"`
	WebsiteUrl  NullableString `json:"website_url,omitempty"`
	Repository  NullableString `json:"repository,omitempty"`
	LogoUrl     NullableString `json:"logo_url,omitempty"`
	IconUrl     NullableString `json:"icon_url,omitempty"`
}

// NewOrganizationResponse instantiates a new OrganizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationResponse(id string, createdAt time.Time, name string, plan string) *OrganizationResponse {
	this := OrganizationResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Name = name
	this.Plan = plan
	return &this
}

// NewOrganizationResponseWithDefaults instantiates a new OrganizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationResponseWithDefaults() *OrganizationResponse {
	this := OrganizationResponse{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *OrganizationResponse) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationResponse) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *OrganizationResponse) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *OrganizationResponse) SetOwner(v string) {
	o.Owner = &v
}

// GetId returns the Id field value
func (o *OrganizationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrganizationResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrganizationResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OrganizationResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value
func (o *OrganizationResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationResponse) SetDescription(v string) {
	o.Description = &v
}

// GetPlan returns the Plan field value
func (o *OrganizationResponse) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *OrganizationResponse) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *OrganizationResponse) SetPlan(v string) {
	o.Plan = v
}

// GetWebsiteUrl returns the WebsiteUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationResponse) GetWebsiteUrl() string {
	if o == nil || o.WebsiteUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebsiteUrl.Get()
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationResponse) GetWebsiteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebsiteUrl.Get(), o.WebsiteUrl.IsSet()
}

// HasWebsiteUrl returns a boolean if a field has been set.
func (o *OrganizationResponse) HasWebsiteUrl() bool {
	if o != nil && o.WebsiteUrl.IsSet() {
		return true
	}

	return false
}

// SetWebsiteUrl gets a reference to the given NullableString and assigns it to the WebsiteUrl field.
func (o *OrganizationResponse) SetWebsiteUrl(v string) {
	o.WebsiteUrl.Set(&v)
}

// SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil
func (o *OrganizationResponse) SetWebsiteUrlNil() {
	o.WebsiteUrl.Set(nil)
}

// UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
func (o *OrganizationResponse) UnsetWebsiteUrl() {
	o.WebsiteUrl.Unset()
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationResponse) GetRepository() string {
	if o == nil || o.Repository.Get() == nil {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationResponse) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *OrganizationResponse) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *OrganizationResponse) SetRepository(v string) {
	o.Repository.Set(&v)
}

// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *OrganizationResponse) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *OrganizationResponse) UnsetRepository() {
	o.Repository.Unset()
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationResponse) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationResponse) GetLogoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *OrganizationResponse) HasLogoUrl() bool {
	if o != nil && o.LogoUrl.IsSet() {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given NullableString and assigns it to the LogoUrl field.
func (o *OrganizationResponse) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}

// SetLogoUrlNil sets the value for LogoUrl to be an explicit nil
func (o *OrganizationResponse) SetLogoUrlNil() {
	o.LogoUrl.Set(nil)
}

// UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
func (o *OrganizationResponse) UnsetLogoUrl() {
	o.LogoUrl.Unset()
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationResponse) GetIconUrl() string {
	if o == nil || o.IconUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationResponse) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *OrganizationResponse) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *OrganizationResponse) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}

// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *OrganizationResponse) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *OrganizationResponse) UnsetIconUrl() {
	o.IconUrl.Unset()
}

func (o OrganizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	if o.WebsiteUrl.IsSet() {
		toSerialize["website_url"] = o.WebsiteUrl.Get()
	}
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if o.LogoUrl.IsSet() {
		toSerialize["logo_url"] = o.LogoUrl.Get()
	}
	if o.IconUrl.IsSet() {
		toSerialize["icon_url"] = o.IconUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationResponse struct {
	value *OrganizationResponse
	isSet bool
}

func (v NullableOrganizationResponse) Get() *OrganizationResponse {
	return v.value
}

func (v *NullableOrganizationResponse) Set(val *OrganizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationResponse(val *OrganizationResponse) *NullableOrganizationResponse {
	return &NullableOrganizationResponse{value: val, isSet: true}
}

func (v NullableOrganizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
