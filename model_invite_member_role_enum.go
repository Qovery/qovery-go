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
)

// InviteMemberRoleEnum the model 'InviteMemberRoleEnum'
type InviteMemberRoleEnum string

// List of InviteMemberRoleEnum
const (
	INVITEMEMBERROLEENUM_OWNER     InviteMemberRoleEnum = "OWNER"
	INVITEMEMBERROLEENUM_ADMIN     InviteMemberRoleEnum = "ADMIN"
	INVITEMEMBERROLEENUM_DEVELOPER InviteMemberRoleEnum = "DEVELOPER"
	INVITEMEMBERROLEENUM_VIEWER    InviteMemberRoleEnum = "VIEWER"
)

// All allowed values of InviteMemberRoleEnum enum
var AllowedInviteMemberRoleEnumEnumValues = []InviteMemberRoleEnum{
	"OWNER",
	"ADMIN",
	"DEVELOPER",
	"VIEWER",
}

func (v *InviteMemberRoleEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InviteMemberRoleEnum(value)
	for _, existing := range AllowedInviteMemberRoleEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InviteMemberRoleEnum", value)
}

// NewInviteMemberRoleEnumFromValue returns a pointer to a valid InviteMemberRoleEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInviteMemberRoleEnumFromValue(v string) (*InviteMemberRoleEnum, error) {
	ev := InviteMemberRoleEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InviteMemberRoleEnum: valid values are %v", v, AllowedInviteMemberRoleEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InviteMemberRoleEnum) IsValid() bool {
	for _, existing := range AllowedInviteMemberRoleEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InviteMemberRoleEnum value
func (v InviteMemberRoleEnum) Ptr() *InviteMemberRoleEnum {
	return &v
}

type NullableInviteMemberRoleEnum struct {
	value *InviteMemberRoleEnum
	isSet bool
}

func (v NullableInviteMemberRoleEnum) Get() *InviteMemberRoleEnum {
	return v.value
}

func (v *NullableInviteMemberRoleEnum) Set(val *InviteMemberRoleEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteMemberRoleEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteMemberRoleEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteMemberRoleEnum(val *InviteMemberRoleEnum) *NullableInviteMemberRoleEnum {
	return &NullableInviteMemberRoleEnum{value: val, isSet: true}
}

func (v NullableInviteMemberRoleEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteMemberRoleEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
