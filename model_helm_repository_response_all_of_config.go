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

// checks if the HelmRepositoryResponseAllOfConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmRepositoryResponseAllOfConfig{}

// HelmRepositoryResponseAllOfConfig struct for HelmRepositoryResponseAllOfConfig
type HelmRepositoryResponseAllOfConfig struct {
	Username             *string `json:"username,omitempty"`
	Region               *string `json:"region,omitempty"`
	AccessKeyId          *string `json:"access_key_id,omitempty"`
	ScalewayAccessKey    *string `json:"scaleway_access_key,omitempty"`
	ScalewayProjectId    *string `json:"scaleway_project_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HelmRepositoryResponseAllOfConfig HelmRepositoryResponseAllOfConfig

// NewHelmRepositoryResponseAllOfConfig instantiates a new HelmRepositoryResponseAllOfConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmRepositoryResponseAllOfConfig() *HelmRepositoryResponseAllOfConfig {
	this := HelmRepositoryResponseAllOfConfig{}
	return &this
}

// NewHelmRepositoryResponseAllOfConfigWithDefaults instantiates a new HelmRepositoryResponseAllOfConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmRepositoryResponseAllOfConfigWithDefaults() *HelmRepositoryResponseAllOfConfig {
	this := HelmRepositoryResponseAllOfConfig{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *HelmRepositoryResponseAllOfConfig) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRepositoryResponseAllOfConfig) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *HelmRepositoryResponseAllOfConfig) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *HelmRepositoryResponseAllOfConfig) SetUsername(v string) {
	o.Username = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *HelmRepositoryResponseAllOfConfig) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRepositoryResponseAllOfConfig) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *HelmRepositoryResponseAllOfConfig) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *HelmRepositoryResponseAllOfConfig) SetRegion(v string) {
	o.Region = &v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *HelmRepositoryResponseAllOfConfig) GetAccessKeyId() string {
	if o == nil || IsNil(o.AccessKeyId) {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRepositoryResponseAllOfConfig) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKeyId) {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *HelmRepositoryResponseAllOfConfig) HasAccessKeyId() bool {
	if o != nil && !IsNil(o.AccessKeyId) {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *HelmRepositoryResponseAllOfConfig) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetScalewayAccessKey returns the ScalewayAccessKey field value if set, zero value otherwise.
func (o *HelmRepositoryResponseAllOfConfig) GetScalewayAccessKey() string {
	if o == nil || IsNil(o.ScalewayAccessKey) {
		var ret string
		return ret
	}
	return *o.ScalewayAccessKey
}

// GetScalewayAccessKeyOk returns a tuple with the ScalewayAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRepositoryResponseAllOfConfig) GetScalewayAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ScalewayAccessKey) {
		return nil, false
	}
	return o.ScalewayAccessKey, true
}

// HasScalewayAccessKey returns a boolean if a field has been set.
func (o *HelmRepositoryResponseAllOfConfig) HasScalewayAccessKey() bool {
	if o != nil && !IsNil(o.ScalewayAccessKey) {
		return true
	}

	return false
}

// SetScalewayAccessKey gets a reference to the given string and assigns it to the ScalewayAccessKey field.
func (o *HelmRepositoryResponseAllOfConfig) SetScalewayAccessKey(v string) {
	o.ScalewayAccessKey = &v
}

// GetScalewayProjectId returns the ScalewayProjectId field value if set, zero value otherwise.
func (o *HelmRepositoryResponseAllOfConfig) GetScalewayProjectId() string {
	if o == nil || IsNil(o.ScalewayProjectId) {
		var ret string
		return ret
	}
	return *o.ScalewayProjectId
}

// GetScalewayProjectIdOk returns a tuple with the ScalewayProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRepositoryResponseAllOfConfig) GetScalewayProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScalewayProjectId) {
		return nil, false
	}
	return o.ScalewayProjectId, true
}

// HasScalewayProjectId returns a boolean if a field has been set.
func (o *HelmRepositoryResponseAllOfConfig) HasScalewayProjectId() bool {
	if o != nil && !IsNil(o.ScalewayProjectId) {
		return true
	}

	return false
}

// SetScalewayProjectId gets a reference to the given string and assigns it to the ScalewayProjectId field.
func (o *HelmRepositoryResponseAllOfConfig) SetScalewayProjectId(v string) {
	o.ScalewayProjectId = &v
}

func (o HelmRepositoryResponseAllOfConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmRepositoryResponseAllOfConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.AccessKeyId) {
		toSerialize["access_key_id"] = o.AccessKeyId
	}
	if !IsNil(o.ScalewayAccessKey) {
		toSerialize["scaleway_access_key"] = o.ScalewayAccessKey
	}
	if !IsNil(o.ScalewayProjectId) {
		toSerialize["scaleway_project_id"] = o.ScalewayProjectId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmRepositoryResponseAllOfConfig) UnmarshalJSON(data []byte) (err error) {
	varHelmRepositoryResponseAllOfConfig := _HelmRepositoryResponseAllOfConfig{}

	err = json.Unmarshal(data, &varHelmRepositoryResponseAllOfConfig)

	if err != nil {
		return err
	}

	*o = HelmRepositoryResponseAllOfConfig(varHelmRepositoryResponseAllOfConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "region")
		delete(additionalProperties, "access_key_id")
		delete(additionalProperties, "scaleway_access_key")
		delete(additionalProperties, "scaleway_project_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmRepositoryResponseAllOfConfig struct {
	value *HelmRepositoryResponseAllOfConfig
	isSet bool
}

func (v NullableHelmRepositoryResponseAllOfConfig) Get() *HelmRepositoryResponseAllOfConfig {
	return v.value
}

func (v *NullableHelmRepositoryResponseAllOfConfig) Set(val *HelmRepositoryResponseAllOfConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmRepositoryResponseAllOfConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmRepositoryResponseAllOfConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmRepositoryResponseAllOfConfig(val *HelmRepositoryResponseAllOfConfig) *NullableHelmRepositoryResponseAllOfConfig {
	return &NullableHelmRepositoryResponseAllOfConfig{value: val, isSet: true}
}

func (v NullableHelmRepositoryResponseAllOfConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmRepositoryResponseAllOfConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
