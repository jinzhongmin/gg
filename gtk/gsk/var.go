package gsk

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/cc/dl"
	"github.com/jinzhongmin/gg/glib/glib"
	"github.com/jinzhongmin/gg/glib/gobject"
	"github.com/jinzhongmin/gg/graphene"
)

type UPtr = unsafe.Pointer
type IPtr = uintptr

// func vfn(name string, fptr interface{})               { cc.RFunc(fptr, name) }
func anyptr(a interface{}) UPtr                       { return (*(*[2]UPtr)(UPtr(&a)))[1] }
func toGobj(a interface{}) *gobject.GObject           { return (*gobject.GObject)(anyptr(a)) }
func gobjGet(a gobject.GObjectIface) *gobject.GObject { return gobject.GetGObjectIface(a) }

func init() { dl.Open("libgtk-4*") }

var (
	gsk_transform_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gsk_transform_get_type"}
	gsk_transform_ref              = cc.DlFunc[func(self *Transform) *Transform, *Transform]{Name: "gsk_transform_ref"}
	gsk_transform_unref            = cc.DlFunc[func(self *Transform), cc.Void]{Name: "gsk_transform_unref"}
	gsk_transform_print            = cc.DlFunc[func(self *Transform, str *glib.GString), cc.Void]{Name: "gsk_transform_print"}
	gsk_transform_to_string        = cc.DlFunc[func(self *Transform) cc.String, cc.String]{Name: "gsk_transform_to_string"}
	gsk_transform_parse            = cc.DlFunc[func(str cc.String, out_transform **Transform) bool, bool]{Name: "gsk_transform_parse"}
	gsk_transform_to_matrix        = cc.DlFunc[func(self *Transform, out_matrix *graphene.Matrix), cc.Void]{Name: "gsk_transform_to_matrix"}
	gsk_transform_to_2d            = cc.DlFunc[func(self *Transform, out_xx, out_yx, out_xy, out_yy, out_dx, out_dy *float32), cc.Void]{Name: "gsk_transform_to_2d"}
	gsk_transform_to_2d_components = cc.DlFunc[func(self *Transform, out_skew_x, out_skew_y, out_scale_x, out_scale_y, out_angle, out_dx, out_dy *float32), cc.Void]{Name: "gsk_transform_to_2d_components"}
	gsk_transform_to_affine        = cc.DlFunc[func(self *Transform, out_scale_x, out_scale_y, out_dx, out_dy *float32), cc.Void]{Name: "gsk_transform_to_affine"}
	gsk_transform_to_translate     = cc.DlFunc[func(self *Transform, out_dx, out_dy *float32), cc.Void]{Name: "gsk_transform_to_translate"}
	gsk_transform_get_category     = cc.DlFunc[func(self *Transform) TransformCategory, TransformCategory]{Name: "gsk_transform_get_category"}
	gsk_transform_equal            = cc.DlFunc[func(first, second *Transform) bool, bool]{Name: "gsk_transform_equal"}
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
)
