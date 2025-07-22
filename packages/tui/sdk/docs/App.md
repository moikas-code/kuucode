# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** |  | 
**Git** | **bool** |  | 
**Path** | [**AppPath**](AppPath.md) |  | 
**Time** | [**AppTime**](AppTime.md) |  | 

## Methods

### NewApp

`func NewApp(hostname string, git bool, path AppPath, time AppTime, ) *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *App) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *App) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *App) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetGit

`func (o *App) GetGit() bool`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *App) GetGitOk() (*bool, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *App) SetGit(v bool)`

SetGit sets Git field to given value.


### GetPath

`func (o *App) GetPath() AppPath`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *App) GetPathOk() (*AppPath, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *App) SetPath(v AppPath)`

SetPath sets Path field to given value.


### GetTime

`func (o *App) GetTime() AppTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *App) GetTimeOk() (*AppTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *App) SetTime(v AppTime)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


