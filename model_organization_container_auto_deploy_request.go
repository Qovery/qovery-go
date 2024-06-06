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

// checks if the OrganizationContainerAutoDeployRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationContainerAutoDeployRequest{}

// OrganizationContainerAutoDeployRequest struct for OrganizationContainerAutoDeployRequest
type OrganizationContainerAutoDeployRequest struct {
	// the container image name to deploy
	ImageName *string `json:"image_name,omitempty"`
	// the new tag to deploy
	Tag                  *string `json:"tag,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationContainerAutoDeployRequest OrganizationContainerAutoDeployRequest

// NewOrganizationContainerAutoDeployRequest instantiates a new OrganizationContainerAutoDeployRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationContainerAutoDeployRequest() *OrganizationContainerAutoDeployRequest {
	this := OrganizationContainerAutoDeployRequest{}
	return &this
}

// NewOrganizationContainerAutoDeployRequestWithDefaults instantiates a new OrganizationContainerAutoDeployRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationContainerAutoDeployRequestWithDefaults() *OrganizationContainerAutoDeployRequest {
	this := OrganizationContainerAutoDeployRequest{}
	return &this
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *OrganizationContainerAutoDeployRequest) GetImageName() string {
	if o == nil || IsNil(o.ImageName) {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationContainerAutoDeployRequest) GetImageNameOk() (*string, bool) {
	if o == nil || IsNil(o.ImageName) {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *OrganizationContainerAutoDeployRequest) HasImageName() bool {
	if o != nil && !IsNil(o.ImageName) {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *OrganizationContainerAutoDeployRequest) SetImageName(v string) {
	o.ImageName = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *OrganizationContainerAutoDeployRequest) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationContainerAutoDeployRequest) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *OrganizationContainerAutoDeployRequest) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *OrganizationContainerAutoDeployRequest) SetTag(v string) {
	o.Tag = &v
}

func (o OrganizationContainerAutoDeployRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationContainerAutoDeployRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageName) {
		toSerialize["image_name"] = o.ImageName
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationContainerAutoDeployRequest) UnmarshalJSON(data []byte) (err error) {
	varOrganizationContainerAutoDeployRequest := _OrganizationContainerAutoDeployRequest{}

	err = json.Unmarshal(data, &varOrganizationContainerAutoDeployRequest)

	if err != nil {
		return err
	}

	*o = OrganizationContainerAutoDeployRequest(varOrganizationContainerAutoDeployRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "image_name")
		delete(additionalProperties, "tag")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationContainerAutoDeployRequest struct {
	value *OrganizationContainerAutoDeployRequest
	isSet bool
}

func (v NullableOrganizationContainerAutoDeployRequest) Get() *OrganizationContainerAutoDeployRequest {
	return v.value
}

func (v *NullableOrganizationContainerAutoDeployRequest) Set(val *OrganizationContainerAutoDeployRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationContainerAutoDeployRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationContainerAutoDeployRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationContainerAutoDeployRequest(val *OrganizationContainerAutoDeployRequest) *NullableOrganizationContainerAutoDeployRequest {
	return &NullableOrganizationContainerAutoDeployRequest{value: val, isSet: true}
}

func (v NullableOrganizationContainerAutoDeployRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationContainerAutoDeployRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
