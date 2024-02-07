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

// checks if the AwsCredentialsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsCredentialsRequest{}

// AwsCredentialsRequest struct for AwsCredentialsRequest
type AwsCredentialsRequest struct {
	Name            string `json:"name"`
	AccessKeyId     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
}

type _AwsCredentialsRequest AwsCredentialsRequest

// NewAwsCredentialsRequest instantiates a new AwsCredentialsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsCredentialsRequest(name string, accessKeyId string, secretAccessKey string) *AwsCredentialsRequest {
	this := AwsCredentialsRequest{}
	this.Name = name
	this.AccessKeyId = accessKeyId
	this.SecretAccessKey = secretAccessKey
	return &this
}

// NewAwsCredentialsRequestWithDefaults instantiates a new AwsCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsCredentialsRequestWithDefaults() *AwsCredentialsRequest {
	this := AwsCredentialsRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AwsCredentialsRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AwsCredentialsRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AwsCredentialsRequest) SetName(v string) {
	o.Name = v
}

// GetAccessKeyId returns the AccessKeyId field value
func (o *AwsCredentialsRequest) GetAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
func (o *AwsCredentialsRequest) GetAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKeyId, true
}

// SetAccessKeyId sets field value
func (o *AwsCredentialsRequest) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

// GetSecretAccessKey returns the SecretAccessKey field value
func (o *AwsCredentialsRequest) GetSecretAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value
// and a boolean to check if the value has been set.
func (o *AwsCredentialsRequest) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretAccessKey, true
}

// SetSecretAccessKey sets field value
func (o *AwsCredentialsRequest) SetSecretAccessKey(v string) {
	o.SecretAccessKey = v
}

func (o AwsCredentialsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsCredentialsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["access_key_id"] = o.AccessKeyId
	toSerialize["secret_access_key"] = o.SecretAccessKey
	return toSerialize, nil
}

func (o *AwsCredentialsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"access_key_id",
		"secret_access_key",
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

	varAwsCredentialsRequest := _AwsCredentialsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAwsCredentialsRequest)

	if err != nil {
		return err
	}

	*o = AwsCredentialsRequest(varAwsCredentialsRequest)

	return err
}

type NullableAwsCredentialsRequest struct {
	value *AwsCredentialsRequest
	isSet bool
}

func (v NullableAwsCredentialsRequest) Get() *AwsCredentialsRequest {
	return v.value
}

func (v *NullableAwsCredentialsRequest) Set(val *AwsCredentialsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsCredentialsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsCredentialsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsCredentialsRequest(val *AwsCredentialsRequest) *NullableAwsCredentialsRequest {
	return &NullableAwsCredentialsRequest{value: val, isSet: true}
}

func (v NullableAwsCredentialsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsCredentialsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
