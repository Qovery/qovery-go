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

// checks if the ApplicationNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationNetwork{}

// ApplicationNetwork struct for ApplicationNetwork
type ApplicationNetwork struct {
	// Specify if the sticky session option (also called persistant session) is activated or not for this application. If activated, user will be redirected by the load balancer to the same instance each time he access to the application.
	StickySession        *bool `json:"sticky_session,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationNetwork ApplicationNetwork

// NewApplicationNetwork instantiates a new ApplicationNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationNetwork() *ApplicationNetwork {
	this := ApplicationNetwork{}
	var stickySession bool = false
	this.StickySession = &stickySession
	return &this
}

// NewApplicationNetworkWithDefaults instantiates a new ApplicationNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationNetworkWithDefaults() *ApplicationNetwork {
	this := ApplicationNetwork{}
	var stickySession bool = false
	this.StickySession = &stickySession
	return &this
}

// GetStickySession returns the StickySession field value if set, zero value otherwise.
func (o *ApplicationNetwork) GetStickySession() bool {
	if o == nil || IsNil(o.StickySession) {
		var ret bool
		return ret
	}
	return *o.StickySession
}

// GetStickySessionOk returns a tuple with the StickySession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationNetwork) GetStickySessionOk() (*bool, bool) {
	if o == nil || IsNil(o.StickySession) {
		return nil, false
	}
	return o.StickySession, true
}

// HasStickySession returns a boolean if a field has been set.
func (o *ApplicationNetwork) HasStickySession() bool {
	if o != nil && !IsNil(o.StickySession) {
		return true
	}

	return false
}

// SetStickySession gets a reference to the given bool and assigns it to the StickySession field.
func (o *ApplicationNetwork) SetStickySession(v bool) {
	o.StickySession = &v
}

func (o ApplicationNetwork) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StickySession) {
		toSerialize["sticky_session"] = o.StickySession
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationNetwork) UnmarshalJSON(data []byte) (err error) {
	varApplicationNetwork := _ApplicationNetwork{}

	err = json.Unmarshal(data, &varApplicationNetwork)

	if err != nil {
		return err
	}

	*o = ApplicationNetwork(varApplicationNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sticky_session")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationNetwork struct {
	value *ApplicationNetwork
	isSet bool
}

func (v NullableApplicationNetwork) Get() *ApplicationNetwork {
	return v.value
}

func (v *NullableApplicationNetwork) Set(val *ApplicationNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationNetwork(val *ApplicationNetwork) *NullableApplicationNetwork {
	return &NullableApplicationNetwork{value: val, isSet: true}
}

func (v NullableApplicationNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
