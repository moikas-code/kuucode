# FileSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | [**FilePartSourceText**](FilePartSourceText.md) |  | 
**Type** | **string** |  | 
**Path** | **string** |  | 

## Methods

### NewFileSource

`func NewFileSource(text FilePartSourceText, type_ string, path string, ) *FileSource`

NewFileSource instantiates a new FileSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSourceWithDefaults

`func NewFileSourceWithDefaults() *FileSource`

NewFileSourceWithDefaults instantiates a new FileSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *FileSource) GetText() FilePartSourceText`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *FileSource) GetTextOk() (*FilePartSourceText, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *FileSource) SetText(v FilePartSourceText)`

SetText sets Text field to given value.


### GetType

`func (o *FileSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileSource) SetType(v string)`

SetType sets Type field to given value.


### GetPath

`func (o *FileSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileSource) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


