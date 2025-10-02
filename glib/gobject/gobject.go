package gobject

import (
	"reflect"
	"runtime"
	"sync"

	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/glib/glib"
)

// #region Binding

type GBinding struct{ GObject }

func GTypeGBindingFlags() GType { return g_binding_flags_get_type.Fn()() }
func GTypeGBinding() GType      { return g_binding_get_type.Fn()() }

func (b *GBinding) GetFlags() GBindingFlags   { return g_binding_get_flags.Fn()(b) }
func (b *GBinding) GetSource() *GObject       { return g_binding_get_source.Fn()(b) }
func (b *GBinding) DupSource() *GObject       { return g_binding_dup_source.Fn()(b) }
func (b *GBinding) GetTarget() *GObject       { return g_binding_get_target.Fn()(b) }
func (b *GBinding) DupTarget() *GObject       { return g_binding_dup_target.Fn()(b) }
func (b *GBinding) GetSourceProperty() string { return g_binding_get_source_property.Fn()(b).String() }
func (b *GBinding) GetTargetProperty() string { return g_binding_get_target_property.Fn()(b).String() }
func (b *GBinding) Unbind()                   { g_binding_unbind.Fn()(b) }

func (source *GObject) BindProperty(source_property string,
	target GObjectIface, target_property string, flags GBindingFlags) *GBinding {
	cSourceProp := cc.NewString(source_property)
	cTargetProp := cc.NewString(target_property)
	defer cSourceProp.Free()
	defer cTargetProp.Free()

	return g_object_bind_property.Fn()(
		GetGObjectIface(source), cSourceProp,
		GetGObjectIface(target), cTargetProp, flags)
}

func (source *GObject) BindPropertyFull(source_property string,
	target GObjectIface, target_property string, flags GBindingFlags,
	transform_to, transform_from func(binding *GBinding, fromValue, toValue *GValue, _ uptr) bool) *GBinding {
	cSourceProp := cc.NewString(source_property)
	cTargetProp := cc.NewString(target_property)
	defer cSourceProp.Free()
	defer cTargetProp.Free()
	return g_object_bind_property_full.Fn()(
		GetGObjectIface(source), cSourceProp,
		GetGObjectIface(target), cTargetProp, flags,
		vcbu(transform_to), vcbu(transform_from), nil, nil)
}

// #endregion

// #region BindingGroup

type GBindingGroup struct{ GObject }

func GTypeGBindingGroup() GType          { return g_binding_group_get_type.Fn()() }
func GBindingGroupNew() *GBindingGroup   { return g_binding_group_new.Fn()() }
func (g *GBindingGroup) DupSource() uptr { return g_binding_group_dup_source.Fn()(g) }
func (g *GBindingGroup) SetSource(source interface{}) {
	g_binding_group_set_source.Fn()(g, anyptr(source))
}
func (g *GBindingGroup) Bind(sourceProperty string, target interface{}, targetProperty string, flags GBindingFlags) {
	cSource := cc.NewString(sourceProperty)
	cTarget := cc.NewString(targetProperty)
	defer cSource.Free()
	defer cTarget.Free()
	g_binding_group_bind.Fn()(g, cSource, anyptr(target), cTarget, flags)
}
func (g *GBindingGroup) BindFull(sourceProperty string, target interface{}, targetProperty string, flags GBindingFlags,
	transformTo, transformFrom func(binding *GBinding, fromValue, toValue *GValue, _ uptr) bool) {
	cSource := cc.NewString(sourceProperty)
	cTarget := cc.NewString(targetProperty)
	defer cSource.Free()
	defer cTarget.Free()
	g_binding_group_bind_full.Fn()(g, cSource, anyptr(target), cTarget, flags, vcbu(transformTo), vcbu(transformFrom), nil, nil)
}

// #endregion

// #region Boxed

func GBoxedCopy(boxedType GType, srcBoxed interface{}) uptr {
	return g_boxed_copy.Fn()(boxedType, anyptr(srcBoxed))
}
func GBoxedFree(boxedType GType, boxed interface{}) {
	g_boxed_free.Fn()(boxedType, anyptr(boxed))
}
func (value *GValue) SetBoxed(boxed interface{}) *GValue {
	g_value_set_boxed.Fn()(value, anyptr(boxed))
	return value
}
func (value *GValue) SetStaticBoxed(boxed interface{}) *GValue {
	g_value_set_static_boxed.Fn()(value, anyptr(boxed))
	return value
}
func (value *GValue) TakeBoxed(boxed interface{}) *GValue {
	g_value_take_boxed.Fn()(value, anyptr(boxed))
	return value
}

// func (value *GValue) SetBoxedTakeOwnership(boxed interface{}) *GValue {
// 	g_value_set_boxed_take_ownership(value, anyptr(boxed))
// 	return value
// }

func (value *GValue) GetBoxed() uptr { return g_value_get_boxed.Fn()(value) }
func (value *GValue) DupBoxed() uptr { return g_value_dup_boxed.Fn()(value) }
func GBoxedTypeRegisterStatic(name string, boxedCopy func(uptr) uptr, boxedFree func(uptr)) GType {
	cName := cc.NewString(name)
	defer cName.Free()
	return g_boxed_type_register_static.Fn()(cName, vcbu(boxedCopy), vcbu(boxedFree))
}

// #endregion

// #region Enums

type GEnumClass struct {
	TypeClass GTypeClass
	Minimum   int32
	Maximum   int32
	ValuesLen uint32
	ValuesPtr uptr
}

type GFlagsClass struct {
	TypeClass GTypeClass
	Mask      uint32
	ValuesLen uint32
	ValuesPtr uptr
}

type GEnumValue struct {
	Value     int32
	ValueName cc.String
	ValueNick cc.String
}

type GFlagsValue struct {
	Value     uint32
	ValueName cc.String
	ValueNick cc.String
}

func (c *GEnumClass) GetValue(value int32) *GEnumValue {
	return g_enum_get_value.Fn()(c, value)
}
func (c *GEnumClass) GetValueByName(name string) *GEnumValue {
	cName := cc.NewString(name)
	defer cName.Free()
	return g_enum_get_value_by_name.Fn()(c, cName)
}
func (c *GEnumClass) GetValueByNick(nick string) *GEnumValue {
	cNick := cc.NewString(nick)
	defer cNick.Free()
	return g_enum_get_value_by_nick.Fn()(c, cNick)
}
func (c *GFlagsClass) GetFirstValue(value uint32) *GFlagsValue {
	return g_flags_get_first_value.Fn()(c, value)
}
func (c *GFlagsClass) GetValueByName(name string) *GFlagsValue {
	cName := cc.NewString(name)
	defer cName.Free()
	return g_flags_get_value_by_name.Fn()(c, cName)
}
func (c *GFlagsClass) GetValueByNick(nick string) *GFlagsValue {
	cNick := cc.NewString(nick)
	defer cNick.Free()
	return g_flags_get_value_by_nick.Fn()(c, cNick)
}
func EnumToString(enumType GType, value int32) string {
	return g_enum_to_string.Fn()(enumType, value).String()
}
func FlagsToString(flagsType GType, value uint32) string {
	return g_flags_to_string.Fn()(flagsType, value).String()
}
func (value *GValue) SetEnum(v_enum int32) *GValue {
	g_value_set_enum.Fn()(value, v_enum)
	return value
}
func (value *GValue) GetEnum() int32 { return g_value_get_enum.Fn()(value) }
func (value *GValue) SetFlags(v_flags uint32) *GValue {
	g_value_set_flags.Fn()(value, v_flags)
	return value
}
func (value *GValue) GetFlags() uint32 {
	return g_value_get_flags.Fn()(value)
}
func GEnumRegisterStatic(name string, values []GEnumValue) GType {
	values = append(values, GEnumValue{})
	cName := cc.NewString(name)
	defer cName.Free()
	return g_enum_register_static.Fn()(cName, carry(values))
}
func GFlagsRegisterStatic(name string, values []GFlagsValue) GType {
	values = append(values, GFlagsValue{})
	cName := cc.NewString(name)
	defer cName.Free()
	return g_flags_register_static.Fn()(cName, carry(values))
}
func GEnumCompleteTypeInfo(gEnumType GType, info *GTypeInfo, values []GEnumValue) {
	g_enum_complete_type_info.Fn()(gEnumType, info, carry(values))
}
func GFlagsCompleteTypeInfo(gFlagsType GType, info *GTypeInfo, values []GFlagsValue) {
	g_flags_complete_type_info.Fn()(gFlagsType, info, carry(values))
}

// #endregion

// #region GObject

type GObjectObj struct {
	GTypeInstance GTypeInstance
	RefCount      uint32
	QData         uptr
}

type GObjectClass struct {
	GTypeClass          GTypeClass
	ConstructProperties *glib.GSList

	Constructor               cc.Func // GObject* (*constructor) (GType type, guint n_construct_properties, GObjectConstructParam *construct_properties);
	SetProperty               cc.Func // void (*set_property) (GObject *object, guint property_id, const GValue *value, GParamSpec *pspec);
	GetProperty               cc.Func // void (*get_property) (GObject *object, guint property_id, GValue *value, GParamSpec *pspec);
	Dispose                   cc.Func // void (*dispose) (GObject *object);
	Finalize                  cc.Func // void (*finalize) (GObject *object);
	DispatchPropertiesChanged cc.Func // void (*dispatch_properties_changed) (GObject *object, guint n_pspecs, GParamSpec **pspecs);
	Notify                    cc.Func // void (*notify) (GObject *object, GParamSpec *pspec);
	Constructed               cc.Func // void (*constructed) (GObject *object);

	Flags                uint64
	NConstructProperties uint64
	Specs                uptr
	NSpecs               uint64
	_                    [3]uptr //dummy
}

type GObjectIface interface {
	GetGObjectIface() *GObject
}

func GetGObjectIface(iface GObjectIface) *GObject {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetGObjectIface()
}

func (g *GObject) GetGObjectIface() *GObject { return g }

type GObject struct{}

type GInitiallyUnowned GObject
type GInitiallyUnownedObj GObjectObj
type GInitiallyUnownedClass GObjectClass

type GObjectConstructParam struct {
	ParamSpec *GParamSpec
	Value     *GValue
}

func (o *GObject) Type() GType            { return ((*GObjectObj)(uptr(o))).GTypeInstance.GType() }
func (oclass *GObjectClass) GType() GType { return oclass.GTypeClass.GType }
func GTypeInitiallyUnowned() GType        { return g_initially_unowned_get_type.Fn()() }
func (c *GObjectClass) InstallProperty(property_id uint32, pspec *GParamSpec) {
	g_object_class_install_property.Fn()(c, property_id, pspec)
}
func (c *GObjectClass) FindProperty(property_name string) *GParamSpec {
	cName := cc.NewString(property_name)
	defer cName.Free()
	return g_object_class_find_property.Fn()(c, cName)
}
func (c *GObjectClass) ListProperties() []*GParamSpec {
	var n uint32
	ptr := g_object_class_list_properties.Fn()(c, &n)
	if ptr == nil || n == 0 {
		return nil
	}
	return *(*[]*GParamSpec)(slice(ptr, int(n)))
}
func (c *GObjectClass) OverrideProperty(property_id uint32, name string) {
	cName := cc.NewString(name)
	defer cName.Free()
	g_object_class_override_property.Fn()(c, property_id, cName)
}
func (c *GObjectClass) InstallProperties(pspecs []*GParamSpec) {
	g_object_class_install_properties.Fn()(c, uint32(len(pspecs)), carry(pspecs))
}
func GObjectInterfaceInstallProperty(g_iface uptr, pspec *GParamSpec) {
	g_object_interface_install_property.Fn()(g_iface, pspec)
}

func GObjectInterfaceFindProperty(g_iface uptr, property_name string) *GParamSpec {
	cName := cc.NewString(property_name)
	defer cName.Free()
	return g_object_interface_find_property.Fn()(g_iface, cName)
}
func GObjectInterfaceListProperties(g_iface uptr) []*GParamSpec {
	var n uint32
	ptr := g_object_interface_list_properties.Fn()(g_iface, &n)
	if ptr == nil || n == 0 {
		return nil
	}
	return *(*[]*GParamSpec)(slice(ptr, int(n)))
}

func GTypeGObject() GType { return g_object_get_type.Fn()() }
func NewGObject(t GType, args ...interface{}) *GObject {
	return g_object_new.FnVa()(t, append(args, uptr(nil))...)
}
func NewGObjectWithProperties(objectType GType, names []string, values []GValue) *GObject {
	cnames := cc.NewStrings(names)
	defer cnames.Free()
	obj := g_object_new_with_properties.Fn()(objectType, uint(len(names)), cnames, carry(values))
	return obj
}
func (object *GObject) SetProperty(propertyName string, value *GValue) {
	prop := cc.NewString(propertyName)
	defer prop.Free()
	g_object_set_property.Fn()(object, prop, value)
}
func (object *GObject) Set(propertyName string, value interface{}) {
	prop := cc.NewString(propertyName)
	defer prop.Free()
	g_object_set_property.Fn()(object, prop, NewGValue(value))
}
func (object *GObject) Get(propertyName string) interface{} {
	value := new(GValue)
	g_object_get_property.Fn()(object, propertyName, value)
	runtime.SetFinalizer(value, func(v *GValue) { v.Unset() })
	return value.Get()
}
func (object *GObject) FreezeNotify() { g_object_freeze_notify.Fn()(object) }
func (object *GObject) Notify(propertyName string) {
	cProp := cc.NewString(propertyName)
	defer cProp.Free()
	g_object_notify.Fn()(object, cProp)
}
func (object *GObject) NotifyByPspec(pspec *GParamSpec) { g_object_notify_by_pspec.Fn()(object, pspec) }
func (object *GObject) ThawNotify()                     { g_object_thaw_notify.Fn()(object) }
func (object *GObject) IsFloating() bool                { return g_object_is_floating.Fn()(anyptr(object)) }
func (object *GObject) RefSink() *GObject               { return (*GObject)(g_object_ref_sink.Fn()(anyptr(object))) }
func (object *GObject) TakeRef() *GObject               { return (*GObject)(g_object_take_ref.Fn()(anyptr(object))) }
func (object *GObject) Ref() *GObject                   { return (*GObject)(g_object_ref.Fn()(anyptr(object))) }
func (object *GObject) Unref()                          { g_object_unref.Fn()(anyptr(object)) }
func (object *GObject) WeakRef(notify func(_ uptr, obj *GObject)) {
	var fn uptr
	fn = vcbu(func(p uptr, obj *GObject) {
		if notify != nil {
			notify(p, obj)
		}
		cc.CbkCloseLate(fn)
	})

	g_object_weak_ref.Fn()(object, fn, nil)
}
func (object *GObject) WeakUnref(notify func(data uptr, obj *GObject), data interface{}) {
	g_object_weak_unref.Fn()(object, vcbu(notify), anyptr(data))
}
func (object *GObject) AddWeakPointer() (weakPtr uptr) {
	g_object_add_weak_pointer.Fn()(object, &weakPtr)
	return
}
func (object *GObject) RemoveWeakPointer(ptr uptr) { g_object_remove_weak_pointer.Fn()(object, &ptr) }

func (object *GObject) AddToggleRef(notify func(_ uptr, obj *GObject, isLastRef bool)) {
	g_object_add_toggle_ref.Fn()(object, vcbu(notify), nil)
}
func (object *GObject) RemoveToggleRef(notify func(_ uptr, obj *GObject, isLastRef bool)) {
	g_object_remove_toggle_ref.Fn()(object, vcbu(notify), nil)
}
func (object *GObject) GetQData(quark glib.GQuark) uptr {
	return g_object_get_qdata.Fn()(object, quark)
}
func (object *GObject) SetQData(quark glib.GQuark, data interface{}) {
	g_object_set_qdata.Fn()(object, quark, anyptr(data))
}
func (object *GObject) SetQDataFull(quark glib.GQuark, data interface{}, destroy func(uptr)) {
	g_object_set_qdata_full.Fn()(object, quark, anyptr(data), vcbu(destroy))
}
func (object *GObject) StealQData(quark glib.GQuark) uptr {
	return g_object_steal_qdata.Fn()(object, quark)
}
func (object *GObject) DupQData(quark glib.GQuark, dupFunc func(data, _ uptr) uptr) uptr {
	return g_object_dup_qdata.Fn()(object, quark, vcbu(dupFunc), nil)
}

// func (object *GObject) ReplaceQData(quark glib.GQuark,
// 	oldval, newval interface{}, destroy, oldDestroy func(uptr)) bool {
// 	return g_object_replace_qdata(object, quark, anyptr(oldval), anyptr(newval), vcbu(destroy), vcbu(oldDestroy))
// }

func (object *GObject) GetData(key string) uptr {
	cKey := cc.NewString(key)
	defer cKey.Free()
	return g_object_get_data.Fn()(object, cKey)
}
func (object *GObject) SetData(key string, data interface{}) {
	cKey := cc.NewString(key)
	defer cKey.Free()
	g_object_set_data.Fn()(object, cKey, anyptr(data))
}
func (object *GObject) SetDataFull(key string, data interface{}, destroy func(uptr)) {
	cKey := cc.NewString(key)
	defer cKey.Free()
	g_object_set_data_full.Fn()(object, cKey, anyptr(data), vcbu(destroy))
}
func (object *GObject) StealData(key string) uptr {
	cKey := cc.NewString(key)
	defer cKey.Free()
	return g_object_steal_data.Fn()(object, cKey)
}
func (object *GObject) DupData(key string, dupFunc func(data, _ uptr) uptr) uptr {
	cKey := cc.NewString(key)
	defer cKey.Free()
	return g_object_dup_data.Fn()(object, cKey, vcbu(dupFunc), nil)
}

//	func (object *GObject) ReplaceData(key string, oldval, newval interface{}, destroy func(uptr), oldDestroy func(uptr)) bool {
//		return g_object_replace_data(object, key, anyptr(oldval), anyptr(newval), vcbu(destroy), oldDestroy)
//	}
//
//	func (object *GObject) WatchClosure(closure uptr) {
//		g_object_watch_closure(object, closure)
//	}
//
//	func CClosureNewObject(callbackFunc interface{}, object *GObject) uptr {
//		return g_cclosure_new_object(vcbu(callbackFunc), object)
//	}
//
//	func CClosureNewObjectSwap(callbackFunc interface{}, object *GObject) uptr {
//		return g_cclosure_new_object_swap(vcbu(callbackFunc), object)
//	}
//
//	func ClosureNewObject(sizeofClosure uint32, object *GObject) uptr {
//		return g_closure_new_object(sizeofClosure, object)
//	}

func (value *GValue) SetObject(vObject GObjectIface) {
	g_value_set_object.Fn()(value, GetGObjectIface(vObject))
}
func (value *GValue) GetObject() *GObject { return g_value_get_object.Fn()(value) }
func (value *GValue) DupObject() *GObject {
	return g_value_dup_object.Fn()(value)
}
func (object *GObject) SignalConnectObject(detailedSignal string, Handler interface{},
	gobject GObjectIface, connectFlags GConnectFlags) uint64 {
	return g_signal_connect_object.Fn()(object, detailedSignal, vcbu(Handler), GetGObjectIface(gobject), connectFlags)
}
func (object *GObject) ForceFloating() { g_object_force_floating.Fn()(object) }
func (object *GObject) RunDispose()    { g_object_run_dispose.Fn()(object) }
func (value *GValue) TakeObject(vObject GObjectIface) {
	g_value_take_object.Fn()(value, GetGObjectIface(vObject))
}
func (value *GValue) SetObjectTakeOwnership(vObject GObjectIface) {
	g_value_set_object_take_ownership.Fn()(value, GetGObjectIface(vObject))
}

var lifeTimePinner = map[*GObject]any{}

func (object *GObject) Pin(goValue any) {
	lifeTimePinner[object] = goValue
	object.WeakRef(func(_ uptr, obj *GObject) {
		delete(lifeTimePinner, obj)
	})
}
func (object *GObject) GetPinned() (goValue any, ok bool) {
	goValue, ok = lifeTimePinner[object]
	return
}

// #endregion

// #region GObjectCore

type GObjectCore struct{}

func GObjectCoreGet[T any](o GObjectIface) T       { g := o.GetGObjectIface(); return *(*T)(uptr(&g)) }
func (obj *GObjectCore) GType() GType              { return obj.GetGObjectIface().Type() }
func (obj *GObjectCore) Ref()                      { obj.GetGObjectIface().Ref() }
func (obj *GObjectCore) Unref()                    { obj.GetGObjectIface().Unref() }
func (obj *GObjectCore) GetGObjectIface() *GObject { return (*GObject)(uptr(obj)) }
func (obj *GObjectCore) Set(property string, value interface{}) {
	obj.GetGObjectIface().Set(property, value)
}
func (obj *GObjectCore) Get(property string) interface{} { return obj.GetGObjectIface().Get(property) }
func (obj *GObjectCore) WeakRef(notif func(_ uptr, obj *GObject)) {
	obj.GetGObjectIface().WeakRef(notif)
}
func (obj *GObjectCore) SignalConnect(signal string, fn interface{}, data interface{}) uint64 {
	return obj.GetGObjectIface().SignalConnect(signal, fn, data)
}
func (obj *GObjectCore) Emit(signal string, args ...interface{}) {
	obj.GetGObjectIface().SignalEmit(signal, args...)
}
func (obj *GObjectCore) Pin(goValue any) { obj.GetGObjectIface().Pin(goValue) }
func (obj *GObjectCore) GetPinned() (goValue any, ok bool) {
	return obj.GetGObjectIface().GetPinned()
}

// #endregion

// #region Param

type GParamSpecObj struct {
	GTypeInstance GTypeInstance

	Name      cc.String
	Flags     GParamFlags
	ValueType GType
	OwnerType GType

	Nick     cc.String
	Blure    cc.String
	QData    uptr
	RefCount uint32
	ParamID  uint32
}

type GParamSpecClass struct {
	GtypeClass GTypeClass

	ValueType GType

	Finalize        cc.Func // void (*finalize) (GParamSpec *pspec);
	ValueSetDefault cc.Func // void (*value_set_default) (GParamSpec *pspec,GValue *value);
	ValueValidate   cc.Func // gboolean (*value_validate) (GParamSpec *pspec, GValue *value);
	ValuesCmp       cc.Func // gint (*values_cmp) (GParamSpec *pspec, const GValue *value1, const GValue *value2);
	ValueIsValid    cc.Func // gboolean (*value_is_valid) (GParamSpec *pspec, const GValue *value);

	_ [3]uptr // dummy
}

type GParamSpec struct{ GObject }

func (pspec *GParamSpec) Ref() *GParamSpec     { return g_param_spec_ref.Fn()(pspec) }
func (pspec *GParamSpec) Unref()               { g_param_spec_unref.Fn()(pspec) }
func (pspec *GParamSpec) Sink()                { g_param_spec_sink.Fn()(pspec) }
func (pspec *GParamSpec) RefSink() *GParamSpec { return g_param_spec_ref_sink.Fn()(pspec) }
func (pspec *GParamSpec) GetQData(quark glib.GQuark) uptr {
	return g_param_spec_get_qdata.Fn()(pspec, quark)
}
func (pspec *GParamSpec) SetQData(quark glib.GQuark, data interface{}) {
	g_param_spec_set_qdata.Fn()(pspec, quark, anyptr(data))
}
func (pspec *GParamSpec) SetQDataFull(quark glib.GQuark, data interface{}, destroy func(uptr)) {
	g_param_spec_set_qdata_full.Fn()(pspec, quark, anyptr(data), vcbu(destroy))
}
func (pspec *GParamSpec) StealQData(quark glib.GQuark) uptr {
	return g_param_spec_steal_qdata.Fn()(pspec, quark)
}
func (pspec *GParamSpec) GetRedirectTarget() *GParamSpec {
	return g_param_spec_get_redirect_target.Fn()(pspec)
}
func (pspec *GParamSpec) ValueSetDefault(value *GValue) { g_param_value_set_default.Fn()(pspec, value) }
func (pspec *GParamSpec) ValueDefaults(value *GValue) bool {
	return g_param_value_defaults.Fn()(pspec, value)
}
func (pspec *GParamSpec) ValueValidate(value *GValue) bool {
	return g_param_value_validate.Fn()(pspec, value)
}
func (pspec *GParamSpec) ValueIsValid(value *GValue) bool {
	return g_param_value_is_valid.Fn()(pspec, value)
}
func (pspec *GParamSpec) ValueConvert(srcValue, destValue *GValue, strictValidation bool) bool {
	return g_param_value_convert.Fn()(pspec, srcValue, destValue, strictValidation)
}
func (pspec *GParamSpec) ValuesCmp(value1, value2 *GValue) int32 {
	return g_param_values_cmp.Fn()(pspec, value1, value2)
}
func (pspec *GParamSpec) GetName() string         { return g_param_spec_get_name.Fn()(pspec).String() }
func (pspec *GParamSpec) GetNick() string         { return g_param_spec_get_nick.Fn()(pspec).String() }
func (pspec *GParamSpec) GetBlurb() string        { return g_param_spec_get_blurb.Fn()(pspec).String() }
func (value *GValue) SetParam(param *GParamSpec)  { g_value_set_param.Fn()(value, param) }
func (value *GValue) GetParam() *GParamSpec       { return g_value_get_param.Fn()(value) }
func (value *GValue) DupParam() *GParamSpec       { return g_value_dup_param.Fn()(value) }
func (value *GValue) TakeParam(param *GParamSpec) { g_value_take_param.Fn()(value, param) }

// func (value *GValue) SetParamTakeOwnership(param *GParamSpec) {
// 	g_value_set_param_take_ownership(value, param)
// }

func (pspec *GParamSpec) GetDefaultValue() *GValue  { return g_param_spec_get_default_value.Fn()(pspec) }
func (pspec *GParamSpec) GetNameQuark() glib.GQuark { return g_param_spec_get_name_quark.Fn()(pspec) }

type GParamSpecTypeInfo struct {
	InstanceSize uint16
	NumPreallocs uint16
	InstanceInit cc.Func // void (*instance_init) (GParamSpec *pspec);

	ValueType       GType
	Finalize        cc.Func // void (*finalize) (GParamSpec *pspec);
	ValueSetDefault cc.Func // void (*value_set_default) (GParamSpec *pspec, GValue *value);
	ValueValidate   cc.Func // gboolean (*value_validate) (GParamSpec *pspec, GValue *value);
	ValuesCmp       cc.Func // gint (*values_cmp) (GParamSpec *pspec, const GValue *value1, const GValue *value2);
}

func GParamTypeRegisterStatic(name string, pspecInfo *GParamSpecTypeInfo) GType {
	cName := cc.NewString(name)
	defer cName.Free()
	return g_param_type_register_static.Fn()(cName, pspecInfo)
}
func GParamSpecIsValidName(name string) bool {
	cName := cc.NewString(name)
	defer cName.Free()
	return g_param_spec_is_valid_name.Fn()(cName)
}

func GParamSpecChar(name, nick, blurb string, minimum, maximum, defaultValue int8, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_char.Fn()(cName, cNick, cBlurb, minimum, maximum, defaultValue, flags)
}
func GParamSpecUchar(name, nick, blurb string, minimum, maximum, defaultValue uint8, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_uchar.Fn()(cName, cNick, cBlurb, minimum, maximum, defaultValue, flags)
}
func GParamSpecBoolean(name, nick, blurb string, defaultValue bool, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_boolean.Fn()(cName, cNick, cBlurb, defaultValue, flags)
}
func GParamSpecInt(name, nick, blurb string, minimum, maximum, defaultValue int32, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_int.Fn()(cName, cNick, cBlurb, minimum, maximum, defaultValue, flags)
}
func GParamSpecUint(name, nick, blurb string, minimum, maximum, defaultValue uint32, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_uint.Fn()(cName, cNick, cBlurb, minimum, maximum, defaultValue, flags)
}
func GParamSpecLong(name, nick, blurb string, minimum, maximum, defaultValue int64, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_long.Fn()(cName, cNick, cBlurb, minimum, maximum, defaultValue, flags)
}
func GParamSpecUlong(name, nick, blurb string, minimum, maximum, defaultValue uint64, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_ulong.Fn()(cName, cNick, cBlurb, minimum, maximum, defaultValue, flags)
}
func GParamSpecInt64(name, nick, blurb string, minimum, maximum, defaultValue int64, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_int64.Fn()(cName, cNick, cBlurb, minimum, maximum, defaultValue, flags)
}
func GParamSpecUint64(name, nick, blurb string, minimum, maximum, defaultValue uint64, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_uint64.Fn()(cName, cNick, cBlurb, minimum, maximum, defaultValue, flags)
}
func GParamSpecUnichar(name, nick, blurb string, defaultValue uint32, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_unichar.Fn()(cName, cNick, cBlurb, defaultValue, flags)
}
func GParamSpecEnum(name, nick, blurb string, enumType GType, defaultValue int32, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_enum.Fn()(cName, cNick, cBlurb, enumType, defaultValue, flags)
}
func GParamSpecFlags(name, nick, blurb string, flagsType GType, defaultValue uint32, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_flags.Fn()(cName, cNick, cBlurb, flagsType, defaultValue, flags)
}
func GParamSpecFloat(name, nick, blurb string, minimum, maximum, defaultValue float32, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_float.Fn()(cName, cNick, cBlurb, minimum, maximum, defaultValue, flags)
}
func GParamSpecDouble(name, nick, blurb string, minimum, maximum, defaultValue float64, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_double.Fn()(cName, cNick, cBlurb, minimum, maximum, defaultValue, flags)
}
func GParamSpecString(name, nick, blurb, defaultValue string, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb, cDef := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb), cc.NewString(defaultValue)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	defer cDef.Free()
	return g_param_spec_string.Fn()(cName, cNick, cBlurb, cDef, flags)
}
func GParamSpecParam(name, nick, blurb string, paramType GType, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_param.Fn()(cName, cNick, cBlurb, paramType, flags)
}
func GParamSpecBoxed(name, nick, blurb string, boxedType GType, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_boxed.Fn()(cName, cNick, cBlurb, boxedType, flags)
}
func GParamSpecPointer(name, nick, blurb string, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_pointer.Fn()(cName, cNick, cBlurb, flags)
}
func GParamSpecValueArray(name, nick, blurb string, elementSpec *GParamSpec, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_value_array.Fn()(cName, cNick, cBlurb, elementSpec, flags)
}
func GParamSpecObject(name, nick, blurb string, objectType GType, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_object.Fn()(cName, cNick, cBlurb, objectType, flags)
}
func GParamSpecOverride(name string, overridden *GParamSpec) *GParamSpec {
	cName := cc.NewString(name)
	defer cName.Free()
	return g_param_spec_override.Fn()(cName, overridden)
}
func GParamSpecGType(name, nick, blurb string, isAType GType, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_gtype.Fn()(cName, cNick, cBlurb, isAType, flags)
}
func GParamSpecVariant(name, nick, blurb string, vtype, defaultValue uptr, flags GParamFlags) *GParamSpec {
	cName, cNick, cBlurb := cc.NewString(name), cc.NewString(nick), cc.NewString(blurb)
	defer cName.Free()
	defer cNick.Free()
	defer cBlurb.Free()
	return g_param_spec_variant.Fn()(cName, cNick, cBlurb, vtype, defaultValue, flags)
}

// #endregion

// #region Signal

type GSignalInvocationHint struct {
	SignalId uint32
	Detail   glib.GQuark
	RunType  GSignalFlags
}

type GSignalQuery struct {
	SignalID      uint32
	SignalName    cc.String
	IType         GType
	SignalFlags   GSignalFlags
	ReturnType    GType
	NumParams     uint32
	PtrParamTypes *GType
}

func (sq GSignalQuery) ParamTypes() []GType {
	if sq.NumParams == 0 {
		return nil
	}
	return *(*[]GType)(slice(uptr(sq.PtrParamTypes), int(sq.NumParams)))
}

func NewSignal(signalName string, itype GType,
	signalFlags GSignalFlags, returnType GType, paramTypes ...GType) uint32 {
	return NewSignalFull(signalName, itype, signalFlags, 0, nil, nil,
		0, returnType, paramTypes...)
}
func NewSignalFull(signalName string, itype GType, signalFlags GSignalFlags, classOffset uint32,
	accumulator func(*GValue, *GValue, uptr) bool, accuData interface{},
	cMarshaller iptr, returnType GType, paramTypes ...GType) uint32 {

	cSig := cc.NewString(signalName)
	defer cSig.Free()
	return g_signal_new.FnVa()(
		cSig,
		itype,
		signalFlags,
		classOffset,
		vcbu(accumulator),
		anyptr(accuData),
		cMarshaller,
		returnType,
		uint32(len(paramTypes)),
		paramTypes...,
	)
}
func SignalEmit(obj GObjectIface, signalId uint32, detail glib.GQuark, args ...interface{}) {
	g_signal_emit.FnVa()(GetGObjectIface(obj), signalId, detail, args...)
}
func SignalEmitByName(obj GObjectIface, signal string, args ...interface{}) {
	cSignal := cc.NewString(signal)
	defer cSignal.Free()
	g_signal_emit_by_name.FnVa()(GetGObjectIface(obj), cSignal, args...)
}
func (obj *GObject) SignalEmitID(signalId uint32, detail glib.GQuark, args ...interface{}) {
	g_signal_emit.FnVa()(obj, signalId, detail, args...)
}
func (obj *GObject) SignalEmit(signal string, args ...interface{}) {
	SignalEmitByName(obj, signal, args...)
}
func SignalLookup(name string, itype GType) (sigID uint32) {
	cName := cc.NewString(name)
	defer cName.Free()
	return g_signal_lookup.Fn()(cName, itype)
}
func SignalName(sigId uint32) string                { return g_signal_name.Fn()(sigId).String() }
func SignalQueryInfo(sigId uint32) (q GSignalQuery) { g_signal_query.Fn()(sigId, &q); return }
func SignalListIds(typ GType) []uint32 {
	var nIds uint32
	ptr := g_signal_list_ids.Fn()(typ, &nIds)
	if ptr == nil || nIds == 0 {
		return nil
	}
	return *(*[]uint32)(slice(uptr(ptr), int(nIds)))
}
func (typ GType) SignalLookup(name string) (sigID uint32) { sigID = SignalLookup(name, typ); return }
func (typ GType) SignalListIds() (ids []uint32)           { return SignalListIds(typ) }
func SignalIsValidName(name string) bool {
	cName := cc.NewString(name)
	defer cName.Free()
	return g_signal_is_valid_name.Fn()(cName)
}
func SignalParseName(detailedSignal string, itype GType, forceDetailQuark bool) (signalId uint32, detail glib.GQuark, success bool) {
	cSig := cc.NewString(detailedSignal)
	defer cSig.Free()
	success = g_signal_parse_name.Fn()(cSig, itype, &signalId, &detail, forceDetailQuark)
	return
}
func SignalGetInvocationHint(obj GObjectIface) *GSignalInvocationHint {
	return g_signal_get_invocation_hint.Fn()(GetGObjectIface(obj))
}
func SignalStopEmission(obj GObjectIface, signalId uint32, detail glib.GQuark) {
	g_signal_stop_emission.Fn()(GetGObjectIface(obj), signalId, detail)
}
func SignalStopEmissionByName(obj GObjectIface, detailedSignal string) {
	sig := cc.NewString(detailedSignal)
	defer sig.Free()
	g_signal_stop_emission_by_name.Fn()(GetGObjectIface(obj), sig)
}
func SignalAddEmissionHook(signalId uint32, detail glib.GQuark,
	hookFunc func(*GSignalInvocationHint, []GValue, uptr) bool, dataDestroy func(uptr)) (hookID uint64) {
	hf := uptr(nil)
	if anyptr(hookFunc) != nil {
		hf = vcbu(func(hint *GSignalInvocationHint, n uint32, p uptr, data uptr) bool {
			return hookFunc(hint, *(*[]GValue)(slice(p, int(n))), data)
		})
	}
	return g_signal_add_emission_hook.Fn()(
		signalId,
		detail,
		hf,
		nil,
		vcbu(dataDestroy),
	)
}
func SignalRemoveEmissionHook(signalId uint32, hookId uint64) {
	g_signal_remove_emission_hook.Fn()(signalId, hookId)
}
func (obj *GObject) SignalGetInvocationHint() *GSignalInvocationHint {
	return SignalGetInvocationHint(obj)
}
func (obj *GObject) SignalStopEmission(signalId uint32, detail glib.GQuark) {
	SignalStopEmission(obj, signalId, detail)
}
func (obj *GObject) SignalStopEmissionByName(detailedSignal string) {
	SignalStopEmissionByName(obj, detailedSignal)
}
func SignalHasHandlerPending(obj GObjectIface, signalId uint32, detail glib.GQuark, mayBeBlocked bool) bool {
	return g_signal_has_handler_pending.Fn()(GetGObjectIface(obj), signalId, detail, mayBeBlocked)
}

// type destroyData struct {
// 	handler uptr
// 	destroy func(data uptr)
// 	data    interface{}
// }

// var destroyLock sync.RWMutex
// var destroyStore = []*destroyData{}
// var destroyCallback uptr

// func init() {
// 	destroyCallback = vcbu(func(u uptr) {
// 		if u == nil {
// 			return
// 		}
// 		id := int(iptr(u))
// 		destroyLock.RLock()
// 		dd := destroyStore[id]
// 		destroyLock.RUnlock()

// 		if dd.destroy != nil {
// 			dd.destroy(anyptr(dd.data))
// 		}
// 		dd.data = nil
// 		dd.destroy = nil
// 		cc.CbkClose(dd.handler)
// 		destroyStore[id] = nil
// 	})
// }

// func SignalConnectData(obj GObjectIface, signal string,
// 	fn interface{}, data interface{}, data_destroy func(data uptr), flags GConnectFlags) uint64 {

// 	handler := vcbu(fn)

// 	sigs := cc.NewString(signal)
// 	defer sigs.Free()

// 	dd := new(destroyData)
// 	dd.handler = handler
// 	dd.destroy = data_destroy
// 	dd.data = data

// 	destroyLock.Lock()
// 	id := len(destroyStore)
// 	destroyStore = append(destroyStore, dd)
// 	destroyLock.Unlock()

// 	return g_signal_connect_data.Fn()(GetGObjectIface(obj), sigs, handler, uptr(iptr(id)), destroyCallback, flags)
// }

// func connect_custom(gobj *GObject, id uptr) {

// }

// func SignalConnectDataCfunc(obj GObjectIface, signal string, handler uptr, flags GConnectFlags) uint64 {

// 	sigs := cc.NewString(signal)
// 	defer sigs.Free()

// 	return g_signal_connect_data.Fn()(GetGObjectIface(obj), sigs, handler, nil, nil, flags)
// }

func SignalConnectData(obj GObjectIface, signal string,
	fn interface{}, data interface{}, data_destroy func(data uptr), flags GConnectFlags) uint64 {

	handler := vcbu(fn)

	var close func()
	close = func() func() { return func() { data = close; data = nil; close = nil } }()

	var destroy uptr
	destroy = vcbu(func(u uptr) {
		if data_destroy != nil {
			data_destroy(u)
		}
		cc.CbkClose(handler)
		cc.CbkCloseLate(destroy)
		close()
	})

	sigs := cc.NewString(signal)
	defer sigs.Free()

	return g_signal_connect_data.Fn()(GetGObjectIface(obj), sigs, handler, anyptr(data), destroy, flags)
}

func SignalConnectDataCfunc(obj GObjectIface, signal string,
	fn uptr, data interface{}, data_destroy func(data uptr), flags GConnectFlags) uint64 {

	var close func()
	close = func() func() { return func() { data = close; data = nil; close = nil } }()

	var destroy uptr
	destroy = vcbu(func(u uptr) {
		if data_destroy != nil {
			data_destroy(u)
		}
		cc.CbkCloseLate(destroy)
		close()
	})

	sigs := cc.NewString(signal)
	defer sigs.Free()

	return g_signal_connect_data.Fn()(GetGObjectIface(obj), sigs, fn, anyptr(data), destroy, flags)
}

func (obj *GObject) SignalHasHandlerPending(signalId uint32, detail glib.GQuark, mayBeBlocked bool) bool {
	return SignalHasHandlerPending(obj, signalId, detail, mayBeBlocked)
}
func (obj *GObject) SignalConnectData(signal string,
	fn interface{}, data interface{}, data_destroy func(data uptr), flags GConnectFlags) uint64 {
	return SignalConnectData(obj, signal, fn, data, data_destroy, flags)
}
func SignalHandlerBlock(obj GObjectIface, sigId uint64) {
	g_signal_handler_block.Fn()(GetGObjectIface(obj), sigId)
}
func SignalHandlerUnblock(obj GObjectIface, sigId uint64) {
	g_signal_handler_unblock.Fn()(GetGObjectIface(obj), sigId)
}
func SignalHandlerDisconnect(obj GObjectIface, sigId uint64) {
	g_signal_handler_disconnect.Fn()(GetGObjectIface(obj), sigId)
}
func SignalHandlerIsConnected(obj GObjectIface, sigId uint64) bool {
	return g_signal_handler_is_connected.Fn()(GetGObjectIface(obj), sigId)
}

func (instance *GObject) SignalHandlerBlock(sigId uint64) {
	SignalHandlerBlock(instance, sigId)
}
func (instance *GObject) SignalHandlerUnblock(sigId uint64) {
	SignalHandlerUnblock(instance, sigId)
}
func (instance *GObject) SignalHandlerDisconnect(sigId uint64) {
	SignalHandlerDisconnect(instance, sigId)
}
func (instance *GObject) SignalHandlerIsConnected(sigId uint64) bool {
	return SignalHandlerIsConnected(instance, sigId)
}
func ClearSignalHandler(sigId *uint64, obj GObjectIface) {
	g_clear_signal_handler.Fn()(sigId, GetGObjectIface(obj))
}
func (obj *GObject) ClearSignalHandler(sigId *uint64) {
	ClearSignalHandler(sigId, obj)
}
func SignalOverrideClassHandler(signalName string, iType GType, handlerFn interface{}) {
	cName := cc.NewString(signalName)
	defer cName.Free()
	g_signal_override_class_handler.Fn()(cName, iType, vcbu(handlerFn))
}
func (t GType) SignalOverrideClassHandler(signalName string, handlerFn interface{}) {
	cName := cc.NewString(signalName)
	defer cName.Free()
	g_signal_override_class_handler.Fn()(cName, t, vcbu(handlerFn))
}
func SignalChainFromOverriddenHandler(obj GObjectIface, args ...interface{}) {
	g_signal_chain_from_overridden_handler.FnVa()(GetGObjectIface(obj), args...)
}
func (obj *GObject) ChainFromOverriddenHandler(args ...interface{}) {
	g_signal_chain_from_overridden_handler.FnVa()(GetGObjectIface(obj), args...)
}

func SignalConnect(obj *GObject, signal string, fn interface{}, data interface{}) uint64 {
	return SignalConnectData(obj, signal, fn, data, nil, GConnectFlagDefault)
}
func (obj *GObject) SignalConnect(signal string, fn interface{}, data interface{}) uint64 {
	return SignalConnectData(obj, signal, fn, data, nil, GConnectFlagDefault)
}
func SignalConnectAfter(obj *GObject, signal string, fn interface{}, data interface{}) uint64 {
	return SignalConnectData(obj, signal, fn, data, nil, GConnectFlagAfter)
}
func (obj *GObject) SignalConnectAfter(signal string, fn interface{}, data interface{}) uint64 {
	return SignalConnectData(obj, signal, fn, data, nil, GConnectFlagAfter)
}
func SignalConnectSwapped(obj *GObject, signal string, fn interface{}, data interface{}) uint64 {
	return SignalConnectData(obj, signal, fn, data, nil, GConnectFlagSwapped)
}
func (obj *GObject) SignalConnectSwapped(signal string, fn interface{}, data interface{}) uint64 {
	return SignalConnectData(obj, signal, fn, data, nil, GConnectFlagSwapped)
}

// #endregion

// #region SignalGroup

type SignalGroup struct{ GObject }

func GTypeSignalGroup() GType                      { return g_signal_group_get_type.Fn()() }
func NewSignalGroup(targetType GType) *SignalGroup { return g_signal_group_new.Fn()(targetType) }

func (group *SignalGroup) SetTarget(target GObjectIface) {
	g_signal_group_set_target.Fn()(group, GetGObjectIface(target))
}
func (group *SignalGroup) DupTarget() *GObject {
	return g_signal_group_dup_target.Fn()(group)
}
func (group *SignalGroup) Block()   { g_signal_group_block.Fn()(group) }
func (group *SignalGroup) Unblock() { g_signal_group_unblock.Fn()(group) }
func (group *SignalGroup) ConnectGObject(detailedSignal string, obj GObjectIface, handlerFn interface{}, flags GConnectFlags) {
	cSigName := cc.NewString(detailedSignal)
	defer cSigName.Free()
	fnp := vcbu(handlerFn)
	group.WeakRef(func(fn uptr) func(_ uptr, obj *GObject) {
		return func(_ uptr, obj *GObject) {
			cc.CbkCloseLate(fn)
		}
	}(fnp))
	g_signal_group_connect_object.Fn()(group, cSigName, fnp, GetGObjectIface(obj), flags)
}
func (group *SignalGroup) ConnectData(detailedSignal string, handlerFn,
	data interface{}, notify func(_ uptr), flags GConnectFlags) {
	cSigName := cc.NewString(detailedSignal)
	defer cSigName.Free()
	c := vcbu(handlerFn)
	var d uptr
	d = vcbu(func(u uptr) {
		if notify != nil {
			notify(u)
		}
		cc.CbkCloseLate(c)
		cc.CbkCloseLate(d)
	})

	g_signal_group_connect_data.Fn()(group, cSigName, c, anyptr(data), d, flags)
}
func (group *SignalGroup) Connect(detailedSignal string, handlerFn interface{}) {
	cSigName := cc.NewString(detailedSignal)
	defer cSigName.Free()
	fnp := vcbu(handlerFn)
	group.WeakRef(func(fn uptr) func(_ uptr, obj *GObject) {
		return func(_ uptr, obj *GObject) {
			cc.CbkCloseLate(fn)
		}
	}(fnp))
	g_signal_group_connect.Fn()(group, cSigName, fnp, nil)
}
func (group *SignalGroup) ConnectAfter(detailedSignal string, handlerFn interface{}) {
	cSigName := cc.NewString(detailedSignal)
	defer cSigName.Free()
	fnp := vcbu(handlerFn)
	group.WeakRef(func(fn uptr) func(_ uptr, obj *GObject) {
		return func(_ uptr, obj *GObject) {
			cc.CbkCloseLate(fn)
		}
	}(fnp))
	g_signal_group_connect_after.Fn()(group, cSigName, fnp, nil)
}
func (group *SignalGroup) ConnectSwapped(detailedSignal string, handlerFn interface{}) {
	cSigName := cc.NewString(detailedSignal)
	defer cSigName.Free()
	fnp := vcbu(handlerFn)
	group.WeakRef(func(fn uptr) func(_ uptr, obj *GObject) {
		return func(_ uptr, obj *GObject) {
			cc.CbkCloseLate(fn)
		}
	}(fnp))
	g_signal_group_connect_swapped.Fn()(group, cSigName, fnp, nil)
}

// #endregion

// #region Type

type GType uint64

const (
	Ginvalid GType = iota << 2
	Gvoid
	Ginterface
	Gchar
	Guchar
	Gbool
	Gint32
	Guint32
	Glong
	Gulong
	Gint64
	Guint64
	Genum
	Gflags
	Gfloat
	Gdouble
	Gstring
	Gpointer
	Gboxed
	Gparam
	Gobject
	Gvariant
)

type GTypeClass struct{ GType GType }

type GTypeInstanceIface interface {
	GetGTypeInstanceIface() *GTypeInstance
}

func GetGTypeInstanceIface(iface GTypeInstanceIface) *GTypeInstance {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetGTypeInstanceIface()
}

func (ti *GTypeInstance) GetGTypeInstanceIface() *GTypeInstance { return ti }

type GTypeInstance struct{ GTypeClass *GTypeClass }
type GTypeInterface struct {
	GType         GType
	GInstanceType GType
}
type GTypeQuery struct {
	GType        GType
	GTypeName    cc.String
	ClassSize    uint32
	InstanceSize uint32
}

func (ins GTypeInstance) GType() GType { return ins.GTypeClass.GType }

func GTypeInit()                                   { g_type_init.Fn()() }
func GTypeInitWithDebugFlags(flag GTypeDebugFlags) { g_type_init_with_debug_flags.Fn()(flag) }
func (t GType) String() string                     { return g_type_name.Fn()(t).String() }
func (t GType) Name() string                       { return g_type_name.Fn()(t).String() }
func (t GType) QName() glib.GQuark                 { return g_type_qname.Fn()(t) }
func GTypeFromName(name string) GType {
	cName := cc.NewString(name)
	defer cName.Free()
	return g_type_from_name.Fn()(cName)
}
func (t GType) Parent() GType              { return g_type_parent.Fn()(t) }
func (t GType) Depth() uint32              { return g_type_depth.Fn()(t) }
func GTypeNextBase(leaf, root GType) GType { return g_type_next_base.Fn()(leaf, root) }

func (t GType) ClassGet() *GObjectClass            { return g_type_class_get.Fn()(t) }
func (t GType) ClassRef() *GObjectClass            { return g_type_class_ref.Fn()(t) }
func (t GType) ClassPeek() *GObjectClass           { return g_type_class_peek.Fn()(t) }
func (t GType) ClassPeekStatic() *GObjectClass     { return g_type_class_peek_static.Fn()(t) }
func (cs *GObjectClass) Unref()                    { g_type_class_unref.Fn()(cs) }
func (cs *GObjectClass) PeekParent() *GObjectClass { return g_type_class_peek_parent.Fn()(cs) }

func (cs *GObjectClass) InterfacePeek(t GType) uptr { return g_type_interface_peek.Fn()(cs, t) }
func GTypeInterfacePeekParent(ifacePtr interface{}) uptr {
	return g_type_interface_peek_parent.Fn()(anyptr(ifacePtr))
}
func (t GType) DefaultInterfaceGet() uptr  { return g_type_default_interface_get.Fn()(t) }
func (t GType) DefaultInterfaceRef() uptr  { return g_type_default_interface_ref.Fn()(t) }
func (t GType) DefaultInterfaceKeep() uptr { return g_type_default_interface_peek.Fn()(t) }
func DefaultInterfaceUnref(obj uptr)       { g_type_default_interface_unref.Fn()(obj) }
func (t GType) Children() []GType {
	n := int32(0)
	p := g_type_children.Fn()(t, &n)
	if p == nil || n == 0 {
		return nil
	}
	defer cc.Free(p)

	ret := make([]GType, n)
	copy(ret, *(*[]GType)(slice(p, int(n))))

	return ret
}
func (t GType) Interfaces() []GType {
	n := int32(0)
	p := g_type_interfaces.Fn()(t, &n)
	if p == nil || n == 0 {
		return nil
	}
	defer cc.Free(p)

	ret := make([]GType, n)
	copy(ret, *(*[]GType)(slice(p, int(n))))

	return ret
}
func (t GType) SetQData(quark glib.GQuark, ptr interface{}) {
	g_type_set_qdata.Fn()(t, quark, anyptr(ptr))
}
func (t GType) GetQData(quark glib.GQuark) uptr { return g_type_get_qdata.Fn()(t, quark) }
func (t GType) Query() (query GTypeQuery)       { g_type_query.Fn()(t, &query); return }
func (t GType) GetInstanceCount() int32         { return g_type_get_instance_count.Fn()(t) }

type GTypeInfo struct {
	ClassSize uint16

	BaseInit     cc.Func // void (*GBaseInitFunc) (gpointer g_class);
	BaseFinalize cc.Func // void (*GBaseFinalizeFunc) (gpointer g_class);

	ClassInit     cc.Func // void (*GClassInitFunc) (gpointer g_class, gpointer class_data);
	ClassFinalize cc.Func // void (*GClassFinalizeFunc) (gpointer g_class, gpointer class_data);
	ClassData     uptr

	InstanceSize uint16
	NumPreallocs uint16
	InstanceInit cc.Func // void (*GInstanceInitFunc) (GTypeInstance *instance, gpointer g_class);

	ValueTable *GTypeValueTable
}

type GTypeFundamentalInfo struct{ TypeFlags GTypeFundamentalFlags }
type GInterfaceInfo struct {
	InterfaceInit     cc.Func // void (*GInterfaceInitFunc) (gpointer g_iface, gpointer iface_data);
	InterfaceFinalize cc.Func // void (*GInterfaceFinalizeFunc) (gpointer g_iface, gpointer iface_data);
	InterfaceData     uptr
}
type GTypeValueTable struct {
	ValueInitFunc        cc.Func // void (* GTypeValueInitFunc) (GValue *value);
	ValueFreeFunc        cc.Func // void (* GTypeValueFreeFunc) (GValue *value);
	ValueCopyFunc        cc.Func // void (* GTypeValueCopyFunc) (const GValue *src_value, GValue *dest_value);
	ValuePeekPointerFunc cc.Func // gpointer (* GTypeValuePeekPointerFunc) (const GValue *value);

	CollectFormat    cc.String
	ValueCollectFunc cc.Func // gchar * (* GTypeValueCollectFunc) (GValue *value, guint n_collect_values, GTypeCValue *collect_values, guint collect_flags);

	LcopyFormat    cc.String
	ValueLCopyFunc cc.Func // gchar * (* GTypeValueLCopyFunc) (const GValue *value, guint n_collect_values, GTypeCValue *collect_values, guint collect_flags);
}

func GTypeRegisterStatic(parent GType, typeName string, info *GTypeInfo, flags GTypeFlags) GType {
	cTName := cc.NewString(typeName)
	defer cTName.Free()
	return g_type_register_static.Fn()(parent, cTName, info, flags)
}
func GTypeRegisterStaticSimple(parent GType, typeName string,
	classSize uint32, classInit func(GClass, _ uptr),
	instaceSize uint32, instaceInit func(Instace, GClass uptr), flags GTypeFlags) GType {
	cTName := cc.NewString(typeName)
	defer cTName.Free()
	return g_type_register_static_simple.Fn()(parent, cTName,
		classSize, vcbu(classInit), instaceSize, vcbu(instaceInit), flags)
}
func GTypeRegisterFundamental(gtype GType, typeName string,
	info *GTypeInfo, finfo *GTypeFundamentalInfo, flags GTypeFlags) {
	cTName := cc.NewString(typeName)
	defer cTName.Free()
	g_type_register_fundamental.Fn()(gtype, cTName, info, finfo, flags)
}
func (gtype GType) AddInterfaceStatic(iface GType, info *GInterfaceInfo) {
	g_type_add_interface_static.Fn()(gtype, iface, info)
}
func (interfaceType GType) InterfaceAddPrerequisite(prerequisiteType GType) {
	g_type_interface_add_prerequisite.Fn()(interfaceType, prerequisiteType)
}
func (interfaceType GType) InterfaceInstantiatablePrerequisite() GType {
	return g_type_interface_instantiatable_prerequisite.Fn()(interfaceType)
}
func (classType GType) AddInstancePrivate(privateSize uint32) int32 {
	return g_type_add_instance_private.Fn()(classType, uint64(privateSize))
}
func (instance *GTypeInstance) GetPrivate(privateType GType) uptr {
	return g_type_instance_get_private.Fn()(instance, privateType)
}
func (classType GType) AddClassPrivate(privateSize uint32) {
	g_type_add_class_private.Fn()(classType, uint64(privateSize))
}
func (kclass *GTypeClass) GetPrivate(privateType GType) uptr {
	return g_type_class_get_private.Fn()(kclass, privateType)
}
func (t GType) Ensure()                            { g_type_ensure.Fn()(t) }
func GTypeGetTypeRegistrationSerial()              { g_type_get_type_registration_serial.Fn()() }
func GTypeFundamentalNext() GType                  { return g_type_fundamental_next.Fn()() }
func GTypeFundamental(t GType) GType               { return g_type_fundamental.Fn()(t) }
func (t GType) Fundamental() GType                 { return g_type_fundamental.Fn()(t) }
func GTypeCreateInstance(typ GType) *GTypeInstance { return g_type_create_instance.Fn()(typ) }
func (t GType) CreateInstance() *GTypeInstance     { return g_type_create_instance.Fn()(t) }
func GTypeFreeInstance(inc GTypeInstanceIface)     { g_type_free_instance.Fn()(GetGTypeInstanceIface(inc)) }
func (inc *GTypeInstance) Free()                   { g_type_free_instance.Fn()(inc) }
func GTypeAddClassCacheFunc(fn func(_ uptr, gClass *GTypeClass) bool) (funcID uintptr) {
	p := vcbu(fn)
	g_type_add_class_cache_func.Fn()(nil, p)
	return uintptr(p)
}
func GTypeRemoveClassCacheFunc(funcID iptr)  { g_type_remove_class_cache_func.Fn()(0, funcID) }
func (inc *GTypeInstance) GTypeName() string { return g_type_name_from_instance.Fn()(inc).String() }
func GTypeNameFromInstance(inc GTypeInstanceIface) string {
	return g_type_name_from_instance.Fn()(GetGTypeInstanceIface(inc)).String()
}
func GTypeNameFromClass(gclass *GTypeClass) string {
	return g_type_name_from_class.Fn()(gclass).String()
}
func (gclass *GTypeClass) GTypeName() string { return g_type_name_from_class.Fn()(gclass).String() }

// #endregion

// #region GValue

type GValue struct {
	GType GType
	Data  [2]uptr
}

func InitGValue(typ GType) (val *GValue) {
	v := GValue{}
	g_value_init.Fn()(&v, typ)
	return &v
}
func NewGValue(a interface{}) *GValue {
	var gValue *GValue

	switch v := a.(type) {
	case int8:
		gValue = InitGValue(Gchar)
		gValue.SetChar(v)
	case uint8:
		gValue = InitGValue(Guchar)
		gValue.SetUChar(v)
	case bool:
		gValue = InitGValue(Gbool)
		gValue.SetBoolean(v)
	case int32:
		gValue = InitGValue(Gint32)
		gValue.SetInt(v)
	case uint32:
		gValue = InitGValue(Guint32)
		gValue.SetUInt(v)
	case int64:
		gValue = InitGValue(Gint64)
		gValue.SetInt64(v)
	case uint64:
		gValue = InitGValue(Guint64)
		gValue.SetUInt64(v)
	case float32:
		gValue = InitGValue(Gfloat)
		gValue.SetFloat(v)
	case float64:
		gValue = InitGValue(Gdouble)
		gValue.SetDouble(v)
	case string:
		gValue = InitGValue(Gstring)
		gValue.SetString(v)
	// case GType:
	// 	gValue = NewGValue(Gparam)
	// 	gValue.SetGType(v)
	case *glib.GVariant:
		gValue = InitGValue(Gvariant)
		gValue.SetVariant(v)
	default:
		// 使用反射判断是否为指针类型
		rv := reflect.ValueOf(a)
		if rv.Kind() == reflect.Ptr ||
			rv.Kind() == reflect.Pointer ||
			rv.Kind() == reflect.UnsafePointer ||
			rv.Kind() == reflect.Uintptr {
			gValue = InitGValue(Gpointer)
			gValue.SetPointer(anyptr(a))
		} else {
			panic("unsupported type for GValue")
		}
	}
	runtime.SetFinalizer(gValue, func(v *GValue) { v.Unset() })
	return gValue
}
func (value *GValue) Get() interface{} {
	switch value.GType {
	case Gchar:
		return value.GetChar()
	case Guchar:
		return value.GetUChar()
	case Gbool:
		return value.GetBoolean()
	case Gint32:
		return value.GetInt()
	case Guint32:
		return value.GetUInt()
	case Glong:
		return value.GetLong()
	case Gulong:
		return value.GetULong()
	case Gint64:
		return value.GetInt64()
	case Guint64:
		return value.GetUInt64()
	case Gfloat:
		return value.GetFloat()
	case Gdouble:
		return value.GetDouble()
	case Gstring:
		return value.GetString()
	case Gpointer:
		return value.GetPointer()
	case Gvariant:
		return value.GetVariant()
	default:
		panic("unsupported GValue type")
	}
}
func (value *GValue) Set(v interface{}) {
	switch val := v.(type) {
	case int8:
		value.SetChar(val)
	case uint8:
		value.SetUChar(val)
	case bool:
		value.SetBoolean(val)
	case int32:
		value.SetInt(val)
	case uint32:
		value.SetUInt(val)
	case int64:
		value.SetInt64(val)
	case uint64:
		value.SetUInt64(val)
	case float32:
		value.SetFloat(val)
	case float64:
		value.SetDouble(val)
	case string:
		value.SetString(val)
	case *glib.GVariant:
		value.SetVariant(val)
	case uptr:
		value.SetPointer(val)
	default:
		// 使用反射判断是否为指针类型
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Ptr ||
			rv.Kind() == reflect.Pointer ||
			rv.Kind() == reflect.UnsafePointer ||
			rv.Kind() == reflect.Uintptr {
			value.SetPointer(anyptr(v))
		} else {
			panic("unsupported type for GValue.Set")
		}
	}
}
func (value *GValue) Init(gType GType) *GValue  { return g_value_init.Fn()(value, gType) }
func (value *GValue) Copy(dest *GValue) *GValue { g_value_copy.Fn()(value, dest); return value }
func (value *GValue) Reset() *GValue            { return g_value_reset.Fn()(value) }
func (value *GValue) Unset() *GValue            { g_value_unset.Fn()(value); return value }
func (value *GValue) SetInstance(instance GTypeInstanceIface) *GValue {
	g_value_set_instance.Fn()(value, GetGTypeInstanceIface(instance))
	return value
}
func (value *GValue) InitFromInstance(instance GTypeInstanceIface) *GValue {
	g_value_init_from_instance.Fn()(value, GetGTypeInstanceIface(instance))
	return value
}
func InitGValueFromInstance(instance GTypeInstanceIface) *GValue {
	val := GValue{}
	g_value_init_from_instance.Fn()(&val, GetGTypeInstanceIface(instance))
	return &val
}
func (value *GValue) Compatible(dest *GValue) bool {
	return g_value_type_compatible.Fn()(value.GType, dest.GType)
}
func (value *GValue) Transformable(dest *GValue) bool {
	return g_value_type_transformable.Fn()(value.GType, dest.GType)
}
func (value *GValue) Transform(dest *GValue) bool {
	return g_value_transform.Fn()(value, dest)
}
func GValueRegisterTransformFunc(srcType, destType GType,
	transformFunc func(src, dst *GValue)) {
	g_value_register_transform_func.Fn()(srcType, destType, vcbu(transformFunc))
}
func GValueTypeCompatible(srcType, destType GType) bool {
	return g_value_type_compatible.Fn()(srcType, destType)
}
func GValueTypeTransformable(srcType, destType GType) bool {
	return g_value_type_transformable.Fn()(srcType, destType)
}

func (value *GValue) SetChar(v int8)      { g_value_set_char.Fn()(value, v) }
func (value *GValue) GetChar() int8       { return g_value_get_char.Fn()(value) }
func (value *GValue) SetSChar(v int8)     { g_value_set_schar.Fn()(value, v) }
func (value *GValue) GetSChar() int8      { return g_value_get_schar.Fn()(value) }
func (value *GValue) SetUChar(v uint8)    { g_value_set_uchar.Fn()(value, v) }
func (value *GValue) GetUChar() uint8     { return g_value_get_uchar.Fn()(value) }
func (value *GValue) SetBoolean(v bool)   { g_value_set_boolean.Fn()(value, v) }
func (value *GValue) GetBoolean() bool    { return g_value_get_boolean.Fn()(value) }
func (value *GValue) SetInt(v int32)      { g_value_set_int.Fn()(value, v) }
func (value *GValue) GetInt() int32       { return g_value_get_int.Fn()(value) }
func (value *GValue) SetUInt(v uint32)    { g_value_set_uint.Fn()(value, v) }
func (value *GValue) GetUInt() uint32     { return g_value_get_uint.Fn()(value) }
func (value *GValue) SetLong(v int64)     { g_value_set_long.Fn()(value, v) }
func (value *GValue) GetLong() int64      { return g_value_get_long.Fn()(value) }
func (value *GValue) SetULong(v uint64)   { g_value_set_ulong.Fn()(value, v) }
func (value *GValue) GetULong() uint64    { return g_value_get_ulong.Fn()(value) }
func (value *GValue) SetInt64(v int64)    { g_value_set_int64.Fn()(value, v) }
func (value *GValue) GetInt64() int64     { return g_value_get_int64.Fn()(value) }
func (value *GValue) SetUInt64(v uint64)  { g_value_set_uint64.Fn()(value, v) }
func (value *GValue) GetUInt64() uint64   { return g_value_get_uint64.Fn()(value) }
func (value *GValue) SetFloat(v float32)  { g_value_set_float.Fn()(value, v) }
func (value *GValue) GetFloat() float32   { return g_value_get_float.Fn()(value) }
func (value *GValue) SetDouble(v float64) { g_value_set_double.Fn()(value, v) }
func (value *GValue) GetDouble() float64  { return g_value_get_double.Fn()(value) }
func (value *GValue) SetString(v string) {
	c := cc.NewString(v)
	defer c.Free()
	g_value_set_string.Fn()(value, c)
}
func (value *GValue) SetStaticString(v string) {
	g_value_set_static_string.Fn()(value, cc.NewString(v))
}
func (value *GValue) SetInternedString(v string) {
	g_value_set_interned_string.Fn()(value, cc.NewString(v))
}
func (value *GValue) GetString() string { return g_value_get_string.Fn()(value).String() }
func (value *GValue) DupString() string {
	ss := g_value_dup_string.Fn()(value)
	defer ss.Free()
	return ss.String()
}
func (value *GValue) StealString() cc.String       { return g_value_steal_string.Fn()(value) }
func (value *GValue) SetPointer(v uptr)            { g_value_set_pointer.Fn()(value, v) }
func (value *GValue) GetPointer() uptr             { return g_value_get_pointer.Fn()(value) }
func (value *GValue) SetGType(v GType)             { g_value_set_gtype.Fn()(value, v) }
func (value *GValue) GetGType() GType              { return g_value_get_gtype.Fn()(value) }
func (value *GValue) SetVariant(v *glib.GVariant)  { g_value_set_variant.Fn()(value, v) }
func (value *GValue) TakeVariant(v *glib.GVariant) { g_value_take_variant.Fn()(value, v) }
func (value *GValue) GetVariant() *glib.GVariant   { return g_value_get_variant.Fn()(value) }
func (value *GValue) DupVariant() *glib.GVariant   { return g_value_dup_variant.Fn()(value) }

// #endregion

// #region GTypes
func GTypeDate() GType               { return g_date_get_type.Fn()() }
func GTypeStrv() GType               { return g_strv_get_type.Fn()() }
func GTypeGString() GType            { return g_gstring_get_type.Fn()() }
func GTypeHashTable() GType          { return g_hash_table_get_type.Fn()() }
func GTypeArray() GType              { return g_array_get_type.Fn()() }
func GTypeByteArray() GType          { return g_byte_array_get_type.Fn()() }
func GTypePtrArray() GType           { return g_ptr_array_get_type.Fn()() }
func GTypeBytes() GType              { return g_bytes_get_type.Fn()() }
func GTypeVariantType() GType        { return g_variant_type_get_gtype.Fn()() }
func GTypeRegex() GType              { return g_regex_get_type.Fn()() }
func GTypeMatchInfo() GType          { return g_match_info_get_type.Fn()() }
func GTypeError() GType              { return g_error_get_type.Fn()() }
func GTypeDateTime() GType           { return g_date_time_get_type.Fn()() }
func GTypeTimeZone() GType           { return g_time_zone_get_type.Fn()() }
func GTypeIOChannel() GType          { return g_io_channel_get_type.Fn()() }
func GTypeIOCondition() GType        { return g_io_condition_get_type.Fn()() }
func GTypeVariantBuilder() GType     { return g_variant_builder_get_type.Fn()() }
func GTypeVariantDict() GType        { return g_variant_dict_get_type.Fn()() }
func GTypeKeyFile() GType            { return g_key_file_get_type.Fn()() }
func GTypeMainLoop() GType           { return g_main_loop_get_type.Fn()() }
func GTypeMainContext() GType        { return g_main_context_get_type.Fn()() }
func GTypeSource() GType             { return g_source_get_type.Fn()() }
func GTypePollfd() GType             { return g_pollfd_get_type.Fn()() }
func GTypeThread() GType             { return g_thread_get_type.Fn()() }
func GTypeChecksum() GType           { return g_checksum_get_type.Fn()() }
func GTypeMarkupParseContext() GType { return g_markup_parse_context_get_type.Fn()() }
func GTypeMappedFile() GType         { return g_mapped_file_get_type.Fn()() }
func GTypeOptionGroup() GType        { return g_option_group_get_type.Fn()() }
func GTypeUri() GType                { return g_uri_get_type.Fn()() }
func GTypeTree() GType               { return g_tree_get_type.Fn()() }
func GTypePatternSpec() GType        { return g_pattern_spec_get_type.Fn()() }
func GTypeBookmarkFile() GType       { return g_bookmark_file_get_type.Fn()() }
func GTypeHmac() GType               { return g_hmac_get_type.Fn()() }
func GTypeDir() GType                { return g_dir_get_type.Fn()() }
func GTypeRand() GType               { return g_rand_get_type.Fn()() }
func GTypeStrvBuilder() GType        { return g_strv_builder_get_type.Fn()() }

// #endregion

// #region custom

// func kindToGtype(kind reflect.Kind) GType {
// 	switch kind {
// 	case reflect.Bool:
// 		return Gbool
// 	case reflect.Int:
// 		return Gint64
// 	case reflect.Int8:
// 		return Gchar
// 	case reflect.Int16:
// 		panic("int16 not suport")
// 	case reflect.Int32:
// 		return Gint32
// 	case reflect.Int64:
// 		return Gint64

// 	case reflect.Uint:
// 		return Guint64
// 	case reflect.Uint8:
// 		return Guchar
// 	case reflect.Uint16:
// 		panic("uint16 not suport")
// 	case reflect.Uint32:
// 		return Guint32
// 	case reflect.Uint64:
// 		return Guint64

// 	case reflect.Float32:
// 		return Gfloat
// 	case reflect.Float64:
// 		return Gdouble

// 	default:
// 		return Gpointer

// 	}
// }

// type SigcAttach[Obj GObjectIface, Fn any] struct {
// 	connectMake func(obj Obj, fn Fn) uint64
// 	emitFunc    func(obj Obj) Fn
// }

// func (sig *SigcAttach[Obj, Fn]) Connect(obj Obj, fn Fn) uint64 { return sig.connectMake(obj, fn) }
// func (sig *SigcAttach[Obj, Fn]) Emit(obj Obj) Fn               { return sig.emitFunc(obj) }

// func NewSigcAttach[Obj GObjectIface, Fn any](gtype GType, signalId string, signalFlags GSignalFlags) *SigcAttach[Obj, Fn] {
// 	var a Fn
// 	fnType := reflect.TypeOf(a)
// 	fnOutNum := fnType.NumOut()
// 	if fnType.Kind() != reflect.Func {
// 		panic(fmt.Errorf("Type T must be a func, but got %v", fnType))
// 	}
// 	if fnOutNum > 1 {
// 		panic(fmt.Errorf("Type T must be a func with at most one return value, but got %d return values", fnType.NumOut()))
// 	}

// 	fnInNum := fnType.NumIn()

// 	// not exist
// 	if exist := SignalLookup(signalId, gtype) != 0; !exist {
// 		// out type
// 		cOut := Gvoid
// 		if fnOutNum > 0 {
// 			cOut = kindToGtype(fnType.Out(0).Kind())
// 		}
// 		cIns := make([]GType, fnInNum)
// 		for i := 0; i < fnInNum; i++ {
// 			cIns[i] = kindToGtype(fnType.In(i).Kind())
// 		}
// 		NewSignal(signalId, gtype, signalFlags, cOut, cIns...)
// 	}

// 	fnOut := []reflect.Type{}
// 	var connectType reflect.Type
// 	{
// 		if fnOutNum > 0 {
// 			fnOut = append(fnOut, fnType.Out(0))
// 		}

// 		goIns := []reflect.Type{reflect.TypeOf(unsafe.Pointer(nil))}
// 		for i := 0; i < fnInNum; i++ {
// 			goIns = append(goIns, fnType.In(i))
// 		}
// 		connectType = reflect.FuncOf(goIns, fnOut, false)
// 	}

// 	rSignal := new(SigcAttach[Obj, Fn])
// 	rSignal.connectMake = func(obj Obj, fn Fn) uint64 {
// 		realFn := reflect.ValueOf(fn)
// 		connect := reflect.MakeFunc(connectType,
// 			func(args []reflect.Value) (results []reflect.Value) {
// 				return realFn.Call(args[1:])
// 			}).Interface()
// 		return (*(**GObject)(uptr(&obj))).SignalConnect(signalId, connect, nil)
// 	}

// 	var emitOut reflect.Type
// 	if fnOutNum > 0 {
// 		emitOut = fnOut[0]
// 	}

// 	var emitFn func(obj Obj, as []interface{}) []reflect.Value
// 	if fnOutNum > 0 {
// 		emitFn = func(obj Obj, as []interface{}) []reflect.Value {
// 			v := reflect.New(emitOut)
// 			(*(**GObject)(uptr(&obj))).SignalEmit(signalId, append(as, v.Interface())...)
// 			return []reflect.Value{v.Elem()}
// 		}
// 	} else {
// 		emitFn = func(obj Obj, as []interface{}) []reflect.Value {
// 			(*(**GObject)(uptr(&obj))).SignalEmit(signalId, as...)
// 			return []reflect.Value{}
// 		}
// 	}
// 	rSignal.emitFunc = func(obj Obj) Fn {
// 		fn := reflect.MakeFunc(fnType, func(args []reflect.Value) (results []reflect.Value) {
// 			as := make([]interface{}, len(args))
// 			for i := 0; i < len(args); i++ {
// 				as[i] = args[i].Interface()
// 			}
// 			return emitFn(obj, as)
// 			// if fnOutNum > 0 {
// 			// 	v := reflect.New(emitOut)
// 			// 	(*(**GObject)(uptr(&obj))).SignalEmit(signalId, append(as, v.Interface())...)
// 			// 	return []reflect.Value{v.Elem()}
// 			// } else {
// 			// 	(*(**GObject)(uptr(&obj))).SignalEmit(signalId, as...)
// 			// 	return []reflect.Value{}
// 			// }
// 		})
// 		return fn.Interface().(Fn)
// 	}

// 	return rSignal
// }

// type SigcObj struct {
// 	gtype func() GType
// 	obj   *GObject
// }
// type SigcSignal[Fn any] struct {
// 	obj *GObject
// 	sig SigcAttach[*GObject, Fn]
// }

// func (sig *SigcSignal[Fn]) Connect(fn Fn) uint64 { return sig.sig.connectMake(sig.obj, fn) }
// func (sig *SigcSignal[Fn]) Emit() Fn             { return sig.sig.emitFunc(sig.obj) }

// func NewSigcObj(ObjName string) *SigcObj {
// 	type obj struct{ parent GObjectObj }
// 	type class struct{ parent GObjectClass }
// 	const classSize = uint32(unsafe.Sizeof(class{}))
// 	const objSize = uint32(unsafe.Sizeof(obj{}))

// 	typ := func(name string) func() GType {
// 		var gtype GType
// 		return func() GType {
// 			if gtype == 0 {
// 				gtype = GTypeRegisterStaticSimple(Gobject, name,
// 					classSize, func(GClass, _ uptr) {},
// 					objSize, func(Instace, GClass uptr) {}, GTypeFlagNone)
// 			}
// 			return gtype
// 		}
// 	}(ObjName)
// 	return &SigcObj{gtype: typ, obj: NewGObject(typ())}
// }
// func NewSigcSignal[Fn any](obj *SigcObj, signal string, signalFlags GSignalFlags) *SigcSignal[Fn] {
// 	sig := NewSigcAttach[*GObject, Fn](obj.gtype(), signal, signalFlags)
// 	return &SigcSignal[Fn]{obj: obj.obj, sig: *sig}
// }

type Sigc[Fn any] struct {
	emit        Fn
	emitIdel    Fn
	connectLock sync.RWMutex
	connectList []reflect.Value
}

func NewSigc[Fn any]() *Sigc[Fn] {
	var fn Fn
	sigc := &Sigc[Fn]{
		connectList: []reflect.Value{},
	}

	sigc.connectLock = sync.RWMutex{}
	sigc.emit = reflect.MakeFunc(reflect.TypeOf(fn),
		func(args []reflect.Value) (results []reflect.Value) {
			sigc.connectLock.RLock()
			defer sigc.connectLock.RUnlock()

			for i := range sigc.connectList {
				fn := sigc.connectList[i]
				if fn.CanInt() {
					continue
				}
				results = fn.Call(args)
			}

			return results
		}).Interface().(Fn)
	sigc.emitIdel = reflect.MakeFunc(reflect.TypeOf(fn),
		func(args []reflect.Value) (results []reflect.Value) {
			sigc.connectLock.RLock()
			defer sigc.connectLock.RUnlock()

			wg := sync.WaitGroup{}
			wg.Add(1)
			glib.IdleAddOnce(func() {
				for i := range sigc.connectList {
					fn := sigc.connectList[i]
					if fn.CanInt() {
						continue
					}
					results = fn.Call(args)
				}
				wg.Done()
			})
			wg.Wait()
			return results
		}).Interface().(Fn)

	return sigc
}
func (sigc *Sigc[Fn]) Connect(fn Fn, liveWith GObjectIface) *SigcID[Fn] {
	sigc.connectLock.Lock()
	defer sigc.connectLock.Unlock()

	v := reflect.ValueOf(fn)
	sigc.connectList = append(sigc.connectList, v)

	idx := len(sigc.connectList) - 1

	id := new(SigcID[Fn])
	id.sigc = sigc
	id.id = idx
	id.fn = v

	if GetGObjectIface(liveWith) != nil {
		liveWith.GetGObjectIface().WeakRef(func(_ uptr, obj *GObject) {
			sigc.connectLock.Lock()
			defer sigc.connectLock.Unlock()
			sigc.connectList[idx] = reflect.ValueOf(0)
		})
	}

	return id
}
func (sigc *Sigc[Fn]) Emit() Fn     { return sigc.emit }
func (sigc *Sigc[Fn]) EmitIdel() Fn { return sigc.emitIdel }

type SigcID[Fn any] struct {
	sigc *Sigc[Fn]
	id   int
	fn   reflect.Value
}

func (id *SigcID[Fn]) Block() {
	id.sigc.connectLock.Lock()
	defer id.sigc.connectLock.Unlock()
	id.sigc.connectList[id.id] = reflect.ValueOf(0)
}
func (id *SigcID[Fn]) Unblock() {
	id.sigc.connectLock.Lock()
	defer id.sigc.connectLock.Unlock()
	id.sigc.connectList[id.id] = id.fn
}
func (id *SigcID[Fn]) Disconnect() {
	id.sigc.connectLock.Lock()
	defer id.sigc.connectLock.Unlock()
	id.sigc.connectList[id.id] = reflect.ValueOf(0)
	id = nil
}

// #endregion
