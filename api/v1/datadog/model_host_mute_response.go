/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// HostMuteResponse struct for HostMuteResponse
type HostMuteResponse struct {
	Action *string `json:"action,omitempty"`
	// POSIX timestamp in seconds when the host is unmuted.
	End      *int64  `json:"end,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Message  *string `json:"message,omitempty"`
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *HostMuteResponse) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HostMuteResponse) GetActionOk() (string, bool) {
	if o == nil || o.Action == nil {
		var ret string
		return ret, false
	}
	return *o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *HostMuteResponse) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *HostMuteResponse) SetAction(v string) {
	o.Action = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *HostMuteResponse) GetEnd() int64 {
	if o == nil || o.End == nil {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HostMuteResponse) GetEndOk() (int64, bool) {
	if o == nil || o.End == nil {
		var ret int64
		return ret, false
	}
	return *o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *HostMuteResponse) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *HostMuteResponse) SetEnd(v int64) {
	o.End = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *HostMuteResponse) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HostMuteResponse) GetHostnameOk() (string, bool) {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret, false
	}
	return *o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *HostMuteResponse) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *HostMuteResponse) SetHostname(v string) {
	o.Hostname = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HostMuteResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HostMuteResponse) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HostMuteResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HostMuteResponse) SetMessage(v string) {
	o.Message = &v
}

type NullableHostMuteResponse struct {
	Value        HostMuteResponse
	ExplicitNull bool
}

func (v NullableHostMuteResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableHostMuteResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
