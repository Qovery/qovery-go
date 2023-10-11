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

// checks if the CustomDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDomain{}

// CustomDomain struct for CustomDomain
type CustomDomain struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// your custom domain
	Domain string `json:"domain"`
	// to control if a certificate has to be generated for this custom domain by Qovery. The default value is `true`. This flag should be set to `false` if a CDN or other entities are managing the certificate for the specified domain and the traffic is proxied by the CDN to Qovery.
	GenerateCertificate bool `json:"generate_certificate"`
	// URL provided by Qovery. You must create a CNAME on your DNS provider using that URL
	ValidationDomain *string                 `json:"validation_domain,omitempty"`
	Status           *CustomDomainStatusEnum `json:"status,omitempty"`
}

// NewCustomDomain instantiates a new CustomDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDomain(id string, createdAt time.Time, domain string, generateCertificate bool) *CustomDomain {
	this := CustomDomain{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Domain = domain
	this.GenerateCertificate = generateCertificate
	return &this
}

// NewCustomDomainWithDefaults instantiates a new CustomDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDomainWithDefaults() *CustomDomain {
	this := CustomDomain{}
	return &this
}

// GetId returns the Id field value
func (o *CustomDomain) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomDomain) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CustomDomain) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CustomDomain) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomDomain) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomDomain) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CustomDomain) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDomain returns the Domain field value
func (o *CustomDomain) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *CustomDomain) SetDomain(v string) {
	o.Domain = v
}

// GetGenerateCertificate returns the GenerateCertificate field value
func (o *CustomDomain) GetGenerateCertificate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GenerateCertificate
}

// GetGenerateCertificateOk returns a tuple with the GenerateCertificate field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetGenerateCertificateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GenerateCertificate, true
}

// SetGenerateCertificate sets field value
func (o *CustomDomain) SetGenerateCertificate(v bool) {
	o.GenerateCertificate = v
}

// GetValidationDomain returns the ValidationDomain field value if set, zero value otherwise.
func (o *CustomDomain) GetValidationDomain() string {
	if o == nil || IsNil(o.ValidationDomain) {
		var ret string
		return ret
	}
	return *o.ValidationDomain
}

// GetValidationDomainOk returns a tuple with the ValidationDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetValidationDomainOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationDomain) {
		return nil, false
	}
	return o.ValidationDomain, true
}

// HasValidationDomain returns a boolean if a field has been set.
func (o *CustomDomain) HasValidationDomain() bool {
	if o != nil && !IsNil(o.ValidationDomain) {
		return true
	}

	return false
}

// SetValidationDomain gets a reference to the given string and assigns it to the ValidationDomain field.
func (o *CustomDomain) SetValidationDomain(v string) {
	o.ValidationDomain = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomDomain) GetStatus() CustomDomainStatusEnum {
	if o == nil || IsNil(o.Status) {
		var ret CustomDomainStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetStatusOk() (*CustomDomainStatusEnum, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomDomain) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CustomDomainStatusEnum and assigns it to the Status field.
func (o *CustomDomain) SetStatus(v CustomDomainStatusEnum) {
	o.Status = &v
}

func (o CustomDomain) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["domain"] = o.Domain
	toSerialize["generate_certificate"] = o.GenerateCertificate
	if !IsNil(o.ValidationDomain) {
		toSerialize["validation_domain"] = o.ValidationDomain
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableCustomDomain struct {
	value *CustomDomain
	isSet bool
}

func (v NullableCustomDomain) Get() *CustomDomain {
	return v.value
}

func (v *NullableCustomDomain) Set(val *CustomDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDomain(val *CustomDomain) *NullableCustomDomain {
	return &NullableCustomDomain{value: val, isSet: true}
}

func (v NullableCustomDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
