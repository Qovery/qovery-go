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

// checks if the OrganizationGithubAppConnectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationGithubAppConnectRequest{}

// OrganizationGithubAppConnectRequest struct for OrganizationGithubAppConnectRequest
type OrganizationGithubAppConnectRequest struct {
	InstallationId string `json:"installation_id"`
	Code           string `json:"code"`
}

type _OrganizationGithubAppConnectRequest OrganizationGithubAppConnectRequest

// NewOrganizationGithubAppConnectRequest instantiates a new OrganizationGithubAppConnectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationGithubAppConnectRequest(installationId string, code string) *OrganizationGithubAppConnectRequest {
	this := OrganizationGithubAppConnectRequest{}
	this.InstallationId = installationId
	this.Code = code
	return &this
}

// NewOrganizationGithubAppConnectRequestWithDefaults instantiates a new OrganizationGithubAppConnectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationGithubAppConnectRequestWithDefaults() *OrganizationGithubAppConnectRequest {
	this := OrganizationGithubAppConnectRequest{}
	return &this
}

// GetInstallationId returns the InstallationId field value
func (o *OrganizationGithubAppConnectRequest) GetInstallationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value
// and a boolean to check if the value has been set.
func (o *OrganizationGithubAppConnectRequest) GetInstallationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallationId, true
}

// SetInstallationId sets field value
func (o *OrganizationGithubAppConnectRequest) SetInstallationId(v string) {
	o.InstallationId = v
}

// GetCode returns the Code field value
func (o *OrganizationGithubAppConnectRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OrganizationGithubAppConnectRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OrganizationGithubAppConnectRequest) SetCode(v string) {
	o.Code = v
}

func (o OrganizationGithubAppConnectRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationGithubAppConnectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["installation_id"] = o.InstallationId
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

func (o *OrganizationGithubAppConnectRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"installation_id",
		"code",
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

	varOrganizationGithubAppConnectRequest := _OrganizationGithubAppConnectRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationGithubAppConnectRequest)

	if err != nil {
		return err
	}

	*o = OrganizationGithubAppConnectRequest(varOrganizationGithubAppConnectRequest)

	return err
}

type NullableOrganizationGithubAppConnectRequest struct {
	value *OrganizationGithubAppConnectRequest
	isSet bool
}

func (v NullableOrganizationGithubAppConnectRequest) Get() *OrganizationGithubAppConnectRequest {
	return v.value
}

func (v *NullableOrganizationGithubAppConnectRequest) Set(val *OrganizationGithubAppConnectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationGithubAppConnectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationGithubAppConnectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationGithubAppConnectRequest(val *OrganizationGithubAppConnectRequest) *NullableOrganizationGithubAppConnectRequest {
	return &NullableOrganizationGithubAppConnectRequest{value: val, isSet: true}
}

func (v NullableOrganizationGithubAppConnectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationGithubAppConnectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
