/*
CouchDB API

*Note* This is not a definitive implementation of the CouchDB API, it's missing a lot of the endpoints for server/database managment and everything for attachments all COPY operations and probably a few other parts.  It also targets golang, as such the '#/definitions/Document' is intentionally left empty to generate a go `interface`, which you can then cast to a `map[string]interface{}`. 

API version: 4.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchdb_go

import (
	"encoding/json"
)

// checks if the Query type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Query{}

// Query struct for Query
type Query struct {
	// Maximum number of results returned. Default is 25. Optional
	Limit *int32 `json:"limit,omitempty"`
	// Skip the first ‘n’ results, where ‘n’ is the value specified. Optional
	Skip *int32 `json:"skip,omitempty"`
	// Name of the index. If no name is provided, a name will be generated automatically. Optional
	Name *string `json:"name,omitempty"`
	Sort map[string]interface{} `json:"sort,omitempty"`
	// Instruct a query to use a specific index. Specified either as \"<design_document>\" or [\"<design_document>\", \"<index_name>\"]. Optional
	Field []string `json:"field,omitempty"`
	// JSON array specifying which fields of each object should be returned. If it is omitted, the entire object is returned. More information provided in the section on filtering fields. Optional
	UseIndex []string `json:"use_index,omitempty"`
	// Include conflicted documents if true. Intended use is to easily find conflicted documents, without an index or view. Default is false. Optional
	Conflitsc *bool `json:"conflitsc,omitempty"`
	// A string that enables you to specify which page of results you require. Used for paging through result sets. Every query returns an opaque string under the bookmark key that can then be passed back in a query to get the next page of results. If any part of the selector query changes between requests, the results are undefined. Optional, default: null 
	Bookmark *string `json:"bookmark,omitempty"`
	// Whether to update the index prior to returning the result. Default is true. Optional
	Update *bool `json:"update,omitempty"`
	// Whether or not the view results should be returned from a “stable” set of shards. Optional
	Stable *bool `json:"stable,omitempty"`
	// Combination of update=false and stable=true options. Possible options: \"ok\", false (default). Optional Note that this parameter is deprecated. Use stable and update instead. See Views Generation for more details. 
	Stale *string `json:"stale,omitempty"`
	// Include execution statistics in the query response. Optional, default: false 
	ExecutionStats *bool `json:"execution_stats,omitempty"`
	Selector map[string]interface{} `json:"selector,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Query Query

// NewQuery instantiates a new Query object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuery() *Query {
	this := Query{}
	return &this
}

// NewQueryWithDefaults instantiates a new Query object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryWithDefaults() *Query {
	this := Query{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *Query) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *Query) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *Query) SetLimit(v int32) {
	o.Limit = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *Query) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *Query) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *Query) SetSkip(v int32) {
	o.Skip = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Query) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Query) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Query) SetName(v string) {
	o.Name = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *Query) GetSort() map[string]interface{} {
	if o == nil || IsNil(o.Sort) {
		var ret map[string]interface{}
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetSortOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Sort) {
		return map[string]interface{}{}, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *Query) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given map[string]interface{} and assigns it to the Sort field.
func (o *Query) SetSort(v map[string]interface{}) {
	o.Sort = v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *Query) GetField() []string {
	if o == nil || IsNil(o.Field) {
		var ret []string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetFieldOk() ([]string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *Query) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given []string and assigns it to the Field field.
func (o *Query) SetField(v []string) {
	o.Field = v
}

// GetUseIndex returns the UseIndex field value if set, zero value otherwise.
func (o *Query) GetUseIndex() []string {
	if o == nil || IsNil(o.UseIndex) {
		var ret []string
		return ret
	}
	return o.UseIndex
}

// GetUseIndexOk returns a tuple with the UseIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetUseIndexOk() ([]string, bool) {
	if o == nil || IsNil(o.UseIndex) {
		return nil, false
	}
	return o.UseIndex, true
}

// HasUseIndex returns a boolean if a field has been set.
func (o *Query) HasUseIndex() bool {
	if o != nil && !IsNil(o.UseIndex) {
		return true
	}

	return false
}

// SetUseIndex gets a reference to the given []string and assigns it to the UseIndex field.
func (o *Query) SetUseIndex(v []string) {
	o.UseIndex = v
}

// GetConflitsc returns the Conflitsc field value if set, zero value otherwise.
func (o *Query) GetConflitsc() bool {
	if o == nil || IsNil(o.Conflitsc) {
		var ret bool
		return ret
	}
	return *o.Conflitsc
}

// GetConflitscOk returns a tuple with the Conflitsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetConflitscOk() (*bool, bool) {
	if o == nil || IsNil(o.Conflitsc) {
		return nil, false
	}
	return o.Conflitsc, true
}

// HasConflitsc returns a boolean if a field has been set.
func (o *Query) HasConflitsc() bool {
	if o != nil && !IsNil(o.Conflitsc) {
		return true
	}

	return false
}

// SetConflitsc gets a reference to the given bool and assigns it to the Conflitsc field.
func (o *Query) SetConflitsc(v bool) {
	o.Conflitsc = &v
}

// GetBookmark returns the Bookmark field value if set, zero value otherwise.
func (o *Query) GetBookmark() string {
	if o == nil || IsNil(o.Bookmark) {
		var ret string
		return ret
	}
	return *o.Bookmark
}

// GetBookmarkOk returns a tuple with the Bookmark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetBookmarkOk() (*string, bool) {
	if o == nil || IsNil(o.Bookmark) {
		return nil, false
	}
	return o.Bookmark, true
}

// HasBookmark returns a boolean if a field has been set.
func (o *Query) HasBookmark() bool {
	if o != nil && !IsNil(o.Bookmark) {
		return true
	}

	return false
}

// SetBookmark gets a reference to the given string and assigns it to the Bookmark field.
func (o *Query) SetBookmark(v string) {
	o.Bookmark = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *Query) GetUpdate() bool {
	if o == nil || IsNil(o.Update) {
		var ret bool
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *Query) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given bool and assigns it to the Update field.
func (o *Query) SetUpdate(v bool) {
	o.Update = &v
}

// GetStable returns the Stable field value if set, zero value otherwise.
func (o *Query) GetStable() bool {
	if o == nil || IsNil(o.Stable) {
		var ret bool
		return ret
	}
	return *o.Stable
}

// GetStableOk returns a tuple with the Stable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetStableOk() (*bool, bool) {
	if o == nil || IsNil(o.Stable) {
		return nil, false
	}
	return o.Stable, true
}

// HasStable returns a boolean if a field has been set.
func (o *Query) HasStable() bool {
	if o != nil && !IsNil(o.Stable) {
		return true
	}

	return false
}

// SetStable gets a reference to the given bool and assigns it to the Stable field.
func (o *Query) SetStable(v bool) {
	o.Stable = &v
}

// GetStale returns the Stale field value if set, zero value otherwise.
func (o *Query) GetStale() string {
	if o == nil || IsNil(o.Stale) {
		var ret string
		return ret
	}
	return *o.Stale
}

// GetStaleOk returns a tuple with the Stale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetStaleOk() (*string, bool) {
	if o == nil || IsNil(o.Stale) {
		return nil, false
	}
	return o.Stale, true
}

// HasStale returns a boolean if a field has been set.
func (o *Query) HasStale() bool {
	if o != nil && !IsNil(o.Stale) {
		return true
	}

	return false
}

// SetStale gets a reference to the given string and assigns it to the Stale field.
func (o *Query) SetStale(v string) {
	o.Stale = &v
}

// GetExecutionStats returns the ExecutionStats field value if set, zero value otherwise.
func (o *Query) GetExecutionStats() bool {
	if o == nil || IsNil(o.ExecutionStats) {
		var ret bool
		return ret
	}
	return *o.ExecutionStats
}

// GetExecutionStatsOk returns a tuple with the ExecutionStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetExecutionStatsOk() (*bool, bool) {
	if o == nil || IsNil(o.ExecutionStats) {
		return nil, false
	}
	return o.ExecutionStats, true
}

// HasExecutionStats returns a boolean if a field has been set.
func (o *Query) HasExecutionStats() bool {
	if o != nil && !IsNil(o.ExecutionStats) {
		return true
	}

	return false
}

// SetExecutionStats gets a reference to the given bool and assigns it to the ExecutionStats field.
func (o *Query) SetExecutionStats(v bool) {
	o.ExecutionStats = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *Query) GetSelector() map[string]interface{} {
	if o == nil || IsNil(o.Selector) {
		var ret map[string]interface{}
		return ret
	}
	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetSelectorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Selector) {
		return map[string]interface{}{}, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *Query) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given map[string]interface{} and assigns it to the Selector field.
func (o *Query) SetSelector(v map[string]interface{}) {
	o.Selector = v
}

func (o Query) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Query) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.UseIndex) {
		toSerialize["use_index"] = o.UseIndex
	}
	if !IsNil(o.Conflitsc) {
		toSerialize["conflitsc"] = o.Conflitsc
	}
	if !IsNil(o.Bookmark) {
		toSerialize["bookmark"] = o.Bookmark
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Stable) {
		toSerialize["stable"] = o.Stable
	}
	if !IsNil(o.Stale) {
		toSerialize["stale"] = o.Stale
	}
	if !IsNil(o.ExecutionStats) {
		toSerialize["execution_stats"] = o.ExecutionStats
	}
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Query) UnmarshalJSON(bytes []byte) (err error) {
	varQuery := _Query{}

	if err = json.Unmarshal(bytes, &varQuery); err == nil {
		*o = Query(varQuery)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "limit")
		delete(additionalProperties, "skip")
		delete(additionalProperties, "name")
		delete(additionalProperties, "sort")
		delete(additionalProperties, "field")
		delete(additionalProperties, "use_index")
		delete(additionalProperties, "conflitsc")
		delete(additionalProperties, "bookmark")
		delete(additionalProperties, "update")
		delete(additionalProperties, "stable")
		delete(additionalProperties, "stale")
		delete(additionalProperties, "execution_stats")
		delete(additionalProperties, "selector")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuery struct {
	value *Query
	isSet bool
}

func (v NullableQuery) Get() *Query {
	return v.value
}

func (v *NullableQuery) Set(val *Query) {
	v.value = val
	v.isSet = true
}

func (v NullableQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuery(val *Query) *NullableQuery {
	return &NullableQuery{value: val, isSet: true}
}

func (v NullableQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


