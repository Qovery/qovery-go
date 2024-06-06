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

// checks if the Database type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Database{}

// Database struct for Database
type Database struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// name is case insensitive
	Name string `json:"name"`
	// give a description to this database
	Description   *string                    `json:"description,omitempty"`
	Type          DatabaseTypeEnum           `json:"type"`
	Version       string                     `json:"version"`
	Mode          DatabaseModeEnum           `json:"mode"`
	Accessibility *DatabaseAccessibilityEnum `json:"accessibility,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu This field will be ignored for managed DB (instance type will be used instead).
	Cpu *int32 `json:"cpu,omitempty"`
	// Database instance type to be used for this database. The list of values can be retrieved via the endpoint /{CloudProvider}/managedDatabase/instanceType/{region}/{dbType}. This field is null for container DB.
	InstanceType *string `json:"instance_type,omitempty"`
	// unit is MB. 1024 MB = 1GB This field will be ignored for managed DB (instance type will be used instead). Default value is linked to the database type: - MANAGED: `100` - CONTAINER   - POSTGRES: `100`   - REDIS: `100`   - MYSQL: `512`   - MONGODB: `256`
	Memory *int32 `json:"memory,omitempty"`
	// unit is GB
	Storage           *int32                                 `json:"storage,omitempty"`
	AnnotationsGroups []OrganizationAnnotationsGroupResponse `json:"annotations_groups,omitempty"`
	LabelsGroups      []OrganizationLabelsGroupResponse      `json:"labels_groups,omitempty"`
	Environment       ReferenceObject                        `json:"environment"`
	Host              *string                                `json:"host,omitempty"`
	Port              *int32                                 `json:"port,omitempty"`
	// Maximum cpu that can be allocated to the database based on organization cluster configuration. unit is millicores (m). 1000m = 1 cpu
	MaximumCpu *int32 `json:"maximum_cpu,omitempty"`
	// Maximum memory that can be allocated to the database based on organization cluster configuration. unit is MB. 1024 MB = 1GB
	MaximumMemory *int32 `json:"maximum_memory,omitempty"`
	// indicates if the database disk is encrypted or not
	DiskEncrypted        *bool `json:"disk_encrypted,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Database Database

// NewDatabase instantiates a new Database object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabase(id string, createdAt time.Time, name string, type_ DatabaseTypeEnum, version string, mode DatabaseModeEnum, environment ReferenceObject) *Database {
	this := Database{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Name = name
	this.Type = type_
	this.Version = version
	this.Mode = mode
	var accessibility DatabaseAccessibilityEnum = DATABASEACCESSIBILITYENUM_PRIVATE
	this.Accessibility = &accessibility
	var cpu int32 = 250
	this.Cpu = &cpu
	var storage int32 = 10
	this.Storage = &storage
	this.Environment = environment
	return &this
}

// NewDatabaseWithDefaults instantiates a new Database object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseWithDefaults() *Database {
	this := Database{}
	var accessibility DatabaseAccessibilityEnum = DATABASEACCESSIBILITYENUM_PRIVATE
	this.Accessibility = &accessibility
	var cpu int32 = 250
	this.Cpu = &cpu
	var storage int32 = 10
	this.Storage = &storage
	return &this
}

// GetId returns the Id field value
func (o *Database) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Database) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Database) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Database) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Database) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Database) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Database) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Database) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Database) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value
func (o *Database) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Database) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Database) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Database) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Database) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Database) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *Database) GetType() DatabaseTypeEnum {
	if o == nil {
		var ret DatabaseTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Database) GetTypeOk() (*DatabaseTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Database) SetType(v DatabaseTypeEnum) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *Database) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Database) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Database) SetVersion(v string) {
	o.Version = v
}

// GetMode returns the Mode field value
func (o *Database) GetMode() DatabaseModeEnum {
	if o == nil {
		var ret DatabaseModeEnum
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *Database) GetModeOk() (*DatabaseModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *Database) SetMode(v DatabaseModeEnum) {
	o.Mode = v
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *Database) GetAccessibility() DatabaseAccessibilityEnum {
	if o == nil || IsNil(o.Accessibility) {
		var ret DatabaseAccessibilityEnum
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetAccessibilityOk() (*DatabaseAccessibilityEnum, bool) {
	if o == nil || IsNil(o.Accessibility) {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *Database) HasAccessibility() bool {
	if o != nil && !IsNil(o.Accessibility) {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given DatabaseAccessibilityEnum and assigns it to the Accessibility field.
func (o *Database) SetAccessibility(v DatabaseAccessibilityEnum) {
	o.Accessibility = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Database) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Database) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *Database) SetCpu(v int32) {
	o.Cpu = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *Database) GetInstanceType() string {
	if o == nil || IsNil(o.InstanceType) {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceType) {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *Database) HasInstanceType() bool {
	if o != nil && !IsNil(o.InstanceType) {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *Database) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Database) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Database) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *Database) SetMemory(v int32) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *Database) GetStorage() int32 {
	if o == nil || IsNil(o.Storage) {
		var ret int32
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetStorageOk() (*int32, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *Database) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given int32 and assigns it to the Storage field.
func (o *Database) SetStorage(v int32) {
	o.Storage = &v
}

// GetAnnotationsGroups returns the AnnotationsGroups field value if set, zero value otherwise.
func (o *Database) GetAnnotationsGroups() []OrganizationAnnotationsGroupResponse {
	if o == nil || IsNil(o.AnnotationsGroups) {
		var ret []OrganizationAnnotationsGroupResponse
		return ret
	}
	return o.AnnotationsGroups
}

// GetAnnotationsGroupsOk returns a tuple with the AnnotationsGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetAnnotationsGroupsOk() ([]OrganizationAnnotationsGroupResponse, bool) {
	if o == nil || IsNil(o.AnnotationsGroups) {
		return nil, false
	}
	return o.AnnotationsGroups, true
}

// HasAnnotationsGroups returns a boolean if a field has been set.
func (o *Database) HasAnnotationsGroups() bool {
	if o != nil && !IsNil(o.AnnotationsGroups) {
		return true
	}

	return false
}

// SetAnnotationsGroups gets a reference to the given []OrganizationAnnotationsGroupResponse and assigns it to the AnnotationsGroups field.
func (o *Database) SetAnnotationsGroups(v []OrganizationAnnotationsGroupResponse) {
	o.AnnotationsGroups = v
}

// GetLabelsGroups returns the LabelsGroups field value if set, zero value otherwise.
func (o *Database) GetLabelsGroups() []OrganizationLabelsGroupResponse {
	if o == nil || IsNil(o.LabelsGroups) {
		var ret []OrganizationLabelsGroupResponse
		return ret
	}
	return o.LabelsGroups
}

// GetLabelsGroupsOk returns a tuple with the LabelsGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetLabelsGroupsOk() ([]OrganizationLabelsGroupResponse, bool) {
	if o == nil || IsNil(o.LabelsGroups) {
		return nil, false
	}
	return o.LabelsGroups, true
}

// HasLabelsGroups returns a boolean if a field has been set.
func (o *Database) HasLabelsGroups() bool {
	if o != nil && !IsNil(o.LabelsGroups) {
		return true
	}

	return false
}

// SetLabelsGroups gets a reference to the given []OrganizationLabelsGroupResponse and assigns it to the LabelsGroups field.
func (o *Database) SetLabelsGroups(v []OrganizationLabelsGroupResponse) {
	o.LabelsGroups = v
}

// GetEnvironment returns the Environment field value
func (o *Database) GetEnvironment() ReferenceObject {
	if o == nil {
		var ret ReferenceObject
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *Database) GetEnvironmentOk() (*ReferenceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *Database) SetEnvironment(v ReferenceObject) {
	o.Environment = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *Database) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Database) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *Database) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Database) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Database) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *Database) SetPort(v int32) {
	o.Port = &v
}

// GetMaximumCpu returns the MaximumCpu field value if set, zero value otherwise.
func (o *Database) GetMaximumCpu() int32 {
	if o == nil || IsNil(o.MaximumCpu) {
		var ret int32
		return ret
	}
	return *o.MaximumCpu
}

// GetMaximumCpuOk returns a tuple with the MaximumCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetMaximumCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumCpu) {
		return nil, false
	}
	return o.MaximumCpu, true
}

// HasMaximumCpu returns a boolean if a field has been set.
func (o *Database) HasMaximumCpu() bool {
	if o != nil && !IsNil(o.MaximumCpu) {
		return true
	}

	return false
}

// SetMaximumCpu gets a reference to the given int32 and assigns it to the MaximumCpu field.
func (o *Database) SetMaximumCpu(v int32) {
	o.MaximumCpu = &v
}

// GetMaximumMemory returns the MaximumMemory field value if set, zero value otherwise.
func (o *Database) GetMaximumMemory() int32 {
	if o == nil || IsNil(o.MaximumMemory) {
		var ret int32
		return ret
	}
	return *o.MaximumMemory
}

// GetMaximumMemoryOk returns a tuple with the MaximumMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetMaximumMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumMemory) {
		return nil, false
	}
	return o.MaximumMemory, true
}

// HasMaximumMemory returns a boolean if a field has been set.
func (o *Database) HasMaximumMemory() bool {
	if o != nil && !IsNil(o.MaximumMemory) {
		return true
	}

	return false
}

// SetMaximumMemory gets a reference to the given int32 and assigns it to the MaximumMemory field.
func (o *Database) SetMaximumMemory(v int32) {
	o.MaximumMemory = &v
}

// GetDiskEncrypted returns the DiskEncrypted field value if set, zero value otherwise.
func (o *Database) GetDiskEncrypted() bool {
	if o == nil || IsNil(o.DiskEncrypted) {
		var ret bool
		return ret
	}
	return *o.DiskEncrypted
}

// GetDiskEncryptedOk returns a tuple with the DiskEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetDiskEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.DiskEncrypted) {
		return nil, false
	}
	return o.DiskEncrypted, true
}

// HasDiskEncrypted returns a boolean if a field has been set.
func (o *Database) HasDiskEncrypted() bool {
	if o != nil && !IsNil(o.DiskEncrypted) {
		return true
	}

	return false
}

// SetDiskEncrypted gets a reference to the given bool and assigns it to the DiskEncrypted field.
func (o *Database) SetDiskEncrypted(v bool) {
	o.DiskEncrypted = &v
}

func (o Database) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Database) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version
	toSerialize["mode"] = o.Mode
	if !IsNil(o.Accessibility) {
		toSerialize["accessibility"] = o.Accessibility
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.InstanceType) {
		toSerialize["instance_type"] = o.InstanceType
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	if !IsNil(o.AnnotationsGroups) {
		toSerialize["annotations_groups"] = o.AnnotationsGroups
	}
	if !IsNil(o.LabelsGroups) {
		toSerialize["labels_groups"] = o.LabelsGroups
	}
	toSerialize["environment"] = o.Environment
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.MaximumCpu) {
		toSerialize["maximum_cpu"] = o.MaximumCpu
	}
	if !IsNil(o.MaximumMemory) {
		toSerialize["maximum_memory"] = o.MaximumMemory
	}
	if !IsNil(o.DiskEncrypted) {
		toSerialize["disk_encrypted"] = o.DiskEncrypted
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Database) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"name",
		"type",
		"version",
		"mode",
		"environment",
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

	varDatabase := _Database{}

	err = json.Unmarshal(data, &varDatabase)

	if err != nil {
		return err
	}

	*o = Database(varDatabase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "version")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "accessibility")
		delete(additionalProperties, "cpu")
		delete(additionalProperties, "instance_type")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "storage")
		delete(additionalProperties, "annotations_groups")
		delete(additionalProperties, "labels_groups")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "host")
		delete(additionalProperties, "port")
		delete(additionalProperties, "maximum_cpu")
		delete(additionalProperties, "maximum_memory")
		delete(additionalProperties, "disk_encrypted")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDatabase struct {
	value *Database
	isSet bool
}

func (v NullableDatabase) Get() *Database {
	return v.value
}

func (v *NullableDatabase) Set(val *Database) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabase(val *Database) *NullableDatabase {
	return &NullableDatabase{value: val, isSet: true}
}

func (v NullableDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
