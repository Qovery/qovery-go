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

// checks if the ClusterLogsError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterLogsError{}

// ClusterLogsError Present only for `error` log
type ClusterLogsError struct {
	// log error tag
	Tag *string `json:"tag,omitempty"`
	// log details about the error
	UserLogMessage *string `json:"user_log_message,omitempty"`
	// link to our documentation
	Link *string `json:"link,omitempty"`
	// hint the user can follow
	HintMessage     *string                          `json:"hint_message,omitempty"`
	EventDetails    *ClusterLogsErrorEventDetails    `json:"event_details,omitempty"`
	UnderlyingError *ClusterLogsErrorUnderlyingError `json:"underlying_error,omitempty"`
}

// NewClusterLogsError instantiates a new ClusterLogsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterLogsError() *ClusterLogsError {
	this := ClusterLogsError{}
	return &this
}

// NewClusterLogsErrorWithDefaults instantiates a new ClusterLogsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterLogsErrorWithDefaults() *ClusterLogsError {
	this := ClusterLogsError{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ClusterLogsError) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsError) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ClusterLogsError) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *ClusterLogsError) SetTag(v string) {
	o.Tag = &v
}

// GetUserLogMessage returns the UserLogMessage field value if set, zero value otherwise.
func (o *ClusterLogsError) GetUserLogMessage() string {
	if o == nil || IsNil(o.UserLogMessage) {
		var ret string
		return ret
	}
	return *o.UserLogMessage
}

// GetUserLogMessageOk returns a tuple with the UserLogMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsError) GetUserLogMessageOk() (*string, bool) {
	if o == nil || IsNil(o.UserLogMessage) {
		return nil, false
	}
	return o.UserLogMessage, true
}

// HasUserLogMessage returns a boolean if a field has been set.
func (o *ClusterLogsError) HasUserLogMessage() bool {
	if o != nil && !IsNil(o.UserLogMessage) {
		return true
	}

	return false
}

// SetUserLogMessage gets a reference to the given string and assigns it to the UserLogMessage field.
func (o *ClusterLogsError) SetUserLogMessage(v string) {
	o.UserLogMessage = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *ClusterLogsError) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsError) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *ClusterLogsError) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *ClusterLogsError) SetLink(v string) {
	o.Link = &v
}

// GetHintMessage returns the HintMessage field value if set, zero value otherwise.
func (o *ClusterLogsError) GetHintMessage() string {
	if o == nil || IsNil(o.HintMessage) {
		var ret string
		return ret
	}
	return *o.HintMessage
}

// GetHintMessageOk returns a tuple with the HintMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsError) GetHintMessageOk() (*string, bool) {
	if o == nil || IsNil(o.HintMessage) {
		return nil, false
	}
	return o.HintMessage, true
}

// HasHintMessage returns a boolean if a field has been set.
func (o *ClusterLogsError) HasHintMessage() bool {
	if o != nil && !IsNil(o.HintMessage) {
		return true
	}

	return false
}

// SetHintMessage gets a reference to the given string and assigns it to the HintMessage field.
func (o *ClusterLogsError) SetHintMessage(v string) {
	o.HintMessage = &v
}

// GetEventDetails returns the EventDetails field value if set, zero value otherwise.
func (o *ClusterLogsError) GetEventDetails() ClusterLogsErrorEventDetails {
	if o == nil || IsNil(o.EventDetails) {
		var ret ClusterLogsErrorEventDetails
		return ret
	}
	return *o.EventDetails
}

// GetEventDetailsOk returns a tuple with the EventDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsError) GetEventDetailsOk() (*ClusterLogsErrorEventDetails, bool) {
	if o == nil || IsNil(o.EventDetails) {
		return nil, false
	}
	return o.EventDetails, true
}

// HasEventDetails returns a boolean if a field has been set.
func (o *ClusterLogsError) HasEventDetails() bool {
	if o != nil && !IsNil(o.EventDetails) {
		return true
	}

	return false
}

// SetEventDetails gets a reference to the given ClusterLogsErrorEventDetails and assigns it to the EventDetails field.
func (o *ClusterLogsError) SetEventDetails(v ClusterLogsErrorEventDetails) {
	o.EventDetails = &v
}

// GetUnderlyingError returns the UnderlyingError field value if set, zero value otherwise.
func (o *ClusterLogsError) GetUnderlyingError() ClusterLogsErrorUnderlyingError {
	if o == nil || IsNil(o.UnderlyingError) {
		var ret ClusterLogsErrorUnderlyingError
		return ret
	}
	return *o.UnderlyingError
}

// GetUnderlyingErrorOk returns a tuple with the UnderlyingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsError) GetUnderlyingErrorOk() (*ClusterLogsErrorUnderlyingError, bool) {
	if o == nil || IsNil(o.UnderlyingError) {
		return nil, false
	}
	return o.UnderlyingError, true
}

// HasUnderlyingError returns a boolean if a field has been set.
func (o *ClusterLogsError) HasUnderlyingError() bool {
	if o != nil && !IsNil(o.UnderlyingError) {
		return true
	}

	return false
}

// SetUnderlyingError gets a reference to the given ClusterLogsErrorUnderlyingError and assigns it to the UnderlyingError field.
func (o *ClusterLogsError) SetUnderlyingError(v ClusterLogsErrorUnderlyingError) {
	o.UnderlyingError = &v
}

func (o ClusterLogsError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterLogsError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.UserLogMessage) {
		toSerialize["user_log_message"] = o.UserLogMessage
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.HintMessage) {
		toSerialize["hint_message"] = o.HintMessage
	}
	if !IsNil(o.EventDetails) {
		toSerialize["event_details"] = o.EventDetails
	}
	if !IsNil(o.UnderlyingError) {
		toSerialize["underlying_error"] = o.UnderlyingError
	}
	return toSerialize, nil
}

type NullableClusterLogsError struct {
	value *ClusterLogsError
	isSet bool
}

func (v NullableClusterLogsError) Get() *ClusterLogsError {
	return v.value
}

func (v *NullableClusterLogsError) Set(val *ClusterLogsError) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterLogsError) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterLogsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterLogsError(val *ClusterLogsError) *NullableClusterLogsError {
	return &NullableClusterLogsError{value: val, isSet: true}
}

func (v NullableClusterLogsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterLogsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
