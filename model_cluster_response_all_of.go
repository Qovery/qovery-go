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

// ClusterResponseAllOf struct for ClusterResponseAllOf
type ClusterResponseAllOf struct {
	// This is an estimation of the cost this cluster will represent on your cloud proider bill, based on your current configuration
	EstimatedCloudProviderCost *int32             `json:"estimated_cloud_provider_cost,omitempty"`
	Status                     *ClusterStatusEnum `json:"status,omitempty"`
	HasAccess                  *bool              `json:"has_access,omitempty"`
	Version                    *string            `json:"version,omitempty"`
	IsDefault                  *bool              `json:"is_default,omitempty"`
}

// NewClusterResponseAllOf instantiates a new ClusterResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterResponseAllOf() *ClusterResponseAllOf {
	this := ClusterResponseAllOf{}
	return &this
}

// NewClusterResponseAllOfWithDefaults instantiates a new ClusterResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterResponseAllOfWithDefaults() *ClusterResponseAllOf {
	this := ClusterResponseAllOf{}
	return &this
}

// GetEstimatedCloudProviderCost returns the EstimatedCloudProviderCost field value if set, zero value otherwise.
func (o *ClusterResponseAllOf) GetEstimatedCloudProviderCost() int32 {
	if o == nil || o.EstimatedCloudProviderCost == nil {
		var ret int32
		return ret
	}
	return *o.EstimatedCloudProviderCost
}

// GetEstimatedCloudProviderCostOk returns a tuple with the EstimatedCloudProviderCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponseAllOf) GetEstimatedCloudProviderCostOk() (*int32, bool) {
	if o == nil || o.EstimatedCloudProviderCost == nil {
		return nil, false
	}
	return o.EstimatedCloudProviderCost, true
}

// HasEstimatedCloudProviderCost returns a boolean if a field has been set.
func (o *ClusterResponseAllOf) HasEstimatedCloudProviderCost() bool {
	if o != nil && o.EstimatedCloudProviderCost != nil {
		return true
	}

	return false
}

// SetEstimatedCloudProviderCost gets a reference to the given int32 and assigns it to the EstimatedCloudProviderCost field.
func (o *ClusterResponseAllOf) SetEstimatedCloudProviderCost(v int32) {
	o.EstimatedCloudProviderCost = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterResponseAllOf) GetStatus() ClusterStatusEnum {
	if o == nil || o.Status == nil {
		var ret ClusterStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponseAllOf) GetStatusOk() (*ClusterStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterResponseAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ClusterStatusEnum and assigns it to the Status field.
func (o *ClusterResponseAllOf) SetStatus(v ClusterStatusEnum) {
	o.Status = &v
}

// GetHasAccess returns the HasAccess field value if set, zero value otherwise.
func (o *ClusterResponseAllOf) GetHasAccess() bool {
	if o == nil || o.HasAccess == nil {
		var ret bool
		return ret
	}
	return *o.HasAccess
}

// GetHasAccessOk returns a tuple with the HasAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponseAllOf) GetHasAccessOk() (*bool, bool) {
	if o == nil || o.HasAccess == nil {
		return nil, false
	}
	return o.HasAccess, true
}

// HasHasAccess returns a boolean if a field has been set.
func (o *ClusterResponseAllOf) HasHasAccess() bool {
	if o != nil && o.HasAccess != nil {
		return true
	}

	return false
}

// SetHasAccess gets a reference to the given bool and assigns it to the HasAccess field.
func (o *ClusterResponseAllOf) SetHasAccess(v bool) {
	o.HasAccess = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ClusterResponseAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponseAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ClusterResponseAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ClusterResponseAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ClusterResponseAllOf) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponseAllOf) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ClusterResponseAllOf) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ClusterResponseAllOf) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o ClusterResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EstimatedCloudProviderCost != nil {
		toSerialize["estimated_cloud_provider_cost"] = o.EstimatedCloudProviderCost
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.HasAccess != nil {
		toSerialize["has_access"] = o.HasAccess
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.IsDefault != nil {
		toSerialize["is_default"] = o.IsDefault
	}
	return json.Marshal(toSerialize)
}

type NullableClusterResponseAllOf struct {
	value *ClusterResponseAllOf
	isSet bool
}

func (v NullableClusterResponseAllOf) Get() *ClusterResponseAllOf {
	return v.value
}

func (v *NullableClusterResponseAllOf) Set(val *ClusterResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterResponseAllOf(val *ClusterResponseAllOf) *NullableClusterResponseAllOf {
	return &NullableClusterResponseAllOf{value: val, isSet: true}
}

func (v NullableClusterResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
