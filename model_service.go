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
	"time"
)

// checks if the Service type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Service{}

// Service struct for Service
type Service struct {
	// uuid of the associated service (application, database, job, gateway...)
	Id        string           `json:"id"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt *time.Time       `json:"updated_at,omitempty"`
	Type      *ServiceTypeEnum `json:"type,omitempty"`
	// name of the service
	Name *string `json:"name,omitempty"`
	// Git commit ID corresponding to the deployed version of the application
	DeployedCommitId *string `json:"deployed_commit_id,omitempty"`
	// uuid of the user that made the last update
	LastUpdatedBy *string `json:"last_updated_by,omitempty"`
	// global overview of resources consumption of the service
	ConsumedResourcesInPercent *float32 `json:"consumed_resources_in_percent,omitempty"`
	// describes the typology of service (container, postgresl, redis...)
	ServiceTypology *string `json:"service_typology,omitempty"`
	// for databases this field exposes the database version
	ServiceVersion *string `json:"service_version,omitempty"`
	ToUpdate       *bool   `json:"to_update,omitempty"`
}

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService(id string, createdAt time.Time) *Service {
	this := Service{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetId returns the Id field value
func (o *Service) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Service) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Service) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Service) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Service) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Service) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Service) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Service) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Service) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Service) GetType() ServiceTypeEnum {
	if o == nil || IsNil(o.Type) {
		var ret ServiceTypeEnum
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTypeOk() (*ServiceTypeEnum, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Service) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ServiceTypeEnum and assigns it to the Type field.
func (o *Service) SetType(v ServiceTypeEnum) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Service) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Service) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Service) SetName(v string) {
	o.Name = &v
}

// GetDeployedCommitId returns the DeployedCommitId field value if set, zero value otherwise.
func (o *Service) GetDeployedCommitId() string {
	if o == nil || IsNil(o.DeployedCommitId) {
		var ret string
		return ret
	}
	return *o.DeployedCommitId
}

// GetDeployedCommitIdOk returns a tuple with the DeployedCommitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDeployedCommitIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeployedCommitId) {
		return nil, false
	}
	return o.DeployedCommitId, true
}

// HasDeployedCommitId returns a boolean if a field has been set.
func (o *Service) HasDeployedCommitId() bool {
	if o != nil && !IsNil(o.DeployedCommitId) {
		return true
	}

	return false
}

// SetDeployedCommitId gets a reference to the given string and assigns it to the DeployedCommitId field.
func (o *Service) SetDeployedCommitId(v string) {
	o.DeployedCommitId = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *Service) GetLastUpdatedBy() string {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *Service) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *Service) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetConsumedResourcesInPercent returns the ConsumedResourcesInPercent field value if set, zero value otherwise.
func (o *Service) GetConsumedResourcesInPercent() float32 {
	if o == nil || IsNil(o.ConsumedResourcesInPercent) {
		var ret float32
		return ret
	}
	return *o.ConsumedResourcesInPercent
}

// GetConsumedResourcesInPercentOk returns a tuple with the ConsumedResourcesInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetConsumedResourcesInPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.ConsumedResourcesInPercent) {
		return nil, false
	}
	return o.ConsumedResourcesInPercent, true
}

// HasConsumedResourcesInPercent returns a boolean if a field has been set.
func (o *Service) HasConsumedResourcesInPercent() bool {
	if o != nil && !IsNil(o.ConsumedResourcesInPercent) {
		return true
	}

	return false
}

// SetConsumedResourcesInPercent gets a reference to the given float32 and assigns it to the ConsumedResourcesInPercent field.
func (o *Service) SetConsumedResourcesInPercent(v float32) {
	o.ConsumedResourcesInPercent = &v
}

// GetServiceTypology returns the ServiceTypology field value if set, zero value otherwise.
func (o *Service) GetServiceTypology() string {
	if o == nil || IsNil(o.ServiceTypology) {
		var ret string
		return ret
	}
	return *o.ServiceTypology
}

// GetServiceTypologyOk returns a tuple with the ServiceTypology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetServiceTypologyOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceTypology) {
		return nil, false
	}
	return o.ServiceTypology, true
}

// HasServiceTypology returns a boolean if a field has been set.
func (o *Service) HasServiceTypology() bool {
	if o != nil && !IsNil(o.ServiceTypology) {
		return true
	}

	return false
}

// SetServiceTypology gets a reference to the given string and assigns it to the ServiceTypology field.
func (o *Service) SetServiceTypology(v string) {
	o.ServiceTypology = &v
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise.
func (o *Service) GetServiceVersion() string {
	if o == nil || IsNil(o.ServiceVersion) {
		var ret string
		return ret
	}
	return *o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetServiceVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceVersion) {
		return nil, false
	}
	return o.ServiceVersion, true
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *Service) HasServiceVersion() bool {
	if o != nil && !IsNil(o.ServiceVersion) {
		return true
	}

	return false
}

// SetServiceVersion gets a reference to the given string and assigns it to the ServiceVersion field.
func (o *Service) SetServiceVersion(v string) {
	o.ServiceVersion = &v
}

// GetToUpdate returns the ToUpdate field value if set, zero value otherwise.
func (o *Service) GetToUpdate() bool {
	if o == nil || IsNil(o.ToUpdate) {
		var ret bool
		return ret
	}
	return *o.ToUpdate
}

// GetToUpdateOk returns a tuple with the ToUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetToUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.ToUpdate) {
		return nil, false
	}
	return o.ToUpdate, true
}

// HasToUpdate returns a boolean if a field has been set.
func (o *Service) HasToUpdate() bool {
	if o != nil && !IsNil(o.ToUpdate) {
		return true
	}

	return false
}

// SetToUpdate gets a reference to the given bool and assigns it to the ToUpdate field.
func (o *Service) SetToUpdate(v bool) {
	o.ToUpdate = &v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Service) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DeployedCommitId) {
		toSerialize["deployed_commit_id"] = o.DeployedCommitId
	}
	if !IsNil(o.LastUpdatedBy) {
		toSerialize["last_updated_by"] = o.LastUpdatedBy
	}
	if !IsNil(o.ConsumedResourcesInPercent) {
		toSerialize["consumed_resources_in_percent"] = o.ConsumedResourcesInPercent
	}
	if !IsNil(o.ServiceTypology) {
		toSerialize["service_typology"] = o.ServiceTypology
	}
	if !IsNil(o.ServiceVersion) {
		toSerialize["service_version"] = o.ServiceVersion
	}
	if !IsNil(o.ToUpdate) {
		toSerialize["to_update"] = o.ToUpdate
	}
	return toSerialize, nil
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
