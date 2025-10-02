package cairo

type Status int32
type Content int32
type Operator int32
type Antialias int32
type FillRule int32
type LineCap int32
type LineJoin int32
type TextClusterFlags int32
type FontSlant int32
type FontWeight int32
type SubpixelOrder int32
type HintStyle int32
type HintMetrics int32
type FontType int32
type PathDataType int32
type DeviceType int32
type SurfaceType int32
type Format int32
type PatternType int32
type Extend int32
type Filter int32
type RegionOverlap int32
type Dither int32
type ColorMode int32
type SurfaceObserverMode int32
type PDFVersion int32
type PDFOutlineFlags int32
type PDFMetadata int32
type PSLevel int32
type ScriptMode int32
type SVGVersion int32
type SVGUnit int32

const (
	StatusSuccess                 Status = 0
	StatusNoMemory                Status = 1
	StatusInvalidRestore          Status = 2
	StatusInvalidPopGroup         Status = 3
	StatusNoCurrentPoint          Status = 4
	StatusInvalidMatrix           Status = 5
	StatusInvalidStatus           Status = 6
	StatusNullPointer             Status = 7
	StatusInvalidString           Status = 8
	StatusInvalidPathData         Status = 9
	StatusReadError               Status = 10
	StatusWriteError              Status = 11
	StatusSurfaceFinished         Status = 12
	StatusSurfaceTypeMismatch     Status = 13
	StatusPatternTypeMismatch     Status = 14
	StatusInvalidContent          Status = 15
	StatusInvalidFormat           Status = 16
	StatusInvalidVisual           Status = 17
	StatusFileNotFound            Status = 18
	StatusInvalidDash             Status = 19
	StatusInvalidDscComment       Status = 20
	StatusInvalidIndex            Status = 21
	StatusClipNotRepresentable    Status = 22
	StatusTempFileError           Status = 23
	StatusInvalidStride           Status = 24
	StatusFontTypeMismatch        Status = 25
	StatusUserFontImmutable       Status = 26
	StatusUserFontError           Status = 27
	StatusNegativeCount           Status = 28
	StatusInvalidClusters         Status = 29
	StatusInvalidSlant            Status = 30
	StatusInvalidWeight           Status = 31
	StatusInvalidSize             Status = 32
	StatusUserFontNotImplemented  Status = 33
	StatusDeviceTypeMismatch      Status = 34
	StatusDeviceError             Status = 35
	StatusInvalidMeshConstruction Status = 36
	StatusDeviceFinished          Status = 37
	StatusJbig2GlobalMissing      Status = 38

	ContentColor      Content = 4096
	ContentAlpha      Content = 8192
	ContentColorAlpha Content = 12288

	OperatorClear         Operator = 0
	OperatorSource        Operator = 1
	OperatorOver          Operator = 2
	OperatorIn            Operator = 3
	OperatorOut           Operator = 4
	OperatorAtop          Operator = 5
	OperatorDest          Operator = 6
	OperatorDestOver      Operator = 7
	OperatorDestIn        Operator = 8
	OperatorDestOut       Operator = 9
	OperatorDestAtop      Operator = 10
	OperatorXor           Operator = 11
	OperatorAdd           Operator = 12
	OperatorSaturate      Operator = 13
	OperatorMultiply      Operator = 14
	OperatorScreen        Operator = 15
	OperatorOverlay       Operator = 16
	OperatorDarken        Operator = 17
	OperatorLighten       Operator = 18
	OperatorColorDodge    Operator = 19
	OperatorColorBurn     Operator = 20
	OperatorHardLight     Operator = 21
	OperatorSoftLight     Operator = 22
	OperatorDifference    Operator = 23
	OperatorExclusion     Operator = 24
	OperatorHslHue        Operator = 25
	OperatorHslSaturation Operator = 26
	OperatorHslColor      Operator = 27
	OperatorHslLuminosity Operator = 28

	AntialiasDefault  Antialias = 0
	AntialiasNone     Antialias = 1
	AntialiasGray     Antialias = 2
	AntialiasSubpixel Antialias = 3
	AntialiasFast     Antialias = 4
	AntialiasGood     Antialias = 5
	AntialiasBest     Antialias = 6

	FillRuleWinding FillRule = 0
	FillRuleEvenOdd FillRule = 1

	LineCapButt   LineCap = 0
	LineCapRound  LineCap = 1
	LineCapSquare LineCap = 2

	LineJoinMiter LineJoin = 0
	LineJoinRound LineJoin = 1
	LineJoinBevel LineJoin = 2

	TextClusterFlagsBackward TextClusterFlags = 1

	FontSlantNormal  FontSlant = 0
	FontSlantItalic  FontSlant = 1
	FontSlantOblique FontSlant = 2

	FontWeightNormal FontWeight = 0
	FontWeightBold   FontWeight = 1

	SubpixelOrderDefault SubpixelOrder = 0
	SubpixelOrderRgb     SubpixelOrder = 1
	SubpixelOrderBgr     SubpixelOrder = 2
	SubpixelOrderVrgb    SubpixelOrder = 3
	SubpixelOrderVbgr    SubpixelOrder = 4

	HintStyleDefault HintStyle = 0
	HintStyleNone    HintStyle = 1
	HintStyleSlight  HintStyle = 2
	HintStyleMedium  HintStyle = 3
	HintStyleFull    HintStyle = 4

	HintMetricsDefault HintMetrics = 0
	HintMetricsOff     HintMetrics = 1
	HintMetricsOn      HintMetrics = 2

	FontTypeToy    FontType = 0
	FontTypeFt     FontType = 1
	FontTypeWin32  FontType = 2
	FontTypeQuartz FontType = 3
	FontTypeUser   FontType = 4

	PathDataTypeMoveTo    PathDataType = 0
	PathDataTypeLineTo    PathDataType = 1
	PathDataTypeCurveTo   PathDataType = 2
	PathDataTypeClosePath PathDataType = 3

	DeviceTypeDrm     DeviceType = 0
	DeviceTypeGl      DeviceType = 1
	DeviceTypeScript  DeviceType = 2
	DeviceTypeXcb     DeviceType = 3
	DeviceTypeXlib    DeviceType = 4
	DeviceTypeXml     DeviceType = 5
	DeviceTypeCogl    DeviceType = 6
	DeviceTypeWin32   DeviceType = 7
	DeviceTypeInvalid DeviceType = -1

	SurfaceTypeImage         SurfaceType = 0
	SurfaceTypePdf           SurfaceType = 1
	SurfaceTypePs            SurfaceType = 2
	SurfaceTypeXlib          SurfaceType = 3
	SurfaceTypeXcb           SurfaceType = 4
	SurfaceTypeGlitz         SurfaceType = 5
	SurfaceTypeQuartz        SurfaceType = 6
	SurfaceTypeWin32         SurfaceType = 7
	SurfaceTypeBeos          SurfaceType = 8
	SurfaceTypeDirectfb      SurfaceType = 9
	SurfaceTypeSvg           SurfaceType = 10
	SurfaceTypeOs2           SurfaceType = 11
	SurfaceTypeWin32Printing SurfaceType = 12
	SurfaceTypeQuartzImage   SurfaceType = 13
	SurfaceTypeScript        SurfaceType = 14
	SurfaceTypeQt            SurfaceType = 15
	SurfaceTypeRecording     SurfaceType = 16
	SurfaceTypeVg            SurfaceType = 17
	SurfaceTypeGl            SurfaceType = 18
	SurfaceTypeDrm           SurfaceType = 19
	SurfaceTypeTee           SurfaceType = 20
	SurfaceTypeXml           SurfaceType = 21
	SurfaceTypeSkia          SurfaceType = 22
	SurfaceTypeSubsurface    SurfaceType = 23
	SurfaceTypeCogl          SurfaceType = 24

	FormatInvalid  Format = -1
	FormatARGB32   Format = 0
	FormatRGB24    Format = 1
	FormatA8       Format = 2
	FormatA1       Format = 3
	FormatRGB16565 Format = 4
	FormatRGB30    Format = 5

	PatternTypeSolid        PatternType = 0
	PatternTypeSurface      PatternType = 1
	PatternTypeLinear       PatternType = 2
	PatternTypeRadial       PatternType = 3
	PatternTypeMesh         PatternType = 4
	PatternTypeRasterSource PatternType = 5

	ExtendNone    Extend = 0
	ExtendRepeat  Extend = 1
	ExtendReflect Extend = 2
	ExtendPad     Extend = 3

	FilterFast     Filter = 0
	FilterGood     Filter = 1
	FilterBest     Filter = 2
	FilterNearest  Filter = 3
	FilterBilinear Filter = 4
	FilterGaussian Filter = 5

	RegionOverlapIn   RegionOverlap = 0
	RegionOverlapOut  RegionOverlap = 1
	RegionOverlapPart RegionOverlap = 2

	DitherNone    Dither = 0
	DitherDefault Dither = 1
	DitherFast    Dither = 2
	DitherGood    Dither = 3
	DitherBest    Dither = 4

	ColorModeDefault ColorMode = 0
	ColorModeNoColor ColorMode = 1
	ColorModeColor   ColorMode = 2

	SurfaceObserverModeNormal     SurfaceObserverMode = 0
	SurfaceObserverModeOperations SurfaceObserverMode = 0x1

	PDFVersion14 PDFVersion = 0
	PDFVersion15 PDFVersion = 1
	PDFVersion16 PDFVersion = 2
	PDFVersion17 PDFVersion = 3

	PDFOutlineFlagOpen   PDFOutlineFlags = 0x1
	PDFOutlineFlagBold   PDFOutlineFlags = 0x2
	PDFOutlineFlagItalic PDFOutlineFlags = 0x4

	PDFMetadataTitle      PDFMetadata = 0
	PDFMetadataAuthor     PDFMetadata = 1
	PDFMetadataSubject    PDFMetadata = 2
	PDFMetadataKeywords   PDFMetadata = 3
	PDFMetadataCreator    PDFMetadata = 4
	PDFMetadataCreateDate PDFMetadata = 5
	PDFMetadataModDate    PDFMetadata = 6

	PSLevel2 PSLevel = 0
	PSLevel3 PSLevel = 1

	ScriptModeAscii  ScriptMode = 0
	ScriptModeBinary ScriptMode = 1

	SVGVersion11   SVGVersion = 0
	SVGVersion12   SVGVersion = 1
	SVGUnitUser    SVGUnit    = 0
	SVGUnitEm      SVGUnit    = 1
	SVGUnitEx      SVGUnit    = 2
	SVGUnitPx      SVGUnit    = 3
	SVGUnitIn      SVGUnit    = 4
	SVGUnitCm      SVGUnit    = 5
	SVGUnitMm      SVGUnit    = 6
	SVGUnitPt      SVGUnit    = 7
	SVGUnitPc      SVGUnit    = 8
	SVGUnitPercent SVGUnit    = 9
)

type MimeType string

const (
	MimeTypeJPEG           MimeType = "image/jpeg"
	MimeTypePNG            MimeType = "image/png"
	MimeTypeJP2            MimeType = "image/jp2"
	MimeTypeURI            MimeType = "text/x-uri"
	MimeTypeUNIQUEID       MimeType = "application/x-cairo.uuid"
	MimeTypeJBIG2          MimeType = "application/x-cairo.jbig2"
	MimeTypeJBIG2GLOBAL    MimeType = "application/x-cairo.jbig2-global"
	MimeTypeJBIG2GLOBALID  MimeType = "application/x-cairo.jbig2-global-id"
	MimeTypeCCITTFAX       MimeType = "image/g3fax"
	MimeTypeCCITTFAXPARAMS MimeType = "application/x-cairo.ccitt.params"
	MimeTypeEPS            MimeType = "application/postscript"
	MimeTypeEPSPARAMS      MimeType = "application/x-cairo.eps.params"
)
