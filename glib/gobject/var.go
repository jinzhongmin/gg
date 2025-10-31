package gobject

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/glib/glib"
)

type uptr = unsafe.Pointer
type iptr = uintptr

type ints interface {
	~int8 | ~int16 | ~int32 | ~int | ~int64 |
		~uint8 | ~uint16 | ~uint32 | ~uint | ~uint64
}

func slice[T any, N ints](p *T, n N) []T { return unsafe.Slice(p, n) }
func anyptr(a interface{}) uptr          { return (*(*[2]uptr)(uptr(&a)))[1] }
func cbool(b bool) int32 {
	if b {
		return 1
	}
	return 0
}
func carry[T any](slice []T) *T {
	if len(slice) > 0 {
		return &slice[0]
	}
	return nil
}
func carryLen[T any, E ints](slice []T) (*T, E) {
	p, l := (*T)(nil), uint64(len(slice))
	if l > 0 {
		p = &slice[0]
	}
	return p, E(l)
}

// func vcbu(fn interface{}) uptr { return cc.Cbk(fn) }

func init() { cc.Open("libgobject-2*") }

var (
	// #region Binding
	g_binding_flags_get_type      = cc.DlFunc[func() GType, GType]{Name: "g_binding_flags_get_type"}
	g_binding_get_type            = cc.DlFunc[func() GType, GType]{Name: "g_binding_get_type"}
	g_binding_get_flags           = cc.DlFunc[func(binding *GBinding) GBindingFlags, GBindingFlags]{Name: "g_binding_get_flags"}
	g_binding_get_source          = cc.DlFunc[func(binding *GBinding) *GObject, *GObject]{Name: "g_binding_get_source"}
	g_binding_dup_source          = cc.DlFunc[func(binding *GBinding) *GObject, *GObject]{Name: "g_binding_dup_source"}
	g_binding_get_target          = cc.DlFunc[func(binding *GBinding) *GObject, *GObject]{Name: "g_binding_get_target"}
	g_binding_dup_target          = cc.DlFunc[func(binding *GBinding) *GObject, *GObject]{Name: "g_binding_dup_target"}
	g_binding_get_source_property = cc.DlFunc[func(binding *GBinding) cc.String, cc.String]{Name: "g_binding_get_source_property"}
	g_binding_get_target_property = cc.DlFunc[func(binding *GBinding) cc.String, cc.String]{Name: "g_binding_get_target_property"}
	g_binding_unbind              = cc.DlFunc[func(binding *GBinding), cc.Void]{Name: "g_binding_unbind"}
	g_object_bind_property        = cc.DlFunc[func(source *GObject, source_property cc.String, target *GObject, target_property cc.String, flags GBindingFlags) *GBinding, *GBinding]{Name: "g_object_bind_property"}
	g_object_bind_property_full   = cc.DlFunc[func(source *GObject, source_property cc.String, target *GObject, target_property cc.String, flags GBindingFlags, transform_to uptr, transform_from uptr, user_data uptr, notify uptr) *GBinding, *GBinding]{Name: "g_object_bind_property_full"}
	// #endregion

	// #region BindingGroup
	g_binding_group_get_type   = cc.DlFunc[func() GType, GType]{Name: "g_binding_group_get_type"}
	g_binding_group_new        = cc.DlFunc[func() *GBindingGroup, *GBindingGroup]{Name: "g_binding_group_new"}
	g_binding_group_dup_source = cc.DlFunc[func(self *GBindingGroup) uptr, uptr]{Name: "g_binding_group_dup_source"}
	g_binding_group_set_source = cc.DlFunc[func(self *GBindingGroup, source uptr), cc.Void]{Name: "g_binding_group_set_source"}
	g_binding_group_bind       = cc.DlFunc[func(self *GBindingGroup, source_property cc.String, target uptr, target_property cc.String, flags GBindingFlags), cc.Void]{Name: "g_binding_group_bind"}
	g_binding_group_bind_full  = cc.DlFunc[func(self *GBindingGroup, source_property cc.String, target uptr, target_property cc.String, flags GBindingFlags, transform_to uptr, transform_from uptr, user_data uptr, user_data_destroy uptr), cc.Void]{Name: "g_binding_group_bind_full"}
	// #endregion

	// #region Boxed
	g_boxed_copy             = cc.DlFunc[func(boxed_type GType, src_boxed uptr) uptr, uptr]{Name: "g_boxed_copy"}
	g_boxed_free             = cc.DlFunc[func(boxed_type GType, boxed uptr), cc.Void]{Name: "g_boxed_free"}
	g_value_set_boxed        = cc.DlFunc[func(value *GValue, v_boxed uptr), cc.Void]{Name: "g_value_set_boxed"}
	g_value_set_static_boxed = cc.DlFunc[func(value *GValue, v_boxed uptr), cc.Void]{Name: "g_value_set_static_boxed"}
	g_value_take_boxed       = cc.DlFunc[func(value *GValue, v_boxed uptr), cc.Void]{Name: "g_value_take_boxed"}
	// g_value_set_boxed_take_ownership #value *GValue, v_boxed uptr)
	g_value_get_boxed            = cc.DlFunc[func(value *GValue) uptr, uptr]{Name: "g_value_get_boxed"}
	g_value_dup_boxed            = cc.DlFunc[func(value *GValue) uptr, uptr]{Name: "g_value_dup_boxed"}
	g_boxed_type_register_static = cc.DlFunc[func(name cc.String, boxed_copy uptr, boxed_free uptr) GType, GType]{Name: "g_boxed_type_register_static"}
	// #endregion

	// #region Enums
	g_enum_get_value           = cc.DlFunc[func(enum_class *GEnumClass, value int32) *GEnumValue, *GEnumValue]{Name: "g_enum_get_value"}
	g_enum_get_value_by_name   = cc.DlFunc[func(enum_class *GEnumClass, name cc.String) *GEnumValue, *GEnumValue]{Name: "g_enum_get_value_by_name"}
	g_enum_get_value_by_nick   = cc.DlFunc[func(enum_class *GEnumClass, nick cc.String) *GEnumValue, *GEnumValue]{Name: "g_enum_get_value_by_nick"}
	g_flags_get_first_value    = cc.DlFunc[func(flags_class *GFlagsClass, value uint32) *GFlagsValue, *GFlagsValue]{Name: "g_flags_get_first_value"}
	g_flags_get_value_by_name  = cc.DlFunc[func(flags_class *GFlagsClass, name cc.String) *GFlagsValue, *GFlagsValue]{Name: "g_flags_get_value_by_name"}
	g_flags_get_value_by_nick  = cc.DlFunc[func(flags_class *GFlagsClass, nick cc.String) *GFlagsValue, *GFlagsValue]{Name: "g_flags_get_value_by_nick"}
	g_enum_to_string           = cc.DlFunc[func(g_enum_type GType, value int32) cc.String, cc.String]{Name: "g_enum_to_string"}
	g_flags_to_string          = cc.DlFunc[func(flags_type GType, value uint32) cc.String, cc.String]{Name: "g_flags_to_string"}
	g_value_set_enum           = cc.DlFunc[func(value *GValue, v_enum int32), cc.Void]{Name: "g_value_set_enum"}
	g_value_get_enum           = cc.DlFunc[func(value *GValue) int32, int32]{Name: "g_value_get_enum"}
	g_value_set_flags          = cc.DlFunc[func(value *GValue, v_flags uint32), cc.Void]{Name: "g_value_set_flags"}
	g_value_get_flags          = cc.DlFunc[func(value *GValue) uint32, uint32]{Name: "g_value_get_flags"}
	g_enum_register_static     = cc.DlFunc[func(name cc.String, values *GEnumValue) GType, GType]{Name: "g_enum_register_static"}
	g_flags_register_static    = cc.DlFunc[func(name cc.String, values *GFlagsValue) GType, GType]{Name: "g_flags_register_static"}
	g_enum_complete_type_info  = cc.DlFunc[func(g_enum_type GType, info *GTypeInfo, values *GEnumValue), cc.Void]{Name: "g_enum_complete_type_info"}
	g_flags_complete_type_info = cc.DlFunc[func(g_flags_type GType, info *GTypeInfo, values *GFlagsValue), cc.Void]{Name: "g_flags_complete_type_info"}
	// #endregion

	// #region Object
	g_initially_unowned_get_type        = cc.DlFunc[func() GType, GType]{Name: "g_initially_unowned_get_type"}
	g_object_class_install_property     = cc.DlFunc[func(oclass *GObjectClass, property_id uint32, pspec *GParamSpec), cc.Void]{Name: "g_object_class_install_property"}
	g_object_class_find_property        = cc.DlFunc[func(oclass *GObjectClass, property_name cc.String) *GParamSpec, *GParamSpec]{Name: "g_object_class_find_property"}
	g_object_class_list_properties      = cc.DlFunc[func(oclass *GObjectClass, n_properties *uint32) **GParamSpec, **GParamSpec]{Name: "g_object_class_list_properties"}
	g_object_class_override_property    = cc.DlFunc[func(oclass *GObjectClass, property_id uint32, name cc.String), cc.Void]{Name: "g_object_class_override_property"}
	g_object_class_install_properties   = cc.DlFunc[func(oclass *GObjectClass, n_pspecs uint32, pspecs **GParamSpec), cc.Void]{Name: "g_object_class_install_properties"}
	g_object_interface_install_property = cc.DlFunc[func(g_iface *GTypeInterface, pspec *GParamSpec), cc.Void]{Name: "g_object_interface_install_property"}
	g_object_interface_find_property    = cc.DlFunc[func(g_iface *GTypeInterface, property_name cc.String) *GParamSpec, *GParamSpec]{Name: "g_object_interface_find_property"}
	g_object_interface_list_properties  = cc.DlFunc[func(g_iface *GTypeInterface, n_properties_p *uint32) **GParamSpec, **GParamSpec]{Name: "g_object_interface_list_properties"}
	g_object_get_type                   = cc.DlFunc[func() GType, GType]{Name: "g_object_get_type"}
	g_object_new                        = cc.DlFunc[func(t GType, args ...interface{}) *GObject, *GObject]{Name: "g_object_new", Va: true}
	g_object_new_with_properties        = cc.DlFunc[func(object_type GType, n_properties uint, names cc.Strings, values *GValue) *GObject, *GObject]{Name: "g_object_new_with_properties"}
	g_object_set_property               = cc.DlFunc[func(object *GObject, propertyName cc.String, value *GValue), cc.Void]{Name: "g_object_set_property"}
	g_object_get_property               = cc.DlFunc[func(object *GObject, propertyName cc.String, value *GValue), cc.Void]{Name: "g_object_get_property"}
	g_object_freeze_notify              = cc.DlFunc[func(object *GObject), cc.Void]{Name: "g_object_freeze_notify"}
	g_object_notify                     = cc.DlFunc[func(object *GObject, property_name cc.String), cc.Void]{Name: "g_object_notify"}
	g_object_notify_by_pspec            = cc.DlFunc[func(object *GObject, pspec *GParamSpec), cc.Void]{Name: "g_object_notify_by_pspec"}
	g_object_thaw_notify                = cc.DlFunc[func(object *GObject), cc.Void]{Name: "g_object_thaw_notify"}
	g_object_is_floating                = cc.DlFunc[func(object *GObject) int32, int32]{Name: "g_object_is_floating"}
	g_object_ref_sink                   = cc.DlFunc[func(object *GObject) uptr, uptr]{Name: "g_object_ref_sink"}
	g_object_take_ref                   = cc.DlFunc[func(object *GObject) uptr, uptr]{Name: "g_object_take_ref"}
	g_object_ref                        = cc.DlFunc[func(object *GObject) uptr, uptr]{Name: "g_object_ref"}
	g_object_unref                      = cc.DlFunc[func(object *GObject), cc.Void]{Name: "g_object_unref"}
	g_object_weak_ref                   = cc.DlFunc[func(object *GObject, notify, data uptr), cc.Void]{Name: "g_object_weak_ref"}
	g_object_weak_unref                 = cc.DlFunc[func(object *GObject, notify, data uptr), cc.Void]{Name: "g_object_weak_unref"}
	g_object_add_weak_pointer           = cc.DlFunc[func(object *GObject, weak_pointer_location *uptr), cc.Void]{Name: "g_object_add_weak_pointer"}
	g_object_remove_weak_pointer        = cc.DlFunc[func(object *GObject, weak_pointer_location *uptr), cc.Void]{Name: "g_object_remove_weak_pointer"}
	g_object_add_toggle_ref             = cc.DlFunc[func(object *GObject, notify, data uptr), cc.Void]{Name: "g_object_add_toggle_ref"}
	g_object_remove_toggle_ref          = cc.DlFunc[func(object *GObject, notify, data uptr), cc.Void]{Name: "g_object_remove_toggle_ref"}
	g_object_get_qdata                  = cc.DlFunc[func(object *GObject, quark glib.GQuark) uptr, uptr]{Name: "g_object_get_qdata"}
	g_object_set_qdata                  = cc.DlFunc[func(object *GObject, quark glib.GQuark, data uptr), cc.Void]{Name: "g_object_set_qdata"}
	g_object_set_qdata_full             = cc.DlFunc[func(object *GObject, quark glib.GQuark, data, destroy uptr), cc.Void]{Name: "g_object_set_qdata_full"}
	g_object_steal_qdata                = cc.DlFunc[func(object *GObject, quark glib.GQuark) uptr, uptr]{Name: "g_object_steal_qdata"}
	g_object_dup_qdata                  = cc.DlFunc[func(object *GObject, quark glib.GQuark, dup_func, user_data uptr) uptr, uptr]{Name: "g_object_dup_qdata"}
	// g_object_replace_qdata              #object *GObject, quark glib.GQuark, oldval, newval, destroy, old_destroy uptr) bool
	g_object_get_data      = cc.DlFunc[func(object *GObject, key cc.String) uptr, uptr]{Name: "g_object_get_data"}
	g_object_set_data      = cc.DlFunc[func(object *GObject, key cc.String, data uptr), cc.Void]{Name: "g_object_set_data"}
	g_object_set_data_full = cc.DlFunc[func(object *GObject, key cc.String, data uptr, destroy uptr), cc.Void]{Name: "g_object_set_data_full"}
	g_object_steal_data    = cc.DlFunc[func(object *GObject, key cc.String) uptr, uptr]{Name: "g_object_steal_data"}
	g_object_dup_data      = cc.DlFunc[func(object *GObject, key cc.String, dup_func, user_data uptr) uptr, uptr]{Name: "g_object_dup_data"}
	// g_object_replace_data               #object *GObject, key cc.CStr, oldval, newval, destroy uptr, old_destroy *uptr) bool
	// g_object_watch_closure              #object *GObject, closure uptr)
	// g_cclosure_new_object             #callback_func uptr, object *GObject) uptr
	// g_cclosure_new_object_swap        #callback_func uptr, object *GObject) uptr
	// g_closure_new_object              #sizeof_closure uint32, object *GObject) uptr
	g_value_set_object                = cc.DlFunc[func(value *GValue, v_object *GObject), cc.Void]{Name: "g_value_set_object"}
	g_value_get_object                = cc.DlFunc[func(value *GValue) *GObject, *GObject]{Name: "g_value_get_object"}
	g_value_dup_object                = cc.DlFunc[func(value *GValue) *GObject, *GObject]{Name: "g_value_dup_object"}
	g_signal_connect_object           = cc.DlFunc[func(instance *GObject, detailed_signal cc.String, c_handler uptr, gobject *GObject, connect_flags GConnectFlags) uint64, uint64]{Name: "g_signal_connect_object"}
	g_object_force_floating           = cc.DlFunc[func(object *GObject), cc.Void]{Name: "g_object_force_floating"}
	g_object_run_dispose              = cc.DlFunc[func(object *GObject), cc.Void]{Name: "g_object_run_dispose"}
	g_value_take_object               = cc.DlFunc[func(value *GValue, v_object *GObject), cc.Void]{Name: "g_value_take_object"}
	g_value_set_object_take_ownership = cc.DlFunc[func(value *GValue, v_object *GObject), cc.Void]{Name: "g_value_set_object_take_ownership"}
	// g_object_compat_control           #what uint64, data uptr) uint64
	// #endregion

	// #region ParamSpec
	g_param_spec_ref                 = cc.DlFunc[func(pspec *GParamSpec) *GParamSpec, *GParamSpec]{Name: "g_param_spec_ref"}
	g_param_spec_unref               = cc.DlFunc[func(pspec *GParamSpec), cc.Void]{Name: "g_param_spec_unref"}
	g_param_spec_sink                = cc.DlFunc[func(pspec *GParamSpec), cc.Void]{Name: "g_param_spec_sink"}
	g_param_spec_ref_sink            = cc.DlFunc[func(pspec *GParamSpec) *GParamSpec, *GParamSpec]{Name: "g_param_spec_ref_sink"}
	g_param_spec_get_qdata           = cc.DlFunc[func(pspec *GParamSpec, quark glib.GQuark) uptr, uptr]{Name: "g_param_spec_get_qdata"}
	g_param_spec_set_qdata           = cc.DlFunc[func(pspec *GParamSpec, quark glib.GQuark, data uptr), cc.Void]{Name: "g_param_spec_set_qdata"}
	g_param_spec_set_qdata_full      = cc.DlFunc[func(pspec *GParamSpec, quark glib.GQuark, data uptr, destroy uptr), cc.Void]{Name: "g_param_spec_set_qdata_full"}
	g_param_spec_steal_qdata         = cc.DlFunc[func(pspec *GParamSpec, quark glib.GQuark) uptr, uptr]{Name: "g_param_spec_steal_qdata"}
	g_param_spec_get_redirect_target = cc.DlFunc[func(pspec *GParamSpec) *GParamSpec, *GParamSpec]{Name: "g_param_spec_get_redirect_target"}
	g_param_value_set_default        = cc.DlFunc[func(pspec *GParamSpec, value *GValue), cc.Void]{Name: "g_param_value_set_default"}
	g_param_value_defaults           = cc.DlFunc[func(pspec *GParamSpec, value *GValue) int32, int32]{Name: "g_param_value_defaults"}
	g_param_value_validate           = cc.DlFunc[func(pspec *GParamSpec, value *GValue) int32, int32]{Name: "g_param_value_validate"}
	g_param_value_is_valid           = cc.DlFunc[func(pspec *GParamSpec, value *GValue) int32, int32]{Name: "g_param_value_is_valid"}
	g_param_value_convert            = cc.DlFunc[func(pspec *GParamSpec, src_value *GValue, dest_value *GValue, strict_validation int32) int32, int32]{Name: "g_param_value_convert"}
	g_param_values_cmp               = cc.DlFunc[func(pspec *GParamSpec, value1 *GValue, value2 *GValue) int32, int32]{Name: "g_param_values_cmp"}
	g_param_spec_get_name            = cc.DlFunc[func(pspec *GParamSpec) cc.String, cc.String]{Name: "g_param_spec_get_name"}
	g_param_spec_get_nick            = cc.DlFunc[func(pspec *GParamSpec) cc.String, cc.String]{Name: "g_param_spec_get_nick"}
	g_param_spec_get_blurb           = cc.DlFunc[func(pspec *GParamSpec) cc.String, cc.String]{Name: "g_param_spec_get_blurb"}
	g_value_set_param                = cc.DlFunc[func(value *GValue, param *GParamSpec), cc.Void]{Name: "g_value_set_param"}
	g_value_get_param                = cc.DlFunc[func(value *GValue) *GParamSpec, *GParamSpec]{Name: "g_value_get_param"}
	g_value_dup_param                = cc.DlFunc[func(value *GValue) *GParamSpec, *GParamSpec]{Name: "g_value_dup_param"}
	g_value_take_param               = cc.DlFunc[func(value *GValue, param *GParamSpec), cc.Void]{Name: "g_value_take_param"}
	g_value_set_param_take_ownership = cc.DlFunc[func(value *GValue, param *GParamSpec), cc.Void]{Name: "g_value_set_param_take_ownership"}
	g_param_spec_get_default_value   = cc.DlFunc[func(pspec *GParamSpec) *GValue, *GValue]{Name: "g_param_spec_get_default_value"}
	g_param_spec_get_name_quark      = cc.DlFunc[func(pspec *GParamSpec) glib.GQuark, glib.GQuark]{Name: "g_param_spec_get_name_quark"}
	g_param_type_register_static     = cc.DlFunc[func(name cc.String, pspec_info *GParamSpecTypeInfo) GType, GType]{Name: "g_param_type_register_static"}
	g_param_spec_is_valid_name       = cc.DlFunc[func(name cc.String) int32, int32]{Name: "g_param_spec_is_valid_name"}
	g_param_spec_char                = cc.DlFunc[func(name, nick, blurb cc.String, minimum, maximum, default_value int8, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_char"}
	g_param_spec_uchar               = cc.DlFunc[func(name, nick, blurb cc.String, minimum, maximum, default_value uint8, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_uchar"}
	g_param_spec_boolean             = cc.DlFunc[func(name, nick, blurb cc.String, default_value int32, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_boolean"}
	g_param_spec_int                 = cc.DlFunc[func(name, nick, blurb cc.String, minimum, maximum, default_value int32, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_int"}
	g_param_spec_uint                = cc.DlFunc[func(name, nick, blurb cc.String, minimum, maximum, default_value uint32, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_uint"}
	g_param_spec_long                = cc.DlFunc[func(name, nick, blurb cc.String, minimum, maximum, default_value int64, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_long"}
	g_param_spec_ulong               = cc.DlFunc[func(name, nick, blurb cc.String, minimum, maximum, default_value uint64, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_ulong"}
	g_param_spec_int64               = cc.DlFunc[func(name, nick, blurb cc.String, minimum, maximum, default_value int64, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_int64"}
	g_param_spec_uint64              = cc.DlFunc[func(name, nick, blurb cc.String, minimum, maximum, default_value uint64, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_uint64"}
	g_param_spec_unichar             = cc.DlFunc[func(name, nick, blurb cc.String, default_value uint32, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_unichar"}
	g_param_spec_enum                = cc.DlFunc[func(name, nick, blurb cc.String, enum_type GType, default_value int32, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_enum"}
	g_param_spec_flags               = cc.DlFunc[func(name, nick, blurb cc.String, flags_type GType, default_value uint32, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_flags"}
	g_param_spec_float               = cc.DlFunc[func(name, nick, blurb cc.String, minimum, maximum, default_value float32, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_float"}
	g_param_spec_double              = cc.DlFunc[func(name, nick, blurb cc.String, minimum, maximum, default_value float64, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_double"}
	g_param_spec_string              = cc.DlFunc[func(name, nick, blurb, default_value cc.String, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_string"}
	g_param_spec_param               = cc.DlFunc[func(name, nick, blurb cc.String, param_type GType, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_param"}
	g_param_spec_boxed               = cc.DlFunc[func(name, nick, blurb cc.String, boxed_type GType, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_boxed"}
	g_param_spec_pointer             = cc.DlFunc[func(name, nick, blurb cc.String, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_pointer"}
	g_param_spec_value_array         = cc.DlFunc[func(name, nick, blurb cc.String, element_spec *GParamSpec, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_value_array"}
	g_param_spec_object              = cc.DlFunc[func(name, nick, blurb cc.String, object_type GType, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_object"}
	g_param_spec_override            = cc.DlFunc[func(name cc.String, overridden *GParamSpec) *GParamSpec, *GParamSpec]{Name: "g_param_spec_override"}
	g_param_spec_gtype               = cc.DlFunc[func(name, nick, blurb cc.String, is_a_type GType, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_gtype"}
	g_param_spec_variant             = cc.DlFunc[func(name, nick, blurb cc.String, vtype uptr, default_value uptr, flags GParamFlags) *GParamSpec, *GParamSpec]{Name: "g_param_spec_variant"}
	// #endregion

	// #region Signal
	g_signal_new                           = cc.DlFunc[func(signalName cc.String, itype GType, signalFlags GSignalFlags, classOffset uint32, accumulator uptr, accuData uptr, cMarshaller iptr, returnType GType, nParams uint32, paramTypes ...GType) uint32, uint32]{Name: "g_signal_new"}
	g_signal_emit                          = cc.DlFunc[func(*GObject, uint32, glib.GQuark, ...interface{}), cc.Void]{Name: "g_signal_emit", Va: true}
	g_signal_emit_by_name                  = cc.DlFunc[func(*GObject, cc.String, ...interface{}), cc.Void]{Name: "g_signal_emit_by_name", Va: true}
	g_signal_lookup                        = cc.DlFunc[func(name cc.String, itype GType) uint32, uint32]{Name: "g_signal_lookup"}
	g_signal_name                          = cc.DlFunc[func(signalId uint32) cc.String, cc.String]{Name: "g_signal_name"}
	g_signal_query                         = cc.DlFunc[func(signalId uint32, query *GSignalQuery), cc.Void]{Name: "g_signal_query"}
	g_signal_list_ids                      = cc.DlFunc[func(itype GType, nIds *uint32) *uint32, *uint32]{Name: "g_signal_list_ids"}
	g_signal_is_valid_name                 = cc.DlFunc[func(name cc.String) int32, int32]{Name: "g_signal_is_valid_name"}
	g_signal_parse_name                    = cc.DlFunc[func(detailedSignal cc.String, itype GType, signalIdP *uint32, detailP *glib.GQuark, forceDetailQuark int32) int32, int32]{Name: "g_signal_parse_name"}
	g_signal_get_invocation_hint           = cc.DlFunc[func(instance *GObject) *GSignalInvocationHint, *GSignalInvocationHint]{Name: "g_signal_get_invocation_hint"}
	g_signal_stop_emission                 = cc.DlFunc[func(instance *GObject, signalId uint32, detail glib.GQuark), cc.Void]{Name: "g_signal_stop_emission"}
	g_signal_stop_emission_by_name         = cc.DlFunc[func(instance *GObject, detailedSignal cc.String), cc.Void]{Name: "g_signal_stop_emission_by_name"}
	g_signal_add_emission_hook             = cc.DlFunc[func(signalId uint32, detail glib.GQuark, hookFunc uptr, hookData uptr, dataDestroy uptr) uint64, uint64]{Name: "g_signal_add_emission_hook"}
	g_signal_remove_emission_hook          = cc.DlFunc[func(signalId uint32, hookId uint64), cc.Void]{Name: "g_signal_remove_emission_hook"}
	g_signal_has_handler_pending           = cc.DlFunc[func(instance *GObject, signalId uint32, detail glib.GQuark, mayBeBlocked int32) int32, int32]{Name: "g_signal_has_handler_pending"}
	g_signal_connect_data                  = cc.DlFunc[func(*GObject, cc.String, uptr, uptr, uptr, GConnectFlags) uint64, uint64]{Name: "g_signal_connect_data"}
	g_signal_handler_block                 = cc.DlFunc[func(instance *GObject, handlerId uint64), cc.Void]{Name: "g_signal_handler_block"}
	g_signal_handler_unblock               = cc.DlFunc[func(instance *GObject, handlerId uint64), cc.Void]{Name: "g_signal_handler_unblock"}
	g_signal_handler_disconnect            = cc.DlFunc[func(instance *GObject, handlerId uint64), cc.Void]{Name: "g_signal_handler_disconnect"}
	g_signal_handler_is_connected          = cc.DlFunc[func(instance *GObject, handlerId uint64) int32, int32]{Name: "g_signal_handler_is_connected"}
	g_clear_signal_handler                 = cc.DlFunc[func(handlerIdPtr *uint64, instance *GObject), cc.Void]{Name: "g_clear_signal_handler"}
	g_signal_override_class_handler        = cc.DlFunc[func(signalName cc.String, instanceType GType, classHandler uptr), cc.Void]{Name: "g_signal_override_class_handler"}
	g_signal_chain_from_overridden_handler = cc.DlFunc[func(instance *GObject, args ...interface{}), cc.Void]{Name: "g_signal_chain_from_overridden_handler", Va: true}
	// #endregion

	// #region SignalGroup
	g_signal_group_get_type        = cc.DlFunc[func() GType, GType]{Name: "g_signal_group_get_type"}
	g_signal_group_new             = cc.DlFunc[func(targetType GType) *SignalGroup, *SignalGroup]{Name: "g_signal_group_new"}
	g_signal_group_set_target      = cc.DlFunc[func(self *SignalGroup, target *GObject), cc.Void]{Name: "g_signal_group_set_target"}
	g_signal_group_dup_target      = cc.DlFunc[func(self *SignalGroup) *GObject, *GObject]{Name: "g_signal_group_dup_target"}
	g_signal_group_block           = cc.DlFunc[func(self *SignalGroup), cc.Void]{Name: "g_signal_group_block"}
	g_signal_group_unblock         = cc.DlFunc[func(self *SignalGroup), cc.Void]{Name: "g_signal_group_unblock"}
	g_signal_group_connect_object  = cc.DlFunc[func(self *SignalGroup, detailedSignal cc.String, cHandler uptr, obj *GObject, flags GConnectFlags), cc.Void]{Name: "g_signal_group_connect_object"}
	g_signal_group_connect_data    = cc.DlFunc[func(self *SignalGroup, detailedSignal cc.String, cHandler uptr, data uptr, dest uptr, flags GConnectFlags), cc.Void]{Name: "g_signal_group_connect_data"}
	g_signal_group_connect         = cc.DlFunc[func(self *SignalGroup, detailedSignal cc.String, cHandler uptr, data uptr), cc.Void]{Name: "g_signal_group_connect"}
	g_signal_group_connect_after   = cc.DlFunc[func(self *SignalGroup, detailedSignal cc.String, cHandler uptr, data uptr), cc.Void]{Name: "g_signal_group_connect_after"}
	g_signal_group_connect_swapped = cc.DlFunc[func(self *SignalGroup, detailedSignal cc.String, cHandler uptr, data uptr), cc.Void]{Name: "g_signal_group_connect_swapped"}
	// #endregion

	// #region SignalGroup
	g_type_init                                  = cc.DlFunc[func(), cc.Void]{Name: "g_type_init"}
	g_type_init_with_debug_flags                 = cc.DlFunc[func(GTypeDebugFlags), cc.Void]{Name: "g_type_init_with_debug_flags"}
	g_type_name                                  = cc.DlFunc[func(GType) cc.String, cc.String]{Name: "g_type_name"}
	g_type_qname                                 = cc.DlFunc[func(GType) glib.GQuark, glib.GQuark]{Name: "g_type_qname"}
	g_type_from_name                             = cc.DlFunc[func(cc.String) GType, GType]{Name: "g_type_from_name"}
	g_type_parent                                = cc.DlFunc[func(GType) GType, GType]{Name: "g_type_parent"}
	g_type_depth                                 = cc.DlFunc[func(GType) uint32, uint32]{Name: "g_type_depth"}
	g_type_next_base                             = cc.DlFunc[func(GType, GType) GType, GType]{Name: "g_type_next_base"}
	g_type_class_get                             = cc.DlFunc[func(GType) *GObjectClass, *GObjectClass]{Name: "g_type_class_get"}
	g_type_class_ref                             = cc.DlFunc[func(GType) *GObjectClass, *GObjectClass]{Name: "g_type_class_ref"}
	g_type_class_peek                            = cc.DlFunc[func(GType) *GObjectClass, *GObjectClass]{Name: "g_type_class_peek"}
	g_type_class_peek_static                     = cc.DlFunc[func(GType) *GObjectClass, *GObjectClass]{Name: "g_type_class_peek_static"}
	g_type_class_unref                           = cc.DlFunc[func(*GObjectClass), cc.Void]{Name: "g_type_class_unref"}
	g_type_class_peek_parent                     = cc.DlFunc[func(*GObjectClass) *GObjectClass, *GObjectClass]{Name: "g_type_class_peek_parent"}
	g_type_interface_peek                        = cc.DlFunc[func(*GObjectClass, GType) uptr, uptr]{Name: "g_type_interface_peek"}
	g_type_interface_peek_parent                 = cc.DlFunc[func(uptr) uptr, uptr]{Name: "g_type_interface_peek_parent"}
	g_type_default_interface_get                 = cc.DlFunc[func(GType) uptr, uptr]{Name: "g_type_default_interface_get"}
	g_type_default_interface_ref                 = cc.DlFunc[func(GType) uptr, uptr]{Name: "g_type_default_interface_ref"}
	g_type_default_interface_peek                = cc.DlFunc[func(GType) uptr, uptr]{Name: "g_type_default_interface_peek"}
	g_type_default_interface_unref               = cc.DlFunc[func(uptr), cc.Void]{Name: "g_type_default_interface_unref"}
	g_type_children                              = cc.DlFunc[func(GType, *int32) *GType, *GType]{Name: "g_type_children"}
	g_type_interfaces                            = cc.DlFunc[func(GType, *int32) *GType, *GType]{Name: "g_type_interfaces"}
	g_type_set_qdata                             = cc.DlFunc[func(GType, glib.GQuark, uptr), cc.Void]{Name: "g_type_set_qdata"}
	g_type_get_qdata                             = cc.DlFunc[func(GType, glib.GQuark) uptr, uptr]{Name: "g_type_get_qdata"}
	g_type_query                                 = cc.DlFunc[func(GType, *GTypeQuery), cc.Void]{Name: "g_type_query"}
	g_type_get_instance_count                    = cc.DlFunc[func(GType) int32, int32]{Name: "g_type_get_instance_count"}
	g_type_register_static                       = cc.DlFunc[func(GType, cc.String, *GTypeInfo, GTypeFlags) GType, GType]{Name: "g_type_register_static"}
	g_type_register_static_simple                = cc.DlFunc[func(GType, cc.String, uint32, uptr, uint32, uptr, GTypeFlags) GType, GType]{Name: "g_type_register_static_simple"}
	g_type_register_fundamental                  = cc.DlFunc[func(GType, cc.String, *GTypeInfo, *GTypeFundamentalInfo, GTypeFlags), cc.Void]{Name: "g_type_register_fundamental"}
	g_type_add_interface_static                  = cc.DlFunc[func(GType, GType, *GInterfaceInfo), cc.Void]{Name: "g_type_add_interface_static"}
	g_type_interface_add_prerequisite            = cc.DlFunc[func(GType, GType), cc.Void]{Name: "g_type_interface_add_prerequisite"}
	g_type_interface_instantiatable_prerequisite = cc.DlFunc[func(GType) GType, GType]{Name: "g_type_interface_instantiatable_prerequisite"}
	g_type_add_instance_private                  = cc.DlFunc[func(GType, uint64) int32, int32]{Name: "g_type_add_instance_private"}
	g_type_instance_get_private                  = cc.DlFunc[func(*GTypeInstance, GType) uptr, uptr]{Name: "g_type_instance_get_private"}
	g_type_add_class_private                     = cc.DlFunc[func(GType, uint64), cc.Void]{Name: "g_type_add_class_private"}
	g_type_class_get_private                     = cc.DlFunc[func(*GTypeClass, GType) uptr, uptr]{Name: "g_type_class_get_private"}
	g_type_ensure                                = cc.DlFunc[func(GType), cc.Void]{Name: "g_type_ensure"}
	g_type_get_type_registration_serial          = cc.DlFunc[func() uint32, uint32]{Name: "g_type_get_type_registration_serial"}
	g_type_fundamental_next                      = cc.DlFunc[func() GType, GType]{Name: "g_type_fundamental_next"}
	g_type_fundamental                           = cc.DlFunc[func(GType) GType, GType]{Name: "g_type_fundamental"}
	g_type_create_instance                       = cc.DlFunc[func(GType) *GTypeInstance, *GTypeInstance]{Name: "g_type_create_instance"}
	g_type_free_instance                         = cc.DlFunc[func(*GTypeInstance), cc.Void]{Name: "g_type_free_instance"}
	g_type_add_class_cache_func                  = cc.DlFunc[func(_, cb uptr), cc.Void]{Name: "g_type_add_class_cache_func"}
	g_type_remove_class_cache_func               = cc.DlFunc[func(_, cb iptr), cc.Void]{Name: "g_type_remove_class_cache_func"}
	g_type_name_from_instance                    = cc.DlFunc[func(*GTypeInstance) cc.String, cc.String]{Name: "g_type_name_from_instance"}
	g_type_name_from_class                       = cc.DlFunc[func(*GTypeClass) cc.String, cc.String]{Name: "g_type_name_from_class"}
	// #endregion

	// #region GValue
	g_value_init                    = cc.DlFunc[func(*GValue, GType) *GValue, *GValue]{Name: "g_value_init"}
	g_value_copy                    = cc.DlFunc[func(*GValue, *GValue), cc.Void]{Name: "g_value_copy"}
	g_value_reset                   = cc.DlFunc[func(*GValue) *GValue, *GValue]{Name: "g_value_reset"}
	g_value_unset                   = cc.DlFunc[func(*GValue), cc.Void]{Name: "g_value_unset"}
	g_value_set_instance            = cc.DlFunc[func(*GValue, *GTypeInstance), cc.Void]{Name: "g_value_set_instance"}
	g_value_init_from_instance      = cc.DlFunc[func(*GValue, *GTypeInstance), cc.Void]{Name: "g_value_init_from_instance"}
	g_value_type_compatible         = cc.DlFunc[func(GType, GType) int32, int32]{Name: "g_value_type_compatible"}
	g_value_type_transformable      = cc.DlFunc[func(GType, GType) int32, int32]{Name: "g_value_type_transformable"}
	g_value_transform               = cc.DlFunc[func(*GValue, *GValue) int32, int32]{Name: "g_value_transform"}
	g_value_register_transform_func = cc.DlFunc[func(GType, GType, uptr), cc.Void]{Name: "g_value_register_transform_func"}
	g_value_set_char                = cc.DlFunc[func(*GValue, int8), cc.Void]{Name: "g_value_set_char"}
	g_value_get_char                = cc.DlFunc[func(*GValue) int8, int8]{Name: "g_value_get_char"}
	g_value_set_schar               = cc.DlFunc[func(*GValue, int8), cc.Void]{Name: "g_value_set_schar"}
	g_value_get_schar               = cc.DlFunc[func(*GValue) int8, int8]{Name: "g_value_get_schar"}
	g_value_set_uchar               = cc.DlFunc[func(*GValue, uint8), cc.Void]{Name: "g_value_set_uchar"}
	g_value_get_uchar               = cc.DlFunc[func(*GValue) uint8, uint8]{Name: "g_value_get_uchar"}
	g_value_set_boolean             = cc.DlFunc[func(*GValue, int32), cc.Void]{Name: "g_value_set_boolean"}
	g_value_get_boolean             = cc.DlFunc[func(*GValue) int32, int32]{Name: "g_value_get_boolean"}
	g_value_set_int                 = cc.DlFunc[func(*GValue, int32), cc.Void]{Name: "g_value_set_int"}
	g_value_get_int                 = cc.DlFunc[func(*GValue) int32, int32]{Name: "g_value_get_int"}
	g_value_set_uint                = cc.DlFunc[func(*GValue, uint32), cc.Void]{Name: "g_value_set_uint"}
	g_value_get_uint                = cc.DlFunc[func(*GValue) uint32, uint32]{Name: "g_value_get_uint"}
	g_value_set_long                = cc.DlFunc[func(*GValue, int64), cc.Void]{Name: "g_value_set_long"}
	g_value_get_long                = cc.DlFunc[func(*GValue) int64, int64]{Name: "g_value_get_long"}
	g_value_set_ulong               = cc.DlFunc[func(*GValue, uint64), cc.Void]{Name: "g_value_set_ulong"}
	g_value_get_ulong               = cc.DlFunc[func(*GValue) uint64, uint64]{Name: "g_value_get_ulong"}
	g_value_set_int64               = cc.DlFunc[func(*GValue, int64), cc.Void]{Name: "g_value_set_int64"}
	g_value_get_int64               = cc.DlFunc[func(*GValue) int64, int64]{Name: "g_value_get_int64"}
	g_value_set_uint64              = cc.DlFunc[func(*GValue, uint64), cc.Void]{Name: "g_value_set_uint64"}
	g_value_get_uint64              = cc.DlFunc[func(*GValue) uint64, uint64]{Name: "g_value_get_uint64"}
	g_value_set_float               = cc.DlFunc[func(*GValue, float32), cc.Void]{Name: "g_value_set_float"}
	g_value_get_float               = cc.DlFunc[func(*GValue) float32, float32]{Name: "g_value_get_float"}
	g_value_set_double              = cc.DlFunc[func(*GValue, float64), cc.Void]{Name: "g_value_set_double"}
	g_value_get_double              = cc.DlFunc[func(*GValue) float64, float64]{Name: "g_value_get_double"}
	g_value_set_string              = cc.DlFunc[func(*GValue, cc.String), cc.Void]{Name: "g_value_set_string"}
	g_value_set_static_string       = cc.DlFunc[func(*GValue, cc.String), cc.Void]{Name: "g_value_set_static_string"}
	g_value_set_interned_string     = cc.DlFunc[func(*GValue, cc.String), cc.Void]{Name: "g_value_set_interned_string"}
	g_value_get_string              = cc.DlFunc[func(*GValue) cc.String, cc.String]{Name: "g_value_get_string"}
	g_value_dup_string              = cc.DlFunc[func(*GValue) cc.String, cc.String]{Name: "g_value_dup_string"}
	g_value_steal_string            = cc.DlFunc[func(*GValue) cc.String, cc.String]{Name: "g_value_steal_string"}
	g_value_set_pointer             = cc.DlFunc[func(*GValue, uptr), cc.Void]{Name: "g_value_set_pointer"}
	g_value_get_pointer             = cc.DlFunc[func(*GValue) uptr, uptr]{Name: "g_value_get_pointer"}
	g_gtype_get_type                = cc.DlFunc[func() GType, GType]{Name: "g_gtype_get_type"}
	g_value_set_gtype               = cc.DlFunc[func(*GValue, GType), cc.Void]{Name: "g_value_set_gtype"}
	g_value_get_gtype               = cc.DlFunc[func(*GValue) GType, GType]{Name: "g_value_get_gtype"}
	g_value_set_variant             = cc.DlFunc[func(*GValue, *glib.GVariant), cc.Void]{Name: "g_value_set_variant"}
	g_value_take_variant            = cc.DlFunc[func(*GValue, *glib.GVariant), cc.Void]{Name: "g_value_take_variant"}
	g_value_get_variant             = cc.DlFunc[func(*GValue) *glib.GVariant, *glib.GVariant]{Name: "g_value_get_variant"}
	g_value_dup_variant             = cc.DlFunc[func(*GValue) *glib.GVariant, *glib.GVariant]{Name: "g_value_dup_variant"}
	// #endregion

	// #region GTypes
	g_date_get_type                 = cc.DlFunc[func() GType, GType]{Name: "g_date_get_type"}
	g_strv_get_type                 = cc.DlFunc[func() GType, GType]{Name: "g_strv_get_type"}
	g_gstring_get_type              = cc.DlFunc[func() GType, GType]{Name: "g_gstring_get_type"}
	g_hash_table_get_type           = cc.DlFunc[func() GType, GType]{Name: "g_hash_table_get_type"}
	g_array_get_type                = cc.DlFunc[func() GType, GType]{Name: "g_array_get_type"}
	g_byte_array_get_type           = cc.DlFunc[func() GType, GType]{Name: "g_byte_array_get_type"}
	g_ptr_array_get_type            = cc.DlFunc[func() GType, GType]{Name: "g_ptr_array_get_type"}
	g_bytes_get_type                = cc.DlFunc[func() GType, GType]{Name: "g_bytes_get_type"}
	g_variant_type_get_gtype        = cc.DlFunc[func() GType, GType]{Name: "g_variant_type_get_gtype"}
	g_regex_get_type                = cc.DlFunc[func() GType, GType]{Name: "g_regex_get_type"}
	g_match_info_get_type           = cc.DlFunc[func() GType, GType]{Name: "g_match_info_get_type"}
	g_error_get_type                = cc.DlFunc[func() GType, GType]{Name: "g_error_get_type"}
	g_date_time_get_type            = cc.DlFunc[func() GType, GType]{Name: "g_date_time_get_type"}
	g_time_zone_get_type            = cc.DlFunc[func() GType, GType]{Name: "g_time_zone_get_type"}
	g_io_channel_get_type           = cc.DlFunc[func() GType, GType]{Name: "g_io_channel_get_type"}
	g_io_condition_get_type         = cc.DlFunc[func() GType, GType]{Name: "g_io_condition_get_type"}
	g_variant_builder_get_type      = cc.DlFunc[func() GType, GType]{Name: "g_variant_builder_get_type"}
	g_variant_dict_get_type         = cc.DlFunc[func() GType, GType]{Name: "g_variant_dict_get_type"}
	g_key_file_get_type             = cc.DlFunc[func() GType, GType]{Name: "g_key_file_get_type"}
	g_main_loop_get_type            = cc.DlFunc[func() GType, GType]{Name: "g_main_loop_get_type"}
	g_main_context_get_type         = cc.DlFunc[func() GType, GType]{Name: "g_main_context_get_type"}
	g_source_get_type               = cc.DlFunc[func() GType, GType]{Name: "g_source_get_type"}
	g_pollfd_get_type               = cc.DlFunc[func() GType, GType]{Name: "g_pollfd_get_type"}
	g_thread_get_type               = cc.DlFunc[func() GType, GType]{Name: "g_thread_get_type"}
	g_checksum_get_type             = cc.DlFunc[func() GType, GType]{Name: "g_checksum_get_type"}
	g_markup_parse_context_get_type = cc.DlFunc[func() GType, GType]{Name: "g_markup_parse_context_get_type"}
	g_mapped_file_get_type          = cc.DlFunc[func() GType, GType]{Name: "g_mapped_file_get_type"}
	g_option_group_get_type         = cc.DlFunc[func() GType, GType]{Name: "g_option_group_get_type"}
	g_uri_get_type                  = cc.DlFunc[func() GType, GType]{Name: "g_uri_get_type"}
	g_tree_get_type                 = cc.DlFunc[func() GType, GType]{Name: "g_tree_get_type"}
	g_pattern_spec_get_type         = cc.DlFunc[func() GType, GType]{Name: "g_pattern_spec_get_type"}
	g_bookmark_file_get_type        = cc.DlFunc[func() GType, GType]{Name: "g_bookmark_file_get_type"}
	g_hmac_get_type                 = cc.DlFunc[func() GType, GType]{Name: "g_hmac_get_type"}
	g_dir_get_type                  = cc.DlFunc[func() GType, GType]{Name: "g_dir_get_type"}
	g_rand_get_type                 = cc.DlFunc[func() GType, GType]{Name: "g_rand_get_type"}
	g_strv_builder_get_type         = cc.DlFunc[func() GType, GType]{Name: "g_strv_builder_get_type"}
	// #endregion

)
