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
)

// OrganizationRequest struct for OrganizationRequest
type OrganizationRequest struct {
	// name is case insensitive
	Name        string         `json:"name"`
	Description *string        `json:"description,omitempty"`
	Plan        string         `json:"plan"`
	WebsiteUrl  NullableString `json:"website_url,omitempty"`
	Repository  NullableString `json:"repository,omitempty"`
	LogoUrl     NullableString `json:"logo_url,omitempty"`
	IconUrl     NullableString `json:"icon_url,omitempty"`
}

// NewOrganizationRequest instantiates a new OrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationRequest(name string, plan string) *OrganizationRequest {
	this := OrganizationRequest{}
	this.Name = name
	this.Plan = plan
	return &this
}

// NewOrganizationRequestWithDefaults instantiates a new OrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationRequestWithDefaults() *OrganizationRequest {
	this := OrganizationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *OrganizationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPlan returns the Plan field value
func (o *OrganizationRequest) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *OrganizationRequest) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *OrganizationRequest) SetPlan(v string) {
	o.Plan = v
}

// GetWebsiteUrl returns the WebsiteUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationRequest) GetWebsiteUrl() string {
	if o == nil || o.WebsiteUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebsiteUrl.Get()
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationRequest) GetWebsiteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebsiteUrl.Get(), o.WebsiteUrl.IsSet()
}

// HasWebsiteUrl returns a boolean if a field has been set.
func (o *OrganizationRequest) HasWebsiteUrl() bool {
	if o != nil && o.WebsiteUrl.IsSet() {
		return true
	}

	return false
}

// SetWebsiteUrl gets a reference to the given NullableString and assigns it to the WebsiteUrl field.
func (o *OrganizationRequest) SetWebsiteUrl(v string) {
	o.WebsiteUrl.Set(&v)
}

// SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil
func (o *OrganizationRequest) SetWebsiteUrlNil() {
	o.WebsiteUrl.Set(nil)
}

// UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
func (o *OrganizationRequest) UnsetWebsiteUrl() {
	o.WebsiteUrl.Unset()
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationRequest) GetRepository() string {
	if o == nil || o.Repository.Get() == nil {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationRequest) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *OrganizationRequest) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *OrganizationRequest) SetRepository(v string) {
	o.Repository.Set(&v)
}

// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *OrganizationRequest) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *OrganizationRequest) UnsetRepository() {
	o.Repository.Unset()
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationRequest) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationRequest) GetLogoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *OrganizationRequest) HasLogoUrl() bool {
	if o != nil && o.LogoUrl.IsSet() {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given NullableString and assigns it to the LogoUrl field.
func (o *OrganizationRequest) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}

// SetLogoUrlNil sets the value for LogoUrl to be an explicit nil
func (o *OrganizationRequest) SetLogoUrlNil() {
	o.LogoUrl.Set(nil)
}

// UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
func (o *OrganizationRequest) UnsetLogoUrl() {
	o.LogoUrl.Unset()
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationRequest) GetIconUrl() string {
	if o == nil || o.IconUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationRequest) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *OrganizationRequest) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *OrganizationRequest) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}

// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *OrganizationRequest) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *OrganizationRequest) UnsetIconUrl() {
	o.IconUrl.Unset()
}

func (o OrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableOrganizationRequest struct {
	value *OrganizationRequest
	isSet bool
}

func (v NullableOrganizationRequest) Get() *OrganizationRequest {
	return v.value
}

func (v *NullableOrganizationRequest) Set(val *OrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationRequest(val *OrganizationRequest) *NullableOrganizationRequest {
	return &NullableOrganizationRequest{value: val, isSet: true}
}

func (v NullableOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
