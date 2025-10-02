package graphene

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
)

type uptr = unsafe.Pointer
type iptr = uintptr

func init() { cc.Open("libgraphene-1*") }

var (
	// #region Point
	graphene_point_alloc           = cc.DlFunc[func() *Point, *Point]{Name: "graphene_point_alloc"}
	graphene_point_free            = cc.DlFunc[func(p *Point), cc.Void]{Name: "graphene_point_free"}
	graphene_point_init            = cc.DlFunc[func(p *Point, x float32, y float32) *Point, *Point]{Name: "graphene_point_init"}
	graphene_point_init_from_point = cc.DlFunc[func(p *Point, src *Point) *Point, *Point]{Name: "graphene_point_init_from_point"}
	graphene_point_init_from_vec2  = cc.DlFunc[func(p *Point, src *Vec2) *Point, *Point]{Name: "graphene_point_init_from_vec2"}
	graphene_point_equal           = cc.DlFunc[func(a *Point, b *Point) bool, bool]{Name: "graphene_point_equal"}
	graphene_point_distance        = cc.DlFunc[func(a *Point, b *Point, dX *float32, dY *float32) float32, float32]{Name: "graphene_point_distance"}
	graphene_point_near            = cc.DlFunc[func(a *Point, b *Point, epsilon float32) bool, bool]{Name: "graphene_point_near"}
	graphene_point_interpolate     = cc.DlFunc[func(a *Point, b *Point, factor float64, res *Point), cc.Void]{Name: "graphene_point_interpolate"}
	graphene_point_to_vec2         = cc.DlFunc[func(p *Point, v *Vec2), cc.Void]{Name: "graphene_point_to_vec2"}
	graphene_point_zero            = cc.DlFunc[func() *Point, *Point]{Name: "graphene_point_zero"}
	// #endregion

	// #region Point3D
	graphene_point3d_alloc              = cc.DlFunc[func() *Point3D, *Point3D]{Name: "graphene_point3d_alloc"}
	graphene_point3d_free               = cc.DlFunc[func(p *Point3D), cc.Void]{Name: "graphene_point3d_free"}
	graphene_point3d_init               = cc.DlFunc[func(p *Point3D, x, y, z float32) *Point3D, *Point3D]{Name: "graphene_point3d_init"}
	graphene_point3d_init_from_point    = cc.DlFunc[func(p *Point3D, src *Point3D) *Point3D, *Point3D]{Name: "graphene_point3d_init_from_point"}
	graphene_point3d_init_from_vec3     = cc.DlFunc[func(p *Point3D, v *Vec3) *Point3D, *Point3D]{Name: "graphene_point3d_init_from_vec3"}
	graphene_point3d_to_vec3            = cc.DlFunc[func(p *Point3D, v *Vec3), cc.Void]{Name: "graphene_point3d_to_vec3"}
	graphene_point3d_equal              = cc.DlFunc[func(a, b *Point3D) bool, bool]{Name: "graphene_point3d_equal"}
	graphene_point3d_near               = cc.DlFunc[func(a, b *Point3D, epsilon float32) bool, bool]{Name: "graphene_point3d_near"}
	graphene_point3d_scale              = cc.DlFunc[func(p *Point3D, factor float32, res *Point3D), cc.Void]{Name: "graphene_point3d_scale"}
	graphene_point3d_cross              = cc.DlFunc[func(a, b *Point3D, res *Point3D), cc.Void]{Name: "graphene_point3d_cross"}
	graphene_point3d_dot                = cc.DlFunc[func(a, b *Point3D) float32, float32]{Name: "graphene_point3d_dot"}
	graphene_point3d_length             = cc.DlFunc[func(p *Point3D) float32, float32]{Name: "graphene_point3d_length"}
	graphene_point3d_normalize          = cc.DlFunc[func(p *Point3D, res *Point3D), cc.Void]{Name: "graphene_point3d_normalize"}
	graphene_point3d_distance           = cc.DlFunc[func(a, b *Point3D, delta *Vec3) float32, float32]{Name: "graphene_point3d_distance"}
	graphene_point3d_interpolate        = cc.DlFunc[func(a, b *Point3D, factor float64, res *Point3D), cc.Void]{Name: "graphene_point3d_interpolate"}
	graphene_point3d_normalize_viewport = cc.DlFunc[func(p *Point3D, viewport *Rect, zNear, zFar float32, res *Point3D), cc.Void]{Name: "graphene_point3d_normalize_viewport"}
	graphene_point3d_zero               = cc.DlFunc[func() *Point3D, *Point3D]{Name: "graphene_point3d_zero"}
	// #endregion

	// #region Rect
	graphene_rect_alloc            = cc.DlFunc[func() *Rect, *Rect]{Name: "graphene_rect_alloc"}
	graphene_rect_free             = cc.DlFunc[func(r *Rect), cc.Void]{Name: "graphene_rect_free"}
	graphene_rect_init             = cc.DlFunc[func(r *Rect, x, y, width, height float32) *Rect, *Rect]{Name: "graphene_rect_init"}
	graphene_rect_init_from_rect   = cc.DlFunc[func(r *Rect, src *Rect) *Rect, *Rect]{Name: "graphene_rect_init_from_rect"}
	graphene_rect_equal            = cc.DlFunc[func(a, b *Rect) bool, bool]{Name: "graphene_rect_equal"}
	graphene_rect_normalize        = cc.DlFunc[func(r *Rect) *Rect, *Rect]{Name: "graphene_rect_normalize"}
	graphene_rect_normalize_r      = cc.DlFunc[func(r *Rect, res *Rect), cc.Void]{Name: "graphene_rect_normalize_r"}
	graphene_rect_get_center       = cc.DlFunc[func(r *Rect, p *Point), cc.Void]{Name: "graphene_rect_get_center"}
	graphene_rect_get_top_left     = cc.DlFunc[func(r *Rect, p *Point), cc.Void]{Name: "graphene_rect_get_top_left"}
	graphene_rect_get_top_right    = cc.DlFunc[func(r *Rect, p *Point), cc.Void]{Name: "graphene_rect_get_top_right"}
	graphene_rect_get_bottom_right = cc.DlFunc[func(r *Rect, p *Point), cc.Void]{Name: "graphene_rect_get_bottom_right"}
	graphene_rect_get_bottom_left  = cc.DlFunc[func(r *Rect, p *Point), cc.Void]{Name: "graphene_rect_get_bottom_left"}
	graphene_rect_get_vertices     = cc.DlFunc[func(r *Rect, vertices *Vec2), cc.Void]{Name: "graphene_rect_get_vertices"}
	graphene_rect_get_x            = cc.DlFunc[func(r *Rect) float32, float32]{Name: "graphene_rect_get_x"}
	graphene_rect_get_y            = cc.DlFunc[func(r *Rect) float32, float32]{Name: "graphene_rect_get_y"}
	graphene_rect_get_width        = cc.DlFunc[func(r *Rect) float32, float32]{Name: "graphene_rect_get_width"}
	graphene_rect_get_height       = cc.DlFunc[func(r *Rect) float32, float32]{Name: "graphene_rect_get_height"}
	graphene_rect_get_area         = cc.DlFunc[func(r *Rect) float32, float32]{Name: "graphene_rect_get_area"}
	graphene_rect_union            = cc.DlFunc[func(a, b *Rect, res *Rect), cc.Void]{Name: "graphene_rect_union"}
	graphene_rect_intersection     = cc.DlFunc[func(a, b *Rect, res *Rect) bool, bool]{Name: "graphene_rect_intersection"}
	graphene_rect_contains_point   = cc.DlFunc[func(r *Rect, p *Point) int32, int32]{Name: "graphene_rect_contains_point"}
	graphene_rect_contains_rect    = cc.DlFunc[func(a, b *Rect) bool, bool]{Name: "graphene_rect_contains_rect"}
	graphene_rect_offset           = cc.DlFunc[func(r *Rect, dx, dy float32) *Rect, *Rect]{Name: "graphene_rect_offset"}
	graphene_rect_offset_r         = cc.DlFunc[func(r *Rect, dx, dy float32, res *Rect), cc.Void]{Name: "graphene_rect_offset_r"}
	graphene_rect_inset            = cc.DlFunc[func(r *Rect, dx, dy float32) *Rect, *Rect]{Name: "graphene_rect_inset"}
	graphene_rect_inset_r          = cc.DlFunc[func(r *Rect, dx, dy float32, res *Rect), cc.Void]{Name: "graphene_rect_inset_r"}
	graphene_rect_round            = cc.DlFunc[func(r *Rect, res *Rect), cc.Void]{Name: "graphene_rect_round"}
	graphene_rect_round_extents    = cc.DlFunc[func(r *Rect, res *Rect), cc.Void]{Name: "graphene_rect_round_extents"}
	graphene_rect_interpolate      = cc.DlFunc[func(a, b *Rect, factor float64, res *Rect), cc.Void]{Name: "graphene_rect_interpolate"}
	graphene_rect_expand           = cc.DlFunc[func(r *Rect, p *Point, res *Rect), cc.Void]{Name: "graphene_rect_expand"}
	graphene_rect_zero             = cc.DlFunc[func() *Rect, *Rect]{Name: "graphene_rect_zero"}
	graphene_rect_scale            = cc.DlFunc[func(r *Rect, sh, sv float32, res *Rect), cc.Void]{Name: "graphene_rect_scale"}
	// #endregion

	// #region Vec2
	graphene_vec2_alloc           = cc.DlFunc[func() *Vec2, *Vec2]{Name: "graphene_vec2_alloc"}
	graphene_vec2_free            = cc.DlFunc[func(v *Vec2), cc.Void]{Name: "graphene_vec2_free"}
	graphene_vec2_init            = cc.DlFunc[func(v *Vec2, x float32, y float32) *Vec2, *Vec2]{Name: "graphene_vec2_init"}
	graphene_vec2_init_from_vec2  = cc.DlFunc[func(v *Vec2, src *Vec2) *Vec2, *Vec2]{Name: "graphene_vec2_init_from_vec2"}
	graphene_vec2_init_from_float = cc.DlFunc[func(v *Vec2, src *float32) *Vec2, *Vec2]{Name: "graphene_vec2_init_from_float"}
	graphene_vec2_to_float        = cc.DlFunc[func(v *Vec2, dest *float32), cc.Void]{Name: "graphene_vec2_to_float"}
	graphene_vec2_add             = cc.DlFunc[func(a *Vec2, b *Vec2, res *Vec2), cc.Void]{Name: "graphene_vec2_add"}
	graphene_vec2_subtract        = cc.DlFunc[func(a *Vec2, b *Vec2, res *Vec2), cc.Void]{Name: "graphene_vec2_subtract"}
	graphene_vec2_multiply        = cc.DlFunc[func(a *Vec2, b *Vec2, res *Vec2), cc.Void]{Name: "graphene_vec2_multiply"}
	graphene_vec2_divide          = cc.DlFunc[func(a *Vec2, b *Vec2, res *Vec2), cc.Void]{Name: "graphene_vec2_divide"}
	graphene_vec2_dot             = cc.DlFunc[func(a *Vec2, b *Vec2) float32, float32]{Name: "graphene_vec2_dot"}
	graphene_vec2_length          = cc.DlFunc[func(v *Vec2) float32, float32]{Name: "graphene_vec2_length"}
	graphene_vec2_normalize       = cc.DlFunc[func(v *Vec2, res *Vec2), cc.Void]{Name: "graphene_vec2_normalize"}
	graphene_vec2_scale           = cc.DlFunc[func(v *Vec2, factor float32, res *Vec2), cc.Void]{Name: "graphene_vec2_scale"}
	graphene_vec2_negate          = cc.DlFunc[func(v *Vec2, res *Vec2), cc.Void]{Name: "graphene_vec2_negate"}
	graphene_vec2_near            = cc.DlFunc[func(v1 *Vec2, v2 *Vec2, epsilon float32) bool, bool]{Name: "graphene_vec2_near"}
	graphene_vec2_equal           = cc.DlFunc[func(v1 *Vec2, v2 *Vec2) bool, bool]{Name: "graphene_vec2_equal"}
	graphene_vec2_min             = cc.DlFunc[func(a *Vec2, b *Vec2, res *Vec2), cc.Void]{Name: "graphene_vec2_min"}
	graphene_vec2_max             = cc.DlFunc[func(a *Vec2, b *Vec2, res *Vec2), cc.Void]{Name: "graphene_vec2_max"}
	graphene_vec2_interpolate     = cc.DlFunc[func(v1 *Vec2, v2 *Vec2, factor float64, res *Vec2), cc.Void]{Name: "graphene_vec2_interpolate"}
	graphene_vec2_get_x           = cc.DlFunc[func(v *Vec2) float32, float32]{Name: "graphene_vec2_get_x"}
	graphene_vec2_get_y           = cc.DlFunc[func(v *Vec2) float32, float32]{Name: "graphene_vec2_get_y"}
	graphene_vec2_zero            = cc.DlFunc[func() *Vec2, *Vec2]{Name: "graphene_vec2_zero"}
	graphene_vec2_one             = cc.DlFunc[func() *Vec2, *Vec2]{Name: "graphene_vec2_one"}
	graphene_vec2_x_axis          = cc.DlFunc[func() *Vec2, *Vec2]{Name: "graphene_vec2_x_axis"}
	graphene_vec2_y_axis          = cc.DlFunc[func() *Vec2, *Vec2]{Name: "graphene_vec2_y_axis"}
	// #endregion

	// #region Vec3
	graphene_vec3_alloc           = cc.DlFunc[func() *Vec3, *Vec3]{Name: "graphene_vec3_alloc"}
	graphene_vec3_free            = cc.DlFunc[func(v *Vec3), cc.Void]{Name: "graphene_vec3_free"}
	graphene_vec3_init            = cc.DlFunc[func(v *Vec3, x float32, y float32, z float32) *Vec3, *Vec3]{Name: "graphene_vec3_init"}
	graphene_vec3_init_from_vec3  = cc.DlFunc[func(v *Vec3, src *Vec3) *Vec3, *Vec3]{Name: "graphene_vec3_init_from_vec3"}
	graphene_vec3_init_from_float = cc.DlFunc[func(v *Vec3, src *float32) *Vec3, *Vec3]{Name: "graphene_vec3_init_from_float"}
	graphene_vec3_to_float        = cc.DlFunc[func(v *Vec3, dest *float32), cc.Void]{Name: "graphene_vec3_to_float"}
	graphene_vec3_add             = cc.DlFunc[func(a *Vec3, b *Vec3, res *Vec3), cc.Void]{Name: "graphene_vec3_add"}
	graphene_vec3_subtract        = cc.DlFunc[func(a *Vec3, b *Vec3, res *Vec3), cc.Void]{Name: "graphene_vec3_subtract"}
	graphene_vec3_multiply        = cc.DlFunc[func(a *Vec3, b *Vec3, res *Vec3), cc.Void]{Name: "graphene_vec3_multiply"}
	graphene_vec3_divide          = cc.DlFunc[func(a *Vec3, b *Vec3, res *Vec3), cc.Void]{Name: "graphene_vec3_divide"}
	graphene_vec3_cross           = cc.DlFunc[func(a *Vec3, b *Vec3, res *Vec3), cc.Void]{Name: "graphene_vec3_cross"}
	graphene_vec3_dot             = cc.DlFunc[func(a *Vec3, b *Vec3) float32, float32]{Name: "graphene_vec3_dot"}
	graphene_vec3_length          = cc.DlFunc[func(v *Vec3) float32, float32]{Name: "graphene_vec3_length"}
	graphene_vec3_normalize       = cc.DlFunc[func(v *Vec3, res *Vec3), cc.Void]{Name: "graphene_vec3_normalize"}
	graphene_vec3_scale           = cc.DlFunc[func(v *Vec3, factor float32, res *Vec3), cc.Void]{Name: "graphene_vec3_scale"}
	graphene_vec3_negate          = cc.DlFunc[func(v *Vec3, res *Vec3), cc.Void]{Name: "graphene_vec3_negate"}
	graphene_vec3_equal           = cc.DlFunc[func(v1 *Vec3, v2 *Vec3) bool, bool]{Name: "graphene_vec3_equal"}
	graphene_vec3_near            = cc.DlFunc[func(v1 *Vec3, v2 *Vec3, epsilon float32) bool, bool]{Name: "graphene_vec3_near"}
	graphene_vec3_min             = cc.DlFunc[func(a *Vec3, b *Vec3, res *Vec3), cc.Void]{Name: "graphene_vec3_min"}
	graphene_vec3_max             = cc.DlFunc[func(a *Vec3, b *Vec3, res *Vec3), cc.Void]{Name: "graphene_vec3_max"}
	graphene_vec3_interpolate     = cc.DlFunc[func(v1 *Vec3, v2 *Vec3, factor float64, res *Vec3), cc.Void]{Name: "graphene_vec3_interpolate"}
	graphene_vec3_get_x           = cc.DlFunc[func(v *Vec3) float32, float32]{Name: "graphene_vec3_get_x"}
	graphene_vec3_get_y           = cc.DlFunc[func(v *Vec3) float32, float32]{Name: "graphene_vec3_get_y"}
	graphene_vec3_get_z           = cc.DlFunc[func(v *Vec3) float32, float32]{Name: "graphene_vec3_get_z"}
	graphene_vec3_get_xy          = cc.DlFunc[func(v *Vec3, res *Vec2), cc.Void]{Name: "graphene_vec3_get_xy"}
	graphene_vec3_get_xy0         = cc.DlFunc[func(v *Vec3, res *Vec3), cc.Void]{Name: "graphene_vec3_get_xy0"}
	graphene_vec3_get_xyz0        = cc.DlFunc[func(v *Vec3, res *Vec4), cc.Void]{Name: "graphene_vec3_get_xyz0"}
	graphene_vec3_get_xyz1        = cc.DlFunc[func(v *Vec3, res *Vec4), cc.Void]{Name: "graphene_vec3_get_xyz1"}
	graphene_vec3_get_xyzw        = cc.DlFunc[func(v *Vec3, w float32, res *Vec4), cc.Void]{Name: "graphene_vec3_get_xyzw"}
	graphene_vec3_zero            = cc.DlFunc[func() *Vec3, *Vec3]{Name: "graphene_vec3_zero"}
	graphene_vec3_one             = cc.DlFunc[func() *Vec3, *Vec3]{Name: "graphene_vec3_one"}
	graphene_vec3_x_axis          = cc.DlFunc[func() *Vec3, *Vec3]{Name: "graphene_vec3_x_axis"}
	graphene_vec3_y_axis          = cc.DlFunc[func() *Vec3, *Vec3]{Name: "graphene_vec3_y_axis"}
	graphene_vec3_z_axis          = cc.DlFunc[func() *Vec3, *Vec3]{Name: "graphene_vec3_z_axis"}
	// #endregion

	// #region Vec4
	graphene_vec4_alloc           = cc.DlFunc[func() *Vec4, *Vec4]{Name: "graphene_vec4_alloc"}
	graphene_vec4_free            = cc.DlFunc[func(v *Vec4), cc.Void]{Name: "graphene_vec4_free"}
	graphene_vec4_init            = cc.DlFunc[func(v *Vec4, x float32, y float32, z float32, w float32) *Vec4, *Vec4]{Name: "graphene_vec4_init"}
	graphene_vec4_init_from_vec4  = cc.DlFunc[func(v *Vec4, src *Vec4) *Vec4, *Vec4]{Name: "graphene_vec4_init_from_vec4"}
	graphene_vec4_init_from_vec3  = cc.DlFunc[func(v *Vec4, src *Vec3, w float32) *Vec4, *Vec4]{Name: "graphene_vec4_init_from_vec3"}
	graphene_vec4_init_from_vec2  = cc.DlFunc[func(v *Vec4, src *Vec2, z, w float32) *Vec4, *Vec4]{Name: "graphene_vec4_init_from_vec2"}
	graphene_vec4_init_from_float = cc.DlFunc[func(v *Vec4, src *float32) *Vec4, *Vec4]{Name: "graphene_vec4_init_from_float"}
	graphene_vec4_to_float        = cc.DlFunc[func(v *Vec4, dest *float32), cc.Void]{Name: "graphene_vec4_to_float"}
	graphene_vec4_add             = cc.DlFunc[func(a *Vec4, b *Vec4, res *Vec4), cc.Void]{Name: "graphene_vec4_add"}
	graphene_vec4_subtract        = cc.DlFunc[func(a *Vec4, b *Vec4, res *Vec4), cc.Void]{Name: "graphene_vec4_subtract"}
	graphene_vec4_multiply        = cc.DlFunc[func(a *Vec4, b *Vec4, res *Vec4), cc.Void]{Name: "graphene_vec4_multiply"}
	graphene_vec4_divide          = cc.DlFunc[func(a *Vec4, b *Vec4, res *Vec4), cc.Void]{Name: "graphene_vec4_divide"}
	graphene_vec4_dot             = cc.DlFunc[func(a *Vec4, b *Vec4) float32, float32]{Name: "graphene_vec4_dot"}
	graphene_vec4_length          = cc.DlFunc[func(v *Vec4) float32, float32]{Name: "graphene_vec4_length"}
	graphene_vec4_normalize       = cc.DlFunc[func(v *Vec4, res *Vec4), cc.Void]{Name: "graphene_vec4_normalize"}
	graphene_vec4_scale           = cc.DlFunc[func(v *Vec4, factor float32, res *Vec4), cc.Void]{Name: "graphene_vec4_scale"}
	graphene_vec4_negate          = cc.DlFunc[func(v *Vec4, res *Vec4), cc.Void]{Name: "graphene_vec4_negate"}
	graphene_vec4_equal           = cc.DlFunc[func(v1 *Vec4, v2 *Vec4) bool, bool]{Name: "graphene_vec4_equal"}
	graphene_vec4_near            = cc.DlFunc[func(v1 *Vec4, v2 *Vec4, epsilon float32) bool, bool]{Name: "graphene_vec4_near"}
	graphene_vec4_min             = cc.DlFunc[func(a *Vec4, b *Vec4, res *Vec4), cc.Void]{Name: "graphene_vec4_min"}
	graphene_vec4_max             = cc.DlFunc[func(a *Vec4, b *Vec4, res *Vec4), cc.Void]{Name: "graphene_vec4_max"}
	graphene_vec4_interpolate     = cc.DlFunc[func(v1 *Vec4, v2 *Vec4, factor float64, res *Vec4), cc.Void]{Name: "graphene_vec4_interpolate"}
	graphene_vec4_get_x           = cc.DlFunc[func(v *Vec4) float32, float32]{Name: "graphene_vec4_get_x"}
	graphene_vec4_get_y           = cc.DlFunc[func(v *Vec4) float32, float32]{Name: "graphene_vec4_get_y"}
	graphene_vec4_get_z           = cc.DlFunc[func(v *Vec4) float32, float32]{Name: "graphene_vec4_get_z"}
	graphene_vec4_get_w           = cc.DlFunc[func(v *Vec4) float32, float32]{Name: "graphene_vec4_get_w"}
	graphene_vec4_get_xy          = cc.DlFunc[func(v *Vec4, res *Vec2), cc.Void]{Name: "graphene_vec4_get_xy"}
	graphene_vec4_get_xyz         = cc.DlFunc[func(v *Vec4, res *Vec3), cc.Void]{Name: "graphene_vec4_get_xyz"}
	graphene_vec4_zero            = cc.DlFunc[func() *Vec4, *Vec4]{Name: "graphene_vec4_zero"}
	graphene_vec4_one             = cc.DlFunc[func() *Vec4, *Vec4]{Name: "graphene_vec4_one"}
	graphene_vec4_x_axis          = cc.DlFunc[func() *Vec4, *Vec4]{Name: "graphene_vec4_x_axis"}
	graphene_vec4_y_axis          = cc.DlFunc[func() *Vec4, *Vec4]{Name: "graphene_vec4_y_axis"}
	graphene_vec4_z_axis          = cc.DlFunc[func() *Vec4, *Vec4]{Name: "graphene_vec4_z_axis"}
	graphene_vec4_w_axis          = cc.DlFunc[func() *Vec4, *Vec4]{Name: "graphene_vec4_w_axis"}
	// #endregion

	//#region Matrix
	graphene_matrix_alloc               = cc.DlFunc[func() *Matrix, *Matrix]{Name: "graphene_matrix_alloc"}
	graphene_matrix_free                = cc.DlFunc[func(m *Matrix), cc.Void]{Name: "graphene_matrix_free"}
	graphene_matrix_init_identity       = cc.DlFunc[func(m *Matrix) *Matrix, *Matrix]{Name: "graphene_matrix_init_identity"}
	graphene_matrix_init_from_float     = cc.DlFunc[func(m *Matrix, v *float32) *Matrix, *Matrix]{Name: "graphene_matrix_init_from_float"}
	graphene_matrix_init_from_vec4      = cc.DlFunc[func(m *Matrix, v0, v1, v2, v3 *Vec4) *Matrix, *Matrix]{Name: "graphene_matrix_init_from_vec4"}
	graphene_matrix_init_from_matrix    = cc.DlFunc[func(m *Matrix, src *Matrix) *Matrix, *Matrix]{Name: "graphene_matrix_init_from_matrix"}
	graphene_matrix_init_perspective    = cc.DlFunc[func(m *Matrix, fovy, aspect, z_near, z_far float32) *Matrix, *Matrix]{Name: "graphene_matrix_init_perspective"}
	graphene_matrix_init_ortho          = cc.DlFunc[func(m *Matrix, left, right, top, bottom, z_near, z_far float32) *Matrix, *Matrix]{Name: "graphene_matrix_init_ortho"}
	graphene_matrix_init_look_at        = cc.DlFunc[func(m *Matrix, eye, center, up *Vec3) *Matrix, *Matrix]{Name: "graphene_matrix_init_look_at"}
	graphene_matrix_init_frustum        = cc.DlFunc[func(m *Matrix, left, right, bottom, top, z_near, z_far float32) *Matrix, *Matrix]{Name: "graphene_matrix_init_frustum"}
	graphene_matrix_init_scale          = cc.DlFunc[func(m *Matrix, x, y, z float32) *Matrix, *Matrix]{Name: "graphene_matrix_init_scale"}
	graphene_matrix_init_translate      = cc.DlFunc[func(m *Matrix, p *Point3D) *Matrix, *Matrix]{Name: "graphene_matrix_init_translate"}
	graphene_matrix_init_rotate         = cc.DlFunc[func(m *Matrix, angle float32, axis *Vec3) *Matrix, *Matrix]{Name: "graphene_matrix_init_rotate"}
	graphene_matrix_init_skew           = cc.DlFunc[func(m *Matrix, xSkew, ySkew float32) *Matrix, *Matrix]{Name: "graphene_matrix_init_skew"}
	graphene_matrix_init_from_2d        = cc.DlFunc[func(m *Matrix, xx, yx, xy, yy, x0, y0 float64) *Matrix, *Matrix]{Name: "graphene_matrix_init_from_2d"}
	graphene_matrix_is_identity         = cc.DlFunc[func(m *Matrix) bool, bool]{Name: "graphene_matrix_is_identity"}
	graphene_matrix_is_2d               = cc.DlFunc[func(m *Matrix) bool, bool]{Name: "graphene_matrix_is_2d"}
	graphene_matrix_is_backface_visible = cc.DlFunc[func(m *Matrix) bool, bool]{Name: "graphene_matrix_is_backface_visible"}
	graphene_matrix_is_singular         = cc.DlFunc[func(m *Matrix) bool, bool]{Name: "graphene_matrix_is_singular"}
	graphene_matrix_to_float            = cc.DlFunc[func(m *Matrix, v *float32), cc.Void]{Name: "graphene_matrix_to_float"}
	graphene_matrix_to_2d               = cc.DlFunc[func(m *Matrix, xx, yx, xy, yy, x0, y0 *float64) bool, bool]{Name: "graphene_matrix_to_2d"}
	graphene_matrix_get_row             = cc.DlFunc[func(m *Matrix, index uint, res *Vec4), cc.Void]{Name: "graphene_matrix_get_row"}
	graphene_matrix_get_value           = cc.DlFunc[func(m *Matrix, row, col uint) float32, float32]{Name: "graphene_matrix_get_value"}
	graphene_matrix_multiply            = cc.DlFunc[func(a *Matrix, b *Matrix, res *Matrix), cc.Void]{Name: "graphene_matrix_multiply"}
	graphene_matrix_determinant         = cc.DlFunc[func(m *Matrix) float32, float32]{Name: "graphene_matrix_determinant"}
	graphene_matrix_transform_vec4      = cc.DlFunc[func(m *Matrix, v *Vec4, res *Vec4), cc.Void]{Name: "graphene_matrix_transform_vec4"}
	graphene_matrix_transform_vec3      = cc.DlFunc[func(m *Matrix, v *Vec3, res *Vec3), cc.Void]{Name: "graphene_matrix_transform_vec3"}
	graphene_matrix_transform_point     = cc.DlFunc[func(m *Matrix, p *Point, res *Point), cc.Void]{Name: "graphene_matrix_transform_point"}
	graphene_matrix_transform_point3d   = cc.DlFunc[func(m *Matrix, p *Point3D, res *Point3D), cc.Void]{Name: "graphene_matrix_transform_point3d"}
	graphene_matrix_transform_rect      = cc.DlFunc[func(m *Matrix, r *Rect, res *Quad), cc.Void]{Name: "graphene_matrix_transform_rect"}
	graphene_matrix_transform_bounds    = cc.DlFunc[func(m *Matrix, r *Rect, res *Rect), cc.Void]{Name: "graphene_matrix_transform_bounds"}
	graphene_matrix_transform_sphere    = cc.DlFunc[func(m *Matrix, s *Sphere, res *Sphere), cc.Void]{Name: "graphene_matrix_transform_sphere"}
	graphene_matrix_transform_box       = cc.DlFunc[func(m *Matrix, b *Box, res *Box), cc.Void]{Name: "graphene_matrix_transform_box"}
	graphene_matrix_transform_ray       = cc.DlFunc[func(m *Matrix, r *Ray, res *Ray), cc.Void]{Name: "graphene_matrix_transform_ray"}
	graphene_matrix_project_point       = cc.DlFunc[func(m *Matrix, p *Point, res *Point), cc.Void]{Name: "graphene_matrix_project_point"}
	graphene_matrix_project_rect_bounds = cc.DlFunc[func(m *Matrix, r *Rect, res *Rect), cc.Void]{Name: "graphene_matrix_project_rect_bounds"}
	graphene_matrix_project_rect        = cc.DlFunc[func(m *Matrix, r *Rect, res *Quad), cc.Void]{Name: "graphene_matrix_project_rect"}
	graphene_matrix_untransform_point   = cc.DlFunc[func(m *Matrix, p *Point, bounds *Rect, res *Point) bool, bool]{Name: "graphene_matrix_untransform_point"}
	graphene_matrix_untransform_bounds  = cc.DlFunc[func(m *Matrix, r *Rect, bounds *Rect, res *Rect), cc.Void]{Name: "graphene_matrix_untransform_bounds"}
	graphene_matrix_unproject_point3d   = cc.DlFunc[func(projection *Matrix, modelview *Matrix, point *Point3D, res *Point3D), cc.Void]{Name: "graphene_matrix_unproject_point3d"}
	graphene_matrix_translate           = cc.DlFunc[func(m *Matrix, pos *Point3D), cc.Void]{Name: "graphene_matrix_translate"}
	graphene_matrix_rotate              = cc.DlFunc[func(m *Matrix, angle float32, axis *Vec3), cc.Void]{Name: "graphene_matrix_rotate"}
	graphene_matrix_rotate_x            = cc.DlFunc[func(m *Matrix, angle float32), cc.Void]{Name: "graphene_matrix_rotate_x"}
	graphene_matrix_rotate_y            = cc.DlFunc[func(m *Matrix, angle float32), cc.Void]{Name: "graphene_matrix_rotate_y"}
	graphene_matrix_rotate_z            = cc.DlFunc[func(m *Matrix, angle float32), cc.Void]{Name: "graphene_matrix_rotate_z"}
	graphene_matrix_rotate_quaternion   = cc.DlFunc[func(m *Matrix, q *Quaternion), cc.Void]{Name: "graphene_matrix_rotate_quaternion"}
	graphene_matrix_rotate_euler        = cc.DlFunc[func(m *Matrix, e *Euler), cc.Void]{Name: "graphene_matrix_rotate_euler"}
	graphene_matrix_scale               = cc.DlFunc[func(m *Matrix, factorX, factorY, factorZ float32), cc.Void]{Name: "graphene_matrix_scale"}
	graphene_matrix_skew_xy             = cc.DlFunc[func(m *Matrix, factor float32), cc.Void]{Name: "graphene_matrix_skew_xy"}
	graphene_matrix_skew_xz             = cc.DlFunc[func(m *Matrix, factor float32), cc.Void]{Name: "graphene_matrix_skew_xz"}
	graphene_matrix_skew_yz             = cc.DlFunc[func(m *Matrix, factor float32), cc.Void]{Name: "graphene_matrix_skew_yz"}
	graphene_matrix_transpose           = cc.DlFunc[func(m *Matrix, res *Matrix), cc.Void]{Name: "graphene_matrix_transpose"}
	graphene_matrix_inverse             = cc.DlFunc[func(m *Matrix, res *Matrix) bool, bool]{Name: "graphene_matrix_inverse"}
	graphene_matrix_perspective         = cc.DlFunc[func(m *Matrix, depth float32, res *Matrix), cc.Void]{Name: "graphene_matrix_perspective"}
	graphene_matrix_normalize           = cc.DlFunc[func(m *Matrix, res *Matrix), cc.Void]{Name: "graphene_matrix_normalize"}
	graphene_matrix_get_x_translation   = cc.DlFunc[func(m *Matrix) float32, float32]{Name: "graphene_matrix_get_x_translation"}
	graphene_matrix_get_y_translation   = cc.DlFunc[func(m *Matrix) float32, float32]{Name: "graphene_matrix_get_y_translation"}
	graphene_matrix_get_z_translation   = cc.DlFunc[func(m *Matrix) float32, float32]{Name: "graphene_matrix_get_z_translation"}
	graphene_matrix_get_x_scale         = cc.DlFunc[func(m *Matrix) float32, float32]{Name: "graphene_matrix_get_x_scale"}
	graphene_matrix_get_y_scale         = cc.DlFunc[func(m *Matrix) float32, float32]{Name: "graphene_matrix_get_y_scale"}
	graphene_matrix_get_z_scale         = cc.DlFunc[func(m *Matrix) float32, float32]{Name: "graphene_matrix_get_z_scale"}
	graphene_matrix_interpolate         = cc.DlFunc[func(a *Matrix, b *Matrix, factor float64, res *Matrix), cc.Void]{Name: "graphene_matrix_interpolate"}
	graphene_matrix_near                = cc.DlFunc[func(a *Matrix, b *Matrix, epsilon float32) bool, bool]{Name: "graphene_matrix_near"}
	graphene_matrix_equal               = cc.DlFunc[func(a *Matrix, b *Matrix) bool, bool]{Name: "graphene_matrix_equal"}
	graphene_matrix_equal_fast          = cc.DlFunc[func(a *Matrix, b *Matrix) bool, bool]{Name: "graphene_matrix_equal_fast"}
	graphene_matrix_print               = cc.DlFunc[func(m *Matrix), cc.Void]{Name: "graphene_matrix_print"}
	graphene_matrix_decompose           = cc.DlFunc[func(m *Matrix, translate *Vec3, scale *Vec3, rotate *Quaternion, shear *Vec3, perspective *Vec4) bool, bool]{Name: "graphene_matrix_decompose"}
	//#endregion

	//#region Matrix

	//#endregion

)
