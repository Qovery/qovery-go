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

// ProbeProbeHttp struct for ProbeProbeHttp
type ProbeProbeHttp struct {
	Path   *string `json:"path,omitempty"`
	Scheme *string `json:"scheme,omitempty"`
}

// NewProbeProbeHttp instantiates a new ProbeProbeHttp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProbeProbeHttp() *ProbeProbeHttp {
	this := ProbeProbeHttp{}
	var path string = "/"
	this.Path = &path
	var scheme string = "HTTP"
	this.Scheme = &scheme
	return &this
}

// NewProbeProbeHttpWithDefaults instantiates a new ProbeProbeHttp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProbeProbeHttpWithDefaults() *ProbeProbeHttp {
	this := ProbeProbeHttp{}
	var path string = "/"
	this.Path = &path
	var scheme string = "HTTP"
	this.Scheme = &scheme
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ProbeProbeHttp) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProbeProbeHttp) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ProbeProbeHttp) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ProbeProbeHttp) SetPath(v string) {
	o.Path = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *ProbeProbeHttp) GetScheme() string {
	if o == nil || o.Scheme == nil {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProbeProbeHttp) GetSchemeOk() (*string, bool) {
	if o == nil || o.Scheme == nil {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *ProbeProbeHttp) HasScheme() bool {
	if o != nil && o.Scheme != nil {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *ProbeProbeHttp) SetScheme(v string) {
	o.Scheme = &v
}

func (o ProbeProbeHttp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Scheme != nil {
		toSerialize["scheme"] = o.Scheme
	}
	return json.Marshal(toSerialize)
}

type NullableProbeProbeHttp struct {
	value *ProbeProbeHttp
	isSet bool
}

func (v NullableProbeProbeHttp) Get() *ProbeProbeHttp {
	return v.value
}

func (v *NullableProbeProbeHttp) Set(val *ProbeProbeHttp) {
	v.value = val
	v.isSet = true
}

func (v NullableProbeProbeHttp) IsSet() bool {
	return v.isSet
}

func (v *NullableProbeProbeHttp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProbeProbeHttp(val *ProbeProbeHttp) *NullableProbeProbeHttp {
	return &NullableProbeProbeHttp{value: val, isSet: true}
}

func (v NullableProbeProbeHttp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProbeProbeHttp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
