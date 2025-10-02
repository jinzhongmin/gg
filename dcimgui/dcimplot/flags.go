package dcimplot

type ImPlotCond int32

const (
	ImPlotCond_None   ImPlotCond = 0      // No condition (always set the variable), same as _Always
	ImPlotCond_Always ImPlotCond = 1 << 0 // No condition (always set the variable)
	ImPlotCond_Once   ImPlotCond = 1 << 1 // Set the variable once per runtime session (only the first call will succeed)

)

type ImPlotFlags int32
type ImPlotSubplotFlags int32
type ImAxis int32
type ImPlotAxisFlags int32
type ImPlotScale int32
type ImPlotLegendFlags int32
type ImPlotLocation int32
type ImPlotMouseTextFlags int32
type ImPlotLineFlags int32
type ImPlotScatterFlags int32
type ImPlotStairsFlags int32
type ImPlotShadedFlags int32
type ImPlotBarsFlags int32
type ImPlotBarGroupsFlags int32
type ImPlotErrorBarsFlags int32
type ImPlotStemsFlags int32
type ImPlotInfLinesFlags int32
type ImPlotPieChartFlags int32
type ImPlotHeatmapFlags int32
type ImPlotHistogramFlags int32
type ImPlotDigitalFlags int32
type ImPlotImageFlags int32
type ImPlotTextFlags int32
type ImPlotDummyFlags int32
type ImPlotBin int32
type ImPlotDragToolFlags int32
type ImPlotCol int32
type ImPlotColormap int32
type ImPlotStyleVar int32
type ImPlotMarker int32
type ImPlotColormapScaleFlags int32
type ImGuiMouseButton int32

const (
	ImPlotFlags_None        ImPlotFlags = 0      // default
	ImPlotFlags_NoTitle     ImPlotFlags = 1 << 0 // the plot title will not be displayed (titles are also hidden if preceeded by double hashes, e.g. "##MyPlot")
	ImPlotFlags_NoLegend    ImPlotFlags = 1 << 1 // the legend will not be displayed
	ImPlotFlags_NoMouseText ImPlotFlags = 1 << 2 // the mouse position, in plot coordinates, will not be displayed inside of the plot
	ImPlotFlags_NoInputs    ImPlotFlags = 1 << 3 // the user will not be able to interact with the plot
	ImPlotFlags_NoMenus     ImPlotFlags = 1 << 4 // the user will not be able to open context menus
	ImPlotFlags_NoBoxSelect ImPlotFlags = 1 << 5 // the user will not be able to box-select
	ImPlotFlags_NoFrame     ImPlotFlags = 1 << 6 // the ImGui frame will not be rendered
	ImPlotFlags_Equal       ImPlotFlags = 1 << 7 // x and y axes pairs will be constrained to have the same units/pixel
	ImPlotFlags_Crosshairs  ImPlotFlags = 1 << 8 // the default mouse cursor will be replaced with a crosshair when hovered
	ImPlotFlags_CanvasOnly  ImPlotFlags = ImPlotFlags_NoTitle | ImPlotFlags_NoLegend | ImPlotFlags_NoMenus | ImPlotFlags_NoBoxSelect | ImPlotFlags_NoMouseText
)

const (
	ImPlotSubplotFlags_None       ImPlotSubplotFlags = 0       // default
	ImPlotSubplotFlags_NoTitle    ImPlotSubplotFlags = 1 << 0  // the subplot title will not be displayed (titles are also hidden if preceeded by double hashes, e.g. "##MySubplot")
	ImPlotSubplotFlags_NoLegend   ImPlotSubplotFlags = 1 << 1  // the legend will not be displayed (only applicable if ImPlotSubplotFlags_ShareItems is enabled)
	ImPlotSubplotFlags_NoMenus    ImPlotSubplotFlags = 1 << 2  // the user will not be able to open context menus with right-click
	ImPlotSubplotFlags_NoResize   ImPlotSubplotFlags = 1 << 3  // resize splitters between subplot cells will be not be provided
	ImPlotSubplotFlags_NoAlign    ImPlotSubplotFlags = 1 << 4  // subplot edges will not be aligned vertically or horizontally
	ImPlotSubplotFlags_ShareItems ImPlotSubplotFlags = 1 << 5  // items across all subplots will be shared and rendered into a single legend entry
	ImPlotSubplotFlags_LinkRows   ImPlotSubplotFlags = 1 << 6  // link the y-axis limits of all plots in each row (does not apply to auxiliary axes)
	ImPlotSubplotFlags_LinkCols   ImPlotSubplotFlags = 1 << 7  // link the x-axis limits of all plots in each column (does not apply to auxiliary axes)
	ImPlotSubplotFlags_LinkAllX   ImPlotSubplotFlags = 1 << 8  // link the x-axis limits in every plot in the subplot (does not apply to auxiliary axes)
	ImPlotSubplotFlags_LinkAllY   ImPlotSubplotFlags = 1 << 9  // link the y-axis limits in every plot in the subplot (does not apply to auxiliary axes)
	ImPlotSubplotFlags_ColMajor   ImPlotSubplotFlags = 1 << 10 // subplots are added in column major order instead of the default row major order
)
const (
	// horizontal axes
	ImAxis_X1 ImAxis = iota // enabled by default
	ImAxis_X2               // disabled by default
	ImAxis_X3               // disabled by default
	// vertical axes
	ImAxis_Y1 // enabled by default
	ImAxis_Y2 // disabled by default
	ImAxis_Y3 // disabled by default
	// bookeeping
	ImAxis_COUNT
)

const (
	ImPlotAxisFlags_None          ImPlotAxisFlags = 0       // default
	ImPlotAxisFlags_NoLabel       ImPlotAxisFlags = 1 << 0  // the axis label will not be displayed (axis labels are also hidden if the supplied string name is nullptr)
	ImPlotAxisFlags_NoGridLines   ImPlotAxisFlags = 1 << 1  // no grid lines will be displayed
	ImPlotAxisFlags_NoTickMarks   ImPlotAxisFlags = 1 << 2  // no tick marks will be displayed
	ImPlotAxisFlags_NoTickLabels  ImPlotAxisFlags = 1 << 3  // no text labels will be displayed
	ImPlotAxisFlags_NoInitialFit  ImPlotAxisFlags = 1 << 4  // axis will not be initially fit to data extents on the first rendered frame
	ImPlotAxisFlags_NoMenus       ImPlotAxisFlags = 1 << 5  // the user will not be able to open context menus with right-click
	ImPlotAxisFlags_NoSideSwitch  ImPlotAxisFlags = 1 << 6  // the user will not be able to switch the axis side by dragging it
	ImPlotAxisFlags_NoHighlight   ImPlotAxisFlags = 1 << 7  // the axis will not have its background highlighted when hovered or held
	ImPlotAxisFlags_Opposite      ImPlotAxisFlags = 1 << 8  // axis ticks and labels will be rendered on the conventionally opposite side (i.e, right or top)
	ImPlotAxisFlags_Foreground    ImPlotAxisFlags = 1 << 9  // grid lines will be displayed in the foreground (i.e. on top of data) instead of the background
	ImPlotAxisFlags_Invert        ImPlotAxisFlags = 1 << 10 // the axis will be inverted
	ImPlotAxisFlags_AutoFit       ImPlotAxisFlags = 1 << 11 // axis will be auto-fitting to data extents
	ImPlotAxisFlags_RangeFit      ImPlotAxisFlags = 1 << 12 // axis will only fit points if the point is in the visible range of the **orthogonal** axis
	ImPlotAxisFlags_PanStretch    ImPlotAxisFlags = 1 << 13 // panning in a locked or constrained state will cause the axis to stretch if possible
	ImPlotAxisFlags_LockMin       ImPlotAxisFlags = 1 << 14 // the axis minimum value will be locked when panning/zooming
	ImPlotAxisFlags_LockMax       ImPlotAxisFlags = 1 << 15 // the axis maximum value will be locked when panning/zooming
	ImPlotAxisFlags_Lock          ImPlotAxisFlags = ImPlotAxisFlags_LockMin | ImPlotAxisFlags_LockMax
	ImPlotAxisFlags_NoDecorations ImPlotAxisFlags = ImPlotAxisFlags_NoLabel | ImPlotAxisFlags_NoGridLines | ImPlotAxisFlags_NoTickMarks | ImPlotAxisFlags_NoTickLabels
	ImPlotAxisFlags_AuxDefault    ImPlotAxisFlags = ImPlotAxisFlags_NoGridLines | ImPlotAxisFlags_Opposite
)

const (
	ImPlotScale_Linear ImPlotScale = iota // default linear scale
	ImPlotScale_Time                      // date/time scale
	ImPlotScale_Log10                     // base 10 logartithmic scale
	ImPlotScale_SymLog                    // symmetric log scale
)

const (
	ImPlotLegendFlags_None            ImPlotLegendFlags = 0      // default
	ImPlotLegendFlags_NoButtons       ImPlotLegendFlags = 1 << 0 // legend icons will not function as hide/show buttons
	ImPlotLegendFlags_NoHighlightItem ImPlotLegendFlags = 1 << 1 // plot items will not be highlighted when their legend entry is hovered
	ImPlotLegendFlags_NoHighlightAxis ImPlotLegendFlags = 1 << 2 // axes will not be highlighted when legend entries are hovered (only relevant if x/y-axis count > 1)
	ImPlotLegendFlags_NoMenus         ImPlotLegendFlags = 1 << 3 // the user will not be able to open context menus with right-click
	ImPlotLegendFlags_Outside         ImPlotLegendFlags = 1 << 4 // legend will be rendered outside of the plot area
	ImPlotLegendFlags_Horizontal      ImPlotLegendFlags = 1 << 5 // legend entries will be displayed horizontally
	ImPlotLegendFlags_Sort            ImPlotLegendFlags = 1 << 6 // legend entries will be displayed in alphabetical order
)
const (
	ImPlotLocation_Center    ImPlotLocation = 0                                          // center-center
	ImPlotLocation_North     ImPlotLocation = 1 << 0                                     // top-center
	ImPlotLocation_South     ImPlotLocation = 1 << 1                                     // bottom-center
	ImPlotLocation_West      ImPlotLocation = 1 << 2                                     // center-left
	ImPlotLocation_East      ImPlotLocation = 1 << 3                                     // center-right
	ImPlotLocation_NorthWest ImPlotLocation = ImPlotLocation_North | ImPlotLocation_West // top-left
	ImPlotLocation_NorthEast ImPlotLocation = ImPlotLocation_North | ImPlotLocation_East // top-right
	ImPlotLocation_SouthWest ImPlotLocation = ImPlotLocation_South | ImPlotLocation_West // bottom-left
	ImPlotLocation_SouthEast ImPlotLocation = ImPlotLocation_South | ImPlotLocation_East // bottom-right
)
const (
	ImPlotMouseTextFlags_None       ImPlotMouseTextFlags = 0      // default
	ImPlotMouseTextFlags_NoAuxAxes  ImPlotMouseTextFlags = 1 << 0 // only show the mouse position for primary axes
	ImPlotMouseTextFlags_NoFormat   ImPlotMouseTextFlags = 1 << 1 // axes label formatters won't be used to render text
	ImPlotMouseTextFlags_ShowAlways ImPlotMouseTextFlags = 1 << 2 // always display mouse position even if plot not hovered
)

const (
	ImPlotLineFlags_None     ImPlotLineFlags = 0       // default
	ImPlotLineFlags_Segments ImPlotLineFlags = 1 << 10 // a line segment will be rendered from every two consecutive points
	ImPlotLineFlags_Loop     ImPlotLineFlags = 1 << 11 // the last and first point will be connected to form a closed loop
	ImPlotLineFlags_SkipNaN  ImPlotLineFlags = 1 << 12 // NaNs values will be skipped instead of rendered as missing data
	ImPlotLineFlags_NoClip   ImPlotLineFlags = 1 << 13 // markers (if displayed) on the edge of a plot will not be clipped
	ImPlotLineFlags_Shaded   ImPlotLineFlags = 1 << 14 // a filled region between the line and horizontal origin will be rendered; use PlotShaded for more advanced cases
)

const (
	ImPlotScatterFlags_None   ImPlotScatterFlags = 0       // default
	ImPlotScatterFlags_NoClip ImPlotScatterFlags = 1 << 10 // markers on the edge of a plot will not be clipped
)

const (
	ImPlotStairsFlags_None    ImPlotStairsFlags = 0       // default
	ImPlotStairsFlags_PreStep ImPlotStairsFlags = 1 << 10 // the y value is continued constantly to the left from every x position, i.e. the interval (x[i-1], x[i]] has the value y[i]
	ImPlotStairsFlags_Shaded  ImPlotStairsFlags = 1 << 11 // a filled region between the stairs and horizontal origin will be rendered; use PlotShaded for more advanced cases
)

const (
	ImPlotShadedFlags_None ImPlotShadedFlags = 0 // default
)

const (
	ImPlotBarsFlags_None       ImPlotBarsFlags = 0       // default
	ImPlotBarsFlags_Horizontal ImPlotBarsFlags = 1 << 10 // bars will be rendered horizontally on the current y-axis
)

const (
	ImPlotBarGroupsFlags_None       ImPlotBarGroupsFlags = 0       // default
	ImPlotBarGroupsFlags_Horizontal ImPlotBarGroupsFlags = 1 << 10 // bar groups will be rendered horizontally on the current y-axis
	ImPlotBarGroupsFlags_Stacked    ImPlotBarGroupsFlags = 1 << 11 // items in a group will be stacked on top of each other
)
const (
	ImPlotErrorBarsFlags_None       ImPlotErrorBarsFlags = 0       // default
	ImPlotErrorBarsFlags_Horizontal ImPlotErrorBarsFlags = 1 << 10 // error bars will be rendered horizontally on the current y-axis
)
const (
	ImPlotStemsFlags_None       ImPlotStemsFlags = 0       // default
	ImPlotStemsFlags_Horizontal ImPlotStemsFlags = 1 << 10 // stems will be rendered horizontally on the current y-axis
)

const (
	ImPlotInfLinesFlags_None       ImPlotInfLinesFlags = 0       // default
	ImPlotInfLinesFlags_Horizontal ImPlotInfLinesFlags = 1 << 10 // lines will be rendered horizontally on the current y-axis
)

const (
	ImPlotPieChartFlags_None         ImPlotPieChartFlags = 0       // default
	ImPlotPieChartFlags_Normalize    ImPlotPieChartFlags = 1 << 10 // force normalization of pie chart values (i.e. always make a full circle if sum < 0)
	ImPlotPieChartFlags_IgnoreHidden ImPlotPieChartFlags = 1 << 11 // ignore hidden slices when drawing the pie chart (as if they were not there)
	ImPlotPieChartFlags_Exploding    ImPlotPieChartFlags = 1 << 12 // Explode legend-hovered slice
)
const (
	ImPlotHeatmapFlags_None     ImPlotHeatmapFlags = 0       // default
	ImPlotHeatmapFlags_ColMajor ImPlotHeatmapFlags = 1 << 10 // data will be read in column major order
)

const (
	ImPlotHistogramFlags_None       ImPlotHistogramFlags = 0       // default
	ImPlotHistogramFlags_Horizontal ImPlotHistogramFlags = 1 << 10 // histogram bars will be rendered horizontally (not supported by PlotHistogram2D)
	ImPlotHistogramFlags_Cumulative ImPlotHistogramFlags = 1 << 11 // each bin will contain its count plus the counts of all previous bins (not supported by PlotHistogram2D)
	ImPlotHistogramFlags_Density    ImPlotHistogramFlags = 1 << 12 // counts will be normalized, i.e. the PDF will be visualized, or the CDF will be visualized if Cumulative is also set
	ImPlotHistogramFlags_NoOutliers ImPlotHistogramFlags = 1 << 13 // exclude values outside the specifed histogram range from the count toward normalizing and cumulative counts
	ImPlotHistogramFlags_ColMajor   ImPlotHistogramFlags = 1 << 14 // data will be read in column major order (not supported by PlotHistogram)
)

const (
	ImPlotDigitalFlags_None = 0 // default
)

const (
	ImPlotImageFlags_None = 0 // default
)

const (
	ImPlotTextFlags_None     ImPlotTextFlags = 0       // default
	ImPlotTextFlags_Vertical ImPlotTextFlags = 1 << 10 // text will be rendered vertically
)

const (
	ImPlotDummyFlags_None = 0 // default
)

const (
	ImPlotBin_Sqrt    ImPlotBin = -1 // k = sqrt(n)
	ImPlotBin_Sturges ImPlotBin = -2 // k = 1 + log2(n)
	ImPlotBin_Rice    ImPlotBin = -3 // k = 2 * cbrt(n)
	ImPlotBin_Scott   ImPlotBin = -4 // w = 3.49 * sigma / cbrt(n)
)

const (
	ImPlotDragToolFlags_None      ImPlotDragToolFlags = 0      // default
	ImPlotDragToolFlags_NoCursors ImPlotDragToolFlags = 1 << 0 // drag tools won't change cursor icons when hovered or held
	ImPlotDragToolFlags_NoFit     ImPlotDragToolFlags = 1 << 1 // the drag tool won't be considered for plot fits
	ImPlotDragToolFlags_NoInputs  ImPlotDragToolFlags = 1 << 2 // lock the tool from user inputs
	ImPlotDragToolFlags_Delayed   ImPlotDragToolFlags = 1 << 3 // tool rendering will be delayed one frame; useful when applying position-constraints
)

const (
	// item styling colors
	ImPlotCol_Line          ImPlotCol = iota // plot line/outline color (defaults to next unused color in current colormap)
	ImPlotCol_Fill                           // plot fill color for bars (defaults to the current line color)
	ImPlotCol_MarkerOutline                  // marker outline color (defaults to the current line color)
	ImPlotCol_MarkerFill                     // marker fill color (defaults to the current line color)
	ImPlotCol_ErrorBar                       // error bar color (defaults to ImGuiCol_Text)
	// plot styling colors
	ImPlotCol_FrameBg       // plot frame background color (defaults to ImGuiCol_FrameBg)
	ImPlotCol_PlotBg        // plot area background color (defaults to ImGuiCol_WindowBg)
	ImPlotCol_PlotBorder    // plot area border color (defaults to ImGuiCol_Border)
	ImPlotCol_LegendBg      // legend background color (defaults to ImGuiCol_PopupBg)
	ImPlotCol_LegendBorder  // legend border color (defaults to ImPlotCol_PlotBorder)
	ImPlotCol_LegendText    // legend text color (defaults to ImPlotCol_InlayText)
	ImPlotCol_TitleText     // plot title text color (defaults to ImGuiCol_Text)
	ImPlotCol_InlayText     // color of text appearing inside of plots (defaults to ImGuiCol_Text)
	ImPlotCol_AxisText      // axis label and tick lables color (defaults to ImGuiCol_Text)
	ImPlotCol_AxisGrid      // axis grid color (defaults to 25% ImPlotCol_AxisText)
	ImPlotCol_AxisTick      // axis tick color (defaults to AxisGrid)
	ImPlotCol_AxisBg        // background color of axis hover region (defaults to transparent)
	ImPlotCol_AxisBgHovered // axis hover color (defaults to ImGuiCol_ButtonHovered)
	ImPlotCol_AxisBgActive  // axis active color (defaults to ImGuiCol_ButtonActive)
	ImPlotCol_Selection     // box-selection color (defaults to yellow)
	ImPlotCol_Crosshairs    // crosshairs color (defaults to ImPlotCol_PlotBorder)
	ImPlotCol_COUNT
)

const (
	// item styling variables
	ImPlotStyleVar_LineWeight       ImPlotStyleVar = iota // float,  plot item line weight in pixels
	ImPlotStyleVar_Marker                                 // int,    marker specification
	ImPlotStyleVar_MarkerSize                             // float,  marker size in pixels (roughly the marker's "radius")
	ImPlotStyleVar_MarkerWeight                           // float,  plot outline weight of markers in pixels
	ImPlotStyleVar_FillAlpha                              // float,  alpha modifier applied to all plot item fills
	ImPlotStyleVar_ErrorBarSize                           // float,  error bar whisker width in pixels
	ImPlotStyleVar_ErrorBarWeight                         // float,  error bar whisker weight in pixels
	ImPlotStyleVar_DigitalBitHeight                       // float,  digital channels bit height (at 1) in pixels
	ImPlotStyleVar_DigitalBitGap                          // float,  digital channels bit padding gap in pixels
	// plot styling variables
	ImPlotStyleVar_PlotBorderSize     // float,  thickness of border around plot area
	ImPlotStyleVar_MinorAlpha         // float,  alpha multiplier applied to minor axis grid lines
	ImPlotStyleVar_MajorTickLen       // ImVec2, major tick lengths for X and Y axes
	ImPlotStyleVar_MinorTickLen       // ImVec2, minor tick lengths for X and Y axes
	ImPlotStyleVar_MajorTickSize      // ImVec2, line thickness of major ticks
	ImPlotStyleVar_MinorTickSize      // ImVec2, line thickness of minor ticks
	ImPlotStyleVar_MajorGridSize      // ImVec2, line thickness of major grid lines
	ImPlotStyleVar_MinorGridSize      // ImVec2, line thickness of minor grid lines
	ImPlotStyleVar_PlotPadding        // ImVec2, padding between widget frame and plot area, labels, or outside legends (i.e. main padding)
	ImPlotStyleVar_LabelPadding       // ImVec2, padding between axes labels, tick labels, and plot edge
	ImPlotStyleVar_LegendPadding      // ImVec2, legend padding from plot edges
	ImPlotStyleVar_LegendInnerPadding // ImVec2, legend inner padding from legend edges
	ImPlotStyleVar_LegendSpacing      // ImVec2, spacing between legend entries
	ImPlotStyleVar_MousePosPadding    // ImVec2, padding between plot edge and interior info text
	ImPlotStyleVar_AnnotationPadding  // ImVec2, text padding around annotation labels
	ImPlotStyleVar_FitPadding         // ImVec2, additional fit padding as a percentage of the fit extents (e.g. ImVec2(0.1f,0.1f) adds 10% to the fit extents of X and Y)
	ImPlotStyleVar_PlotDefaultSize    // ImVec2, default size used when ImVec2(0,0) is passed to BeginPlot
	ImPlotStyleVar_PlotMinSize        // ImVec2, minimum size plot frame can be when shrunk
	ImPlotStyleVar_COUNT
)

const (
	ImPlotMarker_None     ImPlotMarker = iota - 1 // no marker
	ImPlotMarker_Circle                           // a circle marker (default)
	ImPlotMarker_Square                           // a square maker
	ImPlotMarker_Diamond                          // a diamond marker
	ImPlotMarker_Up                               // an upward-pointing triangle marker
	ImPlotMarker_Down                             // an downward-pointing triangle marker
	ImPlotMarker_Left                             // an leftward-pointing triangle marker
	ImPlotMarker_Right                            // an rightward-pointing triangle marker
	ImPlotMarker_Cross                            // a cross marker (not fillable)
	ImPlotMarker_Plus                             // a plus marker (not fillable)
	ImPlotMarker_Asterisk                         // a asterisk marker (not fillable)
	ImPlotMarker_COUNT
)

const (
	ImPlotColormap_Deep                   ImPlotColormap = 0  // a.k.a. seaborn deep             (qual=true,  n=10) (default)
	ImPlotColormap_Dark                   ImPlotColormap = 1  // a.k.a. matplotlib "Set1"        (qual=true,  n=9 )
	ImPlotColormap_Pastel                 ImPlotColormap = 2  // a.k.a. matplotlib "Pastel1"     (qual=true,  n=9 )
	ImPlotColormap_Paired                 ImPlotColormap = 3  // a.k.a. matplotlib "Paired"      (qual=true,  n=12)
	ImPlotColormap_Viridis                ImPlotColormap = 4  // a.k.a. matplotlib "viridis"     (qual=false, n=11)
	ImPlotColormap_Plasma                 ImPlotColormap = 5  // a.k.a. matplotlib "plasma"      (qual=false, n=11)
	ImPlotColormap_Hot                    ImPlotColormap = 6  // a.k.a. matplotlib/MATLAB "hot"  (qual=false, n=11)
	ImPlotColormap_Cool                   ImPlotColormap = 7  // a.k.a. matplotlib/MATLAB "cool" (qual=false, n=11)
	ImPlotColormap_Pink                   ImPlotColormap = 8  // a.k.a. matplotlib/MATLAB "pink" (qual=false, n=11)
	ImPlotColormap_Jet                    ImPlotColormap = 9  // a.k.a. MATLAB "jet"             (qual=false, n=11)
	ImPlotColormap_TwilightImPlotColormap                = 10 // a.k.a. matplotlib "twilight"    (qual=false, n=11)
	ImPlotColormap_RdBu                   ImPlotColormap = 11 // red/blue, Color Brewer          (qual=false, n=11)
	ImPlotColormap_BrBG                   ImPlotColormap = 12 // brown/blue-green, Color Brewer  (qual=false, n=11)
	ImPlotColormap_PiYG                   ImPlotColormap = 13 // pink/yellow-green, Color Brewer (qual=false, n=11)
	ImPlotColormap_SpectralImPlotColormap                = 14 // color spectrum, Color Brewer    (qual=false, n=11)
	ImPlotColormap_Greys                  ImPlotColormap = 15 // white/black                     (qual=false, n=2 )
)

const (
	ImPlotColormapScaleFlags_None     ImPlotColormapScaleFlags = 0      // default
	ImPlotColormapScaleFlags_NoLabel  ImPlotColormapScaleFlags = 1 << 0 // the colormap axis label will not be displayed
	ImPlotColormapScaleFlags_Opposite ImPlotColormapScaleFlags = 1 << 1 // render the colormap label and tick labels on the opposite side
	ImPlotColormapScaleFlags_Invert   ImPlotColormapScaleFlags = 1 << 2 // invert the colormap bar and axis scale (this only affects rendering; if you only want to reverse the scale mapping, make scale_min > scale_max)
)

const (
	ImGuiMouseButton_Left   ImGuiMouseButton = 0
	ImGuiMouseButton_Right  ImGuiMouseButton = 1
	ImGuiMouseButton_Middle ImGuiMouseButton = 2
	ImGuiMouseButton_COUNT  ImGuiMouseButton = 5
)
