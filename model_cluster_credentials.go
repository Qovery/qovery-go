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

// checks if the ClusterCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterCredentials{}

// ClusterCredentials struct for ClusterCredentials
type ClusterCredentials struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// NewClusterCredentials instantiates a new ClusterCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterCredentials(id string, name string) *ClusterCredentials {
	this := ClusterCredentials{}
	this.Id = id
	this.Name = name
	return &this
}

// NewClusterCredentialsWithDefaults instantiates a new ClusterCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterCredentialsWithDefaults() *ClusterCredentials {
	this := ClusterCredentials{}
	return &this
}

// GetId returns the Id field value
func (o *ClusterCredentials) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClusterCredentials) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClusterCredentials) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ClusterCredentials) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterCredentials) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterCredentials) SetName(v string) {
	o.Name = v
}

func (o ClusterCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableClusterCredentials struct {
	value *ClusterCredentials
	isSet bool
}

func (v NullableClusterCredentials) Get() *ClusterCredentials {
	return v.value
}

func (v *NullableClusterCredentials) Set(val *ClusterCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterCredentials(val *ClusterCredentials) *NullableClusterCredentials {
	return &NullableClusterCredentials{value: val, isSet: true}
}

func (v NullableClusterCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
