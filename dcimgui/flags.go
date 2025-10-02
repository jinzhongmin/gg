package dcimgui

type ImGuiWindowFlags_ int32

const (
	ImGuiWindowFlags_None                      ImGuiWindowFlags_ = 0
	ImGuiWindowFlags_NoTitleBar                ImGuiWindowFlags_ = 1 << 0
	ImGuiWindowFlags_NoResize                  ImGuiWindowFlags_ = 1 << 1
	ImGuiWindowFlags_NoMove                    ImGuiWindowFlags_ = 1 << 2
	ImGuiWindowFlags_NoScrollbar               ImGuiWindowFlags_ = 1 << 3
	ImGuiWindowFlags_NoScrollWithMouse         ImGuiWindowFlags_ = 1 << 4
	ImGuiWindowFlags_NoCollapse                ImGuiWindowFlags_ = 1 << 5
	ImGuiWindowFlags_AlwaysAutoResize          ImGuiWindowFlags_ = 1 << 6
	ImGuiWindowFlags_NoBackground              ImGuiWindowFlags_ = 1 << 7
	ImGuiWindowFlags_NoSavedSettings           ImGuiWindowFlags_ = 1 << 8
	ImGuiWindowFlags_NoMouseInputs             ImGuiWindowFlags_ = 1 << 9
	ImGuiWindowFlags_MenuBar                   ImGuiWindowFlags_ = 1 << 10
	ImGuiWindowFlags_HorizontalScrollbar       ImGuiWindowFlags_ = 1 << 11
	ImGuiWindowFlags_NoFocusOnAppearing        ImGuiWindowFlags_ = 1 << 12
	ImGuiWindowFlags_NoBringToFrontOnFocus     ImGuiWindowFlags_ = 1 << 13
	ImGuiWindowFlags_AlwaysVerticalScrollbar   ImGuiWindowFlags_ = 1 << 14
	ImGuiWindowFlags_AlwaysHorizontalScrollbar ImGuiWindowFlags_ = 1 << 15
	ImGuiWindowFlags_NoNavInputs               ImGuiWindowFlags_ = 1 << 16
	ImGuiWindowFlags_NoNavFocus                ImGuiWindowFlags_ = 1 << 17
	ImGuiWindowFlags_UnsavedDocument           ImGuiWindowFlags_ = 1 << 18
	ImGuiWindowFlags_NoDocking                 ImGuiWindowFlags_ = 1 << 19
	ImGuiWindowFlags_NoNav                     ImGuiWindowFlags_ = ImGuiWindowFlags_NoNavInputs | ImGuiWindowFlags_NoNavFocus
	ImGuiWindowFlags_NoDecoration              ImGuiWindowFlags_ = ImGuiWindowFlags_NoTitleBar | ImGuiWindowFlags_NoResize | ImGuiWindowFlags_NoScrollbar | ImGuiWindowFlags_NoCollapse
	ImGuiWindowFlags_NoInputs                  ImGuiWindowFlags_ = ImGuiWindowFlags_NoMouseInputs | ImGuiWindowFlags_NoNavInputs | ImGuiWindowFlags_NoNavFocus
	ImGuiWindowFlags_DockNodeHost              ImGuiWindowFlags_ = 1 << 23
	ImGuiWindowFlags_ChildWindow               ImGuiWindowFlags_ = 1 << 24
	ImGuiWindowFlags_Tooltip                   ImGuiWindowFlags_ = 1 << 25
	ImGuiWindowFlags_Popup                     ImGuiWindowFlags_ = 1 << 26
	ImGuiWindowFlags_Modal                     ImGuiWindowFlags_ = 1 << 27
	ImGuiWindowFlags_ChildMenu                 ImGuiWindowFlags_ = 1 << 28
	ImGuiWindowFlags_NavFlattened              ImGuiWindowFlags_ = 1 << 29
	ImGuiWindowFlags_AlwaysUseWindowPadding    ImGuiWindowFlags_ = 1 << 30
)

type ImGuiChildFlags_ int32

const (
	ImGuiChildFlags_None                   ImGuiChildFlags_ = 0
	ImGuiChildFlags_Borders                ImGuiChildFlags_ = 1 << 0
	ImGuiChildFlags_AlwaysUseWindowPadding ImGuiChildFlags_ = 1 << 1
	ImGuiChildFlags_ResizeX                ImGuiChildFlags_ = 1 << 2
	ImGuiChildFlags_ResizeY                ImGuiChildFlags_ = 1 << 3
	ImGuiChildFlags_AutoResizeX            ImGuiChildFlags_ = 1 << 4
	ImGuiChildFlags_AutoResizeY            ImGuiChildFlags_ = 1 << 5
	ImGuiChildFlags_AlwaysAutoResize       ImGuiChildFlags_ = 1 << 6
	ImGuiChildFlags_FrameStyle             ImGuiChildFlags_ = 1 << 7
	ImGuiChildFlags_NavFlattened           ImGuiChildFlags_ = 1 << 8
	ImGuiChildFlags_Border                 ImGuiChildFlags_ = ImGuiChildFlags_Borders
)

type ImGuiItemFlags_ int32

const (
	ImGuiItemFlags_None              ImGuiItemFlags_ = 0
	ImGuiItemFlags_NoTabStop         ImGuiItemFlags_ = 1 << 0
	ImGuiItemFlags_NoNav             ImGuiItemFlags_ = 1 << 1
	ImGuiItemFlags_NoNavDefaultFocus ImGuiItemFlags_ = 1 << 2
	ImGuiItemFlags_ButtonRepeat      ImGuiItemFlags_ = 1 << 3
	ImGuiItemFlags_AutoClosePopups   ImGuiItemFlags_ = 1 << 4
	ImGuiItemFlags_AllowDuplicateId  ImGuiItemFlags_ = 1 << 5
)

type ImGuiInputTextFlags_ int32

const (
	ImGuiInputTextFlags_None                ImGuiInputTextFlags_ = 0
	ImGuiInputTextFlags_CharsDecimal        ImGuiInputTextFlags_ = 1 << 0
	ImGuiInputTextFlags_CharsHexadecimal    ImGuiInputTextFlags_ = 1 << 1
	ImGuiInputTextFlags_CharsScientific     ImGuiInputTextFlags_ = 1 << 2
	ImGuiInputTextFlags_CharsUppercase      ImGuiInputTextFlags_ = 1 << 3
	ImGuiInputTextFlags_CharsNoBlank        ImGuiInputTextFlags_ = 1 << 4
	ImGuiInputTextFlags_AllowTabInput       ImGuiInputTextFlags_ = 1 << 5
	ImGuiInputTextFlags_EnterReturnsTrue    ImGuiInputTextFlags_ = 1 << 6
	ImGuiInputTextFlags_EscapeClearsAll     ImGuiInputTextFlags_ = 1 << 7
	ImGuiInputTextFlags_CtrlEnterForNewLine ImGuiInputTextFlags_ = 1 << 8
	ImGuiInputTextFlags_ReadOnly            ImGuiInputTextFlags_ = 1 << 9
	ImGuiInputTextFlags_Password            ImGuiInputTextFlags_ = 1 << 10
	ImGuiInputTextFlags_AlwaysOverwrite     ImGuiInputTextFlags_ = 1 << 11
	ImGuiInputTextFlags_AutoSelectAll       ImGuiInputTextFlags_ = 1 << 12
	ImGuiInputTextFlags_ParseEmptyRefVal    ImGuiInputTextFlags_ = 1 << 13
	ImGuiInputTextFlags_DisplayEmptyRefVal  ImGuiInputTextFlags_ = 1 << 14
	ImGuiInputTextFlags_NoHorizontalScroll  ImGuiInputTextFlags_ = 1 << 15
	ImGuiInputTextFlags_NoUndoRedo          ImGuiInputTextFlags_ = 1 << 16
	ImGuiInputTextFlags_ElideLeft           ImGuiInputTextFlags_ = 1 << 17
	ImGuiInputTextFlags_CallbackCompletion  ImGuiInputTextFlags_ = 1 << 18
	ImGuiInputTextFlags_CallbackHistory     ImGuiInputTextFlags_ = 1 << 19
	ImGuiInputTextFlags_CallbackAlways      ImGuiInputTextFlags_ = 1 << 20
	ImGuiInputTextFlags_CallbackCharFilter  ImGuiInputTextFlags_ = 1 << 21
	ImGuiInputTextFlags_CallbackResize      ImGuiInputTextFlags_ = 1 << 22
	ImGuiInputTextFlags_CallbackEdit        ImGuiInputTextFlags_ = 1 << 23
)

type ImGuiTreeNodeFlags_ int32

const (
	ImGuiTreeNodeFlags_None                 ImGuiTreeNodeFlags_ = 0
	ImGuiTreeNodeFlags_Selected             ImGuiTreeNodeFlags_ = 1 << 0
	ImGuiTreeNodeFlags_Framed               ImGuiTreeNodeFlags_ = 1 << 1
	ImGuiTreeNodeFlags_AllowOverlap         ImGuiTreeNodeFlags_ = 1 << 2
	ImGuiTreeNodeFlags_NoTreePushOnOpen     ImGuiTreeNodeFlags_ = 1 << 3
	ImGuiTreeNodeFlags_NoAutoOpenOnLog      ImGuiTreeNodeFlags_ = 1 << 4
	ImGuiTreeNodeFlags_DefaultOpen          ImGuiTreeNodeFlags_ = 1 << 5
	ImGuiTreeNodeFlags_OpenOnDoubleClick    ImGuiTreeNodeFlags_ = 1 << 6
	ImGuiTreeNodeFlags_OpenOnArrow          ImGuiTreeNodeFlags_ = 1 << 7
	ImGuiTreeNodeFlags_Leaf                 ImGuiTreeNodeFlags_ = 1 << 8
	ImGuiTreeNodeFlags_Bullet               ImGuiTreeNodeFlags_ = 1 << 9
	ImGuiTreeNodeFlags_FramePadding         ImGuiTreeNodeFlags_ = 1 << 10
	ImGuiTreeNodeFlags_SpanAvailWidth       ImGuiTreeNodeFlags_ = 1 << 11
	ImGuiTreeNodeFlags_SpanFullWidth        ImGuiTreeNodeFlags_ = 1 << 12
	ImGuiTreeNodeFlags_SpanLabelWidth       ImGuiTreeNodeFlags_ = 1 << 13
	ImGuiTreeNodeFlags_SpanAllColumns       ImGuiTreeNodeFlags_ = 1 << 14
	ImGuiTreeNodeFlags_LabelSpanAllColumns  ImGuiTreeNodeFlags_ = 1 << 15
	ImGuiTreeNodeFlags_NavLeftJumpsToParent ImGuiTreeNodeFlags_ = 1 << 17
	ImGuiTreeNodeFlags_CollapsingHeader     ImGuiTreeNodeFlags_ = ImGuiTreeNodeFlags_Framed | ImGuiTreeNodeFlags_NoTreePushOnOpen | ImGuiTreeNodeFlags_NoAutoOpenOnLog
	ImGuiTreeNodeFlags_DrawLinesNone        ImGuiTreeNodeFlags_ = 1 << 18
	ImGuiTreeNodeFlags_DrawLinesFull        ImGuiTreeNodeFlags_ = 1 << 19
	ImGuiTreeNodeFlags_DrawLinesToNodes     ImGuiTreeNodeFlags_ = 1 << 20
	ImGuiTreeNodeFlags_NavLeftJumpsBackHere ImGuiTreeNodeFlags_ = ImGuiTreeNodeFlags_NavLeftJumpsToParent
	ImGuiTreeNodeFlags_SpanTextWidth        ImGuiTreeNodeFlags_ = ImGuiTreeNodeFlags_SpanLabelWidth
	ImGuiTreeNodeFlags_AllowItemOverlap     ImGuiTreeNodeFlags_ = ImGuiTreeNodeFlags_AllowOverlap
)

type ImGuiPopupFlags_ int32

const (
	ImGuiPopupFlags_None                    ImGuiPopupFlags_ = 0
	ImGuiPopupFlags_MouseButtonLeft         ImGuiPopupFlags_ = 0
	ImGuiPopupFlags_MouseButtonRight        ImGuiPopupFlags_ = 1
	ImGuiPopupFlags_MouseButtonMiddle       ImGuiPopupFlags_ = 2
	ImGuiPopupFlags_MouseButtonMask_        ImGuiPopupFlags_ = 0x1F
	ImGuiPopupFlags_MouseButtonDefault_     ImGuiPopupFlags_ = 1
	ImGuiPopupFlags_NoReopen                ImGuiPopupFlags_ = 1 << 5
	ImGuiPopupFlags_NoOpenOverExistingPopup ImGuiPopupFlags_ = 1 << 7
	ImGuiPopupFlags_NoOpenOverItems         ImGuiPopupFlags_ = 1 << 8
	ImGuiPopupFlags_AnyPopupId              ImGuiPopupFlags_ = 1 << 10
	ImGuiPopupFlags_AnyPopupLevel           ImGuiPopupFlags_ = 1 << 11
	ImGuiPopupFlags_AnyPopup                ImGuiPopupFlags_ = ImGuiPopupFlags_AnyPopupId | ImGuiPopupFlags_AnyPopupLevel
)

type ImGuiSelectableFlags_ int32

const (
	ImGuiSelectableFlags_None              ImGuiSelectableFlags_ = 0
	ImGuiSelectableFlags_NoAutoClosePopups ImGuiSelectableFlags_ = 1 << 0
	ImGuiSelectableFlags_SpanAllColumns    ImGuiSelectableFlags_ = 1 << 1
	ImGuiSelectableFlags_AllowDoubleClick  ImGuiSelectableFlags_ = 1 << 2
	ImGuiSelectableFlags_Disabled          ImGuiSelectableFlags_ = 1 << 3
	ImGuiSelectableFlags_AllowOverlap      ImGuiSelectableFlags_ = 1 << 4
	ImGuiSelectableFlags_Highlight         ImGuiSelectableFlags_ = 1 << 5
	ImGuiSelectableFlags_DontClosePopups   ImGuiSelectableFlags_ = ImGuiSelectableFlags_NoAutoClosePopups
	ImGuiSelectableFlags_AllowItemOverlap  ImGuiSelectableFlags_ = ImGuiSelectableFlags_AllowOverlap
)

type ImGuiComboFlags_ int32

const (
	ImGuiComboFlags_None            ImGuiComboFlags_ = 0
	ImGuiComboFlags_PopupAlignLeft  ImGuiComboFlags_ = 1 << 0
	ImGuiComboFlags_HeightSmall     ImGuiComboFlags_ = 1 << 1
	ImGuiComboFlags_HeightRegular   ImGuiComboFlags_ = 1 << 2
	ImGuiComboFlags_HeightLarge     ImGuiComboFlags_ = 1 << 3
	ImGuiComboFlags_HeightLargest   ImGuiComboFlags_ = 1 << 4
	ImGuiComboFlags_NoArrowButton   ImGuiComboFlags_ = 1 << 5
	ImGuiComboFlags_NoPreview       ImGuiComboFlags_ = 1 << 6
	ImGuiComboFlags_WidthFitPreview ImGuiComboFlags_ = 1 << 7
	ImGuiComboFlags_HeightMask_     ImGuiComboFlags_ = ImGuiComboFlags_HeightSmall | ImGuiComboFlags_HeightRegular | ImGuiComboFlags_HeightLarge | ImGuiComboFlags_HeightLargest
)

type ImGuiTabBarFlags_ int32

const (
	ImGuiTabBarFlags_None                         ImGuiTabBarFlags_ = 0
	ImGuiTabBarFlags_Reorderable                  ImGuiTabBarFlags_ = 1 << 0
	ImGuiTabBarFlags_AutoSelectNewTabs            ImGuiTabBarFlags_ = 1 << 1
	ImGuiTabBarFlags_TabListPopupButton           ImGuiTabBarFlags_ = 1 << 2
	ImGuiTabBarFlags_NoCloseWithMiddleMouseButton ImGuiTabBarFlags_ = 1 << 3
	ImGuiTabBarFlags_NoTabListScrollingButtons    ImGuiTabBarFlags_ = 1 << 4
	ImGuiTabBarFlags_NoTooltip                    ImGuiTabBarFlags_ = 1 << 5
	ImGuiTabBarFlags_DrawSelectedOverline         ImGuiTabBarFlags_ = 1 << 6
	ImGuiTabBarFlags_FittingPolicyResizeDown      ImGuiTabBarFlags_ = 1 << 7
	ImGuiTabBarFlags_FittingPolicyScroll          ImGuiTabBarFlags_ = 1 << 8
	ImGuiTabBarFlags_FittingPolicyMask_           ImGuiTabBarFlags_ = ImGuiTabBarFlags_FittingPolicyResizeDown | ImGuiTabBarFlags_FittingPolicyScroll
	ImGuiTabBarFlags_FittingPolicyDefault_        ImGuiTabBarFlags_ = ImGuiTabBarFlags_FittingPolicyResizeDown
)

type ImGuiTabItemFlags_ int32

const (
	ImGuiTabItemFlags_None                         ImGuiTabItemFlags_ = 0
	ImGuiTabItemFlags_UnsavedDocument              ImGuiTabItemFlags_ = 1 << 0
	ImGuiTabItemFlags_SetSelected                  ImGuiTabItemFlags_ = 1 << 1
	ImGuiTabItemFlags_NoCloseWithMiddleMouseButton ImGuiTabItemFlags_ = 1 << 2
	ImGuiTabItemFlags_NoPushId                     ImGuiTabItemFlags_ = 1 << 3
	ImGuiTabItemFlags_NoTooltip                    ImGuiTabItemFlags_ = 1 << 4
	ImGuiTabItemFlags_NoReorder                    ImGuiTabItemFlags_ = 1 << 5
	ImGuiTabItemFlags_Leading                      ImGuiTabItemFlags_ = 1 << 6
	ImGuiTabItemFlags_Trailing                     ImGuiTabItemFlags_ = 1 << 7
	ImGuiTabItemFlags_NoAssumedClosure             ImGuiTabItemFlags_ = 1 << 8
)

type ImGuiFocusedFlags_ int32

const (
	ImGuiFocusedFlags_None                ImGuiFocusedFlags_ = 0
	ImGuiFocusedFlags_ChildWindows        ImGuiFocusedFlags_ = 1 << 0
	ImGuiFocusedFlags_RootWindow          ImGuiFocusedFlags_ = 1 << 1
	ImGuiFocusedFlags_AnyWindow           ImGuiFocusedFlags_ = 1 << 2
	ImGuiFocusedFlags_NoPopupHierarchy    ImGuiFocusedFlags_ = 1 << 3
	ImGuiFocusedFlags_DockHierarchy       ImGuiFocusedFlags_ = 1 << 4
	ImGuiFocusedFlags_RootAndChildWindows ImGuiFocusedFlags_ = ImGuiFocusedFlags_RootWindow | ImGuiFocusedFlags_ChildWindows
)

type ImGuiHoveredFlags_ int32

const (
	ImGuiHoveredFlags_None                         ImGuiHoveredFlags_ = 0
	ImGuiHoveredFlags_ChildWindows                 ImGuiHoveredFlags_ = 1 << 0
	ImGuiHoveredFlags_RootWindow                   ImGuiHoveredFlags_ = 1 << 1
	ImGuiHoveredFlags_AnyWindow                    ImGuiHoveredFlags_ = 1 << 2
	ImGuiHoveredFlags_NoPopupHierarchy             ImGuiHoveredFlags_ = 1 << 3
	ImGuiHoveredFlags_DockHierarchy                ImGuiHoveredFlags_ = 1 << 4
	ImGuiHoveredFlags_AllowWhenBlockedByPopup      ImGuiHoveredFlags_ = 1 << 5
	ImGuiHoveredFlags_AllowWhenBlockedByActiveItem ImGuiHoveredFlags_ = 1 << 7
	ImGuiHoveredFlags_AllowWhenOverlappedByItem    ImGuiHoveredFlags_ = 1 << 8
	ImGuiHoveredFlags_AllowWhenOverlappedByWindow  ImGuiHoveredFlags_ = 1 << 9
	ImGuiHoveredFlags_AllowWhenDisabled            ImGuiHoveredFlags_ = 1 << 10
	ImGuiHoveredFlags_NoNavOverride                ImGuiHoveredFlags_ = 1 << 11
	ImGuiHoveredFlags_AllowWhenOverlapped          ImGuiHoveredFlags_ = ImGuiHoveredFlags_AllowWhenOverlappedByItem | ImGuiHoveredFlags_AllowWhenOverlappedByWindow
	ImGuiHoveredFlags_RectOnly                     ImGuiHoveredFlags_ = ImGuiHoveredFlags_AllowWhenBlockedByPopup | ImGuiHoveredFlags_AllowWhenBlockedByActiveItem | ImGuiHoveredFlags_AllowWhenOverlapped
	ImGuiHoveredFlags_RootAndChildWindows          ImGuiHoveredFlags_ = ImGuiHoveredFlags_RootWindow | ImGuiHoveredFlags_ChildWindows
	ImGuiHoveredFlags_ForTooltip                   ImGuiHoveredFlags_ = 1 << 12
	ImGuiHoveredFlags_Stationary                   ImGuiHoveredFlags_ = 1 << 13
	ImGuiHoveredFlags_DelayNone                    ImGuiHoveredFlags_ = 1 << 14
	ImGuiHoveredFlags_DelayShort                   ImGuiHoveredFlags_ = 1 << 15
	ImGuiHoveredFlags_DelayNormal                  ImGuiHoveredFlags_ = 1 << 16
	ImGuiHoveredFlags_NoSharedDelay                ImGuiHoveredFlags_ = 1 << 17
)

type ImGuiDockNodeFlags_ int32

const (
	ImGuiDockNodeFlags_None                     ImGuiDockNodeFlags_ = 0
	ImGuiDockNodeFlags_KeepAliveOnly            ImGuiDockNodeFlags_ = 1 << 0
	ImGuiDockNodeFlags_NoDockingOverCentralNode ImGuiDockNodeFlags_ = 1 << 2
	ImGuiDockNodeFlags_PassthruCentralNode      ImGuiDockNodeFlags_ = 1 << 3
	ImGuiDockNodeFlags_NoDockingSplit           ImGuiDockNodeFlags_ = 1 << 4
	ImGuiDockNodeFlags_NoResize                 ImGuiDockNodeFlags_ = 1 << 5
	ImGuiDockNodeFlags_AutoHideTabBar           ImGuiDockNodeFlags_ = 1 << 6
	ImGuiDockNodeFlags_NoUndocking              ImGuiDockNodeFlags_ = 1 << 7
	ImGuiDockNodeFlags_NoSplit                  ImGuiDockNodeFlags_ = ImGuiDockNodeFlags_NoDockingSplit
	ImGuiDockNodeFlags_NoDockingInCentralNode   ImGuiDockNodeFlags_ = ImGuiDockNodeFlags_NoDockingOverCentralNode
)

type ImGuiDragDropFlags_ int32

const (
	ImGuiDragDropFlags_None                     ImGuiDragDropFlags_ = 0
	ImGuiDragDropFlags_SourceNoPreviewTooltip   ImGuiDragDropFlags_ = 1 << 0
	ImGuiDragDropFlags_SourceNoDisableHover     ImGuiDragDropFlags_ = 1 << 1
	ImGuiDragDropFlags_SourceNoHoldToOpenOthers ImGuiDragDropFlags_ = 1 << 2
	ImGuiDragDropFlags_SourceAllowNullID        ImGuiDragDropFlags_ = 1 << 3
	ImGuiDragDropFlags_SourceExtern             ImGuiDragDropFlags_ = 1 << 4
	ImGuiDragDropFlags_PayloadAutoExpire        ImGuiDragDropFlags_ = 1 << 5
	ImGuiDragDropFlags_PayloadNoCrossContext    ImGuiDragDropFlags_ = 1 << 6
	ImGuiDragDropFlags_PayloadNoCrossProcess    ImGuiDragDropFlags_ = 1 << 7
	ImGuiDragDropFlags_AcceptBeforeDelivery     ImGuiDragDropFlags_ = 1 << 10
	ImGuiDragDropFlags_AcceptNoDrawDefaultRect  ImGuiDragDropFlags_ = 1 << 11
	ImGuiDragDropFlags_AcceptNoPreviewTooltip   ImGuiDragDropFlags_ = 1 << 12
	ImGuiDragDropFlags_AcceptPeekOnly           ImGuiDragDropFlags_ = ImGuiDragDropFlags_AcceptBeforeDelivery | ImGuiDragDropFlags_AcceptNoDrawDefaultRect
	ImGuiDragDropFlags_SourceAutoExpirePayload  ImGuiDragDropFlags_ = ImGuiDragDropFlags_PayloadAutoExpire
)

type ImGuiDataType_ int32

const (
	ImGuiDataType_S8     ImGuiDataType_ = 0
	ImGuiDataType_U8     ImGuiDataType_ = 1
	ImGuiDataType_S16    ImGuiDataType_ = 2
	ImGuiDataType_U16    ImGuiDataType_ = 3
	ImGuiDataType_S32    ImGuiDataType_ = 4
	ImGuiDataType_U32    ImGuiDataType_ = 5
	ImGuiDataType_S64    ImGuiDataType_ = 6
	ImGuiDataType_U64    ImGuiDataType_ = 7
	ImGuiDataType_Float  ImGuiDataType_ = 8
	ImGuiDataType_Double ImGuiDataType_ = 9
	ImGuiDataType_Bool   ImGuiDataType_ = 10
	ImGuiDataType_String ImGuiDataType_ = 11
	ImGuiDataType_COUNT  ImGuiDataType_ = 12
)

type ImGuiDir int32

const (
	ImGuiDir_None  ImGuiDir = -1
	ImGuiDir_Left  ImGuiDir = 0
	ImGuiDir_Right ImGuiDir = 1
	ImGuiDir_Up    ImGuiDir = 2
	ImGuiDir_Down  ImGuiDir = 3
	ImGuiDir_COUNT ImGuiDir = 4
)

type ImGuiSortDirection int32

const (
	ImGuiSortDirection_None       ImGuiSortDirection = 0
	ImGuiSortDirection_Ascending  ImGuiSortDirection = 1
	ImGuiSortDirection_Descending ImGuiSortDirection = 2
)

type ImGuiKey int32

const (
	ImGuiKey_None                ImGuiKey = 0
	ImGuiKey_NamedKey_BEGIN      ImGuiKey = 512
	ImGuiKey_Tab                 ImGuiKey = 512
	ImGuiKey_LeftArrow           ImGuiKey = 513
	ImGuiKey_RightArrow          ImGuiKey = 514
	ImGuiKey_UpArrow             ImGuiKey = 515
	ImGuiKey_DownArrow           ImGuiKey = 516
	ImGuiKey_PageUp              ImGuiKey = 517
	ImGuiKey_PageDown            ImGuiKey = 518
	ImGuiKey_Home                ImGuiKey = 519
	ImGuiKey_End                 ImGuiKey = 520
	ImGuiKey_Insert              ImGuiKey = 521
	ImGuiKey_Delete              ImGuiKey = 522
	ImGuiKey_Backspace           ImGuiKey = 523
	ImGuiKey_Space               ImGuiKey = 524
	ImGuiKey_Enter               ImGuiKey = 525
	ImGuiKey_Escape              ImGuiKey = 526
	ImGuiKey_LeftCtrl            ImGuiKey = 527
	ImGuiKey_LeftShift           ImGuiKey = 528
	ImGuiKey_LeftAlt             ImGuiKey = 529
	ImGuiKey_LeftSuper           ImGuiKey = 530
	ImGuiKey_RightCtrl           ImGuiKey = 531
	ImGuiKey_RightShift          ImGuiKey = 532
	ImGuiKey_RightAlt            ImGuiKey = 533
	ImGuiKey_RightSuper          ImGuiKey = 534
	ImGuiKey_Menu                ImGuiKey = 535
	ImGuiKey_0                   ImGuiKey = 536
	ImGuiKey_1                   ImGuiKey = 537
	ImGuiKey_2                   ImGuiKey = 538
	ImGuiKey_3                   ImGuiKey = 539
	ImGuiKey_4                   ImGuiKey = 540
	ImGuiKey_5                   ImGuiKey = 541
	ImGuiKey_6                   ImGuiKey = 542
	ImGuiKey_7                   ImGuiKey = 543
	ImGuiKey_8                   ImGuiKey = 544
	ImGuiKey_9                   ImGuiKey = 545
	ImGuiKey_A                   ImGuiKey = 546
	ImGuiKey_B                   ImGuiKey = 547
	ImGuiKey_C                   ImGuiKey = 548
	ImGuiKey_D                   ImGuiKey = 549
	ImGuiKey_E                   ImGuiKey = 550
	ImGuiKey_F                   ImGuiKey = 551
	ImGuiKey_G                   ImGuiKey = 552
	ImGuiKey_H                   ImGuiKey = 553
	ImGuiKey_I                   ImGuiKey = 554
	ImGuiKey_J                   ImGuiKey = 555
	ImGuiKey_K                   ImGuiKey = 556
	ImGuiKey_L                   ImGuiKey = 557
	ImGuiKey_M                   ImGuiKey = 558
	ImGuiKey_N                   ImGuiKey = 559
	ImGuiKey_O                   ImGuiKey = 560
	ImGuiKey_P                   ImGuiKey = 561
	ImGuiKey_Q                   ImGuiKey = 562
	ImGuiKey_R                   ImGuiKey = 563
	ImGuiKey_S                   ImGuiKey = 564
	ImGuiKey_T                   ImGuiKey = 565
	ImGuiKey_U                   ImGuiKey = 566
	ImGuiKey_V                   ImGuiKey = 567
	ImGuiKey_W                   ImGuiKey = 568
	ImGuiKey_X                   ImGuiKey = 569
	ImGuiKey_Y                   ImGuiKey = 570
	ImGuiKey_Z                   ImGuiKey = 571
	ImGuiKey_F1                  ImGuiKey = 572
	ImGuiKey_F2                  ImGuiKey = 573
	ImGuiKey_F3                  ImGuiKey = 574
	ImGuiKey_F4                  ImGuiKey = 575
	ImGuiKey_F5                  ImGuiKey = 576
	ImGuiKey_F6                  ImGuiKey = 577
	ImGuiKey_F7                  ImGuiKey = 578
	ImGuiKey_F8                  ImGuiKey = 579
	ImGuiKey_F9                  ImGuiKey = 580
	ImGuiKey_F10                 ImGuiKey = 581
	ImGuiKey_F11                 ImGuiKey = 582
	ImGuiKey_F12                 ImGuiKey = 583
	ImGuiKey_F13                 ImGuiKey = 584
	ImGuiKey_F14                 ImGuiKey = 585
	ImGuiKey_F15                 ImGuiKey = 586
	ImGuiKey_F16                 ImGuiKey = 587
	ImGuiKey_F17                 ImGuiKey = 588
	ImGuiKey_F18                 ImGuiKey = 589
	ImGuiKey_F19                 ImGuiKey = 590
	ImGuiKey_F20                 ImGuiKey = 591
	ImGuiKey_F21                 ImGuiKey = 592
	ImGuiKey_F22                 ImGuiKey = 593
	ImGuiKey_F23                 ImGuiKey = 594
	ImGuiKey_F24                 ImGuiKey = 595
	ImGuiKey_Apostrophe          ImGuiKey = 596
	ImGuiKey_Comma               ImGuiKey = 597
	ImGuiKey_Minus               ImGuiKey = 598
	ImGuiKey_Period              ImGuiKey = 599
	ImGuiKey_Slash               ImGuiKey = 600
	ImGuiKey_Semicolon           ImGuiKey = 601
	ImGuiKey_Equal               ImGuiKey = 602
	ImGuiKey_LeftBracket         ImGuiKey = 603
	ImGuiKey_Backslash           ImGuiKey = 604
	ImGuiKey_RightBracket        ImGuiKey = 605
	ImGuiKey_GraveAccent         ImGuiKey = 606
	ImGuiKey_CapsLock            ImGuiKey = 607
	ImGuiKey_ScrollLock          ImGuiKey = 608
	ImGuiKey_NumLock             ImGuiKey = 609
	ImGuiKey_PrintScreen         ImGuiKey = 610
	ImGuiKey_Pause               ImGuiKey = 611
	ImGuiKey_Keypad0             ImGuiKey = 612
	ImGuiKey_Keypad1             ImGuiKey = 613
	ImGuiKey_Keypad2             ImGuiKey = 614
	ImGuiKey_Keypad3             ImGuiKey = 615
	ImGuiKey_Keypad4             ImGuiKey = 616
	ImGuiKey_Keypad5             ImGuiKey = 617
	ImGuiKey_Keypad6             ImGuiKey = 618
	ImGuiKey_Keypad7             ImGuiKey = 619
	ImGuiKey_Keypad8             ImGuiKey = 620
	ImGuiKey_Keypad9             ImGuiKey = 621
	ImGuiKey_KeypadDecimal       ImGuiKey = 622
	ImGuiKey_KeypadDivide        ImGuiKey = 623
	ImGuiKey_KeypadMultiply      ImGuiKey = 624
	ImGuiKey_KeypadSubtract      ImGuiKey = 625
	ImGuiKey_KeypadAdd           ImGuiKey = 626
	ImGuiKey_KeypadEnter         ImGuiKey = 627
	ImGuiKey_KeypadEqual         ImGuiKey = 628
	ImGuiKey_AppBack             ImGuiKey = 629
	ImGuiKey_AppForward          ImGuiKey = 630
	ImGuiKey_Oem102              ImGuiKey = 631
	ImGuiKey_GamepadStart        ImGuiKey = 632
	ImGuiKey_GamepadBack         ImGuiKey = 633
	ImGuiKey_GamepadFaceLeft     ImGuiKey = 634
	ImGuiKey_GamepadFaceRight    ImGuiKey = 635
	ImGuiKey_GamepadFaceUp       ImGuiKey = 636
	ImGuiKey_GamepadFaceDown     ImGuiKey = 637
	ImGuiKey_GamepadDpadLeft     ImGuiKey = 638
	ImGuiKey_GamepadDpadRight    ImGuiKey = 639
	ImGuiKey_GamepadDpadUp       ImGuiKey = 640
	ImGuiKey_GamepadDpadDown     ImGuiKey = 641
	ImGuiKey_GamepadL1           ImGuiKey = 642
	ImGuiKey_GamepadR1           ImGuiKey = 643
	ImGuiKey_GamepadL2           ImGuiKey = 644
	ImGuiKey_GamepadR2           ImGuiKey = 645
	ImGuiKey_GamepadL3           ImGuiKey = 646
	ImGuiKey_GamepadR3           ImGuiKey = 647
	ImGuiKey_GamepadLStickLeft   ImGuiKey = 648
	ImGuiKey_GamepadLStickRight  ImGuiKey = 649
	ImGuiKey_GamepadLStickUp     ImGuiKey = 650
	ImGuiKey_GamepadLStickDown   ImGuiKey = 651
	ImGuiKey_GamepadRStickLeft   ImGuiKey = 652
	ImGuiKey_GamepadRStickRight  ImGuiKey = 653
	ImGuiKey_GamepadRStickUp     ImGuiKey = 654
	ImGuiKey_GamepadRStickDown   ImGuiKey = 655
	ImGuiKey_MouseLeft           ImGuiKey = 656
	ImGuiKey_MouseRight          ImGuiKey = 657
	ImGuiKey_MouseMiddle         ImGuiKey = 658
	ImGuiKey_MouseX1             ImGuiKey = 659
	ImGuiKey_MouseX2             ImGuiKey = 660
	ImGuiKey_MouseWheelX         ImGuiKey = 661
	ImGuiKey_MouseWheelY         ImGuiKey = 662
	ImGuiKey_ReservedForModCtrl  ImGuiKey = 663
	ImGuiKey_ReservedForModShift ImGuiKey = 664
	ImGuiKey_ReservedForModAlt   ImGuiKey = 665
	ImGuiKey_ReservedForModSuper ImGuiKey = 666
	ImGuiKey_NamedKey_END        ImGuiKey = 667
	ImGuiKey_NamedKey_COUNT      ImGuiKey = ImGuiKey_NamedKey_END - ImGuiKey_NamedKey_BEGIN
	ImGuiMod_None                ImGuiKey = 0
	ImGuiMod_Ctrl                ImGuiKey = 1 << 12
	ImGuiMod_Shift               ImGuiKey = 1 << 13
	ImGuiMod_Alt                 ImGuiKey = 1 << 14
	ImGuiMod_Super               ImGuiKey = 1 << 15
	ImGuiMod_Mask_               ImGuiKey = 0xF000
	ImGuiKey_COUNT               ImGuiKey = ImGuiKey_NamedKey_END
	ImGuiMod_Shortcut            ImGuiKey = ImGuiMod_Ctrl
	ImGuiKey_ModCtrl             ImGuiKey = ImGuiMod_Ctrl
	ImGuiKey_ModShift            ImGuiKey = ImGuiMod_Shift
	ImGuiKey_ModAlt              ImGuiKey = ImGuiMod_Alt
	ImGuiKey_ModSuper            ImGuiKey = ImGuiMod_Super
)

type ImGuiInputFlags_ int32

const (
	ImGuiInputFlags_None                 ImGuiInputFlags_ = 0
	ImGuiInputFlags_Repeat               ImGuiInputFlags_ = 1 << 0
	ImGuiInputFlags_RouteActive          ImGuiInputFlags_ = 1 << 10
	ImGuiInputFlags_RouteFocused         ImGuiInputFlags_ = 1 << 11
	ImGuiInputFlags_RouteGlobal          ImGuiInputFlags_ = 1 << 12
	ImGuiInputFlags_RouteAlways          ImGuiInputFlags_ = 1 << 13
	ImGuiInputFlags_RouteOverFocused     ImGuiInputFlags_ = 1 << 14
	ImGuiInputFlags_RouteOverActive      ImGuiInputFlags_ = 1 << 15
	ImGuiInputFlags_RouteUnlessBgFocused ImGuiInputFlags_ = 1 << 16
	ImGuiInputFlags_RouteFromRootWindow  ImGuiInputFlags_ = 1 << 17
	ImGuiInputFlags_Tooltip              ImGuiInputFlags_ = 1 << 18
)

type ImGuiConfigFlags_ int32

const (
	ImGuiConfigFlags_None                    ImGuiConfigFlags_ = 0
	ImGuiConfigFlags_NavEnableKeyboard       ImGuiConfigFlags_ = 1 << 0
	ImGuiConfigFlags_NavEnableGamepad        ImGuiConfigFlags_ = 1 << 1
	ImGuiConfigFlags_NoMouse                 ImGuiConfigFlags_ = 1 << 4
	ImGuiConfigFlags_NoMouseCursorChange     ImGuiConfigFlags_ = 1 << 5
	ImGuiConfigFlags_NoKeyboard              ImGuiConfigFlags_ = 1 << 6
	ImGuiConfigFlags_DockingEnable           ImGuiConfigFlags_ = 1 << 7
	ImGuiConfigFlags_ViewportsEnable         ImGuiConfigFlags_ = 1 << 10
	ImGuiConfigFlags_IsSRGB                  ImGuiConfigFlags_ = 1 << 20
	ImGuiConfigFlags_IsTouchScreen           ImGuiConfigFlags_ = 1 << 21
	ImGuiConfigFlags_NavEnableSetMousePos    ImGuiConfigFlags_ = 1 << 2
	ImGuiConfigFlags_NavNoCaptureKeyboard    ImGuiConfigFlags_ = 1 << 3
	ImGuiConfigFlags_DpiEnableScaleFonts     ImGuiConfigFlags_ = 1 << 14
	ImGuiConfigFlags_DpiEnableScaleViewports ImGuiConfigFlags_ = 1 << 15
)

type ImGuiBackendFlags_ int32

const (
	ImGuiBackendFlags_None                    ImGuiBackendFlags_ = 0
	ImGuiBackendFlags_HasGamepad              ImGuiBackendFlags_ = 1 << 0
	ImGuiBackendFlags_HasMouseCursors         ImGuiBackendFlags_ = 1 << 1
	ImGuiBackendFlags_HasSetMousePos          ImGuiBackendFlags_ = 1 << 2
	ImGuiBackendFlags_RendererHasVtxOffset    ImGuiBackendFlags_ = 1 << 3
	ImGuiBackendFlags_RendererHasTextures     ImGuiBackendFlags_ = 1 << 4
	ImGuiBackendFlags_PlatformHasViewports    ImGuiBackendFlags_ = 1 << 10
	ImGuiBackendFlags_HasMouseHoveredViewport ImGuiBackendFlags_ = 1 << 11
	ImGuiBackendFlags_RendererHasViewports    ImGuiBackendFlags_ = 1 << 12
)

type ImGuiCol_ int32

const (
	ImGuiCol_Text                      ImGuiCol_ = 0
	ImGuiCol_TextDisabled              ImGuiCol_ = 1
	ImGuiCol_WindowBg                  ImGuiCol_ = 2
	ImGuiCol_ChildBg                   ImGuiCol_ = 3
	ImGuiCol_PopupBg                   ImGuiCol_ = 4
	ImGuiCol_Border                    ImGuiCol_ = 5
	ImGuiCol_BorderShadow              ImGuiCol_ = 6
	ImGuiCol_FrameBg                   ImGuiCol_ = 7
	ImGuiCol_FrameBgHovered            ImGuiCol_ = 8
	ImGuiCol_FrameBgActive             ImGuiCol_ = 9
	ImGuiCol_TitleBg                   ImGuiCol_ = 10
	ImGuiCol_TitleBgActive             ImGuiCol_ = 11
	ImGuiCol_TitleBgCollapsed          ImGuiCol_ = 12
	ImGuiCol_MenuBarBg                 ImGuiCol_ = 13
	ImGuiCol_ScrollbarBg               ImGuiCol_ = 14
	ImGuiCol_ScrollbarGrab             ImGuiCol_ = 15
	ImGuiCol_ScrollbarGrabHovered      ImGuiCol_ = 16
	ImGuiCol_ScrollbarGrabActive       ImGuiCol_ = 17
	ImGuiCol_CheckMark                 ImGuiCol_ = 18
	ImGuiCol_SliderGrab                ImGuiCol_ = 19
	ImGuiCol_SliderGrabActive          ImGuiCol_ = 20
	ImGuiCol_Button                    ImGuiCol_ = 21
	ImGuiCol_ButtonHovered             ImGuiCol_ = 22
	ImGuiCol_ButtonActive              ImGuiCol_ = 23
	ImGuiCol_Header                    ImGuiCol_ = 24
	ImGuiCol_HeaderHovered             ImGuiCol_ = 25
	ImGuiCol_HeaderActive              ImGuiCol_ = 26
	ImGuiCol_Separator                 ImGuiCol_ = 27
	ImGuiCol_SeparatorHovered          ImGuiCol_ = 28
	ImGuiCol_SeparatorActive           ImGuiCol_ = 29
	ImGuiCol_ResizeGrip                ImGuiCol_ = 30
	ImGuiCol_ResizeGripHovered         ImGuiCol_ = 31
	ImGuiCol_ResizeGripActive          ImGuiCol_ = 32
	ImGuiCol_InputTextCursor           ImGuiCol_ = 33
	ImGuiCol_TabHovered                ImGuiCol_ = 34
	ImGuiCol_Tab                       ImGuiCol_ = 35
	ImGuiCol_TabSelected               ImGuiCol_ = 36
	ImGuiCol_TabSelectedOverline       ImGuiCol_ = 37
	ImGuiCol_TabDimmed                 ImGuiCol_ = 38
	ImGuiCol_TabDimmedSelected         ImGuiCol_ = 39
	ImGuiCol_TabDimmedSelectedOverline ImGuiCol_ = 40
	ImGuiCol_DockingPreview            ImGuiCol_ = 41
	ImGuiCol_DockingEmptyBg            ImGuiCol_ = 42
	ImGuiCol_PlotLines                 ImGuiCol_ = 43
	ImGuiCol_PlotLinesHovered          ImGuiCol_ = 44
	ImGuiCol_PlotHistogram             ImGuiCol_ = 45
	ImGuiCol_PlotHistogramHovered      ImGuiCol_ = 46
	ImGuiCol_TableHeaderBg             ImGuiCol_ = 47
	ImGuiCol_TableBorderStrong         ImGuiCol_ = 48
	ImGuiCol_TableBorderLight          ImGuiCol_ = 49
	ImGuiCol_TableRowBg                ImGuiCol_ = 50
	ImGuiCol_TableRowBgAlt             ImGuiCol_ = 51
	ImGuiCol_TextLink                  ImGuiCol_ = 52
	ImGuiCol_TextSelectedBg            ImGuiCol_ = 53
	ImGuiCol_TreeLines                 ImGuiCol_ = 54
	ImGuiCol_DragDropTarget            ImGuiCol_ = 55
	ImGuiCol_NavCursor                 ImGuiCol_ = 56
	ImGuiCol_NavWindowingHighlight     ImGuiCol_ = 57
	ImGuiCol_NavWindowingDimBg         ImGuiCol_ = 58
	ImGuiCol_ModalWindowDimBg          ImGuiCol_ = 59
	ImGuiCol_COUNT                     ImGuiCol_ = 60
	ImGuiCol_TabActive                 ImGuiCol_ = ImGuiCol_TabSelected
	ImGuiCol_TabUnfocused              ImGuiCol_ = ImGuiCol_TabDimmed
	ImGuiCol_TabUnfocusedActive        ImGuiCol_ = ImGuiCol_TabDimmedSelected
	ImGuiCol_NavHighlight              ImGuiCol_ = ImGuiCol_NavCursor
)

type ImGuiStyleVar_ int32

const (
	ImGuiStyleVar_Alpha                       ImGuiStyleVar_ = 0
	ImGuiStyleVar_DisabledAlpha               ImGuiStyleVar_ = 1
	ImGuiStyleVar_WindowPadding               ImGuiStyleVar_ = 2
	ImGuiStyleVar_WindowRounding              ImGuiStyleVar_ = 3
	ImGuiStyleVar_WindowBorderSize            ImGuiStyleVar_ = 4
	ImGuiStyleVar_WindowMinSize               ImGuiStyleVar_ = 5
	ImGuiStyleVar_WindowTitleAlign            ImGuiStyleVar_ = 6
	ImGuiStyleVar_ChildRounding               ImGuiStyleVar_ = 7
	ImGuiStyleVar_ChildBorderSize             ImGuiStyleVar_ = 8
	ImGuiStyleVar_PopupRounding               ImGuiStyleVar_ = 9
	ImGuiStyleVar_PopupBorderSize             ImGuiStyleVar_ = 10
	ImGuiStyleVar_FramePadding                ImGuiStyleVar_ = 11
	ImGuiStyleVar_FrameRounding               ImGuiStyleVar_ = 12
	ImGuiStyleVar_FrameBorderSize             ImGuiStyleVar_ = 13
	ImGuiStyleVar_ItemSpacing                 ImGuiStyleVar_ = 14
	ImGuiStyleVar_ItemInnerSpacing            ImGuiStyleVar_ = 15
	ImGuiStyleVar_IndentSpacing               ImGuiStyleVar_ = 16
	ImGuiStyleVar_CellPadding                 ImGuiStyleVar_ = 17
	ImGuiStyleVar_ScrollbarSize               ImGuiStyleVar_ = 18
	ImGuiStyleVar_ScrollbarRounding           ImGuiStyleVar_ = 19
	ImGuiStyleVar_GrabMinSize                 ImGuiStyleVar_ = 20
	ImGuiStyleVar_GrabRounding                ImGuiStyleVar_ = 21
	ImGuiStyleVar_ImageBorderSize             ImGuiStyleVar_ = 22
	ImGuiStyleVar_TabRounding                 ImGuiStyleVar_ = 23
	ImGuiStyleVar_TabBorderSize               ImGuiStyleVar_ = 24
	ImGuiStyleVar_TabBarBorderSize            ImGuiStyleVar_ = 25
	ImGuiStyleVar_TabBarOverlineSize          ImGuiStyleVar_ = 26
	ImGuiStyleVar_TableAngledHeadersAngle     ImGuiStyleVar_ = 27
	ImGuiStyleVar_TableAngledHeadersTextAlign ImGuiStyleVar_ = 28
	ImGuiStyleVar_TreeLinesSize               ImGuiStyleVar_ = 29
	ImGuiStyleVar_TreeLinesRounding           ImGuiStyleVar_ = 30
	ImGuiStyleVar_ButtonTextAlign             ImGuiStyleVar_ = 31
	ImGuiStyleVar_SelectableTextAlign         ImGuiStyleVar_ = 32
	ImGuiStyleVar_SeparatorTextBorderSize     ImGuiStyleVar_ = 33
	ImGuiStyleVar_SeparatorTextAlign          ImGuiStyleVar_ = 34
	ImGuiStyleVar_SeparatorTextPadding        ImGuiStyleVar_ = 35
	ImGuiStyleVar_DockingSeparatorSize        ImGuiStyleVar_ = 36
	ImGuiStyleVar_COUNT                       ImGuiStyleVar_ = 37
)

type ImGuiButtonFlags_ int32

const (
	ImGuiButtonFlags_None              ImGuiButtonFlags_ = 0
	ImGuiButtonFlags_MouseButtonLeft   ImGuiButtonFlags_ = 1 << 0
	ImGuiButtonFlags_MouseButtonRight  ImGuiButtonFlags_ = 1 << 1
	ImGuiButtonFlags_MouseButtonMiddle ImGuiButtonFlags_ = 1 << 2
	ImGuiButtonFlags_MouseButtonMask_  ImGuiButtonFlags_ = ImGuiButtonFlags_MouseButtonLeft | ImGuiButtonFlags_MouseButtonRight | ImGuiButtonFlags_MouseButtonMiddle
	ImGuiButtonFlags_EnableNav         ImGuiButtonFlags_ = 1 << 3
)

type ImGuiColorEditFlags_ int32

const (
	ImGuiColorEditFlags_None             ImGuiColorEditFlags_ = 0
	ImGuiColorEditFlags_NoAlpha          ImGuiColorEditFlags_ = 1 << 1
	ImGuiColorEditFlags_NoPicker         ImGuiColorEditFlags_ = 1 << 2
	ImGuiColorEditFlags_NoOptions        ImGuiColorEditFlags_ = 1 << 3
	ImGuiColorEditFlags_NoSmallPreview   ImGuiColorEditFlags_ = 1 << 4
	ImGuiColorEditFlags_NoInputs         ImGuiColorEditFlags_ = 1 << 5
	ImGuiColorEditFlags_NoTooltip        ImGuiColorEditFlags_ = 1 << 6
	ImGuiColorEditFlags_NoLabel          ImGuiColorEditFlags_ = 1 << 7
	ImGuiColorEditFlags_NoSidePreview    ImGuiColorEditFlags_ = 1 << 8
	ImGuiColorEditFlags_NoDragDrop       ImGuiColorEditFlags_ = 1 << 9
	ImGuiColorEditFlags_NoBorder         ImGuiColorEditFlags_ = 1 << 10
	ImGuiColorEditFlags_AlphaOpaque      ImGuiColorEditFlags_ = 1 << 11
	ImGuiColorEditFlags_AlphaNoBg        ImGuiColorEditFlags_ = 1 << 12
	ImGuiColorEditFlags_AlphaPreviewHalf ImGuiColorEditFlags_ = 1 << 13
	ImGuiColorEditFlags_AlphaBar         ImGuiColorEditFlags_ = 1 << 16
	ImGuiColorEditFlags_HDR              ImGuiColorEditFlags_ = 1 << 19
	ImGuiColorEditFlags_DisplayRGB       ImGuiColorEditFlags_ = 1 << 20
	ImGuiColorEditFlags_DisplayHSV       ImGuiColorEditFlags_ = 1 << 21
	ImGuiColorEditFlags_DisplayHex       ImGuiColorEditFlags_ = 1 << 22
	ImGuiColorEditFlags_Uint8            ImGuiColorEditFlags_ = 1 << 23
	ImGuiColorEditFlags_Float            ImGuiColorEditFlags_ = 1 << 24
	ImGuiColorEditFlags_PickerHueBar     ImGuiColorEditFlags_ = 1 << 25
	ImGuiColorEditFlags_PickerHueWheel   ImGuiColorEditFlags_ = 1 << 26
	ImGuiColorEditFlags_InputRGB         ImGuiColorEditFlags_ = 1 << 27
	ImGuiColorEditFlags_InputHSV         ImGuiColorEditFlags_ = 1 << 28
	ImGuiColorEditFlags_DefaultOptions_  ImGuiColorEditFlags_ = ImGuiColorEditFlags_Uint8 | ImGuiColorEditFlags_DisplayRGB | ImGuiColorEditFlags_InputRGB | ImGuiColorEditFlags_PickerHueBar
	ImGuiColorEditFlags_AlphaMask_       ImGuiColorEditFlags_ = ImGuiColorEditFlags_NoAlpha | ImGuiColorEditFlags_AlphaOpaque | ImGuiColorEditFlags_AlphaNoBg | ImGuiColorEditFlags_AlphaPreviewHalf
	ImGuiColorEditFlags_DisplayMask_     ImGuiColorEditFlags_ = ImGuiColorEditFlags_DisplayRGB | ImGuiColorEditFlags_DisplayHSV | ImGuiColorEditFlags_DisplayHex
	ImGuiColorEditFlags_DataTypeMask_    ImGuiColorEditFlags_ = ImGuiColorEditFlags_Uint8 | ImGuiColorEditFlags_Float
	ImGuiColorEditFlags_PickerMask_      ImGuiColorEditFlags_ = ImGuiColorEditFlags_PickerHueWheel | ImGuiColorEditFlags_PickerHueBar
	ImGuiColorEditFlags_InputMask_       ImGuiColorEditFlags_ = ImGuiColorEditFlags_InputRGB | ImGuiColorEditFlags_InputHSV
	ImGuiColorEditFlags_AlphaPreview     ImGuiColorEditFlags_ = 0
)

type ImGuiSliderFlags_ int32

const (
	ImGuiSliderFlags_None            ImGuiSliderFlags_ = 0
	ImGuiSliderFlags_Logarithmic     ImGuiSliderFlags_ = 1 << 5
	ImGuiSliderFlags_NoRoundToFormat ImGuiSliderFlags_ = 1 << 6
	ImGuiSliderFlags_NoInput         ImGuiSliderFlags_ = 1 << 7
	ImGuiSliderFlags_WrapAround      ImGuiSliderFlags_ = 1 << 8
	ImGuiSliderFlags_ClampOnInput    ImGuiSliderFlags_ = 1 << 9
	ImGuiSliderFlags_ClampZeroRange  ImGuiSliderFlags_ = 1 << 10
	ImGuiSliderFlags_NoSpeedTweaks   ImGuiSliderFlags_ = 1 << 11
	ImGuiSliderFlags_AlwaysClamp     ImGuiSliderFlags_ = ImGuiSliderFlags_ClampOnInput | ImGuiSliderFlags_ClampZeroRange
	ImGuiSliderFlags_InvalidMask_    ImGuiSliderFlags_ = 0x7000000F
)

type ImGuiMouseButton_ int32

const (
	ImGuiMouseButton_Left   ImGuiMouseButton_ = 0
	ImGuiMouseButton_Right  ImGuiMouseButton_ = 1
	ImGuiMouseButton_Middle ImGuiMouseButton_ = 2
	ImGuiMouseButton_COUNT  ImGuiMouseButton_ = 5
)

type ImGuiMouseCursor_ int32

const (
	ImGuiMouseCursor_None       ImGuiMouseCursor_ = -1
	ImGuiMouseCursor_Arrow      ImGuiMouseCursor_ = 0
	ImGuiMouseCursor_TextInput  ImGuiMouseCursor_ = 1
	ImGuiMouseCursor_ResizeAll  ImGuiMouseCursor_ = 2
	ImGuiMouseCursor_ResizeNS   ImGuiMouseCursor_ = 3
	ImGuiMouseCursor_ResizeEW   ImGuiMouseCursor_ = 4
	ImGuiMouseCursor_ResizeNESW ImGuiMouseCursor_ = 5
	ImGuiMouseCursor_ResizeNWSE ImGuiMouseCursor_ = 6
	ImGuiMouseCursor_Hand       ImGuiMouseCursor_ = 7
	ImGuiMouseCursor_Wait       ImGuiMouseCursor_ = 8
	ImGuiMouseCursor_Progress   ImGuiMouseCursor_ = 9
	ImGuiMouseCursor_NotAllowed ImGuiMouseCursor_ = 10
	ImGuiMouseCursor_COUNT      ImGuiMouseCursor_ = 11
)

type ImGuiMouseSource int32

const (
	ImGuiMouseSource_Mouse       ImGuiMouseSource = 0
	ImGuiMouseSource_TouchScreen ImGuiMouseSource = 1
	ImGuiMouseSource_Pen         ImGuiMouseSource = 2
	ImGuiMouseSource_COUNT       ImGuiMouseSource = 3
)

type ImGuiCond_ int32

const (
	ImGuiCond_None         ImGuiCond_ = 0
	ImGuiCond_Always       ImGuiCond_ = 1 << 0
	ImGuiCond_Once         ImGuiCond_ = 1 << 1
	ImGuiCond_FirstUseEver ImGuiCond_ = 1 << 2
	ImGuiCond_Appearing    ImGuiCond_ = 1 << 3
)

type ImGuiTableFlags_ int32

const (
	ImGuiTableFlags_None                       ImGuiTableFlags_ = 0
	ImGuiTableFlags_Resizable                  ImGuiTableFlags_ = 1 << 0
	ImGuiTableFlags_Reorderable                ImGuiTableFlags_ = 1 << 1
	ImGuiTableFlags_Hideable                   ImGuiTableFlags_ = 1 << 2
	ImGuiTableFlags_Sortable                   ImGuiTableFlags_ = 1 << 3
	ImGuiTableFlags_NoSavedSettings            ImGuiTableFlags_ = 1 << 4
	ImGuiTableFlags_ContextMenuInBody          ImGuiTableFlags_ = 1 << 5
	ImGuiTableFlags_RowBg                      ImGuiTableFlags_ = 1 << 6
	ImGuiTableFlags_BordersInnerH              ImGuiTableFlags_ = 1 << 7
	ImGuiTableFlags_BordersOuterH              ImGuiTableFlags_ = 1 << 8
	ImGuiTableFlags_BordersInnerV              ImGuiTableFlags_ = 1 << 9
	ImGuiTableFlags_BordersOuterV              ImGuiTableFlags_ = 1 << 10
	ImGuiTableFlags_BordersH                   ImGuiTableFlags_ = ImGuiTableFlags_BordersInnerH | ImGuiTableFlags_BordersOuterH
	ImGuiTableFlags_BordersV                   ImGuiTableFlags_ = ImGuiTableFlags_BordersInnerV | ImGuiTableFlags_BordersOuterV
	ImGuiTableFlags_BordersInner               ImGuiTableFlags_ = ImGuiTableFlags_BordersInnerV | ImGuiTableFlags_BordersInnerH
	ImGuiTableFlags_BordersOuter               ImGuiTableFlags_ = ImGuiTableFlags_BordersOuterV | ImGuiTableFlags_BordersOuterH
	ImGuiTableFlags_Borders                    ImGuiTableFlags_ = ImGuiTableFlags_BordersInner | ImGuiTableFlags_BordersOuter
	ImGuiTableFlags_NoBordersInBody            ImGuiTableFlags_ = 1 << 11
	ImGuiTableFlags_NoBordersInBodyUntilResize ImGuiTableFlags_ = 1 << 12
	ImGuiTableFlags_SizingFixedFit             ImGuiTableFlags_ = 1 << 13
	ImGuiTableFlags_SizingFixedSame            ImGuiTableFlags_ = 2 << 13
	ImGuiTableFlags_SizingStretchProp          ImGuiTableFlags_ = 3 << 13
	ImGuiTableFlags_SizingStretchSame          ImGuiTableFlags_ = 4 << 13
	ImGuiTableFlags_NoHostExtendX              ImGuiTableFlags_ = 1 << 16
	ImGuiTableFlags_NoHostExtendY              ImGuiTableFlags_ = 1 << 17
	ImGuiTableFlags_NoKeepColumnsVisible       ImGuiTableFlags_ = 1 << 18
	ImGuiTableFlags_PreciseWidths              ImGuiTableFlags_ = 1 << 19
	ImGuiTableFlags_NoClip                     ImGuiTableFlags_ = 1 << 20
	ImGuiTableFlags_PadOuterX                  ImGuiTableFlags_ = 1 << 21
	ImGuiTableFlags_NoPadOuterX                ImGuiTableFlags_ = 1 << 22
	ImGuiTableFlags_NoPadInnerX                ImGuiTableFlags_ = 1 << 23
	ImGuiTableFlags_ScrollX                    ImGuiTableFlags_ = 1 << 24
	ImGuiTableFlags_ScrollY                    ImGuiTableFlags_ = 1 << 25
	ImGuiTableFlags_SortMulti                  ImGuiTableFlags_ = 1 << 26
	ImGuiTableFlags_SortTristate               ImGuiTableFlags_ = 1 << 27
	ImGuiTableFlags_HighlightHoveredColumn     ImGuiTableFlags_ = 1 << 28
	ImGuiTableFlags_SizingMask_                ImGuiTableFlags_ = ImGuiTableFlags_SizingFixedFit | ImGuiTableFlags_SizingFixedSame | ImGuiTableFlags_SizingStretchProp | ImGuiTableFlags_SizingStretchSame
)

type ImGuiTableColumnFlags_ int32

const (
	ImGuiTableColumnFlags_None                 ImGuiTableColumnFlags_ = 0
	ImGuiTableColumnFlags_Disabled             ImGuiTableColumnFlags_ = 1 << 0
	ImGuiTableColumnFlags_DefaultHide          ImGuiTableColumnFlags_ = 1 << 1
	ImGuiTableColumnFlags_DefaultSort          ImGuiTableColumnFlags_ = 1 << 2
	ImGuiTableColumnFlags_WidthStretch         ImGuiTableColumnFlags_ = 1 << 3
	ImGuiTableColumnFlags_WidthFixed           ImGuiTableColumnFlags_ = 1 << 4
	ImGuiTableColumnFlags_NoResize             ImGuiTableColumnFlags_ = 1 << 5
	ImGuiTableColumnFlags_NoReorder            ImGuiTableColumnFlags_ = 1 << 6
	ImGuiTableColumnFlags_NoHide               ImGuiTableColumnFlags_ = 1 << 7
	ImGuiTableColumnFlags_NoClip               ImGuiTableColumnFlags_ = 1 << 8
	ImGuiTableColumnFlags_NoSort               ImGuiTableColumnFlags_ = 1 << 9
	ImGuiTableColumnFlags_NoSortAscending      ImGuiTableColumnFlags_ = 1 << 10
	ImGuiTableColumnFlags_NoSortDescending     ImGuiTableColumnFlags_ = 1 << 11
	ImGuiTableColumnFlags_NoHeaderLabel        ImGuiTableColumnFlags_ = 1 << 12
	ImGuiTableColumnFlags_NoHeaderWidth        ImGuiTableColumnFlags_ = 1 << 13
	ImGuiTableColumnFlags_PreferSortAscending  ImGuiTableColumnFlags_ = 1 << 14
	ImGuiTableColumnFlags_PreferSortDescending ImGuiTableColumnFlags_ = 1 << 15
	ImGuiTableColumnFlags_IndentEnable         ImGuiTableColumnFlags_ = 1 << 16
	ImGuiTableColumnFlags_IndentDisable        ImGuiTableColumnFlags_ = 1 << 17
	ImGuiTableColumnFlags_AngledHeader         ImGuiTableColumnFlags_ = 1 << 18
	ImGuiTableColumnFlags_IsEnabled            ImGuiTableColumnFlags_ = 1 << 24
	ImGuiTableColumnFlags_IsVisible            ImGuiTableColumnFlags_ = 1 << 25
	ImGuiTableColumnFlags_IsSorted             ImGuiTableColumnFlags_ = 1 << 26
	ImGuiTableColumnFlags_IsHovered            ImGuiTableColumnFlags_ = 1 << 27
	ImGuiTableColumnFlags_WidthMask_           ImGuiTableColumnFlags_ = ImGuiTableColumnFlags_WidthStretch | ImGuiTableColumnFlags_WidthFixed
	ImGuiTableColumnFlags_IndentMask_          ImGuiTableColumnFlags_ = ImGuiTableColumnFlags_IndentEnable | ImGuiTableColumnFlags_IndentDisable
	ImGuiTableColumnFlags_StatusMask_          ImGuiTableColumnFlags_ = ImGuiTableColumnFlags_IsEnabled | ImGuiTableColumnFlags_IsVisible | ImGuiTableColumnFlags_IsSorted | ImGuiTableColumnFlags_IsHovered
	ImGuiTableColumnFlags_NoDirectResize_      ImGuiTableColumnFlags_ = 1 << 30
)

type ImGuiTableRowFlags_ int32

const (
	ImGuiTableRowFlags_None    ImGuiTableRowFlags_ = 0
	ImGuiTableRowFlags_Headers ImGuiTableRowFlags_ = 1 << 0
)

type ImGuiTableBgTarget_ int32

const (
	ImGuiTableBgTarget_None   ImGuiTableBgTarget_ = 0
	ImGuiTableBgTarget_RowBg0 ImGuiTableBgTarget_ = 1
	ImGuiTableBgTarget_RowBg1 ImGuiTableBgTarget_ = 2
	ImGuiTableBgTarget_CellBg ImGuiTableBgTarget_ = 3
)

type ImGuiMultiSelectFlags_ int32

const (
	ImGuiMultiSelectFlags_None                  ImGuiMultiSelectFlags_ = 0
	ImGuiMultiSelectFlags_SingleSelect          ImGuiMultiSelectFlags_ = 1 << 0
	ImGuiMultiSelectFlags_NoSelectAll           ImGuiMultiSelectFlags_ = 1 << 1
	ImGuiMultiSelectFlags_NoRangeSelect         ImGuiMultiSelectFlags_ = 1 << 2
	ImGuiMultiSelectFlags_NoAutoSelect          ImGuiMultiSelectFlags_ = 1 << 3
	ImGuiMultiSelectFlags_NoAutoClear           ImGuiMultiSelectFlags_ = 1 << 4
	ImGuiMultiSelectFlags_NoAutoClearOnReselect ImGuiMultiSelectFlags_ = 1 << 5
	ImGuiMultiSelectFlags_BoxSelect1d           ImGuiMultiSelectFlags_ = 1 << 6
	ImGuiMultiSelectFlags_BoxSelect2d           ImGuiMultiSelectFlags_ = 1 << 7
	ImGuiMultiSelectFlags_BoxSelectNoScroll     ImGuiMultiSelectFlags_ = 1 << 8
	ImGuiMultiSelectFlags_ClearOnEscape         ImGuiMultiSelectFlags_ = 1 << 9
	ImGuiMultiSelectFlags_ClearOnClickVoid      ImGuiMultiSelectFlags_ = 1 << 10
	ImGuiMultiSelectFlags_ScopeWindow           ImGuiMultiSelectFlags_ = 1 << 11
	ImGuiMultiSelectFlags_ScopeRect             ImGuiMultiSelectFlags_ = 1 << 12
	ImGuiMultiSelectFlags_SelectOnClick         ImGuiMultiSelectFlags_ = 1 << 13
	ImGuiMultiSelectFlags_SelectOnClickRelease  ImGuiMultiSelectFlags_ = 1 << 14
	ImGuiMultiSelectFlags_NavWrapX              ImGuiMultiSelectFlags_ = 1 << 16
)

type ImGuiSelectionRequestType int32

const (
	ImGuiSelectionRequestType_None     ImGuiSelectionRequestType = 0
	ImGuiSelectionRequestType_SetAll   ImGuiSelectionRequestType = 1
	ImGuiSelectionRequestType_SetRange ImGuiSelectionRequestType = 2
)

type ImDrawFlags_ int32

const (
	ImDrawFlags_None                    ImDrawFlags_ = 0
	ImDrawFlags_Closed                  ImDrawFlags_ = 1 << 0
	ImDrawFlags_RoundCornersTopLeft     ImDrawFlags_ = 1 << 4
	ImDrawFlags_RoundCornersTopRight    ImDrawFlags_ = 1 << 5
	ImDrawFlags_RoundCornersBottomLeft  ImDrawFlags_ = 1 << 6
	ImDrawFlags_RoundCornersBottomRight ImDrawFlags_ = 1 << 7
	ImDrawFlags_RoundCornersNone        ImDrawFlags_ = 1 << 8
	ImDrawFlags_RoundCornersTop         ImDrawFlags_ = ImDrawFlags_RoundCornersTopLeft | ImDrawFlags_RoundCornersTopRight
	ImDrawFlags_RoundCornersBottom      ImDrawFlags_ = ImDrawFlags_RoundCornersBottomLeft | ImDrawFlags_RoundCornersBottomRight
	ImDrawFlags_RoundCornersLeft        ImDrawFlags_ = ImDrawFlags_RoundCornersBottomLeft | ImDrawFlags_RoundCornersTopLeft
	ImDrawFlags_RoundCornersRight       ImDrawFlags_ = ImDrawFlags_RoundCornersBottomRight | ImDrawFlags_RoundCornersTopRight
	ImDrawFlags_RoundCornersAll         ImDrawFlags_ = ImDrawFlags_RoundCornersTopLeft | ImDrawFlags_RoundCornersTopRight | ImDrawFlags_RoundCornersBottomLeft | ImDrawFlags_RoundCornersBottomRight
	ImDrawFlags_RoundCornersDefault_    ImDrawFlags_ = ImDrawFlags_RoundCornersAll
	ImDrawFlags_RoundCornersMask_       ImDrawFlags_ = ImDrawFlags_RoundCornersAll | ImDrawFlags_RoundCornersNone
)

type ImDrawListFlags_ int32

const (
	ImDrawListFlags_None                   ImDrawListFlags_ = 0
	ImDrawListFlags_AntiAliasedLines       ImDrawListFlags_ = 1 << 0
	ImDrawListFlags_AntiAliasedLinesUseTex ImDrawListFlags_ = 1 << 1
	ImDrawListFlags_AntiAliasedFill        ImDrawListFlags_ = 1 << 2
	ImDrawListFlags_AllowVtxOffset         ImDrawListFlags_ = 1 << 3
)

type ImTextureFormat int32

const (
	ImTextureFormat_RGBA32 ImTextureFormat = 0
	ImTextureFormat_Alpha8 ImTextureFormat = 1
)

type ImTextureStatus int32

const (
	ImTextureStatus_OK          ImTextureStatus = 0
	ImTextureStatus_Destroyed   ImTextureStatus = 1
	ImTextureStatus_WantCreate  ImTextureStatus = 2
	ImTextureStatus_WantUpdates ImTextureStatus = 3
	ImTextureStatus_WantDestroy ImTextureStatus = 4
)

type ImFontAtlasFlags_ int32

const (
	ImFontAtlasFlags_None               ImFontAtlasFlags_ = 0
	ImFontAtlasFlags_NoPowerOfTwoHeight ImFontAtlasFlags_ = 1 << 0
	ImFontAtlasFlags_NoMouseCursors     ImFontAtlasFlags_ = 1 << 1
	ImFontAtlasFlags_NoBakedLines       ImFontAtlasFlags_ = 1 << 2
)

type ImFontFlags_ int32

const (
	ImFontFlags_None           ImFontFlags_ = 0
	ImFontFlags_NoLoadError    ImFontFlags_ = 1 << 1
	ImFontFlags_NoLoadGlyphs   ImFontFlags_ = 1 << 2
	ImFontFlags_LockBakedSizes ImFontFlags_ = 1 << 3
)

type ImGuiViewportFlags_ int32

const (
	ImGuiViewportFlags_None                ImGuiViewportFlags_ = 0
	ImGuiViewportFlags_IsPlatformWindow    ImGuiViewportFlags_ = 1 << 0
	ImGuiViewportFlags_IsPlatformMonitor   ImGuiViewportFlags_ = 1 << 1
	ImGuiViewportFlags_OwnedByApp          ImGuiViewportFlags_ = 1 << 2
	ImGuiViewportFlags_NoDecoration        ImGuiViewportFlags_ = 1 << 3
	ImGuiViewportFlags_NoTaskBarIcon       ImGuiViewportFlags_ = 1 << 4
	ImGuiViewportFlags_NoFocusOnAppearing  ImGuiViewportFlags_ = 1 << 5
	ImGuiViewportFlags_NoFocusOnClick      ImGuiViewportFlags_ = 1 << 6
	ImGuiViewportFlags_NoInputs            ImGuiViewportFlags_ = 1 << 7
	ImGuiViewportFlags_NoRendererClear     ImGuiViewportFlags_ = 1 << 8
	ImGuiViewportFlags_NoAutoMerge         ImGuiViewportFlags_ = 1 << 9
	ImGuiViewportFlags_TopMost             ImGuiViewportFlags_ = 1 << 10
	ImGuiViewportFlags_CanHostOtherWindows ImGuiViewportFlags_ = 1 << 11
	ImGuiViewportFlags_IsMinimized         ImGuiViewportFlags_ = 1 << 12
	ImGuiViewportFlags_IsFocused           ImGuiViewportFlags_ = 1 << 13
)
