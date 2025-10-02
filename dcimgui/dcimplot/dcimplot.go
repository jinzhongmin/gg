package dcimplot

import (
	"github.com/jinzhongmin/gg/cc"
	ig "github.com/jinzhongmin/gg/dcimgui"
)

func pBool(b *bool) *cc.Bool {
	return (*cc.Bool)(uptr(b))
}

type ImPlotContext struct{}

type ImPlotPoint struct{ X, Y float64 }
type ImPlotRange struct{ Min, Max float64 }
type ImPlotRect struct{ X, Y ImPlotRange }

type ImPlotStyle struct {
	// item styling variables
	LineWeight       float32 // = 1,      item line weight in pixels
	Marker           int32   // = ImPlotMarker_None, marker specification
	MarkerSize       float32 // = 4,      marker size in pixels (roughly the marker's "radius")
	MarkerWeight     float32 // = 1,      outline weight of markers in pixels
	FillAlpha        float32 // = 1,      alpha modifier applied to plot fills
	ErrorBarSize     float32 // = 5,      error bar whisker width in pixels
	ErrorBarWeight   float32 // = 1.5,    error bar whisker weight in pixels
	DigitalBitHeight float32 // = 8,      digital channels bit height (at y = 1.0f) in pixels
	DigitalBitGap    float32 // = 4,      digital channels bit padding gap in pixels
	// plot styling variables
	PlotBorderSize     float32   // = 1,      line thickness of border around plot area
	MinorAlpha         float32   // = 0.25    alpha multiplier applied to minor axis grid lines
	MajorTickLen       ig.ImVec2 // = 10,10   major tick lengths for X and Y axes
	MinorTickLen       ig.ImVec2 // = 5,5     minor tick lengths for X and Y axes
	MajorTickSize      ig.ImVec2 // = 1,1     line thickness of major ticks
	MinorTickSize      ig.ImVec2 // = 1,1     line thickness of minor ticks
	MajorGridSize      ig.ImVec2 // = 1,1     line thickness of major grid lines
	MinorGridSize      ig.ImVec2 // = 1,1     line thickness of minor grid lines
	PlotPadding        ig.ImVec2 // = 10,10   padding between widget frame and plot area, labels, or outside legends (i.e. main padding)
	LabelPadding       ig.ImVec2 // = 5,5     padding between axes labels, tick labels, and plot edge
	LegendPadding      ig.ImVec2 // = 10,10   legend padding from plot edges
	LegendInnerPadding ig.ImVec2 // = 5,5     legend inner padding from legend edges
	LegendSpacing      ig.ImVec2 // = 5,0     spacing between legend entries
	MousePosPadding    ig.ImVec2 // = 10,10   padding between plot edge and int32erior mouse location text
	AnnotationPadding  ig.ImVec2 // = 2,2     text padding around annotation labels
	FitPadding         ig.ImVec2 // = 0,0     additional fit padding as a percentage of the fit extents (e.g. ImVec2(0.1f,0.1f) adds 10% to the fit extents of X and Y)
	PlotDefaultSize    ig.ImVec2 // = 400,300 default size used when ImVec2(0,0) is passed to BeginPlot
	PlotMinSize        ig.ImVec2 // = 200,150 minimum size plot frame can be when shrunk
	// style colors
	Colors [ImPlotCol_COUNT]ig.ImVec4 // Array of styling colors. Indexable with ImPlotCol_ enums.
	// colormap
	Colormap ImPlotColormap // The current colormap. Set this to either an ImPlotColormap_ enum or an index returned by AddColormap.
	// settings/flags
	UseLocalTime   bool // = false,  axis labels will be formatted for your timezone when ImPlotAxisFlag_Time is enabled
	UseISO8601     bool // = false,  dates will be formatted according to ISO 8601 where applicable (e.g. YYYY-MM-DD, YYYY-MM, --MM-DD, etc.)
	Use24HourClock bool // = false,  times will be formatted using a 24 hour clock
}

type ImPlotInputMap struct {
	Pan           ImGuiMouseButton // LMB    enables panning when held,
	PanMod        int32            // none   optional modifier that must be held for panning/fitting
	Fit           ImGuiMouseButton // LMB    initiates fit when double clicked
	Select        ImGuiMouseButton // RMB    begins box selection when pressed and confirms selection when released
	SelectCancel  ImGuiMouseButton // LMB    cancels active box selection when pressed; cannot be same as Select
	SelectMod     int32            // none   optional modifier that must be held for box selection
	SelectHorzMod int32            // Alt    expands active box selection horizontally to plot edge when held
	SelectVertMod int32            // Shift  expands active box selection vertically to plot edge when held
	Menu          ImGuiMouseButton // RMB    opens context menus (if enabled) when clicked
	OverrideMod   int32            // Ctrl   when held, all input is ignored; used to enable axis/plots as DND sources
	ZoomMod       int32            // none   optional modifier that must be held for scroll wheel zooming
	ZoomRate      float32          // 0.1f   zoom rate for scroll (e.g. 0.1f = 10% plot range every scroll click); make negative to invert
}

func CreateContext() *ImPlotContext        { return imPlot_CreateContext.Fn()() }
func (ctx *ImPlotContext) Destroy()        { imPlot_DestroyContext.Fn()(ctx) }
func GetCurrentContext() *ImPlotContext    { return imPlot_GetCurrentContext.Fn()() }
func SetCurrentContext(ctx *ImPlotContext) { imPlot_SetCurrentContext.Fn()(ctx) }
func SetImGuiContext(ctx *ig.ImGuiContext) { imPlot_SetImGuiContext.Fn()(ctx) }

func BeginPlot(titleId cc.String, size ig.ImVec2, flags ImPlotFlags) bool {
	return imPlot_BeginPlot.Fn()(titleId, size, flags) != 0
}
func EndPlot() { imPlot_EndPlot.Fn()() }

func BeginSubplots(titleId cc.String, row, col int32, size ig.ImVec2,
	flags ImPlotSubplotFlags, rowRatios, colRatios []float32) bool {
	rr, cr := (*float32)(nil), (*float32)(nil)
	if rowRatios != nil {
		rr = &rowRatios[0]
	}
	if colRatios != nil {
		cr = &colRatios[0]
	}
	return imPlot_BeginSubplots.Fn()(titleId, row, col, size, flags, rr, cr) != 0
}

func EndSubplots() { imPlot_EndSubplots.Fn()() }
func SetupAxis(axis ImAxis, label cc.String, flags ImPlotAxisFlags) {
	imPlot_SetupAxis.Fn()(axis, label, flags)
}
func SetupAxisLimits(axis ImAxis, vMin, vMax float64, cond ImPlotCond /*=2*/) {
	imPlot_SetupAxisLimits.Fn()(axis, vMin, vMax, cond)
}
func SetupAxisLinks(axis ImAxis, linkMin, linkMax *float64) {
	imPlot_SetupAxisLinks.Fn()(axis, linkMin, linkMax)
}
func SetupAxisFormat(axis ImAxis, fmt cc.String) {
	imPlot_SetupAxisFormat.Fn()(axis, fmt)
}
func SetupAxisFormatCallback(axis ImAxis, formatter cc.Func /*int (*ImPlotFormatter)(double value, char* buff, int size, void* user_data)*/, data uptr) {
	imPlot_SetupAxisFormatCallback.Fn()(axis, formatter, data)
}
func SetupAxisTicks(axis ImAxis, values []float64, labels cc.String /*=0*/, keepDefault bool /*=false*/) {
	v := (*float64)(nil)
	n := int32(len(values))
	if n != 0 {
		v = &values[0]
	}
	imPlot_SetupAxisTicks.Fn()(axis, v, n, labels, cc.MakeBool(keepDefault))
}
func SetupAxisTicksRange(axis ImAxis, vMin, vMax float64, n_ticks int32, labels cc.String /*=0*/, keepDefault bool /*=false*/) {
	imPlot_SetupAxisTicksRange.Fn()(axis, vMin, vMax, n_ticks, labels, cc.MakeBool(keepDefault))
}
func SetupAxisScale(axis ImAxis, scale ImPlotScale) { imPlot_SetupAxisScale.Fn()(axis, scale) }
func SetupAxisScaleTransform(axis ImAxis, forward, inverse cc.Func /*double (*ImPlotTransform)(double value, void* user_data);*/, data uptr) {
	imPlot_SetupAxisScaleTransform.Fn()(axis, forward, inverse, data)
}
func SetupAxisLimitsConstraints(axis ImAxis, vMin, vMax float64) {
	imPlot_SetupAxisLimitsConstraints.Fn()(axis, vMin, vMax)
}
func SetupAxisZoomConstraints(axis ImAxis, vMin, vMax float64) {
	imPlot_SetupAxisZoomConstraints.Fn()(axis, vMin, vMax)
}

func SetupAxes(xLabel, yLabel cc.String, xFlags, yFlags ImPlotAxisFlags) {
	imPlot_SetupAxes.Fn()(xLabel, yLabel, xFlags, yFlags)
}
func SetupAxesLimits(xMin, xMax, yMin, yMax float64, cond ImPlotCond) {
	imPlot_SetupAxesLimits.Fn()(xMin, xMax, yMin, yMax, cond)
}
func SetupLegend(location ImPlotLocation, flags ImPlotLegendFlags) {
	imPlot_SetupLegend.Fn()(location, flags)
}
func SetupMouseText(location ImPlotLocation, flags ImPlotMouseTextFlags) {
	imPlot_SetupMouseText.Fn()(location, flags)
}
func SetupFinish() { imPlot_SetupFinish.Fn()() }

func SetNextAxisLimits(axis ImAxis, vMin, vMax float64, cond ImPlotCond /*=2*/) {
	imPlot_SetNextAxisLimits.Fn()(axis, vMin, vMax, cond)
}
func SetNextAxisLinks(axis ImAxis, linkMin, linkMax *float64) {
	imPlot_SetNextAxisLinks.Fn()(axis, linkMin, linkMax)
}
func SetNextAxisToFit(axis ImAxis) { imPlot_SetNextAxisToFit.Fn()(axis) }
func SetNextAxesLimits(xMin, xMax, yMin, yMax float64, cond ImPlotCond) {
	imPlot_SetNextAxesLimits.Fn()(xMin, xMax, yMin, yMax, cond)
}
func SetNextAxesToFit() { imPlot_SetNextAxesToFit.Fn()() }

func PlotLine[T any](labelId cc.String, dtype ig.ImGuiDataType, values []T, offset int32,
	xscale /*=1*/, xstart /*=0*/ float64, flags /*=0*/ ImPlotLineFlags) {
	v, l := (*T)(nil), int32(len(values))
	if l > 0 {
		v = &values[0]
	}

	imPlot_PlotLine.Fn()(dtype, labelId, uptr(v), l, xscale, xstart, flags, offset)
}
func PlotLine2D[T any](labelId cc.String, dtype ig.ImGuiDataType, xs, ys []T, offset int32,
	flags /*=0*/ ImPlotLineFlags) {
	xv, yv := (*T)(nil), (*T)(nil)
	xl, yl := int32(len(xs)), int32(len(ys))

	count := min(xl, yl)

	if xl > 0 {
		xv = &xs[0]
	}
	if yl > 0 {
		yv = &ys[0]
	}

	imPlot_PlotLine2D.Fn()(dtype, labelId, uptr(xv), uptr(yv), count, flags, offset)
}
func PlotLineG(labelId cc.String, getter cc.Func, /*ImPlotPoint (*ImPlotGetter)(int idx, void* user_data);*/
	count int32, flags ImPlotLineFlags, data uptr) {
	imPlot_PlotLineG.Fn()(labelId, getter, data, count, flags)
}

func PlotScatter[T any](labelId cc.String, dtype ig.ImGuiDataType, values []T, offset int32,
	xscale /*=1*/, xstart float64 /*=0*/, flags /*=0*/ ImPlotScatterFlags) {
	v, l := (*T)(nil), int32(len(values))
	if l > 0 {
		v = &values[0]
	}
	imPlot_PlotScatter.Fn()(dtype, labelId, uptr(v), l, xscale, xstart, flags, offset)
}
func PlotScatter2D[T any](labelId cc.String, dtype ig.ImGuiDataType, xs, ys []T, offset /*=0*/ int32, flags /*=0*/ ImPlotScatterFlags) {

	xv, yv := (*T)(nil), (*T)(nil)
	xl, yl := int32(len(xs)), int32(len(ys))

	count := min(xl, yl)

	if xl > 0 {
		xv = &xs[0]
	}
	if yl > 0 {
		yv = &ys[0]
	}

	imPlot_PlotScatter2D.Fn()(dtype, labelId, uptr(xv), uptr(yv), count, flags, offset)
}
func PlotScatterG(labelId cc.String, getter cc.Func /* ImPlotPoint (*ImPlotGetter)(int idx, void* user_data) */, count int32, flags ImPlotScatterFlags /* = 0 */, data uptr) {
	imPlot_PlotScatterG.Fn()(labelId, getter, data, count, flags)
}

func PlotStairs[T any](labelId cc.String, dtype ig.ImGuiDataType, values []T, offset /*= 0*/ int32,
	xscale /*= 1*/ float64, xstart /*= 0*/ float64, flags /*= 0*/ ImPlotStairsFlags) {
	v, l := (*T)(nil), int32(len(values))
	if l > 0 {
		v = &values[0]
	}
	imPlot_PlotStairs.Fn()(dtype, labelId, uptr(v), l, xscale, xstart, flags, offset)
}
func PlotStairs2D[T any](labelId cc.String, dtype ig.ImGuiDataType, xs, ys []T, offset /*= 0*/ int32, flags /*= 0*/ ImPlotStairsFlags) {
	xv, yv := (*T)(nil), (*T)(nil)
	xl, yl := int32(len(xs)), int32(len(ys))

	count := min(xl, yl)

	if xl > 0 {
		xv = &xs[0]
	}
	if yl > 0 {
		yv = &ys[0]
	}

	imPlot_PlotStairs2D.Fn()(dtype, labelId, uptr(xv), uptr(yv), count, flags, offset)
}
func PlotStairsG(labelId cc.String, getter cc.Func /* ImPlotPoint (*ImPlotGetter)(int idx, void* user_data) */, count int32, flags /*= 0*/ ImPlotStairsFlags, data uptr) {
	imPlot_PlotStairsG.Fn()(labelId, getter, data, count, flags)
}

func PlotShaded[T any](labelId cc.String, dtype ig.ImGuiDataType, values []T, offset int32,
	yref /*=0*/, xscale /*=1*/, xstart /*=0*/ float64, flags /*=0*/ ImPlotShadedFlags) {
	v, l := (*T)(nil), int32(len(values))
	if l > 0 {
		v = &values[0]
	}
	imPlot_PlotShaded.Fn()(dtype, labelId, uptr(v), l, yref, xscale, xstart, flags, offset)
}
func PlotShaded2D[T any](labelId cc.String, dtype ig.ImGuiDataType, xs, ys []T, offset int32,
	yref /*=0*/ float64, flags /*=0*/ ImPlotShadedFlags) {
	xv, yv := (*T)(nil), (*T)(nil)
	xl, yl := int32(len(xs)), int32(len(ys))

	count := min(xl, yl)

	if xl > 0 {
		xv = &xs[0]
	}
	if yl > 0 {
		yv = &ys[0]
	}

	imPlot_PlotShaded2D.Fn()(dtype, labelId, uptr(xv), uptr(yv), count, yref, flags, offset)
}
func PlotShaded2Y[T any](labelId cc.String, dtype ig.ImGuiDataType, xs, ys1, ys2 []T, offset int32,
	flags /*=0*/ ImPlotShadedFlags) {
	xv, yv1, yv2 := (*T)(nil), (*T)(nil), (*T)(nil)
	xl, yl1, yl2 := int32(len(xs)), int32(len(ys1)), int32(len(ys2))

	count := min(xl, yl1, yl2)

	if xl > 0 {
		xv = &xs[0]
	}
	if yl1 > 0 {
		yv1 = &ys1[0]
	}
	if yl2 > 0 {
		yv2 = &ys2[0]
	}

	imPlot_PlotShaded2Y.Fn()(dtype, labelId, uptr(xv), uptr(yv1), uptr(yv2), count, flags, offset)
}
func PlotShadedG(labelId cc.String, getter1 cc.Func, /*ImPlotPoint (*ImPlotGetter)(int idx, void* user_data);*/
	data1 uptr, getter2 cc.Func, /*ImPlotPoint (*ImPlotGetter)(int idx, void* user_data);*/
	data2 uptr, count int32, flags /*=0*/ ImPlotShadedFlags) {
	imPlot_PlotShadedG.Fn()(labelId, getter1, data1, getter2, data2, count, flags)
}

func PlotBars[T any](labelId cc.String, dtype ig.ImGuiDataType, values []T, offset int32,
	bar_size /*=0.67*/, shift /*=0*/ float64, flags /*=0*/ ImPlotBarsFlags) {
	v, l := (*T)(nil), int32(len(values))
	if l > 0 {
		v = &values[0]
	}
	imPlot_PlotBars.Fn()(dtype, labelId, uptr(v), l, bar_size, shift, flags, offset)
}
func PlotBars2D[T any](labelId cc.String, dtype ig.ImGuiDataType, xs, ys []T, offset int32,
	bar_size float64, flags /*=0*/ ImPlotBarsFlags) {
	xv, yv := (*T)(nil), (*T)(nil)
	xl, yl := int32(len(xs)), int32(len(ys))

	count := min(xl, yl)

	if xl > 0 {
		xv = &xs[0]
	}
	if yl > 0 {
		yv = &ys[0]
	}

	imPlot_PlotBars2D.Fn()(dtype, labelId, uptr(xv), uptr(yv), count, bar_size, flags, offset)
}
func PlotBarsG(labelId cc.String, getter cc.Func, /*ImPlotPoint (*ImPlotGetter)(int idx, void* user_data);*/
	count int32, bar_size float64, flags /*=0*/ ImPlotBarsFlags, data uptr) {
	imPlot_PlotBarsG.Fn()(labelId, getter, data, count, bar_size, flags)
}

func PlotBarGroups[T any](labelIds cc.Strings, dtype ig.ImGuiDataType, values []T, itemCount, groupCount int32, groupSize /*=0.67*/, shift /*=0*/ float64, flags /*=0*/ ImPlotBarGroupsFlags) {
	v, l := (*T)(nil), int32(len(values))
	if l > 0 {
		v = &values[0]
	}
	imPlot_PlotBarGroups.Fn()(dtype, labelIds, uptr(v), itemCount, groupCount, groupSize, shift, flags)
}
func PlotErrorBars[T any](labelId cc.String, dtype ig.ImGuiDataType, xs, ys, err []T, offset int32, flags /*=0*/ ImPlotErrorBarsFlags) {
	xv, yv, ev := (*T)(nil), (*T)(nil), (*T)(nil)
	xl, yl, el := int32(len(xs)), int32(len(ys)), int32(len(err))

	count := xl
	if count > yl {
		count = yl
	}
	if count > el {
		count = el
	}

	if xl > 0 {
		xv = &xs[0]
	}
	if yl > 0 {
		yv = &ys[0]
	}
	if el > 0 {
		ev = &err[0]
	}

	imPlot_PlotErrorBars.Fn()(dtype, labelId, uptr(xv), uptr(yv), uptr(ev), count, flags, offset)
}
func PlotErrorBarsNeg[T any](labelId cc.String, dtype ig.ImGuiDataType, xs, ys, neg, pos []T, offset int32, flags /*=0*/ ImPlotErrorBarsFlags) {
	xv, yv, nv, pv := (*T)(nil), (*T)(nil), (*T)(nil), (*T)(nil)
	xl, yl, nl, pl := int32(len(xs)), int32(len(ys)), int32(len(neg)), int32(len(pos))

	count := xl
	if count > yl {
		count = yl
	}
	if count > nl {
		count = nl
	}
	if count > pl {
		count = pl
	}

	if xl > 0 {
		xv = &xs[0]
	}
	if yl > 0 {
		yv = &ys[0]
	}
	if nl > 0 {
		nv = &neg[0]
	}
	if pl > 0 {
		pv = &pos[0]
	}

	imPlot_PlotErrorBarsNeg.Fn()(dtype, labelId, uptr(xv), uptr(yv), uptr(nv), uptr(pv), count, flags, offset)
}
func PlotStems[T any](labelId cc.String, dtype ig.ImGuiDataType, values []T, offset int32, ref /*=0*/, scale /*=1*/, start /*=0*/ float64, flags /*=0*/ ImPlotStemsFlags) {
	v, l := (*T)(nil), int32(len(values))
	if l > 0 {
		v = &values[0]
	}
	imPlot_PlotStems.Fn()(dtype, labelId, uptr(v), l, ref, scale, start, flags, offset)
}
func PlotStems2D[T any](labelId cc.String, dtype ig.ImGuiDataType, xs, ys []T, offset int32, ref /*=0*/ float64, flags /*=0*/ ImPlotStemsFlags) {
	xv, yv := (*T)(nil), (*T)(nil)
	xl, yl := int32(len(xs)), int32(len(ys))

	count := min(xl, yl)

	if xl > 0 {
		xv = &xs[0]
	}
	if yl > 0 {
		yv = &ys[0]
	}

	imPlot_PlotStems2D.Fn()(dtype, labelId, uptr(xv), uptr(yv), count, ref, flags, offset)
}
func PlotInfLines[T any](labelId cc.String, dtype ig.ImGuiDataType, values []T, offset int32, flags /*=0*/ ImPlotInfLinesFlags) {
	v, l := (*T)(nil), int32(len(values))
	if l > 0 {
		v = &values[0]
	}
	imPlot_PlotInfLines.Fn()(dtype, labelId, uptr(v), l, flags, offset)
}
func PlotPieChartf[T any](labelIds cc.Strings, dtype ig.ImGuiDataType, values []T, x, y, radius float64, fmt cc.Func /* void (*ImPlotFormatter)(double value, char* buf, int size, void* user_data) */, fmtData uptr, angle0 /*=90*/ float64, flags /*=0*/ ImPlotPieChartFlags) {
	v, l := (*T)(nil), int32(len(values))
	if l > 0 {
		v = &values[0]
	}

	imPlot_PlotPieChart_Fmt.Fn()(dtype, labelIds, uptr(v), l, x, y, radius, fmt, fmtData, angle0, flags)
}
func PlotPieChart[T any](labelIds cc.Strings, dtype ig.ImGuiDataType, values []T, x, y, radius float64, labelFmt /*="%.1f"*/ cc.String, angle0 /*=90*/ float64, flags /*=0*/ ImPlotPieChartFlags) {
	v, count := (*T)(nil), int32(len(values))
	if count > 0 {
		v = &values[0]
	}
	imPlot_PlotPieChart.Fn()(dtype, labelIds, uptr(v), count, x, y, radius, labelFmt, angle0, flags)
}

func PlotHeatmap[T any](labelId cc.String, dtype ig.ImGuiDataType, values []T, rows, cols int32, scaleMin /*=0*/, scaleMax /*=0*/ float64, labelFmt /*="%.1f"*/ cc.String, boundsMin /*=ImPlotPoint(0,0)*/, boundsMax /*=ImPlotPoint(1,1)*/ ImPlotPoint, flags /*=0*/ ImPlotHeatmapFlags) {
	v, l := (*T)(nil), int32(len(values))
	if l > 0 {
		v = &values[0]
	}
	imPlot_PlotHeatmap.Fn()(dtype, labelId, uptr(v), rows, cols, scaleMin, scaleMax, labelFmt, boundsMin, boundsMax, flags)
}
func PlotHistogram[T any](labelId cc.String, dtype ig.ImGuiDataType, values []T, bins /*=ImPlotBin_Sturges*/ int32, barScale /*=1.0*/ float64, range_ ImPlotRange, flags /*=0*/ ImPlotHistogramFlags) float64 {
	v, l := (*T)(nil), int32(len(values))
	if l > 0 {
		v = &values[0]
	}
	return imPlot_PlotHistogram.Fn()(dtype, labelId, uptr(v), l, bins, barScale, range_, flags)
}

func PlotHistogram2D[T any](labelId cc.String, dtype ig.ImGuiDataType, xs, ys []T, xBins /*=ImPlotBin_Sturges*/, yBins /*=ImPlotBin_Sturges*/ int32, range_ ImPlotRect, flags /*=0*/ ImPlotHistogramFlags) float64 {
	xv, yv := (*T)(nil), (*T)(nil)
	xl, yl := int32(len(xs)), int32(len(ys))

	count := min(xl, yl)
	if xl > 0 {
		xv = &xs[0]
	}
	if yl > 0 {
		yv = &ys[0]
	}
	return imPlot_PlotHistogram2D.Fn()(dtype, labelId, uptr(xv), uptr(yv), count, xBins, yBins, range_, flags)
}

func PlotDigital[T any](labelId cc.String, dtype ig.ImGuiDataType, xs, ys []T, offset int32, flags /*=0*/ ImPlotDigitalFlags) {
	xv, yv := (*T)(nil), (*T)(nil)
	xl, yl := int32(len(xs)), int32(len(ys))

	count := min(xl, yl)

	if xl > 0 {
		xv = &xs[0]
	}
	if yl > 0 {
		yv = &ys[0]
	}
	imPlot_PlotDigital.Fn()(dtype, labelId, uptr(xv), uptr(yv), count, flags, offset)
}
func PlotDigitalG(labelId cc.String, getter cc.Func /* ImPlotPoint (*ImPlotGetter)(int idx, void* user_data) */, data uptr, count int32, flags /*=0*/ ImPlotDigitalFlags) {
	imPlot_PlotDigitalG.Fn()(labelId, getter, data, count, flags)
}
func PlotImage(labelId cc.String, texRef ig.ImTextureID, boundsMin, boundsMax ImPlotPoint, uv0 /*=ImVec2(0,0)*/, uv1 /*=ImVec2(1,1)*/ ig.ImVec2, tintCol /*=ImVec4(1,1,1,1)*/ ig.ImVec4, flags /*=0*/ ImPlotImageFlags) {
	imPlot_PlotImage.Fn()(labelId, texRef, boundsMin, boundsMax, uv0, uv1, tintCol, flags)
}
func PlotText(text cc.String, x, y float64, pixOffset /*=ImVec2(0,0)*/ ig.ImVec2, flags /*=0*/ ImPlotTextFlags) {
	imPlot_PlotText.Fn()(text, x, y, pixOffset, flags)
}
func PlotDummy(labelId cc.String, flags /*=0*/ ImPlotDummyFlags) {
	imPlot_PlotDummy.Fn()(labelId, flags)
}

func DragPoint(id int32, x, y *float64, col ig.ImVec4, size /*=4*/ float32, flags /*=0*/ ImPlotDragToolFlags, outClicked, outHovered, held *bool) bool {
	oc, oh, h := pBool(outClicked), pBool(outHovered), pBool(held)
	return imPlot_DragPoint.Fn()(id, x, y, col, size, flags, oc, oh, h) != 0
}
func DragLineX(id int32, x *float64, col ig.ImVec4, thickness /*=1*/ float32, flags /*=0*/ ImPlotDragToolFlags, outClicked, outHovered, held *bool) bool {
	oc, oh, h := pBool(outClicked), pBool(outHovered), pBool(held)
	return imPlot_DragLineX.Fn()(id, x, col, thickness, flags, oc, oh, h) != 0
}
func DragLineY(id int32, y *float64, col ig.ImVec4, thickness /*=1*/ float32, flags /*=0*/ ImPlotDragToolFlags, outClicked, outHovered, held *bool) bool {
	oc, oh, h := pBool(outClicked), pBool(outHovered), pBool(held)
	return imPlot_DragLineY.Fn()(id, y, col, thickness, flags, oc, oh, h) != 0
}
func DragRect(id int32, x1, y1, x2, y2 *float64, col ig.ImVec4, flags /*=0*/ ImPlotDragToolFlags, outClicked, outHovered, held *bool) bool {
	oc, oh, h := pBool(outClicked), pBool(outHovered), pBool(held)
	return imPlot_DragRect.Fn()(id, x1, y1, x2, y2, col, flags, oc, oh, h) != 0
}
func Annotation(x, y float64, col ig.ImVec4, pixOffset /*=ImVec2(0,0)*/ ig.ImVec2, clamp /*=false*/ bool, round /*=false*/ bool) {
	imPlot_Annotation_Round.Fn()(x, y, col, pixOffset, cc.MakeBool(clamp), cc.MakeBool(round))
}
func AnnotationText(x, y float64, col ig.ImVec4, pixOffset /*=ImVec2(0,0)*/ ig.ImVec2, clamp /*=false*/ bool, text cc.String) {
	imPlot_Annotation.Fn()(x, y, col, pixOffset, cc.MakeBool(clamp), text)
}
func AnnotationTextf(x, y float64, col ig.ImVec4, pixOffset /*=ImVec2(0,0)*/ ig.ImVec2, clamp /*=false*/ bool, fmt cc.String, args ...any) {
	imPlot_Annotation_Fmt.FnVa()(x, y, col, pixOffset, cc.MakeBool(clamp), fmt, append(args, uptr(nil))...)
}
func TagX(x float64, col ig.ImVec4, round /*=false*/ bool) {
	imPlot_TagX.Fn()(x, col, cc.MakeBool(round))
}
func TagXText(x float64, col ig.ImVec4, text cc.String) {
	imPlot_TagX_Text.Fn()(x, col, text)
}
func TagXTextf(x float64, col ig.ImVec4, fmt cc.String, args ...any) {
	imPlot_TagX_Fmt.FnVa()(x, col, fmt, append(args, uptr(nil))...)
}
func TagY(y float64, col ig.ImVec4, round /*=false*/ bool) {
	imPlot_TagY.Fn()(y, col, cc.MakeBool(round))
}
func TagYText(y float64, col ig.ImVec4, text cc.String) {
	imPlot_TagY_Text.Fn()(y, col, text)
}
func TagYTextf(y float64, col ig.ImVec4, fmt cc.String, args ...any) {
	imPlot_TagY_Fmt.FnVa()(y, col, fmt, append(args, uptr(nil))...)
}

func SetAxis(axis ImAxis)         { imPlot_SetAxis.Fn()(axis) }
func SetAxes(xAxis, yAxis ImAxis) { imPlot_SetAxes.Fn()(xAxis, yAxis) }
func PixelsToPlot(pix ig.ImVec2, xAxis /*=IMPLOT_AUTO*/, yAxis /*=IMPLOT_AUTO*/ ImAxis) ImPlotPoint {
	return imPlot_PixelsToPlot.Fn()(pix, xAxis, yAxis)
}
func PixelsToPlotXY(x, y float32, xAxis /*=IMPLOT_AUTO*/, yAxis /*=IMPLOT_AUTO*/ ImAxis) ImPlotPoint {
	return imPlot_PixelsToPlotXY.Fn()(x, y, xAxis, yAxis)
}
func PlotToPixels(plt ImPlotPoint, xAxis /*=IMPLOT_AUTO*/, yAxis /*=IMPLOT_AUTO*/ ImAxis) ig.ImVec2 {
	return imPlot_PlotToPixels.Fn()(plt, xAxis, yAxis)
}
func PlotToPixelsXY(x, y float64, xAxis /*=IMPLOT_AUTO*/, yAxis /*=IMPLOT_AUTO*/ ImAxis) ig.ImVec2 {
	return imPlot_PlotToPixelsXY.Fn()(x, y, xAxis, yAxis)
}
func GetPlotPos() ig.ImVec2  { return imPlot_GetPlotPos.Fn()() }
func GetPlotSize() ig.ImVec2 { return imPlot_GetPlotSize.Fn()() }
func GetPlotMousePos(xAxis /*=IMPLOT_AUTO*/, yAxis /*=IMPLOT_AUTO*/ ImAxis) ImPlotPoint {
	return imPlot_GetPlotMousePos.Fn()(xAxis, yAxis)
}
func GetPlotLimits(xAxis /*=IMPLOT_AUTO*/, yAxis /*=IMPLOT_AUTO*/ ImAxis) ImPlotRect {
	return imPlot_GetPlotLimits.Fn()(xAxis, yAxis)
}
func IsPlotHovered() bool            { return imPlot_IsPlotHovered.Fn()() != 0 }
func IsAxisHovered(axis ImAxis) bool { return imPlot_IsAxisHovered.Fn()(axis) != 0 }
func IsSubplotsHovered() bool        { return imPlot_IsSubplotsHovered.Fn()() != 0 }
func IsPlotSelected() bool           { return imPlot_IsPlotSelected.Fn()() != 0 }
func GetPlotSelection(xAxis /*=IMPLOT_AUTO*/, yAxis /*=IMPLOT_AUTO*/ ImAxis) ImPlotRect {
	return imPlot_GetPlotSelection.Fn()(xAxis, yAxis)
}
func CancelPlotSelection() {
	imPlot_CancelPlotSelection.Fn()()
}
func HideNextItem(hidden /*=true*/ bool, cond /*=ImPlotCond_Once*/ ImPlotCond) {
	imPlot_HideNextItem.Fn()(cc.MakeBool(hidden), cond)
}
func BeginAlignedPlots(groupId cc.String, vertical /*=true*/ bool) bool {
	return imPlot_BeginAlignedPlots.Fn()(groupId, cc.MakeBool(vertical)) != 0
}
func EndAlignedPlots() { imPlot_EndAlignedPlots.Fn()() }

func BeginLegendPopup(labelId cc.String, mouseButton /*=1*/ ig.ImGuiMouseButton) bool {
	return imPlot_BeginLegendPopup.Fn()(labelId, mouseButton) != 0
}
func EndLegendPopup() { imPlot_EndLegendPopup.Fn()() }
func IsLegendEntryHovered(labelId cc.String) bool {
	return imPlot_IsLegendEntryHovered.Fn()(labelId) != 0
}
func BeginDragDropTargetPlot() bool            { return imPlot_BeginDragDropTargetPlot.Fn()() != 0 }
func BeginDragDropTargetAxis(axis ImAxis) bool { return imPlot_BeginDragDropTargetAxis.Fn()(axis) != 0 }
func BeginDragDropTargetLegend() bool          { return imPlot_BeginDragDropTargetLegend.Fn()() != 0 }
func EndDragDropTarget()                       { imPlot_EndDragDropTarget.Fn()() }
func BeginDragDropSourcePlot(flags /*=0*/ ig.ImGuiDragDropFlags) bool {
	return imPlot_BeginDragDropSourcePlot.Fn()(flags) != 0
}
func BeginDragDropSourceAxis(axis ImAxis, flags /*=0*/ ig.ImGuiDragDropFlags) bool {
	return imPlot_BeginDragDropSourceAxis.Fn()(axis, flags) != 0
}
func BeginDragDropSourceItem(labelId cc.String, flags /*=0*/ ig.ImGuiDragDropFlags) bool {
	return imPlot_BeginDragDropSourceItem.Fn()(labelId, flags) != 0
}
func EndDragDropSource() { imPlot_EndDragDropSource.Fn()() }

func GetStyle() *ImPlotStyle                        { return imPlot_GetStyle.Fn()() }
func StyleColorsAuto(dst /*=nil*/ *ImPlotStyle)     { imPlot_StyleColorsAuto.Fn()(dst) }
func StyleColorsClassic(dst /*=nil*/ *ImPlotStyle)  { imPlot_StyleColorsClassic.Fn()(dst) }
func StyleColorsDark(dst /*=nil*/ *ImPlotStyle)     { imPlot_StyleColorsDark.Fn()(dst) }
func StyleColorsLight(dst /*=nil*/ *ImPlotStyle)    { imPlot_StyleColorsLight.Fn()(dst) }
func PushStyleColorU32(idx ImPlotCol, col ig.ImU32) { imPlot_PushStyleColor_ImU32.Fn()(idx, col) }
func PushStyleColor(idx ImPlotCol, col ig.ImVec4)   { imPlot_PushStyleColor_ImVec4.Fn()(idx, col) }
func PopStyleColor(count /*=1*/ int32)              { imPlot_PopStyleColor.Fn()(count) }
func PushStyleVarF32(idx ImPlotStyleVar, val float32) {
	imPlot_PushStyleVar_Float.Fn()(idx, val)
}
func PushStyleVarI32(idx ImPlotStyleVar, val int32) {
	imPlot_PushStyleVar_Int.Fn()(idx, val)
}
func PushStyleVar(idx ImPlotStyleVar, val ig.ImVec2) {
	imPlot_PushStyleVar_ImVec2.Fn()(idx, val)
}
func PopStyleVar(count /*=1*/ int32) {
	imPlot_PopStyleVar.Fn()(count)
}
func SetNextLineStyle(col /*=IMPLOT_AUTO_COL*/ ig.ImVec4, weight /*=IMPLOT_AUTO*/ float32) {
	imPlot_SetNextLineStyle.Fn()(col, weight)
}
func SetNextFillStyle(col /*=IMPLOT_AUTO_COL*/ ig.ImVec4, alphaMod /*=IMPLOT_AUTO*/ float32) {
	imPlot_SetNextFillStyle.Fn()(col, alphaMod)
}
func SetNextMarkerStyle(marker /*=IMPLOT_AUTO*/ ImPlotMarker, size /*=IMPLOT_AUTO*/ float32, fill /*=IMPLOT_AUTO_COL*/ ig.ImVec4, weight /*=IMPLOT_AUTO*/ float32, outline /*=IMPLOT_AUTO_COL*/ ig.ImVec4) {
	imPlot_SetNextMarkerStyle.Fn()(marker, size, fill, weight, outline)
}
func SetNextErrorBarStyle(col /*=IMPLOT_AUTO_COL*/ ig.ImVec4, size /*=IMPLOT_AUTO*/ float32, weight /*=IMPLOT_AUTO*/ float32) {
	imPlot_SetNextErrorBarStyle.Fn()(col, size, weight)
}
func GetLastItemColor() ig.ImVec4               { return imPlot_GetLastItemColor.Fn()() }
func GetStyleColorName(idx ImPlotCol) cc.String { return imPlot_GetStyleColorName.Fn()(idx) }
func GetMarkerName(idx ImPlotMarker) cc.String  { return imPlot_GetMarkerName.Fn()(idx) }

func AddColormapImVec4(name cc.String, cols []ig.ImVec4, qual /*=true*/ bool) ImPlotColormap {
	c, size := (*ig.ImVec4)(nil), int32(len(cols))
	if size > 0 {
		c = &cols[0]
	}
	return imPlot_AddColormap_ImVec4.Fn()(name, c, size, cc.MakeBool(qual))
}
func AddColormap(name cc.String, cols []ig.ImU32, qual /*=true*/ bool) ImPlotColormap {
	c, size := (*ig.ImU32)(nil), int32(len(cols))
	if size > 0 {
		c = &cols[0]
	}
	return imPlot_AddColormap_ImU32.Fn()(name, c, size, cc.MakeBool(qual))
}
func GetColormapCount() int32                       { return imPlot_GetColormapCount.Fn()() }
func GetColormapName(cmap ImPlotColormap) cc.String { return imPlot_GetColormapName.Fn()(cmap) }
func GetColormap(name cc.String) ImPlotColormap     { return imPlot_GetColormapIndex.Fn()(name) }

func PushColormap(cmap ImPlotColormap) { imPlot_PushColormap_Index.Fn()(cmap) }
func PushColormapName(name cc.String)  { imPlot_PushColormap_Name.Fn()(name) }

func PopColormap(count /*=1*/ int32) { imPlot_PopColormap.Fn()(count) }
func NextColormapColor() ig.ImVec4 {
	return imPlot_NextColormapColor.Fn()()
}
func GetColormapSize(cmap /*=IMPLOT_AUTO*/ ImPlotColormap) int32 {
	return imPlot_GetColormapSize.Fn()(cmap)
}
func GetColormapColor(idx int32, cmap /*=IMPLOT_AUTO*/ ImPlotColormap) ig.ImVec4 {
	return imPlot_GetColormapColor.Fn()(idx, cmap)
}
func SampleColormap(t float32, cmap /*=IMPLOT_AUTO*/ ImPlotColormap) ig.ImVec4 {
	return imPlot_SampleColormap.Fn()(t, cmap)
}
func ColormapScale(label cc.String, scaleMin, scaleMax float64, size /*=ImVec2(0,0)*/ ig.ImVec2, format /*="%g"*/ cc.String, flags /*=0*/ ImPlotColormapScaleFlags, cmap /*=IMPLOT_AUTO*/ ImPlotColormap) {
	imPlot_ColormapScale.Fn()(label, scaleMin, scaleMax, size, format, flags, cmap)
}
func ColormapSlider(label cc.String, t *float32, out /*=nil*/ *ig.ImVec4, format /*=""*/ cc.String, cmap /*=IMPLOT_AUTO*/ ImPlotColormap) bool {
	return imPlot_ColormapSlider.Fn()(label, t, out, format, cmap) != 0
}
func ColormapButton(label cc.String, size /*=ImVec2(0,0)*/ ig.ImVec2, cmap /*=IMPLOT_AUTO*/ ImPlotColormap) bool {
	return imPlot_ColormapButton.Fn()(label, size, cmap) != 0
}
func BustColorCache(plotTitleId /*=nil*/ cc.String) {
	imPlot_BustColorCache.Fn()(plotTitleId)
}

func GetInputMap() *ImPlotInputMap                 { return imPlot_GetInputMap.Fn()() }
func MapInputDefault(dst /*=nil*/ *ImPlotInputMap) { imPlot_MapInputDefault.Fn()(dst) }
func MapInputReverse(dst /*=nil*/ *ImPlotInputMap) { imPlot_MapInputReverse.Fn()(dst) }
func ItemIconImVec4(col ig.ImVec4)                 { imPlot_ItemIcon_ImVec4.Fn()(col) }
func ItemIconImU32(col ig.ImU32)                   { imPlot_ItemIcon_ImU32.Fn()(col) }
func ColormapIcon(cmap ImPlotColormap)             { imPlot_ColormapIcon.Fn()(cmap) }
func GetPlotDrawList() *ig.ImDrawList              { return imPlot_GetPlotDrawList.Fn()() }
func PushPlotClipRect(expand /*=0*/ float32)       { imPlot_PushPlotClipRect.Fn()(expand) }
func PopPlotClipRect()                             { imPlot_PopPlotClipRect.Fn()() }
func ShowStyleSelector(label cc.String) bool       { return imPlot_ShowStyleSelector.Fn()(label) != 0 }
func ShowColormapSelector(label cc.String) bool    { return imPlot_ShowColormapSelector.Fn()(label) != 0 }
func ShowInputMapSelector(label cc.String) bool    { return imPlot_ShowInputMapSelector.Fn()(label) != 0 }
func ShowStyleEditor(ref /*=nil*/ *ImPlotStyle)    { imPlot_ShowStyleEditor.Fn()(ref) }
func ShowUserGuide()                               { imPlot_ShowUserGuide.Fn()() }
func ShowMetricsWindow(pPopen /*=nil*/ *bool) {
	imPlot_ShowMetricsWindow.Fn()(pBool(pPopen))
}

// func ShowDemoWindow(pOpen /*=nil*/ *bool) {
// 	imPlot_ShowDemoWindow.Fn()(pBool(pOpen))
// }
