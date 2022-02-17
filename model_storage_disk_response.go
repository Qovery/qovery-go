/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"time"
)

// StorageDiskResponse struct for StorageDiskResponse
type StorageDiskResponse struct {
	CreatedAt                 *time.Time `json:"created_at,omitempty"`
	StorageId                 *string    `json:"storage_id,omitempty"`
	RequestedInGb             *int32     `json:"requested_in_gb,omitempty"`
	ConsumedInGb              *float32   `json:"consumed_in_gb,omitempty"`
	ConsumedInPercent         *float32   `json:"consumed_in_percent,omitempty"`
	WarningThresholdInPercent *float32   `json:"warning_threshold_in_percent,omitempty"`
	AlertThresholdInPercent   *float32   `json:"alert_threshold_in_percent,omitempty"`
	Status                    *string    `json:"status,omitempty"`
}

// NewStorageDiskResponse instantiates a new StorageDiskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageDiskResponse() *StorageDiskResponse {
	this := StorageDiskResponse{}
	return &this
}

// NewStorageDiskResponseWithDefaults instantiates a new StorageDiskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageDiskResponseWithDefaults() *StorageDiskResponse {
	this := StorageDiskResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *StorageDiskResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StorageDiskResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *StorageDiskResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetStorageId returns the StorageId field value if set, zero value otherwise.
func (o *StorageDiskResponse) GetStorageId() string {
	if o == nil || o.StorageId == nil {
		var ret string
		return ret
	}
	return *o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskResponse) GetStorageIdOk() (*string, bool) {
	if o == nil || o.StorageId == nil {
		return nil, false
	}
	return o.StorageId, true
}

// HasStorageId returns a boolean if a field has been set.
func (o *StorageDiskResponse) HasStorageId() bool {
	if o != nil && o.StorageId != nil {
		return true
	}

	return false
}

// SetStorageId gets a reference to the given string and assigns it to the StorageId field.
func (o *StorageDiskResponse) SetStorageId(v string) {
	o.StorageId = &v
}

// GetRequestedInGb returns the RequestedInGb field value if set, zero value otherwise.
func (o *StorageDiskResponse) GetRequestedInGb() int32 {
	if o == nil || o.RequestedInGb == nil {
		var ret int32
		return ret
	}
	return *o.RequestedInGb
}

// GetRequestedInGbOk returns a tuple with the RequestedInGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskResponse) GetRequestedInGbOk() (*int32, bool) {
	if o == nil || o.RequestedInGb == nil {
		return nil, false
	}
	return o.RequestedInGb, true
}

// HasRequestedInGb returns a boolean if a field has been set.
func (o *StorageDiskResponse) HasRequestedInGb() bool {
	if o != nil && o.RequestedInGb != nil {
		return true
	}

	return false
}

// SetRequestedInGb gets a reference to the given int32 and assigns it to the RequestedInGb field.
func (o *StorageDiskResponse) SetRequestedInGb(v int32) {
	o.RequestedInGb = &v
}

// GetConsumedInGb returns the ConsumedInGb field value if set, zero value otherwise.
func (o *StorageDiskResponse) GetConsumedInGb() float32 {
	if o == nil || o.ConsumedInGb == nil {
		var ret float32
		return ret
	}
	return *o.ConsumedInGb
}

// GetConsumedInGbOk returns a tuple with the ConsumedInGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskResponse) GetConsumedInGbOk() (*float32, bool) {
	if o == nil || o.ConsumedInGb == nil {
		return nil, false
	}
	return o.ConsumedInGb, true
}

// HasConsumedInGb returns a boolean if a field has been set.
func (o *StorageDiskResponse) HasConsumedInGb() bool {
	if o != nil && o.ConsumedInGb != nil {
		return true
	}

	return false
}

// SetConsumedInGb gets a reference to the given float32 and assigns it to the ConsumedInGb field.
func (o *StorageDiskResponse) SetConsumedInGb(v float32) {
	o.ConsumedInGb = &v
}

// GetConsumedInPercent returns the ConsumedInPercent field value if set, zero value otherwise.
func (o *StorageDiskResponse) GetConsumedInPercent() float32 {
	if o == nil || o.ConsumedInPercent == nil {
		var ret float32
		return ret
	}
	return *o.ConsumedInPercent
}

// GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskResponse) GetConsumedInPercentOk() (*float32, bool) {
	if o == nil || o.ConsumedInPercent == nil {
		return nil, false
	}
	return o.ConsumedInPercent, true
}

// HasConsumedInPercent returns a boolean if a field has been set.
func (o *StorageDiskResponse) HasConsumedInPercent() bool {
	if o != nil && o.ConsumedInPercent != nil {
		return true
	}

	return false
}

// SetConsumedInPercent gets a reference to the given float32 and assigns it to the ConsumedInPercent field.
func (o *StorageDiskResponse) SetConsumedInPercent(v float32) {
	o.ConsumedInPercent = &v
}

// GetWarningThresholdInPercent returns the WarningThresholdInPercent field value if set, zero value otherwise.
func (o *StorageDiskResponse) GetWarningThresholdInPercent() float32 {
	if o == nil || o.WarningThresholdInPercent == nil {
		var ret float32
		return ret
	}
	return *o.WarningThresholdInPercent
}

// GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskResponse) GetWarningThresholdInPercentOk() (*float32, bool) {
	if o == nil || o.WarningThresholdInPercent == nil {
		return nil, false
	}
	return o.WarningThresholdInPercent, true
}

// HasWarningThresholdInPercent returns a boolean if a field has been set.
func (o *StorageDiskResponse) HasWarningThresholdInPercent() bool {
	if o != nil && o.WarningThresholdInPercent != nil {
		return true
	}

	return false
}

// SetWarningThresholdInPercent gets a reference to the given float32 and assigns it to the WarningThresholdInPercent field.
func (o *StorageDiskResponse) SetWarningThresholdInPercent(v float32) {
	o.WarningThresholdInPercent = &v
}

// GetAlertThresholdInPercent returns the AlertThresholdInPercent field value if set, zero value otherwise.
func (o *StorageDiskResponse) GetAlertThresholdInPercent() float32 {
	if o == nil || o.AlertThresholdInPercent == nil {
		var ret float32
		return ret
	}
	return *o.AlertThresholdInPercent
}

// GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskResponse) GetAlertThresholdInPercentOk() (*float32, bool) {
	if o == nil || o.AlertThresholdInPercent == nil {
		return nil, false
	}
	return o.AlertThresholdInPercent, true
}

// HasAlertThresholdInPercent returns a boolean if a field has been set.
func (o *StorageDiskResponse) HasAlertThresholdInPercent() bool {
	if o != nil && o.AlertThresholdInPercent != nil {
		return true
	}

	return false
}

// SetAlertThresholdInPercent gets a reference to the given float32 and assigns it to the AlertThresholdInPercent field.
func (o *StorageDiskResponse) SetAlertThresholdInPercent(v float32) {
	o.AlertThresholdInPercent = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StorageDiskResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StorageDiskResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StorageDiskResponse) SetStatus(v string) {
	o.Status = &v
}

func (o StorageDiskResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.StorageId != nil {
		toSerialize["storage_id"] = o.StorageId
	}
	if o.RequestedInGb != nil {
		toSerialize["requested_in_gb"] = o.RequestedInGb
	}
	if o.ConsumedInGb != nil {
		toSerialize["consumed_in_gb"] = o.ConsumedInGb
	}
	if o.ConsumedInPercent != nil {
		toSerialize["consumed_in_percent"] = o.ConsumedInPercent
	}
	if o.WarningThresholdInPercent != nil {
		toSerialize["warning_threshold_in_percent"] = o.WarningThresholdInPercent
	}
	if o.AlertThresholdInPercent != nil {
		toSerialize["alert_threshold_in_percent"] = o.AlertThresholdInPercent
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableStorageDiskResponse struct {
	value *StorageDiskResponse
	isSet bool
}

func (v NullableStorageDiskResponse) Get() *StorageDiskResponse {
	return v.value
}

func (v *NullableStorageDiskResponse) Set(val *StorageDiskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDiskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDiskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDiskResponse(val *StorageDiskResponse) *NullableStorageDiskResponse {
	return &NullableStorageDiskResponse{value: val, isSet: true}
}

func (v NullableStorageDiskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDiskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
