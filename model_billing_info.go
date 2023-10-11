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

// checks if the BillingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInfo{}

// BillingInfo struct for BillingInfo
type BillingInfo struct {
	FirstName NullableString `json:"first_name,omitempty"`
	LastName  NullableString `json:"last_name,omitempty"`
	// email used for billing, and to receive all invoices by email
	Email   NullableString `json:"email,omitempty"`
	Address NullableString `json:"address,omitempty"`
	City    NullableString `json:"city,omitempty"`
	Zip     NullableString `json:"zip,omitempty"`
	// only for US
	State NullableString `json:"state,omitempty"`
	// ISO code of the country
	CountryCode NullableString `json:"country_code,omitempty"`
	// name of the company to bill
	Company   NullableString `json:"company,omitempty"`
	VatNumber NullableString `json:"vat_number,omitempty"`
}

// NewBillingInfo instantiates a new BillingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfo() *BillingInfo {
	this := BillingInfo{}
	return &this
}

// NewBillingInfoWithDefaults instantiates a new BillingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoWithDefaults() *BillingInfo {
	this := BillingInfo{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfo) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfo) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *BillingInfo) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *BillingInfo) SetFirstName(v string) {
	o.FirstName.Set(&v)
}

// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *BillingInfo) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *BillingInfo) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfo) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfo) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *BillingInfo) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *BillingInfo) SetLastName(v string) {
	o.LastName.Set(&v)
}

// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *BillingInfo) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *BillingInfo) UnsetLastName() {
	o.LastName.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfo) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfo) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *BillingInfo) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *BillingInfo) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *BillingInfo) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *BillingInfo) UnsetEmail() {
	o.Email.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfo) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfo) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *BillingInfo) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *BillingInfo) SetAddress(v string) {
	o.Address.Set(&v)
}

// SetAddressNil sets the value for Address to be an explicit nil
func (o *BillingInfo) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *BillingInfo) UnsetAddress() {
	o.Address.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfo) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfo) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *BillingInfo) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *BillingInfo) SetCity(v string) {
	o.City.Set(&v)
}

// SetCityNil sets the value for City to be an explicit nil
func (o *BillingInfo) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *BillingInfo) UnsetCity() {
	o.City.Unset()
}

// GetZip returns the Zip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfo) GetZip() string {
	if o == nil || IsNil(o.Zip.Get()) {
		var ret string
		return ret
	}
	return *o.Zip.Get()
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfo) GetZipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zip.Get(), o.Zip.IsSet()
}

// HasZip returns a boolean if a field has been set.
func (o *BillingInfo) HasZip() bool {
	if o != nil && o.Zip.IsSet() {
		return true
	}

	return false
}

// SetZip gets a reference to the given NullableString and assigns it to the Zip field.
func (o *BillingInfo) SetZip(v string) {
	o.Zip.Set(&v)
}

// SetZipNil sets the value for Zip to be an explicit nil
func (o *BillingInfo) SetZipNil() {
	o.Zip.Set(nil)
}

// UnsetZip ensures that no value is present for Zip, not even an explicit nil
func (o *BillingInfo) UnsetZip() {
	o.Zip.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfo) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfo) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *BillingInfo) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *BillingInfo) SetState(v string) {
	o.State.Set(&v)
}

// SetStateNil sets the value for State to be an explicit nil
func (o *BillingInfo) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *BillingInfo) UnsetState() {
	o.State.Unset()
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfo) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfo) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// HasCountryCode returns a boolean if a field has been set.
func (o *BillingInfo) HasCountryCode() bool {
	if o != nil && o.CountryCode.IsSet() {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given NullableString and assigns it to the CountryCode field.
func (o *BillingInfo) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}

// SetCountryCodeNil sets the value for CountryCode to be an explicit nil
func (o *BillingInfo) SetCountryCodeNil() {
	o.CountryCode.Set(nil)
}

// UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
func (o *BillingInfo) UnsetCountryCode() {
	o.CountryCode.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfo) GetCompany() string {
	if o == nil || IsNil(o.Company.Get()) {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfo) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *BillingInfo) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *BillingInfo) SetCompany(v string) {
	o.Company.Set(&v)
}

// SetCompanyNil sets the value for Company to be an explicit nil
func (o *BillingInfo) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *BillingInfo) UnsetCompany() {
	o.Company.Unset()
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfo) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber.Get()) {
		var ret string
		return ret
	}
	return *o.VatNumber.Get()
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfo) GetVatNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VatNumber.Get(), o.VatNumber.IsSet()
}

// HasVatNumber returns a boolean if a field has been set.
func (o *BillingInfo) HasVatNumber() bool {
	if o != nil && o.VatNumber.IsSet() {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given NullableString and assigns it to the VatNumber field.
func (o *BillingInfo) SetVatNumber(v string) {
	o.VatNumber.Set(&v)
}

// SetVatNumberNil sets the value for VatNumber to be an explicit nil
func (o *BillingInfo) SetVatNumberNil() {
	o.VatNumber.Set(nil)
}

// UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
func (o *BillingInfo) UnsetVatNumber() {
	o.VatNumber.Unset()
}

func (o BillingInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.Zip.IsSet() {
		toSerialize["zip"] = o.Zip.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.CountryCode.IsSet() {
		toSerialize["country_code"] = o.CountryCode.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.VatNumber.IsSet() {
		toSerialize["vat_number"] = o.VatNumber.Get()
	}
	return toSerialize, nil
}

type NullableBillingInfo struct {
	value *BillingInfo
	isSet bool
}

func (v NullableBillingInfo) Get() *BillingInfo {
	return v.value
}

func (v *NullableBillingInfo) Set(val *BillingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfo(val *BillingInfo) *NullableBillingInfo {
	return &NullableBillingInfo{value: val, isSet: true}
}

func (v NullableBillingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
