package dcimgui

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
)

// #region Impl

func ImplOpenGL3Init(glslVersion string) {
	ver := cc.NewString(glslVersion)
	defer ver.Free()
	imGui_ImplOpenGL3_Init.Fn()(ver)
}
func ImplOpenGL3Shutdown()                       { imGui_ImplOpenGL3_Shutdown.Fn()() }
func ImplOpenGL3NewFrame()                       { imGui_ImplOpenGL3_NewFrame.Fn()() }
func ImplOpenGL3RenderDrawData(data *ImDrawData) { imGui_ImplOpenGL3_RenderDrawData.Fn()(data) }

// #endregion

// #region struct

type (
	ImDrawIdx              uint16
	ImGuiID                uint32
	ImS8                   int8
	ImU8                   uint8
	ImS16                  int16
	ImU16                  uint16
	ImS32                  int32
	ImU32                  uint32
	ImS64                  cc.Longlong
	ImU64                  cc.ULonglong
	ImGuiCol               = ImGuiCol_
	ImGuiCond              = ImGuiCond_
	ImGuiDataType          = ImGuiDataType_
	ImGuiMouseButton       = ImGuiMouseButton_
	ImGuiMouseCursor       = ImGuiMouseCursor_
	ImGuiStyleVar          = ImGuiStyleVar_
	ImGuiTableBgTarget     = ImGuiTableBgTarget_
	ImDrawFlags            = ImDrawFlags_
	ImDrawListFlags        = ImDrawListFlags_
	ImFontFlags            = ImFontFlags_
	ImFontAtlasFlags       = ImFontAtlasFlags_
	ImGuiBackendFlags      = ImGuiBackendFlags_
	ImGuiButtonFlags       = ImGuiButtonFlags_
	ImGuiChildFlags        = ImGuiChildFlags_
	ImGuiColorEditFlags    = ImGuiColorEditFlags_
	ImGuiConfigFlags       = ImGuiConfigFlags_
	ImGuiComboFlags        = ImGuiComboFlags_
	ImGuiDockNodeFlags     = ImGuiDockNodeFlags_
	ImGuiDragDropFlags     = ImGuiDragDropFlags_
	ImGuiFocusedFlags      = ImGuiFocusedFlags_
	ImGuiHoveredFlags      = ImGuiHoveredFlags_
	ImGuiInputFlags        = ImGuiInputFlags_
	ImGuiInputTextFlags    = ImGuiInputTextFlags_
	ImGuiItemFlags         = ImGuiItemFlags_
	ImGuiKeyChord          int32
	ImGuiPopupFlags        = ImGuiPopupFlags_
	ImGuiMultiSelectFlags  = ImGuiMultiSelectFlags_
	ImGuiSelectableFlags   = ImGuiSelectableFlags_
	ImGuiSliderFlags       = ImGuiSliderFlags_
	ImGuiTabBarFlags       = ImGuiTabBarFlags_
	ImGuiTabItemFlags      = ImGuiTabItemFlags_
	ImGuiTableFlags        = ImGuiTableFlags_
	ImGuiTableColumnFlags  = ImGuiTableColumnFlags_
	ImGuiTableRowFlags     = ImGuiTableRowFlags_
	ImGuiTreeNodeFlags     = ImGuiTreeNodeFlags_
	ImGuiViewportFlags     = ImGuiViewportFlags_
	ImGuiWindowFlags       = ImGuiWindowFlags_
	ImWchar32              uint32
	ImWchar16              uint16
	ImWchar                ImWchar32
	ImGuiSelectionUserData ImS64
	ImTextureID            ImU64
	ImFontAtlasRectId      int32
	ImFontAtlasCustomRect  ImFontAtlasRect
)
type ImVector[T any] struct {
	Size     int32
	Capacity int32
	Data     *T
}
type ImDrawListSharedData struct{}
type ImFontAtlasBuilder struct{}
type ImFontLoader struct{}
type ImGuiContext struct{}
type ImVec2 struct {
	X float32
	Y float32
}
type ImVec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}
type ImTextureRef struct {
	TexData *ImTextureData
	TexID   ImTextureID
}
type ImGuiTableSortSpecs struct {
	Specs      *ImGuiTableColumnSortSpecs
	SpecsCount int32
	SpecsDirty cc.Bool
}
type ImGuiTableColumnSortSpecs struct {
	ColumnUserID  ImGuiID
	ColumnIndex   ImS16
	SortOrder     ImS16
	SortDirection ImGuiSortDirection
}
type ImGuiStyle struct {
	FontSizeBase                     float32
	FontScaleMain                    float32
	FontScaleDpi                     float32
	Alpha                            float32
	DisabledAlpha                    float32
	WindowPadding                    ImVec2
	WindowRounding                   float32
	WindowBorderSize                 float32
	WindowBorderHoverPadding         float32
	WindowMinSize                    ImVec2
	WindowTitleAlign                 ImVec2
	WindowMenuButtonPosition         ImGuiDir
	ChildRounding                    float32
	ChildBorderSize                  float32
	PopupRounding                    float32
	PopupBorderSize                  float32
	FramePadding                     ImVec2
	FrameRounding                    float32
	FrameBorderSize                  float32
	ItemSpacing                      ImVec2
	ItemInnerSpacing                 ImVec2
	CellPadding                      ImVec2
	TouchExtraPadding                ImVec2
	IndentSpacing                    float32
	ColumnsMinSpacing                float32
	ScrollbarSize                    float32
	ScrollbarRounding                float32
	GrabMinSize                      float32
	GrabRounding                     float32
	LogSliderDeadzone                float32
	ImageBorderSize                  float32
	TabRounding                      float32
	TabBorderSize                    float32
	TabCloseButtonMinWidthSelected   float32
	TabCloseButtonMinWidthUnselected float32
	TabBarBorderSize                 float32
	TabBarOverlineSize               float32
	TableAngledHeadersAngle          float32
	TableAngledHeadersTextAlign      ImVec2
	TreeLinesFlags                   ImGuiTreeNodeFlags
	TreeLinesSize                    float32
	TreeLinesRounding                float32
	ColorButtonPosition              ImGuiDir
	ButtonTextAlign                  ImVec2
	SelectableTextAlign              ImVec2
	SeparatorTextBorderSize          float32
	SeparatorTextAlign               ImVec2
	SeparatorTextPadding             ImVec2
	DisplayWindowPadding             ImVec2
	DisplaySafeAreaPadding           ImVec2
	DockingSeparatorSize             float32
	MouseCursorScale                 float32
	AntiAliasedLines                 cc.Bool
	AntiAliasedLinesUseTex           cc.Bool
	AntiAliasedFill                  cc.Bool
	CurveTessellationTol             float32
	CircleTessellationMaxError       float32
	Colors                           [ImGuiCol_COUNT]ImVec4
	HoverStationaryDelay             float32
	HoverDelayShort                  float32
	HoverDelayNormal                 float32
	HoverFlagsForTooltipMouse        ImGuiHoveredFlags
	HoverFlagsForTooltipNav          ImGuiHoveredFlags
	_MainScale                       float32
	_NextFrameFontSizeBase           float32
}
type ImGuiKeyData struct {
	Down             cc.Bool
	DownDuration     float32
	DownDurationPrev float32
	AnalogValue      float32
}
type ImGuiIO struct {
	ConfigFlags                                   ImGuiConfigFlags
	BackendFlags                                  ImGuiBackendFlags
	DisplaySize                                   ImVec2
	DisplayFramebufferScale                       ImVec2
	DeltaTime                                     float32
	IniSavingRate                                 float32
	IniFilename                                   cc.String
	LogFilename                                   cc.String
	UserData                                      unsafe.Pointer
	Fonts                                         *ImFontAtlas
	FontDefault                                   *ImFont
	FontAllowUserScaling                          cc.Bool
	ConfigNavSwapGamepadButtons                   cc.Bool
	ConfigNavMoveSetMousePos                      cc.Bool
	ConfigNavCaptureKeyboard                      cc.Bool
	ConfigNavEscapeClearFocusItem                 cc.Bool
	ConfigNavEscapeClearFocusWindow               cc.Bool
	ConfigNavCursorVisibleAuto                    cc.Bool
	ConfigNavCursorVisibleAlways                  cc.Bool
	ConfigDockingNoSplit                          cc.Bool
	ConfigDockingWithShift                        cc.Bool
	ConfigDockingAlwaysTabBar                     cc.Bool
	ConfigDockingTransparentPayload               cc.Bool
	ConfigViewportsNoAutoMerge                    cc.Bool
	ConfigViewportsNoTaskBarIcon                  cc.Bool
	ConfigViewportsNoDecoration                   cc.Bool
	ConfigViewportsNoDefaultParent                cc.Bool
	ConfigDpiScaleFonts                           cc.Bool
	ConfigDpiScaleViewports                       cc.Bool
	MouseDrawCursor                               cc.Bool
	ConfigMacOSXBehaviors                         cc.Bool
	ConfigInputTrickleEventQueue                  cc.Bool
	ConfigInputTextCursorBlink                    cc.Bool
	ConfigInputTextEnterKeepActive                cc.Bool
	ConfigDragClickToInputText                    cc.Bool
	ConfigWindowsResizeFromEdges                  cc.Bool
	ConfigWindowsMoveFromTitleBarOnly             cc.Bool
	ConfigWindowsCopyContentsWithCtrlC            cc.Bool
	ConfigScrollbarScrollByPage                   cc.Bool
	ConfigMemoryCompactTimer                      float32
	MouseDoubleClickTime                          float32
	MouseDoubleClickMaxDist                       float32
	MouseDragThreshold                            float32
	KeyRepeatDelay                                float32
	KeyRepeatRate                                 float32
	ConfigErrorRecovery                           cc.Bool
	ConfigErrorRecoveryEnableAssert               cc.Bool
	ConfigErrorRecoveryEnableDebugLog             cc.Bool
	ConfigErrorRecoveryEnableTooltip              cc.Bool
	ConfigDebugIsDebuggerPresent                  cc.Bool
	ConfigDebugHighlightIdConflicts               cc.Bool
	ConfigDebugHighlightIdConflictsShowItemPicker cc.Bool
	ConfigDebugBeginReturnValueOnce               cc.Bool
	ConfigDebugBeginReturnValueLoop               cc.Bool
	ConfigDebugIgnoreFocusLoss                    cc.Bool
	ConfigDebugIniSettings                        cc.Bool
	BackendPlatformName                           cc.String
	BackendRendererName                           cc.String
	BackendPlatformUserData                       unsafe.Pointer
	BackendRendererUserData                       unsafe.Pointer
	BackendLanguageUserData                       unsafe.Pointer
	WantCaptureMouse                              cc.Bool
	WantCaptureKeyboard                           cc.Bool
	WantTextInput                                 cc.Bool
	WantSetMousePos                               cc.Bool
	WantSaveIniSettings                           cc.Bool
	NavActive                                     cc.Bool
	NavVisible                                    cc.Bool
	Framerate                                     float32
	MetricsRenderVertices                         int32
	MetricsRenderIndices                          int32
	MetricsRenderWindows                          int32
	MetricsActiveWindows                          int32
	MouseDelta                                    ImVec2
	Ctx                                           *ImGuiContext
	MousePos                                      ImVec2
	MouseDown                                     [5]cc.Bool
	MouseWheel                                    float32
	MouseWheelH                                   float32
	MouseSource                                   ImGuiMouseSource
	MouseHoveredViewport                          ImGuiID
	KeyCtrl                                       cc.Bool
	KeyShift                                      cc.Bool
	KeyAlt                                        cc.Bool
	KeySuper                                      cc.Bool
	KeyMods                                       ImGuiKeyChord
	KeysData                                      [ImGuiKey_NamedKey_COUNT]ImGuiKeyData
	WantCaptureMouseUnlessPopupClose              cc.Bool
	MousePosPrev                                  ImVec2
	MouseClickedPos                               [5]ImVec2
	MouseClickedTime                              [5]float64
	MouseClicked                                  [5]cc.Bool
	MouseDoubleClicked                            [5]cc.Bool
	MouseClickedCount                             [5]ImU16
	MouseClickedLastCount                         [5]ImU16
	MouseReleased                                 [5]cc.Bool
	MouseReleasedTime                             [5]float64
	MouseDownOwned                                [5]cc.Bool
	MouseDownOwnedUnlessPopupClose                [5]cc.Bool
	MouseWheelRequestAxisSwap                     cc.Bool
	MouseCtrlLeftAsRightClick                     cc.Bool
	MouseDownDuration                             [5]float32
	MouseDownDurationPrev                         [5]float32
	MouseDragMaxDistanceAbs                       [5]ImVec2
	MouseDragMaxDistanceSqr                       [5]float32
	PenPressure                                   float32
	AppFocusLost                                  cc.Bool
	AppAcceptingEvents                            cc.Bool
	InputQueueSurrogate                           ImWchar16
	InputQueueCharacters                          ImVector[ImWchar]
	FontGlobalScale                               float32
	GetClipboardTextFn                            cc.Func // const char* (*GetClipboardTextFn)(void* user_data)
	SetClipboardTextFn                            cc.Func // void (*SetClipboardTextFn)(void* user_data, const char* text)
	ClipboardUserData                             unsafe.Pointer
}
type ImGuiInputTextCallbackData struct {
	Ctx            *ImGuiContext
	EventFlag      ImGuiInputTextFlags
	Flags          ImGuiInputTextFlags
	UserData       unsafe.Pointer
	EventChar      ImWchar
	EventKey       ImGuiKey
	Buf            cc.String
	BufTextLen     int32
	BufSize        int32
	BufDirty       cc.Bool
	CursorPos      int32
	SelectionStart int32
	SelectionEnd   int32
}
type ImGuiSizeCallbackData struct {
	UserData    unsafe.Pointer
	Pos         ImVec2
	CurrentSize ImVec2
	DesiredSize ImVec2
}
type ImGuiWindowClass struct {
	ClassId                    ImGuiID
	ParentViewportId           ImGuiID
	FocusRouteParentWindowId   ImGuiID
	ViewportFlagsOverrideSet   ImGuiViewportFlags
	ViewportFlagsOverrideClear ImGuiViewportFlags
	TabItemFlagsOverrideSet    ImGuiTabItemFlags
	DockNodeFlagsOverrideSet   ImGuiDockNodeFlags
	DockingAlwaysTabBar        cc.Bool
	DockingAllowUnclassed      cc.Bool
}
type ImGuiPayload struct {
	Data           unsafe.Pointer
	DataSize       int32
	SourceId       ImGuiID
	SourceParentId ImGuiID
	DataFrameCount int32
	DataType       [32 + 1]byte
	Preview        cc.Bool
	Delivery       cc.Bool
}
type ImGuiTextRange struct {
	B cc.String
	E cc.String
}
type ImGuiTextFilter struct {
	InputBuf  [256]byte
	Filters   ImVector[ImGuiTextRange]
	CountGrep int32
}
type ImGuiTextBuffer struct {
	Buf ImVector[byte]
}
type ImGuiStoragePair struct {
	Key ImGuiID
	Val ImGuiStoragePairVal
}
type ImGuiStoragePairVal struct {
	// Val_i int32
	// Val_f float32
	// Val_p unsafe.Pointer
	Val *unsafe.Pointer
}
type ImGuiStorage struct {
	Data ImVector[ImGuiStoragePair]
}
type ImGuiListClipper struct {
	Ctx              *ImGuiContext
	DisplayStart     int32
	DisplayEnd       int32
	ItemsCount       int32
	ItemsHeight      float32
	StartPosY        float64
	StartSeekOffsetY float64
	TempData         unsafe.Pointer
}
type ImColor struct {
	Value ImVec4
}
type ImGuiMultiSelectIO struct {
	Requests      ImVector[ImGuiSelectionRequest]
	RangeSrcItem  ImGuiSelectionUserData
	NavIdItem     ImGuiSelectionUserData
	NavIdSelected cc.Bool
	RangeSrcReset cc.Bool
	ItemsCount    int32
}
type ImGuiSelectionRequest struct {
	Type           ImGuiSelectionRequestType
	Selected       cc.Bool
	RangeDirection ImS8
	RangeFirstItem ImGuiSelectionUserData
	RangeLastItem  ImGuiSelectionUserData
}
type ImGuiSelectionBasicStorage struct {
	Size                    int32
	PreserveOrder           cc.Bool
	UserData                unsafe.Pointer
	AdapterIndexToStorageId cc.Func // ImGuiID (*AdapterIndexToStorageId)(ImGuiSelectionBasicStorage* self, int idx)
	_SelectionOrder         int32
	_Storage                ImGuiStorage
}
type ImGuiSelectionExternalStorage struct {
	UserData               unsafe.Pointer
	AdapterSetItemSelected cc.Func // void (*AdapterSetItemSelected)(ImGuiSelectionExternalStorage* self, int idx, bool selected)
}
type ImDrawCmd struct {
	ClipRect               ImVec4
	TexRef                 ImTextureRef
	VtxOffset              uint32
	IdxOffset              uint32
	ElemCount              uint32
	UserCallback           cc.Func // void (*ImDrawCallback)(const ImDrawList* parent_list, const ImDrawCmd* cmd);
	UserCallbackData       unsafe.Pointer
	UserCallbackDataSize   int32
	UserCallbackDataOffset int32
}
type ImDrawVert struct {
	Pos ImVec2
	Uv  ImVec2
	Col ImU32
}
type ImDrawCmdHeader struct {
	ClipRect  ImVec4
	TexRef    ImTextureRef
	VtxOffset uint32
}
type ImDrawChannel struct {
	_CmdBuffer ImVector[ImDrawCmd]
	_IdxBuffer ImVector[ImDrawIdx]
}
type ImDrawListSplitter struct {
	_Current  int32
	_Count    int32
	_Channels ImVector[ImDrawChannel]
}
type ImDrawList struct {
	CmdBuffer         ImVector[ImDrawCmd]
	IdxBuffer         ImVector[ImDrawIdx]
	VtxBuffer         ImVector[ImDrawVert]
	Flags             ImDrawListFlags
	_VtxCurrentIdx    uint32
	_Data             *ImDrawListSharedData
	_VtxWritePtr      *ImDrawVert
	_IdxWritePtr      *ImDrawIdx
	_Path             ImVector[ImVec2]
	_CmdHeader        ImDrawCmdHeader
	_Splitter         ImDrawListSplitter
	_ClipRectStack    ImVector[ImVec4]
	_TextureStack     ImVector[ImTextureRef]
	_CallbacksDataBuf ImVector[ImU8]
	_FringeScale      float32
	_OwnerName        cc.String
}
type ImDrawData struct {
	Valid            cc.Bool
	CmdListsCount    int32
	TotalIdxCount    int32
	TotalVtxCount    int32
	CmdLists         ImVector[*ImDrawList]
	DisplayPos       ImVec2
	DisplaySize      ImVec2
	FramebufferScale ImVec2
	OwnerViewport    *ImGuiViewport
	Textures         *ImVector[*ImTextureData]
}
type ImTextureRect struct {
	X uint16
	Y uint16
	W uint16
	H uint16
}
type ImTextureData struct {
	UniqueID             int32
	Status               ImTextureStatus
	BackendUserData      unsafe.Pointer
	TexID                ImTextureID
	Format               ImTextureFormat
	Width                int32
	Height               int32
	BytesPerPixel        int32
	Pixels               *byte
	UsedRect             ImTextureRect
	UpdateRect           ImTextureRect
	Updates              ImVector[ImTextureRect]
	UnusedFrames         int32
	RefCount             uint16
	UseColors            cc.Bool
	WantDestroyNextFrame cc.Bool
}
type ImFontConfig struct {
	Name                 [40]byte
	FontData             unsafe.Pointer
	FontDataSize         int32
	FontDataOwnedByAtlas cc.Bool
	MergeMode            cc.Bool
	PixelSnapH           cc.Bool
	PixelSnapV           cc.Bool
	OversampleH          ImS8
	OversampleV          ImS8
	EllipsisChar         ImWchar
	SizePixels           float32
	GlyphRanges          *ImWchar
	GlyphExcludeRanges   *ImWchar
	GlyphOffset          ImVec2
	GlyphMinAdvanceX     float32
	GlyphMaxAdvanceX     float32
	GlyphExtraAdvanceX   float32
	FontNo               ImU32
	FontLoaderFlags      uint32
	RasterizerMultiply   float32
	RasterizerDensity    float32
	Flags                ImFontFlags
	DstFont              *ImFont
	FontLoader           *ImFontLoader
	FontLoaderData       unsafe.Pointer
}
type ImFontGlyph struct {
	Colored   uint32
	Visible   uint32
	SourceIdx uint32
	Codepoint uint32
	AdvanceX  float32
	X0        float32
	Y0        float32
	X1        float32
	Y1        float32
	U0        float32
	V0        float32
	U1        float32
	V1        float32
	PackId    int32
}
type ImFontGlyphRangesBuilder struct {
	UsedChars ImVector[ImU32]
}
type ImFontAtlasRect struct {
	X   uint16
	Y   uint16
	W   uint16
	H   uint16
	Uv0 ImVec2
	Uv1 ImVec2
}

const IM_DRAWLIST_TEX_LINES_WIDTH_MAX = 32

type ImFontAtlas struct {
	Flags            ImFontAtlasFlags
	TexDesiredFormat ImTextureFormat
	TexGlyphPadding  int32
	TexMinWidth      int32
	TexMinHeight     int32
	TexMaxWidth      int32
	TexMaxHeight     int32
	UserData         unsafe.Pointer

	TexRef ImTextureRef
	TexID  ImTextureRef // __anonymous_type1   __anonymous_type1

	TexData             *ImTextureData
	TexList             ImVector[*ImTextureData]
	Locked              cc.Bool
	RendererHasTextures cc.Bool
	TexIsBuilt          cc.Bool
	TexPixelsUseColors  cc.Bool
	TexUvScale          ImVec2
	TexUvWhitePixel     ImVec2
	Fonts               ImVector[*ImFont]
	Sources             ImVector[ImFontConfig]
	TexUvLines          [IM_DRAWLIST_TEX_LINES_WIDTH_MAX + 1]ImVec4
	TexNextUniqueID     int32
	FontNextUniqueID    int32
	DrawListSharedDatas ImVector[*ImDrawListSharedData]
	Builder             *ImFontAtlasBuilder
	FontLoader          *ImFontLoader
	FontLoaderName      cc.String
	FontLoaderData      unsafe.Pointer
	FontLoaderFlags     uint32
	RefCount            int32
	OwnerContext        *ImGuiContext
	TempRect            ImFontAtlasRect
}
type __anonymous_type1 struct {
	TexRef ImTextureRef
	TexID  ImTextureRef
}
type ImFontBaked struct {
	IndexAdvanceX       ImVector[float32]
	FallbackAdvanceX    float32
	Size                float32
	RasterizerDensity   float32
	IndexLookup         ImVector[ImU16]
	Glyphs              ImVector[ImFontGlyph]
	FallbackGlyphIndex  int32
	Ascent              float32
	Descent             float32
	MetricsTotalSurface uint32
	WantDestroy         uint32
	LockLoadingFallback uint32
	LastUsedFrame       int32
	BakedId             ImGuiID
	ContainerFont       *ImFont
	FontLoaderDatas     unsafe.Pointer
}

const IM_UNICODE_CODEPOINT_MAX = 0xffff

type ImFont struct {
	LastBaked                *ImFontBaked
	ContainerAtlas           *ImFontAtlas
	Flags                    ImFontFlags
	CurrentRasterizerDensity float32
	FontId                   ImGuiID
	LegacySize               float32
	Sources                  ImVector[*ImFontConfig]
	EllipsisChar             ImWchar
	FallbackChar             ImWchar
	Used8kPagesMap           [(IM_UNICODE_CODEPOINT_MAX + 1) / 8192 / 8]ImU8
	EllipsisAutoBake         cc.Bool
	RemapPairs               ImGuiStorage
	Scale                    float32
}
type ImGuiViewport struct {
	ID                    ImGuiID
	Flags                 ImGuiViewportFlags
	Pos                   ImVec2
	Size                  ImVec2
	FramebufferScale      ImVec2
	WorkPos               ImVec2
	WorkSize              ImVec2
	DpiScale              float32
	ParentViewportId      ImGuiID
	DrawData              *ImDrawData
	RendererUserData      unsafe.Pointer
	PlatformUserData      unsafe.Pointer
	PlatformHandle        unsafe.Pointer
	PlatformHandleRaw     unsafe.Pointer
	PlatformWindowCreated cc.Bool
	PlatformRequestMove   cc.Bool
	PlatformRequestResize cc.Bool
	PlatformRequestClose  cc.Bool
}
type ImGuiPlatformIO struct {
	Platform_GetClipboardTextFn        cc.Func // const char* (*Platform_GetClipboardTextFn)(ImGuiContext* ctx)
	Platform_SetClipboardTextFn        cc.Func // void (*Platform_SetClipboardTextFn)(ImGuiContext* ctx, const char* text)
	Platform_ClipboardUserData         unsafe.Pointer
	Platform_OpenInShellFn             cc.Func // bool (*Platform_OpenInShellFn)(ImGuiContext* ctx, const char* path)
	Platform_OpenInShellUserData       unsafe.Pointer
	Platform_SetImeDataFn              cc.Func // void (*Platform_SetImeDataFn)(ImGuiContext* ctx, ImGuiViewport* viewport, ImGuiPlatformImeData* data)
	Platform_ImeUserData               unsafe.Pointer
	Platform_LocaleDecimalPoint        ImWchar
	Renderer_TextureMaxWidth           int32
	Renderer_TextureMaxHeight          int32
	Renderer_RenderState               unsafe.Pointer
	Platform_CreateWindow              cc.Func // void (*Platform_CreateWindow)(ImGuiViewport* vp)
	Platform_DestroyWindow             cc.Func // void (*Platform_DestroyWindow)(ImGuiViewport* vp)
	Platform_ShowWindow                cc.Func // void (*Platform_ShowWindow)(ImGuiViewport* vp)
	Platform_SetWindowPos              cc.Func // void (*Platform_SetWindowPos)(ImGuiViewport* vp, ImVec2 pos)
	Platform_GetWindowPos              cc.Func // ImVec2 (*Platform_GetWindowPos)(ImGuiViewport* vp)
	Platform_SetWindowSize             cc.Func // void (*Platform_SetWindowSize)(ImGuiViewport* vp, ImVec2 size)
	Platform_GetWindowSize             cc.Func // ImVec2 (*Platform_GetWindowSize)(ImGuiViewport* vp)
	Platform_GetWindowFramebufferScale cc.Func // ImVec2 (*Platform_GetWindowFramebufferScale)(ImGuiViewport* vp)
	Platform_SetWindowFocus            cc.Func // void (*Platform_SetWindowFocus)(ImGuiViewport* vp)
	Platform_GetWindowFocus            cc.Func // bool (*Platform_GetWindowFocus)(ImGuiViewport* vp)
	Platform_GetWindowMinimized        cc.Func // bool (*Platform_GetWindowMinimized)(ImGuiViewport* vp)
	Platform_SetWindowTitle            cc.Func // void (*Platform_SetWindowTitle)(ImGuiViewport* vp, const char* str)
	Platform_SetWindowAlpha            cc.Func // void (*Platform_SetWindowAlpha)(ImGuiViewport* vp, float alpha)
	Platform_UpdateWindow              cc.Func // void (*Platform_UpdateWindow)(ImGuiViewport* vp)
	Platform_RenderWindow              cc.Func // void (*Platform_RenderWindow)(ImGuiViewport* vp, void* render_arg)
	Platform_SwapBuffers               cc.Func // void (*Platform_SwapBuffers)(ImGuiViewport* vp, void* render_arg)
	Platform_GetWindowDpiScale         cc.Func // float (*Platform_GetWindowDpiScale)(ImGuiViewport* vp)
	Platform_OnChangedViewport         cc.Func // void (*Platform_OnChangedViewport)(ImGuiViewport* vp)
	Platform_GetWindowWorkAreaInsets   cc.Func // ImVec4 (*Platform_GetWindowWorkAreaInsets)(ImGuiViewport* vp)
	Platform_CreateVkSurface           cc.Func // int (*Platform_CreateVkSurface)(ImGuiViewport* vp, ImU64 vk_inst, const void* vk_allocators, ImU64* out_vk_surface)
	Renderer_CreateWindow              cc.Func // void (*Renderer_CreateWindow)(ImGuiViewport* vp)
	Renderer_DestroyWindow             cc.Func // void (*Renderer_DestroyWindow)(ImGuiViewport* vp)
	Renderer_SetWindowSize             cc.Func // void (*Renderer_SetWindowSize)(ImGuiViewport* vp, ImVec2 size)
	Renderer_RenderWindow              cc.Func // void (*Renderer_RenderWindow)(ImGuiViewport* vp, void* render_arg)
	Renderer_SwapBuffers               cc.Func // void (*Renderer_SwapBuffers)(ImGuiViewport* vp, void* render_arg)
	Monitors                           ImVector[ImGuiPlatformMonitor]
	Textures                           ImVector[*ImTextureData]
	Viewports                          ImVector[*ImGuiViewport]
}
type ImGuiPlatformMonitor struct {
	MainPos        ImVec2
	MainSize       ImVec2
	WorkPos        ImVec2
	WorkSize       ImVec2
	DpiScale       float32
	PlatformHandle unsafe.Pointer
}
type ImGuiPlatformImeData struct {
	WantVisible     cc.Bool
	WantTextInput   cc.Bool
	InputPos        ImVec2
	InputLineHeight float32
	ViewportId      ImGuiID
}

// #endregion

func CreateContext(fa *ImFontAtlas /*= nil*/) *ImGuiContext { return imGui_CreateContext.Fn()(fa) }
func (ctx *ImGuiContext) Destroy()                          { imGui_DestroyContext.Fn()(ctx) }
func GetCurrentContext() *ImGuiContext                      { return imGui_GetCurrentContext.Fn()() }
func SetCurrentContext(ctx *ImGuiContext)                   { imGui_SetCurrentContext.Fn()(ctx) }

func GetIO() *ImGuiIO                 { return imGui_GetIO.Fn()() }
func GetPlatformIO() *ImGuiPlatformIO { return imGui_GetPlatformIO.Fn()() }
func GetStyle() *ImGuiStyle           { return imGui_GetStyle.Fn()() }
func NewFrame()                       { imGui_NewFrame.Fn()() }
func EndFrame()                       { imGui_EndFrame.Fn()() }
func Render()                         { imGui_Render.Fn()() }
func GetDrawData() *ImDrawData        { return imGui_GetDrawData.Fn()() }

func pBool(b *bool) *cc.Bool { return (*cc.Bool)(uptr(b)) }

func ShowDemoWindow(open *bool)         { imGui_ShowDemoWindow.Fn()(pBool(open)) }
func ShowMetricsWindow(open *bool)      { imGui_ShowMetricsWindow.Fn()(pBool(open)) }
func ShowDebugLogWindow(open *bool)     { imGui_ShowDebugLogWindow.Fn()(pBool(open)) }
func ShowIDStackToolWindow(open *bool)  { imGui_ShowIDStackToolWindow.Fn()(pBool(open)) }
func ShowAboutWindow(open *bool)        { imGui_ShowAboutWindow.Fn()(pBool(open)) }
func ShowStyleEditor(style *ImGuiStyle) { imGui_ShowStyleEditor.Fn()(style) }
func ShowStyleSelector(label cc.String) bool {
	return imGui_ShowStyleSelector.Fn()(label) == 1
}
func ShowFontSelector(label cc.String) { imGui_ShowFontSelector.Fn()(label) }
func ShowUserGuide()                   { imGui_ShowUserGuide.Fn()() }
func GetVersion() string               { return imGui_GetVersion.Fn()().Ref() }

func StyleColorsDark(dst *ImGuiStyle)    { imGui_StyleColorsDark.Fn()(dst) }
func StyleColorsLight(dst *ImGuiStyle)   { imGui_StyleColorsLight.Fn()(dst) }
func StyleColorsClassic(dst *ImGuiStyle) { imGui_StyleColorsClassic.Fn()(dst) }

func Begin(name cc.String, p_open *bool, flags ImGuiWindowFlags) bool {
	return imGui_Begin.Fn()(name, pBool(p_open), flags) != 0
}
func End() { imGui_End.Fn()() }

func BeginChild(str_id cc.String, size ImVec2, /* = ImVec2(0, 0) */
	child_flags ImGuiChildFlags /* = 0 */, window_flags ImGuiWindowFlags /* = 0 */) bool {
	return imGui_BeginChild.Fn()(str_id, size, child_flags, window_flags) != 0
}
func BeginChildID(id ImGuiID, size ImVec2, /* = ImVec2(0, 0) */
	child_flags ImGuiChildFlags /* = 0 */, window_flags ImGuiWindowFlags /* = 0 */) bool {
	return imGui_BeginChildID.Fn()(id, size, child_flags, window_flags) != 0
}
func EndChild() { imGui_EndChild.Fn()() }

func IsWindowAppearing() bool { return imGui_IsWindowAppearing.Fn()() != 0 }
func IsWindowCollapsed() bool { return imGui_IsWindowCollapsed.Fn()() != 0 }
func IsWindowFocused(flags ImGuiFocusedFlags /* = 0 */) bool {
	return imGui_IsWindowFocused.Fn()(flags) != 0
}
func IsWindowHovered(flags ImGuiHoveredFlags /* = 0 */) bool {
	return imGui_IsWindowHovered.Fn()(flags) != 0
}
func GetWindowDrawList() *ImDrawList    { return imGui_GetWindowDrawList.Fn()() }
func GetWindowDpiScale() float32        { return imGui_GetWindowDpiScale.Fn()() }
func GetWindowPos() ImVec2              { return imGui_GetWindowPos.Fn()() }
func GetWindowSize() ImVec2             { return imGui_GetWindowSize.Fn()() }
func GetWindowWidth() float32           { return imGui_GetWindowWidth.Fn()() }
func GetWindowHeight() float32          { return imGui_GetWindowHeight.Fn()() }
func GetWindowViewport() *ImGuiViewport { return imGui_GetWindowViewport.Fn()() }

func SetNextWindowPos(pos ImVec2, cond ImGuiCond /* = 0 */, pivot ImVec2 /* = ImVec2(0, 0) */) {
	imGui_SetNextWindowPos.Fn()(pos, cond, pivot)
}
func SetNextWindowSize(size ImVec2, cond ImGuiCond /* = 0 */) {
	imGui_SetNextWindowSize.Fn()(size, cond)
}
func SetNextWindowSizeConstraints(sizeMin ImVec2, sizeMax ImVec2,
	customCallback cc.Func, /* = NULL,typedef void (*ImGuiSizeCallback)(ImGuiSizeCallbackData* data); */
	customCallbackData *ImGuiSizeCallbackData /* = NULL */) {
	imGui_SetNextWindowSizeConstraints.Fn()(sizeMin, sizeMax, customCallback, customCallbackData)
}
func SetNextWindowContentSize(size ImVec2) { imGui_SetNextWindowContentSize.Fn()(size) }
func SetNextWindowCollapsed(collapsed bool, cond ImGuiCond /* = 0 */) {
	imGui_SetNextWindowCollapsed.Fn()(cc.MakeBool(collapsed), cond)
}
func SetNextWindowFocus()                                 { imGui_SetNextWindowFocus.Fn()() }
func SetNextWindowScroll(scroll ImVec2)                   { imGui_SetNextWindowScroll.Fn()(scroll) }
func SetNextWindowBgAlpha(alpha float32)                  { imGui_SetNextWindowBgAlpha.Fn()(alpha) }
func SetNextWindowViewport(viewportId ImGuiID)            { imGui_SetNextWindowViewport.Fn()(viewportId) }
func SetWindowPos(pos ImVec2, cond ImGuiCond /* = 0 */)   { imGui_SetWindowPos.Fn()(pos, cond) }
func SetWindowSize(size ImVec2, cond ImGuiCond /* = 0 */) { imGui_SetWindowSize.Fn()(size, cond) }
func SetWindowCollapsed(collapsed bool, cond ImGuiCond /* = 0 */) {
	imGui_SetWindowCollapsed.Fn()(cc.MakeBool(collapsed), cond)
}
func SetWindowFocus() { imGui_SetWindowFocus.Fn()() }
func SetWindowPosStr(name cc.String, pos ImVec2, cond ImGuiCond /* = 0 */) {
	imGui_SetWindowPosStr.Fn()(name, pos, cond)
}
func SetWindowSizeStr(name cc.String, size ImVec2, cond ImGuiCond /* = 0 */) {
	imGui_SetWindowSizeStr.Fn()(name, size, cond)
}
func SetWindowCollapsedStr(name cc.String, collapsed bool, cond ImGuiCond /* = 0 */) {
	imGui_SetWindowCollapsedStr.Fn()(name, cc.MakeBool(collapsed), cond)
}
func SetWindowFocusStr(name cc.String) { imGui_SetWindowFocusStr.Fn()(name) }

func GetScrollX() float32                              { return imGui_GetScrollX.Fn()() }
func GetScrollY() float32                              { return imGui_GetScrollY.Fn()() }
func SetScrollX(scrollX float32)                       { imGui_SetScrollX.Fn()(scrollX) }
func SetScrollY(scrollY float32)                       { imGui_SetScrollY.Fn()(scrollY) }
func GetScrollMaxX() float32                           { return imGui_GetScrollMaxX.Fn()() }
func GetScrollMaxY() float32                           { return imGui_GetScrollMaxY.Fn()() }
func SetScrollHereX(centerXRatio float32 /* = 0.5f */) { imGui_SetScrollHereX.Fn()(centerXRatio) }
func SetScrollHereY(centerYRatio float32 /* = 0.5f */) { imGui_SetScrollHereY.Fn()(centerYRatio) }
func SetScrollFromPosX(localX float32, centerXRatio float32 /* = 0.5f */) {
	imGui_SetScrollFromPosX.Fn()(localX, centerXRatio)
}
func SetScrollFromPosY(localY float32, centerYRatio float32 /* = 0.5f */) {
	imGui_SetScrollFromPosY.Fn()(localY, centerYRatio)
}

func PushFontSize(font *ImFont, fontSizeBaseUnscaled float32) {
	imGui_PushFontFloat.Fn()(font, fontSizeBaseUnscaled)
}
func PopFont()                   { imGui_PopFont.Fn()() }
func GetFont() *ImFont           { return imGui_GetFont.Fn()() }
func GetFontSize() float32       { return imGui_GetFontSize.Fn()() }
func GetFontBaked() *ImFontBaked { return imGui_GetFontBaked.Fn()() }

func PushStyleColor(idx ImGuiCol, col ImU32)           { imGui_PushStyleColor.Fn()(idx, col) }
func PushStyleColorImVec4(idx ImGuiCol, col ImVec4)    { imGui_PushStyleColorImVec4.Fn()(idx, col) }
func PopStyleColor(count int /* = 1 */)                { imGui_PopStyleColor.Fn()(count) }
func PushStyleVar(idx ImGuiStyleVar, val float32)      { imGui_PushStyleVar.Fn()(idx, val) }
func PushStyleVarImVec2(idx ImGuiStyleVar, val ImVec2) { imGui_PushStyleVarImVec2.Fn()(idx, val) }
func PushStyleVarX(idx ImGuiStyleVar, valX float32)    { imGui_PushStyleVarX.Fn()(idx, valX) }
func PushStyleVarY(idx ImGuiStyleVar, valY float32)    { imGui_PushStyleVarY.Fn()(idx, valY) }
func PopStyleVar(count int /* = 1 */)                  { imGui_PopStyleVar.Fn()(count) }
func PushItemFlag(option ImGuiItemFlags, enabled bool) {
	imGui_PushItemFlag.Fn()(option, cc.MakeBool(enabled))
}
func PopItemFlag() { imGui_PopItemFlag.Fn()() }

func PushItemWidth(itemWidth float32)                    { imGui_PushItemWidth.Fn()(itemWidth) }
func PopItemWidth()                                      { imGui_PopItemWidth.Fn()() }
func SetNextItemWidth(itemWidth float32)                 { imGui_SetNextItemWidth.Fn()(itemWidth) }
func CalcItemWidth() float32                             { return imGui_CalcItemWidth.Fn()() }
func PushTextWrapPos(wrapLocalPosX float32 /* = 0.0f */) { imGui_PushTextWrapPos.Fn()(wrapLocalPosX) }
func PopTextWrapPos()                                    { imGui_PopTextWrapPos.Fn()() }

func GetFontTexUvWhitePixel() ImVec2 { return imGui_GetFontTexUvWhitePixel.Fn()() }
func GetColorU32(idx ImGuiCol, alphaMul float32 /* = 1.0f */) ImU32 {
	return imGui_GetColorU32.Fn()(idx, alphaMul)
}
func GetColorU32ImVec4(col ImVec4) ImU32 { return imGui_GetColorU32ImVec4.Fn()(col) }
func GetColorU32ImU32(col ImU32, alphaMul float32 /* = 1.0f */) ImU32 {
	return imGui_GetColorU32ImU32.Fn()(col, alphaMul)
}
func GetStyleColorVec4(idx ImGuiCol) *ImVec4 { return imGui_GetStyleColorVec4.Fn()(idx) }

func GetCursorScreenPos() ImVec2    { return imGui_GetCursorScreenPos.Fn()() }
func SetCursorScreenPos(pos ImVec2) { imGui_SetCursorScreenPos.Fn()(pos) }
func GetContentRegionAvail() ImVec2 { return imGui_GetContentRegionAvail.Fn()() }
func GetCursorPos() ImVec2          { return imGui_GetCursorPos.Fn()() }
func GetCursorPosX() float32        { return imGui_GetCursorPosX.Fn()() }
func GetCursorPosY() float32        { return imGui_GetCursorPosY.Fn()() }
func SetCursorPos(localPos ImVec2)  { imGui_SetCursorPos.Fn()(localPos) }
func SetCursorPosX(localX float32)  { imGui_SetCursorPosX.Fn()(localX) }
func SetCursorPosY(localY float32)  { imGui_SetCursorPosY.Fn()(localY) }
func GetCursorStartPos() ImVec2     { return imGui_GetCursorStartPos.Fn()() }

func Separator() { imGui_Separator.Fn()() }
func SameLine(offsetFromStartX /* = 0.0f */, spacing float32 /* = -1.0f */) {
	imGui_SameLine.Fn()(offsetFromStartX, spacing)
}
func NewLine()                              { imGui_NewLine.Fn()() }
func Spacing()                              { imGui_Spacing.Fn()() }
func Dummy(size ImVec2)                     { imGui_Dummy.Fn()(size) }
func Indent(indentW float32 /* = 0.0f */)   { imGui_Indent.Fn()(indentW) }
func Unindent(indentW float32 /* = 0.0f */) { imGui_Unindent.Fn()(indentW) }
func BeginGroup()                           { imGui_BeginGroup.Fn()() }
func EndGroup()                             { imGui_EndGroup.Fn()() }
func AlignTextToFramePadding()              { imGui_AlignTextToFramePadding.Fn()() }
func GetTextLineHeight() float32            { return imGui_GetTextLineHeight.Fn()() }
func GetTextLineHeightWithSpacing() float32 { return imGui_GetTextLineHeightWithSpacing.Fn()() }
func GetFrameHeight() float32               { return imGui_GetFrameHeight.Fn()() }
func GetFrameHeightWithSpacing() float32    { return imGui_GetFrameHeightWithSpacing.Fn()() }

func PushID(strId cc.String)                   { imGui_PushID.Fn()(strId) }
func PushIDStr(strIdBegin, strIdEnd cc.String) { imGui_PushIDStr.Fn()(strIdBegin, strIdEnd) }
func PushIDPtr(ptrId unsafe.Pointer)           { imGui_PushIDPtr.Fn()(ptrId) }
func PushIDInt(intId int32)                    { imGui_PushIDInt.Fn()(intId) }
func PopID()                                   { imGui_PopID.Fn()() }
func GetID(strId cc.String) ImGuiID            { return imGui_GetID.Fn()(strId) }
func GetIDStr(strIdBegin, strIdEnd cc.String) ImGuiID {
	return imGui_GetIDStr.Fn()(strIdBegin, strIdEnd)
}
func GetIDPtr(ptrId unsafe.Pointer) ImGuiID { return imGui_GetIDPtr.Fn()(ptrId) }
func GetIDInt(intId int32) ImGuiID          { return imGui_GetIDInt.Fn()(intId) }

func Text(text cc.String, textEnd cc.String /* = NULL */) {
	imGui_TextUnformatted.Fn()(text, textEnd)
}
func Textf(fmt cc.String, args ...interface{}) {
	imGui_Text.FnVa()(fmt, append(args, uptr(nil))...)
}
func TextColoredf(col ImVec4, fmt cc.String, args ...interface{}) {
	imGui_TextColored.FnVa()(col, fmt, args...)
}
func TextColored(col ImVec4, text cc.String)           { imGui_TextColoredUnformatted.Fn()(col, text) }
func TextDisabledf(fmt cc.String, args ...interface{}) { imGui_TextDisabled.FnVa()(fmt, args...) }
func TextDisabled(text cc.String)                      { imGui_TextDisabledUnformatted.Fn()(text) }
func TextWrappedf(fmt cc.String, args ...interface{})  { imGui_TextWrapped.FnVa()(fmt, args...) }
func TextWrapped(text cc.String)                       { imGui_TextWrappedUnformatted.Fn()(text) }
func LabelTextf(label cc.String, fmt cc.String, args ...interface{}) {
	imGui_LabelText.FnVa()(label, fmt, args...)
}
func LabelText(label cc.String, text cc.String) {
	imGui_LabelTextUnformatted.Fn()(label, text)
}
func BulletTextf(fmt cc.String, args ...interface{}) { imGui_BulletText.FnVa()(fmt, args...) }
func BulletText(text cc.String)                      { imGui_BulletTextUnformatted.Fn()(text) }
func SeparatorText(label cc.String)                  { imGui_SeparatorText.Fn()(label) }

func Button(label cc.String, size ImVec2 /* = ImVec2(0, 0) */) bool {
	return imGui_Button.Fn()(label, size) != 0
}
func SmallButton(label cc.String) bool { return imGui_SmallButton.Fn()(label) != 0 }
func InvisibleButton(str_id cc.String, size ImVec2, flags ImGuiButtonFlags /* = 0 */) bool {
	return imGui_InvisibleButton.Fn()(str_id, size, flags) != 0
}
func ArrowButton(str_id cc.String, dir ImGuiDir) bool {
	return imGui_ArrowButton.Fn()(str_id, dir) != 0
}
func Checkbox(label cc.String, v *bool) bool { return imGui_Checkbox.Fn()(label, pBool(v)) != 0 }
func CheckboxFlagsIntPtr(label cc.String, flags *int32, flags_value int32) bool {
	return imGui_CheckboxFlagsIntPtr.Fn()(label, flags, flags_value) != 0
}
func CheckboxFlagsUintPtr(label cc.String, flags *uint32, flags_value uint32) bool {
	return imGui_CheckboxFlagsUintPtr.Fn()(label, flags, flags_value) != 0
}
func RadioButton(label cc.String, active bool) bool {
	return imGui_RadioButton.Fn()(label, cc.MakeBool(active)) != 0
}
func RadioButtonIntPtr(label cc.String, v *int32, v_button int32) bool {
	return imGui_RadioButtonIntPtr.Fn()(label, v, v_button) != 0
}
func ProgressBar(fraction float32, size_arg ImVec2 /* = ImVec2(-FLT_MIN, 0) */, overlay cc.String /* = NULL */) {
	imGui_ProgressBar.Fn()(fraction, size_arg, overlay)
}
func Bullet()                       { imGui_Bullet.Fn()() }
func TextLink(label cc.String) bool { return imGui_TextLink.Fn()(label) != 0 }
func TextLinkOpenURL(label cc.String, url cc.String /* = NULL */) bool {
	return imGui_TextLinkOpenURL.Fn()(label, url) != 0
}

func Image(tex_ref ImTextureRef, image_size ImVec2, uv0 ImVec2 /* = ImVec2(0, 0) */, uv1 ImVec2 /* = ImVec2(1, 1) */) {
	imGui_Image.Fn()(tex_ref, image_size, uv0, uv1)
}
func ImageWithBg(tex_ref ImTextureRef, image_size ImVec2, uv0 ImVec2 /* = ImVec2(0, 0) */, uv1 ImVec2 /* = ImVec2(1, 1) */, bg_col ImVec4 /* = ImVec4(0, 0, 0, 0) */, tint_col ImVec4 /* = ImVec4(1, 1, 1, 1) */) {
	imGui_ImageWithBg.Fn()(tex_ref, image_size, uv0, uv1, bg_col, tint_col)
}
func ImageButton(str_id cc.String, tex_ref ImTextureRef, image_size ImVec2, uv0 ImVec2 /* = ImVec2(0, 0) */, uv1 ImVec2 /* = ImVec2(1, 1) */, bg_col ImVec4 /* = ImVec4(0, 0, 0, 0) */, tint_col ImVec4 /* = ImVec4(1, 1, 1, 1) */) bool {
	return imGui_ImageButton.Fn()(str_id, tex_ref, image_size, uv0, uv1, bg_col, tint_col) != 0
}

func BeginCombo(label cc.String, preview_value cc.String, flags ImGuiComboFlags /* = 0 */) bool {
	return imGui_BeginCombo.Fn()(label, preview_value, flags) != 0
}
func EndCombo() { imGui_EndCombo.Fn()() }
func ComboChar(label cc.String, current_item *int32, items cc.Strings, items_count int32, popup_max_height_in_items int32 /* = -1 */) bool {
	return imGui_ComboChar.Fn()(label, current_item, items, items_count, popup_max_height_in_items) != 0
}
func Combo(label cc.String, current_item *int32, items_separated_by_zeros cc.String, popup_max_height_in_items int32 /* = -1 */) bool {
	return imGui_Combo.Fn()(label, current_item, items_separated_by_zeros, popup_max_height_in_items) != 0
}
func ComboCallback(label cc.String, current_item *int32,
	getter cc.Func /* const char* (*getter)(void* user_data, int idx)*/, user_data unsafe.Pointer, items_count int32, popup_max_height_in_items int32 /* = -1 */) bool {
	return imGui_ComboCallback.Fn()(label, current_item, getter, user_data, items_count, popup_max_height_in_items) != 0
}

func DragFloat(label cc.String, v *float32, v_speed float32, /* = 1.0f */
	v_min float32 /* = 0.0f */, v_max float32 /* = 0.0f */, format cc.String /* = "%.3f" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_DragFloat.Fn()(label, v, v_speed, v_min, v_max, format, flags) != 0
}
func DragFloat2(label cc.String, v []float32, v_speed float32, /* = 1.0f */
	v_min float32 /* = 0.0f */, v_max float32 /* = 0.0f */, format cc.String /* = "%.3f" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_DragFloat2.Fn()(label, &v[0], v_speed, v_min, v_max, format, flags) != 0
}
func DragFloat3(label cc.String, v []float32, v_speed float32, /* = 1.0f */
	v_min float32 /* = 0.0f */, v_max float32 /* = 0.0f */, format cc.String /* = "%.3f" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_DragFloat3.Fn()(label, &v[0], v_speed, v_min, v_max, format, flags) != 0
}
func DragFloat4(label cc.String, v []float32, v_speed float32, /* = 1.0f */
	v_min float32 /* = 0.0f */, v_max float32 /* = 0.0f */, format cc.String /* = "%.3f" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_DragFloat4.Fn()(label, &v[0], v_speed, v_min, v_max, format, flags) != 0
}
func DragFloatRange2(label cc.String, v_current_min *float32, v_current_max *float32,
	v_speed float32 /* = 1.0f */, v_min float32 /* = 0.0f */, v_max float32, /* = 0.0f */
	format cc.String /* = "%.3f" */, format_max cc.String /* = NULL */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_DragFloatRange2.Fn()(label, v_current_min, v_current_max, v_speed, v_min, v_max, format, format_max, flags) != 0
}
func DragInt(label cc.String, v *int32, v_speed float32, /* = 1.0f */
	v_min int32 /* = 0 */, v_max int32 /* = 0 */, format cc.String /* = "%d" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_DragInt.Fn()(label, v, v_speed, v_min, v_max, format, flags) != 0
}
func DragInt2(label cc.String, v []int32, v_speed float32, /* = 1.0f */
	v_min int32 /* = 0 */, v_max int32 /* = 0 */, format cc.String /* = "%d" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_DragInt2.Fn()(label, &v[0], v_speed, v_min, v_max, format, flags) != 0
}
func DragInt3(label cc.String, v []int32, v_speed float32, /* = 1.0f */
	v_min int32 /* = 0 */, v_max int32 /* = 0 */, format cc.String /* = "%d" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_DragInt3.Fn()(label, &v[0], v_speed, v_min, v_max, format, flags) != 0
}
func DragInt4(label cc.String, v []int32, v_speed float32, /* = 1.0f */
	v_min int32 /* = 0 */, v_max int32 /* = 0 */, format cc.String /* = "%d" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_DragInt4.Fn()(label, &v[0], v_speed, v_min, v_max, format, flags) != 0
}
func DragIntRange2(label cc.String, v_current_min *int32, v_current_max *int32,
	v_speed float32 /* = 1.0f */, v_min int32 /* = 0 */, v_max int32, /* = 0 */
	format cc.String /* = "%d" */, format_max cc.String /* = NULL */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_DragIntRange2.Fn()(label, v_current_min, v_current_max, v_speed, v_min, v_max, format, format_max, flags) != 0
}
func DragScalar[T any](label cc.String, data_type ImGuiDataType, p_data *T,
	v_speed float32 /* = 1.0f */, p_min *T /* = NULL */, p_max *T, /* = NULL */
	format cc.String /* = NULL */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_DragScalar.Fn()(label, data_type, uptr(p_data), v_speed, uptr(p_min), uptr(p_max), format, flags) != 0
}
func DragScalarN[T any](label cc.String, data_type ImGuiDataType, p_data []T,
	components int32, v_speed float32 /* = 1.0f */, p_min *T /* = NULL */, p_max *T, /* = NULL */
	format cc.String /* = NULL */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_DragScalarN.Fn()(label, data_type, uptr(&p_data[0]), components, v_speed, uptr(p_min), uptr(p_max), format, flags) != 0
}

func SliderFloat(label cc.String, v *float32, v_min, v_max float32,
	format cc.String /* = "%.3f" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_SliderFloat.Fn()(label, v, v_min, v_max, format, flags) != 0
}
func SliderFloat2(label cc.String, v []float32, v_min, v_max float32,
	format cc.String /* = "%.3f" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_SliderFloat2.Fn()(label, &v[0], v_min, v_max, format, flags) != 0
}
func SliderFloat3(label cc.String, v []float32, v_min, v_max float32,
	format cc.String /* = "%.3f" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_SliderFloat3.Fn()(label, &v[0], v_min, v_max, format, flags) != 0
}
func SliderFloat4(label cc.String, v []float32, v_min, v_max float32,
	format cc.String /* = "%.3f" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_SliderFloat4.Fn()(label, &v[0], v_min, v_max, format, flags) != 0
}
func SliderAngle(label cc.String, v_rad *float32, v_degrees_min /* = -360.0f */, v_degrees_max float32, /* = +360.0f */
	format cc.String /* = "%.0f deg" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_SliderAngle.Fn()(label, v_rad, v_degrees_min, v_degrees_max, format, flags) != 0
}
func SliderInt(label cc.String, v *int32, v_min, v_max int32,
	format cc.String /* = "%d" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_SliderInt.Fn()(label, v, v_min, v_max, format, flags) != 0
}
func SliderInt2(label cc.String, v []int32, v_min, v_max int32,
	format cc.String /* = "%d" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_SliderInt2.Fn()(label, &v[0], v_min, v_max, format, flags) != 0
}
func SliderInt3(label cc.String, v []int32, v_min, v_max int32,
	format cc.String /* = "%d" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_SliderInt3.Fn()(label, &v[0], v_min, v_max, format, flags) != 0
}
func SliderInt4(label cc.String, v []int32, v_min, v_max int32,
	format cc.String /* = "%d" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_SliderInt4.Fn()(label, &v[0], v_min, v_max, format, flags) != 0
}
func SliderScalar[T any](label cc.String, data_type ImGuiDataType, p_data, p_min, p_max *T,
	format cc.String /* = NULL */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_SliderScalar.Fn()(label, data_type, uptr(p_data), uptr(p_min), uptr(p_max), format, flags) != 0
}
func SliderScalarN[T any](label cc.String, data_type ImGuiDataType, p_data []T, components int32,
	p_min *T, p_max *T, format cc.String /* = NULL */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_SliderScalarN.Fn()(label, data_type, uptr(&p_data[0]), components, uptr(p_min), uptr(p_max), format, flags) != 0
}
func VSliderFloat(label cc.String, size ImVec2, v *float32, v_min, v_max float32,
	format cc.String /* = "%.3f" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_VSliderFloat.Fn()(label, size, v, v_min, v_max, format, flags) != 0
}
func VSliderInt(label cc.String, size ImVec2, v *int32, v_min, v_max int32,
	format cc.String /* = "%d" */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_VSliderInt.Fn()(label, size, v, v_min, v_max, format, flags) != 0
}
func VSliderScalar[T any](label cc.String, size ImVec2, data_type ImGuiDataType, p_data, p_min, p_max *T,
	format cc.String /* = NULL */, flags ImGuiSliderFlags /* = 0 */) bool {
	return imGui_VSliderScalar.Fn()(label, size, data_type, uptr(p_data), uptr(p_min), uptr(p_max), format, flags) != 0
}

func InputText(label, buf cc.String, buf_size uint64, flags ImGuiInputTextFlags, /* = 0 */
	callback cc.Func /* int (*)(ImGuiInputTextCallbackData*) */, user_data unsafe.Pointer /* = NULL */) bool {
	return imGui_InputText.Fn()(label, buf, buf_size, flags, callback, user_data) != 0
}
func InputTextMultiline(label cc.String, buf cc.String, buf_size uint64,
	size ImVec2 /* = ImVec2(0, 0) */, flags ImGuiInputTextFlags, /* = 0 */
	callback cc.Func /* int (*)(ImGuiInputTextCallbackData*) */, user_data unsafe.Pointer /* = NULL */) bool {
	return imGui_InputTextMultiline.Fn()(label, buf, buf_size, size, flags, callback, user_data) != 0
}
func InputTextWithHint(label cc.String, hint cc.String, buf cc.String, buf_size uint64,
	flags ImGuiInputTextFlags /* = 0 */, callback cc.Func, /* int (*)(ImGuiInputTextCallbackData*) */
	user_data unsafe.Pointer /* = NULL */) bool {
	return imGui_InputTextWithHint.Fn()(label, hint, buf, buf_size, flags, callback, user_data) != 0
}
func InputFloat(label cc.String, v *float32, step float32 /* = 0.0f */, step_fast float32, /* = 0.0f */
	format cc.String /* = "%.3f" */, flags ImGuiInputTextFlags /* = 0 */) bool {
	return imGui_InputFloat.Fn()(label, v, step, step_fast, format, flags) != 0
}
func InputFloat2(label cc.String, v []float32, format cc.String, /* = "%.3f" */
	flags ImGuiInputTextFlags /* = 0 */) bool {
	return imGui_InputFloat2.Fn()(label, &v[0], format, flags) != 0
}
func InputFloat3(label cc.String, v []float32, format cc.String, /* = "%.3f" */
	flags ImGuiInputTextFlags /* = 0 */) bool {
	return imGui_InputFloat3.Fn()(label, &v[0], format, flags) != 0
}
func InputFloat4(label cc.String, v []float32, format cc.String, /* = "%.3f" */
	flags ImGuiInputTextFlags /* = 0 */) bool {
	return imGui_InputFloat4.Fn()(label, &v[0], format, flags) != 0
}
func InputInt(label cc.String, v *int32, step int32 /* = 1 */, step_fast int32, /* = 100 */
	flags ImGuiInputTextFlags /* = 0 */) bool {
	return imGui_InputInt.Fn()(label, v, step, step_fast, flags) != 0
}
func InputInt2(label cc.String, v []int32, flags ImGuiInputTextFlags /* = 0 */) bool {
	return imGui_InputInt2.Fn()(label, &v[0], flags) != 0
}
func InputInt3(label cc.String, v []int32, flags ImGuiInputTextFlags /* = 0 */) bool {
	return imGui_InputInt3.Fn()(label, &v[0], flags) != 0
}
func InputInt4(label cc.String, v []int32, flags ImGuiInputTextFlags /* = 0 */) bool {
	return imGui_InputInt4.Fn()(label, &v[0], flags) != 0
}
func InputDouble(label cc.String, v *float64, step float64 /* = 0.0 */, step_fast float64, /* = 0.0 */
	format cc.String /* = "%.6f" */, flags ImGuiInputTextFlags /* = 0 */) bool {
	return imGui_InputDouble.Fn()(label, v, step, step_fast, format, flags) != 0
}
func InputScalar[T any](label cc.String, data_type ImGuiDataType, p_data, p_step /* = NULL */, p_step_fast *T, /* = NULL */
	format cc.String /* = NULL */, flags ImGuiInputTextFlags /* = 0 */) bool {
	return imGui_InputScalar.Fn()(label, data_type, uptr(p_data), uptr(p_step), uptr(p_step_fast), format, flags) != 0
}
func InputScalarN[T any](label cc.String, data_type ImGuiDataType, p_data []T, components int32,
	p_step /* = NULL */, p_step_fast *T /* = NULL */, format cc.String /* = NULL */, flags ImGuiInputTextFlags /* = 0 */) bool {
	return imGui_InputScalarN.Fn()(label, data_type, uptr(&p_data[0]), components, uptr(p_step), uptr(p_step_fast), format, flags) != 0
}

func ColorEdit3(label cc.String, col []float32 /*[3]float32*/, flags ImGuiColorEditFlags /* = 0 */) bool {
	return imGui_ColorEdit3.Fn()(label, &col[0], flags) != 0
}
func ColorEdit4(label cc.String, col []float32 /*[4]float32*/, flags ImGuiColorEditFlags /* = 0 */) bool {
	return imGui_ColorEdit4.Fn()(label, &col[0], flags) != 0
}
func ColorPicker3(label cc.String, col []float32 /*[3]float32*/, flags ImGuiColorEditFlags /* = 0 */) bool {
	return imGui_ColorPicker3.Fn()(label, &col[0], flags) != 0
}
func ColorPicker4(label cc.String, col []float32 /*[4]float32*/, flags ImGuiColorEditFlags /* = 0 */, ref_col *float32 /* = NULL */) bool {
	return imGui_ColorPicker4.Fn()(label, &col[0], flags, ref_col) != 0
}
func ColorButton(desc_id cc.String, col ImVec4, flags ImGuiColorEditFlags /* = 0 */, size ImVec2 /* = ImVec2(0, 0) */) bool {
	return imGui_ColorButton.Fn()(desc_id, col, flags, size) != 0
}
func SetColorEditOptions(flags ImGuiColorEditFlags) { imGui_SetColorEditOptions.Fn()(flags) }

func TreeNode(label cc.String) bool { return imGui_TreeNode.Fn()(label) != 0 }
func TreeNodef(str_id cc.String, fmt cc.String, args ...any) bool {
	return imGui_TreeNodeStr.FnVa()(str_id, fmt, args...) != 0
}
func TreeNodeID(str_id cc.String, text cc.String) bool {
	return imGui_TreeNodeStrUnformatted.Fn()(str_id, text) != 0
}
func TreeNodePtrIDf(ptr_id unsafe.Pointer, fmt cc.String, args ...any) bool {
	return imGui_TreeNodePtr.FnVa()(ptr_id, fmt, args...) != 0
}
func TreeNodePtrID(ptr_id unsafe.Pointer, text cc.String) bool {
	return imGui_TreeNodePtrUnformatted.Fn()(ptr_id, text) != 0
}

//	func TreeNodeV(str_id cc.String, fmt cc.String, args unsafe.Pointer) bool {
//		return imGui_TreeNodeV.Fn()(str_id, fmt, args) != 0
//	}
// func TreeNodeVPtrFmt(ptr_id unsafe.Pointer, fmt cc.String, args unsafe.Pointer) bool {
// 	return imGui_TreeNodeVPtr.Fn()(ptr_id, fmt, args) != 0
// }

func TreeNodeEx(label cc.String, flags ImGuiTreeNodeFlags /* = 0 */) bool {
	return imGui_TreeNodeEx.Fn()(label, flags) != 0
}
func TreeNodeExIDf(str_id cc.String, flags ImGuiTreeNodeFlags, fmt cc.String, args ...any) bool {
	return imGui_TreeNodeExStr.FnVa()(str_id, flags, fmt, args...) != 0
}
func TreeNodeExID(str_id cc.String, flags ImGuiTreeNodeFlags, text cc.String) bool {
	return imGui_TreeNodeExStrUnformatted.Fn()(str_id, flags, text) != 0
}
func TreeNodeExPtrIDf(ptr_id unsafe.Pointer, flags ImGuiTreeNodeFlags, fmt cc.String, args ...any) bool {
	return imGui_TreeNodeExPtr.FnVa()(ptr_id, flags, fmt, args...) != 0
}
func TreeNodeExPtrID(ptr_id unsafe.Pointer, flags ImGuiTreeNodeFlags, text cc.String) bool {
	return imGui_TreeNodeExPtrUnformatted.Fn()(ptr_id, flags, text) != 0
}

//	func TreeNodeExV(str_id cc.String, flags ImGuiTreeNodeFlags, fmt cc.String, args unsafe.Pointer) bool {
//		return imGui_TreeNodeExV.Fn()(str_id, flags, fmt, args) != 0
//	}
//
//	func TreeNodeExVPtr(ptr_id unsafe.Pointer, flags ImGuiTreeNodeFlags, fmt cc.String, args unsafe.Pointer) bool {
//		return imGui_TreeNodeExVPtr.Fn()(ptr_id, flags, fmt, args) != 0
//	}

func TreePush(str_id cc.String)          { imGui_TreePush.Fn()(str_id) }
func TreePushPtr(ptr_id unsafe.Pointer)  { imGui_TreePushPtr.Fn()(ptr_id) }
func TreePop()                           { imGui_TreePop.Fn()() }
func GetTreeNodeToLabelSpacing() float32 { return imGui_GetTreeNodeToLabelSpacing.Fn()() }
func CollapsingHeader(label cc.String, flags ImGuiTreeNodeFlags /* = 0 */) bool {
	return imGui_CollapsingHeader.Fn()(label, flags) != 0
}
func CollapsingHeaderBoolPtr(label cc.String, p_visible *bool, flags ImGuiTreeNodeFlags /* = 0 */) bool {
	return imGui_CollapsingHeaderBoolPtr.Fn()(label, pBool(p_visible), flags) != 0
}
func SetNextItemOpen(is_open bool, cond ImGuiCond /* = 0 */) {
	imGui_SetNextItemOpen.Fn()(cc.MakeBool(is_open), cond)
}
func SetNextItemStorageID(storage_id ImGuiID) { imGui_SetNextItemStorageID.Fn()(storage_id) }

func Selectable(label cc.String, selected bool /* = false */, flags ImGuiSelectableFlags /* = 0 */, size ImVec2 /* = ImVec2(0, 0) */) bool {
	return imGui_Selectable.Fn()(label, cc.MakeBool(selected), flags, size) != 0
}
func SelectableOut(label cc.String, Outselected *bool, flags ImGuiSelectableFlags /* = 0 */, size ImVec2 /* = ImVec2(0, 0) */) bool {
	return imGui_SelectableBoolPtr.Fn()(label, pBool(Outselected), flags, size) != 0
}
func BeginMultiSelect(flags ImGuiMultiSelectFlags, selection_size int32 /* = -1 */, items_count int32 /* = -1 */) *ImGuiMultiSelectIO {
	return imGui_BeginMultiSelect.Fn()(flags, selection_size, items_count)
}
func EndMultiSelect() *ImGuiMultiSelectIO { return imGui_EndMultiSelect.Fn()() }
func SetNextItemSelectionUserData(data ImGuiSelectionUserData) {
	imGui_SetNextItemSelectionUserData.Fn()(data)
}
func IsItemToggledSelection() bool { return imGui_IsItemToggledSelection.Fn()() != 0 }
func BeginListBox(label cc.String, size ImVec2 /* = ImVec2(0, 0) */) bool {
	return imGui_BeginListBox.Fn()(label, size) != 0
}
func EndListBox() { imGui_EndListBox.Fn()() }
func ListBox(label cc.String, current *int32, items []cc.String, height_in_items int32 /* = -1 */) bool {
	return imGui_ListBox.Fn()(label, current, &items[0], int32(len(items)), height_in_items) != 0
}
func ListBoxCallback(label cc.String, current_item *int32,
	getter cc.Func /* const char* (*getter)(void* user_data, int idx)*/, user_data unsafe.Pointer, items_count int32, height_in_items int32 /* = -1 */) bool {
	return imGui_ListBoxCallback.Fn()(label, current_item, getter, user_data, items_count, height_in_items) != 0
}

func PlotLines(label cc.String, values []float32, offset int32, /* = 0 */
	overlay_text cc.String /* = NULL */, scale_min float32 /* = FLT_MAX */, scale_max float32, /* = FLT_MAX */
	graph_size ImVec2 /* = ImVec2(0, 0) */, stride int32 /* = sizeof(float) */) {
	imGui_PlotLines.Fn()(label, &values[0], int32(len(values)), offset, overlay_text, scale_min, scale_max, graph_size, stride)
}
func PlotLinesCallback(label cc.String, values_getter cc.Func, /* float (*values_getter)(void* data, int idx)*/
	data unsafe.Pointer, count int32, offset int32 /* = 0 */, overlay_text cc.String, /* = NULL */
	scale_min float32 /* = FLT_MAX */, scale_max float32 /* = FLT_MAX */, graph_size ImVec2 /* = ImVec2(0, 0) */) {
	imGui_PlotLinesCallback.Fn()(label, values_getter, data, count, offset, overlay_text, scale_min, scale_max, graph_size)
}
func PlotHistogram(label cc.String, values []float32, values_offset int32, /* = 0 */
	overlay_text cc.String /* = NULL */, scale_min float32, /* = FLT_MAX */
	scale_max float32 /* = FLT_MAX */, graph_size ImVec2 /* = ImVec2(0, 0) */, stride int32 /* = sizeof(float) */) {
	imGui_PlotHistogram.Fn()(label, &values[0], int32(len(values)), values_offset, overlay_text, scale_min, scale_max, graph_size, stride)
}
func PlotHistogramCallback(label cc.String, values_getter cc.Func, /* float (*values_getter)(void* data, int idx)*/
	data unsafe.Pointer, values_count int32, values_offset int32 /* = 0 */, overlay_text cc.String, /* = NULL */
	scale_min float32 /* = FLT_MAX */, scale_max float32 /* = FLT_MAX */, graph_size ImVec2 /* = ImVec2(0, 0) */) {
	imGui_PlotHistogramCallback.Fn()(label, values_getter, data, values_count, values_offset, overlay_text, scale_min, scale_max, graph_size)
}
func BeginMenuBar() bool     { return imGui_BeginMenuBar.Fn()() != 0 }
func EndMenuBar()            { imGui_EndMenuBar.Fn()() }
func BeginMainMenuBar() bool { return imGui_BeginMainMenuBar.Fn()() != 0 }
func EndMainMenuBar()        { imGui_EndMainMenuBar.Fn()() }
func BeginMenu(label cc.String, enabled bool /* = true */) bool {
	return imGui_BeginMenu.Fn()(label, cc.MakeBool(enabled)) != 0
}
func EndMenu() { imGui_EndMenu.Fn()() }
func MenuItem(label cc.String, shortcut cc.String /* = NULL */, selected bool /* = false */, enabled bool /* = true */) bool {
	return imGui_MenuItem.Fn()(label, shortcut, cc.MakeBool(selected), cc.MakeBool(enabled)) != 0
}
func MenuItemBoolPtr(label cc.String, shortcut cc.String, p_selected *bool, enabled bool /* = true */) bool {
	return imGui_MenuItemBoolPtr.Fn()(label, shortcut, pBool(p_selected), cc.MakeBool(enabled)) != 0
}

func BeginTooltip() bool                         { return imGui_BeginTooltip.Fn()() != 0 }
func EndTooltip()                                { imGui_EndTooltip.Fn()() }
func SetTooltipf(fmt cc.String, args ...any)     { imGui_SetTooltip.FnVa()(fmt, args...) }
func SetTooltip(text cc.String)                  { imGui_SetTooltipUnformatted.Fn()(text) }
func BeginItemTooltip() bool                     { return imGui_BeginItemTooltip.Fn()() != 0 }
func EndItemTooltip()                            { imGui_EndTooltip.Fn()() }
func SetItemTooltipf(fmt cc.String, args ...any) { imGui_SetItemTooltip.FnVa()(fmt, args...) }
func SetItemTooltip(text cc.String)              { imGui_SetItemTooltipUnformatted.Fn()(text) }

func BeginPopup(str_id cc.String, flags ImGuiPopupFlags /* = 0 */) bool {
	return imGui_BeginPopup.Fn()(str_id, flags) != 0
}
func BeginPopupModal(name cc.String, p_open *bool /* = NULL */, flags ImGuiWindowFlags /* = 0 */) bool {
	return imGui_BeginPopupModal.Fn()(name, pBool(p_open), flags) != 0
}
func EndPopup() { imGui_EndPopup.Fn()() }
func OpenPopup(str_id cc.String, popup_flags ImGuiPopupFlags /* = 0 */) {
	imGui_OpenPopup.Fn()(str_id, popup_flags)
}
func OpenPopupID(id ImGuiID, popup_flags ImGuiPopupFlags /* = 0 */) {
	imGui_OpenPopupID.Fn()(id, popup_flags)
}
func OpenPopupOnItemClick(str_id cc.String /* = NULL */, popup_flags ImGuiPopupFlags /* = 1 */) {
	imGui_OpenPopupOnItemClick.Fn()(str_id, popup_flags)
}
func CloseCurrentPopup() { imGui_CloseCurrentPopup.Fn()() }

func BeginPopupContextItem(str_id cc.String /* = NULL */, popup_flags ImGuiPopupFlags /* = 1 */) bool {
	return imGui_BeginPopupContextItem.Fn()(str_id, popup_flags) != 0
}
func BeginPopupContextWindow(str_id cc.String /* = NULL */, popup_flags ImGuiPopupFlags /* = 1 */) bool {
	return imGui_BeginPopupContextWindow.Fn()(str_id, popup_flags) != 0
}
func BeginPopupContextVoid(str_id cc.String /* = NULL */, popup_flags ImGuiPopupFlags /* = 1 */) bool {
	return imGui_BeginPopupContextVoid.Fn()(str_id, popup_flags) != 0
}
func IsPopupOpen(str_id cc.String, flags ImGuiPopupFlags /* = 0 */) bool {
	return imGui_IsPopupOpen.Fn()(str_id, flags) != 0
}
func BeginTable(str_id cc.String, columns int32, flags ImGuiTableFlags, /* = 0 */
	outer_size ImVec2 /* = ImVec2(0.0f, 0.0f) */, inner_width float32 /* = 0.0f */) bool {
	return imGui_BeginTable.Fn()(str_id, columns, flags, outer_size, inner_width) != 0
}
func EndTable() { imGui_EndTable.Fn()() }
func TableNextRow(row_flags ImGuiTableRowFlags /* = 0 */, min_row_height float32 /* = 0.0f */) {
	imGui_TableNextRow.Fn()(row_flags, min_row_height)
}
func TableNextColumn() bool                   { return imGui_TableNextColumn.Fn()() != 0 }
func TableSetColumnIndex(column_n int32) bool { return imGui_TableSetColumnIndex.Fn()(column_n) != 0 }

func TableSetupColumn(label cc.String, flags ImGuiTableColumnFlags /* = 0 */, init_width_or_weight float32 /* = 0.0f */, user_id ImGuiID /* = 0 */) {
	imGui_TableSetupColumn.Fn()(label, flags, init_width_or_weight, user_id)
}
func TableSetupScrollFreeze(cols int32, rows int32) { imGui_TableSetupScrollFreeze.Fn()(cols, rows) }
func TableHeader(label cc.String)                   { imGui_TableHeader.Fn()(label) }
func TableHeadersRow()                              { imGui_TableHeadersRow.Fn()() }
func TableAngledHeadersRow()                        { imGui_TableAngledHeadersRow.Fn()() }
func TableGetSortSpecs() *ImGuiTableSortSpecs       { return imGui_TableGetSortSpecs.Fn()() }
func TableGetColumnCount() int32                    { return imGui_TableGetColumnCount.Fn()() }
func TableGetColumnIndex() int32                    { return imGui_TableGetColumnIndex.Fn()() }
func TableGetRowIndex() int32                       { return imGui_TableGetRowIndex.Fn()() }
func TableGetColumnName(column_n int32 /* = -1 */) cc.String {
	return imGui_TableGetColumnName.Fn()(column_n)
}
func TableGetColumnFlags(column_n int32 /* = -1 */) ImGuiTableColumnFlags {
	return imGui_TableGetColumnFlags.Fn()(column_n)
}
func TableSetColumnEnabled(column_n int32, v bool) {
	imGui_TableSetColumnEnabled.Fn()(column_n, cc.MakeBool(v))
}
func TableGetHoveredColumn() int32 { return imGui_TableGetHoveredColumn.Fn()() }
func TableSetBgColor(target ImGuiTableBgTarget, color ImU32, column_n int32 /* = -1 */) {
	imGui_TableSetBgColor.Fn()(target, color, column_n)
}
func Columns(count int32 /* = 1 */, id cc.String /* = NULL */, borders bool /* = true */) {
	imGui_Columns.Fn()(count, id, cc.MakeBool(borders))
}
func NextColumn()           { imGui_NextColumn.Fn()() }
func GetColumnIndex() int32 { return imGui_GetColumnIndex.Fn()() }
func GetColumnWidth(column_index int32 /* = -1 */) float32 {
	return imGui_GetColumnWidth.Fn()(column_index)
}
func SetColumnWidth(column_index int32, width float32) {
	imGui_SetColumnWidth.Fn()(column_index, width)
}
func GetColumnOffset(column_index int32 /* = -1 */) float32 {
	return imGui_GetColumnOffset.Fn()(column_index)
}
func SetColumnOffset(column_index int32, offset_x float32) {
	imGui_SetColumnOffset.Fn()(column_index, offset_x)
}
func GetColumnsCount() int32 { return imGui_GetColumnsCount.Fn()() }

func BeginTabBar(str_id cc.String, flags ImGuiTabBarFlags /* = 0 */) bool {
	return imGui_BeginTabBar.Fn()(str_id, flags) != 0
}
func EndTabBar() { imGui_EndTabBar.Fn()() }
func BeginTabItem(label cc.String, p_open *bool, flags ImGuiTabItemFlags /* = 0 */) bool {
	return imGui_BeginTabItem.Fn()(label, pBool(p_open), flags) != 0
}
func EndTabItem() { imGui_EndTabItem.Fn()() }
func TabItemButton(label cc.String, flags ImGuiTabItemFlags /* = 0 */) bool {
	return imGui_TabItemButton.Fn()(label, flags) != 0
}
func SetTabItemClosed(tab_or_docked_window_label cc.String) {
	imGui_SetTabItemClosed.Fn()(tab_or_docked_window_label)
}
func DockSpace(dockspace_id ImGuiID, size ImVec2 /* = ImVec2(0, 0) */, flags ImGuiDockNodeFlags /* = 0 */, window_class *ImGuiWindowClass /* = NULL */) ImGuiID {
	return imGui_DockSpace.Fn()(dockspace_id, size, flags, window_class)
}
func DockSpaceOverViewport(dockspace_id ImGuiID /* = 0 */, viewport *ImGuiViewport, /* = NULL */
	flags ImGuiDockNodeFlags /* = 0 */, window_class *ImGuiWindowClass /* = NULL */) ImGuiID {
	return imGui_DockSpaceOverViewport.Fn()(dockspace_id, viewport, flags, window_class)
}
func SetNextWindowDockID(dock_id ImGuiID, cond ImGuiCond /* = 0 */) {
	imGui_SetNextWindowDockID.Fn()(dock_id, cond)
}
func SetNextWindowClass(window_class *ImGuiWindowClass) { imGui_SetNextWindowClass.Fn()(window_class) }
func GetWindowDockID() ImGuiID                          { return imGui_GetWindowDockID.Fn()() }
func IsWindowDocked() bool                              { return imGui_IsWindowDocked.Fn()() != 0 }

func LogToTTY(auto_open_depth int32 /* = -1 */) { imGui_LogToTTY.Fn()(auto_open_depth) }
func LogToFile(auto_open_depth int32 /* = -1 */, filename cc.String /* = NULL */) {
	imGui_LogToFile.Fn()(auto_open_depth, filename)
}
func LogToClipboard(auto_open_depth int32 /* = -1 */) { imGui_LogToClipboard.Fn()(auto_open_depth) }
func LogFinish()                                      { imGui_LogFinish.Fn()() }
func LogButtons()                                     { imGui_LogButtons.Fn()() }
func LogTextf(fmt cc.String, args ...any)             { imGui_LogText.FnVa()(fmt, args...) }
func LogText(text cc.String)                          { imGui_LogTextUnformatted.Fn()(text) }

func BeginDragDropSource(flags ImGuiDragDropFlags /* = 0 */) bool {
	return imGui_BeginDragDropSource.Fn()(flags) != 0
}
func SetDragDropPayload(type_ cc.String, data unsafe.Pointer, sz uint64, cond ImGuiCond /* = 0 */) bool {
	return imGui_SetDragDropPayload.Fn()(type_, data, sz, cond) != 0
}
func EndDragDropSource()        { imGui_EndDragDropSource.Fn()() }
func BeginDragDropTarget() bool { return imGui_BeginDragDropTarget.Fn()() != 0 }
func AcceptDragDropPayload(type_ cc.String, flags ImGuiDragDropFlags /* = 0 */) *ImGuiPayload {
	return imGui_AcceptDragDropPayload.Fn()(type_, flags)
}
func EndDragDropTarget()                       { imGui_EndDragDropTarget.Fn()() }
func GetDragDropPayload() *ImGuiPayload        { return imGui_GetDragDropPayload.Fn()() }
func BeginDisabled(disabled bool /* = true */) { imGui_BeginDisabled.Fn()(cc.MakeBool(disabled)) }
func EndDisabled()                             { imGui_EndDisabled.Fn()() }

func PushClipRect(clip_rect_min, clip_rect_max ImVec2, intersect_with_current_clip_rect bool) {
	imGui_PushClipRect.Fn()(clip_rect_min, clip_rect_max, cc.MakeBool(intersect_with_current_clip_rect))
}
func PopClipRect()                                { imGui_PopClipRect.Fn()() }
func SetItemDefaultFocus()                        { imGui_SetItemDefaultFocus.Fn()() }
func SetKeyboardFocusHere(offset int32 /* = 0 */) { imGui_SetKeyboardFocusHere.Fn()(offset) }
func SetNavCursorVisible(visible bool)            { imGui_SetNavCursorVisible.Fn()(cc.MakeBool(visible)) }
func SetNextItemAllowOverlap()                    { imGui_SetNextItemAllowOverlap.Fn()() }
func IsItemHovered(flags ImGuiHoveredFlags /* = 0 */) bool {
	return imGui_IsItemHovered.Fn()(flags) != 0
}
func IsItemActive() bool  { return imGui_IsItemActive.Fn()() != 0 }
func IsItemFocused() bool { return imGui_IsItemFocused.Fn()() != 0 }
func IsItemClicked(mouse_button ImGuiMouseButton /* = 0 */) bool {
	return imGui_IsItemClicked.Fn()(mouse_button) != 0
}
func IsItemVisible() bool              { return imGui_IsItemVisible.Fn()() != 0 }
func IsItemEdited() bool               { return imGui_IsItemEdited.Fn()() != 0 }
func IsItemActivated() bool            { return imGui_IsItemActivated.Fn()() != 0 }
func IsItemDeactivated() bool          { return imGui_IsItemDeactivated.Fn()() != 0 }
func IsItemDeactivatedAfterEdit() bool { return imGui_IsItemDeactivatedAfterEdit.Fn()() != 0 }
func IsItemToggledOpen() bool          { return imGui_IsItemToggledOpen.Fn()() != 0 }
func IsAnyItemHovered() bool           { return imGui_IsAnyItemHovered.Fn()() != 0 }
func IsAnyItemActive() bool            { return imGui_IsAnyItemActive.Fn()() != 0 }
func IsAnyItemFocused() bool           { return imGui_IsAnyItemFocused.Fn()() != 0 }
func GetItemID() ImGuiID               { return imGui_GetItemID.Fn()() }
func GetItemRectMin() ImVec2           { return imGui_GetItemRectMin.Fn()() }
func GetItemRectMax() ImVec2           { return imGui_GetItemRectMax.Fn()() }
func GetItemRectSize() ImVec2          { return imGui_GetItemRectSize.Fn()() }

func GetMainViewport() *ImGuiViewport { return imGui_GetMainViewport.Fn()() }
func GetBackgroundDrawList(viewport *ImGuiViewport /* = NULL */) *ImDrawList {
	return imGui_GetBackgroundDrawList.Fn()(viewport)
}
func GetForegroundDrawList(viewport *ImGuiViewport /* = NULL */) *ImDrawList {
	return imGui_GetForegroundDrawList.Fn()(viewport)
}
func IsRectVisibleBySize(size ImVec2) bool { return imGui_IsRectVisibleBySize.Fn()(size) != 0 }
func IsRectVisible(rect_min, rect_max ImVec2) bool {
	return imGui_IsRectVisible.Fn()(rect_min, rect_max) != 0
}
func GetTime() float64                             { return imGui_GetTime.Fn()() }
func GetFrameCount() int32                         { return imGui_GetFrameCount.Fn()() }
func GetDrawListSharedData() *ImDrawListSharedData { return imGui_GetDrawListSharedData.Fn()() }
func GetStyleColorName(idx ImGuiCol) string        { return imGui_GetStyleColorName.Fn()(idx).String() }
func SetStateStorage(storage *ImGuiStorage)        { imGui_SetStateStorage.Fn()(storage) }
func GetStateStorage() *ImGuiStorage               { return imGui_GetStateStorage.Fn()() }
func CalcTextSize(text cc.String, text_end cc.String, /* = NULL */
	hide_text_after_double_hash bool /* = false */, wrap_width float32 /* = -1.0f */) ImVec2 {
	return imGui_CalcTextSize.Fn()(text, text_end, cc.MakeBool(hide_text_after_double_hash), wrap_width)
}
func ColorConvertU32ToFloat4(in ImU32) ImVec4 { return imGui_ColorConvertU32ToFloat4.Fn()(in) }
func ColorConvertFloat4ToU32(in ImVec4) ImU32 { return imGui_ColorConvertFloat4ToU32.Fn()(in) }
func ColorConvertRGBtoHSV(r, g, b float32) (h, s, v float32) {
	imGui_ColorConvertRGBtoHSV.Fn()(r, g, b, &h, &s, &v)
	return
}
func ColorConvertHSVtoRGB(h, s, v float32) (r, g, b float32) {
	imGui_ColorConvertHSVtoRGB.Fn()(h, s, v, &r, &g, &b)
	return
}

func IsKeyDown(key ImGuiKey) bool { return imGui_IsKeyDown.Fn()(key) != 0 }
func IsKeyPressed(key ImGuiKey, repeat bool /* = true */) bool {
	return imGui_IsKeyPressed.Fn()(key, cc.MakeBool(repeat)) != 0
}
func IsKeyReleased(key ImGuiKey) bool { return imGui_IsKeyReleased.Fn()(key) != 0 }
func IsKeyChordPressed(key_chord ImGuiKeyChord) bool {
	return imGui_IsKeyChordPressed.Fn()(key_chord) != 0
}
func GetKeyPressedAmount(key ImGuiKey, repeat_delay, rate float32) int32 {
	return imGui_GetKeyPressedAmount.Fn()(key, repeat_delay, rate)
}
func GetKeyName(key ImGuiKey) string { return imGui_GetKeyName.Fn()(key).String() }
func SetNextFrameWantCaptureKeyboard(want_capture_keyboard bool) {
	imGui_SetNextFrameWantCaptureKeyboard.Fn()(cc.MakeBool(want_capture_keyboard))
}
func Shortcut(key_chord ImGuiKeyChord, flags ImGuiInputFlags /* = 0 */) bool {
	return imGui_Shortcut.Fn()(key_chord, flags) != 0
}
func SetNextItemShortcut(key_chord ImGuiKeyChord, flags ImGuiInputFlags /* = 0 */) {
	imGui_SetNextItemShortcut.Fn()(key_chord, flags)
}
func SetItemKeyOwner(key ImGuiKey) { imGui_SetItemKeyOwner.Fn()(key) }

func IsMouseDown(button ImGuiMouseButton) bool { return imGui_IsMouseDown.Fn()(button) != 0 }
func IsMouseClicked(button ImGuiMouseButton, repeat bool /* = false */) bool {
	return imGui_IsMouseClicked.Fn()(button, cc.MakeBool(repeat)) != 0
}
func IsMouseReleased(button ImGuiMouseButton) bool { return imGui_IsMouseReleased.Fn()(button) != 0 }
func IsMouseDoubleClicked(button ImGuiMouseButton) bool {
	return imGui_IsMouseDoubleClicked.Fn()(button) != 0
}
func IsMouseReleasedWithDelay(button ImGuiMouseButton, delay float32) bool {
	return imGui_IsMouseReleasedWithDelay.Fn()(button, delay) != 0
}
func GetMouseClickedCount(button ImGuiMouseButton) int32 {
	return imGui_GetMouseClickedCount.Fn()(button)
}
func IsMouseHoveringRect(r_min, r_max ImVec2, clip bool /* = true */) bool {
	return imGui_IsMouseHoveringRect.Fn()(r_min, r_max, cc.MakeBool(clip)) != 0
}
func IsMousePosValid(mouse_pos *ImVec2 /* = NULL */) bool {
	return imGui_IsMousePosValid.Fn()(mouse_pos) != 0
}
func IsAnyMouseDown() bool                     { return imGui_IsAnyMouseDown.Fn()() != 0 }
func GetMousePos() ImVec2                      { return imGui_GetMousePos.Fn()() }
func GetMousePosOnOpeningCurrentPopup() ImVec2 { return imGui_GetMousePosOnOpeningCurrentPopup.Fn()() }
func IsMouseDragging(button ImGuiMouseButton, lock_threshold float32 /* = -1.0f */) bool {
	return imGui_IsMouseDragging.Fn()(button, lock_threshold) != 0
}
func GetMouseDragDelta(button ImGuiMouseButton /* = 0 */, lock_threshold float32 /* = -1.0f */) ImVec2 {
	return imGui_GetMouseDragDelta.Fn()(button, lock_threshold)
}
func ResetMouseDragDelta(button ImGuiMouseButton /* = 0 */) { imGui_ResetMouseDragDelta.Fn()(button) }
func GetMouseCursor() ImGuiMouseCursor                      { return imGui_GetMouseCursor.Fn()() }
func SetMouseCursor(cursor_type ImGuiMouseCursor)           { imGui_SetMouseCursor.Fn()(cursor_type) }
func SetNextFrameWantCaptureMouse(want_capture_mouse bool) {
	imGui_SetNextFrameWantCaptureMouse.Fn()(cc.MakeBool(want_capture_mouse))
}
func GetClipboardText() string        { return imGui_GetClipboardText.Fn()().String() }
func SetClipboardText(text cc.String) { imGui_SetClipboardText.Fn()(text) }

func LoadIniSettingsFromDisk(ini_filename string) {
	ini := cc.NewString(ini_filename)
	defer ini.Free()
	imGui_LoadIniSettingsFromDisk.Fn()(ini)
}
func LoadIniSettingsFromMemory(data []byte) {
	imGui_LoadIniSettingsFromMemory.Fn()(&data[0], uint64(len(data)))
}
func SaveIniSettingsToDisk(ini_filename string) {
	ini := cc.NewString(ini_filename)
	defer ini.Free()
	imGui_SaveIniSettingsToDisk.Fn()(ini)
}
func SaveIniSettingsToMemory() []byte {
	size := uint64(0)
	p := imGui_SaveIniSettingsToMemory.Fn()(&size)
	return unsafe.Slice(p, size)
}
func DebugTextEncoding(text cc.String)  { imGui_DebugTextEncoding.Fn()(text) }
func DebugFlashStyleColor(idx ImGuiCol) { imGui_DebugFlashStyleColor.Fn()(idx) }
func DebugStartItemPicker()             { imGui_DebugStartItemPicker.Fn()() }
func DebugCheckVersionAndDataLayout(version_str cc.String, sz_io, sz_style, sz_vec2, sz_vec4, sz_drawvert, sz_drawidx cc.SizeT) bool {
	return imGui_DebugCheckVersionAndDataLayout.Fn()(version_str, sz_io, sz_style, sz_vec2, sz_vec4, sz_drawvert, sz_drawidx) != 0
}
func DebugLogf(fmt cc.String, args ...interface{}) { imGui_DebugLog.FnVa()(fmt, args...) }
func DebugLog(text cc.String)                      { imGui_DebugLogUnformatted.Fn()(text) }
func SetAllocatorFunctions(alloc_func cc.Func, /* void* (*ImGuiMemAllocFunc)(size_t size, void* user_data)*/
	free_func cc.Func /* void (*ImGuiMemFreeFunc)(void* ptr, void* user_data)*/, user_data unsafe.Pointer /* = NULL */) {
	imGui_SetAllocatorFunctions.Fn()(alloc_func, free_func, user_data)
}
func MemAlloc(size cc.SizeT) unsafe.Pointer { return imGui_MemAlloc.Fn()(size) }
func MemFree(ptr unsafe.Pointer)            { imGui_MemFree.Fn()(ptr) }
func UpdatePlatformWindows()                { imGui_UpdatePlatformWindows.Fn()() }
func RenderPlatformWindowsDefault(platform_render_arg unsafe.Pointer /* = NULL */, renderer_render_arg unsafe.Pointer /* = NULL */) {
	imGui_RenderPlatformWindowsDefault.Fn()(platform_render_arg, renderer_render_arg)
}
func DestroyPlatformWindows()                    { imGui_DestroyPlatformWindows.Fn()() }
func FindViewportByID(id ImGuiID) *ImGuiViewport { return imGui_FindViewportByID.Fn()(id) }
func FindViewportByPlatformHandle(platform_handle unsafe.Pointer) *ImGuiViewport {
	return imGui_FindViewportByPlatformHandle.Fn()(platform_handle)
}

func (s *ImGuiStyle) ScaleAllSizes(scaleFactor float32) {
	imGuiStyle_ScaleAllSizes.Fn()(s, scaleFactor)
}

func (io *ImGuiIO) AddKeyEvent(key ImGuiKey, down bool) {
	imGuiIO_AddKeyEvent.Fn()(io, key, cc.MakeBool(down))
}
func (io *ImGuiIO) AddKeyAnalogEvent(key ImGuiKey, down bool, v float32) {
	imGuiIO_AddKeyAnalogEvent.Fn()(io, key, cc.MakeBool(down), v)
}
func (io *ImGuiIO) AddMousePosEvent(x, y float32) { imGuiIO_AddMousePosEvent.Fn()(io, x, y) }
func (io *ImGuiIO) AddMouseButtonEvent(btn int32, down bool) {
	imGuiIO_AddMouseButtonEvent.Fn()(io, btn, cc.MakeBool(down))
}
func (io *ImGuiIO) AddMouseWheelEvent(wheelX, wheelY float32) {
	imGuiIO_AddMouseWheelEvent.Fn()(io, wheelX, wheelY)
}
func (io *ImGuiIO) AddMouseSourceEvent(source ImGuiMouseSource) {
	imGuiIO_AddMouseSourceEvent.Fn()(io, source)
}
func (io *ImGuiIO) AddMouseViewportEvent(id ImGuiID)   { imGuiIO_AddMouseViewportEvent.Fn()(io, id) }
func (io *ImGuiIO) AddFocusEvent(focused bool)         { imGuiIO_AddFocusEvent.Fn()(io, cc.MakeBool(focused)) }
func (io *ImGuiIO) AddInputCharacter(c uint32)         { imGuiIO_AddInputCharacter.Fn()(io, c) }
func (io *ImGuiIO) AddInputCharacterUTF16(c ImWchar16) { imGuiIO_AddInputCharacterUTF16.Fn()(io, c) }
func (io *ImGuiIO) AddInputCharactersUTF8(c cc.String) { imGuiIO_AddInputCharactersUTF8.Fn()(io, c) }
func (io *ImGuiIO) SetKeyEventNativeData(key ImGuiKey, keycode, scancode, legacy_index /*-1*/ int32) {
	imGuiIO_SetKeyEventNativeData.Fn()(io, key, keycode, scancode, legacy_index)
}
func (io *ImGuiIO) SetAppAcceptingEvents(acceptingEvents bool) {
	imGuiIO_SetAppAcceptingEvents.Fn()(io, cc.MakeBool(acceptingEvents))
}
func (io *ImGuiIO) ClearEventsQueue()     { imGuiIO_ClearEventsQueue.Fn()(io) }
func (io *ImGuiIO) ClearInputKeys()       { imGuiIO_ClearInputKeys.Fn()(io) }
func (io *ImGuiIO) ClearInputMouse()      { imGuiIO_ClearInputMouse.Fn()(io) }
func (io *ImGuiIO) ClearInputCharacters() { imGuiIO_ClearInputCharacters.Fn()(io) }

func (itcd *ImGuiInputTextCallbackData) DeleteChars(pos, bytes_count int32) {
	imGuiInputTextCallbackData_DeleteChars.Fn()(itcd, pos, bytes_count)
}
func (itcd *ImGuiInputTextCallbackData) InsertChars(pos int32, text cc.String, text_end cc.String /* = NULL */) {
	imGuiInputTextCallbackData_InsertChars.Fn()(itcd, pos, text, text_end)
}
func (itcd *ImGuiInputTextCallbackData) InsertCharsStr(pos int32, text string) {
	t := cc.NewString(text)
	defer t.Free()
	imGuiInputTextCallbackData_InsertChars.Fn()(itcd, pos, t, 0)
}
func (itcd *ImGuiInputTextCallbackData) SelectAll() { imGuiInputTextCallbackData_SelectAll.Fn()(itcd) }
func (itcd *ImGuiInputTextCallbackData) ClearSelection() {
	imGuiInputTextCallbackData_ClearSelection.Fn()(itcd)
}
func (itcd *ImGuiInputTextCallbackData) HasSelection() bool {
	return imGuiInputTextCallbackData_HasSelection.Fn()(itcd) != 0
}

func (pl *ImGuiPayload) Clear() { imGuiPayload_Clear.Fn()(pl) }
func (pl *ImGuiPayload) IsDataType(typ string) bool {
	t := cc.NewString(typ)
	defer t.Free()
	return imGuiPayload_IsDataType.Fn()(pl, t) != 0
}
func (pl *ImGuiPayload) IsPreview() bool  { return imGuiPayload_IsPreview.Fn()(pl) != 0 }
func (pl *ImGuiPayload) IsDelivery() bool { return imGuiPayload_IsDelivery.Fn()(pl) != 0 }

func (tr *ImGuiTextRange) Empty() bool {
	return imGuiTextFilter_ImGuiTextRange_empty.Fn()(tr) != 0
}
func (tr *ImGuiTextRange) Split(separator int32) (out ImVector[ImGuiTextRange]) {
	imGuiTextFilter_ImGuiTextRange_split.Fn()(tr, separator, &out)
	return
}
func (tf *ImGuiTextFilter) Draw(label cc.String /* = "Filter (inc,-exc)" */, width float32 /* = 0.0f */) bool {
	return imGuiTextFilter_Draw.Fn()(tf, label, width) != 0
}
func (tf *ImGuiTextFilter) PassFilter(text cc.String, text_end cc.String /* = NULL */) bool {
	return imGuiTextFilter_PassFilter.Fn()(tf, text, text_end) != 0
}
func (tf *ImGuiTextFilter) Build()         { imGuiTextFilter_Build.Fn()(tf) }
func (tf *ImGuiTextFilter) Clear()         { imGuiTextFilter_Clear.Fn()(tf) }
func (tf *ImGuiTextFilter) IsActive() bool { return imGuiTextFilter_IsActive.Fn()(tf) != 0 }

func (tb *ImGuiTextBuffer) Begin() cc.String       { return imGuiTextBuffer_begin.Fn()(tb) }
func (tb *ImGuiTextBuffer) End() cc.String         { return imGuiTextBuffer_end.Fn()(tb) }
func (tb *ImGuiTextBuffer) Size() int32            { return imGuiTextBuffer_size.Fn()(tb) }
func (tb *ImGuiTextBuffer) Empty() bool            { return imGuiTextBuffer_empty.Fn()(tb) != 0 }
func (tb *ImGuiTextBuffer) Clear()                 { imGuiTextBuffer_clear.Fn()(tb) }
func (tb *ImGuiTextBuffer) Resize(size int32)      { imGuiTextBuffer_resize.Fn()(tb, size) }
func (tb *ImGuiTextBuffer) Reserve(capacity int32) { imGuiTextBuffer_reserve.Fn()(tb, capacity) }
func (tb *ImGuiTextBuffer) CStr() cc.String        { return imGuiTextBuffer_c_str.Fn()(tb) }
func (tb *ImGuiTextBuffer) Append(str cc.String, str_end cc.String /* = NULL */) {
	imGuiTextBuffer_append.Fn()(tb, str, str_end)
}
func (tb *ImGuiTextBuffer) Appendf(fmt cc.String, args ...interface{}) {
	imGuiTextBuffer_appendf.FnVa()(tb, fmt, args...)
}
func (st *ImGuiStorage) Clear() { imGuiStorage_Clear.Fn()(st) }
func (st *ImGuiStorage) GetInt(key ImGuiID, default_val int32 /* = 0 */) int32 {
	return imGuiStorage_GetInt.Fn()(st, key, default_val)
}
func (st *ImGuiStorage) SetInt(key ImGuiID, val int32) { imGuiStorage_SetInt.Fn()(st, key, val) }
func (st *ImGuiStorage) GetBool(key ImGuiID, default_val bool /* = false */) bool {
	return imGuiStorage_GetBool.Fn()(st, key, cc.MakeBool(default_val)) != 0
}
func (st *ImGuiStorage) SetBool(key ImGuiID, val bool) {
	imGuiStorage_SetBool.Fn()(st, key, cc.MakeBool(val))
}
func (st *ImGuiStorage) GetFloat(key ImGuiID, default_val float32 /* = 0.0f */) float32 {
	return imGuiStorage_GetFloat.Fn()(st, key, default_val)
}
func (st *ImGuiStorage) SetFloat(key ImGuiID, val float32) { imGuiStorage_SetFloat.Fn()(st, key, val) }
func (st *ImGuiStorage) GetPointer(key ImGuiID) unsafe.Pointer {
	return imGuiStorage_GetVoidPtr.Fn()(st, key)
}
func (st *ImGuiStorage) SetPointer(key ImGuiID, val unsafe.Pointer) {
	imGuiStorage_SetVoidPtr.Fn()(st, key, val)
}
func (st *ImGuiStorage) GetIntRef(key ImGuiID, default_val int32 /* = 0 */) *int32 {
	return imGuiStorage_GetIntRef.Fn()(st, key, default_val)
}
func (st *ImGuiStorage) GetBoolRef(key ImGuiID, default_val bool /* = false */) *bool {
	return (*bool)(uptr(imGuiStorage_GetBoolRef.Fn()(st, key, cc.MakeBool(default_val))))
}
func (st *ImGuiStorage) GetFloatRef(key ImGuiID, default_val float32 /* = 0.0f */) *float32 {
	return imGuiStorage_GetFloatRef.Fn()(st, key, default_val)
}
func (st *ImGuiStorage) GetPointerRef(key ImGuiID, default_val unsafe.Pointer /* = NULL */) **unsafe.Pointer {
	return imGuiStorage_GetVoidPtrRef.Fn()(st, key, default_val)
}
func (st *ImGuiStorage) BuildSortByKey()     { imGuiStorage_BuildSortByKey.Fn()(st) }
func (st *ImGuiStorage) SetAllInt(val int32) { imGuiStorage_SetAllInt.Fn()(st, val) }

func (lc *ImGuiListClipper) Begin(items_count int32, items_height float32 /* = -1.0f */) {
	imGuiListClipper_Begin.Fn()(lc, items_count, items_height)
}
func (lc *ImGuiListClipper) End()       { imGuiListClipper_End.Fn()(lc) }
func (lc *ImGuiListClipper) Step() bool { return imGuiListClipper_Step.Fn()(lc) != 0 }
func (lc *ImGuiListClipper) IncludeItemByIndex(item_index int32) {
	imGuiListClipper_IncludeItemByIndex.Fn()(lc, item_index)
}
func (lc *ImGuiListClipper) IncludeItemsByIndex(item_begin, item_end int32) {
	imGuiListClipper_IncludeItemsByIndex.Fn()(lc, item_begin, item_end)
}
func (lc *ImGuiListClipper) SeekCursorForItem(item_index int32) {
	imGuiListClipper_SeekCursorForItem.Fn()(lc, item_index)
}
func (lc *ImGuiListClipper) IncludeRangeByIndices(item_begin, item_end int32) {
	imGuiListClipper_IncludeRangeByIndices.Fn()(lc, item_begin, item_end)
}

const (
	IM_COL32_R_SHIFT = 0
	IM_COL32_G_SHIFT = 8
	IM_COL32_B_SHIFT = 16
	IM_COL32_A_SHIFT = 24
	IM_COL32_A_MASK  = 0xFF000000
)

func ImCol32(R, G, B, A ImU32) ImU32 {
	return (((ImU32)(A) << IM_COL32_A_SHIFT) | ((ImU32)(B) << IM_COL32_B_SHIFT) | ((ImU32)(G) << IM_COL32_G_SHIFT) | ((ImU32)(R) << IM_COL32_R_SHIFT))
}

func (c *ImColor) SetHSV(h, s, v float32, a float32 /* = 1.0f */) { imColor_SetHSV.Fn()(c, h, s, v, a) }
func ImColorHSV(h, s, v float32, a float32 /* = 1.0f */) ImColor  { return imColor_HSV.Fn()(h, s, v, a) }

func (sbs *ImGuiSelectionBasicStorage) ApplyRequests(ms_io *ImGuiMultiSelectIO) {
	imGuiSelectionBasicStorage_ApplyRequests.Fn()(sbs, ms_io)
}
func (sbs *ImGuiSelectionBasicStorage) Contains(id ImGuiID) bool {
	return imGuiSelectionBasicStorage_Contains.Fn()(sbs, id) != 0
}
func (sbs *ImGuiSelectionBasicStorage) Clear() { imGuiSelectionBasicStorage_Clear.Fn()(sbs) }
func (sbs *ImGuiSelectionBasicStorage) Swap(r *ImGuiSelectionBasicStorage) {
	imGuiSelectionBasicStorage_Swap.Fn()(sbs, r)
}
func (sbs *ImGuiSelectionBasicStorage) SetItemSelected(id ImGuiID, selected bool) {
	imGuiSelectionBasicStorage_SetItemSelected.Fn()(sbs, id, cc.MakeBool(selected))
}
func (sbs *ImGuiSelectionBasicStorage) GetNextSelectedItem(opaque_it *unsafe.Pointer, out_id *ImGuiID) bool {
	return imGuiSelectionBasicStorage_GetNextSelectedItem.Fn()(sbs, opaque_it, out_id) != 0
}
func (sbs *ImGuiSelectionBasicStorage) GetStorageIdFromIndex(idx int32) ImGuiID {
	return imGuiSelectionBasicStorage_GetStorageIdFromIndex.Fn()(sbs, idx)
}
func (ses *ImGuiSelectionExternalStorage) ApplyRequests(ms_io *ImGuiMultiSelectIO) {
	imGuiSelectionExternalStorage_ApplyRequests.Fn()(ses, ms_io)
}

func (dc *ImDrawCmd) GetTexID() ImTextureID      { return imDrawCmd_GetTexID.Fn()(dc) }
func (dls *ImDrawListSplitter) Clear()           { imDrawListSplitter_Clear.Fn()(dls) }
func (dls *ImDrawListSplitter) ClearFreeMemory() { imDrawListSplitter_ClearFreeMemory.Fn()(dls) }
func (dls *ImDrawListSplitter) Split(draw_list *ImDrawList, count int32) {
	imDrawListSplitter_Split.Fn()(dls, draw_list, count)
}
func (dls *ImDrawListSplitter) Merge(draw_list *ImDrawList) {
	imDrawListSplitter_Merge.Fn()(dls, draw_list)
}
func (dls *ImDrawListSplitter) SetCurrentChannel(draw_list *ImDrawList, channel_idx int32) {
	imDrawListSplitter_SetCurrentChannel.Fn()(dls, draw_list, channel_idx)
}

func (dl *ImDrawList) PushClipRect(rect_min, rect_max ImVec2, intersect_with_current_clip_rect bool /* = false */) {
	imDrawList_PushClipRect.Fn()(dl, rect_min, rect_max, cc.MakeBool(intersect_with_current_clip_rect))
}
func (dl *ImDrawList) PushClipRectFullScreen()          { imDrawList_PushClipRectFullScreen.Fn()(dl) }
func (dl *ImDrawList) PopClipRect()                     { imDrawList_PopClipRect.Fn()(dl) }
func (dl *ImDrawList) PushTexture(tex_ref ImTextureRef) { imDrawList_PushTexture.Fn()(dl, tex_ref) }
func (dl *ImDrawList) PopTexture()                      { imDrawList_PopTexture.Fn()(dl) }
func (dl *ImDrawList) GetClipRectMin() ImVec2           { return imDrawList_GetClipRectMin.Fn()(dl) }
func (dl *ImDrawList) GetClipRectMax() ImVec2           { return imDrawList_GetClipRectMax.Fn()(dl) }
func (dl *ImDrawList) AddLine(p1, p2 ImVec2, col ImU32, thickness float32 /* = 1.0f */) {
	imDrawList_AddLine.Fn()(dl, p1, p2, col, thickness)
}
func (dl *ImDrawList) AddRect(p_min, p_max ImVec2, col ImU32, rounding float32, /* = 0.0f */
	flags ImDrawFlags /* = 0 */, thickness float32 /* = 1.0f */) {
	imDrawList_AddRect.Fn()(dl, p_min, p_max, col, rounding, flags, thickness)
}
func (dl *ImDrawList) AddRectFilled(p_min, p_max ImVec2, col ImU32, rounding float32, /* = 0.0f */
	flags ImDrawFlags /* = 0 */) {
	imDrawList_AddRectFilled.Fn()(dl, p_min, p_max, col, rounding, flags)
}
func (dl *ImDrawList) AddRectFilledMultiColor(p_min, p_max ImVec2,
	upr_left, upr_right, bot_right, bot_left ImU32) {
	imDrawList_AddRectFilledMultiColor.Fn()(dl, p_min, p_max, upr_left, upr_right, bot_right, bot_left)
}
func (dl *ImDrawList) AddQuad(p1, p2, p3, p4 ImVec2, col ImU32, thickness float32 /* = 1.0f */) {
	imDrawList_AddQuad.Fn()(dl, p1, p2, p3, p4, col, thickness)
}
func (dl *ImDrawList) AddQuadFilled(p1, p2, p3, p4 ImVec2, col ImU32) {
	imDrawList_AddQuadFilled.Fn()(dl, p1, p2, p3, p4, col)
}
func (dl *ImDrawList) AddTriangle(p1, p2, p3 ImVec2, col ImU32, thickness float32 /* = 1.0f */) {
	imDrawList_AddTriangle.Fn()(dl, p1, p2, p3, col, thickness)
}
func (dl *ImDrawList) AddTriangleFilled(p1, p2, p3 ImVec2, col ImU32) {
	imDrawList_AddTriangleFilled.Fn()(dl, p1, p2, p3, col)
}
func (dl *ImDrawList) AddCircle(center ImVec2, radius float32, col ImU32, num_segments int32 /* = 0 */, thickness float32 /* = 1.0f */) {
	imDrawList_AddCircle.Fn()(dl, center, radius, col, num_segments, thickness)
}
func (dl *ImDrawList) AddCircleFilled(center ImVec2, radius float32, col ImU32, num_segments int32 /* = 0 */) {
	imDrawList_AddCircleFilled.Fn()(dl, center, radius, col, num_segments)
}
func (dl *ImDrawList) AddNgon(center ImVec2, radius float32, col ImU32, num_segments int32, thickness float32 /* = 1.0f */) {
	imDrawList_AddNgon.Fn()(dl, center, radius, col, num_segments, thickness)
}
func (dl *ImDrawList) AddNgonFilled(center ImVec2, radius float32, col ImU32, num_segments int32) {
	imDrawList_AddNgonFilled.Fn()(dl, center, radius, col, num_segments)
}
func (dl *ImDrawList) AddEllipse(center ImVec2, radius ImVec2, col ImU32,
	rot float32 /* = 0.0f */, num_segments int32 /* = 0 */, thickness float32 /* = 1.0f */) {
	imDrawList_AddEllipse.Fn()(dl, center, radius, col, rot, num_segments, thickness)
}
func (dl *ImDrawList) AddEllipseFilled(center ImVec2, radius ImVec2, col ImU32,
	rot float32 /* = 0.0f */, num_segments int32 /* = 0 */) {
	imDrawList_AddEllipseFilled.Fn()(dl, center, radius, col, rot, num_segments)
}
func (dl *ImDrawList) AddText(pos ImVec2, col ImU32, text_begin cc.String, text_end cc.String /* = NULL */) {
	imDrawList_AddText.Fn()(dl, pos, col, text_begin, text_end)
}
func (dl *ImDrawList) AddTextImFontPtr(font *ImFont, font_size float32, pos ImVec2, col ImU32,
	text_begin cc.String, text_end cc.String /* = NULL */, wrap_width float32 /* = 0.0f */, cpu_fine_clip_rect *ImVec4 /* = NULL */) {
	imDrawList_AddTextImFontPtr.Fn()(dl, font, font_size, pos, col, text_begin, text_end, wrap_width, cpu_fine_clip_rect)
}
func (dl *ImDrawList) AddBezierCubic(p1, p2, p3, p4 ImVec2, col ImU32, thickness float32, num_segments int32 /* = 0 */) {
	imDrawList_AddBezierCubic.Fn()(dl, p1, p2, p3, p4, col, thickness, num_segments)
}
func (dl *ImDrawList) AddBezierQuadratic(p1, p2, p3 ImVec2, col ImU32, thickness float32, num_segments int32 /* = 0 */) {
	imDrawList_AddBezierQuadratic.Fn()(dl, p1, p2, p3, col, thickness, num_segments)
}

func (dl *ImDrawList) AddPolyline(points []ImVec2, col ImU32, flags ImDrawFlags, thickness float32) {
	imDrawList_AddPolyline.Fn()(dl, &points[0], int32(len(points)), col, flags, thickness)
}
func (dl *ImDrawList) AddConvexPolyFilled(points []ImVec2, col ImU32) {
	imDrawList_AddConvexPolyFilled.Fn()(dl, &points[0], int32(len(points)), col)
}
func (dl *ImDrawList) AddConcavePolyFilled(points []ImVec2, col ImU32) {
	imDrawList_AddConcavePolyFilled.Fn()(dl, &points[0], int32(len(points)), col)
}
func (dl *ImDrawList) AddImage(tex_ref ImTextureRef, p_min, p_max ImVec2, uv_min ImVec2, /* = ImVec2(0, 0) */
	uv_max ImVec2 /* = ImVec2(1, 1) */, col ImU32 /* = IM_COL32_WHITE */) {
	imDrawList_AddImage.Fn()(dl, tex_ref, p_min, p_max, uv_min, uv_max, col)
}
func (dl *ImDrawList) AddImageQuad(tex_ref ImTextureRef, p1, p2, p3, p4 ImVec2, uv1 ImVec2, /* = ImVec2(0, 0) */
	uv2 ImVec2 /* = ImVec2(1, 0) */, uv3 ImVec2 /* = ImVec2(1, 1) */, uv4 ImVec2 /* = ImVec2(0, 1) */, col ImU32 /* = IM_COL32_WHITE */) {
	imDrawList_AddImageQuad.Fn()(dl, tex_ref, p1, p2, p3, p4, uv1, uv2, uv3, uv4, col)
}
func (dl *ImDrawList) AddImageRounded(tex_ref ImTextureRef, p_min, p_max ImVec2, uv_min, uv_max ImVec2,
	col ImU32, rounding float32, flags ImDrawFlags /* = 0 */) {
	imDrawList_AddImageRounded.Fn()(dl, tex_ref, p_min, p_max, uv_min, uv_max, col, rounding, flags)
}
func (dl *ImDrawList) PathClear()            { imDrawList_PathClear.Fn()(dl) }
func (dl *ImDrawList) PathLineTo(pos ImVec2) { imDrawList_PathLineTo.Fn()(dl, pos) }
func (dl *ImDrawList) PathLineToMergeDuplicate(pos ImVec2) {
	imDrawList_PathLineToMergeDuplicate.Fn()(dl, pos)
}
func (dl *ImDrawList) PathFillConvex(col ImU32)  { imDrawList_PathFillConvex.Fn()(dl, col) }
func (dl *ImDrawList) PathFillConcave(col ImU32) { imDrawList_PathFillConcave.Fn()(dl, col) }
func (dl *ImDrawList) PathStroke(col ImU32, flags ImDrawFlags /* = 0 */, thickness float32 /* = 1.0f */) {
	imDrawList_PathStroke.Fn()(dl, col, flags, thickness)
}
func (dl *ImDrawList) PathArcTo(center ImVec2, radius float32, a_min, a_max float32, num_segments int32 /* = 0 */) {
	imDrawList_PathArcTo.Fn()(dl, center, radius, a_min, a_max, num_segments)
}
func (dl *ImDrawList) PathArcToFast(center ImVec2, radius float32, a_min_of_12, a_max_of_12 int32) {
	imDrawList_PathArcToFast.Fn()(dl, center, radius, a_min_of_12, a_max_of_12)
}
func (dl *ImDrawList) PathEllipticalArcTo(center ImVec2, radius ImVec2, rot, a_min, a_max float32, num_segments int32 /* = 0 */) {
	imDrawList_PathEllipticalArcTo.Fn()(dl, center, radius, rot, a_min, a_max, num_segments)
}
func (dl *ImDrawList) PathBezierCubicCurveTo(p2, p3, p4 ImVec2, num_segments int32 /* = 0 */) {
	imDrawList_PathBezierCubicCurveTo.Fn()(dl, p2, p3, p4, num_segments)
}
func (dl *ImDrawList) PathBezierQuadraticCurveTo(p2, p3 ImVec2, num_segments int32 /* = 0 */) {
	imDrawList_PathBezierQuadraticCurveTo.Fn()(dl, p2, p3, num_segments)
}
func (dl *ImDrawList) PathRect(rect_min, rect_max ImVec2, rounding float32 /* = 0.0f */, flags ImDrawFlags /* = 0 */) {
	imDrawList_PathRect.Fn()(dl, rect_min, rect_max, rounding, flags)
}

func (dl *ImDrawList) AddCallback(callback cc.Func /* void (*ImDrawCallback)(const ImDrawList* parent_list, const ImDrawCmd* cmd)*/, userdata unsafe.Pointer, userdata_size cc.SizeT /* = 0 */) {
	imDrawList_AddCallback.Fn()(dl, callback, userdata, userdata_size)
}
func (dl *ImDrawList) AddDrawCmd()                { imDrawList_AddDrawCmd.Fn()(dl) }
func (dl *ImDrawList) CloneOutput() *ImDrawList   { return imDrawList_CloneOutput.Fn()(dl) }
func (dl *ImDrawList) ChannelsSplit(count int32)  { imDrawList_ChannelsSplit.Fn()(dl, count) }
func (dl *ImDrawList) ChannelsMerge()             { imDrawList_ChannelsMerge.Fn()(dl) }
func (dl *ImDrawList) ChannelsSetCurrent(n int32) { imDrawList_ChannelsSetCurrent.Fn()(dl, n) }
func (dl *ImDrawList) PrimReserve(idx_count, vtx_count int32) {
	imDrawList_PrimReserve.Fn()(dl, idx_count, vtx_count)
}
func (dl *ImDrawList) PrimUnreserve(idx_count, vtx_count int32) {
	imDrawList_PrimUnreserve.Fn()(dl, idx_count, vtx_count)
}
func (dl *ImDrawList) PrimRect(a, b ImVec2, col ImU32) { imDrawList_PrimRect.Fn()(dl, a, b, col) }
func (dl *ImDrawList) PrimRectUV(a, b ImVec2, uv_a, uv_b ImVec2, col ImU32) {
	imDrawList_PrimRectUV.Fn()(dl, a, b, uv_a, uv_b, col)
}
func (dl *ImDrawList) PrimQuadUV(a, b, c, d ImVec2, uv_a, uv_b, uv_c, uv_d ImVec2, col ImU32) {
	imDrawList_PrimQuadUV.Fn()(dl, a, b, c, d, uv_a, uv_b, uv_c, uv_d, col)
}
func (dl *ImDrawList) PrimWriteVtx(pos, uv ImVec2, col ImU32) {
	imDrawList_PrimWriteVtx.Fn()(dl, pos, uv, col)
}
func (dl *ImDrawList) PrimWriteIdx(idx ImDrawIdx)         { imDrawList_PrimWriteIdx.Fn()(dl, idx) }
func (dl *ImDrawList) PrimVtx(pos, uv ImVec2, col ImU32)  { imDrawList_PrimVtx.Fn()(dl, pos, uv, col) }
func (dl *ImDrawList) PushTextureID(tex_ref ImTextureRef) { imDrawList_PushTextureID.Fn()(dl, tex_ref) }
func (dl *ImDrawList) PopTextureID()                      { imDrawList_PopTextureID.Fn()(dl) }
func (dl *ImDrawList) SetDrawListSharedData(data *ImDrawListSharedData) {
	imDrawList__SetDrawListSharedData.Fn()(dl, data)
}
func (dl *ImDrawList) ResetForNewFrame()               { imDrawList__ResetForNewFrame.Fn()(dl) }
func (dl *ImDrawList) ClearFreeMemory()                { imDrawList__ClearFreeMemory.Fn()(dl) }
func (dl *ImDrawList) PopUnusedDrawCmd()               { imDrawList__PopUnusedDrawCmd.Fn()(dl) }
func (dl *ImDrawList) TryMergeDrawCmds()               { imDrawList__TryMergeDrawCmds.Fn()(dl) }
func (dl *ImDrawList) OnChangedClipRect()              { imDrawList__OnChangedClipRect.Fn()(dl) }
func (dl *ImDrawList) OnChangedTexture()               { imDrawList__OnChangedTexture.Fn()(dl) }
func (dl *ImDrawList) OnChangedVtxOffset()             { imDrawList__OnChangedVtxOffset.Fn()(dl) }
func (dl *ImDrawList) SetTexture(tex_ref ImTextureRef) { imDrawList__SetTexture.Fn()(dl, tex_ref) }
func (dl *ImDrawList) CalcCircleAutoSegmentCount(radius float32) int32 {
	return imDrawList__CalcCircleAutoSegmentCount.Fn()(dl, radius)
}
func (dl *ImDrawList) PathArcToFastEx(center ImVec2, radius float32, a_min_sample, a_max_sample, a_step int32) {
	imDrawList__PathArcToFastEx.Fn()(dl, center, radius, a_min_sample, a_max_sample, a_step)
}
func (dl *ImDrawList) PathArcToN(center ImVec2, radius float32, a_min, a_max float32, num_segments int32) {
	imDrawList__PathArcToN.Fn()(dl, center, radius, a_min, a_max, num_segments)
}

func (dd *ImDrawData) Clear()                            { imDrawData_Clear.Fn()(dd) }
func (dd *ImDrawData) AddDrawList(draw_list *ImDrawList) { imDrawData_AddDrawList.Fn()(dd, draw_list) }
func (dd *ImDrawData) DeIndexAllBuffers()                { imDrawData_DeIndexAllBuffers.Fn()(dd) }
func (dd *ImDrawData) ScaleClipRects(fb_scale ImVec2)    { imDrawData_ScaleClipRects.Fn()(dd, fb_scale) }

func (td *ImTextureData) Create(format ImTextureFormat, w, h int32) {
	imTextureData_Create.Fn()(td, format, w, h)
}
func (td *ImTextureData) DestroyPixels()            { imTextureData_DestroyPixels.Fn()(td) }
func (td *ImTextureData) GetPixels() unsafe.Pointer { return imTextureData_GetPixels.Fn()(td) }
func (td *ImTextureData) GetPixelsAt(x, y int32) unsafe.Pointer {
	return imTextureData_GetPixelsAt.Fn()(td, x, y)
}
func (td *ImTextureData) GetSizeInBytes() int32            { return imTextureData_GetSizeInBytes.Fn()(td) }
func (td *ImTextureData) GetPitch() int32                  { return imTextureData_GetPitch.Fn()(td) }
func (td *ImTextureData) GetTexRef() ImTextureRef          { return imTextureData_GetTexRef.Fn()(td) }
func (td *ImTextureData) GetTexID() ImTextureID            { return imTextureData_GetTexID.Fn()(td) }
func (td *ImTextureData) SetTexID(tex_id ImTextureID)      { imTextureData_SetTexID.Fn()(td, tex_id) }
func (td *ImTextureData) SetStatus(status ImTextureStatus) { imTextureData_SetStatus.Fn()(td, status) }

func (fgrb *ImFontGlyphRangesBuilder) Clear() { imFontGlyphRangesBuilder_Clear.Fn()(fgrb) }
func (fgrb *ImFontGlyphRangesBuilder) GetBit(n uint64) bool {
	return imFontGlyphRangesBuilder_GetBit.Fn()(fgrb, n) != 0
}
func (fgrb *ImFontGlyphRangesBuilder) SetBit(n uint64) {
	imFontGlyphRangesBuilder_SetBit.Fn()(fgrb, n)
}
func (fgrb *ImFontGlyphRangesBuilder) AddChar(c ImWchar) {
	imFontGlyphRangesBuilder_AddChar.Fn()(fgrb, c)
}
func (fgrb *ImFontGlyphRangesBuilder) AddText(text cc.String, text_end cc.String /* = NULL */) {
	imFontGlyphRangesBuilder_AddText.Fn()(fgrb, text, text_end)
}
func (fgrb *ImFontGlyphRangesBuilder) AddRanges(ranges *ImWchar) {
	imFontGlyphRangesBuilder_AddRanges.Fn()(fgrb, ranges)
}
func (fgrb *ImFontGlyphRangesBuilder) BuildRanges(out_ranges *ImVector[ImWchar]) {
	imFontGlyphRangesBuilder_BuildRanges.Fn()(fgrb, out_ranges)
}

func (fa *ImFontAtlas) AddFont(font_cfg *ImFontConfig) *ImFont {
	return imFontAtlas_AddFont.Fn()(fa, font_cfg)
}
func (fa *ImFontAtlas) AddFontDefault(font_cfg *ImFontConfig /* = NULL */) *ImFont {
	return imFontAtlas_AddFontDefault.Fn()(fa, font_cfg)
}
func (fa *ImFontAtlas) AddFontFromFileTTF(filename cc.String, size_pixels float32, /* = 0.0f */
	font_cfg *ImFontConfig /* = NULL */, glyph_ranges *ImWchar /* = NULL */) *ImFont {
	return imFontAtlas_AddFontFromFileTTF.Fn()(fa, filename, size_pixels, font_cfg, glyph_ranges)
}
func (fa *ImFontAtlas) AddFontFromMemoryTTF(fontData []byte, size_pixels float32, /* = 0.0f */
	font_cfg *ImFontConfig /* = NULL */, glyph_ranges *ImWchar /* = NULL */) *ImFont {

	return imFontAtlas_AddFontFromMemoryTTF.Fn()(fa, &fontData[0], int32(len(fontData)), size_pixels, font_cfg, glyph_ranges)
}
func (fa *ImFontAtlas) AddFontFromMemoryCompressedTTF(fontData []byte, size_pixels float32, /* = 0.0f */
	font_cfg *ImFontConfig /* = NULL */, glyph_ranges *ImWchar /* = NULL */) *ImFont {
	return imFontAtlas_AddFontFromMemoryCompressedTTF.Fn()(fa, &fontData[0], int32(len(fontData)), size_pixels, font_cfg, glyph_ranges)
}
func (fa *ImFontAtlas) AddFontFromMemoryCompressedBase85TTF(fontData cc.String, size_pixels float32, /* = 0.0f */
	font_cfg *ImFontConfig /* = NULL */, glyph_ranges *ImWchar /* = NULL */) *ImFont {
	return imFontAtlas_AddFontFromMemoryCompressedBase85TTF.Fn()(fa, fontData, size_pixels, font_cfg, glyph_ranges)
}
func (fa *ImFontAtlas) RemoveFont(font *ImFont) { imFontAtlas_RemoveFont.Fn()(fa, font) }
func (fa *ImFontAtlas) Clear()                  { imFontAtlas_Clear.Fn()(fa) }
func (fa *ImFontAtlas) CompactCache()           { imFontAtlas_CompactCache.Fn()(fa) }
func (fa *ImFontAtlas) SetFontLoader(font_loader *ImFontLoader) {
	imFontAtlas_SetFontLoader.Fn()(fa, font_loader)
}
func (fa *ImFontAtlas) ClearInputData() { imFontAtlas_ClearInputData.Fn()(fa) }
func (fa *ImFontAtlas) ClearFonts()     { imFontAtlas_ClearFonts.Fn()(fa) }
func (fa *ImFontAtlas) ClearTexData()   { imFontAtlas_ClearTexData.Fn()(fa) }

func (fa *ImFontAtlas) Build() bool { return imFontAtlas_Build.Fn()(fa) != 0 }
func (fa *ImFontAtlas) GetTexDataAsAlpha8() []byte {
	pixles, w, h := (*byte)(nil), int32(0), int32(0)
	imFontAtlas_GetTexDataAsAlpha8.Fn()(fa, &pixles, &w, &h, nil)
	return unsafe.Slice(pixles, w*h*1)
}
func (fa *ImFontAtlas) GetTexDataAsRGBA32() []byte {
	pixles, w, h := (*byte)(nil), int32(0), int32(0)
	imFontAtlas_GetTexDataAsRGBA32.Fn()(fa, &pixles, &w, &h, nil)
	return unsafe.Slice(pixles, w*h*1)
}
func (fa *ImFontAtlas) SetTexID(id ImTextureID) { imFontAtlas_SetTexID.Fn()(fa, id) }
func (fa *ImFontAtlas) SetTexIDImTextureRef(id ImTextureRef) {
	imFontAtlas_SetTexIDImTextureRef.Fn()(fa, id)
}
func (fa *ImFontAtlas) IsBuilt() bool { return imFontAtlas_IsBuilt.Fn()(fa) != 0 }
func (fa *ImFontAtlas) GetGlyphRangesDefault() *ImWchar {
	return imFontAtlas_GetGlyphRangesDefault.Fn()(fa)
}
func (fa *ImFontAtlas) GetGlyphRangesGreek() *ImWchar {
	return imFontAtlas_GetGlyphRangesGreek.Fn()(fa)
}
func (fa *ImFontAtlas) GetGlyphRangesKorean() *ImWchar {
	return imFontAtlas_GetGlyphRangesKorean.Fn()(fa)
}
func (fa *ImFontAtlas) GetGlyphRangesJapanese() *ImWchar {
	return imFontAtlas_GetGlyphRangesJapanese.Fn()(fa)
}
func (fa *ImFontAtlas) GetGlyphRangesChineseFull() *ImWchar {
	return imFontAtlas_GetGlyphRangesChineseFull.Fn()(fa)
}
func (fa *ImFontAtlas) GetGlyphRangesChineseSimplifiedCommon() *ImWchar {
	return imFontAtlas_GetGlyphRangesChineseSimplifiedCommon.Fn()(fa)
}
func (fa *ImFontAtlas) GetGlyphRangesCyrillic() *ImWchar {
	return imFontAtlas_GetGlyphRangesCyrillic.Fn()(fa)
}
func (fa *ImFontAtlas) GetGlyphRangesThai() *ImWchar {
	return imFontAtlas_GetGlyphRangesThai.Fn()(fa)
}
func (fa *ImFontAtlas) GetGlyphRangesVietnamese() *ImWchar {
	return imFontAtlas_GetGlyphRangesVietnamese.Fn()(fa)
}

func (fa *ImFontAtlas) AddCustomRect(width int32, height int32, outRect *ImFontAtlasRect /* = NULL */) ImFontAtlasRectId {
	return imFontAtlas_AddCustomRect.Fn()(fa, width, height, outRect)
}
func (fa *ImFontAtlas) RemoveCustomRect(id ImFontAtlasRectId) {
	imFontAtlas_RemoveCustomRect.Fn()(fa, id)
}
func (fa *ImFontAtlas) GetCustomRect(id ImFontAtlasRectId) (outRect ImFontAtlasRect, ok bool) {
	ok = imFontAtlas_GetCustomRect.Fn()(fa, id, &outRect) != 0
	return
}
func (fa *ImFontAtlas) AddCustomRectRegular(w int32, h int32) ImFontAtlasRectId {
	return imFontAtlas_AddCustomRectRegular.Fn()(fa, w, h)
}
func (fa *ImFontAtlas) GetCustomRectByIndex(id ImFontAtlasRectId) *ImFontAtlasRect {
	return imFontAtlas_GetCustomRectByIndex.Fn()(fa, id)
}
func (fa *ImFontAtlas) CalcCustomRectUV(r *ImFontAtlasRect, out_uv_min, out_uv_max *ImVec2) {
	imFontAtlas_CalcCustomRectUV.Fn()(fa, r, out_uv_min, out_uv_max)
}
func (fa *ImFontAtlas) AddCustomRectFontGlyph(font *ImFont, codepoint ImWchar, w int32, h int32, advance_x float32, offset ImVec2 /* = ImVec2(0, 0) */) ImFontAtlasRectId {
	return imFontAtlas_AddCustomRectFontGlyph.Fn()(fa, font, codepoint, w, h, advance_x, offset)
}
func (fa *ImFontAtlas) AddCustomRectFontGlyphForSize(font *ImFont, font_size float32, codepoint ImWchar, w int32, h int32, advance_x float32, offset ImVec2 /* = ImVec2(0, 0) */) ImFontAtlasRectId {
	return imFontAtlas_AddCustomRectFontGlyphForSize.Fn()(fa, font, font_size, codepoint, w, h, advance_x, offset)
}

func (fb *ImFontBaked) ClearOutputData()                 { imFontBaked_ClearOutputData.Fn()(fb) }
func (fb *ImFontBaked) FindGlyph(c ImWchar) *ImFontGlyph { return imFontBaked_FindGlyph.Fn()(fb, c) }
func (fb *ImFontBaked) FindGlyphNoFallback(c ImWchar) *ImFontGlyph {
	return imFontBaked_FindGlyphNoFallback.Fn()(fb, c)
}
func (fb *ImFontBaked) GetCharAdvance(c ImWchar) float32 {
	return imFontBaked_GetCharAdvance.Fn()(fb, c)
}
func (fb *ImFontBaked) IsGlyphLoaded(c ImWchar) bool {
	return imFontBaked_IsGlyphLoaded.Fn()(fb, c) != 0
}

func (f *ImFont) IsGlyphInFont(c ImWchar) bool { return imFont_IsGlyphInFont.Fn()(f, c) != 0 }
func (f *ImFont) IsLoaded() bool               { return imFont_IsLoaded.Fn()(f) != 0 }
func (f *ImFont) GetDebugName() cc.String      { return imFont_GetDebugName.Fn()(f) }
func (f *ImFont) GetFontBaked(font_size, density float32 /* = -1.0f */) *ImFontBaked {
	return imFont_GetFontBaked.Fn()(f, font_size, density)
}
func (f *ImFont) CalcTextSizeA(size, max_width, wrap_width float32, text_begin cc.String, text_end cc.String /* = NULL */, remaining *cc.String /* = NULL */) ImVec2 {
	return imFont_CalcTextSizeA.Fn()(f, size, max_width, wrap_width, text_begin, text_end, remaining)
}
func (f *ImFont) CalcWordWrapPosition(size float32, text cc.String, text_end cc.String, wrap_width float32) cc.String {
	return imFont_CalcWordWrapPosition.Fn()(f, size, text, text_end, wrap_width)
}
func (f *ImFont) RenderChar(draw_list *ImDrawList, size float32, pos ImVec2, col ImU32, c ImWchar, cpu_fine_clip *ImVec4 /* = NULL */) {
	imFont_RenderChar.Fn()(f, draw_list, size, pos, col, c, cpu_fine_clip)
}
func (f *ImFont) RenderText(draw_list *ImDrawList, size float32, pos ImVec2, col ImU32, clip_rect ImVec4, text_begin cc.String, text_end cc.String, wrap_width float32 /* = 0.0f */, cpu_fine_clip bool) {
	imFont_RenderText.Fn()(f, draw_list, size, pos, col, clip_rect, text_begin, text_end, wrap_width, cc.MakeBool(cpu_fine_clip))
}
func (f *ImFont) CalcWordWrapPositionA(scale float32, text cc.String, text_end cc.String, wrap_width float32) cc.String {
	return imFont_CalcWordWrapPositionA.Fn()(f, scale, text, text_end, wrap_width)
}
func (f *ImFont) ClearOutputData() { imFont_ClearOutputData.Fn()(f) }
func (f *ImFont) AddRemapChar(from_codepoint, to_codepoint ImWchar) {
	imFont_AddRemapChar.Fn()(f, from_codepoint, to_codepoint)
}
func (f *ImFont) IsGlyphRangeUnused(c_begin, c_last uint32) bool {
	return imFont_IsGlyphRangeUnused.Fn()(f, c_begin, c_last) != 0
}

func (vp *ImGuiViewport) GetCenter() ImVec2     { return imGuiViewport_GetCenter.Fn()(vp) }
func (vp *ImGuiViewport) GetWorkCenter() ImVec2 { return imGuiViewport_GetWorkCenter.Fn()(vp) }
func PushFont(font *ImFont)                     { imGui_PushFont.Fn()(font) }
func SetWindowFontScale(scale float32)          { imGui_SetWindowFontScale.Fn()(scale) }
