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
)

// checks if the OrganizationLabelsGroupAssociatedItemsResponseListResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationLabelsGroupAssociatedItemsResponseListResultsInner{}

// OrganizationLabelsGroupAssociatedItemsResponseListResultsInner struct for OrganizationLabelsGroupAssociatedItemsResponseListResultsInner
type OrganizationLabelsGroupAssociatedItemsResponseListResultsInner struct {
	ClusterId            *string                       `json:"cluster_id,omitempty"`
	ClusterName          *string                       `json:"cluster_name,omitempty"`
	ProjectId            *string                       `json:"project_id,omitempty"`
	ProjectName          *string                       `json:"project_name,omitempty"`
	EnvironmentId        *string                       `json:"environment_id,omitempty"`
	EnvironmentName      *string                       `json:"environment_name,omitempty"`
	ItemId               string                        `json:"item_id"`
	ItemName             string                        `json:"item_name"`
	ItemType             LabelsGroupAssociatedItemType `json:"item_type"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationLabelsGroupAssociatedItemsResponseListResultsInner OrganizationLabelsGroupAssociatedItemsResponseListResultsInner

// NewOrganizationLabelsGroupAssociatedItemsResponseListResultsInner instantiates a new OrganizationLabelsGroupAssociatedItemsResponseListResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationLabelsGroupAssociatedItemsResponseListResultsInner(itemId string, itemName string, itemType LabelsGroupAssociatedItemType) *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner {
	this := OrganizationLabelsGroupAssociatedItemsResponseListResultsInner{}
	this.ItemId = itemId
	this.ItemName = itemName
	this.ItemType = itemType
	return &this
}

// NewOrganizationLabelsGroupAssociatedItemsResponseListResultsInnerWithDefaults instantiates a new OrganizationLabelsGroupAssociatedItemsResponseListResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationLabelsGroupAssociatedItemsResponseListResultsInnerWithDefaults() *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner {
	this := OrganizationLabelsGroupAssociatedItemsResponseListResultsInner{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetProjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) HasProjectName() bool {
	if o != nil && !IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetEnvironmentName() string {
	if o == nil || IsNil(o.EnvironmentName) {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentName) {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) HasEnvironmentName() bool {
	if o != nil && !IsNil(o.EnvironmentName) {
		return true
	}

	return false
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetItemId returns the ItemId field value
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) SetItemId(v string) {
	o.ItemId = v
}

// GetItemName returns the ItemName field value
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetItemName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemName
}

// GetItemNameOk returns a tuple with the ItemName field value
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetItemNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemName, true
}

// SetItemName sets field value
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) SetItemName(v string) {
	o.ItemName = v
}

// GetItemType returns the ItemType field value
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetItemType() LabelsGroupAssociatedItemType {
	if o == nil {
		var ret LabelsGroupAssociatedItemType
		return ret
	}

	return o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) GetItemTypeOk() (*LabelsGroupAssociatedItemType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemType, true
}

// SetItemType sets field value
func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) SetItemType(v LabelsGroupAssociatedItemType) {
	o.ItemType = v
}

func (o OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterId) {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if !IsNil(o.ClusterName) {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.ProjectName) {
		toSerialize["project_name"] = o.ProjectName
	}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if !IsNil(o.EnvironmentName) {
		toSerialize["environment_name"] = o.EnvironmentName
	}
	toSerialize["item_id"] = o.ItemId
	toSerialize["item_name"] = o.ItemName
	toSerialize["item_type"] = o.ItemType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"item_id",
		"item_name",
		"item_type",
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

	varOrganizationLabelsGroupAssociatedItemsResponseListResultsInner := _OrganizationLabelsGroupAssociatedItemsResponseListResultsInner{}

	err = json.Unmarshal(data, &varOrganizationLabelsGroupAssociatedItemsResponseListResultsInner)

	if err != nil {
		return err
	}

	*o = OrganizationLabelsGroupAssociatedItemsResponseListResultsInner(varOrganizationLabelsGroupAssociatedItemsResponseListResultsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cluster_id")
		delete(additionalProperties, "cluster_name")
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "project_name")
		delete(additionalProperties, "environment_id")
		delete(additionalProperties, "environment_name")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "item_name")
		delete(additionalProperties, "item_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationLabelsGroupAssociatedItemsResponseListResultsInner struct {
	value *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner
	isSet bool
}

func (v NullableOrganizationLabelsGroupAssociatedItemsResponseListResultsInner) Get() *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner {
	return v.value
}

func (v *NullableOrganizationLabelsGroupAssociatedItemsResponseListResultsInner) Set(val *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationLabelsGroupAssociatedItemsResponseListResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationLabelsGroupAssociatedItemsResponseListResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationLabelsGroupAssociatedItemsResponseListResultsInner(val *OrganizationLabelsGroupAssociatedItemsResponseListResultsInner) *NullableOrganizationLabelsGroupAssociatedItemsResponseListResultsInner {
	return &NullableOrganizationLabelsGroupAssociatedItemsResponseListResultsInner{value: val, isSet: true}
}

func (v NullableOrganizationLabelsGroupAssociatedItemsResponseListResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationLabelsGroupAssociatedItemsResponseListResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
