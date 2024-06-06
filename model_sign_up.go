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

// checks if the SignUp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignUp{}

// SignUp struct for SignUp
type SignUp struct {
	Id                    string           `json:"id"`
	CreatedAt             time.Time        `json:"created_at"`
	UpdatedAt             *time.Time       `json:"updated_at,omitempty"`
	FirstName             string           `json:"first_name"`
	LastName              string           `json:"last_name"`
	UserEmail             string           `json:"user_email"`
	TypeOfUse             TypeOfUseEnum    `json:"type_of_use"`
	QoveryUsage           string           `json:"qovery_usage"`
	CompanyName           NullableString   `json:"company_name,omitempty"`
	CompanySize           *CompanySizeEnum `json:"company_size,omitempty"`
	UserRole              NullableString   `json:"user_role,omitempty"`
	QoveryUsageOther      NullableString   `json:"qovery_usage_other,omitempty"`
	UserQuestions         NullableString   `json:"user_questions,omitempty"`
	CurrentStep           NullableString   `json:"current_step,omitempty"`
	DxAuth                NullableBool     `json:"dx_auth,omitempty"`
	InfrastructureHosting NullableString   `json:"infrastructure_hosting,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _SignUp SignUp

// NewSignUp instantiates a new SignUp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignUp(id string, createdAt time.Time, firstName string, lastName string, userEmail string, typeOfUse TypeOfUseEnum, qoveryUsage string) *SignUp {
	this := SignUp{}
	this.Id = id
	this.CreatedAt = createdAt
	this.FirstName = firstName
	this.LastName = lastName
	this.UserEmail = userEmail
	this.TypeOfUse = typeOfUse
	this.QoveryUsage = qoveryUsage
	return &this
}

// NewSignUpWithDefaults instantiates a new SignUp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignUpWithDefaults() *SignUp {
	this := SignUp{}
	return &this
}

// GetId returns the Id field value
func (o *SignUp) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SignUp) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SignUp) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SignUp) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SignUp) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SignUp) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SignUp) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignUp) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SignUp) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SignUp) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetFirstName returns the FirstName field value
func (o *SignUp) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *SignUp) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *SignUp) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *SignUp) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *SignUp) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *SignUp) SetLastName(v string) {
	o.LastName = v
}

// GetUserEmail returns the UserEmail field value
func (o *SignUp) GetUserEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value
// and a boolean to check if the value has been set.
func (o *SignUp) GetUserEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEmail, true
}

// SetUserEmail sets field value
func (o *SignUp) SetUserEmail(v string) {
	o.UserEmail = v
}

// GetTypeOfUse returns the TypeOfUse field value
func (o *SignUp) GetTypeOfUse() TypeOfUseEnum {
	if o == nil {
		var ret TypeOfUseEnum
		return ret
	}

	return o.TypeOfUse
}

// GetTypeOfUseOk returns a tuple with the TypeOfUse field value
// and a boolean to check if the value has been set.
func (o *SignUp) GetTypeOfUseOk() (*TypeOfUseEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeOfUse, true
}

// SetTypeOfUse sets field value
func (o *SignUp) SetTypeOfUse(v TypeOfUseEnum) {
	o.TypeOfUse = v
}

// GetQoveryUsage returns the QoveryUsage field value
func (o *SignUp) GetQoveryUsage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QoveryUsage
}

// GetQoveryUsageOk returns a tuple with the QoveryUsage field value
// and a boolean to check if the value has been set.
func (o *SignUp) GetQoveryUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QoveryUsage, true
}

// SetQoveryUsage sets field value
func (o *SignUp) SetQoveryUsage(v string) {
	o.QoveryUsage = v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignUp) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyName.Get()
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignUp) GetCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyName.Get(), o.CompanyName.IsSet()
}

// HasCompanyName returns a boolean if a field has been set.
func (o *SignUp) HasCompanyName() bool {
	if o != nil && o.CompanyName.IsSet() {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given NullableString and assigns it to the CompanyName field.
func (o *SignUp) SetCompanyName(v string) {
	o.CompanyName.Set(&v)
}

// SetCompanyNameNil sets the value for CompanyName to be an explicit nil
func (o *SignUp) SetCompanyNameNil() {
	o.CompanyName.Set(nil)
}

// UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
func (o *SignUp) UnsetCompanyName() {
	o.CompanyName.Unset()
}

// GetCompanySize returns the CompanySize field value if set, zero value otherwise.
func (o *SignUp) GetCompanySize() CompanySizeEnum {
	if o == nil || IsNil(o.CompanySize) {
		var ret CompanySizeEnum
		return ret
	}
	return *o.CompanySize
}

// GetCompanySizeOk returns a tuple with the CompanySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignUp) GetCompanySizeOk() (*CompanySizeEnum, bool) {
	if o == nil || IsNil(o.CompanySize) {
		return nil, false
	}
	return o.CompanySize, true
}

// HasCompanySize returns a boolean if a field has been set.
func (o *SignUp) HasCompanySize() bool {
	if o != nil && !IsNil(o.CompanySize) {
		return true
	}

	return false
}

// SetCompanySize gets a reference to the given CompanySizeEnum and assigns it to the CompanySize field.
func (o *SignUp) SetCompanySize(v CompanySizeEnum) {
	o.CompanySize = &v
}

// GetUserRole returns the UserRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignUp) GetUserRole() string {
	if o == nil || IsNil(o.UserRole.Get()) {
		var ret string
		return ret
	}
	return *o.UserRole.Get()
}

// GetUserRoleOk returns a tuple with the UserRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignUp) GetUserRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserRole.Get(), o.UserRole.IsSet()
}

// HasUserRole returns a boolean if a field has been set.
func (o *SignUp) HasUserRole() bool {
	if o != nil && o.UserRole.IsSet() {
		return true
	}

	return false
}

// SetUserRole gets a reference to the given NullableString and assigns it to the UserRole field.
func (o *SignUp) SetUserRole(v string) {
	o.UserRole.Set(&v)
}

// SetUserRoleNil sets the value for UserRole to be an explicit nil
func (o *SignUp) SetUserRoleNil() {
	o.UserRole.Set(nil)
}

// UnsetUserRole ensures that no value is present for UserRole, not even an explicit nil
func (o *SignUp) UnsetUserRole() {
	o.UserRole.Unset()
}

// GetQoveryUsageOther returns the QoveryUsageOther field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignUp) GetQoveryUsageOther() string {
	if o == nil || IsNil(o.QoveryUsageOther.Get()) {
		var ret string
		return ret
	}
	return *o.QoveryUsageOther.Get()
}

// GetQoveryUsageOtherOk returns a tuple with the QoveryUsageOther field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignUp) GetQoveryUsageOtherOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QoveryUsageOther.Get(), o.QoveryUsageOther.IsSet()
}

// HasQoveryUsageOther returns a boolean if a field has been set.
func (o *SignUp) HasQoveryUsageOther() bool {
	if o != nil && o.QoveryUsageOther.IsSet() {
		return true
	}

	return false
}

// SetQoveryUsageOther gets a reference to the given NullableString and assigns it to the QoveryUsageOther field.
func (o *SignUp) SetQoveryUsageOther(v string) {
	o.QoveryUsageOther.Set(&v)
}

// SetQoveryUsageOtherNil sets the value for QoveryUsageOther to be an explicit nil
func (o *SignUp) SetQoveryUsageOtherNil() {
	o.QoveryUsageOther.Set(nil)
}

// UnsetQoveryUsageOther ensures that no value is present for QoveryUsageOther, not even an explicit nil
func (o *SignUp) UnsetQoveryUsageOther() {
	o.QoveryUsageOther.Unset()
}

// GetUserQuestions returns the UserQuestions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignUp) GetUserQuestions() string {
	if o == nil || IsNil(o.UserQuestions.Get()) {
		var ret string
		return ret
	}
	return *o.UserQuestions.Get()
}

// GetUserQuestionsOk returns a tuple with the UserQuestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignUp) GetUserQuestionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserQuestions.Get(), o.UserQuestions.IsSet()
}

// HasUserQuestions returns a boolean if a field has been set.
func (o *SignUp) HasUserQuestions() bool {
	if o != nil && o.UserQuestions.IsSet() {
		return true
	}

	return false
}

// SetUserQuestions gets a reference to the given NullableString and assigns it to the UserQuestions field.
func (o *SignUp) SetUserQuestions(v string) {
	o.UserQuestions.Set(&v)
}

// SetUserQuestionsNil sets the value for UserQuestions to be an explicit nil
func (o *SignUp) SetUserQuestionsNil() {
	o.UserQuestions.Set(nil)
}

// UnsetUserQuestions ensures that no value is present for UserQuestions, not even an explicit nil
func (o *SignUp) UnsetUserQuestions() {
	o.UserQuestions.Unset()
}

// GetCurrentStep returns the CurrentStep field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignUp) GetCurrentStep() string {
	if o == nil || IsNil(o.CurrentStep.Get()) {
		var ret string
		return ret
	}
	return *o.CurrentStep.Get()
}

// GetCurrentStepOk returns a tuple with the CurrentStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignUp) GetCurrentStepOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentStep.Get(), o.CurrentStep.IsSet()
}

// HasCurrentStep returns a boolean if a field has been set.
func (o *SignUp) HasCurrentStep() bool {
	if o != nil && o.CurrentStep.IsSet() {
		return true
	}

	return false
}

// SetCurrentStep gets a reference to the given NullableString and assigns it to the CurrentStep field.
func (o *SignUp) SetCurrentStep(v string) {
	o.CurrentStep.Set(&v)
}

// SetCurrentStepNil sets the value for CurrentStep to be an explicit nil
func (o *SignUp) SetCurrentStepNil() {
	o.CurrentStep.Set(nil)
}

// UnsetCurrentStep ensures that no value is present for CurrentStep, not even an explicit nil
func (o *SignUp) UnsetCurrentStep() {
	o.CurrentStep.Unset()
}

// GetDxAuth returns the DxAuth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignUp) GetDxAuth() bool {
	if o == nil || IsNil(o.DxAuth.Get()) {
		var ret bool
		return ret
	}
	return *o.DxAuth.Get()
}

// GetDxAuthOk returns a tuple with the DxAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignUp) GetDxAuthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DxAuth.Get(), o.DxAuth.IsSet()
}

// HasDxAuth returns a boolean if a field has been set.
func (o *SignUp) HasDxAuth() bool {
	if o != nil && o.DxAuth.IsSet() {
		return true
	}

	return false
}

// SetDxAuth gets a reference to the given NullableBool and assigns it to the DxAuth field.
func (o *SignUp) SetDxAuth(v bool) {
	o.DxAuth.Set(&v)
}

// SetDxAuthNil sets the value for DxAuth to be an explicit nil
func (o *SignUp) SetDxAuthNil() {
	o.DxAuth.Set(nil)
}

// UnsetDxAuth ensures that no value is present for DxAuth, not even an explicit nil
func (o *SignUp) UnsetDxAuth() {
	o.DxAuth.Unset()
}

// GetInfrastructureHosting returns the InfrastructureHosting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignUp) GetInfrastructureHosting() string {
	if o == nil || IsNil(o.InfrastructureHosting.Get()) {
		var ret string
		return ret
	}
	return *o.InfrastructureHosting.Get()
}

// GetInfrastructureHostingOk returns a tuple with the InfrastructureHosting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignUp) GetInfrastructureHostingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfrastructureHosting.Get(), o.InfrastructureHosting.IsSet()
}

// HasInfrastructureHosting returns a boolean if a field has been set.
func (o *SignUp) HasInfrastructureHosting() bool {
	if o != nil && o.InfrastructureHosting.IsSet() {
		return true
	}

	return false
}

// SetInfrastructureHosting gets a reference to the given NullableString and assigns it to the InfrastructureHosting field.
func (o *SignUp) SetInfrastructureHosting(v string) {
	o.InfrastructureHosting.Set(&v)
}

// SetInfrastructureHostingNil sets the value for InfrastructureHosting to be an explicit nil
func (o *SignUp) SetInfrastructureHostingNil() {
	o.InfrastructureHosting.Set(nil)
}

// UnsetInfrastructureHosting ensures that no value is present for InfrastructureHosting, not even an explicit nil
func (o *SignUp) UnsetInfrastructureHosting() {
	o.InfrastructureHosting.Unset()
}

func (o SignUp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignUp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["first_name"] = o.FirstName
	toSerialize["last_name"] = o.LastName
	toSerialize["user_email"] = o.UserEmail
	toSerialize["type_of_use"] = o.TypeOfUse
	toSerialize["qovery_usage"] = o.QoveryUsage
	if o.CompanyName.IsSet() {
		toSerialize["company_name"] = o.CompanyName.Get()
	}
	if !IsNil(o.CompanySize) {
		toSerialize["company_size"] = o.CompanySize
	}
	if o.UserRole.IsSet() {
		toSerialize["user_role"] = o.UserRole.Get()
	}
	if o.QoveryUsageOther.IsSet() {
		toSerialize["qovery_usage_other"] = o.QoveryUsageOther.Get()
	}
	if o.UserQuestions.IsSet() {
		toSerialize["user_questions"] = o.UserQuestions.Get()
	}
	if o.CurrentStep.IsSet() {
		toSerialize["current_step"] = o.CurrentStep.Get()
	}
	if o.DxAuth.IsSet() {
		toSerialize["dx_auth"] = o.DxAuth.Get()
	}
	if o.InfrastructureHosting.IsSet() {
		toSerialize["infrastructure_hosting"] = o.InfrastructureHosting.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SignUp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"first_name",
		"last_name",
		"user_email",
		"type_of_use",
		"qovery_usage",
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

	varSignUp := _SignUp{}

	err = json.Unmarshal(data, &varSignUp)

	if err != nil {
		return err
	}

	*o = SignUp(varSignUp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "first_name")
		delete(additionalProperties, "last_name")
		delete(additionalProperties, "user_email")
		delete(additionalProperties, "type_of_use")
		delete(additionalProperties, "qovery_usage")
		delete(additionalProperties, "company_name")
		delete(additionalProperties, "company_size")
		delete(additionalProperties, "user_role")
		delete(additionalProperties, "qovery_usage_other")
		delete(additionalProperties, "user_questions")
		delete(additionalProperties, "current_step")
		delete(additionalProperties, "dx_auth")
		delete(additionalProperties, "infrastructure_hosting")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSignUp struct {
	value *SignUp
	isSet bool
}

func (v NullableSignUp) Get() *SignUp {
	return v.value
}

func (v *NullableSignUp) Set(val *SignUp) {
	v.value = val
	v.isSet = true
}

func (v NullableSignUp) IsSet() bool {
	return v.isSet
}

func (v *NullableSignUp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignUp(val *SignUp) *NullableSignUp {
	return &NullableSignUp{value: val, isSet: true}
}

func (v NullableSignUp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignUp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
