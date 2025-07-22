# EventStorageWriteProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Content** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewEventStorageWriteProperties

`func NewEventStorageWriteProperties(key string, ) *EventStorageWriteProperties`

NewEventStorageWriteProperties instantiates a new EventStorageWriteProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventStorageWritePropertiesWithDefaults

`func NewEventStorageWritePropertiesWithDefaults() *EventStorageWriteProperties`

NewEventStorageWritePropertiesWithDefaults instantiates a new EventStorageWriteProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *EventStorageWriteProperties) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EventStorageWriteProperties) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EventStorageWriteProperties) SetKey(v string)`

SetKey sets Key field to given value.


### GetContent

`func (o *EventStorageWriteProperties) GetContent() interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EventStorageWriteProperties) GetContentOk() (*interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EventStorageWriteProperties) SetContent(v interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *EventStorageWriteProperties) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *EventStorageWriteProperties) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *EventStorageWriteProperties) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


