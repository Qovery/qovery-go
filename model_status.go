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

// Status struct for Status
type Status struct {
	Id *string `json:"id,omitempty"`
	// Status is a state machine. It starts with `BUILDING` or `DEPLOYING` state (or `INITIALIZED`if auto-deploy is deactivated). Then finish with `*_ERROR` or any termination state.
	State string `json:"state"`
	// message related to the state
	Message                 NullableString `json:"message,omitempty"`
	ServiceDeploymentStatus NullableString `json:"service_deployment_status,omitempty"`
}

// NewStatus instantiates a new Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus(state string) *Status {
	this := Status{}
	this.State = state
	return &this
}

// NewStatusWithDefaults instantiates a new Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusWithDefaults() *Status {
	this := Status{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Status) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Status) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Status) SetId(v string) {
	o.Id = &v
}

// GetState returns the State field value
func (o *Status) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Status) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Status) SetState(v string) {
	o.State = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Status) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Status) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *Status) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *Status) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil
func (o *Status) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *Status) UnsetMessage() {
	o.Message.Unset()
}

// GetServiceDeploymentStatus returns the ServiceDeploymentStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Status) GetServiceDeploymentStatus() string {
	if o == nil || o.ServiceDeploymentStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceDeploymentStatus.Get()
}

// GetServiceDeploymentStatusOk returns a tuple with the ServiceDeploymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Status) GetServiceDeploymentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceDeploymentStatus.Get(), o.ServiceDeploymentStatus.IsSet()
}

// HasServiceDeploymentStatus returns a boolean if a field has been set.
func (o *Status) HasServiceDeploymentStatus() bool {
	if o != nil && o.ServiceDeploymentStatus.IsSet() {
		return true
	}

	return false
}

// SetServiceDeploymentStatus gets a reference to the given NullableString and assigns it to the ServiceDeploymentStatus field.
func (o *Status) SetServiceDeploymentStatus(v string) {
	o.ServiceDeploymentStatus.Set(&v)
}

// SetServiceDeploymentStatusNil sets the value for ServiceDeploymentStatus to be an explicit nil
func (o *Status) SetServiceDeploymentStatusNil() {
	o.ServiceDeploymentStatus.Set(nil)
}

// UnsetServiceDeploymentStatus ensures that no value is present for ServiceDeploymentStatus, not even an explicit nil
func (o *Status) UnsetServiceDeploymentStatus() {
	o.ServiceDeploymentStatus.Unset()
}

func (o Status) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.ServiceDeploymentStatus.IsSet() {
		toSerialize["service_deployment_status"] = o.ServiceDeploymentStatus.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableStatus struct {
	value *Status
	isSet bool
}

func (v NullableStatus) Get() *Status {
	return v.value
}

func (v *NullableStatus) Set(val *Status) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus(val *Status) *NullableStatus {
	return &NullableStatus{value: val, isSet: true}
}

func (v NullableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
