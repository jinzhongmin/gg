package graphene

// #region Box
type Box struct{}

// #endregion

// #region Euler
type Euler struct{}

// #endregion

// #region Matrix

type Matrix struct{}

func MatrixAlloc() *Matrix { return graphene_matrix_alloc.Fn()() }
func (m *Matrix) Free()    { graphene_matrix_free.Fn()(m) }
func (m *Matrix) InitIdentity() *Matrix {
	return graphene_matrix_init_identity.Fn()(m)
}
func (m *Matrix) InitFromFloat(v [16]float32) *Matrix {
	return graphene_matrix_init_from_float.Fn()(m, &v[0])
}
func (m *Matrix) InitFromVec4(v0, v1, v2, v3 *Vec4) *Matrix {
	return graphene_matrix_init_from_vec4.Fn()(m, v0, v1, v2, v3)
}

func (m *Matrix) InitFromMatrix(src *Matrix) *Matrix {
	return graphene_matrix_init_from_matrix.Fn()(m, src)
}

func (m *Matrix) InitPerspective(fovy, aspect, zNear, zFar float32) *Matrix {
	return graphene_matrix_init_perspective.Fn()(m, fovy, aspect, zNear, zFar)
}

func (m *Matrix) InitOrtho(left, right, top, bottom, zNear, zFar float32) *Matrix {
	return graphene_matrix_init_ortho.Fn()(m, left, right, top, bottom, zNear, zFar)
}

func (m *Matrix) InitLookAt(eye, center, up *Vec3) *Matrix {
	return graphene_matrix_init_look_at.Fn()(m, eye, center, up)
}

func (m *Matrix) InitFrustum(left, right, bottom, top, zNear, zFar float32) *Matrix {
	return graphene_matrix_init_frustum.Fn()(m, left, right, bottom, top, zNear, zFar)
}
func (m *Matrix) InitScale(x, y, z float32) *Matrix {
	return graphene_matrix_init_scale.Fn()(m, x, y, z)
}

func (m *Matrix) InitTranslate(p *Point3D) *Matrix {
	return graphene_matrix_init_translate.Fn()(m, p)
}

func (m *Matrix) InitRotate(angle float32, axis *Vec3) *Matrix {
	return graphene_matrix_init_rotate.Fn()(m, angle, axis)
}

func (m *Matrix) InitSkew(xSkew, ySkew float32) *Matrix {
	return graphene_matrix_init_skew.Fn()(m, xSkew, ySkew)
}

func (m *Matrix) InitFrom2D(xx, yx, xy, yy, x0, y0 float64) *Matrix {
	return graphene_matrix_init_from_2d.Fn()(m, xx, yx, xy, yy, x0, y0)
}
func (m *Matrix) IsIdentity() bool {
	return graphene_matrix_is_identity.Fn()(m)
}

func (m *Matrix) Is2D() bool {
	return graphene_matrix_is_2d.Fn()(m)
}

func (m *Matrix) IsBackfaceVisible() bool {
	return graphene_matrix_is_backface_visible.Fn()(m)
}

func (m *Matrix) IsSingular() bool {
	return graphene_matrix_is_singular.Fn()(m)
}

func (m *Matrix) ToFloat() (v [16]float32) {
	graphene_matrix_to_float.Fn()(m, &v[0])
	return v
}

func (m *Matrix) To2D() (xx, yx, xy, yy, x0, y0 float64, success bool) {
	success = graphene_matrix_to_2d.Fn()(m, &xx, &yx, &xy, &yy, &x0, &y0)
	return
}
func (m *Matrix) GetRowTo(index uint, out *Vec4) {
	graphene_matrix_get_row.Fn()(m, index, out)
}

func (m *Matrix) GetRow(index uint) *Vec4 {
	res := Vec4Alloc()
	graphene_matrix_get_row.Fn()(m, index, res)
	return res
}

func (m *Matrix) GetValue(row, col uint) float32 {
	return graphene_matrix_get_value.Fn()(m, row, col)
}
func (m *Matrix) Multiply(b *Matrix) {
	graphene_matrix_multiply.Fn()(m, b, m)
}
func (m *Matrix) MultiplyTo(b *Matrix, out *Matrix) {
	graphene_matrix_multiply.Fn()(m, b, out)
}
func (m *Matrix) Determinant() float32 {
	return graphene_matrix_determinant.Fn()(m)
}
func (m *Matrix) TransformVec4(v *Vec4, res *Vec4) {
	graphene_matrix_transform_vec4.Fn()(m, v, res)
}
func (m *Matrix) TransformVec3(v *Vec3, res *Vec3) {
	graphene_matrix_transform_vec3.Fn()(m, v, res)
}
func (m *Matrix) TransformPoint(p *Point, res *Point) {
	graphene_matrix_transform_point.Fn()(m, p, res)
}
func (m *Matrix) TransformPoint3D(p *Point3D, res *Point3D) {
	graphene_matrix_transform_point3d.Fn()(m, p, res)
}
func (m *Matrix) TransformRect(r *Rect, res *Quad) {
	graphene_matrix_transform_rect.Fn()(m, r, res)
}

func (m *Matrix) TransformBounds(r *Rect, res *Rect) {
	graphene_matrix_transform_bounds.Fn()(m, r, res)
}

func (m *Matrix) TransformSphere(s *Sphere, res *Sphere) {
	graphene_matrix_transform_sphere.Fn()(m, s, res)
}

func (m *Matrix) TransformBox(b *Box, res *Box) {
	graphene_matrix_transform_box.Fn()(m, b, res)
}

func (m *Matrix) TransformRay(r *Ray, res *Ray) {
	graphene_matrix_transform_ray.Fn()(m, r, res)
}
func (m *Matrix) ProjectPoint(p *Point, res *Point) {
	graphene_matrix_project_point.Fn()(m, p, res)
}

func (m *Matrix) ProjectRectBounds(r *Rect, res *Rect) {
	graphene_matrix_project_rect_bounds.Fn()(m, r, res)
}

func (m *Matrix) ProjectRect(r *Rect, res *Quad) {
	graphene_matrix_project_rect.Fn()(m, r, res)
}

func (m *Matrix) UntransformPoint(p *Point, bounds *Rect, res *Point) bool {
	return graphene_matrix_untransform_point.Fn()(m, p, bounds, res)
}

func (m *Matrix) UntransformBounds(r *Rect, bounds *Rect, res *Rect) {
	graphene_matrix_untransform_bounds.Fn()(m, r, bounds, res)
}

func (m *Matrix) UnprojectPoint3D(projection *Matrix, modelview *Matrix, point *Point3D, res *Point3D) {
	graphene_matrix_unproject_point3d.Fn()(projection, modelview, point, res)
}

func (m *Matrix) Translate(pos *Point3D) {
	graphene_matrix_translate.Fn()(m, pos)
}

func (m *Matrix) Rotate(angle float32, axis *Vec3) {
	graphene_matrix_rotate.Fn()(m, angle, axis)
}

func (m *Matrix) RotateX(angle float32) {
	graphene_matrix_rotate_x.Fn()(m, angle)
}

func (m *Matrix) RotateY(angle float32) {
	graphene_matrix_rotate_y.Fn()(m, angle)
}

func (m *Matrix) RotateZ(angle float32) {
	graphene_matrix_rotate_z.Fn()(m, angle)
}

func (m *Matrix) RotateQuaternion(q *Quaternion) {
	graphene_matrix_rotate_quaternion.Fn()(m, q)
}

func (m *Matrix) RotateEuler(e *Euler) {
	graphene_matrix_rotate_euler.Fn()(m, e)
}

func (m *Matrix) Scale(factorX, factorY, factorZ float32) {
	graphene_matrix_scale.Fn()(m, factorX, factorY, factorZ)
}
func (m *Matrix) SkewXY(factor float32) {
	graphene_matrix_skew_xy.Fn()(m, factor)
}

func (m *Matrix) SkewXZ(factor float32) {
	graphene_matrix_skew_xz.Fn()(m, factor)
}

func (m *Matrix) SkewYZ(factor float32) {
	graphene_matrix_skew_yz.Fn()(m, factor)
}
func (m *Matrix) Transpose() {
	graphene_matrix_transpose.Fn()(m, m)
}
func (m *Matrix) TransposeTo(out *Matrix) {
	graphene_matrix_transpose.Fn()(m, out)
}
func (m *Matrix) Inverse(res *Matrix) bool {
	return graphene_matrix_inverse.Fn()(m, res)
}
func (m *Matrix) Perspective(depth float32) {
	graphene_matrix_perspective.Fn()(m, depth, m)
}
func (m *Matrix) PerspectiveTo(depth float32, out *Matrix) {
	graphene_matrix_perspective.Fn()(m, depth, out)
}
func (m *Matrix) Normalize() {
	graphene_matrix_normalize.Fn()(m, m)
}
func (m *Matrix) NormalizeTo(out *Matrix) {
	graphene_matrix_normalize.Fn()(m, out)
}
func (m *Matrix) GetXTranslation() float32 {
	return graphene_matrix_get_x_translation.Fn()(m)
}

func (m *Matrix) GetYTranslation() float32 {
	return graphene_matrix_get_y_translation.Fn()(m)
}

func (m *Matrix) GetZTranslation() float32 {
	return graphene_matrix_get_z_translation.Fn()(m)
}

func (m *Matrix) GetXScale() float32 {
	return graphene_matrix_get_x_scale.Fn()(m)
}

func (m *Matrix) GetYScale() float32 {
	return graphene_matrix_get_y_scale.Fn()(m)
}

func (m *Matrix) GetZScale() float32 {
	return graphene_matrix_get_z_scale.Fn()(m)
}

func (m *Matrix) Interpolate(b *Matrix, factor float64, res *Matrix) {
	graphene_matrix_interpolate.Fn()(m, b, factor, res)
}
func (m *Matrix) Near(b *Matrix, epsilon float32) bool {
	return graphene_matrix_near.Fn()(m, b, epsilon)
}

func (m *Matrix) Equal(b *Matrix) bool {
	return graphene_matrix_equal.Fn()(m, b)
}

func (m *Matrix) EqualFast(b *Matrix) bool {
	return graphene_matrix_equal_fast.Fn()(m, b)
}

func (m *Matrix) Print() {
	graphene_matrix_print.Fn()(m)
}

func (m *Matrix) Decompose(translate *Vec3, scale *Vec3, rotate *Quaternion, shear *Vec3, perspective *Vec4) bool {
	return graphene_matrix_decompose.Fn()(m, translate, scale, rotate, shear, perspective)
}

//#endregion

// #region Point

type Point struct{ X, Y float32 }

func PointAlloc() *Point                         { return graphene_point_alloc.Fn()() }
func (p *Point) Free()                           { graphene_point_free.Fn()(p) }
func (p *Point) Init(x, y float32) *Point        { return graphene_point_init.Fn()(p, x, y) }
func (p *Point) InitFromPoint(src *Point) *Point { return graphene_point_init_from_point.Fn()(p, src) }
func (p *Point) InitFromVec2(src *Vec2) *Point   { return graphene_point_init_from_vec2.Fn()(p, src) }
func (a *Point) Equal(b *Point) bool             { return graphene_point_equal.Fn()(a, b) }
func (a *Point) Distance(b *Point) (distance, dX, dY float32) {
	distance = graphene_point_distance.Fn()(a, b, &dX, &dY)
	return
}
func (a *Point) Near(b *Point, epsilon float32) bool { return graphene_point_near.Fn()(a, b, epsilon) }
func (a *Point) Interpolate(b *Point, factor float64, res *Point) {
	graphene_point_interpolate.Fn()(a, b, factor, res)
}
func (p *Point) Vec2To(out *Vec2) { graphene_point_to_vec2.Fn()(p, out) }
func (p *Point) Vec2() *Vec2      { v := Vec2Alloc(); graphene_point_to_vec2.Fn()(p, v); return v }
func PointZero() *Point           { return graphene_point_zero.Fn()() }

//#endregion

// #region Point3D

type Point3D struct{ X, Y, Z float32 }

func Point3DAlloc() *Point3D                     { return graphene_point3d_alloc.Fn()() }
func (p *Point3D) Free()                         { graphene_point3d_free.Fn()(p) }
func (p *Point3D) Init(x, y, z float32) *Point3D { return graphene_point3d_init.Fn()(p, x, y, z) }
func (p *Point3D) InitFromPoint(src *Point3D) *Point3D {
	return graphene_point3d_init_from_point.Fn()(p, src)
}
func (p *Point3D) InitFromVec3(v *Vec3) *Point3D { return graphene_point3d_init_from_vec3.Fn()(p, v) }
func (p *Point3D) ToVec3(out *Vec3)              { graphene_point3d_to_vec3.Fn()(p, out) }
func (a *Point3D) Equal(b *Point3D) bool         { return graphene_point3d_equal.Fn()(a, b) }
func (a *Point3D) Near(b *Point3D, epsilon float32) bool {
	return graphene_point3d_near.Fn()(a, b, epsilon)
}
func (p *Point3D) Scale(factor float32, res *Point3D) { graphene_point3d_scale.Fn()(p, factor, res) }
func (a *Point3D) Cross(b *Point3D, res *Point3D)     { graphene_point3d_cross.Fn()(a, b, res) }
func (a *Point3D) Dot(b *Point3D) float32             { return graphene_point3d_dot.Fn()(a, b) }
func (p *Point3D) Length() float32                    { return graphene_point3d_length.Fn()(p) }
func (p *Point3D) Normalize(res *Point3D)             { graphene_point3d_normalize.Fn()(p, res) }
func (a *Point3D) Distance(b *Point3D, delta *Vec3) float32 {
	return graphene_point3d_distance.Fn()(a, b, delta)
}
func (a *Point3D) Interpolate(b *Point3D, factor float64, res *Point3D) {
	graphene_point3d_interpolate.Fn()(a, b, factor, res)
}
func (p *Point3D) NormalizeViewport(viewport *Rect, zNear, zFar float32, res *Point3D) {
	graphene_point3d_normalize_viewport.Fn()(p, viewport, zNear, zFar, res)
}
func Point3DZero() *Point3D { return graphene_point3d_zero.Fn()() }

//#endregion

// #region Quad
type Quad struct{}

//#endregion

// #region Quaternion
type Quaternion struct{}

// #endregion

// #region Ray
type Ray struct{}

// #endregion

// #region Rect

type Rect struct {
	Origin Point
	Size   Size
}

func RectAlloc() *Rect { return graphene_rect_alloc.Fn()() }
func (r *Rect) Free()  { graphene_rect_free.Fn()(r) }
func (r *Rect) Init(x, y, width, height float32) *Rect {
	return graphene_rect_init.Fn()(r, x, y, width, height)
}
func (r *Rect) InitFromRect(src *Rect) *Rect { return graphene_rect_init_from_rect.Fn()(r, src) }
func (a *Rect) Equal(b *Rect) bool           { return graphene_rect_equal.Fn()(a, b) }
func (r *Rect) Normalize() *Rect             { return graphene_rect_normalize.Fn()(r) }
func (r *Rect) NormalizeTo(res *Rect)        { graphene_rect_normalize_r.Fn()(r, res) }
func (r *Rect) GetCenter(p *Point)           { graphene_rect_get_center.Fn()(r, p) }
func (r *Rect) GetTopLeft(p *Point)          { graphene_rect_get_top_left.Fn()(r, p) }
func (r *Rect) GetTopRight(p *Point)         { graphene_rect_get_top_right.Fn()(r, p) }
func (r *Rect) GetBottomRight(p *Point)      { graphene_rect_get_bottom_right.Fn()(r, p) }
func (r *Rect) GetBottomLeft(p *Point)       { graphene_rect_get_bottom_left.Fn()(r, p) }
func (r *Rect) GetVertices() (vertices [4]Vec2) {
	graphene_rect_get_vertices.Fn()(r, &vertices[0])
	return
}
func (r *Rect) GetX() float32            { return graphene_rect_get_x.Fn()(r) }
func (r *Rect) GetY() float32            { return graphene_rect_get_y.Fn()(r) }
func (r *Rect) GetWidth() float32        { return graphene_rect_get_width.Fn()(r) }
func (r *Rect) GetHeight() float32       { return graphene_rect_get_height.Fn()(r) }
func (r *Rect) GetArea() float32         { return graphene_rect_get_area.Fn()(r) }
func (a *Rect) Union(b *Rect, res *Rect) { graphene_rect_union.Fn()(a, b, res) }
func (a *Rect) Intersection(b *Rect) (res Rect, ok bool) {
	ok = graphene_rect_intersection.Fn()(a, b, &res)
	return
}
func (r *Rect) ContainsPoint(p *Point) bool        { return graphene_rect_contains_point.Fn()(r, p) != 0 }
func (a *Rect) ContainsRect(b *Rect) bool          { return graphene_rect_contains_rect.Fn()(a, b) }
func (r *Rect) Offset(dx, dy float32) *Rect        { return graphene_rect_offset.Fn()(r, dx, dy) }
func (r *Rect) OffsetTo(dx, dy float32, res *Rect) { graphene_rect_offset_r.Fn()(r, dx, dy, res) }
func (r *Rect) Inset(dx, dy float32) *Rect         { return graphene_rect_inset.Fn()(r, dx, dy) }
func (r *Rect) InsetTo(dx, dy float32, res *Rect)  { graphene_rect_inset_r.Fn()(r, dx, dy, res) }
func (r *Rect) RoundTo(res *Rect)                  { graphene_rect_round.Fn()(r, res) }
func (r *Rect) RoundExtentsTo(res *Rect)           { graphene_rect_round_extents.Fn()(r, res) }
func (a *Rect) Interpolate(b *Rect, factor float64, res *Rect) {
	graphene_rect_interpolate.Fn()(a, b, factor, res)
}
func (r *Rect) Expand(p *Point, res *Rect)      { graphene_rect_expand.Fn()(r, p, res) }
func RectZero() *Rect                           { return graphene_rect_zero.Fn()() }
func (r *Rect) Scale(sh, sv float32, res *Rect) { graphene_rect_scale.Fn()(r, sh, sv, res) }

// #endregion

// #region Size

type Size struct{ Width, Height float32 }

// #endregion

// #region Sphere
type Sphere struct{}

// #endregion

// #region Vec2

type Vec2 struct{}

func Vec2Alloc() *Vec2                       { return graphene_vec2_alloc.Fn()() }
func (v *Vec2) Free()                        { graphene_vec2_free.Fn()(v) }
func (v *Vec2) Init(x, y float32) *Vec2      { return graphene_vec2_init.Fn()(v, x, y) }
func (v *Vec2) InitFromVec2(src *Vec2) *Vec2 { return graphene_vec2_init_from_vec2.Fn()(v, src) }
func (v *Vec2) InitFromFloat(src [2]float32) *Vec2 {
	return graphene_vec2_init_from_float.Fn()(v, &src[0])
}
func (v *Vec2) ToFloat() (dest [2]float32) { graphene_vec2_to_float.Fn()(v, &dest[0]); return }

func (v *Vec2) Add(b *Vec2)             { graphene_vec2_add.Fn()(v, b, v) }
func (v *Vec2) Subtract(b *Vec2)        { graphene_vec2_subtract.Fn()(v, b, v) }
func (v *Vec2) Multiply(b *Vec2)        { graphene_vec2_multiply.Fn()(v, b, v) }
func (v *Vec2) Divide(b *Vec2)          { graphene_vec2_divide.Fn()(v, b, v) }
func (v *Vec2) Dot(b *Vec2) float32     { return graphene_vec2_dot.Fn()(v, b) }
func (a *Vec2) AddTo(b, out *Vec2)      { graphene_vec2_add.Fn()(a, b, out) }
func (a *Vec2) SubtractTo(b, out *Vec2) { graphene_vec2_subtract.Fn()(a, b, out) }
func (a *Vec2) MultiplyTo(b, out *Vec2) { graphene_vec2_multiply.Fn()(a, b, out) }
func (a *Vec2) DivideTo(b, out *Vec2)   { graphene_vec2_divide.Fn()(a, b, out) }

func (v *Vec2) Length() float32                      { return graphene_vec2_length.Fn()(v) }
func (v *Vec2) Normalize()                           { graphene_vec2_normalize.Fn()(v, v) }
func (v *Vec2) Scale(factor float32)                 { graphene_vec2_scale.Fn()(v, factor, v) }
func (v *Vec2) Negate()                              { graphene_vec2_negate.Fn()(v, v) }
func (v1 *Vec2) Near(v2 *Vec2, epsilon float32) bool { return graphene_vec2_near.Fn()(v1, v2, epsilon) }
func (v1 *Vec2) Equal(v2 *Vec2) bool                 { return graphene_vec2_equal.Fn()(v1, v2) }
func (v *Vec2) NormalizeTo(out *Vec2)                { graphene_vec2_normalize.Fn()(v, out) }
func (v *Vec2) ScaleTo(factor float32, out *Vec2)    { graphene_vec2_scale.Fn()(v, factor, out) }
func (v *Vec2) NegateTo(out *Vec2)                   { graphene_vec2_negate.Fn()(v, out) }

func (a *Vec2) Min(b *Vec2) { graphene_vec2_min.Fn()(a, b, a) }
func (a *Vec2) Max(b *Vec2) { graphene_vec2_max.Fn()(a, b, a) }
func (v1 *Vec2) Interpolate(v2 *Vec2, factor float64) {
	graphene_vec2_interpolate.Fn()(v1, v2, factor, v1)
}
func (a *Vec2) MinTo(b, out *Vec2) { graphene_vec2_min.Fn()(a, b, out) }
func (a *Vec2) MaxTo(b, out *Vec2) { graphene_vec2_max.Fn()(a, b, out) }
func (v1 *Vec2) InterpolateTo(v2 *Vec2, factor float64, out *Vec2) {
	graphene_vec2_interpolate.Fn()(v1, v2, factor, out)
}
func (v *Vec2) GetX() float32 { return graphene_vec2_get_x.Fn()(v) }
func (v *Vec2) GetY() float32 { return graphene_vec2_get_y.Fn()(v) }

func Vec2Zero() *Vec2  { return graphene_vec2_zero.Fn()() }
func Vec2One() *Vec2   { return graphene_vec2_one.Fn()() }
func Vec2XAxis() *Vec2 { return graphene_vec2_x_axis.Fn()() }
func Vec2YAxis() *Vec2 { return graphene_vec2_y_axis.Fn()() }

//#endregion

// #region Vec3

type Vec3 struct{}

func Vec3Alloc() *Vec3                       { return graphene_vec3_alloc.Fn()() }
func (v *Vec3) Free()                        { graphene_vec3_free.Fn()(v) }
func (v *Vec3) Init(x, y, z float32) *Vec3   { return graphene_vec3_init.Fn()(v, x, y, z) }
func (v *Vec3) InitFromVec3(src *Vec3) *Vec3 { return graphene_vec3_init_from_vec3.Fn()(v, src) }
func (v *Vec3) InitFromFloat(src [3]float32) *Vec3 {
	return graphene_vec3_init_from_float.Fn()(v, &src[0])
}

func (v *Vec3) ToFloat() (dest [3]float32) { graphene_vec3_to_float.Fn()(v, &dest[0]); return }
func (v *Vec3) Add(b *Vec3)                { graphene_vec3_add.Fn()(v, b, v) }
func (v *Vec3) Subtract(b *Vec3)           { graphene_vec3_subtract.Fn()(v, b, v) }
func (v *Vec3) Multiply(b *Vec3)           { graphene_vec3_multiply.Fn()(v, b, v) }
func (v *Vec3) Divide(b *Vec3)             { graphene_vec3_divide.Fn()(v, b, v) }
func (v *Vec3) Cross(b *Vec3)              { graphene_vec3_cross.Fn()(v, b, v) }
func (v *Vec3) Dot(b *Vec3) float32        { return graphene_vec3_dot.Fn()(v, b) }

func (a *Vec3) AddTo(b, out *Vec3)      { graphene_vec3_add.Fn()(a, b, out) }
func (a *Vec3) SubtractTo(b, out *Vec3) { graphene_vec3_subtract.Fn()(a, b, out) }
func (a *Vec3) MultiplyTo(b, out *Vec3) { graphene_vec3_multiply.Fn()(a, b, out) }
func (a *Vec3) DivideTo(b, out *Vec3)   { graphene_vec3_divide.Fn()(a, b, out) }
func (a *Vec3) CrossTo(b, out *Vec3)    { graphene_vec3_cross.Fn()(a, b, out) }

func (v *Vec3) Length() float32      { return graphene_vec3_length.Fn()(v) }
func (v *Vec3) Normalize()           { graphene_vec3_normalize.Fn()(v, v) }
func (v *Vec3) Scale(factor float32) { graphene_vec3_scale.Fn()(v, factor, v) }
func (v *Vec3) Negate()              { graphene_vec3_negate.Fn()(v, v) }

func (v *Vec3) NormalizeTo(out *Vec3)             { graphene_vec3_normalize.Fn()(v, out) }
func (v *Vec3) ScaleTo(factor float32, out *Vec3) { graphene_vec3_scale.Fn()(v, factor, out) }
func (v *Vec3) NegateTo(out *Vec3)                { graphene_vec3_negate.Fn()(v, out) }

func (v1 *Vec3) Equal(v2 *Vec3) bool                 { return graphene_vec3_equal.Fn()(v1, v2) }
func (v1 *Vec3) Near(v2 *Vec3, epsilon float32) bool { return graphene_vec3_near.Fn()(v1, v2, epsilon) }
func (a *Vec3) Min(b *Vec3)                          { graphene_vec3_min.Fn()(a, b, a) }
func (a *Vec3) Max(b *Vec3)                          { graphene_vec3_max.Fn()(a, b, a) }
func (v1 *Vec3) Interpolate(v2 *Vec3, factor float64) {
	graphene_vec3_interpolate.Fn()(v1, v2, factor, v1)
}

func (a *Vec3) MinTo(b, out *Vec3) { graphene_vec3_min.Fn()(a, b, out) }
func (a *Vec3) MaxTo(b, out *Vec3) { graphene_vec3_max.Fn()(a, b, out) }
func (v1 *Vec3) InterpolateTo(v2 *Vec3, factor float64, out *Vec3) {
	graphene_vec3_interpolate.Fn()(v1, v2, factor, out)
}

func (v *Vec3) GetX() float32       { return graphene_vec3_get_x.Fn()(v) }
func (v *Vec3) GetY() float32       { return graphene_vec3_get_y.Fn()(v) }
func (v *Vec3) GetZ() float32       { return graphene_vec3_get_z.Fn()(v) }
func (v *Vec3) GetXY() *Vec2        { v2 := Vec2Alloc(); graphene_vec3_get_xy.Fn()(v, v2); return v2 }
func (v *Vec3) GetXYTo(out *Vec2)   { graphene_vec3_get_xy.Fn()(v, out) }
func (v *Vec3) GetXY0() *Vec3       { v3 := Vec3Alloc(); graphene_vec3_get_xy0.Fn()(v, v3); return v3 }
func (v *Vec3) GetXY0ToSelf()       { graphene_vec3_get_xy0.Fn()(v, v) }
func (v *Vec3) GetXY0To(out *Vec3)  { graphene_vec3_get_xy0.Fn()(v, out) }
func (v *Vec3) GetXYZ0() *Vec4      { v4 := Vec4Alloc(); graphene_vec3_get_xyz0.Fn()(v, v4); return v4 }
func (v *Vec3) GetXYZ0To(out *Vec4) { graphene_vec3_get_xyz0.Fn()(v, out) }
func (v *Vec3) GetXYZ1() *Vec4      { v4 := Vec4Alloc(); graphene_vec3_get_xyz1.Fn()(v, v4); return v4 }
func (v *Vec3) GetXYZ1To(out *Vec4) { graphene_vec3_get_xyz1.Fn()(v, out) }
func (v *Vec3) GetXYZW(w float32) *Vec4 {
	v4 := Vec4Alloc()
	graphene_vec3_get_xyzw.Fn()(v, w, v4)
	return v4
}
func (v *Vec3) GetXYZWTo(w float32, out *Vec4) { graphene_vec3_get_xyzw.Fn()(v, w, out) }
func Vec3Zero() *Vec3                          { return graphene_vec3_zero.Fn()() }
func Vec3One() *Vec3                           { return graphene_vec3_one.Fn()() }
func Vec3XAxis() *Vec3                         { return graphene_vec3_x_axis.Fn()() }
func Vec3YAxis() *Vec3                         { return graphene_vec3_y_axis.Fn()() }
func Vec3ZAxis() *Vec3                         { return graphene_vec3_z_axis.Fn()() }

//#endregion

// #region Vec4

type Vec4 struct{}

func Vec4Alloc() *Vec4                        { return graphene_vec4_alloc.Fn()() }
func (v *Vec4) Free()                         { graphene_vec4_free.Fn()(v) }
func (v *Vec4) Init(x, y, z, w float32) *Vec4 { return graphene_vec4_init.Fn()(v, x, y, z, w) }
func (v *Vec4) InitFromVec4(src *Vec4) *Vec4  { return graphene_vec4_init_from_vec4.Fn()(v, src) }
func (v *Vec4) InitFromVec3(src *Vec3, w float32) *Vec4 {
	return graphene_vec4_init_from_vec3.Fn()(v, src, w)
}
func (v *Vec4) InitFromVec2(src *Vec2, z, w float32) *Vec4 {
	return graphene_vec4_init_from_vec2.Fn()(v, src, z, w)
}
func (v *Vec4) InitFromFloat(src [4]float32) *Vec4 {
	return graphene_vec4_init_from_float.Fn()(v, &src[0])
}
func (v *Vec4) ToFloat() (dest [4]float32) { graphene_vec4_to_float.Fn()(v, &dest[0]); return }

func (v *Vec4) Add(b *Vec4)         { graphene_vec4_add.Fn()(v, b, v) }
func (v *Vec4) Subtract(b *Vec4)    { graphene_vec4_subtract.Fn()(v, b, v) }
func (v *Vec4) Multiply(b *Vec4)    { graphene_vec4_multiply.Fn()(v, b, v) }
func (v *Vec4) Divide(b *Vec4)      { graphene_vec4_divide.Fn()(v, b, v) }
func (v *Vec4) Dot(b *Vec4) float32 { return graphene_vec4_dot.Fn()(v, b) }

func (a *Vec4) AddTo(b, out *Vec4)      { graphene_vec4_add.Fn()(a, b, out) }
func (a *Vec4) SubtractTo(b, out *Vec4) { graphene_vec4_subtract.Fn()(a, b, out) }
func (a *Vec4) MultiplyTo(b, out *Vec4) { graphene_vec4_multiply.Fn()(a, b, out) }
func (a *Vec4) DivideTo(b, out *Vec4)   { graphene_vec4_divide.Fn()(a, b, out) }

func (v *Vec4) Length() float32                      { return graphene_vec4_length.Fn()(v) }
func (v *Vec4) Normalize()                           { graphene_vec4_normalize.Fn()(v, v) }
func (v *Vec4) Scale(factor float32)                 { graphene_vec4_scale.Fn()(v, factor, v) }
func (v *Vec4) Negate()                              { graphene_vec4_negate.Fn()(v, v) }
func (v1 *Vec4) Equal(v2 *Vec4) bool                 { return graphene_vec4_equal.Fn()(v1, v2) }
func (v1 *Vec4) Near(v2 *Vec4, epsilon float32) bool { return graphene_vec4_near.Fn()(v1, v2, epsilon) }

func (v *Vec4) NormalizeTo(out *Vec4)             { graphene_vec4_normalize.Fn()(v, out) }
func (v *Vec4) ScaleTo(factor float32, out *Vec4) { graphene_vec4_scale.Fn()(v, factor, out) }
func (v *Vec4) NegateTo(out *Vec4)                { graphene_vec4_negate.Fn()(v, out) }

func (a *Vec4) Min(b *Vec4) { graphene_vec4_min.Fn()(a, b, a) }
func (a *Vec4) Max(b *Vec4) { graphene_vec4_max.Fn()(a, b, a) }
func (v1 *Vec4) Interpolate(v2 *Vec4, factor float64) {
	graphene_vec4_interpolate.Fn()(v1, v2, factor, v1)
}

func (a *Vec4) MinTo(b, out *Vec4) { graphene_vec4_min.Fn()(a, b, out) }
func (a *Vec4) MaxTo(b, out *Vec4) { graphene_vec4_max.Fn()(a, b, out) }
func (v1 *Vec4) InterpolateTo(v2 *Vec4, factor float64, out *Vec4) {
	graphene_vec4_interpolate.Fn()(v1, v2, factor, out)
}

func (v *Vec4) GetX() float32      { return graphene_vec4_get_x.Fn()(v) }
func (v *Vec4) GetY() float32      { return graphene_vec4_get_y.Fn()(v) }
func (v *Vec4) GetZ() float32      { return graphene_vec4_get_z.Fn()(v) }
func (v *Vec4) GetW() float32      { return graphene_vec4_get_w.Fn()(v) }
func (v *Vec4) GetXYTo(out *Vec2)  { graphene_vec4_get_xy.Fn()(v, out) }
func (v *Vec4) GetXY() *Vec2       { res := Vec2Alloc(); graphene_vec4_get_xy.Fn()(v, res); return res }
func (v *Vec4) GetXYZTo(out *Vec3) { graphene_vec4_get_xyz.Fn()(v, out) }
func (v *Vec4) GetXYZ() *Vec3      { res := Vec3Alloc(); graphene_vec4_get_xyz.Fn()(v, res); return res }

func Vec4Zero() *Vec4  { return graphene_vec4_zero.Fn()() }
func Vec4One() *Vec4   { return graphene_vec4_one.Fn()() }
func Vec4XAxis() *Vec4 { return graphene_vec4_x_axis.Fn()() }
func Vec4YAxis() *Vec4 { return graphene_vec4_y_axis.Fn()() }
func Vec4ZAxis() *Vec4 { return graphene_vec4_z_axis.Fn()() }
func Vec4WAxis() *Vec4 { return graphene_vec4_w_axis.Fn()() }

//#endregion
