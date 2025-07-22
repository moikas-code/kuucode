# GetSessionByIdMessage200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | [**Message**](Message.md) |  | 
**Parts** | [**[]Part**](Part.md) |  | 

## Methods

### NewGetSessionByIdMessage200ResponseInner

`func NewGetSessionByIdMessage200ResponseInner(info Message, parts []Part, ) *GetSessionByIdMessage200ResponseInner`

NewGetSessionByIdMessage200ResponseInner instantiates a new GetSessionByIdMessage200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSessionByIdMessage200ResponseInnerWithDefaults

`func NewGetSessionByIdMessage200ResponseInnerWithDefaults() *GetSessionByIdMessage200ResponseInner`

NewGetSessionByIdMessage200ResponseInnerWithDefaults instantiates a new GetSessionByIdMessage200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *GetSessionByIdMessage200ResponseInner) GetInfo() Message`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GetSessionByIdMessage200ResponseInner) GetInfoOk() (*Message, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GetSessionByIdMessage200ResponseInner) SetInfo(v Message)`

SetInfo sets Info field to given value.


### GetParts

`func (o *GetSessionByIdMessage200ResponseInner) GetParts() []Part`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *GetSessionByIdMessage200ResponseInner) GetPartsOk() (*[]Part, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *GetSessionByIdMessage200ResponseInner) SetParts(v []Part)`

SetParts sets Parts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


