# EventPermissionUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Properties** | [**PermissionInfo**](PermissionInfo.md) |  | 

## Methods

### NewEventPermissionUpdated

`func NewEventPermissionUpdated(type_ string, properties PermissionInfo, ) *EventPermissionUpdated`

NewEventPermissionUpdated instantiates a new EventPermissionUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPermissionUpdatedWithDefaults

`func NewEventPermissionUpdatedWithDefaults() *EventPermissionUpdated`

NewEventPermissionUpdatedWithDefaults instantiates a new EventPermissionUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventPermissionUpdated) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventPermissionUpdated) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventPermissionUpdated) SetType(v string)`

SetType sets Type field to given value.


### GetProperties

`func (o *EventPermissionUpdated) GetProperties() PermissionInfo`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventPermissionUpdated) GetPropertiesOk() (*PermissionInfo, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventPermissionUpdated) SetProperties(v PermissionInfo)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


