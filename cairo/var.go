package cairo

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/glib/gobject"
)

type uptr = unsafe.Pointer
type iptr = uintptr

func slice(p uptr, n int) uptr  { return uptr(&[2]uptr{uptr(p), uptr(iptr(n))}) }
func vcbu(fn interface{}) uptr  { return cc.Cbk(fn) }
func anyptr(a interface{}) uptr { return (*(*[2]uptr)(uptr(&a)))[1] }
func carry[T any](arry []T) *T {
	if len(arry) == 0 {
		return nil
	}
	return &arry[0]
}

func init() { cc.Open("libcairo-gobject*") }

var (

	// #region cairo.h
	cairo_version                               = cc.DlFunc[func() int32, int32]{Name: "cairo_version"}
	cairo_version_string                        = cc.DlFunc[func() string, string]{Name: "cairo_version_string"}
	cairo_pattern_set_dither                    = cc.DlFunc[func(pattern *Pattern, dither Dither), cc.Void]{Name: "cairo_pattern_set_dither"}
	cairo_pattern_get_dither                    = cc.DlFunc[func(pattern *Pattern) Dither, Dither]{Name: "cairo_pattern_get_dither"}
	cairo_create                                = cc.DlFunc[func(target *Surface) *Context, *Context]{Name: "cairo_create"}
	cairo_reference                             = cc.DlFunc[func(cr *Context) *Context, *Context]{Name: "cairo_reference"}
	cairo_destroy                               = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_destroy"}
	cairo_get_reference_count                   = cc.DlFunc[func(cr *Context) uint32, uint32]{Name: "cairo_get_reference_count"}
	cairo_get_user_data                         = cc.DlFunc[func(cr *Context, key *UserDataKey) uptr, uptr]{Name: "cairo_get_user_data"}
	cairo_set_user_data                         = cc.DlFunc[func(cr *Context, key *UserDataKey, user_data uptr, destroy uptr) Status, Status]{Name: "cairo_set_user_data"}
	cairo_save                                  = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_save"}
	cairo_restore                               = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_restore"}
	cairo_push_group                            = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_push_group"}
	cairo_push_group_with_content               = cc.DlFunc[func(cr *Context, content int32), cc.Void]{Name: "cairo_push_group_with_content"}
	cairo_pop_group                             = cc.DlFunc[func(cr *Context) *Pattern, *Pattern]{Name: "cairo_pop_group"}
	cairo_pop_group_to_source                   = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_pop_group_to_source"}
	cairo_set_operator                          = cc.DlFunc[func(cr *Context, op Operator), cc.Void]{Name: "cairo_set_operator"}
	cairo_set_source                            = cc.DlFunc[func(cr *Context, source *Pattern), cc.Void]{Name: "cairo_set_source"}
	cairo_set_source_rgb                        = cc.DlFunc[func(cr *Context, red, green, blue float64), cc.Void]{Name: "cairo_set_source_rgb"}
	cairo_set_source_rgba                       = cc.DlFunc[func(cr *Context, red, green, blue, alpha float64), cc.Void]{Name: "cairo_set_source_rgba"}
	cairo_set_source_surface                    = cc.DlFunc[func(cr *Context, surface *Surface, x, y float64), cc.Void]{Name: "cairo_set_source_surface"}
	cairo_set_tolerance                         = cc.DlFunc[func(cr *Context, tolerance float64), cc.Void]{Name: "cairo_set_tolerance"}
	cairo_set_antialias                         = cc.DlFunc[func(cr *Context, antialias Antialias), cc.Void]{Name: "cairo_set_antialias"}
	cairo_set_fill_rule                         = cc.DlFunc[func(cr *Context, fill_rule FillRule), cc.Void]{Name: "cairo_set_fill_rule"}
	cairo_set_line_width                        = cc.DlFunc[func(cr *Context, width float64), cc.Void]{Name: "cairo_set_line_width"}
	cairo_set_hairline                          = cc.DlFunc[func(cr *Context, set_hairline Bool), cc.Void]{Name: "cairo_set_hairline"}
	cairo_set_line_cap                          = cc.DlFunc[func(cr *Context, line_cap LineCap), cc.Void]{Name: "cairo_set_line_cap"}
	cairo_set_line_join                         = cc.DlFunc[func(cr *Context, line_join LineJoin), cc.Void]{Name: "cairo_set_line_join"}
	cairo_set_dash                              = cc.DlFunc[func(cr *Context, dashes *float64, num_dashes int32, offset float64), cc.Void]{Name: "cairo_set_dash"}
	cairo_set_miter_limit                       = cc.DlFunc[func(cr *Context, limit float64), cc.Void]{Name: "cairo_set_miter_limit"}
	cairo_translate                             = cc.DlFunc[func(cr *Context, tx, ty float64), cc.Void]{Name: "cairo_translate"}
	cairo_scale                                 = cc.DlFunc[func(cr *Context, sx, sy float64), cc.Void]{Name: "cairo_scale"}
	cairo_rotate                                = cc.DlFunc[func(cr *Context, angle float64), cc.Void]{Name: "cairo_rotate"}
	cairo_transform                             = cc.DlFunc[func(cr *Context, matrix *Matrix), cc.Void]{Name: "cairo_transform"}
	cairo_set_matrix                            = cc.DlFunc[func(cr *Context, matrix *Matrix), cc.Void]{Name: "cairo_set_matrix"}
	cairo_identity_matrix                       = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_identity_matrix"}
	cairo_user_to_device                        = cc.DlFunc[func(cr *Context, x, y *float64), cc.Void]{Name: "cairo_user_to_device"}
	cairo_user_to_device_distance               = cc.DlFunc[func(cr *Context, dx, dy *float64), cc.Void]{Name: "cairo_user_to_device_distance"}
	cairo_device_to_user                        = cc.DlFunc[func(cr *Context, x, y *float64), cc.Void]{Name: "cairo_device_to_user"}
	cairo_device_to_user_distance               = cc.DlFunc[func(cr *Context, dx, dy *float64), cc.Void]{Name: "cairo_device_to_user_distance"}
	cairo_new_path                              = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_new_path"}
	cairo_move_to                               = cc.DlFunc[func(cr *Context, x, y float64), cc.Void]{Name: "cairo_move_to"}
	cairo_new_sub_path                          = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_new_sub_path"}
	cairo_line_to                               = cc.DlFunc[func(cr *Context, x, y float64), cc.Void]{Name: "cairo_line_to"}
	cairo_curve_to                              = cc.DlFunc[func(cr *Context, x1, y1, x2, y2, x3, y3 float64), cc.Void]{Name: "cairo_curve_to"}
	cairo_arc                                   = cc.DlFunc[func(cr *Context, xc, yc, radius, angle1, angle2 float64), cc.Void]{Name: "cairo_arc"}
	cairo_arc_negative                          = cc.DlFunc[func(cr *Context, xc, yc, radius, angle1, angle2 float64), cc.Void]{Name: "cairo_arc_negative"}
	cairo_rel_move_to                           = cc.DlFunc[func(cr *Context, dx, dy float64), cc.Void]{Name: "cairo_rel_move_to"}
	cairo_rel_line_to                           = cc.DlFunc[func(cr *Context, dx, dy float64), cc.Void]{Name: "cairo_rel_line_to"}
	cairo_rel_curve_to                          = cc.DlFunc[func(cr *Context, dx1, dy1, dx2, dy2, dx3, dy3 float64), cc.Void]{Name: "cairo_rel_curve_to"}
	cairo_rectangle                             = cc.DlFunc[func(cr *Context, x, y, width, height float64), cc.Void]{Name: "cairo_rectangle"}
	cairo_close_path                            = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_close_path"}
	cairo_path_extents                          = cc.DlFunc[func(cr *Context, x1, y1, x2, y2 *float64), cc.Void]{Name: "cairo_path_extents"}
	cairo_paint                                 = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_paint"}
	cairo_paint_with_alpha                      = cc.DlFunc[func(cr *Context, alpha float64), cc.Void]{Name: "cairo_paint_with_alpha"}
	cairo_mask                                  = cc.DlFunc[func(cr *Context, pattern *Pattern), cc.Void]{Name: "cairo_mask"}
	cairo_mask_surface                          = cc.DlFunc[func(cr *Context, surface *Surface, surface_x, surface_y float64), cc.Void]{Name: "cairo_mask_surface"}
	cairo_stroke                                = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_stroke"}
	cairo_stroke_preserve                       = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_stroke_preserve"}
	cairo_fill                                  = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_fill"}
	cairo_fill_preserve                         = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_fill_preserve"}
	cairo_copy_page                             = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_copy_page"}
	cairo_show_page                             = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_show_page"}
	cairo_in_stroke                             = cc.DlFunc[func(cr *Context, x, y float64) Bool, Bool]{Name: "cairo_in_stroke"}
	cairo_in_fill                               = cc.DlFunc[func(cr *Context, x, y float64) Bool, Bool]{Name: "cairo_in_fill"}
	cairo_in_clip                               = cc.DlFunc[func(cr *Context, x, y float64) Bool, Bool]{Name: "cairo_in_clip"}
	cairo_stroke_extents                        = cc.DlFunc[func(cr *Context, x1, y1, x2, y2 *float64), cc.Void]{Name: "cairo_stroke_extents"}
	cairo_fill_extents                          = cc.DlFunc[func(cr *Context, x1, y1, x2, y2 *float64), cc.Void]{Name: "cairo_fill_extents"}
	cairo_reset_clip                            = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_reset_clip"}
	cairo_clip                                  = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_clip"}
	cairo_clip_preserve                         = cc.DlFunc[func(cr *Context), cc.Void]{Name: "cairo_clip_preserve"}
	cairo_clip_extents                          = cc.DlFunc[func(cr *Context, x1, y1, x2, y2 *float64), cc.Void]{Name: "cairo_clip_extents"}
	cairo_copy_clip_rectangle_list              = cc.DlFunc[func(cr *Context) *RectangleList, *RectangleList]{Name: "cairo_copy_clip_rectangle_list"}
	cairo_rectangle_list_destroy                = cc.DlFunc[func(rectangle_list *RectangleList), cc.Void]{Name: "cairo_rectangle_list_destroy"}
	cairo_tag_begin                             = cc.DlFunc[func(cr *Context, tag_name, attributes cc.String), cc.Void]{Name: "cairo_tag_begin"}
	cairo_tag_end                               = cc.DlFunc[func(cr *Context, tag_name cc.String), cc.Void]{Name: "cairo_tag_end"}
	cairo_glyph_allocate                        = cc.DlFunc[func(num_glyphs int32) *Glyph, *Glyph]{Name: "cairo_glyph_allocate"}
	cairo_glyph_free                            = cc.DlFunc[func(glyphs *Glyph), cc.Void]{Name: "cairo_glyph_free"}
	cairo_text_cluster_allocate                 = cc.DlFunc[func(num_clusters int32) *TextCluster, *TextCluster]{Name: "cairo_text_cluster_allocate"}
	cairo_text_cluster_free                     = cc.DlFunc[func(clusters *TextCluster), cc.Void]{Name: "cairo_text_cluster_free"}
	cairo_font_options_create                   = cc.DlFunc[func() *FontOptions, *FontOptions]{Name: "cairo_font_options_create"}
	cairo_font_options_copy                     = cc.DlFunc[func(original *FontOptions) *FontOptions, *FontOptions]{Name: "cairo_font_options_copy"}
	cairo_font_options_destroy                  = cc.DlFunc[func(options *FontOptions), cc.Void]{Name: "cairo_font_options_destroy"}
	cairo_font_options_status                   = cc.DlFunc[func(options *FontOptions) Status, Status]{Name: "cairo_font_options_status"}
	cairo_font_options_merge                    = cc.DlFunc[func(options *FontOptions, other *FontOptions), cc.Void]{Name: "cairo_font_options_merge"}
	cairo_font_options_equal                    = cc.DlFunc[func(options *FontOptions, other *FontOptions) Bool, Bool]{Name: "cairo_font_options_equal"}
	cairo_font_options_hash                     = cc.DlFunc[func(options *FontOptions) uint64, uint64]{Name: "cairo_font_options_hash"}
	cairo_font_options_set_antialias            = cc.DlFunc[func(options *FontOptions, antialias Antialias), cc.Void]{Name: "cairo_font_options_set_antialias"}
	cairo_font_options_get_antialias            = cc.DlFunc[func(options *FontOptions) Antialias, Antialias]{Name: "cairo_font_options_get_antialias"}
	cairo_font_options_set_subpixel_order       = cc.DlFunc[func(options *FontOptions, subpixel_order SubpixelOrder), cc.Void]{Name: "cairo_font_options_set_subpixel_order"}
	cairo_font_options_get_subpixel_order       = cc.DlFunc[func(options *FontOptions) SubpixelOrder, SubpixelOrder]{Name: "cairo_font_options_get_subpixel_order"}
	cairo_font_options_set_hint_style           = cc.DlFunc[func(options *FontOptions, hint_style HintStyle), cc.Void]{Name: "cairo_font_options_set_hint_style"}
	cairo_font_options_get_hint_style           = cc.DlFunc[func(options *FontOptions) HintStyle, HintStyle]{Name: "cairo_font_options_get_hint_style"}
	cairo_font_options_set_hint_metrics         = cc.DlFunc[func(options *FontOptions, hint_metrics HintMetrics), cc.Void]{Name: "cairo_font_options_set_hint_metrics"}
	cairo_font_options_get_hint_metrics         = cc.DlFunc[func(options *FontOptions) HintMetrics, HintMetrics]{Name: "cairo_font_options_get_hint_metrics"}
	cairo_font_options_get_variations           = cc.DlFunc[func(options *FontOptions) cc.String, cc.String]{Name: "cairo_font_options_get_variations"}
	cairo_font_options_set_variations           = cc.DlFunc[func(options *FontOptions, variations cc.String), cc.Void]{Name: "cairo_font_options_set_variations"}
	cairo_font_options_set_color_mode           = cc.DlFunc[func(options *FontOptions, color_mode ColorMode), cc.Void]{Name: "cairo_font_options_set_color_mode"}
	cairo_font_options_get_color_mode           = cc.DlFunc[func(options *FontOptions) ColorMode, ColorMode]{Name: "cairo_font_options_get_color_mode"}
	cairo_font_options_get_color_palette        = cc.DlFunc[func(options *FontOptions) uint32, uint32]{Name: "cairo_font_options_get_color_palette"}
	cairo_font_options_set_color_palette        = cc.DlFunc[func(options *FontOptions, palette_index uint32), cc.Void]{Name: "cairo_font_options_set_color_palette"}
	cairo_font_options_set_custom_palette_color = cc.DlFunc[func(options *FontOptions, index uint32, red, green, blue, alpha float64), cc.Void]{Name: "cairo_font_options_set_custom_palette_color"}
	cairo_font_options_get_custom_palette_color = cc.DlFunc[func(options *FontOptions, index uint32, red, green, blue, alpha *float64) Status, Status]{Name: "cairo_font_options_get_custom_palette_color"}
	cairo_select_font_face                      = cc.DlFunc[func(cr *Context, family cc.String, slant FontSlant, weight FontWeight), cc.Void]{Name: "cairo_select_font_face"}
	cairo_set_font_size                         = cc.DlFunc[func(cr *Context, size float64), cc.Void]{Name: "cairo_set_font_size"}
	cairo_set_font_matrix                       = cc.DlFunc[func(cr *Context, matrix *Matrix), cc.Void]{Name: "cairo_set_font_matrix"}
	cairo_get_font_matrix                       = cc.DlFunc[func(cr *Context, matrix *Matrix), cc.Void]{Name: "cairo_get_font_matrix"}
	cairo_set_font_options                      = cc.DlFunc[func(cr *Context, options *FontOptions), cc.Void]{Name: "cairo_set_font_options"}
	cairo_get_font_options                      = cc.DlFunc[func(cr *Context, options *FontOptions), cc.Void]{Name: "cairo_get_font_options"}
	cairo_set_font_face                         = cc.DlFunc[func(cr *Context, font_face *FontFace), cc.Void]{Name: "cairo_set_font_face"}
	cairo_get_font_face                         = cc.DlFunc[func(cr *Context) *FontFace, *FontFace]{Name: "cairo_get_font_face"}
	cairo_set_scaled_font                       = cc.DlFunc[func(cr *Context, scaled_font *ScaledFont), cc.Void]{Name: "cairo_set_scaled_font"}
	cairo_get_scaled_font                       = cc.DlFunc[func(cr *Context) *ScaledFont, *ScaledFont]{Name: "cairo_get_scaled_font"}
	cairo_show_text                             = cc.DlFunc[func(cr *Context, utf8 cc.String), cc.Void]{Name: "cairo_show_text"}
	cairo_show_glyphs                           = cc.DlFunc[func(cr *Context, glyphs *Glyph, num_glyphs int32), cc.Void]{Name: "cairo_show_glyphs"}
	cairo_show_text_glyphs                      = cc.DlFunc[func(cr *Context, utf8 cc.String, utf8_len int32, glyphs *Glyph, num_glyphs int32, clusters *TextCluster, num_clusters int32, cluster_flags int32), cc.Void]{Name: "cairo_show_text_glyphs"}
	cairo_text_path                             = cc.DlFunc[func(cr *Context, utf8 cc.String), cc.Void]{Name: "cairo_text_path"}
	cairo_glyph_path                            = cc.DlFunc[func(cr *Context, glyphs *Glyph, num_glyphs int32), cc.Void]{Name: "cairo_glyph_path"}
	cairo_text_extents                          = cc.DlFunc[func(cr *Context, utf8 cc.String, extents *TextExtents), cc.Void]{Name: "cairo_text_extents"}
	cairo_glyph_extents                         = cc.DlFunc[func(cr *Context, glyphs *Glyph, num_glyphs int32, extents *TextExtents), cc.Void]{Name: "cairo_glyph_extents"}
	cairo_font_extents                          = cc.DlFunc[func(cr *Context, extents *FontExtents), cc.Void]{Name: "cairo_font_extents"}
	cairo_font_face_reference                   = cc.DlFunc[func(font_face *FontFace) *FontFace, *FontFace]{Name: "cairo_font_face_reference"}
	cairo_font_face_destroy                     = cc.DlFunc[func(font_face *FontFace), cc.Void]{Name: "cairo_font_face_destroy"}
	cairo_font_face_get_reference_count         = cc.DlFunc[func(font_face *FontFace) uint32, uint32]{Name: "cairo_font_face_get_reference_count"}
	cairo_font_face_status                      = cc.DlFunc[func(font_face *FontFace) Status, Status]{Name: "cairo_font_face_status"}
	cairo_font_face_get_type                    = cc.DlFunc[func(font_face *FontFace) FontType, FontType]{Name: "cairo_font_face_get_type"}
	cairo_font_face_get_user_data               = cc.DlFunc[func(font_face *FontFace, key *UserDataKey) uptr, uptr]{Name: "cairo_font_face_get_user_data"}
	cairo_font_face_set_user_data               = cc.DlFunc[func(font_face *FontFace, key *UserDataKey, user_data uptr, destroy uptr) Status, Status]{Name: "cairo_font_face_set_user_data"}

	cairo_scaled_font_create              = cc.DlFunc[func(font_face *FontFace, font_matrix, ctm *Matrix, options *FontOptions) *ScaledFont, *ScaledFont]{Name: "cairo_scaled_font_create"}
	cairo_scaled_font_reference           = cc.DlFunc[func(scaled_font *ScaledFont) *ScaledFont, *ScaledFont]{Name: "cairo_scaled_font_reference"}
	cairo_scaled_font_destroy             = cc.DlFunc[func(scaled_font *ScaledFont), cc.Void]{Name: "cairo_scaled_font_destroy"}
	cairo_scaled_font_get_reference_count = cc.DlFunc[func(scaled_font *ScaledFont) uint32, uint32]{Name: "cairo_scaled_font_get_reference_count"}
	cairo_scaled_font_status              = cc.DlFunc[func(scaled_font *ScaledFont) Status, Status]{Name: "cairo_scaled_font_status"}
	cairo_scaled_font_get_type            = cc.DlFunc[func(scaled_font *ScaledFont) FontType, FontType]{Name: "cairo_scaled_font_get_type"}
	cairo_scaled_font_get_user_data       = cc.DlFunc[func(scaled_font *ScaledFont, key *UserDataKey) uptr, uptr]{Name: "cairo_scaled_font_get_user_data"}
	cairo_scaled_font_set_user_data       = cc.DlFunc[func(scaled_font *ScaledFont, key *UserDataKey, user_data uptr, destroy uptr) Status, Status]{Name: "cairo_scaled_font_set_user_data"}
	cairo_scaled_font_extents             = cc.DlFunc[func(scaled_font *ScaledFont, extents *FontExtents), cc.Void]{Name: "cairo_scaled_font_extents"}
	cairo_scaled_font_text_extents        = cc.DlFunc[func(scaled_font *ScaledFont, utf8 cc.String, extents *TextExtents), cc.Void]{Name: "cairo_scaled_font_text_extents"}
	cairo_scaled_font_glyph_extents       = cc.DlFunc[func(scaled_font *ScaledFont, glyphs *Glyph, num_glyphs int32, extents *TextExtents), cc.Void]{Name: "cairo_scaled_font_glyph_extents"}
	cairo_scaled_font_text_to_glyphs      = cc.DlFunc[func(scaled_font *ScaledFont, x, y float64, utf8 cc.String, utf8_len int32, glyphs **Glyph, num_glyphs *int32, clusters **TextCluster, num_clusters *int32, cluster_flags *int32) Status, Status]{Name: "cairo_scaled_font_text_to_glyphs"}
	cairo_scaled_font_get_font_face       = cc.DlFunc[func(scaled_font *ScaledFont) *FontFace, *FontFace]{Name: "cairo_scaled_font_get_font_face"}
	cairo_scaled_font_get_font_matrix     = cc.DlFunc[func(scaled_font *ScaledFont, font_matrix *Matrix), cc.Void]{Name: "cairo_scaled_font_get_font_matrix"}
	cairo_scaled_font_get_ctm             = cc.DlFunc[func(scaled_font *ScaledFont, ctm *Matrix), cc.Void]{Name: "cairo_scaled_font_get_ctm"}
	cairo_scaled_font_get_scale_matrix    = cc.DlFunc[func(scaled_font *ScaledFont, scale_matrix *Matrix), cc.Void]{Name: "cairo_scaled_font_get_scale_matrix"}
	cairo_scaled_font_get_font_options    = cc.DlFunc[func(scaled_font *ScaledFont, options *FontOptions), cc.Void]{Name: "cairo_scaled_font_get_font_options"}

	cairo_toy_font_face_create     = cc.DlFunc[func(family cc.String, slant FontSlant, weight FontWeight) *ToyFontFace, *ToyFontFace]{Name: "cairo_toy_font_face_create"}
	cairo_toy_font_face_get_family = cc.DlFunc[func(font_face *ToyFontFace) cc.String, cc.String]{Name: "cairo_toy_font_face_get_family"}
	cairo_toy_font_face_get_slant  = cc.DlFunc[func(font_face *ToyFontFace) FontSlant, FontSlant]{Name: "cairo_toy_font_face_get_slant"}
	cairo_toy_font_face_get_weight = cc.DlFunc[func(font_face *ToyFontFace) FontWeight, FontWeight]{Name: "cairo_toy_font_face_get_weight"}

	cairo_user_font_face_create                      = cc.DlFunc[func() *UserFontFace, *UserFontFace]{Name: "cairo_user_font_face_create"}
	cairo_user_font_face_set_init_func               = cc.DlFunc[func(font_face *UserFontFace, init_func uptr), cc.Void]{Name: "cairo_user_font_face_set_init_func"}
	cairo_user_font_face_set_render_glyph_func       = cc.DlFunc[func(font_face *UserFontFace, render_glyph_func uptr), cc.Void]{Name: "cairo_user_font_face_set_render_glyph_func"}
	cairo_user_font_face_set_render_color_glyph_func = cc.DlFunc[func(font_face *UserFontFace, render_glyph_func uptr), cc.Void]{Name: "cairo_user_font_face_set_render_color_glyph_func"}
	cairo_user_font_face_set_text_to_glyphs_func     = cc.DlFunc[func(font_face *UserFontFace, text_to_glyphs_func uptr), cc.Void]{Name: "cairo_user_font_face_set_text_to_glyphs_func"}
	cairo_user_font_face_set_unicode_to_glyph_func   = cc.DlFunc[func(font_face *UserFontFace, unicode_to_glyph_func uptr), cc.Void]{Name: "cairo_user_font_face_set_unicode_to_glyph_func"}
	cairo_user_font_face_get_init_func               = cc.DlFunc[func(font_face *UserFontFace) cc.Func, cc.Func]{Name: "cairo_user_font_face_get_init_func"}
	cairo_user_font_face_get_render_glyph_func       = cc.DlFunc[func(font_face *UserFontFace) cc.Func, cc.Func]{Name: "cairo_user_font_face_get_render_glyph_func"}
	cairo_user_font_face_get_render_color_glyph_func = cc.DlFunc[func(font_face *UserFontFace) cc.Func, cc.Func]{Name: "cairo_user_font_face_get_render_color_glyph_func"}
	cairo_user_font_face_get_text_to_glyphs_func     = cc.DlFunc[func(font_face *UserFontFace) cc.Func, cc.Func]{Name: "cairo_user_font_face_get_text_to_glyphs_func"}
	cairo_user_font_face_get_unicode_to_glyph_func   = cc.DlFunc[func(font_face *UserFontFace) cc.Func, cc.Func]{Name: "cairo_user_font_face_get_unicode_to_glyph_func"}

	cairo_user_scaled_font_get_foreground_marker = cc.DlFunc[func(scaled_font *UserScaledFont) *Pattern, *Pattern]{Name: "cairo_user_scaled_font_get_foreground_marker"}
	cairo_user_scaled_font_get_foreground_source = cc.DlFunc[func(scaled_font *UserScaledFont) *Pattern, *Pattern]{Name: "cairo_user_scaled_font_get_foreground_source"}

	cairo_get_operator      = cc.DlFunc[func(cr *Context) Operator, Operator]{Name: "cairo_get_operator"}
	cairo_get_source        = cc.DlFunc[func(cr *Context) *Pattern, *Pattern]{Name: "cairo_get_source"}
	cairo_get_tolerance     = cc.DlFunc[func(cr *Context) float64, float64]{Name: "cairo_get_tolerance"}
	cairo_get_antialias     = cc.DlFunc[func(cr *Context) Antialias, Antialias]{Name: "cairo_get_antialias"}
	cairo_has_current_point = cc.DlFunc[func(cr *Context) Bool, Bool]{Name: "cairo_has_current_point"}
	cairo_get_current_point = cc.DlFunc[func(cr *Context, x, y *float64), cc.Void]{Name: "cairo_get_current_point"}
	cairo_get_fill_rule     = cc.DlFunc[func(cr *Context) FillRule, FillRule]{Name: "cairo_get_fill_rule"}
	cairo_get_line_width    = cc.DlFunc[func(cr *Context) float64, float64]{Name: "cairo_get_line_width"}
	cairo_get_hairline      = cc.DlFunc[func(cr *Context) Bool, Bool]{Name: "cairo_get_hairline"}
	cairo_get_line_cap      = cc.DlFunc[func(cr *Context) LineCap, LineCap]{Name: "cairo_get_line_cap"}
	cairo_get_line_join     = cc.DlFunc[func(cr *Context) LineJoin, LineJoin]{Name: "cairo_get_line_join"}
	cairo_get_miter_limit   = cc.DlFunc[func(cr *Context) float64, float64]{Name: "cairo_get_miter_limit"}
	cairo_get_dash_count    = cc.DlFunc[func(cr *Context) int32, int32]{Name: "cairo_get_dash_count"}
	cairo_get_dash          = cc.DlFunc[func(cr *Context, dashes, offset *float64), cc.Void]{Name: "cairo_get_dash"}
	cairo_get_matrix        = cc.DlFunc[func(cr *Context, matrix *Matrix), cc.Void]{Name: "cairo_get_matrix"}
	cairo_get_target        = cc.DlFunc[func(cr *Context) *Surface, *Surface]{Name: "cairo_get_target"}
	cairo_get_group_target  = cc.DlFunc[func(cr *Context) *Surface, *Surface]{Name: "cairo_get_group_target"}

	cairo_copy_path        = cc.DlFunc[func(cr *Context) *Path, *Path]{Name: "cairo_copy_path"}
	cairo_copy_path_flat   = cc.DlFunc[func(cr *Context) *Path, *Path]{Name: "cairo_copy_path_flat"}
	cairo_append_path      = cc.DlFunc[func(cr *Context, path *Path), cc.Void]{Name: "cairo_append_path"}
	cairo_path_destroy     = cc.DlFunc[func(path *Path), cc.Void]{Name: "cairo_path_destroy"}
	cairo_status           = cc.DlFunc[func(cr *Context) Status, Status]{Name: "cairo_status"}
	cairo_status_to_string = cc.DlFunc[func(status Status) cc.String, cc.String]{Name: "cairo_status_to_string"}
	cairo_device_reference = cc.DlFunc[func(device *Device) *Device, *Device]{Name: "cairo_device_reference"}

	cairo_device_get_type            = cc.DlFunc[func(device *Device) DeviceType, DeviceType]{Name: "cairo_device_get_type"}
	cairo_device_status              = cc.DlFunc[func(device *Device) Status, Status]{Name: "cairo_device_status"}
	cairo_device_acquire             = cc.DlFunc[func(device *Device) Status, Status]{Name: "cairo_device_acquire"}
	cairo_device_release             = cc.DlFunc[func(device *Device), cc.Void]{Name: "cairo_device_release"}
	cairo_device_flush               = cc.DlFunc[func(device *Device), cc.Void]{Name: "cairo_device_flush"}
	cairo_device_finish              = cc.DlFunc[func(device *Device), cc.Void]{Name: "cairo_device_finish"}
	cairo_device_destroy             = cc.DlFunc[func(device *Device), cc.Void]{Name: "cairo_device_destroy"}
	cairo_device_get_reference_count = cc.DlFunc[func(device *Device) uint32, uint32]{Name: "cairo_device_get_reference_count"}
	cairo_device_get_user_data       = cc.DlFunc[func(device *Device, key *UserDataKey) uptr, uptr]{Name: "cairo_device_get_user_data"}
	cairo_device_set_user_data       = cc.DlFunc[func(device *Device, key *UserDataKey, user_data uptr, destroy uptr) Status, Status]{Name: "cairo_device_set_user_data"}

	cairo_surface_create_similar       = cc.DlFunc[func(other *Surface, content int32, width, height int32) *Surface, *Surface]{Name: "cairo_surface_create_similar"}
	cairo_surface_create_similar_image = cc.DlFunc[func(other *Surface, format int32, width, height int32) *Surface, *Surface]{Name: "cairo_surface_create_similar_image"}
	cairo_surface_map_to_image         = cc.DlFunc[func(surface *Surface, extents *RectangleInt) *Surface, *Surface]{Name: "cairo_surface_map_to_image"}
	cairo_surface_unmap_image          = cc.DlFunc[func(surface *Surface, image *Surface), cc.Void]{Name: "cairo_surface_unmap_image"}
	cairo_surface_create_for_rectangle = cc.DlFunc[func(target *Surface, x, y, width, height float64) *Surface, *Surface]{Name: "cairo_surface_create_for_rectangle"}

	cairo_surface_create_observer              = cc.DlFunc[func(target *Surface, mode SurfaceObserverMode) *SurfaceObserver, *SurfaceObserver]{Name: "cairo_surface_create_observer"}
	cairo_surface_observer_add_paint_callback  = cc.DlFunc[func(abstract_surface *SurfaceObserver, fn uptr, data uptr) Status, Status]{Name: "cairo_surface_observer_add_paint_callback"}
	cairo_surface_observer_add_mask_callback   = cc.DlFunc[func(abstract_surface *SurfaceObserver, fn uptr, data uptr) Status, Status]{Name: "cairo_surface_observer_add_mask_callback"}
	cairo_surface_observer_add_fill_callback   = cc.DlFunc[func(abstract_surface *SurfaceObserver, fn uptr, data uptr) Status, Status]{Name: "cairo_surface_observer_add_fill_callback"}
	cairo_surface_observer_add_stroke_callback = cc.DlFunc[func(abstract_surface *SurfaceObserver, fn uptr, data uptr) Status, Status]{Name: "cairo_surface_observer_add_stroke_callback"}
	cairo_surface_observer_add_glyphs_callback = cc.DlFunc[func(abstract_surface *SurfaceObserver, fn uptr, data uptr) Status, Status]{Name: "cairo_surface_observer_add_glyphs_callback"}
	cairo_surface_observer_add_flush_callback  = cc.DlFunc[func(abstract_surface *SurfaceObserver, fn uptr, data uptr) Status, Status]{Name: "cairo_surface_observer_add_flush_callback"}
	cairo_surface_observer_add_finish_callback = cc.DlFunc[func(abstract_surface *SurfaceObserver, fn uptr, data uptr) Status, Status]{Name: "cairo_surface_observer_add_finish_callback"}
	cairo_surface_observer_print               = cc.DlFunc[func(abstract_surface *SurfaceObserver, write_func uptr, closure uptr) Status, Status]{Name: "cairo_surface_observer_print"}
	cairo_surface_observer_elapsed             = cc.DlFunc[func(abstract_surface *SurfaceObserver) float64, float64]{Name: "cairo_surface_observer_elapsed"}

	cairo_device_observer_print          = cc.DlFunc[func(abstract_device *Device, write_func uptr, closure uptr) Status, Status]{Name: "cairo_device_observer_print"}
	cairo_device_observer_elapsed        = cc.DlFunc[func(abstract_device *Device) float64, float64]{Name: "cairo_device_observer_elapsed"}
	cairo_device_observer_paint_elapsed  = cc.DlFunc[func(abstract_device *Device) float64, float64]{Name: "cairo_device_observer_paint_elapsed"}
	cairo_device_observer_mask_elapsed   = cc.DlFunc[func(abstract_device *Device) float64, float64]{Name: "cairo_device_observer_mask_elapsed"}
	cairo_device_observer_fill_elapsed   = cc.DlFunc[func(abstract_device *Device) float64, float64]{Name: "cairo_device_observer_fill_elapsed"}
	cairo_device_observer_stroke_elapsed = cc.DlFunc[func(abstract_device *Device) float64, float64]{Name: "cairo_device_observer_stroke_elapsed"}
	cairo_device_observer_glyphs_elapsed = cc.DlFunc[func(abstract_device *Device) float64, float64]{Name: "cairo_device_observer_glyphs_elapsed"}

	cairo_surface_reference           = cc.DlFunc[func(surface *Surface) *Surface, *Surface]{Name: "cairo_surface_reference"}
	cairo_surface_finish              = cc.DlFunc[func(surface *Surface), cc.Void]{Name: "cairo_surface_finish"}
	cairo_surface_destroy             = cc.DlFunc[func(surface *Surface), cc.Void]{Name: "cairo_surface_destroy"}
	cairo_surface_get_device          = cc.DlFunc[func(surface *Surface) *Device, *Device]{Name: "cairo_surface_get_device"}
	cairo_surface_get_reference_count = cc.DlFunc[func(surface *Surface) uint32, uint32]{Name: "cairo_surface_get_reference_count"}
	cairo_surface_status              = cc.DlFunc[func(surface *Surface) Status, Status]{Name: "cairo_surface_status"}

	cairo_surface_get_type                = cc.DlFunc[func(surface *Surface) int32, int32]{Name: "cairo_surface_get_type"}
	cairo_surface_get_content             = cc.DlFunc[func(surface *Surface) int32, int32]{Name: "cairo_surface_get_content"}
	cairo_surface_write_to_png            = cc.DlFunc[func(surface *Surface, filename cc.String) Status, Status]{Name: "cairo_surface_write_to_png"}
	cairo_surface_write_to_png_stream     = cc.DlFunc[func(surface *Surface, write_func uptr, closure uptr) Status, Status]{Name: "cairo_surface_write_to_png_stream"}
	cairo_surface_get_user_data           = cc.DlFunc[func(surface *Surface, key *UserDataKey) uptr, uptr]{Name: "cairo_surface_get_user_data"}
	cairo_surface_set_user_data           = cc.DlFunc[func(surface *Surface, key *UserDataKey, user_data uptr, destroy uptr) Status, Status]{Name: "cairo_surface_set_user_data"}
	cairo_surface_get_mime_data           = cc.DlFunc[func(surface *Surface, mime_type cc.String, data **byte, length *uint64), cc.Void]{Name: "cairo_surface_get_mime_data"}
	cairo_surface_set_mime_data           = cc.DlFunc[func(surface *Surface, mime_type cc.String, data *byte, length uint64, destroy uptr, closure uptr) Status, Status]{Name: "cairo_surface_set_mime_data"}
	cairo_surface_supports_mime_type      = cc.DlFunc[func(surface *Surface, mime_type cc.String) Bool, Bool]{Name: "cairo_surface_supports_mime_type"}
	cairo_surface_get_font_options        = cc.DlFunc[func(surface *Surface, options *FontOptions), cc.Void]{Name: "cairo_surface_get_font_options"}
	cairo_surface_flush                   = cc.DlFunc[func(surface *Surface), cc.Void]{Name: "cairo_surface_flush"}
	cairo_surface_mark_dirty              = cc.DlFunc[func(surface *Surface), cc.Void]{Name: "cairo_surface_mark_dirty"}
	cairo_surface_mark_dirty_rectangle    = cc.DlFunc[func(surface *Surface, x, y, width, height int32), cc.Void]{Name: "cairo_surface_mark_dirty_rectangle"}
	cairo_surface_set_device_scale        = cc.DlFunc[func(surface *Surface, x_scale, y_scale float64), cc.Void]{Name: "cairo_surface_set_device_scale"}
	cairo_surface_get_device_scale        = cc.DlFunc[func(surface *Surface, x_scale, y_scale *float64), cc.Void]{Name: "cairo_surface_get_device_scale"}
	cairo_surface_set_device_offset       = cc.DlFunc[func(surface *Surface, x_offset, y_offset float64), cc.Void]{Name: "cairo_surface_set_device_offset"}
	cairo_surface_get_device_offset       = cc.DlFunc[func(surface *Surface, x_offset, y_offset *float64), cc.Void]{Name: "cairo_surface_get_device_offset"}
	cairo_surface_set_fallback_resolution = cc.DlFunc[func(surface *Surface, x_ppi, y_ppi float64), cc.Void]{Name: "cairo_surface_set_fallback_resolution"}
	cairo_surface_get_fallback_resolution = cc.DlFunc[func(surface *Surface, x_ppi, y_ppi *float64), cc.Void]{Name: "cairo_surface_get_fallback_resolution"}
	cairo_surface_copy_page               = cc.DlFunc[func(surface *Surface), cc.Void]{Name: "cairo_surface_copy_page"}
	cairo_surface_show_page               = cc.DlFunc[func(surface *Surface), cc.Void]{Name: "cairo_surface_show_page"}
	cairo_surface_has_show_text_glyphs    = cc.DlFunc[func(surface *Surface) Bool, Bool]{Name: "cairo_surface_has_show_text_glyphs"}

	cairo_image_surface_create                 = cc.DlFunc[func(format Format, width, height int32) *ImageSurface, *ImageSurface]{Name: "cairo_image_surface_create"}
	cairo_format_stride_for_width              = cc.DlFunc[func(format Format, width int32) int32, int32]{Name: "cairo_format_stride_for_width"}
	cairo_image_surface_create_for_data        = cc.DlFunc[func(data *byte, format Format, width, height, stride int32) *ImageSurface, *ImageSurface]{Name: "cairo_image_surface_create_for_data"}
	cairo_image_surface_get_data               = cc.DlFunc[func(surface *ImageSurface) *byte, *byte]{Name: "cairo_image_surface_get_data"}
	cairo_image_surface_get_format             = cc.DlFunc[func(surface *ImageSurface) int32, int32]{Name: "cairo_image_surface_get_format"}
	cairo_image_surface_get_width              = cc.DlFunc[func(surface *ImageSurface) int32, int32]{Name: "cairo_image_surface_get_width"}
	cairo_image_surface_get_height             = cc.DlFunc[func(surface *ImageSurface) int32, int32]{Name: "cairo_image_surface_get_height"}
	cairo_image_surface_get_stride             = cc.DlFunc[func(surface *ImageSurface) int32, int32]{Name: "cairo_image_surface_get_stride"}
	cairo_image_surface_create_from_png        = cc.DlFunc[func(filename cc.String) *ImageSurface, *ImageSurface]{Name: "cairo_image_surface_create_from_png"}
	cairo_image_surface_create_from_png_stream = cc.DlFunc[func(read_func uptr, closure uptr) *ImageSurface, *ImageSurface]{Name: "cairo_image_surface_create_from_png_stream"}

	cairo_recording_surface_create      = cc.DlFunc[func(content int32, extents *Rectangle) *RecordingSurface, *RecordingSurface]{Name: "cairo_recording_surface_create"}
	cairo_recording_surface_ink_extents = cc.DlFunc[func(surface *RecordingSurface, x0, y0, width, height *float64), cc.Void]{Name: "cairo_recording_surface_ink_extents"}
	cairo_recording_surface_get_extents = cc.DlFunc[func(surface *RecordingSurface, extents *Rectangle) Bool, Bool]{Name: "cairo_recording_surface_get_extents"}

	cairo_pattern_create_raster_source            = cc.DlFunc[func(user_data uptr, content Content, width, height int32) *RasterSourcePattern, *RasterSourcePattern]{Name: "cairo_pattern_create_raster_source"}
	cairo_raster_source_pattern_set_callback_data = cc.DlFunc[func(pattern *RasterSourcePattern, data uptr), cc.Void]{Name: "cairo_raster_source_pattern_set_callback_data"}
	cairo_raster_source_pattern_get_callback_data = cc.DlFunc[func(pattern *RasterSourcePattern) uptr, uptr]{Name: "cairo_raster_source_pattern_get_callback_data"}
	cairo_raster_source_pattern_set_acquire       = cc.DlFunc[func(pattern *RasterSourcePattern, acquire uptr, release uptr), cc.Void]{Name: "cairo_raster_source_pattern_set_acquire"}
	cairo_raster_source_pattern_get_acquire       = cc.DlFunc[func(pattern *RasterSourcePattern, acquire *cc.Func, release *cc.Func), cc.Void]{Name: "cairo_raster_source_pattern_get_acquire"}
	cairo_raster_source_pattern_set_snapshot      = cc.DlFunc[func(pattern *RasterSourcePattern, snapshot uptr), cc.Void]{Name: "cairo_raster_source_pattern_set_snapshot"}
	cairo_raster_source_pattern_get_snapshot      = cc.DlFunc[func(pattern *RasterSourcePattern) cc.Func, cc.Func]{Name: "cairo_raster_source_pattern_get_snapshot"}
	cairo_raster_source_pattern_set_copy          = cc.DlFunc[func(pattern *RasterSourcePattern, copy uptr), cc.Void]{Name: "cairo_raster_source_pattern_set_copy"}
	cairo_raster_source_pattern_get_copy          = cc.DlFunc[func(pattern *RasterSourcePattern) cc.Func, cc.Func]{Name: "cairo_raster_source_pattern_get_copy"}
	cairo_raster_source_pattern_set_finish        = cc.DlFunc[func(pattern *RasterSourcePattern, finish uptr), cc.Void]{Name: "cairo_raster_source_pattern_set_finish"}
	cairo_raster_source_pattern_get_finish        = cc.DlFunc[func(pattern *RasterSourcePattern) cc.Func, cc.Func]{Name: "cairo_raster_source_pattern_get_finish"}

	cairo_pattern_create_rgb          = cc.DlFunc[func(red, green, blue float64) *SolidPattern, *SolidPattern]{Name: "cairo_pattern_create_rgb"}
	cairo_pattern_create_rgba         = cc.DlFunc[func(red, green, blue, alpha float64) *SolidPattern, *SolidPattern]{Name: "cairo_pattern_create_rgba"}
	cairo_pattern_create_for_surface  = cc.DlFunc[func(surface *Surface) *Pattern, *Pattern]{Name: "cairo_pattern_create_for_surface"}
	cairo_pattern_create_linear       = cc.DlFunc[func(x0, y0, x1, y1 float64) *Pattern, *Pattern]{Name: "cairo_pattern_create_linear"}
	cairo_pattern_create_radial       = cc.DlFunc[func(cx0, cy0, radius0, cx1, cy1, radius1 float64) *Pattern, *Pattern]{Name: "cairo_pattern_create_radial"}
	cairo_pattern_create_mesh         = cc.DlFunc[func() *Pattern, *Pattern]{Name: "cairo_pattern_create_mesh"}
	cairo_pattern_reference           = cc.DlFunc[func(pattern *Pattern) *Pattern, *Pattern]{Name: "cairo_pattern_reference"}
	cairo_pattern_destroy             = cc.DlFunc[func(pattern *Pattern), cc.Void]{Name: "cairo_pattern_destroy"}
	cairo_pattern_get_reference_count = cc.DlFunc[func(pattern *Pattern) uint32, uint32]{Name: "cairo_pattern_get_reference_count"}
	cairo_pattern_status              = cc.DlFunc[func(pattern *Pattern) Status, Status]{Name: "cairo_pattern_status"}
	cairo_pattern_get_user_data       = cc.DlFunc[func(pattern *Pattern, key *UserDataKey) uptr, uptr]{Name: "cairo_pattern_get_user_data"}
	cairo_pattern_set_user_data       = cc.DlFunc[func(pattern *Pattern, key *UserDataKey, user_data uptr, destroy uptr) Status, Status]{Name: "cairo_pattern_set_user_data"}

	cairo_pattern_get_type                   = cc.DlFunc[func(pattern *Pattern) PatternType, PatternType]{Name: "cairo_pattern_get_type"}
	cairo_pattern_add_color_stop_rgb         = cc.DlFunc[func(pattern *GradientPattern, offset, red, green, blue float64), cc.Void]{Name: "cairo_pattern_add_color_stop_rgb"}
	cairo_pattern_add_color_stop_rgba        = cc.DlFunc[func(pattern *GradientPattern, offset, red, green, blue, alpha float64), cc.Void]{Name: "cairo_pattern_add_color_stop_rgba"}
	cairo_mesh_pattern_begin_patch           = cc.DlFunc[func(pattern *MeshPattern), cc.Void]{Name: "cairo_mesh_pattern_begin_patch"}
	cairo_mesh_pattern_end_patch             = cc.DlFunc[func(pattern *MeshPattern), cc.Void]{Name: "cairo_mesh_pattern_end_patch"}
	cairo_mesh_pattern_curve_to              = cc.DlFunc[func(pattern *MeshPattern, x1, y1, x2, y2, x3, y3 float64), cc.Void]{Name: "cairo_mesh_pattern_curve_to"}
	cairo_mesh_pattern_line_to               = cc.DlFunc[func(pattern *MeshPattern, x, y float64), cc.Void]{Name: "cairo_mesh_pattern_line_to"}
	cairo_mesh_pattern_move_to               = cc.DlFunc[func(pattern *MeshPattern, x, y float64), cc.Void]{Name: "cairo_mesh_pattern_move_to"}
	cairo_mesh_pattern_set_control_point     = cc.DlFunc[func(pattern *MeshPattern, point_num uint32, x, y float64), cc.Void]{Name: "cairo_mesh_pattern_set_control_point"}
	cairo_mesh_pattern_set_corner_color_rgb  = cc.DlFunc[func(pattern *MeshPattern, corner_num uint32, red, green, blue float64), cc.Void]{Name: "cairo_mesh_pattern_set_corner_color_rgb"}
	cairo_mesh_pattern_set_corner_color_rgba = cc.DlFunc[func(pattern *MeshPattern, corner_num uint32, red, green, blue, alpha float64), cc.Void]{Name: "cairo_mesh_pattern_set_corner_color_rgba"}
	cairo_pattern_set_matrix                 = cc.DlFunc[func(pattern *Pattern, matrix *Matrix), cc.Void]{Name: "cairo_pattern_set_matrix"}
	cairo_pattern_get_matrix                 = cc.DlFunc[func(pattern *Pattern, matrix *Matrix), cc.Void]{Name: "cairo_pattern_get_matrix"}

	cairo_pattern_set_extend                 = cc.DlFunc[func(pattern *Pattern, extend Extend), cc.Void]{Name: "cairo_pattern_set_extend"}
	cairo_pattern_get_extend                 = cc.DlFunc[func(pattern *Pattern) Extend, Extend]{Name: "cairo_pattern_get_extend"}
	cairo_pattern_set_filter                 = cc.DlFunc[func(pattern *Pattern, filter Filter), cc.Void]{Name: "cairo_pattern_set_filter"}
	cairo_pattern_get_filter                 = cc.DlFunc[func(pattern *Pattern) Filter, Filter]{Name: "cairo_pattern_get_filter"}
	cairo_pattern_get_rgba                   = cc.DlFunc[func(pattern *SolidPattern, red, green, blue, alpha *float64) Status, Status]{Name: "cairo_pattern_get_rgba"}
	cairo_pattern_get_surface                = cc.DlFunc[func(pattern *SurfacePattern, surface **Surface) Status, Status]{Name: "cairo_pattern_get_surface"}
	cairo_pattern_get_color_stop_rgba        = cc.DlFunc[func(pattern *GradientPattern, index int32, offset, red, green, blue, alpha *float64) Status, Status]{Name: "cairo_pattern_get_color_stop_rgba"}
	cairo_pattern_get_color_stop_count       = cc.DlFunc[func(pattern *GradientPattern, count *int32) Status, Status]{Name: "cairo_pattern_get_color_stop_count"}
	cairo_pattern_get_linear_points          = cc.DlFunc[func(pattern *LinearPattern, x0, y0, x1, y1 *float64) Status, Status]{Name: "cairo_pattern_get_linear_points"}
	cairo_pattern_get_radial_circles         = cc.DlFunc[func(pattern *RadialPattern, x0, y0, r0, x1, y1, r1 *float64) Status, Status]{Name: "cairo_pattern_get_radial_circles"}
	cairo_mesh_pattern_get_patch_count       = cc.DlFunc[func(pattern *MeshPattern, count *uint32) Status, Status]{Name: "cairo_mesh_pattern_get_patch_count"}
	cairo_mesh_pattern_get_path              = cc.DlFunc[func(pattern *MeshPattern, patch_num uint32) *Path, *Path]{Name: "cairo_mesh_pattern_get_path"}
	cairo_mesh_pattern_get_corner_color_rgba = cc.DlFunc[func(pattern *MeshPattern, patch_num, corner_num uint32, red, green, blue, alpha *float64) Status, Status]{Name: "cairo_mesh_pattern_get_corner_color_rgba"}
	cairo_mesh_pattern_get_control_point     = cc.DlFunc[func(pattern *MeshPattern, patch_num, point_num uint32, x, y *float64) Status, Status]{Name: "cairo_mesh_pattern_get_control_point"}

	cairo_matrix_init               = cc.DlFunc[func(matrix *Matrix, xx, yx, xy, yy, x0, y0 float64), cc.Void]{Name: "cairo_matrix_init"}
	cairo_matrix_init_identity      = cc.DlFunc[func(matrix *Matrix), cc.Void]{Name: "cairo_matrix_init_identity"}
	cairo_matrix_init_translate     = cc.DlFunc[func(matrix *Matrix, tx, ty float64), cc.Void]{Name: "cairo_matrix_init_translate"}
	cairo_matrix_init_scale         = cc.DlFunc[func(matrix *Matrix, sx, sy float64), cc.Void]{Name: "cairo_matrix_init_scale"}
	cairo_matrix_init_rotate        = cc.DlFunc[func(matrix *Matrix, radians float64), cc.Void]{Name: "cairo_matrix_init_rotate"}
	cairo_matrix_translate          = cc.DlFunc[func(matrix *Matrix, tx, ty float64), cc.Void]{Name: "cairo_matrix_translate"}
	cairo_matrix_scale              = cc.DlFunc[func(matrix *Matrix, sx, sy float64), cc.Void]{Name: "cairo_matrix_scale"}
	cairo_matrix_rotate             = cc.DlFunc[func(matrix *Matrix, radians float64), cc.Void]{Name: "cairo_matrix_rotate"}
	cairo_matrix_invert             = cc.DlFunc[func(matrix *Matrix) Status, Status]{Name: "cairo_matrix_invert"}
	cairo_matrix_multiply           = cc.DlFunc[func(result, a, b *Matrix), cc.Void]{Name: "cairo_matrix_multiply"}
	cairo_matrix_transform_distance = cc.DlFunc[func(matrix *Matrix, dx, dy *float64), cc.Void]{Name: "cairo_matrix_transform_distance"}
	cairo_matrix_transform_point    = cc.DlFunc[func(matrix *Matrix, x, y *float64), cc.Void]{Name: "cairo_matrix_transform_point"}

	cairo_region_create              = cc.DlFunc[func() *Region, *Region]{Name: "cairo_region_create"}
	cairo_region_create_rectangle    = cc.DlFunc[func(rectangle *RectangleInt) *Region, *Region]{Name: "cairo_region_create_rectangle"}
	cairo_region_create_rectangles   = cc.DlFunc[func(rects *RectangleInt, count int32) *Region, *Region]{Name: "cairo_region_create_rectangles"}
	cairo_region_copy                = cc.DlFunc[func(original *Region) *Region, *Region]{Name: "cairo_region_copy"}
	cairo_region_reference           = cc.DlFunc[func(region *Region) *Region, *Region]{Name: "cairo_region_reference"}
	cairo_region_destroy             = cc.DlFunc[func(region *Region), cc.Void]{Name: "cairo_region_destroy"}
	cairo_region_equal               = cc.DlFunc[func(a, b *Region) Bool, Bool]{Name: "cairo_region_equal"}
	cairo_region_status              = cc.DlFunc[func(region *Region) Status, Status]{Name: "cairo_region_status"}
	cairo_region_get_extents         = cc.DlFunc[func(region *Region, extents *RectangleInt), cc.Void]{Name: "cairo_region_get_extents"}
	cairo_region_num_rectangles      = cc.DlFunc[func(region *Region) int32, int32]{Name: "cairo_region_num_rectangles"}
	cairo_region_get_rectangle       = cc.DlFunc[func(region *Region, nth int32, rectangle *RectangleInt), cc.Void]{Name: "cairo_region_get_rectangle"}
	cairo_region_is_empty            = cc.DlFunc[func(region *Region) Bool, Bool]{Name: "cairo_region_is_empty"}
	cairo_region_contains_rectangle  = cc.DlFunc[func(region *Region, rectangle *RectangleInt) int32, int32]{Name: "cairo_region_contains_rectangle"}
	cairo_region_contains_point      = cc.DlFunc[func(region *Region, x, y int32) Bool, Bool]{Name: "cairo_region_contains_point"}
	cairo_region_translate           = cc.DlFunc[func(region *Region, dx, dy int32), cc.Void]{Name: "cairo_region_translate"}
	cairo_region_subtract            = cc.DlFunc[func(dst, other *Region) Status, Status]{Name: "cairo_region_subtract"}
	cairo_region_subtract_rectangle  = cc.DlFunc[func(dst *Region, rectangle *RectangleInt) Status, Status]{Name: "cairo_region_subtract_rectangle"}
	cairo_region_intersect           = cc.DlFunc[func(dst, other *Region) Status, Status]{Name: "cairo_region_intersect"}
	cairo_region_intersect_rectangle = cc.DlFunc[func(dst *Region, rectangle *RectangleInt) Status, Status]{Name: "cairo_region_intersect_rectangle"}
	cairo_region_union               = cc.DlFunc[func(dst, other *Region) Status, Status]{Name: "cairo_region_union"}
	cairo_region_union_rectangle     = cc.DlFunc[func(dst *Region, rectangle *RectangleInt) Status, Status]{Name: "cairo_region_union_rectangle"}
	cairo_region_xor                 = cc.DlFunc[func(dst, other *Region) Status, Status]{Name: "cairo_region_xor"}
	cairo_region_xor_rectangle       = cc.DlFunc[func(dst *Region, rectangle *RectangleInt) Status, Status]{Name: "cairo_region_xor_rectangle"}

	cairo_debug_reset_static_data = cc.DlFunc[func(), cc.Void]{Name: "cairo_debug_reset_static_data"}
	// #endregion

	// #region cairo-gobject.h
	cairo_gobject_context_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_context_get_type"}
	cairo_gobject_device_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_device_get_type"}
	cairo_gobject_matrix_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_matrix_get_type"}
	cairo_gobject_pattern_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_pattern_get_type"}
	cairo_gobject_surface_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_surface_get_type"}
	cairo_gobject_rectangle_get_type     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_rectangle_get_type"}
	cairo_gobject_scaled_font_get_type   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_scaled_font_get_type"}
	cairo_gobject_font_face_get_type     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_font_face_get_type"}
	cairo_gobject_font_options_get_type  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_font_options_get_type"}
	cairo_gobject_rectangle_int_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_rectangle_int_get_type"}
	cairo_gobject_region_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_region_get_type"}
	cairo_gobject_glyph_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_glyph_get_type"}
	cairo_gobject_text_cluster_get_type  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_text_cluster_get_type"}

	cairo_gobject_status_get_type             = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_status_get_type"}
	cairo_gobject_content_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_content_get_type"}
	cairo_gobject_operator_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_operator_get_type"}
	cairo_gobject_antialias_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_antialias_get_type"}
	cairo_gobject_fill_rule_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_fill_rule_get_type"}
	cairo_gobject_line_cap_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_line_cap_get_type"}
	cairo_gobject_line_join_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_line_join_get_type"}
	cairo_gobject_text_cluster_flags_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_text_cluster_flags_get_type"}
	cairo_gobject_font_slant_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_font_slant_get_type"}
	cairo_gobject_font_weight_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_font_weight_get_type"}
	cairo_gobject_subpixel_order_get_type     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_subpixel_order_get_type"}
	cairo_gobject_hint_style_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_hint_style_get_type"}
	cairo_gobject_hint_metrics_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_hint_metrics_get_type"}
	cairo_gobject_font_type_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_font_type_get_type"}
	cairo_gobject_path_data_type_get_type     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_path_data_type_get_type"}
	cairo_gobject_device_type_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_device_type_get_type"}
	cairo_gobject_surface_type_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_surface_type_get_type"}
	cairo_gobject_format_get_type             = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_format_get_type"}
	cairo_gobject_pattern_type_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_pattern_type_get_type"}
	cairo_gobject_extend_get_type             = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_extend_get_type"}
	cairo_gobject_filter_get_type             = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_filter_get_type"}
	cairo_gobject_region_overlap_get_type     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "cairo_gobject_region_overlap_get_type"}
	// #endregion

	// #region cairo-pdf.h
	cairo_pdf_surface_create              = cc.DlFunc[func(filename cc.String, width_in_points, height_in_points float64) *PDFSurface, *PDFSurface]{Name: "cairo_pdf_surface_create"}
	cairo_pdf_surface_create_for_stream   = cc.DlFunc[func(write_func uptr, closure uptr, width_in_points, height_in_points float64) *Surface, *Surface]{Name: "cairo_pdf_surface_create_for_stream"}
	cairo_pdf_surface_restrict_to_version = cc.DlFunc[func(surface *PDFSurface, version PDFVersion), cc.Void]{Name: "cairo_pdf_surface_restrict_to_version"}
	cairo_pdf_get_versions                = cc.DlFunc[func(versions **PDFVersion, num_versions *int32), cc.Void]{Name: "cairo_pdf_get_versions"}
	cairo_pdf_version_to_string           = cc.DlFunc[func(version PDFVersion) cc.String, cc.String]{Name: "cairo_pdf_version_to_string"}
	cairo_pdf_surface_set_size            = cc.DlFunc[func(surface *PDFSurface, width_in_points, height_in_points float64), cc.Void]{Name: "cairo_pdf_surface_set_size"}
	cairo_pdf_surface_add_outline         = cc.DlFunc[func(surface *PDFSurface, parent_id int32, utf8 cc.String, link_attribs cc.String, flags PDFOutlineFlags) int32, int32]{Name: "cairo_pdf_surface_add_outline"}
	cairo_pdf_surface_set_metadata        = cc.DlFunc[func(surface *PDFSurface, metadata PDFMetadata, utf8 cc.String), cc.Void]{Name: "cairo_pdf_surface_set_metadata"}
	cairo_pdf_surface_set_custom_metadata = cc.DlFunc[func(surface *PDFSurface, name cc.String, value cc.String), cc.Void]{Name: "cairo_pdf_surface_set_custom_metadata"}
	cairo_pdf_surface_set_page_label      = cc.DlFunc[func(surface *PDFSurface, utf8 cc.String), cc.Void]{Name: "cairo_pdf_surface_set_page_label"}
	cairo_pdf_surface_set_thumbnail_size  = cc.DlFunc[func(surface *PDFSurface, width, height int32), cc.Void]{Name: "cairo_pdf_surface_set_thumbnail_size"}
	// #endregion

	// #region cairo-ps.h
	cairo_ps_surface_create               = cc.DlFunc[func(filename cc.String, width_in_points, height_in_points float64) *PSSurface, *PSSurface]{Name: "cairo_ps_surface_create"}
	cairo_ps_surface_create_for_stream    = cc.DlFunc[func(write_func uptr, closure uptr, width_in_points, height_in_points float64) *PSSurface, *PSSurface]{Name: "cairo_ps_surface_create_for_stream"}
	cairo_ps_surface_restrict_to_level    = cc.DlFunc[func(surface *PSSurface, level PSLevel), cc.Void]{Name: "cairo_ps_surface_restrict_to_level"}
	cairo_ps_get_levels                   = cc.DlFunc[func(levels **PSLevel, num_levels *int32), cc.Void]{Name: "cairo_ps_get_levels"}
	cairo_ps_level_to_string              = cc.DlFunc[func(level PSLevel) cc.String, cc.String]{Name: "cairo_ps_level_to_string"}
	cairo_ps_surface_set_eps              = cc.DlFunc[func(surface *PSSurface, eps Bool), cc.Void]{Name: "cairo_ps_surface_set_eps"}
	cairo_ps_surface_get_eps              = cc.DlFunc[func(surface *PSSurface) Bool, Bool]{Name: "cairo_ps_surface_get_eps"}
	cairo_ps_surface_set_size             = cc.DlFunc[func(surface *PSSurface, width_in_points, height_in_points float64), cc.Void]{Name: "cairo_ps_surface_set_size"}
	cairo_ps_surface_dsc_comment          = cc.DlFunc[func(surface *PSSurface, comment cc.String), cc.Void]{Name: "cairo_ps_surface_dsc_comment"}
	cairo_ps_surface_dsc_begin_setup      = cc.DlFunc[func(surface *PSSurface), cc.Void]{Name: "cairo_ps_surface_dsc_begin_setup"}
	cairo_ps_surface_dsc_begin_page_setup = cc.DlFunc[func(surface *PSSurface), cc.Void]{Name: "cairo_ps_surface_dsc_begin_page_setup"}

	// #endregion

	// #region cairo-script.h
	cairo_script_create                    = cc.DlFunc[func(filename cc.String) *Script, *Script]{Name: "cairo_script_create"}
	cairo_script_create_for_stream         = cc.DlFunc[func(write_func uptr, closure uptr) *Script, *Script]{Name: "cairo_script_create_for_stream"}
	cairo_script_write_comment             = cc.DlFunc[func(script *Script, comment cc.String, length int32), cc.Void]{Name: "cairo_script_write_comment"}
	cairo_script_set_mode                  = cc.DlFunc[func(script *Script, mode ScriptMode), cc.Void]{Name: "cairo_script_set_mode"}
	cairo_script_get_mode                  = cc.DlFunc[func(script *Script) ScriptMode, ScriptMode]{Name: "cairo_script_get_mode"}
	cairo_script_surface_create            = cc.DlFunc[func(script *Script, content Content, width, height float64) *Surface, *Surface]{Name: "cairo_script_surface_create"}
	cairo_script_surface_create_for_target = cc.DlFunc[func(script *Script, target *Surface) *Surface, *Surface]{Name: "cairo_script_surface_create_for_target"}
	cairo_script_from_recording_surface    = cc.DlFunc[func(script *Script, recording_surface *Surface) Status, Status]{Name: "cairo_script_from_recording_surface"}
	// #endregion

	// #region cairo-svg.h
	cairo_svg_surface_create              = cc.DlFunc[func(filename cc.String, width_in_points, height_in_points float64) *SVGSurface, *SVGSurface]{Name: "cairo_svg_surface_create"}
	cairo_svg_surface_create_for_stream   = cc.DlFunc[func(write_func uptr, closure uptr, width_in_points, height_in_points float64) *SVGSurface, *SVGSurface]{Name: "cairo_svg_surface_create_for_stream"}
	cairo_svg_surface_restrict_to_version = cc.DlFunc[func(surface *SVGSurface, version SVGVersion), cc.Void]{Name: "cairo_svg_surface_restrict_to_version"}
	cairo_svg_get_versions                = cc.DlFunc[func(versions **SVGVersion, num_versions *int32), cc.Void]{Name: "cairo_svg_get_versions"}
	cairo_svg_version_to_string           = cc.DlFunc[func(version SVGVersion) cc.String, cc.String]{Name: "cairo_svg_version_to_string"}
	cairo_svg_surface_set_document_unit   = cc.DlFunc[func(surface *SVGSurface, unit SVGUnit), cc.Void]{Name: "cairo_svg_surface_set_document_unit"}
	cairo_svg_surface_get_document_unit   = cc.DlFunc[func(surface *SVGSurface) SVGUnit, SVGUnit]{Name: "cairo_svg_surface_get_document_unit"}
	// #endregion

)
