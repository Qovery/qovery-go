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
	"time"
)

// ClusterResponse struct for ClusterResponse
type ClusterResponse struct {
	// This is an estimation of the cost this cluster will represent on your cloud proider bill, based on your current configuration
	EstimatedCloudProviderCost *float32   `json:"estimated_cloud_provider_cost,omitempty"`
	Status                     *string    `json:"status,omitempty"`
	HasAccess                  *bool      `json:"has_access,omitempty"`
	Version                    *string    `json:"version,omitempty"`
	IsDefault                  *bool      `json:"is_default,omitempty"`
	Id                         string     `json:"id"`
	CreatedAt                  time.Time  `json:"created_at"`
	UpdatedAt                  *time.Time `json:"updated_at,omitempty"`
	// name is case-insensitive
	Name          string         `json:"name"`
	Description   NullableString `json:"description,omitempty"`
	CloudProvider string         `json:"cloud_provider"`
	Region        string         `json:"region"`
	AutoUpdate    *bool          `json:"auto_update,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *float32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory              *float32             `json:"memory,omitempty"`
	MinRunningNodes     *int32               `json:"min_running_nodes,omitempty"`
	MaxRunningNodes     *int32               `json:"max_running_nodes,omitempty"`
	Title               *string              `json:"title,omitempty"`
	CostPerMonthInCents NullableInt32        `json:"cost_per_month_in_cents,omitempty"`
	CostPerMonth        NullableFloat32      `json:"cost_per_month,omitempty"`
	CurrencyCode        NullableString       `json:"currency_code,omitempty"`
	ValueType           *string              `json:"value_type,omitempty"`
	Value               NullableString       `json:"value,omitempty"`
	IsValueUpdatable    *bool                `json:"is_value_updatable,omitempty"`
	AcceptedValues      []interface{} `json:"accepted_values,omitempty"`
}

// NewClusterResponse instantiates a new ClusterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterResponse(id string, createdAt time.Time, name string, cloudProvider string, region string) *ClusterResponse {
	this := ClusterResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Name = name
	this.CloudProvider = cloudProvider
	this.Region = region
	var cpu float32 = 250
	this.Cpu = &cpu
	var memory float32 = 256
	this.Memory = &memory
	var minRunningNodes int32 = 1
	this.MinRunningNodes = &minRunningNodes
	var maxRunningNodes int32 = 1
	this.MaxRunningNodes = &maxRunningNodes
	var isValueUpdatable bool = false
	this.IsValueUpdatable = &isValueUpdatable
	return &this
}

// NewClusterResponseWithDefaults instantiates a new ClusterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterResponseWithDefaults() *ClusterResponse {
	this := ClusterResponse{}
	var cpu float32 = 250
	this.Cpu = &cpu
	var memory float32 = 256
	this.Memory = &memory
	var minRunningNodes int32 = 1
	this.MinRunningNodes = &minRunningNodes
	var maxRunningNodes int32 = 1
	this.MaxRunningNodes = &maxRunningNodes
	var isValueUpdatable bool = false
	this.IsValueUpdatable = &isValueUpdatable
	return &this
}

// GetEstimatedCloudProviderCost returns the EstimatedCloudProviderCost field value if set, zero value otherwise.
func (o *ClusterResponse) GetEstimatedCloudProviderCost() float32 {
	if o == nil || o.EstimatedCloudProviderCost == nil {
		var ret float32
		return ret
	}
	return *o.EstimatedCloudProviderCost
}

// GetEstimatedCloudProviderCostOk returns a tuple with the EstimatedCloudProviderCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetEstimatedCloudProviderCostOk() (*float32, bool) {
	if o == nil || o.EstimatedCloudProviderCost == nil {
		return nil, false
	}
	return o.EstimatedCloudProviderCost, true
}

// HasEstimatedCloudProviderCost returns a boolean if a field has been set.
func (o *ClusterResponse) HasEstimatedCloudProviderCost() bool {
	if o != nil && o.EstimatedCloudProviderCost != nil {
		return true
	}

	return false
}

// SetEstimatedCloudProviderCost gets a reference to the given float32 and assigns it to the EstimatedCloudProviderCost field.
func (o *ClusterResponse) SetEstimatedCloudProviderCost(v float32) {
	o.EstimatedCloudProviderCost = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ClusterResponse) SetStatus(v string) {
	o.Status = &v
}

// GetHasAccess returns the HasAccess field value if set, zero value otherwise.
func (o *ClusterResponse) GetHasAccess() bool {
	if o == nil || o.HasAccess == nil {
		var ret bool
		return ret
	}
	return *o.HasAccess
}

// GetHasAccessOk returns a tuple with the HasAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetHasAccessOk() (*bool, bool) {
	if o == nil || o.HasAccess == nil {
		return nil, false
	}
	return o.HasAccess, true
}

// HasHasAccess returns a boolean if a field has been set.
func (o *ClusterResponse) HasHasAccess() bool {
	if o != nil && o.HasAccess != nil {
		return true
	}

	return false
}

// SetHasAccess gets a reference to the given bool and assigns it to the HasAccess field.
func (o *ClusterResponse) SetHasAccess(v bool) {
	o.HasAccess = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ClusterResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ClusterResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ClusterResponse) SetVersion(v string) {
	o.Version = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ClusterResponse) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ClusterResponse) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ClusterResponse) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetId returns the Id field value
func (o *ClusterResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClusterResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ClusterResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ClusterResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ClusterResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ClusterResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ClusterResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value
func (o *ClusterResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ClusterResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ClusterResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ClusterResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetCloudProvider returns the CloudProvider field value
func (o *ClusterResponse) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ClusterResponse) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetRegion returns the Region field value
func (o *ClusterResponse) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ClusterResponse) SetRegion(v string) {
	o.Region = v
}

// GetAutoUpdate returns the AutoUpdate field value if set, zero value otherwise.
func (o *ClusterResponse) GetAutoUpdate() bool {
	if o == nil || o.AutoUpdate == nil {
		var ret bool
		return ret
	}
	return *o.AutoUpdate
}

// GetAutoUpdateOk returns a tuple with the AutoUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetAutoUpdateOk() (*bool, bool) {
	if o == nil || o.AutoUpdate == nil {
		return nil, false
	}
	return o.AutoUpdate, true
}

// HasAutoUpdate returns a boolean if a field has been set.
func (o *ClusterResponse) HasAutoUpdate() bool {
	if o != nil && o.AutoUpdate != nil {
		return true
	}

	return false
}

// SetAutoUpdate gets a reference to the given bool and assigns it to the AutoUpdate field.
func (o *ClusterResponse) SetAutoUpdate(v bool) {
	o.AutoUpdate = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ClusterResponse) GetCpu() float32 {
	if o == nil || o.Cpu == nil {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetCpuOk() (*float32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ClusterResponse) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the Cpu field.
func (o *ClusterResponse) SetCpu(v float32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ClusterResponse) GetMemory() float32 {
	if o == nil || o.Memory == nil {
		var ret float32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetMemoryOk() (*float32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ClusterResponse) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given float32 and assigns it to the Memory field.
func (o *ClusterResponse) SetMemory(v float32) {
	o.Memory = &v
}

// GetMinRunningNodes returns the MinRunningNodes field value if set, zero value otherwise.
func (o *ClusterResponse) GetMinRunningNodes() int32 {
	if o == nil || o.MinRunningNodes == nil {
		var ret int32
		return ret
	}
	return *o.MinRunningNodes
}

// GetMinRunningNodesOk returns a tuple with the MinRunningNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetMinRunningNodesOk() (*int32, bool) {
	if o == nil || o.MinRunningNodes == nil {
		return nil, false
	}
	return o.MinRunningNodes, true
}

// HasMinRunningNodes returns a boolean if a field has been set.
func (o *ClusterResponse) HasMinRunningNodes() bool {
	if o != nil && o.MinRunningNodes != nil {
		return true
	}

	return false
}

// SetMinRunningNodes gets a reference to the given int32 and assigns it to the MinRunningNodes field.
func (o *ClusterResponse) SetMinRunningNodes(v int32) {
	o.MinRunningNodes = &v
}

// GetMaxRunningNodes returns the MaxRunningNodes field value if set, zero value otherwise.
func (o *ClusterResponse) GetMaxRunningNodes() int32 {
	if o == nil || o.MaxRunningNodes == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunningNodes
}

// GetMaxRunningNodesOk returns a tuple with the MaxRunningNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetMaxRunningNodesOk() (*int32, bool) {
	if o == nil || o.MaxRunningNodes == nil {
		return nil, false
	}
	return o.MaxRunningNodes, true
}

// HasMaxRunningNodes returns a boolean if a field has been set.
func (o *ClusterResponse) HasMaxRunningNodes() bool {
	if o != nil && o.MaxRunningNodes != nil {
		return true
	}

	return false
}

// SetMaxRunningNodes gets a reference to the given int32 and assigns it to the MaxRunningNodes field.
func (o *ClusterResponse) SetMaxRunningNodes(v int32) {
	o.MaxRunningNodes = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ClusterResponse) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ClusterResponse) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ClusterResponse) SetTitle(v string) {
	o.Title = &v
}

// GetCostPerMonthInCents returns the CostPerMonthInCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterResponse) GetCostPerMonthInCents() int32 {
	if o == nil || o.CostPerMonthInCents.Get() == nil {
		var ret int32
		return ret
	}
	return *o.CostPerMonthInCents.Get()
}

// GetCostPerMonthInCentsOk returns a tuple with the CostPerMonthInCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterResponse) GetCostPerMonthInCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CostPerMonthInCents.Get(), o.CostPerMonthInCents.IsSet()
}

// HasCostPerMonthInCents returns a boolean if a field has been set.
func (o *ClusterResponse) HasCostPerMonthInCents() bool {
	if o != nil && o.CostPerMonthInCents.IsSet() {
		return true
	}

	return false
}

// SetCostPerMonthInCents gets a reference to the given NullableInt32 and assigns it to the CostPerMonthInCents field.
func (o *ClusterResponse) SetCostPerMonthInCents(v int32) {
	o.CostPerMonthInCents.Set(&v)
}

// SetCostPerMonthInCentsNil sets the value for CostPerMonthInCents to be an explicit nil
func (o *ClusterResponse) SetCostPerMonthInCentsNil() {
	o.CostPerMonthInCents.Set(nil)
}

// UnsetCostPerMonthInCents ensures that no value is present for CostPerMonthInCents, not even an explicit nil
func (o *ClusterResponse) UnsetCostPerMonthInCents() {
	o.CostPerMonthInCents.Unset()
}

// GetCostPerMonth returns the CostPerMonth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterResponse) GetCostPerMonth() float32 {
	if o == nil || o.CostPerMonth.Get() == nil {
		var ret float32
		return ret
	}
	return *o.CostPerMonth.Get()
}

// GetCostPerMonthOk returns a tuple with the CostPerMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterResponse) GetCostPerMonthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CostPerMonth.Get(), o.CostPerMonth.IsSet()
}

// HasCostPerMonth returns a boolean if a field has been set.
func (o *ClusterResponse) HasCostPerMonth() bool {
	if o != nil && o.CostPerMonth.IsSet() {
		return true
	}

	return false
}

// SetCostPerMonth gets a reference to the given NullableFloat32 and assigns it to the CostPerMonth field.
func (o *ClusterResponse) SetCostPerMonth(v float32) {
	o.CostPerMonth.Set(&v)
}

// SetCostPerMonthNil sets the value for CostPerMonth to be an explicit nil
func (o *ClusterResponse) SetCostPerMonthNil() {
	o.CostPerMonth.Set(nil)
}

// UnsetCostPerMonth ensures that no value is present for CostPerMonth, not even an explicit nil
func (o *ClusterResponse) UnsetCostPerMonth() {
	o.CostPerMonth.Unset()
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterResponse) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode.Get()
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyCode.Get(), o.CurrencyCode.IsSet()
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ClusterResponse) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given NullableString and assigns it to the CurrencyCode field.
func (o *ClusterResponse) SetCurrencyCode(v string) {
	o.CurrencyCode.Set(&v)
}

// SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil
func (o *ClusterResponse) SetCurrencyCodeNil() {
	o.CurrencyCode.Set(nil)
}

// UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
func (o *ClusterResponse) UnsetCurrencyCode() {
	o.CurrencyCode.Unset()
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *ClusterResponse) GetValueType() string {
	if o == nil || o.ValueType == nil {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetValueTypeOk() (*string, bool) {
	if o == nil || o.ValueType == nil {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *ClusterResponse) HasValueType() bool {
	if o != nil && o.ValueType != nil {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *ClusterResponse) SetValueType(v string) {
	o.ValueType = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterResponse) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ClusterResponse) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ClusterResponse) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *ClusterResponse) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ClusterResponse) UnsetValue() {
	o.Value.Unset()
}

// GetIsValueUpdatable returns the IsValueUpdatable field value if set, zero value otherwise.
func (o *ClusterResponse) GetIsValueUpdatable() bool {
	if o == nil || o.IsValueUpdatable == nil {
		var ret bool
		return ret
	}
	return *o.IsValueUpdatable
}

// GetIsValueUpdatableOk returns a tuple with the IsValueUpdatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetIsValueUpdatableOk() (*bool, bool) {
	if o == nil || o.IsValueUpdatable == nil {
		return nil, false
	}
	return o.IsValueUpdatable, true
}

// HasIsValueUpdatable returns a boolean if a field has been set.
func (o *ClusterResponse) HasIsValueUpdatable() bool {
	if o != nil && o.IsValueUpdatable != nil {
		return true
	}

	return false
}

// SetIsValueUpdatable gets a reference to the given bool and assigns it to the IsValueUpdatable field.
func (o *ClusterResponse) SetIsValueUpdatable(v bool) {
	o.IsValueUpdatable = &v
}

// GetAcceptedValues returns the AcceptedValues field value if set, zero value otherwise.
func (o *ClusterResponse) GetAcceptedValues() []interface{} {
	if o == nil || o.AcceptedValues == nil {
		var ret []interface{}
		return ret
	}
	return o.AcceptedValues
}

// GetAcceptedValuesOk returns a tuple with the AcceptedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetAcceptedValuesOk() ([]interface{}, bool) {
	if o == nil || o.AcceptedValues == nil {
		return nil, false
	}
	return o.AcceptedValues, true
}

// HasAcceptedValues returns a boolean if a field has been set.
func (o *ClusterResponse) HasAcceptedValues() bool {
	if o != nil && o.AcceptedValues != nil {
		return true
	}

	return false
}

// SetAcceptedValues gets a reference to the given []interface{} and assigns it to the AcceptedValues field.
func (o *ClusterResponse) SetAcceptedValues(v []interface{}) {
	o.AcceptedValues = v
}

func (o ClusterResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EstimatedCloudProviderCost != nil {
		toSerialize["estimated_cloud_provider_cost"] = o.EstimatedCloudProviderCost
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.HasAccess != nil {
		toSerialize["has_access"] = o.HasAccess
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.IsDefault != nil {
		toSerialize["is_default"] = o.IsDefault
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
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if o.AutoUpdate != nil {
		toSerialize["auto_update"] = o.AutoUpdate
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.MinRunningNodes != nil {
		toSerialize["min_running_nodes"] = o.MinRunningNodes
	}
	if o.MaxRunningNodes != nil {
		toSerialize["max_running_nodes"] = o.MaxRunningNodes
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
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
	if o.ValueType != nil {
		toSerialize["value_type"] = o.ValueType
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.IsValueUpdatable != nil {
		toSerialize["is_value_updatable"] = o.IsValueUpdatable
	}
	if o.AcceptedValues != nil {
		toSerialize["accepted_values"] = o.AcceptedValues
	}
	return json.Marshal(toSerialize)
}

type NullableClusterResponse struct {
	value *ClusterResponse
	isSet bool
}

func (v NullableClusterResponse) Get() *ClusterResponse {
	return v.value
}

func (v *NullableClusterResponse) Set(val *ClusterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterResponse(val *ClusterResponse) *NullableClusterResponse {
	return &NullableClusterResponse{value: val, isSet: true}
}

func (v NullableClusterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
