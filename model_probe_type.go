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

// checks if the ProbeType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProbeType{}

// ProbeType struct for ProbeType
type ProbeType struct {
	Tcp                  NullableProbeTypeTcp  `json:"tcp,omitempty"`
	Http                 NullableProbeTypeHttp `json:"http,omitempty"`
	Exec                 NullableProbeTypeExec `json:"exec,omitempty"`
	Grpc                 NullableProbeTypeGrpc `json:"grpc,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProbeType ProbeType

// NewProbeType instantiates a new ProbeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProbeType() *ProbeType {
	this := ProbeType{}
	return &this
}

// NewProbeTypeWithDefaults instantiates a new ProbeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProbeTypeWithDefaults() *ProbeType {
	this := ProbeType{}
	return &this
}

// GetTcp returns the Tcp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProbeType) GetTcp() ProbeTypeTcp {
	if o == nil || IsNil(o.Tcp.Get()) {
		var ret ProbeTypeTcp
		return ret
	}
	return *o.Tcp.Get()
}

// GetTcpOk returns a tuple with the Tcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProbeType) GetTcpOk() (*ProbeTypeTcp, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tcp.Get(), o.Tcp.IsSet()
}

// HasTcp returns a boolean if a field has been set.
func (o *ProbeType) HasTcp() bool {
	if o != nil && o.Tcp.IsSet() {
		return true
	}

	return false
}

// SetTcp gets a reference to the given NullableProbeTypeTcp and assigns it to the Tcp field.
func (o *ProbeType) SetTcp(v ProbeTypeTcp) {
	o.Tcp.Set(&v)
}

// SetTcpNil sets the value for Tcp to be an explicit nil
func (o *ProbeType) SetTcpNil() {
	o.Tcp.Set(nil)
}

// UnsetTcp ensures that no value is present for Tcp, not even an explicit nil
func (o *ProbeType) UnsetTcp() {
	o.Tcp.Unset()
}

// GetHttp returns the Http field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProbeType) GetHttp() ProbeTypeHttp {
	if o == nil || IsNil(o.Http.Get()) {
		var ret ProbeTypeHttp
		return ret
	}
	return *o.Http.Get()
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProbeType) GetHttpOk() (*ProbeTypeHttp, bool) {
	if o == nil {
		return nil, false
	}
	return o.Http.Get(), o.Http.IsSet()
}

// HasHttp returns a boolean if a field has been set.
func (o *ProbeType) HasHttp() bool {
	if o != nil && o.Http.IsSet() {
		return true
	}

	return false
}

// SetHttp gets a reference to the given NullableProbeTypeHttp and assigns it to the Http field.
func (o *ProbeType) SetHttp(v ProbeTypeHttp) {
	o.Http.Set(&v)
}

// SetHttpNil sets the value for Http to be an explicit nil
func (o *ProbeType) SetHttpNil() {
	o.Http.Set(nil)
}

// UnsetHttp ensures that no value is present for Http, not even an explicit nil
func (o *ProbeType) UnsetHttp() {
	o.Http.Unset()
}

// GetExec returns the Exec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProbeType) GetExec() ProbeTypeExec {
	if o == nil || IsNil(o.Exec.Get()) {
		var ret ProbeTypeExec
		return ret
	}
	return *o.Exec.Get()
}

// GetExecOk returns a tuple with the Exec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProbeType) GetExecOk() (*ProbeTypeExec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Exec.Get(), o.Exec.IsSet()
}

// HasExec returns a boolean if a field has been set.
func (o *ProbeType) HasExec() bool {
	if o != nil && o.Exec.IsSet() {
		return true
	}

	return false
}

// SetExec gets a reference to the given NullableProbeTypeExec and assigns it to the Exec field.
func (o *ProbeType) SetExec(v ProbeTypeExec) {
	o.Exec.Set(&v)
}

// SetExecNil sets the value for Exec to be an explicit nil
func (o *ProbeType) SetExecNil() {
	o.Exec.Set(nil)
}

// UnsetExec ensures that no value is present for Exec, not even an explicit nil
func (o *ProbeType) UnsetExec() {
	o.Exec.Unset()
}

// GetGrpc returns the Grpc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProbeType) GetGrpc() ProbeTypeGrpc {
	if o == nil || IsNil(o.Grpc.Get()) {
		var ret ProbeTypeGrpc
		return ret
	}
	return *o.Grpc.Get()
}

// GetGrpcOk returns a tuple with the Grpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProbeType) GetGrpcOk() (*ProbeTypeGrpc, bool) {
	if o == nil {
		return nil, false
	}
	return o.Grpc.Get(), o.Grpc.IsSet()
}

// HasGrpc returns a boolean if a field has been set.
func (o *ProbeType) HasGrpc() bool {
	if o != nil && o.Grpc.IsSet() {
		return true
	}

	return false
}

// SetGrpc gets a reference to the given NullableProbeTypeGrpc and assigns it to the Grpc field.
func (o *ProbeType) SetGrpc(v ProbeTypeGrpc) {
	o.Grpc.Set(&v)
}

// SetGrpcNil sets the value for Grpc to be an explicit nil
func (o *ProbeType) SetGrpcNil() {
	o.Grpc.Set(nil)
}

// UnsetGrpc ensures that no value is present for Grpc, not even an explicit nil
func (o *ProbeType) UnsetGrpc() {
	o.Grpc.Unset()
}

func (o ProbeType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProbeType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Tcp.IsSet() {
		toSerialize["tcp"] = o.Tcp.Get()
	}
	if o.Http.IsSet() {
		toSerialize["http"] = o.Http.Get()
	}
	if o.Exec.IsSet() {
		toSerialize["exec"] = o.Exec.Get()
	}
	if o.Grpc.IsSet() {
		toSerialize["grpc"] = o.Grpc.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProbeType) UnmarshalJSON(data []byte) (err error) {
	varProbeType := _ProbeType{}

	err = json.Unmarshal(data, &varProbeType)

	if err != nil {
		return err
	}

	*o = ProbeType(varProbeType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tcp")
		delete(additionalProperties, "http")
		delete(additionalProperties, "exec")
		delete(additionalProperties, "grpc")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProbeType struct {
	value *ProbeType
	isSet bool
}

func (v NullableProbeType) Get() *ProbeType {
	return v.value
}

func (v *NullableProbeType) Set(val *ProbeType) {
	v.value = val
	v.isSet = true
}

func (v NullableProbeType) IsSet() bool {
	return v.isSet
}

func (v *NullableProbeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProbeType(val *ProbeType) *NullableProbeType {
	return &NullableProbeType{value: val, isSet: true}
}

func (v NullableProbeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProbeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
