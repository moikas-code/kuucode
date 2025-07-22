# FilePartInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Mime** | **string** |  | 
**Filename** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**Source** | Pointer to [**FilePartSource**](FilePartSource.md) |  | [optional] 

## Methods

### NewFilePartInput

`func NewFilePartInput(type_ string, mime string, url string, ) *FilePartInput`

NewFilePartInput instantiates a new FilePartInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilePartInputWithDefaults

`func NewFilePartInputWithDefaults() *FilePartInput`

NewFilePartInputWithDefaults instantiates a new FilePartInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilePartInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilePartInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilePartInput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FilePartInput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *FilePartInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilePartInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilePartInput) SetType(v string)`

SetType sets Type field to given value.


### GetMime

`func (o *FilePartInput) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *FilePartInput) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *FilePartInput) SetMime(v string)`

SetMime sets Mime field to given value.


### GetFilename

`func (o *FilePartInput) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FilePartInput) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FilePartInput) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *FilePartInput) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetUrl

`func (o *FilePartInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FilePartInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FilePartInput) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSource

`func (o *FilePartInput) GetSource() FilePartSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FilePartInput) GetSourceOk() (*FilePartSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FilePartInput) SetSource(v FilePartSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *FilePartInput) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


