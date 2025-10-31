package gsk

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cairo"
	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/glib/glib"
	"github.com/jinzhongmin/gg/glib/gobject"
	"github.com/jinzhongmin/gg/graphene"
	"github.com/jinzhongmin/gg/gtk/gdk"
	"github.com/jinzhongmin/gg/pango"
)

func cbool(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// #region CairoRenderer

type CairoRenderer struct{ Renderer }

func GTypeCairoRenderer() gobject.GType { return gsk_cairo_renderer_get_type.Fn()() }
func NewCairoRenderer() *Renderer       { return gsk_cairo_renderer_new.Fn()() }

// #endregion

// #region Path

type Path struct{}

func GTypePath() gobject.GType { return gsk_path_get_type.Fn()() }

func (p *Path) Ref() *Path                 { return gsk_path_ref.Fn()(p) }
func (p *Path) Unref()                     { gsk_path_unref.Fn()(p) }
func (p *Path) Print(string *glib.GString) { gsk_path_print.Fn()(p, string) }
func (p *Path) ToString() string           { return gsk_path_to_string.Fn()(p).String() }
func ParsePath(s string) *Path {
	ss := cc.NewString(s)
	defer ss.Free()
	return gsk_path_parse.Fn()(ss)
}
func (p *Path) ToCairo(cr *cairo.Context) { gsk_path_to_cairo.Fn()(p, cr) }
func (p *Path) IsEmpty() bool             { return gsk_path_is_empty.Fn()(p) != 0 }
func (p *Path) IsClosed() bool            { return gsk_path_is_closed.Fn()(p) != 0 }
func (p *Path) GetBounds() (bounds graphene.Rect, ok bool) {
	ok = gsk_path_get_bounds.Fn()(p, &bounds) != 0
	return
}
func (p *Path) GetStrokeBounds(stroke *Stroke) (bounds graphene.Rect, ok bool) {
	ok = gsk_path_get_stroke_bounds.Fn()(p, stroke, &bounds) != 0
	return
}
func (p *Path) InFill(point *graphene.Point, fillRule FillRule) bool {
	return gsk_path_in_fill.Fn()(p, point, fillRule) != 0
}
func (p *Path) GetStartPoint(result *PathPoint) bool {
	return gsk_path_get_start_point.Fn()(p, result) != 0
}
func (p *Path) GetEndPoint(result *PathPoint) bool {
	return gsk_path_get_end_point.Fn()(p, result) != 0
}
func (p *Path) GetClosestPoint(point *graphene.Point, threshold float64, result *PathPoint) (distance float64, ok bool) {
	ok = gsk_path_get_closest_point.Fn()(p, point, threshold, result, &distance) != 0
	return
}
func (p *Path) Foreach(flags PathForeachFlags, fn func(op PathOperation, pts []graphene.Point, weight float32) bool) bool {
	var cb UPtr
	if fn != nil {
		cb = cc.CbkRaw[func(PathOperation, *graphene.Point, uint64, float32, UPtr) int32](func(out, ins UPtr) {
			ps := unsafe.Slice((*UPtr)(ins), 5)
			op := *(*PathOperation)(ps[0])
			pts := *(**graphene.Point)(ps[1])
			npts := *(*uint64)(ps[2])
			weight := *(*float32)(ps[3])

			if fn(op, unsafe.Slice(pts, npts), weight) {
				*(*int32)(out) = 1
			} else {
				*(*int32)(out) = 2
			}
		})
		defer cc.CbkClose(cb)
	}

	return gsk_path_foreach.Fn()(p, flags, cb, nil) != 0
}

// #endregion

// #region PathBuilder

type PathBuilder struct{}

func GTypePathBuilder() gobject.GType    { return gsk_path_builder_get_type.Fn()() }
func NewPathBuilder() *PathBuilder       { return gsk_path_builder_new.Fn()() }
func (b *PathBuilder) Ref() *PathBuilder { return gsk_path_builder_ref.Fn()(b) }
func (b *PathBuilder) Unref()            { gsk_path_builder_unref.Fn()(b) }
func (b *PathBuilder) FreeToPath() *Path { return gsk_path_builder_free_to_path.Fn()(b) }
func (b *PathBuilder) ToPath() *Path     { return gsk_path_builder_to_path.Fn()(b) }
func (b *PathBuilder) GetCurrentPoint() *graphene.Point {
	return gsk_path_builder_get_current_point.Fn()(b)
}
func (b *PathBuilder) AddPath(path *Path)             { gsk_path_builder_add_path.Fn()(b, path) }
func (b *PathBuilder) AddReversePath(path *Path)      { gsk_path_builder_add_reverse_path.Fn()(b, path) }
func (b *PathBuilder) AddCairoPath(path *cairo.Path)  { gsk_path_builder_add_cairo_path.Fn()(b, path) }
func (b *PathBuilder) AddLayout(layout *pango.Layout) { gsk_path_builder_add_layout.Fn()(b, layout) }
func (b *PathBuilder) AddRect(rect *graphene.Rect)    { gsk_path_builder_add_rect.Fn()(b, rect) }
func (b *PathBuilder) AddRoundedRect(rect *RoundedRect) {
	gsk_path_builder_add_rounded_rect.Fn()(b, rect)
}
func (b *PathBuilder) AddCircle(center *graphene.Point, radius float64) {
	gsk_path_builder_add_circle.Fn()(b, center, radius)
}
func (b *PathBuilder) AddSegment(path *Path, start, end *PathPoint) {
	gsk_path_builder_add_segment.Fn()(b, path, start, end)
}
func (b *PathBuilder) MoveTo(x, y float64)    { gsk_path_builder_move_to.Fn()(b, x, y) }
func (b *PathBuilder) RelMoveTo(x, y float64) { gsk_path_builder_rel_move_to.Fn()(b, x, y) }
func (b *PathBuilder) LineTo(x, y float64)    { gsk_path_builder_line_to.Fn()(b, x, y) }
func (b *PathBuilder) RelLineTo(x, y float64) { gsk_path_builder_rel_line_to.Fn()(b, x, y) }
func (b *PathBuilder) QuadTo(x1, y1, x2, y2 float64) {
	gsk_path_builder_quad_to.Fn()(b, x1, y1, x2, y2)
}
func (b *PathBuilder) RelQuadTo(x1, y1, x2, y2 float64) {
	gsk_path_builder_rel_quad_to.Fn()(b, x1, y1, x2, y2)
}
func (b *PathBuilder) CubicTo(x1, y1, x2, y2, x3, y3 float64) {
	gsk_path_builder_cubic_to.Fn()(b, x1, y1, x2, y2, x3, y3)
}
func (b *PathBuilder) RelCubicTo(x1, y1, x2, y2, x3, y3 float64) {
	gsk_path_builder_rel_cubic_to.Fn()(b, x1, y1, x2, y2, x3, y3)
}
func (b *PathBuilder) ConicTo(x1, y1, x2, y2, weight float64) {
	gsk_path_builder_conic_to.Fn()(b, x1, y1, x2, y2, weight)
}
func (b *PathBuilder) RelConicTo(x1, y1, x2, y2, weight float64) {
	gsk_path_builder_rel_conic_to.Fn()(b, x1, y1, x2, y2, weight)
}
func (b *PathBuilder) ArcTo(x1, y1, x2, y2 float64) { gsk_path_builder_arc_to.Fn()(b, x1, y1, x2, y2) }
func (b *PathBuilder) RelArcTo(x1, y1, x2, y2 float64) {
	gsk_path_builder_rel_arc_to.Fn()(b, x1, y1, x2, y2)
}
func (b *PathBuilder) SvgArcTo(rx, ry, xAxisRotation float64, largeArc, positiveSweep bool, x, y float64) {
	gsk_path_builder_svg_arc_to.Fn()(b, rx, ry, xAxisRotation, cbool(largeArc), cbool(positiveSweep), x, y)
}
func (b *PathBuilder) RelSvgArcTo(rx, ry, xAxisRotation float64, largeArc, positiveSweep bool, x, y float64) {
	gsk_path_builder_rel_svg_arc_to.Fn()(b, rx, ry, xAxisRotation, cbool(largeArc), cbool(positiveSweep), x, y)
}
func (b *PathBuilder) HtmlArcTo(x1, y1, x2, y2, radius float64) {
	gsk_path_builder_html_arc_to.Fn()(b, x1, y1, x2, y2, radius)
}
func (b *PathBuilder) RelHtmlArcTo(x1, y1, x2, y2, radius float64) {
	gsk_path_builder_rel_html_arc_to.Fn()(b, x1, y1, x2, y2, radius)
}
func (b *PathBuilder) Close() { gsk_path_builder_close.Fn()(b) }

// #endregion

// #region PathMeasure

type PathMeasure struct{}

func GTypePathMeasure() gobject.GType        { return gsk_path_measure_get_type.Fn()() }
func NewPathMeasure(path *Path) *PathMeasure { return gsk_path_measure_new.Fn()(path) }
func NewPathMeasureWithTolerance(path *Path, tolerance float64) *PathMeasure {
	return gsk_path_measure_new_with_tolerance.Fn()(path, tolerance)
}
func (measure *PathMeasure) Ref() *PathMeasure { return gsk_path_measure_ref.Fn()(measure) }
func (measure *PathMeasure) Unref()            { gsk_path_measure_unref.Fn()(measure) }
func (measure *PathMeasure) GetPath() *Path    { return gsk_path_measure_get_path.Fn()(measure) }
func (measure *PathMeasure) GetTolerance() float64 {
	return gsk_path_measure_get_tolerance.Fn()(measure)
}
func (measure *PathMeasure) GetLength() float64 { return gsk_path_measure_get_length.Fn()(measure) }
func (measure *PathMeasure) GetPoint(distance float64, result *PathPoint) bool {
	return gsk_path_measure_get_point.Fn()(measure, distance, result) != 0
}

// #endregion

// #region PathPoint

type PathPoint struct{}

func GTypePathPoint() gobject.GType       { return gsk_path_point_get_type.Fn()() }
func (point *PathPoint) Copy() *PathPoint { return gsk_path_point_copy.Fn()(point) }
func (point *PathPoint) Free()            { gsk_path_point_free.Fn()(point) }
func (point *PathPoint) Equal(other *PathPoint) bool {
	return gsk_path_point_equal.Fn()(point, other) != 0
}
func (point *PathPoint) Compare(other *PathPoint) int32 {
	return gsk_path_point_compare.Fn()(point, other)
}
func (point *PathPoint) GetPosition(path *Path, position *graphene.Point) {
	gsk_path_point_get_position.Fn()(point, path, position)
}
func (point *PathPoint) GetTangent(path *Path, direction PathDirection, tangent *graphene.Vec2) {
	gsk_path_point_get_tangent.Fn()(point, path, direction, tangent)
}
func (point *PathPoint) GetRotation(path *Path, direction PathDirection) float64 {
	return gsk_path_point_get_rotation.Fn()(point, path, direction)
}
func (point *PathPoint) GetCurvature(path *Path, direction PathDirection, center *graphene.Point) float64 {
	return gsk_path_point_get_curvature.Fn()(point, path, direction, center)
}
func (point *PathPoint) GetDistance(measure *PathMeasure) float64 {
	return gsk_path_point_get_distance.Fn()(point, measure)
}

// #endregion

// #region Renderer

type Renderer struct{ gobject.GObjectCore }

func GTypeRenderer() gobject.GType { return gsk_renderer_get_type.Fn()() }
func NewRendererForSurface(surface *gdk.Surface) *Renderer {
	return gsk_renderer_new_for_surface.Fn()(surface)
}
func (renderer *Renderer) GetSurface() *gdk.Surface { return gsk_renderer_get_surface.Fn()(renderer) }
func (renderer *Renderer) Realize(surface *gdk.Surface) error {
	var err *glib.GError
	ok := gsk_renderer_realize.Fn()(renderer, surface, &err) != 0
	if ok {
		return nil
	}
	return err.TakeError()
}
func (renderer *Renderer) RealizeForDisplay(display *gdk.Display) error {
	var err *glib.GError
	ok := gsk_renderer_realize_for_display.Fn()(renderer, display, &err) != 0
	if ok {
		return nil
	}
	return err.TakeError()
}
func (renderer *Renderer) Unrealize()       { gsk_renderer_unrealize.Fn()(renderer) }
func (renderer *Renderer) IsRealized() bool { return gsk_renderer_is_realized.Fn()(renderer) != 0 }
func (renderer *Renderer) RenderTexture(root RenderNodeIface, viewport *graphene.Rect) *gdk.Texture {
	return gsk_renderer_render_texture.Fn()(renderer, GetRenderNodeIface(root), viewport)
}
func (renderer *Renderer) Render(root RenderNodeIface, region *cairo.Region) {
	gsk_renderer_render.Fn()(renderer, GetRenderNodeIface(root), region)
}

// #endregion

// #region RenderNode

type RenderNodeIface interface{ GetRenderNodeIface() *RenderNode }

func GetRenderNodeIface(ifac RenderNodeIface) *RenderNode {
	if anyptr(ifac) == nil {
		return nil
	}
	return ifac.GetRenderNodeIface()
}

func (provider *RenderNode) GetRenderNodeIface() *RenderNode { return provider }

type ColorStop struct {
	Offset float32
	Color  gdk.RGBA
}

type Shadow struct {
	Color          gdk.RGBA
	Dx, Dy, Radius float32
}

type ParseLocation struct {
	Bytes     uint64
	Chars     uint64
	Lines     uint64
	LineBytes uint64
	LineChars uint64
}

type RenderNode struct{ gobject.GTypeInstance }

func GTypeRenderNode() gobject.GType                 { return render_node_get_type.Fn()() }
func SerializationErrorQuark() uint32                { return serialization_error_quark.Fn()() }
func (node *RenderNode) Ref() *RenderNode            { return render_node_ref.Fn()(node) }
func (node *RenderNode) Unref()                      { render_node_unref.Fn()(node) }
func (node *RenderNode) GetNodeType() RenderNodeType { return render_node_get_node_type.Fn()(node) }
func (node *RenderNode) GetBounds() (bounds graphene.Rect) {
	render_node_get_bounds.Fn()(node, &bounds)
	return
}
func (node *RenderNode) GetOpaqueRect() (outOpaque graphene.Rect, ok bool) {
	ok = render_node_get_opaque_rect.Fn()(node, &outOpaque) != 0
	return
}
func (node *RenderNode) Draw(cr *cairo.Context)  { render_node_draw.Fn()(node, cr) }
func (node *RenderNode) Serialize() *glib.GBytes { return render_node_serialize.Fn()(node) }
func (node *RenderNode) WriteToFile(filename string) error {
	fn := cc.NewString(filename)
	defer fn.Free()
	var err *glib.GError
	ok := render_node_write_to_file.Fn()(node, fn, &err) != 0
	if ok {
		return nil
	}
	return err.TakeError()
}
func DeserializeRenderNode(bytes *glib.GBytes,
	errorFunc func(start, end *ParseLocation, err *glib.GError)) *RenderNode {
	ef := UPtr(nil)
	if errorFunc != nil {
		ef = cc.CbkRaw[func(start, end *ParseLocation, err *glib.GError, p UPtr)](func(out, ins UPtr) {
			is := unsafe.Slice((*UPtr)(ins), 4)
			start, end, err := (*ParseLocation)(is[0]), (*ParseLocation)(is[1]), (*glib.GError)(is[2])
			errorFunc(start, end, err)
		})
		defer cc.CbkClose(ef)
	}
	return render_node_deserialize.Fn()(bytes, ef, nil)
}

// DebugNode
type DebugNode struct{ RenderNode }

func GTypeDebugNode() gobject.GType { return debug_node_get_type.Fn()() }
func NewDebugNode(child RenderNodeIface, message string) *DebugNode {
	msg := cc.NewString(message)
	defer msg.Free()
	return debug_node_new.Fn()(GetRenderNodeIface(child), msg)
}
func (node *DebugNode) GetDebugChild() *RenderNode { return debug_node_get_child.Fn()(node) }
func (node *DebugNode) GetDebugMessage() string    { return debug_node_get_message.Fn()(node).String() }

// ColorNode
type ColorNode struct{ RenderNode }

func GTypeColorNode() gobject.GType { return color_node_get_type.Fn()() }
func NewColorNode(rgba *gdk.RGBA, bounds *graphene.Rect) *ColorNode {
	return color_node_new.Fn()(rgba, bounds)
}
func (node *ColorNode) GetColor() *gdk.RGBA { return color_node_get_color.Fn()(node) }

// TextureNode
type TextureNode struct{ RenderNode }

func GTypeTextureNode() gobject.GType { return texture_node_get_type.Fn()() }
func NewTextureNode(texture *gdk.Texture, bounds *graphene.Rect) *TextureNode {
	return texture_node_new.Fn()(texture, bounds)
}
func (node *TextureNode) GetTexture() *gdk.Texture { return texture_node_get_texture.Fn()(node) }

// TextureScaleNode
type TextureScaleNode struct{ RenderNode }

func GTypeTextureScaleNode() gobject.GType { return texture_scale_node_get_type.Fn()() }
func NewTextureScaleNode(texture *gdk.Texture, bounds *graphene.Rect, filter ScalingFilter) *TextureScaleNode {
	return texture_scale_node_new.Fn()(texture, bounds, filter)
}
func (node *TextureScaleNode) GetScaleTexture() *gdk.Texture {
	return texture_scale_node_get_texture.Fn()(node)
}
func (node *TextureScaleNode) GetScaleFilter() ScalingFilter {
	return texture_scale_node_get_filter.Fn()(node)
}

// LinearGradientNode
type LinearGradientNode struct{ RenderNode }

func GTypeLinearGradientNode() gobject.GType { return linear_gradient_node_get_type.Fn()() }
func NewLinearGradientNode(bounds *graphene.Rect, start, end *graphene.Point, colorStops []ColorStop) *LinearGradientNode {
	p, n := (*ColorStop)(nil), uint64(0)
	if n > 0 {
		p = &colorStops[0]
	}
	return linear_gradient_node_new.Fn()(bounds, start, end, p, n)
}
func (node *LinearGradientNode) GetLinearStart() *graphene.Point {
	return linear_gradient_node_get_start.Fn()(node)
}
func (node *LinearGradientNode) GetLinearEnd() *graphene.Point {
	return linear_gradient_node_get_end.Fn()(node)
}
func (node *LinearGradientNode) GetNLinearColorStops() uint64 {
	return linear_gradient_node_get_n_color_stops.Fn()(node)
}
func (node *LinearGradientNode) GetLinearColorStops() []ColorStop {
	n := uint64(0)
	cs := linear_gradient_node_get_color_stops.Fn()(node, &n)
	return unsafe.Slice(cs, n)
}

// RepeatingLinearGradientNode
type RepeatingLinearGradientNode struct{ RenderNode }

func GTypeRepeatingLinearGradientNode() gobject.GType {
	return repeating_linear_gradient_node_get_type.Fn()()
}
func NewRepeatingLinearGradientNode(bounds *graphene.Rect, start, end *graphene.Point, colorStops []ColorStop) *RepeatingLinearGradientNode {
	p, n := (*ColorStop)(nil), uint64(0)
	if n > 0 {
		p = &colorStops[0]
	}
	return repeating_linear_gradient_node_new.Fn()(bounds, start, end, p, n)
}

// ConicGradientNode
type ConicGradientNode struct{ RenderNode }

func GTypeConicGradientNode() gobject.GType { return conic_gradient_node_get_type.Fn()() }
func NewConicGradientNode(bounds *graphene.Rect, center *graphene.Point, rotation float64, colorStops []ColorStop) *ConicGradientNode {
	p, n := (*ColorStop)(nil), uint64(0)
	if n > 0 {
		p = &colorStops[0]
	}
	return conic_gradient_node_new.Fn()(bounds, center, rotation, p, n)
}
func (node *ConicGradientNode) GetConicCenter() *graphene.Point {
	return conic_gradient_node_get_center.Fn()(node)
}
func (node *ConicGradientNode) GetConicRotation() float64 {
	return conic_gradient_node_get_rotation.Fn()(node)
}
func (node *ConicGradientNode) GetConicAngle() float64 {
	return conic_gradient_node_get_angle.Fn()(node)
}
func (node *ConicGradientNode) GetNConicColorStops() uint64 {
	return conic_gradient_node_get_n_color_stops.Fn()(node)
}
func (node *ConicGradientNode) GetConicColorStops(nStops *uint64) []ColorStop {
	n := uint64(0)
	cs := conic_gradient_node_get_color_stops.Fn()(node, &n)
	return unsafe.Slice(cs, n)
}

// RadialGradientNode
type RadialGradientNode struct{ RenderNode }

func GTypeRadialGradientNode() gobject.GType { return radial_gradient_node_get_type.Fn()() }
func NewRadialGradientNode(bounds *graphene.Rect, center *graphene.Point, hradius, vradius, start, end float64, colorStops []ColorStop) *RadialGradientNode {
	p, n := (*ColorStop)(nil), uint64(0)
	if n > 0 {
		p = &colorStops[0]
	}
	return radial_gradient_node_new.Fn()(bounds, center, hradius, vradius, start, end, p, n)
}
func (node *RadialGradientNode) GetNRadialColorStops() uint64 {
	return radial_gradient_node_get_n_color_stops.Fn()(node)
}
func (node *RadialGradientNode) GetRadialColorStops(nStops *uint64) *ColorStop {
	return radial_gradient_node_get_color_stops.Fn()(node, nStops)
}
func (node *RadialGradientNode) GetRadialCenter() *graphene.Point {
	return radial_gradient_node_get_center.Fn()(node)
}
func (node *RadialGradientNode) GetRadialHRadius() float64 {
	return radial_gradient_node_get_hradius.Fn()(node)
}
func (node *RadialGradientNode) GetRadialVRadius() float64 {
	return radial_gradient_node_get_vradius.Fn()(node)
}
func (node *RadialGradientNode) GetRadialStart() float64 {
	return radial_gradient_node_get_start.Fn()(node)
}
func (node *RadialGradientNode) GetRadialEnd() float64 {
	return radial_gradient_node_get_end.Fn()(node)
}

// RepeatingRadialGradientNode
type RepeatingRadialGradientNode struct{ RenderNode }

func GTypeRepeatingRadialGradientNode() gobject.GType {
	return repeating_radial_gradient_node_get_type.Fn()()
}
func NewRepeatingRadialGradientNode(bounds *graphene.Rect, center *graphene.Point, hradius, vradius, start, end float64, colorStops []ColorStop) *RepeatingRadialGradientNode {
	p, n := (*ColorStop)(nil), uint64(0)
	if n > 0 {
		p = &colorStops[0]
	}
	return repeating_radial_gradient_node_new.Fn()(bounds, center, hradius, vradius, start, end, p, n)
}

// BorderNode
type BorderNode struct{ RenderNode }

func GTypeBorderNode() gobject.GType { return border_node_get_type.Fn()() }
func NewBorderNode(outline *RoundedRect, borderWidth [4]float64, borderColor [4]gdk.RGBA) *BorderNode {

	return border_node_new.Fn()(outline, &borderWidth[0], &borderColor[0])
}
func (node *BorderNode) GetBorderOutline() *RoundedRect { return border_node_get_outline.Fn()(node) }
func (node *BorderNode) GetBorderWidths() [4]float64 {
	p := border_node_get_widths.Fn()(node)
	return ([4]float64)(unsafe.Slice(p, 4))
}
func (node *BorderNode) GetBorderColors() [4]gdk.RGBA {
	p := border_node_get_colors.Fn()(node)
	return ([4]gdk.RGBA)(unsafe.Slice(p, 4))
}

// InsetShadowNode
type InsetShadowNode struct{ RenderNode }

func GTypeInsetShadowNode() gobject.GType { return inset_shadow_node_get_type.Fn()() }
func NewInsetShadowNode(outline *RoundedRect, color *gdk.RGBA, dx, dy, spread, blurRadius float64) *InsetShadowNode {
	return inset_shadow_node_new.Fn()(outline, color, dx, dy, spread, blurRadius)
}
func (node *InsetShadowNode) GetInsetShadowOutline() *RoundedRect {
	return inset_shadow_node_get_outline.Fn()(node)
}
func (node *InsetShadowNode) GetInsetShadowColor() *gdk.RGBA {
	return inset_shadow_node_get_color.Fn()(node)
}
func (node *InsetShadowNode) GetInsetShadowDx() float64 { return inset_shadow_node_get_dx.Fn()(node) }
func (node *InsetShadowNode) GetInsetShadowDy() float64 { return inset_shadow_node_get_dy.Fn()(node) }
func (node *InsetShadowNode) GetInsetShadowSpread() float64 {
	return inset_shadow_node_get_spread.Fn()(node)
}
func (node *InsetShadowNode) GetInsetShadowBlurRadius() float64 {
	return inset_shadow_node_get_blur_radius.Fn()(node)
}

// OutsetShadowNode
type OutsetShadowNode struct{ RenderNode }

func GTypeOutsetShadowNode() gobject.GType { return outset_shadow_node_get_type.Fn()() }
func NewOutsetShadowNode(outline *RoundedRect, color *gdk.RGBA, dx, dy, spread, blurRadius float64) *OutsetShadowNode {
	return outset_shadow_node_new.Fn()(outline, color, dx, dy, spread, blurRadius)
}
func (node *OutsetShadowNode) GetOutsetShadowOutline() *RoundedRect {
	return outset_shadow_node_get_outline.Fn()(node)
}
func (node *OutsetShadowNode) GetOutsetShadowColor() *gdk.RGBA {
	return outset_shadow_node_get_color.Fn()(node)
}
func (node *OutsetShadowNode) GetOutsetShadowDx() float64 {
	return outset_shadow_node_get_dx.Fn()(node)
}
func (node *OutsetShadowNode) GetOutsetShadowDy() float64 {
	return outset_shadow_node_get_dy.Fn()(node)
}
func (node *OutsetShadowNode) GetOutsetShadowSpread() float64 {
	return outset_shadow_node_get_spread.Fn()(node)
}
func (node *OutsetShadowNode) GetOutsetShadowBlurRadius() float64 {
	return outset_shadow_node_get_blur_radius.Fn()(node)
}

// CairoNode
type CairoNode struct{ RenderNode }

func GTypeCairoNode() gobject.GType                 { return cairo_node_get_type.Fn()() }
func NewCairoNode(bounds *graphene.Rect) *CairoNode { return cairo_node_new.Fn()(bounds) }
func (node *CairoNode) GetCairoDrawContext() *cairo.Context {
	return cairo_node_get_draw_context.Fn()(node)
}
func (node *CairoNode) GetCairoSurface() *cairo.Surface { return cairo_node_get_surface.Fn()(node) }

// ContainerNode
type ContainerNode struct{ RenderNode }

func GTypeContainerNode() gobject.GType { return container_node_get_type.Fn()() }
func NewContainerNode(children []RenderNodeIface) *ContainerNode {
	p, n := (**RenderNode)(nil), uint32(len(children))

	chs := make([]*RenderNode, n)
	for i := range chs {
		chs[i] = GetRenderNodeIface(children[i])
	}
	if n > 0 {
		p = &chs[0]
	}
	return container_node_new.Fn()(p, n)
}
func (node *ContainerNode) GetNContainerChildren() uint32 {
	return container_node_get_n_children.Fn()(node)
}
func (node *ContainerNode) GetContainerChild(idx uint32) *RenderNode {
	return container_node_get_child.Fn()(node, idx)
}

// TransformNode
type TransformNode struct{ RenderNode }

func GTypeTransformNode() gobject.GType { return transform_node_get_type.Fn()() }
func NewTransformNode(child RenderNodeIface, transform *Transform) *TransformNode {
	return transform_node_new.Fn()(GetRenderNodeIface(child), transform)
}
func (node *TransformNode) GetTransformChild() *RenderNode {
	return transform_node_get_child.Fn()(node)
}
func (node *TransformNode) GetTransform() *Transform { return transform_node_get_transform.Fn()(node) }

// OpacityNode
type OpacityNode struct{ RenderNode }

func GTypeOpacityNode() gobject.GType { return opacity_node_get_type.Fn()() }
func NewOpacityNode(child RenderNodeIface, opacity float64) *OpacityNode {
	return opacity_node_new.Fn()(GetRenderNodeIface(child), opacity)
}
func (node *OpacityNode) GetOpacityChild() *RenderNode { return opacity_node_get_child.Fn()(node) }
func (node *OpacityNode) GetOpacity() float64          { return opacity_node_get_opacity.Fn()(node) }

// ColorMatrixNode
type ColorMatrixNode struct{ RenderNode }

func GTypeColorMatrixNode() gobject.GType { return color_matrix_node_get_type.Fn()() }
func NewColorMatrixNode(child RenderNodeIface, colorMatrix *graphene.Matrix, colorOffset *graphene.Vec4) *ColorMatrixNode {
	return color_matrix_node_new.Fn()(GetRenderNodeIface(child), colorMatrix, colorOffset)
}
func (node *ColorMatrixNode) GetColorMatrixChild() *RenderNode {
	return color_matrix_node_get_child.Fn()(node)
}
func (node *ColorMatrixNode) GetColorMatrix() *graphene.Matrix {
	return color_matrix_node_get_color_matrix.Fn()(node)
}
func (node *ColorMatrixNode) GetColorOffset() *graphene.Vec4 {
	return color_matrix_node_get_color_offset.Fn()(node)
}

// RepeatNode
type RepeatNode struct{ RenderNode }

func GTypeRepeatNode() gobject.GType { return repeat_node_get_type.Fn()() }
func NewRepeatNode(bounds *graphene.Rect, child RenderNodeIface, childBounds *graphene.Rect) *RepeatNode {
	return repeat_node_new.Fn()(bounds, GetRenderNodeIface(child), childBounds)
}
func (node *RepeatNode) GetRepeatChild() *RenderNode { return repeat_node_get_child.Fn()(node) }
func (node *RepeatNode) GetRepeatChildBounds() *graphene.Rect {
	return repeat_node_get_child_bounds.Fn()(node)
}

// ClipNode
type ClipNode struct{ RenderNode }

func GTypeClipNode() gobject.GType { return clip_node_get_type.Fn()() }
func NewClipNode(child RenderNodeIface, clip *graphene.Rect) *ClipNode {
	return clip_node_new.Fn()(GetRenderNodeIface(child), clip)
}
func (node *ClipNode) GetClipChild() *RenderNode { return clip_node_get_child.Fn()(node) }
func (node *ClipNode) GetClip() *graphene.Rect   { return clip_node_get_clip.Fn()(node) }

// RoundedClipNode
type RoundedClipNode struct{ RenderNode }

func GTypeRoundedClipNode() gobject.GType { return rounded_clip_node_get_type.Fn()() }
func NewRoundedClipNode(child RenderNodeIface, clip *RoundedRect) *RoundedClipNode {
	return rounded_clip_node_new.Fn()(GetRenderNodeIface(child), clip)
}
func (node *RoundedClipNode) GetRoundedClipChild() *RenderNode {
	return rounded_clip_node_get_child.Fn()(node)
}
func (node *RoundedClipNode) GetRoundedClip() *RoundedRect {
	return rounded_clip_node_get_clip.Fn()(node)
}

// FillNode
type FillNode struct{ RenderNode }

func GTypeFillNode() gobject.GType { return fill_node_get_type.Fn()() }
func NewFillNode(child RenderNodeIface, path *Path, fillRule FillRule) *FillNode {
	return fill_node_new.Fn()(GetRenderNodeIface(child), path, fillRule)
}
func (node *FillNode) GetFillChild() *RenderNode { return fill_node_get_child.Fn()(node) }
func (node *FillNode) GetFillPath() *Path        { return fill_node_get_path.Fn()(node) }
func (node *FillNode) GetFillRule() FillRule     { return fill_node_get_fill_rule.Fn()(node) }

// StrokeNode
type StrokeNode struct{ RenderNode }

func GTypeStrokeNode() gobject.GType { return stroke_node_get_type.Fn()() }
func NewStrokeNode(child RenderNodeIface, path *Path, stroke *Stroke) *StrokeNode {
	return stroke_node_new.Fn()(GetRenderNodeIface(child), path, stroke)
}
func (node *StrokeNode) GetStrokeChild() *RenderNode { return stroke_node_get_child.Fn()(node) }
func (node *StrokeNode) GetStrokePath() *Path        { return stroke_node_get_path.Fn()(node) }
func (node *StrokeNode) GetStroke() *Stroke          { return stroke_node_get_stroke.Fn()(node) }

// ShadowNode
type ShadowNode struct{ RenderNode }

func GTypeShadowNode() gobject.GType { return shadow_node_get_type.Fn()() }
func NewShadowNode(child RenderNodeIface, shadows []Shadow) *ShadowNode {
	p, n := (*Shadow)(nil), uint64(0)
	if n > 0 {
		p = &shadows[0]
	}
	return shadow_node_new.Fn()(GetRenderNodeIface(child), p, n)
}
func (node *ShadowNode) GetShadowChild() *RenderNode { return shadow_node_get_child.Fn()(node) }
func (node *ShadowNode) GetShadow(i uint64) *Shadow  { return shadow_node_get_shadow.Fn()(node, i) }
func (node *ShadowNode) GetNShadows() uint64         { return shadow_node_get_n_shadows.Fn()(node) }

// BlendNode
type BlendNode struct{ RenderNode }

func GTypeBlendNode() gobject.GType { return blend_node_get_type.Fn()() }
func NewBlendNode(bottom, top RenderNodeIface, blendMode BlendMode) *BlendNode {
	return blend_node_new.Fn()(GetRenderNodeIface(bottom), GetRenderNodeIface(top), blendMode)
}
func (node *BlendNode) GetBottomChild() *RenderNode { return blend_node_get_bottom_child.Fn()(node) }
func (node *BlendNode) GetTopChild() *RenderNode    { return blend_node_get_top_child.Fn()(node) }
func (node *BlendNode) GetBlendMode() BlendMode     { return blend_node_get_blend_mode.Fn()(node) }

// CrossFadeNode
type CrossFadeNode struct{ RenderNode }

func GTypeCrossFadeNode() gobject.GType { return cross_fade_node_get_type.Fn()() }
func NewCrossFadeNode(start, end RenderNodeIface, progress float64) *CrossFadeNode {
	return cross_fade_node_new.Fn()(GetRenderNodeIface(start), GetRenderNodeIface(end), progress)
}
func (node *CrossFadeNode) GetStartChild() *RenderNode {
	return cross_fade_node_get_start_child.Fn()(node)
}
func (node *CrossFadeNode) GetEndChild() *RenderNode { return cross_fade_node_get_end_child.Fn()(node) }
func (node *CrossFadeNode) GetProgress() float64     { return cross_fade_node_get_progress.Fn()(node) }

// TextNode
type TextNode struct{ RenderNode }

func GTypeTextNode() gobject.GType { return text_node_get_type.Fn()() }
func NewTextNode(font *pango.Font, glyphs *pango.GlyphString, color *gdk.RGBA, offset *graphene.Point) *TextNode {
	return text_node_new.Fn()(font, glyphs, color, offset)
}
func (node *TextNode) GetTextFont() *pango.Font { return text_node_get_font.Fn()(node) }
func (node *TextNode) HasColorGlyphs() bool     { return text_node_has_color_glyphs.Fn()(node) != 0 }
func (node *TextNode) GetNumGlyphs() uint32     { return text_node_get_num_glyphs.Fn()(node) }
func (node *TextNode) GetGlyphs() []pango.GlyphInfo {
	n := uint32(0)
	p := text_node_get_glyphs.Fn()(node, &n)
	return unsafe.Slice(p, n)
}
func (node *TextNode) GetTextColor() *gdk.RGBA        { return text_node_get_color.Fn()(node) }
func (node *TextNode) GetTextOffset() *graphene.Point { return text_node_get_offset.Fn()(node) }

// BlurNode
type BlurNode struct{ RenderNode }

func GTypeBlurNode() gobject.GType { return blur_node_get_type.Fn()() }
func NewBlurNode(child RenderNodeIface, radius float64) *BlurNode {
	return blur_node_new.Fn()(GetRenderNodeIface(child), radius)
}
func (node *BlurNode) GetBlurChild() *RenderNode { return blur_node_get_child.Fn()(node) }
func (node *BlurNode) GetBlurRadius() float64    { return blur_node_get_radius.Fn()(node) }

// MaskNode
type MaskNode struct{ RenderNode }

func GTypeMaskNode() gobject.GType { return mask_node_get_type.Fn()() }
func NewMaskNode(source, mask RenderNodeIface, maskMode MaskMode) *MaskNode {
	return mask_node_new.Fn()(GetRenderNodeIface(source), GetRenderNodeIface(mask), maskMode)
}
func (node *MaskNode) GetMaskSource() *RenderNode { return mask_node_get_source.Fn()(node) }
func (node *MaskNode) GetMask() *RenderNode       { return mask_node_get_mask.Fn()(node) }
func (node *MaskNode) GetMaskMode() MaskMode      { return mask_node_get_mask_mode.Fn()(node) }

// GLShaderNode (已过时，按要求生成但标记)

type GLShader struct{}

type GLShaderNode struct{ RenderNode }

func GTypeGLShaderNode() gobject.GType { return gl_shader_node_get_type.Fn()() }
func NewGLShaderNode(shader *GLShader, bounds *graphene.Rect, args *glib.GBytes, children []RenderNodeIface) *GLShaderNode {
	p, n := (**RenderNode)(nil), uint32(len(children))
	chs := make([]*RenderNode, n)
	for i := range chs {
		chs[i] = GetRenderNodeIface(children[i])
	}
	if n > 0 {
		p = &chs[0]
	}
	return gl_shader_node_new.Fn()(shader, bounds, args, p, n)
}
func (node *GLShaderNode) GetNGLShaderChildren() uint32 {
	return gl_shader_node_get_n_children.Fn()(node)
}
func (node *GLShaderNode) GetGLShaderChild(idx uint32) *RenderNode {
	return gl_shader_node_get_child.Fn()(node, idx)
}
func (node *GLShaderNode) GetGLShaderArgs() *glib.GBytes { return gl_shader_node_get_args.Fn()(node) }
func (node *GLShaderNode) GetGLShader() *GLShader        { return gl_shader_node_get_shader.Fn()(node) }

// SubsurfaceNode
type SubsurfaceNode struct{ RenderNode }

func GTypeSubsurfaceNode() gobject.GType { return subsurface_node_get_type.Fn()() }
func NewSubsurfaceNode(child RenderNodeIface, subsurface UPtr) *SubsurfaceNode {
	return subsurface_node_new.Fn()(GetRenderNodeIface(child), subsurface)
}
func (node *SubsurfaceNode) GetSubsurfaceChild() *RenderNode {
	return subsurface_node_get_child.Fn()(node)
}
func (node *SubsurfaceNode) GetSubsurface() UPtr { return subsurface_node_get_subsurface.Fn()(node) }

// Value functions

func (node *RenderNode) SetTo(value *gobject.GValue) {
	value_set_render_node.Fn()(value, GetRenderNodeIface(node))
}
func (node *RenderNode) TakeTo(value *gobject.GValue) {
	value_take_render_node.Fn()(value, GetRenderNodeIface(node))
}

func ValueSetRenderNode(value *gobject.GValue, node RenderNodeIface) {
	value_set_render_node.Fn()(value, GetRenderNodeIface(node))
}
func ValueTakeRenderNode(value *gobject.GValue, node RenderNodeIface) {
	value_take_render_node.Fn()(value, GetRenderNodeIface(node))
}

func ValueGetRenderNode(value *gobject.GValue) *RenderNode { return value_get_render_node.Fn()(value) }
func ValueDupRenderNode(value *gobject.GValue) *RenderNode { return value_dup_render_node.Fn()(value) }

// #endregion

// #region RoundedRect

type RoundedRect struct {
	Bounds graphene.Rect
	Corner [4]graphene.Size
}

func (rr *RoundedRect) Init(bounds *graphene.Rect, topLeft, topRight, bottomRight, bottomLeft *graphene.Size) *RoundedRect {
	return rounded_rect_init.Fn()(rr, bounds, topLeft, topRight, bottomRight, bottomLeft)
}
func (rr *RoundedRect) InitCopy(src *RoundedRect) *RoundedRect {
	return rounded_rect_init_copy.Fn()(rr, src)
}
func (rr *RoundedRect) InitFromRect(bounds *graphene.Rect, radius float64) *RoundedRect {
	return rounded_rect_init_from_rect.Fn()(rr, bounds, radius)
}
func (rr *RoundedRect) Normalize() *RoundedRect { return rounded_rect_normalize.Fn()(rr) }
func (rr *RoundedRect) Offset(dx, dy float64) *RoundedRect {
	return rounded_rect_offset.Fn()(rr, dx, dy)
}
func (rr *RoundedRect) Shrink(top, right, bottom, left float64) *RoundedRect {
	return rounded_rect_shrink.Fn()(rr, top, right, bottom, left)
}
func (rr *RoundedRect) IsRectilinear() bool { return rounded_rect_is_rectilinear.Fn()(rr) != 0 }
func (rr *RoundedRect) ContainsPoint(point *graphene.Point) bool {
	return rounded_rect_contains_point.Fn()(rr, point) != 0
}
func (rr *RoundedRect) ContainsRect(rect *graphene.Rect) bool {
	return rounded_rect_contains_rect.Fn()(rr, rect) != 0
}
func (rr *RoundedRect) IntersectsRect(rect *graphene.Rect) bool {
	return rounded_rect_intersects_rect.Fn()(rr, rect) != 0
}

// #endregion

// #region Stroke

type Stroke struct{}

func GTypeStroke() gobject.GType          { return stroke_get_type.Fn()() }
func NewStroke(lineWidth float64) *Stroke { return stroke_new.Fn()(lineWidth) }
func (s *Stroke) Copy() *Stroke           { return stroke_copy.Fn()(s) }
func (s *Stroke) Free()                   { stroke_free.Fn()(s) }
func (s *Stroke) Equal(other *Stroke) bool {
	return stroke_equal.Fn()(s, other) != 0
}
func (s *Stroke) SetLineWidth(lineWidth float64) { stroke_set_line_width.Fn()(s, lineWidth) }
func (s *Stroke) GetLineWidth() float64          { return stroke_get_line_width.Fn()(s) }
func (s *Stroke) SetLineCap(lineCap LineCap)     { stroke_set_line_cap.Fn()(s, lineCap) }
func (s *Stroke) GetLineCap() LineCap            { return stroke_get_line_cap.Fn()(s) }
func (s *Stroke) SetLineJoin(lineJoin LineJoin)  { stroke_set_line_join.Fn()(s, lineJoin) }
func (s *Stroke) GetLineJoin() LineJoin          { return stroke_get_line_join.Fn()(s) }
func (s *Stroke) SetMiterLimit(limit float64)    { stroke_set_miter_limit.Fn()(s, limit) }
func (s *Stroke) GetMiterLimit() float64         { return stroke_get_miter_limit.Fn()(s) }
func (s *Stroke) SetDash(dash []float64) {
	p, n := (*float64)(nil), uint64(len(dash))
	if n > 0 {
		p = &dash[0]
	}
	stroke_set_dash.Fn()(s, p, n)
}
func (s *Stroke) GetDash() (dash []float64) {
	nDash := uint64(0)
	ptr := stroke_get_dash.Fn()(s, &nDash)
	if ptr != nil && nDash > 0 {
		dash = unsafe.Slice(ptr, nDash)
	}
	return
}
func (s *Stroke) SetDashOffset(offset float64) { stroke_set_dash_offset.Fn()(s, offset) }
func (s *Stroke) GetDashOffset() float64       { return stroke_get_dash_offset.Fn()(s) }
func (s *Stroke) ToCairo(cr *cairo.Context)    { stroke_to_cairo.Fn()(s, cr) }

// #endregion

// #region Transform

type Transform struct{}

func GTypeGskTransform() gobject.GType { return gsk_transform_get_type.Fn()() }

func (t *Transform) Ref() *Transform          { return gsk_transform_ref.Fn()(t) }
func (t *Transform) Unref()                   { gsk_transform_unref.Fn()(t) }
func (t *Transform) Print(gstr *glib.GString) { gsk_transform_print.Fn()(t, gstr) }
func (t *Transform) ToString() string         { return gsk_transform_to_string.Fn()(t).String() }
func NewTransformFromParse(str string) (tran *Transform, ok bool) {
	cstr := cc.NewString(str)
	defer cstr.Free()
	var out *Transform
	ok = gsk_transform_parse.Fn()(cstr, &out) != 0
	return out, ok
}
func (t *Transform) ToMatrix(out *graphene.Matrix) { gsk_transform_to_matrix.Fn()(t, out) }
func (t *Transform) To2D() (xx, yx, xy, yy, dx, dy float32) {
	gsk_transform_to_2d.Fn()(t, &xx, &yx, &xy, &yy, &dx, &dy)
	return
}
func (t *Transform) To2DComponents() (skew_x, skew_y, scale_x, scale_y, angle, dx, dy float32) {
	gsk_transform_to_2d_components.Fn()(t, &skew_x, &skew_y, &scale_x, &scale_y, &angle, &dx, &dy)
	return
}
func (t *Transform) ToAffine() (scale_x, scale_y, dx, dy float32) {
	gsk_transform_to_affine.Fn()(t, &scale_x, &scale_y, &dx, &dy)
	return
}
func (t *Transform) ToTranslate() (dx, dy float32) {
	gsk_transform_to_translate.Fn()(t, &dx, &dy)
	return
}
func (t *Transform) GetCategory() TransformCategory { return gsk_transform_get_category.Fn()(t) }
func (t *Transform) Equal(second *Transform) bool   { return gsk_transform_equal.Fn()(t, second) != 0 }
func NewTransform() *Transform                      { return gsk_transform_new.Fn()() }
func (t *Transform) Transform(other *Transform) *Transform {
	return gsk_transform_transform.Fn()(t, other)
}
func (t *Transform) Invert() *Transform { return gsk_transform_invert.Fn()(t) }
func (t *Transform) Matrix(matrix *graphene.Matrix) *Transform {
	return gsk_transform_matrix.Fn()(t, matrix)
}
func (t *Transform) Translate(point *graphene.Point) *Transform {
	return gsk_transform_translate.Fn()(t, point)
}
func (t *Transform) Translate3D(point *graphene.Point3D) *Transform {
	return gsk_transform_translate_3d.Fn()(t, point)
}
func (t *Transform) Skew(skew_x, skew_y float32) *Transform {
	return gsk_transform_skew.Fn()(t, skew_x, skew_y)
}
func (t *Transform) Rotate(angle float32) *Transform { return gsk_transform_rotate.Fn()(t, angle) }
func (t *Transform) Rotate3D(angle float32, axis *graphene.Vec3) *Transform {
	return gsk_transform_rotate_3d.Fn()(t, angle, axis)
}
func (t *Transform) Scale(factor_x, factor_y float32) *Transform {
	return gsk_transform_scale.Fn()(t, factor_x, factor_y)
}
func (t *Transform) Scale3D(factor_x, factor_y, factor_z float32) *Transform {
	return gsk_transform_scale_3d.Fn()(t, factor_x, factor_y, factor_z)
}
func (t *Transform) Perspective(depth float32) *Transform {
	return gsk_transform_perspective.Fn()(t, depth)
}
func (t *Transform) TransformBounds(rect *graphene.Rect, out_rect *graphene.Rect) {
	gsk_transform_transform_bounds.Fn()(t, rect, out_rect)
}
func (t *Transform) TransformPoint(point *graphene.Point, out_point *graphene.Point) {
	gsk_transform_transform_point.Fn()(t, point, out_point)
}

// #endregion

type GLRenderer struct{ Renderer }
type VulkanRenderer struct{ Renderer }

func GTypeGLRenderer() gobject.GType { return gl_renderer_get_type.Fn()() }
func NewGLRenderer() *GLRenderer     { return gl_renderer_new.Fn()() }

func GTypeVulkanRenderer() gobject.GType { return vulkan_renderer_get_type.Fn()() }
func NewVulkanRenderer() *VulkanRenderer { return vulkan_renderer_new.Fn()() }
