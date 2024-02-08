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

// checks if the ClusterRequestFeaturesInnerValueOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterRequestFeaturesInnerValueOneOf{}

// ClusterRequestFeaturesInnerValueOneOf struct for ClusterRequestFeaturesInnerValueOneOf
type ClusterRequestFeaturesInnerValueOneOf struct {
	AwsVpcEksId                string   `json:"aws_vpc_eks_id"`
	EksSubnetsZoneAIds         []string `json:"eks_subnets_zone_a_ids"`
	EksSubnetsZoneBIds         []string `json:"eks_subnets_zone_b_ids"`
	EksSubnetsZoneCIds         []string `json:"eks_subnets_zone_c_ids"`
	DocumentdbSubnetsZoneAIds  []string `json:"documentdb_subnets_zone_a_ids,omitempty"`
	DocumentdbSubnetsZoneBIds  []string `json:"documentdb_subnets_zone_b_ids,omitempty"`
	DocumentdbSubnetsZoneCIds  []string `json:"documentdb_subnets_zone_c_ids,omitempty"`
	ElasticacheSubnetsZoneAIds []string `json:"elasticache_subnets_zone_a_ids,omitempty"`
	ElasticacheSubnetsZoneBIds []string `json:"elasticache_subnets_zone_b_ids,omitempty"`
	ElasticacheSubnetsZoneCIds []string `json:"elasticache_subnets_zone_c_ids,omitempty"`
	RdsSubnetsZoneAIds         []string `json:"rds_subnets_zone_a_ids,omitempty"`
	RdsSubnetsZoneBIds         []string `json:"rds_subnets_zone_b_ids,omitempty"`
	RdsSubnetsZoneCIds         []string `json:"rds_subnets_zone_c_ids,omitempty"`
}

// NewClusterRequestFeaturesInnerValueOneOf instantiates a new ClusterRequestFeaturesInnerValueOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRequestFeaturesInnerValueOneOf(awsVpcEksId string, eksSubnetsZoneAIds []string, eksSubnetsZoneBIds []string, eksSubnetsZoneCIds []string) *ClusterRequestFeaturesInnerValueOneOf {
	this := ClusterRequestFeaturesInnerValueOneOf{}
	this.AwsVpcEksId = awsVpcEksId
	this.EksSubnetsZoneAIds = eksSubnetsZoneAIds
	this.EksSubnetsZoneBIds = eksSubnetsZoneBIds
	this.EksSubnetsZoneCIds = eksSubnetsZoneCIds
	return &this
}

// NewClusterRequestFeaturesInnerValueOneOfWithDefaults instantiates a new ClusterRequestFeaturesInnerValueOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterRequestFeaturesInnerValueOneOfWithDefaults() *ClusterRequestFeaturesInnerValueOneOf {
	this := ClusterRequestFeaturesInnerValueOneOf{}
	return &this
}

// GetAwsVpcEksId returns the AwsVpcEksId field value
func (o *ClusterRequestFeaturesInnerValueOneOf) GetAwsVpcEksId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsVpcEksId
}

// GetAwsVpcEksIdOk returns a tuple with the AwsVpcEksId field value
// and a boolean to check if the value has been set.
func (o *ClusterRequestFeaturesInnerValueOneOf) GetAwsVpcEksIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsVpcEksId, true
}

// SetAwsVpcEksId sets field value
func (o *ClusterRequestFeaturesInnerValueOneOf) SetAwsVpcEksId(v string) {
	o.AwsVpcEksId = v
}

// GetEksSubnetsZoneAIds returns the EksSubnetsZoneAIds field value
func (o *ClusterRequestFeaturesInnerValueOneOf) GetEksSubnetsZoneAIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EksSubnetsZoneAIds
}

// GetEksSubnetsZoneAIdsOk returns a tuple with the EksSubnetsZoneAIds field value
// and a boolean to check if the value has been set.
func (o *ClusterRequestFeaturesInnerValueOneOf) GetEksSubnetsZoneAIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EksSubnetsZoneAIds, true
}

// SetEksSubnetsZoneAIds sets field value
func (o *ClusterRequestFeaturesInnerValueOneOf) SetEksSubnetsZoneAIds(v []string) {
	o.EksSubnetsZoneAIds = v
}

// GetEksSubnetsZoneBIds returns the EksSubnetsZoneBIds field value
func (o *ClusterRequestFeaturesInnerValueOneOf) GetEksSubnetsZoneBIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EksSubnetsZoneBIds
}

// GetEksSubnetsZoneBIdsOk returns a tuple with the EksSubnetsZoneBIds field value
// and a boolean to check if the value has been set.
func (o *ClusterRequestFeaturesInnerValueOneOf) GetEksSubnetsZoneBIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EksSubnetsZoneBIds, true
}

// SetEksSubnetsZoneBIds sets field value
func (o *ClusterRequestFeaturesInnerValueOneOf) SetEksSubnetsZoneBIds(v []string) {
	o.EksSubnetsZoneBIds = v
}

// GetEksSubnetsZoneCIds returns the EksSubnetsZoneCIds field value
func (o *ClusterRequestFeaturesInnerValueOneOf) GetEksSubnetsZoneCIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EksSubnetsZoneCIds
}

// GetEksSubnetsZoneCIdsOk returns a tuple with the EksSubnetsZoneCIds field value
// and a boolean to check if the value has been set.
func (o *ClusterRequestFeaturesInnerValueOneOf) GetEksSubnetsZoneCIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EksSubnetsZoneCIds, true
}

// SetEksSubnetsZoneCIds sets field value
func (o *ClusterRequestFeaturesInnerValueOneOf) SetEksSubnetsZoneCIds(v []string) {
	o.EksSubnetsZoneCIds = v
}

// GetDocumentdbSubnetsZoneAIds returns the DocumentdbSubnetsZoneAIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterRequestFeaturesInnerValueOneOf) GetDocumentdbSubnetsZoneAIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DocumentdbSubnetsZoneAIds
}

// GetDocumentdbSubnetsZoneAIdsOk returns a tuple with the DocumentdbSubnetsZoneAIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterRequestFeaturesInnerValueOneOf) GetDocumentdbSubnetsZoneAIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DocumentdbSubnetsZoneAIds) {
		return nil, false
	}
	return o.DocumentdbSubnetsZoneAIds, true
}

// HasDocumentdbSubnetsZoneAIds returns a boolean if a field has been set.
func (o *ClusterRequestFeaturesInnerValueOneOf) HasDocumentdbSubnetsZoneAIds() bool {
	if o != nil && IsNil(o.DocumentdbSubnetsZoneAIds) {
		return true
	}

	return false
}

// SetDocumentdbSubnetsZoneAIds gets a reference to the given []string and assigns it to the DocumentdbSubnetsZoneAIds field.
func (o *ClusterRequestFeaturesInnerValueOneOf) SetDocumentdbSubnetsZoneAIds(v []string) {
	o.DocumentdbSubnetsZoneAIds = v
}

// GetDocumentdbSubnetsZoneBIds returns the DocumentdbSubnetsZoneBIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterRequestFeaturesInnerValueOneOf) GetDocumentdbSubnetsZoneBIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DocumentdbSubnetsZoneBIds
}

// GetDocumentdbSubnetsZoneBIdsOk returns a tuple with the DocumentdbSubnetsZoneBIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterRequestFeaturesInnerValueOneOf) GetDocumentdbSubnetsZoneBIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DocumentdbSubnetsZoneBIds) {
		return nil, false
	}
	return o.DocumentdbSubnetsZoneBIds, true
}

// HasDocumentdbSubnetsZoneBIds returns a boolean if a field has been set.
func (o *ClusterRequestFeaturesInnerValueOneOf) HasDocumentdbSubnetsZoneBIds() bool {
	if o != nil && IsNil(o.DocumentdbSubnetsZoneBIds) {
		return true
	}

	return false
}

// SetDocumentdbSubnetsZoneBIds gets a reference to the given []string and assigns it to the DocumentdbSubnetsZoneBIds field.
func (o *ClusterRequestFeaturesInnerValueOneOf) SetDocumentdbSubnetsZoneBIds(v []string) {
	o.DocumentdbSubnetsZoneBIds = v
}

// GetDocumentdbSubnetsZoneCIds returns the DocumentdbSubnetsZoneCIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterRequestFeaturesInnerValueOneOf) GetDocumentdbSubnetsZoneCIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DocumentdbSubnetsZoneCIds
}

// GetDocumentdbSubnetsZoneCIdsOk returns a tuple with the DocumentdbSubnetsZoneCIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterRequestFeaturesInnerValueOneOf) GetDocumentdbSubnetsZoneCIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DocumentdbSubnetsZoneCIds) {
		return nil, false
	}
	return o.DocumentdbSubnetsZoneCIds, true
}

// HasDocumentdbSubnetsZoneCIds returns a boolean if a field has been set.
func (o *ClusterRequestFeaturesInnerValueOneOf) HasDocumentdbSubnetsZoneCIds() bool {
	if o != nil && IsNil(o.DocumentdbSubnetsZoneCIds) {
		return true
	}

	return false
}

// SetDocumentdbSubnetsZoneCIds gets a reference to the given []string and assigns it to the DocumentdbSubnetsZoneCIds field.
func (o *ClusterRequestFeaturesInnerValueOneOf) SetDocumentdbSubnetsZoneCIds(v []string) {
	o.DocumentdbSubnetsZoneCIds = v
}

// GetElasticacheSubnetsZoneAIds returns the ElasticacheSubnetsZoneAIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterRequestFeaturesInnerValueOneOf) GetElasticacheSubnetsZoneAIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ElasticacheSubnetsZoneAIds
}

// GetElasticacheSubnetsZoneAIdsOk returns a tuple with the ElasticacheSubnetsZoneAIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterRequestFeaturesInnerValueOneOf) GetElasticacheSubnetsZoneAIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ElasticacheSubnetsZoneAIds) {
		return nil, false
	}
	return o.ElasticacheSubnetsZoneAIds, true
}

// HasElasticacheSubnetsZoneAIds returns a boolean if a field has been set.
func (o *ClusterRequestFeaturesInnerValueOneOf) HasElasticacheSubnetsZoneAIds() bool {
	if o != nil && IsNil(o.ElasticacheSubnetsZoneAIds) {
		return true
	}

	return false
}

// SetElasticacheSubnetsZoneAIds gets a reference to the given []string and assigns it to the ElasticacheSubnetsZoneAIds field.
func (o *ClusterRequestFeaturesInnerValueOneOf) SetElasticacheSubnetsZoneAIds(v []string) {
	o.ElasticacheSubnetsZoneAIds = v
}

// GetElasticacheSubnetsZoneBIds returns the ElasticacheSubnetsZoneBIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterRequestFeaturesInnerValueOneOf) GetElasticacheSubnetsZoneBIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ElasticacheSubnetsZoneBIds
}

// GetElasticacheSubnetsZoneBIdsOk returns a tuple with the ElasticacheSubnetsZoneBIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterRequestFeaturesInnerValueOneOf) GetElasticacheSubnetsZoneBIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ElasticacheSubnetsZoneBIds) {
		return nil, false
	}
	return o.ElasticacheSubnetsZoneBIds, true
}

// HasElasticacheSubnetsZoneBIds returns a boolean if a field has been set.
func (o *ClusterRequestFeaturesInnerValueOneOf) HasElasticacheSubnetsZoneBIds() bool {
	if o != nil && IsNil(o.ElasticacheSubnetsZoneBIds) {
		return true
	}

	return false
}

// SetElasticacheSubnetsZoneBIds gets a reference to the given []string and assigns it to the ElasticacheSubnetsZoneBIds field.
func (o *ClusterRequestFeaturesInnerValueOneOf) SetElasticacheSubnetsZoneBIds(v []string) {
	o.ElasticacheSubnetsZoneBIds = v
}

// GetElasticacheSubnetsZoneCIds returns the ElasticacheSubnetsZoneCIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterRequestFeaturesInnerValueOneOf) GetElasticacheSubnetsZoneCIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ElasticacheSubnetsZoneCIds
}

// GetElasticacheSubnetsZoneCIdsOk returns a tuple with the ElasticacheSubnetsZoneCIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterRequestFeaturesInnerValueOneOf) GetElasticacheSubnetsZoneCIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ElasticacheSubnetsZoneCIds) {
		return nil, false
	}
	return o.ElasticacheSubnetsZoneCIds, true
}

// HasElasticacheSubnetsZoneCIds returns a boolean if a field has been set.
func (o *ClusterRequestFeaturesInnerValueOneOf) HasElasticacheSubnetsZoneCIds() bool {
	if o != nil && IsNil(o.ElasticacheSubnetsZoneCIds) {
		return true
	}

	return false
}

// SetElasticacheSubnetsZoneCIds gets a reference to the given []string and assigns it to the ElasticacheSubnetsZoneCIds field.
func (o *ClusterRequestFeaturesInnerValueOneOf) SetElasticacheSubnetsZoneCIds(v []string) {
	o.ElasticacheSubnetsZoneCIds = v
}

// GetRdsSubnetsZoneAIds returns the RdsSubnetsZoneAIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterRequestFeaturesInnerValueOneOf) GetRdsSubnetsZoneAIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RdsSubnetsZoneAIds
}

// GetRdsSubnetsZoneAIdsOk returns a tuple with the RdsSubnetsZoneAIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterRequestFeaturesInnerValueOneOf) GetRdsSubnetsZoneAIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RdsSubnetsZoneAIds) {
		return nil, false
	}
	return o.RdsSubnetsZoneAIds, true
}

// HasRdsSubnetsZoneAIds returns a boolean if a field has been set.
func (o *ClusterRequestFeaturesInnerValueOneOf) HasRdsSubnetsZoneAIds() bool {
	if o != nil && IsNil(o.RdsSubnetsZoneAIds) {
		return true
	}

	return false
}

// SetRdsSubnetsZoneAIds gets a reference to the given []string and assigns it to the RdsSubnetsZoneAIds field.
func (o *ClusterRequestFeaturesInnerValueOneOf) SetRdsSubnetsZoneAIds(v []string) {
	o.RdsSubnetsZoneAIds = v
}

// GetRdsSubnetsZoneBIds returns the RdsSubnetsZoneBIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterRequestFeaturesInnerValueOneOf) GetRdsSubnetsZoneBIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RdsSubnetsZoneBIds
}

// GetRdsSubnetsZoneBIdsOk returns a tuple with the RdsSubnetsZoneBIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterRequestFeaturesInnerValueOneOf) GetRdsSubnetsZoneBIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RdsSubnetsZoneBIds) {
		return nil, false
	}
	return o.RdsSubnetsZoneBIds, true
}

// HasRdsSubnetsZoneBIds returns a boolean if a field has been set.
func (o *ClusterRequestFeaturesInnerValueOneOf) HasRdsSubnetsZoneBIds() bool {
	if o != nil && IsNil(o.RdsSubnetsZoneBIds) {
		return true
	}

	return false
}

// SetRdsSubnetsZoneBIds gets a reference to the given []string and assigns it to the RdsSubnetsZoneBIds field.
func (o *ClusterRequestFeaturesInnerValueOneOf) SetRdsSubnetsZoneBIds(v []string) {
	o.RdsSubnetsZoneBIds = v
}

// GetRdsSubnetsZoneCIds returns the RdsSubnetsZoneCIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterRequestFeaturesInnerValueOneOf) GetRdsSubnetsZoneCIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RdsSubnetsZoneCIds
}

// GetRdsSubnetsZoneCIdsOk returns a tuple with the RdsSubnetsZoneCIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterRequestFeaturesInnerValueOneOf) GetRdsSubnetsZoneCIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RdsSubnetsZoneCIds) {
		return nil, false
	}
	return o.RdsSubnetsZoneCIds, true
}

// HasRdsSubnetsZoneCIds returns a boolean if a field has been set.
func (o *ClusterRequestFeaturesInnerValueOneOf) HasRdsSubnetsZoneCIds() bool {
	if o != nil && IsNil(o.RdsSubnetsZoneCIds) {
		return true
	}

	return false
}

// SetRdsSubnetsZoneCIds gets a reference to the given []string and assigns it to the RdsSubnetsZoneCIds field.
func (o *ClusterRequestFeaturesInnerValueOneOf) SetRdsSubnetsZoneCIds(v []string) {
	o.RdsSubnetsZoneCIds = v
}

func (o ClusterRequestFeaturesInnerValueOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterRequestFeaturesInnerValueOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aws_vpc_eks_id"] = o.AwsVpcEksId
	toSerialize["eks_subnets_zone_a_ids"] = o.EksSubnetsZoneAIds
	toSerialize["eks_subnets_zone_b_ids"] = o.EksSubnetsZoneBIds
	toSerialize["eks_subnets_zone_c_ids"] = o.EksSubnetsZoneCIds
	if o.DocumentdbSubnetsZoneAIds != nil {
		toSerialize["documentdb_subnets_zone_a_ids"] = o.DocumentdbSubnetsZoneAIds
	}
	if o.DocumentdbSubnetsZoneBIds != nil {
		toSerialize["documentdb_subnets_zone_b_ids"] = o.DocumentdbSubnetsZoneBIds
	}
	if o.DocumentdbSubnetsZoneCIds != nil {
		toSerialize["documentdb_subnets_zone_c_ids"] = o.DocumentdbSubnetsZoneCIds
	}
	if o.ElasticacheSubnetsZoneAIds != nil {
		toSerialize["elasticache_subnets_zone_a_ids"] = o.ElasticacheSubnetsZoneAIds
	}
	if o.ElasticacheSubnetsZoneBIds != nil {
		toSerialize["elasticache_subnets_zone_b_ids"] = o.ElasticacheSubnetsZoneBIds
	}
	if o.ElasticacheSubnetsZoneCIds != nil {
		toSerialize["elasticache_subnets_zone_c_ids"] = o.ElasticacheSubnetsZoneCIds
	}
	if o.RdsSubnetsZoneAIds != nil {
		toSerialize["rds_subnets_zone_a_ids"] = o.RdsSubnetsZoneAIds
	}
	if o.RdsSubnetsZoneBIds != nil {
		toSerialize["rds_subnets_zone_b_ids"] = o.RdsSubnetsZoneBIds
	}
	if o.RdsSubnetsZoneCIds != nil {
		toSerialize["rds_subnets_zone_c_ids"] = o.RdsSubnetsZoneCIds
	}
	return toSerialize, nil
}

type NullableClusterRequestFeaturesInnerValueOneOf struct {
	value *ClusterRequestFeaturesInnerValueOneOf
	isSet bool
}

func (v NullableClusterRequestFeaturesInnerValueOneOf) Get() *ClusterRequestFeaturesInnerValueOneOf {
	return v.value
}

func (v *NullableClusterRequestFeaturesInnerValueOneOf) Set(val *ClusterRequestFeaturesInnerValueOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRequestFeaturesInnerValueOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRequestFeaturesInnerValueOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRequestFeaturesInnerValueOneOf(val *ClusterRequestFeaturesInnerValueOneOf) *NullableClusterRequestFeaturesInnerValueOneOf {
	return &NullableClusterRequestFeaturesInnerValueOneOf{value: val, isSet: true}
}

func (v NullableClusterRequestFeaturesInnerValueOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRequestFeaturesInnerValueOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
