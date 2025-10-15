package gsk

type BlendMode int32
type Corner int32
type FillRule int32
type GLUniformType int32
type LineCap int32
type LineJoin int32
type MaskMode int32
type PathDirection int32
type PathOperation int32
type RenderNodeType int32
type ScalingFilter int32
type SerializationError int32
type TransformCategory int32
type PathForeachFlags int32

const (
	BlendModeDefault    BlendMode = 0
	BlendModeMultiply   BlendMode = 1
	BlendModeScreen     BlendMode = 2
	BlendModeOverlay    BlendMode = 3
	BlendModeDarken     BlendMode = 4
	BlendModeLighten    BlendMode = 5
	BlendModeColorDodge BlendMode = 6
	BlendModeColorBurn  BlendMode = 7
	BlendModeHardLight  BlendMode = 8
	BlendModeSoftLight  BlendMode = 9
	BlendModeDifference BlendMode = 10
	BlendModeExclusion  BlendMode = 11
	BlendModeColor      BlendMode = 12
	BlendModeHue        BlendMode = 13
	BlendModeSaturation BlendMode = 14
	BlendModeLuminosity BlendMode = 15

	CornerTopLeft     Corner = 0
	CornerTopRight    Corner = 1
	CornerBottomRight Corner = 2
	CornerBottomLeft  Corner = 3

	FillRuleWinding FillRule = 0
	FillRuleEvenOdd FillRule = 1

	GLUniformTypeNone  GLUniformType = 0
	GLUniformTypeFloat GLUniformType = 1
	GLUniformTypeInt   GLUniformType = 2
	GLUniformTypeUint  GLUniformType = 3
	GLUniformTypeBool  GLUniformType = 4
	GLUniformTypeVec2  GLUniformType = 5
	GLUniformTypeVec3  GLUniformType = 6
	GLUniformTypeVec4  GLUniformType = 7

	LineCapButt   LineCap = 0
	LineCapRound  LineCap = 1
	LineCapSquare LineCap = 2

	LineJoinMiter LineJoin = 0
	LineJoinRound LineJoin = 1
	LineJoinBevel LineJoin = 2

	MaskModeAlpha             MaskMode = 0
	MaskModeInvertedAlpha     MaskMode = 1
	MaskModeLuminance         MaskMode = 2
	MaskModeInvertedLuminance MaskMode = 3

	PathDirectionFromStart PathDirection = 0
	PathDirectionToStart   PathDirection = 1
	PathDirectionToEnd     PathDirection = 2
	PathDirectionFromEnd   PathDirection = 3

	PathOperationMove  PathOperation = 0
	PathOperationClose PathOperation = 1
	PathOperationLine  PathOperation = 2
	PathOperationQuad  PathOperation = 3
	PathOperationCubic PathOperation = 4
	PathOperationConic PathOperation = 5

	RenderNodeTypeNotARenderNode              RenderNodeType = 0
	RenderNodeTypeContainerNode               RenderNodeType = 1
	RenderNodeTypeCairoNode                   RenderNodeType = 2
	RenderNodeTypeColorNode                   RenderNodeType = 3
	RenderNodeTypeLinearGradientNode          RenderNodeType = 4
	RenderNodeTypeRepeatingLinearGradientNode RenderNodeType = 5
	RenderNodeTypeRadialGradientNode          RenderNodeType = 6
	RenderNodeTypeRepeatingRadialGradientNode RenderNodeType = 7
	RenderNodeTypeConicGradientNode           RenderNodeType = 8
	RenderNodeTypeBorderNode                  RenderNodeType = 9
	RenderNodeTypeTextureNode                 RenderNodeType = 10
	RenderNodeTypeInsetShadowNode             RenderNodeType = 11
	RenderNodeTypeOutsetShadowNode            RenderNodeType = 12
	RenderNodeTypeTransformNode               RenderNodeType = 13
	RenderNodeTypeOpacityNode                 RenderNodeType = 14
	RenderNodeTypeColorMatrixNode             RenderNodeType = 15
	RenderNodeTypeRepeatNode                  RenderNodeType = 16
	RenderNodeTypeClipNode                    RenderNodeType = 17
	RenderNodeTypeRoundedClipNode             RenderNodeType = 18
	RenderNodeTypeShadowNode                  RenderNodeType = 19
	RenderNodeTypeBlendNode                   RenderNodeType = 20
	RenderNodeTypeCrossFadeNode               RenderNodeType = 21
	RenderNodeTypeTextNode                    RenderNodeType = 22
	RenderNodeTypeBlurNode                    RenderNodeType = 23
	RenderNodeTypeDebugNode                   RenderNodeType = 24
	RenderNodeTypeGlShaderNode                RenderNodeType = 25
	RenderNodeTypeTextureScaleNode            RenderNodeType = 26
	RenderNodeTypeMaskNode                    RenderNodeType = 27
	RenderNodeTypeFillNode                    RenderNodeType = 28
	RenderNodeTypeStrokeNode                  RenderNodeType = 29
	RenderNodeTypeSubsurfaceNode              RenderNodeType = 30

	ScalingFilterLinear    ScalingFilter = 0
	ScalingFilterNearest   ScalingFilter = 1
	ScalingFilterTrilinear ScalingFilter = 2

	SerializationErrorUnsupportedFormat  SerializationError = 0
	SerializationErrorUnsupportedVersion SerializationError = 1
	SerializationErrorInvalidData        SerializationError = 2

	TransformCategoryUnknown     TransformCategory = 0
	TransformCategoryAny         TransformCategory = 1
	TransformCategory3d          TransformCategory = 2
	TransformCategory2d          TransformCategory = 3
	TransformCategory2dAffine    TransformCategory = 4
	TransformCategory2dTranslate TransformCategory = 5
	TransformCategoryIdentity    TransformCategory = 6

	PathForeachFlagsOnlyLines PathForeachFlags = 0
	PathForeachFlagsQuad      PathForeachFlags = 1
	PathForeachFlagsCubic     PathForeachFlags = 2
	PathForeachFlagsConic     PathForeachFlags = 4
)
