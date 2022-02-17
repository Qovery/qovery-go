/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet. 

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// ClusterFeatureRequest struct for ClusterFeatureRequest
type ClusterFeatureRequest struct {
	Features []ClusterFeatureRequestFeatures `json:"features,omitempty"`
}

// NewClusterFeatureRequest instantiates a new ClusterFeatureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterFeatureRequest() *ClusterFeatureRequest {
	this := ClusterFeatureRequest{}
	return &this
}

// NewClusterFeatureRequestWithDefaults instantiates a new ClusterFeatureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterFeatureRequestWithDefaults() *ClusterFeatureRequest {
	this := ClusterFeatureRequest{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ClusterFeatureRequest) GetFeatures() []ClusterFeatureRequestFeatures {
	if o == nil || o.Features == nil {
		var ret []ClusterFeatureRequestFeatures
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeatureRequest) GetFeaturesOk() ([]ClusterFeatureRequestFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ClusterFeatureRequest) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []ClusterFeatureRequestFeatures and assigns it to the Features field.
func (o *ClusterFeatureRequest) SetFeatures(v []ClusterFeatureRequestFeatures) {
	o.Features = v
}

func (o ClusterFeatureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	return json.Marshal(toSerialize)
}

type NullableClusterFeatureRequest struct {
	value *ClusterFeatureRequest
	isSet bool
}

func (v NullableClusterFeatureRequest) Get() *ClusterFeatureRequest {
	return v.value
}

func (v *NullableClusterFeatureRequest) Set(val *ClusterFeatureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterFeatureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterFeatureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterFeatureRequest(val *ClusterFeatureRequest) *NullableClusterFeatureRequest {
	return &NullableClusterFeatureRequest{value: val, isSet: true}
}

func (v NullableClusterFeatureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterFeatureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


