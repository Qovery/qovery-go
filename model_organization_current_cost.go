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

// checks if the OrganizationCurrentCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationCurrentCost{}

// OrganizationCurrentCost struct for OrganizationCurrentCost
type OrganizationCurrentCost struct {
	Plan *PlanEnum `json:"plan,omitempty"`
	// number of days remaining before the end of the trial period
	RemainingTrialDay    *int32            `json:"remaining_trial_day,omitempty"`
	RemainingCredits     *RemainingCredits `json:"remaining_credits,omitempty"`
	Cost                 *Cost             `json:"cost,omitempty"`
	PaidUsage            *PaidUsage        `json:"paid_usage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationCurrentCost OrganizationCurrentCost

// NewOrganizationCurrentCost instantiates a new OrganizationCurrentCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationCurrentCost() *OrganizationCurrentCost {
	this := OrganizationCurrentCost{}
	return &this
}

// NewOrganizationCurrentCostWithDefaults instantiates a new OrganizationCurrentCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationCurrentCostWithDefaults() *OrganizationCurrentCost {
	this := OrganizationCurrentCost{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *OrganizationCurrentCost) GetPlan() PlanEnum {
	if o == nil || IsNil(o.Plan) {
		var ret PlanEnum
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCurrentCost) GetPlanOk() (*PlanEnum, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *OrganizationCurrentCost) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given PlanEnum and assigns it to the Plan field.
func (o *OrganizationCurrentCost) SetPlan(v PlanEnum) {
	o.Plan = &v
}

// GetRemainingTrialDay returns the RemainingTrialDay field value if set, zero value otherwise.
func (o *OrganizationCurrentCost) GetRemainingTrialDay() int32 {
	if o == nil || IsNil(o.RemainingTrialDay) {
		var ret int32
		return ret
	}
	return *o.RemainingTrialDay
}

// GetRemainingTrialDayOk returns a tuple with the RemainingTrialDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCurrentCost) GetRemainingTrialDayOk() (*int32, bool) {
	if o == nil || IsNil(o.RemainingTrialDay) {
		return nil, false
	}
	return o.RemainingTrialDay, true
}

// HasRemainingTrialDay returns a boolean if a field has been set.
func (o *OrganizationCurrentCost) HasRemainingTrialDay() bool {
	if o != nil && !IsNil(o.RemainingTrialDay) {
		return true
	}

	return false
}

// SetRemainingTrialDay gets a reference to the given int32 and assigns it to the RemainingTrialDay field.
func (o *OrganizationCurrentCost) SetRemainingTrialDay(v int32) {
	o.RemainingTrialDay = &v
}

// GetRemainingCredits returns the RemainingCredits field value if set, zero value otherwise.
func (o *OrganizationCurrentCost) GetRemainingCredits() RemainingCredits {
	if o == nil || IsNil(o.RemainingCredits) {
		var ret RemainingCredits
		return ret
	}
	return *o.RemainingCredits
}

// GetRemainingCreditsOk returns a tuple with the RemainingCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCurrentCost) GetRemainingCreditsOk() (*RemainingCredits, bool) {
	if o == nil || IsNil(o.RemainingCredits) {
		return nil, false
	}
	return o.RemainingCredits, true
}

// HasRemainingCredits returns a boolean if a field has been set.
func (o *OrganizationCurrentCost) HasRemainingCredits() bool {
	if o != nil && !IsNil(o.RemainingCredits) {
		return true
	}

	return false
}

// SetRemainingCredits gets a reference to the given RemainingCredits and assigns it to the RemainingCredits field.
func (o *OrganizationCurrentCost) SetRemainingCredits(v RemainingCredits) {
	o.RemainingCredits = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *OrganizationCurrentCost) GetCost() Cost {
	if o == nil || IsNil(o.Cost) {
		var ret Cost
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCurrentCost) GetCostOk() (*Cost, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *OrganizationCurrentCost) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given Cost and assigns it to the Cost field.
func (o *OrganizationCurrentCost) SetCost(v Cost) {
	o.Cost = &v
}

// GetPaidUsage returns the PaidUsage field value if set, zero value otherwise.
func (o *OrganizationCurrentCost) GetPaidUsage() PaidUsage {
	if o == nil || IsNil(o.PaidUsage) {
		var ret PaidUsage
		return ret
	}
	return *o.PaidUsage
}

// GetPaidUsageOk returns a tuple with the PaidUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCurrentCost) GetPaidUsageOk() (*PaidUsage, bool) {
	if o == nil || IsNil(o.PaidUsage) {
		return nil, false
	}
	return o.PaidUsage, true
}

// HasPaidUsage returns a boolean if a field has been set.
func (o *OrganizationCurrentCost) HasPaidUsage() bool {
	if o != nil && !IsNil(o.PaidUsage) {
		return true
	}

	return false
}

// SetPaidUsage gets a reference to the given PaidUsage and assigns it to the PaidUsage field.
func (o *OrganizationCurrentCost) SetPaidUsage(v PaidUsage) {
	o.PaidUsage = &v
}

func (o OrganizationCurrentCost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationCurrentCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.RemainingTrialDay) {
		toSerialize["remaining_trial_day"] = o.RemainingTrialDay
	}
	if !IsNil(o.RemainingCredits) {
		toSerialize["remaining_credits"] = o.RemainingCredits
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.PaidUsage) {
		toSerialize["paid_usage"] = o.PaidUsage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationCurrentCost) UnmarshalJSON(data []byte) (err error) {
	varOrganizationCurrentCost := _OrganizationCurrentCost{}

	err = json.Unmarshal(data, &varOrganizationCurrentCost)

	if err != nil {
		return err
	}

	*o = OrganizationCurrentCost(varOrganizationCurrentCost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "plan")
		delete(additionalProperties, "remaining_trial_day")
		delete(additionalProperties, "remaining_credits")
		delete(additionalProperties, "cost")
		delete(additionalProperties, "paid_usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationCurrentCost struct {
	value *OrganizationCurrentCost
	isSet bool
}

func (v NullableOrganizationCurrentCost) Get() *OrganizationCurrentCost {
	return v.value
}

func (v *NullableOrganizationCurrentCost) Set(val *OrganizationCurrentCost) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCurrentCost) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCurrentCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCurrentCost(val *OrganizationCurrentCost) *NullableOrganizationCurrentCost {
	return &NullableOrganizationCurrentCost{value: val, isSet: true}
}

func (v NullableOrganizationCurrentCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCurrentCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
