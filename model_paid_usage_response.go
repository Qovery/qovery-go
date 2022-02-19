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
	"time"
)

// PaidUsageResponse struct for PaidUsageResponse
type PaidUsageResponse struct {
	MaxDeploymentsPerMonth *int32     `json:"max_deployments_per_month,omitempty"`
	ConsumedDeployments    *int32     `json:"consumed_deployments,omitempty"`
	MonthlyPlanCost        *float32   `json:"monthly_plan_cost,omitempty"`
	MonthlyPlanCostInCents *int32     `json:"monthly_plan_cost_in_cents,omitempty"`
	RemainingDeployments   *int32     `json:"remaining_deployments,omitempty"`
	DeploymentsExceeded    *bool      `json:"deployments_exceeded,omitempty"`
	RenewalAt              *time.Time `json:"renewal_at,omitempty"`
}

// NewPaidUsageResponse instantiates a new PaidUsageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaidUsageResponse() *PaidUsageResponse {
	this := PaidUsageResponse{}
	return &this
}

// NewPaidUsageResponseWithDefaults instantiates a new PaidUsageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaidUsageResponseWithDefaults() *PaidUsageResponse {
	this := PaidUsageResponse{}
	return &this
}

// GetMaxDeploymentsPerMonth returns the MaxDeploymentsPerMonth field value if set, zero value otherwise.
func (o *PaidUsageResponse) GetMaxDeploymentsPerMonth() int32 {
	if o == nil || o.MaxDeploymentsPerMonth == nil {
		var ret int32
		return ret
	}
	return *o.MaxDeploymentsPerMonth
}

// GetMaxDeploymentsPerMonthOk returns a tuple with the MaxDeploymentsPerMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidUsageResponse) GetMaxDeploymentsPerMonthOk() (*int32, bool) {
	if o == nil || o.MaxDeploymentsPerMonth == nil {
		return nil, false
	}
	return o.MaxDeploymentsPerMonth, true
}

// HasMaxDeploymentsPerMonth returns a boolean if a field has been set.
func (o *PaidUsageResponse) HasMaxDeploymentsPerMonth() bool {
	if o != nil && o.MaxDeploymentsPerMonth != nil {
		return true
	}

	return false
}

// SetMaxDeploymentsPerMonth gets a reference to the given int32 and assigns it to the MaxDeploymentsPerMonth field.
func (o *PaidUsageResponse) SetMaxDeploymentsPerMonth(v int32) {
	o.MaxDeploymentsPerMonth = &v
}

// GetConsumedDeployments returns the ConsumedDeployments field value if set, zero value otherwise.
func (o *PaidUsageResponse) GetConsumedDeployments() int32 {
	if o == nil || o.ConsumedDeployments == nil {
		var ret int32
		return ret
	}
	return *o.ConsumedDeployments
}

// GetConsumedDeploymentsOk returns a tuple with the ConsumedDeployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidUsageResponse) GetConsumedDeploymentsOk() (*int32, bool) {
	if o == nil || o.ConsumedDeployments == nil {
		return nil, false
	}
	return o.ConsumedDeployments, true
}

// HasConsumedDeployments returns a boolean if a field has been set.
func (o *PaidUsageResponse) HasConsumedDeployments() bool {
	if o != nil && o.ConsumedDeployments != nil {
		return true
	}

	return false
}

// SetConsumedDeployments gets a reference to the given int32 and assigns it to the ConsumedDeployments field.
func (o *PaidUsageResponse) SetConsumedDeployments(v int32) {
	o.ConsumedDeployments = &v
}

// GetMonthlyPlanCost returns the MonthlyPlanCost field value if set, zero value otherwise.
func (o *PaidUsageResponse) GetMonthlyPlanCost() float32 {
	if o == nil || o.MonthlyPlanCost == nil {
		var ret float32
		return ret
	}
	return *o.MonthlyPlanCost
}

// GetMonthlyPlanCostOk returns a tuple with the MonthlyPlanCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidUsageResponse) GetMonthlyPlanCostOk() (*float32, bool) {
	if o == nil || o.MonthlyPlanCost == nil {
		return nil, false
	}
	return o.MonthlyPlanCost, true
}

// HasMonthlyPlanCost returns a boolean if a field has been set.
func (o *PaidUsageResponse) HasMonthlyPlanCost() bool {
	if o != nil && o.MonthlyPlanCost != nil {
		return true
	}

	return false
}

// SetMonthlyPlanCost gets a reference to the given float32 and assigns it to the MonthlyPlanCost field.
func (o *PaidUsageResponse) SetMonthlyPlanCost(v float32) {
	o.MonthlyPlanCost = &v
}

// GetMonthlyPlanCostInCents returns the MonthlyPlanCostInCents field value if set, zero value otherwise.
func (o *PaidUsageResponse) GetMonthlyPlanCostInCents() int32 {
	if o == nil || o.MonthlyPlanCostInCents == nil {
		var ret int32
		return ret
	}
	return *o.MonthlyPlanCostInCents
}

// GetMonthlyPlanCostInCentsOk returns a tuple with the MonthlyPlanCostInCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidUsageResponse) GetMonthlyPlanCostInCentsOk() (*int32, bool) {
	if o == nil || o.MonthlyPlanCostInCents == nil {
		return nil, false
	}
	return o.MonthlyPlanCostInCents, true
}

// HasMonthlyPlanCostInCents returns a boolean if a field has been set.
func (o *PaidUsageResponse) HasMonthlyPlanCostInCents() bool {
	if o != nil && o.MonthlyPlanCostInCents != nil {
		return true
	}

	return false
}

// SetMonthlyPlanCostInCents gets a reference to the given int32 and assigns it to the MonthlyPlanCostInCents field.
func (o *PaidUsageResponse) SetMonthlyPlanCostInCents(v int32) {
	o.MonthlyPlanCostInCents = &v
}

// GetRemainingDeployments returns the RemainingDeployments field value if set, zero value otherwise.
func (o *PaidUsageResponse) GetRemainingDeployments() int32 {
	if o == nil || o.RemainingDeployments == nil {
		var ret int32
		return ret
	}
	return *o.RemainingDeployments
}

// GetRemainingDeploymentsOk returns a tuple with the RemainingDeployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidUsageResponse) GetRemainingDeploymentsOk() (*int32, bool) {
	if o == nil || o.RemainingDeployments == nil {
		return nil, false
	}
	return o.RemainingDeployments, true
}

// HasRemainingDeployments returns a boolean if a field has been set.
func (o *PaidUsageResponse) HasRemainingDeployments() bool {
	if o != nil && o.RemainingDeployments != nil {
		return true
	}

	return false
}

// SetRemainingDeployments gets a reference to the given int32 and assigns it to the RemainingDeployments field.
func (o *PaidUsageResponse) SetRemainingDeployments(v int32) {
	o.RemainingDeployments = &v
}

// GetDeploymentsExceeded returns the DeploymentsExceeded field value if set, zero value otherwise.
func (o *PaidUsageResponse) GetDeploymentsExceeded() bool {
	if o == nil || o.DeploymentsExceeded == nil {
		var ret bool
		return ret
	}
	return *o.DeploymentsExceeded
}

// GetDeploymentsExceededOk returns a tuple with the DeploymentsExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidUsageResponse) GetDeploymentsExceededOk() (*bool, bool) {
	if o == nil || o.DeploymentsExceeded == nil {
		return nil, false
	}
	return o.DeploymentsExceeded, true
}

// HasDeploymentsExceeded returns a boolean if a field has been set.
func (o *PaidUsageResponse) HasDeploymentsExceeded() bool {
	if o != nil && o.DeploymentsExceeded != nil {
		return true
	}

	return false
}

// SetDeploymentsExceeded gets a reference to the given bool and assigns it to the DeploymentsExceeded field.
func (o *PaidUsageResponse) SetDeploymentsExceeded(v bool) {
	o.DeploymentsExceeded = &v
}

// GetRenewalAt returns the RenewalAt field value if set, zero value otherwise.
func (o *PaidUsageResponse) GetRenewalAt() time.Time {
	if o == nil || o.RenewalAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RenewalAt
}

// GetRenewalAtOk returns a tuple with the RenewalAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidUsageResponse) GetRenewalAtOk() (*time.Time, bool) {
	if o == nil || o.RenewalAt == nil {
		return nil, false
	}
	return o.RenewalAt, true
}

// HasRenewalAt returns a boolean if a field has been set.
func (o *PaidUsageResponse) HasRenewalAt() bool {
	if o != nil && o.RenewalAt != nil {
		return true
	}

	return false
}

// SetRenewalAt gets a reference to the given time.Time and assigns it to the RenewalAt field.
func (o *PaidUsageResponse) SetRenewalAt(v time.Time) {
	o.RenewalAt = &v
}

func (o PaidUsageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxDeploymentsPerMonth != nil {
		toSerialize["max_deployments_per_month"] = o.MaxDeploymentsPerMonth
	}
	if o.ConsumedDeployments != nil {
		toSerialize["consumed_deployments"] = o.ConsumedDeployments
	}
	if o.MonthlyPlanCost != nil {
		toSerialize["monthly_plan_cost"] = o.MonthlyPlanCost
	}
	if o.MonthlyPlanCostInCents != nil {
		toSerialize["monthly_plan_cost_in_cents"] = o.MonthlyPlanCostInCents
	}
	if o.RemainingDeployments != nil {
		toSerialize["remaining_deployments"] = o.RemainingDeployments
	}
	if o.DeploymentsExceeded != nil {
		toSerialize["deployments_exceeded"] = o.DeploymentsExceeded
	}
	if o.RenewalAt != nil {
		toSerialize["renewal_at"] = o.RenewalAt
	}
	return json.Marshal(toSerialize)
}

type NullablePaidUsageResponse struct {
	value *PaidUsageResponse
	isSet bool
}

func (v NullablePaidUsageResponse) Get() *PaidUsageResponse {
	return v.value
}

func (v *NullablePaidUsageResponse) Set(val *PaidUsageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaidUsageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaidUsageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaidUsageResponse(val *PaidUsageResponse) *NullablePaidUsageResponse {
	return &NullablePaidUsageResponse{value: val, isSet: true}
}

func (v NullablePaidUsageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaidUsageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
