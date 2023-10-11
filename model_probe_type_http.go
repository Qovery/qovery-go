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

// checks if the ProbeTypeHttp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProbeTypeHttp{}

// ProbeTypeHttp struct for ProbeTypeHttp
type ProbeTypeHttp struct {
	Path   *string `json:"path,omitempty"`
	Scheme *string `json:"scheme,omitempty"`
	Port   *int32  `json:"port,omitempty"`
}

// NewProbeTypeHttp instantiates a new ProbeTypeHttp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProbeTypeHttp() *ProbeTypeHttp {
	this := ProbeTypeHttp{}
	var path string = "/"
	this.Path = &path
	var scheme string = "HTTP"
	this.Scheme = &scheme
	return &this
}

// NewProbeTypeHttpWithDefaults instantiates a new ProbeTypeHttp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProbeTypeHttpWithDefaults() *ProbeTypeHttp {
	this := ProbeTypeHttp{}
	var path string = "/"
	this.Path = &path
	var scheme string = "HTTP"
	this.Scheme = &scheme
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ProbeTypeHttp) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProbeTypeHttp) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ProbeTypeHttp) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ProbeTypeHttp) SetPath(v string) {
	o.Path = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *ProbeTypeHttp) GetScheme() string {
	if o == nil || IsNil(o.Scheme) {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProbeTypeHttp) GetSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.Scheme) {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *ProbeTypeHttp) HasScheme() bool {
	if o != nil && !IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *ProbeTypeHttp) SetScheme(v string) {
	o.Scheme = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ProbeTypeHttp) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProbeTypeHttp) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ProbeTypeHttp) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ProbeTypeHttp) SetPort(v int32) {
	o.Port = &v
}

func (o ProbeTypeHttp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProbeTypeHttp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableProbeTypeHttp struct {
	value *ProbeTypeHttp
	isSet bool
}

func (v NullableProbeTypeHttp) Get() *ProbeTypeHttp {
	return v.value
}

func (v *NullableProbeTypeHttp) Set(val *ProbeTypeHttp) {
	v.value = val
	v.isSet = true
}

func (v NullableProbeTypeHttp) IsSet() bool {
	return v.isSet
}

func (v *NullableProbeTypeHttp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProbeTypeHttp(val *ProbeTypeHttp) *NullableProbeTypeHttp {
	return &NullableProbeTypeHttp{value: val, isSet: true}
}

func (v NullableProbeTypeHttp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProbeTypeHttp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
