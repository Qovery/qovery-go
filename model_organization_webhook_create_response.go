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
	"time"
)

// checks if the OrganizationWebhookCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationWebhookCreateResponse{}

// OrganizationWebhookCreateResponse struct for OrganizationWebhookCreateResponse
type OrganizationWebhookCreateResponse struct {
	Id        string                       `json:"id"`
	CreatedAt time.Time                    `json:"created_at"`
	UpdatedAt *time.Time                   `json:"updated_at,omitempty"`
	Kind      *OrganizationWebhookKindEnum `json:"kind,omitempty"`
	// Set the public HTTP or HTTPS endpoint that will receive the specified events. The target URL must starts with `http://` or `https://`
	TargetUrl       *string `json:"target_url,omitempty"`
	TargetSecretSet *bool   `json:"target_secret_set,omitempty"`
	Description     *string `json:"description,omitempty"`
	// Turn on or off your endpoint.
	Enabled *bool                          `json:"enabled,omitempty"`
	Events  []OrganizationWebhookEventEnum `json:"events,omitempty"`
	// Specify the project names you want to filter to.  This webhook will be triggered only if the event is coming from the specified Project IDs. Notes: 1. Wildcard is accepted E.g. `product*`. 2. Name is case insensitive.
	ProjectNamesFilter []string `json:"project_names_filter,omitempty"`
	// Specify the environment modes you want to filter to. This webhook will be triggered only if the event is coming from an environment with the specified mode.
	EnvironmentTypesFilter []EnvironmentModeEnum `json:"environment_types_filter,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _OrganizationWebhookCreateResponse OrganizationWebhookCreateResponse

// NewOrganizationWebhookCreateResponse instantiates a new OrganizationWebhookCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationWebhookCreateResponse(id string, createdAt time.Time) *OrganizationWebhookCreateResponse {
	this := OrganizationWebhookCreateResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewOrganizationWebhookCreateResponseWithDefaults instantiates a new OrganizationWebhookCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWebhookCreateResponseWithDefaults() *OrganizationWebhookCreateResponse {
	this := OrganizationWebhookCreateResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationWebhookCreateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationWebhookCreateResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationWebhookCreateResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationWebhookCreateResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OrganizationWebhookCreateResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetKind() OrganizationWebhookKindEnum {
	if o == nil || IsNil(o.Kind) {
		var ret OrganizationWebhookKindEnum
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetKindOk() (*OrganizationWebhookKindEnum, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given OrganizationWebhookKindEnum and assigns it to the Kind field.
func (o *OrganizationWebhookCreateResponse) SetKind(v OrganizationWebhookKindEnum) {
	o.Kind = &v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetTargetUrl() string {
	if o == nil || IsNil(o.TargetUrl) {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetTargetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUrl) {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasTargetUrl() bool {
	if o != nil && !IsNil(o.TargetUrl) {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *OrganizationWebhookCreateResponse) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

// GetTargetSecretSet returns the TargetSecretSet field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetTargetSecretSet() bool {
	if o == nil || IsNil(o.TargetSecretSet) {
		var ret bool
		return ret
	}
	return *o.TargetSecretSet
}

// GetTargetSecretSetOk returns a tuple with the TargetSecretSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetTargetSecretSetOk() (*bool, bool) {
	if o == nil || IsNil(o.TargetSecretSet) {
		return nil, false
	}
	return o.TargetSecretSet, true
}

// HasTargetSecretSet returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasTargetSecretSet() bool {
	if o != nil && !IsNil(o.TargetSecretSet) {
		return true
	}

	return false
}

// SetTargetSecretSet gets a reference to the given bool and assigns it to the TargetSecretSet field.
func (o *OrganizationWebhookCreateResponse) SetTargetSecretSet(v bool) {
	o.TargetSecretSet = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationWebhookCreateResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationWebhookCreateResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetEvents() []OrganizationWebhookEventEnum {
	if o == nil || IsNil(o.Events) {
		var ret []OrganizationWebhookEventEnum
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetEventsOk() ([]OrganizationWebhookEventEnum, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []OrganizationWebhookEventEnum and assigns it to the Events field.
func (o *OrganizationWebhookCreateResponse) SetEvents(v []OrganizationWebhookEventEnum) {
	o.Events = v
}

// GetProjectNamesFilter returns the ProjectNamesFilter field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetProjectNamesFilter() []string {
	if o == nil || IsNil(o.ProjectNamesFilter) {
		var ret []string
		return ret
	}
	return o.ProjectNamesFilter
}

// GetProjectNamesFilterOk returns a tuple with the ProjectNamesFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetProjectNamesFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectNamesFilter) {
		return nil, false
	}
	return o.ProjectNamesFilter, true
}

// HasProjectNamesFilter returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasProjectNamesFilter() bool {
	if o != nil && !IsNil(o.ProjectNamesFilter) {
		return true
	}

	return false
}

// SetProjectNamesFilter gets a reference to the given []string and assigns it to the ProjectNamesFilter field.
func (o *OrganizationWebhookCreateResponse) SetProjectNamesFilter(v []string) {
	o.ProjectNamesFilter = v
}

// GetEnvironmentTypesFilter returns the EnvironmentTypesFilter field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetEnvironmentTypesFilter() []EnvironmentModeEnum {
	if o == nil || IsNil(o.EnvironmentTypesFilter) {
		var ret []EnvironmentModeEnum
		return ret
	}
	return o.EnvironmentTypesFilter
}

// GetEnvironmentTypesFilterOk returns a tuple with the EnvironmentTypesFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetEnvironmentTypesFilterOk() ([]EnvironmentModeEnum, bool) {
	if o == nil || IsNil(o.EnvironmentTypesFilter) {
		return nil, false
	}
	return o.EnvironmentTypesFilter, true
}

// HasEnvironmentTypesFilter returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasEnvironmentTypesFilter() bool {
	if o != nil && !IsNil(o.EnvironmentTypesFilter) {
		return true
	}

	return false
}

// SetEnvironmentTypesFilter gets a reference to the given []EnvironmentModeEnum and assigns it to the EnvironmentTypesFilter field.
func (o *OrganizationWebhookCreateResponse) SetEnvironmentTypesFilter(v []EnvironmentModeEnum) {
	o.EnvironmentTypesFilter = v
}

func (o OrganizationWebhookCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationWebhookCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.TargetUrl) {
		toSerialize["target_url"] = o.TargetUrl
	}
	if !IsNil(o.TargetSecretSet) {
		toSerialize["target_secret_set"] = o.TargetSecretSet
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.ProjectNamesFilter) {
		toSerialize["project_names_filter"] = o.ProjectNamesFilter
	}
	if !IsNil(o.EnvironmentTypesFilter) {
		toSerialize["environment_types_filter"] = o.EnvironmentTypesFilter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationWebhookCreateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
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

	varOrganizationWebhookCreateResponse := _OrganizationWebhookCreateResponse{}

	err = json.Unmarshal(data, &varOrganizationWebhookCreateResponse)

	if err != nil {
		return err
	}

	*o = OrganizationWebhookCreateResponse(varOrganizationWebhookCreateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "kind")
		delete(additionalProperties, "target_url")
		delete(additionalProperties, "target_secret_set")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "events")
		delete(additionalProperties, "project_names_filter")
		delete(additionalProperties, "environment_types_filter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationWebhookCreateResponse struct {
	value *OrganizationWebhookCreateResponse
	isSet bool
}

func (v NullableOrganizationWebhookCreateResponse) Get() *OrganizationWebhookCreateResponse {
	return v.value
}

func (v *NullableOrganizationWebhookCreateResponse) Set(val *OrganizationWebhookCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationWebhookCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationWebhookCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationWebhookCreateResponse(val *OrganizationWebhookCreateResponse) *NullableOrganizationWebhookCreateResponse {
	return &NullableOrganizationWebhookCreateResponse{value: val, isSet: true}
}

func (v NullableOrganizationWebhookCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationWebhookCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
