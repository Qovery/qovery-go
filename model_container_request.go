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

// checks if the ContainerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerRequest{}

// ContainerRequest struct for ContainerRequest
type ContainerRequest struct {
	Storage []ServiceStorageRequestStorageInner `json:"storage,omitempty"`
	Ports   []ServicePortRequestPortsInner      `json:"ports,omitempty"`
	// name is case insensitive
	Name string `json:"name"`
	// give a description to this container
	Description *string `json:"description,omitempty"`
	// id of the linked registry
	RegistryId string `json:"registry_id"`
	// The image name pattern differs according to chosen container registry provider:   * `ECR`: `repository` * `SCALEWAY_CR`: `namespace/image` * `DOCKER_HUB`: `image` or `repository/image` * `PUBLIC_ECR`: `registry_alias/repository`
	ImageName string `json:"image_name"`
	// tag of the image container
	Tag       string   `json:"tag"`
	Arguments []string `json:"arguments,omitempty"`
	// optional entrypoint when launching container
	Entrypoint *string `json:"entrypoint,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *int32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *int32 `json:"memory,omitempty"`
	// Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no container running.
	MinRunningInstances *int32 `json:"min_running_instances,omitempty"`
	// Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.
	MaxRunningInstances *int32      `json:"max_running_instances,omitempty"`
	Healthchecks        Healthcheck `json:"healthchecks"`
	// Indicates if the 'environment preview option' is enabled for this container.   If enabled, a preview environment will be automatically cloned when `/preview` endpoint is called.   If not specified, it takes the value of the `auto_preview` property from the associated environment.
	AutoPreview *bool `json:"auto_preview,omitempty"`
	// Specify if the container will be automatically updated after receiving a new image tag.  The new image tag shall be communicated via the \"Auto Deploy container\" endpoint https://api-doc.qovery.com/#tag/Containers/operation/autoDeployContainerEnvironments
	AutoDeploy NullableBool `json:"auto_deploy,omitempty"`
}

// NewContainerRequest instantiates a new ContainerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRequest(name string, registryId string, imageName string, tag string, healthchecks Healthcheck) *ContainerRequest {
	this := ContainerRequest{}
	this.Name = name
	this.RegistryId = registryId
	this.ImageName = imageName
	this.Tag = tag
	var cpu int32 = 500
	this.Cpu = &cpu
	var memory int32 = 512
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	this.Healthchecks = healthchecks
	return &this
}

// NewContainerRequestWithDefaults instantiates a new ContainerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRequestWithDefaults() *ContainerRequest {
	this := ContainerRequest{}
	var cpu int32 = 500
	this.Cpu = &cpu
	var memory int32 = 512
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	return &this
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ContainerRequest) GetStorage() []ServiceStorageRequestStorageInner {
	if o == nil || IsNil(o.Storage) {
		var ret []ServiceStorageRequestStorageInner
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetStorageOk() ([]ServiceStorageRequestStorageInner, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ContainerRequest) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []ServiceStorageRequestStorageInner and assigns it to the Storage field.
func (o *ContainerRequest) SetStorage(v []ServiceStorageRequestStorageInner) {
	o.Storage = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ContainerRequest) GetPorts() []ServicePortRequestPortsInner {
	if o == nil || IsNil(o.Ports) {
		var ret []ServicePortRequestPortsInner
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetPortsOk() ([]ServicePortRequestPortsInner, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ContainerRequest) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ServicePortRequestPortsInner and assigns it to the Ports field.
func (o *ContainerRequest) SetPorts(v []ServicePortRequestPortsInner) {
	o.Ports = v
}

// GetName returns the Name field value
func (o *ContainerRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ContainerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ContainerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ContainerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetRegistryId returns the RegistryId field value
func (o *ContainerRequest) GetRegistryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetRegistryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistryId, true
}

// SetRegistryId sets field value
func (o *ContainerRequest) SetRegistryId(v string) {
	o.RegistryId = v
}

// GetImageName returns the ImageName field value
func (o *ContainerRequest) GetImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageName, true
}

// SetImageName sets field value
func (o *ContainerRequest) SetImageName(v string) {
	o.ImageName = v
}

// GetTag returns the Tag field value
func (o *ContainerRequest) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *ContainerRequest) SetTag(v string) {
	o.Tag = v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *ContainerRequest) GetArguments() []string {
	if o == nil || IsNil(o.Arguments) {
		var ret []string
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetArgumentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Arguments) {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *ContainerRequest) HasArguments() bool {
	if o != nil && !IsNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []string and assigns it to the Arguments field.
func (o *ContainerRequest) SetArguments(v []string) {
	o.Arguments = v
}

// GetEntrypoint returns the Entrypoint field value if set, zero value otherwise.
func (o *ContainerRequest) GetEntrypoint() string {
	if o == nil || IsNil(o.Entrypoint) {
		var ret string
		return ret
	}
	return *o.Entrypoint
}

// GetEntrypointOk returns a tuple with the Entrypoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetEntrypointOk() (*string, bool) {
	if o == nil || IsNil(o.Entrypoint) {
		return nil, false
	}
	return o.Entrypoint, true
}

// HasEntrypoint returns a boolean if a field has been set.
func (o *ContainerRequest) HasEntrypoint() bool {
	if o != nil && !IsNil(o.Entrypoint) {
		return true
	}

	return false
}

// SetEntrypoint gets a reference to the given string and assigns it to the Entrypoint field.
func (o *ContainerRequest) SetEntrypoint(v string) {
	o.Entrypoint = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ContainerRequest) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ContainerRequest) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *ContainerRequest) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ContainerRequest) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ContainerRequest) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *ContainerRequest) SetMemory(v int32) {
	o.Memory = &v
}

// GetMinRunningInstances returns the MinRunningInstances field value if set, zero value otherwise.
func (o *ContainerRequest) GetMinRunningInstances() int32 {
	if o == nil || IsNil(o.MinRunningInstances) {
		var ret int32
		return ret
	}
	return *o.MinRunningInstances
}

// GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetMinRunningInstancesOk() (*int32, bool) {
	if o == nil || IsNil(o.MinRunningInstances) {
		return nil, false
	}
	return o.MinRunningInstances, true
}

// HasMinRunningInstances returns a boolean if a field has been set.
func (o *ContainerRequest) HasMinRunningInstances() bool {
	if o != nil && !IsNil(o.MinRunningInstances) {
		return true
	}

	return false
}

// SetMinRunningInstances gets a reference to the given int32 and assigns it to the MinRunningInstances field.
func (o *ContainerRequest) SetMinRunningInstances(v int32) {
	o.MinRunningInstances = &v
}

// GetMaxRunningInstances returns the MaxRunningInstances field value if set, zero value otherwise.
func (o *ContainerRequest) GetMaxRunningInstances() int32 {
	if o == nil || IsNil(o.MaxRunningInstances) {
		var ret int32
		return ret
	}
	return *o.MaxRunningInstances
}

// GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetMaxRunningInstancesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRunningInstances) {
		return nil, false
	}
	return o.MaxRunningInstances, true
}

// HasMaxRunningInstances returns a boolean if a field has been set.
func (o *ContainerRequest) HasMaxRunningInstances() bool {
	if o != nil && !IsNil(o.MaxRunningInstances) {
		return true
	}

	return false
}

// SetMaxRunningInstances gets a reference to the given int32 and assigns it to the MaxRunningInstances field.
func (o *ContainerRequest) SetMaxRunningInstances(v int32) {
	o.MaxRunningInstances = &v
}

// GetHealthchecks returns the Healthchecks field value
func (o *ContainerRequest) GetHealthchecks() Healthcheck {
	if o == nil {
		var ret Healthcheck
		return ret
	}

	return o.Healthchecks
}

// GetHealthchecksOk returns a tuple with the Healthchecks field value
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetHealthchecksOk() (*Healthcheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Healthchecks, true
}

// SetHealthchecks sets field value
func (o *ContainerRequest) SetHealthchecks(v Healthcheck) {
	o.Healthchecks = v
}

// GetAutoPreview returns the AutoPreview field value if set, zero value otherwise.
func (o *ContainerRequest) GetAutoPreview() bool {
	if o == nil || IsNil(o.AutoPreview) {
		var ret bool
		return ret
	}
	return *o.AutoPreview
}

// GetAutoPreviewOk returns a tuple with the AutoPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequest) GetAutoPreviewOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoPreview) {
		return nil, false
	}
	return o.AutoPreview, true
}

// HasAutoPreview returns a boolean if a field has been set.
func (o *ContainerRequest) HasAutoPreview() bool {
	if o != nil && !IsNil(o.AutoPreview) {
		return true
	}

	return false
}

// SetAutoPreview gets a reference to the given bool and assigns it to the AutoPreview field.
func (o *ContainerRequest) SetAutoPreview(v bool) {
	o.AutoPreview = &v
}

// GetAutoDeploy returns the AutoDeploy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerRequest) GetAutoDeploy() bool {
	if o == nil || IsNil(o.AutoDeploy.Get()) {
		var ret bool
		return ret
	}
	return *o.AutoDeploy.Get()
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerRequest) GetAutoDeployOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoDeploy.Get(), o.AutoDeploy.IsSet()
}

// HasAutoDeploy returns a boolean if a field has been set.
func (o *ContainerRequest) HasAutoDeploy() bool {
	if o != nil && o.AutoDeploy.IsSet() {
		return true
	}

	return false
}

// SetAutoDeploy gets a reference to the given NullableBool and assigns it to the AutoDeploy field.
func (o *ContainerRequest) SetAutoDeploy(v bool) {
	o.AutoDeploy.Set(&v)
}

// SetAutoDeployNil sets the value for AutoDeploy to be an explicit nil
func (o *ContainerRequest) SetAutoDeployNil() {
	o.AutoDeploy.Set(nil)
}

// UnsetAutoDeploy ensures that no value is present for AutoDeploy, not even an explicit nil
func (o *ContainerRequest) UnsetAutoDeploy() {
	o.AutoDeploy.Unset()
}

func (o ContainerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["registry_id"] = o.RegistryId
	toSerialize["image_name"] = o.ImageName
	toSerialize["tag"] = o.Tag
	if !IsNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}
	if !IsNil(o.Entrypoint) {
		toSerialize["entrypoint"] = o.Entrypoint
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.MinRunningInstances) {
		toSerialize["min_running_instances"] = o.MinRunningInstances
	}
	if !IsNil(o.MaxRunningInstances) {
		toSerialize["max_running_instances"] = o.MaxRunningInstances
	}
	toSerialize["healthchecks"] = o.Healthchecks
	if !IsNil(o.AutoPreview) {
		toSerialize["auto_preview"] = o.AutoPreview
	}
	if o.AutoDeploy.IsSet() {
		toSerialize["auto_deploy"] = o.AutoDeploy.Get()
	}
	return toSerialize, nil
}

type NullableContainerRequest struct {
	value *ContainerRequest
	isSet bool
}

func (v NullableContainerRequest) Get() *ContainerRequest {
	return v.value
}

func (v *NullableContainerRequest) Set(val *ContainerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRequest(val *ContainerRequest) *NullableContainerRequest {
	return &NullableContainerRequest{value: val, isSet: true}
}

func (v NullableContainerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
