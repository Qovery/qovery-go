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

// checks if the HelmRequestAllOfValuesOverride type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmRequestAllOfValuesOverride{}

// HelmRequestAllOfValuesOverride Specify helm values you want to set or override
type HelmRequestAllOfValuesOverride struct {
	// The input is in json array format: [ [$KEY,$VALUE], [...] ]
	Set [][]string `json:"set,omitempty"`
	// The input is in json array format: [ [$KEY,$VALUE], [...] ]
	SetString [][]string `json:"set_string,omitempty"`
	// The input is in json array format: [ [$KEY,$VALUE], [...] ]
	SetJson              [][]string                                 `json:"set_json,omitempty"`
	File                 NullableHelmRequestAllOfValuesOverrideFile `json:"file,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HelmRequestAllOfValuesOverride HelmRequestAllOfValuesOverride

// NewHelmRequestAllOfValuesOverride instantiates a new HelmRequestAllOfValuesOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmRequestAllOfValuesOverride() *HelmRequestAllOfValuesOverride {
	this := HelmRequestAllOfValuesOverride{}
	return &this
}

// NewHelmRequestAllOfValuesOverrideWithDefaults instantiates a new HelmRequestAllOfValuesOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmRequestAllOfValuesOverrideWithDefaults() *HelmRequestAllOfValuesOverride {
	this := HelmRequestAllOfValuesOverride{}
	return &this
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *HelmRequestAllOfValuesOverride) GetSet() [][]string {
	if o == nil || IsNil(o.Set) {
		var ret [][]string
		return ret
	}
	return o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRequestAllOfValuesOverride) GetSetOk() ([][]string, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *HelmRequestAllOfValuesOverride) HasSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given [][]string and assigns it to the Set field.
func (o *HelmRequestAllOfValuesOverride) SetSet(v [][]string) {
	o.Set = v
}

// GetSetString returns the SetString field value if set, zero value otherwise.
func (o *HelmRequestAllOfValuesOverride) GetSetString() [][]string {
	if o == nil || IsNil(o.SetString) {
		var ret [][]string
		return ret
	}
	return o.SetString
}

// GetSetStringOk returns a tuple with the SetString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRequestAllOfValuesOverride) GetSetStringOk() ([][]string, bool) {
	if o == nil || IsNil(o.SetString) {
		return nil, false
	}
	return o.SetString, true
}

// HasSetString returns a boolean if a field has been set.
func (o *HelmRequestAllOfValuesOverride) HasSetString() bool {
	if o != nil && !IsNil(o.SetString) {
		return true
	}

	return false
}

// SetSetString gets a reference to the given [][]string and assigns it to the SetString field.
func (o *HelmRequestAllOfValuesOverride) SetSetString(v [][]string) {
	o.SetString = v
}

// GetSetJson returns the SetJson field value if set, zero value otherwise.
func (o *HelmRequestAllOfValuesOverride) GetSetJson() [][]string {
	if o == nil || IsNil(o.SetJson) {
		var ret [][]string
		return ret
	}
	return o.SetJson
}

// GetSetJsonOk returns a tuple with the SetJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRequestAllOfValuesOverride) GetSetJsonOk() ([][]string, bool) {
	if o == nil || IsNil(o.SetJson) {
		return nil, false
	}
	return o.SetJson, true
}

// HasSetJson returns a boolean if a field has been set.
func (o *HelmRequestAllOfValuesOverride) HasSetJson() bool {
	if o != nil && !IsNil(o.SetJson) {
		return true
	}

	return false
}

// SetSetJson gets a reference to the given [][]string and assigns it to the SetJson field.
func (o *HelmRequestAllOfValuesOverride) SetSetJson(v [][]string) {
	o.SetJson = v
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmRequestAllOfValuesOverride) GetFile() HelmRequestAllOfValuesOverrideFile {
	if o == nil || IsNil(o.File.Get()) {
		var ret HelmRequestAllOfValuesOverrideFile
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmRequestAllOfValuesOverride) GetFileOk() (*HelmRequestAllOfValuesOverrideFile, bool) {
	if o == nil {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *HelmRequestAllOfValuesOverride) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given NullableHelmRequestAllOfValuesOverrideFile and assigns it to the File field.
func (o *HelmRequestAllOfValuesOverride) SetFile(v HelmRequestAllOfValuesOverrideFile) {
	o.File.Set(&v)
}

// SetFileNil sets the value for File to be an explicit nil
func (o *HelmRequestAllOfValuesOverride) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *HelmRequestAllOfValuesOverride) UnsetFile() {
	o.File.Unset()
}

func (o HelmRequestAllOfValuesOverride) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmRequestAllOfValuesOverride) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Set) {
		toSerialize["set"] = o.Set
	}
	if !IsNil(o.SetString) {
		toSerialize["set_string"] = o.SetString
	}
	if !IsNil(o.SetJson) {
		toSerialize["set_json"] = o.SetJson
	}
	if o.File.IsSet() {
		toSerialize["file"] = o.File.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmRequestAllOfValuesOverride) UnmarshalJSON(data []byte) (err error) {
	varHelmRequestAllOfValuesOverride := _HelmRequestAllOfValuesOverride{}

	err = json.Unmarshal(data, &varHelmRequestAllOfValuesOverride)

	if err != nil {
		return err
	}

	*o = HelmRequestAllOfValuesOverride(varHelmRequestAllOfValuesOverride)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "set")
		delete(additionalProperties, "set_string")
		delete(additionalProperties, "set_json")
		delete(additionalProperties, "file")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmRequestAllOfValuesOverride struct {
	value *HelmRequestAllOfValuesOverride
	isSet bool
}

func (v NullableHelmRequestAllOfValuesOverride) Get() *HelmRequestAllOfValuesOverride {
	return v.value
}

func (v *NullableHelmRequestAllOfValuesOverride) Set(val *HelmRequestAllOfValuesOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmRequestAllOfValuesOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmRequestAllOfValuesOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmRequestAllOfValuesOverride(val *HelmRequestAllOfValuesOverride) *NullableHelmRequestAllOfValuesOverride {
	return &NullableHelmRequestAllOfValuesOverride{value: val, isSet: true}
}

func (v NullableHelmRequestAllOfValuesOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmRequestAllOfValuesOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
