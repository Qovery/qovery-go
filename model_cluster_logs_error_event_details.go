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

// checks if the ClusterLogsErrorEventDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterLogsErrorEventDetails{}

// ClusterLogsErrorEventDetails struct for ClusterLogsErrorEventDetails
type ClusterLogsErrorEventDetails struct {
	// cloud provider used
	ProviderKind         *string                                  `json:"provider_kind,omitempty"`
	Region               *string                                  `json:"region,omitempty"`
	Transmitter          *ClusterLogsErrorEventDetailsTransmitter `json:"transmitter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClusterLogsErrorEventDetails ClusterLogsErrorEventDetails

// NewClusterLogsErrorEventDetails instantiates a new ClusterLogsErrorEventDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterLogsErrorEventDetails() *ClusterLogsErrorEventDetails {
	this := ClusterLogsErrorEventDetails{}
	return &this
}

// NewClusterLogsErrorEventDetailsWithDefaults instantiates a new ClusterLogsErrorEventDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterLogsErrorEventDetailsWithDefaults() *ClusterLogsErrorEventDetails {
	this := ClusterLogsErrorEventDetails{}
	return &this
}

// GetProviderKind returns the ProviderKind field value if set, zero value otherwise.
func (o *ClusterLogsErrorEventDetails) GetProviderKind() string {
	if o == nil || IsNil(o.ProviderKind) {
		var ret string
		return ret
	}
	return *o.ProviderKind
}

// GetProviderKindOk returns a tuple with the ProviderKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsErrorEventDetails) GetProviderKindOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderKind) {
		return nil, false
	}
	return o.ProviderKind, true
}

// HasProviderKind returns a boolean if a field has been set.
func (o *ClusterLogsErrorEventDetails) HasProviderKind() bool {
	if o != nil && !IsNil(o.ProviderKind) {
		return true
	}

	return false
}

// SetProviderKind gets a reference to the given string and assigns it to the ProviderKind field.
func (o *ClusterLogsErrorEventDetails) SetProviderKind(v string) {
	o.ProviderKind = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ClusterLogsErrorEventDetails) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsErrorEventDetails) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ClusterLogsErrorEventDetails) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ClusterLogsErrorEventDetails) SetRegion(v string) {
	o.Region = &v
}

// GetTransmitter returns the Transmitter field value if set, zero value otherwise.
func (o *ClusterLogsErrorEventDetails) GetTransmitter() ClusterLogsErrorEventDetailsTransmitter {
	if o == nil || IsNil(o.Transmitter) {
		var ret ClusterLogsErrorEventDetailsTransmitter
		return ret
	}
	return *o.Transmitter
}

// GetTransmitterOk returns a tuple with the Transmitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsErrorEventDetails) GetTransmitterOk() (*ClusterLogsErrorEventDetailsTransmitter, bool) {
	if o == nil || IsNil(o.Transmitter) {
		return nil, false
	}
	return o.Transmitter, true
}

// HasTransmitter returns a boolean if a field has been set.
func (o *ClusterLogsErrorEventDetails) HasTransmitter() bool {
	if o != nil && !IsNil(o.Transmitter) {
		return true
	}

	return false
}

// SetTransmitter gets a reference to the given ClusterLogsErrorEventDetailsTransmitter and assigns it to the Transmitter field.
func (o *ClusterLogsErrorEventDetails) SetTransmitter(v ClusterLogsErrorEventDetailsTransmitter) {
	o.Transmitter = &v
}

func (o ClusterLogsErrorEventDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterLogsErrorEventDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProviderKind) {
		toSerialize["provider_kind"] = o.ProviderKind
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Transmitter) {
		toSerialize["transmitter"] = o.Transmitter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClusterLogsErrorEventDetails) UnmarshalJSON(data []byte) (err error) {
	varClusterLogsErrorEventDetails := _ClusterLogsErrorEventDetails{}

	err = json.Unmarshal(data, &varClusterLogsErrorEventDetails)

	if err != nil {
		return err
	}

	*o = ClusterLogsErrorEventDetails(varClusterLogsErrorEventDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "provider_kind")
		delete(additionalProperties, "region")
		delete(additionalProperties, "transmitter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClusterLogsErrorEventDetails struct {
	value *ClusterLogsErrorEventDetails
	isSet bool
}

func (v NullableClusterLogsErrorEventDetails) Get() *ClusterLogsErrorEventDetails {
	return v.value
}

func (v *NullableClusterLogsErrorEventDetails) Set(val *ClusterLogsErrorEventDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterLogsErrorEventDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterLogsErrorEventDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterLogsErrorEventDetails(val *ClusterLogsErrorEventDetails) *NullableClusterLogsErrorEventDetails {
	return &NullableClusterLogsErrorEventDetails{value: val, isSet: true}
}

func (v NullableClusterLogsErrorEventDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterLogsErrorEventDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
