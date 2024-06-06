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

// checks if the DockerfileCheckResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DockerfileCheckResponse{}

// DockerfileCheckResponse struct for DockerfileCheckResponse
type DockerfileCheckResponse struct {
	DockerfilePath string `json:"dockerfile_path"`
	// All ARG variable declared in the Dockerfile
	Arg []string `json:"arg,omitempty"`
	// All image repositories we found declared in the Dockerfile
	Repositories         []string `json:"repositories,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DockerfileCheckResponse DockerfileCheckResponse

// NewDockerfileCheckResponse instantiates a new DockerfileCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerfileCheckResponse(dockerfilePath string) *DockerfileCheckResponse {
	this := DockerfileCheckResponse{}
	this.DockerfilePath = dockerfilePath
	return &this
}

// NewDockerfileCheckResponseWithDefaults instantiates a new DockerfileCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerfileCheckResponseWithDefaults() *DockerfileCheckResponse {
	this := DockerfileCheckResponse{}
	return &this
}

// GetDockerfilePath returns the DockerfilePath field value
func (o *DockerfileCheckResponse) GetDockerfilePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerfilePath
}

// GetDockerfilePathOk returns a tuple with the DockerfilePath field value
// and a boolean to check if the value has been set.
func (o *DockerfileCheckResponse) GetDockerfilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerfilePath, true
}

// SetDockerfilePath sets field value
func (o *DockerfileCheckResponse) SetDockerfilePath(v string) {
	o.DockerfilePath = v
}

// GetArg returns the Arg field value if set, zero value otherwise.
func (o *DockerfileCheckResponse) GetArg() []string {
	if o == nil || IsNil(o.Arg) {
		var ret []string
		return ret
	}
	return o.Arg
}

// GetArgOk returns a tuple with the Arg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerfileCheckResponse) GetArgOk() ([]string, bool) {
	if o == nil || IsNil(o.Arg) {
		return nil, false
	}
	return o.Arg, true
}

// HasArg returns a boolean if a field has been set.
func (o *DockerfileCheckResponse) HasArg() bool {
	if o != nil && !IsNil(o.Arg) {
		return true
	}

	return false
}

// SetArg gets a reference to the given []string and assigns it to the Arg field.
func (o *DockerfileCheckResponse) SetArg(v []string) {
	o.Arg = v
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *DockerfileCheckResponse) GetRepositories() []string {
	if o == nil || IsNil(o.Repositories) {
		var ret []string
		return ret
	}
	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerfileCheckResponse) GetRepositoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Repositories) {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *DockerfileCheckResponse) HasRepositories() bool {
	if o != nil && !IsNil(o.Repositories) {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given []string and assigns it to the Repositories field.
func (o *DockerfileCheckResponse) SetRepositories(v []string) {
	o.Repositories = v
}

func (o DockerfileCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DockerfileCheckResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dockerfile_path"] = o.DockerfilePath
	if !IsNil(o.Arg) {
		toSerialize["arg"] = o.Arg
	}
	if !IsNil(o.Repositories) {
		toSerialize["repositories"] = o.Repositories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DockerfileCheckResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dockerfile_path",
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

	varDockerfileCheckResponse := _DockerfileCheckResponse{}

	err = json.Unmarshal(data, &varDockerfileCheckResponse)

	if err != nil {
		return err
	}

	*o = DockerfileCheckResponse(varDockerfileCheckResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dockerfile_path")
		delete(additionalProperties, "arg")
		delete(additionalProperties, "repositories")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDockerfileCheckResponse struct {
	value *DockerfileCheckResponse
	isSet bool
}

func (v NullableDockerfileCheckResponse) Get() *DockerfileCheckResponse {
	return v.value
}

func (v *NullableDockerfileCheckResponse) Set(val *DockerfileCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerfileCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerfileCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerfileCheckResponse(val *DockerfileCheckResponse) *NullableDockerfileCheckResponse {
	return &NullableDockerfileCheckResponse{value: val, isSet: true}
}

func (v NullableDockerfileCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerfileCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
