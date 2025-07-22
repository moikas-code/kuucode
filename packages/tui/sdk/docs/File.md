# File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** |  | 
**Added** | **int32** |  | 
**Removed** | **int32** |  | 
**Status** | **string** |  | 

## Methods

### NewFile

`func NewFile(path string, added int32, removed int32, status string, ) *File`

NewFile instantiates a new File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileWithDefaults

`func NewFileWithDefaults() *File`

NewFileWithDefaults instantiates a new File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *File) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *File) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *File) SetPath(v string)`

SetPath sets Path field to given value.


### GetAdded

`func (o *File) GetAdded() int32`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *File) GetAddedOk() (*int32, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *File) SetAdded(v int32)`

SetAdded sets Added field to given value.


### GetRemoved

`func (o *File) GetRemoved() int32`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *File) GetRemovedOk() (*int32, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *File) SetRemoved(v int32)`

SetRemoved sets Removed field to given value.


### GetStatus

`func (o *File) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *File) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *File) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


