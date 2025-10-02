package dcimgui

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
)

type uptr = unsafe.Pointer

func init() { cc.Open("libdcimgui*") }

var (
	imGui_ImplOpenGL3_Init           = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_ImplOpenGL3_Init"}
	imGui_ImplOpenGL3_Shutdown       = cc.DlFunc[func(), cc.Void]{Name: "ImGui_ImplOpenGL3_Shutdown"}
	imGui_ImplOpenGL3_NewFrame       = cc.DlFunc[func(), cc.Void]{Name: "ImGui_ImplOpenGL3_NewFrame"}
	imGui_ImplOpenGL3_RenderDrawData = cc.DlFunc[func(*ImDrawData), cc.Void]{Name: "ImGui_ImplOpenGL3_RenderDrawData"}

	imGui_CreateContext     = cc.DlFunc[func(*ImFontAtlas) *ImGuiContext, *ImGuiContext]{Name: "ImGui_CreateContext"}
	imGui_DestroyContext    = cc.DlFunc[func(*ImGuiContext), cc.Void]{Name: "ImGui_DestroyContext"}
	imGui_GetCurrentContext = cc.DlFunc[func() *ImGuiContext, *ImGuiContext]{Name: "ImGui_GetCurrentContext"}
	imGui_SetCurrentContext = cc.DlFunc[func(*ImGuiContext), cc.Void]{Name: "ImGui_SetCurrentContext"}

	imGui_GetIO         = cc.DlFunc[func() *ImGuiIO, *ImGuiIO]{Name: "ImGui_GetIO"}
	imGui_GetPlatformIO = cc.DlFunc[func() *ImGuiPlatformIO, *ImGuiPlatformIO]{Name: "ImGui_GetPlatformIO"}
	imGui_GetStyle      = cc.DlFunc[func() *ImGuiStyle, *ImGuiStyle]{Name: "ImGui_GetStyle"}
	imGui_NewFrame      = cc.DlFunc[func(), cc.Void]{Name: "ImGui_NewFrame"}
	imGui_EndFrame      = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndFrame"}
	imGui_Render        = cc.DlFunc[func(), cc.Void]{Name: "ImGui_Render"}
	imGui_GetDrawData   = cc.DlFunc[func() *ImDrawData, *ImDrawData]{Name: "ImGui_GetDrawData"}

	imGui_ShowDemoWindow        = cc.DlFunc[func(*cc.Bool), cc.Void]{Name: "ImGui_ShowDemoWindow"}
	imGui_ShowMetricsWindow     = cc.DlFunc[func(*cc.Bool), cc.Void]{Name: "ImGui_ShowMetricsWindow"}
	imGui_ShowDebugLogWindow    = cc.DlFunc[func(*cc.Bool), cc.Void]{Name: "ImGui_ShowDebugLogWindow"}
	imGui_ShowIDStackToolWindow = cc.DlFunc[func(*cc.Bool), cc.Void]{Name: "ImGui_ShowIDStackToolWindow"}
	imGui_ShowAboutWindow       = cc.DlFunc[func(*cc.Bool), cc.Void]{Name: "ImGui_ShowAboutWindow"}
	imGui_ShowStyleEditor       = cc.DlFunc[func(*ImGuiStyle), cc.Void]{Name: "ImGui_ShowStyleEditor"}
	imGui_ShowStyleSelector     = cc.DlFunc[func(cc.String) cc.Bool, cc.Bool]{Name: "ImGui_ShowStyleSelector"}
	imGui_ShowFontSelector      = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_ShowFontSelector"}
	imGui_ShowUserGuide         = cc.DlFunc[func(), cc.Void]{Name: "ImGui_ShowUserGuide"}
	imGui_GetVersion            = cc.DlFunc[func() cc.String, cc.String]{Name: "ImGui_GetVersion"}

	imGui_StyleColorsDark    = cc.DlFunc[func(*ImGuiStyle), cc.Void]{Name: "ImGui_StyleColorsDark"}
	imGui_StyleColorsLight   = cc.DlFunc[func(*ImGuiStyle), cc.Void]{Name: "ImGui_StyleColorsLight"}
	imGui_StyleColorsClassic = cc.DlFunc[func(*ImGuiStyle), cc.Void]{Name: "ImGui_StyleColorsClassic"}

	imGui_Begin = cc.DlFunc[func(cc.String, *cc.Bool, ImGuiWindowFlags) cc.Bool, cc.Bool]{Name: "ImGui_Begin"}
	imGui_End   = cc.DlFunc[func() cc.Void, cc.Void]{Name: "ImGui_End"}

	// Child Windows
	imGui_BeginChild   = cc.DlFunc[func(cc.String, ImVec2, ImGuiChildFlags, ImGuiWindowFlags) cc.Bool, cc.Bool]{Name: "ImGui_BeginChild"}
	imGui_BeginChildID = cc.DlFunc[func(ImGuiID, ImVec2, ImGuiChildFlags, ImGuiWindowFlags) cc.Bool, cc.Bool]{Name: "ImGui_BeginChildID"}
	imGui_EndChild     = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndChild"}

	// Window Status & Properties
	imGui_IsWindowAppearing = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsWindowAppearing"}
	imGui_IsWindowCollapsed = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsWindowCollapsed"}
	imGui_IsWindowFocused   = cc.DlFunc[func(ImGuiFocusedFlags) cc.Bool, cc.Bool]{Name: "ImGui_IsWindowFocused"}
	imGui_IsWindowHovered   = cc.DlFunc[func(ImGuiHoveredFlags) cc.Bool, cc.Bool]{Name: "ImGui_IsWindowHovered"}
	imGui_GetWindowDrawList = cc.DlFunc[func() *ImDrawList, *ImDrawList]{Name: "ImGui_GetWindowDrawList"}
	imGui_GetWindowDpiScale = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetWindowDpiScale"}
	imGui_GetWindowPos      = cc.DlFunc[func() ImVec2, ImVec2]{Name: "ImGui_GetWindowPos"}
	imGui_GetWindowSize     = cc.DlFunc[func() ImVec2, ImVec2]{Name: "ImGui_GetWindowSize"}
	imGui_GetWindowWidth    = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetWindowWidth"}
	imGui_GetWindowHeight   = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetWindowHeight"}
	imGui_GetWindowViewport = cc.DlFunc[func() *ImGuiViewport, *ImGuiViewport]{Name: "ImGui_GetWindowViewport"}

	imGui_SetNextWindowPos             = cc.DlFunc[func(ImVec2, ImGuiCond, ImVec2), cc.Void]{Name: "ImGui_SetNextWindowPos"}
	imGui_SetNextWindowSize            = cc.DlFunc[func(ImVec2, ImGuiCond), cc.Void]{Name: "ImGui_SetNextWindowSize"}
	imGui_SetNextWindowSizeConstraints = cc.DlFunc[func(ImVec2, ImVec2, cc.Func, *ImGuiSizeCallbackData), cc.Void]{Name: "ImGui_SetNextWindowSizeConstraints"}
	imGui_SetNextWindowContentSize     = cc.DlFunc[func(ImVec2), cc.Void]{Name: "ImGui_SetNextWindowContentSize"}
	imGui_SetNextWindowCollapsed       = cc.DlFunc[func(cc.Bool, ImGuiCond), cc.Void]{Name: "ImGui_SetNextWindowCollapsed"}
	imGui_SetNextWindowFocus           = cc.DlFunc[func(), cc.Void]{Name: "ImGui_SetNextWindowFocus"}
	imGui_SetNextWindowScroll          = cc.DlFunc[func(ImVec2), cc.Void]{Name: "ImGui_SetNextWindowScroll"}
	imGui_SetNextWindowBgAlpha         = cc.DlFunc[func(float32), cc.Void]{Name: "ImGui_SetNextWindowBgAlpha"}
	imGui_SetNextWindowViewport        = cc.DlFunc[func(ImGuiID), cc.Void]{Name: "ImGui_SetNextWindowViewport"}
	imGui_SetWindowPos                 = cc.DlFunc[func(ImVec2, ImGuiCond), cc.Void]{Name: "ImGui_SetWindowPos"}
	imGui_SetWindowSize                = cc.DlFunc[func(ImVec2, ImGuiCond), cc.Void]{Name: "ImGui_SetWindowSize"}
	imGui_SetWindowCollapsed           = cc.DlFunc[func(cc.Bool, ImGuiCond), cc.Void]{Name: "ImGui_SetWindowCollapsed"}
	imGui_SetWindowFocus               = cc.DlFunc[func(), cc.Void]{Name: "ImGui_SetWindowFocus"}
	imGui_SetWindowPosStr              = cc.DlFunc[func(cc.String, ImVec2, ImGuiCond), cc.Void]{Name: "ImGui_SetWindowPosStr"}
	imGui_SetWindowSizeStr             = cc.DlFunc[func(cc.String, ImVec2, ImGuiCond), cc.Void]{Name: "ImGui_SetWindowSizeStr"}
	imGui_SetWindowCollapsedStr        = cc.DlFunc[func(cc.String, cc.Bool, ImGuiCond), cc.Void]{Name: "ImGui_SetWindowCollapsedStr"}
	imGui_SetWindowFocusStr            = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_SetWindowFocusStr"}

	imGui_GetScrollX        = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetScrollX"}
	imGui_GetScrollY        = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetScrollY"}
	imGui_SetScrollX        = cc.DlFunc[func(float32), cc.Void]{Name: "ImGui_SetScrollX"}
	imGui_SetScrollY        = cc.DlFunc[func(float32), cc.Void]{Name: "ImGui_SetScrollY"}
	imGui_GetScrollMaxX     = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetScrollMaxX"}
	imGui_GetScrollMaxY     = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetScrollMaxY"}
	imGui_SetScrollHereX    = cc.DlFunc[func(float32), cc.Void]{Name: "ImGui_SetScrollHereX"}
	imGui_SetScrollHereY    = cc.DlFunc[func(float32), cc.Void]{Name: "ImGui_SetScrollHereY"}
	imGui_SetScrollFromPosX = cc.DlFunc[func(float32, float32), cc.Void]{Name: "ImGui_SetScrollFromPosX"}
	imGui_SetScrollFromPosY = cc.DlFunc[func(float32, float32), cc.Void]{Name: "ImGui_SetScrollFromPosY"}

	imGui_PushFontFloat = cc.DlFunc[func(*ImFont, float32), cc.Void]{Name: "ImGui_PushFontFloat"}
	imGui_PopFont       = cc.DlFunc[func(), cc.Void]{Name: "ImGui_PopFont"}
	imGui_GetFont       = cc.DlFunc[func() *ImFont, *ImFont]{Name: "ImGui_GetFont"}
	imGui_GetFontSize   = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetFontSize"}
	imGui_GetFontBaked  = cc.DlFunc[func() *ImFontBaked, *ImFontBaked]{Name: "ImGui_GetFontBaked"}

	imGui_PushStyleColor       = cc.DlFunc[func(ImGuiCol, ImU32), cc.Void]{Name: "ImGui_PushStyleColor"}
	imGui_PushStyleColorImVec4 = cc.DlFunc[func(ImGuiCol, ImVec4), cc.Void]{Name: "ImGui_PushStyleColorImVec4"}
	imGui_PopStyleColor        = cc.DlFunc[func(int), cc.Void]{Name: "ImGui_PopStyleColor"}
	imGui_PushStyleVar         = cc.DlFunc[func(ImGuiStyleVar, float32), cc.Void]{Name: "ImGui_PushStyleVar"}
	imGui_PushStyleVarImVec2   = cc.DlFunc[func(ImGuiStyleVar, ImVec2), cc.Void]{Name: "ImGui_PushStyleVarImVec2"}
	imGui_PushStyleVarX        = cc.DlFunc[func(ImGuiStyleVar, float32), cc.Void]{Name: "ImGui_PushStyleVarX"}
	imGui_PushStyleVarY        = cc.DlFunc[func(ImGuiStyleVar, float32), cc.Void]{Name: "ImGui_PushStyleVarY"}
	imGui_PopStyleVar          = cc.DlFunc[func(int), cc.Void]{Name: "ImGui_PopStyleVar"}
	imGui_PushItemFlag         = cc.DlFunc[func(ImGuiItemFlags, cc.Bool), cc.Void]{Name: "ImGui_PushItemFlag"}
	imGui_PopItemFlag          = cc.DlFunc[func(), cc.Void]{Name: "ImGui_PopItemFlag"}
	imGui_PushItemWidth        = cc.DlFunc[func(float32), cc.Void]{Name: "ImGui_PushItemWidth"}
	imGui_PopItemWidth         = cc.DlFunc[func(), cc.Void]{Name: "ImGui_PopItemWidth"}
	imGui_SetNextItemWidth     = cc.DlFunc[func(float32), cc.Void]{Name: "ImGui_SetNextItemWidth"}
	imGui_CalcItemWidth        = cc.DlFunc[func() float32, float32]{Name: "ImGui_CalcItemWidth"}
	imGui_PushTextWrapPos      = cc.DlFunc[func(float32), cc.Void]{Name: "ImGui_PushTextWrapPos"}
	imGui_PopTextWrapPos       = cc.DlFunc[func(), cc.Void]{Name: "ImGui_PopTextWrapPos"}

	imGui_GetFontTexUvWhitePixel = cc.DlFunc[func() ImVec2, ImVec2]{Name: "ImGui_GetFontTexUvWhitePixel"}
	imGui_GetColorU32            = cc.DlFunc[func(ImGuiCol, float32) ImU32, ImU32]{Name: "ImGui_GetColorU32"}
	imGui_GetColorU32ImVec4      = cc.DlFunc[func(ImVec4) ImU32, ImU32]{Name: "ImGui_GetColorU32ImVec4"}
	imGui_GetColorU32ImU32       = cc.DlFunc[func(ImU32, float32) ImU32, ImU32]{Name: "ImGui_GetColorU32ImU32"}
	imGui_GetStyleColorVec4      = cc.DlFunc[func(ImGuiCol) *ImVec4, *ImVec4]{Name: "ImGui_GetStyleColorVec4"}

	imGui_GetCursorScreenPos    = cc.DlFunc[func() ImVec2, ImVec2]{Name: "ImGui_GetCursorScreenPos"}
	imGui_SetCursorScreenPos    = cc.DlFunc[func(ImVec2), cc.Void]{Name: "ImGui_SetCursorScreenPos"}
	imGui_GetContentRegionAvail = cc.DlFunc[func() ImVec2, ImVec2]{Name: "ImGui_GetContentRegionAvail"}
	imGui_GetCursorPos          = cc.DlFunc[func() ImVec2, ImVec2]{Name: "ImGui_GetCursorPos"}
	imGui_GetCursorPosX         = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetCursorPosX"}
	imGui_GetCursorPosY         = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetCursorPosY"}
	imGui_SetCursorPos          = cc.DlFunc[func(ImVec2), cc.Void]{Name: "ImGui_SetCursorPos"}
	imGui_SetCursorPosX         = cc.DlFunc[func(float32), cc.Void]{Name: "ImGui_SetCursorPosX"}
	imGui_SetCursorPosY         = cc.DlFunc[func(float32), cc.Void]{Name: "ImGui_SetCursorPosY"}
	imGui_GetCursorStartPos     = cc.DlFunc[func() ImVec2, ImVec2]{Name: "ImGui_GetCursorStartPos"}

	imGui_Separator                    = cc.DlFunc[func(), cc.Void]{Name: "ImGui_Separator"}
	imGui_SameLine                     = cc.DlFunc[func(float32, float32), cc.Void]{Name: "ImGui_SameLine"}
	imGui_NewLine                      = cc.DlFunc[func(), cc.Void]{Name: "ImGui_NewLine"}
	imGui_Spacing                      = cc.DlFunc[func(), cc.Void]{Name: "ImGui_Spacing"}
	imGui_Dummy                        = cc.DlFunc[func(ImVec2), cc.Void]{Name: "ImGui_Dummy"}
	imGui_Indent                       = cc.DlFunc[func(float32), cc.Void]{Name: "ImGui_Indent"}
	imGui_Unindent                     = cc.DlFunc[func(float32), cc.Void]{Name: "ImGui_Unindent"}
	imGui_BeginGroup                   = cc.DlFunc[func(), cc.Void]{Name: "ImGui_BeginGroup"}
	imGui_EndGroup                     = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndGroup"}
	imGui_AlignTextToFramePadding      = cc.DlFunc[func(), cc.Void]{Name: "ImGui_AlignTextToFramePadding"}
	imGui_GetTextLineHeight            = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetTextLineHeight"}
	imGui_GetTextLineHeightWithSpacing = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetTextLineHeightWithSpacing"}
	imGui_GetFrameHeight               = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetFrameHeight"}
	imGui_GetFrameHeightWithSpacing    = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetFrameHeightWithSpacing"}

	imGui_PushID    = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_PushID"}
	imGui_PushIDStr = cc.DlFunc[func(cc.String, cc.String), cc.Void]{Name: "ImGui_PushIDStr"}
	imGui_PushIDPtr = cc.DlFunc[func(unsafe.Pointer), cc.Void]{Name: "ImGui_PushIDPtr"}
	imGui_PushIDInt = cc.DlFunc[func(int32), cc.Void]{Name: "ImGui_PushIDInt"}
	imGui_PopID     = cc.DlFunc[func(), cc.Void]{Name: "ImGui_PopID"}
	imGui_GetID     = cc.DlFunc[func(cc.String) ImGuiID, ImGuiID]{Name: "ImGui_GetID"}
	imGui_GetIDStr  = cc.DlFunc[func(cc.String, cc.String) ImGuiID, ImGuiID]{Name: "ImGui_GetIDStr"}
	imGui_GetIDPtr  = cc.DlFunc[func(unsafe.Pointer) ImGuiID, ImGuiID]{Name: "ImGui_GetIDPtr"}
	imGui_GetIDInt  = cc.DlFunc[func(int32) ImGuiID, ImGuiID]{Name: "ImGui_GetIDInt"}

	imGui_TextUnformatted         = cc.DlFunc[func(cc.String, cc.String), cc.Void]{Name: "ImGui_TextUnformatted"}
	imGui_Text                    = cc.DlFunc[func(cc.String, ...interface{}), cc.Void]{Name: "ImGui_Text", Va: true}
	imGui_TextColored             = cc.DlFunc[func(ImVec4, cc.String, ...interface{}), cc.Void]{Name: "ImGui_TextColored"}
	imGui_TextColoredUnformatted  = cc.DlFunc[func(ImVec4, cc.String), cc.Void]{Name: "ImGui_TextColoredUnformatted"}
	imGui_TextDisabled            = cc.DlFunc[func(cc.String, ...interface{}), cc.Void]{Name: "ImGui_TextDisabled", Va: true}
	imGui_TextDisabledUnformatted = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_TextDisabledUnformatted"}
	imGui_TextWrapped             = cc.DlFunc[func(cc.String, ...interface{}), cc.Void]{Name: "ImGui_TextWrapped", Va: true}
	imGui_TextWrappedUnformatted  = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_TextWrappedUnformatted"}
	imGui_LabelText               = cc.DlFunc[func(cc.String, cc.String, ...interface{}), cc.Void]{Name: "ImGui_LabelText", Va: true}
	imGui_LabelTextUnformatted    = cc.DlFunc[func(cc.String, cc.String), cc.Void]{Name: "ImGui_LabelTextUnformatted"}
	imGui_BulletText              = cc.DlFunc[func(cc.String, ...interface{}), cc.Void]{Name: "ImGui_BulletText"}
	imGui_BulletTextUnformatted   = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_BulletTextUnformatted"}
	imGui_SeparatorText           = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_SeparatorText"}

	imGui_Button               = cc.DlFunc[func(cc.String, ImVec2) cc.Bool, cc.Bool]{Name: "ImGui_Button"}
	imGui_SmallButton          = cc.DlFunc[func(cc.String) cc.Bool, cc.Bool]{Name: "ImGui_SmallButton"}
	imGui_InvisibleButton      = cc.DlFunc[func(cc.String, ImVec2, ImGuiButtonFlags) cc.Bool, cc.Bool]{Name: "ImGui_InvisibleButton"}
	imGui_ArrowButton          = cc.DlFunc[func(cc.String, ImGuiDir) cc.Bool, cc.Bool]{Name: "ImGui_ArrowButton"}
	imGui_Checkbox             = cc.DlFunc[func(cc.String, *cc.Bool) cc.Bool, cc.Bool]{Name: "ImGui_Checkbox"}
	imGui_CheckboxFlagsIntPtr  = cc.DlFunc[func(cc.String, *int32, int32) cc.Bool, cc.Bool]{Name: "ImGui_CheckboxFlagsIntPtr"}
	imGui_CheckboxFlagsUintPtr = cc.DlFunc[func(cc.String, *uint32, uint32) cc.Bool, cc.Bool]{Name: "ImGui_CheckboxFlagsUintPtr"}
	imGui_RadioButton          = cc.DlFunc[func(cc.String, cc.Bool) cc.Bool, cc.Bool]{Name: "ImGui_RadioButton"}
	imGui_RadioButtonIntPtr    = cc.DlFunc[func(cc.String, *int32, int32) cc.Bool, cc.Bool]{Name: "ImGui_RadioButtonIntPtr"}
	imGui_ProgressBar          = cc.DlFunc[func(float32, ImVec2, cc.String) cc.Void, cc.Void]{Name: "ImGui_ProgressBar"}
	imGui_Bullet               = cc.DlFunc[func() cc.Void, cc.Void]{Name: "ImGui_Bullet"}
	imGui_TextLink             = cc.DlFunc[func(cc.String) cc.Bool, cc.Bool]{Name: "ImGui_TextLink"}
	imGui_TextLinkOpenURL      = cc.DlFunc[func(cc.String, cc.String) cc.Bool, cc.Bool]{Name: "ImGui_TextLinkOpenURL"}

	imGui_Image       = cc.DlFunc[func(ImTextureRef, ImVec2, ImVec2, ImVec2) cc.Void, cc.Void]{Name: "ImGui_Image"}
	imGui_ImageWithBg = cc.DlFunc[func(ImTextureRef, ImVec2, ImVec2, ImVec2, ImVec4, ImVec4) cc.Void, cc.Void]{Name: "ImGui_ImageWithBg"}
	imGui_ImageButton = cc.DlFunc[func(cc.String, ImTextureRef, ImVec2, ImVec2, ImVec2, ImVec4, ImVec4) cc.Bool, cc.Bool]{Name: "ImGui_ImageButton"}

	imGui_BeginCombo    = cc.DlFunc[func(cc.String, cc.String, ImGuiComboFlags) cc.Bool, cc.Bool]{Name: "ImGui_BeginCombo"}
	imGui_EndCombo      = cc.DlFunc[func() cc.Void, cc.Void]{Name: "ImGui_EndCombo"}
	imGui_ComboChar     = cc.DlFunc[func(cc.String, *int32, cc.Strings, int32, int32) cc.Bool, cc.Bool]{Name: "ImGui_ComboChar"}
	imGui_Combo         = cc.DlFunc[func(cc.String, *int32, cc.String, int32) cc.Bool, cc.Bool]{Name: "ImGui_Combo"}
	imGui_ComboCallback = cc.DlFunc[func(cc.String, *int32, cc.Func, unsafe.Pointer, int32, int32) cc.Bool, cc.Bool]{Name: "ImGui_ComboCallback"}

	imGui_DragFloat       = cc.DlFunc[func(cc.String, *float32, float32, float32, float32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_DragFloat"}
	imGui_DragFloat2      = cc.DlFunc[func(cc.String, *float32, float32, float32, float32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_DragFloat2"}
	imGui_DragFloat3      = cc.DlFunc[func(cc.String, *float32, float32, float32, float32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_DragFloat3"}
	imGui_DragFloat4      = cc.DlFunc[func(cc.String, *float32, float32, float32, float32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_DragFloat4"}
	imGui_DragFloatRange2 = cc.DlFunc[func(cc.String, *float32, *float32, float32, float32, float32, cc.String, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_DragFloatRange2"}
	imGui_DragInt         = cc.DlFunc[func(cc.String, *int32, float32, int32, int32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_DragInt"}
	imGui_DragInt2        = cc.DlFunc[func(cc.String, *int32, float32, int32, int32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_DragInt2"}
	imGui_DragInt3        = cc.DlFunc[func(cc.String, *int32, float32, int32, int32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_DragInt3"}
	imGui_DragInt4        = cc.DlFunc[func(cc.String, *int32, float32, int32, int32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_DragInt4"}
	imGui_DragIntRange2   = cc.DlFunc[func(cc.String, *int32, *int32, float32, int32, int32, cc.String, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_DragIntRange2"}
	imGui_DragScalar      = cc.DlFunc[func(cc.String, ImGuiDataType, unsafe.Pointer, float32, unsafe.Pointer, unsafe.Pointer, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_DragScalar"}
	imGui_DragScalarN     = cc.DlFunc[func(cc.String, ImGuiDataType, unsafe.Pointer, int32, float32, unsafe.Pointer, unsafe.Pointer, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_DragScalarN"}

	imGui_SliderFloat   = cc.DlFunc[func(cc.String, *float32, float32, float32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_SliderFloat"}
	imGui_SliderFloat2  = cc.DlFunc[func(cc.String, *float32, float32, float32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_SliderFloat2"}
	imGui_SliderFloat3  = cc.DlFunc[func(cc.String, *float32, float32, float32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_SliderFloat3"}
	imGui_SliderFloat4  = cc.DlFunc[func(cc.String, *float32, float32, float32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_SliderFloat4"}
	imGui_SliderAngle   = cc.DlFunc[func(cc.String, *float32, float32, float32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_SliderAngle"}
	imGui_SliderInt     = cc.DlFunc[func(cc.String, *int32, int32, int32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_SliderInt"}
	imGui_SliderInt2    = cc.DlFunc[func(cc.String, *int32, int32, int32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_SliderInt2"}
	imGui_SliderInt3    = cc.DlFunc[func(cc.String, *int32, int32, int32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_SliderInt3"}
	imGui_SliderInt4    = cc.DlFunc[func(cc.String, *int32, int32, int32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_SliderInt4"}
	imGui_SliderScalar  = cc.DlFunc[func(cc.String, ImGuiDataType, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_SliderScalar"}
	imGui_SliderScalarN = cc.DlFunc[func(cc.String, ImGuiDataType, unsafe.Pointer, int32, unsafe.Pointer, unsafe.Pointer, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_SliderScalarN"}
	imGui_VSliderFloat  = cc.DlFunc[func(cc.String, ImVec2, *float32, float32, float32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_VSliderFloat"}
	imGui_VSliderInt    = cc.DlFunc[func(cc.String, ImVec2, *int32, int32, int32, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_VSliderInt"}
	imGui_VSliderScalar = cc.DlFunc[func(cc.String, ImVec2, ImGuiDataType, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, cc.String, ImGuiSliderFlags) cc.Bool, cc.Bool]{Name: "ImGui_VSliderScalar"}

	imGui_InputText          = cc.DlFunc[func(cc.String, cc.String, uint64, ImGuiInputTextFlags, cc.Func, unsafe.Pointer) cc.Bool, cc.Bool]{Name: "ImGui_InputText"}
	imGui_InputTextMultiline = cc.DlFunc[func(cc.String, cc.String, uint64, ImVec2, ImGuiInputTextFlags, cc.Func, unsafe.Pointer) cc.Bool, cc.Bool]{Name: "ImGui_InputTextMultiline"}
	imGui_InputTextWithHint  = cc.DlFunc[func(cc.String, cc.String, cc.String, uint64, ImGuiInputTextFlags, cc.Func, unsafe.Pointer) cc.Bool, cc.Bool]{Name: "ImGui_InputTextWithHint"}
	imGui_InputFloat         = cc.DlFunc[func(cc.String, *float32, float32, float32, cc.String, ImGuiInputTextFlags) cc.Bool, cc.Bool]{Name: "ImGui_InputFloat"}
	imGui_InputFloat2        = cc.DlFunc[func(cc.String, *float32, cc.String, ImGuiInputTextFlags) cc.Bool, cc.Bool]{Name: "ImGui_InputFloat2"}
	imGui_InputFloat3        = cc.DlFunc[func(cc.String, *float32, cc.String, ImGuiInputTextFlags) cc.Bool, cc.Bool]{Name: "ImGui_InputFloat3"}
	imGui_InputFloat4        = cc.DlFunc[func(cc.String, *float32, cc.String, ImGuiInputTextFlags) cc.Bool, cc.Bool]{Name: "ImGui_InputFloat4"}
	imGui_InputInt           = cc.DlFunc[func(cc.String, *int32, int32, int32, ImGuiInputTextFlags) cc.Bool, cc.Bool]{Name: "ImGui_InputInt"}
	imGui_InputInt2          = cc.DlFunc[func(cc.String, *int32, ImGuiInputTextFlags) cc.Bool, cc.Bool]{Name: "ImGui_InputInt2"}
	imGui_InputInt3          = cc.DlFunc[func(cc.String, *int32, ImGuiInputTextFlags) cc.Bool, cc.Bool]{Name: "ImGui_InputInt3"}
	imGui_InputInt4          = cc.DlFunc[func(cc.String, *int32, ImGuiInputTextFlags) cc.Bool, cc.Bool]{Name: "ImGui_InputInt4"}
	imGui_InputDouble        = cc.DlFunc[func(cc.String, *float64, float64, float64, cc.String, ImGuiInputTextFlags) cc.Bool, cc.Bool]{Name: "ImGui_InputDouble"}
	imGui_InputScalar        = cc.DlFunc[func(cc.String, ImGuiDataType, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, cc.String, ImGuiInputTextFlags) cc.Bool, cc.Bool]{Name: "ImGui_InputScalar"}
	imGui_InputScalarN       = cc.DlFunc[func(cc.String, ImGuiDataType, unsafe.Pointer, int32, unsafe.Pointer, unsafe.Pointer, cc.String, ImGuiInputTextFlags) cc.Bool, cc.Bool]{Name: "ImGui_InputScalarN"}

	imGui_ColorEdit3          = cc.DlFunc[func(cc.String, *float32, ImGuiColorEditFlags) cc.Bool, cc.Bool]{Name: "ImGui_ColorEdit3"}
	imGui_ColorEdit4          = cc.DlFunc[func(cc.String, *float32, ImGuiColorEditFlags) cc.Bool, cc.Bool]{Name: "ImGui_ColorEdit4"}
	imGui_ColorPicker3        = cc.DlFunc[func(cc.String, *float32, ImGuiColorEditFlags) cc.Bool, cc.Bool]{Name: "ImGui_ColorPicker3"}
	imGui_ColorPicker4        = cc.DlFunc[func(cc.String, *float32, ImGuiColorEditFlags, *float32) cc.Bool, cc.Bool]{Name: "ImGui_ColorPicker4"}
	imGui_ColorButton         = cc.DlFunc[func(cc.String, ImVec4, ImGuiColorEditFlags, ImVec2) cc.Bool, cc.Bool]{Name: "ImGui_ColorButton"}
	imGui_SetColorEditOptions = cc.DlFunc[func(ImGuiColorEditFlags), cc.Void]{Name: "ImGui_SetColorEditOptions"}

	imGui_TreeNode               = cc.DlFunc[func(cc.String) cc.Bool, cc.Bool]{Name: "ImGui_TreeNode"}
	imGui_TreeNodeStr            = cc.DlFunc[func(cc.String, cc.String, ...any) cc.Bool, cc.Bool]{Name: "ImGui_TreeNodeStr", Va: true}
	imGui_TreeNodeStrUnformatted = cc.DlFunc[func(cc.String, cc.String) cc.Bool, cc.Bool]{Name: "ImGui_TreeNodeStrUnformatted"}
	imGui_TreeNodePtr            = cc.DlFunc[func(unsafe.Pointer, cc.String, ...any) cc.Bool, cc.Bool]{Name: "ImGui_TreeNodePtr", Va: true}
	imGui_TreeNodePtrUnformatted = cc.DlFunc[func(unsafe.Pointer, cc.String) cc.Bool, cc.Bool]{Name: "ImGui_TreeNodePtrUnformatted"}
	// imGui_TreeNodeV                 = cc.DlFunc[func(cc.String, cc.String, unsafe.Pointer) cc.Bool, cc.Bool]{Name: "ImGui_TreeNodeV"}
	// imGui_TreeNodeVPtr              = cc.DlFunc[func(unsafe.Pointer, cc.String, unsafe.Pointer) cc.Bool, cc.Bool]{Name: "ImGui_TreeNodeVPtr"}
	imGui_TreeNodeEx               = cc.DlFunc[func(cc.String, ImGuiTreeNodeFlags) cc.Bool, cc.Bool]{Name: "ImGui_TreeNodeEx"}
	imGui_TreeNodeExStr            = cc.DlFunc[func(cc.String, ImGuiTreeNodeFlags, cc.String, ...any) cc.Bool, cc.Bool]{Name: "ImGui_TreeNodeExStr", Va: true}
	imGui_TreeNodeExStrUnformatted = cc.DlFunc[func(cc.String, ImGuiTreeNodeFlags, cc.String) cc.Bool, cc.Bool]{Name: "ImGui_TreeNodeExStrUnformatted"}
	imGui_TreeNodeExPtr            = cc.DlFunc[func(unsafe.Pointer, ImGuiTreeNodeFlags, cc.String, ...any) cc.Bool, cc.Bool]{Name: "ImGui_TreeNodeExPtr", Va: true}
	imGui_TreeNodeExPtrUnformatted = cc.DlFunc[func(unsafe.Pointer, ImGuiTreeNodeFlags, cc.String) cc.Bool, cc.Bool]{Name: "ImGui_TreeNodeExPtrUnformatted"}
	// imGui_TreeNodeExV               = cc.DlFunc[func(cc.String, ImGuiTreeNodeFlags, cc.String, unsafe.Pointer) cc.Bool, cc.Bool]{Name: "ImGui_TreeNodeExV"}
	// imGui_TreeNodeExVPtr            = cc.DlFunc[func(unsafe.Pointer, ImGuiTreeNodeFlags, cc.String, unsafe.Pointer) cc.Bool, cc.Bool]{Name: "ImGui_TreeNodeExVPtr"}
	imGui_TreePush                  = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_TreePush"}
	imGui_TreePushPtr               = cc.DlFunc[func(unsafe.Pointer), cc.Void]{Name: "ImGui_TreePushPtr"}
	imGui_TreePop                   = cc.DlFunc[func(), cc.Void]{Name: "ImGui_TreePop"}
	imGui_GetTreeNodeToLabelSpacing = cc.DlFunc[func() float32, float32]{Name: "ImGui_GetTreeNodeToLabelSpacing"}
	imGui_CollapsingHeader          = cc.DlFunc[func(cc.String, ImGuiTreeNodeFlags) cc.Bool, cc.Bool]{Name: "ImGui_CollapsingHeader"}
	imGui_CollapsingHeaderBoolPtr   = cc.DlFunc[func(cc.String, *cc.Bool, ImGuiTreeNodeFlags) cc.Bool, cc.Bool]{Name: "ImGui_CollapsingHeaderBoolPtr"}
	imGui_SetNextItemOpen           = cc.DlFunc[func(cc.Bool, ImGuiCond), cc.Void]{Name: "ImGui_SetNextItemOpen"}
	imGui_SetNextItemStorageID      = cc.DlFunc[func(ImGuiID), cc.Void]{Name: "ImGui_SetNextItemStorageID"}

	imGui_Selectable                   = cc.DlFunc[func(cc.String, cc.Bool, ImGuiSelectableFlags, ImVec2) cc.Bool, cc.Bool]{Name: "ImGui_Selectable"}
	imGui_SelectableBoolPtr            = cc.DlFunc[func(cc.String, *cc.Bool, ImGuiSelectableFlags, ImVec2) cc.Bool, cc.Bool]{Name: "ImGui_SelectableBoolPtr"}
	imGui_BeginMultiSelect             = cc.DlFunc[func(ImGuiMultiSelectFlags, int32, int32) *ImGuiMultiSelectIO, *ImGuiMultiSelectIO]{Name: "ImGui_BeginMultiSelect"}
	imGui_EndMultiSelect               = cc.DlFunc[func() *ImGuiMultiSelectIO, *ImGuiMultiSelectIO]{Name: "ImGui_EndMultiSelect"}
	imGui_SetNextItemSelectionUserData = cc.DlFunc[func(ImGuiSelectionUserData), cc.Void]{Name: "ImGui_SetNextItemSelectionUserData"}
	imGui_IsItemToggledSelection       = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsItemToggledSelection"}
	imGui_BeginListBox                 = cc.DlFunc[func(cc.String, ImVec2) cc.Bool, cc.Bool]{Name: "ImGui_BeginListBox"}
	imGui_EndListBox                   = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndListBox"}
	imGui_ListBox                      = cc.DlFunc[func(cc.String, *int32, *cc.String, int32, int32) cc.Bool, cc.Bool]{Name: "ImGui_ListBox"}
	imGui_ListBoxCallback              = cc.DlFunc[func(cc.String, *int32, cc.Func, unsafe.Pointer, int32, int32) cc.Bool, cc.Bool]{Name: "ImGui_ListBoxCallback"}

	imGui_PlotLines             = cc.DlFunc[func(cc.String, *float32, int32, int32, cc.String, float32, float32, ImVec2, int32), cc.Void]{Name: "ImGui_PlotLines"}
	imGui_PlotLinesCallback     = cc.DlFunc[func(cc.String, cc.Func, unsafe.Pointer, int32, int32, cc.String, float32, float32, ImVec2), cc.Void]{Name: "ImGui_PlotLinesCallback"}
	imGui_PlotHistogram         = cc.DlFunc[func(cc.String, *float32, int32, int32, cc.String, float32, float32, ImVec2, int32), cc.Void]{Name: "ImGui_PlotHistogram"}
	imGui_PlotHistogramCallback = cc.DlFunc[func(cc.String, cc.Func, unsafe.Pointer, int32, int32, cc.String, float32, float32, ImVec2), cc.Void]{Name: "ImGui_PlotHistogramCallback"}
	imGui_BeginMenuBar          = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_BeginMenuBar"}
	imGui_EndMenuBar            = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndMenuBar"}
	imGui_BeginMainMenuBar      = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_BeginMainMenuBar"}
	imGui_EndMainMenuBar        = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndMainMenuBar"}
	imGui_BeginMenu             = cc.DlFunc[func(cc.String, cc.Bool) cc.Bool, cc.Bool]{Name: "ImGui_BeginMenu"}
	imGui_EndMenu               = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndMenu"}
	imGui_MenuItem              = cc.DlFunc[func(cc.String, cc.String, cc.Bool, cc.Bool) cc.Bool, cc.Bool]{Name: "ImGui_MenuItem"}
	imGui_MenuItemBoolPtr       = cc.DlFunc[func(cc.String, cc.String, *cc.Bool, cc.Bool) cc.Bool, cc.Bool]{Name: "ImGui_MenuItemBoolPtr"}

	imGui_BeginTooltip          = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_BeginTooltip"}
	imGui_EndTooltip            = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndTooltip"}
	imGui_SetTooltip            = cc.DlFunc[func(cc.String, ...any) cc.Void, cc.Void]{Name: "ImGui_SetTooltip", Va: true}
	imGui_SetTooltipUnformatted = cc.DlFunc[func(cc.String) cc.Void, cc.Void]{Name: "ImGui_SetTooltipUnformatted"}
	// imGui_SetTooltipV               = cc.DlFunc[func(cc.String, unsafe.Pointer) cc.Void, cc.Void]{Name: "ImGui_SetTooltipV"}
	imGui_BeginItemTooltip          = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_BeginItemTooltip"}
	imGui_SetItemTooltip            = cc.DlFunc[func(cc.String, ...any) cc.Void, cc.Void]{Name: "ImGui_SetItemTooltip", Va: true}
	imGui_SetItemTooltipUnformatted = cc.DlFunc[func(cc.String) cc.Void, cc.Void]{Name: "ImGui_SetItemTooltipUnformatted"}
	// imGui_SetItemTooltipV           = cc.DlFunc[func(cc.String, unsafe.Pointer) cc.Void, cc.Void]{Name: "ImGui_SetItemTooltipV"}

	imGui_BeginPopup           = cc.DlFunc[func(cc.String, ImGuiPopupFlags) cc.Bool, cc.Bool]{Name: "ImGui_BeginPopup"}
	imGui_BeginPopupModal      = cc.DlFunc[func(cc.String, *cc.Bool, ImGuiWindowFlags) cc.Bool, cc.Bool]{Name: "ImGui_BeginPopupModal"}
	imGui_EndPopup             = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndPopup"}
	imGui_OpenPopup            = cc.DlFunc[func(cc.String, ImGuiPopupFlags) cc.Void, cc.Void]{Name: "ImGui_OpenPopup"}
	imGui_OpenPopupID          = cc.DlFunc[func(ImGuiID, ImGuiPopupFlags) cc.Void, cc.Void]{Name: "ImGui_OpenPopupID"}
	imGui_OpenPopupOnItemClick = cc.DlFunc[func(cc.String, ImGuiPopupFlags) cc.Void, cc.Void]{Name: "ImGui_OpenPopupOnItemClick"}
	imGui_CloseCurrentPopup    = cc.DlFunc[func(), cc.Void]{Name: "ImGui_CloseCurrentPopup"}

	imGui_BeginPopupContextItem   = cc.DlFunc[func(cc.String, ImGuiPopupFlags) cc.Bool, cc.Bool]{Name: "ImGui_BeginPopupContextItem"}
	imGui_BeginPopupContextWindow = cc.DlFunc[func(cc.String, ImGuiPopupFlags) cc.Bool, cc.Bool]{Name: "ImGui_BeginPopupContextWindow"}
	imGui_BeginPopupContextVoid   = cc.DlFunc[func(cc.String, ImGuiPopupFlags) cc.Bool, cc.Bool]{Name: "ImGui_BeginPopupContextVoid"}
	imGui_IsPopupOpen             = cc.DlFunc[func(cc.String, ImGuiPopupFlags) cc.Bool, cc.Bool]{Name: "ImGui_IsPopupOpen"}
	imGui_BeginTable              = cc.DlFunc[func(cc.String, int32, ImGuiTableFlags, ImVec2, float32) cc.Bool, cc.Bool]{Name: "ImGui_BeginTable"}
	imGui_EndTable                = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndTable"}
	imGui_TableNextRow            = cc.DlFunc[func(ImGuiTableRowFlags, float32), cc.Void]{Name: "ImGui_TableNextRow"}
	imGui_TableNextColumn         = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_TableNextColumn"}
	imGui_TableSetColumnIndex     = cc.DlFunc[func(int32) cc.Bool, cc.Bool]{Name: "ImGui_TableSetColumnIndex"}

	imGui_TableSetupColumn       = cc.DlFunc[func(cc.String, ImGuiTableColumnFlags, float32, ImGuiID), cc.Void]{Name: "ImGui_TableSetupColumn"}
	imGui_TableSetupScrollFreeze = cc.DlFunc[func(int32, int32), cc.Void]{Name: "ImGui_TableSetupScrollFreeze"}
	imGui_TableHeader            = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_TableHeader"}
	imGui_TableHeadersRow        = cc.DlFunc[func(), cc.Void]{Name: "ImGui_TableHeadersRow"}
	imGui_TableAngledHeadersRow  = cc.DlFunc[func(), cc.Void]{Name: "ImGui_TableAngledHeadersRow"}
	imGui_TableGetSortSpecs      = cc.DlFunc[func() *ImGuiTableSortSpecs, *ImGuiTableSortSpecs]{Name: "ImGui_TableGetSortSpecs"}
	imGui_TableGetColumnCount    = cc.DlFunc[func() int32, int32]{Name: "ImGui_TableGetColumnCount"}
	imGui_TableGetColumnIndex    = cc.DlFunc[func() int32, int32]{Name: "ImGui_TableGetColumnIndex"}
	imGui_TableGetRowIndex       = cc.DlFunc[func() int32, int32]{Name: "ImGui_TableGetRowIndex"}
	imGui_TableGetColumnName     = cc.DlFunc[func(int32) cc.String, cc.String]{Name: "ImGui_TableGetColumnName"}
	imGui_TableGetColumnFlags    = cc.DlFunc[func(int32) ImGuiTableColumnFlags, ImGuiTableColumnFlags]{Name: "ImGui_TableGetColumnFlags"}
	imGui_TableSetColumnEnabled  = cc.DlFunc[func(int32, cc.Bool), cc.Void]{Name: "ImGui_TableSetColumnEnabled"}
	imGui_TableGetHoveredColumn  = cc.DlFunc[func() int32, int32]{Name: "ImGui_TableGetHoveredColumn"}
	imGui_TableSetBgColor        = cc.DlFunc[func(ImGuiTableBgTarget, ImU32, int32), cc.Void]{Name: "ImGui_TableSetBgColor"}
	imGui_Columns                = cc.DlFunc[func(int32, cc.String, cc.Bool), cc.Void]{Name: "ImGui_Columns"}
	imGui_NextColumn             = cc.DlFunc[func(), cc.Void]{Name: "ImGui_NextColumn"}
	imGui_GetColumnIndex         = cc.DlFunc[func() int32, int32]{Name: "ImGui_GetColumnIndex"}
	imGui_GetColumnWidth         = cc.DlFunc[func(int32) float32, float32]{Name: "ImGui_GetColumnWidth"}
	imGui_SetColumnWidth         = cc.DlFunc[func(int32, float32), cc.Void]{Name: "ImGui_SetColumnWidth"}
	imGui_GetColumnOffset        = cc.DlFunc[func(int32) float32, float32]{Name: "ImGui_GetColumnOffset"}
	imGui_SetColumnOffset        = cc.DlFunc[func(int32, float32), cc.Void]{Name: "ImGui_SetColumnOffset"}
	imGui_GetColumnsCount        = cc.DlFunc[func() int32, int32]{Name: "ImGui_GetColumnsCount"}

	imGui_BeginTabBar           = cc.DlFunc[func(cc.String, ImGuiTabBarFlags) cc.Bool, cc.Bool]{Name: "ImGui_BeginTabBar"}
	imGui_EndTabBar             = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndTabBar"}
	imGui_BeginTabItem          = cc.DlFunc[func(cc.String, *cc.Bool, ImGuiTabItemFlags) cc.Bool, cc.Bool]{Name: "ImGui_BeginTabItem"}
	imGui_EndTabItem            = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndTabItem"}
	imGui_TabItemButton         = cc.DlFunc[func(cc.String, ImGuiTabItemFlags) cc.Bool, cc.Bool]{Name: "ImGui_TabItemButton"}
	imGui_SetTabItemClosed      = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_SetTabItemClosed"}
	imGui_DockSpace             = cc.DlFunc[func(ImGuiID, ImVec2, ImGuiDockNodeFlags, *ImGuiWindowClass) ImGuiID, ImGuiID]{Name: "ImGui_DockSpace"}
	imGui_DockSpaceOverViewport = cc.DlFunc[func(ImGuiID, *ImGuiViewport, ImGuiDockNodeFlags, *ImGuiWindowClass) ImGuiID, ImGuiID]{Name: "ImGui_DockSpaceOverViewport"}
	imGui_SetNextWindowDockID   = cc.DlFunc[func(ImGuiID, ImGuiCond), cc.Void]{Name: "ImGui_SetNextWindowDockID"}
	imGui_SetNextWindowClass    = cc.DlFunc[func(*ImGuiWindowClass), cc.Void]{Name: "ImGui_SetNextWindowClass"}
	imGui_GetWindowDockID       = cc.DlFunc[func() ImGuiID, ImGuiID]{Name: "ImGui_GetWindowDockID"}
	imGui_IsWindowDocked        = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsWindowDocked"}

	imGui_LogToTTY           = cc.DlFunc[func(int32), cc.Void]{Name: "ImGui_LogToTTY"}
	imGui_LogToFile          = cc.DlFunc[func(int32, cc.String), cc.Void]{Name: "ImGui_LogToFile"}
	imGui_LogToClipboard     = cc.DlFunc[func(int32), cc.Void]{Name: "ImGui_LogToClipboard"}
	imGui_LogFinish          = cc.DlFunc[func(), cc.Void]{Name: "ImGui_LogFinish"}
	imGui_LogButtons         = cc.DlFunc[func(), cc.Void]{Name: "ImGui_LogButtons"}
	imGui_LogText            = cc.DlFunc[func(cc.String, ...any) cc.Void, cc.Void]{Name: "ImGui_LogText", Va: true}
	imGui_LogTextUnformatted = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_LogTextUnformatted"}
	// imGui_LogTextV = cc.DlFunc[func(cc.String, unsafe.Pointer) cc.Void, cc.Void]{Name: "ImGui_LogTextV"}

	imGui_BeginDragDropSource   = cc.DlFunc[func(ImGuiDragDropFlags) cc.Bool, cc.Bool]{Name: "ImGui_BeginDragDropSource"}
	imGui_SetDragDropPayload    = cc.DlFunc[func(cc.String, unsafe.Pointer, uint64, ImGuiCond) cc.Bool, cc.Bool]{Name: "ImGui_SetDragDropPayload"}
	imGui_EndDragDropSource     = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndDragDropSource"}
	imGui_BeginDragDropTarget   = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_BeginDragDropTarget"}
	imGui_AcceptDragDropPayload = cc.DlFunc[func(cc.String, ImGuiDragDropFlags) *ImGuiPayload, *ImGuiPayload]{Name: "ImGui_AcceptDragDropPayload"}
	imGui_EndDragDropTarget     = cc.DlFunc[func(), cc.Void]{Name: "ImGui_EndDragDropTarget"}
	imGui_GetDragDropPayload    = cc.DlFunc[func() *ImGuiPayload, *ImGuiPayload]{Name: "ImGui_GetDragDropPayload"}
	imGui_BeginDisabled         = cc.DlFunc[func(cc.Bool) cc.Void, cc.Void]{Name: "ImGui_BeginDisabled"}
	imGui_EndDisabled           = cc.DlFunc[func() cc.Void, cc.Void]{Name: "ImGui_EndDisabled"}

	imGui_PushClipRect               = cc.DlFunc[func(ImVec2, ImVec2, cc.Bool), cc.Void]{Name: "ImGui_PushClipRect"}
	imGui_PopClipRect                = cc.DlFunc[func(), cc.Void]{Name: "ImGui_PopClipRect"}
	imGui_SetItemDefaultFocus        = cc.DlFunc[func(), cc.Void]{Name: "ImGui_SetItemDefaultFocus"}
	imGui_SetKeyboardFocusHere       = cc.DlFunc[func(int32), cc.Void]{Name: "ImGui_SetKeyboardFocusHere"}
	imGui_SetNavCursorVisible        = cc.DlFunc[func(cc.Bool), cc.Void]{Name: "ImGui_SetNavCursorVisible"}
	imGui_SetNextItemAllowOverlap    = cc.DlFunc[func(), cc.Void]{Name: "ImGui_SetNextItemAllowOverlap"}
	imGui_IsItemHovered              = cc.DlFunc[func(ImGuiHoveredFlags) cc.Bool, cc.Bool]{Name: "ImGui_IsItemHovered"}
	imGui_IsItemActive               = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsItemActive"}
	imGui_IsItemFocused              = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsItemFocused"}
	imGui_IsItemClicked              = cc.DlFunc[func(ImGuiMouseButton) cc.Bool, cc.Bool]{Name: "ImGui_IsItemClicked"}
	imGui_IsItemVisible              = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsItemVisible"}
	imGui_IsItemEdited               = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsItemEdited"}
	imGui_IsItemActivated            = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsItemActivated"}
	imGui_IsItemDeactivated          = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsItemDeactivated"}
	imGui_IsItemDeactivatedAfterEdit = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsItemDeactivatedAfterEdit"}
	imGui_IsItemToggledOpen          = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsItemToggledOpen"}
	imGui_IsAnyItemHovered           = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsAnyItemHovered"}
	imGui_IsAnyItemActive            = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsAnyItemActive"}
	imGui_IsAnyItemFocused           = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsAnyItemFocused"}
	imGui_GetItemID                  = cc.DlFunc[func() ImGuiID, ImGuiID]{Name: "ImGui_GetItemID"}
	imGui_GetItemRectMin             = cc.DlFunc[func() ImVec2, ImVec2]{Name: "ImGui_GetItemRectMin"}
	imGui_GetItemRectMax             = cc.DlFunc[func() ImVec2, ImVec2]{Name: "ImGui_GetItemRectMax"}
	imGui_GetItemRectSize            = cc.DlFunc[func() ImVec2, ImVec2]{Name: "ImGui_GetItemRectSize"}

	imGui_GetMainViewport         = cc.DlFunc[func() *ImGuiViewport, *ImGuiViewport]{Name: "ImGui_GetMainViewport"}
	imGui_GetBackgroundDrawList   = cc.DlFunc[func(*ImGuiViewport) *ImDrawList, *ImDrawList]{Name: "ImGui_GetBackgroundDrawList"}
	imGui_GetForegroundDrawList   = cc.DlFunc[func(*ImGuiViewport) *ImDrawList, *ImDrawList]{Name: "ImGui_GetForegroundDrawList"}
	imGui_IsRectVisibleBySize     = cc.DlFunc[func(ImVec2) cc.Bool, cc.Bool]{Name: "ImGui_IsRectVisibleBySize"}
	imGui_IsRectVisible           = cc.DlFunc[func(ImVec2, ImVec2) cc.Bool, cc.Bool]{Name: "ImGui_IsRectVisible"}
	imGui_GetTime                 = cc.DlFunc[func() float64, float64]{Name: "ImGui_GetTime"}
	imGui_GetFrameCount           = cc.DlFunc[func() int32, int32]{Name: "ImGui_GetFrameCount"}
	imGui_GetDrawListSharedData   = cc.DlFunc[func() *ImDrawListSharedData, *ImDrawListSharedData]{Name: "ImGui_GetDrawListSharedData"}
	imGui_GetStyleColorName       = cc.DlFunc[func(ImGuiCol) cc.String, cc.String]{Name: "ImGui_GetStyleColorName"}
	imGui_SetStateStorage         = cc.DlFunc[func(*ImGuiStorage), cc.Void]{Name: "ImGui_SetStateStorage"}
	imGui_GetStateStorage         = cc.DlFunc[func() *ImGuiStorage, *ImGuiStorage]{Name: "ImGui_GetStateStorage"}
	imGui_CalcTextSize            = cc.DlFunc[func(cc.String, cc.String, cc.Bool, float32) ImVec2, ImVec2]{Name: "ImGui_CalcTextSize"}
	imGui_ColorConvertU32ToFloat4 = cc.DlFunc[func(ImU32) ImVec4, ImVec4]{Name: "ImGui_ColorConvertU32ToFloat4"}
	imGui_ColorConvertFloat4ToU32 = cc.DlFunc[func(ImVec4) ImU32, ImU32]{Name: "ImGui_ColorConvertFloat4ToU32"}
	imGui_ColorConvertRGBtoHSV    = cc.DlFunc[func(float32, float32, float32, *float32, *float32, *float32), cc.Void]{Name: "ImGui_ColorConvertRGBtoHSV"}
	imGui_ColorConvertHSVtoRGB    = cc.DlFunc[func(float32, float32, float32, *float32, *float32, *float32), cc.Void]{Name: "ImGui_ColorConvertHSVtoRGB"}

	imGui_IsKeyDown                       = cc.DlFunc[func(ImGuiKey) cc.Bool, cc.Bool]{Name: "ImGui_IsKeyDown"}
	imGui_IsKeyPressed                    = cc.DlFunc[func(ImGuiKey, cc.Bool) cc.Bool, cc.Bool]{Name: "ImGui_IsKeyPressed"}
	imGui_IsKeyReleased                   = cc.DlFunc[func(ImGuiKey) cc.Bool, cc.Bool]{Name: "ImGui_IsKeyReleased"}
	imGui_IsKeyChordPressed               = cc.DlFunc[func(ImGuiKeyChord) cc.Bool, cc.Bool]{Name: "ImGui_IsKeyChordPressed"}
	imGui_GetKeyPressedAmount             = cc.DlFunc[func(ImGuiKey, float32, float32) int32, int32]{Name: "ImGui_GetKeyPressedAmount"}
	imGui_GetKeyName                      = cc.DlFunc[func(ImGuiKey) cc.String, cc.String]{Name: "ImGui_GetKeyName"}
	imGui_SetNextFrameWantCaptureKeyboard = cc.DlFunc[func(cc.Bool), cc.Void]{Name: "ImGui_SetNextFrameWantCaptureKeyboard"}
	imGui_Shortcut                        = cc.DlFunc[func(ImGuiKeyChord, ImGuiInputFlags) cc.Bool, cc.Bool]{Name: "ImGui_Shortcut"}
	imGui_SetNextItemShortcut             = cc.DlFunc[func(ImGuiKeyChord, ImGuiInputFlags), cc.Void]{Name: "ImGui_SetNextItemShortcut"}
	imGui_SetItemKeyOwner                 = cc.DlFunc[func(ImGuiKey), cc.Void]{Name: "ImGui_SetItemKeyOwner"}

	imGui_IsMouseDown                      = cc.DlFunc[func(ImGuiMouseButton) cc.Bool, cc.Bool]{Name: "ImGui_IsMouseDown"}
	imGui_IsMouseClicked                   = cc.DlFunc[func(ImGuiMouseButton, cc.Bool) cc.Bool, cc.Bool]{Name: "ImGui_IsMouseClicked"}
	imGui_IsMouseReleased                  = cc.DlFunc[func(ImGuiMouseButton) cc.Bool, cc.Bool]{Name: "ImGui_IsMouseReleased"}
	imGui_IsMouseDoubleClicked             = cc.DlFunc[func(ImGuiMouseButton) cc.Bool, cc.Bool]{Name: "ImGui_IsMouseDoubleClicked"}
	imGui_IsMouseReleasedWithDelay         = cc.DlFunc[func(ImGuiMouseButton, float32) cc.Bool, cc.Bool]{Name: "ImGui_IsMouseReleasedWithDelay"}
	imGui_GetMouseClickedCount             = cc.DlFunc[func(ImGuiMouseButton) int32, int32]{Name: "ImGui_GetMouseClickedCount"}
	imGui_IsMouseHoveringRect              = cc.DlFunc[func(ImVec2, ImVec2, cc.Bool) cc.Bool, cc.Bool]{Name: "ImGui_IsMouseHoveringRect"}
	imGui_IsMousePosValid                  = cc.DlFunc[func(*ImVec2) cc.Bool, cc.Bool]{Name: "ImGui_IsMousePosValid"}
	imGui_IsAnyMouseDown                   = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImGui_IsAnyMouseDown"}
	imGui_GetMousePos                      = cc.DlFunc[func() ImVec2, ImVec2]{Name: "ImGui_GetMousePos"}
	imGui_GetMousePosOnOpeningCurrentPopup = cc.DlFunc[func() ImVec2, ImVec2]{Name: "ImGui_GetMousePosOnOpeningCurrentPopup"}
	imGui_IsMouseDragging                  = cc.DlFunc[func(ImGuiMouseButton, float32) cc.Bool, cc.Bool]{Name: "ImGui_IsMouseDragging"}
	imGui_GetMouseDragDelta                = cc.DlFunc[func(ImGuiMouseButton, float32) ImVec2, ImVec2]{Name: "ImGui_GetMouseDragDelta"}
	imGui_ResetMouseDragDelta              = cc.DlFunc[func(ImGuiMouseButton), cc.Void]{Name: "ImGui_ResetMouseDragDelta"}
	imGui_GetMouseCursor                   = cc.DlFunc[func() ImGuiMouseCursor, ImGuiMouseCursor]{Name: "ImGui_GetMouseCursor"}
	imGui_SetMouseCursor                   = cc.DlFunc[func(ImGuiMouseCursor), cc.Void]{Name: "ImGui_SetMouseCursor"}
	imGui_SetNextFrameWantCaptureMouse     = cc.DlFunc[func(cc.Bool), cc.Void]{Name: "ImGui_SetNextFrameWantCaptureMouse"}
	imGui_GetClipboardText                 = cc.DlFunc[func() cc.String, cc.String]{Name: "ImGui_GetClipboardText"}
	imGui_SetClipboardText                 = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_SetClipboardText"}

	imGui_LoadIniSettingsFromDisk        = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_LoadIniSettingsFromDisk"}
	imGui_LoadIniSettingsFromMemory      = cc.DlFunc[func(*byte, uint64), cc.Void]{Name: "ImGui_LoadIniSettingsFromMemory"}
	imGui_SaveIniSettingsToDisk          = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_SaveIniSettingsToDisk"}
	imGui_SaveIniSettingsToMemory        = cc.DlFunc[func(*uint64) *byte, *byte]{Name: "ImGui_SaveIniSettingsToMemory"}
	imGui_DebugTextEncoding              = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_DebugTextEncoding"}
	imGui_DebugFlashStyleColor           = cc.DlFunc[func(ImGuiCol), cc.Void]{Name: "ImGui_DebugFlashStyleColor"}
	imGui_DebugStartItemPicker           = cc.DlFunc[func(), cc.Void]{Name: "ImGui_DebugStartItemPicker"}
	imGui_DebugCheckVersionAndDataLayout = cc.DlFunc[func(cc.String, cc.SizeT, cc.SizeT, cc.SizeT, cc.SizeT, cc.SizeT, cc.SizeT) cc.Bool, cc.Bool]{Name: "ImGui_DebugCheckVersionAndDataLayout"}
	imGui_DebugLog                       = cc.DlFunc[func(cc.String, ...interface{}), cc.Void]{Name: "ImGui_DebugLog", Va: true}
	imGui_DebugLogUnformatted            = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImGui_DebugLogUnformatted"}
	// imGui_DebugLogV                      = cc.DlFunc[func(cc.String, unsafe.Pointer), cc.Void]{Name: "ImGui_DebugLogV", Va: true}
	imGui_SetAllocatorFunctions = cc.DlFunc[func(cc.Func, cc.Func, unsafe.Pointer), cc.Void]{Name: "ImGui_SetAllocatorFunctions"}
	// imGui_GetAllocatorFunctions        = cc.DlFunc[func(*cc.Func, *cc.Func, *unsafe.Pointer), cc.Void]{Name: "ImGui_GetAllocatorFunctions"}
	imGui_MemAlloc                     = cc.DlFunc[func(cc.SizeT) unsafe.Pointer, unsafe.Pointer]{Name: "ImGui_MemAlloc"}
	imGui_MemFree                      = cc.DlFunc[func(unsafe.Pointer), cc.Void]{Name: "ImGui_MemFree"}
	imGui_UpdatePlatformWindows        = cc.DlFunc[func(), cc.Void]{Name: "ImGui_UpdatePlatformWindows"}
	imGui_RenderPlatformWindowsDefault = cc.DlFunc[func(unsafe.Pointer, unsafe.Pointer), cc.Void]{Name: "ImGui_RenderPlatformWindowsDefault"}
	imGui_DestroyPlatformWindows       = cc.DlFunc[func(), cc.Void]{Name: "ImGui_DestroyPlatformWindows"}
	imGui_FindViewportByID             = cc.DlFunc[func(ImGuiID) *ImGuiViewport, *ImGuiViewport]{Name: "ImGui_FindViewportByID"}
	imGui_FindViewportByPlatformHandle = cc.DlFunc[func(unsafe.Pointer) *ImGuiViewport, *ImGuiViewport]{Name: "ImGui_FindViewportByPlatformHandle"}

	imGuiStyle_ScaleAllSizes = cc.DlFunc[func(*ImGuiStyle, float32), cc.Void]{Name: "ImGuiStyle_ScaleAllSizes"}

	imGuiIO_AddKeyEvent            = cc.DlFunc[func(*ImGuiIO, ImGuiKey, cc.Bool), cc.Void]{Name: "ImGuiIO_AddKeyEvent"}
	imGuiIO_AddKeyAnalogEvent      = cc.DlFunc[func(*ImGuiIO, ImGuiKey, cc.Bool, float32), cc.Void]{Name: "ImGuiIO_AddKeyAnalogEvent"}
	imGuiIO_AddMousePosEvent       = cc.DlFunc[func(*ImGuiIO, float32, float32), cc.Void]{Name: "ImGuiIO_AddMousePosEvent"}
	imGuiIO_AddMouseButtonEvent    = cc.DlFunc[func(*ImGuiIO, int32, cc.Bool), cc.Void]{Name: "ImGuiIO_AddMouseButtonEvent"}
	imGuiIO_AddMouseWheelEvent     = cc.DlFunc[func(*ImGuiIO, float32, float32), cc.Void]{Name: "ImGuiIO_AddMouseWheelEvent"}
	imGuiIO_AddMouseSourceEvent    = cc.DlFunc[func(*ImGuiIO, ImGuiMouseSource), cc.Void]{Name: "ImGuiIO_AddMouseSourceEvent"}
	imGuiIO_AddMouseViewportEvent  = cc.DlFunc[func(*ImGuiIO, ImGuiID), cc.Void]{Name: "ImGuiIO_AddMouseViewportEvent"}
	imGuiIO_AddFocusEvent          = cc.DlFunc[func(*ImGuiIO, cc.Bool), cc.Void]{Name: "ImGuiIO_AddFocusEvent"}
	imGuiIO_AddInputCharacter      = cc.DlFunc[func(*ImGuiIO, uint32), cc.Void]{Name: "ImGuiIO_AddInputCharacter"}
	imGuiIO_AddInputCharacterUTF16 = cc.DlFunc[func(*ImGuiIO, ImWchar16), cc.Void]{Name: "ImGuiIO_AddInputCharacterUTF16"}
	imGuiIO_AddInputCharactersUTF8 = cc.DlFunc[func(*ImGuiIO, cc.String), cc.Void]{Name: "ImGuiIO_AddInputCharactersUTF8"}
	imGuiIO_SetKeyEventNativeData  = cc.DlFunc[func(*ImGuiIO, ImGuiKey, int32, int32, int32), cc.Void]{Name: "ImGuiIO_SetKeyEventNativeData"}
	imGuiIO_SetAppAcceptingEvents  = cc.DlFunc[func(*ImGuiIO, cc.Bool), cc.Void]{Name: "ImGuiIO_SetAppAcceptingEvents"}
	imGuiIO_ClearEventsQueue       = cc.DlFunc[func(*ImGuiIO), cc.Void]{Name: "ImGuiIO_ClearEventsQueue"}
	imGuiIO_ClearInputKeys         = cc.DlFunc[func(*ImGuiIO), cc.Void]{Name: "ImGuiIO_ClearInputKeys"}
	imGuiIO_ClearInputMouse        = cc.DlFunc[func(*ImGuiIO), cc.Void]{Name: "ImGuiIO_ClearInputMouse"}
	imGuiIO_ClearInputCharacters   = cc.DlFunc[func(*ImGuiIO), cc.Void]{Name: "ImGuiIO_ClearInputCharacters"}

	imGuiInputTextCallbackData_DeleteChars    = cc.DlFunc[func(*ImGuiInputTextCallbackData, int32, int32), cc.Void]{Name: "ImGuiInputTextCallbackData_DeleteChars"}
	imGuiInputTextCallbackData_InsertChars    = cc.DlFunc[func(*ImGuiInputTextCallbackData, int32, cc.String, cc.String), cc.Void]{Name: "ImGuiInputTextCallbackData_InsertChars"}
	imGuiInputTextCallbackData_SelectAll      = cc.DlFunc[func(*ImGuiInputTextCallbackData), cc.Void]{Name: "ImGuiInputTextCallbackData_SelectAll"}
	imGuiInputTextCallbackData_ClearSelection = cc.DlFunc[func(*ImGuiInputTextCallbackData), cc.Void]{Name: "ImGuiInputTextCallbackData_ClearSelection"}
	imGuiInputTextCallbackData_HasSelection   = cc.DlFunc[func(*ImGuiInputTextCallbackData) cc.Bool, cc.Bool]{Name: "ImGuiInputTextCallbackData_HasSelection"}

	imGuiPayload_Clear      = cc.DlFunc[func(*ImGuiPayload), cc.Void]{Name: "ImGuiPayload_Clear"}
	imGuiPayload_IsDataType = cc.DlFunc[func(*ImGuiPayload, cc.String) cc.Bool, cc.Bool]{Name: "ImGuiPayload_IsDataType"}
	imGuiPayload_IsPreview  = cc.DlFunc[func(*ImGuiPayload) cc.Bool, cc.Bool]{Name: "ImGuiPayload_IsPreview"}
	imGuiPayload_IsDelivery = cc.DlFunc[func(*ImGuiPayload) cc.Bool, cc.Bool]{Name: "ImGuiPayload_IsDelivery"}

	imGuiTextFilter_ImGuiTextRange_empty = cc.DlFunc[func(*ImGuiTextRange) cc.Bool, cc.Bool]{Name: "ImGuiTextFilter_ImGuiTextRange_empty"}
	imGuiTextFilter_ImGuiTextRange_split = cc.DlFunc[func(*ImGuiTextRange, int32, *ImVector[ImGuiTextRange]), cc.Void]{Name: "ImGuiTextFilter_ImGuiTextRange_split"}
	imGuiTextFilter_Draw                 = cc.DlFunc[func(*ImGuiTextFilter, cc.String, float32) cc.Bool, cc.Bool]{Name: "ImGuiTextFilter_Draw"}
	imGuiTextFilter_PassFilter           = cc.DlFunc[func(*ImGuiTextFilter, cc.String, cc.String) cc.Bool, cc.Bool]{Name: "ImGuiTextFilter_PassFilter"}
	imGuiTextFilter_Build                = cc.DlFunc[func(*ImGuiTextFilter), cc.Void]{Name: "ImGuiTextFilter_Build"}
	imGuiTextFilter_Clear                = cc.DlFunc[func(*ImGuiTextFilter), cc.Void]{Name: "ImGuiTextFilter_Clear"}
	imGuiTextFilter_IsActive             = cc.DlFunc[func(*ImGuiTextFilter) cc.Bool, cc.Bool]{Name: "ImGuiTextFilter_IsActive"}

	imGuiTextBuffer_begin   = cc.DlFunc[func(*ImGuiTextBuffer) cc.String, cc.String]{Name: "ImGuiTextBuffer_begin"}
	imGuiTextBuffer_end     = cc.DlFunc[func(*ImGuiTextBuffer) cc.String, cc.String]{Name: "ImGuiTextBuffer_end"}
	imGuiTextBuffer_size    = cc.DlFunc[func(*ImGuiTextBuffer) int32, int32]{Name: "ImGuiTextBuffer_size"}
	imGuiTextBuffer_empty   = cc.DlFunc[func(*ImGuiTextBuffer) cc.Bool, cc.Bool]{Name: "ImGuiTextBuffer_empty"}
	imGuiTextBuffer_clear   = cc.DlFunc[func(*ImGuiTextBuffer), cc.Void]{Name: "ImGuiTextBuffer_clear"}
	imGuiTextBuffer_resize  = cc.DlFunc[func(*ImGuiTextBuffer, int32), cc.Void]{Name: "ImGuiTextBuffer_resize"}
	imGuiTextBuffer_reserve = cc.DlFunc[func(*ImGuiTextBuffer, int32), cc.Void]{Name: "ImGuiTextBuffer_reserve"}
	imGuiTextBuffer_c_str   = cc.DlFunc[func(*ImGuiTextBuffer) cc.String, cc.String]{Name: "ImGuiTextBuffer_c_str"}
	imGuiTextBuffer_append  = cc.DlFunc[func(*ImGuiTextBuffer, cc.String, cc.String), cc.Void]{Name: "ImGuiTextBuffer_append"}
	imGuiTextBuffer_appendf = cc.DlFunc[func(*ImGuiTextBuffer, cc.String, ...interface{}), cc.Void]{Name: "ImGuiTextBuffer_appendf", Va: true}
	// imGuiTextBuffer_appendfv = cc.DlFunc[func(*ImGuiTextBuffer, cc.String, unsafe.Pointer), cc.Void]{Name: "ImGuiTextBuffer_appendfv", Va: true}

	imGuiStorage_Clear          = cc.DlFunc[func(*ImGuiStorage), cc.Void]{Name: "ImGuiStorage_Clear"}
	imGuiStorage_GetInt         = cc.DlFunc[func(*ImGuiStorage, ImGuiID, int32) int32, int32]{Name: "ImGuiStorage_GetInt"}
	imGuiStorage_SetInt         = cc.DlFunc[func(*ImGuiStorage, ImGuiID, int32), cc.Void]{Name: "ImGuiStorage_SetInt"}
	imGuiStorage_GetBool        = cc.DlFunc[func(*ImGuiStorage, ImGuiID, cc.Bool) cc.Bool, cc.Bool]{Name: "ImGuiStorage_GetBool"}
	imGuiStorage_SetBool        = cc.DlFunc[func(*ImGuiStorage, ImGuiID, cc.Bool), cc.Void]{Name: "ImGuiStorage_SetBool"}
	imGuiStorage_GetFloat       = cc.DlFunc[func(*ImGuiStorage, ImGuiID, float32) float32, float32]{Name: "ImGuiStorage_GetFloat"}
	imGuiStorage_SetFloat       = cc.DlFunc[func(*ImGuiStorage, ImGuiID, float32), cc.Void]{Name: "ImGuiStorage_SetFloat"}
	imGuiStorage_GetVoidPtr     = cc.DlFunc[func(*ImGuiStorage, ImGuiID) unsafe.Pointer, unsafe.Pointer]{Name: "ImGuiStorage_GetVoidPtr"}
	imGuiStorage_SetVoidPtr     = cc.DlFunc[func(*ImGuiStorage, ImGuiID, unsafe.Pointer), cc.Void]{Name: "ImGuiStorage_SetVoidPtr"}
	imGuiStorage_GetIntRef      = cc.DlFunc[func(*ImGuiStorage, ImGuiID, int32) *int32, *int32]{Name: "ImGuiStorage_GetIntRef"}
	imGuiStorage_GetBoolRef     = cc.DlFunc[func(*ImGuiStorage, ImGuiID, cc.Bool) *cc.Bool, *cc.Bool]{Name: "ImGuiStorage_GetBoolRef"}
	imGuiStorage_GetFloatRef    = cc.DlFunc[func(*ImGuiStorage, ImGuiID, float32) *float32, *float32]{Name: "ImGuiStorage_GetFloatRef"}
	imGuiStorage_GetVoidPtrRef  = cc.DlFunc[func(*ImGuiStorage, ImGuiID, unsafe.Pointer) **unsafe.Pointer, **unsafe.Pointer]{Name: "ImGuiStorage_GetVoidPtrRef"}
	imGuiStorage_BuildSortByKey = cc.DlFunc[func(*ImGuiStorage), cc.Void]{Name: "ImGuiStorage_BuildSortByKey"}
	imGuiStorage_SetAllInt      = cc.DlFunc[func(*ImGuiStorage, int32), cc.Void]{Name: "ImGuiStorage_SetAllInt"}

	imGuiListClipper_Begin                 = cc.DlFunc[func(*ImGuiListClipper, int32, float32), cc.Void]{Name: "ImGuiListClipper_Begin"}
	imGuiListClipper_End                   = cc.DlFunc[func(*ImGuiListClipper), cc.Void]{Name: "ImGuiListClipper_End"}
	imGuiListClipper_Step                  = cc.DlFunc[func(*ImGuiListClipper) cc.Bool, cc.Bool]{Name: "ImGuiListClipper_Step"}
	imGuiListClipper_IncludeItemByIndex    = cc.DlFunc[func(*ImGuiListClipper, int32), cc.Void]{Name: "ImGuiListClipper_IncludeItemByIndex"}
	imGuiListClipper_IncludeItemsByIndex   = cc.DlFunc[func(*ImGuiListClipper, int32, int32), cc.Void]{Name: "ImGuiListClipper_IncludeItemsByIndex"}
	imGuiListClipper_SeekCursorForItem     = cc.DlFunc[func(*ImGuiListClipper, int32), cc.Void]{Name: "ImGuiListClipper_SeekCursorForItem"}
	imGuiListClipper_IncludeRangeByIndices = cc.DlFunc[func(*ImGuiListClipper, int32, int32), cc.Void]{Name: "ImGuiListClipper_IncludeRangeByIndices"}

	imColor_SetHSV = cc.DlFunc[func(*ImColor, float32, float32, float32, float32), cc.Void]{Name: "ImColor_SetHSV"}
	imColor_HSV    = cc.DlFunc[func(float32, float32, float32, float32) ImColor, ImColor]{Name: "ImColor_HSV"}

	imGuiSelectionBasicStorage_ApplyRequests         = cc.DlFunc[func(*ImGuiSelectionBasicStorage, *ImGuiMultiSelectIO), cc.Void]{Name: "ImGuiSelectionBasicStorage_ApplyRequests"}
	imGuiSelectionBasicStorage_Contains              = cc.DlFunc[func(*ImGuiSelectionBasicStorage, ImGuiID) cc.Bool, cc.Bool]{Name: "ImGuiSelectionBasicStorage_Contains"}
	imGuiSelectionBasicStorage_Clear                 = cc.DlFunc[func(*ImGuiSelectionBasicStorage), cc.Void]{Name: "ImGuiSelectionBasicStorage_Clear"}
	imGuiSelectionBasicStorage_Swap                  = cc.DlFunc[func(*ImGuiSelectionBasicStorage, *ImGuiSelectionBasicStorage), cc.Void]{Name: "ImGuiSelectionBasicStorage_Swap"}
	imGuiSelectionBasicStorage_SetItemSelected       = cc.DlFunc[func(*ImGuiSelectionBasicStorage, ImGuiID, cc.Bool), cc.Void]{Name: "ImGuiSelectionBasicStorage_SetItemSelected"}
	imGuiSelectionBasicStorage_GetNextSelectedItem   = cc.DlFunc[func(*ImGuiSelectionBasicStorage, *unsafe.Pointer, *ImGuiID) cc.Bool, cc.Bool]{Name: "ImGuiSelectionBasicStorage_GetNextSelectedItem"}
	imGuiSelectionBasicStorage_GetStorageIdFromIndex = cc.DlFunc[func(*ImGuiSelectionBasicStorage, int32) ImGuiID, ImGuiID]{Name: "ImGuiSelectionBasicStorage_GetStorageIdFromIndex"}
	imGuiSelectionExternalStorage_ApplyRequests      = cc.DlFunc[func(*ImGuiSelectionExternalStorage, *ImGuiMultiSelectIO), cc.Void]{Name: "ImGuiSelectionExternalStorage_ApplyRequests"}

	imDrawCmd_GetTexID                   = cc.DlFunc[func(*ImDrawCmd) ImTextureID, ImTextureID]{Name: "ImDrawCmd_GetTexID"}
	imDrawListSplitter_Clear             = cc.DlFunc[func(*ImDrawListSplitter), cc.Void]{Name: "ImDrawListSplitter_Clear"}
	imDrawListSplitter_ClearFreeMemory   = cc.DlFunc[func(*ImDrawListSplitter), cc.Void]{Name: "ImDrawListSplitter_ClearFreeMemory"}
	imDrawListSplitter_Split             = cc.DlFunc[func(*ImDrawListSplitter, *ImDrawList, int32), cc.Void]{Name: "ImDrawListSplitter_Split"}
	imDrawListSplitter_Merge             = cc.DlFunc[func(*ImDrawListSplitter, *ImDrawList), cc.Void]{Name: "ImDrawListSplitter_Merge"}
	imDrawListSplitter_SetCurrentChannel = cc.DlFunc[func(*ImDrawListSplitter, *ImDrawList, int32), cc.Void]{Name: "ImDrawListSplitter_SetCurrentChannel"}

	imDrawList_PushClipRect            = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, cc.Bool), cc.Void]{Name: "ImDrawList_PushClipRect"}
	imDrawList_PushClipRectFullScreen  = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList_PushClipRectFullScreen"}
	imDrawList_PopClipRect             = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList_PopClipRect"}
	imDrawList_PushTexture             = cc.DlFunc[func(*ImDrawList, ImTextureRef), cc.Void]{Name: "ImDrawList_PushTexture"}
	imDrawList_PopTexture              = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList_PopTexture"}
	imDrawList_GetClipRectMin          = cc.DlFunc[func(*ImDrawList) ImVec2, ImVec2]{Name: "ImDrawList_GetClipRectMin"}
	imDrawList_GetClipRectMax          = cc.DlFunc[func(*ImDrawList) ImVec2, ImVec2]{Name: "ImDrawList_GetClipRectMax"}
	imDrawList_AddLine                 = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImU32, float32), cc.Void]{Name: "ImDrawList_AddLine"}
	imDrawList_AddRect                 = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImU32, float32, ImDrawFlags, float32), cc.Void]{Name: "ImDrawList_AddRect"}
	imDrawList_AddRectFilled           = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImU32, float32, ImDrawFlags), cc.Void]{Name: "ImDrawList_AddRectFilled"}
	imDrawList_AddRectFilledMultiColor = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImU32, ImU32, ImU32, ImU32), cc.Void]{Name: "ImDrawList_AddRectFilledMultiColor"}
	imDrawList_AddQuad                 = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImVec2, ImVec2, ImU32, float32), cc.Void]{Name: "ImDrawList_AddQuad"}
	imDrawList_AddQuadFilled           = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImVec2, ImVec2, ImU32), cc.Void]{Name: "ImDrawList_AddQuadFilled"}
	imDrawList_AddTriangle             = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImVec2, ImU32, float32), cc.Void]{Name: "ImDrawList_AddTriangle"}
	imDrawList_AddTriangleFilled       = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImVec2, ImU32), cc.Void]{Name: "ImDrawList_AddTriangleFilled"}
	imDrawList_AddCircle               = cc.DlFunc[func(*ImDrawList, ImVec2, float32, ImU32, int32, float32), cc.Void]{Name: "ImDrawList_AddCircle"}
	imDrawList_AddCircleFilled         = cc.DlFunc[func(*ImDrawList, ImVec2, float32, ImU32, int32), cc.Void]{Name: "ImDrawList_AddCircleFilled"}
	imDrawList_AddNgon                 = cc.DlFunc[func(*ImDrawList, ImVec2, float32, ImU32, int32, float32), cc.Void]{Name: "ImDrawList_AddNgon"}
	imDrawList_AddNgonFilled           = cc.DlFunc[func(*ImDrawList, ImVec2, float32, ImU32, int32), cc.Void]{Name: "ImDrawList_AddNgonFilled"}
	imDrawList_AddEllipse              = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImU32, float32, int32, float32), cc.Void]{Name: "ImDrawList_AddEllipse"}
	imDrawList_AddEllipseFilled        = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImU32, float32, int32), cc.Void]{Name: "ImDrawList_AddEllipseFilled"}
	imDrawList_AddText                 = cc.DlFunc[func(*ImDrawList, ImVec2, ImU32, cc.String, cc.String), cc.Void]{Name: "ImDrawList_AddText"}
	imDrawList_AddTextImFontPtr        = cc.DlFunc[func(*ImDrawList, *ImFont, float32, ImVec2, ImU32, cc.String, cc.String, float32, *ImVec4), cc.Void]{Name: "ImDrawList_AddTextImFontPtr"}
	imDrawList_AddBezierCubic          = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImVec2, ImVec2, ImU32, float32, int32), cc.Void]{Name: "ImDrawList_AddBezierCubic"}
	imDrawList_AddBezierQuadratic      = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImVec2, ImU32, float32, int32), cc.Void]{Name: "ImDrawList_AddBezierQuadratic"}

	imDrawList_AddPolyline                = cc.DlFunc[func(*ImDrawList, *ImVec2, int32, ImU32, ImDrawFlags, float32), cc.Void]{Name: "ImDrawList_AddPolyline"}
	imDrawList_AddConvexPolyFilled        = cc.DlFunc[func(*ImDrawList, *ImVec2, int32, ImU32), cc.Void]{Name: "ImDrawList_AddConvexPolyFilled"}
	imDrawList_AddConcavePolyFilled       = cc.DlFunc[func(*ImDrawList, *ImVec2, int32, ImU32), cc.Void]{Name: "ImDrawList_AddConcavePolyFilled"}
	imDrawList_AddImage                   = cc.DlFunc[func(*ImDrawList, ImTextureRef, ImVec2, ImVec2, ImVec2, ImVec2, ImU32), cc.Void]{Name: "ImDrawList_AddImage"}
	imDrawList_AddImageQuad               = cc.DlFunc[func(*ImDrawList, ImTextureRef, ImVec2, ImVec2, ImVec2, ImVec2, ImVec2, ImVec2, ImVec2, ImVec2, ImU32), cc.Void]{Name: "ImDrawList_AddImageQuad"}
	imDrawList_AddImageRounded            = cc.DlFunc[func(*ImDrawList, ImTextureRef, ImVec2, ImVec2, ImVec2, ImVec2, ImU32, float32, ImDrawFlags), cc.Void]{Name: "ImDrawList_AddImageRounded"}
	imDrawList_PathClear                  = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList_PathClear"}
	imDrawList_PathLineTo                 = cc.DlFunc[func(*ImDrawList, ImVec2), cc.Void]{Name: "ImDrawList_PathLineTo"}
	imDrawList_PathLineToMergeDuplicate   = cc.DlFunc[func(*ImDrawList, ImVec2), cc.Void]{Name: "ImDrawList_PathLineToMergeDuplicate"}
	imDrawList_PathFillConvex             = cc.DlFunc[func(*ImDrawList, ImU32), cc.Void]{Name: "ImDrawList_PathFillConvex"}
	imDrawList_PathFillConcave            = cc.DlFunc[func(*ImDrawList, ImU32), cc.Void]{Name: "ImDrawList_PathFillConcave"}
	imDrawList_PathStroke                 = cc.DlFunc[func(*ImDrawList, ImU32, ImDrawFlags, float32), cc.Void]{Name: "ImDrawList_PathStroke"}
	imDrawList_PathArcTo                  = cc.DlFunc[func(*ImDrawList, ImVec2, float32, float32, float32, int32), cc.Void]{Name: "ImDrawList_PathArcTo"}
	imDrawList_PathArcToFast              = cc.DlFunc[func(*ImDrawList, ImVec2, float32, int32, int32), cc.Void]{Name: "ImDrawList_PathArcToFast"}
	imDrawList_PathEllipticalArcTo        = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, float32, float32, float32, int32), cc.Void]{Name: "ImDrawList_PathEllipticalArcTo"}
	imDrawList_PathBezierCubicCurveTo     = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImVec2, int32), cc.Void]{Name: "ImDrawList_PathBezierCubicCurveTo"}
	imDrawList_PathBezierQuadraticCurveTo = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, int32), cc.Void]{Name: "ImDrawList_PathBezierQuadraticCurveTo"}
	imDrawList_PathRect                   = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, float32, ImDrawFlags), cc.Void]{Name: "ImDrawList_PathRect"}

	imDrawList_AddCallback                 = cc.DlFunc[func(*ImDrawList, cc.Func, unsafe.Pointer, cc.SizeT), cc.Void]{Name: "ImDrawList_AddCallback"}
	imDrawList_AddDrawCmd                  = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList_AddDrawCmd"}
	imDrawList_CloneOutput                 = cc.DlFunc[func(*ImDrawList) *ImDrawList, *ImDrawList]{Name: "ImDrawList_CloneOutput"}
	imDrawList_ChannelsSplit               = cc.DlFunc[func(*ImDrawList, int32), cc.Void]{Name: "ImDrawList_ChannelsSplit"}
	imDrawList_ChannelsMerge               = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList_ChannelsMerge"}
	imDrawList_ChannelsSetCurrent          = cc.DlFunc[func(*ImDrawList, int32), cc.Void]{Name: "ImDrawList_ChannelsSetCurrent"}
	imDrawList_PrimReserve                 = cc.DlFunc[func(*ImDrawList, int32, int32), cc.Void]{Name: "ImDrawList_PrimReserve"}
	imDrawList_PrimUnreserve               = cc.DlFunc[func(*ImDrawList, int32, int32), cc.Void]{Name: "ImDrawList_PrimUnreserve"}
	imDrawList_PrimRect                    = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImU32), cc.Void]{Name: "ImDrawList_PrimRect"}
	imDrawList_PrimRectUV                  = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImVec2, ImVec2, ImU32), cc.Void]{Name: "ImDrawList_PrimRectUV"}
	imDrawList_PrimQuadUV                  = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImVec2, ImVec2, ImVec2, ImVec2, ImVec2, ImVec2, ImU32), cc.Void]{Name: "ImDrawList_PrimQuadUV"}
	imDrawList_PrimWriteVtx                = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImU32), cc.Void]{Name: "ImDrawList_PrimWriteVtx"}
	imDrawList_PrimWriteIdx                = cc.DlFunc[func(*ImDrawList, ImDrawIdx), cc.Void]{Name: "ImDrawList_PrimWriteIdx"}
	imDrawList_PrimVtx                     = cc.DlFunc[func(*ImDrawList, ImVec2, ImVec2, ImU32), cc.Void]{Name: "ImDrawList_PrimVtx"}
	imDrawList_PushTextureID               = cc.DlFunc[func(*ImDrawList, ImTextureRef), cc.Void]{Name: "ImDrawList_PushTextureID"}
	imDrawList_PopTextureID                = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList_PopTextureID"}
	imDrawList__SetDrawListSharedData      = cc.DlFunc[func(*ImDrawList, *ImDrawListSharedData), cc.Void]{Name: "ImDrawList__SetDrawListSharedData"}
	imDrawList__ResetForNewFrame           = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList__ResetForNewFrame"}
	imDrawList__ClearFreeMemory            = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList__ClearFreeMemory"}
	imDrawList__PopUnusedDrawCmd           = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList__PopUnusedDrawCmd"}
	imDrawList__TryMergeDrawCmds           = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList__TryMergeDrawCmds"}
	imDrawList__OnChangedClipRect          = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList__OnChangedClipRect"}
	imDrawList__OnChangedTexture           = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList__OnChangedTexture"}
	imDrawList__OnChangedVtxOffset         = cc.DlFunc[func(*ImDrawList), cc.Void]{Name: "ImDrawList__OnChangedVtxOffset"}
	imDrawList__SetTexture                 = cc.DlFunc[func(*ImDrawList, ImTextureRef), cc.Void]{Name: "ImDrawList__SetTexture"}
	imDrawList__CalcCircleAutoSegmentCount = cc.DlFunc[func(*ImDrawList, float32) int32, int32]{Name: "ImDrawList__CalcCircleAutoSegmentCount"}
	imDrawList__PathArcToFastEx            = cc.DlFunc[func(*ImDrawList, ImVec2, float32, int32, int32, int32), cc.Void]{Name: "ImDrawList__PathArcToFastEx"}
	imDrawList__PathArcToN                 = cc.DlFunc[func(*ImDrawList, ImVec2, float32, float32, float32, int32), cc.Void]{Name: "ImDrawList__PathArcToN"}

	imDrawData_Clear             = cc.DlFunc[func(*ImDrawData), cc.Void]{Name: "ImDrawData_Clear"}
	imDrawData_AddDrawList       = cc.DlFunc[func(*ImDrawData, *ImDrawList), cc.Void]{Name: "ImDrawData_AddDrawList"}
	imDrawData_DeIndexAllBuffers = cc.DlFunc[func(*ImDrawData), cc.Void]{Name: "ImDrawData_DeIndexAllBuffers"}
	imDrawData_ScaleClipRects    = cc.DlFunc[func(*ImDrawData, ImVec2), cc.Void]{Name: "ImDrawData_ScaleClipRects"}

	imTextureData_Create         = cc.DlFunc[func(*ImTextureData, ImTextureFormat, int32, int32), cc.Void]{Name: "ImTextureData_Create"}
	imTextureData_DestroyPixels  = cc.DlFunc[func(*ImTextureData), cc.Void]{Name: "ImTextureData_DestroyPixels"}
	imTextureData_GetPixels      = cc.DlFunc[func(*ImTextureData) unsafe.Pointer, unsafe.Pointer]{Name: "ImTextureData_GetPixels"}
	imTextureData_GetPixelsAt    = cc.DlFunc[func(*ImTextureData, int32, int32) unsafe.Pointer, unsafe.Pointer]{Name: "ImTextureData_GetPixelsAt"}
	imTextureData_GetSizeInBytes = cc.DlFunc[func(*ImTextureData) int32, int32]{Name: "ImTextureData_GetSizeInBytes"}
	imTextureData_GetPitch       = cc.DlFunc[func(*ImTextureData) int32, int32]{Name: "ImTextureData_GetPitch"}
	imTextureData_GetTexRef      = cc.DlFunc[func(*ImTextureData) ImTextureRef, ImTextureRef]{Name: "ImTextureData_GetTexRef"}
	imTextureData_GetTexID       = cc.DlFunc[func(*ImTextureData) ImTextureID, ImTextureID]{Name: "ImTextureData_GetTexID"}
	imTextureData_SetTexID       = cc.DlFunc[func(*ImTextureData, ImTextureID), cc.Void]{Name: "ImTextureData_SetTexID"}
	imTextureData_SetStatus      = cc.DlFunc[func(*ImTextureData, ImTextureStatus), cc.Void]{Name: "ImTextureData_SetStatus"}

	imFontGlyphRangesBuilder_Clear       = cc.DlFunc[func(*ImFontGlyphRangesBuilder), cc.Void]{Name: "ImFontGlyphRangesBuilder_Clear"}
	imFontGlyphRangesBuilder_GetBit      = cc.DlFunc[func(*ImFontGlyphRangesBuilder, uint64) cc.Bool, cc.Bool]{Name: "ImFontGlyphRangesBuilder_GetBit"}
	imFontGlyphRangesBuilder_SetBit      = cc.DlFunc[func(*ImFontGlyphRangesBuilder, uint64), cc.Void]{Name: "ImFontGlyphRangesBuilder_SetBit"}
	imFontGlyphRangesBuilder_AddChar     = cc.DlFunc[func(*ImFontGlyphRangesBuilder, ImWchar), cc.Void]{Name: "ImFontGlyphRangesBuilder_AddChar"}
	imFontGlyphRangesBuilder_AddText     = cc.DlFunc[func(*ImFontGlyphRangesBuilder, cc.String, cc.String), cc.Void]{Name: "ImFontGlyphRangesBuilder_AddText"}
	imFontGlyphRangesBuilder_AddRanges   = cc.DlFunc[func(*ImFontGlyphRangesBuilder, *ImWchar), cc.Void]{Name: "ImFontGlyphRangesBuilder_AddRanges"}
	imFontGlyphRangesBuilder_BuildRanges = cc.DlFunc[func(*ImFontGlyphRangesBuilder, *ImVector[ImWchar]), cc.Void]{Name: "ImFontGlyphRangesBuilder_BuildRanges"}

	imFontAtlas_AddFont                              = cc.DlFunc[func(*ImFontAtlas, *ImFontConfig) *ImFont, *ImFont]{Name: "ImFontAtlas_AddFont"}
	imFontAtlas_AddFontDefault                       = cc.DlFunc[func(*ImFontAtlas, *ImFontConfig) *ImFont, *ImFont]{Name: "ImFontAtlas_AddFontDefault"}
	imFontAtlas_AddFontFromFileTTF                   = cc.DlFunc[func(*ImFontAtlas, cc.String, float32, *ImFontConfig, *ImWchar) *ImFont, *ImFont]{Name: "ImFontAtlas_AddFontFromFileTTF"}
	imFontAtlas_AddFontFromMemoryTTF                 = cc.DlFunc[func(*ImFontAtlas, *byte, int32, float32, *ImFontConfig, *ImWchar) *ImFont, *ImFont]{Name: "ImFontAtlas_AddFontFromMemoryTTF"}
	imFontAtlas_AddFontFromMemoryCompressedTTF       = cc.DlFunc[func(*ImFontAtlas, *byte, int32, float32, *ImFontConfig, *ImWchar) *ImFont, *ImFont]{Name: "ImFontAtlas_AddFontFromMemoryCompressedTTF"}
	imFontAtlas_AddFontFromMemoryCompressedBase85TTF = cc.DlFunc[func(*ImFontAtlas, cc.String, float32, *ImFontConfig, *ImWchar) *ImFont, *ImFont]{Name: "ImFontAtlas_AddFontFromMemoryCompressedBase85TTF"}
	imFontAtlas_RemoveFont                           = cc.DlFunc[func(*ImFontAtlas, *ImFont), cc.Void]{Name: "ImFontAtlas_RemoveFont"}
	imFontAtlas_Clear                                = cc.DlFunc[func(*ImFontAtlas), cc.Void]{Name: "ImFontAtlas_Clear"}
	imFontAtlas_CompactCache                         = cc.DlFunc[func(*ImFontAtlas), cc.Void]{Name: "ImFontAtlas_CompactCache"}
	imFontAtlas_SetFontLoader                        = cc.DlFunc[func(*ImFontAtlas, *ImFontLoader), cc.Void]{Name: "ImFontAtlas_SetFontLoader"}
	imFontAtlas_ClearInputData                       = cc.DlFunc[func(*ImFontAtlas), cc.Void]{Name: "ImFontAtlas_ClearInputData"}
	imFontAtlas_ClearFonts                           = cc.DlFunc[func(*ImFontAtlas), cc.Void]{Name: "ImFontAtlas_ClearFonts"}
	imFontAtlas_ClearTexData                         = cc.DlFunc[func(*ImFontAtlas), cc.Void]{Name: "ImFontAtlas_ClearTexData"}

	imFontAtlas_Build                                 = cc.DlFunc[func(*ImFontAtlas) cc.Bool, cc.Bool]{Name: "ImFontAtlas_Build"}
	imFontAtlas_GetTexDataAsAlpha8                    = cc.DlFunc[func(*ImFontAtlas, **uint8, *int32, *int32, *int32), cc.Void]{Name: "ImFontAtlas_GetTexDataAsAlpha8"}
	imFontAtlas_GetTexDataAsRGBA32                    = cc.DlFunc[func(*ImFontAtlas, **uint8, *int32, *int32, *int32), cc.Void]{Name: "ImFontAtlas_GetTexDataAsRGBA32"}
	imFontAtlas_SetTexID                              = cc.DlFunc[func(*ImFontAtlas, ImTextureID), cc.Void]{Name: "ImFontAtlas_SetTexID"}
	imFontAtlas_SetTexIDImTextureRef                  = cc.DlFunc[func(*ImFontAtlas, ImTextureRef), cc.Void]{Name: "ImFontAtlas_SetTexIDImTextureRef"}
	imFontAtlas_IsBuilt                               = cc.DlFunc[func(*ImFontAtlas) cc.Bool, cc.Bool]{Name: "ImFontAtlas_IsBuilt"}
	imFontAtlas_GetGlyphRangesDefault                 = cc.DlFunc[func(*ImFontAtlas) *ImWchar, *ImWchar]{Name: "ImFontAtlas_GetGlyphRangesDefault"}
	imFontAtlas_GetGlyphRangesGreek                   = cc.DlFunc[func(*ImFontAtlas) *ImWchar, *ImWchar]{Name: "ImFontAtlas_GetGlyphRangesGreek"}
	imFontAtlas_GetGlyphRangesKorean                  = cc.DlFunc[func(*ImFontAtlas) *ImWchar, *ImWchar]{Name: "ImFontAtlas_GetGlyphRangesKorean"}
	imFontAtlas_GetGlyphRangesJapanese                = cc.DlFunc[func(*ImFontAtlas) *ImWchar, *ImWchar]{Name: "ImFontAtlas_GetGlyphRangesJapanese"}
	imFontAtlas_GetGlyphRangesChineseFull             = cc.DlFunc[func(*ImFontAtlas) *ImWchar, *ImWchar]{Name: "ImFontAtlas_GetGlyphRangesChineseFull"}
	imFontAtlas_GetGlyphRangesChineseSimplifiedCommon = cc.DlFunc[func(*ImFontAtlas) *ImWchar, *ImWchar]{Name: "ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon"}
	imFontAtlas_GetGlyphRangesCyrillic                = cc.DlFunc[func(*ImFontAtlas) *ImWchar, *ImWchar]{Name: "ImFontAtlas_GetGlyphRangesCyrillic"}
	imFontAtlas_GetGlyphRangesThai                    = cc.DlFunc[func(*ImFontAtlas) *ImWchar, *ImWchar]{Name: "ImFontAtlas_GetGlyphRangesThai"}
	imFontAtlas_GetGlyphRangesVietnamese              = cc.DlFunc[func(*ImFontAtlas) *ImWchar, *ImWchar]{Name: "ImFontAtlas_GetGlyphRangesVietnamese"}

	imFontAtlas_AddCustomRect                 = cc.DlFunc[func(*ImFontAtlas, int32, int32, *ImFontAtlasRect) ImFontAtlasRectId, ImFontAtlasRectId]{Name: "ImFontAtlas_AddCustomRect"}
	imFontAtlas_RemoveCustomRect              = cc.DlFunc[func(*ImFontAtlas, ImFontAtlasRectId), cc.Void]{Name: "ImFontAtlas_RemoveCustomRect"}
	imFontAtlas_GetCustomRect                 = cc.DlFunc[func(*ImFontAtlas, ImFontAtlasRectId, *ImFontAtlasRect) cc.Bool, cc.Bool]{Name: "ImFontAtlas_GetCustomRect"}
	imFontAtlas_AddCustomRectRegular          = cc.DlFunc[func(*ImFontAtlas, int32, int32) ImFontAtlasRectId, ImFontAtlasRectId]{Name: "ImFontAtlas_AddCustomRectRegular"}
	imFontAtlas_GetCustomRectByIndex          = cc.DlFunc[func(*ImFontAtlas, ImFontAtlasRectId) *ImFontAtlasRect, *ImFontAtlasRect]{Name: "ImFontAtlas_GetCustomRectByIndex"}
	imFontAtlas_CalcCustomRectUV              = cc.DlFunc[func(*ImFontAtlas, *ImFontAtlasRect, *ImVec2, *ImVec2), cc.Void]{Name: "ImFontAtlas_CalcCustomRectUV"}
	imFontAtlas_AddCustomRectFontGlyph        = cc.DlFunc[func(*ImFontAtlas, *ImFont, ImWchar, int32, int32, float32, ImVec2) ImFontAtlasRectId, ImFontAtlasRectId]{Name: "ImFontAtlas_AddCustomRectFontGlyph"}
	imFontAtlas_AddCustomRectFontGlyphForSize = cc.DlFunc[func(*ImFontAtlas, *ImFont, float32, ImWchar, int32, int32, float32, ImVec2) ImFontAtlasRectId, ImFontAtlasRectId]{Name: "ImFontAtlas_AddCustomRectFontGlyphForSize"}

	imFontBaked_ClearOutputData     = cc.DlFunc[func(*ImFontBaked), cc.Void]{Name: "ImFontBaked_ClearOutputData"}
	imFontBaked_FindGlyph           = cc.DlFunc[func(*ImFontBaked, ImWchar) *ImFontGlyph, *ImFontGlyph]{Name: "ImFontBaked_FindGlyph"}
	imFontBaked_FindGlyphNoFallback = cc.DlFunc[func(*ImFontBaked, ImWchar) *ImFontGlyph, *ImFontGlyph]{Name: "ImFontBaked_FindGlyphNoFallback"}
	imFontBaked_GetCharAdvance      = cc.DlFunc[func(*ImFontBaked, ImWchar) float32, float32]{Name: "ImFontBaked_GetCharAdvance"}
	imFontBaked_IsGlyphLoaded       = cc.DlFunc[func(*ImFontBaked, ImWchar) cc.Bool, cc.Bool]{Name: "ImFontBaked_IsGlyphLoaded"}

	imFont_IsGlyphInFont         = cc.DlFunc[func(*ImFont, ImWchar) cc.Bool, cc.Bool]{Name: "ImFont_IsGlyphInFont"}
	imFont_IsLoaded              = cc.DlFunc[func(*ImFont) cc.Bool, cc.Bool]{Name: "ImFont_IsLoaded"}
	imFont_GetDebugName          = cc.DlFunc[func(*ImFont) cc.String, cc.String]{Name: "ImFont_GetDebugName"}
	imFont_GetFontBaked          = cc.DlFunc[func(*ImFont, float32, float32) *ImFontBaked, *ImFontBaked]{Name: "ImFont_GetFontBaked"}
	imFont_CalcTextSizeA         = cc.DlFunc[func(*ImFont, float32, float32, float32, cc.String, cc.String, *cc.String) ImVec2, ImVec2]{Name: "ImFont_CalcTextSizeA"}
	imFont_CalcWordWrapPosition  = cc.DlFunc[func(*ImFont, float32, cc.String, cc.String, float32) cc.String, cc.String]{Name: "ImFont_CalcWordWrapPosition"}
	imFont_RenderChar            = cc.DlFunc[func(*ImFont, *ImDrawList, float32, ImVec2, ImU32, ImWchar, *ImVec4), cc.Void]{Name: "ImFont_RenderChar"}
	imFont_RenderText            = cc.DlFunc[func(*ImFont, *ImDrawList, float32, ImVec2, ImU32, ImVec4, cc.String, cc.String, float32, cc.Bool), cc.Void]{Name: "ImFont_RenderText"}
	imFont_CalcWordWrapPositionA = cc.DlFunc[func(*ImFont, float32, cc.String, cc.String, float32) cc.String, cc.String]{Name: "ImFont_CalcWordWrapPositionA"}
	imFont_ClearOutputData       = cc.DlFunc[func(*ImFont), cc.Void]{Name: "ImFont_ClearOutputData"}
	imFont_AddRemapChar          = cc.DlFunc[func(*ImFont, ImWchar, ImWchar), cc.Void]{Name: "ImFont_AddRemapChar"}
	imFont_IsGlyphRangeUnused    = cc.DlFunc[func(*ImFont, uint32, uint32) cc.Bool, cc.Bool]{Name: "ImFont_IsGlyphRangeUnused"}

	imGuiViewport_GetCenter     = cc.DlFunc[func(*ImGuiViewport) ImVec2, ImVec2]{Name: "ImGuiViewport_GetCenter"}
	imGuiViewport_GetWorkCenter = cc.DlFunc[func(*ImGuiViewport) ImVec2, ImVec2]{Name: "ImGuiViewport_GetWorkCenter"}
	imGui_PushFont              = cc.DlFunc[func(*ImFont), cc.Void]{Name: "ImGui_PushFont"}
	imGui_SetWindowFontScale    = cc.DlFunc[func(float32), cc.Void]{Name: "ImGui_SetWindowFontScale"}
)
