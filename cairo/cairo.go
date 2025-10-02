package cairo

import (
	"runtime"

	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/glib/gobject"
)

// #region cairo.h

func Version() int32        { return cairo_version.Fn()() }
func VersionString() string { return cairo_version_string.Fn()() }

type Bool int32
type Context struct{}

type SurfaceIface interface {
	GetSurfaceIface() *Surface
}

func GetSurfaceIface(iface SurfaceIface) *Surface {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetSurfaceIface()
}

func (s *Surface) GetSurfaceIface() *Surface { return s }

type Surface struct{}

type DeviceIface interface {
	GetDeviceIface() *Device
}

func GetDeviceIface(iface DeviceIface) *Device {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetDeviceIface()
}
func (d *Device) GetDeviceIface() *Device { return d }

type Device struct{}
type Matrix struct{ XX, YX, XY, YY, X0, Y0 float64 }

type PatternIface interface {
	GetPatternIface() *Pattern
}

func GetPatternIface(iface PatternIface) *Pattern {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetPatternIface()
}
func (p *Pattern) GetPatternIface() *Pattern { return p }

type Pattern struct{}
type UserDataKey struct{ Unused int32 }

func (pattern *Pattern) SetDither(dither Dither) { cairo_pattern_set_dither.Fn()(pattern, dither) }
func (pattern *Pattern) GetDither() Dither       { return cairo_pattern_get_dither.Fn()(pattern) }

type RectangleInt struct {
	X, Y          int32
	Width, Height int32
}

func CreateContext(target SurfaceIface) *Context      { return cairo_create.Fn()(GetSurfaceIface(target)) }
func (cr *Context) Reference() *Context               { return cairo_reference.Fn()(cr) }
func (cr *Context) Destroy()                          { cairo_destroy.Fn()(cr) }
func (cr *Context) GetReferenceCount() uint32         { return cairo_get_reference_count.Fn()(cr) }
func (cr *Context) GetUserData(key *UserDataKey) uptr { return cairo_get_user_data.Fn()(cr, key) }
func (cr *Context) SetUserData(key *UserDataKey, user_data interface{}, destroy func(uptr)) Status {
	return cairo_set_user_data.Fn()(cr, key, anyptr(user_data), vcbu(destroy))
}
func (cr *Context) Save()      { cairo_save.Fn()(cr) }
func (cr *Context) Restore()   { cairo_restore.Fn()(cr) }
func (cr *Context) PushGroup() { cairo_push_group.Fn()(cr) }
func (cr *Context) PushGroupWithContent(content int32) {
	cairo_push_group_with_content.Fn()(cr, content)
}
func (cr *Context) PopGroup() *Pattern { return cairo_pop_group.Fn()(cr) }
func (cr *Context) PopGroupToSource()  { cairo_pop_group_to_source.Fn()(cr) }

func (cr *Context) SetOperator(op Operator)   { cairo_set_operator.Fn()(cr, op) }
func (cr *Context) SetSource(source *Pattern) { cairo_set_source.Fn()(cr, source) }
func (cr *Context) SetSourceRGB(red, green, blue float64) {
	cairo_set_source_rgb.Fn()(cr, red, green, blue)
}
func (cr *Context) SetSourceRGBA(red, green, blue, alpha float64) {
	cairo_set_source_rgba.Fn()(cr, red, green, blue, alpha)
}
func (cr *Context) SetSourceSurface(surface SurfaceIface, x, y float64) {
	cairo_set_source_surface.Fn()(cr, GetSurfaceIface(surface), x, y)
}
func (cr *Context) SetTolerance(tolerance float64) { cairo_set_tolerance.Fn()(cr, tolerance) }

func (cr *Context) SetAntialias(antialias Antialias) { cairo_set_antialias.Fn()(cr, antialias) }
func (cr *Context) SetFillRule(fill_rule FillRule)   { cairo_set_fill_rule.Fn()(cr, fill_rule) }
func (cr *Context) SetLineWidth(width float64)       { cairo_set_line_width.Fn()(cr, width) }
func (cr *Context) SetHairline(set_hairline Bool)    { cairo_set_hairline.Fn()(cr, set_hairline) }

func (cr *Context) SetLineCap(line_cap LineCap)    { cairo_set_line_cap.Fn()(cr, line_cap) }
func (cr *Context) SetLineJoin(line_join LineJoin) { cairo_set_line_join.Fn()(cr, line_join) }
func (cr *Context) SetDash(dashes []float64, num_dashes int32, offset float64) {
	cairo_set_dash.Fn()(cr, carry(dashes), num_dashes, offset)
}
func (cr *Context) SetMiterLimit(limit float64) { cairo_set_miter_limit.Fn()(cr, limit) }
func (cr *Context) Translate(tx, ty float64)    { cairo_translate.Fn()(cr, tx, ty) }
func (cr *Context) Scale(sx, sy float64)        { cairo_scale.Fn()(cr, sx, sy) }
func (cr *Context) Rotate(angle float64)        { cairo_rotate.Fn()(cr, angle) }
func (cr *Context) Transform(matrix *Matrix)    { cairo_transform.Fn()(cr, matrix) }
func (cr *Context) SetMatrix(matrix *Matrix)    { cairo_set_matrix.Fn()(cr, matrix) }
func (cr *Context) IdentityMatrix()             { cairo_identity_matrix.Fn()(cr) }
func (cr *Context) UserToDevice(x, y *float64)  { cairo_user_to_device.Fn()(cr, x, y) }
func (cr *Context) UserToDeviceDistance(dx, dy *float64) {
	cairo_user_to_device_distance.Fn()(cr, dx, dy)
}
func (cr *Context) DeviceToUser(x, y *float64) { cairo_device_to_user.Fn()(cr, x, y) }
func (cr *Context) DeviceToUserDistance(dx, dy *float64) {
	cairo_device_to_user_distance.Fn()(cr, dx, dy)
}
func (cr *Context) NewPath()            { cairo_new_path.Fn()(cr) }
func (cr *Context) MoveTo(x, y float64) { cairo_move_to.Fn()(cr, x, y) }
func (cr *Context) NewSubPath()         { cairo_new_sub_path.Fn()(cr) }
func (cr *Context) LineTo(x, y float64) { cairo_line_to.Fn()(cr, x, y) }
func (cr *Context) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	cairo_curve_to.Fn()(cr, x1, y1, x2, y2, x3, y3)
}
func (cr *Context) Arc(xc, yc, radius, angle1, angle2 float64) {
	cairo_arc.Fn()(cr, xc, yc, radius, angle1, angle2)
}
func (cr *Context) ArcNegative(xc, yc, radius, angle1, angle2 float64) {
	cairo_arc_negative.Fn()(cr, xc, yc, radius, angle1, angle2)
}
func (cr *Context) RelMoveTo(dx, dy float64) { cairo_rel_move_to.Fn()(cr, dx, dy) }
func (cr *Context) RelLineTo(dx, dy float64) { cairo_rel_line_to.Fn()(cr, dx, dy) }
func (cr *Context) RelCurveTo(dx1, dy1, dx2, dy2, dx3, dy3 float64) {
	cairo_rel_curve_to.Fn()(cr, dx1, dy1, dx2, dy2, dx3, dy3)
}
func (cr *Context) Rectangle(x, y, width, height float64) {
	cairo_rectangle.Fn()(cr, x, y, width, height)
}
func (cr *Context) ClosePath() { cairo_close_path.Fn()(cr) }
func (cr *Context) PathExtents() (x1, y1, x2, y2 float64) {
	cairo_path_extents.Fn()(cr, &x1, &y1, &x2, &y2)
	return
}
func (cr *Context) Paint()                       { cairo_paint.Fn()(cr) }
func (cr *Context) PaintWithAlpha(alpha float64) { cairo_paint_with_alpha.Fn()(cr, alpha) }
func (cr *Context) Mask(pattern PatternIface)    { cairo_mask.Fn()(cr, GetPatternIface(pattern)) }
func (cr *Context) MaskSurface(surface SurfaceIface, surface_x, surface_y float64) {
	cairo_mask_surface.Fn()(cr, GetSurfaceIface(surface), surface_x, surface_y)
}
func (cr *Context) Stroke()                    { cairo_stroke.Fn()(cr) }
func (cr *Context) StrokePreserve()            { cairo_stroke_preserve.Fn()(cr) }
func (cr *Context) Fill()                      { cairo_fill.Fn()(cr) }
func (cr *Context) FillPreserve()              { cairo_fill_preserve.Fn()(cr) }
func (cr *Context) CopyPage()                  { cairo_copy_page.Fn()(cr) }
func (cr *Context) ShowPage()                  { cairo_show_page.Fn()(cr) }
func (cr *Context) InStroke(x, y float64) Bool { return cairo_in_stroke.Fn()(cr, x, y) }
func (cr *Context) InFill(x, y float64) Bool   { return cairo_in_fill.Fn()(cr, x, y) }
func (cr *Context) InClip(x, y float64) Bool   { return cairo_in_clip.Fn()(cr, x, y) }
func (cr *Context) StrokeExtents() (x1, y1, x2, y2 float64) {
	cairo_stroke_extents.Fn()(cr, &x1, &y1, &x2, &y2)
	return
}
func (cr *Context) FillExtents() (x1, y1, x2, y2 float64) {
	cairo_fill_extents.Fn()(cr, &x1, &y1, &x2, &y2)
	return
}
func (cr *Context) ResetClip()    { cairo_reset_clip.Fn()(cr) }
func (cr *Context) Clip()         { cairo_clip.Fn()(cr) }
func (cr *Context) ClipPreserve() { cairo_clip_preserve.Fn()(cr) }
func (cr *Context) ClipExtents() (x1, y1, x2, y2 float64) {
	cairo_clip_extents.Fn()(cr, &x1, &y1, &x2, &y2)
	return
}

type Rectangle struct {
	X, Y          float64
	Width, Height float64
}

type RectangleList struct {
	Status        Status
	Rectangles    *Rectangle
	NumRectangles int32
}

func (rl *RectangleList) GetList() []Rectangle {
	return *(*[]Rectangle)(slice(uptr(rl.Rectangles), int(rl.NumRectangles)))
}

func (cr *Context) CopyClipRectangleList() *RectangleList {
	return cairo_copy_clip_rectangle_list.Fn()(cr)
}
func RectangleListDestroy(rectangle_list *RectangleList) {
	cairo_rectangle_list_destroy.Fn()(rectangle_list)
}
func (cr *Context) TagBegin(tag_name, attributes string) {
	cTag, cAttr := cc.NewString(tag_name), cc.NewString(attributes)
	defer cTag.Free()
	defer cAttr.Free()
	cairo_tag_begin.Fn()(cr, cTag, cAttr)
}
func (cr *Context) TagEnd(tag_name string) {
	cTag := cc.NewString(tag_name)
	defer cTag.Free()
	cairo_tag_end.Fn()(cr, cTag)
}

type ScaledFontIface interface {
	GetScaledFontIface() *ScaledFont
}

func GetScaledFontIface(iface ScaledFontIface) *ScaledFont {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetScaledFontIface()
}
func (sf *ScaledFont) GetScaledFontIface() *ScaledFont { return sf }

type ScaledFont struct{}

type FontFaceIface interface {
	GetFontFaceIface() *FontFace
}

func GetFontFaceIface(iface FontFaceIface) *FontFace {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetFontFaceIface()
}
func (ff *FontFace) GetFontFaceIface() *FontFace { return ff }

type FontFace struct{}

type Glyph struct {
	Index uint64
	X, Y  float64
}

func GlyphAllocate(num_glyphs int32) *Glyph { return cairo_glyph_allocate.Fn()(num_glyphs) }
func GlyphFree(glyphs *Glyph)               { cairo_glyph_free.Fn()(glyphs) }

type TextCluster struct {
	NumBytes  int32
	NumGlyphs int32
}

func TextClusterAllocate(num_clusters int32) *TextCluster {
	return cairo_text_cluster_allocate.Fn()(num_clusters)
}
func TextClusterFree(clusters *TextCluster) { cairo_text_cluster_free.Fn()(clusters) }

type TextExtents struct {
	Xbearing      float64
	Ybearing      float64
	Width, Height float64
	Xadvance      float64
	Yadvance      float64
}

type FontExtents struct {
	Ascent      float64
	Descent     float64
	Height      float64
	MaxXadvance float64
	MaxYadvance float64
}

type FontOptions struct{}

func CreateFontOptions() *FontOptions            { return cairo_font_options_create.Fn()() }
func (fo *FontOptions) Copy() *FontOptions       { return cairo_font_options_copy.Fn()(fo) }
func (fo *FontOptions) Destroy()                 { cairo_font_options_destroy.Fn()(fo) }
func (fo *FontOptions) Status() Status           { return cairo_font_options_status.Fn()(fo) }
func (fo *FontOptions) Merge(other *FontOptions) { cairo_font_options_merge.Fn()(fo, other) }
func (fo *FontOptions) Equal(other *FontOptions) Bool {
	return cairo_font_options_equal.Fn()(fo, other)
}
func (fo *FontOptions) Hash() uint64 { return cairo_font_options_hash.Fn()(fo) }
func (fo *FontOptions) SetAntialias(antialias Antialias) {
	cairo_font_options_set_antialias.Fn()(fo, antialias)
}
func (fo *FontOptions) GetAntialias() Antialias {
	return Antialias(cairo_font_options_get_antialias.Fn()(fo))
}
func (fo *FontOptions) SetSubpixelOrder(subpixel_order SubpixelOrder) {
	cairo_font_options_set_subpixel_order.Fn()(fo, subpixel_order)
}
func (fo *FontOptions) GetSubpixelOrder() SubpixelOrder {
	return SubpixelOrder(cairo_font_options_get_subpixel_order.Fn()(fo))
}
func (fo *FontOptions) SetHintStyle(hint_style HintStyle) {
	cairo_font_options_set_hint_style.Fn()(fo, hint_style)
}
func (fo *FontOptions) GetHintStyle() HintStyle {
	return HintStyle(cairo_font_options_get_hint_style.Fn()(fo))
}
func (fo *FontOptions) SetHintMetrics(hint_metrics HintMetrics) {
	cairo_font_options_set_hint_metrics.Fn()(fo, hint_metrics)
}
func (fo *FontOptions) GetHintMetrics() HintMetrics {
	return HintMetrics(cairo_font_options_get_hint_metrics.Fn()(fo))
}
func (fo *FontOptions) GetVariations() string {
	return cairo_font_options_get_variations.Fn()(fo).String()
}
func (fo *FontOptions) SetVariations(variations string) {
	cVar := cc.NewString(variations)
	defer cVar.Free()
	cairo_font_options_set_variations.Fn()(fo, cVar)
}
func (fo *FontOptions) SetColorMode(color_mode ColorMode) {
	cairo_font_options_set_color_mode.Fn()(fo, color_mode)
}
func (fo *FontOptions) GetColorMode() ColorMode { return cairo_font_options_get_color_mode.Fn()(fo) }
func (fo *FontOptions) GetColorPalette() uint32 { return cairo_font_options_get_color_palette.Fn()(fo) }
func (fo *FontOptions) SetColorPalette(palette_index uint32) {
	cairo_font_options_set_color_palette.Fn()(fo, palette_index)
}
func (fo *FontOptions) SetCustomPaletteColor(index uint32, red, green, blue, alpha float64) {
	cairo_font_options_set_custom_palette_color.Fn()(fo, index, red, green, blue, alpha)
}
func (fo *FontOptions) GetCustomPaletteColor(index uint32) (red, green, blue, alpha float64, status Status) {
	status = cairo_font_options_get_custom_palette_color.Fn()(fo, index, &red, &green, &blue, &alpha)
	return
}

func (cr *Context) SelectFontFace(family string, slant FontSlant, weight FontWeight) {
	cFamily := cc.NewString(family)
	defer cFamily.Free()
	cairo_select_font_face.Fn()(cr, cFamily, slant, weight)
}
func (cr *Context) SetFontSize(size float64)            { cairo_set_font_size.Fn()(cr, size) }
func (cr *Context) SetFontMatrix(matrix *Matrix)        { cairo_set_font_matrix.Fn()(cr, matrix) }
func (cr *Context) GetFontMatrix() (matrix Matrix)      { cairo_get_font_matrix.Fn()(cr, &matrix); return }
func (cr *Context) SetFontOptions(options *FontOptions) { cairo_set_font_options.Fn()(cr, options) }
func (cr *Context) GetFontOptions() (options FontOptions) {
	cairo_get_font_options.Fn()(cr, &options)
	return
}
func (cr *Context) SetFontFace(font_face FontFaceIface) {
	cairo_set_font_face.Fn()(cr, GetFontFaceIface(font_face))
}
func (cr *Context) GetFontFace() *FontFace { return cairo_get_font_face.Fn()(cr) }
func (cr *Context) SetScaledFont(scaled_font *ScaledFont) {
	cairo_set_scaled_font.Fn()(cr, scaled_font)
}
func (cr *Context) GetScaledFont() *ScaledFont { return cairo_get_scaled_font.Fn()(cr) }
func (cr *Context) ShowText(utf8 string) {
	cUtf8 := cc.NewString(utf8)
	defer cUtf8.Free()
	cairo_show_text.Fn()(cr, cUtf8)
}
func (cr *Context) ShowGlyphs(glyphs []Glyph) {
	cairo_show_glyphs.Fn()(cr, carry(glyphs), int32(len(glyphs)))
}
func (cr *Context) ShowTextGlyphs(utf8 string, glyphs []Glyph, clusters []TextCluster, cluster_flags int32) {
	cUtf8 := cc.NewString(utf8)
	defer cUtf8.Free()
	cairo_show_text_glyphs.Fn()(cr,
		cUtf8, int32(len(utf8)),
		carry(glyphs), int32(len(glyphs)),
		carry(clusters), int32(len(clusters)),
		cluster_flags)
}
func (cr *Context) TextPath(utf8 string) {
	cUtf8 := cc.NewString(utf8)
	defer cUtf8.Free()
	cairo_text_path.Fn()(cr, cUtf8)
}
func (cr *Context) GlyphPath(glyphs []Glyph) {
	cairo_glyph_path.Fn()(cr, carry(glyphs), int32(len(glyphs)))
}
func (cr *Context) TextExtents(utf8 string, extents *TextExtents) {
	cUtf8 := cc.NewString(utf8)
	defer cUtf8.Free()
	cairo_text_extents.Fn()(cr, cUtf8, extents)
}
func (cr *Context) GlyphExtents(glyphs []Glyph, extents *TextExtents) {
	cairo_glyph_extents.Fn()(cr, carry(glyphs), int32(len(glyphs)), extents)
}
func (cr *Context) FontExtents(extents *FontExtents) { cairo_font_extents.Fn()(cr, extents) }

func (ff *FontFace) Reference() *FontFace      { return cairo_font_face_reference.Fn()(ff) }
func (ff *FontFace) Destroy()                  { cairo_font_face_destroy.Fn()(ff) }
func (ff *FontFace) GetReferenceCount() uint32 { return cairo_font_face_get_reference_count.Fn()(ff) }
func (ff *FontFace) Status() Status            { return cairo_font_face_status.Fn()(ff) }
func (ff *FontFace) GetType() FontType         { return cairo_font_face_get_type.Fn()(ff) }
func (ff *FontFace) GetUserData(key *UserDataKey) uptr {
	return cairo_font_face_get_user_data.Fn()(ff, key)
}
func (ff *FontFace) SetUserData(key *UserDataKey, user_data interface{}, destroy func(uptr)) Status {
	return cairo_font_face_set_user_data.Fn()(ff, key, anyptr(user_data), vcbu(destroy))
}

func CreateScaledFont(font_face FontFaceIface, font_matrix, ctm *Matrix, options *FontOptions) *ScaledFont {
	return cairo_scaled_font_create.Fn()(GetFontFaceIface(font_face), font_matrix, ctm, options)
}
func (sf *ScaledFont) Reference() *ScaledFont { return cairo_scaled_font_reference.Fn()(sf) }
func (sf *ScaledFont) Destroy()               { cairo_scaled_font_destroy.Fn()(sf) }
func (sf *ScaledFont) GetReferenceCount() uint32 {
	return cairo_scaled_font_get_reference_count.Fn()(sf)
}
func (sf *ScaledFont) Status() Status    { return cairo_scaled_font_status.Fn()(sf) }
func (sf *ScaledFont) GetType() FontType { return cairo_scaled_font_get_type.Fn()(sf) }
func (sf *ScaledFont) GetUserData(key *UserDataKey) uptr {
	return cairo_scaled_font_get_user_data.Fn()(sf, key)
}
func (sf *ScaledFont) SetUserData(key *UserDataKey, user_data interface{}, destroy func(uptr)) Status {
	return cairo_scaled_font_set_user_data.Fn()(sf, key, anyptr(user_data), vcbu(destroy))
}
func (sf *ScaledFont) Extents(extents *FontExtents) { cairo_scaled_font_extents.Fn()(sf, extents) }
func (sf *ScaledFont) TextExtents(utf8 string, extents *TextExtents) {
	cUtf8 := cc.NewString(utf8)
	defer cUtf8.Free()
	cairo_scaled_font_text_extents.Fn()(sf, cUtf8, extents)
}
func (sf *ScaledFont) GlyphExtents(glyphs []Glyph, extents *TextExtents) {
	cairo_scaled_font_glyph_extents.Fn()(sf, carry(glyphs), int32(len(glyphs)), extents)
}
func (sf *ScaledFont) TextToGlyphs(x, y float64, utf8 string) (glyphs []Glyph, clusters []TextCluster, cluster_flags int32, status Status) {
	cUtf8 := cc.NewString(utf8)
	defer cUtf8.Free()
	var glyphsPtr *Glyph
	var numGlyphs int32
	var clustersPtr *TextCluster
	var numClusters int32
	var flags int32
	status = cairo_scaled_font_text_to_glyphs.Fn()(sf, x, y, cUtf8, int32(len(utf8)),
		&glyphsPtr, &numGlyphs, &clustersPtr, &numClusters, &flags)
	if glyphsPtr != nil && numGlyphs > 0 {
		glyphs = *(*[]Glyph)(slice(uptr(glyphsPtr), int(numGlyphs)))
	}
	if clustersPtr != nil && numClusters > 0 {
		clusters = *(*[]TextCluster)(slice(uptr(clustersPtr), int(numClusters)))
	}
	cluster_flags = flags
	return
}
func (sf *ScaledFont) GetFontFace() *FontFace { return cairo_scaled_font_get_font_face.Fn()(sf) }
func (sf *ScaledFont) GetFontMatrix() (font_matrix Matrix) {
	cairo_scaled_font_get_font_matrix.Fn()(sf, &font_matrix)
	return
}
func (sf *ScaledFont) GetCTM() (ctm Matrix) { cairo_scaled_font_get_ctm.Fn()(sf, &ctm); return }
func (sf *ScaledFont) GetScaleMatrix() (scale_matrix Matrix) {
	cairo_scaled_font_get_scale_matrix.Fn()(sf, &scale_matrix)
	return
}
func (sf *ScaledFont) GetFontOptions(options *FontOptions) {
	cairo_scaled_font_get_font_options.Fn()(sf, options)
}

type ToyFontFace struct{ FontFace }

func CreateToyFontFace(family string, slant FontSlant, weight FontWeight) *ToyFontFace {
	cFamily := cc.NewString(family)
	defer cFamily.Free()
	return cairo_toy_font_face_create.Fn()(cFamily, slant, weight)
}
func (ff *ToyFontFace) GetFamily() string     { return cairo_toy_font_face_get_family.Fn()(ff).String() }
func (ff *ToyFontFace) GetSlant() FontSlant   { return cairo_toy_font_face_get_slant.Fn()(ff) }
func (ff *ToyFontFace) GetWeight() FontWeight { return cairo_toy_font_face_get_weight.Fn()(ff) }

type UserFontFace struct{ FontFace }

func CreateUserFontFace() *UserFontFace { return cairo_user_font_face_create.Fn()() }

func (ff *UserFontFace) SetInitFunc(init_func func(scaledFont *ScaledFont, cr *Context, extents *FontExtents) Status) {
	cairo_user_font_face_set_init_func.Fn()(ff, vcbu(init_func))
}
func (ff *UserFontFace) SetRenderGlyphFunc(
	render_glyph_func func(scaledFont *ScaledFont, glyph uint64, cr *Context, extents *TextExtents) Status) {
	cairo_user_font_face_set_render_glyph_func.Fn()(ff, vcbu(render_glyph_func))
}
func (ff *UserFontFace) SetRenderColorGlyphFunc(
	render_glyph_func func(scaledFont *ScaledFont, glyph uint64, cr *Context, extents *TextExtents) Status) {
	cairo_user_font_face_set_render_color_glyph_func.Fn()(ff, vcbu(render_glyph_func))
}
func (ff *UserFontFace) SetTextToGlyphsFunc(
	text_to_glyphs_func func(scaledFont *ScaledFont,
		glyphs **Glyph, numGlyphs *int32,
		clusters **TextCluster, numClusters *int32, clusterFlags *TextClusterFlags) Status) {
	cairo_user_font_face_set_text_to_glyphs_func.Fn()(ff, vcbu(text_to_glyphs_func))
}
func (ff *UserFontFace) SetUnicodeToGlyphFunc(
	unicode_to_glyph_func func(scaledFont *ScaledFont, unicode uint64, glyphIndex *uint64) Status) {
	cairo_user_font_face_set_unicode_to_glyph_func.Fn()(ff, vcbu(unicode_to_glyph_func))
}

func (ff *UserFontFace) GetInitFunc() cc.Func { return cairo_user_font_face_get_init_func.Fn()(ff) }
func (ff *UserFontFace) GetRenderGlyphFunc() cc.Func {
	return cairo_user_font_face_get_render_glyph_func.Fn()(ff)
}
func (ff *UserFontFace) GetRenderColorGlyphFunc() cc.Func {
	return cairo_user_font_face_get_render_color_glyph_func.Fn()(ff)
}
func (ff *UserFontFace) GetTextToGlyphsFunc() cc.Func {
	return cairo_user_font_face_get_text_to_glyphs_func.Fn()(ff)
}
func (ff *UserFontFace) GetUnicodeToGlyphFunc() cc.Func {
	return cairo_user_font_face_get_unicode_to_glyph_func.Fn()(ff)
}

type UserScaledFont struct{ ScaledFont }

func (sf *UserScaledFont) GetForegroundMarker() *Pattern {
	return cairo_user_scaled_font_get_foreground_marker.Fn()(sf)
}
func (sf *UserScaledFont) GetForegroundSource() *Pattern {
	return cairo_user_scaled_font_get_foreground_source.Fn()(sf)
}

func (cr *Context) GetOperator() Operator           { return cairo_get_operator.Fn()(cr) }
func (cr *Context) GetSource() *Pattern             { return cairo_get_source.Fn()(cr) }
func (cr *Context) GetTolerance() float64           { return cairo_get_tolerance.Fn()(cr) }
func (cr *Context) GetAntialias() Antialias         { return cairo_get_antialias.Fn()(cr) }
func (cr *Context) HasCurrentPoint() Bool           { return cairo_has_current_point.Fn()(cr) }
func (cr *Context) GetCurrentPoint() (x, y float64) { cairo_get_current_point.Fn()(cr, &x, &y); return }
func (cr *Context) GetFillRule() FillRule           { return cairo_get_fill_rule.Fn()(cr) }
func (cr *Context) GetLineWidth() float64           { return cairo_get_line_width.Fn()(cr) }
func (cr *Context) GetHairline() Bool               { return cairo_get_hairline.Fn()(cr) }
func (cr *Context) GetLineCap() LineCap             { return cairo_get_line_cap.Fn()(cr) }
func (cr *Context) GetLineJoin() LineJoin           { return cairo_get_line_join.Fn()(cr) }
func (cr *Context) GetMiterLimit() float64          { return cairo_get_miter_limit.Fn()(cr) }
func (cr *Context) GetDashCount() int32             { return cairo_get_dash_count.Fn()(cr) }
func (cr *Context) GetDash() (dashes []float64, offset float64) {
	n := cairo_get_dash_count.Fn()(cr)
	dashes = make([]float64, n)
	cairo_get_dash.Fn()(cr, carry(dashes), &offset)
	return
}
func (cr *Context) GetMatrix() (matrix Matrix) { cairo_get_matrix.Fn()(cr, &matrix); return }
func (cr *Context) GetTarget() *Surface        { return cairo_get_target.Fn()(cr) }
func (cr *Context) GetGroupTarget() *Surface   { return cairo_get_group_target.Fn()(cr) }

type PathData struct {
	_ [2]float64
}
type PathDataHeader struct {
	DataType PathDataType
	Lenght   int32
}
type PathDataPoint struct{ X, Y float64 }

func (data *PathData) ToPoint() *PathDataPoint   { return (*PathDataPoint)(uptr(data)) }
func (data *PathData) ToHeader() *PathDataHeader { return (*PathDataHeader)(uptr(data)) }

type Path struct {
	Status  Status
	Data    *PathData
	NumData int32
}

func (p *Path) PathData() []PathData { return *(*[]PathData)(slice(uptr(p.Data), int(p.NumData))) }

func (cr *Context) CopyPath() *Path       { return cairo_copy_path.Fn()(cr) }
func (cr *Context) CopyPathFlat() *Path   { return cairo_copy_path_flat.Fn()(cr) }
func (cr *Context) AppendPath(path *Path) { cairo_append_path.Fn()(cr, path) }
func PathDestroy(path *Path)              { cairo_path_destroy.Fn()(path) }
func (cr *Context) Status() Status        { return cairo_status.Fn()(cr) }
func StatusToString(status Status) string { return cairo_status_to_string.Fn()(status).String() }
func (st Status) String() string          { return cairo_status_to_string.Fn()(st).String() }
func (d *Device) Reference() *Device      { return cairo_device_reference.Fn()(d) }

func (d *Device) GetType() DeviceType               { return cairo_device_get_type.Fn()(d) }
func (d *Device) Status() Status                    { return cairo_device_status.Fn()(d) }
func (d *Device) Acquire() Status                   { return cairo_device_acquire.Fn()(d) }
func (d *Device) Release()                          { cairo_device_release.Fn()(d) }
func (d *Device) Flush()                            { cairo_device_flush.Fn()(d) }
func (d *Device) Finish()                           { cairo_device_finish.Fn()(d) }
func (d *Device) Destroy()                          { cairo_device_destroy.Fn()(d) }
func (d *Device) GetReferenceCount() uint32         { return cairo_device_get_reference_count.Fn()(d) }
func (d *Device) GetUserData(key *UserDataKey) uptr { return cairo_device_get_user_data.Fn()(d, key) }
func (d *Device) SetUserData(key *UserDataKey, user_data interface{}, destroy func(uptr)) Status {
	return cairo_device_set_user_data.Fn()(d, key, anyptr(user_data), vcbu(destroy))
}

func (s *Surface) CreateSimilar(content int32, width, height int32) *Surface {
	return cairo_surface_create_similar.Fn()(s, content, width, height)
}
func (s *Surface) CreateSimilarImage(format int32, width, height int32) *Surface {
	return cairo_surface_create_similar_image.Fn()(s, format, width, height)
}
func (s *Surface) MapToImage(extents *RectangleInt) *Surface {
	return cairo_surface_map_to_image.Fn()(s, extents)
}
func (s *Surface) UnmapImage(image SurfaceIface) {
	cairo_surface_unmap_image.Fn()(s, GetSurfaceIface(image))
}
func (s *Surface) CreateForRectangle(x, y, width, height float64) *Surface {
	return cairo_surface_create_for_rectangle.Fn()(s, x, y, width, height)
}

type SurfaceObserver struct{ Surface }

func (s *Surface) CreateObserver(mode SurfaceObserverMode) *SurfaceObserver {
	return cairo_surface_create_observer.Fn()(s, mode)
}
func (s *SurfaceObserver) AddPaintCallback(
	fn func(observer, target *Surface, data uptr), data interface{}) Status {
	return cairo_surface_observer_add_paint_callback.Fn()(s, vcbu(fn), anyptr(data))
}
func (s *SurfaceObserver) AddMaskCallback(
	fn func(observer, target *Surface, data uptr), data interface{}) Status {
	return cairo_surface_observer_add_mask_callback.Fn()(s, vcbu(fn), anyptr(data))
}
func (s *SurfaceObserver) AddFillCallback(
	fn func(observer, target *Surface, data uptr), data interface{}) Status {
	return cairo_surface_observer_add_fill_callback.Fn()(s, vcbu(fn), anyptr(data))
}
func (s *SurfaceObserver) AddStrokeCallback(
	fn func(observer, target *Surface, data uptr), data interface{}) Status {
	return cairo_surface_observer_add_stroke_callback.Fn()(s, vcbu(fn), anyptr(data))
}
func (s *SurfaceObserver) AddGlyphsCallback(
	fn func(observer, target *Surface, data uptr), data interface{}) Status {
	return cairo_surface_observer_add_glyphs_callback.Fn()(s, vcbu(fn), anyptr(data))
}
func (s *SurfaceObserver) AddFlushCallback(
	fn func(observer, target *Surface, data uptr), data interface{}) Status {
	return cairo_surface_observer_add_flush_callback.Fn()(s, vcbu(fn), anyptr(data))
}
func (s *SurfaceObserver) AddFinishCallback(
	fn func(observer, target *Surface, data uptr), data interface{}) Status {
	return cairo_surface_observer_add_finish_callback.Fn()(s, vcbu(fn), anyptr(data))
}
func (s *SurfaceObserver) Print(
	write_func func(closure, dataCharPtr uptr, lenght uint32), closure uptr) Status {
	return cairo_surface_observer_print.Fn()(s, vcbu(write_func), closure)
}
func (s *SurfaceObserver) Elapsed() float64 { return cairo_surface_observer_elapsed.Fn()(s) }

func (d *Device) ObserverPrint(
	write_func func(closure, dataCharPtr uptr, lenght uint32), closure uptr) Status {
	return cairo_device_observer_print.Fn()(d, vcbu(write_func), closure)
}
func (d *Device) ObserverElapsed() float64       { return cairo_device_observer_elapsed.Fn()(d) }
func (d *Device) ObserverPaintElapsed() float64  { return cairo_device_observer_paint_elapsed.Fn()(d) }
func (d *Device) ObserverMaskElapsed() float64   { return cairo_device_observer_mask_elapsed.Fn()(d) }
func (d *Device) ObserverFillElapsed() float64   { return cairo_device_observer_fill_elapsed.Fn()(d) }
func (d *Device) ObserverStrokeElapsed() float64 { return cairo_device_observer_stroke_elapsed.Fn()(d) }
func (d *Device) ObserverGlyphsElapsed() float64 { return cairo_device_observer_glyphs_elapsed.Fn()(d) }

func (s *Surface) Reference() *Surface       { return cairo_surface_reference.Fn()(s) }
func (s *Surface) Finish()                   { cairo_surface_finish.Fn()(s) }
func (s *Surface) Destroy()                  { cairo_surface_destroy.Fn()(s) }
func (s *Surface) GetDevice() *Device        { return cairo_surface_get_device.Fn()(s) }
func (s *Surface) GetReferenceCount() uint32 { return cairo_surface_get_reference_count.Fn()(s) }
func (s *Surface) Status() Status            { return cairo_surface_status.Fn()(s) }

func (s *Surface) GetType() SurfaceType { return SurfaceType(cairo_surface_get_type.Fn()(s)) }
func (s *Surface) GetContent() Content  { return Content(cairo_surface_get_content.Fn()(s)) }
func (s *Surface) WriteToPNG(filename string) Status {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	return cairo_surface_write_to_png.Fn()(s, cFilename)
}
func (s *Surface) WriteToPNGStream(write_func func(closure, dataCharPtr uptr, lenght uint32), closure uptr) Status {
	return cairo_surface_write_to_png_stream.Fn()(s, vcbu(write_func), anyptr(closure))
}
func (s *Surface) GetUserData(key *UserDataKey) uptr { return cairo_surface_get_user_data.Fn()(s, key) }
func (s *Surface) SetUserData(key *UserDataKey, user_data interface{}, destroy func(uptr)) Status {
	return cairo_surface_set_user_data.Fn()(s, key, anyptr(user_data), vcbu(destroy))
}
func (s *Surface) GetMimeData(mime_type MimeType) (data []byte) {
	cMime := cc.NewString(string(mime_type))
	defer cMime.Free()
	var ptr *byte
	var length uint64
	cairo_surface_get_mime_data.Fn()(s, cMime, &ptr, &length)
	if ptr != nil && length > 0 {
		data = *(*[]byte)(slice(uptr(ptr), int(length)))
	}
	return
}
func (s *Surface) SetMimeData(mime_type MimeType, data []byte, destroy func(uptr), closure uptr) Status {
	cMime := cc.NewString(string(mime_type))
	defer cMime.Free()
	pin := runtime.Pinner{}
	pin.Pin(data)
	d := func(p uptr) {
		defer pin.Unpin()
		if destroy != nil {
			destroy(p)
		}
	}
	return cairo_surface_set_mime_data.Fn()(s, cMime, carry(data), uint64(len(data)), vcbu(d), closure)
}
func (s *Surface) SupportsMimeType(mime_type MimeType) Bool {
	cMime := cc.NewString(string(mime_type))
	defer cMime.Free()
	return cairo_surface_supports_mime_type.Fn()(s, cMime)
}
func (s *Surface) GetFontOptions(options *FontOptions) {
	cairo_surface_get_font_options.Fn()(s, options)
}
func (s *Surface) Flush()     { cairo_surface_flush.Fn()(s) }
func (s *Surface) MarkDirty() { cairo_surface_mark_dirty.Fn()(s) }
func (s *Surface) MarkDirtyRectangle(x, y, width, height int32) {
	cairo_surface_mark_dirty_rectangle.Fn()(s, x, y, width, height)
}
func (s *Surface) SetDeviceScale(x_scale, y_scale float64) {
	cairo_surface_set_device_scale.Fn()(s, x_scale, y_scale)
}
func (s *Surface) GetDeviceScale() (x_scale, y_scale float64) {
	cairo_surface_get_device_scale.Fn()(s, &x_scale, &y_scale)
	return
}
func (s *Surface) SetDeviceOffset(x_offset, y_offset float64) {
	cairo_surface_set_device_offset.Fn()(s, x_offset, y_offset)
}
func (s *Surface) GetDeviceOffset() (x_offset, y_offset float64) {
	cairo_surface_get_device_offset.Fn()(s, &x_offset, &y_offset)
	return
}
func (s *Surface) SetFallbackResolution(x_ppi, y_ppi float64) {
	cairo_surface_set_fallback_resolution.Fn()(s, x_ppi, y_ppi)
}
func (s *Surface) GetFallbackResolution() (x_ppi, y_ppi float64) {
	cairo_surface_get_fallback_resolution.Fn()(s, &x_ppi, &y_ppi)
	return
}
func (s *Surface) CopyPage()               { cairo_surface_copy_page.Fn()(s) }
func (s *Surface) ShowPage()               { cairo_surface_show_page.Fn()(s) }
func (s *Surface) HasShowTextGlyphs() Bool { return cairo_surface_has_show_text_glyphs.Fn()(s) }

type ImageSurface struct{ Surface }

func CreateImageSurface(format Format, width, height int32) *ImageSurface {
	return cairo_image_surface_create.Fn()(format, width, height)
}
func FormatStrideForWidth(format Format, width int32) int32 {
	return cairo_format_stride_for_width.Fn()(format, width)
}
func CreateImageSurfaceForData(data []byte, format Format, width, height, stride int32) *ImageSurface {
	return cairo_image_surface_create_for_data.Fn()(carry(data), format, width, height, stride)
}
func (s *ImageSurface) GetData() *byte   { return cairo_image_surface_get_data.Fn()(s) }
func (s *ImageSurface) GetFormat() int32 { return cairo_image_surface_get_format.Fn()(s) }
func (s *ImageSurface) GetWidth() int32  { return cairo_image_surface_get_width.Fn()(s) }
func (s *ImageSurface) GetHeight() int32 { return cairo_image_surface_get_height.Fn()(s) }
func (s *ImageSurface) GetStride() int32 { return cairo_image_surface_get_stride.Fn()(s) }
func CreateImageSurfaceFromPNG(filename string) *ImageSurface {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	return cairo_image_surface_create_from_png.Fn()(cFilename)
}
func CreateImageSurfaceFromPNGStream(
	read_func func(closure uptr, data uptr, length uint32) Status, closure uptr) *ImageSurface {
	return cairo_image_surface_create_from_png_stream.Fn()(vcbu(read_func), anyptr(closure))
}

type RecordingSurface struct{ Surface }

func CreateRecordingSurface(content int32, extents *Rectangle) *RecordingSurface {
	return cairo_recording_surface_create.Fn()(content, extents)
}
func (s *RecordingSurface) InkExtents() (x0, y0, width, height float64) {
	cairo_recording_surface_ink_extents.Fn()(s, &x0, &y0, &width, &height)
	return
}
func (s *RecordingSurface) GetExtents(extents *Rectangle) Bool {
	return cairo_recording_surface_get_extents.Fn()(s, extents)
}

type RasterSourcePattern struct{ Pattern }

func CreatePatternRasterSource(user_data interface{}, content Content, width, height int32) *RasterSourcePattern {
	return cairo_pattern_create_raster_source.Fn()(anyptr(user_data), content, width, height)
}
func (p *RasterSourcePattern) SetCallbackData(data interface{}) {
	cairo_raster_source_pattern_set_callback_data.Fn()(p, anyptr(data))
}
func (p *RasterSourcePattern) GetCallbackData() uptr {
	return cairo_raster_source_pattern_get_callback_data.Fn()(p)
}
func (p *RasterSourcePattern) SetAcquire(
	acquire func(pattern *Pattern, callbackData uptr, target *Surface, extents *RectangleInt) *Surface,
	release func(pattern *Pattern, callbackData uptr, surface *Surface)) {
	cairo_raster_source_pattern_set_acquire.Fn()(p, vcbu(acquire), vcbu(release))
}
func (p *RasterSourcePattern) GetAcquire() (acquire, release cc.Func) {
	cairo_raster_source_pattern_get_acquire.Fn()(p, &acquire, &release)
	return
}
func (p *RasterSourcePattern) SetSnapshot(snapshot func(pattern *Pattern, callbackData uptr) Status) {
	cairo_raster_source_pattern_set_snapshot.Fn()(p, vcbu(snapshot))
}
func (p *RasterSourcePattern) GetSnapshot() cc.Func {
	return cairo_raster_source_pattern_get_snapshot.Fn()(p)
}
func (p *RasterSourcePattern) SetCopy(copy func(pattern *Pattern, callbackData uptr, srcOther *Pattern) Status) {
	cairo_raster_source_pattern_set_copy.Fn()(p, vcbu(copy))
}
func (p *RasterSourcePattern) GetCopy() cc.Func {
	return cairo_raster_source_pattern_get_copy.Fn()(p)
}
func (p *RasterSourcePattern) SetFinish(finish func(pattern *Pattern, callbackData uptr)) {
	cairo_raster_source_pattern_set_finish.Fn()(p, vcbu(finish))
}
func (p *RasterSourcePattern) GetFinish() cc.Func {
	return cairo_raster_source_pattern_get_finish.Fn()(p)
}

type SolidPattern struct{ Pattern }

func CreatePatternRGB(red, green, blue float64) *SolidPattern {
	return cairo_pattern_create_rgb.Fn()(red, green, blue)
}
func CreatePatternRGBA(red, green, blue, alpha float64) *SolidPattern {
	return cairo_pattern_create_rgba.Fn()(red, green, blue, alpha)
}

type SurfacePattern struct{ Pattern }

func CreatePatternForSurface(surface SurfaceIface) *Pattern {
	return cairo_pattern_create_for_surface.Fn()(GetSurfaceIface(surface))
}

type GradientPattern struct{ Pattern }

type LinearPattern struct{ GradientPattern }

func CreatePatternLinear(x0, y0, x1, y1 float64) *Pattern {
	return cairo_pattern_create_linear.Fn()(x0, y0, x1, y1)
}

type RadialPattern struct{ GradientPattern }

func CreatePatternRadial(cx0, cy0, radius0, cx1, cy1, radius1 float64) *Pattern {
	return cairo_pattern_create_radial.Fn()(cx0, cy0, radius0, cx1, cy1, radius1)
}

type MeshPattern struct{ Pattern }

func CreatePatternMesh() *Pattern {
	return cairo_pattern_create_mesh.Fn()()
}
func (p *Pattern) Reference() *Pattern {
	return cairo_pattern_reference.Fn()(p)
}
func (p *Pattern) Destroy() {
	cairo_pattern_destroy.Fn()(p)
}
func (p *Pattern) GetReferenceCount() uint32 {
	return cairo_pattern_get_reference_count.Fn()(p)
}
func (p *Pattern) Status() Status {
	return cairo_pattern_status.Fn()(p)
}
func (p *Pattern) GetUserData(key *UserDataKey) uptr {
	return cairo_pattern_get_user_data.Fn()(p, key)
}
func (p *Pattern) SetUserData(key *UserDataKey, user_data interface{}, destroy func(uptr)) Status {
	return cairo_pattern_set_user_data.Fn()(p, key, anyptr(user_data), vcbu(destroy))
}

func (p *Pattern) GetType() PatternType { return cairo_pattern_get_type.Fn()(p) }
func (p *GradientPattern) AddColorStopRGB(offset, red, green, blue float64) {
	cairo_pattern_add_color_stop_rgb.Fn()(p, offset, red, green, blue)
}
func (p *GradientPattern) AddColorStopRGBA(offset, red, green, blue, alpha float64) {
	cairo_pattern_add_color_stop_rgba.Fn()(p, offset, red, green, blue, alpha)
}

func (p *MeshPattern) BeginPatch() { cairo_mesh_pattern_begin_patch.Fn()(p) }
func (p *MeshPattern) EndPatch()   { cairo_mesh_pattern_end_patch.Fn()(p) }
func (p *MeshPattern) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	cairo_mesh_pattern_curve_to.Fn()(p, x1, y1, x2, y2, x3, y3)
}
func (p *MeshPattern) LineTo(x, y float64) { cairo_mesh_pattern_line_to.Fn()(p, x, y) }
func (p *MeshPattern) MoveTo(x, y float64) { cairo_mesh_pattern_move_to.Fn()(p, x, y) }
func (p *MeshPattern) SetControlPoint(point_num uint32, x, y float64) {
	cairo_mesh_pattern_set_control_point.Fn()(p, point_num, x, y)
}
func (p *MeshPattern) SetCornerColorRGB(corner_num uint32, red, green, blue float64) {
	cairo_mesh_pattern_set_corner_color_rgb.Fn()(p, corner_num, red, green, blue)
}
func (p *MeshPattern) SetCornerColorRGBA(corner_num uint32, red, green, blue, alpha float64) {
	cairo_mesh_pattern_set_corner_color_rgba.Fn()(p, corner_num, red, green, blue, alpha)
}

func (p *Pattern) SetMatrix(matrix *Matrix)   { cairo_pattern_set_matrix.Fn()(p, matrix) }
func (p *Pattern) GetMatrix() (matrix Matrix) { cairo_pattern_get_matrix.Fn()(p, &matrix); return }

func (p *Pattern) SetExtend(extend Extend) { cairo_pattern_set_extend.Fn()(p, extend) }
func (p *Pattern) GetExtend() Extend       { return cairo_pattern_get_extend.Fn()(p) }
func (p *Pattern) SetFilter(filter Filter) { cairo_pattern_set_filter.Fn()(p, filter) }
func (p *Pattern) GetFilter() Filter       { return cairo_pattern_get_filter.Fn()(p) }
func (p *SolidPattern) GetRGBA() (red, green, blue, alpha float64, status Status) {
	status = cairo_pattern_get_rgba.Fn()(p, &red, &green, &blue, &alpha)
	return
}
func (p *SurfacePattern) GetSurface() (surface *Surface, status Status) {
	status = cairo_pattern_get_surface.Fn()(p, &surface)
	return
}
func (p *GradientPattern) GetColorStopRGBA(index int32) (offset, red, green, blue, alpha float64, status Status) {
	status = cairo_pattern_get_color_stop_rgba.Fn()(p, index, &offset, &red, &green, &blue, &alpha)
	return
}
func (p *GradientPattern) GetColorStopCount() (count int32, status Status) {
	status = cairo_pattern_get_color_stop_count.Fn()(p, &count)
	return
}
func (p *LinearPattern) GetLinearPoints() (x0, y0, x1, y1 float64, status Status) {
	status = cairo_pattern_get_linear_points.Fn()(p, &x0, &y0, &x1, &y1)
	return
}
func (p *RadialPattern) GetRadialCircles() (x0, y0, r0, x1, y1, r1 float64, status Status) {
	status = cairo_pattern_get_radial_circles.Fn()(p, &x0, &y0, &r0, &x1, &y1, &r1)
	return
}
func (p *MeshPattern) MeshGetPatchCount() (count uint32, status Status) {
	status = cairo_mesh_pattern_get_patch_count.Fn()(p, &count)
	return
}
func (p *MeshPattern) MeshGetPath(patch_num uint32) *Path {
	return cairo_mesh_pattern_get_path.Fn()(p, patch_num)
}
func (p *MeshPattern) MeshGetCornerColorRGBA(patch_num, corner_num uint32) (red, green, blue, alpha float64, status Status) {
	status = cairo_mesh_pattern_get_corner_color_rgba.Fn()(p, patch_num, corner_num, &red, &green, &blue, &alpha)
	return
}
func (p *MeshPattern) MeshGetControlPoint(patch_num, point_num uint32) (x, y float64, status Status) {
	status = cairo_mesh_pattern_get_control_point.Fn()(p, patch_num, point_num, &x, &y)
	return
}

func MatrixInit(xx, yx, xy, yy, x0, y0 float64) (matrix Matrix) {
	cairo_matrix_init.Fn()(&matrix, xx, yx, xy, yy, x0, y0)
	return
}
func MatrixInitIdentity() (matrix Matrix) {
	cairo_matrix_init_identity.Fn()(&matrix)
	return
}
func MatrixInitTranslate(tx, ty float64) (matrix Matrix) {
	cairo_matrix_init_translate.Fn()(&matrix, tx, ty)
	return
}
func MatrixInitScale(sx, sy float64) (matrix Matrix) {
	cairo_matrix_init_scale.Fn()(&matrix, sx, sy)
	return
}
func MatrixInitRotate(radians float64) (matrix Matrix) {
	cairo_matrix_init_rotate.Fn()(&matrix, radians)
	return
}
func (matrix *Matrix) Translate(tx, ty float64) {
	cairo_matrix_translate.Fn()(matrix, tx, ty)
}
func (matrix *Matrix) Scale(sx, sy float64) {
	cairo_matrix_scale.Fn()(matrix, sx, sy)
}
func (matrix *Matrix) Rotate(radians float64) {
	cairo_matrix_rotate.Fn()(matrix, radians)
}
func (matrix *Matrix) Invert() Status {
	return cairo_matrix_invert.Fn()(matrix)
}
func MatrixMultiply(result, a, b *Matrix) {
	cairo_matrix_multiply.Fn()(result, a, b)
}
func (matrix *Matrix) Multiply(b *Matrix) (result *Matrix) {
	result = &Matrix{}
	cairo_matrix_multiply.Fn()(result, matrix, b)
	return
}
func (matrix *Matrix) TransformDistance(dx, dy *float64) {
	cairo_matrix_transform_distance.Fn()(matrix, dx, dy)
}
func (matrix *Matrix) TransformPoint(x, y *float64) {
	cairo_matrix_transform_point.Fn()(matrix, x, y)
}

type Region struct{}

func CreateRegion() *Region { return cairo_region_create.Fn()() }
func CreateRegionRectangle(rectangle *RectangleInt) *Region {
	return cairo_region_create_rectangle.Fn()(rectangle)
}
func CreateRegionRectangles(rects *RectangleInt, count int32) *Region {
	return cairo_region_create_rectangles.Fn()(rects, count)
}
func (r *Region) Copy() *Region                    { return cairo_region_copy.Fn()(r) }
func (r *Region) Reference() *Region               { return cairo_region_reference.Fn()(r) }
func (r *Region) Destroy()                         { cairo_region_destroy.Fn()(r) }
func RegionEqual(a, b *Region) Bool                { return cairo_region_equal.Fn()(a, b) }
func (r *Region) Equal(b *Region) Bool             { return cairo_region_equal.Fn()(r, b) }
func (r *Region) Status() Status                   { return cairo_region_status.Fn()(r) }
func (r *Region) GetExtents(extents *RectangleInt) { cairo_region_get_extents.Fn()(r, extents) }
func (r *Region) NumRectangles() int32             { return cairo_region_num_rectangles.Fn()(r) }
func (r *Region) GetRectangle(nth int32, rectangle *RectangleInt) {
	cairo_region_get_rectangle.Fn()(r, nth, rectangle)
}
func (r *Region) IsEmpty() Bool { return cairo_region_is_empty.Fn()(r) }
func (r *Region) ContainsRectangle(rectangle *RectangleInt) int32 {
	return cairo_region_contains_rectangle.Fn()(r, rectangle)
}
func (r *Region) ContainsPoint(x, y int32) Bool { return cairo_region_contains_point.Fn()(r, x, y) }
func (r *Region) Translate(dx, dy int32)        { cairo_region_translate.Fn()(r, dx, dy) }
func (r *Region) Subtract(other *Region) Status { return cairo_region_subtract.Fn()(r, other) }
func (r *Region) SubtractRectangle(rectangle *RectangleInt) Status {
	return cairo_region_subtract_rectangle.Fn()(r, rectangle)
}
func (r *Region) Intersect(other *Region) Status { return cairo_region_intersect.Fn()(r, other) }
func (r *Region) IntersectRectangle(rectangle *RectangleInt) Status {
	return cairo_region_intersect_rectangle.Fn()(r, rectangle)
}
func (r *Region) Union(other *Region) Status { return cairo_region_union.Fn()(r, other) }
func (r *Region) UnionRectangle(rectangle *RectangleInt) Status {
	return cairo_region_union_rectangle.Fn()(r, rectangle)
}
func (r *Region) Xor(other *Region) Status { return cairo_region_xor.Fn()(r, other) }
func (r *Region) XorRectangle(rectangle *RectangleInt) Status {
	return cairo_region_xor_rectangle.Fn()(r, rectangle)
}

func DebugResetStaticData() { cairo_debug_reset_static_data.Fn()() }

// #endregion

// #region cairo-gobject.h

func GTypeContext() gobject.GType      { return cairo_gobject_context_get_type.Fn()() }
func GTypeDevice() gobject.GType       { return cairo_gobject_device_get_type.Fn()() }
func GTypeMatrix() gobject.GType       { return cairo_gobject_matrix_get_type.Fn()() }
func GTypePattern() gobject.GType      { return cairo_gobject_pattern_get_type.Fn()() }
func GTypeSurface() gobject.GType      { return cairo_gobject_surface_get_type.Fn()() }
func GTypeRectangle() gobject.GType    { return cairo_gobject_rectangle_get_type.Fn()() }
func GTypeScaledFont() gobject.GType   { return cairo_gobject_scaled_font_get_type.Fn()() }
func GTypeFontFace() gobject.GType     { return cairo_gobject_font_face_get_type.Fn()() }
func GTypeFontOptions() gobject.GType  { return cairo_gobject_font_options_get_type.Fn()() }
func GTypeRectangleInt() gobject.GType { return cairo_gobject_rectangle_int_get_type.Fn()() }
func GTypeRegion() gobject.GType       { return cairo_gobject_region_get_type.Fn()() }
func GTypeGlyph() gobject.GType        { return cairo_gobject_glyph_get_type.Fn()() }
func GTypeTextCluster() gobject.GType  { return cairo_gobject_text_cluster_get_type.Fn()() }

func GTypeStatus() gobject.GType           { return cairo_gobject_status_get_type.Fn()() }
func GTypeContent() gobject.GType          { return cairo_gobject_content_get_type.Fn()() }
func GTypeOperator() gobject.GType         { return cairo_gobject_operator_get_type.Fn()() }
func GTypeAntialias() gobject.GType        { return cairo_gobject_antialias_get_type.Fn()() }
func GTypeFillRule() gobject.GType         { return cairo_gobject_fill_rule_get_type.Fn()() }
func GTypeLineCap() gobject.GType          { return cairo_gobject_line_cap_get_type.Fn()() }
func GTypeLineJoin() gobject.GType         { return cairo_gobject_line_join_get_type.Fn()() }
func GTypeTextClusterFlags() gobject.GType { return cairo_gobject_text_cluster_flags_get_type.Fn()() }
func GTypeFontSlant() gobject.GType        { return cairo_gobject_font_slant_get_type.Fn()() }
func GTypeFontWeight() gobject.GType       { return cairo_gobject_font_weight_get_type.Fn()() }
func GTypeSubpixelOrder() gobject.GType    { return cairo_gobject_subpixel_order_get_type.Fn()() }
func GTypeHintStyle() gobject.GType        { return cairo_gobject_hint_style_get_type.Fn()() }
func GTypeHintMetrics() gobject.GType      { return cairo_gobject_hint_metrics_get_type.Fn()() }
func GTypeFontType() gobject.GType         { return cairo_gobject_font_type_get_type.Fn()() }
func GTypePathDataType() gobject.GType     { return cairo_gobject_path_data_type_get_type.Fn()() }
func GTypeDeviceType() gobject.GType       { return cairo_gobject_device_type_get_type.Fn()() }
func GTypeSurfaceType() gobject.GType      { return cairo_gobject_surface_type_get_type.Fn()() }
func GTypeFormat() gobject.GType           { return cairo_gobject_format_get_type.Fn()() }
func GTypePatternType() gobject.GType      { return cairo_gobject_pattern_type_get_type.Fn()() }
func GTypeExtend() gobject.GType           { return cairo_gobject_extend_get_type.Fn()() }
func GTypeFilter() gobject.GType           { return cairo_gobject_filter_get_type.Fn()() }
func GTypeRegionOverlap() gobject.GType    { return cairo_gobject_region_overlap_get_type.Fn()() }

// #endregion

// #region cairo-pdf.h

type PDFSurface struct{ Surface }

func CreatePdfSurface(filename string, width_in_points, height_in_points float64) *PDFSurface {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	return cairo_pdf_surface_create.Fn()(cFilename, width_in_points, height_in_points)
}
func PDFSurfaceCreateForStream(write_func func(closure, data uptr, length uint32) Status, closure interface{},
	width_in_points, height_in_points float64) *Surface {
	return cairo_pdf_surface_create_for_stream.Fn()(vcbu(write_func), anyptr(closure), width_in_points, height_in_points)
}
func (surface *PDFSurface) RestrictToVersion(version PDFVersion) {
	cairo_pdf_surface_restrict_to_version.Fn()(surface, version)
}
func PDFGetVersions() (versions []PDFVersion) {
	var ptr *PDFVersion
	var n int32
	cairo_pdf_get_versions.Fn()(&ptr, &n)
	if ptr != nil && n > 0 {
		versions = *(*[]PDFVersion)(slice(uptr(ptr), int(n)))
	}
	return
}
func PDFVersionToString(version PDFVersion) string {
	return cairo_pdf_version_to_string.Fn()(version).String()
}
func (surface *PDFSurface) SetSize(width_in_points, height_in_points float64) {
	cairo_pdf_surface_set_size.Fn()(surface, width_in_points, height_in_points)
}
func (surface *PDFSurface) AddOutline(parent_id int32, utf8, link_attribs string, flags PDFOutlineFlags) int32 {
	cUtf8 := cc.NewString(utf8)
	cLink := cc.NewString(link_attribs)
	defer cUtf8.Free()
	defer cLink.Free()
	return cairo_pdf_surface_add_outline.Fn()(surface, parent_id, cUtf8, cLink, flags)
}
func (surface *PDFSurface) SetMetadata(metadata PDFMetadata, utf8 string) {
	cUtf8 := cc.NewString(utf8)
	defer cUtf8.Free()
	cairo_pdf_surface_set_metadata.Fn()(surface, metadata, cUtf8)
}
func (surface *PDFSurface) SetCustomMetadata(name, value string) {
	cName := cc.NewString(name)
	cValue := cc.NewString(value)
	defer cName.Free()
	defer cValue.Free()
	cairo_pdf_surface_set_custom_metadata.Fn()(surface, cName, cValue)
}
func (surface *PDFSurface) SetPageLabel(utf8 string) {
	cUtf8 := cc.NewString(utf8)
	defer cUtf8.Free()
	cairo_pdf_surface_set_page_label.Fn()(surface, cUtf8)
}
func (surface *PDFSurface) SetThumbnailSize(width, height int32) {
	cairo_pdf_surface_set_thumbnail_size.Fn()(surface, width, height)
}

// #endregion

// #region cairo-ps.h

type PSSurface struct{ Surface }

func CreatePSSurface(filename string, width_in_points, height_in_points float64) *PSSurface {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	return cairo_ps_surface_create.Fn()(cFilename, width_in_points, height_in_points)
}

func CreatePSSurfaceForStream(
	write_func func(closure, data uptr, length uint32) Status,
	closure interface{}, width_in_points, height_in_points float64) *PSSurface {
	return cairo_ps_surface_create_for_stream.Fn()(vcbu(write_func), anyptr(closure), width_in_points, height_in_points)
}

func (s *PSSurface) RestrictToLevel(level PSLevel) {
	cairo_ps_surface_restrict_to_level.Fn()(s, level)
}

func PSGetLevels() (levels []PSLevel) {
	var ptr *PSLevel
	var n int32
	cairo_ps_get_levels.Fn()(&ptr, &n)
	if ptr != nil && n > 0 {
		levels = *(*[]PSLevel)(slice(uptr(ptr), int(n)))
	}
	return
}

func PsLevelToString(level PSLevel) string {
	return cairo_ps_level_to_string.Fn()(level).String()
}

func (level PSLevel) String() string {
	return cairo_ps_level_to_string.Fn()(level).String()
}
func (s *PSSurface) SetEps(eps bool) {
	var v Bool
	if eps {
		v = 1
	}
	cairo_ps_surface_set_eps.Fn()(s, v)
}
func (s *PSSurface) GetEps() Bool { return cairo_ps_surface_get_eps.Fn()(s) }
func (s *PSSurface) SetSize(width_in_points, height_in_points float64) {
	cairo_ps_surface_set_size.Fn()(s, width_in_points, height_in_points)
}
func (s *PSSurface) DscComment(comment string) {
	cComment := cc.NewString(comment)
	defer cComment.Free()
	cairo_ps_surface_dsc_comment.Fn()(s, cComment)
}
func (s *PSSurface) DscBeginSetup()     { cairo_ps_surface_dsc_begin_setup.Fn()(s) }
func (s *PSSurface) DscBeginPageSetup() { cairo_ps_surface_dsc_begin_page_setup.Fn()(s) }

// #endregion

// #region cairo-script.h

type Script struct{ Device }

func CreateScriptDevice(filename string) *Script {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	return cairo_script_create.Fn()(cFilename)
}
func CreateScriptDeviceForStream(
	write_func func(closure, data uptr, length uint32) Status,
	closure interface{}) *Script {
	return cairo_script_create_for_stream.Fn()(vcbu(write_func), anyptr(closure))
}
func (d *Script) WriteComment(comment string) {
	cComment := cc.NewString(comment)
	defer cComment.Free()
	cairo_script_write_comment.Fn()(d, cComment, int32(len(comment)))
}
func (d *Script) SetMode(mode ScriptMode) { cairo_script_set_mode.Fn()(d, mode) }
func (d *Script) GetMode() ScriptMode     { return cairo_script_get_mode.Fn()(d) }
func (d *Script) SurfaceCreate(content Content, width, height float64) *Surface {
	return cairo_script_surface_create.Fn()(d, content, width, height)
}
func (d *Script) SurfaceCreateForTarget(target SurfaceIface) *Surface {
	return cairo_script_surface_create_for_target.Fn()(d, GetSurfaceIface(target))
}
func (d *Script) FromRecordingSurface(recording_surface SurfaceIface) Status {
	return cairo_script_from_recording_surface.Fn()(d, GetSurfaceIface(recording_surface))
}

// #endregion

// #region cairo-svg.h

type SVGSurface struct{ Surface }

func CreateSVGSurface(filename string, width_in_points, height_in_points float64) *SVGSurface {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	return cairo_svg_surface_create.Fn()(cFilename, width_in_points, height_in_points)
}

func CreateSVGSurfaceForStream(write_func func(closure, data uptr, length uint32) Status, closure interface{}, width_in_points, height_in_points float64) *SVGSurface {
	return cairo_svg_surface_create_for_stream.Fn()(vcbu(write_func), anyptr(closure), width_in_points, height_in_points)
}
func (s *SVGSurface) RestrictToVersion(version SVGVersion) {
	cairo_svg_surface_restrict_to_version.Fn()(s, version)
}
func SVGGetVersions() (versions []SVGVersion) {
	var ptr *SVGVersion
	var n int32
	cairo_svg_get_versions.Fn()(&ptr, &n)
	if ptr != nil && n > 0 {
		versions = *(*[]SVGVersion)(slice(uptr(ptr), int(n)))
	}
	return
}
func SVGVersionToString(version SVGVersion) string {
	return cairo_svg_version_to_string.Fn()(version).String()
}
func (version SVGVersion) String() string {
	return cairo_svg_version_to_string.Fn()(version).String()
}
func (s *SVGSurface) SetDocumentUnit(unit SVGUnit) { cairo_svg_surface_set_document_unit.Fn()(s, unit) }
func (s *SVGSurface) GetDocumentUnit() SVGUnit     { return cairo_svg_surface_get_document_unit.Fn()(s) }

// #endregion
