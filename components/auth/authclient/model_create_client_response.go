/*
Auth API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: AUTH_VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authclient

import (
	"encoding/json"
)

// checks if the CreateClientResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateClientResponse{}

// CreateClientResponse struct for CreateClientResponse
type CreateClientResponse struct {
	Data interface{} `json:"data,omitempty"`
}

// NewCreateClientResponse instantiates a new CreateClientResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClientResponse() *CreateClientResponse {
	this := CreateClientResponse{}
	return &this
}

// NewCreateClientResponseWithDefaults instantiates a new CreateClientResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClientResponseWithDefaults() *CreateClientResponse {
	this := CreateClientResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClientResponse) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClientResponse) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateClientResponse) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *CreateClientResponse) SetData(v interface{}) {
	o.Data = v
}

func (o CreateClientResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateClientResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateClientResponse struct {
	value *CreateClientResponse
	isSet bool
}

func (v NullableCreateClientResponse) Get() *CreateClientResponse {
	return v.value
}

func (v *NullableCreateClientResponse) Set(val *CreateClientResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClientResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClientResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClientResponse(val *CreateClientResponse) *NullableCreateClientResponse {
	return &NullableCreateClientResponse{value: val, isSet: true}
}

func (v NullableCreateClientResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClientResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


