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

// ApplicationPortRequestPorts struct for ApplicationPortRequestPorts
type ApplicationPortRequestPorts struct {
	Name NullableString `json:"name,omitempty"`
	// The listening port of your application
	InternalPort int32 `json:"internal_port"`
	// The exposed port for your application. This is optional. If not set a default port will be used.
	ExternalPort *int32 `json:"external_port,omitempty"`
	// Expose the port to the world
	PubliclyAccessible bool    `json:"publicly_accessible"`
	Protocol           *string `json:"protocol,omitempty"`
}

// NewApplicationPortRequestPorts instantiates a new ApplicationPortRequestPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationPortRequestPorts(internalPort int32, publiclyAccessible bool) *ApplicationPortRequestPorts {
	this := ApplicationPortRequestPorts{}
	this.InternalPort = internalPort
	this.PubliclyAccessible = publiclyAccessible
	var protocol string = "HTTP"
	this.Protocol = &protocol
	return &this
}

// NewApplicationPortRequestPortsWithDefaults instantiates a new ApplicationPortRequestPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationPortRequestPortsWithDefaults() *ApplicationPortRequestPorts {
	this := ApplicationPortRequestPorts{}
	var protocol string = "HTTP"
	this.Protocol = &protocol
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationPortRequestPorts) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationPortRequestPorts) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationPortRequestPorts) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ApplicationPortRequestPorts) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ApplicationPortRequestPorts) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ApplicationPortRequestPorts) UnsetName() {
	o.Name.Unset()
}

// GetInternalPort returns the InternalPort field value
func (o *ApplicationPortRequestPorts) GetInternalPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InternalPort
}

// GetInternalPortOk returns a tuple with the InternalPort field value
// and a boolean to check if the value has been set.
func (o *ApplicationPortRequestPorts) GetInternalPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternalPort, true
}

// SetInternalPort sets field value
func (o *ApplicationPortRequestPorts) SetInternalPort(v int32) {
	o.InternalPort = v
}

// GetExternalPort returns the ExternalPort field value if set, zero value otherwise.
func (o *ApplicationPortRequestPorts) GetExternalPort() int32 {
	if o == nil || o.ExternalPort == nil {
		var ret int32
		return ret
	}
	return *o.ExternalPort
}

// GetExternalPortOk returns a tuple with the ExternalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPortRequestPorts) GetExternalPortOk() (*int32, bool) {
	if o == nil || o.ExternalPort == nil {
		return nil, false
	}
	return o.ExternalPort, true
}

// HasExternalPort returns a boolean if a field has been set.
func (o *ApplicationPortRequestPorts) HasExternalPort() bool {
	if o != nil && o.ExternalPort != nil {
		return true
	}

	return false
}

// SetExternalPort gets a reference to the given int32 and assigns it to the ExternalPort field.
func (o *ApplicationPortRequestPorts) SetExternalPort(v int32) {
	o.ExternalPort = &v
}

// GetPubliclyAccessible returns the PubliclyAccessible field value
func (o *ApplicationPortRequestPorts) GetPubliclyAccessible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PubliclyAccessible
}

// GetPubliclyAccessibleOk returns a tuple with the PubliclyAccessible field value
// and a boolean to check if the value has been set.
func (o *ApplicationPortRequestPorts) GetPubliclyAccessibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PubliclyAccessible, true
}

// SetPubliclyAccessible sets field value
func (o *ApplicationPortRequestPorts) SetPubliclyAccessible(v bool) {
	o.PubliclyAccessible = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ApplicationPortRequestPorts) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPortRequestPorts) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ApplicationPortRequestPorts) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ApplicationPortRequestPorts) SetProtocol(v string) {
	o.Protocol = &v
}

func (o ApplicationPortRequestPorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["internal_port"] = o.InternalPort
	}
	if o.ExternalPort != nil {
		toSerialize["external_port"] = o.ExternalPort
	}
	if true {
		toSerialize["publicly_accessible"] = o.PubliclyAccessible
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationPortRequestPorts struct {
	value *ApplicationPortRequestPorts
	isSet bool
}

func (v NullableApplicationPortRequestPorts) Get() *ApplicationPortRequestPorts {
	return v.value
}

func (v *NullableApplicationPortRequestPorts) Set(val *ApplicationPortRequestPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationPortRequestPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationPortRequestPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationPortRequestPorts(val *ApplicationPortRequestPorts) *NullableApplicationPortRequestPorts {
	return &NullableApplicationPortRequestPorts{value: val, isSet: true}
}

func (v NullableApplicationPortRequestPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationPortRequestPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
