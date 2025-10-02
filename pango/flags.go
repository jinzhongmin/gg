package pango

type Alignment int32
type AttrType int32
type BaselineShift int32
type BidiType int32
type CoverageLevel int32
type Direction int32
type EllipsizeMode int32
type FontScale int32
type Gravity int32
type GravityHint int32
type LayoutDeserializeError int32
type Overline int32
type RenderPart int32
type Script int32
type Stretch int32
type Style int32
type TabAlign int32
type TextTransform int32
type Underline int32
type Variant int32
type Weight int32
type WrapMode int32
type FontMask int32
type LayoutDeserializeFlags int32
type LayoutSerializeFlags int32
type ShapeFlags int32
type ShowFlags int32

const (
	AlignmentLeft   Alignment = 0
	AlignmentCenter Alignment = 1
	AlignmentRight  Alignment = 2

	AttrTypeInvalid            AttrType = 0
	AttrTypeLanguage           AttrType = 1
	AttrTypeFamily             AttrType = 2
	AttrTypeStyle              AttrType = 3
	AttrTypeWeight             AttrType = 4
	AttrTypeVariant            AttrType = 5
	AttrTypeStretch            AttrType = 6
	AttrTypeSize               AttrType = 7
	AttrTypeFontDesc           AttrType = 8
	AttrTypeForeground         AttrType = 9
	AttrTypeBackground         AttrType = 10
	AttrTypeUnderline          AttrType = 11
	AttrTypeStrikethrough      AttrType = 12
	AttrTypeRise               AttrType = 13
	AttrTypeShape              AttrType = 14
	AttrTypeScale              AttrType = 15
	AttrTypeFallback           AttrType = 16
	AttrTypeLetterSpacing      AttrType = 17
	AttrTypeUnderlineColor     AttrType = 18
	AttrTypeStrikethroughColor AttrType = 19
	AttrTypeAbsoluteSize       AttrType = 20
	AttrTypeGravity            AttrType = 21
	AttrTypeGravityHint        AttrType = 22
	AttrTypeFontFeatures       AttrType = 23
	AttrTypeForegroundAlpha    AttrType = 24
	AttrTypeBackgroundAlpha    AttrType = 25
	AttrTypeAllowBreaks        AttrType = 26
	AttrTypeShow               AttrType = 27
	AttrTypeInsertHyphens      AttrType = 28
	AttrTypeOverline           AttrType = 29
	AttrTypeOverlineColor      AttrType = 30
	AttrTypeLineHeight         AttrType = 31
	AttrTypeAbsoluteLineHeight AttrType = 32
	AttrTypeTextTransform      AttrType = 33
	AttrTypeWord               AttrType = 34
	AttrTypeSentence           AttrType = 35
	AttrTypeBaselineShift      AttrType = 36
	AttrTypeFontScale          AttrType = 37

	BaselineShiftNone        BaselineShift = 0
	BaselineShiftSuperscript BaselineShift = 1
	BaselineShiftSubscript   BaselineShift = 2

	BidiTypeL   BidiType = 0
	BidiTypeLre BidiType = 1
	BidiTypeLro BidiType = 2
	BidiTypeR   BidiType = 3
	BidiTypeAl  BidiType = 4
	BidiTypeRle BidiType = 5
	BidiTypeRlo BidiType = 6
	BidiTypePdf BidiType = 7
	BidiTypeEn  BidiType = 8
	BidiTypeEs  BidiType = 9
	BidiTypeEt  BidiType = 10
	BidiTypeAn  BidiType = 11
	BidiTypeCs  BidiType = 12
	BidiTypeNsm BidiType = 13
	BidiTypeBn  BidiType = 14
	BidiTypeB   BidiType = 15
	BidiTypeS   BidiType = 16
	BidiTypeWs  BidiType = 17
	BidiTypeOn  BidiType = 18
	BidiTypeLri BidiType = 19
	BidiTypeRli BidiType = 20
	BidiTypeFsi BidiType = 21
	BidiTypePdi BidiType = 22

	CoverageLevelNone        CoverageLevel = 0
	CoverageLevelFallback    CoverageLevel = 1
	CoverageLevelApproximate CoverageLevel = 2
	CoverageLevelExact       CoverageLevel = 3

	DirectionLtr     Direction = 0
	DirectionRtl     Direction = 1
	DirectionTtbLtr  Direction = 2
	DirectionTtbRtl  Direction = 3
	DirectionWeakLtr Direction = 4
	DirectionWeakRtl Direction = 5
	DirectionNeutral Direction = 6

	EllipsizeModeNone   EllipsizeMode = 0
	EllipsizeModeStart  EllipsizeMode = 1
	EllipsizeModeMiddle EllipsizeMode = 2
	EllipsizeModeEnd    EllipsizeMode = 3

	FontScaleNone        FontScale = 0
	FontScaleSuperscript FontScale = 1
	FontScaleSubscript   FontScale = 2
	FontScaleSmallCaps   FontScale = 3

	GravitySouth Gravity = 0
	GravityEast  Gravity = 1
	GravityNorth Gravity = 2
	GravityWest  Gravity = 3
	GravityAuto  Gravity = 4

	GravityHintNatural GravityHint = 0
	GravityHintStrong  GravityHint = 1
	GravityHintLine    GravityHint = 2

	LayoutDeserializeErrorInvalid      LayoutDeserializeError = 0
	LayoutDeserializeErrorInvalidValue LayoutDeserializeError = 1
	LayoutDeserializeErrorMissingValue LayoutDeserializeError = 2

	OverlineNone   Overline = 0
	OverlineSingle Overline = 1

	RenderPartForeground    RenderPart = 0
	RenderPartBackground    RenderPart = 1
	RenderPartUnderline     RenderPart = 2
	RenderPartStrikethrough RenderPart = 3
	RenderPartOverline      RenderPart = 4

	ScriptInvalidCode          Script = -1
	ScriptCommon               Script = 0
	ScriptInherited            Script = 1
	ScriptArabic               Script = 2
	ScriptArmenian             Script = 3
	ScriptBengali              Script = 4
	ScriptBopomofo             Script = 5
	ScriptCherokee             Script = 6
	ScriptCoptic               Script = 7
	ScriptCyrillic             Script = 8
	ScriptDeseret              Script = 9
	ScriptDevanagari           Script = 10
	ScriptEthiopic             Script = 11
	ScriptGeorgian             Script = 12
	ScriptGothic               Script = 13
	ScriptGreek                Script = 14
	ScriptGujarati             Script = 15
	ScriptGurmukhi             Script = 16
	ScriptHan                  Script = 17
	ScriptHangul               Script = 18
	ScriptHebrew               Script = 19
	ScriptHiragana             Script = 20
	ScriptKannada              Script = 21
	ScriptKatakana             Script = 22
	ScriptKhmer                Script = 23
	ScriptLao                  Script = 24
	ScriptLatin                Script = 25
	ScriptMalayalam            Script = 26
	ScriptMongolian            Script = 27
	ScriptMyanmar              Script = 28
	ScriptOgham                Script = 29
	ScriptOldItalic            Script = 30
	ScriptOriya                Script = 31
	ScriptRunic                Script = 32
	ScriptSinhala              Script = 33
	ScriptSyriac               Script = 34
	ScriptTamil                Script = 35
	ScriptTelugu               Script = 36
	ScriptThaana               Script = 37
	ScriptThai                 Script = 38
	ScriptTibetan              Script = 39
	ScriptCanadianAboriginal   Script = 40
	ScriptYi                   Script = 41
	ScriptTagalog              Script = 42
	ScriptHanunoo              Script = 43
	ScriptBuhid                Script = 44
	ScriptTagbanwa             Script = 45
	ScriptBraille              Script = 46
	ScriptCypriot              Script = 47
	ScriptLimbu                Script = 48
	ScriptOsmanya              Script = 49
	ScriptShavian              Script = 50
	ScriptLinearB              Script = 51
	ScriptTaiLe                Script = 52
	ScriptUgaritic             Script = 53
	ScriptNewTaiLue            Script = 54
	ScriptBuginese             Script = 55
	ScriptGlagolitic           Script = 56
	ScriptTifinagh             Script = 57
	ScriptSylotiNagri          Script = 58
	ScriptOldPersian           Script = 59
	ScriptKharoshthi           Script = 60
	ScriptUnknown              Script = 61
	ScriptBalinese             Script = 62
	ScriptCuneiform            Script = 63
	ScriptPhoenician           Script = 64
	ScriptPhagsPa              Script = 65
	ScriptNko                  Script = 66
	ScriptKayahLi              Script = 67
	ScriptLepcha               Script = 68
	ScriptRejang               Script = 69
	ScriptSundanese            Script = 70
	ScriptSaurashtra           Script = 71
	ScriptCham                 Script = 72
	ScriptOlChiki              Script = 73
	ScriptVai                  Script = 74
	ScriptCarian               Script = 75
	ScriptLycian               Script = 76
	ScriptLydian               Script = 77
	ScriptBatak                Script = 78
	ScriptBrahmi               Script = 79
	ScriptMandaic              Script = 80
	ScriptChakma               Script = 81
	ScriptMeroiticCursive      Script = 82
	ScriptMeroiticHieroglyphs  Script = 83
	ScriptMiao                 Script = 84
	ScriptSharada              Script = 85
	ScriptSoraSompeng          Script = 86
	ScriptTakri                Script = 87
	ScriptBassaVah             Script = 88
	ScriptCaucasianAlbanian    Script = 89
	ScriptDuployan             Script = 90
	ScriptElbasan              Script = 91
	ScriptGrantha              Script = 92
	ScriptKhojki               Script = 93
	ScriptKhudawadi            Script = 94
	ScriptLinearA              Script = 95
	ScriptMahajani             Script = 96
	ScriptManichaean           Script = 97
	ScriptMendeKikakui         Script = 98
	ScriptModi                 Script = 99
	ScriptMro                  Script = 100
	ScriptNabataean            Script = 101
	ScriptOldNorthArabian      Script = 102
	ScriptOldPermic            Script = 103
	ScriptPahawhHmong          Script = 104
	ScriptPalmyrene            Script = 105
	ScriptPauCinHau            Script = 106
	ScriptPsalterPahlavi       Script = 107
	ScriptSiddham              Script = 108
	ScriptTirhuta              Script = 109
	ScriptWarangCiti           Script = 110
	ScriptAhom                 Script = 111
	ScriptAnatolianHieroglyphs Script = 112
	ScriptHatran               Script = 113
	ScriptMultani              Script = 114
	ScriptOldHungarian         Script = 115
	ScriptSignwriting          Script = 116

	StretchUltraCondensed Stretch = 0
	StretchExtraCondensed Stretch = 1
	StretchCondensed      Stretch = 2
	StretchSemiCondensed  Stretch = 3
	StretchNormal         Stretch = 4
	StretchSemiExpanded   Stretch = 5
	StretchExpanded       Stretch = 6
	StretchExtraExpanded  Stretch = 7
	StretchUltraExpanded  Stretch = 8

	StyleNormal  Style = 0
	StyleOblique Style = 1
	StyleItalic  Style = 2

	TabAlignLeft    TabAlign = 0
	TabAlignRight   TabAlign = 1
	TabAlignCenter  TabAlign = 2
	TabAlignDecimal TabAlign = 3

	TextTransformNone       TextTransform = 0
	TextTransformLowercase  TextTransform = 1
	TextTransformUppercase  TextTransform = 2
	TextTransformCapitalize TextTransform = 3

	UnderlineNone       Underline = 0
	UnderlineSingle     Underline = 1
	UnderlineDouble     Underline = 2
	UnderlineLow        Underline = 3
	UnderlineError      Underline = 4
	UnderlineSingleLine Underline = 5
	UnderlineDoubleLine Underline = 6
	UnderlineErrorLine  Underline = 7

	VariantNormal        Variant = 0
	VariantSmallCaps     Variant = 1
	VariantAllSmallCaps  Variant = 2
	VariantPetiteCaps    Variant = 3
	VariantAllPetiteCaps Variant = 4
	VariantUnicase       Variant = 5
	VariantTitleCaps     Variant = 6

	WeightThin       Weight = 100
	WeightUltralight Weight = 200
	WeightLight      Weight = 300
	WeightSemilight  Weight = 350
	WeightBook       Weight = 380
	WeightNormal     Weight = 400
	WeightMedium     Weight = 500
	WeightSemibold   Weight = 600
	WeightBold       Weight = 700
	WeightUltrabold  Weight = 800
	WeightHeavy      Weight = 900
	WeightUltraheavy Weight = 1000

	WrapModeWord     WrapMode = 0
	WrapModeChar     WrapMode = 1
	WrapModeWordChar WrapMode = 2
	WrapModeNone     WrapMode = 3

	FontMaskFamily     FontMask = 1
	FontMaskStyle      FontMask = 2
	FontMaskVariant    FontMask = 4
	FontMaskWeight     FontMask = 8
	FontMaskStretch    FontMask = 16
	FontMaskSize       FontMask = 32
	FontMaskGravity    FontMask = 64
	FontMaskVariations FontMask = 128
	FontMaskFeatures   FontMask = 256

	LayoutDeserializeFlagsDefault LayoutDeserializeFlags = 0
	LayoutDeserializeFlagsContext LayoutDeserializeFlags = 1

	LayoutSerializeFlagsDefault LayoutSerializeFlags = 0
	LayoutSerializeFlagsContext LayoutSerializeFlags = 1
	LayoutSerializeFlagsOutput  LayoutSerializeFlags = 2

	ShapeFlagsNone           ShapeFlags = 0
	ShapeFlagsRoundPositions ShapeFlags = 1

	ShowFlagsNone       ShowFlags = 0
	ShowFlagsSpaces     ShowFlags = 1
	ShowFlagsLineBreaks ShowFlags = 2
	ShowFlagsIgnorables ShowFlags = 4
)
