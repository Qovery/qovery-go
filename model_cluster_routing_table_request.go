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

// ClusterRoutingTableRequest struct for ClusterRoutingTableRequest
type ClusterRoutingTableRequest struct {
	Routes []ClusterRoutingTableRequestRoutes `json:"routes"`
}

// NewClusterRoutingTableRequest instantiates a new ClusterRoutingTableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRoutingTableRequest(routes []ClusterRoutingTableRequestRoutes) *ClusterRoutingTableRequest {
	this := ClusterRoutingTableRequest{}
	this.Routes = routes
	return &this
}

// NewClusterRoutingTableRequestWithDefaults instantiates a new ClusterRoutingTableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterRoutingTableRequestWithDefaults() *ClusterRoutingTableRequest {
	this := ClusterRoutingTableRequest{}
	return &this
}

// GetRoutes returns the Routes field value
func (o *ClusterRoutingTableRequest) GetRoutes() []ClusterRoutingTableRequestRoutes {
	if o == nil {
		var ret []ClusterRoutingTableRequestRoutes
		return ret
	}

	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value
// and a boolean to check if the value has been set.
func (o *ClusterRoutingTableRequest) GetRoutesOk() ([]ClusterRoutingTableRequestRoutes, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Routes, true
}

// SetRoutes sets field value
func (o *ClusterRoutingTableRequest) SetRoutes(v []ClusterRoutingTableRequestRoutes) {
	o.Routes = v
}

func (o ClusterRoutingTableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["routes"] = o.Routes
	}
	return json.Marshal(toSerialize)
}

type NullableClusterRoutingTableRequest struct {
	value *ClusterRoutingTableRequest
	isSet bool
}

func (v NullableClusterRoutingTableRequest) Get() *ClusterRoutingTableRequest {
	return v.value
}

func (v *NullableClusterRoutingTableRequest) Set(val *ClusterRoutingTableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRoutingTableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRoutingTableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRoutingTableRequest(val *ClusterRoutingTableRequest) *NullableClusterRoutingTableRequest {
	return &NullableClusterRoutingTableRequest{value: val, isSet: true}
}

func (v NullableClusterRoutingTableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRoutingTableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


