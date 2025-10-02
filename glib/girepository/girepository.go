package girepository

import (
	"errors"
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/glib/glib"
	"github.com/jinzhongmin/gg/glib/gobject"
)

type uptr = unsafe.Pointer
type iptr = uintptr

func slice(p uptr, n int) uptr { return uptr(&[2]uptr{uptr(p), uptr(iptr(n))}) }

// func vfn(name string, fptr interface{}) { cc.RFunc(fptr, name) }

// region ArgInfo

type ArgInfo struct{ BaseInfo }

func (info *ArgInfo) GetDirection() Direction {
	return gi_arg_info_get_direction.Fn()(info)
}

func (info *ArgInfo) IsReturnValue() bool {
	return gi_arg_info_is_return_value.Fn()(info)
}

func (info *ArgInfo) IsOptional() bool {
	return gi_arg_info_is_optional.Fn()(info)
}

func (info *ArgInfo) IsCallerAllocates() bool {
	return gi_arg_info_is_caller_allocates.Fn()(info)
}

func (info *ArgInfo) MayBeNull() bool {
	return gi_arg_info_may_be_null.Fn()(info)
}

func (info *ArgInfo) IsSkip() bool {
	return gi_arg_info_is_skip.Fn()(info)
}

func (info *ArgInfo) GetOwnershipTransfer() Transfer {
	return gi_arg_info_get_ownership_transfer.Fn()(info)
}

func (info *ArgInfo) GetScope() ScopeType {
	return gi_arg_info_get_scope.Fn()(info)
}

func (info *ArgInfo) GetClosureIndex() (uint32, bool) {
	var index uint32
	ok := gi_arg_info_get_closure_index.Fn()(info, &index)
	return index, ok
}

func (info *ArgInfo) GetDestroyIndex() (uint32, bool) {
	var index uint32
	ok := gi_arg_info_get_destroy_index.Fn()(info, &index)
	return index, ok
}

func (info *ArgInfo) GetTypeInfo() *TypeInfo {
	return gi_arg_info_get_type_info.Fn()(info)
}

func (info *ArgInfo) LoadTypeInfo(typeInfo *TypeInfo) {
	gi_arg_info_load_type_info.Fn()(info, typeInfo)
}

// endregion

// region GIBaseInfo

type BaseInfo struct{ gobject.GTypeInstance }

func (info *BaseInfo) Ref() *BaseInfo {
	return gi_base_info_ref.Fn()(info)
}

func (info *BaseInfo) Unref() {
	gi_base_info_unref.Fn()(info)
}

func (info *BaseInfo) Clear() {
	gi_base_info_clear.Fn()(info)
}

func (info *BaseInfo) GetName() string {
	return gi_base_info_get_name.Fn()(info)
}

func (info *BaseInfo) GetNamespace() string {
	return gi_base_info_get_namespace.Fn()(info)
}

func (info *BaseInfo) IsDeprecated() bool {
	return gi_base_info_is_deprecated.Fn()(info)
}

func (info *BaseInfo) GetAttribute(name string) string {
	return gi_base_info_get_attribute.Fn()(info, name)
}

func (info *BaseInfo) IterateAttributes(iter *AttributeIter) (ok bool, name, value string) {
	var pName, pValue cc.String
	ok = gi_base_info_iterate_attributes.Fn()(info, iter, &pName, &pValue) != 0
	if ok {
		name = pName.String()
		value = pValue.String()
	}
	return
}

func (info *BaseInfo) GetContainer() *BaseInfo {
	return gi_base_info_get_container.Fn()(info)
}

func (info *BaseInfo) GetTypelib() *Typelib {
	return gi_base_info_get_typelib.Fn()(info)
}

func (info1 *BaseInfo) Equal(info2 *BaseInfo) bool {
	return gi_base_info_equal.Fn()(info1, info2)
}

// endregion

// region CallableInfo

type CallableInfo struct{ BaseInfo }

func (info *CallableInfo) IsMethod() bool {
	return gi_callable_info_is_method.Fn()(info)
}

func (info *CallableInfo) CanThrowGError() bool {
	return gi_callable_info_can_throw_gerror.Fn()(info)
}

func (info *CallableInfo) GetReturnType() *TypeInfo {
	return gi_callable_info_get_return_type.Fn()(info)
}

func (info *CallableInfo) LoadReturnType(typeInfo *TypeInfo) {
	gi_callable_info_load_return_type.Fn()(info, typeInfo)
}

func (info *CallableInfo) GetReturnAttribute(name string) string {
	return gi_callable_info_get_return_attribute.Fn()(info, name)
}

func (info *CallableInfo) IterateReturnAttributes(iter *AttributeIter) (ok bool, name, value string) {
	var pName, pValue cc.String
	ok = gi_callable_info_iterate_return_attributes.Fn()(info, iter, &pName, &pValue)
	if ok {
		name = pName.String()
		value = pValue.String()
	}
	return
}

func (info *CallableInfo) GetCallerOwns() Transfer {
	return gi_callable_info_get_caller_owns.Fn()(info)
}

func (info *CallableInfo) MayReturnNull() bool {
	return gi_callable_info_may_return_null.Fn()(info)
}

func (info *CallableInfo) SkipReturn() bool {
	return gi_callable_info_skip_return.Fn()(info)
}

func (info *CallableInfo) GetNArgs() uint32 {
	return gi_callable_info_get_n_args.Fn()(info)
}

func (info *CallableInfo) GetArg(n uint32) *ArgInfo {
	return gi_callable_info_get_arg.Fn()(info, n)
}

func (info *CallableInfo) LoadArg(n uint32, arg *ArgInfo) {
	gi_callable_info_load_arg.Fn()(info, n, arg)
}

// func (info *CallableInfo) Invoke(function uptr, inArgs uptr, nInArgs uint, outArgs uptr, nOutArgs uint, returnValue uptr) (bool, error) {
// 	var gerr *glib.GError
// 	success := gi_callable_info_invoke(info, function, inArgs, nInArgs, outArgs, nOutArgs, returnValue, &gerr)
// 	if gerr != nil {
// 		return false, errors.New(gerr.Message())
// 	}
// 	return success, nil
// }

func (info *CallableInfo) GetInstanceOwnershipTransfer() Transfer {
	return gi_callable_info_get_instance_ownership_transfer.Fn()(info)
}

// endregion

// region CallbackInfo

type CallbackInfo struct{ CallableInfo }

// endregion

// region ConstantInfo

type Argument struct{}

type ConstantInfo struct{ BaseInfo }

func (info *ConstantInfo) GetTypeInfo() *TypeInfo {
	return gi_constant_info_get_type_info.Fn()(info)
}

func (info *ConstantInfo) FreeValue(value *Argument) {
	gi_constant_info_free_value.Fn()(info, value)
}

func (info *ConstantInfo) GetValue(value *Argument) uint {
	return gi_constant_info_get_value.Fn()(info, value)
}

// endregion

// region GIEnumInfo

type EnumInfo struct{ RegisteredTypeInfo }

func (info *EnumInfo) GetNValues() uint32 {
	return gi_enum_info_get_n_values.Fn()(info)
}

func (info *EnumInfo) GetValue(n uint32) *ValueInfo {
	return gi_enum_info_get_value.Fn()(info, n)
}

func (info *EnumInfo) GetNMethods() uint32 {
	return gi_enum_info_get_n_methods.Fn()(info)
}

func (info *EnumInfo) GetMethod(n uint32) *FunctionInfo {
	return gi_enum_info_get_method.Fn()(info, n)
}

func (info *EnumInfo) GetStorageType() TypeTag {
	return gi_enum_info_get_storage_type.Fn()(info)
}

func (info *EnumInfo) GetErrorDomain() string {
	return gi_enum_info_get_error_domain.Fn()(info)
}

// endregion

// region FieldInfo

type FieldInfo struct{ BaseInfo }

func (info *FieldInfo) GetFlags() FieldInfoFlags {
	return gi_field_info_get_flags.Fn()(info)
}

func (info *FieldInfo) GetSize() uint {
	return gi_field_info_get_size.Fn()(info)
}

func (info *FieldInfo) GetOffset() uint {
	return gi_field_info_get_offset.Fn()(info)
}

func (info *FieldInfo) GetTypeInfo() *TypeInfo {
	return gi_field_info_get_type_info.Fn()(info)
}

func (info *FieldInfo) GetField(mem uptr, value *Argument) bool {
	return gi_field_info_get_field.Fn()(info, mem, value)
}

func (info *FieldInfo) SetField(mem uptr, value *Argument) bool {
	return gi_field_info_set_field.Fn()(info, mem, value)
}

// endregion

// region FlagsInfo

type FlagsInfo struct{ EnumInfo }

// endregion

// region FunctionInfo

type FunctionInfo struct{ CallableInfo }

func (info *FunctionInfo) GetSymbol() string {
	return gi_function_info_get_symbol.Fn()(info)
}

func (info *FunctionInfo) GetFlags() FunctionInfoFlags {
	return gi_function_info_get_flags.Fn()(info)
}

func (info *FunctionInfo) GetProperty() *PropertyInfo {
	return gi_function_info_get_property.Fn()(info)
}

func (info *FunctionInfo) GetVFunc() *VFuncInfo {
	return gi_function_info_get_vfunc.Fn()(info)
}

// endregion

// region GIInterfaceInfo

type InterfaceInfo struct{ RegisteredTypeInfo }

func (info *InterfaceInfo) GetNPrerequisites() uint32 {
	return gi_interface_info_get_n_prerequisites.Fn()(info)
}

func (info *InterfaceInfo) GetPrerequisite(n uint32) *BaseInfo {
	return gi_interface_info_get_prerequisite.Fn()(info, n)
}

func (info *InterfaceInfo) GetNProperties() uint32 {
	return gi_interface_info_get_n_properties.Fn()(info)
}

func (info *InterfaceInfo) GetProperty(n uint32) *PropertyInfo {
	return gi_interface_info_get_property.Fn()(info, n)
}

func (info *InterfaceInfo) GetNMethods() uint32 {
	return gi_interface_info_get_n_methods.Fn()(info)
}

func (info *InterfaceInfo) GetMethod(n uint32) *FunctionInfo {
	return gi_interface_info_get_method.Fn()(info, n)
}

func (info *InterfaceInfo) FindMethod(name string) *FunctionInfo {
	return gi_interface_info_find_method.Fn()(info, name)
}

func (info *InterfaceInfo) GetNSignals() uint32 {
	return gi_interface_info_get_n_signals.Fn()(info)
}

func (info *InterfaceInfo) GetSignal(n uint32) *SignalInfo {
	return gi_interface_info_get_signal.Fn()(info, n)
}

func (info *InterfaceInfo) FindSignal(name string) *SignalInfo {
	return gi_interface_info_find_signal.Fn()(info, name)
}

func (info *InterfaceInfo) GetNVFuncs() uint32 {
	return gi_interface_info_get_n_vfuncs.Fn()(info)
}

func (info *InterfaceInfo) GetVFunc(n uint32) *VFuncInfo {
	return gi_interface_info_get_vfunc.Fn()(info, n)
}

func (info *InterfaceInfo) FindVFunc(name string) *VFuncInfo {
	return gi_interface_info_find_vfunc.Fn()(info, name)
}

func (info *InterfaceInfo) GetNConstants() uint32 {
	return gi_interface_info_get_n_constants.Fn()(info)
}

func (info *InterfaceInfo) GetConstant(n uint32) *ConstantInfo {
	return gi_interface_info_get_constant.Fn()(info, n)
}

func (info *InterfaceInfo) GetIfaceStruct() *StructInfo {
	return gi_interface_info_get_iface_struct.Fn()(info)
}

// endregion

// region ObjectInfo

type ObjectInfo struct{ RegisteredTypeInfo }

func (info *ObjectInfo) GetTypeName() string {
	return gi_object_info_get_type_name.Fn()(info)
}

func (info *ObjectInfo) GetTypeInitFunctionName() string {
	return gi_object_info_get_type_init_function_name.Fn()(info)
}

func (info *ObjectInfo) GetAbstract() bool {
	return gi_object_info_get_abstract.Fn()(info)
}

func (info *ObjectInfo) GetFinal() bool {
	return gi_object_info_get_final.Fn()(info)
}

func (info *ObjectInfo) GetFundamental() bool {
	return gi_object_info_get_fundamental.Fn()(info)
}

func (info *ObjectInfo) GetParent() *ObjectInfo {
	return gi_object_info_get_parent.Fn()(info)
}

func (info *ObjectInfo) GetNInterfaces() uint32 {
	return gi_object_info_get_n_interfaces.Fn()(info)
}

func (info *ObjectInfo) GetInterface(n uint32) *InterfaceInfo {
	return gi_object_info_get_interface.Fn()(info, n)
}

func (info *ObjectInfo) GetNFields() uint32 {
	return gi_object_info_get_n_fields.Fn()(info)
}

func (info *ObjectInfo) GetField(n uint32) *FieldInfo {
	return gi_object_info_get_field.Fn()(info, n)
}

func (info *ObjectInfo) GetNProperties() uint32 {
	return gi_object_info_get_n_properties.Fn()(info)
}

func (info *ObjectInfo) GetProperty(n uint32) *PropertyInfo {
	return gi_object_info_get_property.Fn()(info, n)
}

func (info *ObjectInfo) GetNMethods() uint32 {
	return gi_object_info_get_n_methods.Fn()(info)
}

func (info *ObjectInfo) GetMethod(n uint32) *FunctionInfo {
	return gi_object_info_get_method.Fn()(info, n)
}

func (info *ObjectInfo) FindMethod(name string) *FunctionInfo {
	return gi_object_info_find_method.Fn()(info, name)
}

func (info *ObjectInfo) FindMethodUsingInterfaces(name string) (*FunctionInfo, *BaseInfo) {
	var declarer *BaseInfo
	method := gi_object_info_find_method_using_interfaces.Fn()(info, name, &declarer)
	return method, declarer
}

func (info *ObjectInfo) GetNSignals() uint32 {
	return gi_object_info_get_n_signals.Fn()(info)
}

func (info *ObjectInfo) GetSignal(n uint32) *SignalInfo {
	return gi_object_info_get_signal.Fn()(info, n)
}

func (info *ObjectInfo) FindSignal(name string) *SignalInfo {
	return gi_object_info_find_signal.Fn()(info, name)
}

func (info *ObjectInfo) GetNVFuncs() uint32 {
	return gi_object_info_get_n_vfuncs.Fn()(info)
}

func (info *ObjectInfo) GetVFunc(n uint32) *VFuncInfo {
	return gi_object_info_get_vfunc.Fn()(info, n)
}

func (info *ObjectInfo) FindVFunc(name string) *VFuncInfo {
	return gi_object_info_find_vfunc.Fn()(info, name)
}

func (info *ObjectInfo) FindVFuncUsingInterfaces(name string) (*VFuncInfo, *BaseInfo) {
	var declarer *BaseInfo
	vfunc := gi_object_info_find_vfunc_using_interfaces.Fn()(info, name, &declarer)
	return vfunc, declarer
}

func (info *ObjectInfo) GetNConstants() uint32 {
	return gi_object_info_get_n_constants.Fn()(info)
}

func (info *ObjectInfo) GetConstant(n uint32) *ConstantInfo {
	return gi_object_info_get_constant.Fn()(info, n)
}

func (info *ObjectInfo) GetClassStruct() *StructInfo {
	return gi_object_info_get_class_struct.Fn()(info)
}

func (info *ObjectInfo) GetRefFunctionName() string {
	return gi_object_info_get_ref_function_name.Fn()(info)
}

func (info *ObjectInfo) GetRefFunctionPointer() uptr {
	return gi_object_info_get_ref_function_pointer.Fn()(info)
}

func (info *ObjectInfo) GetUnrefFunctionName() string {
	return gi_object_info_get_unref_function_name.Fn()(info)
}

func (info *ObjectInfo) GetUnrefFunctionPointer() uptr {
	return gi_object_info_get_unref_function_pointer.Fn()(info)
}

func (info *ObjectInfo) GetSetValueFunctionName() string {
	return gi_object_info_get_set_value_function_name.Fn()(info)
}

func (info *ObjectInfo) GetSetValueFunctionPointer() uptr {
	return gi_object_info_get_set_value_function_pointer.Fn()(info)
}

func (info *ObjectInfo) GetGetValueFunctionName() string {
	return gi_object_info_get_get_value_function_name.Fn()(info)
}

func (info *ObjectInfo) GetGetValueFunctionPointer() uptr {
	return gi_object_info_get_get_value_function_pointer.Fn()(info)
}

// endregion

// region PropertyInfo

type PropertyInfo struct{ BaseInfo }

func (info *PropertyInfo) GetFlags() int32 {
	return gi_property_info_get_flags.Fn()(info)
}

func (info *PropertyInfo) GetTypeInfo() *TypeInfo {
	return gi_property_info_get_type_info.Fn()(info)
}

func (info *PropertyInfo) GetOwnershipTransfer() int32 {
	return gi_property_info_get_ownership_transfer.Fn()(info)
}

func (info *PropertyInfo) GetSetter() *FunctionInfo {
	return gi_property_info_get_setter.Fn()(info)
}

func (info *PropertyInfo) GetGetter() *FunctionInfo {
	return gi_property_info_get_getter.Fn()(info)
}

// endregion

// region RegisteredTypeInfo

type RegisteredTypeInfo struct{ BaseInfo }

func (info *RegisteredTypeInfo) GetTypeName() string {
	return gi_registered_type_info_get_type_name.Fn()(info)
}

func (info *RegisteredTypeInfo) GetTypeInitFunctionName() string {
	return gi_registered_type_info_get_type_init_function_name.Fn()(info)
}

func (info *RegisteredTypeInfo) GetGType() gobject.GType {
	return gi_registered_type_info_get_g_type.Fn()(info)
}

func (info *RegisteredTypeInfo) IsBoxed() bool {
	return gi_registered_type_info_is_boxed.Fn()(info)
}

// endregion

// region Repository

type Repository struct{}

func NewRepository() *Repository { return gi_repository_new.Fn()() }
func (repo *Repository) PrependSearchPath(directory string) {
	gi_repository_prepend_search_path.Fn()(repo, directory)
}
func (repo *Repository) PrependLibraryPath(directory string) {
	gi_repository_prepend_library_path.Fn()(repo, directory)
}
func (repo *Repository) GetSearchPath() []string {
	var nPaths uint64
	paths := gi_repository_get_search_path.Fn()(repo, &nPaths)
	return paths.Strings(nPaths)
}
func (repo *Repository) GetLibraryPath() []string {
	var nPaths uint64
	paths := gi_repository_get_library_path.Fn()(repo, &nPaths)
	return paths.Strings(nPaths)
}
func (repo *Repository) LoadTypelib(typelib *Typelib, flags RepositoryLoadFlags) (string, error) {
	var gerr *glib.GError
	result := gi_repository_load_typelib.Fn()(repo, typelib, flags, &gerr)
	if gerr != nil {
		return "", errors.New(gerr.Message.String())
	}
	return result, nil
}
func (repo *Repository) IsRegistered(namespace, version string) bool {
	return gi_repository_is_registered.Fn()(repo, namespace, version)
}
func (repo *Repository) FindByName(namespace, name string) *BaseInfo {
	return gi_repository_find_by_name.Fn()(repo, namespace, name)
}
func (repo *Repository) EnumerateVersions(namespace string) []string {
	var nVersions uint64
	versions := gi_repository_enumerate_versions.Fn()(repo, namespace, &nVersions)
	return versions.Strings(nVersions)
}
func (repo *Repository) Require(namespace, version string, flags RepositoryLoadFlags) (*Typelib, error) {
	var gerr *glib.GError
	typelib := gi_repository_require.Fn()(repo, namespace, version, flags, &gerr)
	if gerr != nil {
		return nil, errors.New(gerr.Message.String())
	}
	return typelib, nil
}
func (repo *Repository) RequirePrivate(namespace, version, module string, flags RepositoryLoadFlags) (*Typelib, error) {
	var gerr *glib.GError
	typelib := gi_repository_require_private.Fn()(repo, namespace, version, module, flags, &gerr)
	if gerr != nil {
		return nil, errors.New(gerr.Message.String())
	}
	return typelib, nil
}
func (repo *Repository) GetImmediateDependencies(namespace string) []string {
	var nDeps uint64
	deps := gi_repository_get_immediate_dependencies.Fn()(repo, namespace, &nDeps)
	return deps.Strings(nDeps)
}
func (repo *Repository) GetDependencies(namespace string) []string {
	var nDeps uint64
	deps := gi_repository_get_dependencies.Fn()(repo, namespace, &nDeps)
	return deps.Strings(nDeps)
}
func (repo *Repository) GetLoadedNamespaces() []string {
	var nNamespacesOut uint64
	versions := gi_repository_get_loaded_namespaces.Fn()(repo, &nNamespacesOut)
	return versions.Strings(nNamespacesOut)
}
func (repo *Repository) FindByGType(gtype gobject.GType) *BaseInfo {
	return gi_repository_find_by_gtype.Fn()(repo, gtype)
}
func (repo *Repository) GetObjectGTypeInterfaces(gtype gobject.GType) ([]*InterfaceInfo, error) {
	var nIfaces uint
	var ifaces **InterfaceInfo
	gi_repository_get_object_gtype_interfaces.Fn()(repo, gtype, &nIfaces, &ifaces)
	if ifaces == nil {
		return nil, nil
	}
	return *((*[]*InterfaceInfo)(slice(uptr(ifaces), int(nIfaces)))), nil
}
func (repo *Repository) GetNInfos(namespace string) uint32 {
	return gi_repository_get_n_infos.Fn()(repo, namespace)
}
func (repo *Repository) GetInfo(namespace string, index uint32) *BaseInfo {
	return gi_repository_get_info.Fn()(repo, namespace, index)
}
func (repo *Repository) FindByErrorDomain(quark glib.GQuark) *EnumInfo {
	return gi_repository_find_by_error_domain.Fn()(repo, quark)
}
func (repo *Repository) GetTypelibPath(namespace string) string {
	return gi_repository_get_typelib_path.Fn()(repo, namespace)
}
func (repo *Repository) GetSharedLibraries(namespace string) []string {
	var nLibs uint64
	libs := gi_repository_get_shared_libraries.Fn()(repo, namespace, &nLibs)
	return libs.Strings(nLibs)
}
func (repo *Repository) GetCPrefix(namespace string) string {
	return gi_repository_get_c_prefix.Fn()(repo, namespace)
}
func (repo *Repository) GetVersion(namespace string) string {
	return gi_repository_get_version.Fn()(repo, namespace)
}
func (repo *Repository) GetOptionGroup() *glib.GOptionGroup {
	return gi_repository_get_option_group.Fn()()
}
func (repo *Repository) Dump(inputFilename, outputFilename string) (bool, error) {
	var gerr *glib.GError
	success := gi_repository_dump.Fn()(inputFilename, outputFilename, &gerr)
	if gerr != nil {
		return false, errors.New(gerr.Message.String())
	}
	return success, nil
}
func (repo *Repository) ErrorQuark() glib.GQuark {
	return gi_repository_error_quark.Fn()()
}

// endregion

// region SignalInfo

type SignalInfo struct{ CallableInfo }

func (info *SignalInfo) GetFlags() int32 {
	return gi_signal_info_get_flags.Fn()(info)
}

func (info *SignalInfo) GetClassClosure() *VFuncInfo {
	return gi_signal_info_get_class_closure.Fn()(info)
}

func (info *SignalInfo) TrueStopsEmit() bool {
	return gi_signal_info_true_stops_emit.Fn()(info)
}

// endregion

// region StructInfo

type StructInfo struct{ RegisteredTypeInfo }

func (info *StructInfo) GetNFields() uint32 {
	return gi_struct_info_get_n_fields.Fn()(info)
}

func (info *StructInfo) GetField(n uint32) *FieldInfo {
	return gi_struct_info_get_field.Fn()(info, n)
}

func (info *StructInfo) FindField(name string) *FieldInfo {
	return gi_struct_info_find_field.Fn()(info, name)
}

func (info *StructInfo) GetNMethods() uint32 {
	return gi_struct_info_get_n_methods.Fn()(info)
}

func (info *StructInfo) GetMethod(n uint32) *FunctionInfo {
	return gi_struct_info_get_method.Fn()(info, n)
}

func (info *StructInfo) FindMethod(name string) *FunctionInfo {
	return gi_struct_info_find_method.Fn()(info, name)
}

func (info *StructInfo) GetSize() uint {
	return gi_struct_info_get_size.Fn()(info)
}

func (info *StructInfo) GetAlignment() uint {
	return gi_struct_info_get_alignment.Fn()(info)
}

func (info *StructInfo) IsGTypeStruct() bool {
	return gi_struct_info_is_gtype_struct.Fn()(info)
}

func (info *StructInfo) IsForeign() bool {
	return gi_struct_info_is_foreign.Fn()(info)
}

func (info *StructInfo) GetCopyFunctionName() string {
	return gi_struct_info_get_copy_function_name.Fn()(info)
}

func (info *StructInfo) GetFreeFunctionName() string {
	return gi_struct_info_get_free_function_name.Fn()(info)
}

// endregion

// region TypeInfo

type TypeInfo struct{ BaseInfo }

func TypeTagToString(tag int32) string {
	return gi_type_tag_to_string.Fn()(tag)
}

func (info *TypeInfo) IsPointer() bool {
	return gi_type_info_is_pointer.Fn()(info)
}

func (info *TypeInfo) GetTag() int32 {
	return gi_type_info_get_tag.Fn()(info)
}
func (info *TypeInfo) GetTagString() string {
	return gi_type_tag_to_string.Fn()(gi_type_info_get_tag.Fn()(info))
}

func (info *TypeInfo) GetParamType(n uint32) *TypeInfo {
	return gi_type_info_get_param_type.Fn()(info, n)
}

func (info *TypeInfo) GetInterface() *BaseInfo {
	return gi_type_info_get_interface.Fn()(info)
}

func (info *TypeInfo) GetArrayLengthIndex() (uint32, bool) {
	var index uint32
	ok := gi_type_info_get_array_length_index.Fn()(info, &index)
	return index, ok
}

func (info *TypeInfo) GetArrayFixedSize() (uint, bool) {
	var size uint
	ok := gi_type_info_get_array_fixed_size.Fn()(info, &size)
	return size, ok
}

func (info *TypeInfo) IsZeroTerminated() bool {
	return gi_type_info_is_zero_terminated.Fn()(info)
}

func (info *TypeInfo) GetArrayType() int32 {
	return gi_type_info_get_array_type.Fn()(info)
}

func (info *TypeInfo) GetStorageType() int32 {
	return gi_type_info_get_storage_type.Fn()(info)
}

func (info *TypeInfo) ArgumentFromHashPointer(hashPointer uptr, arg *Argument) {
	gi_type_info_argument_from_hash_pointer.Fn()(info, hashPointer, arg)
}

func (info *TypeInfo) HashPointerFromArgument(arg *Argument) uptr {
	return gi_type_info_hash_pointer_from_argument.Fn()(info, arg)
}

func TypeTagArgumentFromHashPointer(storageType int32, hashPointer uptr, arg *Argument) {
	gi_type_tag_argument_from_hash_pointer.Fn()(storageType, hashPointer, arg)
}

func TypeTagHashPointerFromArgument(storageType int32, arg *Argument) uptr {
	return gi_type_tag_hash_pointer_from_argument.Fn()(storageType, arg)
}

// endregion

// region UnionInfo

type UnionInfo struct{ RegisteredTypeInfo }

func (info *UnionInfo) GetNFields() uint32 {
	return gi_union_info_get_n_fields.Fn()(info)
}

func (info *UnionInfo) GetField(n uint32) *FieldInfo {
	return gi_union_info_get_field.Fn()(info, n)
}

func (info *UnionInfo) GetNMethods() uint32 {
	return gi_union_info_get_n_methods.Fn()(info)
}

func (info *UnionInfo) GetMethod(n uint32) *FunctionInfo {
	return gi_union_info_get_method.Fn()(info, n)
}

func (info *UnionInfo) IsDiscriminated() bool {
	return gi_union_info_is_discriminated.Fn()(info)
}

func (info *UnionInfo) GetDiscriminatorOffset() (uint, bool) {
	var offset uint
	ok := gi_union_info_get_discriminator_offset.Fn()(info, &offset)
	return offset, ok
}

func (info *UnionInfo) GetDiscriminatorType() *TypeInfo {
	return gi_union_info_get_discriminator_type.Fn()(info)
}

func (info *UnionInfo) GetDiscriminator(n uint) *ConstantInfo {
	return gi_union_info_get_discriminator.Fn()(info, n)
}

func (info *UnionInfo) FindMethod(name string) *FunctionInfo {
	return gi_union_info_find_method.Fn()(info, name)
}

func (info *UnionInfo) GetSize() uint {
	return gi_union_info_get_size.Fn()(info)
}

func (info *UnionInfo) GetAlignment() uint {
	return gi_union_info_get_alignment.Fn()(info)
}

func (info *UnionInfo) GetCopyFunctionName() string {
	return gi_union_info_get_copy_function_name.Fn()(info)
}

func (info *UnionInfo) GetFreeFunctionName() string {
	return gi_union_info_get_free_function_name.Fn()(info)
}

// endregion

// region UnresolvedInfo

type UnresolvedInfo struct{ BaseInfo }

// endregion

// region ValueInfo

type ValueInfo struct{ BaseInfo }

func (info *ValueInfo) GetValue() int64 {
	return gi_value_info_get_value.Fn()(info)
}

// endregion

// region VFuncInfo

type VFuncInfo struct{ CallableInfo }

func (info *VFuncInfo) GetFlags() int32 {
	return gi_vfunc_info_get_flags.Fn()(info)
}

func (info *VFuncInfo) GetOffset() uint {
	return gi_vfunc_info_get_offset.Fn()(info)
}

func (info *VFuncInfo) GetSignal() *SignalInfo {
	return gi_vfunc_info_get_signal.Fn()(info)
}

func (info *VFuncInfo) GetInvoker() *FunctionInfo {
	return gi_vfunc_info_get_invoker.Fn()(info)
}

func (info *VFuncInfo) GetAddress(implementorGType gobject.GType) (uptr, error) {
	var gerr *glib.GError
	addr := gi_vfunc_info_get_address.Fn()(info, implementorGType, &gerr)
	if gerr != nil {
		return nil, errors.New(gerr.Message.String())
	}
	return addr, nil
}

// endregion

// region AttributeIter

type AttributeIter struct {
	data   uptr
	_dummy uptr
}

// endregion

// region BaseInfoStack

type BaseInfoStack struct{}

// endregion

// region Typelib

type Typelib struct{ BaseInfo }

func NewTypelibFromBytes(bytes uptr) (*Typelib, error) {
	var gerr *glib.GError
	typelib := gi_typelib_new_from_bytes.Fn()(bytes, &gerr)
	if gerr != nil {
		return nil, errors.New(gerr.Message.String())
	}
	return typelib, nil
}

func (typelib *Typelib) Ref() *Typelib {
	return gi_typelib_ref.Fn()(typelib)
}

func (typelib *Typelib) Unref() {
	gi_typelib_unref.Fn()(typelib)
}

func (typelib *Typelib) Symbol(symbolName string) (uptr, bool) {
	var symbol uptr
	ok := gi_typelib_symbol.Fn()(typelib, symbolName, &symbol)
	return symbol, ok
}

func (typelib *Typelib) GetNamespace() string {
	return gi_typelib_get_namespace.Fn()(typelib)
}

// endregion

// region Type Getters

func GTypeCallableInfo() gobject.GType {
	return gi_callable_info_get_type.Fn()()
}

func GTypeFunctionInfo() gobject.GType {
	return gi_function_info_get_type.Fn()()
}

func GTypeCallbackInfo() gobject.GType {
	return gi_callback_info_get_type.Fn()()
}

func GTypeRegisteredTypeInfo() gobject.GType {
	return gi_registered_type_info_get_type.Fn()()
}

func GTypeStructInfo() gobject.GType {
	return gi_struct_info_get_type.Fn()()
}

func GTypeUnionInfo() gobject.GType {
	return gi_union_info_get_type.Fn()()
}

func GTypeEnumInfo() gobject.GType {
	return gi_enum_info_get_type.Fn()()
}

func GTypeFlagsInfo() gobject.GType {
	return gi_flags_info_get_type.Fn()()
}

func GTypeObjectInfo() gobject.GType {
	return gi_object_info_get_type.Fn()()
}

func GTypeInterfaceInfo() gobject.GType {
	return gi_interface_info_get_type.Fn()()
}

func GTypeConstantInfo() gobject.GType {
	return gi_constant_info_get_type.Fn()()
}

func GTypeValueInfo() gobject.GType {
	return gi_value_info_get_type.Fn()()
}

func GTypeSignalInfo() gobject.GType {
	return gi_signal_info_get_type.Fn()()
}

func GTypeVFuncInfo() gobject.GType {
	return gi_vfunc_info_get_type.Fn()()
}

func GTypePropertyInfo() gobject.GType {
	return gi_property_info_get_type.Fn()()
}

func GTypeFieldInfo() gobject.GType {
	return gi_field_info_get_type.Fn()()
}

// endregion
