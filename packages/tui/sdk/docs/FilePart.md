# FilePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SessionID** | **string** |  | 
**MessageID** | **string** |  | 
**Type** | **string** |  | 
**Mime** | **string** |  | 
**Filename** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**Source** | Pointer to [**FilePartSource**](FilePartSource.md) |  | [optional] 

## Methods

### NewFilePart

`func NewFilePart(id string, sessionID string, messageID string, type_ string, mime string, url string, ) *FilePart`

NewFilePart instantiates a new FilePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilePartWithDefaults

`func NewFilePartWithDefaults() *FilePart`

NewFilePartWithDefaults instantiates a new FilePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilePart) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilePart) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilePart) SetId(v string)`

SetId sets Id field to given value.


### GetSessionID

`func (o *FilePart) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *FilePart) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *FilePart) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.


### GetMessageID

`func (o *FilePart) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *FilePart) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *FilePart) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.


### GetType

`func (o *FilePart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilePart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilePart) SetType(v string)`

SetType sets Type field to given value.


### GetMime

`func (o *FilePart) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *FilePart) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *FilePart) SetMime(v string)`

SetMime sets Mime field to given value.


### GetFilename

`func (o *FilePart) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FilePart) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FilePart) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *FilePart) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetUrl

`func (o *FilePart) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FilePart) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FilePart) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSource

`func (o *FilePart) GetSource() FilePartSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FilePart) GetSourceOk() (*FilePartSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FilePart) SetSource(v FilePartSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *FilePart) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


