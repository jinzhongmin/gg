package dcimplot

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
	ig "github.com/jinzhongmin/gg/dcimgui"
)

type uptr = unsafe.Pointer

// func init() { cc.Open("libdcimplot*") }

var (
	imPlot_CreateContext     = cc.DlFunc[func() *ImPlotContext, *ImPlotContext]{Name: "ImPlot_CreateContext"}
	imPlot_DestroyContext    = cc.DlFunc[func(*ImPlotContext), cc.Void]{Name: "ImPlot_DestroyContext"}
	imPlot_GetCurrentContext = cc.DlFunc[func() *ImPlotContext, *ImPlotContext]{Name: "ImPlot_GetCurrentContext"}
	imPlot_SetCurrentContext = cc.DlFunc[func(*ImPlotContext), cc.Void]{Name: "ImPlot_SetCurrentContext"}
	imPlot_SetImGuiContext   = cc.DlFunc[func(*ig.ImGuiContext), cc.Void]{Name: "ImPlot_SetImGuiContext"}

	imPlot_BeginPlot     = cc.DlFunc[func(cc.String, ig.ImVec2, ImPlotFlags) cc.Bool, cc.Bool]{Name: "ImPlot_BeginPlot"}
	imPlot_EndPlot       = cc.DlFunc[func(), cc.Void]{Name: "ImPlot_EndPlot"}
	imPlot_BeginSubplots = cc.DlFunc[func(cc.String, int32, int32, ig.ImVec2, ImPlotSubplotFlags, *float32, *float32) cc.Bool, cc.Bool]{Name: "ImPlot_BeginSubplots"}
	imPlot_EndSubplots   = cc.DlFunc[func(), cc.Void]{Name: "ImPlot_EndSubplots"}

	imPlot_SetupAxis                  = cc.DlFunc[func(ImAxis, cc.String, ImPlotAxisFlags), cc.Void]{Name: "ImPlot_SetupAxis"}
	imPlot_SetupAxisLimits            = cc.DlFunc[func(ImAxis, float64, float64, ImPlotCond), cc.Void]{Name: "ImPlot_SetupAxisLimits"}
	imPlot_SetupAxisLinks             = cc.DlFunc[func(ImAxis, *float64, *float64), cc.Void]{Name: "ImPlot_SetupAxisLinks"}
	imPlot_SetupAxisFormat            = cc.DlFunc[func(ImAxis, cc.String), cc.Void]{Name: "ImPlot_SetupAxisFormat"}
	imPlot_SetupAxisFormatCallback    = cc.DlFunc[func(ImAxis, cc.Func, uptr), cc.Void]{Name: "ImPlot_SetupAxisFormatCallback"}
	imPlot_SetupAxisTicks             = cc.DlFunc[func(ImAxis, *float64, int32, cc.String, cc.Bool), cc.Void]{Name: "ImPlot_SetupAxisTicks"}
	imPlot_SetupAxisTicksRange        = cc.DlFunc[func(ImAxis, float64, float64, int32, cc.String, cc.Bool), cc.Void]{Name: "ImPlot_SetupAxisTicksRange"}
	imPlot_SetupAxisScale             = cc.DlFunc[func(ImAxis, ImPlotScale), cc.Void]{Name: "ImPlot_SetupAxisScale"}
	imPlot_SetupAxisScaleTransform    = cc.DlFunc[func(ImAxis, cc.Func, cc.Func, uptr), cc.Void]{Name: "ImPlot_SetupAxisScaleTransform"}
	imPlot_SetupAxisLimitsConstraints = cc.DlFunc[func(ImAxis, float64, float64), cc.Void]{Name: "ImPlot_SetupAxisLimitsConstraints"}
	imPlot_SetupAxisZoomConstraints   = cc.DlFunc[func(ImAxis, float64, float64), cc.Void]{Name: "ImPlot_SetupAxisZoomConstraints"}

	imPlot_SetupAxes       = cc.DlFunc[func(cc.String, cc.String, ImPlotAxisFlags, ImPlotAxisFlags), cc.Void]{Name: "ImPlot_SetupAxes"}
	imPlot_SetupAxesLimits = cc.DlFunc[func(float64, float64, float64, float64, ImPlotCond), cc.Void]{Name: "ImPlot_SetupAxesLimits"}
	imPlot_SetupLegend     = cc.DlFunc[func(ImPlotLocation, ImPlotLegendFlags), cc.Void]{Name: "ImPlot_SetupLegend"}
	imPlot_SetupMouseText  = cc.DlFunc[func(ImPlotLocation, ImPlotMouseTextFlags), cc.Void]{Name: "ImPlot_SetupMouseText"}
	imPlot_SetupFinish     = cc.DlFunc[func(), cc.Void]{Name: "ImPlot_SetupFinish"}

	imPlot_SetNextAxisLimits = cc.DlFunc[func(ImAxis, float64, float64, ImPlotCond), cc.Void]{Name: "ImPlot_SetNextAxisLimits"}
	imPlot_SetNextAxisLinks  = cc.DlFunc[func(ImAxis, *float64, *float64), cc.Void]{Name: "ImPlot_SetNextAxisLinks"}
	imPlot_SetNextAxisToFit  = cc.DlFunc[func(ImAxis), cc.Void]{Name: "ImPlot_SetNextAxisToFit"}
	imPlot_SetNextAxesLimits = cc.DlFunc[func(float64, float64, float64, float64, ImPlotCond), cc.Void]{Name: "ImPlot_SetNextAxesLimits"}
	imPlot_SetNextAxesToFit  = cc.DlFunc[func(), cc.Void]{Name: "ImPlot_SetNextAxesToFit"}

	imPlot_PlotLine   = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, int32, float64, float64, ImPlotLineFlags, int32), cc.Void]{Name: "ImPlot_PlotLine"}
	imPlot_PlotLine2D = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, uptr, int32, ImPlotLineFlags, int32), cc.Void]{Name: "ImPlot_PlotLine2D"}
	imPlot_PlotLineG  = cc.DlFunc[func(cc.String, cc.Func, uptr, int32, ImPlotLineFlags), cc.Void]{Name: "ImPlot_PlotLineG"}

	imPlot_PlotScatter   = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, int32, float64, float64, ImPlotScatterFlags, int32), cc.Void]{Name: "ImPlot_PlotScatter"}
	imPlot_PlotScatter2D = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, uptr, int32, ImPlotScatterFlags, int32), cc.Void]{Name: "ImPlot_PlotScatter2D"}
	imPlot_PlotScatterG  = cc.DlFunc[func(cc.String, cc.Func, uptr, int32, ImPlotScatterFlags), cc.Void]{Name: "ImPlot_PlotScatterG"}

	imPlot_PlotStairs   = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, int32, float64, float64, ImPlotStairsFlags, int32), cc.Void]{Name: "ImPlot_PlotStairs"}
	imPlot_PlotStairs2D = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, uptr, int32, ImPlotStairsFlags, int32), cc.Void]{Name: "ImPlot_PlotStairs2D"}
	imPlot_PlotStairsG  = cc.DlFunc[func(cc.String, cc.Func, uptr, int32, ImPlotStairsFlags), cc.Void]{Name: "ImPlot_PlotStairsG"}

	imPlot_PlotShaded   = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, int32, float64, float64, float64, ImPlotShadedFlags, int32), cc.Void]{Name: "ImPlot_PlotShaded"}
	imPlot_PlotShaded2D = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, uptr, int32, float64, ImPlotShadedFlags, int32), cc.Void]{Name: "ImPlot_PlotShaded2D"}
	imPlot_PlotShaded2Y = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, uptr, uptr, int32, ImPlotShadedFlags, int32), cc.Void]{Name: "ImPlot_PlotShaded2Y"}
	imPlot_PlotShadedG  = cc.DlFunc[func(cc.String, cc.Func, uptr, cc.Func, uptr, int32, ImPlotShadedFlags), cc.Void]{Name: "ImPlot_PlotShadedG"}

	imPlot_PlotBars   = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, int32, float64, float64, ImPlotBarsFlags, int32), cc.Void]{Name: "ImPlot_PlotBars"}
	imPlot_PlotBars2D = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, uptr, int32, float64, ImPlotBarsFlags, int32), cc.Void]{Name: "ImPlot_PlotBars2D"}
	imPlot_PlotBarsG  = cc.DlFunc[func(cc.String, cc.Func, uptr, int32, float64, ImPlotBarsFlags), cc.Void]{Name: "ImPlot_PlotBarsG"}

	imPlot_PlotBarGroups    = cc.DlFunc[func(ig.ImGuiDataType, cc.Strings, uptr, int32, int32, float64, float64, ImPlotBarGroupsFlags), cc.Void]{Name: "ImPlot_PlotBarGroups"}
	imPlot_PlotErrorBars    = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, uptr, uptr, int32, ImPlotErrorBarsFlags, int32), cc.Void]{Name: "ImPlot_PlotErrorBars"}
	imPlot_PlotErrorBarsNeg = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, uptr, uptr, uptr, int32, ImPlotErrorBarsFlags, int32), cc.Void]{Name: "ImPlot_PlotErrorBarsNeg"}
	imPlot_PlotStems        = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, int32, float64, float64, float64, ImPlotStemsFlags, int32), cc.Void]{Name: "ImPlot_PlotStems"}
	imPlot_PlotStems2D      = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, uptr, int32, float64, ImPlotStemsFlags, int32), cc.Void]{Name: "ImPlot_PlotStems2D"}
	imPlot_PlotInfLines     = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, int32, ImPlotInfLinesFlags, int32), cc.Void]{Name: "ImPlot_PlotInfLines"}
	imPlot_PlotPieChart_Fmt = cc.DlFunc[func(ig.ImGuiDataType, cc.Strings, uptr, int32, float64, float64, float64, cc.Func, uptr, float64, ImPlotPieChartFlags), cc.Void]{Name: "ImPlot_PlotPieChart_Fmt"}
	imPlot_PlotPieChart     = cc.DlFunc[func(ig.ImGuiDataType, cc.Strings, uptr, int32, float64, float64, float64, cc.String, float64, ImPlotPieChartFlags), cc.Void]{Name: "ImPlot_PlotPieChart"}

	imPlot_PlotHeatmap     = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, int32, int32, float64, float64, cc.String, ImPlotPoint, ImPlotPoint, ImPlotHeatmapFlags), cc.Void]{Name: "ImPlot_PlotHeatmap"}
	imPlot_PlotHistogram   = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, int32, int32, float64, ImPlotRange, ImPlotHistogramFlags) float64, float64]{Name: "ImPlot_PlotHistogram"}
	imPlot_PlotHistogram2D = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, uptr, int32, int32, int32, ImPlotRect, ImPlotHistogramFlags) float64, float64]{Name: "ImPlot_PlotHistogram2D"}
	imPlot_PlotDigital     = cc.DlFunc[func(ig.ImGuiDataType, cc.String, uptr, uptr, int32, ImPlotDigitalFlags, int32), cc.Void]{Name: "ImPlot_PlotDigital"}
	imPlot_PlotDigitalG    = cc.DlFunc[func(cc.String, cc.Func, uptr, int32, ImPlotDigitalFlags), cc.Void]{Name: "ImPlot_PlotDigitalG"}
	imPlot_PlotImage       = cc.DlFunc[func(cc.String, ig.ImTextureID, ImPlotPoint, ImPlotPoint, ig.ImVec2, ig.ImVec2, ig.ImVec4, ImPlotImageFlags), cc.Void]{Name: "ImPlot_PlotImage"}
	imPlot_PlotText        = cc.DlFunc[func(cc.String, float64, float64, ig.ImVec2, ImPlotTextFlags), cc.Void]{Name: "ImPlot_PlotText"}
	imPlot_PlotDummy       = cc.DlFunc[func(cc.String, ImPlotDummyFlags), cc.Void]{Name: "ImPlot_PlotDummy"}

	imPlot_DragPoint        = cc.DlFunc[func(int32, *float64, *float64, ig.ImVec4, float32, ImPlotDragToolFlags, *cc.Bool, *cc.Bool, *cc.Bool) cc.Bool, cc.Bool]{Name: "ImPlot_DragPoint"}
	imPlot_DragLineX        = cc.DlFunc[func(int32, *float64, ig.ImVec4, float32, ImPlotDragToolFlags, *cc.Bool, *cc.Bool, *cc.Bool) cc.Bool, cc.Bool]{Name: "ImPlot_DragLineX"}
	imPlot_DragLineY        = cc.DlFunc[func(int32, *float64, ig.ImVec4, float32, ImPlotDragToolFlags, *cc.Bool, *cc.Bool, *cc.Bool) cc.Bool, cc.Bool]{Name: "ImPlot_DragLineY"}
	imPlot_DragRect         = cc.DlFunc[func(int32, *float64, *float64, *float64, *float64, ig.ImVec4, ImPlotDragToolFlags, *cc.Bool, *cc.Bool, *cc.Bool) cc.Bool, cc.Bool]{Name: "ImPlot_DragRect"}
	imPlot_Annotation_Round = cc.DlFunc[func(float64, float64, ig.ImVec4, ig.ImVec2, cc.Bool, cc.Bool), cc.Void]{Name: "ImPlot_Annotation_Round"}
	imPlot_Annotation       = cc.DlFunc[func(float64, float64, ig.ImVec4, ig.ImVec2, cc.Bool, cc.String) cc.Void, cc.Void]{Name: "ImPlot_Annotation"}
	imPlot_Annotation_Fmt   = cc.DlFunc[func(float64, float64, ig.ImVec4, ig.ImVec2, cc.Bool, cc.String, ...any), cc.Void]{Name: "ImPlot_Annotation_Fmt", Va: true}
	imPlot_TagX             = cc.DlFunc[func(float64, ig.ImVec4, cc.Bool), cc.Void]{Name: "ImPlot_TagX"}
	imPlot_TagX_Text        = cc.DlFunc[func(float64, ig.ImVec4, cc.String) cc.Void, cc.Void]{Name: "ImPlot_TagX_Fmt"}
	imPlot_TagX_Fmt         = cc.DlFunc[func(float64, ig.ImVec4, cc.String, ...any) cc.Void, cc.Void]{Name: "ImPlot_TagX_Fmt", Va: true}
	imPlot_TagY             = cc.DlFunc[func(float64, ig.ImVec4, cc.Bool), cc.Void]{Name: "ImPlot_TagY"}
	imPlot_TagY_Text        = cc.DlFunc[func(float64, ig.ImVec4, cc.String) cc.Void, cc.Void]{Name: "ImPlot_TagY_Fmt"}
	imPlot_TagY_Fmt         = cc.DlFunc[func(float64, ig.ImVec4, cc.String, ...any) cc.Void, cc.Void]{Name: "ImPlot_TagY_Fmt", Va: true}

	imPlot_SetAxis             = cc.DlFunc[func(ImAxis), cc.Void]{Name: "ImPlot_SetAxis"}
	imPlot_SetAxes             = cc.DlFunc[func(ImAxis, ImAxis), cc.Void]{Name: "ImPlot_SetAxes"}
	imPlot_PixelsToPlot        = cc.DlFunc[func(ig.ImVec2, ImAxis, ImAxis) ImPlotPoint, ImPlotPoint]{Name: "ImPlot_PixelsToPlot"}
	imPlot_PixelsToPlotXY      = cc.DlFunc[func(float32, float32, ImAxis, ImAxis) ImPlotPoint, ImPlotPoint]{Name: "ImPlot_PixelsToPlotXY"}
	imPlot_PlotToPixels        = cc.DlFunc[func(ImPlotPoint, ImAxis, ImAxis) ig.ImVec2, ig.ImVec2]{Name: "ImPlot_PlotToPixels"}
	imPlot_PlotToPixelsXY      = cc.DlFunc[func(float64, float64, ImAxis, ImAxis) ig.ImVec2, ig.ImVec2]{Name: "ImPlot_PlotToPixelsXY"}
	imPlot_GetPlotPos          = cc.DlFunc[func() ig.ImVec2, ig.ImVec2]{Name: "ImPlot_GetPlotPos"}
	imPlot_GetPlotSize         = cc.DlFunc[func() ig.ImVec2, ig.ImVec2]{Name: "ImPlot_GetPlotSize"}
	imPlot_GetPlotMousePos     = cc.DlFunc[func(ImAxis, ImAxis) ImPlotPoint, ImPlotPoint]{Name: "ImPlot_GetPlotMousePos"}
	imPlot_GetPlotLimits       = cc.DlFunc[func(ImAxis, ImAxis) ImPlotRect, ImPlotRect]{Name: "ImPlot_GetPlotLimits"}
	imPlot_IsPlotHovered       = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImPlot_IsPlotHovered"}
	imPlot_IsAxisHovered       = cc.DlFunc[func(ImAxis) cc.Bool, cc.Bool]{Name: "ImPlot_IsAxisHovered"}
	imPlot_IsSubplotsHovered   = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImPlot_IsSubplotsHovered"}
	imPlot_IsPlotSelected      = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImPlot_IsPlotSelected"}
	imPlot_GetPlotSelection    = cc.DlFunc[func(ImAxis, ImAxis) ImPlotRect, ImPlotRect]{Name: "ImPlot_GetPlotSelection"}
	imPlot_CancelPlotSelection = cc.DlFunc[func(), cc.Void]{Name: "ImPlot_CancelPlotSelection"}
	imPlot_HideNextItem        = cc.DlFunc[func(cc.Bool, ImPlotCond), cc.Void]{Name: "ImPlot_HideNextItem"}
	imPlot_BeginAlignedPlots   = cc.DlFunc[func(cc.String, cc.Bool) cc.Bool, cc.Bool]{Name: "ImPlot_BeginAlignedPlots"}
	imPlot_EndAlignedPlots     = cc.DlFunc[func(), cc.Void]{Name: "ImPlot_EndAlignedPlots"}

	imPlot_BeginLegendPopup          = cc.DlFunc[func(cc.String, ig.ImGuiMouseButton) cc.Bool, cc.Bool]{Name: "ImPlot_BeginLegendPopup"}
	imPlot_EndLegendPopup            = cc.DlFunc[func(), cc.Void]{Name: "ImPlot_EndLegendPopup"}
	imPlot_IsLegendEntryHovered      = cc.DlFunc[func(cc.String) cc.Bool, cc.Bool]{Name: "ImPlot_IsLegendEntryHovered"}
	imPlot_BeginDragDropTargetPlot   = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImPlot_BeginDragDropTargetPlot"}
	imPlot_BeginDragDropTargetAxis   = cc.DlFunc[func(ImAxis) cc.Bool, cc.Bool]{Name: "ImPlot_BeginDragDropTargetAxis"}
	imPlot_BeginDragDropTargetLegend = cc.DlFunc[func() cc.Bool, cc.Bool]{Name: "ImPlot_BeginDragDropTargetLegend"}
	imPlot_EndDragDropTarget         = cc.DlFunc[func(), cc.Void]{Name: "ImPlot_EndDragDropTarget"}
	imPlot_BeginDragDropSourcePlot   = cc.DlFunc[func(ig.ImGuiDragDropFlags) cc.Bool, cc.Bool]{Name: "ImPlot_BeginDragDropSourcePlot"}
	imPlot_BeginDragDropSourceAxis   = cc.DlFunc[func(ImAxis, ig.ImGuiDragDropFlags) cc.Bool, cc.Bool]{Name: "ImPlot_BeginDragDropSourceAxis"}
	imPlot_BeginDragDropSourceItem   = cc.DlFunc[func(cc.String, ig.ImGuiDragDropFlags) cc.Bool, cc.Bool]{Name: "ImPlot_BeginDragDropSourceItem"}
	imPlot_EndDragDropSource         = cc.DlFunc[func(), cc.Void]{Name: "ImPlot_EndDragDropSource"}

	imPlot_GetStyle              = cc.DlFunc[func() *ImPlotStyle, *ImPlotStyle]{Name: "ImPlot_GetStyle"}
	imPlot_StyleColorsAuto       = cc.DlFunc[func(*ImPlotStyle), cc.Void]{Name: "ImPlot_StyleColorsAuto"}
	imPlot_StyleColorsClassic    = cc.DlFunc[func(*ImPlotStyle), cc.Void]{Name: "ImPlot_StyleColorsClassic"}
	imPlot_StyleColorsDark       = cc.DlFunc[func(*ImPlotStyle), cc.Void]{Name: "ImPlot_StyleColorsDark"}
	imPlot_StyleColorsLight      = cc.DlFunc[func(*ImPlotStyle), cc.Void]{Name: "ImPlot_StyleColorsLight"}
	imPlot_PushStyleColor_ImU32  = cc.DlFunc[func(ImPlotCol, ig.ImU32), cc.Void]{Name: "ImPlot_PushStyleColor_ImU32"}
	imPlot_PushStyleColor_ImVec4 = cc.DlFunc[func(ImPlotCol, ig.ImVec4), cc.Void]{Name: "ImPlot_PushStyleColor_ImVec4"}
	imPlot_PopStyleColor         = cc.DlFunc[func(int32), cc.Void]{Name: "ImPlot_PopStyleColor"}
	imPlot_PushStyleVar_Float    = cc.DlFunc[func(ImPlotStyleVar, float32), cc.Void]{Name: "ImPlot_PushStyleVar_Float"}
	imPlot_PushStyleVar_Int      = cc.DlFunc[func(ImPlotStyleVar, int32), cc.Void]{Name: "ImPlot_PushStyleVar_Int"}
	imPlot_PushStyleVar_ImVec2   = cc.DlFunc[func(ImPlotStyleVar, ig.ImVec2), cc.Void]{Name: "ImPlot_PushStyleVar_ImVec2"}
	imPlot_PopStyleVar           = cc.DlFunc[func(int32), cc.Void]{Name: "ImPlot_PopStyleVar"}
	imPlot_SetNextLineStyle      = cc.DlFunc[func(ig.ImVec4, float32), cc.Void]{Name: "ImPlot_SetNextLineStyle"}
	imPlot_SetNextFillStyle      = cc.DlFunc[func(ig.ImVec4, float32), cc.Void]{Name: "ImPlot_SetNextFillStyle"}
	imPlot_SetNextMarkerStyle    = cc.DlFunc[func(ImPlotMarker, float32, ig.ImVec4, float32, ig.ImVec4), cc.Void]{Name: "ImPlot_SetNextMarkerStyle"}
	imPlot_SetNextErrorBarStyle  = cc.DlFunc[func(ig.ImVec4, float32, float32), cc.Void]{Name: "ImPlot_SetNextErrorBarStyle"}
	imPlot_GetLastItemColor      = cc.DlFunc[func() ig.ImVec4, ig.ImVec4]{Name: "ImPlot_GetLastItemColor"}
	imPlot_GetStyleColorName     = cc.DlFunc[func(ImPlotCol) cc.String, cc.String]{Name: "ImPlot_GetStyleColorName"}
	imPlot_GetMarkerName         = cc.DlFunc[func(ImPlotMarker) cc.String, cc.String]{Name: "ImPlot_GetMarkerName"}

	imPlot_AddColormap_ImVec4 = cc.DlFunc[func(cc.String, *ig.ImVec4, int32, cc.Bool) ImPlotColormap, ImPlotColormap]{Name: "ImPlot_AddColormap_ImVec4"}
	imPlot_AddColormap_ImU32  = cc.DlFunc[func(cc.String, *ig.ImU32, int32, cc.Bool) ImPlotColormap, ImPlotColormap]{Name: "ImPlot_AddColormap_ImU32"}
	imPlot_GetColormapCount   = cc.DlFunc[func() int32, int32]{Name: "ImPlot_GetColormapCount"}
	imPlot_GetColormapName    = cc.DlFunc[func(ImPlotColormap) cc.String, cc.String]{Name: "ImPlot_GetColormapName"}
	imPlot_GetColormapIndex   = cc.DlFunc[func(cc.String) ImPlotColormap, ImPlotColormap]{Name: "ImPlot_GetColormapIndex"}
	imPlot_PushColormap_Index = cc.DlFunc[func(ImPlotColormap), cc.Void]{Name: "ImPlot_PushColormap_Index"}
	imPlot_PushColormap_Name  = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImPlot_PushColormap_Name"}
	imPlot_PopColormap        = cc.DlFunc[func(int32), cc.Void]{Name: "ImPlot_PopColormap"}
	imPlot_NextColormapColor  = cc.DlFunc[func() ig.ImVec4, ig.ImVec4]{Name: "ImPlot_NextColormapColor"}
	imPlot_GetColormapSize    = cc.DlFunc[func(ImPlotColormap) int32, int32]{Name: "ImPlot_GetColormapSize"}
	imPlot_GetColormapColor   = cc.DlFunc[func(int32, ImPlotColormap) ig.ImVec4, ig.ImVec4]{Name: "ImPlot_GetColormapColor"}
	imPlot_SampleColormap     = cc.DlFunc[func(float32, ImPlotColormap) ig.ImVec4, ig.ImVec4]{Name: "ImPlot_SampleColormap"}
	imPlot_ColormapScale      = cc.DlFunc[func(cc.String, float64, float64, ig.ImVec2, cc.String, ImPlotColormapScaleFlags, ImPlotColormap), cc.Void]{Name: "ImPlot_ColormapScale"}
	imPlot_ColormapSlider     = cc.DlFunc[func(cc.String, *float32, *ig.ImVec4, cc.String, ImPlotColormap) cc.Bool, cc.Bool]{Name: "ImPlot_ColormapSlider"}
	imPlot_ColormapButton     = cc.DlFunc[func(cc.String, ig.ImVec2, ImPlotColormap) cc.Bool, cc.Bool]{Name: "ImPlot_ColormapButton"}
	imPlot_BustColorCache     = cc.DlFunc[func(cc.String), cc.Void]{Name: "ImPlot_BustColorCache"}

	imPlot_GetInputMap          = cc.DlFunc[func() *ImPlotInputMap, *ImPlotInputMap]{Name: "ImPlot_GetInputMap"}
	imPlot_MapInputDefault      = cc.DlFunc[func(*ImPlotInputMap), cc.Void]{Name: "ImPlot_MapInputDefault"}
	imPlot_MapInputReverse      = cc.DlFunc[func(*ImPlotInputMap), cc.Void]{Name: "ImPlot_MapInputReverse"}
	imPlot_ItemIcon_ImVec4      = cc.DlFunc[func(ig.ImVec4), cc.Void]{Name: "ImPlot_ItemIcon_ImVec4"}
	imPlot_ItemIcon_ImU32       = cc.DlFunc[func(ig.ImU32), cc.Void]{Name: "ImPlot_ItemIcon_ImU32"}
	imPlot_ColormapIcon         = cc.DlFunc[func(ImPlotColormap), cc.Void]{Name: "ImPlot_ColormapIcon"}
	imPlot_GetPlotDrawList      = cc.DlFunc[func() *ig.ImDrawList, *ig.ImDrawList]{Name: "ImPlot_GetPlotDrawList"}
	imPlot_PushPlotClipRect     = cc.DlFunc[func(float32), cc.Void]{Name: "ImPlot_PushPlotClipRect"}
	imPlot_PopPlotClipRect      = cc.DlFunc[func(), cc.Void]{Name: "ImPlot_PopPlotClipRect"}
	imPlot_ShowStyleSelector    = cc.DlFunc[func(cc.String) cc.Bool, cc.Bool]{Name: "ImPlot_ShowStyleSelector"}
	imPlot_ShowColormapSelector = cc.DlFunc[func(cc.String) cc.Bool, cc.Bool]{Name: "ImPlot_ShowColormapSelector"}
	imPlot_ShowInputMapSelector = cc.DlFunc[func(cc.String) cc.Bool, cc.Bool]{Name: "ImPlot_ShowInputMapSelector"}
	imPlot_ShowStyleEditor      = cc.DlFunc[func(*ImPlotStyle), cc.Void]{Name: "ImPlot_ShowStyleEditor"}
	imPlot_ShowUserGuide        = cc.DlFunc[func(), cc.Void]{Name: "ImPlot_ShowUserGuide"}
	imPlot_ShowMetricsWindow    = cc.DlFunc[func(*cc.Bool), cc.Void]{Name: "ImPlot_ShowMetricsWindow"}
	// imPlot_ShowDemoWindow       = cc.DlFunc[func(*cc.Bool), cc.Void]{Name: "ImPlot_ShowDemoWindow"}
)
