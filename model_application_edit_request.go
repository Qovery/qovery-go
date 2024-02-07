/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ApplicationEditRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationEditRequest{}

// ApplicationEditRequest struct for ApplicationEditRequest
type ApplicationEditRequest struct {
	Storage []ServiceStorageRequestStorageInner `json:"storage,omitempty"`
	// name is case insensitive
	Name *string `json:"name,omitempty"`
	// give a description to this application
	Description   *string                          `json:"description,omitempty"`
	GitRepository *ApplicationGitRepositoryRequest `json:"git_repository,omitempty"`
	BuildMode     *BuildModeEnum                   `json:"build_mode,omitempty"`
	// The path of the associated Dockerfile
	DockerfilePath    *string                       `json:"dockerfile_path,omitempty"`
	BuildpackLanguage NullableBuildPackLanguageEnum `json:"buildpack_language,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *int32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *int32 `json:"memory,omitempty"`
	// Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no application running.
	MinRunningInstances *int32 `json:"min_running_instances,omitempty"`
	// Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.
	MaxRunningInstances *int32      `json:"max_running_instances,omitempty"`
	Healthchecks        Healthcheck `json:"healthchecks"`
	// Specify if the environment preview option is activated or not for this application.   If activated, a preview environment will be automatically cloned at each pull request.   If not specified, it takes the value of the `auto_preview` property from the associated environment.
	AutoPreview *bool         `json:"auto_preview,omitempty"`
	Ports       []ServicePort `json:"ports,omitempty"`
	Arguments   []string      `json:"arguments,omitempty"`
	// optional entrypoint when launching container
	Entrypoint *string `json:"entrypoint,omitempty"`
	// Specify if the application will be automatically updated after receiving a new commit.
	AutoDeploy NullableBool `json:"auto_deploy,omitempty"`
}

type _ApplicationEditRequest ApplicationEditRequest

// NewApplicationEditRequest instantiates a new ApplicationEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationEditRequest(healthchecks Healthcheck) *ApplicationEditRequest {
	this := ApplicationEditRequest{}
	var buildMode BuildModeEnum = BUILDMODEENUM_BUILDPACKS
	this.BuildMode = &buildMode
	var cpu int32 = 500
	this.Cpu = &cpu
	var memory int32 = 512
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	this.Healthchecks = healthchecks
	var autoPreview bool = true
	this.AutoPreview = &autoPreview
	return &this
}

// NewApplicationEditRequestWithDefaults instantiates a new ApplicationEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationEditRequestWithDefaults() *ApplicationEditRequest {
	this := ApplicationEditRequest{}
	var buildMode BuildModeEnum = BUILDMODEENUM_BUILDPACKS
	this.BuildMode = &buildMode
	var cpu int32 = 500
	this.Cpu = &cpu
	var memory int32 = 512
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	var autoPreview bool = true
	this.AutoPreview = &autoPreview
	return &this
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetStorage() []ServiceStorageRequestStorageInner {
	if o == nil || IsNil(o.Storage) {
		var ret []ServiceStorageRequestStorageInner
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetStorageOk() ([]ServiceStorageRequestStorageInner, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []ServiceStorageRequestStorageInner and assigns it to the Storage field.
func (o *ApplicationEditRequest) SetStorage(v []ServiceStorageRequestStorageInner) {
	o.Storage = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationEditRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationEditRequest) SetDescription(v string) {
	o.Description = &v
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetGitRepository() ApplicationGitRepositoryRequest {
	if o == nil || IsNil(o.GitRepository) {
		var ret ApplicationGitRepositoryRequest
		return ret
	}
	return *o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetGitRepositoryOk() (*ApplicationGitRepositoryRequest, bool) {
	if o == nil || IsNil(o.GitRepository) {
		return nil, false
	}
	return o.GitRepository, true
}

// HasGitRepository returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasGitRepository() bool {
	if o != nil && !IsNil(o.GitRepository) {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given ApplicationGitRepositoryRequest and assigns it to the GitRepository field.
func (o *ApplicationEditRequest) SetGitRepository(v ApplicationGitRepositoryRequest) {
	o.GitRepository = &v
}

// GetBuildMode returns the BuildMode field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetBuildMode() BuildModeEnum {
	if o == nil || IsNil(o.BuildMode) {
		var ret BuildModeEnum
		return ret
	}
	return *o.BuildMode
}

// GetBuildModeOk returns a tuple with the BuildMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetBuildModeOk() (*BuildModeEnum, bool) {
	if o == nil || IsNil(o.BuildMode) {
		return nil, false
	}
	return o.BuildMode, true
}

// HasBuildMode returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasBuildMode() bool {
	if o != nil && !IsNil(o.BuildMode) {
		return true
	}

	return false
}

// SetBuildMode gets a reference to the given BuildModeEnum and assigns it to the BuildMode field.
func (o *ApplicationEditRequest) SetBuildMode(v BuildModeEnum) {
	o.BuildMode = &v
}

// GetDockerfilePath returns the DockerfilePath field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetDockerfilePath() string {
	if o == nil || IsNil(o.DockerfilePath) {
		var ret string
		return ret
	}
	return *o.DockerfilePath
}

// GetDockerfilePathOk returns a tuple with the DockerfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetDockerfilePathOk() (*string, bool) {
	if o == nil || IsNil(o.DockerfilePath) {
		return nil, false
	}
	return o.DockerfilePath, true
}

// HasDockerfilePath returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasDockerfilePath() bool {
	if o != nil && !IsNil(o.DockerfilePath) {
		return true
	}

	return false
}

// SetDockerfilePath gets a reference to the given string and assigns it to the DockerfilePath field.
func (o *ApplicationEditRequest) SetDockerfilePath(v string) {
	o.DockerfilePath = &v
}

// GetBuildpackLanguage returns the BuildpackLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationEditRequest) GetBuildpackLanguage() BuildPackLanguageEnum {
	if o == nil || IsNil(o.BuildpackLanguage.Get()) {
		var ret BuildPackLanguageEnum
		return ret
	}
	return *o.BuildpackLanguage.Get()
}

// GetBuildpackLanguageOk returns a tuple with the BuildpackLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationEditRequest) GetBuildpackLanguageOk() (*BuildPackLanguageEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildpackLanguage.Get(), o.BuildpackLanguage.IsSet()
}

// HasBuildpackLanguage returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasBuildpackLanguage() bool {
	if o != nil && o.BuildpackLanguage.IsSet() {
		return true
	}

	return false
}

// SetBuildpackLanguage gets a reference to the given NullableBuildPackLanguageEnum and assigns it to the BuildpackLanguage field.
func (o *ApplicationEditRequest) SetBuildpackLanguage(v BuildPackLanguageEnum) {
	o.BuildpackLanguage.Set(&v)
}

// SetBuildpackLanguageNil sets the value for BuildpackLanguage to be an explicit nil
func (o *ApplicationEditRequest) SetBuildpackLanguageNil() {
	o.BuildpackLanguage.Set(nil)
}

// UnsetBuildpackLanguage ensures that no value is present for BuildpackLanguage, not even an explicit nil
func (o *ApplicationEditRequest) UnsetBuildpackLanguage() {
	o.BuildpackLanguage.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *ApplicationEditRequest) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *ApplicationEditRequest) SetMemory(v int32) {
	o.Memory = &v
}

// GetMinRunningInstances returns the MinRunningInstances field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetMinRunningInstances() int32 {
	if o == nil || IsNil(o.MinRunningInstances) {
		var ret int32
		return ret
	}
	return *o.MinRunningInstances
}

// GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetMinRunningInstancesOk() (*int32, bool) {
	if o == nil || IsNil(o.MinRunningInstances) {
		return nil, false
	}
	return o.MinRunningInstances, true
}

// HasMinRunningInstances returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasMinRunningInstances() bool {
	if o != nil && !IsNil(o.MinRunningInstances) {
		return true
	}

	return false
}

// SetMinRunningInstances gets a reference to the given int32 and assigns it to the MinRunningInstances field.
func (o *ApplicationEditRequest) SetMinRunningInstances(v int32) {
	o.MinRunningInstances = &v
}

// GetMaxRunningInstances returns the MaxRunningInstances field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetMaxRunningInstances() int32 {
	if o == nil || IsNil(o.MaxRunningInstances) {
		var ret int32
		return ret
	}
	return *o.MaxRunningInstances
}

// GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetMaxRunningInstancesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRunningInstances) {
		return nil, false
	}
	return o.MaxRunningInstances, true
}

// HasMaxRunningInstances returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasMaxRunningInstances() bool {
	if o != nil && !IsNil(o.MaxRunningInstances) {
		return true
	}

	return false
}

// SetMaxRunningInstances gets a reference to the given int32 and assigns it to the MaxRunningInstances field.
func (o *ApplicationEditRequest) SetMaxRunningInstances(v int32) {
	o.MaxRunningInstances = &v
}

// GetHealthchecks returns the Healthchecks field value
func (o *ApplicationEditRequest) GetHealthchecks() Healthcheck {
	if o == nil {
		var ret Healthcheck
		return ret
	}

	return o.Healthchecks
}

// GetHealthchecksOk returns a tuple with the Healthchecks field value
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetHealthchecksOk() (*Healthcheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Healthchecks, true
}

// SetHealthchecks sets field value
func (o *ApplicationEditRequest) SetHealthchecks(v Healthcheck) {
	o.Healthchecks = v
}

// GetAutoPreview returns the AutoPreview field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetAutoPreview() bool {
	if o == nil || IsNil(o.AutoPreview) {
		var ret bool
		return ret
	}
	return *o.AutoPreview
}

// GetAutoPreviewOk returns a tuple with the AutoPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetAutoPreviewOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoPreview) {
		return nil, false
	}
	return o.AutoPreview, true
}

// HasAutoPreview returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasAutoPreview() bool {
	if o != nil && !IsNil(o.AutoPreview) {
		return true
	}

	return false
}

// SetAutoPreview gets a reference to the given bool and assigns it to the AutoPreview field.
func (o *ApplicationEditRequest) SetAutoPreview(v bool) {
	o.AutoPreview = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetPorts() []ServicePort {
	if o == nil || IsNil(o.Ports) {
		var ret []ServicePort
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetPortsOk() ([]ServicePort, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ServicePort and assigns it to the Ports field.
func (o *ApplicationEditRequest) SetPorts(v []ServicePort) {
	o.Ports = v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetArguments() []string {
	if o == nil || IsNil(o.Arguments) {
		var ret []string
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetArgumentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Arguments) {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasArguments() bool {
	if o != nil && !IsNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []string and assigns it to the Arguments field.
func (o *ApplicationEditRequest) SetArguments(v []string) {
	o.Arguments = v
}

// GetEntrypoint returns the Entrypoint field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetEntrypoint() string {
	if o == nil || IsNil(o.Entrypoint) {
		var ret string
		return ret
	}
	return *o.Entrypoint
}

// GetEntrypointOk returns a tuple with the Entrypoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetEntrypointOk() (*string, bool) {
	if o == nil || IsNil(o.Entrypoint) {
		return nil, false
	}
	return o.Entrypoint, true
}

// HasEntrypoint returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasEntrypoint() bool {
	if o != nil && !IsNil(o.Entrypoint) {
		return true
	}

	return false
}

// SetEntrypoint gets a reference to the given string and assigns it to the Entrypoint field.
func (o *ApplicationEditRequest) SetEntrypoint(v string) {
	o.Entrypoint = &v
}

// GetAutoDeploy returns the AutoDeploy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationEditRequest) GetAutoDeploy() bool {
	if o == nil || IsNil(o.AutoDeploy.Get()) {
		var ret bool
		return ret
	}
	return *o.AutoDeploy.Get()
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationEditRequest) GetAutoDeployOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoDeploy.Get(), o.AutoDeploy.IsSet()
}

// HasAutoDeploy returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasAutoDeploy() bool {
	if o != nil && o.AutoDeploy.IsSet() {
		return true
	}

	return false
}

// SetAutoDeploy gets a reference to the given NullableBool and assigns it to the AutoDeploy field.
func (o *ApplicationEditRequest) SetAutoDeploy(v bool) {
	o.AutoDeploy.Set(&v)
}

// SetAutoDeployNil sets the value for AutoDeploy to be an explicit nil
func (o *ApplicationEditRequest) SetAutoDeployNil() {
	o.AutoDeploy.Set(nil)
}

// UnsetAutoDeploy ensures that no value is present for AutoDeploy, not even an explicit nil
func (o *ApplicationEditRequest) UnsetAutoDeploy() {
	o.AutoDeploy.Unset()
}

func (o ApplicationEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationEditRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.GitRepository) {
		toSerialize["git_repository"] = o.GitRepository
	}
	if !IsNil(o.BuildMode) {
		toSerialize["build_mode"] = o.BuildMode
	}
	if !IsNil(o.DockerfilePath) {
		toSerialize["dockerfile_path"] = o.DockerfilePath
	}
	if o.BuildpackLanguage.IsSet() {
		toSerialize["buildpack_language"] = o.BuildpackLanguage.Get()
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
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}
	if !IsNil(o.Entrypoint) {
		toSerialize["entrypoint"] = o.Entrypoint
	}
	if o.AutoDeploy.IsSet() {
		toSerialize["auto_deploy"] = o.AutoDeploy.Get()
	}
	return toSerialize, nil
}

func (o *ApplicationEditRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"healthchecks",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApplicationEditRequest := _ApplicationEditRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationEditRequest)

	if err != nil {
		return err
	}

	*o = ApplicationEditRequest(varApplicationEditRequest)

	return err
}

type NullableApplicationEditRequest struct {
	value *ApplicationEditRequest
	isSet bool
}

func (v NullableApplicationEditRequest) Get() *ApplicationEditRequest {
	return v.value
}

func (v *NullableApplicationEditRequest) Set(val *ApplicationEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationEditRequest(val *ApplicationEditRequest) *NullableApplicationEditRequest {
	return &NullableApplicationEditRequest{value: val, isSet: true}
}

func (v NullableApplicationEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
