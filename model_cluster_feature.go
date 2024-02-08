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

// checks if the ClusterFeature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterFeature{}

// ClusterFeature struct for ClusterFeature
type ClusterFeature struct {
	Id                   *string                             `json:"id,omitempty"`
	Title                *string                             `json:"title,omitempty"`
	Description          NullableString                      `json:"description,omitempty"`
	CostPerMonthInCents  NullableInt32                       `json:"cost_per_month_in_cents,omitempty"`
	CostPerMonth         NullableFloat32                     `json:"cost_per_month,omitempty"`
	CurrencyCode         NullableString                      `json:"currency_code,omitempty"`
	ValueType            *string                             `json:"value_type,omitempty"`
	Value                NullableClusterFeatureValue         `json:"value,omitempty"`
	IsValueUpdatable     *bool                               `json:"is_value_updatable,omitempty"`
	AcceptedValues       []ClusterFeatureAcceptedValuesInner `json:"accepted_values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClusterFeature ClusterFeature

// NewClusterFeature instantiates a new ClusterFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterFeature() *ClusterFeature {
	this := ClusterFeature{}
	var isValueUpdatable bool = false
	this.IsValueUpdatable = &isValueUpdatable
	return &this
}

// NewClusterFeatureWithDefaults instantiates a new ClusterFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterFeatureWithDefaults() *ClusterFeature {
	this := ClusterFeature{}
	var isValueUpdatable bool = false
	this.IsValueUpdatable = &isValueUpdatable
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterFeature) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeature) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterFeature) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClusterFeature) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ClusterFeature) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeature) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ClusterFeature) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ClusterFeature) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeature) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeature) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterFeature) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ClusterFeature) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ClusterFeature) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ClusterFeature) UnsetDescription() {
	o.Description.Unset()
}

// GetCostPerMonthInCents returns the CostPerMonthInCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeature) GetCostPerMonthInCents() int32 {
	if o == nil || IsNil(o.CostPerMonthInCents.Get()) {
		var ret int32
		return ret
	}
	return *o.CostPerMonthInCents.Get()
}

// GetCostPerMonthInCentsOk returns a tuple with the CostPerMonthInCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeature) GetCostPerMonthInCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CostPerMonthInCents.Get(), o.CostPerMonthInCents.IsSet()
}

// HasCostPerMonthInCents returns a boolean if a field has been set.
func (o *ClusterFeature) HasCostPerMonthInCents() bool {
	if o != nil && o.CostPerMonthInCents.IsSet() {
		return true
	}

	return false
}

// SetCostPerMonthInCents gets a reference to the given NullableInt32 and assigns it to the CostPerMonthInCents field.
func (o *ClusterFeature) SetCostPerMonthInCents(v int32) {
	o.CostPerMonthInCents.Set(&v)
}

// SetCostPerMonthInCentsNil sets the value for CostPerMonthInCents to be an explicit nil
func (o *ClusterFeature) SetCostPerMonthInCentsNil() {
	o.CostPerMonthInCents.Set(nil)
}

// UnsetCostPerMonthInCents ensures that no value is present for CostPerMonthInCents, not even an explicit nil
func (o *ClusterFeature) UnsetCostPerMonthInCents() {
	o.CostPerMonthInCents.Unset()
}

// GetCostPerMonth returns the CostPerMonth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeature) GetCostPerMonth() float32 {
	if o == nil || IsNil(o.CostPerMonth.Get()) {
		var ret float32
		return ret
	}
	return *o.CostPerMonth.Get()
}

// GetCostPerMonthOk returns a tuple with the CostPerMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeature) GetCostPerMonthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CostPerMonth.Get(), o.CostPerMonth.IsSet()
}

// HasCostPerMonth returns a boolean if a field has been set.
func (o *ClusterFeature) HasCostPerMonth() bool {
	if o != nil && o.CostPerMonth.IsSet() {
		return true
	}

	return false
}

// SetCostPerMonth gets a reference to the given NullableFloat32 and assigns it to the CostPerMonth field.
func (o *ClusterFeature) SetCostPerMonth(v float32) {
	o.CostPerMonth.Set(&v)
}

// SetCostPerMonthNil sets the value for CostPerMonth to be an explicit nil
func (o *ClusterFeature) SetCostPerMonthNil() {
	o.CostPerMonth.Set(nil)
}

// UnsetCostPerMonth ensures that no value is present for CostPerMonth, not even an explicit nil
func (o *ClusterFeature) UnsetCostPerMonth() {
	o.CostPerMonth.Unset()
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeature) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencyCode.Get()
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeature) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyCode.Get(), o.CurrencyCode.IsSet()
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ClusterFeature) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given NullableString and assigns it to the CurrencyCode field.
func (o *ClusterFeature) SetCurrencyCode(v string) {
	o.CurrencyCode.Set(&v)
}

// SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil
func (o *ClusterFeature) SetCurrencyCodeNil() {
	o.CurrencyCode.Set(nil)
}

// UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
func (o *ClusterFeature) UnsetCurrencyCode() {
	o.CurrencyCode.Unset()
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *ClusterFeature) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeature) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *ClusterFeature) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *ClusterFeature) SetValueType(v string) {
	o.ValueType = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeature) GetValue() ClusterFeatureValue {
	if o == nil || IsNil(o.Value.Get()) {
		var ret ClusterFeatureValue
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeature) GetValueOk() (*ClusterFeatureValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ClusterFeature) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableClusterFeatureValue and assigns it to the Value field.
func (o *ClusterFeature) SetValue(v ClusterFeatureValue) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *ClusterFeature) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ClusterFeature) UnsetValue() {
	o.Value.Unset()
}

// GetIsValueUpdatable returns the IsValueUpdatable field value if set, zero value otherwise.
func (o *ClusterFeature) GetIsValueUpdatable() bool {
	if o == nil || IsNil(o.IsValueUpdatable) {
		var ret bool
		return ret
	}
	return *o.IsValueUpdatable
}

// GetIsValueUpdatableOk returns a tuple with the IsValueUpdatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeature) GetIsValueUpdatableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsValueUpdatable) {
		return nil, false
	}
	return o.IsValueUpdatable, true
}

// HasIsValueUpdatable returns a boolean if a field has been set.
func (o *ClusterFeature) HasIsValueUpdatable() bool {
	if o != nil && !IsNil(o.IsValueUpdatable) {
		return true
	}

	return false
}

// SetIsValueUpdatable gets a reference to the given bool and assigns it to the IsValueUpdatable field.
func (o *ClusterFeature) SetIsValueUpdatable(v bool) {
	o.IsValueUpdatable = &v
}

// GetAcceptedValues returns the AcceptedValues field value if set, zero value otherwise.
func (o *ClusterFeature) GetAcceptedValues() []ClusterFeatureAcceptedValuesInner {
	if o == nil || IsNil(o.AcceptedValues) {
		var ret []ClusterFeatureAcceptedValuesInner
		return ret
	}
	return o.AcceptedValues
}

// GetAcceptedValuesOk returns a tuple with the AcceptedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeature) GetAcceptedValuesOk() ([]ClusterFeatureAcceptedValuesInner, bool) {
	if o == nil || IsNil(o.AcceptedValues) {
		return nil, false
	}
	return o.AcceptedValues, true
}

// HasAcceptedValues returns a boolean if a field has been set.
func (o *ClusterFeature) HasAcceptedValues() bool {
	if o != nil && !IsNil(o.AcceptedValues) {
		return true
	}

	return false
}

// SetAcceptedValues gets a reference to the given []ClusterFeatureAcceptedValuesInner and assigns it to the AcceptedValues field.
func (o *ClusterFeature) SetAcceptedValues(v []ClusterFeatureAcceptedValuesInner) {
	o.AcceptedValues = v
}

func (o ClusterFeature) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterFeature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.CostPerMonthInCents.IsSet() {
		toSerialize["cost_per_month_in_cents"] = o.CostPerMonthInCents.Get()
	}
	if o.CostPerMonth.IsSet() {
		toSerialize["cost_per_month"] = o.CostPerMonth.Get()
	}
	if o.CurrencyCode.IsSet() {
		toSerialize["currency_code"] = o.CurrencyCode.Get()
	}
	if !IsNil(o.ValueType) {
		toSerialize["value_type"] = o.ValueType
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if !IsNil(o.IsValueUpdatable) {
		toSerialize["is_value_updatable"] = o.IsValueUpdatable
	}
	if !IsNil(o.AcceptedValues) {
		toSerialize["accepted_values"] = o.AcceptedValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClusterFeature) UnmarshalJSON(data []byte) (err error) {
	varClusterFeature := _ClusterFeature{}

	err = json.Unmarshal(data, &varClusterFeature)

	if err != nil {
		return err
	}

	*o = ClusterFeature(varClusterFeature)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "title")
		delete(additionalProperties, "description")
		delete(additionalProperties, "cost_per_month_in_cents")
		delete(additionalProperties, "cost_per_month")
		delete(additionalProperties, "currency_code")
		delete(additionalProperties, "value_type")
		delete(additionalProperties, "value")
		delete(additionalProperties, "is_value_updatable")
		delete(additionalProperties, "accepted_values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClusterFeature struct {
	value *ClusterFeature
	isSet bool
}

func (v NullableClusterFeature) Get() *ClusterFeature {
	return v.value
}

func (v *NullableClusterFeature) Set(val *ClusterFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterFeature(val *ClusterFeature) *NullableClusterFeature {
	return &NullableClusterFeature{value: val, isSet: true}
}

func (v NullableClusterFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
