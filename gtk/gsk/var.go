package gsk

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cairo"
	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/cc/dl"
	"github.com/jinzhongmin/gg/glib/glib"
	"github.com/jinzhongmin/gg/glib/gobject"
	"github.com/jinzhongmin/gg/graphene"
	"github.com/jinzhongmin/gg/gtk/gdk"
	"github.com/jinzhongmin/gg/pango"
)

type UPtr = unsafe.Pointer
type IPtr = uintptr

// func vfn(name string, fptr interface{})               { cc.RFunc(fptr, name) }
func anyptr(a interface{}) UPtr { return (*(*[2]UPtr)(UPtr(&a)))[1] }

func init() { dl.Open("libgtk-4*") }

var (
	// #region CairoRenderer
	gsk_cairo_renderer_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_cairo_renderer_get_type"}
	gsk_cairo_renderer_new      = cc.DlFunc[func() *Renderer, *Renderer]{Name: "gsk_cairo_renderer_new"}
	// #endregion

	// #region Path
	gsk_path_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_path_get_type"}
	gsk_path_ref               = cc.DlFunc[func(*Path) *Path, *Path]{Name: "gsk_path_ref"}
	gsk_path_unref             = cc.DlFunc[func(*Path), cc.Void]{Name: "gsk_path_unref"}
	gsk_path_print             = cc.DlFunc[func(*Path, *glib.GString), cc.Void]{Name: "gsk_path_print"}
	gsk_path_to_string         = cc.DlFunc[func(*Path) cc.String, cc.String]{Name: "gsk_path_to_string"}
	gsk_path_parse             = cc.DlFunc[func(cc.String) *Path, *Path]{Name: "gsk_path_parse"}
	gsk_path_to_cairo          = cc.DlFunc[func(*Path, *cairo.Context), cc.Void]{Name: "gsk_path_to_cairo"}
	gsk_path_is_empty          = cc.DlFunc[func(*Path) int32, int32]{Name: "gsk_path_is_empty"}
	gsk_path_is_closed         = cc.DlFunc[func(*Path) int32, int32]{Name: "gsk_path_is_closed"}
	gsk_path_get_bounds        = cc.DlFunc[func(*Path, *graphene.Rect) int32, int32]{Name: "gsk_path_get_bounds"}
	gsk_path_get_stroke_bounds = cc.DlFunc[func(*Path, *Stroke, *graphene.Rect) int32, int32]{Name: "gsk_path_get_stroke_bounds"}
	gsk_path_in_fill           = cc.DlFunc[func(*Path, *graphene.Point, FillRule) int32, int32]{Name: "gsk_path_in_fill"}
	gsk_path_get_start_point   = cc.DlFunc[func(*Path, *PathPoint) int32, int32]{Name: "gsk_path_get_start_point"}
	gsk_path_get_end_point     = cc.DlFunc[func(*Path, *PathPoint) int32, int32]{Name: "gsk_path_get_end_point"}
	gsk_path_get_closest_point = cc.DlFunc[func(*Path, *graphene.Point, float64, *PathPoint, *float64) int32, int32]{Name: "gsk_path_get_closest_point"}
	gsk_path_foreach           = cc.DlFunc[func(*Path, PathForeachFlags, UPtr, UPtr) int32, int32]{Name: "gsk_path_foreach"}
	// #endregion

	// #region PathBuilder
	gsk_path_builder_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_path_builder_get_type"}
	gsk_path_builder_new               = cc.DlFunc[func() *PathBuilder, *PathBuilder]{Name: "gsk_path_builder_new"}
	gsk_path_builder_ref               = cc.DlFunc[func(*PathBuilder) *PathBuilder, *PathBuilder]{Name: "gsk_path_builder_ref"}
	gsk_path_builder_unref             = cc.DlFunc[func(*PathBuilder), cc.Void]{Name: "gsk_path_builder_unref"}
	gsk_path_builder_free_to_path      = cc.DlFunc[func(*PathBuilder) *Path, *Path]{Name: "gsk_path_builder_free_to_path"}
	gsk_path_builder_to_path           = cc.DlFunc[func(*PathBuilder) *Path, *Path]{Name: "gsk_path_builder_to_path"}
	gsk_path_builder_get_current_point = cc.DlFunc[func(*PathBuilder) *graphene.Point, *graphene.Point]{Name: "gsk_path_builder_get_current_point"}
	gsk_path_builder_add_path          = cc.DlFunc[func(*PathBuilder, *Path), cc.Void]{Name: "gsk_path_builder_add_path"}
	gsk_path_builder_add_reverse_path  = cc.DlFunc[func(*PathBuilder, *Path), cc.Void]{Name: "gsk_path_builder_add_reverse_path"}
	gsk_path_builder_add_cairo_path    = cc.DlFunc[func(*PathBuilder, *cairo.Path), cc.Void]{Name: "gsk_path_builder_add_cairo_path"}
	gsk_path_builder_add_layout        = cc.DlFunc[func(*PathBuilder, *pango.Layout), cc.Void]{Name: "gsk_path_builder_add_layout"}
	gsk_path_builder_add_rect          = cc.DlFunc[func(*PathBuilder, *graphene.Rect), cc.Void]{Name: "gsk_path_builder_add_rect"}
	gsk_path_builder_add_rounded_rect  = cc.DlFunc[func(*PathBuilder, *RoundedRect), cc.Void]{Name: "gsk_path_builder_add_rounded_rect"}
	gsk_path_builder_add_circle        = cc.DlFunc[func(*PathBuilder, *graphene.Point, float64), cc.Void]{Name: "gsk_path_builder_add_circle"}
	gsk_path_builder_add_segment       = cc.DlFunc[func(*PathBuilder, *Path, *PathPoint, *PathPoint), cc.Void]{Name: "gsk_path_builder_add_segment"}
	gsk_path_builder_move_to           = cc.DlFunc[func(*PathBuilder, float64, float64), cc.Void]{Name: "gsk_path_builder_move_to"}
	gsk_path_builder_rel_move_to       = cc.DlFunc[func(*PathBuilder, float64, float64), cc.Void]{Name: "gsk_path_builder_rel_move_to"}
	gsk_path_builder_line_to           = cc.DlFunc[func(*PathBuilder, float64, float64), cc.Void]{Name: "gsk_path_builder_line_to"}
	gsk_path_builder_rel_line_to       = cc.DlFunc[func(*PathBuilder, float64, float64), cc.Void]{Name: "gsk_path_builder_rel_line_to"}
	gsk_path_builder_quad_to           = cc.DlFunc[func(*PathBuilder, float64, float64, float64, float64), cc.Void]{Name: "gsk_path_builder_quad_to"}
	gsk_path_builder_rel_quad_to       = cc.DlFunc[func(*PathBuilder, float64, float64, float64, float64), cc.Void]{Name: "gsk_path_builder_rel_quad_to"}
	gsk_path_builder_cubic_to          = cc.DlFunc[func(*PathBuilder, float64, float64, float64, float64, float64, float64), cc.Void]{Name: "gsk_path_builder_cubic_to"}
	gsk_path_builder_rel_cubic_to      = cc.DlFunc[func(*PathBuilder, float64, float64, float64, float64, float64, float64), cc.Void]{Name: "gsk_path_builder_rel_cubic_to"}
	gsk_path_builder_conic_to          = cc.DlFunc[func(*PathBuilder, float64, float64, float64, float64, float64), cc.Void]{Name: "gsk_path_builder_conic_to"}
	gsk_path_builder_rel_conic_to      = cc.DlFunc[func(*PathBuilder, float64, float64, float64, float64, float64), cc.Void]{Name: "gsk_path_builder_rel_conic_to"}
	gsk_path_builder_arc_to            = cc.DlFunc[func(*PathBuilder, float64, float64, float64, float64), cc.Void]{Name: "gsk_path_builder_arc_to"}
	gsk_path_builder_rel_arc_to        = cc.DlFunc[func(*PathBuilder, float64, float64, float64, float64), cc.Void]{Name: "gsk_path_builder_rel_arc_to"}
	gsk_path_builder_svg_arc_to        = cc.DlFunc[func(*PathBuilder, float64, float64, float64, int32, int32, float64, float64), cc.Void]{Name: "gsk_path_builder_svg_arc_to"}
	gsk_path_builder_rel_svg_arc_to    = cc.DlFunc[func(*PathBuilder, float64, float64, float64, int32, int32, float64, float64), cc.Void]{Name: "gsk_path_builder_rel_svg_arc_to"}
	gsk_path_builder_html_arc_to       = cc.DlFunc[func(*PathBuilder, float64, float64, float64, float64, float64), cc.Void]{Name: "gsk_path_builder_html_arc_to"}
	gsk_path_builder_rel_html_arc_to   = cc.DlFunc[func(*PathBuilder, float64, float64, float64, float64, float64), cc.Void]{Name: "gsk_path_builder_rel_html_arc_to"}
	gsk_path_builder_close             = cc.DlFunc[func(*PathBuilder), cc.Void]{Name: "gsk_path_builder_close"}
	// #endregion

	// #region PathMeasure
	gsk_path_measure_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_path_measure_get_type"}
	gsk_path_measure_new                = cc.DlFunc[func(*Path) *PathMeasure, *PathMeasure]{Name: "gsk_path_measure_new"}
	gsk_path_measure_new_with_tolerance = cc.DlFunc[func(*Path, float64) *PathMeasure, *PathMeasure]{Name: "gsk_path_measure_new_with_tolerance"}
	gsk_path_measure_ref                = cc.DlFunc[func(*PathMeasure) *PathMeasure, *PathMeasure]{Name: "gsk_path_measure_ref"}
	gsk_path_measure_unref              = cc.DlFunc[func(*PathMeasure), cc.Void]{Name: "gsk_path_measure_unref"}
	gsk_path_measure_get_path           = cc.DlFunc[func(*PathMeasure) *Path, *Path]{Name: "gsk_path_measure_get_path"}
	gsk_path_measure_get_tolerance      = cc.DlFunc[func(*PathMeasure) float64, float64]{Name: "gsk_path_measure_get_tolerance"}
	gsk_path_measure_get_length         = cc.DlFunc[func(*PathMeasure) float64, float64]{Name: "gsk_path_measure_get_length"}
	gsk_path_measure_get_point          = cc.DlFunc[func(*PathMeasure, float64, *PathPoint) int32, int32]{Name: "gsk_path_measure_get_point"}
	// #endregion

	// #region PathPoint
	gsk_path_point_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_path_point_get_type"}
	gsk_path_point_copy          = cc.DlFunc[func(*PathPoint) *PathPoint, *PathPoint]{Name: "gsk_path_point_copy"}
	gsk_path_point_free          = cc.DlFunc[func(*PathPoint), cc.Void]{Name: "gsk_path_point_free"}
	gsk_path_point_equal         = cc.DlFunc[func(*PathPoint, *PathPoint) int32, int32]{Name: "gsk_path_point_equal"}
	gsk_path_point_compare       = cc.DlFunc[func(*PathPoint, *PathPoint) int32, int32]{Name: "gsk_path_point_compare"}
	gsk_path_point_get_position  = cc.DlFunc[func(*PathPoint, *Path, *graphene.Point), cc.Void]{Name: "gsk_path_point_get_position"}
	gsk_path_point_get_tangent   = cc.DlFunc[func(*PathPoint, *Path, PathDirection, *graphene.Vec2), cc.Void]{Name: "gsk_path_point_get_tangent"}
	gsk_path_point_get_rotation  = cc.DlFunc[func(*PathPoint, *Path, PathDirection) float64, float64]{Name: "gsk_path_point_get_rotation"}
	gsk_path_point_get_curvature = cc.DlFunc[func(*PathPoint, *Path, PathDirection, *graphene.Point) float64, float64]{Name: "gsk_path_point_get_curvature"}
	gsk_path_point_get_distance  = cc.DlFunc[func(*PathPoint, *PathMeasure) float64, float64]{Name: "gsk_path_point_get_distance"}
	// #endregion

	// #region Renderer
	gsk_renderer_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_renderer_get_type"}
	gsk_renderer_new_for_surface     = cc.DlFunc[func(*gdk.Surface) *Renderer, *Renderer]{Name: "gsk_renderer_new_for_surface"}
	gsk_renderer_get_surface         = cc.DlFunc[func(*Renderer) *gdk.Surface, *gdk.Surface]{Name: "gsk_renderer_get_surface"}
	gsk_renderer_realize             = cc.DlFunc[func(*Renderer, *gdk.Surface, **glib.GError) int32, int32]{Name: "gsk_renderer_realize"}
	gsk_renderer_realize_for_display = cc.DlFunc[func(*Renderer, *gdk.Display, **glib.GError) int32, int32]{Name: "gsk_renderer_realize_for_display"}
	gsk_renderer_unrealize           = cc.DlFunc[func(*Renderer), cc.Void]{Name: "gsk_renderer_unrealize"}
	gsk_renderer_is_realized         = cc.DlFunc[func(*Renderer) int32, int32]{Name: "gsk_renderer_is_realized"}
	gsk_renderer_render_texture      = cc.DlFunc[func(*Renderer, *RenderNode, *graphene.Rect) *gdk.Texture, *gdk.Texture]{Name: "gsk_renderer_render_texture"}
	gsk_renderer_render              = cc.DlFunc[func(*Renderer, *RenderNode, *cairo.Region), cc.Void]{Name: "gsk_renderer_render"}
	// #endregion

	// #region RenderNode

	render_node_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_render_node_get_type"}
	serialization_error_quark   = cc.DlFunc[func() uint32, uint32]{Name: "gsk_serialization_error_quark"}
	render_node_ref             = cc.DlFunc[func(*RenderNode) *RenderNode, *RenderNode]{Name: "gsk_render_node_ref"}
	render_node_unref           = cc.DlFunc[func(*RenderNode), cc.Void]{Name: "gsk_render_node_unref"}
	render_node_get_node_type   = cc.DlFunc[func(*RenderNode) RenderNodeType, RenderNodeType]{Name: "gsk_render_node_get_node_type"}
	render_node_get_bounds      = cc.DlFunc[func(*RenderNode, *graphene.Rect), cc.Void]{Name: "gsk_render_node_get_bounds"}
	render_node_get_opaque_rect = cc.DlFunc[func(*RenderNode, *graphene.Rect) int32, int32]{Name: "gsk_render_node_get_opaque_rect"}
	render_node_draw            = cc.DlFunc[func(*RenderNode, *cairo.Context), cc.Void]{Name: "gsk_render_node_draw"}
	render_node_serialize       = cc.DlFunc[func(*RenderNode) *glib.GBytes, *glib.GBytes]{Name: "gsk_render_node_serialize"}
	render_node_write_to_file   = cc.DlFunc[func(*RenderNode, cc.String, **glib.GError) int32, int32]{Name: "gsk_render_node_write_to_file"}
	render_node_deserialize     = cc.DlFunc[func(*glib.GBytes, UPtr, UPtr) *RenderNode, *RenderNode]{Name: "gsk_render_node_deserialize"}

	// DebugNode
	debug_node_get_type    = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_debug_node_get_type"}
	debug_node_new         = cc.DlFunc[func(*RenderNode, cc.String) *DebugNode, *DebugNode]{Name: "gsk_debug_node_new"}
	debug_node_get_child   = cc.DlFunc[func(*DebugNode) *RenderNode, *RenderNode]{Name: "gsk_debug_node_get_child"}
	debug_node_get_message = cc.DlFunc[func(*DebugNode) cc.String, cc.String]{Name: "gsk_debug_node_get_message"}

	// ColorNode
	color_node_get_type  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_color_node_get_type"}
	color_node_new       = cc.DlFunc[func(*gdk.RGBA, *graphene.Rect) *ColorNode, *ColorNode]{Name: "gsk_color_node_new"}
	color_node_get_color = cc.DlFunc[func(*ColorNode) *gdk.RGBA, *gdk.RGBA]{Name: "gsk_color_node_get_color"}

	// TextureNode
	texture_node_get_type    = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_texture_node_get_type"}
	texture_node_new         = cc.DlFunc[func(*gdk.Texture, *graphene.Rect) *TextureNode, *TextureNode]{Name: "gsk_texture_node_new"}
	texture_node_get_texture = cc.DlFunc[func(*TextureNode) *gdk.Texture, *gdk.Texture]{Name: "gsk_texture_node_get_texture"}

	// TextureScaleNode
	texture_scale_node_get_type    = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_texture_scale_node_get_type"}
	texture_scale_node_new         = cc.DlFunc[func(*gdk.Texture, *graphene.Rect, ScalingFilter) *TextureScaleNode, *TextureScaleNode]{Name: "gsk_texture_scale_node_new"}
	texture_scale_node_get_texture = cc.DlFunc[func(*TextureScaleNode) *gdk.Texture, *gdk.Texture]{Name: "gsk_texture_scale_node_get_texture"}
	texture_scale_node_get_filter  = cc.DlFunc[func(*TextureScaleNode) ScalingFilter, ScalingFilter]{Name: "gsk_texture_scale_node_get_filter"}

	// LinearGradientNode
	linear_gradient_node_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_linear_gradient_node_get_type"}
	linear_gradient_node_new               = cc.DlFunc[func(*graphene.Rect, *graphene.Point, *graphene.Point, *ColorStop, uint64) *LinearGradientNode, *LinearGradientNode]{Name: "gsk_linear_gradient_node_new"}
	linear_gradient_node_get_start         = cc.DlFunc[func(*LinearGradientNode) *graphene.Point, *graphene.Point]{Name: "gsk_linear_gradient_node_get_start"}
	linear_gradient_node_get_end           = cc.DlFunc[func(*LinearGradientNode) *graphene.Point, *graphene.Point]{Name: "gsk_linear_gradient_node_get_end"}
	linear_gradient_node_get_n_color_stops = cc.DlFunc[func(*LinearGradientNode) uint64, uint64]{Name: "gsk_linear_gradient_node_get_n_color_stops"}
	linear_gradient_node_get_color_stops   = cc.DlFunc[func(*LinearGradientNode, *uint64) *ColorStop, *ColorStop]{Name: "gsk_linear_gradient_node_get_color_stops"}

	// RepeatingLinearGradientNode
	repeating_linear_gradient_node_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_repeating_linear_gradient_node_get_type"}
	repeating_linear_gradient_node_new      = cc.DlFunc[func(*graphene.Rect, *graphene.Point, *graphene.Point, *ColorStop, uint64) *RepeatingLinearGradientNode, *RepeatingLinearGradientNode]{Name: "gsk_repeating_linear_gradient_node_new"}

	// ConicGradientNode
	conic_gradient_node_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_conic_gradient_node_get_type"}
	conic_gradient_node_new               = cc.DlFunc[func(*graphene.Rect, *graphene.Point, float64, *ColorStop, uint64) *ConicGradientNode, *ConicGradientNode]{Name: "gsk_conic_gradient_node_new"}
	conic_gradient_node_get_center        = cc.DlFunc[func(*ConicGradientNode) *graphene.Point, *graphene.Point]{Name: "gsk_conic_gradient_node_get_center"}
	conic_gradient_node_get_rotation      = cc.DlFunc[func(*ConicGradientNode) float64, float64]{Name: "gsk_conic_gradient_node_get_rotation"}
	conic_gradient_node_get_angle         = cc.DlFunc[func(*ConicGradientNode) float64, float64]{Name: "gsk_conic_gradient_node_get_angle"}
	conic_gradient_node_get_n_color_stops = cc.DlFunc[func(*ConicGradientNode) uint64, uint64]{Name: "gsk_conic_gradient_node_get_n_color_stops"}
	conic_gradient_node_get_color_stops   = cc.DlFunc[func(*ConicGradientNode, *uint64) *ColorStop, *ColorStop]{Name: "gsk_conic_gradient_node_get_color_stops"}

	// RadialGradientNode
	radial_gradient_node_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_radial_gradient_node_get_type"}
	radial_gradient_node_new               = cc.DlFunc[func(*graphene.Rect, *graphene.Point, float64, float64, float64, float64, *ColorStop, uint64) *RadialGradientNode, *RadialGradientNode]{Name: "gsk_radial_gradient_node_new"}
	radial_gradient_node_get_n_color_stops = cc.DlFunc[func(*RadialGradientNode) uint64, uint64]{Name: "gsk_radial_gradient_node_get_n_color_stops"}
	radial_gradient_node_get_color_stops   = cc.DlFunc[func(*RadialGradientNode, *uint64) *ColorStop, *ColorStop]{Name: "gsk_radial_gradient_node_get_color_stops"}
	radial_gradient_node_get_center        = cc.DlFunc[func(*RadialGradientNode) *graphene.Point, *graphene.Point]{Name: "gsk_radial_gradient_node_get_center"}
	radial_gradient_node_get_hradius       = cc.DlFunc[func(*RadialGradientNode) float64, float64]{Name: "gsk_radial_gradient_node_get_hradius"}
	radial_gradient_node_get_vradius       = cc.DlFunc[func(*RadialGradientNode) float64, float64]{Name: "gsk_radial_gradient_node_get_vradius"}
	radial_gradient_node_get_start         = cc.DlFunc[func(*RadialGradientNode) float64, float64]{Name: "gsk_radial_gradient_node_get_start"}
	radial_gradient_node_get_end           = cc.DlFunc[func(*RadialGradientNode) float64, float64]{Name: "gsk_radial_gradient_node_get_end"}

	// RepeatingRadialGradientNode
	repeating_radial_gradient_node_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_repeating_radial_gradient_node_get_type"}
	repeating_radial_gradient_node_new      = cc.DlFunc[func(*graphene.Rect, *graphene.Point, float64, float64, float64, float64, *ColorStop, uint64) *RepeatingRadialGradientNode, *RepeatingRadialGradientNode]{Name: "gsk_repeating_radial_gradient_node_new"}

	// BorderNode
	border_node_get_type    = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_border_node_get_type"}
	border_node_new         = cc.DlFunc[func(*RoundedRect, *float64, *gdk.RGBA) *BorderNode, *BorderNode]{Name: "gsk_border_node_new"}
	border_node_get_outline = cc.DlFunc[func(*BorderNode) *RoundedRect, *RoundedRect]{Name: "gsk_border_node_get_outline"}
	border_node_get_widths  = cc.DlFunc[func(*BorderNode) *float64, *float64]{Name: "gsk_border_node_get_widths"}
	border_node_get_colors  = cc.DlFunc[func(*BorderNode) *gdk.RGBA, *gdk.RGBA]{Name: "gsk_border_node_get_colors"}

	// InsetShadowNode
	inset_shadow_node_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_inset_shadow_node_get_type"}
	inset_shadow_node_new             = cc.DlFunc[func(*RoundedRect, *gdk.RGBA, float64, float64, float64, float64) *InsetShadowNode, *InsetShadowNode]{Name: "gsk_inset_shadow_node_new"}
	inset_shadow_node_get_outline     = cc.DlFunc[func(*InsetShadowNode) *RoundedRect, *RoundedRect]{Name: "gsk_inset_shadow_node_get_outline"}
	inset_shadow_node_get_color       = cc.DlFunc[func(*InsetShadowNode) *gdk.RGBA, *gdk.RGBA]{Name: "gsk_inset_shadow_node_get_color"}
	inset_shadow_node_get_dx          = cc.DlFunc[func(*InsetShadowNode) float64, float64]{Name: "gsk_inset_shadow_node_get_dx"}
	inset_shadow_node_get_dy          = cc.DlFunc[func(*InsetShadowNode) float64, float64]{Name: "gsk_inset_shadow_node_get_dy"}
	inset_shadow_node_get_spread      = cc.DlFunc[func(*InsetShadowNode) float64, float64]{Name: "gsk_inset_shadow_node_get_spread"}
	inset_shadow_node_get_blur_radius = cc.DlFunc[func(*InsetShadowNode) float64, float64]{Name: "gsk_inset_shadow_node_get_blur_radius"}

	// OutsetShadowNode
	outset_shadow_node_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_outset_shadow_node_get_type"}
	outset_shadow_node_new             = cc.DlFunc[func(*RoundedRect, *gdk.RGBA, float64, float64, float64, float64) *OutsetShadowNode, *OutsetShadowNode]{Name: "gsk_outset_shadow_node_new"}
	outset_shadow_node_get_outline     = cc.DlFunc[func(*OutsetShadowNode) *RoundedRect, *RoundedRect]{Name: "gsk_outset_shadow_node_get_outline"}
	outset_shadow_node_get_color       = cc.DlFunc[func(*OutsetShadowNode) *gdk.RGBA, *gdk.RGBA]{Name: "gsk_outset_shadow_node_get_color"}
	outset_shadow_node_get_dx          = cc.DlFunc[func(*OutsetShadowNode) float64, float64]{Name: "gsk_outset_shadow_node_get_dx"}
	outset_shadow_node_get_dy          = cc.DlFunc[func(*OutsetShadowNode) float64, float64]{Name: "gsk_outset_shadow_node_get_dy"}
	outset_shadow_node_get_spread      = cc.DlFunc[func(*OutsetShadowNode) float64, float64]{Name: "gsk_outset_shadow_node_get_spread"}
	outset_shadow_node_get_blur_radius = cc.DlFunc[func(*OutsetShadowNode) float64, float64]{Name: "gsk_outset_shadow_node_get_blur_radius"}

	// CairoNode
	cairo_node_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_cairo_node_get_type"}
	cairo_node_new              = cc.DlFunc[func(*graphene.Rect) *CairoNode, *CairoNode]{Name: "gsk_cairo_node_new"}
	cairo_node_get_draw_context = cc.DlFunc[func(*CairoNode) *cairo.Context, *cairo.Context]{Name: "gsk_cairo_node_get_draw_context"}
	cairo_node_get_surface      = cc.DlFunc[func(*CairoNode) *cairo.Surface, *cairo.Surface]{Name: "gsk_cairo_node_get_surface"}

	// ContainerNode
	container_node_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_container_node_get_type"}
	container_node_new            = cc.DlFunc[func(**RenderNode, uint32) *ContainerNode, *ContainerNode]{Name: "gsk_container_node_new"}
	container_node_get_n_children = cc.DlFunc[func(*ContainerNode) uint32, uint32]{Name: "gsk_container_node_get_n_children"}
	container_node_get_child      = cc.DlFunc[func(*ContainerNode, uint32) *RenderNode, *RenderNode]{Name: "gsk_container_node_get_child"}

	// TransformNode
	transform_node_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_transform_node_get_type"}
	transform_node_new           = cc.DlFunc[func(*RenderNode, *Transform) *TransformNode, *TransformNode]{Name: "gsk_transform_node_new"}
	transform_node_get_child     = cc.DlFunc[func(*TransformNode) *RenderNode, *RenderNode]{Name: "gsk_transform_node_get_child"}
	transform_node_get_transform = cc.DlFunc[func(*TransformNode) *Transform, *Transform]{Name: "gsk_transform_node_get_transform"}

	// OpacityNode
	opacity_node_get_type    = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_opacity_node_get_type"}
	opacity_node_new         = cc.DlFunc[func(*RenderNode, float64) *OpacityNode, *OpacityNode]{Name: "gsk_opacity_node_new"}
	opacity_node_get_child   = cc.DlFunc[func(*OpacityNode) *RenderNode, *RenderNode]{Name: "gsk_opacity_node_get_child"}
	opacity_node_get_opacity = cc.DlFunc[func(*OpacityNode) float64, float64]{Name: "gsk_opacity_node_get_opacity"}

	// ColorMatrixNode
	color_matrix_node_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_color_matrix_node_get_type"}
	color_matrix_node_new              = cc.DlFunc[func(*RenderNode, *graphene.Matrix, *graphene.Vec4) *ColorMatrixNode, *ColorMatrixNode]{Name: "gsk_color_matrix_node_new"}
	color_matrix_node_get_child        = cc.DlFunc[func(*ColorMatrixNode) *RenderNode, *RenderNode]{Name: "gsk_color_matrix_node_get_child"}
	color_matrix_node_get_color_matrix = cc.DlFunc[func(*ColorMatrixNode) *graphene.Matrix, *graphene.Matrix]{Name: "gsk_color_matrix_node_get_color_matrix"}
	color_matrix_node_get_color_offset = cc.DlFunc[func(*ColorMatrixNode) *graphene.Vec4, *graphene.Vec4]{Name: "gsk_color_matrix_node_get_color_offset"}

	// RepeatNode
	repeat_node_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_repeat_node_get_type"}
	repeat_node_new              = cc.DlFunc[func(*graphene.Rect, *RenderNode, *graphene.Rect) *RepeatNode, *RepeatNode]{Name: "gsk_repeat_node_new"}
	repeat_node_get_child        = cc.DlFunc[func(*RepeatNode) *RenderNode, *RenderNode]{Name: "gsk_repeat_node_get_child"}
	repeat_node_get_child_bounds = cc.DlFunc[func(*RepeatNode) *graphene.Rect, *graphene.Rect]{Name: "gsk_repeat_node_get_child_bounds"}

	// ClipNode
	clip_node_get_type  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_clip_node_get_type"}
	clip_node_new       = cc.DlFunc[func(*RenderNode, *graphene.Rect) *ClipNode, *ClipNode]{Name: "gsk_clip_node_new"}
	clip_node_get_child = cc.DlFunc[func(*ClipNode) *RenderNode, *RenderNode]{Name: "gsk_clip_node_get_child"}
	clip_node_get_clip  = cc.DlFunc[func(*ClipNode) *graphene.Rect, *graphene.Rect]{Name: "gsk_clip_node_get_clip"}

	// RoundedClipNode
	rounded_clip_node_get_type  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_rounded_clip_node_get_type"}
	rounded_clip_node_new       = cc.DlFunc[func(*RenderNode, *RoundedRect) *RoundedClipNode, *RoundedClipNode]{Name: "gsk_rounded_clip_node_new"}
	rounded_clip_node_get_child = cc.DlFunc[func(*RoundedClipNode) *RenderNode, *RenderNode]{Name: "gsk_rounded_clip_node_get_child"}
	rounded_clip_node_get_clip  = cc.DlFunc[func(*RoundedClipNode) *RoundedRect, *RoundedRect]{Name: "gsk_rounded_clip_node_get_clip"}

	// FillNode
	fill_node_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_fill_node_get_type"}
	fill_node_new           = cc.DlFunc[func(*RenderNode, *Path, FillRule) *FillNode, *FillNode]{Name: "gsk_fill_node_new"}
	fill_node_get_child     = cc.DlFunc[func(*FillNode) *RenderNode, *RenderNode]{Name: "gsk_fill_node_get_child"}
	fill_node_get_path      = cc.DlFunc[func(*FillNode) *Path, *Path]{Name: "gsk_fill_node_get_path"}
	fill_node_get_fill_rule = cc.DlFunc[func(*FillNode) FillRule, FillRule]{Name: "gsk_fill_node_get_fill_rule"}

	// StrokeNode
	stroke_node_get_type   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_stroke_node_get_type"}
	stroke_node_new        = cc.DlFunc[func(*RenderNode, *Path, *Stroke) *StrokeNode, *StrokeNode]{Name: "gsk_stroke_node_new"}
	stroke_node_get_child  = cc.DlFunc[func(*StrokeNode) *RenderNode, *RenderNode]{Name: "gsk_stroke_node_get_child"}
	stroke_node_get_path   = cc.DlFunc[func(*StrokeNode) *Path, *Path]{Name: "gsk_stroke_node_get_path"}
	stroke_node_get_stroke = cc.DlFunc[func(*StrokeNode) *Stroke, *Stroke]{Name: "gsk_stroke_node_get_stroke"}

	// ShadowNode
	shadow_node_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_shadow_node_get_type"}
	shadow_node_new           = cc.DlFunc[func(*RenderNode, *Shadow, uint64) *ShadowNode, *ShadowNode]{Name: "gsk_shadow_node_new"}
	shadow_node_get_child     = cc.DlFunc[func(*ShadowNode) *RenderNode, *RenderNode]{Name: "gsk_shadow_node_get_child"}
	shadow_node_get_shadow    = cc.DlFunc[func(*ShadowNode, uint64) *Shadow, *Shadow]{Name: "gsk_shadow_node_get_shadow"}
	shadow_node_get_n_shadows = cc.DlFunc[func(*ShadowNode) uint64, uint64]{Name: "gsk_shadow_node_get_n_shadows"}

	// BlendNode
	blend_node_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_blend_node_get_type"}
	blend_node_new              = cc.DlFunc[func(*RenderNode, *RenderNode, BlendMode) *BlendNode, *BlendNode]{Name: "gsk_blend_node_new"}
	blend_node_get_bottom_child = cc.DlFunc[func(*BlendNode) *RenderNode, *RenderNode]{Name: "gsk_blend_node_get_bottom_child"}
	blend_node_get_top_child    = cc.DlFunc[func(*BlendNode) *RenderNode, *RenderNode]{Name: "gsk_blend_node_get_top_child"}
	blend_node_get_blend_mode   = cc.DlFunc[func(*BlendNode) BlendMode, BlendMode]{Name: "gsk_blend_node_get_blend_mode"}

	// CrossFadeNode
	cross_fade_node_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_cross_fade_node_get_type"}
	cross_fade_node_new             = cc.DlFunc[func(*RenderNode, *RenderNode, float64) *CrossFadeNode, *CrossFadeNode]{Name: "gsk_cross_fade_node_new"}
	cross_fade_node_get_start_child = cc.DlFunc[func(*CrossFadeNode) *RenderNode, *RenderNode]{Name: "gsk_cross_fade_node_get_start_child"}
	cross_fade_node_get_end_child   = cc.DlFunc[func(*CrossFadeNode) *RenderNode, *RenderNode]{Name: "gsk_cross_fade_node_get_end_child"}
	cross_fade_node_get_progress    = cc.DlFunc[func(*CrossFadeNode) float64, float64]{Name: "gsk_cross_fade_node_get_progress"}

	// TextNode
	text_node_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_text_node_get_type"}
	text_node_new              = cc.DlFunc[func(*pango.Font, *pango.GlyphString, *gdk.RGBA, *graphene.Point) *TextNode, *TextNode]{Name: "gsk_text_node_new"}
	text_node_get_font         = cc.DlFunc[func(*TextNode) *pango.Font, *pango.Font]{Name: "gsk_text_node_get_font"}
	text_node_has_color_glyphs = cc.DlFunc[func(*TextNode) int32, int32]{Name: "gsk_text_node_has_color_glyphs"}
	text_node_get_num_glyphs   = cc.DlFunc[func(*TextNode) uint32, uint32]{Name: "gsk_text_node_get_num_glyphs"}
	text_node_get_glyphs       = cc.DlFunc[func(*TextNode, *uint32) *pango.GlyphInfo, *pango.GlyphInfo]{Name: "gsk_text_node_get_glyphs"}
	text_node_get_color        = cc.DlFunc[func(*TextNode) *gdk.RGBA, *gdk.RGBA]{Name: "gsk_text_node_get_color"}
	text_node_get_offset       = cc.DlFunc[func(*TextNode) *graphene.Point, *graphene.Point]{Name: "gsk_text_node_get_offset"}

	// BlurNode
	blur_node_get_type   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_blur_node_get_type"}
	blur_node_new        = cc.DlFunc[func(*RenderNode, float64) *BlurNode, *BlurNode]{Name: "gsk_blur_node_new"}
	blur_node_get_child  = cc.DlFunc[func(*BlurNode) *RenderNode, *RenderNode]{Name: "gsk_blur_node_get_child"}
	blur_node_get_radius = cc.DlFunc[func(*BlurNode) float64, float64]{Name: "gsk_blur_node_get_radius"}

	// MaskNode
	mask_node_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_mask_node_get_type"}
	mask_node_new           = cc.DlFunc[func(*RenderNode, *RenderNode, MaskMode) *MaskNode, *MaskNode]{Name: "gsk_mask_node_new"}
	mask_node_get_source    = cc.DlFunc[func(*MaskNode) *RenderNode, *RenderNode]{Name: "gsk_mask_node_get_source"}
	mask_node_get_mask      = cc.DlFunc[func(*MaskNode) *RenderNode, *RenderNode]{Name: "gsk_mask_node_get_mask"}
	mask_node_get_mask_mode = cc.DlFunc[func(*MaskNode) MaskMode, MaskMode]{Name: "gsk_mask_node_get_mask_mode"}

	// GLShaderNode (已过时)
	gl_shader_node_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_gl_shader_node_get_type"}
	gl_shader_node_new            = cc.DlFunc[func(*GLShader, *graphene.Rect, *glib.GBytes, **RenderNode, uint32) *GLShaderNode, *GLShaderNode]{Name: "gsk_gl_shader_node_new"}
	gl_shader_node_get_n_children = cc.DlFunc[func(*GLShaderNode) uint32, uint32]{Name: "gsk_gl_shader_node_get_n_children"}
	gl_shader_node_get_child      = cc.DlFunc[func(*GLShaderNode, uint32) *RenderNode, *RenderNode]{Name: "gsk_gl_shader_node_get_child"}
	gl_shader_node_get_args       = cc.DlFunc[func(*GLShaderNode) *glib.GBytes, *glib.GBytes]{Name: "gsk_gl_shader_node_get_args"}
	gl_shader_node_get_shader     = cc.DlFunc[func(*GLShaderNode) *GLShader, *GLShader]{Name: "gsk_gl_shader_node_get_shader"}

	// SubsurfaceNode
	subsurface_node_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_subsurface_node_get_type"}
	subsurface_node_new            = cc.DlFunc[func(*RenderNode, UPtr) *SubsurfaceNode, *SubsurfaceNode]{Name: "gsk_subsurface_node_new"}
	subsurface_node_get_child      = cc.DlFunc[func(*SubsurfaceNode) *RenderNode, *RenderNode]{Name: "gsk_subsurface_node_get_child"}
	subsurface_node_get_subsurface = cc.DlFunc[func(*SubsurfaceNode) UPtr, UPtr]{Name: "gsk_subsurface_node_get_subsurface"}

	// Value functions
	value_set_render_node  = cc.DlFunc[func(*gobject.GValue, *RenderNode), cc.Void]{Name: "gsk_value_set_render_node"}
	value_take_render_node = cc.DlFunc[func(*gobject.GValue, *RenderNode), cc.Void]{Name: "gsk_value_take_render_node"}
	value_get_render_node  = cc.DlFunc[func(*gobject.GValue) *RenderNode, *RenderNode]{Name: "gsk_value_get_render_node"}
	value_dup_render_node  = cc.DlFunc[func(*gobject.GValue) *RenderNode, *RenderNode]{Name: "gsk_value_dup_render_node"}
	// #endregion

	// #region RoundedRect
	rounded_rect_init            = cc.DlFunc[func(*RoundedRect, *graphene.Rect, *graphene.Size, *graphene.Size, *graphene.Size, *graphene.Size) *RoundedRect, *RoundedRect]{Name: "gsk_rounded_rect_init"}
	rounded_rect_init_copy       = cc.DlFunc[func(*RoundedRect, *RoundedRect) *RoundedRect, *RoundedRect]{Name: "gsk_rounded_rect_init_copy"}
	rounded_rect_init_from_rect  = cc.DlFunc[func(*RoundedRect, *graphene.Rect, float64) *RoundedRect, *RoundedRect]{Name: "gsk_rounded_rect_init_from_rect"}
	rounded_rect_normalize       = cc.DlFunc[func(*RoundedRect) *RoundedRect, *RoundedRect]{Name: "gsk_rounded_rect_normalize"}
	rounded_rect_offset          = cc.DlFunc[func(*RoundedRect, float64, float64) *RoundedRect, *RoundedRect]{Name: "gsk_rounded_rect_offset"}
	rounded_rect_shrink          = cc.DlFunc[func(*RoundedRect, float64, float64, float64, float64) *RoundedRect, *RoundedRect]{Name: "gsk_rounded_rect_shrink"}
	rounded_rect_is_rectilinear  = cc.DlFunc[func(*RoundedRect) int32, int32]{Name: "gsk_rounded_rect_is_rectilinear"}
	rounded_rect_contains_point  = cc.DlFunc[func(*RoundedRect, *graphene.Point) int32, int32]{Name: "gsk_rounded_rect_contains_point"}
	rounded_rect_contains_rect   = cc.DlFunc[func(*RoundedRect, *graphene.Rect) int32, int32]{Name: "gsk_rounded_rect_contains_rect"}
	rounded_rect_intersects_rect = cc.DlFunc[func(*RoundedRect, *graphene.Rect) int32, int32]{Name: "gsk_rounded_rect_intersects_rect"}
	// #endregion

	// #region Stroke
	stroke_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_stroke_get_type"}
	stroke_new             = cc.DlFunc[func(float64) *Stroke, *Stroke]{Name: "gsk_stroke_new"}
	stroke_copy            = cc.DlFunc[func(*Stroke) *Stroke, *Stroke]{Name: "gsk_stroke_copy"}
	stroke_free            = cc.DlFunc[func(*Stroke), cc.Void]{Name: "gsk_stroke_free"}
	stroke_equal           = cc.DlFunc[func(*Stroke, *Stroke) int32, int32]{Name: "gsk_stroke_equal"}
	stroke_set_line_width  = cc.DlFunc[func(*Stroke, float64), cc.Void]{Name: "gsk_stroke_set_line_width"}
	stroke_get_line_width  = cc.DlFunc[func(*Stroke) float64, float64]{Name: "gsk_stroke_get_line_width"}
	stroke_set_line_cap    = cc.DlFunc[func(*Stroke, LineCap), cc.Void]{Name: "gsk_stroke_set_line_cap"}
	stroke_get_line_cap    = cc.DlFunc[func(*Stroke) LineCap, LineCap]{Name: "gsk_stroke_get_line_cap"}
	stroke_set_line_join   = cc.DlFunc[func(*Stroke, LineJoin), cc.Void]{Name: "gsk_stroke_set_line_join"}
	stroke_get_line_join   = cc.DlFunc[func(*Stroke) LineJoin, LineJoin]{Name: "gsk_stroke_get_line_join"}
	stroke_set_miter_limit = cc.DlFunc[func(*Stroke, float64), cc.Void]{Name: "gsk_stroke_set_miter_limit"}
	stroke_get_miter_limit = cc.DlFunc[func(*Stroke) float64, float64]{Name: "gsk_stroke_get_miter_limit"}
	stroke_set_dash        = cc.DlFunc[func(*Stroke, *float64, uint64), cc.Void]{Name: "gsk_stroke_set_dash"}
	stroke_get_dash        = cc.DlFunc[func(*Stroke, *uint64) *float64, *float64]{Name: "gsk_stroke_get_dash"}
	stroke_set_dash_offset = cc.DlFunc[func(*Stroke, float64), cc.Void]{Name: "gsk_stroke_set_dash_offset"}
	stroke_get_dash_offset = cc.DlFunc[func(*Stroke) float64, float64]{Name: "gsk_stroke_get_dash_offset"}
	stroke_to_cairo        = cc.DlFunc[func(*Stroke, *cairo.Context), cc.Void]{Name: "gsk_stroke_to_cairo"}
	// #endregion

	// #region Transform
	gsk_transform_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_transform_get_type"}
	gsk_transform_ref              = cc.DlFunc[func(self *Transform) *Transform, *Transform]{Name: "gsk_transform_ref"}
	gsk_transform_unref            = cc.DlFunc[func(self *Transform), cc.Void]{Name: "gsk_transform_unref"}
	gsk_transform_print            = cc.DlFunc[func(self *Transform, str *glib.GString), cc.Void]{Name: "gsk_transform_print"}
	gsk_transform_to_string        = cc.DlFunc[func(self *Transform) cc.String, cc.String]{Name: "gsk_transform_to_string"}
	gsk_transform_parse            = cc.DlFunc[func(str cc.String, out_transform **Transform) int32, int32]{Name: "gsk_transform_parse"}
	gsk_transform_to_matrix        = cc.DlFunc[func(self *Transform, out_matrix *graphene.Matrix), cc.Void]{Name: "gsk_transform_to_matrix"}
	gsk_transform_to_2d            = cc.DlFunc[func(self *Transform, out_xx, out_yx, out_xy, out_yy, out_dx, out_dy *float32), cc.Void]{Name: "gsk_transform_to_2d"}
	gsk_transform_to_2d_components = cc.DlFunc[func(self *Transform, out_skew_x, out_skew_y, out_scale_x, out_scale_y, out_angle, out_dx, out_dy *float32), cc.Void]{Name: "gsk_transform_to_2d_components"}
	gsk_transform_to_affine        = cc.DlFunc[func(self *Transform, out_scale_x, out_scale_y, out_dx, out_dy *float32), cc.Void]{Name: "gsk_transform_to_affine"}
	gsk_transform_to_translate     = cc.DlFunc[func(self *Transform, out_dx, out_dy *float32), cc.Void]{Name: "gsk_transform_to_translate"}
	gsk_transform_get_category     = cc.DlFunc[func(self *Transform) TransformCategory, TransformCategory]{Name: "gsk_transform_get_category"}
	gsk_transform_equal            = cc.DlFunc[func(first, second *Transform) int32, int32]{Name: "gsk_transform_equal"}
	gsk_transform_new              = cc.DlFunc[func() *Transform, *Transform]{Name: "gsk_transform_new"}
	gsk_transform_transform        = cc.DlFunc[func(next, other *Transform) *Transform, *Transform]{Name: "gsk_transform_transform"}
	gsk_transform_invert           = cc.DlFunc[func(self *Transform) *Transform, *Transform]{Name: "gsk_transform_invert"}
	gsk_transform_matrix           = cc.DlFunc[func(next *Transform, matrix *graphene.Matrix) *Transform, *Transform]{Name: "gsk_transform_matrix"}
	gsk_transform_translate        = cc.DlFunc[func(next *Transform, point *graphene.Point) *Transform, *Transform]{Name: "gsk_transform_translate"}
	gsk_transform_translate_3d     = cc.DlFunc[func(next *Transform, point *graphene.Point3D) *Transform, *Transform]{Name: "gsk_transform_translate_3d"}
	gsk_transform_skew             = cc.DlFunc[func(next *Transform, skew_x, skew_y float32) *Transform, *Transform]{Name: "gsk_transform_skew"}
	gsk_transform_rotate           = cc.DlFunc[func(next *Transform, angle float32) *Transform, *Transform]{Name: "gsk_transform_rotate"}
	gsk_transform_rotate_3d        = cc.DlFunc[func(next *Transform, angle float32, axis *graphene.Vec3) *Transform, *Transform]{Name: "gsk_transform_rotate_3d"}
	gsk_transform_scale            = cc.DlFunc[func(next *Transform, factor_x, factor_y float32) *Transform, *Transform]{Name: "gsk_transform_scale"}
	gsk_transform_scale_3d         = cc.DlFunc[func(next *Transform, factor_x, factor_y, factor_z float32) *Transform, *Transform]{Name: "gsk_transform_scale_3d"}
	gsk_transform_perspective      = cc.DlFunc[func(next *Transform, depth float32) *Transform, *Transform]{Name: "gsk_transform_perspective"}
	gsk_transform_transform_bounds = cc.DlFunc[func(self *Transform, rect *graphene.Rect, out_rect *graphene.Rect), cc.Void]{Name: "gsk_transform_transform_bounds"}
	gsk_transform_transform_point  = cc.DlFunc[func(self *Transform, point *graphene.Point, out_point *graphene.Point), cc.Void]{Name: "gsk_transform_transform_point"}
	// #endregion

	gl_renderer_new          = cc.DlFunc[func() *GLRenderer, *GLRenderer]{Name: "gsk_gl_renderer_new"}
	gl_renderer_get_type     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_gl_renderer_get_type"}
	vulkan_renderer_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_vulkan_renderer_get_type"}
	vulkan_renderer_new      = cc.DlFunc[func() *VulkanRenderer, *VulkanRenderer]{Name: "gsk_vulkan_renderer_new"}
)
