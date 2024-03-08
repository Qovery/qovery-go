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

// ClusterFeatureValue - struct for ClusterFeatureValue
type ClusterFeatureValue struct {
	ClusterFeatureAwsExistingVpc *ClusterFeatureAwsExistingVpc
	ClusterFeatureGcpExistingVpc *ClusterFeatureGcpExistingVpc
	Bool                         *bool
	String                       *string
}

// ClusterFeatureAwsExistingVpcAsClusterFeatureValue is a convenience function that returns ClusterFeatureAwsExistingVpc wrapped in ClusterFeatureValue
func ClusterFeatureAwsExistingVpcAsClusterFeatureValue(v *ClusterFeatureAwsExistingVpc) ClusterFeatureValue {
	return ClusterFeatureValue{
		ClusterFeatureAwsExistingVpc: v,
	}
}

// ClusterFeatureGcpExistingVpcAsClusterFeatureValue is a convenience function that returns ClusterFeatureGcpExistingVpc wrapped in ClusterFeatureValue
func ClusterFeatureGcpExistingVpcAsClusterFeatureValue(v *ClusterFeatureGcpExistingVpc) ClusterFeatureValue {
	return ClusterFeatureValue{
		ClusterFeatureGcpExistingVpc: v,
	}
}

// boolAsClusterFeatureValue is a convenience function that returns bool wrapped in ClusterFeatureValue
func BoolAsClusterFeatureValue(v *bool) ClusterFeatureValue {
	return ClusterFeatureValue{
		Bool: v,
	}
}

// stringAsClusterFeatureValue is a convenience function that returns string wrapped in ClusterFeatureValue
func StringAsClusterFeatureValue(v *string) ClusterFeatureValue {
	return ClusterFeatureValue{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ClusterFeatureValue) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into ClusterFeatureAwsExistingVpc
	err = newStrictDecoder(data).Decode(&dst.ClusterFeatureAwsExistingVpc)
	if err == nil {
		jsonClusterFeatureAwsExistingVpc, _ := json.Marshal(dst.ClusterFeatureAwsExistingVpc)
		if string(jsonClusterFeatureAwsExistingVpc) == "{}" { // empty struct
			dst.ClusterFeatureAwsExistingVpc = nil
		} else {
			match++
		}
	} else {
		dst.ClusterFeatureAwsExistingVpc = nil
	}

	// try to unmarshal data into ClusterFeatureGcpExistingVpc
	err = newStrictDecoder(data).Decode(&dst.ClusterFeatureGcpExistingVpc)
	if err == nil {
		jsonClusterFeatureGcpExistingVpc, _ := json.Marshal(dst.ClusterFeatureGcpExistingVpc)
		if string(jsonClusterFeatureGcpExistingVpc) == "{}" { // empty struct
			dst.ClusterFeatureGcpExistingVpc = nil
		} else {
			match++
		}
	} else {
		dst.ClusterFeatureGcpExistingVpc = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ClusterFeatureAwsExistingVpc = nil
		dst.ClusterFeatureGcpExistingVpc = nil
		dst.Bool = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ClusterFeatureValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ClusterFeatureValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ClusterFeatureValue) MarshalJSON() ([]byte, error) {
	if src.ClusterFeatureAwsExistingVpc != nil {
		return json.Marshal(&src.ClusterFeatureAwsExistingVpc)
	}

	if src.ClusterFeatureGcpExistingVpc != nil {
		return json.Marshal(&src.ClusterFeatureGcpExistingVpc)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ClusterFeatureValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ClusterFeatureAwsExistingVpc != nil {
		return obj.ClusterFeatureAwsExistingVpc
	}

	if obj.ClusterFeatureGcpExistingVpc != nil {
		return obj.ClusterFeatureGcpExistingVpc
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableClusterFeatureValue struct {
	value *ClusterFeatureValue
	isSet bool
}

func (v NullableClusterFeatureValue) Get() *ClusterFeatureValue {
	return v.value
}

func (v *NullableClusterFeatureValue) Set(val *ClusterFeatureValue) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterFeatureValue) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterFeatureValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterFeatureValue(val *ClusterFeatureValue) *NullableClusterFeatureValue {
	return &NullableClusterFeatureValue{value: val, isSet: true}
}

func (v NullableClusterFeatureValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterFeatureValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
