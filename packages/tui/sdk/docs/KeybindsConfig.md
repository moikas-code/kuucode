# KeybindsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leader** | **string** | Leader key for keybind combinations | [default to "ctrl+x"]
**AppHelp** | **string** | Show help dialog | [default to "<leader>h"]
**SwitchMode** | **string** | Next mode | [default to "tab"]
**SwitchModeReverse** | **string** | Previous Mode | [default to "shift+tab"]
**EditorOpen** | **string** | Open external editor | [default to "<leader>e"]
**SessionExport** | **string** | Export session to editor | [default to "<leader>x"]
**SessionNew** | **string** | Create a new session | [default to "<leader>n"]
**SessionList** | **string** | List all sessions | [default to "<leader>l"]
**SessionShare** | **string** | Share current session | [default to "<leader>s"]
**SessionUnshare** | **string** | Unshare current session | [default to "<leader>u"]
**SessionInterrupt** | **string** | Interrupt current session | [default to "esc"]
**SessionCompact** | **string** | Compact the session | [default to "<leader>c"]
**ToolDetails** | **string** | Toggle tool details | [default to "<leader>d"]
**ModelList** | **string** | List available models | [default to "<leader>m"]
**ThemeList** | **string** | List available themes | [default to "<leader>t"]
**FileList** | **string** | List files | [default to "<leader>f"]
**FileClose** | **string** | Close file | [default to "esc"]
**FileSearch** | **string** | Search file | [default to "<leader>/"]
**FileDiffToggle** | **string** | Split/unified diff | [default to "<leader>v"]
**ProjectInit** | **string** | Create/update .agentrc | [default to "<leader>i"]
**InputClear** | **string** | Clear input field | [default to "ctrl+c"]
**InputPaste** | **string** | Paste from clipboard | [default to "ctrl+v"]
**InputSubmit** | **string** | Submit input | [default to "enter"]
**InputNewline** | **string** | Insert newline in input | [default to "shift+enter,ctrl+j"]
**MessagesPageUp** | **string** | Scroll messages up by one page | [default to "pgup"]
**MessagesPageDown** | **string** | Scroll messages down by one page | [default to "pgdown"]
**MessagesHalfPageUp** | **string** | Scroll messages up by half page | [default to "ctrl+alt+u"]
**MessagesHalfPageDown** | **string** | Scroll messages down by half page | [default to "ctrl+alt+d"]
**MessagesPrevious** | **string** | Navigate to previous message | [default to "ctrl+up"]
**MessagesNext** | **string** | Navigate to next message | [default to "ctrl+down"]
**MessagesFirst** | **string** | Navigate to first message | [default to "ctrl+g"]
**MessagesLast** | **string** | Navigate to last message | [default to "ctrl+alt+g"]
**MessagesLayoutToggle** | **string** | Toggle layout | [default to "<leader>p"]
**MessagesCopy** | **string** | Copy message | [default to "<leader>y"]
**MessagesRevert** | **string** | Revert message | [default to "<leader>r"]
**AppExit** | **string** | Exit the application | [default to "ctrl+c,<leader>q"]

## Methods

### NewKeybindsConfig

`func NewKeybindsConfig(leader string, appHelp string, switchMode string, switchModeReverse string, editorOpen string, sessionExport string, sessionNew string, sessionList string, sessionShare string, sessionUnshare string, sessionInterrupt string, sessionCompact string, toolDetails string, modelList string, themeList string, fileList string, fileClose string, fileSearch string, fileDiffToggle string, projectInit string, inputClear string, inputPaste string, inputSubmit string, inputNewline string, messagesPageUp string, messagesPageDown string, messagesHalfPageUp string, messagesHalfPageDown string, messagesPrevious string, messagesNext string, messagesFirst string, messagesLast string, messagesLayoutToggle string, messagesCopy string, messagesRevert string, appExit string, ) *KeybindsConfig`

NewKeybindsConfig instantiates a new KeybindsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeybindsConfigWithDefaults

`func NewKeybindsConfigWithDefaults() *KeybindsConfig`

NewKeybindsConfigWithDefaults instantiates a new KeybindsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeader

`func (o *KeybindsConfig) GetLeader() string`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *KeybindsConfig) GetLeaderOk() (*string, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *KeybindsConfig) SetLeader(v string)`

SetLeader sets Leader field to given value.


### GetAppHelp

`func (o *KeybindsConfig) GetAppHelp() string`

GetAppHelp returns the AppHelp field if non-nil, zero value otherwise.

### GetAppHelpOk

`func (o *KeybindsConfig) GetAppHelpOk() (*string, bool)`

GetAppHelpOk returns a tuple with the AppHelp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppHelp

`func (o *KeybindsConfig) SetAppHelp(v string)`

SetAppHelp sets AppHelp field to given value.


### GetSwitchMode

`func (o *KeybindsConfig) GetSwitchMode() string`

GetSwitchMode returns the SwitchMode field if non-nil, zero value otherwise.

### GetSwitchModeOk

`func (o *KeybindsConfig) GetSwitchModeOk() (*string, bool)`

GetSwitchModeOk returns a tuple with the SwitchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMode

`func (o *KeybindsConfig) SetSwitchMode(v string)`

SetSwitchMode sets SwitchMode field to given value.


### GetSwitchModeReverse

`func (o *KeybindsConfig) GetSwitchModeReverse() string`

GetSwitchModeReverse returns the SwitchModeReverse field if non-nil, zero value otherwise.

### GetSwitchModeReverseOk

`func (o *KeybindsConfig) GetSwitchModeReverseOk() (*string, bool)`

GetSwitchModeReverseOk returns a tuple with the SwitchModeReverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchModeReverse

`func (o *KeybindsConfig) SetSwitchModeReverse(v string)`

SetSwitchModeReverse sets SwitchModeReverse field to given value.


### GetEditorOpen

`func (o *KeybindsConfig) GetEditorOpen() string`

GetEditorOpen returns the EditorOpen field if non-nil, zero value otherwise.

### GetEditorOpenOk

`func (o *KeybindsConfig) GetEditorOpenOk() (*string, bool)`

GetEditorOpenOk returns a tuple with the EditorOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorOpen

`func (o *KeybindsConfig) SetEditorOpen(v string)`

SetEditorOpen sets EditorOpen field to given value.


### GetSessionExport

`func (o *KeybindsConfig) GetSessionExport() string`

GetSessionExport returns the SessionExport field if non-nil, zero value otherwise.

### GetSessionExportOk

`func (o *KeybindsConfig) GetSessionExportOk() (*string, bool)`

GetSessionExportOk returns a tuple with the SessionExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionExport

`func (o *KeybindsConfig) SetSessionExport(v string)`

SetSessionExport sets SessionExport field to given value.


### GetSessionNew

`func (o *KeybindsConfig) GetSessionNew() string`

GetSessionNew returns the SessionNew field if non-nil, zero value otherwise.

### GetSessionNewOk

`func (o *KeybindsConfig) GetSessionNewOk() (*string, bool)`

GetSessionNewOk returns a tuple with the SessionNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionNew

`func (o *KeybindsConfig) SetSessionNew(v string)`

SetSessionNew sets SessionNew field to given value.


### GetSessionList

`func (o *KeybindsConfig) GetSessionList() string`

GetSessionList returns the SessionList field if non-nil, zero value otherwise.

### GetSessionListOk

`func (o *KeybindsConfig) GetSessionListOk() (*string, bool)`

GetSessionListOk returns a tuple with the SessionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionList

`func (o *KeybindsConfig) SetSessionList(v string)`

SetSessionList sets SessionList field to given value.


### GetSessionShare

`func (o *KeybindsConfig) GetSessionShare() string`

GetSessionShare returns the SessionShare field if non-nil, zero value otherwise.

### GetSessionShareOk

`func (o *KeybindsConfig) GetSessionShareOk() (*string, bool)`

GetSessionShareOk returns a tuple with the SessionShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionShare

`func (o *KeybindsConfig) SetSessionShare(v string)`

SetSessionShare sets SessionShare field to given value.


### GetSessionUnshare

`func (o *KeybindsConfig) GetSessionUnshare() string`

GetSessionUnshare returns the SessionUnshare field if non-nil, zero value otherwise.

### GetSessionUnshareOk

`func (o *KeybindsConfig) GetSessionUnshareOk() (*string, bool)`

GetSessionUnshareOk returns a tuple with the SessionUnshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionUnshare

`func (o *KeybindsConfig) SetSessionUnshare(v string)`

SetSessionUnshare sets SessionUnshare field to given value.


### GetSessionInterrupt

`func (o *KeybindsConfig) GetSessionInterrupt() string`

GetSessionInterrupt returns the SessionInterrupt field if non-nil, zero value otherwise.

### GetSessionInterruptOk

`func (o *KeybindsConfig) GetSessionInterruptOk() (*string, bool)`

GetSessionInterruptOk returns a tuple with the SessionInterrupt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInterrupt

`func (o *KeybindsConfig) SetSessionInterrupt(v string)`

SetSessionInterrupt sets SessionInterrupt field to given value.


### GetSessionCompact

`func (o *KeybindsConfig) GetSessionCompact() string`

GetSessionCompact returns the SessionCompact field if non-nil, zero value otherwise.

### GetSessionCompactOk

`func (o *KeybindsConfig) GetSessionCompactOk() (*string, bool)`

GetSessionCompactOk returns a tuple with the SessionCompact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCompact

`func (o *KeybindsConfig) SetSessionCompact(v string)`

SetSessionCompact sets SessionCompact field to given value.


### GetToolDetails

`func (o *KeybindsConfig) GetToolDetails() string`

GetToolDetails returns the ToolDetails field if non-nil, zero value otherwise.

### GetToolDetailsOk

`func (o *KeybindsConfig) GetToolDetailsOk() (*string, bool)`

GetToolDetailsOk returns a tuple with the ToolDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolDetails

`func (o *KeybindsConfig) SetToolDetails(v string)`

SetToolDetails sets ToolDetails field to given value.


### GetModelList

`func (o *KeybindsConfig) GetModelList() string`

GetModelList returns the ModelList field if non-nil, zero value otherwise.

### GetModelListOk

`func (o *KeybindsConfig) GetModelListOk() (*string, bool)`

GetModelListOk returns a tuple with the ModelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelList

`func (o *KeybindsConfig) SetModelList(v string)`

SetModelList sets ModelList field to given value.


### GetThemeList

`func (o *KeybindsConfig) GetThemeList() string`

GetThemeList returns the ThemeList field if non-nil, zero value otherwise.

### GetThemeListOk

`func (o *KeybindsConfig) GetThemeListOk() (*string, bool)`

GetThemeListOk returns a tuple with the ThemeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeList

`func (o *KeybindsConfig) SetThemeList(v string)`

SetThemeList sets ThemeList field to given value.


### GetFileList

`func (o *KeybindsConfig) GetFileList() string`

GetFileList returns the FileList field if non-nil, zero value otherwise.

### GetFileListOk

`func (o *KeybindsConfig) GetFileListOk() (*string, bool)`

GetFileListOk returns a tuple with the FileList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileList

`func (o *KeybindsConfig) SetFileList(v string)`

SetFileList sets FileList field to given value.


### GetFileClose

`func (o *KeybindsConfig) GetFileClose() string`

GetFileClose returns the FileClose field if non-nil, zero value otherwise.

### GetFileCloseOk

`func (o *KeybindsConfig) GetFileCloseOk() (*string, bool)`

GetFileCloseOk returns a tuple with the FileClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileClose

`func (o *KeybindsConfig) SetFileClose(v string)`

SetFileClose sets FileClose field to given value.


### GetFileSearch

`func (o *KeybindsConfig) GetFileSearch() string`

GetFileSearch returns the FileSearch field if non-nil, zero value otherwise.

### GetFileSearchOk

`func (o *KeybindsConfig) GetFileSearchOk() (*string, bool)`

GetFileSearchOk returns a tuple with the FileSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSearch

`func (o *KeybindsConfig) SetFileSearch(v string)`

SetFileSearch sets FileSearch field to given value.


### GetFileDiffToggle

`func (o *KeybindsConfig) GetFileDiffToggle() string`

GetFileDiffToggle returns the FileDiffToggle field if non-nil, zero value otherwise.

### GetFileDiffToggleOk

`func (o *KeybindsConfig) GetFileDiffToggleOk() (*string, bool)`

GetFileDiffToggleOk returns a tuple with the FileDiffToggle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDiffToggle

`func (o *KeybindsConfig) SetFileDiffToggle(v string)`

SetFileDiffToggle sets FileDiffToggle field to given value.


### GetProjectInit

`func (o *KeybindsConfig) GetProjectInit() string`

GetProjectInit returns the ProjectInit field if non-nil, zero value otherwise.

### GetProjectInitOk

`func (o *KeybindsConfig) GetProjectInitOk() (*string, bool)`

GetProjectInitOk returns a tuple with the ProjectInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectInit

`func (o *KeybindsConfig) SetProjectInit(v string)`

SetProjectInit sets ProjectInit field to given value.


### GetInputClear

`func (o *KeybindsConfig) GetInputClear() string`

GetInputClear returns the InputClear field if non-nil, zero value otherwise.

### GetInputClearOk

`func (o *KeybindsConfig) GetInputClearOk() (*string, bool)`

GetInputClearOk returns a tuple with the InputClear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputClear

`func (o *KeybindsConfig) SetInputClear(v string)`

SetInputClear sets InputClear field to given value.


### GetInputPaste

`func (o *KeybindsConfig) GetInputPaste() string`

GetInputPaste returns the InputPaste field if non-nil, zero value otherwise.

### GetInputPasteOk

`func (o *KeybindsConfig) GetInputPasteOk() (*string, bool)`

GetInputPasteOk returns a tuple with the InputPaste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPaste

`func (o *KeybindsConfig) SetInputPaste(v string)`

SetInputPaste sets InputPaste field to given value.


### GetInputSubmit

`func (o *KeybindsConfig) GetInputSubmit() string`

GetInputSubmit returns the InputSubmit field if non-nil, zero value otherwise.

### GetInputSubmitOk

`func (o *KeybindsConfig) GetInputSubmitOk() (*string, bool)`

GetInputSubmitOk returns a tuple with the InputSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSubmit

`func (o *KeybindsConfig) SetInputSubmit(v string)`

SetInputSubmit sets InputSubmit field to given value.


### GetInputNewline

`func (o *KeybindsConfig) GetInputNewline() string`

GetInputNewline returns the InputNewline field if non-nil, zero value otherwise.

### GetInputNewlineOk

`func (o *KeybindsConfig) GetInputNewlineOk() (*string, bool)`

GetInputNewlineOk returns a tuple with the InputNewline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputNewline

`func (o *KeybindsConfig) SetInputNewline(v string)`

SetInputNewline sets InputNewline field to given value.


### GetMessagesPageUp

`func (o *KeybindsConfig) GetMessagesPageUp() string`

GetMessagesPageUp returns the MessagesPageUp field if non-nil, zero value otherwise.

### GetMessagesPageUpOk

`func (o *KeybindsConfig) GetMessagesPageUpOk() (*string, bool)`

GetMessagesPageUpOk returns a tuple with the MessagesPageUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesPageUp

`func (o *KeybindsConfig) SetMessagesPageUp(v string)`

SetMessagesPageUp sets MessagesPageUp field to given value.


### GetMessagesPageDown

`func (o *KeybindsConfig) GetMessagesPageDown() string`

GetMessagesPageDown returns the MessagesPageDown field if non-nil, zero value otherwise.

### GetMessagesPageDownOk

`func (o *KeybindsConfig) GetMessagesPageDownOk() (*string, bool)`

GetMessagesPageDownOk returns a tuple with the MessagesPageDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesPageDown

`func (o *KeybindsConfig) SetMessagesPageDown(v string)`

SetMessagesPageDown sets MessagesPageDown field to given value.


### GetMessagesHalfPageUp

`func (o *KeybindsConfig) GetMessagesHalfPageUp() string`

GetMessagesHalfPageUp returns the MessagesHalfPageUp field if non-nil, zero value otherwise.

### GetMessagesHalfPageUpOk

`func (o *KeybindsConfig) GetMessagesHalfPageUpOk() (*string, bool)`

GetMessagesHalfPageUpOk returns a tuple with the MessagesHalfPageUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesHalfPageUp

`func (o *KeybindsConfig) SetMessagesHalfPageUp(v string)`

SetMessagesHalfPageUp sets MessagesHalfPageUp field to given value.


### GetMessagesHalfPageDown

`func (o *KeybindsConfig) GetMessagesHalfPageDown() string`

GetMessagesHalfPageDown returns the MessagesHalfPageDown field if non-nil, zero value otherwise.

### GetMessagesHalfPageDownOk

`func (o *KeybindsConfig) GetMessagesHalfPageDownOk() (*string, bool)`

GetMessagesHalfPageDownOk returns a tuple with the MessagesHalfPageDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesHalfPageDown

`func (o *KeybindsConfig) SetMessagesHalfPageDown(v string)`

SetMessagesHalfPageDown sets MessagesHalfPageDown field to given value.


### GetMessagesPrevious

`func (o *KeybindsConfig) GetMessagesPrevious() string`

GetMessagesPrevious returns the MessagesPrevious field if non-nil, zero value otherwise.

### GetMessagesPreviousOk

`func (o *KeybindsConfig) GetMessagesPreviousOk() (*string, bool)`

GetMessagesPreviousOk returns a tuple with the MessagesPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesPrevious

`func (o *KeybindsConfig) SetMessagesPrevious(v string)`

SetMessagesPrevious sets MessagesPrevious field to given value.


### GetMessagesNext

`func (o *KeybindsConfig) GetMessagesNext() string`

GetMessagesNext returns the MessagesNext field if non-nil, zero value otherwise.

### GetMessagesNextOk

`func (o *KeybindsConfig) GetMessagesNextOk() (*string, bool)`

GetMessagesNextOk returns a tuple with the MessagesNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesNext

`func (o *KeybindsConfig) SetMessagesNext(v string)`

SetMessagesNext sets MessagesNext field to given value.


### GetMessagesFirst

`func (o *KeybindsConfig) GetMessagesFirst() string`

GetMessagesFirst returns the MessagesFirst field if non-nil, zero value otherwise.

### GetMessagesFirstOk

`func (o *KeybindsConfig) GetMessagesFirstOk() (*string, bool)`

GetMessagesFirstOk returns a tuple with the MessagesFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesFirst

`func (o *KeybindsConfig) SetMessagesFirst(v string)`

SetMessagesFirst sets MessagesFirst field to given value.


### GetMessagesLast

`func (o *KeybindsConfig) GetMessagesLast() string`

GetMessagesLast returns the MessagesLast field if non-nil, zero value otherwise.

### GetMessagesLastOk

`func (o *KeybindsConfig) GetMessagesLastOk() (*string, bool)`

GetMessagesLastOk returns a tuple with the MessagesLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesLast

`func (o *KeybindsConfig) SetMessagesLast(v string)`

SetMessagesLast sets MessagesLast field to given value.


### GetMessagesLayoutToggle

`func (o *KeybindsConfig) GetMessagesLayoutToggle() string`

GetMessagesLayoutToggle returns the MessagesLayoutToggle field if non-nil, zero value otherwise.

### GetMessagesLayoutToggleOk

`func (o *KeybindsConfig) GetMessagesLayoutToggleOk() (*string, bool)`

GetMessagesLayoutToggleOk returns a tuple with the MessagesLayoutToggle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesLayoutToggle

`func (o *KeybindsConfig) SetMessagesLayoutToggle(v string)`

SetMessagesLayoutToggle sets MessagesLayoutToggle field to given value.


### GetMessagesCopy

`func (o *KeybindsConfig) GetMessagesCopy() string`

GetMessagesCopy returns the MessagesCopy field if non-nil, zero value otherwise.

### GetMessagesCopyOk

`func (o *KeybindsConfig) GetMessagesCopyOk() (*string, bool)`

GetMessagesCopyOk returns a tuple with the MessagesCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesCopy

`func (o *KeybindsConfig) SetMessagesCopy(v string)`

SetMessagesCopy sets MessagesCopy field to given value.


### GetMessagesRevert

`func (o *KeybindsConfig) GetMessagesRevert() string`

GetMessagesRevert returns the MessagesRevert field if non-nil, zero value otherwise.

### GetMessagesRevertOk

`func (o *KeybindsConfig) GetMessagesRevertOk() (*string, bool)`

GetMessagesRevertOk returns a tuple with the MessagesRevert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesRevert

`func (o *KeybindsConfig) SetMessagesRevert(v string)`

SetMessagesRevert sets MessagesRevert field to given value.


### GetAppExit

`func (o *KeybindsConfig) GetAppExit() string`

GetAppExit returns the AppExit field if non-nil, zero value otherwise.

### GetAppExitOk

`func (o *KeybindsConfig) GetAppExitOk() (*string, bool)`

GetAppExitOk returns a tuple with the AppExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppExit

`func (o *KeybindsConfig) SetAppExit(v string)`

SetAppExit sets AppExit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


