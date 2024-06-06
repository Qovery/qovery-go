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

// checks if the ApplicationResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationResponseList{}

// ApplicationResponseList struct for ApplicationResponseList
type ApplicationResponseList struct {
	Results              []Application `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationResponseList ApplicationResponseList

// NewApplicationResponseList instantiates a new ApplicationResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResponseList() *ApplicationResponseList {
	this := ApplicationResponseList{}
	return &this
}

// NewApplicationResponseListWithDefaults instantiates a new ApplicationResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResponseListWithDefaults() *ApplicationResponseList {
	this := ApplicationResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ApplicationResponseList) GetResults() []Application {
	if o == nil || IsNil(o.Results) {
		var ret []Application
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseList) GetResultsOk() ([]Application, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ApplicationResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Application and assigns it to the Results field.
func (o *ApplicationResponseList) SetResults(v []Application) {
	o.Results = v
}

func (o ApplicationResponseList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationResponseList) UnmarshalJSON(data []byte) (err error) {
	varApplicationResponseList := _ApplicationResponseList{}

	err = json.Unmarshal(data, &varApplicationResponseList)

	if err != nil {
		return err
	}

	*o = ApplicationResponseList(varApplicationResponseList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationResponseList struct {
	value *ApplicationResponseList
	isSet bool
}

func (v NullableApplicationResponseList) Get() *ApplicationResponseList {
	return v.value
}

func (v *NullableApplicationResponseList) Set(val *ApplicationResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResponseList(val *ApplicationResponseList) *NullableApplicationResponseList {
	return &NullableApplicationResponseList{value: val, isSet: true}
}

func (v NullableApplicationResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
