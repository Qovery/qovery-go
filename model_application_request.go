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
	"fmt"
)

// checks if the ApplicationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationRequest{}

// ApplicationRequest struct for ApplicationRequest
type ApplicationRequest struct {
	Storage []ServiceStorageRequestStorageInner `json:"storage,omitempty"`
	Ports   []ServicePortRequestPortsInner      `json:"ports,omitempty"`
	// name is case insensitive
	Name string `json:"name"`
	// give a description to this application
	Description   NullableString                  `json:"description,omitempty"`
	GitRepository ApplicationGitRepositoryRequest `json:"git_repository"`
	BuildMode     *BuildModeEnum                  `json:"build_mode,omitempty"`
	// The path of the associated Dockerfile. Only if you are using build_mode = DOCKER
	DockerfilePath NullableString `json:"dockerfile_path,omitempty"`
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
	AutoPreview *bool    `json:"auto_preview,omitempty"`
	Arguments   []string `json:"arguments,omitempty"`
	// optional entrypoint when launching container
	Entrypoint *string `json:"entrypoint,omitempty"`
	// Specify if the application will be automatically updated after receiving a new commit.
	AutoDeploy        NullableBool               `json:"auto_deploy,omitempty"`
	AnnotationsGroups []ServiceAnnotationRequest `json:"annotations_groups,omitempty"`
	LabelsGroups      []ServiceLabelRequest      `json:"labels_groups,omitempty"`
	// Icon URI representing the application.
	IconUri              *string `json:"icon_uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationRequest ApplicationRequest

// NewApplicationRequest instantiates a new ApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationRequest(name string, gitRepository ApplicationGitRepositoryRequest, healthchecks Healthcheck) *ApplicationRequest {
	this := ApplicationRequest{}
	this.Name = name
	this.GitRepository = gitRepository
	var buildMode BuildModeEnum = BUILDMODEENUM_DOCKER
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

// NewApplicationRequestWithDefaults instantiates a new ApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationRequestWithDefaults() *ApplicationRequest {
	this := ApplicationRequest{}
	var buildMode BuildModeEnum = BUILDMODEENUM_DOCKER
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
func (o *ApplicationRequest) GetStorage() []ServiceStorageRequestStorageInner {
	if o == nil || IsNil(o.Storage) {
		var ret []ServiceStorageRequestStorageInner
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetStorageOk() ([]ServiceStorageRequestStorageInner, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ApplicationRequest) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []ServiceStorageRequestStorageInner and assigns it to the Storage field.
func (o *ApplicationRequest) SetStorage(v []ServiceStorageRequestStorageInner) {
	o.Storage = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ApplicationRequest) GetPorts() []ServicePortRequestPortsInner {
	if o == nil || IsNil(o.Ports) {
		var ret []ServicePortRequestPortsInner
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetPortsOk() ([]ServicePortRequestPortsInner, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ApplicationRequest) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ServicePortRequestPortsInner and assigns it to the Ports field.
func (o *ApplicationRequest) SetPorts(v []ServicePortRequestPortsInner) {
	o.Ports = v
}

// GetName returns the Name field value
func (o *ApplicationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ApplicationRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ApplicationRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ApplicationRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetGitRepository returns the GitRepository field value
func (o *ApplicationRequest) GetGitRepository() ApplicationGitRepositoryRequest {
	if o == nil {
		var ret ApplicationGitRepositoryRequest
		return ret
	}

	return o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetGitRepositoryOk() (*ApplicationGitRepositoryRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitRepository, true
}

// SetGitRepository sets field value
func (o *ApplicationRequest) SetGitRepository(v ApplicationGitRepositoryRequest) {
	o.GitRepository = v
}

// GetBuildMode returns the BuildMode field value if set, zero value otherwise.
func (o *ApplicationRequest) GetBuildMode() BuildModeEnum {
	if o == nil || IsNil(o.BuildMode) {
		var ret BuildModeEnum
		return ret
	}
	return *o.BuildMode
}

// GetBuildModeOk returns a tuple with the BuildMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetBuildModeOk() (*BuildModeEnum, bool) {
	if o == nil || IsNil(o.BuildMode) {
		return nil, false
	}
	return o.BuildMode, true
}

// HasBuildMode returns a boolean if a field has been set.
func (o *ApplicationRequest) HasBuildMode() bool {
	if o != nil && !IsNil(o.BuildMode) {
		return true
	}

	return false
}

// SetBuildMode gets a reference to the given BuildModeEnum and assigns it to the BuildMode field.
func (o *ApplicationRequest) SetBuildMode(v BuildModeEnum) {
	o.BuildMode = &v
}

// GetDockerfilePath returns the DockerfilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationRequest) GetDockerfilePath() string {
	if o == nil || IsNil(o.DockerfilePath.Get()) {
		var ret string
		return ret
	}
	return *o.DockerfilePath.Get()
}

// GetDockerfilePathOk returns a tuple with the DockerfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationRequest) GetDockerfilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DockerfilePath.Get(), o.DockerfilePath.IsSet()
}

// HasDockerfilePath returns a boolean if a field has been set.
func (o *ApplicationRequest) HasDockerfilePath() bool {
	if o != nil && o.DockerfilePath.IsSet() {
		return true
	}

	return false
}

// SetDockerfilePath gets a reference to the given NullableString and assigns it to the DockerfilePath field.
func (o *ApplicationRequest) SetDockerfilePath(v string) {
	o.DockerfilePath.Set(&v)
}

// SetDockerfilePathNil sets the value for DockerfilePath to be an explicit nil
func (o *ApplicationRequest) SetDockerfilePathNil() {
	o.DockerfilePath.Set(nil)
}

// UnsetDockerfilePath ensures that no value is present for DockerfilePath, not even an explicit nil
func (o *ApplicationRequest) UnsetDockerfilePath() {
	o.DockerfilePath.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ApplicationRequest) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ApplicationRequest) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *ApplicationRequest) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ApplicationRequest) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ApplicationRequest) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *ApplicationRequest) SetMemory(v int32) {
	o.Memory = &v
}

// GetMinRunningInstances returns the MinRunningInstances field value if set, zero value otherwise.
func (o *ApplicationRequest) GetMinRunningInstances() int32 {
	if o == nil || IsNil(o.MinRunningInstances) {
		var ret int32
		return ret
	}
	return *o.MinRunningInstances
}

// GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetMinRunningInstancesOk() (*int32, bool) {
	if o == nil || IsNil(o.MinRunningInstances) {
		return nil, false
	}
	return o.MinRunningInstances, true
}

// HasMinRunningInstances returns a boolean if a field has been set.
func (o *ApplicationRequest) HasMinRunningInstances() bool {
	if o != nil && !IsNil(o.MinRunningInstances) {
		return true
	}

	return false
}

// SetMinRunningInstances gets a reference to the given int32 and assigns it to the MinRunningInstances field.
func (o *ApplicationRequest) SetMinRunningInstances(v int32) {
	o.MinRunningInstances = &v
}

// GetMaxRunningInstances returns the MaxRunningInstances field value if set, zero value otherwise.
func (o *ApplicationRequest) GetMaxRunningInstances() int32 {
	if o == nil || IsNil(o.MaxRunningInstances) {
		var ret int32
		return ret
	}
	return *o.MaxRunningInstances
}

// GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetMaxRunningInstancesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRunningInstances) {
		return nil, false
	}
	return o.MaxRunningInstances, true
}

// HasMaxRunningInstances returns a boolean if a field has been set.
func (o *ApplicationRequest) HasMaxRunningInstances() bool {
	if o != nil && !IsNil(o.MaxRunningInstances) {
		return true
	}

	return false
}

// SetMaxRunningInstances gets a reference to the given int32 and assigns it to the MaxRunningInstances field.
func (o *ApplicationRequest) SetMaxRunningInstances(v int32) {
	o.MaxRunningInstances = &v
}

// GetHealthchecks returns the Healthchecks field value
func (o *ApplicationRequest) GetHealthchecks() Healthcheck {
	if o == nil {
		var ret Healthcheck
		return ret
	}

	return o.Healthchecks
}

// GetHealthchecksOk returns a tuple with the Healthchecks field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetHealthchecksOk() (*Healthcheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Healthchecks, true
}

// SetHealthchecks sets field value
func (o *ApplicationRequest) SetHealthchecks(v Healthcheck) {
	o.Healthchecks = v
}

// GetAutoPreview returns the AutoPreview field value if set, zero value otherwise.
func (o *ApplicationRequest) GetAutoPreview() bool {
	if o == nil || IsNil(o.AutoPreview) {
		var ret bool
		return ret
	}
	return *o.AutoPreview
}

// GetAutoPreviewOk returns a tuple with the AutoPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetAutoPreviewOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoPreview) {
		return nil, false
	}
	return o.AutoPreview, true
}

// HasAutoPreview returns a boolean if a field has been set.
func (o *ApplicationRequest) HasAutoPreview() bool {
	if o != nil && !IsNil(o.AutoPreview) {
		return true
	}

	return false
}

// SetAutoPreview gets a reference to the given bool and assigns it to the AutoPreview field.
func (o *ApplicationRequest) SetAutoPreview(v bool) {
	o.AutoPreview = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *ApplicationRequest) GetArguments() []string {
	if o == nil || IsNil(o.Arguments) {
		var ret []string
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetArgumentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Arguments) {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *ApplicationRequest) HasArguments() bool {
	if o != nil && !IsNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []string and assigns it to the Arguments field.
func (o *ApplicationRequest) SetArguments(v []string) {
	o.Arguments = v
}

// GetEntrypoint returns the Entrypoint field value if set, zero value otherwise.
func (o *ApplicationRequest) GetEntrypoint() string {
	if o == nil || IsNil(o.Entrypoint) {
		var ret string
		return ret
	}
	return *o.Entrypoint
}

// GetEntrypointOk returns a tuple with the Entrypoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetEntrypointOk() (*string, bool) {
	if o == nil || IsNil(o.Entrypoint) {
		return nil, false
	}
	return o.Entrypoint, true
}

// HasEntrypoint returns a boolean if a field has been set.
func (o *ApplicationRequest) HasEntrypoint() bool {
	if o != nil && !IsNil(o.Entrypoint) {
		return true
	}

	return false
}

// SetEntrypoint gets a reference to the given string and assigns it to the Entrypoint field.
func (o *ApplicationRequest) SetEntrypoint(v string) {
	o.Entrypoint = &v
}

// GetAutoDeploy returns the AutoDeploy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationRequest) GetAutoDeploy() bool {
	if o == nil || IsNil(o.AutoDeploy.Get()) {
		var ret bool
		return ret
	}
	return *o.AutoDeploy.Get()
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationRequest) GetAutoDeployOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoDeploy.Get(), o.AutoDeploy.IsSet()
}

// HasAutoDeploy returns a boolean if a field has been set.
func (o *ApplicationRequest) HasAutoDeploy() bool {
	if o != nil && o.AutoDeploy.IsSet() {
		return true
	}

	return false
}

// SetAutoDeploy gets a reference to the given NullableBool and assigns it to the AutoDeploy field.
func (o *ApplicationRequest) SetAutoDeploy(v bool) {
	o.AutoDeploy.Set(&v)
}

// SetAutoDeployNil sets the value for AutoDeploy to be an explicit nil
func (o *ApplicationRequest) SetAutoDeployNil() {
	o.AutoDeploy.Set(nil)
}

// UnsetAutoDeploy ensures that no value is present for AutoDeploy, not even an explicit nil
func (o *ApplicationRequest) UnsetAutoDeploy() {
	o.AutoDeploy.Unset()
}

// GetAnnotationsGroups returns the AnnotationsGroups field value if set, zero value otherwise.
func (o *ApplicationRequest) GetAnnotationsGroups() []ServiceAnnotationRequest {
	if o == nil || IsNil(o.AnnotationsGroups) {
		var ret []ServiceAnnotationRequest
		return ret
	}
	return o.AnnotationsGroups
}

// GetAnnotationsGroupsOk returns a tuple with the AnnotationsGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetAnnotationsGroupsOk() ([]ServiceAnnotationRequest, bool) {
	if o == nil || IsNil(o.AnnotationsGroups) {
		return nil, false
	}
	return o.AnnotationsGroups, true
}

// HasAnnotationsGroups returns a boolean if a field has been set.
func (o *ApplicationRequest) HasAnnotationsGroups() bool {
	if o != nil && !IsNil(o.AnnotationsGroups) {
		return true
	}

	return false
}

// SetAnnotationsGroups gets a reference to the given []ServiceAnnotationRequest and assigns it to the AnnotationsGroups field.
func (o *ApplicationRequest) SetAnnotationsGroups(v []ServiceAnnotationRequest) {
	o.AnnotationsGroups = v
}

// GetLabelsGroups returns the LabelsGroups field value if set, zero value otherwise.
func (o *ApplicationRequest) GetLabelsGroups() []ServiceLabelRequest {
	if o == nil || IsNil(o.LabelsGroups) {
		var ret []ServiceLabelRequest
		return ret
	}
	return o.LabelsGroups
}

// GetLabelsGroupsOk returns a tuple with the LabelsGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetLabelsGroupsOk() ([]ServiceLabelRequest, bool) {
	if o == nil || IsNil(o.LabelsGroups) {
		return nil, false
	}
	return o.LabelsGroups, true
}

// HasLabelsGroups returns a boolean if a field has been set.
func (o *ApplicationRequest) HasLabelsGroups() bool {
	if o != nil && !IsNil(o.LabelsGroups) {
		return true
	}

	return false
}

// SetLabelsGroups gets a reference to the given []ServiceLabelRequest and assigns it to the LabelsGroups field.
func (o *ApplicationRequest) SetLabelsGroups(v []ServiceLabelRequest) {
	o.LabelsGroups = v
}

// GetIconUri returns the IconUri field value if set, zero value otherwise.
func (o *ApplicationRequest) GetIconUri() string {
	if o == nil || IsNil(o.IconUri) {
		var ret string
		return ret
	}
	return *o.IconUri
}

// GetIconUriOk returns a tuple with the IconUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetIconUriOk() (*string, bool) {
	if o == nil || IsNil(o.IconUri) {
		return nil, false
	}
	return o.IconUri, true
}

// HasIconUri returns a boolean if a field has been set.
func (o *ApplicationRequest) HasIconUri() bool {
	if o != nil && !IsNil(o.IconUri) {
		return true
	}

	return false
}

// SetIconUri gets a reference to the given string and assigns it to the IconUri field.
func (o *ApplicationRequest) SetIconUri(v string) {
	o.IconUri = &v
}

func (o ApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["git_repository"] = o.GitRepository
	if !IsNil(o.BuildMode) {
		toSerialize["build_mode"] = o.BuildMode
	}
	if o.DockerfilePath.IsSet() {
		toSerialize["dockerfile_path"] = o.DockerfilePath.Get()
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
	if !IsNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}
	if !IsNil(o.Entrypoint) {
		toSerialize["entrypoint"] = o.Entrypoint
	}
	if o.AutoDeploy.IsSet() {
		toSerialize["auto_deploy"] = o.AutoDeploy.Get()
	}
	if !IsNil(o.AnnotationsGroups) {
		toSerialize["annotations_groups"] = o.AnnotationsGroups
	}
	if !IsNil(o.LabelsGroups) {
		toSerialize["labels_groups"] = o.LabelsGroups
	}
	if !IsNil(o.IconUri) {
		toSerialize["icon_uri"] = o.IconUri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"git_repository",
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

	varApplicationRequest := _ApplicationRequest{}

	err = json.Unmarshal(data, &varApplicationRequest)

	if err != nil {
		return err
	}

	*o = ApplicationRequest(varApplicationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "storage")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "git_repository")
		delete(additionalProperties, "build_mode")
		delete(additionalProperties, "dockerfile_path")
		delete(additionalProperties, "cpu")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "min_running_instances")
		delete(additionalProperties, "max_running_instances")
		delete(additionalProperties, "healthchecks")
		delete(additionalProperties, "auto_preview")
		delete(additionalProperties, "arguments")
		delete(additionalProperties, "entrypoint")
		delete(additionalProperties, "auto_deploy")
		delete(additionalProperties, "annotations_groups")
		delete(additionalProperties, "labels_groups")
		delete(additionalProperties, "icon_uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationRequest struct {
	value *ApplicationRequest
	isSet bool
}

func (v NullableApplicationRequest) Get() *ApplicationRequest {
	return v.value
}

func (v *NullableApplicationRequest) Set(val *ApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationRequest(val *ApplicationRequest) *NullableApplicationRequest {
	return &NullableApplicationRequest{value: val, isSet: true}
}

func (v NullableApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
