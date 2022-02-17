/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"time"
)

// EnvironmentVariableResponse struct for EnvironmentVariableResponse
type EnvironmentVariableResponse struct {
	OverriddenVariable *map[string]interface{} `json:"overridden_variable,omitempty"`
	AliasedVariable    *map[string]interface{} `json:"aliased_variable,omitempty"`
	Scope              string                  `json:"scope"`
	ServiceName        *string                 `json:"service_name,omitempty"`
	Id                 string                  `json:"id"`
	CreatedAt          time.Time               `json:"created_at"`
	UpdatedAt          *time.Time              `json:"updated_at,omitempty"`
	// key is case sensitive
	Key string `json:"key"`
	// value of the env variable.
	Value string `json:"value"`
}

// NewEnvironmentVariableResponse instantiates a new EnvironmentVariableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariableResponse(scope string, id string, createdAt time.Time, key string, value string) *EnvironmentVariableResponse {
	this := EnvironmentVariableResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Key = key
	this.Value = value
	return &this
}

// NewEnvironmentVariableResponseWithDefaults instantiates a new EnvironmentVariableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableResponseWithDefaults() *EnvironmentVariableResponse {
	this := EnvironmentVariableResponse{}
	return &this
}

// GetOverriddenVariable returns the OverriddenVariable field value if set, zero value otherwise.
func (o *EnvironmentVariableResponse) GetOverriddenVariable() map[string]interface{} {
	if o == nil || o.OverriddenVariable == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.OverriddenVariable
}

// GetOverriddenVariableOk returns a tuple with the OverriddenVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponse) GetOverriddenVariableOk() (*map[string]interface{}, bool) {
	if o == nil || o.OverriddenVariable == nil {
		return nil, false
	}
	return o.OverriddenVariable, true
}

// HasOverriddenVariable returns a boolean if a field has been set.
func (o *EnvironmentVariableResponse) HasOverriddenVariable() bool {
	if o != nil && o.OverriddenVariable != nil {
		return true
	}

	return false
}

// SetOverriddenVariable gets a reference to the given map[string]interface{} and assigns it to the OverriddenVariable field.
func (o *EnvironmentVariableResponse) SetOverriddenVariable(v map[string]interface{}) {
	o.OverriddenVariable = &v
}

// GetAliasedVariable returns the AliasedVariable field value if set, zero value otherwise.
func (o *EnvironmentVariableResponse) GetAliasedVariable() map[string]interface{} {
	if o == nil || o.AliasedVariable == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AliasedVariable
}

// GetAliasedVariableOk returns a tuple with the AliasedVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponse) GetAliasedVariableOk() (*map[string]interface{}, bool) {
	if o == nil || o.AliasedVariable == nil {
		return nil, false
	}
	return o.AliasedVariable, true
}

// HasAliasedVariable returns a boolean if a field has been set.
func (o *EnvironmentVariableResponse) HasAliasedVariable() bool {
	if o != nil && o.AliasedVariable != nil {
		return true
	}

	return false
}

// SetAliasedVariable gets a reference to the given map[string]interface{} and assigns it to the AliasedVariable field.
func (o *EnvironmentVariableResponse) SetAliasedVariable(v map[string]interface{}) {
	o.AliasedVariable = &v
}

// GetScope returns the Scope field value
func (o *EnvironmentVariableResponse) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponse) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *EnvironmentVariableResponse) SetScope(v string) {
	o.Scope = v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *EnvironmentVariableResponse) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponse) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *EnvironmentVariableResponse) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *EnvironmentVariableResponse) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetId returns the Id field value
func (o *EnvironmentVariableResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentVariableResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EnvironmentVariableResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EnvironmentVariableResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentVariableResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentVariableResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentVariableResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetKey returns the Key field value
func (o *EnvironmentVariableResponse) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EnvironmentVariableResponse) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *EnvironmentVariableResponse) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EnvironmentVariableResponse) SetValue(v string) {
	o.Value = v
}

func (o EnvironmentVariableResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OverriddenVariable != nil {
		toSerialize["overridden_variable"] = o.OverriddenVariable
	}
	if o.AliasedVariable != nil {
		toSerialize["aliased_variable"] = o.AliasedVariable
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentVariableResponse struct {
	value *EnvironmentVariableResponse
	isSet bool
}

func (v NullableEnvironmentVariableResponse) Get() *EnvironmentVariableResponse {
	return v.value
}

func (v *NullableEnvironmentVariableResponse) Set(val *EnvironmentVariableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariableResponse(val *EnvironmentVariableResponse) *NullableEnvironmentVariableResponse {
	return &NullableEnvironmentVariableResponse{value: val, isSet: true}
}

func (v NullableEnvironmentVariableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
