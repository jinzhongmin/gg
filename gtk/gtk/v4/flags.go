package gtk

type AccessibleAnnouncementPriority int32
type AccessibleAutocomplete int32
type AccessibleInvalidState int32
type AccessiblePlatformState int32
type AccessibleProperty int32
type AccessibleRelation int32
type AccessibleRole int32
type AccessibleSort int32
type AccessibleState int32
type AccessibleTextContentChange int32
type AccessibleTextGranularity int32
type AccessibleTristate int32
type Align int32
type ArrowType int32
type AssistantPageType int32
type BaselinePosition int32
type BorderStyle int32
type BuilderError int32
type ButtonsType int32
type CellRendererAccelMode int32
type CellRendererMode int32
type Collation int32
type ConstraintAttribute int32
type ConstraintRelation int32
type ConstraintStrength int32
type ConstraintVflParserError int32
type ContentFit int32
type CornerType int32
type CssParserError int32
type CssParserWarning int32
type DeleteType int32
type DialogError int32
type DirectionType int32
type EditableProperties int32
type EntryIconPosition int32
type EventSequenceState int32
type FileChooserAction int32
type FileChooserError int32
type FilterChange int32
type FilterMatch int32
type FontLevel int32
type FontRendering int32
type GraphicsOffloadEnabled int32
type IconSize int32
type IconThemeError int32
type IconViewDropPosition int32
type ImageType int32
type InputPurpose int32
type InscriptionOverflow int32
type Justification int32
type LevelBarMode int32
type License int32
type ListTabBehavior int32
type MessageType int32
type MovementStep int32
type NaturalWrapMode int32
type NotebookTab int32
type NumberUpLayout int32
type Ordering int32
type Orientation int32
type Overflow int32
type PackType int32
type PadActionType int32
type PageOrientation int32
type PageSet int32
type PanDirection int32
type PolicyType int32
type PositionType int32
type PrintDuplex int32
type PrintError int32
type PrintOperationAction int32
type PrintOperationResult int32
type PrintPages int32
type PrintQuality int32
type PrintStatus int32
type PropagationLimit int32
type PropagationPhase int32
type RecentManagerError int32
type ResponseType int32
type RevealerTransitionType int32
type ScrollStep int32
type ScrollType int32
type ScrollablePolicy int32
type SelectionMode int32
type SensitivityType int32
type ShortcutScope int32
type ShortcutType int32
type SizeGroupMode int32
type SizeRequestMode int32
type SortType int32
type SorterChange int32
type SorterOrder int32
type SpinButtonUpdatePolicy int32
type SpinType int32
type StackTransitionType int32
type StringFilterMatchMode int32
type SymbolicColor int32
type SystemSetting int32
type TextDirection int32
type TextExtendSelection int32
type TextViewLayer int32
type TextWindowType int32
type TreeViewColumnSizing int32
type TreeViewDropPosition int32
type TreeViewGridLines int32
type Unit int32
type WrapMode int32
type ApplicationInhibitFlags int32
type BuilderClosureFlags int32
type CellRendererState int32
type DebugFlags int32
type DialogFlags int32
type EventControllerScrollFlags int32
type FontChooserLevel int32
type IconLookupFlags int32
type InputHints int32
type ListScrollFlags int32
type PickFlags int32
type PopoverMenuFlags int32
type ShortcutActionFlags int32
type StateFlags int32
type StyleContextPrintFlags int32
type TextBufferNotifyFlags int32
type TextSearchFlags int32
type TreeModelFlags int32

const (
	AccessibleAnnouncementPriorityLow    AccessibleAnnouncementPriority = 0
	AccessibleAnnouncementPriorityMedium AccessibleAnnouncementPriority = 1
	AccessibleAnnouncementPriorityHigh   AccessibleAnnouncementPriority = 2

	AccessibleAutocompleteNone   AccessibleAutocomplete = 0
	AccessibleAutocompleteInline AccessibleAutocomplete = 1
	AccessibleAutocompleteList   AccessibleAutocomplete = 2
	AccessibleAutocompleteBoth   AccessibleAutocomplete = 3

	AccessibleInvalidStateFalse    AccessibleInvalidState = 0
	AccessibleInvalidStateTrue     AccessibleInvalidState = 1
	AccessibleInvalidStateGrammar  AccessibleInvalidState = 2
	AccessibleInvalidStateSpelling AccessibleInvalidState = 3

	AccessiblePlatformStateFocusable AccessiblePlatformState = 0
	AccessiblePlatformStateFocused   AccessiblePlatformState = 1
	AccessiblePlatformStateActive    AccessiblePlatformState = 2

	AccessiblePropertyAutocomplete    AccessibleProperty = 0
	AccessiblePropertyDescription     AccessibleProperty = 1
	AccessiblePropertyHasPopup        AccessibleProperty = 2
	AccessiblePropertyKeyShortcuts    AccessibleProperty = 3
	AccessiblePropertyLabel           AccessibleProperty = 4
	AccessiblePropertyLevel           AccessibleProperty = 5
	AccessiblePropertyModal           AccessibleProperty = 6
	AccessiblePropertyMultiLine       AccessibleProperty = 7
	AccessiblePropertyMultiSelectable AccessibleProperty = 8
	AccessiblePropertyOrientation     AccessibleProperty = 9
	AccessiblePropertyPlaceholder     AccessibleProperty = 10
	AccessiblePropertyReadOnly        AccessibleProperty = 11
	AccessiblePropertyRequired        AccessibleProperty = 12
	AccessiblePropertyRoleDescription AccessibleProperty = 13
	AccessiblePropertySort            AccessibleProperty = 14
	AccessiblePropertyValueMax        AccessibleProperty = 15
	AccessiblePropertyValueMin        AccessibleProperty = 16
	AccessiblePropertyValueNow        AccessibleProperty = 17
	AccessiblePropertyValueText       AccessibleProperty = 18
	AccessiblePropertyHelpText        AccessibleProperty = 19

	AccessibleRelationActiveDescendant AccessibleRelation = 0
	AccessibleRelationColCount         AccessibleRelation = 1
	AccessibleRelationColIndex         AccessibleRelation = 2
	AccessibleRelationColIndexText     AccessibleRelation = 3
	AccessibleRelationColSpan          AccessibleRelation = 4
	AccessibleRelationControls         AccessibleRelation = 5
	AccessibleRelationDescribedBy      AccessibleRelation = 6
	AccessibleRelationDetails          AccessibleRelation = 7
	AccessibleRelationErrorMessage     AccessibleRelation = 8
	AccessibleRelationFlowTo           AccessibleRelation = 9
	AccessibleRelationLabelledBy       AccessibleRelation = 10
	AccessibleRelationOwns             AccessibleRelation = 11
	AccessibleRelationPosInSet         AccessibleRelation = 12
	AccessibleRelationRowCount         AccessibleRelation = 13
	AccessibleRelationRowIndex         AccessibleRelation = 14
	AccessibleRelationRowIndexText     AccessibleRelation = 15
	AccessibleRelationRowSpan          AccessibleRelation = 16
	AccessibleRelationSetSize          AccessibleRelation = 17

	AccessibleRoleAlert            AccessibleRole = 0
	AccessibleRoleAlertDialog      AccessibleRole = 1
	AccessibleRoleBanner           AccessibleRole = 2
	AccessibleRoleButton           AccessibleRole = 3
	AccessibleRoleCaption          AccessibleRole = 4
	AccessibleRoleCell             AccessibleRole = 5
	AccessibleRoleCheckbox         AccessibleRole = 6
	AccessibleRoleColumnHeader     AccessibleRole = 7
	AccessibleRoleComboBox         AccessibleRole = 8
	AccessibleRoleCommand          AccessibleRole = 9
	AccessibleRoleComposite        AccessibleRole = 10
	AccessibleRoleDialog           AccessibleRole = 11
	AccessibleRoleDocument         AccessibleRole = 12
	AccessibleRoleFeed             AccessibleRole = 13
	AccessibleRoleForm             AccessibleRole = 14
	AccessibleRoleGeneric          AccessibleRole = 15
	AccessibleRoleGrid             AccessibleRole = 16
	AccessibleRoleGridCell         AccessibleRole = 17
	AccessibleRoleGroup            AccessibleRole = 18
	AccessibleRoleHeading          AccessibleRole = 19
	AccessibleRoleImg              AccessibleRole = 20
	AccessibleRoleInput            AccessibleRole = 21
	AccessibleRoleLabel            AccessibleRole = 22
	AccessibleRoleLandmark         AccessibleRole = 23
	AccessibleRoleLegend           AccessibleRole = 24
	AccessibleRoleLink             AccessibleRole = 25
	AccessibleRoleList             AccessibleRole = 26
	AccessibleRoleListBox          AccessibleRole = 27
	AccessibleRoleListItem         AccessibleRole = 28
	AccessibleRoleLog              AccessibleRole = 29
	AccessibleRoleMain             AccessibleRole = 30
	AccessibleRoleMarquee          AccessibleRole = 31
	AccessibleRoleMath             AccessibleRole = 32
	AccessibleRoleMeter            AccessibleRole = 33
	AccessibleRoleMenu             AccessibleRole = 34
	AccessibleRoleMenuBar          AccessibleRole = 35
	AccessibleRoleMenuItem         AccessibleRole = 36
	AccessibleRoleMenuItemCheckbox AccessibleRole = 37
	AccessibleRoleMenuItemRadio    AccessibleRole = 38
	AccessibleRoleNavigation       AccessibleRole = 39
	AccessibleRoleNone             AccessibleRole = 40
	AccessibleRoleNote             AccessibleRole = 41
	AccessibleRoleOption           AccessibleRole = 42
	AccessibleRolePresentation     AccessibleRole = 43
	AccessibleRoleProgressBar      AccessibleRole = 44
	AccessibleRoleRadio            AccessibleRole = 45
	AccessibleRoleRadioGroup       AccessibleRole = 46
	AccessibleRoleRange            AccessibleRole = 47
	AccessibleRoleRegion           AccessibleRole = 48
	AccessibleRoleRow              AccessibleRole = 49
	AccessibleRoleRowGroup         AccessibleRole = 50
	AccessibleRoleRowHeader        AccessibleRole = 51
	AccessibleRoleScrollbar        AccessibleRole = 52
	AccessibleRoleSearch           AccessibleRole = 53
	AccessibleRoleSearchBox        AccessibleRole = 54
	AccessibleRoleSection          AccessibleRole = 55
	AccessibleRoleSectionHead      AccessibleRole = 56
	AccessibleRoleSelect           AccessibleRole = 57
	AccessibleRoleSeparator        AccessibleRole = 58
	AccessibleRoleSlider           AccessibleRole = 59
	AccessibleRoleSpinButton       AccessibleRole = 60
	AccessibleRoleStatus           AccessibleRole = 61
	AccessibleRoleStructure        AccessibleRole = 62
	AccessibleRoleSwitch           AccessibleRole = 63
	AccessibleRoleTab              AccessibleRole = 64
	AccessibleRoleTable            AccessibleRole = 65
	AccessibleRoleTabList          AccessibleRole = 66
	AccessibleRoleTabPanel         AccessibleRole = 67
	AccessibleRoleTextBox          AccessibleRole = 68
	AccessibleRoleTime             AccessibleRole = 69
	AccessibleRoleTimer            AccessibleRole = 70
	AccessibleRoleToolbar          AccessibleRole = 71
	AccessibleRoleTooltip          AccessibleRole = 72
	AccessibleRoleTree             AccessibleRole = 73
	AccessibleRoleTreeGrid         AccessibleRole = 74
	AccessibleRoleTreeItem         AccessibleRole = 75
	AccessibleRoleWidget           AccessibleRole = 76
	AccessibleRoleWindow           AccessibleRole = 77
	AccessibleRoleToggleButton     AccessibleRole = 78
	AccessibleRoleApplication      AccessibleRole = 79
	AccessibleRoleParagraph        AccessibleRole = 80
	AccessibleRoleBlockQuote       AccessibleRole = 81
	AccessibleRoleArticle          AccessibleRole = 82
	AccessibleRoleComment          AccessibleRole = 83
	AccessibleRoleTerminal         AccessibleRole = 84

	AccessibleSortNone       AccessibleSort = 0
	AccessibleSortAscending  AccessibleSort = 1
	AccessibleSortDescending AccessibleSort = 2
	AccessibleSortOther      AccessibleSort = 3

	AccessibleStateBusy     AccessibleState = 0
	AccessibleStateChecked  AccessibleState = 1
	AccessibleStateDisabled AccessibleState = 2
	AccessibleStateExpanded AccessibleState = 3
	AccessibleStateHidden   AccessibleState = 4
	AccessibleStateInvalid  AccessibleState = 5
	AccessibleStatePressed  AccessibleState = 6
	AccessibleStateSelected AccessibleState = 7
	AccessibleStateVisited  AccessibleState = 8

	AccessibleTextContentChangeInsert AccessibleTextContentChange = 0
	AccessibleTextContentChangeRemove AccessibleTextContentChange = 1

	AccessibleTextGranularityCharacter AccessibleTextGranularity = 0
	AccessibleTextGranularityWord      AccessibleTextGranularity = 1
	AccessibleTextGranularitySentence  AccessibleTextGranularity = 2
	AccessibleTextGranularityLine      AccessibleTextGranularity = 3
	AccessibleTextGranularityParagraph AccessibleTextGranularity = 4

	AccessibleTristateFalse AccessibleTristate = 0
	AccessibleTristateTrue  AccessibleTristate = 1
	AccessibleTristateMixed AccessibleTristate = 2

	AlignFill           Align = 0
	AlignStart          Align = 1
	AlignEnd            Align = 2
	AlignCenter         Align = 3
	AlignBaselineFill   Align = 4
	AlignBaseline       Align = 4
	AlignBaselineCenter Align = 5

	ArrowTypeUp    ArrowType = 0
	ArrowTypeDown  ArrowType = 1
	ArrowTypeLeft  ArrowType = 2
	ArrowTypeRight ArrowType = 3
	ArrowTypeNone  ArrowType = 4

	AssistantPageTypeContent  AssistantPageType = 0
	AssistantPageTypeIntro    AssistantPageType = 1
	AssistantPageTypeConfirm  AssistantPageType = 2
	AssistantPageTypeSummary  AssistantPageType = 3
	AssistantPageTypeProgress AssistantPageType = 4
	AssistantPageTypeCustom   AssistantPageType = 5

	BaselinePositionTop    BaselinePosition = 0
	BaselinePositionCenter BaselinePosition = 1
	BaselinePositionBottom BaselinePosition = 2

	BorderStyleNone   BorderStyle = 0
	BorderStyleHidden BorderStyle = 1
	BorderStyleSolid  BorderStyle = 2
	BorderStyleInset  BorderStyle = 3
	BorderStyleOutset BorderStyle = 4
	BorderStyleDotted BorderStyle = 5
	BorderStyleDashed BorderStyle = 6
	BorderStyleDouble BorderStyle = 7
	BorderStyleGroove BorderStyle = 8
	BorderStyleRidge  BorderStyle = 9

	BuilderErrorInvalidTypeFunction  BuilderError = 0
	BuilderErrorUnhandledTag         BuilderError = 1
	BuilderErrorMissingAttribute     BuilderError = 2
	BuilderErrorInvalidAttribute     BuilderError = 3
	BuilderErrorInvalidTag           BuilderError = 4
	BuilderErrorMissingPropertyValue BuilderError = 5
	BuilderErrorInvalidValue         BuilderError = 6
	BuilderErrorVersionMismatch      BuilderError = 7
	BuilderErrorDuplicateId          BuilderError = 8
	BuilderErrorObjectTypeRefused    BuilderError = 9
	BuilderErrorTemplateMismatch     BuilderError = 10
	BuilderErrorInvalidProperty      BuilderError = 11
	BuilderErrorInvalidSignal        BuilderError = 12
	BuilderErrorInvalidId            BuilderError = 13
	BuilderErrorInvalidFunction      BuilderError = 14

	ButtonsTypeNone     ButtonsType = 0
	ButtonsTypeOk       ButtonsType = 1
	ButtonsTypeClose    ButtonsType = 2
	ButtonsTypeCancel   ButtonsType = 3
	ButtonsTypeYesNo    ButtonsType = 4
	ButtonsTypeOkCancel ButtonsType = 5

	CellRendererAccelModeGtk   CellRendererAccelMode = 0
	CellRendererAccelModeOther CellRendererAccelMode = 1

	CellRendererModeInert       CellRendererMode = 0
	CellRendererModeActivatable CellRendererMode = 1
	CellRendererModeEditable    CellRendererMode = 2

	CollationNone     Collation = 0
	CollationUnicode  Collation = 1
	CollationFilename Collation = 2

	ConstraintAttributeNone     ConstraintAttribute = 0
	ConstraintAttributeLeft     ConstraintAttribute = 1
	ConstraintAttributeRight    ConstraintAttribute = 2
	ConstraintAttributeTop      ConstraintAttribute = 3
	ConstraintAttributeBottom   ConstraintAttribute = 4
	ConstraintAttributeStart    ConstraintAttribute = 5
	ConstraintAttributeEnd      ConstraintAttribute = 6
	ConstraintAttributeWidth    ConstraintAttribute = 7
	ConstraintAttributeHeight   ConstraintAttribute = 8
	ConstraintAttributeCenterX  ConstraintAttribute = 9
	ConstraintAttributeCenterY  ConstraintAttribute = 10
	ConstraintAttributeBaseline ConstraintAttribute = 11

	ConstraintRelationLe ConstraintRelation = -1
	ConstraintRelationEq ConstraintRelation = 0
	ConstraintRelationGe ConstraintRelation = 1

	ConstraintStrengthRequired ConstraintStrength = 1001001000
	ConstraintStrengthStrong   ConstraintStrength = 1000000000
	ConstraintStrengthMedium   ConstraintStrength = 1000
	ConstraintStrengthWeak     ConstraintStrength = 1

	ConstraintVflParserErrorSymbol    ConstraintVflParserError = 0
	ConstraintVflParserErrorAttribute ConstraintVflParserError = 1
	ConstraintVflParserErrorView      ConstraintVflParserError = 2
	ConstraintVflParserErrorMetric    ConstraintVflParserError = 3
	ConstraintVflParserErrorPriority  ConstraintVflParserError = 4
	ConstraintVflParserErrorRelation  ConstraintVflParserError = 5

	ContentFitFill      ContentFit = 0
	ContentFitContain   ContentFit = 1
	ContentFitCover     ContentFit = 2
	ContentFitScaleDown ContentFit = 3

	CornerTypeTopLeft     CornerType = 0
	CornerTypeBottomLeft  CornerType = 1
	CornerTypeTopRight    CornerType = 2
	CornerTypeBottomRight CornerType = 3

	CssParserErrorFailed       CssParserError = 0
	CssParserErrorSyntax       CssParserError = 1
	CssParserErrorImport       CssParserError = 2
	CssParserErrorName         CssParserError = 3
	CssParserErrorUnknownValue CssParserError = 4

	CssParserWarningDeprecated    CssParserWarning = 0
	CssParserWarningSyntax        CssParserWarning = 1
	CssParserWarningUnimplemented CssParserWarning = 2

	DeleteTypeChars           DeleteType = 0
	DeleteTypeWordEnds        DeleteType = 1
	DeleteTypeWords           DeleteType = 2
	DeleteTypeDisplayLines    DeleteType = 3
	DeleteTypeDisplayLineEnds DeleteType = 4
	DeleteTypeParagraphEnds   DeleteType = 5
	DeleteTypeParagraphs      DeleteType = 6
	DeleteTypeWhitespace      DeleteType = 7

	DialogErrorFailed    DialogError = 0
	DialogErrorCancelled DialogError = 1
	DialogErrorDismissed DialogError = 2

	DirectionTypeTabForward  DirectionType = 0
	DirectionTypeTabBackward DirectionType = 1
	DirectionTypeUp          DirectionType = 2
	DirectionTypeDown        DirectionType = 3
	DirectionTypeLeft        DirectionType = 4
	DirectionTypeRight       DirectionType = 5

	EditablePropertiesPropText           EditableProperties = 0
	EditablePropertiesPropCursorPosition EditableProperties = 1
	EditablePropertiesPropSelectionBound EditableProperties = 2
	EditablePropertiesPropEditable       EditableProperties = 3
	EditablePropertiesPropWidthChars     EditableProperties = 4
	EditablePropertiesPropMaxWidthChars  EditableProperties = 5
	EditablePropertiesPropXalign         EditableProperties = 6
	EditablePropertiesPropEnableUndo     EditableProperties = 7
	EditablePropertiesNumProperties      EditableProperties = 8

	EntryIconPositionPrimary   EntryIconPosition = 0
	EntryIconPositionSecondary EntryIconPosition = 1

	EventSequenceStateNone    EventSequenceState = 0
	EventSequenceStateClaimed EventSequenceState = 1
	EventSequenceStateDenied  EventSequenceState = 2

	FileChooserActionOpen         FileChooserAction = 0
	FileChooserActionSave         FileChooserAction = 1
	FileChooserActionSelectFolder FileChooserAction = 2

	FileChooserErrorNonexistent        FileChooserError = 0
	FileChooserErrorBadFilename        FileChooserError = 1
	FileChooserErrorAlreadyExists      FileChooserError = 2
	FileChooserErrorIncompleteHostname FileChooserError = 3

	FilterChangeDifferent  FilterChange = 0
	FilterChangeLessStrict FilterChange = 1
	FilterChangeMoreStrict FilterChange = 2

	FilterMatchSome FilterMatch = 0
	FilterMatchNone FilterMatch = 1
	FilterMatchAll  FilterMatch = 2

	FontLevelFamily   FontLevel = 0
	FontLevelFace     FontLevel = 1
	FontLevelFont     FontLevel = 2
	FontLevelFeatures FontLevel = 3

	FontRenderingAutomatic FontRendering = 0
	FontRenderingManual    FontRendering = 1

	GraphicsOffloadEnabledEnabled  GraphicsOffloadEnabled = 0
	GraphicsOffloadEnabledDisabled GraphicsOffloadEnabled = 1

	IconSizeInherit IconSize = 0
	IconSizeNormal  IconSize = 1
	IconSizeLarge   IconSize = 2

	IconThemeErrorNotFound IconThemeError = 0
	IconThemeErrorFailed   IconThemeError = 1

	IconViewDropPositionNoDrop    IconViewDropPosition = 0
	IconViewDropPositionDropInto  IconViewDropPosition = 1
	IconViewDropPositionDropLeft  IconViewDropPosition = 2
	IconViewDropPositionDropRight IconViewDropPosition = 3
	IconViewDropPositionDropAbove IconViewDropPosition = 4
	IconViewDropPositionDropBelow IconViewDropPosition = 5

	ImageTypeEmpty     ImageType = 0
	ImageTypeIconName  ImageType = 1
	ImageTypeGicon     ImageType = 2
	ImageTypePaintable ImageType = 3

	InputPurposeFreeForm InputPurpose = 0
	InputPurposeAlpha    InputPurpose = 1
	InputPurposeDigits   InputPurpose = 2
	InputPurposeNumber   InputPurpose = 3
	InputPurposePhone    InputPurpose = 4
	InputPurposeUrl      InputPurpose = 5
	InputPurposeEmail    InputPurpose = 6
	InputPurposeName     InputPurpose = 7
	InputPurposePassword InputPurpose = 8
	InputPurposePin      InputPurpose = 9
	InputPurposeTerminal InputPurpose = 10

	InscriptionOverflowClip            InscriptionOverflow = 0
	InscriptionOverflowEllipsizeStart  InscriptionOverflow = 1
	InscriptionOverflowEllipsizeMiddle InscriptionOverflow = 2
	InscriptionOverflowEllipsizeEnd    InscriptionOverflow = 3

	JustificationLeft   Justification = 0
	JustificationRight  Justification = 1
	JustificationCenter Justification = 2
	JustificationFill   Justification = 3

	LevelBarModeContinuous LevelBarMode = 0
	LevelBarModeDiscrete   LevelBarMode = 1

	LicenseUnknown    License = 0
	LicenseCustom     License = 1
	LicenseGpl20      License = 2
	LicenseGpl30      License = 3
	LicenseLgpl21     License = 4
	LicenseLgpl30     License = 5
	LicenseBsd        License = 6
	LicenseMitX11     License = 7
	LicenseArtistic   License = 8
	LicenseGpl20Only  License = 9
	LicenseGpl30Only  License = 10
	LicenseLgpl21Only License = 11
	LicenseLgpl30Only License = 12
	LicenseAgpl30     License = 13
	LicenseAgpl30Only License = 14
	LicenseBsd3       License = 15
	LicenseApache20   License = 16
	LicenseMpl20      License = 17
	License0bsd       License = 18

	ListTabBehaviorAll  ListTabBehavior = 0
	ListTabBehaviorItem ListTabBehavior = 1
	ListTabBehaviorCell ListTabBehavior = 2

	MessageTypeInfo     MessageType = 0
	MessageTypeWarning  MessageType = 1
	MessageTypeQuestion MessageType = 2
	MessageTypeError    MessageType = 3
	MessageTypeOther    MessageType = 4

	MovementStepLogicalPositions MovementStep = 0
	MovementStepVisualPositions  MovementStep = 1
	MovementStepWords            MovementStep = 2
	MovementStepDisplayLines     MovementStep = 3
	MovementStepDisplayLineEnds  MovementStep = 4
	MovementStepParagraphs       MovementStep = 5
	MovementStepParagraphEnds    MovementStep = 6
	MovementStepPages            MovementStep = 7
	MovementStepBufferEnds       MovementStep = 8
	MovementStepHorizontalPages  MovementStep = 9

	NaturalWrapModeInherit NaturalWrapMode = 0
	NaturalWrapModeNone    NaturalWrapMode = 1
	NaturalWrapModeWord    NaturalWrapMode = 2

	NotebookTabFirst NotebookTab = 0
	NotebookTabLast  NotebookTab = 1

	NumberUpLayoutLrtb NumberUpLayout = 0
	NumberUpLayoutLrbt NumberUpLayout = 1
	NumberUpLayoutRltb NumberUpLayout = 2
	NumberUpLayoutRlbt NumberUpLayout = 3
	NumberUpLayoutTblr NumberUpLayout = 4
	NumberUpLayoutTbrl NumberUpLayout = 5
	NumberUpLayoutBtlr NumberUpLayout = 6
	NumberUpLayoutBtrl NumberUpLayout = 7

	OrderingSmaller Ordering = -1
	OrderingEqual   Ordering = 0
	OrderingLarger  Ordering = 1

	OrientationHorizontal Orientation = 0
	OrientationVertical   Orientation = 1

	OverflowVisible Overflow = 0
	OverflowHidden  Overflow = 1

	PackTypeStart PackType = 0
	PackTypeEnd   PackType = 1

	PadActionTypeButton PadActionType = 0
	PadActionTypeRing   PadActionType = 1
	PadActionTypeStrip  PadActionType = 2

	PageOrientationPortrait         PageOrientation = 0
	PageOrientationLandscape        PageOrientation = 1
	PageOrientationReversePortrait  PageOrientation = 2
	PageOrientationReverseLandscape PageOrientation = 3

	PageSetAll  PageSet = 0
	PageSetEven PageSet = 1
	PageSetOdd  PageSet = 2

	PanDirectionLeft  PanDirection = 0
	PanDirectionRight PanDirection = 1
	PanDirectionUp    PanDirection = 2
	PanDirectionDown  PanDirection = 3

	PolicyTypeAlways    PolicyType = 0
	PolicyTypeAutomatic PolicyType = 1
	PolicyTypeNever     PolicyType = 2
	PolicyTypeExternal  PolicyType = 3

	PositionTypeLeft   PositionType = 0
	PositionTypeRight  PositionType = 1
	PositionTypeTop    PositionType = 2
	PositionTypeBottom PositionType = 3

	PrintDuplexSimplex    PrintDuplex = 0
	PrintDuplexHorizontal PrintDuplex = 1
	PrintDuplexVertical   PrintDuplex = 2

	PrintErrorGeneral       PrintError = 0
	PrintErrorInternalError PrintError = 1
	PrintErrorNomem         PrintError = 2
	PrintErrorInvalidFile   PrintError = 3

	PrintOperationActionPrintDialog PrintOperationAction = 0
	PrintOperationActionPrint       PrintOperationAction = 1
	PrintOperationActionPreview     PrintOperationAction = 2
	PrintOperationActionExport      PrintOperationAction = 3

	PrintOperationResultError      PrintOperationResult = 0
	PrintOperationResultApply      PrintOperationResult = 1
	PrintOperationResultCancel     PrintOperationResult = 2
	PrintOperationResultInProgress PrintOperationResult = 3

	PrintPagesAll       PrintPages = 0
	PrintPagesCurrent   PrintPages = 1
	PrintPagesRanges    PrintPages = 2
	PrintPagesSelection PrintPages = 3

	PrintQualityLow    PrintQuality = 0
	PrintQualityNormal PrintQuality = 1
	PrintQualityHigh   PrintQuality = 2
	PrintQualityDraft  PrintQuality = 3

	PrintStatusInitial         PrintStatus = 0
	PrintStatusPreparing       PrintStatus = 1
	PrintStatusGeneratingData  PrintStatus = 2
	PrintStatusSendingData     PrintStatus = 3
	PrintStatusPending         PrintStatus = 4
	PrintStatusPendingIssue    PrintStatus = 5
	PrintStatusPrinting        PrintStatus = 6
	PrintStatusFinished        PrintStatus = 7
	PrintStatusFinishedAborted PrintStatus = 8

	PropagationLimitNone       PropagationLimit = 0
	PropagationLimitSameNative PropagationLimit = 1

	PropagationPhaseNone    PropagationPhase = 0
	PropagationPhaseCapture PropagationPhase = 1
	PropagationPhaseBubble  PropagationPhase = 2
	PropagationPhaseTarget  PropagationPhase = 3

	RecentManagerErrorNotFound        RecentManagerError = 0
	RecentManagerErrorInvalidUri      RecentManagerError = 1
	RecentManagerErrorInvalidEncoding RecentManagerError = 2
	RecentManagerErrorNotRegistered   RecentManagerError = 3
	RecentManagerErrorRead            RecentManagerError = 4
	RecentManagerErrorWrite           RecentManagerError = 5
	RecentManagerErrorUnknown         RecentManagerError = 6

	ResponseTypeNone        ResponseType = -1
	ResponseTypeReject      ResponseType = -2
	ResponseTypeAccept      ResponseType = -3
	ResponseTypeDeleteEvent ResponseType = -4
	ResponseTypeOk          ResponseType = -5
	ResponseTypeCancel      ResponseType = -6
	ResponseTypeClose       ResponseType = -7
	ResponseTypeYes         ResponseType = -8
	ResponseTypeNo          ResponseType = -9
	ResponseTypeApply       ResponseType = -10
	ResponseTypeHelp        ResponseType = -11

	RevealerTransitionTypeNone       RevealerTransitionType = 0
	RevealerTransitionTypeCrossfade  RevealerTransitionType = 1
	RevealerTransitionTypeSlideRight RevealerTransitionType = 2
	RevealerTransitionTypeSlideLeft  RevealerTransitionType = 3
	RevealerTransitionTypeSlideUp    RevealerTransitionType = 4
	RevealerTransitionTypeSlideDown  RevealerTransitionType = 5
	RevealerTransitionTypeSwingRight RevealerTransitionType = 6
	RevealerTransitionTypeSwingLeft  RevealerTransitionType = 7
	RevealerTransitionTypeSwingUp    RevealerTransitionType = 8
	RevealerTransitionTypeSwingDown  RevealerTransitionType = 9

	ScrollStepSteps           ScrollStep = 0
	ScrollStepPages           ScrollStep = 1
	ScrollStepEnds            ScrollStep = 2
	ScrollStepHorizontalSteps ScrollStep = 3
	ScrollStepHorizontalPages ScrollStep = 4
	ScrollStepHorizontalEnds  ScrollStep = 5

	ScrollTypeNone         ScrollType = 0
	ScrollTypeJump         ScrollType = 1
	ScrollTypeStepBackward ScrollType = 2
	ScrollTypeStepForward  ScrollType = 3
	ScrollTypePageBackward ScrollType = 4
	ScrollTypePageForward  ScrollType = 5
	ScrollTypeStepUp       ScrollType = 6
	ScrollTypeStepDown     ScrollType = 7
	ScrollTypePageUp       ScrollType = 8
	ScrollTypePageDown     ScrollType = 9
	ScrollTypeStepLeft     ScrollType = 10
	ScrollTypeStepRight    ScrollType = 11
	ScrollTypePageLeft     ScrollType = 12
	ScrollTypePageRight    ScrollType = 13
	ScrollTypeStart        ScrollType = 14
	ScrollTypeEnd          ScrollType = 15

	ScrollablePolicyMinimum ScrollablePolicy = 0
	ScrollablePolicyNatural ScrollablePolicy = 1

	SelectionModeNone     SelectionMode = 0
	SelectionModeSingle   SelectionMode = 1
	SelectionModeBrowse   SelectionMode = 2
	SelectionModeMultiple SelectionMode = 3

	SensitivityTypeAuto SensitivityType = 0
	SensitivityTypeOn   SensitivityType = 1
	SensitivityTypeOff  SensitivityType = 2

	ShortcutScopeLocal   ShortcutScope = 0
	ShortcutScopeManaged ShortcutScope = 1
	ShortcutScopeGlobal  ShortcutScope = 2

	ShortcutTypeAccelerator                   ShortcutType = 0
	ShortcutTypeGesturePinch                  ShortcutType = 1
	ShortcutTypeGestureStretch                ShortcutType = 2
	ShortcutTypeGestureRotateClockwise        ShortcutType = 3
	ShortcutTypeGestureRotateCounterclockwise ShortcutType = 4
	ShortcutTypeGestureTwoFingerSwipeLeft     ShortcutType = 5
	ShortcutTypeGestureTwoFingerSwipeRight    ShortcutType = 6
	ShortcutTypeGesture                       ShortcutType = 7
	ShortcutTypeGestureSwipeLeft              ShortcutType = 8
	ShortcutTypeGestureSwipeRight             ShortcutType = 9

	SizeGroupModeNone       SizeGroupMode = 0
	SizeGroupModeHorizontal SizeGroupMode = 1
	SizeGroupModeVertical   SizeGroupMode = 2
	SizeGroupModeBoth       SizeGroupMode = 3

	SizeRequestModeHeightForWidth SizeRequestMode = 0
	SizeRequestModeWidthForHeight SizeRequestMode = 1
	SizeRequestModeConstantSize   SizeRequestMode = 2

	SortTypeAscending  SortType = 0
	SortTypeDescending SortType = 1

	SorterChangeDifferent  SorterChange = 0
	SorterChangeInverted   SorterChange = 1
	SorterChangeLessStrict SorterChange = 2
	SorterChangeMoreStrict SorterChange = 3

	SorterOrderPartial SorterOrder = 0
	SorterOrderNone    SorterOrder = 1
	SorterOrderTotal   SorterOrder = 2

	SpinButtonUpdatePolicyAlways  SpinButtonUpdatePolicy = 0
	SpinButtonUpdatePolicyIfValid SpinButtonUpdatePolicy = 1

	SpinTypeStepForward  SpinType = 0
	SpinTypeStepBackward SpinType = 1
	SpinTypePageForward  SpinType = 2
	SpinTypePageBackward SpinType = 3
	SpinTypeHome         SpinType = 4
	SpinTypeEnd          SpinType = 5
	SpinTypeUserDefined  SpinType = 6

	StackTransitionTypeNone            StackTransitionType = 0
	StackTransitionTypeCrossfade       StackTransitionType = 1
	StackTransitionTypeSlideRight      StackTransitionType = 2
	StackTransitionTypeSlideLeft       StackTransitionType = 3
	StackTransitionTypeSlideUp         StackTransitionType = 4
	StackTransitionTypeSlideDown       StackTransitionType = 5
	StackTransitionTypeSlideLeftRight  StackTransitionType = 6
	StackTransitionTypeSlideUpDown     StackTransitionType = 7
	StackTransitionTypeOverUp          StackTransitionType = 8
	StackTransitionTypeOverDown        StackTransitionType = 9
	StackTransitionTypeOverLeft        StackTransitionType = 10
	StackTransitionTypeOverRight       StackTransitionType = 11
	StackTransitionTypeUnderUp         StackTransitionType = 12
	StackTransitionTypeUnderDown       StackTransitionType = 13
	StackTransitionTypeUnderLeft       StackTransitionType = 14
	StackTransitionTypeUnderRight      StackTransitionType = 15
	StackTransitionTypeOverUpDown      StackTransitionType = 16
	StackTransitionTypeOverDownUp      StackTransitionType = 17
	StackTransitionTypeOverLeftRight   StackTransitionType = 18
	StackTransitionTypeOverRightLeft   StackTransitionType = 19
	StackTransitionTypeRotateLeft      StackTransitionType = 20
	StackTransitionTypeRotateRight     StackTransitionType = 21
	StackTransitionTypeRotateLeftRight StackTransitionType = 22

	StringFilterMatchModeExact     StringFilterMatchMode = 0
	StringFilterMatchModeSubstring StringFilterMatchMode = 1
	StringFilterMatchModePrefix    StringFilterMatchMode = 2

	SymbolicColorForeground SymbolicColor = 0
	SymbolicColorError      SymbolicColor = 1
	SymbolicColorWarning    SymbolicColor = 2
	SymbolicColorSuccess    SymbolicColor = 3

	SystemSettingDpi        SystemSetting = 0
	SystemSettingFontName   SystemSetting = 1
	SystemSettingFontConfig SystemSetting = 2
	SystemSettingDisplay    SystemSetting = 3
	SystemSettingIconTheme  SystemSetting = 4

	TextDirectionNone TextDirection = 0
	TextDirectionLtr  TextDirection = 1
	TextDirectionRtl  TextDirection = 2

	TextExtendSelectionWord TextExtendSelection = 0
	TextExtendSelectionLine TextExtendSelection = 1

	TextViewLayerBelowText TextViewLayer = 0
	TextViewLayerAboveText TextViewLayer = 1

	TextWindowTypeWidget TextWindowType = 1
	TextWindowTypeText   TextWindowType = 2
	TextWindowTypeLeft   TextWindowType = 3
	TextWindowTypeRight  TextWindowType = 4
	TextWindowTypeTop    TextWindowType = 5
	TextWindowTypeBottom TextWindowType = 6

	TreeViewColumnSizingGrowOnly TreeViewColumnSizing = 0
	TreeViewColumnSizingAutosize TreeViewColumnSizing = 1
	TreeViewColumnSizingFixed    TreeViewColumnSizing = 2

	TreeViewDropPositionBefore       TreeViewDropPosition = 0
	TreeViewDropPositionAfter        TreeViewDropPosition = 1
	TreeViewDropPositionIntoOrBefore TreeViewDropPosition = 2
	TreeViewDropPositionIntoOrAfter  TreeViewDropPosition = 3

	TreeViewGridLinesNone       TreeViewGridLines = 0
	TreeViewGridLinesHorizontal TreeViewGridLines = 1
	TreeViewGridLinesVertical   TreeViewGridLines = 2
	TreeViewGridLinesBoth       TreeViewGridLines = 3

	UnitNone   Unit = 0
	UnitPoints Unit = 1
	UnitInch   Unit = 2
	UnitMm     Unit = 3

	WrapModeNone     WrapMode = 0
	WrapModeChar     WrapMode = 1
	WrapModeWord     WrapMode = 2
	WrapModeWordChar WrapMode = 3

	ApplicationInhibitFlagsLogout  ApplicationInhibitFlags = 1
	ApplicationInhibitFlagsSwitch  ApplicationInhibitFlags = 2
	ApplicationInhibitFlagsSuspend ApplicationInhibitFlags = 4
	ApplicationInhibitFlagsIdle    ApplicationInhibitFlags = 8

	BuilderClosureFlagsSwapped BuilderClosureFlags = 1

	CellRendererStateSelected    CellRendererState = 1
	CellRendererStatePrelit      CellRendererState = 2
	CellRendererStateInsensitive CellRendererState = 4
	CellRendererStateSorted      CellRendererState = 8
	CellRendererStateFocused     CellRendererState = 16
	CellRendererStateExpandable  CellRendererState = 32
	CellRendererStateExpanded    CellRendererState = 64

	DebugFlagsText           DebugFlags = 1
	DebugFlagsTree           DebugFlags = 2
	DebugFlagsKeybindings    DebugFlags = 4
	DebugFlagsModules        DebugFlags = 8
	DebugFlagsGeometry       DebugFlags = 16
	DebugFlagsIcontheme      DebugFlags = 32
	DebugFlagsPrinting       DebugFlags = 64
	DebugFlagsBuilder        DebugFlags = 128
	DebugFlagsSizeRequest    DebugFlags = 256
	DebugFlagsNoCssCache     DebugFlags = 512
	DebugFlagsInteractive    DebugFlags = 1024
	DebugFlagsActions        DebugFlags = 4096
	DebugFlagsLayout         DebugFlags = 8192
	DebugFlagsSnapshot       DebugFlags = 16384
	DebugFlagsConstraints    DebugFlags = 32768
	DebugFlagsBuilderObjects DebugFlags = 65536
	DebugFlagsA11y           DebugFlags = 131072
	DebugFlagsIconfallback   DebugFlags = 262144
	DebugFlagsInvertTextDir  DebugFlags = 524288
	DebugFlagsCss            DebugFlags = 1048576

	DialogFlagsModal             DialogFlags = 1
	DialogFlagsDestroyWithParent DialogFlags = 2
	DialogFlagsUseHeaderBar      DialogFlags = 4

	EventControllerScrollFlagsNone       EventControllerScrollFlags = 0
	EventControllerScrollFlagsVertical   EventControllerScrollFlags = 1
	EventControllerScrollFlagsHorizontal EventControllerScrollFlags = 2
	EventControllerScrollFlagsDiscrete   EventControllerScrollFlags = 4
	EventControllerScrollFlagsKinetic    EventControllerScrollFlags = 8
	EventControllerScrollFlagsBothAxes   EventControllerScrollFlags = 3

	FontChooserLevelFamily     FontChooserLevel = 0
	FontChooserLevelStyle      FontChooserLevel = 1
	FontChooserLevelSize       FontChooserLevel = 2
	FontChooserLevelVariations FontChooserLevel = 4
	FontChooserLevelFeatures   FontChooserLevel = 8

	IconLookupFlagsForceRegular  IconLookupFlags = 1
	IconLookupFlagsForceSymbolic IconLookupFlags = 2
	IconLookupFlagsPreload       IconLookupFlags = 4

	InputHintsNone               InputHints = 0
	InputHintsSpellcheck         InputHints = 1
	InputHintsNoSpellcheck       InputHints = 2
	InputHintsWordCompletion     InputHints = 4
	InputHintsLowercase          InputHints = 8
	InputHintsUppercaseChars     InputHints = 16
	InputHintsUppercaseWords     InputHints = 32
	InputHintsUppercaseSentences InputHints = 64
	InputHintsInhibitOsk         InputHints = 128
	InputHintsVerticalWriting    InputHints = 256
	InputHintsEmoji              InputHints = 512
	InputHintsNoEmoji            InputHints = 1024
	InputHintsPrivate            InputHints = 2048

	ListScrollFlagsNone   ListScrollFlags = 0
	ListScrollFlagsFocus  ListScrollFlags = 1
	ListScrollFlagsSelect ListScrollFlags = 2

	PickFlagsDefault       PickFlags = 0
	PickFlagsInsensitive   PickFlags = 1
	PickFlagsNonTargetable PickFlags = 2

	PopoverMenuFlagsSliding PopoverMenuFlags = 0
	PopoverMenuFlagsNested  PopoverMenuFlags = 1

	ShortcutActionFlagsExclusive ShortcutActionFlags = 1

	StateFlagsNormal       StateFlags = 0
	StateFlagsActive       StateFlags = 1
	StateFlagsPrelight     StateFlags = 2
	StateFlagsSelected     StateFlags = 4
	StateFlagsInsensitive  StateFlags = 8
	StateFlagsInconsistent StateFlags = 16
	StateFlagsFocused      StateFlags = 32
	StateFlagsBackdrop     StateFlags = 64
	StateFlagsDirLtr       StateFlags = 128
	StateFlagsDirRtl       StateFlags = 256
	StateFlagsLink         StateFlags = 512
	StateFlagsVisited      StateFlags = 1024
	StateFlagsChecked      StateFlags = 2048
	StateFlagsDropActive   StateFlags = 4096
	StateFlagsFocusVisible StateFlags = 8192
	StateFlagsFocusWithin  StateFlags = 16384

	StyleContextPrintFlagsNone       StyleContextPrintFlags = 0
	StyleContextPrintFlagsRecurse    StyleContextPrintFlags = 1
	StyleContextPrintFlagsShowStyle  StyleContextPrintFlags = 2
	StyleContextPrintFlagsShowChange StyleContextPrintFlags = 4

	TextBufferNotifyFlagsBeforeInsert TextBufferNotifyFlags = 1
	TextBufferNotifyFlagsAfterInsert  TextBufferNotifyFlags = 2
	TextBufferNotifyFlagsBeforeDelete TextBufferNotifyFlags = 4
	TextBufferNotifyFlagsAfterDelete  TextBufferNotifyFlags = 8

	TextSearchFlagsVisibleOnly     TextSearchFlags = 1
	TextSearchFlagsTextOnly        TextSearchFlags = 2
	TextSearchFlagsCaseInsensitive TextSearchFlags = 4

	TreeModelFlagsItersPersist TreeModelFlags = 1
	TreeModelFlagsListOnly     TreeModelFlags = 2
)
