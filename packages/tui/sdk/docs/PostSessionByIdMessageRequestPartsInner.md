# PostSessionByIdMessageRequestPartsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Text** | **string** |  | 
**Synthetic** | Pointer to **bool** |  | [optional] 
**Time** | Pointer to [**TextPartTime**](TextPartTime.md) |  | [optional] 
**Mime** | **string** |  | 
**Filename** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**Source** | Pointer to [**FilePartSource**](FilePartSource.md) |  | [optional] 

## Methods

### NewPostSessionByIdMessageRequestPartsInner

`func NewPostSessionByIdMessageRequestPartsInner(type_ string, text string, mime string, url string, ) *PostSessionByIdMessageRequestPartsInner`

NewPostSessionByIdMessageRequestPartsInner instantiates a new PostSessionByIdMessageRequestPartsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSessionByIdMessageRequestPartsInnerWithDefaults

`func NewPostSessionByIdMessageRequestPartsInnerWithDefaults() *PostSessionByIdMessageRequestPartsInner`

NewPostSessionByIdMessageRequestPartsInnerWithDefaults instantiates a new PostSessionByIdMessageRequestPartsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostSessionByIdMessageRequestPartsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostSessionByIdMessageRequestPartsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostSessionByIdMessageRequestPartsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostSessionByIdMessageRequestPartsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostSessionByIdMessageRequestPartsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostSessionByIdMessageRequestPartsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostSessionByIdMessageRequestPartsInner) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *PostSessionByIdMessageRequestPartsInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostSessionByIdMessageRequestPartsInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostSessionByIdMessageRequestPartsInner) SetText(v string)`

SetText sets Text field to given value.


### GetSynthetic

`func (o *PostSessionByIdMessageRequestPartsInner) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *PostSessionByIdMessageRequestPartsInner) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *PostSessionByIdMessageRequestPartsInner) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *PostSessionByIdMessageRequestPartsInner) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTime

`func (o *PostSessionByIdMessageRequestPartsInner) GetTime() TextPartTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PostSessionByIdMessageRequestPartsInner) GetTimeOk() (*TextPartTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PostSessionByIdMessageRequestPartsInner) SetTime(v TextPartTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *PostSessionByIdMessageRequestPartsInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetMime

`func (o *PostSessionByIdMessageRequestPartsInner) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *PostSessionByIdMessageRequestPartsInner) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *PostSessionByIdMessageRequestPartsInner) SetMime(v string)`

SetMime sets Mime field to given value.


### GetFilename

`func (o *PostSessionByIdMessageRequestPartsInner) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PostSessionByIdMessageRequestPartsInner) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PostSessionByIdMessageRequestPartsInner) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *PostSessionByIdMessageRequestPartsInner) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetUrl

`func (o *PostSessionByIdMessageRequestPartsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PostSessionByIdMessageRequestPartsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PostSessionByIdMessageRequestPartsInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSource

`func (o *PostSessionByIdMessageRequestPartsInner) GetSource() FilePartSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PostSessionByIdMessageRequestPartsInner) GetSourceOk() (*FilePartSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PostSessionByIdMessageRequestPartsInner) SetSource(v FilePartSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *PostSessionByIdMessageRequestPartsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


