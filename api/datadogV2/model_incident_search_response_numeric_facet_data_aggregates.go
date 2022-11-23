// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// IncidentSearchResponseNumericFacetDataAggregates Aggregate information for numeric incident data.
type IncidentSearchResponseNumericFacetDataAggregates struct {
	// Maximum value of the numeric aggregates.
	Max *int32 `json:"max,omitempty"`
	// Minimum value of the numeric aggregates.
	Min *int32 `json:"min,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewIncidentSearchResponseNumericFacetDataAggregates instantiates a new IncidentSearchResponseNumericFacetDataAggregates object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentSearchResponseNumericFacetDataAggregates() *IncidentSearchResponseNumericFacetDataAggregates {
	this := IncidentSearchResponseNumericFacetDataAggregates{}
	return &this
}

// NewIncidentSearchResponseNumericFacetDataAggregatesWithDefaults instantiates a new IncidentSearchResponseNumericFacetDataAggregates object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentSearchResponseNumericFacetDataAggregatesWithDefaults() *IncidentSearchResponseNumericFacetDataAggregates {
	this := IncidentSearchResponseNumericFacetDataAggregates{}
	return &this
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *IncidentSearchResponseNumericFacetDataAggregates) GetMax() int32 {
	if o == nil || o.Max == nil {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseNumericFacetDataAggregates) GetMaxOk() (*int32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *IncidentSearchResponseNumericFacetDataAggregates) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *IncidentSearchResponseNumericFacetDataAggregates) SetMax(v int32) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *IncidentSearchResponseNumericFacetDataAggregates) GetMin() int32 {
	if o == nil || o.Min == nil {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseNumericFacetDataAggregates) GetMinOk() (*int32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *IncidentSearchResponseNumericFacetDataAggregates) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *IncidentSearchResponseNumericFacetDataAggregates) SetMin(v int32) {
	o.Min = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentSearchResponseNumericFacetDataAggregates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentSearchResponseNumericFacetDataAggregates) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Max *int32 `json:"max,omitempty"`
		Min *int32 `json:"min,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Max = all.Max
	o.Min = all.Min
	return nil
}
