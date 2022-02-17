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

// EnvironmentRequest struct for EnvironmentRequest
type EnvironmentRequest struct {
	// name is case insensitive
	Name string `json:"name"`
	Cluster *string `json:"cluster,omitempty"`
	Mode *string `json:"mode,omitempty"`
}

// NewEnvironmentRequest instantiates a new EnvironmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentRequest(name string) *EnvironmentRequest {
	this := EnvironmentRequest{}
	this.Name = name
	return &this
}

// NewEnvironmentRequestWithDefaults instantiates a new EnvironmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentRequestWithDefaults() *EnvironmentRequest {
	this := EnvironmentRequest{}
	return &this
}

// GetName returns the Name field value
func (o *EnvironmentRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentRequest) SetName(v string) {
	o.Name = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *EnvironmentRequest) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRequest) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *EnvironmentRequest) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *EnvironmentRequest) SetCluster(v string) {
	o.Cluster = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *EnvironmentRequest) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRequest) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *EnvironmentRequest) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *EnvironmentRequest) SetMode(v string) {
	o.Mode = &v
}

func (o EnvironmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentRequest struct {
	value *EnvironmentRequest
	isSet bool
}

func (v NullableEnvironmentRequest) Get() *EnvironmentRequest {
	return v.value
}

func (v *NullableEnvironmentRequest) Set(val *EnvironmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentRequest(val *EnvironmentRequest) *NullableEnvironmentRequest {
	return &NullableEnvironmentRequest{value: val, isSet: true}
}

func (v NullableEnvironmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


