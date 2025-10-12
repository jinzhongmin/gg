package gdk

type AxisUse int32
type CicpRange int32
type CrossingMode int32
type DevicePadFeature int32
type DeviceToolType int32
type DmabufError int32
type DragCancelReason int32
type EventType int32
type FullscreenMode int32
type GLError int32
type Gravity int32
type InputSource int32
type KeyMatch int32
type MemoryFormat int32
type NotifyType int32
type ScrollDirection int32
type ScrollUnit int32
type SubpixelLayout int32
type SurfaceEdge int32
type TextureError int32
type TitlebarGesture int32
type TouchpadGesturePhase int32
type VulkanError int32
type AnchorHints int32
type AxisFlags int32
type DragAction int32
type FrameClockPhase int32
type GLAPI int32
type ModifierType int32
type PaintableFlags int32
type SeatCapabilities int32
type ToplevelState int32

const (
	AxisUseIgnore   AxisUse = 0
	AxisUseX        AxisUse = 1
	AxisUseY        AxisUse = 2
	AxisUseDeltaX   AxisUse = 3
	AxisUseDeltaY   AxisUse = 4
	AxisUsePressure AxisUse = 5
	AxisUseXtilt    AxisUse = 6
	AxisUseYtilt    AxisUse = 7
	AxisUseWheel    AxisUse = 8
	AxisUseDistance AxisUse = 9
	AxisUseRotation AxisUse = 10
	AxisUseSlider   AxisUse = 11
	AxisUseLast     AxisUse = 12

	CicpRangeNarrow CicpRange = 0
	CicpRangeFull   CicpRange = 1

	CrossingModeNormal       CrossingMode = 0
	CrossingModeGrab         CrossingMode = 1
	CrossingModeUngrab       CrossingMode = 2
	CrossingModeGtkGrab      CrossingMode = 3
	CrossingModeGtkUngrab    CrossingMode = 4
	CrossingModeStateChanged CrossingMode = 5
	CrossingModeTouchBegin   CrossingMode = 6
	CrossingModeTouchEnd     CrossingMode = 7
	CrossingModeDeviceSwitch CrossingMode = 8

	DevicePadFeatureButton DevicePadFeature = 0
	DevicePadFeatureRing   DevicePadFeature = 1
	DevicePadFeatureStrip  DevicePadFeature = 2

	DeviceToolTypeUnknown  DeviceToolType = 0
	DeviceToolTypePen      DeviceToolType = 1
	DeviceToolTypeEraser   DeviceToolType = 2
	DeviceToolTypeBrush    DeviceToolType = 3
	DeviceToolTypePencil   DeviceToolType = 4
	DeviceToolTypeAirbrush DeviceToolType = 5
	DeviceToolTypeMouse    DeviceToolType = 6
	DeviceToolTypeLens     DeviceToolType = 7

	DmabufErrorNotAvailable      DmabufError = 0
	DmabufErrorUnsupportedFormat DmabufError = 1
	DmabufErrorCreationFailed    DmabufError = 2

	DragCancelReasonNoTarget      DragCancelReason = 0
	DragCancelReasonUserCancelled DragCancelReason = 1
	DragCancelReasonError         DragCancelReason = 2

	EventTypeDelete           EventType = 0
	EventTypeMotionNotify     EventType = 1
	EventTypeButtonPress      EventType = 2
	EventTypeButtonRelease    EventType = 3
	EventTypeKeyPress         EventType = 4
	EventTypeKeyRelease       EventType = 5
	EventTypeEnterNotify      EventType = 6
	EventTypeLeaveNotify      EventType = 7
	EventTypeFocusChange      EventType = 8
	EventTypeProximityIn      EventType = 9
	EventTypeProximityOut     EventType = 10
	EventTypeDragEnter        EventType = 11
	EventTypeDragLeave        EventType = 12
	EventTypeDragMotion       EventType = 13
	EventTypeDropStart        EventType = 14
	EventTypeScroll           EventType = 15
	EventTypeGrabBroken       EventType = 16
	EventTypeTouchBegin       EventType = 17
	EventTypeTouchUpdate      EventType = 18
	EventTypeTouchEnd         EventType = 19
	EventTypeTouchCancel      EventType = 20
	EventTypeTouchpadSwipe    EventType = 21
	EventTypeTouchpadPinch    EventType = 22
	EventTypePadButtonPress   EventType = 23
	EventTypePadButtonRelease EventType = 24
	EventTypePadRing          EventType = 25
	EventTypePadStrip         EventType = 26
	EventTypePadGroupMode     EventType = 27
	EventTypeTouchpadHold     EventType = 28
	EventTypeEventLast        EventType = 29

	FullscreenModeCurrentMonitor FullscreenMode = 0
	FullscreenModeAllMonitors    FullscreenMode = 1

	GLErrorNotAvailable       GLError = 0
	GLErrorUnsupportedFormat  GLError = 1
	GLErrorUnsupportedProfile GLError = 2
	GLErrorCompilationFailed  GLError = 3
	GLErrorLinkFailed         GLError = 4

	GravityNorthWest Gravity = 1
	GravityNorth     Gravity = 2
	GravityNorthEast Gravity = 3
	GravityWest      Gravity = 4
	GravityCenter    Gravity = 5
	GravityEast      Gravity = 6
	GravitySouthWest Gravity = 7
	GravitySouth     Gravity = 8
	GravitySouthEast Gravity = 9
	GravityStatic    Gravity = 10

	InputSourceMouse       InputSource = 0
	InputSourcePen         InputSource = 1
	InputSourceKeyboard    InputSource = 2
	InputSourceTouchscreen InputSource = 3
	InputSourceTouchpad    InputSource = 4
	InputSourceTrackpoint  InputSource = 5
	InputSourceTabletPad   InputSource = 6

	KeyMatchNone    KeyMatch = 0
	KeyMatchPartial KeyMatch = 1
	KeyMatchExact   KeyMatch = 2

	MemoryFormatB8g8r8a8Premultiplied          MemoryFormat = 0
	MemoryFormatA8r8g8b8Premultiplied          MemoryFormat = 1
	MemoryFormatR8g8b8a8Premultiplied          MemoryFormat = 2
	MemoryFormatB8g8r8a8                       MemoryFormat = 3
	MemoryFormatA8r8g8b8                       MemoryFormat = 4
	MemoryFormatR8g8b8a8                       MemoryFormat = 5
	MemoryFormatA8b8g8r8                       MemoryFormat = 6
	MemoryFormatR8g8b8                         MemoryFormat = 7
	MemoryFormatB8g8r8                         MemoryFormat = 8
	MemoryFormatR16g16b16                      MemoryFormat = 9
	MemoryFormatR16g16b16a16Premultiplied      MemoryFormat = 10
	MemoryFormatR16g16b16a16                   MemoryFormat = 11
	MemoryFormatR16g16b16Float                 MemoryFormat = 12
	MemoryFormatR16g16b16a16FloatPremultiplied MemoryFormat = 13
	MemoryFormatR16g16b16a16Float              MemoryFormat = 14
	MemoryFormatR32g32b32Float                 MemoryFormat = 15
	MemoryFormatR32g32b32a32FloatPremultiplied MemoryFormat = 16
	MemoryFormatR32g32b32a32Float              MemoryFormat = 17
	MemoryFormatG8a8Premultiplied              MemoryFormat = 18
	MemoryFormatG8a8                           MemoryFormat = 19
	MemoryFormatG8                             MemoryFormat = 20
	MemoryFormatG16a16Premultiplied            MemoryFormat = 21
	MemoryFormatG16a16                         MemoryFormat = 22
	MemoryFormatG16                            MemoryFormat = 23
	MemoryFormatA8                             MemoryFormat = 24
	MemoryFormatA16                            MemoryFormat = 25
	MemoryFormatA16Float                       MemoryFormat = 26
	MemoryFormatA32Float                       MemoryFormat = 27
	MemoryFormatA8b8g8r8Premultiplied          MemoryFormat = 28
	MemoryFormatB8g8r8x8                       MemoryFormat = 29
	MemoryFormatX8r8g8b8                       MemoryFormat = 30
	MemoryFormatR8g8b8x8                       MemoryFormat = 31
	MemoryFormatX8b8g8r8                       MemoryFormat = 32
	MemoryFormatNFormats                       MemoryFormat = 33

	NotifyTypeAncestor         NotifyType = 0
	NotifyTypeVirtual          NotifyType = 1
	NotifyTypeInferior         NotifyType = 2
	NotifyTypeNonlinear        NotifyType = 3
	NotifyTypeNonlinearVirtual NotifyType = 4
	NotifyTypeUnknown          NotifyType = 5

	ScrollDirectionUp     ScrollDirection = 0
	ScrollDirectionDown   ScrollDirection = 1
	ScrollDirectionLeft   ScrollDirection = 2
	ScrollDirectionRight  ScrollDirection = 3
	ScrollDirectionSmooth ScrollDirection = 4

	ScrollUnitWheel   ScrollUnit = 0
	ScrollUnitSurface ScrollUnit = 1

	SubpixelLayoutUnknown       SubpixelLayout = 0
	SubpixelLayoutNone          SubpixelLayout = 1
	SubpixelLayoutHorizontalRgb SubpixelLayout = 2
	SubpixelLayoutHorizontalBgr SubpixelLayout = 3
	SubpixelLayoutVerticalRgb   SubpixelLayout = 4
	SubpixelLayoutVerticalBgr   SubpixelLayout = 5

	SurfaceEdgeNorthWest SurfaceEdge = 0
	SurfaceEdgeNorth     SurfaceEdge = 1
	SurfaceEdgeNorthEast SurfaceEdge = 2
	SurfaceEdgeWest      SurfaceEdge = 3
	SurfaceEdgeEast      SurfaceEdge = 4
	SurfaceEdgeSouthWest SurfaceEdge = 5
	SurfaceEdgeSouth     SurfaceEdge = 6
	SurfaceEdgeSouthEast SurfaceEdge = 7

	TextureErrorTooLarge           TextureError = 0
	TextureErrorCorruptImage       TextureError = 1
	TextureErrorUnsupportedContent TextureError = 2
	TextureErrorUnsupportedFormat  TextureError = 3

	TitlebarGestureDoubleClick TitlebarGesture = 1
	TitlebarGestureRightClick  TitlebarGesture = 2
	TitlebarGestureMiddleClick TitlebarGesture = 3

	TouchpadGesturePhaseBegin  TouchpadGesturePhase = 0
	TouchpadGesturePhaseUpdate TouchpadGesturePhase = 1
	TouchpadGesturePhaseEnd    TouchpadGesturePhase = 2
	TouchpadGesturePhaseCancel TouchpadGesturePhase = 3

	VulkanErrorUnsupported  VulkanError = 0
	VulkanErrorNotAvailable VulkanError = 1

	AnchorHintsFlipX   AnchorHints = 1
	AnchorHintsFlipY   AnchorHints = 2
	AnchorHintsSlideX  AnchorHints = 4
	AnchorHintsSlideY  AnchorHints = 8
	AnchorHintsResizeX AnchorHints = 16
	AnchorHintsResizeY AnchorHints = 32
	AnchorHintsFlip    AnchorHints = 3
	AnchorHintsSlide   AnchorHints = 12
	AnchorHintsResize  AnchorHints = 48

	AxisFlagsX        AxisFlags = 2
	AxisFlagsY        AxisFlags = 4
	AxisFlagsDeltaX   AxisFlags = 8
	AxisFlagsDeltaY   AxisFlags = 16
	AxisFlagsPressure AxisFlags = 32
	AxisFlagsXtilt    AxisFlags = 64
	AxisFlagsYtilt    AxisFlags = 128
	AxisFlagsWheel    AxisFlags = 256
	AxisFlagsDistance AxisFlags = 512
	AxisFlagsRotation AxisFlags = 1024
	AxisFlagsSlider   AxisFlags = 2048

	DragActionCopy DragAction = 1
	DragActionMove DragAction = 2
	DragActionLink DragAction = 4
	DragActionAsk  DragAction = 8

	FrameClockPhaseNone         FrameClockPhase = 0
	FrameClockPhaseFlushEvents  FrameClockPhase = 1
	FrameClockPhaseBeforePaint  FrameClockPhase = 2
	FrameClockPhaseUpdate       FrameClockPhase = 4
	FrameClockPhaseLayout       FrameClockPhase = 8
	FrameClockPhasePaint        FrameClockPhase = 16
	FrameClockPhaseResumeEvents FrameClockPhase = 32
	FrameClockPhaseAfterPaint   FrameClockPhase = 64

	GLAPIGl   GLAPI = 1
	GLAPIGles GLAPI = 2

	ModifierTypeNoModifierMask ModifierType = 0
	ModifierTypeShiftMask      ModifierType = 1
	ModifierTypeLockMask       ModifierType = 2
	ModifierTypeControlMask    ModifierType = 4
	ModifierTypeAltMask        ModifierType = 8
	ModifierTypeButton1Mask    ModifierType = 256
	ModifierTypeButton2Mask    ModifierType = 512
	ModifierTypeButton3Mask    ModifierType = 1024
	ModifierTypeButton4Mask    ModifierType = 2048
	ModifierTypeButton5Mask    ModifierType = 4096
	ModifierTypeSuperMask      ModifierType = 67108864
	ModifierTypeHyperMask      ModifierType = 134217728
	ModifierTypeMetaMask       ModifierType = 268435456

	PaintableFlagsSize     PaintableFlags = 1
	PaintableFlagsContents PaintableFlags = 2

	SeatCapabilitiesNone         SeatCapabilities = 0
	SeatCapabilitiesPointer      SeatCapabilities = 1
	SeatCapabilitiesTouch        SeatCapabilities = 2
	SeatCapabilitiesTabletStylus SeatCapabilities = 4
	SeatCapabilitiesKeyboard     SeatCapabilities = 8
	SeatCapabilitiesTabletPad    SeatCapabilities = 16
	SeatCapabilitiesAllPointing  SeatCapabilities = 7
	SeatCapabilitiesAll          SeatCapabilities = 31

	ToplevelStateMinimized       ToplevelState = 1
	ToplevelStateMaximized       ToplevelState = 2
	ToplevelStateSticky          ToplevelState = 4
	ToplevelStateFullscreen      ToplevelState = 8
	ToplevelStateAbove           ToplevelState = 16
	ToplevelStateBelow           ToplevelState = 32
	ToplevelStateFocused         ToplevelState = 64
	ToplevelStateTiled           ToplevelState = 128
	ToplevelStateTopTiled        ToplevelState = 256
	ToplevelStateTopResizable    ToplevelState = 512
	ToplevelStateRightTiled      ToplevelState = 1024
	ToplevelStateRightResizable  ToplevelState = 2048
	ToplevelStateBottomTiled     ToplevelState = 4096
	ToplevelStateBottomResizable ToplevelState = 8192
	ToplevelStateLeftTiled       ToplevelState = 16384
	ToplevelStateLeftResizable   ToplevelState = 32768
	ToplevelStateSuspended       ToplevelState = 65536
)
