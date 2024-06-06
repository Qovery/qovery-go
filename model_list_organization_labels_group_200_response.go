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

// checks if the ListOrganizationLabelsGroup200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationLabelsGroup200Response{}

// ListOrganizationLabelsGroup200Response struct for ListOrganizationLabelsGroup200Response
type ListOrganizationLabelsGroup200Response struct {
	Results              []OrganizationLabelsGroupEnrichedResponse `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListOrganizationLabelsGroup200Response ListOrganizationLabelsGroup200Response

// NewListOrganizationLabelsGroup200Response instantiates a new ListOrganizationLabelsGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationLabelsGroup200Response() *ListOrganizationLabelsGroup200Response {
	this := ListOrganizationLabelsGroup200Response{}
	return &this
}

// NewListOrganizationLabelsGroup200ResponseWithDefaults instantiates a new ListOrganizationLabelsGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationLabelsGroup200ResponseWithDefaults() *ListOrganizationLabelsGroup200Response {
	this := ListOrganizationLabelsGroup200Response{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListOrganizationLabelsGroup200Response) GetResults() []OrganizationLabelsGroupEnrichedResponse {
	if o == nil || IsNil(o.Results) {
		var ret []OrganizationLabelsGroupEnrichedResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationLabelsGroup200Response) GetResultsOk() ([]OrganizationLabelsGroupEnrichedResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListOrganizationLabelsGroup200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OrganizationLabelsGroupEnrichedResponse and assigns it to the Results field.
func (o *ListOrganizationLabelsGroup200Response) SetResults(v []OrganizationLabelsGroupEnrichedResponse) {
	o.Results = v
}

func (o ListOrganizationLabelsGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrganizationLabelsGroup200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListOrganizationLabelsGroup200Response) UnmarshalJSON(data []byte) (err error) {
	varListOrganizationLabelsGroup200Response := _ListOrganizationLabelsGroup200Response{}

	err = json.Unmarshal(data, &varListOrganizationLabelsGroup200Response)

	if err != nil {
		return err
	}

	*o = ListOrganizationLabelsGroup200Response(varListOrganizationLabelsGroup200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListOrganizationLabelsGroup200Response struct {
	value *ListOrganizationLabelsGroup200Response
	isSet bool
}

func (v NullableListOrganizationLabelsGroup200Response) Get() *ListOrganizationLabelsGroup200Response {
	return v.value
}

func (v *NullableListOrganizationLabelsGroup200Response) Set(val *ListOrganizationLabelsGroup200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationLabelsGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationLabelsGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationLabelsGroup200Response(val *ListOrganizationLabelsGroup200Response) *NullableListOrganizationLabelsGroup200Response {
	return &NullableListOrganizationLabelsGroup200Response{value: val, isSet: true}
}

func (v NullableListOrganizationLabelsGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationLabelsGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
