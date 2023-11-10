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

// checks if the HelmRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmRequest{}

// HelmRequest struct for HelmRequest
type HelmRequest struct {
	// name is case insensitive
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Maximum number of seconds allowed for helm to run before killing it and mark it as failed
	TimeoutSec *int32 `json:"timeout_sec,omitempty"`
	// Indicates if the 'environment preview option' is enabled.   If enabled, a preview environment will be automatically cloned when `/preview` endpoint is called or when a new commit is updated. If not specified, it takes the value of the `auto_preview` property from the associated environment.
	AutoPreview NullableBool `json:"auto_preview,omitempty"`
	// Specify if the helm will be automatically updated after receiving a new image tag or a new commit according to the source type.
	AutoDeploy bool                   `json:"auto_deploy"`
	Source     HelmRequestAllOfSource `json:"source"`
	// The extra arguments to pass to helm
	Arguments []string `json:"arguments"`
	// If we should allow the chart to deploy object outside his specified namespace. Setting this flag to true, requires special rights
	AllowClusterWideResources *bool                          `json:"allow_cluster_wide_resources,omitempty"`
	ValuesOverride            HelmRequestAllOfValuesOverride `json:"values_override"`
}

// NewHelmRequest instantiates a new HelmRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmRequest(name string, autoDeploy bool, source HelmRequestAllOfSource, arguments []string, valuesOverride HelmRequestAllOfValuesOverride) *HelmRequest {
	this := HelmRequest{}
	this.Name = name
	var timeoutSec int32 = 600
	this.TimeoutSec = &timeoutSec
	this.AutoDeploy = autoDeploy
	this.Source = source
	this.Arguments = arguments
	var allowClusterWideResources bool = false
	this.AllowClusterWideResources = &allowClusterWideResources
	this.ValuesOverride = valuesOverride
	return &this
}

// NewHelmRequestWithDefaults instantiates a new HelmRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmRequestWithDefaults() *HelmRequest {
	this := HelmRequest{}
	var timeoutSec int32 = 600
	this.TimeoutSec = &timeoutSec
	var allowClusterWideResources bool = false
	this.AllowClusterWideResources = &allowClusterWideResources
	return &this
}

// GetName returns the Name field value
func (o *HelmRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HelmRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HelmRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HelmRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HelmRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HelmRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTimeoutSec returns the TimeoutSec field value if set, zero value otherwise.
func (o *HelmRequest) GetTimeoutSec() int32 {
	if o == nil || IsNil(o.TimeoutSec) {
		var ret int32
		return ret
	}
	return *o.TimeoutSec
}

// GetTimeoutSecOk returns a tuple with the TimeoutSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRequest) GetTimeoutSecOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeoutSec) {
		return nil, false
	}
	return o.TimeoutSec, true
}

// HasTimeoutSec returns a boolean if a field has been set.
func (o *HelmRequest) HasTimeoutSec() bool {
	if o != nil && !IsNil(o.TimeoutSec) {
		return true
	}

	return false
}

// SetTimeoutSec gets a reference to the given int32 and assigns it to the TimeoutSec field.
func (o *HelmRequest) SetTimeoutSec(v int32) {
	o.TimeoutSec = &v
}

// GetAutoPreview returns the AutoPreview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmRequest) GetAutoPreview() bool {
	if o == nil || IsNil(o.AutoPreview.Get()) {
		var ret bool
		return ret
	}
	return *o.AutoPreview.Get()
}

// GetAutoPreviewOk returns a tuple with the AutoPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmRequest) GetAutoPreviewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoPreview.Get(), o.AutoPreview.IsSet()
}

// HasAutoPreview returns a boolean if a field has been set.
func (o *HelmRequest) HasAutoPreview() bool {
	if o != nil && o.AutoPreview.IsSet() {
		return true
	}

	return false
}

// SetAutoPreview gets a reference to the given NullableBool and assigns it to the AutoPreview field.
func (o *HelmRequest) SetAutoPreview(v bool) {
	o.AutoPreview.Set(&v)
}

// SetAutoPreviewNil sets the value for AutoPreview to be an explicit nil
func (o *HelmRequest) SetAutoPreviewNil() {
	o.AutoPreview.Set(nil)
}

// UnsetAutoPreview ensures that no value is present for AutoPreview, not even an explicit nil
func (o *HelmRequest) UnsetAutoPreview() {
	o.AutoPreview.Unset()
}

// GetAutoDeploy returns the AutoDeploy field value
func (o *HelmRequest) GetAutoDeploy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoDeploy
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value
// and a boolean to check if the value has been set.
func (o *HelmRequest) GetAutoDeployOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoDeploy, true
}

// SetAutoDeploy sets field value
func (o *HelmRequest) SetAutoDeploy(v bool) {
	o.AutoDeploy = v
}

// GetSource returns the Source field value
func (o *HelmRequest) GetSource() HelmRequestAllOfSource {
	if o == nil {
		var ret HelmRequestAllOfSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *HelmRequest) GetSourceOk() (*HelmRequestAllOfSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *HelmRequest) SetSource(v HelmRequestAllOfSource) {
	o.Source = v
}

// GetArguments returns the Arguments field value
func (o *HelmRequest) GetArguments() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *HelmRequest) GetArgumentsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Arguments, true
}

// SetArguments sets field value
func (o *HelmRequest) SetArguments(v []string) {
	o.Arguments = v
}

// GetAllowClusterWideResources returns the AllowClusterWideResources field value if set, zero value otherwise.
func (o *HelmRequest) GetAllowClusterWideResources() bool {
	if o == nil || IsNil(o.AllowClusterWideResources) {
		var ret bool
		return ret
	}
	return *o.AllowClusterWideResources
}

// GetAllowClusterWideResourcesOk returns a tuple with the AllowClusterWideResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRequest) GetAllowClusterWideResourcesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowClusterWideResources) {
		return nil, false
	}
	return o.AllowClusterWideResources, true
}

// HasAllowClusterWideResources returns a boolean if a field has been set.
func (o *HelmRequest) HasAllowClusterWideResources() bool {
	if o != nil && !IsNil(o.AllowClusterWideResources) {
		return true
	}

	return false
}

// SetAllowClusterWideResources gets a reference to the given bool and assigns it to the AllowClusterWideResources field.
func (o *HelmRequest) SetAllowClusterWideResources(v bool) {
	o.AllowClusterWideResources = &v
}

// GetValuesOverride returns the ValuesOverride field value
func (o *HelmRequest) GetValuesOverride() HelmRequestAllOfValuesOverride {
	if o == nil {
		var ret HelmRequestAllOfValuesOverride
		return ret
	}

	return o.ValuesOverride
}

// GetValuesOverrideOk returns a tuple with the ValuesOverride field value
// and a boolean to check if the value has been set.
func (o *HelmRequest) GetValuesOverrideOk() (*HelmRequestAllOfValuesOverride, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValuesOverride, true
}

// SetValuesOverride sets field value
func (o *HelmRequest) SetValuesOverride(v HelmRequestAllOfValuesOverride) {
	o.ValuesOverride = v
}

func (o HelmRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TimeoutSec) {
		toSerialize["timeout_sec"] = o.TimeoutSec
	}
	if o.AutoPreview.IsSet() {
		toSerialize["auto_preview"] = o.AutoPreview.Get()
	}
	toSerialize["auto_deploy"] = o.AutoDeploy
	toSerialize["source"] = o.Source
	toSerialize["arguments"] = o.Arguments
	if !IsNil(o.AllowClusterWideResources) {
		toSerialize["allow_cluster_wide_resources"] = o.AllowClusterWideResources
	}
	toSerialize["values_override"] = o.ValuesOverride
	return toSerialize, nil
}

type NullableHelmRequest struct {
	value *HelmRequest
	isSet bool
}

func (v NullableHelmRequest) Get() *HelmRequest {
	return v.value
}

func (v *NullableHelmRequest) Set(val *HelmRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmRequest(val *HelmRequest) *NullableHelmRequest {
	return &NullableHelmRequest{value: val, isSet: true}
}

func (v NullableHelmRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
